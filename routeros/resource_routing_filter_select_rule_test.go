package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterSelectRule = "routeros_routing_filter_select_rule.test_routing_filter_select_rule"
const testRoutingFilterSelectRulePrfx = "routeros_routing_filter_select_rule.test_routing_filter_select_rule_prfx"
const testRoutingFilterSelectRuleTake = "routeros_routing_filter_select_rule.test_routing_filter_select_rule_take"
const testRoutingFilterSelectRuleJump = "routeros_routing_filter_select_rule.test_routing_filter_select_rule_jump"

// The do-select-*/do-group-* properties only parse as '<property-token>><ordering-selector>'; RouterOS
// answers "expected >" when the ordering selector is omitted. The rules are chained with depends_on
// because concurrent REST writes to /routing/filter/select-rule wedge the RouterOS 7.23 www-ssl server.
func TestAccRoutingFilterSelectRuleTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/filter/select-rule",
					"routeros_routing_filter_select_rule"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterSelectRuleConfig("bgp-local-pref>largest-none-best"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterSelectRule),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRule, "chain", "test_routing_filter_select_rule_x"),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRule, "comment", "acc test select rule"),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRule, "do_select_num", "bgp-local-pref>largest-none-best"),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRule, "disabled", "false"),

							testResourcePrimaryInstanceId(testRoutingFilterSelectRulePrfx),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRulePrfx, "do_select_prfx", "gw>largest-none-best"),

							testResourcePrimaryInstanceId(testRoutingFilterSelectRuleTake),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRuleTake, "do_take", "1"),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRuleTake, "disabled", "true"),

							testResourcePrimaryInstanceId(testRoutingFilterSelectRuleJump),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRuleJump, "do_jump", "test_routing_filter_select_rule_y"),
						),
					},
					{
						Config: testAccRoutingFilterSelectRuleConfig("bgp-weight>smallest-none-worst"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterSelectRule),
							resource.TestCheckResourceAttr(testRoutingFilterSelectRule, "do_select_num", "bgp-weight>smallest-none-worst"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterSelectRuleConfig(doSelectNum string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_filter_select_rule" "test_routing_filter_select_rule" {
  chain         = "test_routing_filter_select_rule_x"
  comment       = "acc test select rule"
  do_select_num = "%v"
}

resource "routeros_routing_filter_select_rule" "test_routing_filter_select_rule_prfx" {
  chain          = "test_routing_filter_select_rule_x"
  do_select_prfx = "gw>largest-none-best"

  depends_on = [routeros_routing_filter_select_rule.test_routing_filter_select_rule]
}

resource "routeros_routing_filter_select_rule" "test_routing_filter_select_rule_take" {
  chain    = "test_routing_filter_select_rule_x"
  do_take  = "1"
  disabled = true

  depends_on = [routeros_routing_filter_select_rule.test_routing_filter_select_rule_prfx]
}

resource "routeros_routing_filter_select_rule" "test_routing_filter_select_rule_jump" {
  chain   = "test_routing_filter_select_rule_x"
  do_jump = "test_routing_filter_select_rule_y"

  depends_on = [routeros_routing_filter_select_rule.test_routing_filter_select_rule_take]
}
`, providerConfig, doSelectNum)
}
