package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/ppp-server` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,ppp-server,add"
	  authentication  comment  copy-from  data-channel  disabled  max-mru  max-mtu  modem-init
	  mrru  name  null-modem  port  profile  ring-count

	Read-only fields reported by the same device: running
*/

// ResourceInterfacePppServer PPP server interface on a serial port or a modem.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328072/PPP
func ResourceInterfacePppServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ppp-server"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Authentication methods that the server will accept.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"data_channel": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The modem channel that carries the PPP data connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"max_mru": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Receive Unit. Maximum packet size that the interface will be able to receive " +
				"without packet fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Transmission Unit. Maximum packet size that the interface will be able to send " +
				"without packet fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"modem_init": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Modem initialization string sent before the modem starts to answer calls.",
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, " +
				"it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the " +
				"tunnel. Set to `disabled` to turn Multilink PPP off.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Descriptive name of the interface."),
		"null_modem": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables null-modem mode. When enabled, no modem initialization strings are sent and calls " +
				"are not answered.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The serial port the modem is connected to, as listed by the `/port` menu.",
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which PPP profile configuration will be used for the incoming connections.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ring_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of rings to wait before answering the call.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
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
