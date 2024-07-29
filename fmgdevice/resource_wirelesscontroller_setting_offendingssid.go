// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure offending SSID.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSettingOffendingSsid() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSettingOffendingSsidCreate,
		Read:   resourceWirelessControllerSettingOffendingSsidRead,
		Update: resourceWirelessControllerSettingOffendingSsidUpdate,
		Delete: resourceWirelessControllerSettingOffendingSsidDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ssid_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerSettingOffendingSsidCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWirelessControllerSettingOffendingSsid(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSettingOffendingSsid resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerSettingOffendingSsid(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSettingOffendingSsid resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSettingOffendingSsidRead(d, m)
}

func resourceWirelessControllerSettingOffendingSsidUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWirelessControllerSettingOffendingSsid(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSettingOffendingSsid resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSettingOffendingSsid(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSettingOffendingSsid resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSettingOffendingSsidRead(d, m)
}

func resourceWirelessControllerSettingOffendingSsidDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteWirelessControllerSettingOffendingSsid(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSettingOffendingSsid resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSettingOffendingSsidRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadWirelessControllerSettingOffendingSsid(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSettingOffendingSsid resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSettingOffendingSsid(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSettingOffendingSsid resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSettingOffendingSsidAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerSettingOffendingSsidId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidSsidPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSettingOffendingSsid(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenWirelessControllerSettingOffendingSsidAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "WirelessControllerSettingOffendingSsid-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWirelessControllerSettingOffendingSsidId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerSettingOffendingSsid-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ssid_pattern", flattenWirelessControllerSettingOffendingSsidSsidPattern2edl(o["ssid-pattern"], d, "ssid_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssid-pattern"], "WirelessControllerSettingOffendingSsid-SsidPattern"); ok {
			if err = d.Set("ssid_pattern", vv); err != nil {
				return fmt.Errorf("Error reading ssid_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssid_pattern: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSettingOffendingSsidFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSettingOffendingSsidAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerSettingOffendingSsidId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidSsidPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSettingOffendingSsid(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandWirelessControllerSettingOffendingSsidAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerSettingOffendingSsidId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ssid_pattern"); ok || d.HasChange("ssid_pattern") {
		t, err := expandWirelessControllerSettingOffendingSsidSsidPattern2edl(d, v, "ssid_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid-pattern"] = t
		}
	}

	return &obj, nil
}
