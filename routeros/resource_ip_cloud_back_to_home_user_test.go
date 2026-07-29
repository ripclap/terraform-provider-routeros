package routeros

import (
	"testing"
)

// Skipped: the MikroTik Cloud subsystem does not exist on RouterOS x86 ("Cloud services not supported
// on x86"); Back To Home needs hardware with a working cloud registration.
func TestAccIpCloudBackToHomeUserTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource is only available on hardware that supports MikroTik Cloud; " +
		"the '/ip/cloud/back-to-home-user' menu does not exist on RouterOS x86.")
}
