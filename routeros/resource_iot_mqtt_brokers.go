package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  `GET /rest/iot/mqtt/brokers` returns `[]` on the RouterOS 7.23 - no broker is configured, so no live record could be captured. The shape
  below lists the wire field names reported by the device; the values are illustrative.

  {
    ".id": "*1",
    "address": "192.0.2.10",
    "auto-connect": "true",
    "certificate": "none",
    "client-id": "ccr",
    "connected": "false",
    "keep-alive": "60",
    "name": "broker1",
    "parallel-scripts-limit": "off",
    "password": "",
    "port": "1883",
    "ssl": "false",
    "username": "",
    "will-message": "",
    "will-qos": "0",
    "will-retain": "false",
    "will-topic": ""
  }
*/

// ResourceIotMqttBrokers MQTT client broker connections of the `iot` package.
// https://help.mikrotik.com/docs/spaces/ROS/pages/46759978/MQTT
func ResourceIotMqttBrokers() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/iot/mqtt/brokers"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP address or hostname of the broker. Must be set for the client to be able to connect.",
		},
		"auto_connect": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Keeps the connection to the broker up and reconnects automatically when it drops.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the certificate from `/certificate` used for the SSL connection.",
		},
		"client_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Unique client ID used for the MQTT connection.",
		},
		"connected": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the client is currently connected to the broker.",
		},
		"keep_alive": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Time in seconds between the client pings sent to the broker. MikroTik documents the range " +
				"as 30..64800, default `60`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Descriptive name of the broker."),
		"parallel_scripts_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum number of scripts that the `on-message` subscription handler may run in parallel. " +
				"Accepts `off` (no limit tracking, the default) or a number; MikroTik documents the numeric range as " +
				"3..1000.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used to authenticate to the broker, if the broker requires one.",
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Network port used by the broker. Default: `1883`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssl": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Use a TLS/SSL protected connection to the broker.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Username used to authenticate to the broker, if the broker requires one.",
		},
		"will_message": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Last Will and Testament message published by the broker on an unexpected disconnect.",
		},
		"will_qos": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Quality of Service value used for the Last Will and Testament message. MQTT defines " +
				"`0`, `1` and `2`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"will_retain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the broker should retain the Last Will and Testament message.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"will_topic": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Topic the Last Will and Testament message is published to.",
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
