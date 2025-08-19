// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomRadiusServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomRadiusServerCreate,
		Read:   resourceSystemVdomRadiusServerRead,
		Update: resourceSystemVdomRadiusServerUpdate,
		Delete: resourceSystemVdomRadiusServerDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"radius_server_vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomRadiusServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemVdomRadiusServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomRadiusServer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomRadiusServer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomRadiusServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomRadiusServerRead(d, m)
}

func resourceSystemVdomRadiusServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemVdomRadiusServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomRadiusServer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomRadiusServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomRadiusServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVdomRadiusServerRead(d, m)
}

func resourceSystemVdomRadiusServerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemVdomRadiusServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomRadiusServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRadiusServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomRadiusServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomRadiusServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomRadiusServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomRadiusServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomRadiusServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystemVdomRadiusServerRadiusServerVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomRadiusServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomRadiusServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemVdomRadiusServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemVdomRadiusServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("radius_server_vdom", flattenSystemVdomRadiusServerRadiusServerVdom(o["radius-server-vdom"], d, "radius_server_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server-vdom"], "SystemVdomRadiusServer-RadiusServerVdom"); ok {
			if err = d.Set("radius_server_vdom", vv); err != nil {
				return fmt.Errorf("Error reading radius_server_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server_vdom: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemVdomRadiusServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemVdomRadiusServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomRadiusServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomRadiusServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystemVdomRadiusServerRadiusServerVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomRadiusServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomRadiusServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemVdomRadiusServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("radius_server_vdom"); ok || d.HasChange("radius_server_vdom") {
		t, err := expandSystemVdomRadiusServerRadiusServerVdom(d, v, "radius_server_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server-vdom"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemVdomRadiusServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
