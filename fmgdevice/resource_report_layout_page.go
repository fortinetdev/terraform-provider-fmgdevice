// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure report page.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutPageUpdate,
		Read:   resourceReportLayoutPageRead,
		Update: resourceReportLayoutPageUpdate,
		Delete: resourceReportLayoutPageDelete,

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
			"column_break_before": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"footer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"header": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"page_break_before": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"paper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceReportLayoutPageUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectReportLayoutPage(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPage resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutPage(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutPage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ReportLayoutPage")

	return resourceReportLayoutPageRead(d, m)
}

func resourceReportLayoutPageDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteReportLayoutPage(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutPage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutPageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayoutPage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutPage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutPage resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutPageColumnBreakBefore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "footer_item"
	if _, ok := i["footer-item"]; ok {
		result["footer_item"] = flattenReportLayoutPageFooterFooterItem2edl(i["footer-item"], d, pre_append)
	}

	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageFooterStyle2edl(i["style"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageFooterFooterItem2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutPageFooterFooterItemContent2edl(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageFooterFooterItemDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageFooterFooterItemId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageFooterFooterItemImgSrc2edl(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageFooterFooterItemStyle2edl(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageFooterFooterItemType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageFooterFooterItemContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemImgSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooterFooterItemType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageHeader2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_item"
	if _, ok := i["header-item"]; ok {
		result["header_item"] = flattenReportLayoutPageHeaderHeaderItem2edl(i["header-item"], d, pre_append)
	}

	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageHeaderStyle2edl(i["style"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageHeaderHeaderItem2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutPageHeaderHeaderItemContent2edl(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemImgSrc2edl(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemStyle2edl(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageHeaderHeaderItemContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemImgSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageHeaderHeaderItemType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPagePageBreakBefore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPagePaper2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportLayoutPage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("column_break_before", flattenReportLayoutPageColumnBreakBefore2edl(o["column-break-before"], d, "column_break_before")); err != nil {
		if vv, ok := fortiAPIPatch(o["column-break-before"], "ReportLayoutPage-ColumnBreakBefore"); ok {
			if err = d.Set("column_break_before", vv); err != nil {
				return fmt.Errorf("Error reading column_break_before: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading column_break_before: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("footer", flattenReportLayoutPageFooter2edl(o["footer"], d, "footer")); err != nil {
			if vv, ok := fortiAPIPatch(o["footer"], "ReportLayoutPage-Footer"); ok {
				if err = d.Set("footer", vv); err != nil {
					return fmt.Errorf("Error reading footer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading footer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("footer"); ok {
			if err = d.Set("footer", flattenReportLayoutPageFooter2edl(o["footer"], d, "footer")); err != nil {
				if vv, ok := fortiAPIPatch(o["footer"], "ReportLayoutPage-Footer"); ok {
					if err = d.Set("footer", vv); err != nil {
						return fmt.Errorf("Error reading footer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading footer: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("header", flattenReportLayoutPageHeader2edl(o["header"], d, "header")); err != nil {
			if vv, ok := fortiAPIPatch(o["header"], "ReportLayoutPage-Header"); ok {
				if err = d.Set("header", vv); err != nil {
					return fmt.Errorf("Error reading header: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header"); ok {
			if err = d.Set("header", flattenReportLayoutPageHeader2edl(o["header"], d, "header")); err != nil {
				if vv, ok := fortiAPIPatch(o["header"], "ReportLayoutPage-Header"); ok {
					if err = d.Set("header", vv); err != nil {
						return fmt.Errorf("Error reading header: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header: %v", err)
				}
			}
		}
	}

	if err = d.Set("options", flattenReportLayoutPageOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ReportLayoutPage-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("page_break_before", flattenReportLayoutPagePageBreakBefore2edl(o["page-break-before"], d, "page_break_before")); err != nil {
		if vv, ok := fortiAPIPatch(o["page-break-before"], "ReportLayoutPage-PageBreakBefore"); ok {
			if err = d.Set("page_break_before", vv); err != nil {
				return fmt.Errorf("Error reading page_break_before: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading page_break_before: %v", err)
		}
	}

	if err = d.Set("paper", flattenReportLayoutPagePaper2edl(o["paper"], d, "paper")); err != nil {
		if vv, ok := fortiAPIPatch(o["paper"], "ReportLayoutPage-Paper"); ok {
			if err = d.Set("paper", vv); err != nil {
				return fmt.Errorf("Error reading paper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading paper: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutPageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutPageColumnBreakBefore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "footer_item"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageFooterFooterItem2edl(d, i["footer_item"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["footer-item"] = t
		}
	}
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["style"], _ = expandReportLayoutPageFooterStyle2edl(d, i["style"], pre_append)
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItem2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutPageFooterFooterItemContent2edl(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageFooterFooterItemDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageFooterFooterItemId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageFooterFooterItemImgSrc2edl(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageFooterFooterItemStyle2edl(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageFooterFooterItemType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItemContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemImgSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooterFooterItemType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_item"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageHeaderHeaderItem2edl(d, i["header_item"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["header-item"] = t
		}
	}
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["style"], _ = expandReportLayoutPageHeaderStyle2edl(d, i["style"], pre_append)
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItem2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutPageHeaderHeaderItemContent2edl(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageHeaderHeaderItemDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageHeaderHeaderItemId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageHeaderHeaderItemImgSrc2edl(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageHeaderHeaderItemStyle2edl(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageHeaderHeaderItemType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItemContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemImgSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageHeaderHeaderItemType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPagePageBreakBefore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPagePaper2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayoutPage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("column_break_before"); ok || d.HasChange("column_break_before") {
		t, err := expandReportLayoutPageColumnBreakBefore2edl(d, v, "column_break_before")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-break-before"] = t
		}
	}

	if v, ok := d.GetOk("footer"); ok || d.HasChange("footer") {
		t, err := expandReportLayoutPageFooter2edl(d, v, "footer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["footer"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandReportLayoutPageHeader2edl(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandReportLayoutPageOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("page_break_before"); ok || d.HasChange("page_break_before") {
		t, err := expandReportLayoutPagePageBreakBefore2edl(d, v, "page_break_before")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page-break-before"] = t
		}
	}

	if v, ok := d.GetOk("paper"); ok || d.HasChange("paper") {
		t, err := expandReportLayoutPagePaper2edl(d, v, "paper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["paper"] = t
		}
	}

	return &obj, nil
}
