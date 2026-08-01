package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/ppp-client` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,ppp-client,add"
	  add-default-route  allow  apn  comment  copy-from  data-channel  default-route-distance
	  dial-command  dial-on-demand  disabled  info-channel  keepalive-timeout  max-mru  max-mtu
	  modem-init  mrru  name  network-mode  null-modem  password  phone  pin  port  profile
	  remote-address  use-peer-dns  user

	Read-only fields reported by the same device: running

	`port` completes to the serial ports of the device (`serial0` on RouterOS 7.23). The
	channel selectors `data-channel` and `info-channel` are kept as strings because RouterOS does
	not report their empty value shape on a device without a modem.
*/

// ResourceInterfacePppClient PPP client interface over a serial port or a modem.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328072/PPP
func ResourceInterfacePppClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ppp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to add the PPP remote address as a default route.",
		},
		"allow": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Allowed authentication methods.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"apn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The Access Point Name used by the modem to establish the packet data connection.",
		},
		KeyComment: PropCommentRw,
		"data_channel": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The modem channel that carries the PPP data connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_route_distance": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Distance value applied to the auto created default route, if add-default-route is selected.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"dial_command": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The modem AT dial command used to establish the connection.",
		},
		"dial_on_demand": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Dial out only when outbound traffic is generated.",
		},
		KeyDisabled: PropDisabledRw,
		"info_channel": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The modem channel used to query the modem status information.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"keepalive_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Tunnel keepalive timeout in seconds.",
			DiffSuppressFunc: TimeEqual,
		},
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
			Description: "Modem initialization string sent before dialling.",
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
		"network_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The cellular network mode the modem is allowed to use.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "lte-m", "nb-iot"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"null_modem": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables null-modem mode. When enabled, no modem initialization strings are sent and no " +
				"dialling takes place.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used for authentication.",
		},
		"phone": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Phone number used for dial out.",
		},
		"pin": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The PIN code of the SIM card.",
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The serial port the modem is connected to, as listed by the `/port` menu.",
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which PPP profile configuration will be used when establishing the connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The IP address assigned to the remote end of the link.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"use_peer_dns": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to use the DNS server settings advertised by the remote server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used for authentication.",
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
