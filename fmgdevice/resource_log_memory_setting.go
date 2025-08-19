// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for memory buffer.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemorySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemorySettingUpdate,
		Read:   resourceLogMemorySettingRead,
		Update: resourceLogMemorySettingUpdate,
		Delete: resourceLogMemorySettingDelete,

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
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogMemorySettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectLogMemorySetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemorySetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogMemorySetting")

	return resourceLogMemorySettingRead(d, m)
}

func resourceLogMemorySettingDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteLogMemorySetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogMemorySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemorySettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogMemorySetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemorySetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource from API: %v", err)
	}
	return nil
}

func flattenLogMemorySettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemorySettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogMemorySetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("diskfull", flattenLogMemorySettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskfull"], "LogMemorySetting-Diskfull"); ok {
			if err = d.Set("diskfull", vv); err != nil {
				return fmt.Errorf("Error reading diskfull: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("status", flattenLogMemorySettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogMemorySetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenLogMemorySettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogMemorySettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemorySettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemorySetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("diskfull"); ok || d.HasChange("diskfull") {
		t, err := expandLogMemorySettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogMemorySettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
