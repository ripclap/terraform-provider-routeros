package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/wifi/network` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,wifi,network,add"
	  beacon-interval  comment  copy-from  datapath.bridge  datapath.bridge-cost
	  datapath.bridge-horizon  datapath.client-isolation  datapath.interface-list
	  datapath.openflow-switch  datapath.traffic-processing  datapath.vlan-id  disabled
	  dtim-period  hide-ssid  hw-protection-mode  labels  max-clients  mlo  mode
	  multicast-enhance  qos-classifier  security.<...>  ssid  station-roaming

	The menu reports no read-only property and no `name` property; the entries are addressed by
	their `.id`.

	The `datapath.` and `security.` groups are inline sub-configurations without a corresponding
	profile reference, so they are exposed as maps in the same way `routeros_wifi` exposes its
	inline settings. The full list of `security.` keys reported by the device is:
	  authentication-types  beacon-protection  connect-group  connect-priority  dh-groups
	  disable-pmkid  eap-accounting  eap-anonymous-identity  eap-certificate-mode  eap-methods
	  eap-password  eap-tls-certificate  eap-username  encryption  ft  ft-mobility-domain
	  ft-nas-identifier  ft-over-ds  ft-preserve-vlanid  ft-r0-key-lifetime
	  ft-reassociation-deadline  group-encryption  group-key-update  management-encryption
	  management-protection  multi-passphrase-group  owe-transition-interface  passphrase
	  sae-anti-clogging-threshold  sae-max-failure-rate  sae-pwe  wps
*/

// ResourceWifiNetwork A WiFi network of the label based WiFi provisioning model.
// https://help.mikrotik.com/docs/spaces/ROS/pages/224559120/WiFi
func ResourceWifiNetwork() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/network"),
		MetaId:           PropId(Id),

		"beacon_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time interval between beacon frames.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyComment: PropCommentRw,
		"datapath": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Datapath inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"dtim_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "A period at which to transmit multicast traffic, when there are client devices in power " +
				"save mode connected to the AP.",
			ValidateFunc:     validation.IntBetween(1, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hide_ssid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Setting it to yes removes this network from the list of wireless networks that are shown " +
				"by some client software. Changing this setting does not improve the security of the wireless " +
				"network, because the SSID is included in other frames sent by the AP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_protection_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Frame protection support mode used to avoid the hidden node problem.",
			ValidateFunc:     validation.StringInSlice([]string{"cts-to-self", "none", "rts-cts"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"labels": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Labels that select the radios this network is provisioned on. Accepts a comma separated " +
				"list of labels, `all`, and the `+`/`-` prefixes to add or remove a label from the selection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_clients": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of associated client devices.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mlo": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables Multi-Link Operation (802.11be) for this network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "An option to specify the operational mode of the network.",
			ValidateFunc: validation.StringInSlice([]string{"ap", "station", "station-bridge", "station-pseudobridge"}, false),
		},
		"multicast_enhance": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "An option to enable converting every multicast-address IP or IPv6 packet into multiple " +
				"unicast-addressed frames for each connected station.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"qos_classifier": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to specify the QoS classifier.",
			ValidateFunc:     validation.StringInSlice([]string{"dscp-high-3-bits", "priority"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"security": {
			Type:             schema.TypeMap,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			Description:      "Security inline settings.",
			ValidateDiagFunc: ValidationMapKeyNames,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SSID (service set identifier) of the network.",
		},
		"station_roaming": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables 802.11k neighbour reports for the client devices of this network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		Description: `*<span style="color:red">The '/interface/wifi/network' menu was verified on RouterOS 7.23.2; ` +
			`VERIFY: the exact minimum RouterOS version could not be established.</span>*`,
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
