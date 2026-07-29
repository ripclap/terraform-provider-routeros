package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*3D",
	  "app-store-url": "",
	  "category": "search",
	  "container-command-lines": "solr:none:docker.io/solr:latest",
	  "custom": "false",
	  "default-network": "internal",
	  "devices": "",
	  "disabled": "true",
	  "environment": "",
	  "extra-mounts": "",
	  "firewall-redirects": "8983:8983:tcp:web",
	  "from-app-store": "false",
	  "interface": "none",
	  "name": "solr",
	  "network": "default",
	  "network-outgoing-access": "true",
	  "network-pvid": "1",
	  "project-page": "https://solr.apache.org",
	  "required-hw-devices": "",
	  "required-mounts": "data",
	  "running": "false",
	  "secrets": "",
	  "status": "",
	  "ui-url": "",
	  "use-https": "true"
	}

	Reference device: RouterOS 7.23.2, `GET /rest/app` (one of 104 entries; the
	bulky `yaml`, `configs`, `cmds` and `description` properties were left out of the sample).

	Settable arguments, `/console/inspect request=syntax path="app,add"` and `path="app,set"`:
	  auto-update  container-command-lines  devices  disabled  environment  extra-mounts
	  firewall-redirects  network  network-outgoing-access  network-pvid  required-hw-devices
	  required-mounts  secrets  use-https  yaml

	`name` is not settable, it is taken from the `yaml` document, and neither are the remaining
	properties reported by the menu. The volatile ones (`app-size`, `cmds`, `configs`, `cpu-usage`,
	`data-size`, `default-credentials`, `memory-current`, `status`,
	`variables-to-use-in-environment`) are skipped, the stable ones are exposed as computed.
*/

// ResourceApp A containerised application deployed from a YAML description.
// https://help.mikrotik.com/docs/spaces/ROS/pages/343244823/Apps
func ResourceApp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/app"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("app_size", "cmds", "configs", "cpu_usage", "data_size",
			"default_credentials", "memory_current", "status", "variables_to_use_in_environment"),

		"app_store_url": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the app store entry the application was installed from.",
		},
		"auto_update": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether this application is updated automatically when a new container image is " +
				"published. Overrides the global `/app/settings` value.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"category": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Category the application declares in its YAML document, used for grouping in WebFig.",
		},
		"container_command_lines": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of `container:command:image` entries overriding the command line " +
				"and the image of the individual containers of the application.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"custom": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the application was defined locally instead of being taken from an app store.",
		},
		"default_network": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Network the YAML document of the application asks for, used when `network=default`.",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description taken from the YAML document of the application.",
		},
		"devices": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of host devices that are passed through to the containers of the " +
				"application, for example a PCI address, `serial0` or a disk name.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"environment": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of `container:VARIABLE=value` entries passed to the containers of " +
				"the application as environment variables.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"extra_mounts": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Comma separated list of additional mounts attached to the containers of the application.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"firewall_redirects": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of `router-port:container-port:protocol:container` entries that " +
				"redirect traffic from the router to the application.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"from_app_store": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the application was installed from an app store.",
		},
		KeyInterface: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the veth interface RouterOS created for the application.",
		},
		"ip_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Address assigned to the application on the selected network.",
		},
		KeyName: {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Name of the application. It is taken from the `name` field of the YAML document and " +
				"cannot be set separately.",
		},
		"network": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Network the application is attached to:" +
				"\n  * default - the network requested by the YAML document of the application," +
				"\n  * internal - an isolated network behind NAT," +
				"\n  * lan - the bridge configured as `lan_bridge` in `/app/settings`," +
				"\n  * the name of a bridge.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"network_outgoing_access": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the application may open outgoing connections. When disabled RouterOS adds a " +
				"mangle drop rule for its traffic.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"network_pvid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Port VLAN ID of the veth interface of the application in the bridge.",
			ValidateFunc:     validation.IntBetween(1, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"project_page": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Home page of the project, taken from the YAML document of the application.",
		},
		"required_hw_devices": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of hardware devices that must be present before the application is " +
				"allowed to start.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"required_mounts": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of mounts that must be available before the application is allowed " +
				"to start.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"secrets": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Comma separated list of `app__secret:=value` entries holding the secrets the application " +
				"asks for, such as generated administrator passwords.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ui_url": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URL of the web interface of the application, when it provides one.",
		},
		"use_https": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the URLs of the application use HTTPS with the certificate reported by " +
				"`/app/settings`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"yaml": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The YAML document describing the application: its name, description, the containers it " +
				"consists of, the mounts and the RouterOS configuration it needs.",
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
