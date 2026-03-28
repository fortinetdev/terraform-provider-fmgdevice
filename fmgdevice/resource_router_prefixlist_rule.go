// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv4 prefix list rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterPrefixListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPrefixListRuleCreate,
		Read:   resourceRouterPrefixListRuleRead,
		Update: resourceRouterPrefixListRuleUpdate,
		Delete: resourceRouterPrefixListRuleDelete,

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
			"prefix_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceRouterPrefixListRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	prefix_list := d.Get("prefix_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list"] = prefix_list

	obj, err := getObjectRouterPrefixListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterPrefixListRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterPrefixListRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterPrefixListRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterPrefixListRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterPrefixListRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterPrefixListRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterPrefixListRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterPrefixListRuleRead(d, m)
}

func resourceRouterPrefixListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	prefix_list := d.Get("prefix_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list"] = prefix_list

	obj, err := getObjectRouterPrefixListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterPrefixListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterPrefixListRuleRead(d, m)
}

func resourceRouterPrefixListRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	prefix_list := d.Get("prefix_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list"] = prefix_list

	wsParams["adom"] = adomv

	err = c.DeleteRouterPrefixListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixListRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	prefix_list := d.Get("prefix_list").(string)
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
	if prefix_list == "" {
		prefix_list = importOptionChecking(m.(*FortiClient).Cfg, "prefix_list")
		if prefix_list == "" {
			return fmt.Errorf("Parameter prefix_list is missing")
		}
		if err = d.Set("prefix_list", prefix_list); err != nil {
			return fmt.Errorf("Error set params prefix_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list"] = prefix_list

	o, err := c.ReadRouterPrefixListRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterPrefixListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixListRule resource from API: %v", err)
	}
	return nil
}

func flattenRouterPrefixListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleGe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRuleLe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixListRulePrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectRouterPrefixListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterPrefixListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterPrefixListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("flags", flattenRouterPrefixListRuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterPrefixListRule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("ge", flattenRouterPrefixListRuleGe2edl(o["ge"], d, "ge")); err != nil {
		if vv, ok := fortiAPIPatch(o["ge"], "RouterPrefixListRule-Ge"); ok {
			if err = d.Set("ge", vv); err != nil {
				return fmt.Errorf("Error reading ge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterPrefixListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterPrefixListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("le", flattenRouterPrefixListRuleLe2edl(o["le"], d, "le")); err != nil {
		if vv, ok := fortiAPIPatch(o["le"], "RouterPrefixListRule-Le"); ok {
			if err = d.Set("le", vv); err != nil {
				return fmt.Errorf("Error reading le: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterPrefixListRulePrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterPrefixListRule-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenRouterPrefixListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterPrefixListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleGe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRuleLe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixListRulePrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterPrefixListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterPrefixListRuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok || d.HasChange("ge") {
		t, err := expandRouterPrefixListRuleGe2edl(d, v, "ge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterPrefixListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("le"); ok || d.HasChange("le") {
		t, err := expandRouterPrefixListRuleLe2edl(d, v, "le")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterPrefixListRulePrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
