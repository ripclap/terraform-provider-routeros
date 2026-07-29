package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingIsisInterfaceTemplateBcast = "routeros_routing_isis_interface_template.test_isis_iftmpl_x_bcast"
const testRoutingIsisInterfaceTemplatePtp = "routeros_routing_isis_interface_template.test_isis_iftmpl_x_ptp"

func TestAccRoutingIsisInterfaceTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/isis/interface-template", "routeros_routing_isis_interface_template"),
					testCheckResourceDestroy("/routing/isis/instance", "routeros_routing_isis_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingIsisInterfaceTemplateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIsisInterfaceTemplateBcast),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "instance", "test_isis_iftmpl_x"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "interfaces.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingIsisInterfaceTemplateBcast, "interfaces.*", "ether5"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "levels.#", "2"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "bcast_l1_hello_interval", "9"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "bcast_l1_metric", "25"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "bcast_l1_priority", "70"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "ptp", "false"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "passive", "false"),

							testResourcePrimaryInstanceId(testRoutingIsisInterfaceTemplatePtp),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp", "true"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp_hello_3way", "true"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp_hello_interval", "10"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp_l2_metric", "44"),
						),
					},
					{
						Config: testAccRoutingIsisInterfaceTemplateUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIsisInterfaceTemplateBcast),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "bcast_l1_metric", "26"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "comment", "test_isis_iftmpl_x comment"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplateBcast, "disabled", "true"),

							testResourcePrimaryInstanceId(testRoutingIsisInterfaceTemplatePtp),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp_hello_3way", "false"),
							resource.TestCheckResourceAttr(testRoutingIsisInterfaceTemplatePtp, "ptp_l2_metric", "45"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIsisInterfaceTemplateConfig() string {
	return providerConfig + `

resource "routeros_routing_isis_instance" "test_isis_iftmpl_x" {
	name      = "test_isis_iftmpl_x"
	system_id = "90ab.cdef.1011"
	areas     = ["49.1012"]
}

resource "routeros_routing_isis_interface_template" "test_isis_iftmpl_x_bcast" {
	instance                   = routeros_routing_isis_instance.test_isis_iftmpl_x.name
	interfaces                 = ["ether5"]
	levels                     = ["l1", "l2"]
	bcast_l1_hello_interval    = "9"
	bcast_l1_hello_interval_dr = "3"
	bcast_l1_hello_multiplier  = 4
	bcast_l1_csnp_interval     = "11"
	bcast_l1_psnp_interval     = "2"
	bcast_l1_metric            = 25
	bcast_l1_priority          = 70
}

resource "routeros_routing_isis_interface_template" "test_isis_iftmpl_x_ptp" {
	instance             = routeros_routing_isis_instance.test_isis_iftmpl_x.name
	interfaces           = ["ether6"]
	levels               = ["l2"]
	ptp                  = true
	ptp_hello_3way       = true
	ptp_hello_interval   = "10"
	ptp_hello_multiplier = 5
	ptp_l2_csnp_interval = "12"
	ptp_l2_psnp_interval = "3"
	ptp_l2_metric        = 44
}

`
}

func testAccRoutingIsisInterfaceTemplateUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_isis_instance" "test_isis_iftmpl_x" {
	name      = "test_isis_iftmpl_x"
	system_id = "90ab.cdef.1011"
	areas     = ["49.1012"]
}

resource "routeros_routing_isis_interface_template" "test_isis_iftmpl_x_bcast" {
	instance                   = routeros_routing_isis_instance.test_isis_iftmpl_x.name
	interfaces                 = ["ether5"]
	levels                     = ["l1", "l2"]
	comment                    = "test_isis_iftmpl_x comment"
	disabled                   = true
	bcast_l1_hello_interval    = "9"
	bcast_l1_hello_interval_dr = "3"
	bcast_l1_hello_multiplier  = 4
	bcast_l1_csnp_interval     = "11"
	bcast_l1_psnp_interval     = "2"
	bcast_l1_metric            = 26
	bcast_l1_priority          = 70
}

resource "routeros_routing_isis_interface_template" "test_isis_iftmpl_x_ptp" {
	instance             = routeros_routing_isis_instance.test_isis_iftmpl_x.name
	interfaces           = ["ether6"]
	levels               = ["l2"]
	ptp                  = true
	ptp_hello_3way       = false
	ptp_hello_interval   = "10"
	ptp_hello_multiplier = 5
	ptp_l2_csnp_interval = "12"
	ptp_l2_psnp_interval = "3"
	ptp_l2_metric        = 45
}

`
}
