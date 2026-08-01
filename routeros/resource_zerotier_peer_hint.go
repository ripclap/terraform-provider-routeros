package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on RouterOS 7.23 (RouterOS 7.23: the default
  instance "zt1" exists but is disabled and no peers are known), so no live JSON sample can be
  shown. The field set below was taken from the device itself:

    /console/inspect request=syntax path="zerotier,peer,hint,add"
      -> addresses, comment, disabled, identity, instance
    /zerotier/peer/hint/print detail            -> Flags: X - DISABLED
    /zerotier/peer/hint/print proplist=<name>   -> "inactive" and "dynamic" do NOT exist here
    /console/inspect request=completion path="zerotier,peer,hint"
      -> "add instance=" offers the instance names ("zt1"); "add addresses=192.0.2.1/9993"
         offers "," so addresses is a comma separated list.
*/

// ResourceZerotierPeerHint pins the transport addresses of a ZeroTier peer.
// https://help.mikrotik.com/docs/spaces/ROS/pages/83755083/ZeroTier
func ResourceZerotierPeerHint() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/zerotier/peer/hint"),
		MetaId:           PropId(Id),

		"addresses": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Transport addresses the peer should be contacted on. The console confirms this is a " +
				"comma separated list. The exact accepted notation of a single element (`ip/port`) could " +
				"not be observed - no peers are known on RouterOS 7.23 and the console performs no value " +
				"check on this attribute.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"identity": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Identity of the peer the hint applies to, i.e. its ZeroTier address / public identity " +
				"as listed in `/zerotier/peer`.",
		},
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ZeroTier instance the hint is installed on.",
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
