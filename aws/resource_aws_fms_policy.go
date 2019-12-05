package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-gp/service/fms"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAwsFmsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsFmsPolicyCreate,
		Read:   resourceAwsFmsPolicyRead,
		Update: resourceAwsFmsPolicyUpdate,
		Delete: resourceAwsFmsPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_accounts": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exclude_accounts": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exclude_resource_tags": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"remediation_enabled": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_tags": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 8,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
			"resource_type_list": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 128,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_service_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"WAF",
								"SHIELD_ADVANCED",
								"SECURITY_GROUPS_COMMON",
								"SECURITY_GROUPS_CONTENT_AUDIT",
								"SECURITY_GROUPS_USAGE_AUDIT",
							}, false),
						},
						"managed_service_data": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAwsFmsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).fmsconn

	input := &fms.
	return resourceAwsFmsPolicyRead(d, meta)
}

func resourceAwsFmsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	// conn := meta.(*AWSClient).fmsconn

	return nil
}

func resourceAwsFmsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	// conn := meta.(*AWSClient).fmsconn

	return nil
}

func resourceAwsFmsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	// conn := meta.(*AWSClient).fmsconn

	return nil
}
