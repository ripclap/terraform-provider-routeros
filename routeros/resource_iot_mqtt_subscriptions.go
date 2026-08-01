package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  `GET /rest/iot/mqtt/subscriptions` returns `[]` on the RouterOS 7.23 - no subscription is configured, so no live record could be captured. The
  shape below lists the wire field names reported by the device; the values are illustrative.

  {
    ".id": "*1",
    "broker": "broker1",
    "on-message": ":log info $msgData",
    "qos": "0",
    "topic": "sensors/#"
  }
*/

// ResourceIotMqttSubscriptions MQTT topic subscriptions of the `iot` package.
// https://help.mikrotik.com/docs/spaces/ROS/pages/46759978/MQTT
func ResourceIotMqttSubscriptions() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/iot/mqtt/subscriptions"),
		MetaId:           PropId(Id),

		"broker": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the broker from `/iot/mqtt/brokers` to subscribe on.",
		},
		"on_message": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Script source executed every time a message arrives on the subscribed topic. The variables " +
				"`$msgData` and `$msgTopic` are available inside the script.",
		},
		"qos": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Quality of Service level requested for the subscription. MQTT defines `0`, `1` and `2`; " +
				"default `0`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"topic": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Topic filter to subscribe to. The MQTT wildcards `+` and `#` are supported.",
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
