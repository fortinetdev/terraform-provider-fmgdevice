// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicyLocalAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicyLocalAccessCreate,
		Read:   resourceSwitchControllerSecurityPolicyLocalAccessRead,
		Update: resourceSwitchControllerSecurityPolicyLocalAccessUpdate,
		Delete: resourceSwitchControllerSecurityPolicyLocalAccessDelete,

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
			"internal_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mgmt_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicyLocalAccessCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyLocalAccess(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSecurityPolicyLocalAccess(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceSwitchControllerSecurityPolicyLocalAccessUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyLocalAccess(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSecurityPolicyLocalAccess(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceSwitchControllerSecurityPolicyLocalAccessDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSecurityPolicyLocalAccess(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicyLocalAccessRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSecurityPolicyLocalAccess(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicyLocalAccess(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyLocalAccess resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicyLocalAccessName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("internal_allowaccess", flattenSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(o["internal-allowaccess"], d, "internal_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-allowaccess"], "SwitchControllerSecurityPolicyLocalAccess-InternalAllowaccess"); ok {
			if err = d.Set("internal_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading internal_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_allowaccess: %v", err)
		}
	}

	if err = d.Set("mgmt_allowaccess", flattenSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(o["mgmt-allowaccess"], d, "mgmt_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt-allowaccess"], "SwitchControllerSecurityPolicyLocalAccess-MgmtAllowaccess"); ok {
			if err = d.Set("mgmt_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_allowaccess: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSecurityPolicyLocalAccessName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSecurityPolicyLocalAccess-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicyLocalAccessFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicyLocalAccessName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("internal_allowaccess"); ok || d.HasChange("internal_allowaccess") {
		t, err := expandSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d, v, "internal_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_allowaccess"); ok || d.HasChange("mgmt_allowaccess") {
		t, err := expandSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d, v, "mgmt_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSecurityPolicyLocalAccessName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
