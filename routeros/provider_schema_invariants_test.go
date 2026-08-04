package routeros

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Static checks over every registered schema. They need no device, so they run
// in the ordinary unit suite and fail the build rather than being discovered by
// a user writing a configuration.

// Terraform meta-arguments. An attribute with one of these names is rejected by
// Terraform core at schema-load time, which fails every plan against the
// provider rather than only the offending resource.
var reservedAttributeNames = map[string]struct{}{
	"count": {}, "for_each": {}, "depends_on": {}, "lifecycle": {},
	"provider": {}, "provisioner": {}, "connection": {},
}

// Attributes whose name denotes a credential. Anything matching must be
// Sensitive, or Terraform prints it in plan output and writes it to state
// unredacted.
var secretNameRe = regexp.MustCompile(
	`(?i)passphrase|password|(^|_)secret($|_)|private_key|preshared|(^|_)psk($|_)|otp_secret|(^|_)key$`)

// Names that match the pattern but hold no secret. Each is a selector, a name
// reference, a toggle or a number, and marking one Sensitive would redact
// something a reader needs to see in a plan.
var notReallySecret = map[string]struct{}{
	"password_format": {}, "minimum_password_length": {}, "password_authentication": {},
	"multi_passphrase_group": {}, "passphrase_group": {}, "require_message_auth": {},
	"key_type": {}, "key_size": {}, "key_usage": {}, "public_key": {},
	"sort_key": {}, "authentication_key_id": {}, "key_id": {},

	// container env: the variable's name, not its value.
	"key": {},
	// ipsec identity: names of entries in the keys menu, not the material.
	"remote_key": {},
	// which of the static keys to transmit with (key-0 .. key-3).
	"static_transmit_key": {},
	// the identity a pre-shared key is filed under, not the key.
	"psk_id": {},
}

// carriesSecret reports whether an attribute could hold credential material at
// all. Key material is a string; a bool or a number is a flag or a selector,
// however its name reads. This is what separates /certificate's private_key,
// a bool saying whether a key is present, from WireGuard's private_key.
func carriesSecret(a *schema.Schema) bool {
	switch a.Type {
	case schema.TypeBool, schema.TypeInt, schema.TypeFloat:
		return false
	}
	return true
}

// walk visits every attribute of a schema, descending into nested blocks, and
// reports the dotted path so a nested offender is identifiable.
func walkSchema(prefix string, s map[string]*schema.Schema, fn func(path string, name string, a *schema.Schema)) {
	for name, attr := range s {
		path := name
		if prefix != "" {
			path = prefix + "." + name
		}
		fn(path, name, attr)
		if res, ok := attr.Elem.(*schema.Resource); ok {
			walkSchema(path, res.Schema, fn)
		}
	}
}

func TestNoReservedAttributeNames(t *testing.T) {
	p := NewProvider()
	// Only the top level of a resource block is affected; the meta-arguments
	// have no meaning inside a nested block, where the names are ordinary.
	check := func(kind string, m map[string]*schema.Resource) {
		for resName, res := range m {
			for name := range res.Schema {
				if _, bad := reservedAttributeNames[name]; bad {
					t.Errorf("%s %s declares reserved attribute %q; Terraform rejects the whole provider",
						kind, resName, name)
				}
			}
		}
	}
	check("resource", p.ResourcesMap)
	check("data source", p.DataSourcesMap)
}

func TestSecretsAreSensitive(t *testing.T) {
	p := NewProvider()
	var missing []string

	for resName, res := range p.ResourcesMap {
		walkSchema("", res.Schema, func(path, name string, a *schema.Schema) {
			if isMetaAttribute(name) {
				return
			}
			if !secretNameRe.MatchString(name) || !carriesSecret(a) {
				return
			}
			if _, ok := notReallySecret[name]; ok {
				return
			}
			// A read-only attribute the device reports is still printed, so
			// Computed alone does not excuse it.
			if !a.Sensitive {
				missing = append(missing, resName+"."+path)
			}
		})
	}

	if len(missing) > 0 {
		t.Errorf("%d secret-bearing attributes are not marked Sensitive, so Terraform "+
			"prints them in plan output and writes them to state unredacted:", len(missing))
		for _, m := range missing {
			t.Errorf("    %s", m)
		}
	}
}

// A data source mirrors its resource, so a secret must stay redacted there too.
func TestDatasourceSecretsAreSensitive(t *testing.T) {
	p := NewProvider()
	var missing []string

	for dsName, ds := range p.DataSourcesMap {
		walkSchema("", ds.Schema, func(path, name string, a *schema.Schema) {
			if isMetaAttribute(name) || !secretNameRe.MatchString(name) || !carriesSecret(a) {
				return
			}
			if _, ok := notReallySecret[name]; ok {
				return
			}
			if !a.Sensitive {
				missing = append(missing, dsName+"."+path)
			}
		})
	}

	if len(missing) > 0 {
		t.Errorf("%d secret-bearing data source attributes are not Sensitive:", len(missing))
		for _, m := range missing[:min(len(missing), 20)] {
			t.Errorf("    %s", m)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestTransformSetUsesKebabOnTheMikrotikSide enforces the contract stated on
// loadTransformSet: the right-hand side of a transform pair is the device's own
// name, in kebab notation.
//
// The two directions disagree when it is written in snake. Serialization runs
// the mapped name through SnakeToKebab, so a write still reaches the right
// property; deserialization matches the raw incoming kebab name against the
// same map and misses. The attribute then accepts a value, sends it, and never
// reads it back, which shows up as a diff that will not settle.
func TestTransformSetUsesKebabOnTheMikrotikSide(t *testing.T) {
	p := NewProvider()

	check := func(kind string, m map[string]*schema.Resource) {
		for name, res := range m {
			ts, ok := res.Schema[MetaTransformSet]
			if !ok {
				continue
			}
			def, ok := ts.Default.(string)
			if !ok {
				continue
			}
			for tf, mt := range loadTransformSet(def, false) {
				if strings.Contains(mt, "_") {
					t.Errorf("%s %s: transform %q -> %q names the device side in snake; "+
						"write converts it to kebab but read does not, so the value is "+
						"dropped on read. Use %q.",
						kind, name, tf, mt, strings.ReplaceAll(mt, "_", "-"))
				}
			}
		}
	}

	check("resource", p.ResourcesMap)
	check("data source", p.DataSourcesMap)
}
