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

func resourceSystemFederatedUpgradeKnownHaMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFederatedUpgradeKnownHaMembersCreate,
		Read:   resourceSystemFederatedUpgradeKnownHaMembersRead,
		Update: resourceSystemFederatedUpgradeKnownHaMembersUpdate,
		Delete: resourceSystemFederatedUpgradeKnownHaMembersDelete,

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
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemFederatedUpgradeKnownHaMembersCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemFederatedUpgradeKnownHaMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemFederatedUpgradeKnownHaMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemFederatedUpgradeKnownHaMembers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemFederatedUpgradeKnownHaMembers resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemFederatedUpgradeKnownHaMembersRead(d, m)
}

func resourceSystemFederatedUpgradeKnownHaMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemFederatedUpgradeKnownHaMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgradeKnownHaMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFederatedUpgradeKnownHaMembers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgradeKnownHaMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial"))

	return resourceSystemFederatedUpgradeKnownHaMembersRead(d, m)
}

func resourceSystemFederatedUpgradeKnownHaMembersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFederatedUpgradeKnownHaMembers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFederatedUpgradeKnownHaMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeKnownHaMembersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFederatedUpgradeKnownHaMembers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgradeKnownHaMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFederatedUpgradeKnownHaMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgradeKnownHaMembers resource from API: %v", err)
	}
	return nil
}

func flattenSystemFederatedUpgradeKnownHaMembersSerial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFederatedUpgradeKnownHaMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("serial", flattenSystemFederatedUpgradeKnownHaMembersSerial2edl(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemFederatedUpgradeKnownHaMembers-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemFederatedUpgradeKnownHaMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFederatedUpgradeKnownHaMembersSerial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFederatedUpgradeKnownHaMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemFederatedUpgradeKnownHaMembersSerial2edl(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
