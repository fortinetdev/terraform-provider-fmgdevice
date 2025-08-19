// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchCustomCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchCustomCommandCreate,
		Read:   resourceSwitchControllerManagedSwitchCustomCommandRead,
		Update: resourceSwitchControllerManagedSwitchCustomCommandUpdate,
		Delete: resourceSwitchControllerManagedSwitchCustomCommandDelete,

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
			"command_entry": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"command_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchCustomCommandCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchCustomCommand resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerManagedSwitchCustomCommand(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	d.SetId(getStringKey(d, "command_entry"))

	return resourceSwitchControllerManagedSwitchCustomCommandRead(d, m)
}

func resourceSwitchControllerManagedSwitchCustomCommandUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerManagedSwitchCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchCustomCommand resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerManagedSwitchCustomCommand(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "command_entry"))

	return resourceSwitchControllerManagedSwitchCustomCommandRead(d, m)
}

func resourceSwitchControllerManagedSwitchCustomCommandDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerManagedSwitchCustomCommand(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchCustomCommandRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchCustomCommand(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchCustomCommand(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchCustomCommand resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("command_entry", flattenSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(o["command-entry"], d, "command_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-entry"], "SwitchControllerManagedSwitchCustomCommand-CommandEntry"); ok {
			if err = d.Set("command_entry", vv); err != nil {
				return fmt.Errorf("Error reading command_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_entry: %v", err)
		}
	}

	if err = d.Set("command_name", flattenSwitchControllerManagedSwitchCustomCommandCommandName2edl(o["command-name"], d, "command_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-name"], "SwitchControllerManagedSwitchCustomCommand-CommandName"); ok {
			if err = d.Set("command_name", vv); err != nil {
				return fmt.Errorf("Error reading command_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchCustomCommandFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchCustomCommandCommandName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command_entry"); ok || d.HasChange("command_entry") {
		t, err := expandSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(d, v, "command_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-entry"] = t
		}
	}

	if v, ok := d.GetOk("command_name"); ok || d.HasChange("command_name") {
		t, err := expandSwitchControllerManagedSwitchCustomCommandCommandName2edl(d, v, "command_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-name"] = t
		}
	}

	return &obj, nil
}
