package routeros

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Each resource file opens with a response captured from a device. Replaying
// those through the real deserialiser turns them into a corpus: whatever a
// device actually returned once has to survive being read back into state.
//
// Two defects show up this way without any hardware:
//
//   - A value the schema cannot hold. RouterOS answers with word sentinels in
//     fields it documents as numbers - mtu=auto, horizon=none,
//     hop-limit=unspecified - and an attribute typed TypeInt cannot take them,
//     so the value lands nowhere and the plan never settles.
//   - A field the device returns that the schema does not declare, which is
//     silently dropped on read, so drift in it is invisible.

var (
	reReplayPath  = regexp.MustCompile(`PropResourcePath\("([^"]+)"\)`)
	reReplayPair  = regexp.MustCompile(`"([a-z0-9._-]+)"\s*:\s*"([^"]*)"`)
	reReplayBlock = regexp.MustCompile(`(?s)/\*\s*(\{.*?\})\s*\*/`)
)

// capturedResponse is one device response recovered from a source comment.
type capturedResponse struct {
	file string
	item MikrotikItem
}

// withRouterOSVersion pins the package-level version for the duration of a
// test. The deserialiser folds the drift map into its rename table keyed on
// that version, so leaving it empty skips every drift entry and reports
// renamed fields - /container's `repo`, which became `tag` in 7.18 - as
// missing from the schema when they are in fact handled.
func withRouterOSVersion(t *testing.T) {
	t.Helper()
	v := os.Getenv("ROS_VERSION")
	if v == "" {
		v = "7.23"
	}
	prev := RouterOSVersion
	RouterOSVersion = v
	t.Cleanup(func() { RouterOSVersion = prev })
}

var (
	reRegistration = regexp.MustCompile(`"(routeros_[a-z0-9_]+)":\s*(Resource[A-Za-z0-9_]+)\(\)`)
	reConstructor  = regexp.MustCompile(`(?m)^func (Resource[A-Za-z0-9_]+)\(\)`)
)

// deviceSamples maps a registered resource name to the response captured in the
// file that defines it.
//
// Keyed by constructor rather than by menu path on purpose: several menus are
// claimed by more than one file - /interface/ethernet/switch has both a generic
// and a CRS-specific definition - and matching on the path attributes one
// file's sample to the other file's schema, which reports dozens of fields as
// missing that are simply not part of that variant.
func deviceSamples(t *testing.T) map[string]capturedResponse {
	t.Helper()

	files, err := filepath.Glob("resource_*.go")
	if err != nil {
		t.Fatalf("glob: %v", err)
	}

	providerSrc, err := os.ReadFile("provider.go")
	if err != nil {
		t.Fatalf("read provider.go: %v", err)
	}
	ctorForName := map[string]string{}
	for _, m := range reRegistration.FindAllStringSubmatch(string(providerSrc), -1) {
		ctorForName[m[1]] = m[2]
	}

	sampleForCtor := map[string]capturedResponse{}
	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") || reOldVersion.MatchString(path) {
			continue
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		src := string(raw)

		// A file may define several resources, each preceded by its own sample.
		// A sample belongs to the first constructor that follows it and to no
		// other: /tool/mac-server carries one sample but three constructors, and
		// letting the later two inherit it reports that menu's fields as missing
		// from the winbox and ping schemas, which do not have them.
		blocks := reReplayBlock.FindAllStringSubmatchIndex(src, -1)
		if len(blocks) == 0 {
			continue
		}
		ctors := reConstructor.FindAllStringSubmatchIndex(src, -1)
		for _, c := range ctors {
			var chosen []int
			for _, b := range blocks {
				if b[1] > c[0] {
					break
				}
				chosen = b
			}
			if chosen == nil {
				continue
			}
			// Skip if another constructor sits between the sample and this one.
			shadowed := false
			for _, other := range ctors {
				if other[0] > chosen[1] && other[0] < c[0] {
					shadowed = true
					break
				}
			}
			if shadowed {
				continue
			}
			body := src[chosen[2]:chosen[3]]

			item := MikrotikItem{}
			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(body), &parsed); err == nil {
				for k, v := range parsed {
					item[k] = fmt.Sprint(v)
				}
			} else {
				// Some blocks are hand-maintained and not valid JSON; recover
				// the pairs rather than skipping and passing vacuously.
				for _, m := range reReplayPair.FindAllStringSubmatch(body, -1) {
					item[m[1]] = m[2]
				}
			}
			if len(item) == 0 {
				continue
			}
			sampleForCtor[src[c[2]:c[3]]] = capturedResponse{file: path, item: item}
		}
	}

	out := map[string]capturedResponse{}
	for name, ctor := range ctorForName {
		if s, ok := sampleForCtor[ctor]; ok {
			out[name] = s
		}
	}
	return out
}

// TestCapturedResponsesDeserialise replays every captured response through the
// deserialiser and fails on an error diagnostic: the device said this, so the
// provider must be able to read it back.
func TestCapturedResponsesDeserialise(t *testing.T) {
	withRouterOSVersion(t)
	samples := deviceSamples(t)
	if len(samples) == 0 {
		t.Fatal("no captured responses found; the check would pass vacuously")
	}

	p := NewProvider()
	var failures []string
	replayed := 0

	for name, res := range p.ResourcesMap {
		pathAttr, ok := res.Schema[MetaResourcePath]
		if !ok {
			continue
		}
		if _, ok := pathAttr.Default.(string); !ok {
			continue
		}
		sample, ok := samples[name]
		if !ok {
			continue
		}

		replayed++
		d := res.TestResourceData()
		for _, diagnostic := range MikrotikResourceDataToTerraform(sample.item, res.Schema, d) {
			if diagnostic.Severity != 0 { // 0 is Error in diag.Severity
				continue
			}
			failures = append(failures, fmt.Sprintf("%s (%s): %s | %s",
				name, sample.file, diagnostic.Summary, diagnostic.Detail))
		}
	}

	t.Logf("replayed %d captured device responses", replayed)
	if len(failures) == 0 {
		return
	}

	sort.Strings(failures)
	t.Errorf("%d captured device responses fail to deserialise; the device returned "+
		"these values, so the schema must accept them:", len(failures))
	for _, f := range failures {
		t.Errorf("    %s", f)
	}
}

