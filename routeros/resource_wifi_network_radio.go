package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/wifi/network/radio` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,wifi,network,radio,add"
	  channel.band  channel.deprioritize-unii-3-4  channel.frequency  channel.preamble-puncturing
	  channel.reselect-interval  channel.reselect-time  channel.secondary-frequency
	  channel.skip-dfs-channels  channel.width  comment  configuration.antenna-gain
	  configuration.chains  configuration.country  configuration.distance
	  configuration.installation  configuration.tx-chains  configuration.tx-power  copy-from
	  disabled  extra-labels  labels  security.authentication-types

	The menu reports no read-only property and no `name` property; the entries are addressed by
	their `.id`.

	The `channel.`, `configuration.` and `security.` groups are inline sub-configurations without a
	corresponding profile reference, so they are exposed as maps in the same way `routeros_wifi`
	exposes its inline settings.
*/

// ResourceWifiNetworkRadio Radio settings of the label based WiFi provisioning model.
// https://help.mikrotik.com/docs/spaces/ROS/pages/224559120/WiFi
func ResourceWifiNetworkRadio() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/network/radio"),
		MetaId:           PropId(Id),

		"channel": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Channel inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"configuration": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Configuration inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"extra_labels": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Additional labels assigned to the radios matched by this entry, so that other entries can " +
				"select them. Accepts a comma separated list of labels and the `+`/`-` prefixes.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"labels": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Labels that select the radios this entry applies to. Accepts a comma separated list of " +
				"labels, `all`, and the `+`/`-` prefixes to add or remove a label from the selection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"security": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Security inline settings. Only `authentication_types` is accepted by this menu.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		Description: `*<span style="color:red">The '/interface/wifi/network/radio' menu was verified on RouterOS ` +
			`7.23; the exact minimum RouterOS version is not documented.</span>*`,
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
