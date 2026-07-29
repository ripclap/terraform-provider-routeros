package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "allowed-number": "",
	  "channel": "0",
	  "last-ussd": "",
	  "polling": "false",
	  "port": "none",
	  "receive-enabled": "false",
	  "remove-sent-sms-after-send": "false",
	  "secret": "",
	  "sim-pin": "",
	  "sms-storage": "sim",
	  "status": "off"
	}
*/

// ResourceToolSms SMS tool settings, used with a serial or LTE modem attached to the router.
// https://help.mikrotik.com/docs/spaces/ROS/pages/26476608/SMS
func ResourceToolSms() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/sms"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("last_ussd", "status"),

		"allowed_number": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of phone numbers that are allowed to execute SMS commands on the " +
				"router. An empty value allows any number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"channel": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Modem channel that is used to send and receive messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"polling": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Use polling of the modem instead of unsolicited result codes. Required by modems that " +
				"do not notify the router about new messages on their own.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the serial port or of the modem interface used for SMS, for example `serial0`. " +
				"`none` disables the SMS tool.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"receive_enabled": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the router receives and stores incoming messages in `/tool/sms/inbox`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remove_sent_sms_after_send": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether a message is deleted from the modem or SIM storage right after it was sent.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secret": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Secret that a received message must start with before the router treats the rest of the " +
				"message as an SMS command.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sim_pin": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "PIN code of the SIM card.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sms_storage": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Where received messages are stored.",
			ValidateFunc:     validation.StringInSlice([]string{"modem", "sim"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
