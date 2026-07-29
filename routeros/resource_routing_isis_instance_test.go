package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingIsisInstance = "routeros_routing_isis_instance.test_routing_isis_instance_x"

func TestAccRoutingIsisInstanceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/isis/instance", "routeros_routing_isis_instance"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingIsisInstanceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIsisInstance),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "name", "test_routing_isis_instance_x"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "afi", "ip"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "system_id", "90ab.cdef.1010"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "metric_type", "wide"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "areas.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingIsisInstance, "areas.*", "49.1010"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "l1_originate_default", "never"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "l2_originate_default", "always"),
						),
					},
					{
						Config: testAccRoutingIsisInstanceUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIsisInstance),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "metric_type", "both"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "comment", "test_routing_isis_instance_x comment"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "disabled", "true"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "areas.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingIsisInstance, "areas.*", "49.1011"),
							resource.TestCheckResourceAttr(testRoutingIsisInstance, "l1_redistribute.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingIsisInstance, "l1_redistribute.*", "connected"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIsisInstanceConfig() string {
	return providerConfig + `

resource "routeros_routing_isis_instance" "test_routing_isis_instance_x" {
	name                    = "test_routing_isis_instance_x"
	afi                     = "ip"
	system_id               = "90ab.cdef.1010"
	areas                   = ["49.1010"]
	metric_type             = "wide"
	l1_originate_default    = "never"
	l1_lsp_refresh_interval = "900"
	l2_originate_default    = "always"
	l2_lsp_max_age          = "1200"
}

`
}

func testAccRoutingIsisInstanceUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_isis_instance" "test_routing_isis_instance_x" {
	name                    = "test_routing_isis_instance_x"
	afi                     = "ip"
	system_id               = "90ab.cdef.1010"
	areas                   = ["49.1010", "49.1011"]
	metric_type             = "both"
	comment                 = "test_routing_isis_instance_x comment"
	disabled                = true
	l1_originate_default    = "never"
	l1_lsp_refresh_interval = "900"
	l1_redistribute         = ["connected", "static"]
	l2_originate_default    = "always"
	l2_lsp_max_age          = "1200"
}

`
}
