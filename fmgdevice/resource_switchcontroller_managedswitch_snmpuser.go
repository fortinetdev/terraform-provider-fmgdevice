// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit Simple Network Management Protocol (SNMP) users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSnmpUserCreate,
		Read:   resourceSwitchControllerManagedSwitchSnmpUserRead,
		Update: resourceSwitchControllerManagedSwitchSnmpUserUpdate,
		Delete: resourceSwitchControllerManagedSwitchSnmpUserDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"query_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSnmpUser resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchSnmpUser(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSnmpUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchSnmpUserRead(d, m)
}

func resourceSwitchControllerManagedSwitchSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpUser resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchSnmpUser(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerManagedSwitchSnmpUserRead(d, m)
}

func resourceSwitchControllerManagedSwitchSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchSnmpUser(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSnmpUserAuthProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserPrivProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueries2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueryPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_proto", flattenSwitchControllerManagedSwitchSnmpUserAuthProto2edl(o["auth-proto"], d, "auth_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-proto"], "SwitchControllerManagedSwitchSnmpUser-AuthProto"); ok {
			if err = d.Set("auth_proto", vv); err != nil {
				return fmt.Errorf("Error reading auth_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchSnmpUserName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitchSnmpUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenSwitchControllerManagedSwitchSnmpUserPrivProto2edl(o["priv-proto"], d, "priv_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-proto"], "SwitchControllerManagedSwitchSnmpUser-PrivProto"); ok {
			if err = d.Set("priv_proto", vv); err != nil {
				return fmt.Errorf("Error reading priv_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("queries", flattenSwitchControllerManagedSwitchSnmpUserQueries2edl(o["queries"], d, "queries")); err != nil {
		if vv, ok := fortiAPIPatch(o["queries"], "SwitchControllerManagedSwitchSnmpUser-Queries"); ok {
			if err = d.Set("queries", vv); err != nil {
				return fmt.Errorf("Error reading queries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", flattenSwitchControllerManagedSwitchSnmpUserQueryPort2edl(o["query-port"], d, "query_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-port"], "SwitchControllerManagedSwitchSnmpUser-QueryPort"); ok {
			if err = d.Set("query_port", vv); err != nil {
				return fmt.Errorf("Error reading query_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("security_level", flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel2edl(o["security-level"], d, "security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-level"], "SwitchControllerManagedSwitchSnmpUser-SecurityLevel"); ok {
			if err = d.Set("security_level", vv); err != nil {
				return fmt.Errorf("Error reading security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSnmpUserAuthProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserAuthPwd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpUserName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivPwd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueryPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserSecurityLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_proto"); ok || d.HasChange("auth_proto") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserAuthProto2edl(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok || d.HasChange("auth_pwd") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserAuthPwd2edl(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok || d.HasChange("priv_proto") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserPrivProto2edl(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok || d.HasChange("priv_pwd") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserPrivPwd2edl(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok || d.HasChange("queries") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserQueries2edl(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("query_port"); ok || d.HasChange("query_port") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserQueryPort2edl(d, v, "query_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok || d.HasChange("security_level") {
		t, err := expandSwitchControllerManagedSwitchSnmpUserSecurityLevel2edl(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	return &obj, nil
}
