// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure user groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupCreate,
		Read:   resourceUserGroupRead,
		Update: resourceUserGroupUpdate,
		Delete: resourceUserGroupDelete,

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
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"company": &schema.Schema{
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
						"auth_concurrent_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_concurrent_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"authtimeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"company": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expire": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"expire_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"guest": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"company": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"email": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"expiration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"group": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mobile_phone": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
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
									"sponsor": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"user_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"http_digest_realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ldap_memberof": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"logic_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_gui_meta": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"group_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"server_name": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_accounts": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mobile_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"multiple_guest_add": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redir_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sms_custom_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sms_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sponsor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_bookmarks_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sslvpn_cache_cleaner": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_client_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sslvpn_ftp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_http": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_os_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_os_check_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"latest_patch_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"tolerance": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sslvpn_portal": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sslvpn_portal_heading": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_rdp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_samba": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_split_tunneling": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_ssh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_telnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_ip_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_tunnel_startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_virtual_desktop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_vnc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sslvpn_webapp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expire_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"company": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expiration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mobile_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
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
						"sponsor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"http_digest_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_gui_meta": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_accounts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mobile_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"multiple_guest_add": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sponsor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logic_type": &schema.Schema{
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

func resourceUserGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating UserGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating UserGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserGroup resource from API: %v", err)
	}
	return nil
}

func flattenUserGroupAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserGroupDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := i["auth-concurrent-override"]; ok {
			v := flattenUserGroupDynamicMappingAuthConcurrentOverride(i["auth-concurrent-override"], d, pre_append)
			tmp["auth_concurrent_override"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-AuthConcurrentOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := i["auth-concurrent-value"]; ok {
			v := flattenUserGroupDynamicMappingAuthConcurrentValue(i["auth-concurrent-value"], d, pre_append)
			tmp["auth_concurrent_value"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-AuthConcurrentValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := i["authtimeout"]; ok {
			v := flattenUserGroupDynamicMappingAuthtimeout(i["authtimeout"], d, pre_append)
			tmp["authtimeout"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Authtimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenUserGroupDynamicMappingCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenUserGroupDynamicMappingEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire"
		if _, ok := i["expire"]; ok {
			v := flattenUserGroupDynamicMappingExpire(i["expire"], d, pre_append)
			tmp["expire"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Expire")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire_type"
		if _, ok := i["expire-type"]; ok {
			v := flattenUserGroupDynamicMappingExpireType(i["expire-type"], d, pre_append)
			tmp["expire_type"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-ExpireType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_type"
		if _, ok := i["group-type"]; ok {
			v := flattenUserGroupDynamicMappingGroupType(i["group-type"], d, pre_append)
			tmp["group_type"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-GroupType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guest"
		if _, ok := i["guest"]; ok {
			v := flattenUserGroupDynamicMappingGuest(i["guest"], d, pre_append)
			tmp["guest"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Guest")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_digest_realm"
		if _, ok := i["http-digest-realm"]; ok {
			v := flattenUserGroupDynamicMappingHttpDigestRealm(i["http-digest-realm"], d, pre_append)
			tmp["http_digest_realm"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-HttpDigestRealm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserGroupDynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_memberof"
		if _, ok := i["ldap-memberof"]; ok {
			v := flattenUserGroupDynamicMappingLdapMemberof(i["ldap-memberof"], d, pre_append)
			tmp["ldap_memberof"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-LdapMemberof")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logic_type"
		if _, ok := i["logic-type"]; ok {
			v := flattenUserGroupDynamicMappingLogicType(i["logic-type"], d, pre_append)
			tmp["logic_type"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-LogicType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := i["match"]; ok {
			v := flattenUserGroupDynamicMappingMatch(i["match"], d, pre_append)
			tmp["match"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Match")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_accounts"
		if _, ok := i["max-accounts"]; ok {
			v := flattenUserGroupDynamicMappingMaxAccounts(i["max-accounts"], d, pre_append)
			tmp["max_accounts"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-MaxAccounts")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenUserGroupDynamicMappingMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenUserGroupDynamicMappingMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multiple_guest_add"
		if _, ok := i["multiple-guest-add"]; ok {
			v := flattenUserGroupDynamicMappingMultipleGuestAdd(i["multiple-guest-add"], d, pre_append)
			tmp["multiple_guest_add"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-MultipleGuestAdd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenUserGroupDynamicMappingPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redir_url"
		if _, ok := i["redir-url"]; ok {
			v := flattenUserGroupDynamicMappingRedirUrl(i["redir-url"], d, pre_append)
			tmp["redir_url"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-RedirUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := i["sms-custom-server"]; ok {
			v := flattenUserGroupDynamicMappingSmsCustomServer(i["sms-custom-server"], d, pre_append)
			tmp["sms_custom_server"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SmsCustomServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := i["sms-server"]; ok {
			v := flattenUserGroupDynamicMappingSmsServer(i["sms-server"], d, pre_append)
			tmp["sms_server"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SmsServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenUserGroupDynamicMappingSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_bookmarks_group"
		if _, ok := i["sslvpn-bookmarks-group"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnBookmarksGroup(i["sslvpn-bookmarks-group"], d, pre_append)
			tmp["sslvpn_bookmarks_group"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnBookmarksGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_cache_cleaner"
		if _, ok := i["sslvpn-cache-cleaner"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnCacheCleaner(i["sslvpn-cache-cleaner"], d, pre_append)
			tmp["sslvpn_cache_cleaner"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnCacheCleaner")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_client_check"
		if _, ok := i["sslvpn-client-check"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnClientCheck(i["sslvpn-client-check"], d, pre_append)
			tmp["sslvpn_client_check"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnClientCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ftp"
		if _, ok := i["sslvpn-ftp"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnFtp(i["sslvpn-ftp"], d, pre_append)
			tmp["sslvpn_ftp"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnFtp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_http"
		if _, ok := i["sslvpn-http"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnHttp(i["sslvpn-http"], d, pre_append)
			tmp["sslvpn_http"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnHttp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check"
		if _, ok := i["sslvpn-os-check"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnOsCheck(i["sslvpn-os-check"], d, pre_append)
			tmp["sslvpn_os_check"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnOsCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check_list"
		if _, ok := i["sslvpn-os-check-list"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnOsCheckList(i["sslvpn-os-check-list"], d, pre_append)
			tmp["sslvpn_os_check_list"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnOsCheckList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal"
		if _, ok := i["sslvpn-portal"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnPortal(i["sslvpn-portal"], d, pre_append)
			tmp["sslvpn_portal"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnPortal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal_heading"
		if _, ok := i["sslvpn-portal-heading"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnPortalHeading(i["sslvpn-portal-heading"], d, pre_append)
			tmp["sslvpn_portal_heading"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnPortalHeading")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_rdp"
		if _, ok := i["sslvpn-rdp"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnRdp(i["sslvpn-rdp"], d, pre_append)
			tmp["sslvpn_rdp"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnRdp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_samba"
		if _, ok := i["sslvpn-samba"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnSamba(i["sslvpn-samba"], d, pre_append)
			tmp["sslvpn_samba"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnSamba")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_split_tunneling"
		if _, ok := i["sslvpn-split-tunneling"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnSplitTunneling(i["sslvpn-split-tunneling"], d, pre_append)
			tmp["sslvpn_split_tunneling"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnSplitTunneling")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ssh"
		if _, ok := i["sslvpn-ssh"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnSsh(i["sslvpn-ssh"], d, pre_append)
			tmp["sslvpn_ssh"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnSsh")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_telnet"
		if _, ok := i["sslvpn-telnet"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnTelnet(i["sslvpn-telnet"], d, pre_append)
			tmp["sslvpn_telnet"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnTelnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel"
		if _, ok := i["sslvpn-tunnel"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnTunnel(i["sslvpn-tunnel"], d, pre_append)
			tmp["sslvpn_tunnel"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnTunnel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_endip"
		if _, ok := i["sslvpn-tunnel-endip"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnTunnelEndip(i["sslvpn-tunnel-endip"], d, pre_append)
			tmp["sslvpn_tunnel_endip"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnTunnelEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_ip_mode"
		if _, ok := i["sslvpn-tunnel-ip-mode"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnTunnelIpMode(i["sslvpn-tunnel-ip-mode"], d, pre_append)
			tmp["sslvpn_tunnel_ip_mode"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnTunnelIpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_startip"
		if _, ok := i["sslvpn-tunnel-startip"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnTunnelStartip(i["sslvpn-tunnel-startip"], d, pre_append)
			tmp["sslvpn_tunnel_startip"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnTunnelStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_virtual_desktop"
		if _, ok := i["sslvpn-virtual-desktop"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnVirtualDesktop(i["sslvpn-virtual-desktop"], d, pre_append)
			tmp["sslvpn_virtual_desktop"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnVirtualDesktop")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_vnc"
		if _, ok := i["sslvpn-vnc"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnVnc(i["sslvpn-vnc"], d, pre_append)
			tmp["sslvpn_vnc"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnVnc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_webapp"
		if _, ok := i["sslvpn-webapp"]; ok {
			v := flattenUserGroupDynamicMappingSslvpnWebapp(i["sslvpn-webapp"], d, pre_append)
			tmp["sslvpn_webapp"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SslvpnWebapp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := i["sso-attribute-value"]; ok {
			v := flattenUserGroupDynamicMappingSsoAttributeValue(i["sso-attribute-value"], d, pre_append)
			tmp["sso_attribute_value"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-SsoAttributeValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenUserGroupDynamicMappingUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-UserId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := i["user-name"]; ok {
			v := flattenUserGroupDynamicMappingUserName(i["user-name"], d, pre_append)
			tmp["user_name"] = fortiAPISubPartPatch(v, "UserGroup-DynamicMapping-UserName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserGroupDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserGroupDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingExpireType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenUserGroupDynamicMappingGuestComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenUserGroupDynamicMappingGuestCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenUserGroupDynamicMappingGuestEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {
			v := flattenUserGroupDynamicMappingGuestExpiration(i["expiration"], d, pre_append)
			tmp["expiration"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Expiration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			v := flattenUserGroupDynamicMappingGuestGroup(i["group"], d, pre_append)
			tmp["group"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Group")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserGroupDynamicMappingGuestId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenUserGroupDynamicMappingGuestMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenUserGroupDynamicMappingGuestName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenUserGroupDynamicMappingGuestSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenUserGroupDynamicMappingGuestUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Guest-UserId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupDynamicMappingGuestComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingGuestUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingHttpDigestRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingLdapMemberof(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingLogicType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := i["_gui_meta"]; ok {
			v := flattenUserGroupDynamicMappingMatchGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Match-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenUserGroupDynamicMappingMatchGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Match-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserGroupDynamicMappingMatchId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {
			v := flattenUserGroupDynamicMappingMatchServerName(i["server-name"], d, pre_append)
			tmp["server_name"] = fortiAPISubPartPatch(v, "UserGroupDynamicMapping-Match-ServerName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupDynamicMappingMatchGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMatchGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMatchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMatchServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingMaxAccounts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingRedirUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnBookmarksGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingSslvpnCacheCleaner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnClientCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingSslvpnFtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnOsCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnOsCheckList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenUserGroupDynamicMappingSslvpnOsCheckListAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := i["latest-patch-level"]; ok {
		result["latest_patch_level"] = flattenUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(i["latest-patch-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenUserGroupDynamicMappingSslvpnOsCheckListName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "tolerance"
	if _, ok := i["tolerance"]; ok {
		result["tolerance"] = flattenUserGroupDynamicMappingSslvpnOsCheckListTolerance(i["tolerance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenUserGroupDynamicMappingSslvpnOsCheckListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnOsCheckListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnOsCheckListTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupDynamicMappingSslvpnPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnRdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnSamba(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnTunnelEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnTunnelIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnTunnelStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnVirtualDesktop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnVnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSslvpnWebapp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupDynamicMappingUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupExpireType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenUserGroupGuestComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {
			v := flattenUserGroupGuestCompany(i["company"], d, pre_append)
			tmp["company"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Company")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {
			v := flattenUserGroupGuestEmail(i["email"], d, pre_append)
			tmp["email"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Email")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {
			v := flattenUserGroupGuestExpiration(i["expiration"], d, pre_append)
			tmp["expiration"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Expiration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserGroupGuestId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {
			v := flattenUserGroupGuestMobilePhone(i["mobile-phone"], d, pre_append)
			tmp["mobile_phone"] = fortiAPISubPartPatch(v, "UserGroup-Guest-MobilePhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenUserGroupGuestName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {
			v := flattenUserGroupGuestSponsor(i["sponsor"], d, pre_append)
			tmp["sponsor"] = fortiAPISubPartPatch(v, "UserGroup-Guest-Sponsor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {
			v := flattenUserGroupGuestUserId(i["user-id"], d, pre_append)
			tmp["user_id"] = fortiAPISubPartPatch(v, "UserGroup-Guest-UserId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupGuestComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestCompany(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupGuestUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupHttpDigestRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := i["_gui_meta"]; ok {
			v := flattenUserGroupMatchGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "UserGroup-Match-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenUserGroupMatchGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "UserGroup-Match-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserGroupMatchId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserGroup-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {
			v := flattenUserGroupMatchServerName(i["server-name"], d, pre_append)
			tmp["server_name"] = fortiAPISubPartPatch(v, "UserGroup-Match-ServerName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserGroupMatchGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMatchServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupMaxAccounts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupMobilePhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserGroupSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupSponsor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserGroupLogicType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_concurrent_override", flattenUserGroupAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "UserGroup-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenUserGroupAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "UserGroup-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenUserGroupAuthtimeout(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "UserGroup-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("company", flattenUserGroupCompany(o["company"], d, "company")); err != nil {
		if vv, ok := fortiAPIPatch(o["company"], "UserGroup-Company"); ok {
			if err = d.Set("company", vv); err != nil {
				return fmt.Errorf("Error reading company: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading company: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserGroupDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserGroup-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserGroupDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserGroup-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("email", flattenUserGroupEmail(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "UserGroup-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("expire", flattenUserGroupExpire(o["expire"], d, "expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire"], "UserGroup-Expire"); ok {
			if err = d.Set("expire", vv); err != nil {
				return fmt.Errorf("Error reading expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("expire_type", flattenUserGroupExpireType(o["expire-type"], d, "expire_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-type"], "UserGroup-ExpireType"); ok {
			if err = d.Set("expire_type", vv); err != nil {
				return fmt.Errorf("Error reading expire_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_type: %v", err)
		}
	}

	if err = d.Set("group_type", flattenUserGroupGroupType(o["group-type"], d, "group_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-type"], "UserGroup-GroupType"); ok {
			if err = d.Set("group_type", vv); err != nil {
				return fmt.Errorf("Error reading group_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenUserGroupGuest(o["guest"], d, "guest")); err != nil {
			if vv, ok := fortiAPIPatch(o["guest"], "UserGroup-Guest"); ok {
				if err = d.Set("guest", vv); err != nil {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenUserGroupGuest(o["guest"], d, "guest")); err != nil {
				if vv, ok := fortiAPIPatch(o["guest"], "UserGroup-Guest"); ok {
					if err = d.Set("guest", vv); err != nil {
						return fmt.Errorf("Error reading guest: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	if err = d.Set("http_digest_realm", flattenUserGroupHttpDigestRealm(o["http-digest-realm"], d, "http_digest_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-digest-realm"], "UserGroup-HttpDigestRealm"); ok {
			if err = d.Set("http_digest_realm", vv); err != nil {
				return fmt.Errorf("Error reading http_digest_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_digest_realm: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserGroupId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserGroup-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("match", flattenUserGroupMatch(o["match"], d, "match")); err != nil {
			if vv, ok := fortiAPIPatch(o["match"], "UserGroup-Match"); ok {
				if err = d.Set("match", vv); err != nil {
					return fmt.Errorf("Error reading match: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenUserGroupMatch(o["match"], d, "match")); err != nil {
				if vv, ok := fortiAPIPatch(o["match"], "UserGroup-Match"); ok {
					if err = d.Set("match", vv); err != nil {
						return fmt.Errorf("Error reading match: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_accounts", flattenUserGroupMaxAccounts(o["max-accounts"], d, "max_accounts")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-accounts"], "UserGroup-MaxAccounts"); ok {
			if err = d.Set("max_accounts", vv); err != nil {
				return fmt.Errorf("Error reading max_accounts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_accounts: %v", err)
		}
	}

	if err = d.Set("member", flattenUserGroupMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "UserGroup-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("mobile_phone", flattenUserGroupMobilePhone(o["mobile-phone"], d, "mobile_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-phone"], "UserGroup-MobilePhone"); ok {
			if err = d.Set("mobile_phone", vv); err != nil {
				return fmt.Errorf("Error reading mobile_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_phone: %v", err)
		}
	}

	if err = d.Set("multiple_guest_add", flattenUserGroupMultipleGuestAdd(o["multiple-guest-add"], d, "multiple_guest_add")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-guest-add"], "UserGroup-MultipleGuestAdd"); ok {
			if err = d.Set("multiple_guest_add", vv); err != nil {
				return fmt.Errorf("Error reading multiple_guest_add: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_guest_add: %v", err)
		}
	}

	if err = d.Set("name", flattenUserGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password", flattenUserGroupPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "UserGroup-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenUserGroupSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "UserGroup-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenUserGroupSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "UserGroup-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sponsor", flattenUserGroupSponsor(o["sponsor"], d, "sponsor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sponsor"], "UserGroup-Sponsor"); ok {
			if err = d.Set("sponsor", vv); err != nil {
				return fmt.Errorf("Error reading sponsor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sponsor: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenUserGroupSsoAttributeValue(o["sso-attribute-value"], d, "sso_attribute_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value"], "UserGroup-SsoAttributeValue"); ok {
			if err = d.Set("sso_attribute_value", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if err = d.Set("user_id", flattenUserGroupUserId(o["user-id"], d, "user_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-id"], "UserGroup-UserId"); ok {
			if err = d.Set("user_id", vv); err != nil {
				return fmt.Errorf("Error reading user_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("user_name", flattenUserGroupUserName(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "UserGroup-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("logic_type", flattenUserGroupLogicType(o["logic-type"], d, "logic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["logic-type"], "UserGroup-LogicType"); ok {
			if err = d.Set("logic_type", vv); err != nil {
				return fmt.Errorf("Error reading logic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logic_type: %v", err)
		}
	}

	return nil
}

func flattenUserGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserGroupAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandUserGroupDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-override"], _ = expandUserGroupDynamicMappingAuthConcurrentOverride(d, i["auth_concurrent_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-value"], _ = expandUserGroupDynamicMappingAuthConcurrentValue(d, i["auth_concurrent_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authtimeout"], _ = expandUserGroupDynamicMappingAuthtimeout(d, i["authtimeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandUserGroupDynamicMappingCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandUserGroupDynamicMappingEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expire"], _ = expandUserGroupDynamicMappingExpire(d, i["expire"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expire_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expire-type"], _ = expandUserGroupDynamicMappingExpireType(d, i["expire_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-type"], _ = expandUserGroupDynamicMappingGroupType(d, i["group_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guest"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserGroupDynamicMappingGuest(d, i["guest"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["guest"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_digest_realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-digest-realm"], _ = expandUserGroupDynamicMappingHttpDigestRealm(d, i["http_digest_realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserGroupDynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_memberof"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-memberof"], _ = expandUserGroupDynamicMappingLdapMemberof(d, i["ldap_memberof"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logic_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logic-type"], _ = expandUserGroupDynamicMappingLogicType(d, i["logic_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserGroupDynamicMappingMatch(d, i["match"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["match"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_accounts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-accounts"], _ = expandUserGroupDynamicMappingMaxAccounts(d, i["max_accounts"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandUserGroupDynamicMappingMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandUserGroupDynamicMappingMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multiple_guest_add"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["multiple-guest-add"], _ = expandUserGroupDynamicMappingMultipleGuestAdd(d, i["multiple_guest_add"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandUserGroupDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redir_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["redir-url"], _ = expandUserGroupDynamicMappingRedirUrl(d, i["redir_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-custom-server"], _ = expandUserGroupDynamicMappingSmsCustomServer(d, i["sms_custom_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-server"], _ = expandUserGroupDynamicMappingSmsServer(d, i["sms_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandUserGroupDynamicMappingSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_bookmarks_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-bookmarks-group"], _ = expandUserGroupDynamicMappingSslvpnBookmarksGroup(d, i["sslvpn_bookmarks_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_cache_cleaner"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-cache-cleaner"], _ = expandUserGroupDynamicMappingSslvpnCacheCleaner(d, i["sslvpn_cache_cleaner"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_client_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-client-check"], _ = expandUserGroupDynamicMappingSslvpnClientCheck(d, i["sslvpn_client_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ftp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-ftp"], _ = expandUserGroupDynamicMappingSslvpnFtp(d, i["sslvpn_ftp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_http"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-http"], _ = expandUserGroupDynamicMappingSslvpnHttp(d, i["sslvpn_http"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-os-check"], _ = expandUserGroupDynamicMappingSslvpnOsCheck(d, i["sslvpn_os_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_os_check_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserGroupDynamicMappingSslvpnOsCheckList(d, i["sslvpn_os_check_list"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sslvpn-os-check-list"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-portal"], _ = expandUserGroupDynamicMappingSslvpnPortal(d, i["sslvpn_portal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_portal_heading"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-portal-heading"], _ = expandUserGroupDynamicMappingSslvpnPortalHeading(d, i["sslvpn_portal_heading"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_rdp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-rdp"], _ = expandUserGroupDynamicMappingSslvpnRdp(d, i["sslvpn_rdp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_samba"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-samba"], _ = expandUserGroupDynamicMappingSslvpnSamba(d, i["sslvpn_samba"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_split_tunneling"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-split-tunneling"], _ = expandUserGroupDynamicMappingSslvpnSplitTunneling(d, i["sslvpn_split_tunneling"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_ssh"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-ssh"], _ = expandUserGroupDynamicMappingSslvpnSsh(d, i["sslvpn_ssh"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_telnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-telnet"], _ = expandUserGroupDynamicMappingSslvpnTelnet(d, i["sslvpn_telnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel"], _ = expandUserGroupDynamicMappingSslvpnTunnel(d, i["sslvpn_tunnel"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-endip"], _ = expandUserGroupDynamicMappingSslvpnTunnelEndip(d, i["sslvpn_tunnel_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_ip_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-ip-mode"], _ = expandUserGroupDynamicMappingSslvpnTunnelIpMode(d, i["sslvpn_tunnel_ip_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_tunnel_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-tunnel-startip"], _ = expandUserGroupDynamicMappingSslvpnTunnelStartip(d, i["sslvpn_tunnel_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_virtual_desktop"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-virtual-desktop"], _ = expandUserGroupDynamicMappingSslvpnVirtualDesktop(d, i["sslvpn_virtual_desktop"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_vnc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-vnc"], _ = expandUserGroupDynamicMappingSslvpnVnc(d, i["sslvpn_vnc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_webapp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sslvpn-webapp"], _ = expandUserGroupDynamicMappingSslvpnWebapp(d, i["sslvpn_webapp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute-value"], _ = expandUserGroupDynamicMappingSsoAttributeValue(d, i["sso_attribute_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandUserGroupDynamicMappingUserId(d, i["user_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-name"], _ = expandUserGroupDynamicMappingUserName(d, i["user_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandUserGroupDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserGroupDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingExpireType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandUserGroupDynamicMappingGuestComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandUserGroupDynamicMappingGuestCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandUserGroupDynamicMappingGuestEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiration"], _ = expandUserGroupDynamicMappingGuestExpiration(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group"], _ = expandUserGroupDynamicMappingGuestGroup(d, i["group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserGroupDynamicMappingGuestId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandUserGroupDynamicMappingGuestMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandUserGroupDynamicMappingGuestName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandUserGroupDynamicMappingGuestPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandUserGroupDynamicMappingGuestSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandUserGroupDynamicMappingGuestUserId(d, i["user_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupDynamicMappingGuestComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingGuestSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingGuestUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingHttpDigestRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingLdapMemberof(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingLogicType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_gui_meta"], _ = expandUserGroupDynamicMappingMatchGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandUserGroupDynamicMappingMatchGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserGroupDynamicMappingMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-name"], _ = expandUserGroupDynamicMappingMatchServerName(d, i["server_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupDynamicMappingMatchGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMatchGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMatchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMatchServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingMaxAccounts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingMultipleGuestAdd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingRedirUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnBookmarksGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingSslvpnCacheCleaner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnClientCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingSslvpnFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheckList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandUserGroupDynamicMappingSslvpnOsCheckListAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latest-patch-level"], _ = expandUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(d, i["latest_patch_level"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandUserGroupDynamicMappingSslvpnOsCheckListName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "tolerance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tolerance"], _ = expandUserGroupDynamicMappingSslvpnOsCheckListTolerance(d, i["tolerance"], pre_append)
	}

	return result, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheckListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheckListLatestPatchLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheckListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnOsCheckListTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupDynamicMappingSslvpnPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnRdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnSamba(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnTelnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnTunnelEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnTunnelIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnTunnelStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnVirtualDesktop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnVnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSslvpnWebapp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupDynamicMappingUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupExpireType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandUserGroupGuestComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company"], _ = expandUserGroupGuestCompany(d, i["company"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email"], _ = expandUserGroupGuestEmail(d, i["email"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiration"], _ = expandUserGroupGuestExpiration(d, i["expiration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserGroupGuestId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mobile-phone"], _ = expandUserGroupGuestMobilePhone(d, i["mobile_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandUserGroupGuestName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandUserGroupGuestPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sponsor"], _ = expandUserGroupGuestSponsor(d, i["sponsor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-id"], _ = expandUserGroupGuestUserId(d, i["user_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupGuestComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestCompany(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupGuestSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupHttpDigestRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_gui_meta"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_gui_meta"], _ = expandUserGroupMatchGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandUserGroupMatchGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserGroupMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-name"], _ = expandUserGroupMatchServerName(d, i["server_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserGroupMatchGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupMaxAccounts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupMobilePhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMultipleGuestAdd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserGroupSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSponsor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserGroupLogicType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandUserGroupAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandUserGroupAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandUserGroupAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("company"); ok || d.HasChange("company") {
		t, err := expandUserGroupCompany(d, v, "company")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserGroupDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok || d.HasChange("email") {
		t, err := expandUserGroupEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok || d.HasChange("expire") {
		t, err := expandUserGroupExpire(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("expire_type"); ok || d.HasChange("expire_type") {
		t, err := expandUserGroupExpireType(d, v, "expire_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-type"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok || d.HasChange("group_type") {
		t, err := expandUserGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok || d.HasChange("guest") {
		t, err := expandUserGroupGuest(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("http_digest_realm"); ok || d.HasChange("http_digest_realm") {
		t, err := expandUserGroupHttpDigestRealm(d, v, "http_digest_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-digest-realm"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserGroupId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandUserGroupMatch(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("max_accounts"); ok || d.HasChange("max_accounts") {
		t, err := expandUserGroupMaxAccounts(d, v, "max_accounts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-accounts"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandUserGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mobile_phone"); ok || d.HasChange("mobile_phone") {
		t, err := expandUserGroupMobilePhone(d, v, "mobile_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-phone"] = t
		}
	}

	if v, ok := d.GetOk("multiple_guest_add"); ok || d.HasChange("multiple_guest_add") {
		t, err := expandUserGroupMultipleGuestAdd(d, v, "multiple_guest_add")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-guest-add"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserGroupPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandUserGroupSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandUserGroupSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sponsor"); ok || d.HasChange("sponsor") {
		t, err := expandUserGroupSponsor(d, v, "sponsor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sponsor"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok || d.HasChange("sso_attribute_value") {
		t, err := expandUserGroupSsoAttributeValue(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok || d.HasChange("user_id") {
		t, err := expandUserGroupUserId(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandUserGroupUserName(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	if v, ok := d.GetOk("logic_type"); ok || d.HasChange("logic_type") {
		t, err := expandUserGroupLogicType(d, v, "logic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logic-type"] = t
		}
	}

	return &obj, nil
}
