// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivityMatchTenantExtraction() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityMatchTenantExtractionUpdate,
		Read:   resourceCasbUserActivityMatchTenantExtractionRead,
		Update: resourceCasbUserActivityMatchTenantExtractionUpdate,
		Delete: resourceCasbUserActivityMatchTenantExtractionDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
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

func resourceCasbUserActivityMatchTenantExtractionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbUserActivityMatchTenantExtraction(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatchTenantExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("CasbUserActivityMatchTenantExtraction")

	return resourceCasbUserActivityMatchTenantExtractionRead(d, m)
}

func resourceCasbUserActivityMatchTenantExtractionDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteCasbUserActivityMatchTenantExtraction(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityMatchTenantExtractionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbUserActivityMatchTenantExtraction(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivityMatchTenantExtraction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivityMatchTenantExtraction resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityMatchTenantExtractionFilters3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionJq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivityMatchTenantExtraction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenCasbUserActivityMatchTenantExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "CasbUserActivityMatchTenantExtraction-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenCasbUserActivityMatchTenantExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "CasbUserActivityMatchTenantExtraction-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("jq", flattenCasbUserActivityMatchTenantExtractionJq3rdl(o["jq"], d, "jq")); err != nil {
		if vv, ok := fortiAPIPatch(o["jq"], "CasbUserActivityMatchTenantExtraction-Jq"); ok {
			if err = d.Set("jq", vv); err != nil {
				return fmt.Errorf("Error reading jq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jq: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbUserActivityMatchTenantExtractionStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "CasbUserActivityMatchTenantExtraction-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbUserActivityMatchTenantExtractionType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbUserActivityMatchTenantExtraction-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenCasbUserActivityMatchTenantExtractionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityMatchTenantExtractionFilters3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["body-type"], _ = expandCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchTenantExtractionFiltersId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionJq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivityMatchTenantExtraction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandCasbUserActivityMatchTenantExtractionFilters3rdl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("jq"); ok || d.HasChange("jq") {
		t, err := expandCasbUserActivityMatchTenantExtractionJq3rdl(d, v, "jq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jq"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandCasbUserActivityMatchTenantExtractionStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbUserActivityMatchTenantExtractionType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
