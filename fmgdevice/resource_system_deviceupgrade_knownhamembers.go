// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDeviceUpgradeKnownHaMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDeviceUpgradeKnownHaMembersCreate,
		Read:   resourceSystemDeviceUpgradeKnownHaMembersRead,
		Update: resourceSystemDeviceUpgradeKnownHaMembersUpdate,
		Delete: resourceSystemDeviceUpgradeKnownHaMembersDelete,

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
			"device_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemDeviceUpgradeKnownHaMembersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_upgrade := d.Get("device_upgrade").(string)
	paradict["device"] = device_name
	paradict["device_upgrade"] = device_upgrade

	obj, err := getObjectSystemDeviceUpgradeKnownHaMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgradeKnownHaMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDeviceUpgradeKnownHaMembers(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgradeKnownHaMembers resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemDeviceUpgradeKnownHaMembersRead(d, m)
}

func resourceSystemDeviceUpgradeKnownHaMembersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_upgrade := d.Get("device_upgrade").(string)
	paradict["device"] = device_name
	paradict["device_upgrade"] = device_upgrade

	obj, err := getObjectSystemDeviceUpgradeKnownHaMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeKnownHaMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDeviceUpgradeKnownHaMembers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeKnownHaMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemDeviceUpgradeKnownHaMembersRead(d, m)
}

func resourceSystemDeviceUpgradeKnownHaMembersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_upgrade := d.Get("device_upgrade").(string)
	paradict["device"] = device_name
	paradict["device_upgrade"] = device_upgrade

	err = c.DeleteSystemDeviceUpgradeKnownHaMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDeviceUpgradeKnownHaMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDeviceUpgradeKnownHaMembersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_upgrade := d.Get("device_upgrade").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_upgrade == "" {
		device_upgrade = importOptionChecking(m.(*FortiClient).Cfg, "device_upgrade")
		if device_upgrade == "" {
			return fmt.Errorf("Parameter device_upgrade is missing")
		}
		if err = d.Set("device_upgrade", device_upgrade); err != nil {
			return fmt.Errorf("Error set params device_upgrade: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["device_upgrade"] = device_upgrade

	o, err := c.ReadSystemDeviceUpgradeKnownHaMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgradeKnownHaMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDeviceUpgradeKnownHaMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgradeKnownHaMembers resource from API: %v", err)
	}
	return nil
}

func flattenSystemDeviceUpgradeKnownHaMembersSerial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDeviceUpgradeKnownHaMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("serial", flattenSystemDeviceUpgradeKnownHaMembersSerial2edl(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemDeviceUpgradeKnownHaMembers-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemDeviceUpgradeKnownHaMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDeviceUpgradeKnownHaMembersSerial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDeviceUpgradeKnownHaMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemDeviceUpgradeKnownHaMembersSerial2edl(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
