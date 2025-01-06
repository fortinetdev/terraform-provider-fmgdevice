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

func resourceLogSyslogd3Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd3FilterUpdate,
		Read:   resourceLogSyslogd3FilterRead,
		Update: resourceLogSyslogd3FilterUpdate,
		Delete: resourceLogSyslogd3FilterDelete,

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

func resourceLogSyslogd3FilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogd3Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd3Filter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogSyslogd3Filter")

	return resourceLogSyslogd3FilterRead(d, m)
}

func resourceLogSyslogd3FilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogSyslogd3Filter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd3Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd3FilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd3Filter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd3Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Filter resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd3FilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFreeStyle(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenLogSyslogd3FilterFreeStyleCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogSyslogd3Filter-FreeStyle-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenLogSyslogd3FilterFreeStyleFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "LogSyslogd3Filter-FreeStyle-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenLogSyslogd3FilterFreeStyleFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "LogSyslogd3Filter-FreeStyle-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogSyslogd3FilterFreeStyleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogSyslogd3Filter-FreeStyle-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogSyslogd3FilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogSyslogd3Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("anomaly", flattenLogSyslogd3FilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly"], "LogSyslogd3Filter-Anomaly"); ok {
			if err = d.Set("anomaly", vv); err != nil {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogSyslogd3FilterFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "LogSyslogd3Filter-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogSyslogd3FilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "LogSyslogd3Filter-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogSyslogd3FilterFortiSwitch(o["forti-switch"], d, "forti_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["forti-switch"], "LogSyslogd3Filter-FortiSwitch"); ok {
			if err = d.Set("forti_switch", vv); err != nil {
				return fmt.Errorf("Error reading forti_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogSyslogd3FilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-traffic"], "LogSyslogd3Filter-ForwardTraffic"); ok {
			if err = d.Set("forward_traffic", vv); err != nil {
				return fmt.Errorf("Error reading forward_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogSyslogd3FilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
			if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd3Filter-FreeStyle"); ok {
				if err = d.Set("free_style", vv); err != nil {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogSyslogd3FilterFreeStyle(o["free-style"], d, "free_style")); err != nil {
				if vv, ok := fortiAPIPatch(o["free-style"], "LogSyslogd3Filter-FreeStyle"); ok {
					if err = d.Set("free_style", vv); err != nil {
						return fmt.Errorf("Error reading free_style: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp", flattenLogSyslogd3FilterGtp(o["gtp"], d, "gtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp"], "LogSyslogd3Filter-Gtp"); ok {
			if err = d.Set("gtp", vv); err != nil {
				return fmt.Errorf("Error reading gtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogSyslogd3FilterHttpTransaction(o["http-transaction"], d, "http_transaction")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-transaction"], "LogSyslogd3Filter-HttpTransaction"); ok {
			if err = d.Set("http_transaction", vv); err != nil {
				return fmt.Errorf("Error reading http_transaction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogSyslogd3FilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-traffic"], "LogSyslogd3Filter-LocalTraffic"); ok {
			if err = d.Set("local_traffic", vv); err != nil {
				return fmt.Errorf("Error reading local_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogSyslogd3FilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-traffic"], "LogSyslogd3Filter-MulticastTraffic"); ok {
			if err = d.Set("multicast_traffic", vv); err != nil {
				return fmt.Errorf("Error reading multicast_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogSyslogd3FilterSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "LogSyslogd3Filter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogSyslogd3FilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["sniffer-traffic"], "LogSyslogd3Filter-SnifferTraffic"); ok {
			if err = d.Set("sniffer_traffic", vv); err != nil {
				return fmt.Errorf("Error reading sniffer_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogSyslogd3FilterVoip(o["voip"], d, "voip")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip"], "LogSyslogd3Filter-Voip"); ok {
			if err = d.Set("voip", vv); err != nil {
				return fmt.Errorf("Error reading voip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogSyslogd3FilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["ztna-traffic"], "LogSyslogd3Filter-ZtnaTraffic"); ok {
			if err = d.Set("ztna_traffic", vv); err != nil {
				return fmt.Errorf("Error reading ztna_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd3FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogSyslogd3FilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFreeStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandLogSyslogd3FilterFreeStyleCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandLogSyslogd3FilterFreeStyleFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandLogSyslogd3FilterFreeStyleFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogSyslogd3FilterFreeStyleId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogSyslogd3FilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd3Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandLogSyslogd3FilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandLogSyslogd3FilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandLogSyslogd3FilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok || d.HasChange("forti_switch") {
		t, err := expandLogSyslogd3FilterFortiSwitch(d, v, "forti_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forti-switch"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok || d.HasChange("forward_traffic") {
		t, err := expandLogSyslogd3FilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		t, err := expandLogSyslogd3FilterFreeStyle(d, v, "free_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["free-style"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok || d.HasChange("gtp") {
		t, err := expandLogSyslogd3FilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok || d.HasChange("http_transaction") {
		t, err := expandLogSyslogd3FilterHttpTransaction(d, v, "http_transaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok || d.HasChange("local_traffic") {
		t, err := expandLogSyslogd3FilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok || d.HasChange("multicast_traffic") {
		t, err := expandLogSyslogd3FilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandLogSyslogd3FilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok || d.HasChange("sniffer_traffic") {
		t, err := expandLogSyslogd3FilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok || d.HasChange("voip") {
		t, err := expandLogSyslogd3FilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok || d.HasChange("ztna_traffic") {
		t, err := expandLogSyslogd3FilterZtnaTraffic(d, v, "ztna_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-traffic"] = t
		}
	}

	return &obj, nil
}
