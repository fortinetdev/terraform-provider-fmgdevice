// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch SNMP v3 users globally.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpUserCreate,
		Read:   resourceSwitchControllerSnmpUserRead,
		Update: resourceSwitchControllerSnmpUserUpdate,
		Delete: resourceSwitchControllerSnmpUserDelete,

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

func resourceSwitchControllerSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpUser resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSnmpUser(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSnmpUserRead(d, m)
}

func resourceSwitchControllerSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSwitchControllerSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpUser resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpUser(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSnmpUserRead(d, m)
}

func resourceSwitchControllerSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSwitchControllerSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpUserRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_proto", flattenSwitchControllerSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-proto"], "SwitchControllerSnmpUser-AuthProto"); ok {
			if err = d.Set("auth_proto", vv); err != nil {
				return fmt.Errorf("Error reading auth_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSnmpUserName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSnmpUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenSwitchControllerSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-proto"], "SwitchControllerSnmpUser-PrivProto"); ok {
			if err = d.Set("priv_proto", vv); err != nil {
				return fmt.Errorf("Error reading priv_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("queries", flattenSwitchControllerSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if vv, ok := fortiAPIPatch(o["queries"], "SwitchControllerSnmpUser-Queries"); ok {
			if err = d.Set("queries", vv); err != nil {
				return fmt.Errorf("Error reading queries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", flattenSwitchControllerSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-port"], "SwitchControllerSnmpUser-QueryPort"); ok {
			if err = d.Set("query_port", vv); err != nil {
				return fmt.Errorf("Error reading query_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("security_level", flattenSwitchControllerSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-level"], "SwitchControllerSnmpUser-SecurityLevel"); ok {
			if err = d.Set("security_level", vv); err != nil {
				return fmt.Errorf("Error reading security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_proto"); ok || d.HasChange("auth_proto") {
		t, err := expandSwitchControllerSnmpUserAuthProto(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok || d.HasChange("auth_pwd") {
		t, err := expandSwitchControllerSnmpUserAuthPwd(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSnmpUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok || d.HasChange("priv_proto") {
		t, err := expandSwitchControllerSnmpUserPrivProto(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok || d.HasChange("priv_pwd") {
		t, err := expandSwitchControllerSnmpUserPrivPwd(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok || d.HasChange("queries") {
		t, err := expandSwitchControllerSnmpUserQueries(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("query_port"); ok || d.HasChange("query_port") {
		t, err := expandSwitchControllerSnmpUserQueryPort(d, v, "query_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok || d.HasChange("security_level") {
		t, err := expandSwitchControllerSnmpUserSecurityLevel(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	return &obj, nil
}
