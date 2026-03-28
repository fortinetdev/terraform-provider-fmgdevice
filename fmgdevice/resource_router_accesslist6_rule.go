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

func resourceRouterAccessList6Rule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAccessList6RuleCreate,
		Read:   resourceRouterAccessList6RuleRead,
		Update: resourceRouterAccessList6RuleUpdate,
		Delete: resourceRouterAccessList6RuleDelete,

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
			"access_list6": &schema.Schema{
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
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRouterAccessList6RuleCreate(d *schema.ResourceData, m interface{}) error {
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
	access_list6 := d.Get("access_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list6"] = access_list6

	obj, err := getObjectRouterAccessList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList6Rule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterAccessList6Rule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterAccessList6Rule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterAccessList6Rule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterAccessList6Rule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterAccessList6Rule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterAccessList6RuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterAccessList6Rule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterAccessList6RuleRead(d, m)
}

func resourceRouterAccessList6RuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_list6 := d.Get("access_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list6"] = access_list6

	obj, err := getObjectRouterAccessList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6Rule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterAccessList6Rule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList6Rule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterAccessList6RuleRead(d, m)
}

func resourceRouterAccessList6RuleDelete(d *schema.ResourceData, m interface{}) error {
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
	access_list6 := d.Get("access_list6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list6"] = access_list6

	wsParams["adom"] = adomv

	err = c.DeleteRouterAccessList6Rule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessList6Rule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAccessList6RuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_list6 := d.Get("access_list6").(string)
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
	if access_list6 == "" {
		access_list6 = importOptionChecking(m.(*FortiClient).Cfg, "access_list6")
		if access_list6 == "" {
			return fmt.Errorf("Parameter access_list6 is missing")
		}
		if err = d.Set("access_list6", access_list6); err != nil {
			return fmt.Errorf("Error set params access_list6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_list6"] = access_list6

	o, err := c.ReadRouterAccessList6Rule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterAccessList6Rule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessList6Rule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList6Rule resource from API: %v", err)
	}
	return nil
}

func flattenRouterAccessList6RuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleExactMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAccessList6RulePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterAccessList6Rule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterAccessList6RuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterAccessList6Rule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("exact_match", flattenRouterAccessList6RuleExactMatch2edl(o["exact-match"], d, "exact_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["exact-match"], "RouterAccessList6Rule-ExactMatch"); ok {
			if err = d.Set("exact_match", vv); err != nil {
				return fmt.Errorf("Error reading exact_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exact_match: %v", err)
		}
	}

	if err = d.Set("flags", flattenRouterAccessList6RuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "RouterAccessList6Rule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterAccessList6RuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterAccessList6Rule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterAccessList6RulePrefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterAccessList6Rule-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenRouterAccessList6RuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterAccessList6RuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleExactMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessList6RulePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAccessList6Rule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterAccessList6RuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("exact_match"); ok || d.HasChange("exact_match") {
		t, err := expandRouterAccessList6RuleExactMatch2edl(d, v, "exact_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact-match"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandRouterAccessList6RuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterAccessList6RuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterAccessList6RulePrefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
