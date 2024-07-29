// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure application signatures.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationName() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationNameUpdate,
		Read:   resourceApplicationNameRead,
		Update: resourceApplicationNameUpdate,
		Delete: resourceApplicationNameDelete,

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
			"behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_scan_range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"popularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"risk": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_category": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceApplicationNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectApplicationName(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationName resource while getting object: %v", err)
	}

	_, err = c.UpdateApplicationName(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ApplicationName")

	return resourceApplicationNameRead(d, m)
}

func resourceApplicationNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteApplicationName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationName resource from API: %v", err)
	}
	return nil
}

func flattenApplicationNameBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationNameId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameMaxScanRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameParameter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNamePopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameSubCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationNameWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("behavior", flattenApplicationNameBehavior(o["behavior"], d, "behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["behavior"], "ApplicationName-Behavior"); ok {
			if err = d.Set("behavior", vv); err != nil {
				return fmt.Errorf("Error reading behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationNameCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ApplicationName-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenApplicationNameId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ApplicationName-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("max_scan_range", flattenApplicationNameMaxScanRange(o["max-scan-range"], d, "max_scan_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-scan-range"], "ApplicationName-MaxScanRange"); ok {
			if err = d.Set("max_scan_range", vv); err != nil {
				return fmt.Errorf("Error reading max_scan_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_scan_range: %v", err)
		}
	}

	if err = d.Set("name", flattenApplicationNameName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ApplicationName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("parameter", flattenApplicationNameParameter(o["parameter"], d, "parameter")); err != nil {
		if vv, ok := fortiAPIPatch(o["parameter"], "ApplicationName-Parameter"); ok {
			if err = d.Set("parameter", vv); err != nil {
				return fmt.Errorf("Error reading parameter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parameter: %v", err)
		}
	}

	if err = d.Set("popularity", flattenApplicationNamePopularity(o["popularity"], d, "popularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["popularity"], "ApplicationName-Popularity"); ok {
			if err = d.Set("popularity", vv); err != nil {
				return fmt.Errorf("Error reading popularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("protocol", flattenApplicationNameProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ApplicationName-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("risk", flattenApplicationNameRisk(o["risk"], d, "risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["risk"], "ApplicationName-Risk"); ok {
			if err = d.Set("risk", vv); err != nil {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risk: %v", err)
		}
	}

	if err = d.Set("sub_category", flattenApplicationNameSubCategory(o["sub-category"], d, "sub_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-category"], "ApplicationName-SubCategory"); ok {
			if err = d.Set("sub_category", vv); err != nil {
				return fmt.Errorf("Error reading sub_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_category: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationNameTechnology(o["technology"], d, "technology")); err != nil {
		if vv, ok := fortiAPIPatch(o["technology"], "ApplicationName-Technology"); ok {
			if err = d.Set("technology", vv); err != nil {
				return fmt.Errorf("Error reading technology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationNameVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ApplicationName-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	if err = d.Set("weight", flattenApplicationNameWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ApplicationName-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenApplicationNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationNameBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationNameId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameMaxScanRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameParameter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNamePopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("behavior"); ok || d.HasChange("behavior") {
		t, err := expandApplicationNameBehavior(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandApplicationNameCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandApplicationNameId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("max_scan_range"); ok || d.HasChange("max_scan_range") {
		t, err := expandApplicationNameMaxScanRange(d, v, "max_scan_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-scan-range"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandApplicationNameName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("parameter"); ok || d.HasChange("parameter") {
		t, err := expandApplicationNameParameter(d, v, "parameter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok || d.HasChange("popularity") {
		t, err := expandApplicationNamePopularity(d, v, "popularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandApplicationNameProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {
		t, err := expandApplicationNameRisk(d, v, "risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOk("sub_category"); ok || d.HasChange("sub_category") {
		t, err := expandApplicationNameSubCategory(d, v, "sub_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-category"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok || d.HasChange("technology") {
		t, err := expandApplicationNameTechnology(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandApplicationNameVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandApplicationNameWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
