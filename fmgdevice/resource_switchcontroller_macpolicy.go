// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerMacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerMacPolicyCreate,
		Read:   resourceSwitchControllerMacPolicyRead,
		Update: resourceSwitchControllerMacPolicyUpdate,
		Delete: resourceSwitchControllerMacPolicyDelete,

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
			"bounce_port_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bounce_port_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmgcount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"poe_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_policy": &schema.Schema{
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

func resourceSwitchControllerMacPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerMacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerMacPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerMacPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerMacPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerMacPolicyRead(d, m)
}

func resourceSwitchControllerMacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerMacPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerMacPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerMacPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerMacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerMacPolicyRead(d, m)
}

func resourceSwitchControllerMacPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerMacPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerMacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerMacPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerMacPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerMacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerMacPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerMacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerMacPolicyBouncePortDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyBouncePortLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerMacPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyPoeReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerMacPolicyTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerMacPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerMacPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bounce_port_duration", flattenSwitchControllerMacPolicyBouncePortDuration(o["bounce-port-duration"], d, "bounce_port_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-port-duration"], "SwitchControllerMacPolicy-BouncePortDuration"); ok {
			if err = d.Set("bounce_port_duration", vv); err != nil {
				return fmt.Errorf("Error reading bounce_port_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_port_duration: %v", err)
		}
	}

	if err = d.Set("bounce_port_link", flattenSwitchControllerMacPolicyBouncePortLink(o["bounce-port-link"], d, "bounce_port_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-port-link"], "SwitchControllerMacPolicy-BouncePortLink"); ok {
			if err = d.Set("bounce_port_link", vv); err != nil {
				return fmt.Errorf("Error reading bounce_port_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_port_link: %v", err)
		}
	}

	if err = d.Set("fmgcount", flattenSwitchControllerMacPolicyCount(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "SwitchControllerMacPolicy-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerMacPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerMacPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("drop", flattenSwitchControllerMacPolicyDrop(o["drop"], d, "drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop"], "SwitchControllerMacPolicy-Drop"); ok {
			if err = d.Set("drop", vv); err != nil {
				return fmt.Errorf("Error reading drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerMacPolicyFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "SwitchControllerMacPolicy-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerMacPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerMacPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("poe_reset", flattenSwitchControllerMacPolicyPoeReset(o["poe-reset"], d, "poe_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-reset"], "SwitchControllerMacPolicy-PoeReset"); ok {
			if err = d.Set("poe_reset", vv); err != nil {
				return fmt.Errorf("Error reading poe_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_reset: %v", err)
		}
	}

	if err = d.Set("traffic_policy", flattenSwitchControllerMacPolicyTrafficPolicy(o["traffic-policy"], d, "traffic_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-policy"], "SwitchControllerMacPolicy-TrafficPolicy"); ok {
			if err = d.Set("traffic_policy", vv); err != nil {
				return fmt.Errorf("Error reading traffic_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_policy: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerMacPolicyVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerMacPolicy-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerMacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerMacPolicyBouncePortDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyBouncePortLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerMacPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyPoeReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerMacPolicyTrafficPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerMacPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerMacPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bounce_port_duration"); ok || d.HasChange("bounce_port_duration") {
		t, err := expandSwitchControllerMacPolicyBouncePortDuration(d, v, "bounce_port_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-duration"] = t
		}
	}

	if v, ok := d.GetOk("bounce_port_link"); ok || d.HasChange("bounce_port_link") {
		t, err := expandSwitchControllerMacPolicyBouncePortLink(d, v, "bounce_port_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-link"] = t
		}
	}

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandSwitchControllerMacPolicyCount(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerMacPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("drop"); ok || d.HasChange("drop") {
		t, err := expandSwitchControllerMacPolicyDrop(d, v, "drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandSwitchControllerMacPolicyFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerMacPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("poe_reset"); ok || d.HasChange("poe_reset") {
		t, err := expandSwitchControllerMacPolicyPoeReset(d, v, "poe_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-reset"] = t
		}
	}

	if v, ok := d.GetOk("traffic_policy"); ok || d.HasChange("traffic_policy") {
		t, err := expandSwitchControllerMacPolicyTrafficPolicy(d, v, "traffic_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-policy"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerMacPolicyVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
