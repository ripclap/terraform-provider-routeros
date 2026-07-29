package routeros

import (
	"testing"
)

// The optional "user-manager" RouterOS package is required; installing it needs a reboot, which an
// acceptance test may not do, so the test is skipped.
func TestAccUserManagerPaymentTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource requires the optional 'user-manager' RouterOS package.")
}
