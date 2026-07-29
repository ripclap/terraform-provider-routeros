package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchQosSettingsTest_basic is skipped: /interface/ethernet/switch exists only
// on switch-chip boards (QoS settings singleton needs the switch-chip driver), not the x86 test VM.
func TestAccInterfaceEthernetSwitchQosSettingsTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/qos/settings does not exist on x86 RouterOS).")
}
