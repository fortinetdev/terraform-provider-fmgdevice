// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure certificate users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCertificateCreate,
		Read:   resourceUserCertificateRead,
		Update: resourceUserCertificateUpdate,
		Delete: resourceUserCertificateDelete,

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
			"common_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"issuer": &schema.Schema{
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserCertificateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserCertificate(d)
	if err != nil {
		return fmt.Errorf("Error creating UserCertificate resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserCertificate(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserCertificate(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserCertificate resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserCertificate(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserCertificate resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserCertificateRead(d, m)
}

func resourceUserCertificateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserCertificate(d)
	if err != nil {
		return fmt.Errorf("Error updating UserCertificate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserCertificate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserCertificate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserCertificateRead(d, m)
}

func resourceUserCertificateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserCertificate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserCertificateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserCertificate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserCertificate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserCertificate resource from API: %v", err)
	}
	return nil
}

func flattenUserCertificateCommonName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserCertificateId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserCertificateIssuer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserCertificateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserCertificateStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserCertificateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserCertificate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("common_name", flattenUserCertificateCommonName(o["common-name"], d, "common_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["common-name"], "UserCertificate-CommonName"); ok {
			if err = d.Set("common_name", vv); err != nil {
				return fmt.Errorf("Error reading common_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading common_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserCertificateId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserCertificate-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("issuer", flattenUserCertificateIssuer(o["issuer"], d, "issuer")); err != nil {
		if vv, ok := fortiAPIPatch(o["issuer"], "UserCertificate-Issuer"); ok {
			if err = d.Set("issuer", vv); err != nil {
				return fmt.Errorf("Error reading issuer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("name", flattenUserCertificateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserCertificate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenUserCertificateStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserCertificate-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenUserCertificateType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserCertificate-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenUserCertificateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserCertificateCommonName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserCertificateId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserCertificateIssuer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserCertificateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserCertificateStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserCertificateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserCertificate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("common_name"); ok || d.HasChange("common_name") {
		t, err := expandUserCertificateCommonName(d, v, "common_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["common-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserCertificateId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("issuer"); ok || d.HasChange("issuer") {
		t, err := expandUserCertificateIssuer(d, v, "issuer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["issuer"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserCertificateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserCertificateStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserCertificateType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
