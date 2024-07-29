// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch traffic policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerTrafficPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficPolicyCreate,
		Read:   resourceSwitchControllerTrafficPolicyRead,
		Update: resourceSwitchControllerTrafficPolicyUpdate,
		Delete: resourceSwitchControllerTrafficPolicyDelete,

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
			"cos_queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"guaranteed_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policer_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerTrafficPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerTrafficPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerTrafficPolicyRead(d, m)
}

func resourceSwitchControllerTrafficPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerTrafficPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerTrafficPolicyRead(d, m)
}

func resourceSwitchControllerTrafficPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerTrafficPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerTrafficPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficPolicyCosQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyGuaranteedBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyMaximumBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyPolicerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cos_queue", flattenSwitchControllerTrafficPolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos-queue"], "SwitchControllerTrafficPolicy-CosQueue"); ok {
			if err = d.Set("cos_queue", vv); err != nil {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerTrafficPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerTrafficPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenSwitchControllerTrafficPolicyGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth"], "SwitchControllerTrafficPolicy-GuaranteedBandwidth"); ok {
			if err = d.Set("guaranteed_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("guaranteed_burst", flattenSwitchControllerTrafficPolicyGuaranteedBurst(o["guaranteed-burst"], d, "guaranteed_burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-burst"], "SwitchControllerTrafficPolicy-GuaranteedBurst"); ok {
			if err = d.Set("guaranteed_burst", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_burst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerTrafficPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerTrafficPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("maximum_burst", flattenSwitchControllerTrafficPolicyMaximumBurst(o["maximum-burst"], d, "maximum_burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-burst"], "SwitchControllerTrafficPolicy-MaximumBurst"); ok {
			if err = d.Set("maximum_burst", vv); err != nil {
				return fmt.Errorf("Error reading maximum_burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_burst: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerTrafficPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerTrafficPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policer_status", flattenSwitchControllerTrafficPolicyPolicerStatus(o["policer-status"], d, "policer_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["policer-status"], "SwitchControllerTrafficPolicy-PolicerStatus"); ok {
			if err = d.Set("policer_status", vv); err != nil {
				return fmt.Errorf("Error reading policer_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policer_status: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerTrafficPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SwitchControllerTrafficPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerTrafficPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficPolicyCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyGuaranteedBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyMaximumBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyPolicerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandSwitchControllerTrafficPolicyCosQueue(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerTrafficPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth"); ok || d.HasChange("guaranteed_bandwidth") {
		t, err := expandSwitchControllerTrafficPolicyGuaranteedBandwidth(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_burst"); ok || d.HasChange("guaranteed_burst") {
		t, err := expandSwitchControllerTrafficPolicyGuaranteedBurst(d, v, "guaranteed_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-burst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerTrafficPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("maximum_burst"); ok || d.HasChange("maximum_burst") {
		t, err := expandSwitchControllerTrafficPolicyMaximumBurst(d, v, "maximum_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-burst"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerTrafficPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policer_status"); ok || d.HasChange("policer_status") {
		t, err := expandSwitchControllerTrafficPolicyPolicerStatus(d, v, "policer_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policer-status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSwitchControllerTrafficPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
