// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogDiskFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogDiskFilterUpdate,
		Read:   resourceLogDiskFilterRead,
		Update: resourceLogDiskFilterUpdate,
		Delete: resourceLogDiskFilterDelete,

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
			"anomaly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_archive": &schema.Schema{
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
			},
			"forti_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"free_style": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sniffer_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_traffic": &schema.Schema{
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

func resourceLogDiskFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogDiskFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogDiskFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogDiskFilter")

	return resourceLogDiskFilterRead(d, m)
}

func resourceLogDiskFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogDiskFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogDiskFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogDiskFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskFilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterDlpArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenLogDiskFilterFreeStyleCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogDiskFilter-FreeStyle-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenLogDiskFilterFreeStyleFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "LogDiskFilter-FreeStyle-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenLogDiskFilterFreeStyleFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "LogDiskFilter-FreeStyle-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogDiskFilterFreeStyleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogDiskFilter-FreeStyle-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogDiskFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogDiskFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("anomaly", flattenLogDiskFilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly"], "LogDiskFilter-Anomaly"); ok {
			if err = d.Set("anomaly", vv); err != nil {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogDiskFilterDlpArchive(o["dlp-archive"], d, "dlp_archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-archive"], "LogDiskFilter-DlpArchive"); ok {
			if err = d.Set("dlp_archive", vv); err != nil {
				return fmt.Errorf("Error reading dlp_archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogDiskFilterFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogDiskFilter-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogDiskFilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogDiskFilter-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogDiskFilterFortiSwitch(o["forti-switch"], d, "forti_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["forti-switch"], "LogDiskFilter-FortiSwitch"); ok {
			if err = d.Set("forti_switch", vv); err != nil {
				return fmt.Errorf("Error reading forti_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogDiskFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-traffic"], "LogDiskFilter-ForwardTraffic"); ok {
			if err = d.Set("forward_traffic", vv); err != nil {
				return fmt.Errorf("Error reading forward_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogDiskFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
			if vv, ok := fortiAPIPatch(o["free-style"], "LogDiskFilter-FreeStyle"); ok {
				if err = d.Set("free_style", vv); err != nil {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogDiskFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
				if vv, ok := fortiAPIPatch(o["free-style"], "LogDiskFilter-FreeStyle"); ok {
					if err = d.Set("free_style", vv); err != nil {
						return fmt.Errorf("Error reading free_style: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp", flattenLogDiskFilterGtp(o["gtp"], d, "gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp"], "LogDiskFilter-Gtp"); ok {
			if err = d.Set("gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogDiskFilterHttpTransaction(o["http-transaction"], d, "http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-transaction"], "LogDiskFilter-HttpTransaction"); ok {
			if err = d.Set("http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogDiskFilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-traffic"], "LogDiskFilter-LocalTraffic"); ok {
			if err = d.Set("local_traffic", vv); err != nil {
				return fmt.Errorf("Error reading local_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogDiskFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-traffic"], "LogDiskFilter-MulticastTraffic"); ok {
			if err = d.Set("multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogDiskFilterSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "LogDiskFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogDiskFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["sniffer-traffic"], "LogDiskFilter-SnifferTraffic"); ok {
			if err = d.Set("sniffer_traffic", vv); err != nil {
				return fmt.Errorf("Error reading sniffer_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogDiskFilterVoip(o["voip"], d, "voip")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip"], "LogDiskFilter-Voip"); ok {
			if err = d.Set("voip", vv); err != nil {
				return fmt.Errorf("Error reading voip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogDiskFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-traffic"], "LogDiskFilter-ZtnaTraffic"); ok {
			if err = d.Set("ztna_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func flattenLogDiskFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogDiskFilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterDlpArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandLogDiskFilterFreeStyleCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandLogDiskFilterFreeStyleFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandLogDiskFilterFreeStyleFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogDiskFilterFreeStyleId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogDiskFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandLogDiskFilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("dlp_archive"); ok || d.HasChange("dlp_archive") {
		t, err := expandLogDiskFilterDlpArchive(d, v, "dlp_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-archive"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogDiskFilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogDiskFilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok || d.HasChange("forti_switch") {
		t, err := expandLogDiskFilterFortiSwitch(d, v, "forti_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forti-switch"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok || d.HasChange("forward_traffic") {
		t, err := expandLogDiskFilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		t, err := expandLogDiskFilterFreeStyle(d, v, "free_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["free-style"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok || d.HasChange("gtp") {
		t, err := expandLogDiskFilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok || d.HasChange("http_transaction") {
		t, err := expandLogDiskFilterHttpTransaction(d, v, "http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok || d.HasChange("local_traffic") {
		t, err := expandLogDiskFilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok || d.HasChange("multicast_traffic") {
		t, err := expandLogDiskFilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandLogDiskFilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok || d.HasChange("sniffer_traffic") {
		t, err := expandLogDiskFilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok || d.HasChange("voip") {
		t, err := expandLogDiskFilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok || d.HasChange("ztna_traffic") {
		t, err := expandLogDiskFilterZtnaTraffic(d, v, "ztna_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-traffic"] = t
		}
	}

	return &obj, nil
}
