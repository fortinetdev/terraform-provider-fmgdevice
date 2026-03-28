// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SMS server for sending SMS messages to support user authentication.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSmsServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSmsServerCreate,
		Read:   resourceSystemSmsServerRead,
		Update: resourceSystemSmsServerUpdate,
		Delete: resourceSystemSmsServerDelete,

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
			"mail_server": &schema.Schema{
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

func resourceSystemSmsServerCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSmsServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSmsServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSmsServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSmsServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSmsServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSmsServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSmsServerRead(d, m)
}

func resourceSystemSmsServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSmsServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSmsServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSmsServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSmsServerRead(d, m)
}

func resourceSystemSmsServerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemSmsServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSmsServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSmsServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSmsServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSmsServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSmsServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSmsServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemSmsServerMailServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSmsServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSmsServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mail_server", flattenSystemSmsServerMailServer(o["mail-server"], d, "mail_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mail-server"], "SystemSmsServer-MailServer"); ok {
			if err = d.Set("mail_server", vv); err != nil {
				return fmt.Errorf("Error reading mail_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mail_server: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSmsServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSmsServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemSmsServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSmsServerMailServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSmsServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSmsServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mail_server"); ok || d.HasChange("mail_server") {
		t, err := expandSystemSmsServerMailServer(d, v, "mail_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSmsServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
