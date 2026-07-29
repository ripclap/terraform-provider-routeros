package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchQosMapTest_basic is skipped: /interface/ethernet/switch exists only on
// switch-chip boards (QoS tree needs 98DX8xxx/98DX4xxx, CRS5xx+), not the x86 acceptance-test VM.
func TestAccInterfaceEthernetSwitchQosMapTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/qos/map does not exist on x86 RouterOS).")
}
