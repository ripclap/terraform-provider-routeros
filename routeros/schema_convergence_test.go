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
)

// A settings menu returns its whole object on every read, with every field
// populated. An attribute of one declared Optional without either Computed or a
// DiffSuppressFunc therefore never reaches an empty plan: the device reports a
// value, the configuration does not set one, and Terraform proposes removing it
// on every run.
//
// Restricted to settings menus on purpose. A list menu returns only what each
// entry has set, so the same shape there is not a defect - checking those as
// well flags several hundred attributes that converge perfectly well.
//
// The evidence is the comment each resource file opens with, holding a response
// captured from a device, which records exactly which attributes come back. So
// the defect is caught here rather than by an acceptance test that needs
// hardware, or by a user whose plan never converges.
//
// Either remedy is accepted:
//   - DiffSuppressFunc: AlwaysPresentNotUserProvided - ignore the device value
//     when the user did not provide one.
//   - Computed: true - adopt the device value as the planned value.

var (
	reSampleBlock = regexp.MustCompile(`(?s)/\*\s*(\{.*?\})\s*\*/`)
	reSampleKey   = regexp.MustCompile(`"([a-z0-9._-]+)"\s*:`)
	reAttrBlock   = regexp.MustCompile(`(?sm)^\t\t"([a-z0-9_]+)": \{\n(.*?)^\t\t\},`)
	reOldVersion  = regexp.MustCompile(`_v\d+\.go$`)
)

// convergenceExceptions holds attributes where the shape is unavoidable, keyed
// as "file.go:attribute". Add an entry only with the reason the device value
// cannot differ from the configured one.
var convergenceExceptions = map[string]string{}

// sampleAttributes returns the attributes in a file's captured device response,
// in schema form. The block is hand-maintained and some carry trailing
// annotations or a stray quote, so a parse failure falls back to scanning keys
// rather than skipping the file and passing it vacuously.
func sampleAttributes(src string) map[string]bool {
	m := reSampleBlock.FindStringSubmatch(src)
	if m == nil {
		return nil
	}

	out := map[string]bool{}
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(m[1]), &parsed); err == nil {
		for k := range parsed {
			out[KebabToSnake(k)] = true
		}
		return out
	}
	for _, k := range reSampleKey.FindAllStringSubmatch(m[1], -1) {
		out[KebabToSnake(k[1])] = true
	}
	return out
}

func TestSchemaConverges(t *testing.T) {
	files, err := filepath.Glob("resource_*.go")
	if err != nil {
		t.Fatalf("glob: %v", err)
	}

	type finding struct{ file, attr string }
	var findings []finding
	sampled := 0

	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") || reOldVersion.MatchString(path) {
			continue
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		src := string(raw)

		// Settings menus only; identified by the CRUD helpers they use.
		if !strings.Contains(src, "DefaultSystemCreate") &&
			!strings.Contains(src, "DefaultSystemRead") &&
			!strings.Contains(src, "SystemResourceCreateUpdate") {
			continue
		}

		reported := sampleAttributes(src)
		if len(reported) == 0 {
			continue
		}
		sampled++

		for _, m := range reAttrBlock.FindAllStringSubmatch(src, -1) {
			attr, body := m[1], m[2]
			if !reported[attr] {
				continue
			}
			if !strings.Contains(body, "Optional:") ||
				strings.Contains(body, "Computed:") ||
				strings.Contains(body, "Required:") ||
				strings.Contains(body, "DiffSuppressFunc") {
				continue
			}
			if _, ok := convergenceExceptions[path+":"+attr]; ok {
				continue
			}
			findings = append(findings, finding{path, attr})
		}
	}

	if sampled == 0 {
		t.Fatal("no device samples found; the check would pass vacuously")
	}
	t.Logf("checked %d resources carrying a captured device response", sampled)

	if len(findings) == 0 {
		return
	}

	sort.Slice(findings, func(i, j int) bool {
		if findings[i].file != findings[j].file {
			return findings[i].file < findings[j].file
		}
		return findings[i].attr < findings[j].attr
	})

	var b strings.Builder
	fmt.Fprintf(&b, "%d attributes the device reports are Optional with neither Computed nor "+
		"DiffSuppressFunc, so a configuration that omits them never reaches an empty plan:\n",
		len(findings))
	for _, f := range findings {
		fmt.Fprintf(&b, "    %s: %s\n", f.file, f.attr)
	}
	b.WriteString("\nAdd DiffSuppressFunc: AlwaysPresentNotUserProvided, or Computed: true, " +
		"or record an exception in convergenceExceptions with a reason.")
	t.Error(b.String())
}
