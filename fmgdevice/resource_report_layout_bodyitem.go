// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure report body item.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayoutBodyItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutBodyItemCreate,
		Read:   resourceReportLayoutBodyItemRead,
		Update: resourceReportLayoutBodyItemUpdate,
		Delete: resourceReportLayoutBodyItemDelete,

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
			"chart": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"chart_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"column": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drill_down_items": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drill_down_types": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hide": &schema.Schema{
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
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"list_component": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"misc_component": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"table_caption_style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"table_column_widths": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_even_row_style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"table_head_style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"table_odd_row_style": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"text_component": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"top_n": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceReportLayoutBodyItemCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectReportLayoutBodyItem(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutBodyItem resource while getting object: %v", err)
	}

	_, err = c.CreateReportLayoutBodyItem(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayoutBodyItem resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutBodyItemRead(d, m)
}

func resourceReportLayoutBodyItemUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectReportLayoutBodyItem(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutBodyItem resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayoutBodyItem(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayoutBodyItem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceReportLayoutBodyItemRead(d, m)
}

func resourceReportLayoutBodyItemDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteReportLayoutBodyItem(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayoutBodyItem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutBodyItemRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayoutBodyItem(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutBodyItem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayoutBodyItem(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayoutBodyItem resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutBodyItemChart2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemChartOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemColumn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownItems2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownTypes2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemHide2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemImgSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutBodyItemListContent2edl(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-List-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutBodyItemListId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-List-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutBodyItemListContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListComponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemMiscComponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParameters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutBodyItemParametersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenReportLayoutBodyItemParametersName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenReportLayoutBodyItemParametersValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutBodyItemParametersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableCaptionStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableColumnWidths2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableEvenRowStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableHeadStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableOddRowStyle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTextComponent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTitle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTopN2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportLayoutBodyItem(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("chart", flattenReportLayoutBodyItemChart2edl(o["chart"], d, "chart")); err != nil {
		if vv, ok := fortiAPIPatch(o["chart"], "ReportLayoutBodyItem-Chart"); ok {
			if err = d.Set("chart", vv); err != nil {
				return fmt.Errorf("Error reading chart: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chart: %v", err)
		}
	}

	if err = d.Set("chart_options", flattenReportLayoutBodyItemChartOptions2edl(o["chart-options"], d, "chart_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["chart-options"], "ReportLayoutBodyItem-ChartOptions"); ok {
			if err = d.Set("chart_options", vv); err != nil {
				return fmt.Errorf("Error reading chart_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chart_options: %v", err)
		}
	}

	if err = d.Set("column", flattenReportLayoutBodyItemColumn2edl(o["column"], d, "column")); err != nil {
		if vv, ok := fortiAPIPatch(o["column"], "ReportLayoutBodyItem-Column"); ok {
			if err = d.Set("column", vv); err != nil {
				return fmt.Errorf("Error reading column: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading column: %v", err)
		}
	}

	if err = d.Set("content", flattenReportLayoutBodyItemContent2edl(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "ReportLayoutBodyItem-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("description", flattenReportLayoutBodyItemDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ReportLayoutBodyItem-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("drill_down_items", flattenReportLayoutBodyItemDrillDownItems2edl(o["drill-down-items"], d, "drill_down_items")); err != nil {
		if vv, ok := fortiAPIPatch(o["drill-down-items"], "ReportLayoutBodyItem-DrillDownItems"); ok {
			if err = d.Set("drill_down_items", vv); err != nil {
				return fmt.Errorf("Error reading drill_down_items: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drill_down_items: %v", err)
		}
	}

	if err = d.Set("drill_down_types", flattenReportLayoutBodyItemDrillDownTypes2edl(o["drill-down-types"], d, "drill_down_types")); err != nil {
		if vv, ok := fortiAPIPatch(o["drill-down-types"], "ReportLayoutBodyItem-DrillDownTypes"); ok {
			if err = d.Set("drill_down_types", vv); err != nil {
				return fmt.Errorf("Error reading drill_down_types: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drill_down_types: %v", err)
		}
	}

	if err = d.Set("hide", flattenReportLayoutBodyItemHide2edl(o["hide"], d, "hide")); err != nil {
		if vv, ok := fortiAPIPatch(o["hide"], "ReportLayoutBodyItem-Hide"); ok {
			if err = d.Set("hide", vv); err != nil {
				return fmt.Errorf("Error reading hide: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hide: %v", err)
		}
	}

	if err = d.Set("fosid", flattenReportLayoutBodyItemId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ReportLayoutBodyItem-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("img_src", flattenReportLayoutBodyItemImgSrc2edl(o["img-src"], d, "img_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["img-src"], "ReportLayoutBodyItem-ImgSrc"); ok {
			if err = d.Set("img_src", vv); err != nil {
				return fmt.Errorf("Error reading img_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading img_src: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenReportLayoutBodyItemList2edl(o["list"], d, "list")); err != nil {
			if vv, ok := fortiAPIPatch(o["list"], "ReportLayoutBodyItem-List"); ok {
				if err = d.Set("list", vv); err != nil {
					return fmt.Errorf("Error reading list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenReportLayoutBodyItemList2edl(o["list"], d, "list")); err != nil {
				if vv, ok := fortiAPIPatch(o["list"], "ReportLayoutBodyItem-List"); ok {
					if err = d.Set("list", vv); err != nil {
						return fmt.Errorf("Error reading list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if err = d.Set("list_component", flattenReportLayoutBodyItemListComponent2edl(o["list-component"], d, "list_component")); err != nil {
		if vv, ok := fortiAPIPatch(o["list-component"], "ReportLayoutBodyItem-ListComponent"); ok {
			if err = d.Set("list_component", vv); err != nil {
				return fmt.Errorf("Error reading list_component: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading list_component: %v", err)
		}
	}

	if err = d.Set("misc_component", flattenReportLayoutBodyItemMiscComponent2edl(o["misc-component"], d, "misc_component")); err != nil {
		if vv, ok := fortiAPIPatch(o["misc-component"], "ReportLayoutBodyItem-MiscComponent"); ok {
			if err = d.Set("misc_component", vv); err != nil {
				return fmt.Errorf("Error reading misc_component: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading misc_component: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameters", flattenReportLayoutBodyItemParameters2edl(o["parameters"], d, "parameters")); err != nil {
			if vv, ok := fortiAPIPatch(o["parameters"], "ReportLayoutBodyItem-Parameters"); ok {
				if err = d.Set("parameters", vv); err != nil {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading parameters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameters"); ok {
			if err = d.Set("parameters", flattenReportLayoutBodyItemParameters2edl(o["parameters"], d, "parameters")); err != nil {
				if vv, ok := fortiAPIPatch(o["parameters"], "ReportLayoutBodyItem-Parameters"); ok {
					if err = d.Set("parameters", vv); err != nil {
						return fmt.Errorf("Error reading parameters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			}
		}
	}

	if err = d.Set("style", flattenReportLayoutBodyItemStyle2edl(o["style"], d, "style")); err != nil {
		if vv, ok := fortiAPIPatch(o["style"], "ReportLayoutBodyItem-Style"); ok {
			if err = d.Set("style", vv); err != nil {
				return fmt.Errorf("Error reading style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading style: %v", err)
		}
	}

	if err = d.Set("table_caption_style", flattenReportLayoutBodyItemTableCaptionStyle2edl(o["table-caption-style"], d, "table_caption_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-caption-style"], "ReportLayoutBodyItem-TableCaptionStyle"); ok {
			if err = d.Set("table_caption_style", vv); err != nil {
				return fmt.Errorf("Error reading table_caption_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_caption_style: %v", err)
		}
	}

	if err = d.Set("table_column_widths", flattenReportLayoutBodyItemTableColumnWidths2edl(o["table-column-widths"], d, "table_column_widths")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-column-widths"], "ReportLayoutBodyItem-TableColumnWidths"); ok {
			if err = d.Set("table_column_widths", vv); err != nil {
				return fmt.Errorf("Error reading table_column_widths: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_column_widths: %v", err)
		}
	}

	if err = d.Set("table_even_row_style", flattenReportLayoutBodyItemTableEvenRowStyle2edl(o["table-even-row-style"], d, "table_even_row_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-even-row-style"], "ReportLayoutBodyItem-TableEvenRowStyle"); ok {
			if err = d.Set("table_even_row_style", vv); err != nil {
				return fmt.Errorf("Error reading table_even_row_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_even_row_style: %v", err)
		}
	}

	if err = d.Set("table_head_style", flattenReportLayoutBodyItemTableHeadStyle2edl(o["table-head-style"], d, "table_head_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-head-style"], "ReportLayoutBodyItem-TableHeadStyle"); ok {
			if err = d.Set("table_head_style", vv); err != nil {
				return fmt.Errorf("Error reading table_head_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_head_style: %v", err)
		}
	}

	if err = d.Set("table_odd_row_style", flattenReportLayoutBodyItemTableOddRowStyle2edl(o["table-odd-row-style"], d, "table_odd_row_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-odd-row-style"], "ReportLayoutBodyItem-TableOddRowStyle"); ok {
			if err = d.Set("table_odd_row_style", vv); err != nil {
				return fmt.Errorf("Error reading table_odd_row_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_odd_row_style: %v", err)
		}
	}

	if err = d.Set("text_component", flattenReportLayoutBodyItemTextComponent2edl(o["text-component"], d, "text_component")); err != nil {
		if vv, ok := fortiAPIPatch(o["text-component"], "ReportLayoutBodyItem-TextComponent"); ok {
			if err = d.Set("text_component", vv); err != nil {
				return fmt.Errorf("Error reading text_component: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading text_component: %v", err)
		}
	}

	if err = d.Set("title", flattenReportLayoutBodyItemTitle2edl(o["title"], d, "title")); err != nil {
		if vv, ok := fortiAPIPatch(o["title"], "ReportLayoutBodyItem-Title"); ok {
			if err = d.Set("title", vv); err != nil {
				return fmt.Errorf("Error reading title: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading title: %v", err)
		}
	}

	if err = d.Set("top_n", flattenReportLayoutBodyItemTopN2edl(o["top-n"], d, "top_n")); err != nil {
		if vv, ok := fortiAPIPatch(o["top-n"], "ReportLayoutBodyItem-TopN"); ok {
			if err = d.Set("top_n", vv); err != nil {
				return fmt.Errorf("Error reading top_n: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading top_n: %v", err)
		}
	}

	if err = d.Set("type", flattenReportLayoutBodyItemType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ReportLayoutBodyItem-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutBodyItemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutBodyItemChart2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemChartOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemColumn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownItems2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownTypes2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemHide2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemImgSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutBodyItemListContent2edl(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutBodyItemListId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemListContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListComponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemMiscComponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParameters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutBodyItemParametersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandReportLayoutBodyItemParametersName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandReportLayoutBodyItemParametersValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemParametersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableCaptionStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableColumnWidths2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableEvenRowStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableHeadStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableOddRowStyle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTextComponent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTitle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTopN2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayoutBodyItem(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("chart"); ok || d.HasChange("chart") {
		t, err := expandReportLayoutBodyItemChart2edl(d, v, "chart")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chart"] = t
		}
	}

	if v, ok := d.GetOk("chart_options"); ok || d.HasChange("chart_options") {
		t, err := expandReportLayoutBodyItemChartOptions2edl(d, v, "chart_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chart-options"] = t
		}
	}

	if v, ok := d.GetOk("column"); ok || d.HasChange("column") {
		t, err := expandReportLayoutBodyItemColumn2edl(d, v, "column")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandReportLayoutBodyItemContent2edl(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandReportLayoutBodyItemDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("drill_down_items"); ok || d.HasChange("drill_down_items") {
		t, err := expandReportLayoutBodyItemDrillDownItems2edl(d, v, "drill_down_items")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drill-down-items"] = t
		}
	}

	if v, ok := d.GetOk("drill_down_types"); ok || d.HasChange("drill_down_types") {
		t, err := expandReportLayoutBodyItemDrillDownTypes2edl(d, v, "drill_down_types")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drill-down-types"] = t
		}
	}

	if v, ok := d.GetOk("hide"); ok || d.HasChange("hide") {
		t, err := expandReportLayoutBodyItemHide2edl(d, v, "hide")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hide"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandReportLayoutBodyItemId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("img_src"); ok || d.HasChange("img_src") {
		t, err := expandReportLayoutBodyItemImgSrc2edl(d, v, "img_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["img-src"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok || d.HasChange("list") {
		t, err := expandReportLayoutBodyItemList2edl(d, v, "list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("list_component"); ok || d.HasChange("list_component") {
		t, err := expandReportLayoutBodyItemListComponent2edl(d, v, "list_component")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list-component"] = t
		}
	}

	if v, ok := d.GetOk("misc_component"); ok || d.HasChange("misc_component") {
		t, err := expandReportLayoutBodyItemMiscComponent2edl(d, v, "misc_component")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["misc-component"] = t
		}
	}

	if v, ok := d.GetOk("parameters"); ok || d.HasChange("parameters") {
		t, err := expandReportLayoutBodyItemParameters2edl(d, v, "parameters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameters"] = t
		}
	}

	if v, ok := d.GetOk("style"); ok || d.HasChange("style") {
		t, err := expandReportLayoutBodyItemStyle2edl(d, v, "style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style"] = t
		}
	}

	if v, ok := d.GetOk("table_caption_style"); ok || d.HasChange("table_caption_style") {
		t, err := expandReportLayoutBodyItemTableCaptionStyle2edl(d, v, "table_caption_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-caption-style"] = t
		}
	}

	if v, ok := d.GetOk("table_column_widths"); ok || d.HasChange("table_column_widths") {
		t, err := expandReportLayoutBodyItemTableColumnWidths2edl(d, v, "table_column_widths")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-column-widths"] = t
		}
	}

	if v, ok := d.GetOk("table_even_row_style"); ok || d.HasChange("table_even_row_style") {
		t, err := expandReportLayoutBodyItemTableEvenRowStyle2edl(d, v, "table_even_row_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-even-row-style"] = t
		}
	}

	if v, ok := d.GetOk("table_head_style"); ok || d.HasChange("table_head_style") {
		t, err := expandReportLayoutBodyItemTableHeadStyle2edl(d, v, "table_head_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-head-style"] = t
		}
	}

	if v, ok := d.GetOk("table_odd_row_style"); ok || d.HasChange("table_odd_row_style") {
		t, err := expandReportLayoutBodyItemTableOddRowStyle2edl(d, v, "table_odd_row_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-odd-row-style"] = t
		}
	}

	if v, ok := d.GetOk("text_component"); ok || d.HasChange("text_component") {
		t, err := expandReportLayoutBodyItemTextComponent2edl(d, v, "text_component")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text-component"] = t
		}
	}

	if v, ok := d.GetOk("title"); ok || d.HasChange("title") {
		t, err := expandReportLayoutBodyItemTitle2edl(d, v, "title")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["title"] = t
		}
	}

	if v, ok := d.GetOk("top_n"); ok || d.HasChange("top_n") {
		t, err := expandReportLayoutBodyItemTopN2edl(d, v, "top_n")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["top-n"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandReportLayoutBodyItemType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
