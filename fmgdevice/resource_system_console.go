// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure console.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleUpdate,
		Read:   resourceSystemConsoleRead,
		Update: resourceSystemConsoleUpdate,
		Delete: resourceSystemConsoleDelete,

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
			"baudrate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortiexplorer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemConsoleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemConsole(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemConsole(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemConsole")

	return resourceSystemConsoleRead(d, m)
}

func resourceSystemConsoleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemConsole(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemConsole(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConsole(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource from API: %v", err)
	}
	return nil
}

func flattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleFortiexplorer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("baudrate", flattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate")); err != nil {
		if vv, ok := fortiAPIPatch(o["baudrate"], "SystemConsole-Baudrate"); ok {
			if err = d.Set("baudrate", vv); err != nil {
				return fmt.Errorf("Error reading baudrate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading baudrate: %v", err)
		}
	}

	if err = d.Set("fortiexplorer", flattenSystemConsoleFortiexplorer(o["fortiexplorer"], d, "fortiexplorer")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiexplorer"], "SystemConsole-Fortiexplorer"); ok {
			if err = d.Set("fortiexplorer", vv); err != nil {
				return fmt.Errorf("Error reading fortiexplorer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiexplorer: %v", err)
		}
	}

	if err = d.Set("login", flattenSystemConsoleLogin(o["login"], d, "login")); err != nil {
		if vv, ok := fortiAPIPatch(o["login"], "SystemConsole-Login"); ok {
			if err = d.Set("login", vv); err != nil {
				return fmt.Errorf("Error reading login: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemConsoleMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemConsole-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("output", flattenSystemConsoleOutput(o["output"], d, "output")); err != nil {
		if vv, ok := fortiAPIPatch(o["output"], "SystemConsole-Output"); ok {
			if err = d.Set("output", vv); err != nil {
				return fmt.Errorf("Error reading output: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading output: %v", err)
		}
	}

	return nil
}

func flattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemConsoleBaudrate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleFortiexplorer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleOutput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConsole(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("baudrate"); ok || d.HasChange("baudrate") {
		t, err := expandSystemConsoleBaudrate(d, v, "baudrate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["baudrate"] = t
		}
	}

	if v, ok := d.GetOk("fortiexplorer"); ok || d.HasChange("fortiexplorer") {
		t, err := expandSystemConsoleFortiexplorer(d, v, "fortiexplorer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiexplorer"] = t
		}
	}

	if v, ok := d.GetOk("login"); ok || d.HasChange("login") {
		t, err := expandSystemConsoleLogin(d, v, "login")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemConsoleMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("output"); ok || d.HasChange("output") {
		t, err := expandSystemConsoleOutput(d, v, "output")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output"] = t
		}
	}

	return &obj, nil
}
