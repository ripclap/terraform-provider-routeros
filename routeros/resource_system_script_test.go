package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemScriptTask = "routeros_system_script.script"

// testSystemScriptSource is the exact source RouterOS stores and returns. RouterOS rewrites leading
// tabs in a script body to spaces, so the heredoc below must not be indented or it reports a permanent diff.
const testSystemScriptSource = ":log info \"This is a test script created by Terraform.\"\n"

func TestAccSystemScriptTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/script", "routeros_script"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemScriptConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemScriptTask),
							resource.TestCheckResourceAttr(testSystemScriptTask, "name", "my_script"),
							resource.TestCheckResourceAttr(testSystemScriptTask, "source", testSystemScriptSource),
							resource.TestCheckTypeSetElemAttr(testSystemScriptTask, "policy.*", "read"),
							resource.TestCheckTypeSetElemAttr(testSystemScriptTask, "policy.*", "write"),
							resource.TestCheckTypeSetElemAttr(testSystemScriptTask, "policy.*", "test"),
							resource.TestCheckTypeSetElemAttr(testSystemScriptTask, "policy.*", "policy"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemScriptConfig() string {
	return providerConfig + `
resource "routeros_system_script" "script" {
	name   = "my_script"
	source = <<EOF
:log info "This is a test script created by Terraform."
EOF
	policy = ["read", "write", "test", "policy"]
}
`
}
