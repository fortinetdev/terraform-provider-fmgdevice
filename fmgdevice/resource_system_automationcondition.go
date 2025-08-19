// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Condition for automation stitches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationConditionCreate,
		Read:   resourceSystemAutomationConditionRead,
		Update: resourceSystemAutomationConditionUpdate,
		Delete: resourceSystemAutomationConditionDelete,

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
			"condition_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_usage_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"input_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"input_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mem_usage_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vpn_tunnel_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_tunnel_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutomationConditionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationCondition(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationCondition resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationCondition(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationCondition resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationConditionRead(d, m)
}

func resourceSystemAutomationConditionUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationCondition(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationCondition resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationCondition(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationCondition resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationConditionRead(d, m)
}

func resourceSystemAutomationConditionDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemAutomationCondition(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationConditionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemAutomationCondition(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationCondition(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationCondition resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationConditionConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionCpuUsagePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionInputId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionInputState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionMemUsagePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationConditionVpnTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationConditionVpnTunnelState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationCondition(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("condition_type", flattenSystemAutomationConditionConditionType(o["condition-type"], d, "condition_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["condition-type"], "SystemAutomationCondition-ConditionType"); ok {
			if err = d.Set("condition_type", vv); err != nil {
				return fmt.Errorf("Error reading condition_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading condition_type: %v", err)
		}
	}

	if err = d.Set("cpu_usage_percent", flattenSystemAutomationConditionCpuUsagePercent(o["cpu-usage-percent"], d, "cpu_usage_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu-usage-percent"], "SystemAutomationCondition-CpuUsagePercent"); ok {
			if err = d.Set("cpu_usage_percent", vv); err != nil {
				return fmt.Errorf("Error reading cpu_usage_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu_usage_percent: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationConditionDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAutomationCondition-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("input_id", flattenSystemAutomationConditionInputId(o["input-id"], d, "input_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-id"], "SystemAutomationCondition-InputId"); ok {
			if err = d.Set("input_id", vv); err != nil {
				return fmt.Errorf("Error reading input_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_id: %v", err)
		}
	}

	if err = d.Set("input_state", flattenSystemAutomationConditionInputState(o["input-state"], d, "input_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-state"], "SystemAutomationCondition-InputState"); ok {
			if err = d.Set("input_state", vv); err != nil {
				return fmt.Errorf("Error reading input_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_state: %v", err)
		}
	}

	if err = d.Set("mem_usage_percent", flattenSystemAutomationConditionMemUsagePercent(o["mem-usage-percent"], d, "mem_usage_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["mem-usage-percent"], "SystemAutomationCondition-MemUsagePercent"); ok {
			if err = d.Set("mem_usage_percent", vv); err != nil {
				return fmt.Errorf("Error reading mem_usage_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mem_usage_percent: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationConditionName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationCondition-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemAutomationConditionVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemAutomationCondition-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vpn_tunnel_name", flattenSystemAutomationConditionVpnTunnelName(o["vpn-tunnel-name"], d, "vpn_tunnel_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-tunnel-name"], "SystemAutomationCondition-VpnTunnelName"); ok {
			if err = d.Set("vpn_tunnel_name", vv); err != nil {
				return fmt.Errorf("Error reading vpn_tunnel_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_tunnel_name: %v", err)
		}
	}

	if err = d.Set("vpn_tunnel_state", flattenSystemAutomationConditionVpnTunnelState(o["vpn-tunnel-state"], d, "vpn_tunnel_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-tunnel-state"], "SystemAutomationCondition-VpnTunnelState"); ok {
			if err = d.Set("vpn_tunnel_state", vv); err != nil {
				return fmt.Errorf("Error reading vpn_tunnel_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_tunnel_state: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationConditionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationConditionConditionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionCpuUsagePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionInputId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionInputState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionMemUsagePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationConditionVpnTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionVpnTunnelState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationCondition(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("condition_type"); ok || d.HasChange("condition_type") {
		t, err := expandSystemAutomationConditionConditionType(d, v, "condition_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition-type"] = t
		}
	}

	if v, ok := d.GetOk("cpu_usage_percent"); ok || d.HasChange("cpu_usage_percent") {
		t, err := expandSystemAutomationConditionCpuUsagePercent(d, v, "cpu_usage_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-usage-percent"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAutomationConditionDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("input_id"); ok || d.HasChange("input_id") {
		t, err := expandSystemAutomationConditionInputId(d, v, "input_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-id"] = t
		}
	}

	if v, ok := d.GetOk("input_state"); ok || d.HasChange("input_state") {
		t, err := expandSystemAutomationConditionInputState(d, v, "input_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-state"] = t
		}
	}

	if v, ok := d.GetOk("mem_usage_percent"); ok || d.HasChange("mem_usage_percent") {
		t, err := expandSystemAutomationConditionMemUsagePercent(d, v, "mem_usage_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mem-usage-percent"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationConditionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemAutomationConditionVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vpn_tunnel_name"); ok || d.HasChange("vpn_tunnel_name") {
		t, err := expandSystemAutomationConditionVpnTunnelName(d, v, "vpn_tunnel_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-tunnel-name"] = t
		}
	}

	if v, ok := d.GetOk("vpn_tunnel_state"); ok || d.HasChange("vpn_tunnel_state") {
		t, err := expandSystemAutomationConditionVpnTunnelState(d, v, "vpn_tunnel_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-tunnel-state"] = t
		}
	}

	return &obj, nil
}
