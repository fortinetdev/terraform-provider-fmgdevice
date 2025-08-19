// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Parameters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutBodyItemParameters() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutBodyItemParametersCreate,
		Read:   resourceReportLayoutBodyItemParametersRead,
		Update: resourceReportLayoutBodyItemParametersUpdate,
		Delete: resourceReportLayoutBodyItemParametersDelete,

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
			"layout": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body_item": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceReportLayoutBodyItemParametersCreate(d *schema.ResourceData, m interface{}) error {
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
	layout := d.Get("layout").(string)
	body_item := d.Get("body_item").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout
	paradict["body_item"] = body_item

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectReportLayoutBodyItemParameters(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutBodyItemParameters resource while getting object: %v", err)
	}

	_, err = c.CreateReportLayoutBodyItemParameters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutBodyItemParameters resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutBodyItemParametersRead(d, m)
}

func resourceReportLayoutBodyItemParametersUpdate(d *schema.ResourceData, m interface{}) error {
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
	layout := d.Get("layout").(string)
	body_item := d.Get("body_item").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout
	paradict["body_item"] = body_item

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectReportLayoutBodyItemParameters(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutBodyItemParameters resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutBodyItemParameters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutBodyItemParameters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutBodyItemParametersRead(d, m)
}

func resourceReportLayoutBodyItemParametersDelete(d *schema.ResourceData, m interface{}) error {
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
	layout := d.Get("layout").(string)
	body_item := d.Get("body_item").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout
	paradict["body_item"] = body_item

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteReportLayoutBodyItemParameters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutBodyItemParameters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutBodyItemParametersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	layout := d.Get("layout").(string)
	body_item := d.Get("body_item").(string)
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
	if layout == "" {
		layout = importOptionChecking(m.(*FortiClient).Cfg, "layout")
		if layout == "" {
			return fmt.Errorf("Parameter layout is missing")
		}
		if err = d.Set("layout", layout); err != nil {
			return fmt.Errorf("Error set params layout: %v", err)
		}
	}
	if body_item == "" {
		body_item = importOptionChecking(m.(*FortiClient).Cfg, "body_item")
		if body_item == "" {
			return fmt.Errorf("Parameter body_item is missing")
		}
		if err = d.Set("body_item", body_item); err != nil {
			return fmt.Errorf("Error set params body_item: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout
	paradict["body_item"] = body_item

	o, err := c.ReadReportLayoutBodyItemParameters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutBodyItemParameters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutBodyItemParameters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutBodyItemParameters resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutBodyItemParametersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportLayoutBodyItemParameters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenReportLayoutBodyItemParametersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ReportLayoutBodyItemParameters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenReportLayoutBodyItemParametersName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ReportLayoutBodyItemParameters-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenReportLayoutBodyItemParametersValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ReportLayoutBodyItemParameters-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutBodyItemParametersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutBodyItemParametersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayoutBodyItemParameters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandReportLayoutBodyItemParametersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandReportLayoutBodyItemParametersName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandReportLayoutBodyItemParametersValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
