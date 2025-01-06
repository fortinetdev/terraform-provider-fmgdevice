// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filters for remote system server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogdFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogdFilterUpdate,
		Read:   resourceLogSyslogdFilterRead,
		Update: resourceLogSyslogdFilterUpdate,
		Delete: resourceLogSyslogdFilterDelete,

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

func resourceLogSyslogdFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectLogSyslogdFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogdFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogdFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSyslogdFilter")

	return resourceLogSyslogdFilterRead(d, m)
}

func resourceLogSyslogdFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteLogSyslogdFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogdFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogdFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadLogSyslogdFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogdFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogdFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogdFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogdFilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenLogSyslogdFilterFreeStyleCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogSyslogdFilter-FreeStyle-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenLogSyslogdFilterFreeStyleFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "LogSyslogdFilter-FreeStyle-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenLogSyslogdFilterFreeStyleFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "LogSyslogdFilter-FreeStyle-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogSyslogdFilterFreeStyleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogSyslogdFilter-FreeStyle-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogSyslogdFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogdFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogdFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("anomaly", flattenLogSyslogdFilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly"], "LogSyslogdFilter-Anomaly"); ok {
			if err = d.Set("anomaly", vv); err != nil {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogSyslogdFilterFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogSyslogdFilter-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogSyslogdFilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogSyslogdFilter-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogSyslogdFilterFortiSwitch(o["forti-switch"], d, "forti_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["forti-switch"], "LogSyslogdFilter-FortiSwitch"); ok {
			if err = d.Set("forti_switch", vv); err != nil {
				return fmt.Errorf("Error reading forti_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogSyslogdFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-traffic"], "LogSyslogdFilter-ForwardTraffic"); ok {
			if err = d.Set("forward_traffic", vv); err != nil {
				return fmt.Errorf("Error reading forward_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogSyslogdFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
			if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogdFilter-FreeStyle"); ok {
				if err = d.Set("free_style", vv); err != nil {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogSyslogdFilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
				if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogdFilter-FreeStyle"); ok {
					if err = d.Set("free_style", vv); err != nil {
						return fmt.Errorf("Error reading free_style: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp", flattenLogSyslogdFilterGtp(o["gtp"], d, "gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp"], "LogSyslogdFilter-Gtp"); ok {
			if err = d.Set("gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogSyslogdFilterHttpTransaction(o["http-transaction"], d, "http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-transaction"], "LogSyslogdFilter-HttpTransaction"); ok {
			if err = d.Set("http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogSyslogdFilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-traffic"], "LogSyslogdFilter-LocalTraffic"); ok {
			if err = d.Set("local_traffic", vv); err != nil {
				return fmt.Errorf("Error reading local_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogSyslogdFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-traffic"], "LogSyslogdFilter-MulticastTraffic"); ok {
			if err = d.Set("multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogSyslogdFilterSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "LogSyslogdFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogSyslogdFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["sniffer-traffic"], "LogSyslogdFilter-SnifferTraffic"); ok {
			if err = d.Set("sniffer_traffic", vv); err != nil {
				return fmt.Errorf("Error reading sniffer_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogSyslogdFilterVoip(o["voip"], d, "voip")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip"], "LogSyslogdFilter-Voip"); ok {
			if err = d.Set("voip", vv); err != nil {
				return fmt.Errorf("Error reading voip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogSyslogdFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-traffic"], "LogSyslogdFilter-ZtnaTraffic"); ok {
			if err = d.Set("ztna_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogdFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogdFilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandLogSyslogdFilterFreeStyleCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandLogSyslogdFilterFreeStyleFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandLogSyslogdFilterFreeStyleFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogSyslogdFilterFreeStyleId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogSyslogdFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogdFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogdFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandLogSyslogdFilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogSyslogdFilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogSyslogdFilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok || d.HasChange("forti_switch") {
		t, err := expandLogSyslogdFilterFortiSwitch(d, v, "forti_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forti-switch"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok || d.HasChange("forward_traffic") {
		t, err := expandLogSyslogdFilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		t, err := expandLogSyslogdFilterFreeStyle(d, v, "free_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["free-style"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok || d.HasChange("gtp") {
		t, err := expandLogSyslogdFilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok || d.HasChange("http_transaction") {
		t, err := expandLogSyslogdFilterHttpTransaction(d, v, "http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok || d.HasChange("local_traffic") {
		t, err := expandLogSyslogdFilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok || d.HasChange("multicast_traffic") {
		t, err := expandLogSyslogdFilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandLogSyslogdFilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok || d.HasChange("sniffer_traffic") {
		t, err := expandLogSyslogdFilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok || d.HasChange("voip") {
		t, err := expandLogSyslogdFilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok || d.HasChange("ztna_traffic") {
		t, err := expandLogSyslogdFilterZtnaTraffic(d, v, "ztna_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-traffic"] = t
		}
	}

	return &obj, nil
}
