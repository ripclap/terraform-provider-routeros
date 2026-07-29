package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchQosTxManagerTest_basic is skipped: the x86 test VM has no switch
// chip, so /interface/ethernet/switch does not exist ("bad command name switch").
func TestAccInterfaceEthernetSwitchQosTxManagerTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/qos/tx-manager does not exist on x86 RouterOS).")
}
