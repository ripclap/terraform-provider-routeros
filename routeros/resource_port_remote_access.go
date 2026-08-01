package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23, so no JSON sample is available.
*/

// ResourcePortRemoteAccess Makes a serial port reachable over TCP/UDP (RFC 2217 and raw modes).
// https://help.mikrotik.com/docs/spaces/ROS/pages/8978525/Ports
func ResourcePortRemoteAccess() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/port/remote-access"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("active", "active_peer", "active_peer_port", "busy", "logging_active"),

		"channel": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Channel of the serial port that is exported.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"ip_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "TCP or UDP port number the router listens on, or connects to when " +
				"`protocol=tcp-client`.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"local_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Local address the service binds to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"log_file": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the file the data exchanged over the port is logged to.",
		},
		"port": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the `/port` entry that is made remotely accessible, for example `serial0`.",
		},
		"protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Protocol used to carry the serial data:" +
				"\n  * rfc2217 - Telnet Com Port Control Option, the remote side can also change the line settings," +
				"\n  * tcp-server - the router listens for a raw TCP connection," +
				"\n  * tcp-client - the router connects to a remote raw TCP server," +
				"\n  * udp - raw data over UDP.",
			ValidateFunc:     validation.StringInSlice([]string{"rfc2217", "tcp-client", "tcp-server", "udp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_addresses": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of addresses that are allowed to connect to the port, or the " +
				"address to connect to when `protocol=tcp-client`.",
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
