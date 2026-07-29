package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchL3HwSettingsAdvancedTest_basic is skipped: /interface/ethernet/switch
// exists only on switch-chip boards (advanced L3 offload needs CRS3xx/CRS5xx), not the x86 test VM.
func TestAccInterfaceEthernetSwitchL3HwSettingsAdvancedTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/l3hw-settings/advanced does not exist on x86 RouterOS).")
}
