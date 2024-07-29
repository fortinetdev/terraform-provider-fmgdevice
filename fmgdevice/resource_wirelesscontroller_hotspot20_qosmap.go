// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure QoS map set.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20QosMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20QosMapCreate,
		Read:   resourceWirelessControllerHotspot20QosMapRead,
		Update: resourceWirelessControllerHotspot20QosMapUpdate,
		Delete: resourceWirelessControllerHotspot20QosMapDelete,

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
			"dscp_except": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"up": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"dscp_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"up": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceWirelessControllerHotspot20QosMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20QosMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20QosMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerHotspot20QosMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20QosMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20QosMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20QosMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMap resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20QosMapDscpExcept(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpExceptDscp(i["dscp"], d, pre_append)
			tmp["dscp"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpExcept-Dscp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpExceptIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpExcept-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpExceptUp(i["up"], d, pre_append)
			tmp["up"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpExcept-Up")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20QosMapDscpExceptDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := i["high"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpRangeHigh(i["high"], d, pre_append)
			tmp["high"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpRange-High")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpRangeIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpRange-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := i["low"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpRangeLow(i["low"], d, pre_append)
			tmp["low"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpRange-Low")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			v := flattenWirelessControllerHotspot20QosMapDscpRangeUp(i["up"], d, pre_append)
			tmp["up"] = fortiAPISubPartPatch(v, "WirelessControllerHotspot20QosMap-DscpRange-Up")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20QosMapDscpRangeHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("dscp_except", flattenWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
			if vv, ok := fortiAPIPatch(o["dscp-except"], "WirelessControllerHotspot20QosMap-DscpExcept"); ok {
				if err = d.Set("dscp_except", vv); err != nil {
					return fmt.Errorf("Error reading dscp_except: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dscp_except: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_except"); ok {
			if err = d.Set("dscp_except", flattenWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
				if vv, ok := fortiAPIPatch(o["dscp-except"], "WirelessControllerHotspot20QosMap-DscpExcept"); ok {
					if err = d.Set("dscp_except", vv); err != nil {
						return fmt.Errorf("Error reading dscp_except: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dscp_except: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_range", flattenWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["dscp-range"], "WirelessControllerHotspot20QosMap-DscpRange"); ok {
				if err = d.Set("dscp_range", vv); err != nil {
					return fmt.Errorf("Error reading dscp_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dscp_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_range"); ok {
			if err = d.Set("dscp_range", flattenWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["dscp-range"], "WirelessControllerHotspot20QosMap-DscpRange"); ok {
					if err = d.Set("dscp_range", vv); err != nil {
						return fmt.Errorf("Error reading dscp_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dscp_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20QosMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20QosMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20QosMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp"], _ = expandWirelessControllerHotspot20QosMapDscpExceptDscp(d, i["dscp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20QosMapDscpExceptIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["up"], _ = expandWirelessControllerHotspot20QosMapDscpExceptUp(d, i["up"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["high"], _ = expandWirelessControllerHotspot20QosMapDscpRangeHigh(d, i["high"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandWirelessControllerHotspot20QosMapDscpRangeIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["low"], _ = expandWirelessControllerHotspot20QosMapDscpRangeLow(d, i["low"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["up"], _ = expandWirelessControllerHotspot20QosMapDscpRangeUp(d, i["up"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp_except"); ok || d.HasChange("dscp_except") {
		t, err := expandWirelessControllerHotspot20QosMapDscpExcept(d, v, "dscp_except")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-except"] = t
		}
	}

	if v, ok := d.GetOk("dscp_range"); ok || d.HasChange("dscp_range") {
		t, err := expandWirelessControllerHotspot20QosMapDscpRange(d, v, "dscp_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-range"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20QosMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
