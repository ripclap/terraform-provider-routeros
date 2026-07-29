package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "acme-plain": "true",
  "crl-plain": "true",
  "graphs-plain": "true",
  "graphs-secure": "true",
  "index-plain": "true",
  "index-secure": "true",
  "rest-plain": "true",
  "rest-secure": "true",
  "scep-plain": "true",
  "webfig-plain": "true",
  "webfig-secure": "true"
}
*/

/*
	On/off switches for the endpoints served by the `www` (HTTP) and `www-ssl` (HTTPS) services.
*/

// ResourceIpServiceWebserver
// https://help.mikrotik.com/docs/spaces/ROS/pages/103841820/Services
func ResourceIpServiceWebserver() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/service/webserver"),
		MetaId:           PropId(Id),

		"acme_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the ACME (Let's Encrypt) HTTP-01 challenge over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the Certificate Revocation List over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"graphs_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the graphing pages over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"graphs_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the graphing pages over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"index_plain": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Serve the home/login page over HTTP. It can be turned off once `webfig_plain` and " +
				"`graphs_plain` are off.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"index_secure": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Serve the home/login page over HTTPS. It can be turned off once `webfig_secure` and " +
				"`graphs_secure` are off.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rest_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the REST API over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rest_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the REST API over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"scep_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve SCEP (Simple Certificate Enrollment Protocol) over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"webfig_plain": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the WebFig interface over HTTP.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"webfig_secure": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Serve the WebFig interface over HTTPS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
