// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure external resource.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemExternalResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemExternalResourceCreate,
		Read:   resourceSystemExternalResourceRead,
		Update: resourceSystemExternalResourceUpdate,
		Delete: resourceSystemExternalResourceDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"address_comment_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_data_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_name_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_array_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"refresh_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"proxy_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"proxy_username": &schema.Schema{
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

func resourceSystemExternalResourceCreate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemExternalResource resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemExternalResource(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemExternalResource(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemExternalResource resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemExternalResource(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemExternalResource resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemExternalResourceRead(d, m)
}

func resourceSystemExternalResourceUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemExternalResource resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemExternalResource(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemExternalResource resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemExternalResourceRead(d, m)
}

func resourceSystemExternalResourceDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteSystemExternalResource(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemExternalResource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemExternalResourceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSystemExternalResource(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemExternalResource resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemExternalResource(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemExternalResource resource from API: %v", err)
	}
	return nil
}

func flattenSystemExternalResourceAddressCommentField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceAddressDataField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceAddressNameField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemExternalResourceClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenSystemExternalResourceDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "SystemExternalResource-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenSystemExternalResourceDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "SystemExternalResource-DynamicMapping-SourceIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemExternalResourceDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemExternalResourceDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemExternalResourceDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSystemExternalResourceDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SystemExternalResourceDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemExternalResourceDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemExternalResourceInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceObjectArrayPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceRefreshRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceResource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemExternalResourceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceUpdateMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemExternalResource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("address_comment_field", flattenSystemExternalResourceAddressCommentField(o["address-comment-field"], d, "address_comment_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-comment-field"], "SystemExternalResource-AddressCommentField"); ok {
			if err = d.Set("address_comment_field", vv); err != nil {
				return fmt.Errorf("Error reading address_comment_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_comment_field: %v", err)
		}
	}

	if err = d.Set("address_data_field", flattenSystemExternalResourceAddressDataField(o["address-data-field"], d, "address_data_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-data-field"], "SystemExternalResource-AddressDataField"); ok {
			if err = d.Set("address_data_field", vv); err != nil {
				return fmt.Errorf("Error reading address_data_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_data_field: %v", err)
		}
	}

	if err = d.Set("address_name_field", flattenSystemExternalResourceAddressNameField(o["address-name-field"], d, "address_name_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["address-name-field"], "SystemExternalResource-AddressNameField"); ok {
			if err = d.Set("address_name_field", vv); err != nil {
				return fmt.Errorf("Error reading address_name_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address_name_field: %v", err)
		}
	}

	if err = d.Set("category", flattenSystemExternalResourceCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "SystemExternalResource-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenSystemExternalResourceClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "SystemExternalResource-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenSystemExternalResourceClientCertAuth(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-auth"], "SystemExternalResource-ClientCertAuth"); ok {
			if err = d.Set("client_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemExternalResourceComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemExternalResource-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenSystemExternalResourceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "SystemExternalResource-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenSystemExternalResourceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "SystemExternalResource-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenSystemExternalResourceInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemExternalResource-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemExternalResourceInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemExternalResource-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemExternalResourceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemExternalResource-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("namespace", flattenSystemExternalResourceNamespace(o["namespace"], d, "namespace")); err != nil {
		if vv, ok := fortiAPIPatch(o["namespace"], "SystemExternalResource-Namespace"); ok {
			if err = d.Set("namespace", vv); err != nil {
				return fmt.Errorf("Error reading namespace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading namespace: %v", err)
		}
	}

	if err = d.Set("object_array_path", flattenSystemExternalResourceObjectArrayPath(o["object-array-path"], d, "object_array_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["object-array-path"], "SystemExternalResource-ObjectArrayPath"); ok {
			if err = d.Set("object_array_path", vv); err != nil {
				return fmt.Errorf("Error reading object_array_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading object_array_path: %v", err)
		}
	}

	if err = d.Set("refresh_rate", flattenSystemExternalResourceRefreshRate(o["refresh-rate"], d, "refresh_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["refresh-rate"], "SystemExternalResource-RefreshRate"); ok {
			if err = d.Set("refresh_rate", vv); err != nil {
				return fmt.Errorf("Error reading refresh_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading refresh_rate: %v", err)
		}
	}

	if err = d.Set("resource", flattenSystemExternalResourceResource(o["resource"], d, "resource")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource"], "SystemExternalResource-Resource"); ok {
			if err = d.Set("resource", vv); err != nil {
				return fmt.Errorf("Error reading resource: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenSystemExternalResourceServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "SystemExternalResource-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemExternalResourceSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemExternalResource-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenSystemExternalResourceSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "SystemExternalResource-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemExternalResourceStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemExternalResource-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemExternalResourceType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemExternalResource-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("update_method", flattenSystemExternalResourceUpdateMethod(o["update-method"], d, "update_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-method"], "SystemExternalResource-UpdateMethod"); ok {
			if err = d.Set("update_method", vv); err != nil {
				return fmt.Errorf("Error reading update_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_method: %v", err)
		}
	}

	if err = d.Set("user_agent", flattenSystemExternalResourceUserAgent(o["user-agent"], d, "user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-agent"], "SystemExternalResource-UserAgent"); ok {
			if err = d.Set("user_agent", vv); err != nil {
				return fmt.Errorf("Error reading user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemExternalResourceUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemExternalResource-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("uuid", flattenSystemExternalResourceUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "SystemExternalResource-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystemExternalResourceVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystemExternalResource-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemExternalResourceProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "SystemExternalResource-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("proxy_port", flattenSystemExternalResourceProxyPort(o["proxy-port"], d, "proxy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-port"], "SystemExternalResource-ProxyPort"); ok {
			if err = d.Set("proxy_port", vv); err != nil {
				return fmt.Errorf("Error reading proxy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenSystemExternalResourceProxyUsername(o["proxy-username"], d, "proxy_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-username"], "SystemExternalResource-ProxyUsername"); ok {
			if err = d.Set("proxy_username", vv); err != nil {
				return fmt.Errorf("Error reading proxy_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	return nil
}

func flattenSystemExternalResourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemExternalResourceAddressCommentField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceAddressDataField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceAddressNameField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemExternalResourceClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemExternalResourceDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandSystemExternalResourceDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemExternalResourceDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemExternalResourceDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSystemExternalResourceDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemExternalResourceDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemExternalResourceInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceNamespace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceObjectArrayPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourcePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemExternalResourceRefreshRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceResource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemExternalResourceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceUpdateMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceUserAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemExternalResourceProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemExternalResource(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_comment_field"); ok || d.HasChange("address_comment_field") {
		t, err := expandSystemExternalResourceAddressCommentField(d, v, "address_comment_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-comment-field"] = t
		}
	}

	if v, ok := d.GetOk("address_data_field"); ok || d.HasChange("address_data_field") {
		t, err := expandSystemExternalResourceAddressDataField(d, v, "address_data_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-data-field"] = t
		}
	}

	if v, ok := d.GetOk("address_name_field"); ok || d.HasChange("address_name_field") {
		t, err := expandSystemExternalResourceAddressNameField(d, v, "address_name_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-name-field"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandSystemExternalResourceCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandSystemExternalResourceClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_auth"); ok || d.HasChange("client_cert_auth") {
		t, err := expandSystemExternalResourceClientCertAuth(d, v, "client_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemExternalResourceComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandSystemExternalResourceDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemExternalResourceInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemExternalResourceInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemExternalResourceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("namespace"); ok || d.HasChange("namespace") {
		t, err := expandSystemExternalResourceNamespace(d, v, "namespace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["namespace"] = t
		}
	}

	if v, ok := d.GetOk("object_array_path"); ok || d.HasChange("object_array_path") {
		t, err := expandSystemExternalResourceObjectArrayPath(d, v, "object_array_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object-array-path"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemExternalResourcePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("refresh_rate"); ok || d.HasChange("refresh_rate") {
		t, err := expandSystemExternalResourceRefreshRate(d, v, "refresh_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh-rate"] = t
		}
	}

	if v, ok := d.GetOk("resource"); ok || d.HasChange("resource") {
		t, err := expandSystemExternalResourceResource(d, v, "resource")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandSystemExternalResourceServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemExternalResourceSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandSystemExternalResourceSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemExternalResourceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemExternalResourceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_method"); ok || d.HasChange("update_method") {
		t, err := expandSystemExternalResourceUpdateMethod(d, v, "update_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-method"] = t
		}
	}

	if v, ok := d.GetOk("user_agent"); ok || d.HasChange("user_agent") {
		t, err := expandSystemExternalResourceUserAgent(d, v, "user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemExternalResourceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandSystemExternalResourceUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystemExternalResourceVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandSystemExternalResourceProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok || d.HasChange("proxy_password") {
		t, err := expandSystemExternalResourceProxyPassword(d, v, "proxy_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-password"] = t
		}
	}

	if v, ok := d.GetOk("proxy_port"); ok || d.HasChange("proxy_port") {
		t, err := expandSystemExternalResourceProxyPort(d, v, "proxy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-port"] = t
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok || d.HasChange("proxy_username") {
		t, err := expandSystemExternalResourceProxyUsername(d, v, "proxy_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-username"] = t
		}
	}

	return &obj, nil
}
