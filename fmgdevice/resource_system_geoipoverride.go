// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGeoipOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGeoipOverrideCreate,
		Read:   resourceSystemGeoipOverrideRead,
		Update: resourceSystemGeoipOverrideUpdate,
		Delete: resourceSystemGeoipOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip6_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceSystemGeoipOverrideCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemGeoipOverride(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeoipOverride resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemGeoipOverride(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemGeoipOverride(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemGeoipOverride resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemGeoipOverride(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemGeoipOverride resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGeoipOverrideRead(d, m)
}

func resourceSystemGeoipOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemGeoipOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverride resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemGeoipOverride(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGeoipOverrideRead(d, m)
}

func resourceSystemGeoipOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemGeoipOverride(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGeoipOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGeoipOverrideRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemGeoipOverride(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemGeoipOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeoipOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeoipOverride resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeoipOverrideCountryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystemGeoipOverrideIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemGeoipOverrideIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystemGeoipOverrideIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-IpRange-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemGeoipOverrideIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6Range(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystemGeoipOverrideIp6RangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-Ip6Range-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemGeoipOverrideIp6RangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-Ip6Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystemGeoipOverrideIp6RangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystemGeoipOverride-Ip6Range-StartIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemGeoipOverrideIp6RangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6RangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6RangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeoipOverrideName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGeoipOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("country_id", flattenSystemGeoipOverrideCountryId(o["country-id"], d, "country_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-id"], "SystemGeoipOverride-CountryId"); ok {
			if err = d.Set("country_id", vv); err != nil {
				return fmt.Errorf("Error reading country_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemGeoipOverrideDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemGeoipOverride-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SystemGeoipOverride-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SystemGeoipOverride-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_range", flattenSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-range"], "SystemGeoipOverride-Ip6Range"); ok {
				if err = d.Set("ip6_range", vv); err != nil {
					return fmt.Errorf("Error reading ip6_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_range"); ok {
			if err = d.Set("ip6_range", flattenSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-range"], "SystemGeoipOverride-Ip6Range"); ok {
					if err = d.Set("ip6_range", vv); err != nil {
						return fmt.Errorf("Error reading ip6_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemGeoipOverrideName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemGeoipOverride-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemGeoipOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGeoipOverrideCountryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystemGeoipOverrideIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemGeoipOverrideIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystemGeoipOverrideIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemGeoipOverrideIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6Range(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystemGeoipOverrideIp6RangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemGeoipOverrideIp6RangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystemGeoipOverrideIp6RangeStartIp(d, i["start_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemGeoipOverrideIp6RangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeoipOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("country_id"); ok || d.HasChange("country_id") {
		t, err := expandSystemGeoipOverrideCountryId(d, v, "country_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemGeoipOverrideDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemGeoipOverrideIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ip6_range"); ok || d.HasChange("ip6_range") {
		t, err := expandSystemGeoipOverrideIp6Range(d, v, "ip6_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-range"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemGeoipOverrideName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
