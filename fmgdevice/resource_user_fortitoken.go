// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure FortiToken.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserFortitoken() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFortitokenCreate,
		Read:   resourceUserFortitokenRead,
		Update: resourceUserFortitokenUpdate,
		Delete: resourceUserFortitokenDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reg_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserFortitokenCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error creating UserFortitoken resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("serial_number")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserFortitoken(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserFortitoken(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserFortitoken resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserFortitoken(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserFortitoken resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "serial_number"))

	return resourceUserFortitokenRead(d, m)
}

func resourceUserFortitokenUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error updating UserFortitoken resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserFortitoken(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserFortitoken resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial_number"))

	return resourceUserFortitokenRead(d, m)
}

func resourceUserFortitokenDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	wsParams["adom"] = adomv

	err = c.DeleteUserFortitoken(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserFortitoken resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFortitokenRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserFortitoken(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserFortitoken resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFortitoken(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserFortitoken resource from API: %v", err)
	}
	return nil
}

func flattenUserFortitokenComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenRegId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserFortitoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comments", flattenUserFortitokenComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "UserFortitoken-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("license", flattenUserFortitokenLicense(o["license"], d, "license")); err != nil {
		if vv, ok := fortiAPIPatch(o["license"], "UserFortitoken-License"); ok {
			if err = d.Set("license", vv); err != nil {
				return fmt.Errorf("Error reading license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading license: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenUserFortitokenSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "UserFortitoken-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("status", flattenUserFortitokenStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserFortitoken-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("os_ver", flattenUserFortitokenOsVer(o["os-ver"], d, "os_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-ver"], "UserFortitoken-OsVer"); ok {
			if err = d.Set("os_ver", vv); err != nil {
				return fmt.Errorf("Error reading os_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_ver: %v", err)
		}
	}

	if err = d.Set("reg_id", flattenUserFortitokenRegId(o["reg-id"], d, "reg_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["reg-id"], "UserFortitoken-RegId"); ok {
			if err = d.Set("reg_id", vv); err != nil {
				return fmt.Errorf("Error reading reg_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reg_id: %v", err)
		}
	}

	return nil
}

func flattenUserFortitokenFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserFortitokenComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenRegId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserFortitoken(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandUserFortitokenComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("license"); ok || d.HasChange("license") {
		t, err := expandUserFortitokenLicense(d, v, "license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandUserFortitokenSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserFortitokenStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok || d.HasChange("os_ver") {
		t, err := expandUserFortitokenOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-ver"] = t
		}
	}

	if v, ok := d.GetOk("reg_id"); ok || d.HasChange("reg_id") {
		t, err := expandUserFortitokenRegId(d, v, "reg_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-id"] = t
		}
	}

	return &obj, nil
}
