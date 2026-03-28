// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 prefix lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterPrefixList6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPrefixList6Create,
		Read:   resourceRouterPrefixList6Read,
		Update: resourceRouterPrefixList6Update,
		Delete: resourceRouterPrefixList6Delete,

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

func resourceRouterPrefixList6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPrefixList6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterPrefixList6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterPrefixList6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterPrefixList6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateRouterPrefixList6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterPrefixList6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterPrefixList6Read(d, m)
}

func resourceRouterPrefixList6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPrefixList6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterPrefixList6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterPrefixList6Read(d, m)
}

func resourceRouterPrefixList6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterPrefixList6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixList6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterPrefixList6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterPrefixList6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterPrefixList6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6Rule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterPrefixList6RuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenRouterPrefixList6RuleFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := i["ge"]; ok {
			v := flattenRouterPrefixList6RuleGe(i["ge"], d, pre_append)
			tmp["ge"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Ge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterPrefixList6RuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := i["le"]; ok {
			v := flattenRouterPrefixList6RuleLe(i["le"], d, pre_append)
			tmp["le"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Le")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterPrefixList6RulePrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterPrefixList6-Rule-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterPrefixList6RuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleGe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleLe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RulePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterPrefixList6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenRouterPrefixList6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterPrefixList6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterPrefixList6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterPrefixList6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterPrefixList6Rule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "RouterPrefixList6-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterPrefixList6Rule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "RouterPrefixList6-Rule"); ok {
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

func flattenRouterPrefixList6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterPrefixList6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6Rule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandRouterPrefixList6RuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandRouterPrefixList6RuleFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ge"], _ = expandRouterPrefixList6RuleGe(d, i["ge"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterPrefixList6RuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["le"], _ = expandRouterPrefixList6RuleLe(d, i["le"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterPrefixList6RulePrefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterPrefixList6RuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleGe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleLe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RulePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterPrefixList6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterPrefixList6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterPrefixList6Rule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
