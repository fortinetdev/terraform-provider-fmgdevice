// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure custom gui themes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemTheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemThemeCreate,
		Read:   resourceSystemThemeRead,
		Update: resourceSystemThemeUpdate,
		Delete: resourceSystemThemeDelete,

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
			"accent_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"banner_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"banner_msg_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"border_radius": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_to_action_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"font": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"font_weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nav_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nav_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"selected_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"theme_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemThemeCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemTheme(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemTheme resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemTheme(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemTheme(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemTheme resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemTheme(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemTheme resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemThemeRead(d, m)
}

func resourceSystemThemeUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemTheme(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemTheme resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemTheme(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemTheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemThemeRead(d, m)
}

func resourceSystemThemeDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemTheme(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemTheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemThemeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemTheme(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemTheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemTheme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemTheme resource from API: %v", err)
	}
	return nil
}

func flattenSystemThemeAccentColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeBannerMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeBannerMsgSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeBaseTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeBorderRadius(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeCallToActionColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeFont(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeFontWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeHeaderColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeNavColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeNavStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeSelectedColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeTableStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemThemeThemeTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemTheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accent_color", flattenSystemThemeAccentColor(o["accent-color"], d, "accent_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["accent-color"], "SystemTheme-AccentColor"); ok {
			if err = d.Set("accent_color", vv); err != nil {
				return fmt.Errorf("Error reading accent_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accent_color: %v", err)
		}
	}

	if err = d.Set("banner_msg", flattenSystemThemeBannerMsg(o["banner-msg"], d, "banner_msg")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner-msg"], "SystemTheme-BannerMsg"); ok {
			if err = d.Set("banner_msg", vv); err != nil {
				return fmt.Errorf("Error reading banner_msg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner_msg: %v", err)
		}
	}

	if err = d.Set("banner_msg_severity", flattenSystemThemeBannerMsgSeverity(o["banner-msg-severity"], d, "banner_msg_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner-msg-severity"], "SystemTheme-BannerMsgSeverity"); ok {
			if err = d.Set("banner_msg_severity", vv); err != nil {
				return fmt.Errorf("Error reading banner_msg_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner_msg_severity: %v", err)
		}
	}

	if err = d.Set("base_theme", flattenSystemThemeBaseTheme(o["base-theme"], d, "base_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["base-theme"], "SystemTheme-BaseTheme"); ok {
			if err = d.Set("base_theme", vv); err != nil {
				return fmt.Errorf("Error reading base_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base_theme: %v", err)
		}
	}

	if err = d.Set("border_radius", flattenSystemThemeBorderRadius(o["border-radius"], d, "border_radius")); err != nil {
		if vv, ok := fortiAPIPatch(o["border-radius"], "SystemTheme-BorderRadius"); ok {
			if err = d.Set("border_radius", vv); err != nil {
				return fmt.Errorf("Error reading border_radius: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading border_radius: %v", err)
		}
	}

	if err = d.Set("call_to_action_color", flattenSystemThemeCallToActionColor(o["call-to-action-color"], d, "call_to_action_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-to-action-color"], "SystemTheme-CallToActionColor"); ok {
			if err = d.Set("call_to_action_color", vv); err != nil {
				return fmt.Errorf("Error reading call_to_action_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_to_action_color: %v", err)
		}
	}

	if err = d.Set("font", flattenSystemThemeFont(o["font"], d, "font")); err != nil {
		if vv, ok := fortiAPIPatch(o["font"], "SystemTheme-Font"); ok {
			if err = d.Set("font", vv); err != nil {
				return fmt.Errorf("Error reading font: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading font: %v", err)
		}
	}

	if err = d.Set("font_weight", flattenSystemThemeFontWeight(o["font-weight"], d, "font_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["font-weight"], "SystemTheme-FontWeight"); ok {
			if err = d.Set("font_weight", vv); err != nil {
				return fmt.Errorf("Error reading font_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading font_weight: %v", err)
		}
	}

	if err = d.Set("header_color", flattenSystemThemeHeaderColor(o["header-color"], d, "header_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-color"], "SystemTheme-HeaderColor"); ok {
			if err = d.Set("header_color", vv); err != nil {
				return fmt.Errorf("Error reading header_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_color: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemThemeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemTheme-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nav_color", flattenSystemThemeNavColor(o["nav-color"], d, "nav_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["nav-color"], "SystemTheme-NavColor"); ok {
			if err = d.Set("nav_color", vv); err != nil {
				return fmt.Errorf("Error reading nav_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nav_color: %v", err)
		}
	}

	if err = d.Set("nav_style", flattenSystemThemeNavStyle(o["nav-style"], d, "nav_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["nav-style"], "SystemTheme-NavStyle"); ok {
			if err = d.Set("nav_style", vv); err != nil {
				return fmt.Errorf("Error reading nav_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nav_style: %v", err)
		}
	}

	if err = d.Set("selected_color", flattenSystemThemeSelectedColor(o["selected-color"], d, "selected_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["selected-color"], "SystemTheme-SelectedColor"); ok {
			if err = d.Set("selected_color", vv); err != nil {
				return fmt.Errorf("Error reading selected_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selected_color: %v", err)
		}
	}

	if err = d.Set("table_style", flattenSystemThemeTableStyle(o["table-style"], d, "table_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["table-style"], "SystemTheme-TableStyle"); ok {
			if err = d.Set("table_style", vv); err != nil {
				return fmt.Errorf("Error reading table_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading table_style: %v", err)
		}
	}

	if err = d.Set("theme_template", flattenSystemThemeThemeTemplate(o["theme-template"], d, "theme_template")); err != nil {
		if vv, ok := fortiAPIPatch(o["theme-template"], "SystemTheme-ThemeTemplate"); ok {
			if err = d.Set("theme_template", vv); err != nil {
				return fmt.Errorf("Error reading theme_template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading theme_template: %v", err)
		}
	}

	return nil
}

func flattenSystemThemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemThemeAccentColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBannerMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBannerMsgSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBaseTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBorderRadius(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeCallToActionColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeFont(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeFontWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeHeaderColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeNavColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeNavStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeSelectedColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeTableStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeThemeTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemTheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accent_color"); ok || d.HasChange("accent_color") {
		t, err := expandSystemThemeAccentColor(d, v, "accent_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accent-color"] = t
		}
	}

	if v, ok := d.GetOk("banner_msg"); ok || d.HasChange("banner_msg") {
		t, err := expandSystemThemeBannerMsg(d, v, "banner_msg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner-msg"] = t
		}
	}

	if v, ok := d.GetOk("banner_msg_severity"); ok || d.HasChange("banner_msg_severity") {
		t, err := expandSystemThemeBannerMsgSeverity(d, v, "banner_msg_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner-msg-severity"] = t
		}
	}

	if v, ok := d.GetOk("base_theme"); ok || d.HasChange("base_theme") {
		t, err := expandSystemThemeBaseTheme(d, v, "base_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base-theme"] = t
		}
	}

	if v, ok := d.GetOk("border_radius"); ok || d.HasChange("border_radius") {
		t, err := expandSystemThemeBorderRadius(d, v, "border_radius")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-radius"] = t
		}
	}

	if v, ok := d.GetOk("call_to_action_color"); ok || d.HasChange("call_to_action_color") {
		t, err := expandSystemThemeCallToActionColor(d, v, "call_to_action_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-to-action-color"] = t
		}
	}

	if v, ok := d.GetOk("font"); ok || d.HasChange("font") {
		t, err := expandSystemThemeFont(d, v, "font")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font"] = t
		}
	}

	if v, ok := d.GetOk("font_weight"); ok || d.HasChange("font_weight") {
		t, err := expandSystemThemeFontWeight(d, v, "font_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-weight"] = t
		}
	}

	if v, ok := d.GetOk("header_color"); ok || d.HasChange("header_color") {
		t, err := expandSystemThemeHeaderColor(d, v, "header_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-color"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemThemeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nav_color"); ok || d.HasChange("nav_color") {
		t, err := expandSystemThemeNavColor(d, v, "nav_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nav-color"] = t
		}
	}

	if v, ok := d.GetOk("nav_style"); ok || d.HasChange("nav_style") {
		t, err := expandSystemThemeNavStyle(d, v, "nav_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nav-style"] = t
		}
	}

	if v, ok := d.GetOk("selected_color"); ok || d.HasChange("selected_color") {
		t, err := expandSystemThemeSelectedColor(d, v, "selected_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-color"] = t
		}
	}

	if v, ok := d.GetOk("table_style"); ok || d.HasChange("table_style") {
		t, err := expandSystemThemeTableStyle(d, v, "table_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-style"] = t
		}
	}

	if v, ok := d.GetOk("theme_template"); ok || d.HasChange("theme_template") {
		t, err := expandSystemThemeThemeTemplate(d, v, "theme_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme-template"] = t
		}
	}

	return &obj, nil
}
