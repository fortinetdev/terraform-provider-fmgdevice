// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Fortinet Single Sign On (FSSO) server.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFssoPollingUpdate,
		Read:   resourceSystemFssoPollingRead,
		Update: resourceSystemFssoPollingUpdate,
		Delete: resourceSystemFssoPollingDelete,

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
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listening_port": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFssoPolling(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFssoPolling")

	return resourceSystemFssoPollingRead(d, m)
}

func resourceSystemFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFssoPolling(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFssoPollingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFssoPolling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFssoPolling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenSystemFssoPollingGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFssoPolling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_gui_meta", flattenSystemFssoPollingGuiMeta(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "SystemFssoPolling-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemFssoPollingAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystemFssoPolling-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenSystemFssoPollingListeningPort(o["listening-port"], d, "listening_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["listening-port"], "SystemFssoPolling-ListeningPort"); ok {
			if err = d.Set("listening_port", vv); err != nil {
				return fmt.Errorf("Error reading listening_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFssoPollingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFssoPolling-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFssoPollingGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingAuthPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFssoPollingAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFssoPolling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandSystemFssoPollingGuiMeta(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("auth_password"); ok || d.HasChange("auth_password") {
		t, err := expandSystemFssoPollingAuthPassword(d, v, "auth_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystemFssoPollingAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("listening_port"); ok || d.HasChange("listening_port") {
		t, err := expandSystemFssoPollingListeningPort(d, v, "listening_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listening-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFssoPollingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
