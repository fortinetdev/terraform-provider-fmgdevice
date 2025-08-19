// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Policy definitions which can define the behavior on auto configured interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigPolicyCreate,
		Read:   resourceSwitchControllerAutoConfigPolicyRead,
		Update: resourceSwitchControllerAutoConfigPolicyUpdate,
		Delete: resourceSwitchControllerAutoConfigPolicyDelete,

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
			"igmp_flood_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_flood_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"poe_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"storm_control_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerAutoConfigPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerAutoConfigPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerAutoConfigPolicyRead(d, m)
}

func resourceSwitchControllerAutoConfigPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerAutoConfigPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAutoConfigPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerAutoConfigPolicyRead(d, m)
}

func resourceSwitchControllerAutoConfigPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerAutoConfigPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerAutoConfigPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigPolicyIgmpFloodReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyIgmpFloodTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyPoeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAutoConfigPolicyStormControlPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("igmp_flood_report", flattenSwitchControllerAutoConfigPolicyIgmpFloodReport(o["igmp-flood-report"], d, "igmp_flood_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-flood-report"], "SwitchControllerAutoConfigPolicy-IgmpFloodReport"); ok {
			if err = d.Set("igmp_flood_report", vv); err != nil {
				return fmt.Errorf("Error reading igmp_flood_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_flood_report: %v", err)
		}
	}

	if err = d.Set("igmp_flood_traffic", flattenSwitchControllerAutoConfigPolicyIgmpFloodTraffic(o["igmp-flood-traffic"], d, "igmp_flood_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-flood-traffic"], "SwitchControllerAutoConfigPolicy-IgmpFloodTraffic"); ok {
			if err = d.Set("igmp_flood_traffic", vv); err != nil {
				return fmt.Errorf("Error reading igmp_flood_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_flood_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerAutoConfigPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerAutoConfigPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("poe_status", flattenSwitchControllerAutoConfigPolicyPoeStatus(o["poe-status"], d, "poe_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-status"], "SwitchControllerAutoConfigPolicy-PoeStatus"); ok {
			if err = d.Set("poe_status", vv); err != nil {
				return fmt.Errorf("Error reading poe_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_status: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchControllerAutoConfigPolicyQosPolicy(o["qos-policy"], d, "qos_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-policy"], "SwitchControllerAutoConfigPolicy-QosPolicy"); ok {
			if err = d.Set("qos_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("storm_control_policy", flattenSwitchControllerAutoConfigPolicyStormControlPolicy(o["storm-control-policy"], d, "storm_control_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["storm-control-policy"], "SwitchControllerAutoConfigPolicy-StormControlPolicy"); ok {
			if err = d.Set("storm_control_policy", vv); err != nil {
				return fmt.Errorf("Error reading storm_control_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading storm_control_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigPolicyIgmpFloodReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyIgmpFloodTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyPoeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerAutoConfigPolicyStormControlPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("igmp_flood_report"); ok || d.HasChange("igmp_flood_report") {
		t, err := expandSwitchControllerAutoConfigPolicyIgmpFloodReport(d, v, "igmp_flood_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-flood-report"] = t
		}
	}

	if v, ok := d.GetOk("igmp_flood_traffic"); ok || d.HasChange("igmp_flood_traffic") {
		t, err := expandSwitchControllerAutoConfigPolicyIgmpFloodTraffic(d, v, "igmp_flood_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-flood-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerAutoConfigPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("poe_status"); ok || d.HasChange("poe_status") {
		t, err := expandSwitchControllerAutoConfigPolicyPoeStatus(d, v, "poe_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-status"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok || d.HasChange("qos_policy") {
		t, err := expandSwitchControllerAutoConfigPolicyQosPolicy(d, v, "qos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("storm_control_policy"); ok || d.HasChange("storm_control_policy") {
		t, err := expandSwitchControllerAutoConfigPolicyStormControlPolicy(d, v, "storm_control_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control-policy"] = t
		}
	}

	return &obj, nil
}