// supersededSampleFields are fields whose captured response predates a RouterOS
// rename, where the schema carries the newer spelling. The sample that names
// them also carries `use-dn` marked "not present in version 7.9", which dates
// it, and each has a plausible successor already declared:
//
//	in-filter  -> in_filter_chain      (out-filter-chain kept the suffix)
//	network    -> networks
//
// Renaming the schema to match the old sample would break the attribute on
// current firmware, and adding a second attribute would give one device field
// two names. Settling it needs a `print detail` from a device on a current
// release; until then the pairing is recorded rather than guessed at.
var supersededSampleFields = map[string]struct{}{
	"routeros_routing_ospf_instance:in_filter":         {},
	"routeros_routing_ospf_interface_template:network": {},
}

// TestCapturedResponsesHaveNoUnknownFields reports fields a device returned
// that the schema does not declare. Those are dropped on read, so drift in them
// cannot be seen.
func TestCapturedResponsesHaveNoUnknownFields(t *testing.T) {
	withRouterOSVersion(t)
	samples := deviceSamples(t)
	p := NewProvider()

	var gaps []string
	for name, res := range p.ResourcesMap {
		pathAttr, ok := res.Schema[MetaResourcePath]
		if !ok {
			continue
		}
		if _, ok := pathAttr.Default.(string); !ok {
			continue
		}
		sample, ok := samples[name]
		if !ok {
			continue
		}

		d := res.TestResourceData()
		for _, diagnostic := range MikrotikResourceDataToTerraform(sample.item, res.Schema, d) {
			if !strings.Contains(diagnostic.Summary, "not found in the schema") {
				continue
			}
			field := strings.TrimSuffix(strings.TrimPrefix(diagnostic.Summary, "Field '"), "' not found in the schema")
			if _, ok := supersededSampleFields[name+":"+field]; ok {
				t.Logf("skipping %s.%s: captured response predates a rename", name, field)
				continue
			}
			gaps = append(gaps, fmt.Sprintf("%s (%s): %s", name, sample.file, diagnostic.Summary))
		}
	}

	if len(gaps) == 0 {
		return
	}
	sort.Strings(gaps)
	t.Errorf("%d fields present in a captured device response are absent from the "+
		"schema, so they are dropped on read and drift in them is invisible:", len(gaps))
	for _, g := range gaps {
		t.Errorf("    %s", g)
	}
}

// unvalidatableSampleValues are captured values that are not device values.
// Both are placeholders in a hand-written sample rather than something a
// device produced: a redacted key, and an ARP opcode of 0, which is not one
// of the ten the protocol defines.
var unvalidatableSampleValues = map[string]struct{}{
	"routeros_interface_wireless_security_profiles.wpa2_pre_shared_key": {},
	"routeros_interface_bridge_filter.arp_opcode":                       {},
}

// TestValidationAcceptsCapturedValues runs each attribute's own validator over
// the value the device returned for it.
//
// Validation does not run on reads - the SDK applies it to configuration, so a
// value that arrives from the device reaches state either way. What it costs
// is expressiveness: a value the device holds and reports cannot be written in
// HCL, so after an import there is no configuration that matches state.
//
// That bites hardest on an attribute carrying DiffSuppressFunc, where omitting
// it leaves the device value in place. If the sentinel that clears the field
// also fails validation, there is no way to clear it at all.
func TestValidationAcceptsCapturedValues(t *testing.T) {
	withRouterOSVersion(t)
	samples := deviceSamples(t)
	p := NewProvider()

	var rejected []string
	checked := 0

	for name, res := range p.ResourcesMap {
		sample, ok := samples[name]
		if !ok {
			continue
		}
		for field, value := range sample.item {
			if strings.HasPrefix(field, ".") || value == "" {
				continue
			}
			attr, ok := res.Schema[KebabToSnake(field)]
			if !ok || attr.Type != schema.TypeString {
				continue
			}
			if _, skip := unvalidatableSampleValues[name+"."+KebabToSnake(field)]; skip {
				continue
			}

			switch {
			case attr.ValidateFunc != nil:
				checked++
				if _, errs := attr.ValidateFunc(value, KebabToSnake(field)); len(errs) > 0 {
					rejected = append(rejected, fmt.Sprintf("%s.%s rejects %q: %v",
						name, KebabToSnake(field), value, errs[0]))
				}
			case attr.ValidateDiagFunc != nil:
				checked++
				for _, d := range attr.ValidateDiagFunc(value, cty.GetAttrPath(KebabToSnake(field))) {
					if d.Severity != 0 {
						continue
					}
					rejected = append(rejected, fmt.Sprintf("%s.%s rejects %q: %s",
						name, KebabToSnake(field), value, d.Summary))
				}
			}
		}
	}

	t.Logf("ran %d validators over captured values", checked)
	if len(rejected) == 0 {
		return
	}
	sort.Strings(rejected)
	t.Errorf("%d attributes reject a value the device returned:", len(rejected))
	for _, r := range rejected {
		t.Errorf("    %s", r)
	}
}
