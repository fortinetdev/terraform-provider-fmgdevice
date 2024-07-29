// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerVlanPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerVlanPolicyCreate,
		Read:   resourceSwitchControllerVlanPolicyRead,
		Update: resourceSwitchControllerVlanPolicyUpdate,
		Delete: resourceSwitchControllerVlanPolicyDelete,

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
			"allowed_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"allowed_vlans_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discard_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink": &schema.Schema{
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
			"untagged_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerVlanPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerVlanPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlanPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerVlanPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerVlanPolicyRead(d, m)
}

func resourceSwitchControllerVlanPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerVlanPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlanPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerVlanPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlanPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerVlanPolicyRead(d, m)
}

func resourceSwitchControllerVlanPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerVlanPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerVlanPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerVlanPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlanPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerVlanPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlanPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerVlanPolicyAllowedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerVlanPolicyAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyDiscardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerVlanPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerVlanPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerVlanPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowed_vlans", flattenSwitchControllerVlanPolicyAllowedVlans(o["allowed-vlans"], d, "allowed_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans"], "SwitchControllerVlanPolicy-AllowedVlans"); ok {
			if err = d.Set("allowed_vlans", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans: %v", err)
		}
	}

	if err = d.Set("allowed_vlans_all", flattenSwitchControllerVlanPolicyAllowedVlansAll(o["allowed-vlans-all"], d, "allowed_vlans_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans-all"], "SwitchControllerVlanPolicy-AllowedVlansAll"); ok {
			if err = d.Set("allowed_vlans_all", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerVlanPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerVlanPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("discard_mode", flattenSwitchControllerVlanPolicyDiscardMode(o["discard-mode"], d, "discard_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["discard-mode"], "SwitchControllerVlanPolicy-DiscardMode"); ok {
			if err = d.Set("discard_mode", vv); err != nil {
				return fmt.Errorf("Error reading discard_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discard_mode: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerVlanPolicyFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "SwitchControllerVlanPolicy-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerVlanPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerVlanPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("untagged_vlans", flattenSwitchControllerVlanPolicyUntaggedVlans(o["untagged-vlans"], d, "untagged_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["untagged-vlans"], "SwitchControllerVlanPolicy-UntaggedVlans"); ok {
			if err = d.Set("untagged_vlans", vv); err != nil {
				return fmt.Errorf("Error reading untagged_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untagged_vlans: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerVlanPolicyVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerVlanPolicy-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerVlanPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerVlanPolicyAllowedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerVlanPolicyAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyDiscardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerVlanPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerVlanPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerVlanPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowed_vlans"); ok || d.HasChange("allowed_vlans") {
		t, err := expandSwitchControllerVlanPolicyAllowedVlans(d, v, "allowed_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans_all"); ok || d.HasChange("allowed_vlans_all") {
		t, err := expandSwitchControllerVlanPolicyAllowedVlansAll(d, v, "allowed_vlans_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans-all"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerVlanPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("discard_mode"); ok || d.HasChange("discard_mode") {
		t, err := expandSwitchControllerVlanPolicyDiscardMode(d, v, "discard_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discard-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandSwitchControllerVlanPolicyFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerVlanPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("untagged_vlans"); ok || d.HasChange("untagged_vlans") {
		t, err := expandSwitchControllerVlanPolicyUntaggedVlans(d, v, "untagged_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerVlanPolicyVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
