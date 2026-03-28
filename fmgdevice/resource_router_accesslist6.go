// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 access lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterAccessList6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAccessList6Create,
		Read:   resourceRouterAccessList6Read,
		Update: resourceRouterAccessList6Update,
		Delete: resourceRouterAccessList6Delete,

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
						"exact_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceRouterAccessList6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterAccessList6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterAccessList6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterAccessList6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterAccessList6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterAccessList6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterAccessList6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterAccessList6Read(d, m)
}

func resourceRouterAccessList6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterAccessList6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterAccessList6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterAccessList6Read(d, m)
}

func resourceRouterAccessList6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterAccessList6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessList6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAccessList6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterAccessList6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterAccessList6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessList6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterAccessList6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6Rule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterAccessList6RuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "RouterAccessList6-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exact_match"
		if _, ok := i["exact-match"]; ok {
			v := flattenRouterAccessList6RuleExactMatch(i["exact-match"], d, pre_append)
			tmp["exact_match"] = fortiAPISubPartPatch(v, "RouterAccessList6-Rule-ExactMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenRouterAccessList6RuleFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "RouterAccessList6-Rule-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterAccessList6RuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterAccessList6-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterAccessList6RulePrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterAccessList6-Rule-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterAccessList6RuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleExactMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RulePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterAccessList6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenRouterAccessList6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterAccessList6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterAccessList6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterAccessList6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterAccessList6Rule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "RouterAccessList6-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterAccessList6Rule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "RouterAccessList6-Rule"); ok {
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

func flattenRouterAccessList6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterAccessList6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6Rule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandRouterAccessList6RuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exact_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exact-match"], _ = expandRouterAccessList6RuleExactMatch(d, i["exact_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandRouterAccessList6RuleFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterAccessList6RuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterAccessList6RulePrefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterAccessList6RuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleExactMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RulePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAccessList6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterAccessList6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterAccessList6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterAccessList6Rule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
