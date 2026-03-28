// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDeviceUpgradeExemptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDeviceUpgradeExemptionsCreate,
		Read:   resourceSystemDeviceUpgradeExemptionsRead,
		Update: resourceSystemDeviceUpgradeExemptionsUpdate,
		Delete: resourceSystemDeviceUpgradeExemptionsDelete,

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
			"fortinet_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemDeviceUpgradeExemptionsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDeviceUpgradeExemptions(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgradeExemptions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDeviceUpgradeExemptions(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDeviceUpgradeExemptions(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemDeviceUpgradeExemptions resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemDeviceUpgradeExemptions(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemDeviceUpgradeExemptions resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDeviceUpgradeExemptionsRead(d, m)
}

func resourceSystemDeviceUpgradeExemptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDeviceUpgradeExemptions(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeExemptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemDeviceUpgradeExemptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeExemptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDeviceUpgradeExemptionsRead(d, m)
}

func resourceSystemDeviceUpgradeExemptionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDeviceUpgradeExemptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDeviceUpgradeExemptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDeviceUpgradeExemptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDeviceUpgradeExemptions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemDeviceUpgradeExemptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDeviceUpgradeExemptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgradeExemptions resource from API: %v", err)
	}
	return nil
}

func flattenSystemDeviceUpgradeExemptionsFortinetDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeExemptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeExemptionsVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDeviceUpgradeExemptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fortinet_device", flattenSystemDeviceUpgradeExemptionsFortinetDevice(o["fortinet-device"], d, "fortinet_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinet-device"], "SystemDeviceUpgradeExemptions-FortinetDevice"); ok {
			if err = d.Set("fortinet_device", vv); err != nil {
				return fmt.Errorf("Error reading fortinet_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinet_device: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDeviceUpgradeExemptionsId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDeviceUpgradeExemptions-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("version", flattenSystemDeviceUpgradeExemptionsVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "SystemDeviceUpgradeExemptions-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenSystemDeviceUpgradeExemptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDeviceUpgradeExemptionsFortinetDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeExemptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeExemptionsVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDeviceUpgradeExemptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortinet_device"); ok || d.HasChange("fortinet_device") {
		t, err := expandSystemDeviceUpgradeExemptionsFortinetDevice(d, v, "fortinet_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-device"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDeviceUpgradeExemptionsId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandSystemDeviceUpgradeExemptionsVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
