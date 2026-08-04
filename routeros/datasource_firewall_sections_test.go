package routeros

import (
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// keysOf returns the block names a firewall data source declares, ignoring
// the meta attributes that carry serializer directives rather than data.
func keysOf(s map[string]*schema.Schema) []string {
	var out []string
	for k := range s {
		if len(k) > 3 && k[:3] == "___" {
			continue
		}
		out = append(out, k)
	}
	return out
}

// The firewall data sources declare a block per firewall menu and read the
// ones named in a package-level slice. A block that is declared but not listed
// is inert: the emptiness check that guards the read does not count it, so
// asking for that section alone reports that no section was given, and asking
// for it alongside another returns nothing for it without saying why.
func TestFirewallDatasourceReadsEverySectionItDeclares(t *testing.T) {
	for _, tc := range []struct {
		name     string
		declared []string
		sections []string
	}{
		{name: "routeros_ip_firewall", declared: keysOf(DatasourceIPFirewall().Schema), sections: ipFirewallSections},
		{name: "routeros_ipv6_firewall", declared: keysOf(DatasourceIPv6Firewall().Schema), sections: ipv6firewallSections},
	} {
		got, want := append([]string(nil), tc.sections...), append([]string(nil), tc.declared...)
		sort.Strings(got)
		sort.Strings(want)
		if len(got) != len(want) {
			t.Errorf("%s declares %v but reads %v", tc.name, want, got)
			continue
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("%s declares %v but reads %v", tc.name, want, got)
				break
			}
		}
	}
}
