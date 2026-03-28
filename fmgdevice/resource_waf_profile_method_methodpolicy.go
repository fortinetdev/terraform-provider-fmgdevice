// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> HTTP method policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileMethodMethodPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileMethodMethodPolicyCreate,
		Read:   resourceWafProfileMethodMethodPolicyRead,
		Update: resourceWafProfileMethodMethodPolicyUpdate,
		Delete: resourceWafProfileMethodMethodPolicyDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"allowed_methods": &schema.Schema{
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
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileMethodMethodPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileMethodMethodPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfileMethodMethodPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfileMethodMethodPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfileMethodMethodPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfileMethodMethodPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWafProfileMethodMethodPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfileMethodMethodPolicy resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWafProfileMethodMethodPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating WafProfileMethodMethodPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileMethodMethodPolicyRead(d, m)
}

func resourceWafProfileMethodMethodPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWafProfileMethodMethodPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileMethodMethodPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileMethodMethodPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileMethodMethodPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileMethodMethodPolicyRead(d, m)
}

func resourceWafProfileMethodMethodPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileMethodMethodPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileMethodMethodPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileMethodMethodPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWafProfileMethodMethodPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileMethodMethodPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileMethodMethodPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileMethodMethodPolicy resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileMethodMethodPolicyAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyAllowedMethods3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileMethodMethodPolicyId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileMethodMethodPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenWafProfileMethodMethodPolicyAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "WafProfileMethodMethodPolicy-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("allowed_methods", flattenWafProfileMethodMethodPolicyAllowedMethods3rdl(o["allowed-methods"], d, "allowed_methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-methods"], "WafProfileMethodMethodPolicy-AllowedMethods"); ok {
			if err = d.Set("allowed_methods", vv); err != nil {
				return fmt.Errorf("Error reading allowed_methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_methods: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafProfileMethodMethodPolicyId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WafProfileMethodMethodPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("pattern", flattenWafProfileMethodMethodPolicyPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "WafProfileMethodMethodPolicy-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenWafProfileMethodMethodPolicyRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "WafProfileMethodMethodPolicy-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	return nil
}

func flattenWafProfileMethodMethodPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileMethodMethodPolicyAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyAllowedMethods3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileMethodMethodPolicyId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileMethodMethodPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandWafProfileMethodMethodPolicyAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("allowed_methods"); ok || d.HasChange("allowed_methods") {
		t, err := expandWafProfileMethodMethodPolicyAllowedMethods3rdl(d, v, "allowed_methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-methods"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWafProfileMethodMethodPolicyId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandWafProfileMethodMethodPolicyPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandWafProfileMethodMethodPolicyRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	return &obj, nil
}
