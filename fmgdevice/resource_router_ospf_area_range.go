// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF area range configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfAreaRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfAreaRangeCreate,
		Read:   resourceRouterOspfAreaRangeRead,
		Update: resourceRouterOspfAreaRangeUpdate,
		Delete: resourceRouterOspfAreaRangeDelete,

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
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"substitute": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"substitute_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterOspfAreaRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfAreaRange(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaRange resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfAreaRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfAreaRangeRead(d, m)
}

func resourceRouterOspfAreaRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfAreaRange(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaRange resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfAreaRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfAreaRangeRead(d, m)
}

func resourceRouterOspfAreaRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfAreaRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfAreaRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfAreaRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfAreaRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfAreaRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaRange resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAreaRangeAdvertise3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangePrefix3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstitute3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstituteStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfAreaRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advertise", flattenRouterOspfAreaRangeAdvertise3rdl(o["advertise"], d, "advertise")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertise"], "RouterOspfAreaRange-Advertise"); ok {
			if err = d.Set("advertise", vv); err != nil {
				return fmt.Errorf("Error reading advertise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertise: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterOspfAreaRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspfAreaRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterOspfAreaRangePrefix3rdl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterOspfAreaRange-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("substitute", flattenRouterOspfAreaRangeSubstitute3rdl(o["substitute"], d, "substitute")); err != nil {
		if vv, ok := fortiAPIPatch(o["substitute"], "RouterOspfAreaRange-Substitute"); ok {
			if err = d.Set("substitute", vv); err != nil {
				return fmt.Errorf("Error reading substitute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading substitute: %v", err)
		}
	}

	if err = d.Set("substitute_status", flattenRouterOspfAreaRangeSubstituteStatus3rdl(o["substitute-status"], d, "substitute_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["substitute-status"], "RouterOspfAreaRange-SubstituteStatus"); ok {
			if err = d.Set("substitute_status", vv); err != nil {
				return fmt.Errorf("Error reading substitute_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading substitute_status: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfAreaRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAreaRangeAdvertise3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangePrefix3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstitute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstituteStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfAreaRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertise"); ok || d.HasChange("advertise") {
		t, err := expandRouterOspfAreaRangeAdvertise3rdl(d, v, "advertise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertise"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspfAreaRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterOspfAreaRangePrefix3rdl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("substitute"); ok || d.HasChange("substitute") {
		t, err := expandRouterOspfAreaRangeSubstitute3rdl(d, v, "substitute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["substitute"] = t
		}
	}

	if v, ok := d.GetOk("substitute_status"); ok || d.HasChange("substitute_status") {
		t, err := expandRouterOspfAreaRangeSubstituteStatus3rdl(d, v, "substitute_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["substitute-status"] = t
		}
	}

	return &obj, nil
}
