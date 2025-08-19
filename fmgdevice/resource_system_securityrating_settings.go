// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for Security Rating.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSecurityRatingSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSecurityRatingSettingsUpdate,
		Read:   resourceSystemSecurityRatingSettingsRead,
		Update: resourceSystemSecurityRatingSettingsUpdate,
		Delete: resourceSystemSecurityRatingSettingsDelete,

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
			"override_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSecurityRatingSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSecurityRatingSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSecurityRatingSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSecurityRatingSettings")

	return resourceSystemSecurityRatingSettingsRead(d, m)
}

func resourceSystemSecurityRatingSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSecurityRatingSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSecurityRatingSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSecurityRatingSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSecurityRatingSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSecurityRatingSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemSecurityRatingSettingsOverrideSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSecurityRatingSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("override_sync", flattenSystemSecurityRatingSettingsOverrideSync(o["override-sync"], d, "override_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-sync"], "SystemSecurityRatingSettings-OverrideSync"); ok {
			if err = d.Set("override_sync", vv); err != nil {
				return fmt.Errorf("Error reading override_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_sync: %v", err)
		}
	}

	return nil
}

func flattenSystemSecurityRatingSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSecurityRatingSettingsOverrideSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSecurityRatingSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override_sync"); ok || d.HasChange("override_sync") {
		t, err := expandSystemSecurityRatingSettingsOverrideSync(d, v, "override_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-sync"] = t
		}
	}

	return &obj, nil
}
