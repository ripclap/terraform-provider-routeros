package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceDerivedEthernet = "data.routeros_interface_ethernet.test"
const testDatasourceDerivedList = "data.routeros_interface_list.test"

// Derived data sources share one implementation, so exercising a couple of
// them against a device covers the read path for all of them: the menu is read,
// the entries are mapped onto the mirrored schema, and `filter` is passed
// through. Both menus exist on every RouterOS device.
func TestAccDatasourceDerived_basic(t *testing.T) {
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
						Config: `data "routeros_interface_ethernet" "test" {}`,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet(testDatasourceDerivedEthernet, "entries.#"),
							resource.TestCheckResourceAttrSet(testDatasourceDerivedEthernet, "entries.0.id"),
							resource.TestCheckResourceAttrSet(testDatasourceDerivedEthernet, "entries.0.name"),
						),
					},
					{
						Config: `data "routeros_interface_list" "test" {
							filter = {
								name = "all"
							}
						}`,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testDatasourceDerivedList, "entries.#", "1"),
							resource.TestCheckResourceAttr(testDatasourceDerivedList, "entries.0.name", "all"),
						),
					},
				},
			})
		})
	}
}
