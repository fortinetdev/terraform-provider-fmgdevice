// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS egress queue policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerQosQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosQueuePolicyCreate,
		Read:   resourceSwitchControllerQosQueuePolicyRead,
		Update: resourceSwitchControllerQosQueuePolicyUpdate,
		Delete: resourceSwitchControllerQosQueuePolicyDelete,

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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"drop_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ecn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rate_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerQosQueuePolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerQosQueuePolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceSwitchControllerQosQueuePolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQosQueuePolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQueuePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceSwitchControllerQosQueuePolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerQosQueuePolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosQueuePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosQueuePolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerQosQueuePolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQueuePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosQueuePolicyCosQueue(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := i["drop-policy"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(i["drop-policy"], d, pre_append)
			tmp["drop_policy"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-DropPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := i["ecn"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueEcn(i["ecn"], d, pre_append)
			tmp["ecn"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-Ecn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := i["max-rate"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(i["max-rate"], d, pre_append)
			tmp["max_rate"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-MaxRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := i["max-rate-percent"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(i["max-rate-percent"], d, pre_append)
			tmp["max_rate_percent"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-MaxRatePercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := i["min-rate"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueMinRate(i["min-rate"], d, pre_append)
			tmp["min_rate"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-MinRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := i["min-rate-percent"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(i["min-rate-percent"], d, pre_append)
			tmp["min_rate_percent"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-MinRatePercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSwitchControllerQosQueuePolicyCosQueueWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SwitchControllerQosQueuePolicy-CosQueue-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerQosQueuePolicyCosQueueDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueEcn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyRateBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
			if vv, ok := fortiAPIPatch(o["cos-queue"], "SwitchControllerQosQueuePolicy-CosQueue"); ok {
				if err = d.Set("cos_queue", vv); err != nil {
					return fmt.Errorf("Error reading cos_queue: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cos_queue"); ok {
			if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
				if vv, ok := fortiAPIPatch(o["cos-queue"], "SwitchControllerQosQueuePolicy-CosQueue"); ok {
					if err = d.Set("cos_queue", vv); err != nil {
						return fmt.Errorf("Error reading cos_queue: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cos_queue: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSwitchControllerQosQueuePolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerQosQueuePolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rate_by", flattenSwitchControllerQosQueuePolicyRateBy(o["rate-by"], d, "rate_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-by"], "SwitchControllerQosQueuePolicy-RateBy"); ok {
			if err = d.Set("rate_by", vv); err != nil {
				return fmt.Errorf("Error reading rate_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_by: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSwitchControllerQosQueuePolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "SwitchControllerQosQueuePolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosQueuePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerQosQueuePolicyCosQueueDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["drop-policy"], _ = expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d, i["drop_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ecn"], _ = expandSwitchControllerQosQueuePolicyCosQueueEcn(d, i["ecn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d, i["max_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d, i["max_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRate(d, i["min_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d, i["min_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerQosQueuePolicyCosQueueName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSwitchControllerQosQueuePolicyCosQueueWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueEcn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyRateBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandSwitchControllerQosQueuePolicyCosQueue(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerQosQueuePolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rate_by"); ok || d.HasChange("rate_by") {
		t, err := expandSwitchControllerQosQueuePolicyRateBy(d, v, "rate_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-by"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandSwitchControllerQosQueuePolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	return &obj, nil
}
