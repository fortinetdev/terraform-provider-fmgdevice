// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report setting configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportSettingUpdate,
		Read:   resourceReportSettingRead,
		Update: resourceReportSettingUpdate,
		Delete: resourceReportSettingDelete,

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
			"fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pdf_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_source": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"top_n": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"web_browsing_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceReportSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectReportSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateReportSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ReportSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ReportSetting")

	return resourceReportSettingRead(d, m)
}

func resourceReportSettingDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteReportSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ReportSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ReportSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportSetting resource from API: %v", err)
	}
	return nil
}

func flattenReportSettingFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportSettingPdfReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportSettingReportSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenReportSettingTopN(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportSettingWebBrowsingThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fortiview", flattenReportSettingFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "ReportSetting-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("pdf_report", flattenReportSettingPdfReport(o["pdf-report"], d, "pdf_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdf-report"], "ReportSetting-PdfReport"); ok {
			if err = d.Set("pdf_report", vv); err != nil {
				return fmt.Errorf("Error reading pdf_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdf_report: %v", err)
		}
	}

	if err = d.Set("report_source", flattenReportSettingReportSource(o["report-source"], d, "report_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-source"], "ReportSetting-ReportSource"); ok {
			if err = d.Set("report_source", vv); err != nil {
				return fmt.Errorf("Error reading report_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_source: %v", err)
		}
	}

	if err = d.Set("top_n", flattenReportSettingTopN(o["top-n"], d, "top_n")); err != nil {
		if vv, ok := fortiAPIPatch(o["top-n"], "ReportSetting-TopN"); ok {
			if err = d.Set("top_n", vv); err != nil {
				return fmt.Errorf("Error reading top_n: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading top_n: %v", err)
		}
	}

	if err = d.Set("web_browsing_threshold", flattenReportSettingWebBrowsingThreshold(o["web-browsing-threshold"], d, "web_browsing_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-browsing-threshold"], "ReportSetting-WebBrowsingThreshold"); ok {
			if err = d.Set("web_browsing_threshold", vv); err != nil {
				return fmt.Errorf("Error reading web_browsing_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_browsing_threshold: %v", err)
		}
	}

	return nil
}

func flattenReportSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportSettingFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportSettingPdfReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportSettingReportSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandReportSettingTopN(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportSettingWebBrowsingThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortiview"); ok || d.HasChange("fortiview") {
		t, err := expandReportSettingFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("pdf_report"); ok || d.HasChange("pdf_report") {
		t, err := expandReportSettingPdfReport(d, v, "pdf_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdf-report"] = t
		}
	}

	if v, ok := d.GetOk("report_source"); ok || d.HasChange("report_source") {
		t, err := expandReportSettingReportSource(d, v, "report_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-source"] = t
		}
	}

	if v, ok := d.GetOk("top_n"); ok || d.HasChange("top_n") {
		t, err := expandReportSettingTopN(d, v, "top_n")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["top-n"] = t
		}
	}

	if v, ok := d.GetOk("web_browsing_threshold"); ok || d.HasChange("web_browsing_threshold") {
		t, err := expandReportSettingWebBrowsingThreshold(d, v, "web_browsing_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-browsing-threshold"] = t
		}
	}

	return &obj, nil
}
