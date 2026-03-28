// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideofilterProfileFortiguardCategoryFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceVideofilterProfileFortiguardCategoryFiltersCreate,
		Read:   resourceVideofilterProfileFortiguardCategoryFiltersRead,
		Update: resourceVideofilterProfileFortiguardCategoryFiltersUpdate,
		Delete: resourceVideofilterProfileFortiguardCategoryFiltersDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVideofilterProfileFortiguardCategoryFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVideofilterProfileFortiguardCategoryFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating VideofilterProfileFortiguardCategoryFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVideofilterProfileFortiguardCategoryFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVideofilterProfileFortiguardCategoryFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VideofilterProfileFortiguardCategoryFilters resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateVideofilterProfileFortiguardCategoryFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VideofilterProfileFortiguardCategoryFilters resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceVideofilterProfileFortiguardCategoryFiltersRead(d, m)
			} else {
				return fmt.Errorf("Error creating VideofilterProfileFortiguardCategoryFilters resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVideofilterProfileFortiguardCategoryFiltersRead(d, m)
}

func resourceVideofilterProfileFortiguardCategoryFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVideofilterProfileFortiguardCategoryFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfileFortiguardCategoryFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVideofilterProfileFortiguardCategoryFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVideofilterProfileFortiguardCategoryFiltersRead(d, m)
}

func resourceVideofilterProfileFortiguardCategoryFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVideofilterProfileFortiguardCategoryFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterProfileFortiguardCategoryFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVideofilterProfileFortiguardCategoryFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVideofilterProfileFortiguardCategoryFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterProfileFortiguardCategoryFilters resource from API: %v", err)
	}
	return nil
}

func flattenVideofilterProfileFortiguardCategoryFiltersAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenVideofilterProfileFortiguardCategoryFiltersAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "VideofilterProfileFortiguardCategoryFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("category_id", flattenVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(o["category-id"], d, "category_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["category-id"], "VideofilterProfileFortiguardCategoryFilters-CategoryId"); ok {
			if err = d.Set("category_id", vv); err != nil {
				return fmt.Errorf("Error reading category_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category_id: %v", err)
		}
	}

	if err = d.Set("fosid", flattenVideofilterProfileFortiguardCategoryFiltersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VideofilterProfileFortiguardCategoryFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenVideofilterProfileFortiguardCategoryFiltersLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "VideofilterProfileFortiguardCategoryFilters-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	return nil
}

func flattenVideofilterProfileFortiguardCategoryFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVideofilterProfileFortiguardCategoryFiltersAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandVideofilterProfileFortiguardCategoryFiltersAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("category_id"); ok || d.HasChange("category_id") {
		t, err := expandVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(d, v, "category_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category-id"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVideofilterProfileFortiguardCategoryFiltersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandVideofilterProfileFortiguardCategoryFiltersLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	return &obj, nil
}
