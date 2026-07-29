package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRpki = "routeros_routing_rpki.test_routing_rpki"

func TestAccRoutingRpkiTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/rpki", "routeros_routing_rpki"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRpkiConfig("3323", "600"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRpki),
							resource.TestCheckResourceAttr(testRoutingRpki, "address", "192.0.2.5"),
							resource.TestCheckResourceAttr(testRoutingRpki, "group", "test_routing_rpki_grp"),
							resource.TestCheckResourceAttr(testRoutingRpki, "port", "3323"),
							resource.TestCheckResourceAttr(testRoutingRpki, "retry_interval", "600"),
							resource.TestCheckResourceAttr(testRoutingRpki, "refresh_interval", "3600"),
							resource.TestCheckResourceAttr(testRoutingRpki, "expire_interval", "7200"),
							resource.TestCheckResourceAttr(testRoutingRpki, "preference", "1"),
							resource.TestCheckResourceAttr(testRoutingRpki, "comment", "test_routing_rpki"),
						),
					},
					{
						Config: testAccRoutingRpkiConfig("3324", "900"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRpki),
							resource.TestCheckResourceAttr(testRoutingRpki, "port", "3324"),
							resource.TestCheckResourceAttr(testRoutingRpki, "retry_interval", "900"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRpkiConfig(port, retry string) string {
	return providerConfig + `

resource "routeros_routing_rpki" "test_routing_rpki" {
	address          = "192.0.2.5"
	group            = "test_routing_rpki_grp"
	comment          = "test_routing_rpki"
	port             = ` + port + `
	preference       = 1
	refresh_interval = "3600"
	retry_interval   = "` + retry + `"
	expire_interval  = "7200"
	disabled         = true
}
`
}
