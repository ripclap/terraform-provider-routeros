package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRipKeys = "routeros_routing_rip_keys.test_rip_keys_x"

func TestAccRoutingRipKeysTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/rip/keys", "routeros_routing_rip_keys"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRipKeysConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipKeys),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "chain", "test_rip_keys_x_chain"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "key", "ripkeyx_secret1"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "key_id", "7"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "valid_from", "2026-01-01 00:00:00"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "valid_till", "2027-01-01 00:00:00"),
						),
					},
					{
						Config: testAccRoutingRipKeysUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipKeys),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "key", "ripkeyx_secret2"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "key_id", "8"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "valid_till", "2028-01-01 00:00:00"),
							resource.TestCheckResourceAttr(testRoutingRipKeys, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRipKeysConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_keys" "test_rip_keys_x" {
	chain      = "test_rip_keys_x_chain"
	key        = "ripkeyx_secret1"
	key_id     = 7
	valid_from = "2026-01-01 00:00:00"
	valid_till = "2027-01-01 00:00:00"
}

`
}

func testAccRoutingRipKeysUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_keys" "test_rip_keys_x" {
	chain      = "test_rip_keys_x_chain"
	key        = "ripkeyx_secret2"
	key_id     = 8
	disabled   = true
	valid_from = "2026-01-01 00:00:00"
	valid_till = "2028-01-01 00:00:00"
}

`
}
