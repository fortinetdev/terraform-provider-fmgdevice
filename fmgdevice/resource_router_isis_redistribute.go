// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IS-IS redistribute protocols.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsisRedistribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisRedistributeUpdate,
		Read:   resourceRouterIsisRedistributeRead,
		Update: resourceRouterIsisRedistributeUpdate,
		Delete: resourceRouterIsisRedistributeDelete,

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

func resourceRouterIsisRedistributeUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterIsisRedistribute(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisRedistribute resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsisRedistribute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisRedistribute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterIsisRedistribute")

	return resourceRouterIsisRedistributeRead(d, m)
}

func resourceRouterIsisRedistributeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterIsisRedistribute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterIsisRedistribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisRedistributeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterIsisRedistribute(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisRedistribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsisRedistribute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisRedistribute resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisRedistributeLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistributeStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterIsisRedistribute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("level", flattenRouterIsisRedistributeLevel2edl(o["level"], d, "level")); err != nil {
		if vv, ok := fortiAPIPatch(o["level"], "RouterIsisRedistribute-Level"); ok {
			if err = d.Set("level", vv); err != nil {
				return fmt.Errorf("Error reading level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("metric", flattenRouterIsisRedistributeMetric2edl(o["metric"], d, "metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric"], "RouterIsisRedistribute-Metric"); ok {
			if err = d.Set("metric", vv); err != nil {
				return fmt.Errorf("Error reading metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric: %v", err)
		}
	}

	if err = d.Set("metric_type", flattenRouterIsisRedistributeMetricType2edl(o["metric-type"], d, "metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric-type"], "RouterIsisRedistribute-MetricType"); ok {
			if err = d.Set("metric_type", vv); err != nil {
				return fmt.Errorf("Error reading metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric_type: %v", err)
		}
	}

	if err = d.Set("protocol", flattenRouterIsisRedistributeProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "RouterIsisRedistribute-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("routemap", flattenRouterIsisRedistributeRoutemap2edl(o["routemap"], d, "routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["routemap"], "RouterIsisRedistribute-Routemap"); ok {
			if err = d.Set("routemap", vv); err != nil {
				return fmt.Errorf("Error reading routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routemap: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterIsisRedistributeStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterIsisRedistribute-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenRouterIsisRedistributeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterIsisRedistributeLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistributeStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsisRedistribute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandRouterIsisRedistributeLevel2edl(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("metric"); ok || d.HasChange("metric") {
		t, err := expandRouterIsisRedistributeMetric2edl(d, v, "metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric"] = t
		}
	}

	if v, ok := d.GetOk("metric_type"); ok || d.HasChange("metric_type") {
		t, err := expandRouterIsisRedistributeMetricType2edl(d, v, "metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-type"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandRouterIsisRedistributeProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("routemap"); ok || d.HasChange("routemap") {
		t, err := expandRouterIsisRedistributeRoutemap2edl(d, v, "routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routemap"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterIsisRedistributeStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
