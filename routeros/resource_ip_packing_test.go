package routeros

import (
	"testing"
)

// `/ip/packing` (M3P) is only accepted on real MikroTik ethernet/wireless hardware; on a VM RouterOS
// refuses every interface with `packing for this interface is not possible`, so it can't be tested here.
func TestAccIpPackingTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource is only available on real hardware: RouterOS answers " +
		"'packing for this interface is not possible' for every interface type of a virtual machine.")
}
