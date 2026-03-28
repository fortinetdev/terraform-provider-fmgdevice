// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 prefix lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterPrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPrefixListCreate,
		Read:   resourceRouterPrefixListRead,
		Update: resourceRouterPrefixListUpdate,
		Delete: resourceRouterPrefixListDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ge": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"le": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterPrefixListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterPrefixList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterPrefixList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterPrefixList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterPrefixList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterPrefixList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterPrefixListRead(d, m)
}

func resourceRouterPrefixListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterPrefixList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterPrefixListRead(d, m)
}

func resourceRouterPrefixListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterPrefixList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterPrefixList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterPrefixList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList resource from API: %v", err)
	}
	return nil
}

func flattenRouterPrefixListComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterPrefixListRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenRouterPrefixListRuleFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := i["ge"]; ok {
			v := flattenRouterPrefixListRuleGe(i["ge"], d, pre_append)
			tmp["ge"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Ge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterPrefixListRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := i["le"]; ok {
			v := flattenRouterPrefixListRuleLe(i["le"], d, pre_append)
			tmp["le"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Le")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterPrefixListRulePrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterPrefixList-Rule-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterPrefixListRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleGe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleLe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRulePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectRouterPrefixList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenRouterPrefixListComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterPrefixList-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterPrefixListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterPrefixList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterPrefixListRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "RouterPrefixList-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterPrefixListRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "RouterPrefixList-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterPrefixListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterPrefixListComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandRouterPrefixListRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandRouterPrefixListRuleFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ge"], _ = expandRouterPrefixListRuleGe(d, i["ge"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterPrefixListRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["le"], _ = expandRouterPrefixListRuleLe(d, i["le"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterPrefixListRulePrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterPrefixListRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleGe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleLe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRulePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterPrefixListComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterPrefixListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterPrefixListRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
