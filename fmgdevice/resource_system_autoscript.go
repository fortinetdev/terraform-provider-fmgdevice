// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure auto script.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoScriptCreate,
		Read:   resourceSystemAutoScriptRead,
		Update: resourceSystemAutoScriptUpdate,
		Delete: resourceSystemAutoScriptDelete,

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
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"output_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemAutoScriptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutoScript(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutoScript resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutoScript(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutoScript resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutoScriptRead(d, m)
}

func resourceSystemAutoScriptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAutoScript(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoScript resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoScript(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoScript resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutoScriptRead(d, m)
}

func resourceSystemAutoScriptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemAutoScript(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoScript resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoScriptRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutoScript(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoScript resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoScript(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoScript resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoScriptInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interval", flattenSystemAutoScriptInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "SystemAutoScript-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutoScriptName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutoScript-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("output_size", flattenSystemAutoScriptOutputSize(o["output-size"], d, "output_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["output-size"], "SystemAutoScript-OutputSize"); ok {
			if err = d.Set("output_size", vv); err != nil {
				return fmt.Errorf("Error reading output_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("repeat", flattenSystemAutoScriptRepeat(o["repeat"], d, "repeat")); err != nil {
		if vv, ok := fortiAPIPatch(o["repeat"], "SystemAutoScript-Repeat"); ok {
			if err = d.Set("repeat", vv); err != nil {
				return fmt.Errorf("Error reading repeat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemAutoScriptScript(o["script"], d, "script")); err != nil {
		if vv, ok := fortiAPIPatch(o["script"], "SystemAutoScript-Script"); ok {
			if err = d.Set("script", vv); err != nil {
				return fmt.Errorf("Error reading script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("start", flattenSystemAutoScriptStart(o["start"], d, "start")); err != nil {
		if vv, ok := fortiAPIPatch(o["start"], "SystemAutoScript-Start"); ok {
			if err = d.Set("start", vv); err != nil {
				return fmt.Errorf("Error reading start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemAutoScriptTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystemAutoScript-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoScriptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoScriptInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptOutputSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptRepeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoScript(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandSystemAutoScriptInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutoScriptName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("output_size"); ok || d.HasChange("output_size") {
		t, err := expandSystemAutoScriptOutputSize(d, v, "output_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-size"] = t
		}
	}

	if v, ok := d.GetOk("repeat"); ok || d.HasChange("repeat") {
		t, err := expandSystemAutoScriptRepeat(d, v, "repeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["repeat"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok || d.HasChange("script") {
		t, err := expandSystemAutoScriptScript(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok || d.HasChange("start") {
		t, err := expandSystemAutoScriptStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystemAutoScriptTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
