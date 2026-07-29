package routeros

import (
	"testing"
)

// MACsec binds only to a physical Ethernet port (else `input does not match any value of interface`),
// and the test device's only such port, ether1, carries the management transport.
func TestAccInterfaceMacsecTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource requires a spare physical Ethernet port and a MACsec peer, " +
		"so it is only testable on real hardware.")
}
