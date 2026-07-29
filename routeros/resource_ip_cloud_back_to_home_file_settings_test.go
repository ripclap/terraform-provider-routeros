package routeros

import (
	"testing"
)

// Skipped: the MikroTik Cloud subsystem does not exist on RouterOS x86 ("Cloud services not supported
// on x86"); the resource needs hardware that supports MikroTik Cloud.
func TestAccIpCloudBackToHomeFileSettingsTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource is only available on hardware that supports MikroTik Cloud; " +
		"the '/ip/cloud/back-to-home-file/settings' menu does not exist on RouterOS x86.")
}
