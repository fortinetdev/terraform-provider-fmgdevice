// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Action for automation stitches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationActionCreate,
		Read:   resourceSystemAutomationActionRead,
		Update: resourceSystemAutomationActionUpdate,
		Delete: resourceSystemAutomationActionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alicloud_access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_access_key_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"alicloud_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_function": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alicloud_function_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alicloud_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_api_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"aws_api_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_api_stage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"azure_app": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_function": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"email_body": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gcp_function": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_function_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_function_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"execute_security_fabric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"forticare_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_body": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"minimum_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"output_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"security_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tls_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_host_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAutomationActionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutomationAction(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationAction resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationAction(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationAction resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationActionRead(d, m)
}

func resourceSystemAutomationActionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutomationAction(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationAction resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationAction(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationAction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationActionRead(d, m)
}

func resourceSystemAutomationActionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemAutomationAction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationActionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemAutomationAction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationAction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationActionAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionActionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudAccessKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunctionDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiStage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureApp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailBody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionGcpFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpFunctionDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpFunctionRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionExecuteSecurityFabric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHeaders(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionForticareEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpBody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemAutomationActionHttpHeadersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemAutomationAction-HttpHeaders-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			v := flattenSystemAutomationActionHttpHeadersKey(i["key"], d, pre_append)
			tmp["key"] = fortiAPISubPartPatch(v, "SystemAutomationAction-HttpHeaders-Key")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemAutomationActionHttpHeadersValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemAutomationAction-HttpHeaders-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAutomationActionHttpHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeadersKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeadersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMessageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMinimumInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionReplacementMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionSdnConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionSecurityTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionSystemAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionTlsCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationActionUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionVerifyHostCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationAction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accprofile", flattenSystemAutomationActionAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemAutomationAction-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("action_type", flattenSystemAutomationActionActionType(o["action-type"], d, "action_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["action-type"], "SystemAutomationAction-ActionType"); ok {
			if err = d.Set("action_type", vv); err != nil {
				return fmt.Errorf("Error reading action_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("alicloud_access_key_id", flattenSystemAutomationActionAlicloudAccessKeyId(o["alicloud-access-key-id"], d, "alicloud_access_key_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-access-key-id"], "SystemAutomationAction-AlicloudAccessKeyId"); ok {
			if err = d.Set("alicloud_access_key_id", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_access_key_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_access_key_id: %v", err)
		}
	}

	if err = d.Set("alicloud_account_id", flattenSystemAutomationActionAlicloudAccountId(o["alicloud-account-id"], d, "alicloud_account_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-account-id"], "SystemAutomationAction-AlicloudAccountId"); ok {
			if err = d.Set("alicloud_account_id", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_account_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_account_id: %v", err)
		}
	}

	if err = d.Set("alicloud_function", flattenSystemAutomationActionAlicloudFunction(o["alicloud-function"], d, "alicloud_function")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-function"], "SystemAutomationAction-AlicloudFunction"); ok {
			if err = d.Set("alicloud_function", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_function: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_function: %v", err)
		}
	}

	if err = d.Set("alicloud_function_authorization", flattenSystemAutomationActionAlicloudFunctionAuthorization(o["alicloud-function-authorization"], d, "alicloud_function_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-function-authorization"], "SystemAutomationAction-AlicloudFunctionAuthorization"); ok {
			if err = d.Set("alicloud_function_authorization", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_function_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_function_authorization: %v", err)
		}
	}

	if err = d.Set("alicloud_function_domain", flattenSystemAutomationActionAlicloudFunctionDomain(o["alicloud-function-domain"], d, "alicloud_function_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-function-domain"], "SystemAutomationAction-AlicloudFunctionDomain"); ok {
			if err = d.Set("alicloud_function_domain", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_function_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_function_domain: %v", err)
		}
	}

	if err = d.Set("alicloud_region", flattenSystemAutomationActionAlicloudRegion(o["alicloud-region"], d, "alicloud_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-region"], "SystemAutomationAction-AlicloudRegion"); ok {
			if err = d.Set("alicloud_region", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_region: %v", err)
		}
	}

	if err = d.Set("alicloud_service", flattenSystemAutomationActionAlicloudService(o["alicloud-service"], d, "alicloud_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-service"], "SystemAutomationAction-AlicloudService"); ok {
			if err = d.Set("alicloud_service", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_service: %v", err)
		}
	}

	if err = d.Set("alicloud_version", flattenSystemAutomationActionAlicloudVersion(o["alicloud-version"], d, "alicloud_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["alicloud-version"], "SystemAutomationAction-AlicloudVersion"); ok {
			if err = d.Set("alicloud_version", vv); err != nil {
				return fmt.Errorf("Error reading alicloud_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alicloud_version: %v", err)
		}
	}

	if err = d.Set("aws_api_id", flattenSystemAutomationActionAwsApiId(o["aws-api-id"], d, "aws_api_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["aws-api-id"], "SystemAutomationAction-AwsApiId"); ok {
			if err = d.Set("aws_api_id", vv); err != nil {
				return fmt.Errorf("Error reading aws_api_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aws_api_id: %v", err)
		}
	}

	if err = d.Set("aws_api_path", flattenSystemAutomationActionAwsApiPath(o["aws-api-path"], d, "aws_api_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["aws-api-path"], "SystemAutomationAction-AwsApiPath"); ok {
			if err = d.Set("aws_api_path", vv); err != nil {
				return fmt.Errorf("Error reading aws_api_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aws_api_path: %v", err)
		}
	}

	if err = d.Set("aws_api_stage", flattenSystemAutomationActionAwsApiStage(o["aws-api-stage"], d, "aws_api_stage")); err != nil {
		if vv, ok := fortiAPIPatch(o["aws-api-stage"], "SystemAutomationAction-AwsApiStage"); ok {
			if err = d.Set("aws_api_stage", vv); err != nil {
				return fmt.Errorf("Error reading aws_api_stage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aws_api_stage: %v", err)
		}
	}

	if err = d.Set("aws_domain", flattenSystemAutomationActionAwsDomain(o["aws-domain"], d, "aws_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["aws-domain"], "SystemAutomationAction-AwsDomain"); ok {
			if err = d.Set("aws_domain", vv); err != nil {
				return fmt.Errorf("Error reading aws_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aws_domain: %v", err)
		}
	}

	if err = d.Set("aws_region", flattenSystemAutomationActionAwsRegion(o["aws-region"], d, "aws_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["aws-region"], "SystemAutomationAction-AwsRegion"); ok {
			if err = d.Set("aws_region", vv); err != nil {
				return fmt.Errorf("Error reading aws_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("azure_app", flattenSystemAutomationActionAzureApp(o["azure-app"], d, "azure_app")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-app"], "SystemAutomationAction-AzureApp"); ok {
			if err = d.Set("azure_app", vv); err != nil {
				return fmt.Errorf("Error reading azure_app: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_app: %v", err)
		}
	}

	if err = d.Set("azure_domain", flattenSystemAutomationActionAzureDomain(o["azure-domain"], d, "azure_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-domain"], "SystemAutomationAction-AzureDomain"); ok {
			if err = d.Set("azure_domain", vv); err != nil {
				return fmt.Errorf("Error reading azure_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_domain: %v", err)
		}
	}

	if err = d.Set("azure_function", flattenSystemAutomationActionAzureFunction(o["azure-function"], d, "azure_function")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-function"], "SystemAutomationAction-AzureFunction"); ok {
			if err = d.Set("azure_function", vv); err != nil {
				return fmt.Errorf("Error reading azure_function: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_function: %v", err)
		}
	}

	if err = d.Set("azure_function_authorization", flattenSystemAutomationActionAzureFunctionAuthorization(o["azure-function-authorization"], d, "azure_function_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-function-authorization"], "SystemAutomationAction-AzureFunctionAuthorization"); ok {
			if err = d.Set("azure_function_authorization", vv); err != nil {
				return fmt.Errorf("Error reading azure_function_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_function_authorization: %v", err)
		}
	}

	if err = d.Set("delay", flattenSystemAutomationActionDelay(o["delay"], d, "delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay"], "SystemAutomationAction-Delay"); ok {
			if err = d.Set("delay", vv); err != nil {
				return fmt.Errorf("Error reading delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay: %v", err)
		}
	}

	if err = d.Set("email_body", flattenSystemAutomationActionEmailBody(o["email-body"], d, "email_body")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-body"], "SystemAutomationAction-EmailBody"); ok {
			if err = d.Set("email_body", vv); err != nil {
				return fmt.Errorf("Error reading email_body: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_body: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationActionDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAutomationAction-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("email_from", flattenSystemAutomationActionEmailFrom(o["email-from"], d, "email_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-from"], "SystemAutomationAction-EmailFrom"); ok {
			if err = d.Set("email_from", vv); err != nil {
				return fmt.Errorf("Error reading email_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_from: %v", err)
		}
	}

	if err = d.Set("email_subject", flattenSystemAutomationActionEmailSubject(o["email-subject"], d, "email_subject")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-subject"], "SystemAutomationAction-EmailSubject"); ok {
			if err = d.Set("email_subject", vv); err != nil {
				return fmt.Errorf("Error reading email_subject: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_subject: %v", err)
		}
	}

	if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "SystemAutomationAction-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("gcp_function", flattenSystemAutomationActionGcpFunction(o["gcp-function"], d, "gcp_function")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-function"], "SystemAutomationAction-GcpFunction"); ok {
			if err = d.Set("gcp_function", vv); err != nil {
				return fmt.Errorf("Error reading gcp_function: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_function: %v", err)
		}
	}

	if err = d.Set("gcp_function_domain", flattenSystemAutomationActionGcpFunctionDomain(o["gcp-function-domain"], d, "gcp_function_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-function-domain"], "SystemAutomationAction-GcpFunctionDomain"); ok {
			if err = d.Set("gcp_function_domain", vv); err != nil {
				return fmt.Errorf("Error reading gcp_function_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_function_domain: %v", err)
		}
	}

	if err = d.Set("gcp_function_region", flattenSystemAutomationActionGcpFunctionRegion(o["gcp-function-region"], d, "gcp_function_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-function-region"], "SystemAutomationAction-GcpFunctionRegion"); ok {
			if err = d.Set("gcp_function_region", vv); err != nil {
				return fmt.Errorf("Error reading gcp_function_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_function_region: %v", err)
		}
	}

	if err = d.Set("gcp_project", flattenSystemAutomationActionGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-project"], "SystemAutomationAction-GcpProject"); ok {
			if err = d.Set("gcp_project", vv); err != nil {
				return fmt.Errorf("Error reading gcp_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("execute_security_fabric", flattenSystemAutomationActionExecuteSecurityFabric(o["execute-security-fabric"], d, "execute_security_fabric")); err != nil {
		if vv, ok := fortiAPIPatch(o["execute-security-fabric"], "SystemAutomationAction-ExecuteSecurityFabric"); ok {
			if err = d.Set("execute_security_fabric", vv); err != nil {
				return fmt.Errorf("Error reading execute_security_fabric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading execute_security_fabric: %v", err)
		}
	}

	if err = d.Set("headers", flattenSystemAutomationActionHeaders(o["headers"], d, "headers")); err != nil {
		if vv, ok := fortiAPIPatch(o["headers"], "SystemAutomationAction-Headers"); ok {
			if err = d.Set("headers", vv); err != nil {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading headers: %v", err)
		}
	}

	if err = d.Set("forticare_email", flattenSystemAutomationActionForticareEmail(o["forticare-email"], d, "forticare_email")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticare-email"], "SystemAutomationAction-ForticareEmail"); ok {
			if err = d.Set("forticare_email", vv); err != nil {
				return fmt.Errorf("Error reading forticare_email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticare_email: %v", err)
		}
	}

	if err = d.Set("http_body", flattenSystemAutomationActionHttpBody(o["http-body"], d, "http_body")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-body"], "SystemAutomationAction-HttpBody"); ok {
			if err = d.Set("http_body", vv); err != nil {
				return fmt.Errorf("Error reading http_body: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_body: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("http_headers", flattenSystemAutomationActionHttpHeaders(o["http-headers"], d, "http_headers")); err != nil {
			if vv, ok := fortiAPIPatch(o["http-headers"], "SystemAutomationAction-HttpHeaders"); ok {
				if err = d.Set("http_headers", vv); err != nil {
					return fmt.Errorf("Error reading http_headers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http_headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http_headers"); ok {
			if err = d.Set("http_headers", flattenSystemAutomationActionHttpHeaders(o["http-headers"], d, "http_headers")); err != nil {
				if vv, ok := fortiAPIPatch(o["http-headers"], "SystemAutomationAction-HttpHeaders"); ok {
					if err = d.Set("http_headers", vv); err != nil {
						return fmt.Errorf("Error reading http_headers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http_headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("message", flattenSystemAutomationActionMessage(o["message"], d, "message")); err != nil {
		if vv, ok := fortiAPIPatch(o["message"], "SystemAutomationAction-Message"); ok {
			if err = d.Set("message", vv); err != nil {
				return fmt.Errorf("Error reading message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message: %v", err)
		}
	}

	if err = d.Set("message_type", flattenSystemAutomationActionMessageType(o["message-type"], d, "message_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-type"], "SystemAutomationAction-MessageType"); ok {
			if err = d.Set("message_type", vv); err != nil {
				return fmt.Errorf("Error reading message_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_type: %v", err)
		}
	}

	if err = d.Set("method", flattenSystemAutomationActionMethod(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "SystemAutomationAction-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("minimum_interval", flattenSystemAutomationActionMinimumInterval(o["minimum-interval"], d, "minimum_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-interval"], "SystemAutomationAction-MinimumInterval"); ok {
			if err = d.Set("minimum_interval", vv); err != nil {
				return fmt.Errorf("Error reading minimum_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationActionName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationAction-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("output_size", flattenSystemAutomationActionOutputSize(o["output-size"], d, "output_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["output-size"], "SystemAutomationAction-OutputSize"); ok {
			if err = d.Set("output_size", vv); err != nil {
				return fmt.Errorf("Error reading output_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutomationActionPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemAutomationAction-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemAutomationActionProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemAutomationAction-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("replacement_message", flattenSystemAutomationActionReplacementMessage(o["replacement-message"], d, "replacement_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacement-message"], "SystemAutomationAction-ReplacementMessage"); ok {
			if err = d.Set("replacement_message", vv); err != nil {
				return fmt.Errorf("Error reading replacement_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacement_message: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenSystemAutomationActionReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "SystemAutomationAction-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("required", flattenSystemAutomationActionRequired(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "SystemAutomationAction-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemAutomationActionScript(o["script"], d, "script")); err != nil {
		if vv, ok := fortiAPIPatch(o["script"], "SystemAutomationAction-Script"); ok {
			if err = d.Set("script", vv); err != nil {
				return fmt.Errorf("Error reading script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-connector"], "SystemAutomationAction-SdnConnector"); ok {
			if err = d.Set("sdn_connector", vv); err != nil {
				return fmt.Errorf("Error reading sdn_connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	if err = d.Set("security_tag", flattenSystemAutomationActionSecurityTag(o["security-tag"], d, "security_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-tag"], "SystemAutomationAction-SecurityTag"); ok {
			if err = d.Set("security_tag", vv); err != nil {
				return fmt.Errorf("Error reading security_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_tag: %v", err)
		}
	}

	if err = d.Set("system_action", flattenSystemAutomationActionSystemAction(o["system-action"], d, "system_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-action"], "SystemAutomationAction-SystemAction"); ok {
			if err = d.Set("system_action", vv); err != nil {
				return fmt.Errorf("Error reading system_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_action: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemAutomationActionTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystemAutomationAction-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("tls_certificate", flattenSystemAutomationActionTlsCertificate(o["tls-certificate"], d, "tls_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-certificate"], "SystemAutomationAction-TlsCertificate"); ok {
			if err = d.Set("tls_certificate", vv); err != nil {
				return fmt.Errorf("Error reading tls_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_certificate: %v", err)
		}
	}

	if err = d.Set("uri", flattenSystemAutomationActionUri(o["uri"], d, "uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["uri"], "SystemAutomationAction-Uri"); ok {
			if err = d.Set("uri", vv); err != nil {
				return fmt.Errorf("Error reading uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("verify_host_cert", flattenSystemAutomationActionVerifyHostCert(o["verify-host-cert"], d, "verify_host_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-host-cert"], "SystemAutomationAction-VerifyHostCert"); ok {
			if err = d.Set("verify_host_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_host_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_host_cert: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationActionAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionActionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudAccessKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudAccessKeySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionAlicloudAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunctionAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunctionDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionAwsApiPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiStage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionAzureApp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureFunction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureFunctionAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailBody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionGcpFunction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpFunctionDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpFunctionRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionExecuteSecurityFabric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionForticareEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpBody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemAutomationActionHttpHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandSystemAutomationActionHttpHeadersKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemAutomationActionHttpHeadersValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionHttpHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMessageType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMinimumInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionOutputSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionReplacementMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSdnConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionSecurityTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSystemAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionTlsCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationActionUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionVerifyHostCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationAction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemAutomationActionAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("action_type"); ok || d.HasChange("action_type") {
		t, err := expandSystemAutomationActionActionType(d, v, "action_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action-type"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_access_key_id"); ok || d.HasChange("alicloud_access_key_id") {
		t, err := expandSystemAutomationActionAlicloudAccessKeyId(d, v, "alicloud_access_key_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-access-key-id"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_access_key_secret"); ok || d.HasChange("alicloud_access_key_secret") {
		t, err := expandSystemAutomationActionAlicloudAccessKeySecret(d, v, "alicloud_access_key_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-access-key-secret"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_account_id"); ok || d.HasChange("alicloud_account_id") {
		t, err := expandSystemAutomationActionAlicloudAccountId(d, v, "alicloud_account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-account-id"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_function"); ok || d.HasChange("alicloud_function") {
		t, err := expandSystemAutomationActionAlicloudFunction(d, v, "alicloud_function")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_function_authorization"); ok || d.HasChange("alicloud_function_authorization") {
		t, err := expandSystemAutomationActionAlicloudFunctionAuthorization(d, v, "alicloud_function_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function-authorization"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_function_domain"); ok || d.HasChange("alicloud_function_domain") {
		t, err := expandSystemAutomationActionAlicloudFunctionDomain(d, v, "alicloud_function_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function-domain"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_region"); ok || d.HasChange("alicloud_region") {
		t, err := expandSystemAutomationActionAlicloudRegion(d, v, "alicloud_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-region"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_service"); ok || d.HasChange("alicloud_service") {
		t, err := expandSystemAutomationActionAlicloudService(d, v, "alicloud_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-service"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_version"); ok || d.HasChange("alicloud_version") {
		t, err := expandSystemAutomationActionAlicloudVersion(d, v, "alicloud_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-version"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_id"); ok || d.HasChange("aws_api_id") {
		t, err := expandSystemAutomationActionAwsApiId(d, v, "aws_api_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-id"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_key"); ok || d.HasChange("aws_api_key") {
		t, err := expandSystemAutomationActionAwsApiKey(d, v, "aws_api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-key"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_path"); ok || d.HasChange("aws_api_path") {
		t, err := expandSystemAutomationActionAwsApiPath(d, v, "aws_api_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-path"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_stage"); ok || d.HasChange("aws_api_stage") {
		t, err := expandSystemAutomationActionAwsApiStage(d, v, "aws_api_stage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-stage"] = t
		}
	}

	if v, ok := d.GetOk("aws_domain"); ok || d.HasChange("aws_domain") {
		t, err := expandSystemAutomationActionAwsDomain(d, v, "aws_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-domain"] = t
		}
	}

	if v, ok := d.GetOk("aws_region"); ok || d.HasChange("aws_region") {
		t, err := expandSystemAutomationActionAwsRegion(d, v, "aws_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-region"] = t
		}
	}

	if v, ok := d.GetOk("azure_api_key"); ok || d.HasChange("azure_api_key") {
		t, err := expandSystemAutomationActionAzureApiKey(d, v, "azure_api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-api-key"] = t
		}
	}

	if v, ok := d.GetOk("azure_app"); ok || d.HasChange("azure_app") {
		t, err := expandSystemAutomationActionAzureApp(d, v, "azure_app")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-app"] = t
		}
	}

	if v, ok := d.GetOk("azure_domain"); ok || d.HasChange("azure_domain") {
		t, err := expandSystemAutomationActionAzureDomain(d, v, "azure_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-domain"] = t
		}
	}

	if v, ok := d.GetOk("azure_function"); ok || d.HasChange("azure_function") {
		t, err := expandSystemAutomationActionAzureFunction(d, v, "azure_function")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-function"] = t
		}
	}

	if v, ok := d.GetOk("azure_function_authorization"); ok || d.HasChange("azure_function_authorization") {
		t, err := expandSystemAutomationActionAzureFunctionAuthorization(d, v, "azure_function_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-function-authorization"] = t
		}
	}

	if v, ok := d.GetOk("delay"); ok || d.HasChange("delay") {
		t, err := expandSystemAutomationActionDelay(d, v, "delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay"] = t
		}
	}

	if v, ok := d.GetOk("email_body"); ok || d.HasChange("email_body") {
		t, err := expandSystemAutomationActionEmailBody(d, v, "email_body")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-body"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAutomationActionDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("email_from"); ok || d.HasChange("email_from") {
		t, err := expandSystemAutomationActionEmailFrom(d, v, "email_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-from"] = t
		}
	}

	if v, ok := d.GetOk("email_subject"); ok || d.HasChange("email_subject") {
		t, err := expandSystemAutomationActionEmailSubject(d, v, "email_subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-subject"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandSystemAutomationActionEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("gcp_function"); ok || d.HasChange("gcp_function") {
		t, err := expandSystemAutomationActionGcpFunction(d, v, "gcp_function")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function"] = t
		}
	}

	if v, ok := d.GetOk("gcp_function_domain"); ok || d.HasChange("gcp_function_domain") {
		t, err := expandSystemAutomationActionGcpFunctionDomain(d, v, "gcp_function_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function-domain"] = t
		}
	}

	if v, ok := d.GetOk("gcp_function_region"); ok || d.HasChange("gcp_function_region") {
		t, err := expandSystemAutomationActionGcpFunctionRegion(d, v, "gcp_function_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function-region"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project"); ok || d.HasChange("gcp_project") {
		t, err := expandSystemAutomationActionGcpProject(d, v, "gcp_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	}

	if v, ok := d.GetOk("execute_security_fabric"); ok || d.HasChange("execute_security_fabric") {
		t, err := expandSystemAutomationActionExecuteSecurityFabric(d, v, "execute_security_fabric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["execute-security-fabric"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok || d.HasChange("headers") {
		t, err := expandSystemAutomationActionHeaders(d, v, "headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	if v, ok := d.GetOk("forticare_email"); ok || d.HasChange("forticare_email") {
		t, err := expandSystemAutomationActionForticareEmail(d, v, "forticare_email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticare-email"] = t
		}
	}

	if v, ok := d.GetOk("http_body"); ok || d.HasChange("http_body") {
		t, err := expandSystemAutomationActionHttpBody(d, v, "http_body")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-body"] = t
		}
	}

	if v, ok := d.GetOk("http_headers"); ok || d.HasChange("http_headers") {
		t, err := expandSystemAutomationActionHttpHeaders(d, v, "http_headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-headers"] = t
		}
	}

	if v, ok := d.GetOk("message"); ok || d.HasChange("message") {
		t, err := expandSystemAutomationActionMessage(d, v, "message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message"] = t
		}
	}

	if v, ok := d.GetOk("message_type"); ok || d.HasChange("message_type") {
		t, err := expandSystemAutomationActionMessageType(d, v, "message_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-type"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandSystemAutomationActionMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("minimum_interval"); ok || d.HasChange("minimum_interval") {
		t, err := expandSystemAutomationActionMinimumInterval(d, v, "minimum_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationActionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("output_size"); ok || d.HasChange("output_size") {
		t, err := expandSystemAutomationActionOutputSize(d, v, "output_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-size"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemAutomationActionPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemAutomationActionProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("replacement_message"); ok || d.HasChange("replacement_message") {
		t, err := expandSystemAutomationActionReplacementMessage(d, v, "replacement_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacement-message"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandSystemAutomationActionReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandSystemAutomationActionRequired(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok || d.HasChange("script") {
		t, err := expandSystemAutomationActionScript(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("sdn_connector"); ok || d.HasChange("sdn_connector") {
		t, err := expandSystemAutomationActionSdnConnector(d, v, "sdn_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-connector"] = t
		}
	}

	if v, ok := d.GetOk("security_tag"); ok || d.HasChange("security_tag") {
		t, err := expandSystemAutomationActionSecurityTag(d, v, "security_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-tag"] = t
		}
	}

	if v, ok := d.GetOk("system_action"); ok || d.HasChange("system_action") {
		t, err := expandSystemAutomationActionSystemAction(d, v, "system_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-action"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystemAutomationActionTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("tls_certificate"); ok || d.HasChange("tls_certificate") {
		t, err := expandSystemAutomationActionTlsCertificate(d, v, "tls_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-certificate"] = t
		}
	}

	if v, ok := d.GetOk("uri"); ok || d.HasChange("uri") {
		t, err := expandSystemAutomationActionUri(d, v, "uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri"] = t
		}
	}

	if v, ok := d.GetOk("verify_host_cert"); ok || d.HasChange("verify_host_cert") {
		t, err := expandSystemAutomationActionVerifyHostCert(d, v, "verify_host_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-host-cert"] = t
		}
	}

	return &obj, nil
}
