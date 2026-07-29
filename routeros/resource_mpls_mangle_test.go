package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsMangleAddress = "routeros_mpls_mangle.test_mpls_mangle_x"

func TestAccMplsMangleTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/mpls/mangle", "routeros_mpls_mangle"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsMangleConfig("forward", "5"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsMangleAddress),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "chain", "forward"),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "exp", "3"),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "set_exp", "5"),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "set_mark", "test_mpls_mangle_x_mark"),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "comment", "test_mpls_mangle_x"),
						),
					},
					{
						Config: testAccMplsMangleConfig("output", "6"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsMangleAddress),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "chain", "output"),
							resource.TestCheckResourceAttr(testMplsMangleAddress, "set_exp", "6"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsMangleConfig(chain, setExp string) string {
	return providerConfig + `

resource "routeros_mpls_mangle" "test_mpls_mangle_x" {
	chain    = "` + chain + `"
	exp      = 3
	set_exp  = ` + setExp + `
	set_mark = "test_mpls_mangle_x_mark"
	comment  = "test_mpls_mangle_x"
}

`
}
