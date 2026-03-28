// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> URL access list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileUrlAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileUrlAccessCreate,
		Read:   resourceWafProfileUrlAccessRead,
		Update: resourceWafProfileUrlAccessUpdate,
		Delete: resourceWafProfileUrlAccessDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_pattern": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"regex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
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

func resourceWafProfileUrlAccessCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileUrlAccess(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfileUrlAccess resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfileUrlAccess(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfileUrlAccess(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfileUrlAccess resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWafProfileUrlAccess(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfileUrlAccess resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWafProfileUrlAccessRead(d, m)
			} else {
				return fmt.Errorf("Error creating WafProfileUrlAccess resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileUrlAccessRead(d, m)
}

func resourceWafProfileUrlAccessUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileUrlAccess(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileUrlAccess resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileUrlAccess(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileUrlAccess resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileUrlAccessRead(d, m)
}

func resourceWafProfileUrlAccessDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileUrlAccess(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileUrlAccess resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileUrlAccessRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWafProfileUrlAccess(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileUrlAccess resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileUrlAccess(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileUrlAccess resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileUrlAccessAccessPattern2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternNegate2edl(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternPattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternRegex2edl(i["regex"], d, pre_append)
			tmp["regex"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Regex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWafProfileUrlAccessAccessPatternSrcaddr2edl(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WafProfileUrlAccess-AccessPattern-Srcaddr")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileUrlAccessAccessPatternId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileUrlAccessAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileUrlAccessId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileUrlAccess(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("access_pattern", flattenWafProfileUrlAccessAccessPattern2edl(o["access-pattern"], d, "access_pattern")); err != nil {
			if vv, ok := fortiAPIPatch(o["access-pattern"], "WafProfileUrlAccess-AccessPattern"); ok {
				if err = d.Set("access_pattern", vv); err != nil {
					return fmt.Errorf("Error reading access_pattern: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading access_pattern: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_pattern"); ok {
			if err = d.Set("access_pattern", flattenWafProfileUrlAccessAccessPattern2edl(o["access-pattern"], d, "access_pattern")); err != nil {
				if vv, ok := fortiAPIPatch(o["access-pattern"], "WafProfileUrlAccess-AccessPattern"); ok {
					if err = d.Set("access_pattern", vv); err != nil {
						return fmt.Errorf("Error reading access_pattern: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading access_pattern: %v", err)
				}
			}
		}
	}

	if err = d.Set("action", flattenWafProfileUrlAccessAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WafProfileUrlAccess-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("address", flattenWafProfileUrlAccessAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "WafProfileUrlAccess-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafProfileUrlAccessId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WafProfileUrlAccess-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenWafProfileUrlAccessLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WafProfileUrlAccess-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafProfileUrlAccessSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "WafProfileUrlAccess-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func flattenWafProfileUrlAccessFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileUrlAccessAccessPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWafProfileUrlAccessAccessPatternId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandWafProfileUrlAccessAccessPatternNegate2edl(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileUrlAccessAccessPatternPattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regex"], _ = expandWafProfileUrlAccessAccessPatternRegex2edl(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWafProfileUrlAccessAccessPatternSrcaddr2edl(d, i["srcaddr"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileUrlAccessAccessPatternId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileUrlAccessAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileUrlAccessId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileUrlAccess(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_pattern"); ok || d.HasChange("access_pattern") {
		t, err := expandWafProfileUrlAccessAccessPattern2edl(d, v, "access_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-pattern"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWafProfileUrlAccessAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandWafProfileUrlAccessAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWafProfileUrlAccessId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWafProfileUrlAccessLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandWafProfileUrlAccessSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	return &obj, nil
}
