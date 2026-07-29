package routeros

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestStateMigrationStringToBool(t *testing.T) {
	up := stateMigrationStringToBool("ddns_enabled")

	for _, tc := range []struct {
		name   string
		in     map[string]interface{}
		want   interface{}
		absent bool
	}{
		{name: "yes", in: map[string]interface{}{"ddns_enabled": "yes"}, want: true},
		{name: "true", in: map[string]interface{}{"ddns_enabled": "true"}, want: true},
		{name: "one", in: map[string]interface{}{"ddns_enabled": "1"}, want: true},
		{name: "mixed case", in: map[string]interface{}{"ddns_enabled": " YeS "}, want: true},
		{name: "no", in: map[string]interface{}{"ddns_enabled": "no"}, want: false},
		{name: "false", in: map[string]interface{}{"ddns_enabled": "false"}, want: false},
		{name: "empty is dropped", in: map[string]interface{}{"ddns_enabled": ""}, absent: true},
		{name: "missing stays missing", in: map[string]interface{}{}, absent: true},
		{name: "already bool untouched", in: map[string]interface{}{"ddns_enabled": true}, want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := up(context.Background(), tc.in, nil)
			if err != nil {
				t.Fatalf("upgrade returned error: %v", err)
			}

			v, ok := got["ddns_enabled"]
			if tc.absent {
				if ok {
					t.Fatalf("expected key to be absent, got %#v", v)
				}
				return
			}
			if !ok {
				t.Fatal("expected key to be present")
			}
			if v != tc.want {
				t.Fatalf("got %#v (%T), want %#v", v, v, tc.want)
			}
		})
	}
}

// The upgrader is only meaningful if v0 really did hold a string.
func TestIpCloudV0DiffersFromCurrent(t *testing.T) {
	v0 := ResourceIpCloudV0().Schema["ddns_enabled"]
	if v0 == nil {
		t.Fatal("v0 schema is missing ddns_enabled")
	}
	if v0.Type != schema.TypeString {
		t.Fatalf("v0 ddns_enabled should be TypeString, got %s", v0.Type)
	}

	cur := ResourceIpCloud()
	if got := cur.Schema["ddns_enabled"].Type; got != schema.TypeBool {
		t.Fatalf("current ddns_enabled should be TypeBool, got %s", got)
	}
	if cur.SchemaVersion != 1 {
		t.Fatalf("SchemaVersion should be 1, got %d", cur.SchemaVersion)
	}
	if len(cur.StateUpgraders) != 1 || cur.StateUpgraders[0].Version != 0 {
		t.Fatalf("expected exactly one v0 upgrader, got %#v", cur.StateUpgraders)
	}
}
