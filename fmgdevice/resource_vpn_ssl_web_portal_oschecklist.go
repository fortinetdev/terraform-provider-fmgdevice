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

func resourceVpnSslWebPortalOsCheckList() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalOsCheckListUpdate,
		Read:   resourceVpnSslWebPortalOsCheckListRead,
		Update: resourceVpnSslWebPortalOsCheckListUpdate,
		Delete: resourceVpnSslWebPortalOsCheckListDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"latest_patch_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"minor_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceVpnSslWebPortalOsCheckListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebPortalOsCheckList(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalOsCheckList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebPortalOsCheckList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalOsCheckList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnSslWebPortalOsCheckList")

	return resourceVpnSslWebPortalOsCheckListRead(d, m)
}

func resourceVpnSslWebPortalOsCheckListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslWebPortalOsCheckList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortalOsCheckList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalOsCheckListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebPortalOsCheckList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebPortalOsCheckList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortalOsCheckList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortalOsCheckList resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalOsCheckListAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListLatestPatchLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListMinorVersion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListTolerance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortalOsCheckList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenVpnSslWebPortalOsCheckListAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "VpnSslWebPortalOsCheckList-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("latest_patch_level", flattenVpnSslWebPortalOsCheckListLatestPatchLevel2edl(o["latest-patch-level"], d, "latest_patch_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["latest-patch-level"], "VpnSslWebPortalOsCheckList-LatestPatchLevel"); ok {
			if err = d.Set("latest_patch_level", vv); err != nil {
				return fmt.Errorf("Error reading latest_patch_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latest_patch_level: %v", err)
		}
	}

	if err = d.Set("minor_version", flattenVpnSslWebPortalOsCheckListMinorVersion2edl(o["minor-version"], d, "minor_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["minor-version"], "VpnSslWebPortalOsCheckList-MinorVersion"); ok {
			if err = d.Set("minor_version", vv); err != nil {
				return fmt.Errorf("Error reading minor_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minor_version: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnSslWebPortalOsCheckListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnSslWebPortalOsCheckList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tolerance", flattenVpnSslWebPortalOsCheckListTolerance2edl(o["tolerance"], d, "tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["tolerance"], "VpnSslWebPortalOsCheckList-Tolerance"); ok {
			if err = d.Set("tolerance", vv); err != nil {
				return fmt.Errorf("Error reading tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tolerance: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebPortalOsCheckListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebPortalOsCheckListAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListLatestPatchLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListMinorVersion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListTolerance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortalOsCheckList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandVpnSslWebPortalOsCheckListAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("latest_patch_level"); ok || d.HasChange("latest_patch_level") {
		t, err := expandVpnSslWebPortalOsCheckListLatestPatchLevel2edl(d, v, "latest_patch_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latest-patch-level"] = t
		}
	}

	if v, ok := d.GetOk("minor_version"); ok || d.HasChange("minor_version") {
		t, err := expandVpnSslWebPortalOsCheckListMinorVersion2edl(d, v, "minor_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minor-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnSslWebPortalOsCheckListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tolerance"); ok || d.HasChange("tolerance") {
		t, err := expandVpnSslWebPortalOsCheckListTolerance2edl(d, v, "tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tolerance"] = t
		}
	}

	return &obj, nil
}
