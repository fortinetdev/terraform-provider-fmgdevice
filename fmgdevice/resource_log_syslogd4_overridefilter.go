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

func resourceLogSyslogd4OverrideFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd4OverrideFilterUpdate,
		Read:   resourceLogSyslogd4OverrideFilterRead,
		Update: resourceLogSyslogd4OverrideFilterUpdate,
		Delete: resourceLogSyslogd4OverrideFilterDelete,

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

func resourceLogSyslogd4OverrideFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogSyslogd4OverrideFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd4OverrideFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSyslogd4OverrideFilter")

	return resourceLogSyslogd4OverrideFilterRead(d, m)
}

func resourceLogSyslogd4OverrideFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogd4OverrideFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd4OverrideFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd4OverrideFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd4OverrideFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd4OverrideFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd4OverrideFilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenLogSyslogd4OverrideFilterFreeStyleCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogSyslogd4OverrideFilter-FreeStyle-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenLogSyslogd4OverrideFilterFreeStyleFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "LogSyslogd4OverrideFilter-FreeStyle-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenLogSyslogd4OverrideFilterFreeStyleFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "LogSyslogd4OverrideFilter-FreeStyle-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogSyslogd4OverrideFilterFreeStyleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogSyslogd4OverrideFilter-FreeStyle-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogSyslogd4OverrideFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd4OverrideFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("anomaly", flattenLogSyslogd4OverrideFilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly"], "LogSyslogd4OverrideFilter-Anomaly"); ok {
			if err = d.Set("anomaly", vv); err != nil {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogSyslogd4OverrideFilterFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogSyslogd4OverrideFilter-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogSyslogd4OverrideFilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogSyslogd4OverrideFilter-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogSyslogd4OverrideFilterFortiSwitch(o["forti-switch"], d, "forti_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["forti-switch"], "LogSyslogd4OverrideFilter-FortiSwitch"); ok {
			if err = d.Set("forti_switch", vv); err != nil {
				return fmt.Errorf("Error reading forti_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogSyslogd4OverrideFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-traffic"], "LogSyslogd4OverrideFilter-ForwardTraffic"); ok {
			if err = d.Set("forward_traffic", vv); err != nil {
				return fmt.Errorf("Error reading forward_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogSyslogd4OverrideFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
			if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd4OverrideFilter-FreeStyle"); ok {
				if err = d.Set("free_style", vv); err != nil {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogSyslogd4OverrideFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
				if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd4OverrideFilter-FreeStyle"); ok {
					if err = d.Set("free_style", vv); err != nil {
						return fmt.Errorf("Error reading free_style: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp", flattenLogSyslogd4OverrideFilterGtp(o["gtp"], d, "gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp"], "LogSyslogd4OverrideFilter-Gtp"); ok {
			if err = d.Set("gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogSyslogd4OverrideFilterHttpTransaction(o["http-transaction"], d, "http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-transaction"], "LogSyslogd4OverrideFilter-HttpTransaction"); ok {
			if err = d.Set("http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogSyslogd4OverrideFilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-traffic"], "LogSyslogd4OverrideFilter-LocalTraffic"); ok {
			if err = d.Set("local_traffic", vv); err != nil {
				return fmt.Errorf("Error reading local_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogSyslogd4OverrideFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-traffic"], "LogSyslogd4OverrideFilter-MulticastTraffic"); ok {
			if err = d.Set("multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogSyslogd4OverrideFilterSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "LogSyslogd4OverrideFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogSyslogd4OverrideFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["sniffer-traffic"], "LogSyslogd4OverrideFilter-SnifferTraffic"); ok {
			if err = d.Set("sniffer_traffic", vv); err != nil {
				return fmt.Errorf("Error reading sniffer_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogSyslogd4OverrideFilterVoip(o["voip"], d, "voip")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip"], "LogSyslogd4OverrideFilter-Voip"); ok {
			if err = d.Set("voip", vv); err != nil {
				return fmt.Errorf("Error reading voip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogSyslogd4OverrideFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-traffic"], "LogSyslogd4OverrideFilter-ZtnaTraffic"); ok {
			if err = d.Set("ztna_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd4OverrideFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd4OverrideFilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandLogSyslogd4OverrideFilterFreeStyleCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandLogSyslogd4OverrideFilterFreeStyleFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandLogSyslogd4OverrideFilterFreeStyleFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogSyslogd4OverrideFilterFreeStyleId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogSyslogd4OverrideFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd4OverrideFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandLogSyslogd4OverrideFilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogSyslogd4OverrideFilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogSyslogd4OverrideFilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok || d.HasChange("forti_switch") {
		t, err := expandLogSyslogd4OverrideFilterFortiSwitch(d, v, "forti_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forti-switch"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok || d.HasChange("forward_traffic") {
		t, err := expandLogSyslogd4OverrideFilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		t, err := expandLogSyslogd4OverrideFilterFreeStyle(d, v, "free_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["free-style"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok || d.HasChange("gtp") {
		t, err := expandLogSyslogd4OverrideFilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok || d.HasChange("http_transaction") {
		t, err := expandLogSyslogd4OverrideFilterHttpTransaction(d, v, "http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok || d.HasChange("local_traffic") {
		t, err := expandLogSyslogd4OverrideFilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok || d.HasChange("multicast_traffic") {
		t, err := expandLogSyslogd4OverrideFilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandLogSyslogd4OverrideFilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok || d.HasChange("sniffer_traffic") {
		t, err := expandLogSyslogd4OverrideFilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok || d.HasChange("voip") {
		t, err := expandLogSyslogd4OverrideFilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok || d.HasChange("ztna_traffic") {
		t, err := expandLogSyslogd4OverrideFilterZtnaTraffic(d, v, "ztna_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-traffic"] = t
		}
	}

	return &obj, nil
}
