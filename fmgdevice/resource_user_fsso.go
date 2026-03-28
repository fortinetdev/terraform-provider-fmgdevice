// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Fortinet Single Sign On (FSSO) agents.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserFsso() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFssoCreate,
		Read:   resourceUserFssoRead,
		Update: resourceUserFssoUpdate,
		Delete: resourceUserFssoDelete,

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
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_gui_meta": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
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
						"group_poll_interval": &schema.Schema{
							Type:     schema.TypeInt,
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
						"ldap_poll": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldap_poll_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ldap_poll_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ldap_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"logon_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"password2": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"password3": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"password4": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"password5": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port5": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sni": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_server_host_ip_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_trusted_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_info_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"group_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
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
			"ldap_poll": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_poll_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"logon_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password3": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password4": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password5": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port5": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_host_ip_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_trusted_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
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

func resourceUserFssoCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error creating UserFsso resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserFsso(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserFsso(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserFsso resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserFsso(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserFsso resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserFsso(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserFsso(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserFsso resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserFsso(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserFsso resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFsso(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserFsso resource from API: %v", err)
	}
	return nil
}

func flattenUserFssoGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserFssoDynamicMappingGuiMeta(i["_gui_meta"], d, pre_append)
			tmp["_gui_meta"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-GuiMeta")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenUserFssoDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_poll_interval"
		if _, ok := i["group-poll-interval"]; ok {
			v := flattenUserFssoDynamicMappingGroupPollInterval(i["group-poll-interval"], d, pre_append)
			tmp["group_poll_interval"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-GroupPollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserFssoDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserFssoDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll"
		if _, ok := i["ldap-poll"]; ok {
			v := flattenUserFssoDynamicMappingLdapPoll(i["ldap-poll"], d, pre_append)
			tmp["ldap_poll"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-LdapPoll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_filter"
		if _, ok := i["ldap-poll-filter"]; ok {
			v := flattenUserFssoDynamicMappingLdapPollFilter(i["ldap-poll-filter"], d, pre_append)
			tmp["ldap_poll_filter"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-LdapPollFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_interval"
		if _, ok := i["ldap-poll-interval"]; ok {
			v := flattenUserFssoDynamicMappingLdapPollInterval(i["ldap-poll-interval"], d, pre_append)
			tmp["ldap_poll_interval"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-LdapPollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := i["ldap-server"]; ok {
			v := flattenUserFssoDynamicMappingLdapServer(i["ldap-server"], d, pre_append)
			tmp["ldap_server"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-LdapServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_timeout"
		if _, ok := i["logon-timeout"]; ok {
			v := flattenUserFssoDynamicMappingLogonTimeout(i["logon-timeout"], d, pre_append)
			tmp["logon_timeout"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-LogonTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserFssoDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port2"
		if _, ok := i["port2"]; ok {
			v := flattenUserFssoDynamicMappingPort2(i["port2"], d, pre_append)
			tmp["port2"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Port2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port3"
		if _, ok := i["port3"]; ok {
			v := flattenUserFssoDynamicMappingPort3(i["port3"], d, pre_append)
			tmp["port3"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Port3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port4"
		if _, ok := i["port4"]; ok {
			v := flattenUserFssoDynamicMappingPort4(i["port4"], d, pre_append)
			tmp["port4"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Port4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port5"
		if _, ok := i["port5"]; ok {
			v := flattenUserFssoDynamicMappingPort5(i["port5"], d, pre_append)
			tmp["port5"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Port5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserFssoDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server2"
		if _, ok := i["server2"]; ok {
			v := flattenUserFssoDynamicMappingServer2(i["server2"], d, pre_append)
			tmp["server2"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Server2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server3"
		if _, ok := i["server3"]; ok {
			v := flattenUserFssoDynamicMappingServer3(i["server3"], d, pre_append)
			tmp["server3"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Server3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server4"
		if _, ok := i["server4"]; ok {
			v := flattenUserFssoDynamicMappingServer4(i["server4"], d, pre_append)
			tmp["server4"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Server4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server5"
		if _, ok := i["server5"]; ok {
			v := flattenUserFssoDynamicMappingServer5(i["server5"], d, pre_append)
			tmp["server5"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Server5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := i["sni"]; ok {
			v := flattenUserFssoDynamicMappingSni(i["sni"], d, pre_append)
			tmp["sni"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Sni")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserFssoDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip6"
		if _, ok := i["source-ip6"]; ok {
			v := flattenUserFssoDynamicMappingSourceIp6(i["source-ip6"], d, pre_append)
			tmp["source_ip6"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-SourceIp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl"
		if _, ok := i["ssl"]; ok {
			v := flattenUserFssoDynamicMappingSsl(i["ssl"], d, pre_append)
			tmp["ssl"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Ssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_host_ip_check"
		if _, ok := i["ssl-server-host-ip-check"]; ok {
			v := flattenUserFssoDynamicMappingSslServerHostIpCheck(i["ssl-server-host-ip-check"], d, pre_append)
			tmp["ssl_server_host_ip_check"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-SslServerHostIpCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_trusted_cert"
		if _, ok := i["ssl-trusted-cert"]; ok {
			v := flattenUserFssoDynamicMappingSslTrustedCert(i["ssl-trusted-cert"], d, pre_append)
			tmp["ssl_trusted_cert"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-SslTrustedCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenUserFssoDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_server"
		if _, ok := i["user-info-server"]; ok {
			v := flattenUserFssoDynamicMappingUserInfoServer(i["user-info-server"], d, pre_append)
			tmp["user_info_server"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-UserInfoServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserFssoDynamicMappingVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserFsso-DynamicMapping-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserFssoDynamicMappingGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserFssoDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserFssoDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserFssoDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserFssoDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserFssoDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingGroupPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingLdapPoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingLdapPollFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingLdapPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoDynamicMappingLogonTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingPort4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingPort5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingServer5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSslServerHostIpCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingSslTrustedCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoDynamicMappingUserInfoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoDynamicMappingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoGroupPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoLdapPoll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoLdapPollFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoLdapPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoLogonTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSslServerHostIpCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSslTrustedCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoUserInfoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserFsso(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_gui_meta", flattenUserFssoGuiMeta(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "UserFsso-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserFssoDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserFsso-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserFssoDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserFsso-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_poll_interval", flattenUserFssoGroupPollInterval(o["group-poll-interval"], d, "group_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-poll-interval"], "UserFsso-GroupPollInterval"); ok {
			if err = d.Set("group_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading group_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_poll_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserFssoInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserFsso-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserFssoInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserFsso-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ldap_poll", flattenUserFssoLdapPoll(o["ldap-poll"], d, "ldap_poll")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll"], "UserFsso-LdapPoll"); ok {
			if err = d.Set("ldap_poll", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll: %v", err)
		}
	}

	if err = d.Set("ldap_poll_filter", flattenUserFssoLdapPollFilter(o["ldap-poll-filter"], d, "ldap_poll_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-filter"], "UserFsso-LdapPollFilter"); ok {
			if err = d.Set("ldap_poll_filter", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
		}
	}

	if err = d.Set("ldap_poll_interval", flattenUserFssoLdapPollInterval(o["ldap-poll-interval"], d, "ldap_poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-poll-interval"], "UserFsso-LdapPollInterval"); ok {
			if err = d.Set("ldap_poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserFssoLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserFsso-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_timeout", flattenUserFssoLogonTimeout(o["logon-timeout"], d, "logon_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-timeout"], "UserFsso-LogonTimeout"); ok {
			if err = d.Set("logon_timeout", vv); err != nil {
				return fmt.Errorf("Error reading logon_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_timeout: %v", err)
		}
	}

	if err = d.Set("name", flattenUserFssoName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserFsso-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenUserFssoPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserFsso-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port2", flattenUserFssoPort2(o["port2"], d, "port2")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2"], "UserFsso-Port2"); ok {
			if err = d.Set("port2", vv); err != nil {
				return fmt.Errorf("Error reading port2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("port3", flattenUserFssoPort3(o["port3"], d, "port3")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3"], "UserFsso-Port3"); ok {
			if err = d.Set("port3", vv); err != nil {
				return fmt.Errorf("Error reading port3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("port4", flattenUserFssoPort4(o["port4"], d, "port4")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4"], "UserFsso-Port4"); ok {
			if err = d.Set("port4", vv); err != nil {
				return fmt.Errorf("Error reading port4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4: %v", err)
		}
	}

	if err = d.Set("port5", flattenUserFssoPort5(o["port5"], d, "port5")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5"], "UserFsso-Port5"); ok {
			if err = d.Set("port5", vv); err != nil {
				return fmt.Errorf("Error reading port5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5: %v", err)
		}
	}

	if err = d.Set("server", flattenUserFssoServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "UserFsso-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server2", flattenUserFssoServer2(o["server2"], d, "server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["server2"], "UserFsso-Server2"); ok {
			if err = d.Set("server2", vv); err != nil {
				return fmt.Errorf("Error reading server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("server3", flattenUserFssoServer3(o["server3"], d, "server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["server3"], "UserFsso-Server3"); ok {
			if err = d.Set("server3", vv); err != nil {
				return fmt.Errorf("Error reading server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("server4", flattenUserFssoServer4(o["server4"], d, "server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["server4"], "UserFsso-Server4"); ok {
			if err = d.Set("server4", vv); err != nil {
				return fmt.Errorf("Error reading server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server4: %v", err)
		}
	}

	if err = d.Set("server5", flattenUserFssoServer5(o["server5"], d, "server5")); err != nil {
		if vv, ok := fortiAPIPatch(o["server5"], "UserFsso-Server5"); ok {
			if err = d.Set("server5", vv); err != nil {
				return fmt.Errorf("Error reading server5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server5: %v", err)
		}
	}

	if err = d.Set("sni", flattenUserFssoSni(o["sni"], d, "sni")); err != nil {
		if vv, ok := fortiAPIPatch(o["sni"], "UserFsso-Sni"); ok {
			if err = d.Set("sni", vv); err != nil {
				return fmt.Errorf("Error reading sni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserFssoSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "UserFsso-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserFssoSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "UserFsso-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("ssl", flattenUserFssoSsl(o["ssl"], d, "ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl"], "UserFsso-Ssl"); ok {
			if err = d.Set("ssl", vv); err != nil {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_server_host_ip_check", flattenUserFssoSslServerHostIpCheck(o["ssl-server-host-ip-check"], d, "ssl_server_host_ip_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-server-host-ip-check"], "UserFsso-SslServerHostIpCheck"); ok {
			if err = d.Set("ssl_server_host_ip_check", vv); err != nil {
				return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
		}
	}

	if err = d.Set("ssl_trusted_cert", flattenUserFssoSslTrustedCert(o["ssl-trusted-cert"], d, "ssl_trusted_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-trusted-cert"], "UserFsso-SslTrustedCert"); ok {
			if err = d.Set("ssl_trusted_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
		}
	}

	if err = d.Set("type", flattenUserFssoType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserFsso-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_info_server", flattenUserFssoUserInfoServer(o["user-info-server"], d, "user_info_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-info-server"], "UserFsso-UserInfoServer"); ok {
			if err = d.Set("user_info_server", vv); err != nil {
				return fmt.Errorf("Error reading user_info_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_info_server: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserFssoVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "UserFsso-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserFssoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserFssoGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_gui_meta"], _ = expandUserFssoDynamicMappingGuiMeta(d, i["_gui_meta"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserFssoDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-poll-interval"], _ = expandUserFssoDynamicMappingGroupPollInterval(d, i["group_poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserFssoDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserFssoDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll"], _ = expandUserFssoDynamicMappingLdapPoll(d, i["ldap_poll"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll-filter"], _ = expandUserFssoDynamicMappingLdapPollFilter(d, i["ldap_poll_filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-poll-interval"], _ = expandUserFssoDynamicMappingLdapPollInterval(d, i["ldap_poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-server"], _ = expandUserFssoDynamicMappingLdapServer(d, i["ldap_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-timeout"], _ = expandUserFssoDynamicMappingLogonTimeout(d, i["logon_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandUserFssoDynamicMappingPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password2"], _ = expandUserFssoDynamicMappingPassword2(d, i["password2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password3"], _ = expandUserFssoDynamicMappingPassword3(d, i["password3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password4"], _ = expandUserFssoDynamicMappingPassword4(d, i["password4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password5"], _ = expandUserFssoDynamicMappingPassword5(d, i["password5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserFssoDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port2"], _ = expandUserFssoDynamicMappingPort2(d, i["port2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port3"], _ = expandUserFssoDynamicMappingPort3(d, i["port3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port4"], _ = expandUserFssoDynamicMappingPort4(d, i["port4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port5"], _ = expandUserFssoDynamicMappingPort5(d, i["port5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserFssoDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server2"], _ = expandUserFssoDynamicMappingServer2(d, i["server2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server3"], _ = expandUserFssoDynamicMappingServer3(d, i["server3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server4"], _ = expandUserFssoDynamicMappingServer4(d, i["server4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server5"], _ = expandUserFssoDynamicMappingServer5(d, i["server5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sni"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sni"], _ = expandUserFssoDynamicMappingSni(d, i["sni"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserFssoDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip6"], _ = expandUserFssoDynamicMappingSourceIp6(d, i["source_ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl"], _ = expandUserFssoDynamicMappingSsl(d, i["ssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_server_host_ip_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-server-host-ip-check"], _ = expandUserFssoDynamicMappingSslServerHostIpCheck(d, i["ssl_server_host_ip_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_trusted_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ssl-trusted-cert"], _ = expandUserFssoDynamicMappingSslTrustedCert(d, i["ssl_trusted_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandUserFssoDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_info_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-info-server"], _ = expandUserFssoDynamicMappingUserInfoServer(d, i["user_info_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserFssoDynamicMappingVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserFssoDynamicMappingGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandUserFssoDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserFssoDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserFssoDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingGroupPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingLdapPoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingLdapPollFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingLdapPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingLogonTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingPassword4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingPassword5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingPort4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingPort5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingServer5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSslServerHostIpCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingSslTrustedCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoDynamicMappingUserInfoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoDynamicMappingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoGroupPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPoll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPollFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoLogonTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPassword4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPassword5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSslServerHostIpCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSslTrustedCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoUserInfoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserFsso(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandUserFssoGuiMeta(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserFssoDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_poll_interval"); ok || d.HasChange("group_poll_interval") {
		t, err := expandUserFssoGroupPollInterval(d, v, "group_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserFssoInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserFssoInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll"); ok || d.HasChange("ldap_poll") {
		t, err := expandUserFssoLdapPoll(d, v, "ldap_poll")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_filter"); ok || d.HasChange("ldap_poll_filter") {
		t, err := expandUserFssoLdapPollFilter(d, v, "ldap_poll_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-filter"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_interval"); ok || d.HasChange("ldap_poll_interval") {
		t, err := expandUserFssoLdapPollInterval(d, v, "ldap_poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserFssoLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("logon_timeout"); ok || d.HasChange("logon_timeout") {
		t, err := expandUserFssoLogonTimeout(d, v, "logon_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-timeout"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserFssoName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserFssoPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok || d.HasChange("password2") {
		t, err := expandUserFssoPassword2(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok || d.HasChange("password3") {
		t, err := expandUserFssoPassword3(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("password4"); ok || d.HasChange("password4") {
		t, err := expandUserFssoPassword4(d, v, "password4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password4"] = t
		}
	}

	if v, ok := d.GetOk("password5"); ok || d.HasChange("password5") {
		t, err := expandUserFssoPassword5(d, v, "password5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password5"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserFssoPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok || d.HasChange("port2") {
		t, err := expandUserFssoPort2(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok || d.HasChange("port3") {
		t, err := expandUserFssoPort3(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("port4"); ok || d.HasChange("port4") {
		t, err := expandUserFssoPort4(d, v, "port4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4"] = t
		}
	}

	if v, ok := d.GetOk("port5"); ok || d.HasChange("port5") {
		t, err := expandUserFssoPort5(d, v, "port5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandUserFssoServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server2"); ok || d.HasChange("server2") {
		t, err := expandUserFssoServer2(d, v, "server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	}

	if v, ok := d.GetOk("server3"); ok || d.HasChange("server3") {
		t, err := expandUserFssoServer3(d, v, "server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	}

	if v, ok := d.GetOk("server4"); ok || d.HasChange("server4") {
		t, err := expandUserFssoServer4(d, v, "server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server4"] = t
		}
	}

	if v, ok := d.GetOk("server5"); ok || d.HasChange("server5") {
		t, err := expandUserFssoServer5(d, v, "server5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server5"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok || d.HasChange("sni") {
		t, err := expandUserFssoSni(d, v, "sni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandUserFssoSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandUserFssoSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok || d.HasChange("ssl") {
		t, err := expandUserFssoSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_host_ip_check"); ok || d.HasChange("ssl_server_host_ip_check") {
		t, err := expandUserFssoSslServerHostIpCheck(d, v, "ssl_server_host_ip_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-host-ip-check"] = t
		}
	}

	if v, ok := d.GetOk("ssl_trusted_cert"); ok || d.HasChange("ssl_trusted_cert") {
		t, err := expandUserFssoSslTrustedCert(d, v, "ssl_trusted_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-trusted-cert"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserFssoType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_info_server"); ok || d.HasChange("user_info_server") {
		t, err := expandUserFssoUserInfoServer(d, v, "user_info_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-server"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandUserFssoVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
