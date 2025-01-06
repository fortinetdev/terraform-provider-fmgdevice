// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Override filters for remote system server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogd2OverrideFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd2OverrideFilterUpdate,
		Read:   resourceLogSyslogd2OverrideFilterRead,
		Update: resourceLogSyslogd2OverrideFilterUpdate,
		Delete: resourceLogSyslogd2OverrideFilterDelete,

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

func resourceLogSyslogd2OverrideFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogd2OverrideFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2OverrideFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd2OverrideFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd2OverrideFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSyslogd2OverrideFilter")

	return resourceLogSyslogd2OverrideFilterRead(d, m)
}

func resourceLogSyslogd2OverrideFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogd2OverrideFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd2OverrideFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd2OverrideFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd2OverrideFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2OverrideFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd2OverrideFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd2OverrideFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd2OverrideFilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenLogSyslogd2OverrideFilterFreeStyleCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogSyslogd2OverrideFilter-FreeStyle-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenLogSyslogd2OverrideFilterFreeStyleFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "LogSyslogd2OverrideFilter-FreeStyle-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenLogSyslogd2OverrideFilterFreeStyleFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "LogSyslogd2OverrideFilter-FreeStyle-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogSyslogd2OverrideFilterFreeStyleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogSyslogd2OverrideFilter-FreeStyle-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogSyslogd2OverrideFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd2OverrideFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd2OverrideFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("anomaly", flattenLogSyslogd2OverrideFilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly"], "LogSyslogd2OverrideFilter-Anomaly"); ok {
			if err = d.Set("anomaly", vv); err != nil {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogSyslogd2OverrideFilterFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogSyslogd2OverrideFilter-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogSyslogd2OverrideFilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogSyslogd2OverrideFilter-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogSyslogd2OverrideFilterFortiSwitch(o["forti-switch"], d, "forti_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["forti-switch"], "LogSyslogd2OverrideFilter-FortiSwitch"); ok {
			if err = d.Set("forti_switch", vv); err != nil {
				return fmt.Errorf("Error reading forti_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogSyslogd2OverrideFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-traffic"], "LogSyslogd2OverrideFilter-ForwardTraffic"); ok {
			if err = d.Set("forward_traffic", vv); err != nil {
				return fmt.Errorf("Error reading forward_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogSyslogd2OverrideFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
			if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd2OverrideFilter-FreeStyle"); ok {
				if err = d.Set("free_style", vv); err != nil {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogSyslogd2OverrideFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
				if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd2OverrideFilter-FreeStyle"); ok {
					if err = d.Set("free_style", vv); err != nil {
						return fmt.Errorf("Error reading free_style: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp", flattenLogSyslogd2OverrideFilterGtp(o["gtp"], d, "gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp"], "LogSyslogd2OverrideFilter-Gtp"); ok {
			if err = d.Set("gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogSyslogd2OverrideFilterHttpTransaction(o["http-transaction"], d, "http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-transaction"], "LogSyslogd2OverrideFilter-HttpTransaction"); ok {
			if err = d.Set("http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogSyslogd2OverrideFilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-traffic"], "LogSyslogd2OverrideFilter-LocalTraffic"); ok {
			if err = d.Set("local_traffic", vv); err != nil {
				return fmt.Errorf("Error reading local_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogSyslogd2OverrideFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-traffic"], "LogSyslogd2OverrideFilter-MulticastTraffic"); ok {
			if err = d.Set("multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogSyslogd2OverrideFilterSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "LogSyslogd2OverrideFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogSyslogd2OverrideFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["sniffer-traffic"], "LogSyslogd2OverrideFilter-SnifferTraffic"); ok {
			if err = d.Set("sniffer_traffic", vv); err != nil {
				return fmt.Errorf("Error reading sniffer_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogSyslogd2OverrideFilterVoip(o["voip"], d, "voip")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip"], "LogSyslogd2OverrideFilter-Voip"); ok {
			if err = d.Set("voip", vv); err != nil {
				return fmt.Errorf("Error reading voip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogSyslogd2OverrideFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-traffic"], "LogSyslogd2OverrideFilter-ZtnaTraffic"); ok {
			if err = d.Set("ztna_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd2OverrideFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd2OverrideFilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandLogSyslogd2OverrideFilterFreeStyleCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandLogSyslogd2OverrideFilterFreeStyleFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandLogSyslogd2OverrideFilterFreeStyleFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogSyslogd2OverrideFilterFreeStyleId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogSyslogd2OverrideFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd2OverrideFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd2OverrideFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandLogSyslogd2OverrideFilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogSyslogd2OverrideFilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogSyslogd2OverrideFilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok || d.HasChange("forti_switch") {
		t, err := expandLogSyslogd2OverrideFilterFortiSwitch(d, v, "forti_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forti-switch"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok || d.HasChange("forward_traffic") {
		t, err := expandLogSyslogd2OverrideFilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		t, err := expandLogSyslogd2OverrideFilterFreeStyle(d, v, "free_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["free-style"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok || d.HasChange("gtp") {
		t, err := expandLogSyslogd2OverrideFilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok || d.HasChange("http_transaction") {
		t, err := expandLogSyslogd2OverrideFilterHttpTransaction(d, v, "http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok || d.HasChange("local_traffic") {
		t, err := expandLogSyslogd2OverrideFilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok || d.HasChange("multicast_traffic") {
		t, err := expandLogSyslogd2OverrideFilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandLogSyslogd2OverrideFilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok || d.HasChange("sniffer_traffic") {
		t, err := expandLogSyslogd2OverrideFilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok || d.HasChange("voip") {
		t, err := expandLogSyslogd2OverrideFilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok || d.HasChange("ztna_traffic") {
		t, err := expandLogSyslogd2OverrideFilterZtnaTraffic(d, v, "ztna_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-traffic"] = t
		}
	}

	return &obj, nil
}
