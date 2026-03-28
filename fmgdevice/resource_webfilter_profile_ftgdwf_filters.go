// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> FortiGuard filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileFtgdWfFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileFtgdWfFiltersCreate,
		Read:   resourceWebfilterProfileFtgdWfFiltersRead,
		Update: resourceWebfilterProfileFtgdWfFiltersUpdate,
		Delete: resourceWebfilterProfileFtgdWfFiltersDelete,

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
			"auth_usr_grp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
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
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warn_duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warning_duration_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"warning_prompt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterProfileFtgdWfFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileFtgdWfFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfileFtgdWfFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterProfileFtgdWfFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterProfileFtgdWfFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterProfileFtgdWfFilters resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWebfilterProfileFtgdWfFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterProfileFtgdWfFilters resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWebfilterProfileFtgdWfFiltersRead(d, m)
			} else {
				return fmt.Errorf("Error creating WebfilterProfileFtgdWfFilters resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterProfileFtgdWfFiltersRead(d, m)
}

func resourceWebfilterProfileFtgdWfFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileFtgdWfFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWfFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileFtgdWfFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWfFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWebfilterProfileFtgdWfFiltersRead(d, m)
}

func resourceWebfilterProfileFtgdWfFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterProfileFtgdWfFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileFtgdWfFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileFtgdWfFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterProfileFtgdWfFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileFtgdWfFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileFtgdWfFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileFtgdWfFilters resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileFtgdWfFiltersAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersCategory3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarnDuration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfileFtgdWfFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWebfilterProfileFtgdWfFiltersAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WebfilterProfileFtgdWfFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("auth_usr_grp", flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(o["auth-usr-grp"], d, "auth_usr_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-usr-grp"], "WebfilterProfileFtgdWfFilters-AuthUsrGrp"); ok {
			if err = d.Set("auth_usr_grp", vv); err != nil {
				return fmt.Errorf("Error reading auth_usr_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_usr_grp: %v", err)
		}
	}

	if err = d.Set("category", flattenWebfilterProfileFtgdWfFiltersCategory3rdl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "WebfilterProfileFtgdWfFilters-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterProfileFtgdWfFiltersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WebfilterProfileFtgdWfFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenWebfilterProfileFtgdWfFiltersLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "WebfilterProfileFtgdWfFilters-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("override_replacemsg", flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(o["override-replacemsg"], d, "override_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-replacemsg"], "WebfilterProfileFtgdWfFilters-OverrideReplacemsg"); ok {
			if err = d.Set("override_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading override_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_replacemsg: %v", err)
		}
	}

	if err = d.Set("warn_duration", flattenWebfilterProfileFtgdWfFiltersWarnDuration3rdl(o["warn-duration"], d, "warn_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["warn-duration"], "WebfilterProfileFtgdWfFilters-WarnDuration"); ok {
			if err = d.Set("warn_duration", vv); err != nil {
				return fmt.Errorf("Error reading warn_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warn_duration: %v", err)
		}
	}

	if err = d.Set("warning_duration_type", flattenWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(o["warning-duration-type"], d, "warning_duration_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-duration-type"], "WebfilterProfileFtgdWfFilters-WarningDurationType"); ok {
			if err = d.Set("warning_duration_type", vv); err != nil {
				return fmt.Errorf("Error reading warning_duration_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_duration_type: %v", err)
		}
	}

	if err = d.Set("warning_prompt", flattenWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(o["warning-prompt"], d, "warning_prompt")); err != nil {
		if vv, ok := fortiAPIPatch(o["warning-prompt"], "WebfilterProfileFtgdWfFilters-WarningPrompt"); ok {
			if err = d.Set("warning_prompt", vv); err != nil {
				return fmt.Errorf("Error reading warning_prompt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warning_prompt: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileFtgdWfFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileFtgdWfFiltersAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersCategory3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarnDuration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfileFtgdWfFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWebfilterProfileFtgdWfFiltersAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("auth_usr_grp"); ok || d.HasChange("auth_usr_grp") {
		t, err := expandWebfilterProfileFtgdWfFiltersAuthUsrGrp3rdl(d, v, "auth_usr_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-usr-grp"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandWebfilterProfileFtgdWfFiltersCategory3rdl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWebfilterProfileFtgdWfFiltersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandWebfilterProfileFtgdWfFiltersLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("override_replacemsg"); ok || d.HasChange("override_replacemsg") {
		t, err := expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg3rdl(d, v, "override_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("warn_duration"); ok || d.HasChange("warn_duration") {
		t, err := expandWebfilterProfileFtgdWfFiltersWarnDuration3rdl(d, v, "warn_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-duration"] = t
		}
	}

	if v, ok := d.GetOk("warning_duration_type"); ok || d.HasChange("warning_duration_type") {
		t, err := expandWebfilterProfileFtgdWfFiltersWarningDurationType3rdl(d, v, "warning_duration_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-duration-type"] = t
		}
	}

	if v, ok := d.GetOk("warning_prompt"); ok || d.HasChange("warning_prompt") {
		t, err := expandWebfilterProfileFtgdWfFiltersWarningPrompt3rdl(d, v, "warning_prompt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warning-prompt"] = t
		}
	}

	return &obj, nil
}
