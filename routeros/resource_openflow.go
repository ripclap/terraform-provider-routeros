package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  `GET /rest/openflow` returns `[]` on the reference device (RouterOS 7.23.2) -
  no switch instance is configured, so no live record could be captured. The shape below lists
  the wire field names reported by the device; the values are illustrative.

  {
    ".id": "*1",
    "comment": "",
    "controllers": "tcp/192.0.2.10:6653",
    "datapath-id": "1",
    "disabled": "false",
    "isolate-controllers": "false",
    "name": "ofswitch1",
    "openflow-fast-path-bytes": "0",
    "openflow-fast-path-packets": "0",
    "passive-port": "disabled",
    "verify-peer": "none",
    "version": "default"
  }
*/

// ResourceOpenflow https://help.mikrotik.com/docs/spaces/ROS/pages/295239685/Openflow
func ResourceOpenflow() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/openflow"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("openflow_fast_path_bytes", "openflow_fast_path_packets"),

		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the certificate from `/certificate`. Used together with the `verify_peer` " +
				"parameter for TLS connections to the controller.",
		},
		KeyComment: PropCommentRw,
		"controllers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Connections to the OpenFlow controllers, written as `protocol/address:port`. The device " +
				"accepts the protocols `tcp` and `tls`, for example `[\"tcp/192.0.2.10:6653\"]`.",
		},
		"datapath_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Datapath ID advertised to the controller. MikroTik documents it as two parts separated " +
				"with a slash - an implementer-defined number `0..65535` and a MAC address. VERIFY: the console " +
				"reports this argument as a plain number, so the exact accepted syntax on RouterOS 7.23 could not be " +
				"confirmed without writing to the device; it is declared as a string so that both forms are accepted.",
		},
		KeyDisabled: PropDisabledRw,
		"isolate_controllers": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "VERIFY: accepts `yes`/`no` on the device, but the property is not described in the " +
				"MikroTik documentation, so its exact effect is unconfirmed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Reference name of the OpenFlow switch instance."),
		"passive_port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Port the switch listens on for incoming controller connections. Accepts `disabled` " +
				"(the default) or a port number `1..65535`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"verify_peer": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Verify the peer's identity using certificates.",
			ValidateFunc:     validation.StringInSlice([]string{"if-cert-present", "none", "required"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Version of the OpenFlow standard to be used. Default: `default`.",
			ValidateFunc:     validation.StringInSlice([]string{"1", "1.3", "default"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
