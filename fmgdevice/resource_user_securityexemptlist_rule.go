// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure rules for exempting users from captive portal authentication.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserSecurityExemptListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSecurityExemptListRuleCreate,
		Read:   resourceUserSecurityExemptListRuleRead,
		Update: resourceUserSecurityExemptListRuleUpdate,
		Delete: resourceUserSecurityExemptListRuleDelete,

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
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dstaddr": &schema.Schema{
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
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	}
}

func resourceUserSecurityExemptListRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["security_exempt_list"] = security_exempt_list

	obj, err := getObjectUserSecurityExemptListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating UserSecurityExemptListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserSecurityExemptListRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserSecurityExemptListRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserSecurityExemptListRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateUserSecurityExemptListRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserSecurityExemptListRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceUserSecurityExemptListRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating UserSecurityExemptListRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserSecurityExemptListRuleRead(d, m)
}

func resourceUserSecurityExemptListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["security_exempt_list"] = security_exempt_list

	obj, err := getObjectUserSecurityExemptListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating UserSecurityExemptListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserSecurityExemptListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserSecurityExemptListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserSecurityExemptListRuleRead(d, m)
}

func resourceUserSecurityExemptListRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["security_exempt_list"] = security_exempt_list

	wsParams["adom"] = adomv

	err = c.DeleteUserSecurityExemptListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserSecurityExemptListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSecurityExemptListRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	security_exempt_list := d.Get("security_exempt_list").(string)
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
	if security_exempt_list == "" {
		security_exempt_list = importOptionChecking(m.(*FortiClient).Cfg, "security_exempt_list")
		if security_exempt_list == "" {
			return fmt.Errorf("Parameter security_exempt_list is missing")
		}
		if err = d.Set("security_exempt_list", security_exempt_list); err != nil {
			return fmt.Errorf("Error set params security_exempt_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["security_exempt_list"] = security_exempt_list

	o, err := c.ReadUserSecurityExemptListRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserSecurityExemptListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSecurityExemptListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserSecurityExemptListRule resource from API: %v", err)
	}
	return nil
}

func flattenUserSecurityExemptListRuleDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSecurityExemptListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSecurityExemptListRuleService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSecurityExemptListRuleSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectUserSecurityExemptListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstaddr", flattenUserSecurityExemptListRuleDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "UserSecurityExemptListRule-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserSecurityExemptListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserSecurityExemptListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("service", flattenUserSecurityExemptListRuleService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "UserSecurityExemptListRule-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenUserSecurityExemptListRuleSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "UserSecurityExemptListRule-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	return nil
}

func flattenUserSecurityExemptListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserSecurityExemptListRuleDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSecurityExemptListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSecurityExemptListRuleService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSecurityExemptListRuleSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserSecurityExemptListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandUserSecurityExemptListRuleDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserSecurityExemptListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandUserSecurityExemptListRuleService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandUserSecurityExemptListRuleSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	return &obj, nil
}
