// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Web Filter override settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileOverrideUpdate,
		Read:   resourceWebfilterProfileOverrideRead,
		Update: resourceWebfilterProfileOverrideUpdate,
		Delete: resourceWebfilterProfileOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"parent_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ovrd_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_dur": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_dur_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebfilterProfileOverrideUpdate(d *schema.ResourceData, m interface{}) error {
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
	parent_profile := d.Get("parent_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = parent_profile

	obj, err := getObjectWebfilterProfileOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileOverride resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileOverride(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterProfileOverride")

	return resourceWebfilterProfileOverrideRead(d, m)
}

func resourceWebfilterProfileOverrideDelete(d *schema.ResourceData, m interface{}) error {
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
	parent_profile := d.Get("parent_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = parent_profile

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterProfileOverride(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	parent_profile := d.Get("parent_profile").(string)
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
	if parent_profile == "" {
		parent_profile = importOptionChecking(m.(*FortiClient).Cfg, "parent_profile")
		if parent_profile == "" {
			return fmt.Errorf("Parameter parent_profile is missing")
		}
		if err = d.Set("parent_profile", parent_profile); err != nil {
			return fmt.Errorf("Error set params parent_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = parent_profile

	o, err := c.ReadWebfilterProfileOverride(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileOverride resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileOverrideOvrdCookie2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDur2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDurMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdScope2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdUserGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileOverrideProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileOverrideProfileAttribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfileOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ovrd_cookie", flattenWebfilterProfileOverrideOvrdCookie2edl(o["ovrd-cookie"], d, "ovrd_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-cookie"], "WebfilterProfileOverride-OvrdCookie"); ok {
			if err = d.Set("ovrd_cookie", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_cookie: %v", err)
		}
	}

	if err = d.Set("ovrd_dur", flattenWebfilterProfileOverrideOvrdDur2edl(o["ovrd-dur"], d, "ovrd_dur")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-dur"], "WebfilterProfileOverride-OvrdDur"); ok {
			if err = d.Set("ovrd_dur", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_dur: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_dur: %v", err)
		}
	}

	if err = d.Set("ovrd_dur_mode", flattenWebfilterProfileOverrideOvrdDurMode2edl(o["ovrd-dur-mode"], d, "ovrd_dur_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-dur-mode"], "WebfilterProfileOverride-OvrdDurMode"); ok {
			if err = d.Set("ovrd_dur_mode", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_dur_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_dur_mode: %v", err)
		}
	}

	if err = d.Set("ovrd_scope", flattenWebfilterProfileOverrideOvrdScope2edl(o["ovrd-scope"], d, "ovrd_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-scope"], "WebfilterProfileOverride-OvrdScope"); ok {
			if err = d.Set("ovrd_scope", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_scope: %v", err)
		}
	}

	if err = d.Set("ovrd_user_group", flattenWebfilterProfileOverrideOvrdUserGroup2edl(o["ovrd-user-group"], d, "ovrd_user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-user-group"], "WebfilterProfileOverride-OvrdUserGroup"); ok {
			if err = d.Set("ovrd_user_group", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_user_group: %v", err)
		}
	}

	if err = d.Set("profile", flattenWebfilterProfileOverrideProfile2edl(o["profile"], d, "profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile"], "WebfilterProfileOverride-Profile"); ok {
			if err = d.Set("profile", vv); err != nil {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("profile_attribute", flattenWebfilterProfileOverrideProfileAttribute2edl(o["profile-attribute"], d, "profile_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-attribute"], "WebfilterProfileOverride-ProfileAttribute"); ok {
			if err = d.Set("profile_attribute", vv); err != nil {
				return fmt.Errorf("Error reading profile_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_attribute: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenWebfilterProfileOverrideProfileType2edl(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "WebfilterProfileOverride-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileOverrideOvrdCookie2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDur2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDurMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdUserGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileOverrideProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileOverrideProfileAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfileOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ovrd_cookie"); ok || d.HasChange("ovrd_cookie") {
		t, err := expandWebfilterProfileOverrideOvrdCookie2edl(d, v, "ovrd_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-cookie"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_dur"); ok || d.HasChange("ovrd_dur") {
		t, err := expandWebfilterProfileOverrideOvrdDur2edl(d, v, "ovrd_dur")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-dur"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_dur_mode"); ok || d.HasChange("ovrd_dur_mode") {
		t, err := expandWebfilterProfileOverrideOvrdDurMode2edl(d, v, "ovrd_dur_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-dur-mode"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_scope"); ok || d.HasChange("ovrd_scope") {
		t, err := expandWebfilterProfileOverrideOvrdScope2edl(d, v, "ovrd_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-scope"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_user_group"); ok || d.HasChange("ovrd_user_group") {
		t, err := expandWebfilterProfileOverrideOvrdUserGroup2edl(d, v, "ovrd_user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-user-group"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandWebfilterProfileOverrideProfile2edl(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_attribute"); ok || d.HasChange("profile_attribute") {
		t, err := expandWebfilterProfileOverrideProfileAttribute2edl(d, v, "profile_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-attribute"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok || d.HasChange("profile_type") {
		t, err := expandWebfilterProfileOverrideProfileType2edl(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	return &obj, nil
}
