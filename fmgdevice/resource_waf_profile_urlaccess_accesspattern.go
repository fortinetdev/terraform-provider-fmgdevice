// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> URL access pattern.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileUrlAccessAccessPattern() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileUrlAccessAccessPatternCreate,
		Read:   resourceWafProfileUrlAccessAccessPatternRead,
		Update: resourceWafProfileUrlAccessAccessPatternUpdate,
		Delete: resourceWafProfileUrlAccessAccessPatternDelete,

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
			"url_access": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"negate": &schema.Schema{
				Type:     schema.TypeString,
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
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafProfileUrlAccessAccessPatternCreate(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	obj, err := getObjectWafProfileUrlAccessAccessPattern(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfileUrlAccessAccessPattern resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWafProfileUrlAccessAccessPattern(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWafProfileUrlAccessAccessPattern(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WafProfileUrlAccessAccessPattern resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWafProfileUrlAccessAccessPattern(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WafProfileUrlAccessAccessPattern resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWafProfileUrlAccessAccessPatternRead(d, m)
			} else {
				return fmt.Errorf("Error creating WafProfileUrlAccessAccessPattern resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileUrlAccessAccessPatternRead(d, m)
}

func resourceWafProfileUrlAccessAccessPatternUpdate(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	obj, err := getObjectWafProfileUrlAccessAccessPattern(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileUrlAccessAccessPattern resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileUrlAccessAccessPattern(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileUrlAccessAccessPattern resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWafProfileUrlAccessAccessPatternRead(d, m)
}

func resourceWafProfileUrlAccessAccessPatternDelete(d *schema.ResourceData, m interface{}) error {
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
	url_access := d.Get("url_access").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	wsParams["adom"] = adomv

	err = c.DeleteWafProfileUrlAccessAccessPattern(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileUrlAccessAccessPattern resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileUrlAccessAccessPatternRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
	url_access := d.Get("url_access").(string)
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
	if url_access == "" {
		url_access = importOptionChecking(m.(*FortiClient).Cfg, "url_access")
		if url_access == "" {
			return fmt.Errorf("Parameter url_access is missing")
		}
		if err = d.Set("url_access", url_access); err != nil {
			return fmt.Errorf("Error set params url_access: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["url_access"] = url_access

	o, err := c.ReadWafProfileUrlAccessAccessPattern(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileUrlAccessAccessPattern resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileUrlAccessAccessPattern(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileUrlAccessAccessPattern resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileUrlAccessAccessPatternId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternNegate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternRegex3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternSrcaddr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWafProfileUrlAccessAccessPattern(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWafProfileUrlAccessAccessPatternId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WafProfileUrlAccessAccessPattern-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("negate", flattenWafProfileUrlAccessAccessPatternNegate3rdl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "WafProfileUrlAccessAccessPattern-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	if err = d.Set("pattern", flattenWafProfileUrlAccessAccessPatternPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "WafProfileUrlAccessAccessPattern-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("regex", flattenWafProfileUrlAccessAccessPatternRegex3rdl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "WafProfileUrlAccessAccessPattern-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenWafProfileUrlAccessAccessPatternSrcaddr3rdl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "WafProfileUrlAccessAccessPattern-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	return nil
}

func flattenWafProfileUrlAccessAccessPatternFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileUrlAccessAccessPatternId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternNegate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternRegex3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternSrcaddr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWafProfileUrlAccessAccessPattern(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWafProfileUrlAccessAccessPatternId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandWafProfileUrlAccessAccessPatternNegate3rdl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandWafProfileUrlAccessAccessPatternPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandWafProfileUrlAccessAccessPatternRegex3rdl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandWafProfileUrlAccessAccessPatternSrcaddr3rdl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	return &obj, nil
}
