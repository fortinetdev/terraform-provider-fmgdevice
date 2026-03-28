// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure FSSO active directory servers for polling mode.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFssoPollingCreate,
		Read:   resourceUserFssoPollingRead,
		Update: resourceUserFssoPollingUpdate,
		Delete: resourceUserFssoPollingDelete,

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
			"adgrp": &schema.Schema{
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
			"default_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"logon_history": &schema.Schema{
				Type:     schema.TypeInt,
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
			"polling_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smbv1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
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

func resourceUserFssoPollingCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error creating UserFssoPolling resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserFssoPolling(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserFssoPolling(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserFssoPolling resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateUserFssoPolling(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserFssoPolling resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceUserFssoPollingRead(d, m)
			} else {
				return fmt.Errorf("Error creating UserFssoPolling resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserFssoPollingRead(d, m)
}

func resourceUserFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error updating UserFssoPolling resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserFssoPolling(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserFssoPollingRead(d, m)
}

func resourceUserFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserFssoPolling(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoPollingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserFssoPolling(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFssoPolling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenUserFssoPollingGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingAdgrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserFssoPollingAdgrpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserFssoPolling-Adgrp-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserFssoPollingAdgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingDefaultDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserFssoPollingLogonHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingPollingFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingSmbv1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPollingUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserFssoPolling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_gui_meta", flattenUserFssoPollingGuiMeta(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "UserFssoPolling-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("adgrp", flattenUserFssoPollingAdgrp(o["adgrp"], d, "adgrp")); err != nil {
			if vv, ok := fortiAPIPatch(o["adgrp"], "UserFssoPolling-Adgrp"); ok {
				if err = d.Set("adgrp", vv); err != nil {
					return fmt.Errorf("Error reading adgrp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading adgrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("adgrp"); ok {
			if err = d.Set("adgrp", flattenUserFssoPollingAdgrp(o["adgrp"], d, "adgrp")); err != nil {
				if vv, ok := fortiAPIPatch(o["adgrp"], "UserFssoPolling-Adgrp"); ok {
					if err = d.Set("adgrp", vv); err != nil {
						return fmt.Errorf("Error reading adgrp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading adgrp: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_domain", flattenUserFssoPollingDefaultDomain(o["default-domain"], d, "default_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-domain"], "UserFssoPolling-DefaultDomain"); ok {
			if err = d.Set("default_domain", vv); err != nil {
				return fmt.Errorf("Error reading default_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_domain: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserFssoPollingId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserFssoPolling-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserFssoPollingLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserFssoPolling-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_history", flattenUserFssoPollingLogonHistory(o["logon-history"], d, "logon_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-history"], "UserFssoPolling-LogonHistory"); ok {
			if err = d.Set("logon_history", vv); err != nil {
				return fmt.Errorf("Error reading logon_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_history: %v", err)
		}
	}

	if err = d.Set("polling_frequency", flattenUserFssoPollingPollingFrequency(o["polling-frequency"], d, "polling_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["polling-frequency"], "UserFssoPolling-PollingFrequency"); ok {
			if err = d.Set("polling_frequency", vv); err != nil {
				return fmt.Errorf("Error reading polling_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polling_frequency: %v", err)
		}
	}

	if err = d.Set("port", flattenUserFssoPollingPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserFssoPolling-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server", flattenUserFssoPollingServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "UserFssoPolling-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenUserFssoPollingSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-ntlmv1-auth"], "UserFssoPolling-SmbNtlmv1Auth"); ok {
			if err = d.Set("smb_ntlmv1_auth", vv); err != nil {
				return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if err = d.Set("smbv1", flattenUserFssoPollingSmbv1(o["smbv1"], d, "smbv1")); err != nil {
		if vv, ok := fortiAPIPatch(o["smbv1"], "UserFssoPolling-Smbv1"); ok {
			if err = d.Set("smbv1", vv); err != nil {
				return fmt.Errorf("Error reading smbv1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smbv1: %v", err)
		}
	}

	if err = d.Set("status", flattenUserFssoPollingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserFssoPolling-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user", flattenUserFssoPollingUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "UserFssoPolling-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenUserFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserFssoPollingGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingAdgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandUserFssoPollingAdgrpName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserFssoPollingAdgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingDefaultDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPollingLogonHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserFssoPollingPollingFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingSmbv1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserFssoPolling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandUserFssoPollingGuiMeta(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("adgrp"); ok || d.HasChange("adgrp") {
		t, err := expandUserFssoPollingAdgrp(d, v, "adgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adgrp"] = t
		}
	}

	if v, ok := d.GetOk("default_domain"); ok || d.HasChange("default_domain") {
		t, err := expandUserFssoPollingDefaultDomain(d, v, "default_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-domain"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserFssoPollingId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserFssoPollingLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("logon_history"); ok || d.HasChange("logon_history") {
		t, err := expandUserFssoPollingLogonHistory(d, v, "logon_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-history"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandUserFssoPollingPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("polling_frequency"); ok || d.HasChange("polling_frequency") {
		t, err := expandUserFssoPollingPollingFrequency(d, v, "polling_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polling-frequency"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserFssoPollingPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandUserFssoPollingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok || d.HasChange("smb_ntlmv1_auth") {
		t, err := expandUserFssoPollingSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	if v, ok := d.GetOk("smbv1"); ok || d.HasChange("smbv1") {
		t, err := expandUserFssoPollingSmbv1(d, v, "smbv1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smbv1"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserFssoPollingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandUserFssoPollingUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
