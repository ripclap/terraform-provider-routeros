package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on the reference device (RouterOS 7.23.2), so no live
  JSON sample can be shown. The field set below was taken from the device itself:

    /console/inspect request=syntax path="user-manager,payment,add"
      -> currency, method, price, profile, trans-end, trans-start, trans-status, user,
         user-message
    /user-manager/payment/print detail          -> no flags at all
    /user-manager/payment/print proplist=<name> -> "comment" and "dynamic" do NOT exist here
    /user-manager/payment/print where method=<v>
      -> the console accepts only "paypal" and "authorize-net"
    /user-manager/payment/print where trans-status=<v>
      -> the console accepts only "aborted", "approved", "declined", "error", "pending",
         "started", "timeout", "user-approved"
*/

// ResourceUserManagerPayment manages the User Manager payment ledger.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2555940/User+Manager
func ResourceUserManagerPayment() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/payment"),
		MetaId:           PropId(Id),

		"currency": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Currency the payment is denominated in, as configured in `/user-manager/advanced`.",
		},
		"method": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Payment gateway that processed the transaction.",
			ValidateFunc:     validation.StringInSlice([]string{"authorize-net", "paypal"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"price": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Description:  "The amount charged for the profile.",
			ValidateFunc: validation.FloatAtLeast(.0),
		},
		"profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the `/user-manager/profile` that was bought.",
		},
		"trans_end": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Date and time the transaction was completed, in RouterOS format (`jan/02/2006 15:04:05`).",
		},
		"trans_start": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Date and time the transaction was started, in RouterOS format (`jan/02/2006 15:04:05`).",
		},
		"trans_status": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "State of the transaction:" +
				"\n  * started - the transaction has been opened towards the gateway;" +
				"\n  * pending - the gateway has not settled the payment yet;" +
				"\n  * user-approved - the user confirmed the payment at the gateway;" +
				"\n  * approved - the payment was accepted and the profile was granted;" +
				"\n  * declined - the gateway refused the payment;" +
				"\n  * aborted - the transaction was abandoned;" +
				"\n  * timeout - the gateway did not answer in time;" +
				"\n  * error - the transaction failed.",
			ValidateFunc: validation.StringInSlice([]string{"aborted", "approved", "declined", "error",
				"pending", "started", "timeout", "user-approved"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the `/user-manager/user` the payment belongs to.",
		},
		"user_message": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Free-form message shown to the user for this payment.",
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
