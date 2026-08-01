package routeros

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*3",
    "arch": "amd64",
    "interface": "veth1",
    "mounts": "",
    "name": "76e7fc0c-e2c0-4b8c-b3b0-4c985c496a90",
    "os": "linux",
    "repo": "registry-1.docker.io/adguard/adguardhome:latest",
    "root-dir": "sata4-part1/docker/adg",
    "status": "stopped",
    "workdir": "/opt/adguardhome/work"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Properties
func ResourceContainer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("running", "starting", "stopped", "passed_devs", "stopping", "extracting", "config_json"),

		"arch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The architecture of the container image",
		},
		"auto_restart_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specify an interval at which Container will be restarted on Container failure.",
			DiffSuppressFunc: TimeEqual,
		},
		"check_certificate": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables trust chain validation from local certificate store.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cmd": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.",
		},
		KeyComment: PropCommentRw,
		"cpu_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The list of CPUs the container is allowed to run on. Example: 0-3",
		},
		"default_dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of DNS resolvers handed to the container.",
		},
		"devices": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Passes through physical device to the container.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set custom DNS servers",
		},
		"domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container NIS domain name",
		},
		"entrypoint": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An ENTRYPOINT allows to specify executable to run when starting container. Example: /bin/sh",
		},
		"env": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Inline environmental variables to be used with the container. Use `envlist` to reference a list configured under /container envs instead.",
		},
		"envlist": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "list of environmental variables (configured under /container envs ) to be used with container",
		},
		"file": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "container *tar.gz tarball if the container is imported from a file",
		},
		"healthcheck_cmd": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Command executed inside the container to determine whether it is healthy.",
		},
		"healthcheck_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval between health check runs.",
			DiffSuppressFunc: TimeEqual,
		},
		"healthcheck_retries": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of consecutive failed health checks before the container is considered unhealthy.",
		},
		"healthcheck_start_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval between health check runs during the start period.",
			DiffSuppressFunc: TimeEqual,
		},
		"healthcheck_start_period": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time given to the container to start up before failed health checks are counted.",
			DiffSuppressFunc: TimeEqual,
		},
		"healthcheck_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"healthcheck_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time after which a running health check is considered failed.",
			DiffSuppressFunc: TimeEqual,
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container host name",
		},
		"hosts": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Additional entries added to the hosts file of the container.",
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "veth interface to be used with the container",
		},
		"layer_dir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container layers directory.",
		},
		"logging": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "if set to yes, all container-generated output will be shown in the RouterOS log",
		},
		"memory_high": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "RAM usage limit in bytes for a specific container (string value).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"memory_max": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Hard RAM usage limit for a specific container, the device reports `unlimited` when it is not set.",
		},
		"mount": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Mount to be used with this container.",
		},
		"mountlists": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the mount list to be used with this container.",
		},
		"mounts": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Mounts from /container/mounts/ sub-menu to be used with this container",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Assign a name to the container",
		},
		"os": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The OS of the container image",
		},
		"remote_image": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The container image name to be installed if an external registry is used (configured under /container/config set registry-url=...)",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" {
					return false
				}

				if AlwaysPresentNotUserProvided(k, old, new, d) {
					return true
				}

				// Checking the presence of a tag:
				// ~ remote_image = "traefik/whoami:latest" -> "traefik/whoami" # forces replacement
				return old == new || old == new+":latest"
			},
		},
		"restart_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval at which the container is restarted according to the restart policy.",
			DiffSuppressFunc: TimeEqual,
		},
		"restart_max_count": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum number of restart attempts before the container is left stopped.",
		},
		"restart_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Policy that decides when the container is restarted.",
		},
		"root_dir": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Used to save container store outside main memory",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"running": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Container state.",
			Default:     true,
		},
		"shell": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The shell used when executing commands inside the container.",
		},
		"shm_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Size of the shared memory (/dev/shm) of the container.",
		},
		"start_on_boot": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Start the container on boot",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of the container",
		},
		"stop_on_unhealthy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "if set to yes, the container will be stopped when the health check reports it as unhealthy",
		},
		"stop_signal": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Signal to stop the container.",
			Computed:    true,
		},
		"stop_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time to wait for the container to stop before it is killed.",
			DiffSuppressFunc: TimeEqual,
		},
		"tag": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The tag of the container image",
		},
		"tmpfs": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "RAM backed file systems to be mounted inside the container.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets the username used",
		},
		"workdir": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The working directory for cmd entrypoint",
		},
	}

	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Run DefaultRead.
		diags := ResourceRead(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}

		// RouterOS clears `remote-image` once the pull completes, so the image name is
		// reconstructed from `tag`. Containers imported from a tarball also carry a tag,
		// but their source is `file` and no `remote_image` must be derived for them.
		tag, ok := d.Get("tag").(string)
		if ok && tag != "" && d.Get("file").(string) == "" {
			// Get Registry URL.
			res, err := ReadItems(nil, "/container/config", m.(Client))
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
				return diag.FromErr(err)
			}

			if len(*res) == 0 {
				return diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to retrieve the URL of the container registry, the response is empty",
					},
				}
			}

			registryUrl, ok := (*res)[0]["registry-url"]
			if !ok {
				// Key name for an unspecified registry.
				registryUrl, ok = (*res)[0]["assumed-registry-url"]
				if !ok {
					return diag.Diagnostics{
						diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "The `registry-url` was not found in the response",
						},
					}
				}
			}

			u, err := url.Parse(registryUrl)
			if err != nil {
				return diag.FromErr(err)
			}

			// Remove http(s); host:port; path...
			for _, item := range []string{u.Scheme, u.Host, u.Path} {
				tag = strings.TrimPrefix(tag, item)
			}
			// Remove (:////)adguard/adguardhome:latest
			tag = strings.TrimLeft(tag, ":/")

			if err := d.Set("remote_image", strings.TrimPrefix(tag, registryUrl)); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("running", d.Get("status").(string) == "running"); err != nil {
				return diag.FromErr(err)
			}
		}

		return nil
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Run DefaultCreate.
		diags := ResourceCreate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}

		if d.Get("running").(bool) {
			startContainer(ctx, resSchema, d, m)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	resUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		stopContainer(ctx, resSchema, d, m)

		// Run DefaultUpdate.
		diags := ResourceUpdate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}
		if d.Get("running").(bool) {
			startContainer(ctx, resSchema, d, m)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Stop container
		stopContainer(ctx, resSchema, d, m)

		// Run DefaultDelete.
		return ResourceDelete(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CustomizeDiff: containerSourceCheck,
		CreateContext: resCreate,
		ReadContext:   resRead,
		UpdateContext: resUpdate,
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func startContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	stopStateConf := &retry.StateChangeConf{
		Pending: []string{"pulling", "extracting"},
		Target:  []string{"stopped"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err := stopStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be pulled: %s", d.Id(), err)
		return diag.FromErr(err)
	}

	item := MikrotikItem{"number": d.Id()}

	// Start container
	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/start"
	}

	err = m.(Client).SendRequest(crudStart, resUrl, item, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	startStateConf := &retry.StateChangeConf{
		Pending: []string{"stopped"},
		Target:  []string{"running"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err = startStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be started: %s", d.Id(), err)
		return diag.FromErr(err)
	}

	return nil
}

func stopContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item := MikrotikItem{"number": d.Id()}

	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/stop"
	}

	err := m.(Client).SendRequest(crudStop, resUrl, item, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	stopStateConf := &retry.StateChangeConf{
		Pending: []string{"stopping"},
		Target:  []string{"stopped"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutDelete),
	}
	_, err = stopStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be stopped: %s", d.Id(), err)
		return diag.FromErr(err)
	}
	return nil
}

// containerSourceCheck enforces "exactly one of file / remote_image" on non-empty values.
// ExactlyOneOf cannot be used: the device reports file="" alongside a real remote_image,
// and Terraform treats an empty string as a specified value.
func containerSourceCheck(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	file, _ := d.Get("file").(string)
	image, _ := d.Get("remote_image").(string)

	if file != "" && image != "" {
		return fmt.Errorf("only one of `file`, `remote_image` can be set")
	}
	if file == "" && image == "" {
		return fmt.Errorf("one of `file`, `remote_image` must be set")
	}

	return nil
}
