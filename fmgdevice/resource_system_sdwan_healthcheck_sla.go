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

func resourceSystemSdwanHealthCheckSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanHealthCheckSlaCreate,
		Read:   resourceSystemSdwanHealthCheckSlaRead,
		Update: resourceSystemSdwanHealthCheckSlaUpdate,
		Delete: resourceSystemSdwanHealthCheckSlaDelete,

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
			"health_check": &schema.Schema{
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

func resourceSystemSdwanHealthCheckSlaCreate(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check"] = health_check

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSdwanHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckSla resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanHealthCheckSla(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckSla resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanHealthCheckSlaRead(d, m)
}

func resourceSystemSdwanHealthCheckSlaUpdate(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check"] = health_check

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSdwanHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckSla resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanHealthCheckSla(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSdwanHealthCheckSlaRead(d, m)
}

func resourceSystemSdwanHealthCheckSlaDelete(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check"] = health_check

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSdwanHealthCheckSla(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanHealthCheckSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanHealthCheckSlaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	health_check := d.Get("health_check").(string)
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
	if health_check == "" {
		health_check = importOptionChecking(m.(*FortiClient).Cfg, "health_check")
		if health_check == "" {
			return fmt.Errorf("Parameter health_check is missing")
		}
		if err = d.Set("health_check", health_check); err != nil {
			return fmt.Errorf("Error set params health_check: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["health_check"] = health_check

	o, err := c.ReadSystemSdwanHealthCheckSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanHealthCheckSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckSla resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanHealthCheckSlaId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaJitterThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLatencyThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLinkCostFactor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckSlaMosThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityInSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityOutSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanHealthCheckSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSdwanHealthCheckSlaId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSdwanHealthCheckSla-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenSystemSdwanHealthCheckSlaJitterThreshold3rdl(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "SystemSdwanHealthCheckSla-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenSystemSdwanHealthCheckSlaLatencyThreshold3rdl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "SystemSdwanHealthCheckSla-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenSystemSdwanHealthCheckSlaLinkCostFactor3rdl(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "SystemSdwanHealthCheckSla-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("mos_threshold", flattenSystemSdwanHealthCheckSlaMosThreshold3rdl(o["mos-threshold"], d, "mos_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-threshold"], "SystemSdwanHealthCheckSla-MosThreshold"); ok {
			if err = d.Set("mos_threshold", vv); err != nil {
				return fmt.Errorf("Error reading mos_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_threshold: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "SystemSdwanHealthCheckSla-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	if err = d.Set("priority_in_sla", flattenSystemSdwanHealthCheckSlaPriorityInSla3rdl(o["priority-in-sla"], d, "priority_in_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-in-sla"], "SystemSdwanHealthCheckSla-PriorityInSla"); ok {
			if err = d.Set("priority_in_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_in_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_in_sla: %v", err)
		}
	}

	if err = d.Set("priority_out_sla", flattenSystemSdwanHealthCheckSlaPriorityOutSla3rdl(o["priority-out-sla"], d, "priority_out_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-out-sla"], "SystemSdwanHealthCheckSla-PriorityOutSla"); ok {
			if err = d.Set("priority_out_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_out_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_out_sla: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanHealthCheckSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanHealthCheckSlaId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaJitterThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLatencyThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLinkCostFactor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckSlaMosThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityInSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityOutSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanHealthCheckSla(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSdwanHealthCheckSlaId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandSystemSdwanHealthCheckSlaJitterThreshold3rdl(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandSystemSdwanHealthCheckSlaLatencyThreshold3rdl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandSystemSdwanHealthCheckSlaLinkCostFactor3rdl(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("mos_threshold"); ok || d.HasChange("mos_threshold") {
		t, err := expandSystemSdwanHealthCheckSlaMosThreshold3rdl(d, v, "mos_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok || d.HasChange("packetloss_threshold") {
		t, err := expandSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("priority_in_sla"); ok || d.HasChange("priority_in_sla") {
		t, err := expandSystemSdwanHealthCheckSlaPriorityInSla3rdl(d, v, "priority_in_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-in-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority_out_sla"); ok || d.HasChange("priority_out_sla") {
		t, err := expandSystemSdwanHealthCheckSlaPriorityOutSla3rdl(d, v, "priority_out_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-out-sla"] = t
		}
	}

	return &obj, nil
}
