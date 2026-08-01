package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchL3HwSettingsTest_basic is skipped: /interface/ethernet/switch exists
// only on switch-chip boards (L3 offload needs CRS3xx/CRS5xx), not on x86 builds.
func TestAccInterfaceEthernetSwitchL3HwSettingsTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/l3hw-settings does not exist on x86 RouterOS).")
}
