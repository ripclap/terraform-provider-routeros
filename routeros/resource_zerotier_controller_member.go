package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on RouterOS 7.23 (RouterOS 7.23: the default
  instance "zt1" exists but is disabled and /zerotier/controller has no networks), so no live
  JSON sample can be shown. The field set below was taken from the device itself:

    /console/inspect request=syntax path="zerotier,controller,member,add"
      -> authorized, bridge, comment, disabled, ip-address, name, network, zt-address
    /zerotier/controller/member/print detail
      -> Flags: X - DISABLED, I - INACTIVE; A - AUTHORIZED; B - BRIDGE
    /zerotier/controller/member/print proplist=<name>
      -> "inactive" and "last-seen" exist; "instance", "identity", "version", "ip6-address"
         and "dynamic" do NOT exist in this menu.
    /zerotier/controller/member/print where authorized=<v> / bridge=<v>
      -> the console answers "expected yes or no", so both are booleans.
    /console/inspect request=completion path="zerotier,controller,member"
                     input="add ip-address=192.0.2.1"
      -> offers "," so ip-address is a comma separated list.
*/

// ResourceZerotierControllerMember manages the members of a ZeroTier network hosted by the local controller.
// https://help.mikrotik.com/docs/spaces/ROS/pages/83755083/ZeroTier
func ResourceZerotierControllerMember() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/zerotier/controller/member"),
		MetaId:           PropId(Id),

		"authorized": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the member is allowed to join the network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the member is allowed to act as a bridge, that is, to forward traffic for MAC " +
				"addresses other than its own.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"inactive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A flag whether the member is currently inactive.",
		},
		"ip_address": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Addresses the controller assigns to this member. They have to fall inside the ranges " +
				"configured on the controller network.",
		},
		"last_seen": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time elapsed since the controller last heard from this member.",
		},
		KeyName: PropNameOptional("Descriptive name of the member."),
		"network": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier network identifier the member belongs to.",
		},
		"zt_address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The 10 hex digit ZeroTier address (node id) of the member.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
