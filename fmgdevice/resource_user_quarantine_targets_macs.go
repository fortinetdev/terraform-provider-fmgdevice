// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Quarantine MACs.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserQuarantineTargetsMacs() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserQuarantineTargetsMacsCreate,
		Read:   resourceUserQuarantineTargetsMacsRead,
		Update: resourceUserQuarantineTargetsMacsUpdate,
		Delete: resourceUserQuarantineTargetsMacsDelete,

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
			"targets": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entry_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserQuarantineTargetsMacsCreate(d *schema.ResourceData, m interface{}) error {
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
	targets := d.Get("targets").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["targets"] = targets

	obj, err := getObjectUserQuarantineTargetsMacs(d)
	if err != nil {
		return fmt.Errorf("Error creating UserQuarantineTargetsMacs resource while getting object: %v", err)
	}

	_, err = c.CreateUserQuarantineTargetsMacs(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating UserQuarantineTargetsMacs resource: %v", err)
	}

	d.SetId(getStringKey(d, "mac"))

	return resourceUserQuarantineTargetsMacsRead(d, m)
}

func resourceUserQuarantineTargetsMacsUpdate(d *schema.ResourceData, m interface{}) error {
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
	targets := d.Get("targets").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["targets"] = targets

	obj, err := getObjectUserQuarantineTargetsMacs(d)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantineTargetsMacs resource while getting object: %v", err)
	}

	_, err = c.UpdateUserQuarantineTargetsMacs(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantineTargetsMacs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "mac"))

	return resourceUserQuarantineTargetsMacsRead(d, m)
}

func resourceUserQuarantineTargetsMacsDelete(d *schema.ResourceData, m interface{}) error {
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
	targets := d.Get("targets").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["targets"] = targets

	err = c.DeleteUserQuarantineTargetsMacs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting UserQuarantineTargetsMacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineTargetsMacsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	targets := d.Get("targets").(string)
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
	if targets == "" {
		targets = importOptionChecking(m.(*FortiClient).Cfg, "targets")
		if targets == "" {
			return fmt.Errorf("Parameter targets is missing")
		}
		if err = d.Set("targets", targets); err != nil {
			return fmt.Errorf("Error set params targets: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["targets"] = targets

	o, err := c.ReadUserQuarantineTargetsMacs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantineTargetsMacs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserQuarantineTargetsMacs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantineTargetsMacs resource from API: %v", err)
	}
	return nil
}

func flattenUserQuarantineTargetsMacsDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDrop3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsEntryId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return case_insensitive(v, d.Get(pre))
}

func flattenUserQuarantineTargetsMacsParent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserQuarantineTargetsMacs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenUserQuarantineTargetsMacsDescription3rdl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "UserQuarantineTargetsMacs-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("drop", flattenUserQuarantineTargetsMacsDrop3rdl(o["drop"], d, "drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop"], "UserQuarantineTargetsMacs-Drop"); ok {
			if err = d.Set("drop", vv); err != nil {
				return fmt.Errorf("Error reading drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop: %v", err)
		}
	}

	if err = d.Set("entry_id", flattenUserQuarantineTargetsMacsEntryId3rdl(o["entry-id"], d, "entry_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entry-id"], "UserQuarantineTargetsMacs-EntryId"); ok {
			if err = d.Set("entry_id", vv); err != nil {
				return fmt.Errorf("Error reading entry_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entry_id: %v", err)
		}
	}

	if err = d.Set("mac", flattenUserQuarantineTargetsMacsMac3rdl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "UserQuarantineTargetsMacs-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("parent", flattenUserQuarantineTargetsMacsParent3rdl(o["parent"], d, "parent")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent"], "UserQuarantineTargetsMacs-Parent"); ok {
			if err = d.Set("parent", vv); err != nil {
				return fmt.Errorf("Error reading parent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent: %v", err)
		}
	}

	return nil
}

func flattenUserQuarantineTargetsMacsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserQuarantineTargetsMacsDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDrop3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsEntryId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsParent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserQuarantineTargetsMacs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandUserQuarantineTargetsMacsDescription3rdl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("drop"); ok || d.HasChange("drop") {
		t, err := expandUserQuarantineTargetsMacsDrop3rdl(d, v, "drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop"] = t
		}
	}

	if v, ok := d.GetOk("entry_id"); ok || d.HasChange("entry_id") {
		t, err := expandUserQuarantineTargetsMacsEntryId3rdl(d, v, "entry_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry-id"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandUserQuarantineTargetsMacsMac3rdl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("parent"); ok || d.HasChange("parent") {
		t, err := expandUserQuarantineTargetsMacsParent3rdl(d, v, "parent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent"] = t
		}
	}

	return &obj, nil
}
