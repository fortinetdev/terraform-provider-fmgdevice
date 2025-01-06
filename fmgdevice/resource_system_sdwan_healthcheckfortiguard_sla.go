// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Service level agreement (SLA).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanHealthCheckFortiguardSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanHealthCheckFortiguardSlaCreate,
		Read:   resourceSystemSdwanHealthCheckFortiguardSlaRead,
		Update: resourceSystemSdwanHealthCheckFortiguardSlaUpdate,
		Delete: resourceSystemSdwanHealthCheckFortiguardSlaDelete,

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
			"health_check_fortiguard": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"jitter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_cost_factor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mos_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packetloss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_in_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_out_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemSdwanHealthCheckFortiguardSlaCreate(d *schema.ResourceData, m interface{}) error {
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
	health_check_fortiguard := d.Get("health_check_fortiguard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check_fortiguard"] = health_check_fortiguard

	obj, err := getObjectSystemSdwanHealthCheckFortiguardSla(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckFortiguardSla resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanHealthCheckFortiguardSla(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckFortiguardSla resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanHealthCheckFortiguardSlaRead(d, m)
}

func resourceSystemSdwanHealthCheckFortiguardSlaUpdate(d *schema.ResourceData, m interface{}) error {
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
	health_check_fortiguard := d.Get("health_check_fortiguard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check_fortiguard"] = health_check_fortiguard

	obj, err := getObjectSystemSdwanHealthCheckFortiguardSla(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckFortiguardSla resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanHealthCheckFortiguardSla(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckFortiguardSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanHealthCheckFortiguardSlaRead(d, m)
}

func resourceSystemSdwanHealthCheckFortiguardSlaDelete(d *schema.ResourceData, m interface{}) error {
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
	health_check_fortiguard := d.Get("health_check_fortiguard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check_fortiguard"] = health_check_fortiguard

	err = c.DeleteSystemSdwanHealthCheckFortiguardSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanHealthCheckFortiguardSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanHealthCheckFortiguardSlaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	health_check_fortiguard := d.Get("health_check_fortiguard").(string)
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
	if health_check_fortiguard == "" {
		health_check_fortiguard = importOptionChecking(m.(*FortiClient).Cfg, "health_check_fortiguard")
		if health_check_fortiguard == "" {
			return fmt.Errorf("Parameter health_check_fortiguard is missing")
		}
		if err = d.Set("health_check_fortiguard", health_check_fortiguard); err != nil {
			return fmt.Errorf("Error set params health_check_fortiguard: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check_fortiguard"] = health_check_fortiguard

	o, err := c.ReadSystemSdwanHealthCheckFortiguardSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckFortiguardSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanHealthCheckFortiguardSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckFortiguardSla resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanHealthCheckFortiguardSlaId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanHealthCheckFortiguardSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSdwanHealthCheckFortiguardSlaId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSdwanHealthCheckFortiguardSla-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold3rdl(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "SystemSdwanHealthCheckFortiguardSla-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold3rdl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "SystemSdwanHealthCheckFortiguardSla-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor3rdl(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "SystemSdwanHealthCheckFortiguardSla-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("mos_threshold", flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold3rdl(o["mos-threshold"], d, "mos_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-threshold"], "SystemSdwanHealthCheckFortiguardSla-MosThreshold"); ok {
			if err = d.Set("mos_threshold", vv); err != nil {
				return fmt.Errorf("Error reading mos_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_threshold: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold3rdl(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "SystemSdwanHealthCheckFortiguardSla-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	if err = d.Set("priority_in_sla", flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla3rdl(o["priority-in-sla"], d, "priority_in_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-in-sla"], "SystemSdwanHealthCheckFortiguardSla-PriorityInSla"); ok {
			if err = d.Set("priority_in_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_in_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_in_sla: %v", err)
		}
	}

	if err = d.Set("priority_out_sla", flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla3rdl(o["priority-out-sla"], d, "priority_out_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-out-sla"], "SystemSdwanHealthCheckFortiguardSla-PriorityOutSla"); ok {
			if err = d.Set("priority_out_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_out_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_out_sla: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanHealthCheckFortiguardSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanHealthCheckFortiguardSlaId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardSlaMosThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanHealthCheckFortiguardSla(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold3rdl(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold3rdl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor3rdl(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("mos_threshold"); ok || d.HasChange("mos_threshold") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaMosThreshold3rdl(d, v, "mos_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok || d.HasChange("packetloss_threshold") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold3rdl(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("priority_in_sla"); ok || d.HasChange("priority_in_sla") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla3rdl(d, v, "priority_in_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-in-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority_out_sla"); ok || d.HasChange("priority_out_sla") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla3rdl(d, v, "priority_out_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-out-sla"] = t
		}
	}

	return &obj, nil
}
