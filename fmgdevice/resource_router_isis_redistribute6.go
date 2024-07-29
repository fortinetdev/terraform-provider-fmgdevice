// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IS-IS IPv6 redistribution for routing protocols.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsisRedistribute6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisRedistribute6Update,
		Read:   resourceRouterIsisRedistribute6Read,
		Update: resourceRouterIsisRedistribute6Update,
		Delete: resourceRouterIsisRedistribute6Delete,

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
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"routemap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterIsisRedistribute6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterIsisRedistribute6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisRedistribute6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsisRedistribute6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisRedistribute6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterIsisRedistribute6")

	return resourceRouterIsisRedistribute6Read(d, m)
}

func resourceRouterIsisRedistribute6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterIsisRedistribute6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterIsisRedistribute6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisRedistribute6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterIsisRedistribute6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisRedistribute6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsisRedistribute6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisRedistribute6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisRedistribute6Level2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Metric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6MetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Protocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Routemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistribute6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterIsisRedistribute6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("level", flattenRouterIsisRedistribute6Level2edl(o["level"], d, "level")); err != nil {
		if vv, ok := fortiAPIPatch(o["level"], "RouterIsisRedistribute6-Level"); ok {
			if err = d.Set("level", vv); err != nil {
				return fmt.Errorf("Error reading level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("metric", flattenRouterIsisRedistribute6Metric2edl(o["metric"], d, "metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric"], "RouterIsisRedistribute6-Metric"); ok {
			if err = d.Set("metric", vv); err != nil {
				return fmt.Errorf("Error reading metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric: %v", err)
		}
	}

	if err = d.Set("metric_type", flattenRouterIsisRedistribute6MetricType2edl(o["metric-type"], d, "metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric-type"], "RouterIsisRedistribute6-MetricType"); ok {
			if err = d.Set("metric_type", vv); err != nil {
				return fmt.Errorf("Error reading metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric_type: %v", err)
		}
	}

	if err = d.Set("protocol", flattenRouterIsisRedistribute6Protocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "RouterIsisRedistribute6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("routemap", flattenRouterIsisRedistribute6Routemap2edl(o["routemap"], d, "routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["routemap"], "RouterIsisRedistribute6-Routemap"); ok {
			if err = d.Set("routemap", vv); err != nil {
				return fmt.Errorf("Error reading routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routemap: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterIsisRedistribute6Status2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterIsisRedistribute6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenRouterIsisRedistribute6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterIsisRedistribute6Level2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Metric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6MetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Protocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Routemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistribute6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsisRedistribute6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandRouterIsisRedistribute6Level2edl(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("metric"); ok || d.HasChange("metric") {
		t, err := expandRouterIsisRedistribute6Metric2edl(d, v, "metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric"] = t
		}
	}

	if v, ok := d.GetOk("metric_type"); ok || d.HasChange("metric_type") {
		t, err := expandRouterIsisRedistribute6MetricType2edl(d, v, "metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-type"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandRouterIsisRedistribute6Protocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("routemap"); ok || d.HasChange("routemap") {
		t, err := expandRouterIsisRedistribute6Routemap2edl(d, v, "routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routemap"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterIsisRedistribute6Status2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
