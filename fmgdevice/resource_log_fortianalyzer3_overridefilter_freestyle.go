// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Free style filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzer3OverrideFilterFreeStyle() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer3OverrideFilterFreeStyleCreate,
		Read:   resourceLogFortianalyzer3OverrideFilterFreeStyleRead,
		Update: resourceLogFortianalyzer3OverrideFilterFreeStyleUpdate,
		Delete: resourceLogFortianalyzer3OverrideFilterFreeStyleDelete,

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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceLogFortianalyzer3OverrideFilterFreeStyleCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortianalyzer3OverrideFilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error creating LogFortianalyzer3OverrideFilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.CreateLogFortianalyzer3OverrideFilterFreeStyle(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogFortianalyzer3OverrideFilterFreeStyle resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogFortianalyzer3OverrideFilterFreeStyleRead(d, m)
}

func resourceLogFortianalyzer3OverrideFilterFreeStyleUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortianalyzer3OverrideFilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3OverrideFilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer3OverrideFilterFreeStyle(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3OverrideFilterFreeStyle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogFortianalyzer3OverrideFilterFreeStyleRead(d, m)
}

func resourceLogFortianalyzer3OverrideFilterFreeStyleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogFortianalyzer3OverrideFilterFreeStyle(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortianalyzer3OverrideFilterFreeStyle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer3OverrideFilterFreeStyleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortianalyzer3OverrideFilterFreeStyle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideFilterFreeStyle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3OverrideFilterFreeStyle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideFilterFreeStyle resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3OverrideFilterFreeStyleCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideFilterFreeStyleFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideFilterFreeStyleFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideFilterFreeStyleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3OverrideFilterFreeStyle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenLogFortianalyzer3OverrideFilterFreeStyleCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "LogFortianalyzer3OverrideFilterFreeStyle-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzer3OverrideFilterFreeStyleFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogFortianalyzer3OverrideFilterFreeStyle-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzer3OverrideFilterFreeStyleFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogFortianalyzer3OverrideFilterFreeStyle-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogFortianalyzer3OverrideFilterFreeStyleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogFortianalyzer3OverrideFilterFreeStyle-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3OverrideFilterFreeStyleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortianalyzer3OverrideFilterFreeStyleCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideFilterFreeStyleFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideFilterFreeStyleFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideFilterFreeStyleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3OverrideFilterFreeStyle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandLogFortianalyzer3OverrideFilterFreeStyleCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogFortianalyzer3OverrideFilterFreeStyleFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogFortianalyzer3OverrideFilterFreeStyleFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogFortianalyzer3OverrideFilterFreeStyleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
