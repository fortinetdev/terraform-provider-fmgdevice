// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB user activity tenant session extraction.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivityMatchTenantSessionExtraction() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityMatchTenantSessionExtractionUpdate,
		Read:   resourceCasbUserActivityMatchTenantSessionExtractionRead,
		Update: resourceCasbUserActivityMatchTenantSessionExtractionUpdate,
		Delete: resourceCasbUserActivityMatchTenantSessionExtractionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cookie_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"place": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"jq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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

func resourceCasbUserActivityMatchTenantSessionExtractionUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchTenantSessionExtraction(d, false)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantSessionExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatchTenantSessionExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("CasbUserActivityMatchTenantSessionExtraction")

	return resourceCasbUserActivityMatchTenantSessionExtractionRead(d, m)
}

func resourceCasbUserActivityMatchTenantSessionExtractionDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectCasbUserActivityMatchTenantSessionExtraction(d, true)

	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantSessionExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatchTenantSessionExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing CasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityMatchTenantSessionExtractionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
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
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadCasbUserActivityMatchTenantSessionExtraction(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivityMatchTenantSessionExtraction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantSessionExtraction resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityMatchTenantSessionExtractionFilters3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cookie_name"
		if _, ok := i["cookie-name"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(i["cookie-name"], d, pre_append)
			tmp["cookie_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-CookieName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantSessionExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionJq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantSessionExtractionStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivityMatchTenantSessionExtraction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenCasbUserActivityMatchTenantSessionExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "CasbUserActivityMatchTenantSessionExtraction-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenCasbUserActivityMatchTenantSessionExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "CasbUserActivityMatchTenantSessionExtraction-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("jq", flattenCasbUserActivityMatchTenantSessionExtractionJq3rdl(o["jq"], d, "jq")); err != nil {
		if vv, ok := fortiAPIPatch(o["jq"], "CasbUserActivityMatchTenantSessionExtraction-Jq"); ok {
			if err = d.Set("jq", vv); err != nil {
				return fmt.Errorf("Error reading jq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jq: %v", err)
		}
	}

	if err = d.Set("session_match", flattenCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(o["session-match"], d, "session_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-match"], "CasbUserActivityMatchTenantSessionExtraction-SessionMatch"); ok {
			if err = d.Set("session_match", vv); err != nil {
				return fmt.Errorf("Error reading session_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_match: %v", err)
		}
	}

	if err = d.Set("session_source", flattenCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(o["session-source"], d, "session_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-source"], "CasbUserActivityMatchTenantSessionExtraction-SessionSource"); ok {
			if err = d.Set("session_source", vv); err != nil {
				return fmt.Errorf("Error reading session_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_source: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbUserActivityMatchTenantSessionExtractionStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "CasbUserActivityMatchTenantSessionExtraction-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenCasbUserActivityMatchTenantSessionExtractionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityMatchTenantSessionExtractionFilters3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cookie_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cookie-name"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(d, i["cookie_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionJq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantSessionExtractionStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivityMatchTenantSessionExtraction(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if bemptysontable {
		obj["filters"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
			t, err := expandCasbUserActivityMatchTenantSessionExtractionFilters3rdl(d, v, "filters")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filters"] = t
			}
		}
	}

	if v, ok := d.GetOk("jq"); ok || d.HasChange("jq") {
		t, err := expandCasbUserActivityMatchTenantSessionExtractionJq3rdl(d, v, "jq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jq"] = t
		}
	}

	if v, ok := d.GetOk("session_match"); ok || d.HasChange("session_match") {
		t, err := expandCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(d, v, "session_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-match"] = t
		}
	}

	if v, ok := d.GetOk("session_source"); ok || d.HasChange("session_source") {
		t, err := expandCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(d, v, "session_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-source"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandCasbUserActivityMatchTenantSessionExtractionStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
