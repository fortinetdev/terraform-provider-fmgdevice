// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure report page header.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutPageHeader() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutPageHeaderUpdate,
		Read:   resourceReportLayoutPageHeaderRead,
		Update: resourceReportLayoutPageHeaderUpdate,
		Delete: resourceReportLayoutPageHeaderDelete,

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
			"header_item": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceReportLayoutPageHeaderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectReportLayoutPageHeader(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageHeader resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutPageHeader(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageHeader resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ReportLayoutPageHeader")

	return resourceReportLayoutPageHeaderRead(d, m)
}

func resourceReportLayoutPageHeaderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteReportLayoutPageHeader(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutPageHeader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutPageHeaderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayoutPageHeader(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageHeader resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutPageHeader(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageHeader resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutPageHeaderHeaderItem3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemContent3rdl(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemDescription3rdl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemImgSrc3rdl(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemStyle3rdl(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemType3rdl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageHeaderHeaderItemContent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemImgSrc3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemStyle3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageHeaderHeaderItemType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderStyle3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectReportLayoutPageHeader(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("header_item", flattenReportLayoutPageHeaderHeaderItem3rdl(o["header-item"], d, "header_item")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-item"], "ReportLayoutPageHeader-HeaderItem"); ok {
				if err = d.Set("header_item", vv); err != nil {
					return fmt.Errorf("Error reading header_item: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_item: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_item"); ok {
			if err = d.Set("header_item", flattenReportLayoutPageHeaderHeaderItem3rdl(o["header-item"], d, "header_item")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-item"], "ReportLayoutPageHeader-HeaderItem"); ok {
					if err = d.Set("header_item", vv); err != nil {
						return fmt.Errorf("Error reading header_item: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_item: %v", err)
				}
			}
		}
	}

	if err = d.Set("style", flattenReportLayoutPageHeaderStyle3rdl(o["style"], d, "style")); err != nil {
		if vv, ok := fortiAPIPatch(o["style"], "ReportLayoutPageHeader-Style"); ok {
			if err = d.Set("style", vv); err != nil {
				return fmt.Errorf("Error reading style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading style: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutPageHeaderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutPageHeaderHeaderItem3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content"], _ = expandReportLayoutPageHeaderHeaderItemContent3rdl(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageHeaderHeaderItemDescription3rdl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageHeaderHeaderItemId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageHeaderHeaderItemImgSrc3rdl(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageHeaderHeaderItemStyle3rdl(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageHeaderHeaderItemType3rdl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItemContent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemImgSrc3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemStyle3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageHeaderHeaderItemType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderStyle3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectReportLayoutPageHeader(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("header_item"); ok || d.HasChange("header_item") {
		t, err := expandReportLayoutPageHeaderHeaderItem3rdl(d, v, "header_item")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-item"] = t
		}
	}

	if v, ok := d.GetOk("style"); ok || d.HasChange("style") {
		t, err := expandReportLayoutPageHeaderStyle3rdl(d, v, "style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style"] = t
		}
	}

	return &obj, nil
}
