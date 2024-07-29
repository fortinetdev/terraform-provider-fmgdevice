// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP Access Control MIB View configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpMibView() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpMibViewCreate,
		Read:   resourceSystemSnmpMibViewRead,
		Update: resourceSystemSnmpMibViewUpdate,
		Delete: resourceSystemSnmpMibViewDelete,

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
			"exclude": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"include": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemSnmpMibViewCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemSnmpMibView(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpMibView resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpMibView(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpMibView resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSnmpMibViewRead(d, m)
}

func resourceSystemSnmpMibViewUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpMibView(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpMibView resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpMibView(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpMibView resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSnmpMibViewRead(d, m)
}

func resourceSystemSnmpMibViewDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSnmpMibView(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpMibView resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpMibViewRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpMibView(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpMibView resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpMibView(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpMibView resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpMibViewExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpMibViewInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpMibViewName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpMibView(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("exclude", flattenSystemSnmpMibViewExclude(o["exclude"], d, "exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude"], "SystemSnmpMibView-Exclude"); ok {
			if err = d.Set("exclude", vv); err != nil {
				return fmt.Errorf("Error reading exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude: %v", err)
		}
	}

	if err = d.Set("include", flattenSystemSnmpMibViewInclude(o["include"], d, "include")); err != nil {
		if vv, ok := fortiAPIPatch(o["include"], "SystemSnmpMibView-Include"); ok {
			if err = d.Set("include", vv); err != nil {
				return fmt.Errorf("Error reading include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSnmpMibViewName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSnmpMibView-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpMibViewFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpMibViewExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpMibViewInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpMibViewName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpMibView(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("exclude"); ok || d.HasChange("exclude") {
		t, err := expandSystemSnmpMibViewExclude(d, v, "exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude"] = t
		}
	}

	if v, ok := d.GetOk("include"); ok || d.HasChange("include") {
		t, err := expandSystemSnmpMibViewInclude(d, v, "include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSnmpMibViewName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
