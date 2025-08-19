// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerCustomCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerCustomCommandCreate,
		Read:   resourceSwitchControllerCustomCommandRead,
		Update: resourceSwitchControllerCustomCommandUpdate,
		Delete: resourceSwitchControllerCustomCommandDelete,

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
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerCustomCommandCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerCustomCommand resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerCustomCommand(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerCustomCommand resource: %v", err)
	}

	d.SetId(getStringKey(d, "command_name"))

	return resourceSwitchControllerCustomCommandRead(d, m)
}

func resourceSwitchControllerCustomCommandUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerCustomCommand resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerCustomCommand(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerCustomCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "command_name"))

	return resourceSwitchControllerCustomCommandRead(d, m)
}

func resourceSwitchControllerCustomCommandDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerCustomCommand(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerCustomCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerCustomCommandRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerCustomCommand(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerCustomCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerCustomCommand(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerCustomCommand resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerCustomCommandCommand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerCustomCommandDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerCustomCommand(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("command", flattenSwitchControllerCustomCommandCommand(o["command"], d, "command")); err != nil {
		if vv, ok := fortiAPIPatch(o["command"], "SwitchControllerCustomCommand-Command"); ok {
			if err = d.Set("command", vv); err != nil {
				return fmt.Errorf("Error reading command: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	if err = d.Set("command_name", flattenSwitchControllerCustomCommandCommandName(o["command-name"], d, "command_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-name"], "SwitchControllerCustomCommand-CommandName"); ok {
			if err = d.Set("command_name", vv); err != nil {
				return fmt.Errorf("Error reading command_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerCustomCommandDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerCustomCommand-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerCustomCommandFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerCustomCommandCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerCustomCommandDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerCustomCommand(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command"); ok || d.HasChange("command") {
		t, err := expandSwitchControllerCustomCommandCommand(d, v, "command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	if v, ok := d.GetOk("command_name"); ok || d.HasChange("command_name") {
		t, err := expandSwitchControllerCustomCommandCommandName(d, v, "command_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerCustomCommandDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
