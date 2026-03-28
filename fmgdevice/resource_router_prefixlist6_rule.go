// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 prefix list rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterPrefixList6Rule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPrefixList6RuleCreate,
		Read:   resourceRouterPrefixList6RuleRead,
		Update: resourceRouterPrefixList6RuleUpdate,
		Delete: resourceRouterPrefixList6RuleDelete,

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
			"prefix_list6": &schema.Schema{
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
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRouterPrefixList6RuleCreate(d *schema.ResourceData, m interface{}) error {
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
	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list6"] = prefix_list6

	obj, err := getObjectRouterPrefixList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6Rule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterPrefixList6Rule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterPrefixList6Rule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterPrefixList6Rule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterPrefixList6Rule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterPrefixList6Rule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterPrefixList6RuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterPrefixList6Rule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterPrefixList6RuleRead(d, m)
}

func resourceRouterPrefixList6RuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list6"] = prefix_list6

	obj, err := getObjectRouterPrefixList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6Rule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterPrefixList6Rule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6Rule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterPrefixList6RuleRead(d, m)
}

func resourceRouterPrefixList6RuleDelete(d *schema.ResourceData, m interface{}) error {
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
	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list6"] = prefix_list6

	wsParams["adom"] = adomv

	err = c.DeleteRouterPrefixList6Rule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList6Rule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixList6RuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	prefix_list6 := d.Get("prefix_list6").(string)
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
	if prefix_list6 == "" {
		prefix_list6 = importOptionChecking(m.(*FortiClient).Cfg, "prefix_list6")
		if prefix_list6 == "" {
			return fmt.Errorf("Parameter prefix_list6 is missing")
		}
		if err = d.Set("prefix_list6", prefix_list6); err != nil {
			return fmt.Errorf("Error set params prefix_list6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["prefix_list6"] = prefix_list6

	o, err := c.ReadRouterPrefixList6Rule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterPrefixList6Rule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList6Rule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6Rule resource from API: %v", err)
	}
	return nil
}

func flattenRouterPrefixList6RuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleGe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleLe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPrefixList6RulePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterPrefixList6Rule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterPrefixList6RuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterPrefixList6Rule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("flags", flattenRouterPrefixList6RuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterPrefixList6Rule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("ge", flattenRouterPrefixList6RuleGe2edl(o["ge"], d, "ge")); err != nil {
		if vv, ok := fortiAPIPatch(o["ge"], "RouterPrefixList6Rule-Ge"); ok {
			if err = d.Set("ge", vv); err != nil {
				return fmt.Errorf("Error reading ge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterPrefixList6RuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterPrefixList6Rule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("le", flattenRouterPrefixList6RuleLe2edl(o["le"], d, "le")); err != nil {
		if vv, ok := fortiAPIPatch(o["le"], "RouterPrefixList6Rule-Le"); ok {
			if err = d.Set("le", vv); err != nil {
				return fmt.Errorf("Error reading le: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterPrefixList6RulePrefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterPrefixList6Rule-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenRouterPrefixList6RuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterPrefixList6RuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleGe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleLe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RulePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList6Rule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterPrefixList6RuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterPrefixList6RuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok || d.HasChange("ge") {
		t, err := expandRouterPrefixList6RuleGe2edl(d, v, "ge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterPrefixList6RuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("le"); ok || d.HasChange("le") {
		t, err := expandRouterPrefixList6RuleLe2edl(d, v, "le")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterPrefixList6RulePrefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
