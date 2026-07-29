package routeros

import (
	"testing"
)

// The /system/health/settings menu only exists on devices with real health hardware (sensors and a
// controllable fan). On the x86 test VM it is absent (print -> syntax error, GET -> HTTP 500), so no property
// is reachable here.
func TestAccSystemHealthSettingsTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource is only available on real hardware with health sensors and a fan.")
}
