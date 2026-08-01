package routeros

import (
	"testing"
)

// The `/system/ptp/port` menu exists only on PTP capable hardware; on x86 builds
// `/system/ptp/port/print` -> `syntax error`.
func TestAccSystemPtpPortTest_basic(t *testing.T) {
	t.Log("Test skipped, the resource is only available on real hardware with PTP support.")
}
