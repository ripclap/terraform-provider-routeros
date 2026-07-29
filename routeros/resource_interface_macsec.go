package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/macsec` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,macsec,add"
	  cak  ckn  comment  copy-from  disabled  interface  mtu  name  profile

	Read-only fields reported by the same device:
	  hw-offloaded  inactive  running  status, plus the MACsec statistics counters (all the
	  rx-sa, rx-sc, rx-untagged, tx-sa, tx-sc, tx-untagged ... properties), which are skipped.
*/

// ResourceInterfaceMacsec MACsec (IEEE 802.1AE) interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/201523202/MACsec
func ResourceInterfaceMacsec() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/macsec"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("rx_bad_tag", "rx_no_sci", "rx_no_tag", "rx_overrun", "rx_sa_invalid",
			"rx_sa_not_using_sa", "rx_sa_not_valid", "rx_sa_ok", "rx_sa_unused_sa", "rx_sc_decrypted_byte",
			"rx_sc_delayed", "rx_sc_invalid", "rx_sc_late", "rx_sc_not_using_sa", "rx_sc_not_valid", "rx_sc_ok",
			"rx_sc_unchecked", "rx_sc_unused_sa", "rx_sc_validated_byte", "rx_unknown_sci", "rx_untagged",
			"tx_sa_encrypted", "tx_sa_protected", "tx_sc_encrypted_byte", "tx_sc_encrypted_packet",
			"tx_sc_protected_byte", "tx_sc_protected_packet", "tx_too_long", "tx_untagged"),

		"cak": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "A 16-byte pre-shared Connectivity Association Key, written as a 32 character hex string. " +
				"RouterOS generates the key automatically when only the Ethernet interface is specified.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ckn": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A 32-byte Connectivity Association Name, written as a 64 character hex string. RouterOS " +
				"generates the name automatically when only the Ethernet interface is specified.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyHwOffloaded: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the MACsec encryption of this interface is offloaded to the hardware.",
		},
		KeyInactive: PropInactiveRo,
		KeyInterface: {
			Type:     schema.TypeString,
			Required: true,
			Description: "The Ethernet interface MACsec is deployed on. Only one MACsec interface can be created " +
				"per Ethernet port.",
		},
		"mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Layer3 Maximum Transmission Unit of the MACsec interface.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the MACsec interface."),
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The MACsec profile that determines the key server election parameters.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Current state of the MACsec Key Agreement: disabled, initializing, invalid, negotiating " +
				"or open-encrypted.",
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
