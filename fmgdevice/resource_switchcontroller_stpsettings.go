// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch spanning tree protocol (STP).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerStpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStpSettingsUpdate,
		Read:   resourceSwitchControllerStpSettingsRead,
		Update: resourceSwitchControllerStpSettingsUpdate,
		Delete: resourceSwitchControllerStpSettingsDelete,

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
			"forward_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_hops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pending_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSwitchControllerStpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerStpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerStpSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerStpSettings")

	return resourceSwitchControllerStpSettingsRead(d, m)
}

func resourceSwitchControllerStpSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerStpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerStpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerStpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("forward_time", flattenSwitchControllerStpSettingsForwardTime(o["forward-time"], d, "forward_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-time"], "SwitchControllerStpSettings-ForwardTime"); ok {
			if err = d.Set("forward_time", vv); err != nil {
				return fmt.Errorf("Error reading forward_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_time: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSwitchControllerStpSettingsHelloTime(o["hello-time"], d, "hello_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-time"], "SwitchControllerStpSettings-HelloTime"); ok {
			if err = d.Set("hello_time", vv); err != nil {
				return fmt.Errorf("Error reading hello_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSwitchControllerStpSettingsMaxAge(o["max-age"], d, "max_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-age"], "SwitchControllerStpSettings-MaxAge"); ok {
			if err = d.Set("max_age", vv); err != nil {
				return fmt.Errorf("Error reading max_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSwitchControllerStpSettingsMaxHops(o["max-hops"], d, "max_hops")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-hops"], "SwitchControllerStpSettings-MaxHops"); ok {
			if err = d.Set("max_hops", vv); err != nil {
				return fmt.Errorf("Error reading max_hops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerStpSettingsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerStpSettings-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pending_timer", flattenSwitchControllerStpSettingsPendingTimer(o["pending-timer"], d, "pending_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["pending-timer"], "SwitchControllerStpSettings-PendingTimer"); ok {
			if err = d.Set("pending_timer", vv); err != nil {
				return fmt.Errorf("Error reading pending_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pending_timer: %v", err)
		}
	}

	if err = d.Set("revision", flattenSwitchControllerStpSettingsRevision(o["revision"], d, "revision")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision"], "SwitchControllerStpSettings-Revision"); ok {
			if err = d.Set("revision", vv); err != nil {
				return fmt.Errorf("Error reading revision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerStpSettingsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerStpSettings-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("forward_time"); ok || d.HasChange("forward_time") {
		t, err := expandSwitchControllerStpSettingsForwardTime(d, v, "forward_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-time"] = t
		}
	}

	if v, ok := d.GetOk("hello_time"); ok || d.HasChange("hello_time") {
		t, err := expandSwitchControllerStpSettingsHelloTime(d, v, "hello_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-time"] = t
		}
	}

	if v, ok := d.GetOk("max_age"); ok || d.HasChange("max_age") {
		t, err := expandSwitchControllerStpSettingsMaxAge(d, v, "max_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-age"] = t
		}
	}

	if v, ok := d.GetOk("max_hops"); ok || d.HasChange("max_hops") {
		t, err := expandSwitchControllerStpSettingsMaxHops(d, v, "max_hops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-hops"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerStpSettingsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pending_timer"); ok || d.HasChange("pending_timer") {
		t, err := expandSwitchControllerStpSettingsPendingTimer(d, v, "pending_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pending-timer"] = t
		}
	}

	if v, ok := d.GetOk("revision"); ok || d.HasChange("revision") {
		t, err := expandSwitchControllerStpSettingsRevision(d, v, "revision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerStpSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
