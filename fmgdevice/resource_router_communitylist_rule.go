// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Community list rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterCommunityListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterCommunityListRuleCreate,
		Read:   resourceRouterCommunityListRuleRead,
		Update: resourceRouterCommunityListRuleUpdate,
		Delete: resourceRouterCommunityListRuleDelete,

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
			"community_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceRouterCommunityListRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	community_list := d.Get("community_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community_list"] = community_list

	obj, err := getObjectRouterCommunityListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterCommunityListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadRouterCommunityListRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateRouterCommunityListRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating RouterCommunityListRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateRouterCommunityListRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating RouterCommunityListRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceRouterCommunityListRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating RouterCommunityListRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterCommunityListRuleRead(d, m)
}

func resourceRouterCommunityListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	community_list := d.Get("community_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community_list"] = community_list

	obj, err := getObjectRouterCommunityListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateRouterCommunityListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterCommunityListRuleRead(d, m)
}

func resourceRouterCommunityListRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	community_list := d.Get("community_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community_list"] = community_list

	wsParams["adom"] = adomv

	err = c.DeleteRouterCommunityListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterCommunityListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterCommunityListRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	community_list := d.Get("community_list").(string)
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
	if community_list == "" {
		community_list = importOptionChecking(m.(*FortiClient).Cfg, "community_list")
		if community_list == "" {
			return fmt.Errorf("Parameter community_list is missing")
		}
		if err = d.Set("community_list", community_list); err != nil {
			return fmt.Errorf("Error set params community_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community_list"] = community_list

	o, err := c.ReadRouterCommunityListRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading RouterCommunityListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterCommunityListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterCommunityListRule resource from API: %v", err)
	}
	return nil
}

func flattenRouterCommunityListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleRegexp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterCommunityListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenRouterCommunityListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "RouterCommunityListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterCommunityListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterCommunityListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match", flattenRouterCommunityListRuleMatch2edl(o["match"], d, "match")); err != nil {
		if vv, ok := fortiAPIPatch(o["match"], "RouterCommunityListRule-Match"); ok {
			if err = d.Set("match", vv); err != nil {
				return fmt.Errorf("Error reading match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match: %v", err)
		}
	}

	if err = d.Set("regexp", flattenRouterCommunityListRuleRegexp2edl(o["regexp"], d, "regexp")); err != nil {
		if vv, ok := fortiAPIPatch(o["regexp"], "RouterCommunityListRule-Regexp"); ok {
			if err = d.Set("regexp", vv); err != nil {
				return fmt.Errorf("Error reading regexp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regexp: %v", err)
		}
	}

	return nil
}

func flattenRouterCommunityListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterCommunityListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleRegexp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterCommunityListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandRouterCommunityListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterCommunityListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandRouterCommunityListRuleMatch2edl(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("regexp"); ok || d.HasChange("regexp") {
		t, err := expandRouterCommunityListRuleRegexp2edl(d, v, "regexp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regexp"] = t
		}
	}

	return &obj, nil
}
