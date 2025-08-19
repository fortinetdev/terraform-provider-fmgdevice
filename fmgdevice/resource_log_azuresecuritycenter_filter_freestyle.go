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

func resourceLogAzureSecurityCenterFilterFreeStyle() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogAzureSecurityCenterFilterFreeStyleCreate,
		Read:   resourceLogAzureSecurityCenterFilterFreeStyleRead,
		Update: resourceLogAzureSecurityCenterFilterFreeStyleUpdate,
		Delete: resourceLogAzureSecurityCenterFilterFreeStyleDelete,

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

func resourceLogAzureSecurityCenterFilterFreeStyleCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogAzureSecurityCenterFilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error creating LogAzureSecurityCenterFilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.CreateLogAzureSecurityCenterFilterFreeStyle(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LogAzureSecurityCenterFilterFreeStyle resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogAzureSecurityCenterFilterFreeStyleRead(d, m)
}

func resourceLogAzureSecurityCenterFilterFreeStyleUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogAzureSecurityCenterFilterFreeStyle(d)
	if err != nil {
		return fmt.Errorf("Error updating LogAzureSecurityCenterFilterFreeStyle resource while getting object: %v", err)
	}

	_, err = c.UpdateLogAzureSecurityCenterFilterFreeStyle(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogAzureSecurityCenterFilterFreeStyle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceLogAzureSecurityCenterFilterFreeStyleRead(d, m)
}

func resourceLogAzureSecurityCenterFilterFreeStyleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogAzureSecurityCenterFilterFreeStyle(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogAzureSecurityCenterFilterFreeStyle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogAzureSecurityCenterFilterFreeStyleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogAzureSecurityCenterFilterFreeStyle(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogAzureSecurityCenterFilterFreeStyle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogAzureSecurityCenterFilterFreeStyle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogAzureSecurityCenterFilterFreeStyle resource from API: %v", err)
	}
	return nil
}

func flattenLogAzureSecurityCenterFilterFreeStyleCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogAzureSecurityCenterFilterFreeStyleFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogAzureSecurityCenterFilterFreeStyleFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogAzureSecurityCenterFilterFreeStyleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogAzureSecurityCenterFilterFreeStyle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenLogAzureSecurityCenterFilterFreeStyleCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "LogAzureSecurityCenterFilterFreeStyle-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogAzureSecurityCenterFilterFreeStyleFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogAzureSecurityCenterFilterFreeStyle-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogAzureSecurityCenterFilterFreeStyleFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogAzureSecurityCenterFilterFreeStyle-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenLogAzureSecurityCenterFilterFreeStyleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "LogAzureSecurityCenterFilterFreeStyle-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenLogAzureSecurityCenterFilterFreeStyleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogAzureSecurityCenterFilterFreeStyleCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogAzureSecurityCenterFilterFreeStyleFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogAzureSecurityCenterFilterFreeStyleFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogAzureSecurityCenterFilterFreeStyleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogAzureSecurityCenterFilterFreeStyle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandLogAzureSecurityCenterFilterFreeStyleCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogAzureSecurityCenterFilterFreeStyleFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogAzureSecurityCenterFilterFreeStyleFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandLogAzureSecurityCenterFilterFreeStyleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
