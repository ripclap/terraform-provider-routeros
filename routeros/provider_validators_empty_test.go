package routeros

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

// A device reports an unset optional attribute as an empty string. That is not a value
// to validate, and rejecting it made every affected resource unimportable.
func TestSliceValidatorsAcceptEmpty(t *testing.T) {
	valid := []string{"ip", "ipv6"}

	multi := ValidationMultiValInSlice(valid, false, false)
	for _, v := range []string{"", "ip", "ipv6", "ip,ipv6", "ip, ipv6"} {
		if d := multi(v, cty.Path{}); d.HasError() {
			t.Errorf("MultiValInSlice rejected %q: %v", v, d)
		}
	}
	for _, v := range []string{"nope", "ip,nope"} {
		if d := multi(v, cty.Path{}); !d.HasError() {
			t.Errorf("MultiValInSlice accepted %q", v)
		}
	}

	single := ValidationValInSlice(valid, false, false)
	for _, v := range []string{"", "ip", "ipv6"} {
		if d := single(v, cty.Path{}); d.HasError() {
			t.Errorf("ValInSlice rejected %q: %v", v, d)
		}
	}
	if d := single("nope", cty.Path{}); !d.HasError() {
		t.Error("ValInSlice accepted \"nope\"")
	}
}
