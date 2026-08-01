package routeros

import (
	"testing"
)

// `/app/add` fails with `container device-mode needs to be enabled`; enabling that device-mode requires a reboot a device may not be able to take.
func TestAccAppTest_basic(t *testing.T) {
	t.Log("Test skipped, '/app/add' requires the 'container' device-mode, which can only be enabled with a reboot.")
}
