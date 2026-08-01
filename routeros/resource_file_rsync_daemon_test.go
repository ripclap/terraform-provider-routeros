package routeros

import (
	"testing"
)

// The `/file/rsync-daemon` menu is absent on x86 builds; the daemon ships only where file sync exists, so this cannot be exercised.
func TestAccFileRsyncDaemonTest_basic(t *testing.T) {
	t.Log("Test skipped, the '/file/rsync-daemon' menu is not present in this RouterOS build.")
}
