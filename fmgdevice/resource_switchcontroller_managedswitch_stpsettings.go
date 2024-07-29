// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchStpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchStpSettingsUpdate,
		Read:   resourceSwitchControllerManagedSwitchStpSettingsRead,
		Update: resourceSwitchControllerManagedSwitchStpSettingsUpdate,
		Delete: resourceSwitchControllerManagedSwitchStpSettingsDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"local_override": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSwitchControllerManagedSwitchStpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchStpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchStpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchStpSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchStpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerManagedSwitchStpSettings")

	return resourceSwitchControllerManagedSwitchStpSettingsRead(d, m)
}

func resourceSwitchControllerManagedSwitchStpSettingsDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	err = c.DeleteSwitchControllerManagedSwitchStpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchStpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchStpSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchStpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchStpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchStpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchStpSettingsForwardTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsHelloTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsLocalOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxAge2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxHops2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsPendingTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsRevision2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchStpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("forward_time", flattenSwitchControllerManagedSwitchStpSettingsForwardTime2edl(o["forward-time"], d, "forward_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-time"], "SwitchControllerManagedSwitchStpSettings-ForwardTime"); ok {
			if err = d.Set("forward_time", vv); err != nil {
				return fmt.Errorf("Error reading forward_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_time: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSwitchControllerManagedSwitchStpSettingsHelloTime2edl(o["hello-time"], d, "hello_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-time"], "SwitchControllerManagedSwitchStpSettings-HelloTime"); ok {
			if err = d.Set("hello_time", vv); err != nil {
				return fmt.Errorf("Error reading hello_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("local_override", flattenSwitchControllerManagedSwitchStpSettingsLocalOverride2edl(o["local-override"], d, "local_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-override"], "SwitchControllerManagedSwitchStpSettings-LocalOverride"); ok {
			if err = d.Set("local_override", vv); err != nil {
				return fmt.Errorf("Error reading local_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_override: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSwitchControllerManagedSwitchStpSettingsMaxAge2edl(o["max-age"], d, "max_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-age"], "SwitchControllerManagedSwitchStpSettings-MaxAge"); ok {
			if err = d.Set("max_age", vv); err != nil {
				return fmt.Errorf("Error reading max_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSwitchControllerManagedSwitchStpSettingsMaxHops2edl(o["max-hops"], d, "max_hops")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-hops"], "SwitchControllerManagedSwitchStpSettings-MaxHops"); ok {
			if err = d.Set("max_hops", vv); err != nil {
				return fmt.Errorf("Error reading max_hops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchStpSettingsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerManagedSwitchStpSettings-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pending_timer", flattenSwitchControllerManagedSwitchStpSettingsPendingTimer2edl(o["pending-timer"], d, "pending_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["pending-timer"], "SwitchControllerManagedSwitchStpSettings-PendingTimer"); ok {
			if err = d.Set("pending_timer", vv); err != nil {
				return fmt.Errorf("Error reading pending_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pending_timer: %v", err)
		}
	}

	if err = d.Set("revision", flattenSwitchControllerManagedSwitchStpSettingsRevision2edl(o["revision"], d, "revision")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision"], "SwitchControllerManagedSwitchStpSettings-Revision"); ok {
			if err = d.Set("revision", vv); err != nil {
				return fmt.Errorf("Error reading revision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchStpSettingsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchStpSettings-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchStpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchStpSettingsForwardTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsHelloTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsLocalOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxAge2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxHops2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsPendingTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsRevision2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchStpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("forward_time"); ok || d.HasChange("forward_time") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsForwardTime2edl(d, v, "forward_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-time"] = t
		}
	}

	if v, ok := d.GetOk("hello_time"); ok || d.HasChange("hello_time") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsHelloTime2edl(d, v, "hello_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-time"] = t
		}
	}

	if v, ok := d.GetOk("local_override"); ok || d.HasChange("local_override") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsLocalOverride2edl(d, v, "local_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-override"] = t
		}
	}

	if v, ok := d.GetOk("max_age"); ok || d.HasChange("max_age") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsMaxAge2edl(d, v, "max_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-age"] = t
		}
	}

	if v, ok := d.GetOk("max_hops"); ok || d.HasChange("max_hops") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsMaxHops2edl(d, v, "max_hops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-hops"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pending_timer"); ok || d.HasChange("pending_timer") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsPendingTimer2edl(d, v, "pending_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pending-timer"] = t
		}
	}

	if v, ok := d.GetOk("revision"); ok || d.HasChange("revision") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsRevision2edl(d, v, "revision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchStpSettingsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
