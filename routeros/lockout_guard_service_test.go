package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// serviceData builds the minimal ResourceData the service guard reads: the
// name, used when the item being written does not carry one, and the
// acknowledgement flag.
func serviceData(t *testing.T, name string, ack bool) (map[string]*schema.Schema, *schema.ResourceData) {
	t.Helper()
	s := map[string]*schema.Schema{
		KeyName:       {Type: schema.TypeString, Optional: true},
		KeyLockoutAck: PropLockoutAck,
	}
	return s, schema.TestResourceDataRaw(t, s, map[string]interface{}{
		KeyName:       name,
		KeyLockoutAck: ack,
	})
}

func TestCheckServiceLockout(t *testing.T) {
	for _, tc := range []struct {
		name    string
		item    MikrotikItem
		dName   string
		refused bool
	}{
		{"disabling ssh", MikrotikItem{"name": "ssh", "disabled": "true"}, "", true},
		{"disabling api-ssl", MikrotikItem{"name": "api-ssl", "disabled": "true"}, "", true},
		{"disabling www-ssl", MikrotikItem{"name": "www-ssl", "disabled": "true"}, "", true},
		{"disabling winbox", MikrotikItem{"name": "winbox", "disabled": "true"}, "", true},

		{"enabling ssh", MikrotikItem{"name": "ssh", "disabled": "false"}, "", false},
		{"editing port only", MikrotikItem{"name": "ssh", "port": "2222"}, "", false},
		{"disabling telnet is fine", MikrotikItem{"name": "telnet", "disabled": "true"}, "", false},
		{"disabling ftp is fine", MikrotikItem{"name": "ftp", "disabled": "true"}, "", false},

		// /ip/service is keyed by name, so an update may omit it from the body.
		{"name taken from state", MikrotikItem{"disabled": "true"}, "ssh", true},
		{"name from state, not management", MikrotikItem{"disabled": "true"}, "ftp", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, d := serviceData(t, tc.dName, false)
			err := checkServiceLockout(tc.item, d)
			if tc.refused != (err != nil) {
				t.Fatalf("refused=%v, err=%v", tc.refused, err)
			}
		})
	}
}

// The acknowledgement must bypass the guard, and only via the dispatcher --
// that is where it is read.
func TestCheckLockoutAcknowledgement(t *testing.T) {
	item := MikrotikItem{"name": "ssh", "disabled": "true"}

	s, d := serviceData(t, "", false)
	if err := CheckLockout("/ip/service", item, s, d); err == nil {
		t.Fatal("expected a refusal without lockout_ack")
	}

	s, d = serviceData(t, "", true)
	if err := CheckLockout("/ip/service", item, s, d); err != nil {
		t.Fatalf("lockout_ack should permit the change, got: %v", err)
	}
}

// A resource with no lockout_ack attribute must not panic, and must not be
// treated as acknowledged.
func TestCheckLockoutOnUnguardedResource(t *testing.T) {
	s := map[string]*schema.Schema{
		KeyName: {Type: schema.TypeString, Optional: true},
	}
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{KeyName: "x"})

	if err := CheckLockout("/ip/pool", MikrotikItem{"name": "x"}, s, d); err != nil {
		t.Fatalf("an unguarded path should pass: %v", err)
	}
	if lockoutAcknowledged(s, d) {
		t.Fatal("a resource without the attribute must not count as acknowledged")
	}
}
