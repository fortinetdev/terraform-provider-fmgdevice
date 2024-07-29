// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure alias command.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAliasCreate,
		Read:   resourceSystemAliasRead,
		Update: resourceSystemAliasUpdate,
		Delete: resourceSystemAliasDelete,

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
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemAliasCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAlias(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAlias resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAlias(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAlias resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAliasRead(d, m)
}

func resourceSystemAliasUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAlias(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlias resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlias(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlias resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAliasRead(d, m)
}

func resourceSystemAliasDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAlias(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAlias resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAliasRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAlias(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlias resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlias(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlias resource from API: %v", err)
	}
	return nil
}

func flattenSystemAliasCommand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAliasName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAlias(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("command", flattenSystemAliasCommand(o["command"], d, "command")); err != nil {
		if vv, ok := fortiAPIPatch(o["command"], "SystemAlias-Command"); ok {
			if err = d.Set("command", vv); err != nil {
				return fmt.Errorf("Error reading command: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAliasName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAlias-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemAliasFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAliasCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAlias(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command"); ok || d.HasChange("command") {
		t, err := expandSystemAliasCommand(d, v, "command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAliasName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
