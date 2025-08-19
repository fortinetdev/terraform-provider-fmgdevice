// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automation setting configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAutomationSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutomationSettingUpdate,
		Read:   resourceAutomationSettingRead,
		Update: resourceAutomationSettingUpdate,
		Delete: resourceAutomationSettingDelete,

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
			"fabric_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_concurrent_stitches": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAutomationSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAutomationSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating AutomationSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateAutomationSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AutomationSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AutomationSetting")

	return resourceAutomationSettingRead(d, m)
}

func resourceAutomationSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAutomationSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AutomationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAutomationSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAutomationSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AutomationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAutomationSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AutomationSetting resource from API: %v", err)
	}
	return nil
}

func flattenAutomationSettingFabricSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAutomationSettingMaxConcurrentStitches(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAutomationSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fabric_sync", flattenAutomationSettingFabricSync(o["fabric-sync"], d, "fabric_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-sync"], "AutomationSetting-FabricSync"); ok {
			if err = d.Set("fabric_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_sync: %v", err)
		}
	}

	if err = d.Set("max_concurrent_stitches", flattenAutomationSettingMaxConcurrentStitches(o["max-concurrent-stitches"], d, "max_concurrent_stitches")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-stitches"], "AutomationSetting-MaxConcurrentStitches"); ok {
			if err = d.Set("max_concurrent_stitches", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_stitches: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_stitches: %v", err)
		}
	}

	return nil
}

func flattenAutomationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAutomationSettingFabricSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAutomationSettingMaxConcurrentStitches(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAutomationSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fabric_sync"); ok || d.HasChange("fabric_sync") {
		t, err := expandAutomationSettingFabricSync(d, v, "fabric_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-sync"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_stitches"); ok || d.HasChange("max_concurrent_stitches") {
		t, err := expandAutomationSettingMaxConcurrentStitches(d, v, "max_concurrent_stitches")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-stitches"] = t
		}
	}

	return &obj, nil
}
