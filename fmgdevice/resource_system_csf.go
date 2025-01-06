// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfUpdate,
		Read:   resourceSystemCsfRead,
		Update: resourceSystemCsfUpdate,
		Delete: resourceSystemCsfDelete,

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
			"accept_auth_by_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"downstream_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"downstream_accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fabric_connector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accprofile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"configuration_write_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fabric_device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"device_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fabric_object_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_workers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"file_mgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"file_quota_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fixed_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"forticloud_account_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"legacy_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"saml_configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authorization_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"downstream_authorization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ha_members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"upstream_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upstream_port": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemCsfUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCsf(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCsf(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemCsf")

	return resourceSystemCsfRead(d, m)
}

func resourceSystemCsfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemCsf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCsf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfAcceptAuthByCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfAuthorizationRequestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfFabricConnector(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := i["accprofile"]; ok {
			v := flattenSystemCsfFabricConnectorAccprofile(i["accprofile"], d, pre_append)
			tmp["accprofile"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-Accprofile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := i["configuration-write-access"]; ok {
			v := flattenSystemCsfFabricConnectorConfigurationWriteAccess(i["configuration-write-access"], d, pre_append)
			tmp["configuration_write_access"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-ConfigurationWriteAccess")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemCsfFabricConnectorSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-Serial")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSystemCsfFabricConnectorVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SystemCsf-FabricConnector-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfFabricConnectorAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfFabricConnectorConfigurationWriteAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfFabricDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := i["device-ip"]; ok {
			v := flattenSystemCsfFabricDeviceDeviceIp(i["device-ip"], d, pre_append)
			tmp["device_ip"] = fortiAPISubPartPatch(v, "SystemCsf-FabricDevice-DeviceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
		if _, ok := i["https-port"]; ok {
			v := flattenSystemCsfFabricDeviceHttpsPort(i["https-port"], d, pre_append)
			tmp["https_port"] = fortiAPISubPartPatch(v, "SystemCsf-FabricDevice-HttpsPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemCsfFabricDeviceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemCsf-FabricDevice-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfFabricDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricObjectUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricWorkers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileMgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFileQuotaWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfForticloudAccountEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfLegacyAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfLogUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSamlConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenSystemCsfTrustedListAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := i["authorization-type"]; ok {
			v := flattenSystemCsfTrustedListAuthorizationType(i["authorization-type"], d, pre_append)
			tmp["authorization_type"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-AuthorizationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			v := flattenSystemCsfTrustedListCertificate(i["certificate"], d, pre_append)
			tmp["certificate"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Certificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := i["downstream-authorization"]; ok {
			v := flattenSystemCsfTrustedListDownstreamAuthorization(i["downstream-authorization"], d, pre_append)
			tmp["downstream_authorization"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-DownstreamAuthorization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := i["ha-members"]; ok {
			v := flattenSystemCsfTrustedListHaMembers(i["ha-members"], d, pre_append)
			tmp["ha_members"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-HaMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenSystemCsfTrustedListIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemCsfTrustedListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			v := flattenSystemCsfTrustedListSerial(i["serial"], d, pre_append)
			tmp["serial"] = fortiAPISubPartPatch(v, "SystemCsf-TrustedList-Serial")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemCsfTrustedListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAuthorizationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListDownstreamAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListHaMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfTrustedListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemCsfUpstreamInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("accept_auth_by_cert", flattenSystemCsfAcceptAuthByCert(o["accept-auth-by-cert"], d, "accept_auth_by_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-auth-by-cert"], "SystemCsf-AcceptAuthByCert"); ok {
			if err = d.Set("accept_auth_by_cert", vv); err != nil {
				return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
		}
	}

	if err = d.Set("authorization_request_type", flattenSystemCsfAuthorizationRequestType(o["authorization-request-type"], d, "authorization_request_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization-request-type"], "SystemCsf-AuthorizationRequestType"); ok {
			if err = d.Set("authorization_request_type", vv); err != nil {
				return fmt.Errorf("Error reading authorization_request_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization_request_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystemCsfCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystemCsf-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("configuration_sync", flattenSystemCsfConfigurationSync(o["configuration-sync"], d, "configuration_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration-sync"], "SystemCsf-ConfigurationSync"); ok {
			if err = d.Set("configuration_sync", vv); err != nil {
				return fmt.Errorf("Error reading configuration_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration_sync: %v", err)
		}
	}

	if err = d.Set("downstream_access", flattenSystemCsfDownstreamAccess(o["downstream-access"], d, "downstream_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["downstream-access"], "SystemCsf-DownstreamAccess"); ok {
			if err = d.Set("downstream_access", vv); err != nil {
				return fmt.Errorf("Error reading downstream_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downstream_access: %v", err)
		}
	}

	if err = d.Set("downstream_accprofile", flattenSystemCsfDownstreamAccprofile(o["downstream-accprofile"], d, "downstream_accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["downstream-accprofile"], "SystemCsf-DownstreamAccprofile"); ok {
			if err = d.Set("downstream_accprofile", vv); err != nil {
				return fmt.Errorf("Error reading downstream_accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downstream_accprofile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector")); err != nil {
			if vv, ok := fortiAPIPatch(o["fabric-connector"], "SystemCsf-FabricConnector"); ok {
				if err = d.Set("fabric_connector", vv); err != nil {
					return fmt.Errorf("Error reading fabric_connector: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fabric_connector: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_connector"); ok {
			if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector")); err != nil {
				if vv, ok := fortiAPIPatch(o["fabric-connector"], "SystemCsf-FabricConnector"); ok {
					if err = d.Set("fabric_connector", vv); err != nil {
						return fmt.Errorf("Error reading fabric_connector: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fabric_connector: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device")); err != nil {
			if vv, ok := fortiAPIPatch(o["fabric-device"], "SystemCsf-FabricDevice"); ok {
				if err = d.Set("fabric_device", vv); err != nil {
					return fmt.Errorf("Error reading fabric_device: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fabric_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_device"); ok {
			if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device")); err != nil {
				if vv, ok := fortiAPIPatch(o["fabric-device"], "SystemCsf-FabricDevice"); ok {
					if err = d.Set("fabric_device", vv); err != nil {
						return fmt.Errorf("Error reading fabric_device: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fabric_device: %v", err)
				}
			}
		}
	}

	if err = d.Set("fabric_object_unification", flattenSystemCsfFabricObjectUnification(o["fabric-object-unification"], d, "fabric_object_unification")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-unification"], "SystemCsf-FabricObjectUnification"); ok {
			if err = d.Set("fabric_object_unification", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_unification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_unification: %v", err)
		}
	}

	if err = d.Set("fabric_workers", flattenSystemCsfFabricWorkers(o["fabric-workers"], d, "fabric_workers")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-workers"], "SystemCsf-FabricWorkers"); ok {
			if err = d.Set("fabric_workers", vv); err != nil {
				return fmt.Errorf("Error reading fabric_workers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_workers: %v", err)
		}
	}

	if err = d.Set("file_mgmt", flattenSystemCsfFileMgmt(o["file-mgmt"], d, "file_mgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-mgmt"], "SystemCsf-FileMgmt"); ok {
			if err = d.Set("file_mgmt", vv); err != nil {
				return fmt.Errorf("Error reading file_mgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_mgmt: %v", err)
		}
	}

	if err = d.Set("file_quota", flattenSystemCsfFileQuota(o["file-quota"], d, "file_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota"], "SystemCsf-FileQuota"); ok {
			if err = d.Set("file_quota", vv); err != nil {
				return fmt.Errorf("Error reading file_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("file_quota_warning", flattenSystemCsfFileQuotaWarning(o["file-quota-warning"], d, "file_quota_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-quota-warning"], "SystemCsf-FileQuotaWarning"); ok {
			if err = d.Set("file_quota_warning", vv); err != nil {
				return fmt.Errorf("Error reading file_quota_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_quota_warning: %v", err)
		}
	}

	if err = d.Set("forticloud_account_enforcement", flattenSystemCsfForticloudAccountEnforcement(o["forticloud-account-enforcement"], d, "forticloud_account_enforcement")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticloud-account-enforcement"], "SystemCsf-ForticloudAccountEnforcement"); ok {
			if err = d.Set("forticloud_account_enforcement", vv); err != nil {
				return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemCsfGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "SystemCsf-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("legacy_authentication", flattenSystemCsfLegacyAuthentication(o["legacy-authentication"], d, "legacy_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["legacy-authentication"], "SystemCsf-LegacyAuthentication"); ok {
			if err = d.Set("legacy_authentication", vv); err != nil {
				return fmt.Errorf("Error reading legacy_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading legacy_authentication: %v", err)
		}
	}

	if err = d.Set("log_unification", flattenSystemCsfLogUnification(o["log-unification"], d, "log_unification")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-unification"], "SystemCsf-LogUnification"); ok {
			if err = d.Set("log_unification", vv); err != nil {
				return fmt.Errorf("Error reading log_unification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_unification: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemCsfManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-ip"], "SystemCsf-ManagementIp"); ok {
			if err = d.Set("management_ip", vv); err != nil {
				return fmt.Errorf("Error reading management_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemCsfManagementPort(o["management-port"], d, "management_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["management-port"], "SystemCsf-ManagementPort"); ok {
			if err = d.Set("management_port", vv); err != nil {
				return fmt.Errorf("Error reading management_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("saml_configuration_sync", flattenSystemCsfSamlConfigurationSync(o["saml-configuration-sync"], d, "saml_configuration_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-configuration-sync"], "SystemCsf-SamlConfigurationSync"); ok {
			if err = d.Set("saml_configuration_sync", vv); err != nil {
				return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemCsfSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemCsf-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemCsfStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemCsf-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemCsf-TrustedList"); ok {
				if err = d.Set("trusted_list", vv); err != nil {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading trusted_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_list"); ok {
			if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["trusted-list"], "SystemCsf-TrustedList"); ok {
					if err = d.Set("trusted_list", vv); err != nil {
						return fmt.Errorf("Error reading trusted_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("uid", flattenSystemCsfUid(o["uid"], d, "uid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uid"], "SystemCsf-Uid"); ok {
			if err = d.Set("uid", vv); err != nil {
				return fmt.Errorf("Error reading uid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uid: %v", err)
		}
	}

	if err = d.Set("upstream", flattenSystemCsfUpstream(o["upstream"], d, "upstream")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream"], "SystemCsf-Upstream"); ok {
			if err = d.Set("upstream", vv); err != nil {
				return fmt.Errorf("Error reading upstream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream: %v", err)
		}
	}

	if err = d.Set("upstream_ip", flattenSystemCsfUpstreamIp(o["upstream-ip"], d, "upstream_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-ip"], "SystemCsf-UpstreamIp"); ok {
			if err = d.Set("upstream_ip", vv); err != nil {
				return fmt.Errorf("Error reading upstream_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_ip: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenSystemCsfUpstreamInterface(o["upstream-interface"], d, "upstream_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-interface"], "SystemCsf-UpstreamInterface"); ok {
			if err = d.Set("upstream_interface", vv); err != nil {
				return fmt.Errorf("Error reading upstream_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
		}
	}

	if err = d.Set("upstream_interface_select_method", flattenSystemCsfUpstreamInterfaceSelectMethod(o["upstream-interface-select-method"], d, "upstream_interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-interface-select-method"], "SystemCsf-UpstreamInterfaceSelectMethod"); ok {
			if err = d.Set("upstream_interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading upstream_interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_interface_select_method: %v", err)
		}
	}

	if err = d.Set("upstream_port", flattenSystemCsfUpstreamPort(o["upstream-port"], d, "upstream_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-port"], "SystemCsf-UpstreamPort"); ok {
			if err = d.Set("upstream_port", vv); err != nil {
				return fmt.Errorf("Error reading upstream_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_port: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfAcceptAuthByCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfAuthorizationRequestType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfConfigurationSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accprofile"], _ = expandSystemCsfFabricConnectorAccprofile(d, i["accprofile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["configuration-write-access"], _ = expandSystemCsfFabricConnectorConfigurationWriteAccess(d, i["configuration_write_access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemCsfFabricConnectorSerial(d, i["serial"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSystemCsfFabricConnectorVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricConnectorAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricConnectorConfigurationWriteAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_token"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-token"], _ = expandSystemCsfFabricDeviceAccessToken(d, i["access_token"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-ip"], _ = expandSystemCsfFabricDeviceDeviceIp(d, i["device_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["https-port"], _ = expandSystemCsfFabricDeviceHttpsPort(d, i["https_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemCsfFabricDeviceName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricDeviceAccessToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfFabricDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricObjectUnification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricWorkers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileMgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuotaWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFixedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfForticloudAccountEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfLegacyAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfLogUnification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSamlConfigurationSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandSystemCsfTrustedListAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authorization-type"], _ = expandSystemCsfTrustedListAuthorizationType(d, i["authorization_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["certificate"], _ = expandSystemCsfTrustedListCertificate(d, i["certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["downstream-authorization"], _ = expandSystemCsfTrustedListDownstreamAuthorization(d, i["downstream_authorization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-members"], _ = expandSystemCsfTrustedListHaMembers(d, i["ha_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandSystemCsfTrustedListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemCsfTrustedListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["serial"], _ = expandSystemCsfTrustedListSerial(d, i["serial"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemCsfTrustedListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAuthorizationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListDownstreamAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListHaMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfTrustedListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemCsfUpstreamInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_auth_by_cert"); ok || d.HasChange("accept_auth_by_cert") {
		t, err := expandSystemCsfAcceptAuthByCert(d, v, "accept_auth_by_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-auth-by-cert"] = t
		}
	}

	if v, ok := d.GetOk("authorization_request_type"); ok || d.HasChange("authorization_request_type") {
		t, err := expandSystemCsfAuthorizationRequestType(d, v, "authorization_request_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-request-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandSystemCsfCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("configuration_sync"); ok || d.HasChange("configuration_sync") {
		t, err := expandSystemCsfConfigurationSync(d, v, "configuration_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-sync"] = t
		}
	}

	if v, ok := d.GetOk("downstream_access"); ok || d.HasChange("downstream_access") {
		t, err := expandSystemCsfDownstreamAccess(d, v, "downstream_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downstream-access"] = t
		}
	}

	if v, ok := d.GetOk("downstream_accprofile"); ok || d.HasChange("downstream_accprofile") {
		t, err := expandSystemCsfDownstreamAccprofile(d, v, "downstream_accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downstream-accprofile"] = t
		}
	}

	if v, ok := d.GetOk("fabric_connector"); ok || d.HasChange("fabric_connector") {
		t, err := expandSystemCsfFabricConnector(d, v, "fabric_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-connector"] = t
		}
	}

	if v, ok := d.GetOk("fabric_device"); ok || d.HasChange("fabric_device") {
		t, err := expandSystemCsfFabricDevice(d, v, "fabric_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-device"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_unification"); ok || d.HasChange("fabric_object_unification") {
		t, err := expandSystemCsfFabricObjectUnification(d, v, "fabric_object_unification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-unification"] = t
		}
	}

	if v, ok := d.GetOk("fabric_workers"); ok || d.HasChange("fabric_workers") {
		t, err := expandSystemCsfFabricWorkers(d, v, "fabric_workers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-workers"] = t
		}
	}

	if v, ok := d.GetOk("file_mgmt"); ok || d.HasChange("file_mgmt") {
		t, err := expandSystemCsfFileMgmt(d, v, "file_mgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-mgmt"] = t
		}
	}

	if v, ok := d.GetOk("file_quota"); ok || d.HasChange("file_quota") {
		t, err := expandSystemCsfFileQuota(d, v, "file_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota"] = t
		}
	}

	if v, ok := d.GetOk("file_quota_warning"); ok || d.HasChange("file_quota_warning") {
		t, err := expandSystemCsfFileQuotaWarning(d, v, "file_quota_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-quota-warning"] = t
		}
	}

	if v, ok := d.GetOk("fixed_key"); ok || d.HasChange("fixed_key") {
		t, err := expandSystemCsfFixedKey(d, v, "fixed_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixed-key"] = t
		}
	}

	if v, ok := d.GetOk("forticloud_account_enforcement"); ok || d.HasChange("forticloud_account_enforcement") {
		t, err := expandSystemCsfForticloudAccountEnforcement(d, v, "forticloud_account_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticloud-account-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandSystemCsfGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("group_password"); ok || d.HasChange("group_password") {
		t, err := expandSystemCsfGroupPassword(d, v, "group_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-password"] = t
		}
	}

	if v, ok := d.GetOk("legacy_authentication"); ok || d.HasChange("legacy_authentication") {
		t, err := expandSystemCsfLegacyAuthentication(d, v, "legacy_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["legacy-authentication"] = t
		}
	}

	if v, ok := d.GetOk("log_unification"); ok || d.HasChange("log_unification") {
		t, err := expandSystemCsfLogUnification(d, v, "log_unification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-unification"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok || d.HasChange("management_ip") {
		t, err := expandSystemCsfManagementIp(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("management_port"); ok || d.HasChange("management_port") {
		t, err := expandSystemCsfManagementPort(d, v, "management_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port"] = t
		}
	}

	if v, ok := d.GetOk("saml_configuration_sync"); ok || d.HasChange("saml_configuration_sync") {
		t, err := expandSystemCsfSamlConfigurationSync(d, v, "saml_configuration_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-configuration-sync"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemCsfSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemCsfStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_list"); ok || d.HasChange("trusted_list") {
		t, err := expandSystemCsfTrustedList(d, v, "trusted_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-list"] = t
		}
	}

	if v, ok := d.GetOk("uid"); ok || d.HasChange("uid") {
		t, err := expandSystemCsfUid(d, v, "uid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uid"] = t
		}
	}

	if v, ok := d.GetOk("upstream"); ok || d.HasChange("upstream") {
		t, err := expandSystemCsfUpstream(d, v, "upstream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream"] = t
		}
	}

	if v, ok := d.GetOk("upstream_ip"); ok || d.HasChange("upstream_ip") {
		t, err := expandSystemCsfUpstreamIp(d, v, "upstream_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-ip"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok || d.HasChange("upstream_interface") {
		t, err := expandSystemCsfUpstreamInterface(d, v, "upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface_select_method"); ok || d.HasChange("upstream_interface_select_method") {
		t, err := expandSystemCsfUpstreamInterfaceSelectMethod(d, v, "upstream_interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("upstream_port"); ok || d.HasChange("upstream_port") {
		t, err := expandSystemCsfUpstreamPort(d, v, "upstream_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-port"] = t
		}
	}

	return &obj, nil
}
