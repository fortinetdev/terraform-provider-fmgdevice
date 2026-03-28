// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure community lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterCommunityList() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterCommunityListCreate,
		Read:   resourceRouterCommunityListRead,
		Update: resourceRouterCommunityListUpdate,
		Delete: resourceRouterCommunityListDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"regexp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

func resourceRouterCommunityListCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectRouterCommunityList(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterCommunityList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterCommunityList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterCommunityList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterCommunityList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterCommunityList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterCommunityList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterCommunityListRead(d, m)
}

func resourceRouterCommunityListUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectRouterCommunityList(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterCommunityList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterCommunityListRead(d, m)
}

func resourceRouterCommunityListDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteRouterCommunityList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterCommunityList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterCommunityListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterCommunityList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterCommunityList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterCommunityList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterCommunityList resource from API: %v", err)
	}
	return nil
}

func flattenRouterCommunityListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenRouterCommunityListRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "RouterCommunityList-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterCommunityListRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterCommunityList-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := i["match"]; ok {
			v := flattenRouterCommunityListRuleMatch(i["match"], d, pre_append)
			tmp["match"] = fortiAPISubPartPatch(v, "RouterCommunityList-Rule-Match")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			v := flattenRouterCommunityListRuleRegexp(i["regexp"], d, pre_append)
			tmp["regexp"] = fortiAPISubPartPatch(v, "RouterCommunityList-Rule-Regexp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterCommunityListRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterCommunityList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenRouterCommunityListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterCommunityList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterCommunityListRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "RouterCommunityList-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterCommunityListRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "RouterCommunityList-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenRouterCommunityListType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "RouterCommunityList-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenRouterCommunityListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterCommunityListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandRouterCommunityListRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterCommunityListRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match"], _ = expandRouterCommunityListRuleMatch(d, i["match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regexp"], _ = expandRouterCommunityListRuleRegexp(d, i["regexp"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterCommunityListRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleRegexp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterCommunityList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterCommunityListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterCommunityListRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandRouterCommunityListType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
