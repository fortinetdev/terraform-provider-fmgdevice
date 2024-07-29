// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP to MAC binding settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIpmacbindingSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpmacbindingSettingUpdate,
		Read:   resourceFirewallIpmacbindingSettingRead,
		Update: resourceFirewallIpmacbindingSettingUpdate,
		Delete: resourceFirewallIpmacbindingSettingDelete,

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
			"bindthroughfw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bindtofw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefinedhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallIpmacbindingSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIpmacbindingSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallIpmacbindingSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallIpmacbindingSetting")

	return resourceFirewallIpmacbindingSettingRead(d, m)
}

func resourceFirewallIpmacbindingSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallIpmacbindingSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpmacbindingSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpmacbindingSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIpmacbindingSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpmacbindingSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpmacbindingSettingBindthroughfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingBindtofw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingUndefinedhost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIpmacbindingSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bindthroughfw", flattenFirewallIpmacbindingSettingBindthroughfw(o["bindthroughfw"], d, "bindthroughfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["bindthroughfw"], "FirewallIpmacbindingSetting-Bindthroughfw"); ok {
			if err = d.Set("bindthroughfw", vv); err != nil {
				return fmt.Errorf("Error reading bindthroughfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bindthroughfw: %v", err)
		}
	}

	if err = d.Set("bindtofw", flattenFirewallIpmacbindingSettingBindtofw(o["bindtofw"], d, "bindtofw")); err != nil {
		if vv, ok := fortiAPIPatch(o["bindtofw"], "FirewallIpmacbindingSetting-Bindtofw"); ok {
			if err = d.Set("bindtofw", vv); err != nil {
				return fmt.Errorf("Error reading bindtofw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bindtofw: %v", err)
		}
	}

	if err = d.Set("undefinedhost", flattenFirewallIpmacbindingSettingUndefinedhost(o["undefinedhost"], d, "undefinedhost")); err != nil {
		if vv, ok := fortiAPIPatch(o["undefinedhost"], "FirewallIpmacbindingSetting-Undefinedhost"); ok {
			if err = d.Set("undefinedhost", vv); err != nil {
				return fmt.Errorf("Error reading undefinedhost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading undefinedhost: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpmacbindingSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIpmacbindingSettingBindthroughfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingBindtofw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingUndefinedhost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpmacbindingSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bindthroughfw"); ok || d.HasChange("bindthroughfw") {
		t, err := expandFirewallIpmacbindingSettingBindthroughfw(d, v, "bindthroughfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bindthroughfw"] = t
		}
	}

	if v, ok := d.GetOk("bindtofw"); ok || d.HasChange("bindtofw") {
		t, err := expandFirewallIpmacbindingSettingBindtofw(d, v, "bindtofw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bindtofw"] = t
		}
	}

	if v, ok := d.GetOk("undefinedhost"); ok || d.HasChange("undefinedhost") {
		t, err := expandFirewallIpmacbindingSettingUndefinedhost(d, v, "undefinedhost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["undefinedhost"] = t
		}
	}

	return &obj, nil
}
