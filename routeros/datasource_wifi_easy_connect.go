package routeros

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/mdp/qrterminal/v3"
)

func DatasourceWiFiEasyConnect() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceQRGenerate,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Authentication type; can be WEP or WPA or WPA2-EAP, or nopass for no password. " +
					"Or, omit for no password.",
				Default:      "WPA",
				ValidateFunc: validation.StringInSlice([]string{"WEP", "WPA2", "WPA2-EAP", "nopass"}, false),
			},
			"ssid": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Network SSID. Required. Enclose in double quotes if it is an ASCII name, but could " +
					"be interpreted as hex (i.e. \"ABCD\").",
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				Description: "Password, ignored if T is nopass (in which case it may be omitted). Enclose in double " +
					"quotes if it is an ASCII name, but could be interpreted as hex (i.e. \"ABCD\").",
			},
			"hidden": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if the network SSID is hidden.",
			},
			"eap_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "(WPA2-EAP only) EAP method, like TTLS or PWD.",
				ValidateFunc: validation.StringInSlice([]string{"TTLS", "PWD"}, false),
			},
			"eap_anonymous": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "(WPA2-EAP only) Anonymous identity",
			},
			"eap_identity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(WPA2-EAP only) Identity.",
			},
			"eap_phase2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(WPA2-EAP only) Phase 2 method, like `MSCHAPV2`",
			},
			"qr_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "QR Code",
			},
		},
	}
}

var mecardEscape = regexp.MustCompile(`([;,":\\])`)

func datasourceQRGenerate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	text := "WIFI:"

	switch d.Get("type").(string) {
	case "", "nopass":
	case "WPA2-EAP":
		text += fmt.Sprintf("E:%v;I:%v;PH2:%v;",
			d.Get("eap_method").(string),
			d.Get("eap_identity").(string),
			d.Get("eap_phase2").(string),
		)
		if d.Get("eap_anonymous").(bool) {
			text += "A:anon;"
		}
	default:
		text += fmt.Sprintf("T:%v;P:%v;",
			d.Get("type").(string),
			mecardEscape.ReplaceAllString(d.Get("password").(string), "\\$1"),
		)
	}
	// SSID
	text += fmt.Sprintf("S:%v;", mecardEscape.ReplaceAllString(d.Get("ssid").(string), "\\$1"))
	// Hidden
	if val := d.Get("hidden").(bool); val {
		text += fmt.Sprintf("H:%v;", val)
	}

	text += ";"

	// The id identifies the network, not its credentials, so it is derived from
	// the SSID and how the network is joined rather than from the payload. The
	// payload carries the passphrase in clear text - that is what the QR code is
	// for - so putting it through a digest would protect nothing while making
	// the id depend on a secret. A data source is read on every plan, so
	// `qr_code` still follows a password change even though the id does not.
	identity := fmt.Sprintf("%v;%v;%v",
		d.Get("ssid").(string), d.Get("type").(string), d.Get("hidden").(bool))
	d.SetId(fmt.Sprintf("%x", sha256.Sum256([]byte(identity))))

	buf := bytes.NewBuffer(nil)
	qrterminal.GenerateWithConfig(text, qrterminal.Config{
		Writer:    buf,
		Level:     qrterminal.L,
		QuietZone: 1,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	})
	if err := d.Set("qr_code", buf.String()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
