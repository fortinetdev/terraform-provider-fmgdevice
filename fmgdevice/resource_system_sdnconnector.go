// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure connection to SDN Connector.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorCreate,
		Read:   resourceSystemSdnConnectorRead,
		Update: resourceSystemSdnConnectorUpdate,
		Delete: resourceSystemSdnConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alt_resource_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"azure_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"compartment_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compartment_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"compute_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_account_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"region_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"role_arn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"external_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"forwarding_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"gcp_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_project_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gcp_zone_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ibm_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibm_region_gen1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ibm_region_gen2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"microsoft_365": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"private_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"public_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"resource_group": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"peer_nic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"nsx_cert_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oci_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_region_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rest_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rest_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"rest_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rest_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"route_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"next_hop": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"subscription_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"secret_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"vcenter_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"verify_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmx_image_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmx_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSdnConnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSdnConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnector resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdnConnector(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdnConnector(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdnConnector resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdnConnector(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdnConnector resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSdnConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemSdnConnector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnConnector(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnector resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorAccessKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorAltResourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorCompartmentList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "compartment_id"
		if _, ok := i["compartment-id"]; ok {
			v := flattenSystemSdnConnectorCompartmentListCompartmentId(i["compartment-id"], d, pre_append)
			tmp["compartment_id"] = fortiAPISubPartPatch(v, "SystemSdnConnector-CompartmentList-CompartmentId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorCompartmentListCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_id"
		if _, ok := i["external-id"]; ok {
			v := flattenSystemSdnConnectorExternalAccountListExternalId(i["external-id"], d, pre_append)
			tmp["external_id"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ExternalAccountList-ExternalId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region_list"
		if _, ok := i["region-list"]; ok {
			v := flattenSystemSdnConnectorExternalAccountListRegionList(i["region-list"], d, pre_append)
			tmp["region_list"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ExternalAccountList-RegionList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role_arn"
		if _, ok := i["role-arn"]; ok {
			v := flattenSystemSdnConnectorExternalAccountListRoleArn(i["role-arn"], d, pre_append)
			tmp["role_arn"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ExternalAccountList-RoleArn")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorExternalAccountListExternalId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountListRegionList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorExternalAccountListRoleArn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorExternalIpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ExternalIp-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorForwardingRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_name"
		if _, ok := i["rule-name"]; ok {
			v := flattenSystemSdnConnectorForwardingRuleRuleName(i["rule-name"], d, pre_append)
			tmp["rule_name"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ForwardingRule-RuleName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenSystemSdnConnectorForwardingRuleTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "SystemSdnConnector-ForwardingRule-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorForwardingRuleRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorForwardingRuleTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProjectList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gcp_zone_list"
		if _, ok := i["gcp-zone-list"]; ok {
			v := flattenSystemSdnConnectorGcpProjectListGcpZoneList(i["gcp-zone-list"], d, pre_append)
			tmp["gcp_zone_list"] = fortiAPISubPartPatch(v, "SystemSdnConnector-GcpProjectList-GcpZoneList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdnConnectorGcpProjectListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdnConnector-GcpProjectList-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorGcpProjectListGcpZoneList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorGcpProjectListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegionGen1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegionGen2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorMessageServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorMicrosoft365(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSdnConnectorNicIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSdnConnector-Nic-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorNicName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnector-Nic-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_nic"
		if _, ok := i["peer-nic"]; ok {
			v := flattenSystemSdnConnectorNicPeerNic(i["peer-nic"], d, pre_append)
			tmp["peer_nic"] = fortiAPISubPartPatch(v, "SystemSdnConnector-Nic-PeerNic")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorNicIpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnectorNic-Ip-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "private_ip"
		if _, ok := i["private-ip"]; ok {
			v := flattenSystemSdnConnectorNicIpPrivateIp(i["private-ip"], d, pre_append)
			tmp["private_ip"] = fortiAPISubPartPatch(v, "SystemSdnConnectorNic-Ip-PrivateIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := i["public-ip"]; ok {
			v := flattenSystemSdnConnectorNicIpPublicIp(i["public-ip"], d, pre_append)
			tmp["public_ip"] = fortiAPISubPartPatch(v, "SystemSdnConnectorNic-Ip-PublicIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			v := flattenSystemSdnConnectorNicIpResourceGroup(i["resource-group"], d, pre_append)
			tmp["resource_group"] = fortiAPISubPartPatch(v, "SystemSdnConnectorNic-Ip-ResourceGroup")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpPrivateIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicPeerNic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorNsxCertFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := i["region"]; ok {
			v := flattenSystemSdnConnectorOciRegionListRegion(i["region"], d, pre_append)
			tmp["region"] = fortiAPISubPartPatch(v, "SystemSdnConnector-OciRegionList-Region")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorOciRegionListRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRestInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRestSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRestSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorRouteName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnector-Route-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorRouteTableName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnector-RouteTable-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			v := flattenSystemSdnConnectorRouteTableResourceGroup(i["resource-group"], d, pre_append)
			tmp["resource_group"] = fortiAPISubPartPatch(v, "SystemSdnConnector-RouteTable-ResourceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {
			v := flattenSystemSdnConnectorRouteTableRoute(i["route"], d, pre_append)
			tmp["route"] = fortiAPISubPartPatch(v, "SystemSdnConnector-RouteTable-Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := i["subscription-id"]; ok {
			v := flattenSystemSdnConnectorRouteTableSubscriptionId(i["subscription-id"], d, pre_append)
			tmp["subscription_id"] = fortiAPISubPartPatch(v, "SystemSdnConnector-RouteTable-SubscriptionId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdnConnectorRouteTableRouteName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdnConnectorRouteTable-Route-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {
			v := flattenSystemSdnConnectorRouteTableRouteNextHop(i["next-hop"], d, pre_append)
			tmp["next_hop"] = fortiAPISubPartPatch(v, "SystemSdnConnectorRouteTable-Route-NextHop")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorServerList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnConnectorVerifyCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVmxImageUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVmxServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_local_cert", flattenSystemSdnConnectorLocalCert(o["_local_cert"], d, "_local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["_local_cert"], "SystemSdnConnector-LocalCert"); ok {
			if err = d.Set("_local_cert", vv); err != nil {
				return fmt.Errorf("Error reading _local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _local_cert: %v", err)
		}
	}

	if err = d.Set("access_key", flattenSystemSdnConnectorAccessKey(o["access-key"], d, "access_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-key"], "SystemSdnConnector-AccessKey"); ok {
			if err = d.Set("access_key", vv); err != nil {
				return fmt.Errorf("Error reading access_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_key: %v", err)
		}
	}

	if err = d.Set("alt_resource_ip", flattenSystemSdnConnectorAltResourceIp(o["alt-resource-ip"], d, "alt_resource_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-resource-ip"], "SystemSdnConnector-AltResourceIp"); ok {
			if err = d.Set("alt_resource_ip", vv); err != nil {
				return fmt.Errorf("Error reading alt_resource_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_resource_ip: %v", err)
		}
	}

	if err = d.Set("azure_region", flattenSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-region"], "SystemSdnConnector-AzureRegion"); ok {
			if err = d.Set("azure_region", vv); err != nil {
				return fmt.Errorf("Error reading azure_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_region: %v", err)
		}
	}

	if err = d.Set("client_id", flattenSystemSdnConnectorClientId(o["client-id"], d, "client_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-id"], "SystemSdnConnector-ClientId"); ok {
			if err = d.Set("client_id", vv); err != nil {
				return fmt.Errorf("Error reading client_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("compartment_list", flattenSystemSdnConnectorCompartmentList(o["compartment-list"], d, "compartment_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["compartment-list"], "SystemSdnConnector-CompartmentList"); ok {
				if err = d.Set("compartment_list", vv); err != nil {
					return fmt.Errorf("Error reading compartment_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading compartment_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("compartment_list"); ok {
			if err = d.Set("compartment_list", flattenSystemSdnConnectorCompartmentList(o["compartment-list"], d, "compartment_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["compartment-list"], "SystemSdnConnector-CompartmentList"); ok {
					if err = d.Set("compartment_list", vv); err != nil {
						return fmt.Errorf("Error reading compartment_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading compartment_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("compute_generation", flattenSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["compute-generation"], "SystemSdnConnector-ComputeGeneration"); ok {
			if err = d.Set("compute_generation", vv); err != nil {
				return fmt.Errorf("Error reading compute_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemSdnConnectorDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemSdnConnector-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("external_account_list", flattenSystemSdnConnectorExternalAccountList(o["external-account-list"], d, "external_account_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["external-account-list"], "SystemSdnConnector-ExternalAccountList"); ok {
				if err = d.Set("external_account_list", vv); err != nil {
					return fmt.Errorf("Error reading external_account_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading external_account_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_account_list"); ok {
			if err = d.Set("external_account_list", flattenSystemSdnConnectorExternalAccountList(o["external-account-list"], d, "external_account_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["external-account-list"], "SystemSdnConnector-ExternalAccountList"); ok {
					if err = d.Set("external_account_list", vv); err != nil {
						return fmt.Errorf("Error reading external_account_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading external_account_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["external-ip"], "SystemSdnConnector-ExternalIp"); ok {
				if err = d.Set("external_ip", vv); err != nil {
					return fmt.Errorf("Error reading external_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading external_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_ip"); ok {
			if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["external-ip"], "SystemSdnConnector-ExternalIp"); ok {
					if err = d.Set("external_ip", vv); err != nil {
						return fmt.Errorf("Error reading external_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading external_ip: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("forwarding_rule", flattenSystemSdnConnectorForwardingRule(o["forwarding-rule"], d, "forwarding_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["forwarding-rule"], "SystemSdnConnector-ForwardingRule"); ok {
				if err = d.Set("forwarding_rule", vv); err != nil {
					return fmt.Errorf("Error reading forwarding_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading forwarding_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forwarding_rule"); ok {
			if err = d.Set("forwarding_rule", flattenSystemSdnConnectorForwardingRule(o["forwarding-rule"], d, "forwarding_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["forwarding-rule"], "SystemSdnConnector-ForwardingRule"); ok {
					if err = d.Set("forwarding_rule", vv); err != nil {
						return fmt.Errorf("Error reading forwarding_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading forwarding_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("gcp_project", flattenSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-project"], "SystemSdnConnector-GcpProject"); ok {
			if err = d.Set("gcp_project", vv); err != nil {
				return fmt.Errorf("Error reading gcp_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gcp_project_list", flattenSystemSdnConnectorGcpProjectList(o["gcp-project-list"], d, "gcp_project_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["gcp-project-list"], "SystemSdnConnector-GcpProjectList"); ok {
				if err = d.Set("gcp_project_list", vv); err != nil {
					return fmt.Errorf("Error reading gcp_project_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading gcp_project_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gcp_project_list"); ok {
			if err = d.Set("gcp_project_list", flattenSystemSdnConnectorGcpProjectList(o["gcp-project-list"], d, "gcp_project_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["gcp-project-list"], "SystemSdnConnector-GcpProjectList"); ok {
					if err = d.Set("gcp_project_list", vv); err != nil {
						return fmt.Errorf("Error reading gcp_project_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading gcp_project_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_name", flattenSystemSdnConnectorGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "SystemSdnConnector-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("ha_status", flattenSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-status"], "SystemSdnConnector-HaStatus"); ok {
			if err = d.Set("ha_status", vv); err != nil {
				return fmt.Errorf("Error reading ha_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_status: %v", err)
		}
	}

	if err = d.Set("ibm_region", flattenSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region"], "SystemSdnConnector-IbmRegion"); ok {
			if err = d.Set("ibm_region", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen1", flattenSystemSdnConnectorIbmRegionGen1(o["ibm-region-gen1"], d, "ibm_region_gen1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region-gen1"], "SystemSdnConnector-IbmRegionGen1"); ok {
			if err = d.Set("ibm_region_gen1", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen2", flattenSystemSdnConnectorIbmRegionGen2(o["ibm-region-gen2"], d, "ibm_region_gen2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region-gen2"], "SystemSdnConnector-IbmRegionGen2"); ok {
			if err = d.Set("ibm_region_gen2", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
		}
	}

	if err = d.Set("login_endpoint", flattenSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-endpoint"], "SystemSdnConnector-LoginEndpoint"); ok {
			if err = d.Set("login_endpoint", vv); err != nil {
				return fmt.Errorf("Error reading login_endpoint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("message_server_port", flattenSystemSdnConnectorMessageServerPort(o["message-server-port"], d, "message_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-server-port"], "SystemSdnConnector-MessageServerPort"); ok {
			if err = d.Set("message_server_port", vv); err != nil {
				return fmt.Errorf("Error reading message_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_server_port: %v", err)
		}
	}

	if err = d.Set("microsoft_365", flattenSystemSdnConnectorMicrosoft365(o["microsoft-365"], d, "microsoft_365")); err != nil {
		if vv, ok := fortiAPIPatch(o["microsoft-365"], "SystemSdnConnector-Microsoft365"); ok {
			if err = d.Set("microsoft_365", vv); err != nil {
				return fmt.Errorf("Error reading microsoft_365: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading microsoft_365: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdnConnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdnConnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
			if vv, ok := fortiAPIPatch(o["nic"], "SystemSdnConnector-Nic"); ok {
				if err = d.Set("nic", vv); err != nil {
					return fmt.Errorf("Error reading nic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nic"); ok {
			if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
				if vv, ok := fortiAPIPatch(o["nic"], "SystemSdnConnector-Nic"); ok {
					if err = d.Set("nic", vv); err != nil {
						return fmt.Errorf("Error reading nic: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nic: %v", err)
				}
			}
		}
	}

	if err = d.Set("nsx_cert_fingerprint", flattenSystemSdnConnectorNsxCertFingerprint(o["nsx-cert-fingerprint"], d, "nsx_cert_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["nsx-cert-fingerprint"], "SystemSdnConnector-NsxCertFingerprint"); ok {
			if err = d.Set("nsx_cert_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading nsx_cert_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nsx_cert_fingerprint: %v", err)
		}
	}

	if err = d.Set("oci_cert", flattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-cert"], "SystemSdnConnector-OciCert"); ok {
			if err = d.Set("oci_cert", vv); err != nil {
				return fmt.Errorf("Error reading oci_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", flattenSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-fingerprint"], "SystemSdnConnector-OciFingerprint"); ok {
			if err = d.Set("oci_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading oci_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_fingerprint: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("oci_region_list", flattenSystemSdnConnectorOciRegionList(o["oci-region-list"], d, "oci_region_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["oci-region-list"], "SystemSdnConnector-OciRegionList"); ok {
				if err = d.Set("oci_region_list", vv); err != nil {
					return fmt.Errorf("Error reading oci_region_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading oci_region_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("oci_region_list"); ok {
			if err = d.Set("oci_region_list", flattenSystemSdnConnectorOciRegionList(o["oci-region-list"], d, "oci_region_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["oci-region-list"], "SystemSdnConnector-OciRegionList"); ok {
					if err = d.Set("oci_region_list", vv); err != nil {
						return fmt.Errorf("Error reading oci_region_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading oci_region_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("oci_region_type", flattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-region-type"], "SystemSdnConnector-OciRegionType"); ok {
			if err = d.Set("oci_region_type", vv); err != nil {
				return fmt.Errorf("Error reading oci_region_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("private_key", flattenSystemSdnConnectorPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "SystemSdnConnector-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemSdnConnectorProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "SystemSdnConnector-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("region", flattenSystemSdnConnectorRegion(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "SystemSdnConnector-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-group"], "SystemSdnConnector-ResourceGroup"); ok {
			if err = d.Set("resource_group", vv); err != nil {
				return fmt.Errorf("Error reading resource_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	if err = d.Set("resource_url", flattenSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-url"], "SystemSdnConnector-ResourceUrl"); ok {
			if err = d.Set("resource_url", vv); err != nil {
				return fmt.Errorf("Error reading resource_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_url: %v", err)
		}
	}

	if err = d.Set("rest_interface", flattenSystemSdnConnectorRestInterface(o["rest-interface"], d, "rest_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-interface"], "SystemSdnConnector-RestInterface"); ok {
			if err = d.Set("rest_interface", vv); err != nil {
				return fmt.Errorf("Error reading rest_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_interface: %v", err)
		}
	}

	if err = d.Set("rest_sport", flattenSystemSdnConnectorRestSport(o["rest-sport"], d, "rest_sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-sport"], "SystemSdnConnector-RestSport"); ok {
			if err = d.Set("rest_sport", vv); err != nil {
				return fmt.Errorf("Error reading rest_sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_sport: %v", err)
		}
	}

	if err = d.Set("rest_ssl", flattenSystemSdnConnectorRestSsl(o["rest-ssl"], d, "rest_ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-ssl"], "SystemSdnConnector-RestSsl"); ok {
			if err = d.Set("rest_ssl", vv); err != nil {
				return fmt.Errorf("Error reading rest_ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_ssl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
			if vv, ok := fortiAPIPatch(o["route"], "SystemSdnConnector-Route"); ok {
				if err = d.Set("route", vv); err != nil {
					return fmt.Errorf("Error reading route: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route"); ok {
			if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
				if vv, ok := fortiAPIPatch(o["route"], "SystemSdnConnector-Route"); ok {
					if err = d.Set("route", vv); err != nil {
						return fmt.Errorf("Error reading route: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
			if vv, ok := fortiAPIPatch(o["route-table"], "SystemSdnConnector-RouteTable"); ok {
				if err = d.Set("route_table", vv); err != nil {
					return fmt.Errorf("Error reading route_table: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route_table: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_table"); ok {
			if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
				if vv, ok := fortiAPIPatch(o["route-table"], "SystemSdnConnector-RouteTable"); ok {
					if err = d.Set("route_table", vv); err != nil {
						return fmt.Errorf("Error reading route_table: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route_table: %v", err)
				}
			}
		}
	}

	if err = d.Set("secret_token", flattenSystemSdnConnectorSecretToken(o["secret-token"], d, "secret_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["secret-token"], "SystemSdnConnector-SecretToken"); ok {
			if err = d.Set("secret_token", vv); err != nil {
				return fmt.Errorf("Error reading secret_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secret_token: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdnConnectorServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemSdnConnector-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_ca_cert", flattenSystemSdnConnectorServerCaCert(o["server-ca-cert"], d, "server_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ca-cert"], "SystemSdnConnector-ServerCaCert"); ok {
			if err = d.Set("server_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading server_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ca_cert: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenSystemSdnConnectorServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-cert"], "SystemSdnConnector-ServerCert"); ok {
			if err = d.Set("server_cert", vv); err != nil {
				return fmt.Errorf("Error reading server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_list", flattenSystemSdnConnectorServerList(o["server-list"], d, "server_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-list"], "SystemSdnConnector-ServerList"); ok {
			if err = d.Set("server_list", vv); err != nil {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_list: %v", err)
		}
	}

	if err = d.Set("server_port", flattenSystemSdnConnectorServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "SystemSdnConnector-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("service_account", flattenSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-account"], "SystemSdnConnector-ServiceAccount"); ok {
			if err = d.Set("service_account", vv); err != nil {
				return fmt.Errorf("Error reading service_account: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_account: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdnConnectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSdnConnector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("subscription_id", flattenSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscription-id"], "SystemSdnConnector-SubscriptionId"); ok {
			if err = d.Set("subscription_id", vv); err != nil {
				return fmt.Errorf("Error reading subscription_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-id"], "SystemSdnConnector-TenantId"); ok {
			if err = d.Set("tenant_id", vv); err != nil {
				return fmt.Errorf("Error reading tenant_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnConnectorType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemSdnConnector-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "SystemSdnConnector-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", flattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-metadata-iam"], "SystemSdnConnector-UseMetadataIam"); ok {
			if err = d.Set("use_metadata_iam", vv); err != nil {
				return fmt.Errorf("Error reading use_metadata_iam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("user_id", flattenSystemSdnConnectorUserId(o["user-id"], d, "user_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-id"], "SystemSdnConnector-UserId"); ok {
			if err = d.Set("user_id", vv); err != nil {
				return fmt.Errorf("Error reading user_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemSdnConnectorUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemSdnConnector-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_server", flattenSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcenter-server"], "SystemSdnConnector-VcenterServer"); ok {
			if err = d.Set("vcenter_server", vv); err != nil {
				return fmt.Errorf("Error reading vcenter_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", flattenSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcenter-username"], "SystemSdnConnector-VcenterUsername"); ok {
			if err = d.Set("vcenter_username", vv); err != nil {
				return fmt.Errorf("Error reading vcenter_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemSdnConnectorVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemSdnConnector-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("verify_certificate", flattenSystemSdnConnectorVerifyCertificate(o["verify-certificate"], d, "verify_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-certificate"], "SystemSdnConnector-VerifyCertificate"); ok {
			if err = d.Set("verify_certificate", vv); err != nil {
				return fmt.Errorf("Error reading verify_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_certificate: %v", err)
		}
	}

	if err = d.Set("vmx_image_url", flattenSystemSdnConnectorVmxImageUrl(o["vmx-image-url"], d, "vmx_image_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["vmx-image-url"], "SystemSdnConnector-VmxImageUrl"); ok {
			if err = d.Set("vmx_image_url", vv); err != nil {
				return fmt.Errorf("Error reading vmx_image_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vmx_image_url: %v", err)
		}
	}

	if err = d.Set("vmx_service_name", flattenSystemSdnConnectorVmxServiceName(o["vmx-service-name"], d, "vmx_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vmx-service-name"], "SystemSdnConnector-VmxServiceName"); ok {
			if err = d.Set("vmx_service_name", vv); err != nil {
				return fmt.Errorf("Error reading vmx_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vmx_service_name: %v", err)
		}
	}

	if err = d.Set("vpc_id", flattenSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-id"], "SystemSdnConnector-VpcId"); ok {
			if err = d.Set("vpc_id", vv); err != nil {
				return fmt.Errorf("Error reading vpc_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_id: %v", err)
		}
	}

	if err = d.Set("compartment_id", flattenSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["compartment-id"], "SystemSdnConnector-CompartmentId"); ok {
			if err = d.Set("compartment_id", vv); err != nil {
				return fmt.Errorf("Error reading compartment_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("oci_region", flattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-region"], "SystemSdnConnector-OciRegion"); ok {
			if err = d.Set("oci_region", vv); err != nil {
				return fmt.Errorf("Error reading oci_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_region: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAccessKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAltResourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorAzureRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorCompartmentList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "compartment_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["compartment-id"], _ = expandSystemSdnConnectorCompartmentListCompartmentId(d, i["compartment_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorCompartmentListCompartmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorComputeGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-id"], _ = expandSystemSdnConnectorExternalAccountListExternalId(d, i["external_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["region-list"], _ = expandSystemSdnConnectorExternalAccountListRegionList(d, i["region_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role_arn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role-arn"], _ = expandSystemSdnConnectorExternalAccountListRoleArn(d, i["role_arn"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalAccountListExternalId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountListRegionList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorExternalAccountListRoleArn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorExternalIpName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalIpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-name"], _ = expandSystemSdnConnectorForwardingRuleRuleName(d, i["rule_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandSystemSdnConnectorForwardingRuleTarget(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorForwardingRuleRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRuleTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProjectList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gcp_zone_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gcp-zone-list"], _ = expandSystemSdnConnectorGcpProjectListGcpZoneList(d, i["gcp_zone_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdnConnectorGcpProjectListId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorGcpProjectListGcpZoneList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorGcpProjectListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorHaStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegionGen1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegionGen2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorKeyPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorLoginEndpoint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorMessageServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorMicrosoft365(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSdnConnectorNicIp(d, i["ip"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ip"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorNicName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_nic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer-nic"], _ = expandSystemSdnConnectorNicPeerNic(d, i["peer_nic"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorNicIpName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "private_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["private-ip"], _ = expandSystemSdnConnectorNicIpPrivateIp(d, i["private_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["public-ip"], _ = expandSystemSdnConnectorNicIpPublicIp(d, i["public_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["resource-group"], _ = expandSystemSdnConnectorNicIpResourceGroup(d, i["resource_group"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicIpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpPrivateIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpPublicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicPeerNic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNsxCertFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorOciFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["region"], _ = expandSystemSdnConnectorOciRegionListRegion(d, i["region"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorOciRegionListRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRestInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorRestSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRestSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorRouteName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorRouteTableName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["resource-group"], _ = expandSystemSdnConnectorRouteTableResourceGroup(d, i["resource_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSdnConnectorRouteTableRoute(d, i["route"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["route"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subscription-id"], _ = expandSystemSdnConnectorRouteTableSubscriptionId(d, i["subscription_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdnConnectorRouteTableRouteName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["next-hop"], _ = expandSystemSdnConnectorRouteTableRouteNextHop(d, i["next_hop"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRouteNextHop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableSubscriptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorSecretToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServiceAccount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSubscriptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorTenantId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorVcenterServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnConnectorVerifyCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVmxImageUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVmxServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVpcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorCompartmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_local_cert"); ok || d.HasChange("_local_cert") {
		t, err := expandSystemSdnConnectorLocalCert(d, v, "_local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_local_cert"] = t
		}
	}

	if v, ok := d.GetOk("access_key"); ok || d.HasChange("access_key") {
		t, err := expandSystemSdnConnectorAccessKey(d, v, "access_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key"] = t
		}
	}

	if v, ok := d.GetOk("alt_resource_ip"); ok || d.HasChange("alt_resource_ip") {
		t, err := expandSystemSdnConnectorAltResourceIp(d, v, "alt_resource_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-resource-ip"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok || d.HasChange("api_key") {
		t, err := expandSystemSdnConnectorApiKey(d, v, "api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("azure_region"); ok || d.HasChange("azure_region") {
		t, err := expandSystemSdnConnectorAzureRegion(d, v, "azure_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-region"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok || d.HasChange("client_id") {
		t, err := expandSystemSdnConnectorClientId(d, v, "client_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok || d.HasChange("client_secret") {
		t, err := expandSystemSdnConnectorClientSecret(d, v, "client_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("compartment_list"); ok || d.HasChange("compartment_list") {
		t, err := expandSystemSdnConnectorCompartmentList(d, v, "compartment_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-list"] = t
		}
	}

	if v, ok := d.GetOk("compute_generation"); ok || d.HasChange("compute_generation") {
		t, err := expandSystemSdnConnectorComputeGeneration(d, v, "compute_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compute-generation"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemSdnConnectorDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_account_list"); ok || d.HasChange("external_account_list") {
		t, err := expandSystemSdnConnectorExternalAccountList(d, v, "external_account_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-account-list"] = t
		}
	}

	if v, ok := d.GetOk("external_ip"); ok || d.HasChange("external_ip") {
		t, err := expandSystemSdnConnectorExternalIp(d, v, "external_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip"] = t
		}
	}

	if v, ok := d.GetOk("forwarding_rule"); ok || d.HasChange("forwarding_rule") {
		t, err := expandSystemSdnConnectorForwardingRule(d, v, "forwarding_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarding-rule"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project"); ok || d.HasChange("gcp_project") {
		t, err := expandSystemSdnConnectorGcpProject(d, v, "gcp_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project_list"); ok || d.HasChange("gcp_project_list") {
		t, err := expandSystemSdnConnectorGcpProjectList(d, v, "gcp_project_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project-list"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandSystemSdnConnectorGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("ha_status"); ok || d.HasChange("ha_status") {
		t, err := expandSystemSdnConnectorHaStatus(d, v, "ha_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-status"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region"); ok || d.HasChange("ibm_region") {
		t, err := expandSystemSdnConnectorIbmRegion(d, v, "ibm_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region_gen1"); ok || d.HasChange("ibm_region_gen1") {
		t, err := expandSystemSdnConnectorIbmRegionGen1(d, v, "ibm_region_gen1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen1"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region_gen2"); ok || d.HasChange("ibm_region_gen2") {
		t, err := expandSystemSdnConnectorIbmRegionGen2(d, v, "ibm_region_gen2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen2"] = t
		}
	}

	if v, ok := d.GetOk("key_passwd"); ok || d.HasChange("key_passwd") {
		t, err := expandSystemSdnConnectorKeyPasswd(d, v, "key_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-passwd"] = t
		}
	}

	if v, ok := d.GetOk("login_endpoint"); ok || d.HasChange("login_endpoint") {
		t, err := expandSystemSdnConnectorLoginEndpoint(d, v, "login_endpoint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-endpoint"] = t
		}
	}

	if v, ok := d.GetOk("message_server_port"); ok || d.HasChange("message_server_port") {
		t, err := expandSystemSdnConnectorMessageServerPort(d, v, "message_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-server-port"] = t
		}
	}

	if v, ok := d.GetOk("microsoft_365"); ok || d.HasChange("microsoft_365") {
		t, err := expandSystemSdnConnectorMicrosoft365(d, v, "microsoft_365")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["microsoft-365"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdnConnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nic"); ok || d.HasChange("nic") {
		t, err := expandSystemSdnConnectorNic(d, v, "nic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nic"] = t
		}
	}

	if v, ok := d.GetOk("nsx_cert_fingerprint"); ok || d.HasChange("nsx_cert_fingerprint") {
		t, err := expandSystemSdnConnectorNsxCertFingerprint(d, v, "nsx_cert_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nsx-cert-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("oci_cert"); ok || d.HasChange("oci_cert") {
		t, err := expandSystemSdnConnectorOciCert(d, v, "oci_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-cert"] = t
		}
	}

	if v, ok := d.GetOk("oci_fingerprint"); ok || d.HasChange("oci_fingerprint") {
		t, err := expandSystemSdnConnectorOciFingerprint(d, v, "oci_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_list"); ok || d.HasChange("oci_region_list") {
		t, err := expandSystemSdnConnectorOciRegionList(d, v, "oci_region_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-list"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_type"); ok || d.HasChange("oci_region_type") {
		t, err := expandSystemSdnConnectorOciRegionType(d, v, "oci_region_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-type"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemSdnConnectorPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok || d.HasChange("private_key") {
		t, err := expandSystemSdnConnectorPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandSystemSdnConnectorProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok || d.HasChange("region") {
		t, err := expandSystemSdnConnectorRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("resource_group"); ok || d.HasChange("resource_group") {
		t, err := expandSystemSdnConnectorResourceGroup(d, v, "resource_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	}

	if v, ok := d.GetOk("resource_url"); ok || d.HasChange("resource_url") {
		t, err := expandSystemSdnConnectorResourceUrl(d, v, "resource_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-url"] = t
		}
	}

	if v, ok := d.GetOk("rest_interface"); ok || d.HasChange("rest_interface") {
		t, err := expandSystemSdnConnectorRestInterface(d, v, "rest_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-interface"] = t
		}
	}

	if v, ok := d.GetOk("rest_password"); ok || d.HasChange("rest_password") {
		t, err := expandSystemSdnConnectorRestPassword(d, v, "rest_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-password"] = t
		}
	}

	if v, ok := d.GetOk("rest_sport"); ok || d.HasChange("rest_sport") {
		t, err := expandSystemSdnConnectorRestSport(d, v, "rest_sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-sport"] = t
		}
	}

	if v, ok := d.GetOk("rest_ssl"); ok || d.HasChange("rest_ssl") {
		t, err := expandSystemSdnConnectorRestSsl(d, v, "rest_ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-ssl"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok || d.HasChange("route") {
		t, err := expandSystemSdnConnectorRoute(d, v, "route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("route_table"); ok || d.HasChange("route_table") {
		t, err := expandSystemSdnConnectorRouteTable(d, v, "route_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-table"] = t
		}
	}

	if v, ok := d.GetOk("secret_key"); ok || d.HasChange("secret_key") {
		t, err := expandSystemSdnConnectorSecretKey(d, v, "secret_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-key"] = t
		}
	}

	if v, ok := d.GetOk("secret_token"); ok || d.HasChange("secret_token") {
		t, err := expandSystemSdnConnectorSecretToken(d, v, "secret_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-token"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemSdnConnectorServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_ca_cert"); ok || d.HasChange("server_ca_cert") {
		t, err := expandSystemSdnConnectorServerCaCert(d, v, "server_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok || d.HasChange("server_cert") {
		t, err := expandSystemSdnConnectorServerCert(d, v, "server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystemSdnConnectorServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandSystemSdnConnectorServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("service_account"); ok || d.HasChange("service_account") {
		t, err := expandSystemSdnConnectorServiceAccount(d, v, "service_account")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSdnConnectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("subscription_id"); ok || d.HasChange("subscription_id") {
		t, err := expandSystemSdnConnectorSubscriptionId(d, v, "subscription_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscription-id"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok || d.HasChange("tenant_id") {
		t, err := expandSystemSdnConnectorTenantId(d, v, "tenant_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemSdnConnectorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandSystemSdnConnectorUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok || d.HasChange("use_metadata_iam") {
		t, err := expandSystemSdnConnectorUseMetadataIam(d, v, "use_metadata_iam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok || d.HasChange("user_id") {
		t, err := expandSystemSdnConnectorUserId(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemSdnConnectorUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_password"); ok || d.HasChange("vcenter_password") {
		t, err := expandSystemSdnConnectorVcenterPassword(d, v, "vcenter_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-password"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_server"); ok || d.HasChange("vcenter_server") {
		t, err := expandSystemSdnConnectorVcenterServer(d, v, "vcenter_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-server"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_username"); ok || d.HasChange("vcenter_username") {
		t, err := expandSystemSdnConnectorVcenterUsername(d, v, "vcenter_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-username"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSdnConnectorVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("verify_certificate"); ok || d.HasChange("verify_certificate") {
		t, err := expandSystemSdnConnectorVerifyCertificate(d, v, "verify_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-certificate"] = t
		}
	}

	if v, ok := d.GetOk("vmx_image_url"); ok || d.HasChange("vmx_image_url") {
		t, err := expandSystemSdnConnectorVmxImageUrl(d, v, "vmx_image_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vmx-image-url"] = t
		}
	}

	if v, ok := d.GetOk("vmx_service_name"); ok || d.HasChange("vmx_service_name") {
		t, err := expandSystemSdnConnectorVmxServiceName(d, v, "vmx_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vmx-service-name"] = t
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok || d.HasChange("vpc_id") {
		t, err := expandSystemSdnConnectorVpcId(d, v, "vpc_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-id"] = t
		}
	}

	if v, ok := d.GetOk("compartment_id"); ok || d.HasChange("compartment_id") {
		t, err := expandSystemSdnConnectorCompartmentId(d, v, "compartment_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	}

	if v, ok := d.GetOk("oci_region"); ok || d.HasChange("oci_region") {
		t, err := expandSystemSdnConnectorOciRegion(d, v, "oci_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region"] = t
		}
	}

	return &obj, nil
}
