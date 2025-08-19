// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IP and MAC address configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryCreate,
		Read:   resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryRead,
		Update: resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryUpdate,
		Delete: resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryDelete,

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
			"ip_source_guard": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entry_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ip_source_guard := d.Get("ip_source_guard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ip_source_guard"] = ip_source_guard

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchIpSourceGuardBindingEntry(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource: %v", err)
	}

	d.SetId(getStringKey(d, "entry_name"))

	return resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryRead(d, m)
}

func resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ip_source_guard := d.Get("ip_source_guard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ip_source_guard"] = ip_source_guard

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchIpSourceGuardBindingEntry(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "entry_name"))

	return resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryRead(d, m)
}

func resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	ip_source_guard := d.Get("ip_source_guard").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ip_source_guard"] = ip_source_guard

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchIpSourceGuardBindingEntry(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchIpSourceGuardBindingEntryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
	ip_source_guard := d.Get("ip_source_guard").(string)
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
	if ip_source_guard == "" {
		ip_source_guard = importOptionChecking(m.(*FortiClient).Cfg, "ip_source_guard")
		if ip_source_guard == "" {
			return fmt.Errorf("Parameter ip_source_guard is missing")
		}
		if err = d.Set("ip_source_guard", ip_source_guard); err != nil {
			return fmt.Errorf("Error set params ip_source_guard: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["ip_source_guard"] = ip_source_guard

	o, err := c.ReadSwitchControllerManagedSwitchIpSourceGuardBindingEntry(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchIpSourceGuardBindingEntry resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("entry_name", flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName3rdl(o["entry-name"], d, "entry_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["entry-name"], "SwitchControllerManagedSwitchIpSourceGuardBindingEntry-EntryName"); ok {
			if err = d.Set("entry_name", vv); err != nil {
				return fmt.Errorf("Error reading entry_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entry_name: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerManagedSwitchIpSourceGuardBindingEntry-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac3rdl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "SwitchControllerManagedSwitchIpSourceGuardBindingEntry-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("entry_name"); ok || d.HasChange("entry_name") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName3rdl(d, v, "entry_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry-name"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac3rdl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	return &obj, nil
}
