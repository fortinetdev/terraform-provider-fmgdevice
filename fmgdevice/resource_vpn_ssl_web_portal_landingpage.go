// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Landing page options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebPortalLandingPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalLandingPageUpdate,
		Read:   resourceVpnSslWebPortalLandingPageRead,
		Update: resourceVpnSslWebPortalLandingPageUpdate,
		Delete: resourceVpnSslWebPortalLandingPageDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"form_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sso_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnSslWebPortalLandingPageUpdate(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	obj, err := getObjectVpnSslWebPortalLandingPage(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalLandingPage resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebPortalLandingPage(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalLandingPage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnSslWebPortalLandingPage")

	return resourceVpnSslWebPortalLandingPageRead(d, m)
}

func resourceVpnSslWebPortalLandingPageDelete(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	wsParams["adom"] = adomv

	err = c.DeleteVpnSslWebPortalLandingPage(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortalLandingPage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalLandingPageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	portal := d.Get("portal").(string)
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
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	o, err := c.ReadVpnSslWebPortalLandingPage(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebPortalLandingPage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortalLandingPage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortalLandingPage resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalLandingPageFormData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnSslWebPortalLandingPageFormDataName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnSslWebPortalLandingPage-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenVpnSslWebPortalLandingPageFormDataValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "VpnSslWebPortalLandingPage-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnSslWebPortalLandingPageFormDataName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoCredential2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortalLandingPage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("form_data", flattenVpnSslWebPortalLandingPageFormData2edl(o["form-data"], d, "form_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebPortalLandingPage-FormData"); ok {
				if err = d.Set("form_data", vv); err != nil {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenVpnSslWebPortalLandingPageFormData2edl(o["form-data"], d, "form_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["form-data"], "VpnSslWebPortalLandingPage-FormData"); ok {
					if err = d.Set("form_data", vv); err != nil {
						return fmt.Errorf("Error reading form_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("sso", flattenVpnSslWebPortalLandingPageSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "VpnSslWebPortalLandingPage-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("sso_credential", flattenVpnSslWebPortalLandingPageSsoCredential2edl(o["sso-credential"], d, "sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential"], "VpnSslWebPortalLandingPage-SsoCredential"); ok {
			if err = d.Set("sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential: %v", err)
		}
	}

	if err = d.Set("sso_username", flattenVpnSslWebPortalLandingPageSsoUsername2edl(o["sso-username"], d, "sso_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-username"], "VpnSslWebPortalLandingPage-SsoUsername"); ok {
			if err = d.Set("sso_username", vv); err != nil {
				return fmt.Errorf("Error reading sso_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_username: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnSslWebPortalLandingPageUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "VpnSslWebPortalLandingPage-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebPortalLandingPageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebPortalLandingPageFormData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnSslWebPortalLandingPageFormDataName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandVpnSslWebPortalLandingPageFormDataValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalLandingPageFormDataName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoCredential2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebPortalLandingPageSsoUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortalLandingPage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandVpnSslWebPortalLandingPageFormData2edl(d, v, "form_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandVpnSslWebPortalLandingPageSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential"); ok || d.HasChange("sso_credential") {
		t, err := expandVpnSslWebPortalLandingPageSsoCredential2edl(d, v, "sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("sso_password"); ok || d.HasChange("sso_password") {
		t, err := expandVpnSslWebPortalLandingPageSsoPassword2edl(d, v, "sso_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-password"] = t
		}
	}

	if v, ok := d.GetOk("sso_username"); ok || d.HasChange("sso_username") {
		t, err := expandVpnSslWebPortalLandingPageSsoUsername2edl(d, v, "sso_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandVpnSslWebPortalLandingPageUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
