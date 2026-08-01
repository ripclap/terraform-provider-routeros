package routeros

import (
	"testing"
)

// Skipped: the resource needs an '/iot/mqtt/brokers' entry, and deleting a broker hangs the whole
// '/iot/mqtt' subtree of RouterOS 7.23 until reboot (device-side defect; see resource_iot_mqtt_brokers_test.go).
func TestAccIotMqttSubscriptionsTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource needs an '/iot/mqtt/brokers' entry and deleting one hangs the " +
		"'/iot/mqtt' subtree of RouterOS 7.23 until the device is rebooted.")
}
