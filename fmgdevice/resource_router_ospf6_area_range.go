// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF6 area range configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6AreaRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6AreaRangeCreate,
		Read:   resourceRouterOspf6AreaRangeRead,
		Update: resourceRouterOspf6AreaRangeUpdate,
		Delete: resourceRouterOspf6AreaRangeDelete,

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
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advertise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterOspf6AreaRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaRange(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaRange resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6AreaRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspf6AreaRangeRead(d, m)
}

func resourceRouterOspf6AreaRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaRange(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaRange resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6AreaRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspf6AreaRangeRead(d, m)
}

func resourceRouterOspf6AreaRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterOspf6AreaRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6AreaRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6AreaRangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	area := d.Get("area").(string)
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
	if area == "" {
		area = importOptionChecking(m.(*FortiClient).Cfg, "area")
		if area == "" {
			return fmt.Errorf("Parameter area is missing")
		}
		if err = d.Set("area", area); err != nil {
			return fmt.Errorf("Error set params area: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area

	o, err := c.ReadRouterOspf6AreaRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6AreaRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaRange resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AreaRangeAdvertise3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangePrefix63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6AreaRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advertise", flattenRouterOspf6AreaRangeAdvertise3rdl(o["advertise"], d, "advertise")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertise"], "RouterOspf6AreaRange-Advertise"); ok {
			if err = d.Set("advertise", vv); err != nil {
				return fmt.Errorf("Error reading advertise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertise: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterOspf6AreaRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspf6AreaRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterOspf6AreaRangePrefix63rdl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterOspf6AreaRange-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6AreaRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6AreaRangeAdvertise3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangePrefix63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6AreaRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertise"); ok || d.HasChange("advertise") {
		t, err := expandRouterOspf6AreaRangeAdvertise3rdl(d, v, "advertise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertise"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspf6AreaRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterOspf6AreaRangePrefix63rdl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
