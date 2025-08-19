// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP User Configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSnmpUserCreate,
		Read:   resourceWirelessControllerSnmpUserRead,
		Update: resourceWirelessControllerSnmpUserUpdate,
		Delete: resourceWirelessControllerSnmpUserDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_pwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"notify_hosts": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priv_pwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"queries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSnmpUser resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerSnmpUser(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSnmpUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerSnmpUserRead(d, m)
}

func resourceWirelessControllerSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpUser resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSnmpUser(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerSnmpUserRead(d, m)
}

func resourceWirelessControllerSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerSnmpUser(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpUserRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSnmpUserAuthProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserNotifyHosts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSnmpUserPrivProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserQueries2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserSecurityLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserTrapStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_proto", flattenWirelessControllerSnmpUserAuthProto2edl(o["auth-proto"], d, "auth_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-proto"], "WirelessControllerSnmpUser-AuthProto"); ok {
			if err = d.Set("auth_proto", vv); err != nil {
				return fmt.Errorf("Error reading auth_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerSnmpUserName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerSnmpUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notify_hosts", flattenWirelessControllerSnmpUserNotifyHosts2edl(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-hosts"], "WirelessControllerSnmpUser-NotifyHosts"); ok {
			if err = d.Set("notify_hosts", vv); err != nil {
				return fmt.Errorf("Error reading notify_hosts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenWirelessControllerSnmpUserPrivProto2edl(o["priv-proto"], d, "priv_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-proto"], "WirelessControllerSnmpUser-PrivProto"); ok {
			if err = d.Set("priv_proto", vv); err != nil {
				return fmt.Errorf("Error reading priv_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("queries", flattenWirelessControllerSnmpUserQueries2edl(o["queries"], d, "queries")); err != nil {
		if vv, ok := fortiAPIPatch(o["queries"], "WirelessControllerSnmpUser-Queries"); ok {
			if err = d.Set("queries", vv); err != nil {
				return fmt.Errorf("Error reading queries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("security_level", flattenWirelessControllerSnmpUserSecurityLevel2edl(o["security-level"], d, "security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-level"], "WirelessControllerSnmpUser-SecurityLevel"); ok {
			if err = d.Set("security_level", vv); err != nil {
				return fmt.Errorf("Error reading security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	if err = d.Set("status", flattenWirelessControllerSnmpUserStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WirelessControllerSnmpUser-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_status", flattenWirelessControllerSnmpUserTrapStatus2edl(o["trap-status"], d, "trap_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-status"], "WirelessControllerSnmpUser-TrapStatus"); ok {
			if err = d.Set("trap_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_status: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSnmpUserAuthProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserAuthPwd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserNotifyHosts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserPrivProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserPrivPwd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSnmpUserQueries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserSecurityLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserTrapStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_proto"); ok || d.HasChange("auth_proto") {
		t, err := expandWirelessControllerSnmpUserAuthProto2edl(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok || d.HasChange("auth_pwd") {
		t, err := expandWirelessControllerSnmpUserAuthPwd2edl(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerSnmpUserName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts"); ok || d.HasChange("notify_hosts") {
		t, err := expandWirelessControllerSnmpUserNotifyHosts2edl(d, v, "notify_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok || d.HasChange("priv_proto") {
		t, err := expandWirelessControllerSnmpUserPrivProto2edl(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok || d.HasChange("priv_pwd") {
		t, err := expandWirelessControllerSnmpUserPrivPwd2edl(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok || d.HasChange("queries") {
		t, err := expandWirelessControllerSnmpUserQueries2edl(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok || d.HasChange("security_level") {
		t, err := expandWirelessControllerSnmpUserSecurityLevel2edl(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWirelessControllerSnmpUserStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_status"); ok || d.HasChange("trap_status") {
		t, err := expandWirelessControllerSnmpUserTrapStatus2edl(d, v, "trap_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-status"] = t
		}
	}

	return &obj, nil
}
