package routeros

import (
	"testing"
)

// TestAccInterfaceEthernetSwitchQosPriorityFlowControlTest_basic is skipped: /interface/ethernet/switch
// exists only on switch-chip boards (PFC is a per-port 802.1Qbb feature), not the x86 test VM.
func TestAccInterfaceEthernetSwitchQosPriorityFlowControlTest_basic(t *testing.T) {
	t.Skip("Test skipped, the resource is only available on real hardware with a switch chip " +
		"(/interface/ethernet/switch/qos/priority-flow-control does not exist on x86 RouterOS).")
}
