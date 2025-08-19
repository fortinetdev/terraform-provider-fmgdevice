// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Redistribute configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRipngRedistribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipngRedistributeUpdate,
		Read:   resourceRouterRipngRedistributeRead,
		Update: resourceRouterRipngRedistributeUpdate,
		Delete: resourceRouterRipngRedistributeDelete,

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
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
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

func resourceRouterRipngRedistributeUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterRipngRedistribute(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipngRedistribute resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRipngRedistribute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipngRedistribute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterRipngRedistribute")

	return resourceRouterRipngRedistributeRead(d, m)
}

func resourceRouterRipngRedistributeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterRipngRedistribute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRipngRedistribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipngRedistributeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRipngRedistribute(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipngRedistribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRipngRedistribute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipngRedistribute resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipngRedistributeMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeRoutemap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngRedistributeStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRipngRedistribute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("metric", flattenRouterRipngRedistributeMetric2edl(o["metric"], d, "metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric"], "RouterRipngRedistribute-Metric"); ok {
			if err = d.Set("metric", vv); err != nil {
				return fmt.Errorf("Error reading metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterRipngRedistributeName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterRipngRedistribute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("routemap", flattenRouterRipngRedistributeRoutemap2edl(o["routemap"], d, "routemap")); err != nil {
		if vv, ok := fortiAPIPatch(o["routemap"], "RouterRipngRedistribute-Routemap"); ok {
			if err = d.Set("routemap", vv); err != nil {
				return fmt.Errorf("Error reading routemap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routemap: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterRipngRedistributeStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterRipngRedistribute-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenRouterRipngRedistributeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRipngRedistributeMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeRoutemap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngRedistributeStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRipngRedistribute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("metric"); ok || d.HasChange("metric") {
		t, err := expandRouterRipngRedistributeMetric2edl(d, v, "metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterRipngRedistributeName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("routemap"); ok || d.HasChange("routemap") {
		t, err := expandRouterRipngRedistributeRoutemap2edl(d, v, "routemap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routemap"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterRipngRedistributeStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
