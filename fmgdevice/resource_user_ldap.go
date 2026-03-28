// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure LDAP server entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLdapCreate,
		Read:   resourceUserLdapRead,
		Update: resourceUserLdapUpdate,
		Delete: resourceUserLdapDelete,

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
			"account_key_cert_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antiphish": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"cnid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": &schema.Schema{
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
						"account_key_cert_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_processing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_upn_san": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"antiphish": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ca_cert": &schema.Schema{
							Type:     schema.TypeString,
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
						},
						"cnid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_member_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_object_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_object_search_base": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_search_base": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"member_attr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obtain_user_info": &schema.Schema{
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
						"password_attr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_expiry_warning": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_renewal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retrieve_protection_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_identity_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssl_max_proto_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_min_proto_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_notification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_info_exchange_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"validate_server_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"group_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_member_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_object_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_base": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"member_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"obtain_user_info": &schema.Schema{
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
			"password_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password_expiry_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"search_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_exchange_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_key_upn_san": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_max_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validate_server_certificate": &schema.Schema{
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

func resourceUserLdapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserLdap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserLdap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserLdap resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserLdap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserLdap resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserLdap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserLdap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserLdap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource from API: %v", err)
	}
	return nil
}

func flattenUserLdapAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAntiphish(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserLdapDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := i["account-key-cert-field"]; ok {
			v := flattenUserLdapDynamicMappingAccountKeyCertField(i["account-key-cert-field"], d, pre_append)
			tmp["account_key_cert_field"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-AccountKeyCertField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_filter"
		if _, ok := i["account-key-filter"]; ok {
			v := flattenUserLdapDynamicMappingAccountKeyFilter(i["account-key-filter"], d, pre_append)
			tmp["account_key_filter"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-AccountKeyFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_name"
		if _, ok := i["account-key-name"]; ok {
			v := flattenUserLdapDynamicMappingAccountKeyName(i["account-key-name"], d, pre_append)
			tmp["account_key_name"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-AccountKeyName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := i["account-key-processing"]; ok {
			v := flattenUserLdapDynamicMappingAccountKeyProcessing(i["account-key-processing"], d, pre_append)
			tmp["account_key_processing"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-AccountKeyProcessing")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_upn_san"
		if _, ok := i["account-key-upn-san"]; ok {
			v := flattenUserLdapDynamicMappingAccountKeyUpnSan(i["account-key-upn-san"], d, pre_append)
			tmp["account_key_upn_san"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-AccountKeyUpnSan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish"
		if _, ok := i["antiphish"]; ok {
			v := flattenUserLdapDynamicMappingAntiphish(i["antiphish"], d, pre_append)
			tmp["antiphish"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Antiphish")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := i["ca-cert"]; ok {
			v := flattenUserLdapDynamicMappingCaCert(i["ca-cert"], d, pre_append)
			tmp["ca_cert"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-CaCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenUserLdapDynamicMappingClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert_auth"
		if _, ok := i["client-cert-auth"]; ok {
			v := flattenUserLdapDynamicMappingClientCertAuth(i["client-cert-auth"], d, pre_append)
			tmp["client_cert_auth"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-ClientCertAuth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cnid"
		if _, ok := i["cnid"]; ok {
			v := flattenUserLdapDynamicMappingCnid(i["cnid"], d, pre_append)
			tmp["cnid"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Cnid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dn"
		if _, ok := i["dn"]; ok {
			v := flattenUserLdapDynamicMappingDn(i["dn"], d, pre_append)
			tmp["dn"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Dn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenUserLdapDynamicMappingFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenUserLdapDynamicMappingGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_filter"
		if _, ok := i["group-filter"]; ok {
			v := flattenUserLdapDynamicMappingGroupFilter(i["group-filter"], d, pre_append)
			tmp["group_filter"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-GroupFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_member_check"
		if _, ok := i["group-member-check"]; ok {
			v := flattenUserLdapDynamicMappingGroupMemberCheck(i["group-member-check"], d, pre_append)
			tmp["group_member_check"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-GroupMemberCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_filter"
		if _, ok := i["group-object-filter"]; ok {
			v := flattenUserLdapDynamicMappingGroupObjectFilter(i["group-object-filter"], d, pre_append)
			tmp["group_object_filter"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-GroupObjectFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_search_base"
		if _, ok := i["group-object-search-base"]; ok {
			v := flattenUserLdapDynamicMappingGroupObjectSearchBase(i["group-object-search-base"], d, pre_append)
			tmp["group_object_search_base"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-GroupObjectSearchBase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_search_base"
		if _, ok := i["group-search-base"]; ok {
			v := flattenUserLdapDynamicMappingGroupSearchBase(i["group-search-base"], d, pre_append)
			tmp["group_search_base"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-GroupSearchBase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserLdapDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserLdapDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			v := flattenUserLdapDynamicMappingMaxConnections(i["max-connections"], d, pre_append)
			tmp["max_connections"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-MaxConnections")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_attr"
		if _, ok := i["member-attr"]; ok {
			v := flattenUserLdapDynamicMappingMemberAttr(i["member-attr"], d, pre_append)
			tmp["member_attr"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-MemberAttr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obtain_user_info"
		if _, ok := i["obtain-user-info"]; ok {
			v := flattenUserLdapDynamicMappingObtainUserInfo(i["obtain-user-info"], d, pre_append)
			tmp["obtain_user_info"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-ObtainUserInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_attr"
		if _, ok := i["password-attr"]; ok {
			v := flattenUserLdapDynamicMappingPasswordAttr(i["password-attr"], d, pre_append)
			tmp["password_attr"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-PasswordAttr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_expiry_warning"
		if _, ok := i["password-expiry-warning"]; ok {
			v := flattenUserLdapDynamicMappingPasswordExpiryWarning(i["password-expiry-warning"], d, pre_append)
			tmp["password_expiry_warning"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-PasswordExpiryWarning")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := i["password-renewal"]; ok {
			v := flattenUserLdapDynamicMappingPasswordRenewal(i["password-renewal"], d, pre_append)
			tmp["password_renewal"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-PasswordRenewal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserLdapDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retrieve_protection_profile"
		if _, ok := i["retrieve-protection-profile"]; ok {
			v := flattenUserLdapDynamicMappingRetrieveProtectionProfile(i["retrieve-protection-profile"], d, pre_append)
			tmp["retrieve_protection_profile"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-RetrieveProtectionProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_type"
		if _, ok := i["search-type"]; ok {
			v := flattenUserLdapDynamicMappingSearchType(i["search-type"], d, pre_append)
			tmp["search_type"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SearchType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenUserLdapDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secure"
		if _, ok := i["secure"]; ok {
			v := flattenUserLdapDynamicMappingSecure(i["secure"], d, pre_append)
			tmp["secure"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Secure")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserLdapDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := i["server-identity-check"]; ok {
			v := flattenUserLdapDynamicMappingServerIdentityCheck(i["server-identity-check"], d, pre_append)
			tmp["server_identity_check"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-ServerIdentityCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserLdapDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
		if _, ok := i["source-ip-interface"]; ok {
			v := flattenUserLdapDynamicMappingSourceIpInterface(i["source-ip-interface"], d, pre_append)
			tmp["source_ip_interface"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SourceIpInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenUserLdapDynamicMappingSourcePort(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SourcePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_proto_version"
		if _, ok := i["ssl-max-proto-version"]; ok {
			v := flattenUserLdapDynamicMappingSslMaxProtoVersion(i["ssl-max-proto-version"], d, pre_append)
			tmp["ssl_max_proto_version"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SslMaxProtoVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_proto_version"
		if _, ok := i["ssl-min-proto-version"]; ok {
			v := flattenUserLdapDynamicMappingSslMinProtoVersion(i["ssl-min-proto-version"], d, pre_append)
			tmp["ssl_min_proto_version"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-SslMinProtoVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := i["status-ttl"]; ok {
			v := flattenUserLdapDynamicMappingStatusTtl(i["status-ttl"], d, pre_append)
			tmp["status_ttl"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-StatusTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenUserLdapDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-TertiaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := i["two-factor"]; ok {
			v := flattenUserLdapDynamicMappingTwoFactor(i["two-factor"], d, pre_append)
			tmp["two_factor"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-TwoFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := i["two-factor-authentication"]; ok {
			v := flattenUserLdapDynamicMappingTwoFactorAuthentication(i["two-factor-authentication"], d, pre_append)
			tmp["two_factor_authentication"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-TwoFactorAuthentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_filter"
		if _, ok := i["two-factor-filter"]; ok {
			v := flattenUserLdapDynamicMappingTwoFactorFilter(i["two-factor-filter"], d, pre_append)
			tmp["two_factor_filter"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-TwoFactorFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := i["two-factor-notification"]; ok {
			v := flattenUserLdapDynamicMappingTwoFactorNotification(i["two-factor-notification"], d, pre_append)
			tmp["two_factor_notification"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-TwoFactorNotification")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenUserLdapDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_exchange_server"
		if _, ok := i["user-info-exchange-server"]; ok {
			v := flattenUserLdapDynamicMappingUserInfoExchangeServer(i["user-info-exchange-server"], d, pre_append)
			tmp["user_info_exchange_server"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-UserInfoExchangeServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username"
		if _, ok := i["username"]; ok {
			v := flattenUserLdapDynamicMappingUsername(i["username"], d, pre_append)
			tmp["username"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-Username")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "validate_server_certificate"
		if _, ok := i["validate-server-certificate"]; ok {
			v := flattenUserLdapDynamicMappingValidateServerCertificate(i["validate-server-certificate"], d, pre_append)
			tmp["validate_server_certificate"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-ValidateServerCertificate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserLdapDynamicMappingVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserLdap-DynamicMapping-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserLdapDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserLdapDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserLdapDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserLdapDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserLdapDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserLdapDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAccountKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAccountKeyUpnSan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingAntiphish(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapDynamicMappingClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroupFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroupObjectSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingGroupSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingMemberAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingObtainUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingPasswordAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingRetrieveProtectionProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSearchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapDynamicMappingSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSslMaxProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingTwoFactorFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapDynamicMappingUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingValidateServerCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDynamicMappingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapObtainUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPasswordAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSearchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTwoFactorFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAccountKeyUpnSan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSslMaxProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapValidateServerCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserLdap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("account_key_cert_field", flattenUserLdapAccountKeyCertField(o["account-key-cert-field"], d, "account_key_cert_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-cert-field"], "UserLdap-AccountKeyCertField"); ok {
			if err = d.Set("account_key_cert_field", vv); err != nil {
				return fmt.Errorf("Error reading account_key_cert_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenUserLdapAccountKeyFilter(o["account-key-filter"], d, "account_key_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-filter"], "UserLdap-AccountKeyFilter"); ok {
			if err = d.Set("account_key_filter", vv); err != nil {
				return fmt.Errorf("Error reading account_key_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenUserLdapAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-processing"], "UserLdap-AccountKeyProcessing"); ok {
			if err = d.Set("account_key_processing", vv); err != nil {
				return fmt.Errorf("Error reading account_key_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("antiphish", flattenUserLdapAntiphish(o["antiphish"], d, "antiphish")); err != nil {
		if vv, ok := fortiAPIPatch(o["antiphish"], "UserLdap-Antiphish"); ok {
			if err = d.Set("antiphish", vv); err != nil {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading antiphish: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserLdapCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "UserLdap-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenUserLdapClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "UserLdap-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenUserLdapClientCertAuth(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert-auth"], "UserLdap-ClientCertAuth"); ok {
			if err = d.Set("client_cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading client_cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("cnid", flattenUserLdapCnid(o["cnid"], d, "cnid")); err != nil {
		if vv, ok := fortiAPIPatch(o["cnid"], "UserLdap-Cnid"); ok {
			if err = d.Set("cnid", vv); err != nil {
				return fmt.Errorf("Error reading cnid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenUserLdapDn(o["dn"], d, "dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["dn"], "UserLdap-Dn"); ok {
			if err = d.Set("dn", vv); err != nil {
				return fmt.Errorf("Error reading dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserLdapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserLdap-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserLdapDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserLdap-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_filter", flattenUserLdapGroupFilter(o["group-filter"], d, "group_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-filter"], "UserLdap-GroupFilter"); ok {
			if err = d.Set("group_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_filter: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-member-check"], "UserLdap-GroupMemberCheck"); ok {
			if err = d.Set("group_member_check", vv); err != nil {
				return fmt.Errorf("Error reading group_member_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-object-filter"], "UserLdap-GroupObjectFilter"); ok {
			if err = d.Set("group_object_filter", vv); err != nil {
				return fmt.Errorf("Error reading group_object_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenUserLdapGroupSearchBase(o["group-search-base"], d, "group_search_base")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-search-base"], "UserLdap-GroupSearchBase"); ok {
			if err = d.Set("group_search_base", vv); err != nil {
				return fmt.Errorf("Error reading group_search_base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserLdapInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserLdap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserLdapInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserLdap-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenUserLdapMemberAttr(o["member-attr"], d, "member_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-attr"], "UserLdap-MemberAttr"); ok {
			if err = d.Set("member_attr", vv); err != nil {
				return fmt.Errorf("Error reading member_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("name", flattenUserLdapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserLdap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obtain_user_info", flattenUserLdapObtainUserInfo(o["obtain-user-info"], d, "obtain_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["obtain-user-info"], "UserLdap-ObtainUserInfo"); ok {
			if err = d.Set("obtain_user_info", vv); err != nil {
				return fmt.Errorf("Error reading obtain_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obtain_user_info: %v", err)
		}
	}

	if err = d.Set("password_attr", flattenUserLdapPasswordAttr(o["password-attr"], d, "password_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-attr"], "UserLdap-PasswordAttr"); ok {
			if err = d.Set("password_attr", vv); err != nil {
				return fmt.Errorf("Error reading password_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_attr: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-expiry-warning"], "UserLdap-PasswordExpiryWarning"); ok {
			if err = d.Set("password_expiry_warning", vv); err != nil {
				return fmt.Errorf("Error reading password_expiry_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "UserLdap-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("port", flattenUserLdapPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserLdap-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("search_type", flattenUserLdapSearchType(o["search-type"], d, "search_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-type"], "UserLdap-SearchType"); ok {
			if err = d.Set("search_type", vv); err != nil {
				return fmt.Errorf("Error reading search_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_type: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserLdapSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "UserLdap-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secure", flattenUserLdapSecure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "UserLdap-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("server", flattenUserLdapServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "UserLdap-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserLdapServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "UserLdap-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserLdapSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "UserLdap-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenUserLdapSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "UserLdap-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_port", flattenUserLdapSourcePort(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "UserLdap-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserLdapSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "UserLdap-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenUserLdapStatusTtl(o["status-ttl"], d, "status_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ttl"], "UserLdap-StatusTtl"); ok {
			if err = d.Set("status_ttl", vv); err != nil {
				return fmt.Errorf("Error reading status_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "UserLdap-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserLdapTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "UserLdap-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenUserLdapTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "UserLdap-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_filter", flattenUserLdapTwoFactorFilter(o["two-factor-filter"], d, "two_factor_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-filter"], "UserLdap-TwoFactorFilter"); ok {
			if err = d.Set("two_factor_filter", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_filter: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenUserLdapTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "UserLdap-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLdapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserLdap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_exchange_server", flattenUserLdapUserInfoExchangeServer(o["user-info-exchange-server"], d, "user_info_exchange_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-exchange-server"], "UserLdap-UserInfoExchangeServer"); ok {
			if err = d.Set("user_info_exchange_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
		}
	}

	if err = d.Set("username", flattenUserLdapUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "UserLdap-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("account_key_upn_san", flattenUserLdapAccountKeyUpnSan(o["account-key-upn-san"], d, "account_key_upn_san")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-upn-san"], "UserLdap-AccountKeyUpnSan"); ok {
			if err = d.Set("account_key_upn_san", vv); err != nil {
				return fmt.Errorf("Error reading account_key_upn_san: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_upn_san: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserLdapVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "UserLdap-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenUserLdapMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "UserLdap-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("ssl_max_proto_version", flattenUserLdapSslMaxProtoVersion(o["ssl-max-proto-version"], d, "ssl_max_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-max-proto-version"], "UserLdap-SslMaxProtoVersion"); ok {
			if err = d.Set("ssl_max_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_max_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_max_proto_version: %v", err)
		}
	}

	if err = d.Set("validate_server_certificate", flattenUserLdapValidateServerCertificate(o["validate-server-certificate"], d, "validate_server_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["validate-server-certificate"], "UserLdap-ValidateServerCertificate"); ok {
			if err = d.Set("validate_server_certificate", vv); err != nil {
				return fmt.Errorf("Error reading validate_server_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading validate_server_certificate: %v", err)
		}
	}

	return nil
}

func flattenUserLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserLdapAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandUserLdapDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-cert-field"], _ = expandUserLdapDynamicMappingAccountKeyCertField(d, i["account_key_cert_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-filter"], _ = expandUserLdapDynamicMappingAccountKeyFilter(d, i["account_key_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-name"], _ = expandUserLdapDynamicMappingAccountKeyName(d, i["account_key_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-processing"], _ = expandUserLdapDynamicMappingAccountKeyProcessing(d, i["account_key_processing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_upn_san"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-upn-san"], _ = expandUserLdapDynamicMappingAccountKeyUpnSan(d, i["account_key_upn_san"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antiphish"], _ = expandUserLdapDynamicMappingAntiphish(d, i["antiphish"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ca-cert"], _ = expandUserLdapDynamicMappingCaCert(d, i["ca_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandUserLdapDynamicMappingClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert_auth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert-auth"], _ = expandUserLdapDynamicMappingClientCertAuth(d, i["client_cert_auth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cnid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cnid"], _ = expandUserLdapDynamicMappingCnid(d, i["cnid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dn"], _ = expandUserLdapDynamicMappingDn(d, i["dn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandUserLdapDynamicMappingFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandUserLdapDynamicMappingGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-filter"], _ = expandUserLdapDynamicMappingGroupFilter(d, i["group_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_member_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-member-check"], _ = expandUserLdapDynamicMappingGroupMemberCheck(d, i["group_member_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-object-filter"], _ = expandUserLdapDynamicMappingGroupObjectFilter(d, i["group_object_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_object_search_base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-object-search-base"], _ = expandUserLdapDynamicMappingGroupObjectSearchBase(d, i["group_object_search_base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_search_base"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-search-base"], _ = expandUserLdapDynamicMappingGroupSearchBase(d, i["group_search_base"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserLdapDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserLdapDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-connections"], _ = expandUserLdapDynamicMappingMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_attr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member-attr"], _ = expandUserLdapDynamicMappingMemberAttr(d, i["member_attr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obtain_user_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obtain-user-info"], _ = expandUserLdapDynamicMappingObtainUserInfo(d, i["obtain_user_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandUserLdapDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_attr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-attr"], _ = expandUserLdapDynamicMappingPasswordAttr(d, i["password_attr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_expiry_warning"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-expiry-warning"], _ = expandUserLdapDynamicMappingPasswordExpiryWarning(d, i["password_expiry_warning"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-renewal"], _ = expandUserLdapDynamicMappingPasswordRenewal(d, i["password_renewal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserLdapDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retrieve_protection_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retrieve-protection-profile"], _ = expandUserLdapDynamicMappingRetrieveProtectionProfile(d, i["retrieve_protection_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-type"], _ = expandUserLdapDynamicMappingSearchType(d, i["search_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-server"], _ = expandUserLdapDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secure"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secure"], _ = expandUserLdapDynamicMappingSecure(d, i["secure"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserLdapDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-identity-check"], _ = expandUserLdapDynamicMappingServerIdentityCheck(d, i["server_identity_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserLdapDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip-interface"], _ = expandUserLdapDynamicMappingSourceIpInterface(d, i["source_ip_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-port"], _ = expandUserLdapDynamicMappingSourcePort(d, i["source_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_proto_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-max-proto-version"], _ = expandUserLdapDynamicMappingSslMaxProtoVersion(d, i["ssl_max_proto_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_proto_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-min-proto-version"], _ = expandUserLdapDynamicMappingSslMinProtoVersion(d, i["ssl_min_proto_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status-ttl"], _ = expandUserLdapDynamicMappingStatusTtl(d, i["status_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-server"], _ = expandUserLdapDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor"], _ = expandUserLdapDynamicMappingTwoFactor(d, i["two_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-authentication"], _ = expandUserLdapDynamicMappingTwoFactorAuthentication(d, i["two_factor_authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-filter"], _ = expandUserLdapDynamicMappingTwoFactorFilter(d, i["two_factor_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-notification"], _ = expandUserLdapDynamicMappingTwoFactorNotification(d, i["two_factor_notification"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandUserLdapDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_exchange_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-info-exchange-server"], _ = expandUserLdapDynamicMappingUserInfoExchangeServer(d, i["user_info_exchange_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username"], _ = expandUserLdapDynamicMappingUsername(d, i["username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "validate_server_certificate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["validate-server-certificate"], _ = expandUserLdapDynamicMappingValidateServerCertificate(d, i["validate_server_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserLdapDynamicMappingVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserLdapDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandUserLdapDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserLdapDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserLdapDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAccountKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAccountKeyUpnSan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingClientCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroupFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroupObjectSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingGroupSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingMemberAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingObtainUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingPasswordAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingRetrieveProtectionProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSearchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSslMaxProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingTwoFactorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapDynamicMappingUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingValidateServerCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDynamicMappingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapObtainUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapPasswordAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSearchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyUpnSan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSslMaxProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapValidateServerCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_key_cert_field"); ok || d.HasChange("account_key_cert_field") {
		t, err := expandUserLdapAccountKeyCertField(d, v, "account_key_cert_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok || d.HasChange("account_key_filter") {
		t, err := expandUserLdapAccountKeyFilter(d, v, "account_key_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok || d.HasChange("account_key_processing") {
		t, err := expandUserLdapAccountKeyProcessing(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("antiphish"); ok || d.HasChange("antiphish") {
		t, err := expandUserLdapAntiphish(d, v, "antiphish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandUserLdapCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandUserLdapClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("client_cert_auth"); ok || d.HasChange("client_cert_auth") {
		t, err := expandUserLdapClientCertAuth(d, v, "client_cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok || d.HasChange("cnid") {
		t, err := expandUserLdapCnid(d, v, "cnid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok || d.HasChange("dn") {
		t, err := expandUserLdapDn(d, v, "dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserLdapDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok || d.HasChange("group_filter") {
		t, err := expandUserLdapGroupFilter(d, v, "group_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok || d.HasChange("group_member_check") {
		t, err := expandUserLdapGroupMemberCheck(d, v, "group_member_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok || d.HasChange("group_object_filter") {
		t, err := expandUserLdapGroupObjectFilter(d, v, "group_object_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok || d.HasChange("group_search_base") {
		t, err := expandUserLdapGroupSearchBase(d, v, "group_search_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserLdapInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserLdapInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok || d.HasChange("member_attr") {
		t, err := expandUserLdapMemberAttr(d, v, "member_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserLdapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obtain_user_info"); ok || d.HasChange("obtain_user_info") {
		t, err := expandUserLdapObtainUserInfo(d, v, "obtain_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obtain-user-info"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserLdapPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password_attr"); ok || d.HasChange("password_attr") {
		t, err := expandUserLdapPasswordAttr(d, v, "password_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-attr"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok || d.HasChange("password_expiry_warning") {
		t, err := expandUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok || d.HasChange("password_renewal") {
		t, err := expandUserLdapPasswordRenewal(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserLdapPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok || d.HasChange("search_type") {
		t, err := expandUserLdapSearchType(d, v, "search_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandUserLdapSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandUserLdapSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandUserLdapServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandUserLdapServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandUserLdapSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandUserLdapSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandUserLdapSourcePort(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok || d.HasChange("ssl_min_proto_version") {
		t, err := expandUserLdapSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("status_ttl"); ok || d.HasChange("status_ttl") {
		t, err := expandUserLdapStatusTtl(d, v, "status_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandUserLdapTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandUserLdapTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandUserLdapTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_filter"); ok || d.HasChange("two_factor_filter") {
		t, err := expandUserLdapTwoFactorFilter(d, v, "two_factor_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-filter"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandUserLdapTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserLdapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_exchange_server"); ok || d.HasChange("user_info_exchange_server") {
		t, err := expandUserLdapUserInfoExchangeServer(d, v, "user_info_exchange_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-exchange-server"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandUserLdapUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("account_key_upn_san"); ok || d.HasChange("account_key_upn_san") {
		t, err := expandUserLdapAccountKeyUpnSan(d, v, "account_key_upn_san")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-upn-san"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandUserLdapVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandUserLdapMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_proto_version"); ok || d.HasChange("ssl_max_proto_version") {
		t, err := expandUserLdapSslMaxProtoVersion(d, v, "ssl_max_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("validate_server_certificate"); ok || d.HasChange("validate_server_certificate") {
		t, err := expandUserLdapValidateServerCertificate(d, v, "validate_server_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-server-certificate"] = t
		}
	}

	return &obj, nil
}
