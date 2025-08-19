// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiCloud SSO admin users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoForticloudAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoForticloudAdminCreate,
		Read:   resourceSystemSsoForticloudAdminRead,
		Update: resourceSystemSsoForticloudAdminUpdate,
		Delete: resourceSystemSsoForticloudAdminDelete,

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

func resourceSystemSsoForticloudAdminCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSsoForticloudAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoForticloudAdmin resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSsoForticloudAdmin(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoForticloudAdmin resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSsoForticloudAdminRead(d, m)
}

func resourceSystemSsoForticloudAdminUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSsoForticloudAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoForticloudAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSsoForticloudAdmin(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoForticloudAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSsoForticloudAdminRead(d, m)
}

func resourceSystemSsoForticloudAdminDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSsoForticloudAdmin(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoForticloudAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoForticloudAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSsoForticloudAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoForticloudAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoForticloudAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoForticloudAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoForticloudAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSsoForticloudAdminGuiDefaultDashboardTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminGuiIgnoreInvalidSignatureVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminGuiIgnoreReleaseOverviewVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSsoForticloudAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accprofile", flattenSystemSsoForticloudAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemSsoForticloudAdmin-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("gui_default_dashboard_template", flattenSystemSsoForticloudAdminGuiDefaultDashboardTemplate(o["gui-default-dashboard-template"], d, "gui_default_dashboard_template")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-default-dashboard-template"], "SystemSsoForticloudAdmin-GuiDefaultDashboardTemplate"); ok {
			if err = d.Set("gui_default_dashboard_template", vv); err != nil {
				return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
		}
	}

	if err = d.Set("gui_ignore_invalid_signature_version", flattenSystemSsoForticloudAdminGuiIgnoreInvalidSignatureVersion(o["gui-ignore-invalid-signature-version"], d, "gui_ignore_invalid_signature_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-invalid-signature-version"], "SystemSsoForticloudAdmin-GuiIgnoreInvalidSignatureVersion"); ok {
			if err = d.Set("gui_ignore_invalid_signature_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
		}
	}

	if err = d.Set("gui_ignore_release_overview_version", flattenSystemSsoForticloudAdminGuiIgnoreReleaseOverviewVersion(o["gui-ignore-release-overview-version"], d, "gui_ignore_release_overview_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-release-overview-version"], "SystemSsoForticloudAdmin-GuiIgnoreReleaseOverviewVersion"); ok {
			if err = d.Set("gui_ignore_release_overview_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSsoForticloudAdminName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSsoForticloudAdmin-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemSsoForticloudAdminVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemSsoForticloudAdmin-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemSsoForticloudAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSsoForticloudAdminAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSsoForticloudAdminGuiDefaultDashboardTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminGuiIgnoreInvalidSignatureVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminGuiIgnoreReleaseOverviewVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSsoForticloudAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemSsoForticloudAdminAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("gui_default_dashboard_template"); ok || d.HasChange("gui_default_dashboard_template") {
		t, err := expandSystemSsoForticloudAdminGuiDefaultDashboardTemplate(d, v, "gui_default_dashboard_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-default-dashboard-template"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_invalid_signature_version"); ok || d.HasChange("gui_ignore_invalid_signature_version") {
		t, err := expandSystemSsoForticloudAdminGuiIgnoreInvalidSignatureVersion(d, v, "gui_ignore_invalid_signature_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-invalid-signature-version"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_release_overview_version"); ok || d.HasChange("gui_ignore_release_overview_version") {
		t, err := expandSystemSsoForticloudAdminGuiIgnoreReleaseOverviewVersion(d, v, "gui_ignore_release_overview_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-release-overview-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSsoForticloudAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoForticloudAdminVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
