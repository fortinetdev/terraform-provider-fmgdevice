// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterAccessListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAccessListRuleCreate,
		Read:   resourceRouterAccessListRuleRead,
		Update: resourceRouterAccessListRuleUpdate,
		Delete: resourceRouterAccessListRuleDelete,

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
			"access_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterAccessListRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	access_list := d.Get("access_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list"] = access_list

	obj, err := getObjectRouterAccessListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterAccessListRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterAccessListRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterAccessListRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterAccessListRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterAccessListRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterAccessListRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterAccessListRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterAccessListRuleRead(d, m)
}

func resourceRouterAccessListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_list := d.Get("access_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list"] = access_list

	obj, err := getObjectRouterAccessListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterAccessListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterAccessListRuleRead(d, m)
}

func resourceRouterAccessListRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	access_list := d.Get("access_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list"] = access_list

	wsParams["adom"] = adomv

	err = c.DeleteRouterAccessListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAccessListRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_list := d.Get("access_list").(string)
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
	if access_list == "" {
		access_list = importOptionChecking(m.(*FortiClient).Cfg, "access_list")
		if access_list == "" {
			return fmt.Errorf("Parameter access_list is missing")
		}
		if err = d.Set("access_list", access_list); err != nil {
			return fmt.Errorf("Error set params access_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list"] = access_list

	o, err := c.ReadRouterAccessListRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterAccessListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessListRule resource from API: %v", err)
	}
	return nil
}

func flattenRouterAccessListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessListRuleExactMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessListRuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessListRulePrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenRouterAccessListRuleWildcard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectRouterAccessListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterAccessListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterAccessListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("exact_match", flattenRouterAccessListRuleExactMatch2edl(o["exact-match"], d, "exact_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["exact-match"], "RouterAccessListRule-ExactMatch"); ok {
			if err = d.Set("exact_match", vv); err != nil {
				return fmt.Errorf("Error reading exact_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exact_match: %v", err)
		}
	}

	if err = d.Set("flags", flattenRouterAccessListRuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterAccessListRule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterAccessListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterAccessListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterAccessListRulePrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterAccessListRule-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenRouterAccessListRuleWildcard2edl(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "RouterAccessListRule-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	return nil
}

func flattenRouterAccessListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterAccessListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleExactMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRulePrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleWildcard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectRouterAccessListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterAccessListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("exact_match"); ok || d.HasChange("exact_match") {
		t, err := expandRouterAccessListRuleExactMatch2edl(d, v, "exact_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact-match"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterAccessListRuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterAccessListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterAccessListRulePrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandRouterAccessListRuleWildcard2edl(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	return &obj, nil
}
