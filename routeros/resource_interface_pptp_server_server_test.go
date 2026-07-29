package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePptpServerServer = "routeros_interface_pptp_server_server.test_pptp_server_server_x"

func TestAccInterfacePptpServerServerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePptpServerServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpServerServer),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "enabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "default_profile", "default"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "keepalive_timeout", "60"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "max_mru", "1440"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "max_mtu", "1440"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "mrru", "1600"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "authentication.#", "2"),
							resource.TestCheckTypeSetElemAttr(testInterfacePptpServerServer, "authentication.*", "chap"),
							resource.TestCheckTypeSetElemAttr(testInterfacePptpServerServer, "authentication.*", "mschap2"),
						),
					},
					{
						Config: testAccInterfacePptpServerServerUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpServerServer),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "enabled", "false"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "default_profile", "default-encryption"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "keepalive_timeout", "30"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "max_mru", "1450"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "max_mtu", "1450"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "mrru", "disabled"),
							resource.TestCheckResourceAttr(testInterfacePptpServerServer, "authentication.#", "2"),
							resource.TestCheckTypeSetElemAttr(testInterfacePptpServerServer, "authentication.*", "mschap1"),
							resource.TestCheckTypeSetElemAttr(testInterfacePptpServerServer, "authentication.*", "mschap2"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePptpServerServerConfig() string {
	return providerConfig + `

resource "routeros_interface_pptp_server_server" "test_pptp_server_server_x" {
	enabled           = true
	authentication    = ["chap", "mschap2"]
	default_profile   = "default"
	keepalive_timeout = "60"
	max_mru           = 1440
	max_mtu           = 1440
	mrru              = "1600"
}
`
}

func testAccInterfacePptpServerServerUpdatedConfig() string {
	return providerConfig + `

resource "routeros_interface_pptp_server_server" "test_pptp_server_server_x" {
	enabled           = false
	authentication    = ["mschap1", "mschap2"]
	default_profile   = "default-encryption"
	keepalive_timeout = "30"
	max_mru           = 1450
	max_mtu           = 1450
	mrru              = "disabled"
}
`
}
