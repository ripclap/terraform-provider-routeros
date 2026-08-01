package routeros

import (
	"testing"
)

// The `/file/sync` menu is absent on x86 builds (`GET /rest/file/sync` answers `no such command prefix`), so this cannot be exercised.
func TestAccFileSyncTest_basic(t *testing.T) {
	t.Log("Test skipped, the '/file/sync' menu is not present in this RouterOS build.")
}
