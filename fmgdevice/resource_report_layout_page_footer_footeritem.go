// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure report footer item.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutPageFooterFooterItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutPageFooterFooterItemCreate,
		Read:   resourceReportLayoutPageFooterFooterItemRead,
		Update: resourceReportLayoutPageFooterFooterItemUpdate,
		Delete: resourceReportLayoutPageFooterFooterItemDelete,

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
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"img_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceReportLayoutPageFooterFooterItemCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectReportLayoutPageFooterFooterItem(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutPageFooterFooterItem resource while getting object: %v", err)
	}

	_, err = c.CreateReportLayoutPageFooterFooterItem(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutPageFooterFooterItem resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutPageFooterFooterItemRead(d, m)
}

func resourceReportLayoutPageFooterFooterItemUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectReportLayoutPageFooterFooterItem(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageFooterFooterItem resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutPageFooterFooterItem(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageFooterFooterItem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutPageFooterFooterItemRead(d, m)
}

func resourceReportLayoutPageFooterFooterItemDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteReportLayoutPageFooterFooterItem(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutPageFooterFooterItem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutPageFooterFooterItemRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	layout := d.Get("layout").(string)
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["layout"] = layout

	o, err := c.ReadReportLayoutPageFooterFooterItem(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageFooterFooterItem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutPageFooterFooterItem(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageFooterFooterItem resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutPageFooterFooterItemContent4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemDescription4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemImgSrc4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemStyle4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooterFooterItemType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportLayoutPageFooterFooterItem(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("content", flattenReportLayoutPageFooterFooterItemContent4thl(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "ReportLayoutPageFooterFooterItem-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("description", flattenReportLayoutPageFooterFooterItemDescription4thl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ReportLayoutPageFooterFooterItem-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenReportLayoutPageFooterFooterItemId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ReportLayoutPageFooterFooterItem-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("img_src", flattenReportLayoutPageFooterFooterItemImgSrc4thl(o["img-src"], d, "img_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["img-src"], "ReportLayoutPageFooterFooterItem-ImgSrc"); ok {
			if err = d.Set("img_src", vv); err != nil {
				return fmt.Errorf("Error reading img_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading img_src: %v", err)
		}
	}

	if err = d.Set("style", flattenReportLayoutPageFooterFooterItemStyle4thl(o["style"], d, "style")); err != nil {
		if vv, ok := fortiAPIPatch(o["style"], "ReportLayoutPageFooterFooterItem-Style"); ok {
			if err = d.Set("style", vv); err != nil {
				return fmt.Errorf("Error reading style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading style: %v", err)
		}
	}

	if err = d.Set("type", flattenReportLayoutPageFooterFooterItemType4thl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ReportLayoutPageFooterFooterItem-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutPageFooterFooterItemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutPageFooterFooterItemContent4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemDescription4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemImgSrc4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemStyle4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooterFooterItemType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayoutPageFooterFooterItem(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandReportLayoutPageFooterFooterItemContent4thl(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandReportLayoutPageFooterFooterItemDescription4thl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandReportLayoutPageFooterFooterItemId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("img_src"); ok || d.HasChange("img_src") {
		t, err := expandReportLayoutPageFooterFooterItemImgSrc4thl(d, v, "img_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["img-src"] = t
		}
	}

	if v, ok := d.GetOk("style"); ok || d.HasChange("style") {
		t, err := expandReportLayoutPageFooterFooterItemStyle4thl(d, v, "style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandReportLayoutPageFooterFooterItemType4thl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
