// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure report page footer.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutPageFooter() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutPageFooterUpdate,
		Read:   resourceReportLayoutPageFooterRead,
		Update: resourceReportLayoutPageFooterUpdate,
		Delete: resourceReportLayoutPageFooterDelete,

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
			"footer_item": &schema.Schema{
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

func resourceReportLayoutPageFooterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectReportLayoutPageFooter(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageFooter resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutPageFooter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPageFooter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ReportLayoutPageFooter")

	return resourceReportLayoutPageFooterRead(d, m)
}

func resourceReportLayoutPageFooterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteReportLayoutPageFooter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutPageFooter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutPageFooterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayoutPageFooter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageFooter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutPageFooter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPageFooter resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutPageFooterFooterItem3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutPageFooterFooterItemContent3rdl(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageFooterFooterItemDescription3rdl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageFooterFooterItemId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageFooterFooterItemImgSrc3rdl(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageFooterFooterItemStyle3rdl(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageFooterFooterItemType3rdl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageFooterFooterItemContent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemImgSrc3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemStyle3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooterFooterItemType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterStyle3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectReportLayoutPageFooter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("footer_item", flattenReportLayoutPageFooterFooterItem3rdl(o["footer-item"], d, "footer_item")); err != nil {
			if vv, ok := fortiAPIPatch(o["footer-item"], "ReportLayoutPageFooter-FooterItem"); ok {
				if err = d.Set("footer_item", vv); err != nil {
					return fmt.Errorf("Error reading footer_item: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading footer_item: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("footer_item"); ok {
			if err = d.Set("footer_item", flattenReportLayoutPageFooterFooterItem3rdl(o["footer-item"], d, "footer_item")); err != nil {
				if vv, ok := fortiAPIPatch(o["footer-item"], "ReportLayoutPageFooter-FooterItem"); ok {
					if err = d.Set("footer_item", vv); err != nil {
						return fmt.Errorf("Error reading footer_item: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading footer_item: %v", err)
				}
			}
		}
	}

	if err = d.Set("style", flattenReportLayoutPageFooterStyle3rdl(o["style"], d, "style")); err != nil {
		if vv, ok := fortiAPIPatch(o["style"], "ReportLayoutPageFooter-Style"); ok {
			if err = d.Set("style", vv); err != nil {
				return fmt.Errorf("Error reading style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading style: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutPageFooterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutPageFooterFooterItem3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutPageFooterFooterItemContent3rdl(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageFooterFooterItemDescription3rdl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageFooterFooterItemId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageFooterFooterItemImgSrc3rdl(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageFooterFooterItemStyle3rdl(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageFooterFooterItemType3rdl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItemContent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemImgSrc3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemStyle3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooterFooterItemType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterStyle3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectReportLayoutPageFooter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("footer_item"); ok || d.HasChange("footer_item") {
		t, err := expandReportLayoutPageFooterFooterItem3rdl(d, v, "footer_item")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["footer-item"] = t
		}
	}

	if v, ok := d.GetOk("style"); ok || d.HasChange("style") {
		t, err := expandReportLayoutPageFooterStyle3rdl(d, v, "style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style"] = t
		}
	}

	return &obj, nil
}
