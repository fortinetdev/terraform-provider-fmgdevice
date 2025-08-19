// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSO admin users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoAdminCreate,
		Read:   resourceSystemSsoAdminRead,
		Update: resourceSystemSsoAdminUpdate,
		Delete: resourceSystemSsoAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gui_default_dashboard_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ignore_invalid_signature_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ignore_release_overview_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSsoAdminCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSsoAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoAdmin resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSsoAdmin(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoAdmin resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSsoAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSsoAdmin(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSsoAdmin(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSsoAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSsoAdminGuiDefaultDashboardTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiIgnoreInvalidSignatureVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiIgnoreReleaseOverviewVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoAdminVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSsoAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accprofile", flattenSystemSsoAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemSsoAdmin-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("gui_default_dashboard_template", flattenSystemSsoAdminGuiDefaultDashboardTemplate(o["gui-default-dashboard-template"], d, "gui_default_dashboard_template")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-default-dashboard-template"], "SystemSsoAdmin-GuiDefaultDashboardTemplate"); ok {
			if err = d.Set("gui_default_dashboard_template", vv); err != nil {
				return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
		}
	}

	if err = d.Set("gui_ignore_invalid_signature_version", flattenSystemSsoAdminGuiIgnoreInvalidSignatureVersion(o["gui-ignore-invalid-signature-version"], d, "gui_ignore_invalid_signature_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-invalid-signature-version"], "SystemSsoAdmin-GuiIgnoreInvalidSignatureVersion"); ok {
			if err = d.Set("gui_ignore_invalid_signature_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
		}
	}

	if err = d.Set("gui_ignore_release_overview_version", flattenSystemSsoAdminGuiIgnoreReleaseOverviewVersion(o["gui-ignore-release-overview-version"], d, "gui_ignore_release_overview_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-release-overview-version"], "SystemSsoAdmin-GuiIgnoreReleaseOverviewVersion"); ok {
			if err = d.Set("gui_ignore_release_overview_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSsoAdminName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSsoAdmin-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemSsoAdminVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemSsoAdmin-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemSsoAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSsoAdminAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSsoAdminGuiDefaultDashboardTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiIgnoreInvalidSignatureVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiIgnoreReleaseOverviewVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSsoAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemSsoAdminAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("gui_default_dashboard_template"); ok || d.HasChange("gui_default_dashboard_template") {
		t, err := expandSystemSsoAdminGuiDefaultDashboardTemplate(d, v, "gui_default_dashboard_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-default-dashboard-template"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_invalid_signature_version"); ok || d.HasChange("gui_ignore_invalid_signature_version") {
		t, err := expandSystemSsoAdminGuiIgnoreInvalidSignatureVersion(d, v, "gui_ignore_invalid_signature_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-invalid-signature-version"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_release_overview_version"); ok || d.HasChange("gui_ignore_release_overview_version") {
		t, err := expandSystemSsoAdminGuiIgnoreReleaseOverviewVersion(d, v, "gui_ignore_release_overview_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-release-overview-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSsoAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoAdminVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
