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

func resourceLogFortianalyzer3FilterFreeStyle() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer3FilterFreeStyleCreate,
		Read:   resourceLogFortianalyzer3FilterFreeStyleRead,
		Update: resourceLogFortianalyzer3FilterFreeStyleUpdate,
		Delete: resourceLogFortianalyzer3FilterFreeStyleDelete,

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

func resourceLogFortianalyzer3FilterFreeStyleCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortianalyzer3FilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error creating LogFortianalyzer3FilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.CreateLogFortianalyzer3FilterFreeStyle(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogFortianalyzer3FilterFreeStyle resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogFortianalyzer3FilterFreeStyleRead(d, m)
}

func resourceLogFortianalyzer3FilterFreeStyleUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogFortianalyzer3FilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3FilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer3FilterFreeStyle(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3FilterFreeStyle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogFortianalyzer3FilterFreeStyleRead(d, m)
}

func resourceLogFortianalyzer3FilterFreeStyleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogFortianalyzer3FilterFreeStyle(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortianalyzer3FilterFreeStyle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer3FilterFreeStyleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortianalyzer3FilterFreeStyle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3FilterFreeStyle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3FilterFreeStyle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3FilterFreeStyle resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3FilterFreeStyleCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyleFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyleFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3FilterFreeStyle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenLogFortianalyzer3FilterFreeStyleCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "LogFortianalyzer3FilterFreeStyle-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzer3FilterFreeStyleFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogFortianalyzer3FilterFreeStyle-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzer3FilterFreeStyleFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogFortianalyzer3FilterFreeStyle-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogFortianalyzer3FilterFreeStyleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogFortianalyzer3FilterFreeStyle-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3FilterFreeStyleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortianalyzer3FilterFreeStyleCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3FilterFreeStyle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandLogFortianalyzer3FilterFreeStyleCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogFortianalyzer3FilterFreeStyleFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogFortianalyzer3FilterFreeStyleFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogFortianalyzer3FilterFreeStyleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
