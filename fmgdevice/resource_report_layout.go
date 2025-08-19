// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report layout configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayout() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutCreate,
		Read:   resourceReportLayoutRead,
		Update: resourceReportLayoutUpdate,
		Delete: resourceReportLayoutDelete,

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
			"body_item": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"cutoff_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cutoff_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_recipients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_pdf_report": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"schedule_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"style_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subtitle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceReportLayoutCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectReportLayout(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayout resource while getting object: %v", err)
	}

	_, err = c.CreateReportLayout(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayout resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceReportLayoutRead(d, m)
}

func resourceReportLayoutUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectReportLayout(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayout resource while getting object: %v", err)
	}

	_, err = c.UpdateReportLayout(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayout resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceReportLayoutRead(d, m)
}

func resourceReportLayoutDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteReportLayout(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayout resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayout(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayout resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayout(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayout resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutBodyItem(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart"
		if _, ok := i["chart"]; ok {
			v := flattenReportLayoutBodyItemChart(i["chart"], d, pre_append)
			tmp["chart"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Chart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_options"
		if _, ok := i["chart-options"]; ok {
			v := flattenReportLayoutBodyItemChartOptions(i["chart-options"], d, pre_append)
			tmp["chart_options"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-ChartOptions")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if _, ok := i["column"]; ok {
			v := flattenReportLayoutBodyItemColumn(i["column"], d, pre_append)
			tmp["column"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Column")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenReportLayoutBodyItemContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutBodyItemDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_items"
		if _, ok := i["drill-down-items"]; ok {
			v := flattenReportLayoutBodyItemDrillDownItems(i["drill-down-items"], d, pre_append)
			tmp["drill_down_items"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-DrillDownItems")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_types"
		if _, ok := i["drill-down-types"]; ok {
			v := flattenReportLayoutBodyItemDrillDownTypes(i["drill-down-types"], d, pre_append)
			tmp["drill_down_types"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-DrillDownTypes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hide"
		if _, ok := i["hide"]; ok {
			v := flattenReportLayoutBodyItemHide(i["hide"], d, pre_append)
			tmp["hide"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Hide")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutBodyItemId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutBodyItemImgSrc(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			v := flattenReportLayoutBodyItemList(i["list"], d, pre_append)
			tmp["list"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-List")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list_component"
		if _, ok := i["list-component"]; ok {
			v := flattenReportLayoutBodyItemListComponent(i["list-component"], d, pre_append)
			tmp["list_component"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-ListComponent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "misc_component"
		if _, ok := i["misc-component"]; ok {
			v := flattenReportLayoutBodyItemMiscComponent(i["misc-component"], d, pre_append)
			tmp["misc_component"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-MiscComponent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := i["parameters"]; ok {
			v := flattenReportLayoutBodyItemParameters(i["parameters"], d, pre_append)
			tmp["parameters"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Parameters")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutBodyItemStyle(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_caption_style"
		if _, ok := i["table-caption-style"]; ok {
			v := flattenReportLayoutBodyItemTableCaptionStyle(i["table-caption-style"], d, pre_append)
			tmp["table_caption_style"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TableCaptionStyle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_column_widths"
		if _, ok := i["table-column-widths"]; ok {
			v := flattenReportLayoutBodyItemTableColumnWidths(i["table-column-widths"], d, pre_append)
			tmp["table_column_widths"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TableColumnWidths")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_even_row_style"
		if _, ok := i["table-even-row-style"]; ok {
			v := flattenReportLayoutBodyItemTableEvenRowStyle(i["table-even-row-style"], d, pre_append)
			tmp["table_even_row_style"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TableEvenRowStyle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_head_style"
		if _, ok := i["table-head-style"]; ok {
			v := flattenReportLayoutBodyItemTableHeadStyle(i["table-head-style"], d, pre_append)
			tmp["table_head_style"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TableHeadStyle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_odd_row_style"
		if _, ok := i["table-odd-row-style"]; ok {
			v := flattenReportLayoutBodyItemTableOddRowStyle(i["table-odd-row-style"], d, pre_append)
			tmp["table_odd_row_style"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TableOddRowStyle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "text_component"
		if _, ok := i["text-component"]; ok {
			v := flattenReportLayoutBodyItemTextComponent(i["text-component"], d, pre_append)
			tmp["text_component"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TextComponent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := i["title"]; ok {
			v := flattenReportLayoutBodyItemTitle(i["title"], d, pre_append)
			tmp["title"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Title")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "top_n"
		if _, ok := i["top-n"]; ok {
			v := flattenReportLayoutBodyItemTopN(i["top-n"], d, pre_append)
			tmp["top_n"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-TopN")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutBodyItemType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayout-BodyItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutBodyItemChart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemChartOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemColumn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownItems(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownTypes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemHide(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemImgSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutBodyItemListContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-List-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutBodyItemListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-List-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutBodyItemListContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListComponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemMiscComponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParameters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutBodyItemParametersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenReportLayoutBodyItemParametersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenReportLayoutBodyItemParametersValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ReportLayoutBodyItem-Parameters-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutBodyItemParametersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableCaptionStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableColumnWidths(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableEvenRowStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableHeadStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTableOddRowStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutBodyItemTextComponent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTitle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTopN(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutBodyItemType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutCutoffOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutCutoffTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutEmailRecipients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutEmailSend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutMaxPdfReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "column_break_before"
	if _, ok := i["column-break-before"]; ok {
		result["column_break_before"] = flattenReportLayoutPageColumnBreakBefore(i["column-break-before"], d, pre_append)
	}

	pre_append = pre + ".0." + "footer"
	if _, ok := i["footer"]; ok {
		result["footer"] = flattenReportLayoutPageFooter(i["footer"], d, pre_append)
	}

	pre_append = pre + ".0." + "header"
	if _, ok := i["header"]; ok {
		result["header"] = flattenReportLayoutPageHeader(i["header"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenReportLayoutPageOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "page_break_before"
	if _, ok := i["page-break-before"]; ok {
		result["page_break_before"] = flattenReportLayoutPagePageBreakBefore(i["page-break-before"], d, pre_append)
	}

	pre_append = pre + ".0." + "paper"
	if _, ok := i["paper"]; ok {
		result["paper"] = flattenReportLayoutPagePaper(i["paper"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageColumnBreakBefore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "footer_item"
	if _, ok := i["footer-item"]; ok {
		result["footer_item"] = flattenReportLayoutPageFooterFooterItem(i["footer-item"], d, pre_append)
	}

	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageFooterStyle(i["style"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageFooterFooterItem(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutPageFooterFooterItemContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageFooterFooterItemDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageFooterFooterItemId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageFooterFooterItemImgSrc(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageFooterFooterItemStyle(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageFooterFooterItemType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageFooter-FooterItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageFooterFooterItemContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemImgSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageFooterFooterItemType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageFooterStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageHeader(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_item"
	if _, ok := i["header-item"]; ok {
		result["header_item"] = flattenReportLayoutPageHeaderHeaderItem(i["header-item"], d, pre_append)
	}

	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageHeaderStyle(i["style"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageHeaderHeaderItem(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenReportLayoutPageHeaderHeaderItemContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := i["img-src"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemImgSrc(i["img-src"], d, pre_append)
			tmp["img_src"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-ImgSrc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := i["style"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemStyle(i["style"], d, pre_append)
			tmp["style"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Style")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenReportLayoutPageHeaderHeaderItemType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ReportLayoutPageHeader-HeaderItem-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenReportLayoutPageHeaderHeaderItemContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemImgSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageHeaderHeaderItemType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPageOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPagePageBreakBefore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportLayoutPagePaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutScheduleType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutStyleTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutSubtitle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportLayoutTitle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportLayout(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("body_item", flattenReportLayoutBodyItem(o["body-item"], d, "body_item")); err != nil {
			if vv, ok := fortiAPIPatch(o["body-item"], "ReportLayout-BodyItem"); ok {
				if err = d.Set("body_item", vv); err != nil {
					return fmt.Errorf("Error reading body_item: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading body_item: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("body_item"); ok {
			if err = d.Set("body_item", flattenReportLayoutBodyItem(o["body-item"], d, "body_item")); err != nil {
				if vv, ok := fortiAPIPatch(o["body-item"], "ReportLayout-BodyItem"); ok {
					if err = d.Set("body_item", vv); err != nil {
						return fmt.Errorf("Error reading body_item: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading body_item: %v", err)
				}
			}
		}
	}

	if err = d.Set("cutoff_option", flattenReportLayoutCutoffOption(o["cutoff-option"], d, "cutoff_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["cutoff-option"], "ReportLayout-CutoffOption"); ok {
			if err = d.Set("cutoff_option", vv); err != nil {
				return fmt.Errorf("Error reading cutoff_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cutoff_option: %v", err)
		}
	}

	if err = d.Set("cutoff_time", flattenReportLayoutCutoffTime(o["cutoff-time"], d, "cutoff_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["cutoff-time"], "ReportLayout-CutoffTime"); ok {
			if err = d.Set("cutoff_time", vv); err != nil {
				return fmt.Errorf("Error reading cutoff_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cutoff_time: %v", err)
		}
	}

	if err = d.Set("day", flattenReportLayoutDay(o["day"], d, "day")); err != nil {
		if vv, ok := fortiAPIPatch(o["day"], "ReportLayout-Day"); ok {
			if err = d.Set("day", vv); err != nil {
				return fmt.Errorf("Error reading day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("description", flattenReportLayoutDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ReportLayout-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("email_recipients", flattenReportLayoutEmailRecipients(o["email-recipients"], d, "email_recipients")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-recipients"], "ReportLayout-EmailRecipients"); ok {
			if err = d.Set("email_recipients", vv); err != nil {
				return fmt.Errorf("Error reading email_recipients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_recipients: %v", err)
		}
	}

	if err = d.Set("email_send", flattenReportLayoutEmailSend(o["email-send"], d, "email_send")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-send"], "ReportLayout-EmailSend"); ok {
			if err = d.Set("email_send", vv); err != nil {
				return fmt.Errorf("Error reading email_send: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_send: %v", err)
		}
	}

	if err = d.Set("format", flattenReportLayoutFormat(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "ReportLayout-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("max_pdf_report", flattenReportLayoutMaxPdfReport(o["max-pdf-report"], d, "max_pdf_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-pdf-report"], "ReportLayout-MaxPdfReport"); ok {
			if err = d.Set("max_pdf_report", vv); err != nil {
				return fmt.Errorf("Error reading max_pdf_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_pdf_report: %v", err)
		}
	}

	if err = d.Set("name", flattenReportLayoutName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ReportLayout-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenReportLayoutOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ReportLayout-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("page", flattenReportLayoutPage(o["page"], d, "page")); err != nil {
			if vv, ok := fortiAPIPatch(o["page"], "ReportLayout-Page"); ok {
				if err = d.Set("page", vv); err != nil {
					return fmt.Errorf("Error reading page: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading page: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("page"); ok {
			if err = d.Set("page", flattenReportLayoutPage(o["page"], d, "page")); err != nil {
				if vv, ok := fortiAPIPatch(o["page"], "ReportLayout-Page"); ok {
					if err = d.Set("page", vv); err != nil {
						return fmt.Errorf("Error reading page: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading page: %v", err)
				}
			}
		}
	}

	if err = d.Set("schedule_type", flattenReportLayoutScheduleType(o["schedule-type"], d, "schedule_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule-type"], "ReportLayout-ScheduleType"); ok {
			if err = d.Set("schedule_type", vv); err != nil {
				return fmt.Errorf("Error reading schedule_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule_type: %v", err)
		}
	}

	if err = d.Set("style_theme", flattenReportLayoutStyleTheme(o["style-theme"], d, "style_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["style-theme"], "ReportLayout-StyleTheme"); ok {
			if err = d.Set("style_theme", vv); err != nil {
				return fmt.Errorf("Error reading style_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading style_theme: %v", err)
		}
	}

	if err = d.Set("subtitle", flattenReportLayoutSubtitle(o["subtitle"], d, "subtitle")); err != nil {
		if vv, ok := fortiAPIPatch(o["subtitle"], "ReportLayout-Subtitle"); ok {
			if err = d.Set("subtitle", vv); err != nil {
				return fmt.Errorf("Error reading subtitle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subtitle: %v", err)
		}
	}

	if err = d.Set("time", flattenReportLayoutTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "ReportLayout-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("title", flattenReportLayoutTitle(o["title"], d, "title")); err != nil {
		if vv, ok := fortiAPIPatch(o["title"], "ReportLayout-Title"); ok {
			if err = d.Set("title", vv); err != nil {
				return fmt.Errorf("Error reading title: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading title: %v", err)
		}
	}

	return nil
}

func flattenReportLayoutFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportLayoutBodyItem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["chart"], _ = expandReportLayoutBodyItemChart(d, i["chart"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_options"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["chart-options"], _ = expandReportLayoutBodyItemChartOptions(d, i["chart_options"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["column"], _ = expandReportLayoutBodyItemColumn(d, i["column"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["content"], _ = expandReportLayoutBodyItemContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutBodyItemDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_items"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["drill-down-items"], _ = expandReportLayoutBodyItemDrillDownItems(d, i["drill_down_items"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_types"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["drill-down-types"], _ = expandReportLayoutBodyItemDrillDownTypes(d, i["drill_down_types"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hide"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hide"], _ = expandReportLayoutBodyItemHide(d, i["hide"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutBodyItemId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutBodyItemImgSrc(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandReportLayoutBodyItemList(d, i["list"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["list"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list_component"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["list-component"], _ = expandReportLayoutBodyItemListComponent(d, i["list_component"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "misc_component"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["misc-component"], _ = expandReportLayoutBodyItemMiscComponent(d, i["misc_component"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandReportLayoutBodyItemParameters(d, i["parameters"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["parameters"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutBodyItemStyle(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_caption_style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["table-caption-style"], _ = expandReportLayoutBodyItemTableCaptionStyle(d, i["table_caption_style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_column_widths"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["table-column-widths"], _ = expandReportLayoutBodyItemTableColumnWidths(d, i["table_column_widths"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_even_row_style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["table-even-row-style"], _ = expandReportLayoutBodyItemTableEvenRowStyle(d, i["table_even_row_style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_head_style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["table-head-style"], _ = expandReportLayoutBodyItemTableHeadStyle(d, i["table_head_style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_odd_row_style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["table-odd-row-style"], _ = expandReportLayoutBodyItemTableOddRowStyle(d, i["table_odd_row_style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "text_component"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["text-component"], _ = expandReportLayoutBodyItemTextComponent(d, i["text_component"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["title"], _ = expandReportLayoutBodyItemTitle(d, i["title"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "top_n"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["top-n"], _ = expandReportLayoutBodyItemTopN(d, i["top_n"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutBodyItemType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemChart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemChartOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemColumn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownItems(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownTypes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemHide(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemImgSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutBodyItemListContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutBodyItemListId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemListContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListComponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemMiscComponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutBodyItemParametersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandReportLayoutBodyItemParametersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandReportLayoutBodyItemParametersValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemParametersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableCaptionStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableColumnWidths(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableEvenRowStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableHeadStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTableOddRowStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutBodyItemTextComponent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTitle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTopN(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutCutoffOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutCutoffTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutEmailRecipients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutEmailSend(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutMaxPdfReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "column_break_before"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["column-break-before"], _ = expandReportLayoutPageColumnBreakBefore(d, i["column_break_before"], pre_append)
	}
	pre_append = pre + ".0." + "footer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageFooter(d, i["footer"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["footer"] = t
		}
	}
	pre_append = pre + ".0." + "header"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageHeader(d, i["header"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["header"] = t
		}
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandReportLayoutPageOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "page_break_before"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["page-break-before"], _ = expandReportLayoutPagePageBreakBefore(d, i["page_break_before"], pre_append)
	}
	pre_append = pre + ".0." + "paper"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["paper"], _ = expandReportLayoutPagePaper(d, i["paper"], pre_append)
	}

	return result, nil
}

func expandReportLayoutPageColumnBreakBefore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "footer_item"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageFooterFooterItem(d, i["footer_item"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["footer-item"] = t
		}
	}
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["style"], _ = expandReportLayoutPageFooterStyle(d, i["style"], pre_append)
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutPageFooterFooterItemContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageFooterFooterItemDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageFooterFooterItemId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageFooterFooterItemImgSrc(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageFooterFooterItemStyle(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageFooterFooterItemType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItemContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemImgSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageFooterFooterItemType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_item"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandReportLayoutPageHeaderHeaderItem(d, i["header_item"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["header-item"] = t
		}
	}
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["style"], _ = expandReportLayoutPageHeaderStyle(d, i["style"], pre_append)
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["content"], _ = expandReportLayoutPageHeaderHeaderItemContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandReportLayoutPageHeaderHeaderItemDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandReportLayoutPageHeaderHeaderItemId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["img-src"], _ = expandReportLayoutPageHeaderHeaderItemImgSrc(d, i["img_src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["style"], _ = expandReportLayoutPageHeaderHeaderItemStyle(d, i["style"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandReportLayoutPageHeaderHeaderItemType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItemContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemImgSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageHeaderHeaderItemType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPageOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPagePageBreakBefore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportLayoutPagePaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutScheduleType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutStyleTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutSubtitle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutTitle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayout(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("body_item"); ok || d.HasChange("body_item") {
		t, err := expandReportLayoutBodyItem(d, v, "body_item")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-item"] = t
		}
	}

	if v, ok := d.GetOk("cutoff_option"); ok || d.HasChange("cutoff_option") {
		t, err := expandReportLayoutCutoffOption(d, v, "cutoff_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cutoff-option"] = t
		}
	}

	if v, ok := d.GetOk("cutoff_time"); ok || d.HasChange("cutoff_time") {
		t, err := expandReportLayoutCutoffTime(d, v, "cutoff_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cutoff-time"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok || d.HasChange("day") {
		t, err := expandReportLayoutDay(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandReportLayoutDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("email_recipients"); ok || d.HasChange("email_recipients") {
		t, err := expandReportLayoutEmailRecipients(d, v, "email_recipients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-recipients"] = t
		}
	}

	if v, ok := d.GetOk("email_send"); ok || d.HasChange("email_send") {
		t, err := expandReportLayoutEmailSend(d, v, "email_send")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-send"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandReportLayoutFormat(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("max_pdf_report"); ok || d.HasChange("max_pdf_report") {
		t, err := expandReportLayoutMaxPdfReport(d, v, "max_pdf_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-pdf-report"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandReportLayoutName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandReportLayoutOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("page"); ok || d.HasChange("page") {
		t, err := expandReportLayoutPage(d, v, "page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page"] = t
		}
	}

	if v, ok := d.GetOk("schedule_type"); ok || d.HasChange("schedule_type") {
		t, err := expandReportLayoutScheduleType(d, v, "schedule_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-type"] = t
		}
	}

	if v, ok := d.GetOk("style_theme"); ok || d.HasChange("style_theme") {
		t, err := expandReportLayoutStyleTheme(d, v, "style_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style-theme"] = t
		}
	}

	if v, ok := d.GetOk("subtitle"); ok || d.HasChange("subtitle") {
		t, err := expandReportLayoutSubtitle(d, v, "subtitle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subtitle"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandReportLayoutTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("title"); ok || d.HasChange("title") {
		t, err := expandReportLayoutTitle(d, v, "title")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["title"] = t
		}
	}

	return &obj, nil
}
