// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure LoRaWAN profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerLwProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerLwProfileCreate,
		Read:   resourceWirelessControllerLwProfileRead,
		Update: resourceWirelessControllerLwProfileUpdate,
		Delete: resourceWirelessControllerLwProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cups_api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"cups_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cups_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lw_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"tc_api_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"tc_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tc_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerLwProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLwProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerLwProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWirelessControllerLwProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWirelessControllerLwProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WirelessControllerLwProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWirelessControllerLwProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WirelessControllerLwProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerLwProfileRead(d, m)
}

func resourceWirelessControllerLwProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLwProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLwProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWirelessControllerLwProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLwProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerLwProfileRead(d, m)
}

func resourceWirelessControllerLwProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerLwProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerLwProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerLwProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerLwProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WirelessControllerLwProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerLwProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLwProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerLwProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileCupsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileCupsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileLwProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileTcServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileTcServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerLwProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerLwProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerLwProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cups_server", flattenWirelessControllerLwProfileCupsServer(o["cups-server"], d, "cups_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["cups-server"], "WirelessControllerLwProfile-CupsServer"); ok {
			if err = d.Set("cups_server", vv); err != nil {
				return fmt.Errorf("Error reading cups_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cups_server: %v", err)
		}
	}

	if err = d.Set("cups_server_port", flattenWirelessControllerLwProfileCupsServerPort(o["cups-server-port"], d, "cups_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["cups-server-port"], "WirelessControllerLwProfile-CupsServerPort"); ok {
			if err = d.Set("cups_server_port", vv); err != nil {
				return fmt.Errorf("Error reading cups_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cups_server_port: %v", err)
		}
	}

	if err = d.Set("lw_protocol", flattenWirelessControllerLwProfileLwProtocol(o["lw-protocol"], d, "lw_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["lw-protocol"], "WirelessControllerLwProfile-LwProtocol"); ok {
			if err = d.Set("lw_protocol", vv); err != nil {
				return fmt.Errorf("Error reading lw_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lw_protocol: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerLwProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerLwProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tc_server", flattenWirelessControllerLwProfileTcServer(o["tc-server"], d, "tc_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tc-server"], "WirelessControllerLwProfile-TcServer"); ok {
			if err = d.Set("tc_server", vv); err != nil {
				return fmt.Errorf("Error reading tc_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tc_server: %v", err)
		}
	}

	if err = d.Set("tc_server_port", flattenWirelessControllerLwProfileTcServerPort(o["tc-server-port"], d, "tc_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["tc-server-port"], "WirelessControllerLwProfile-TcServerPort"); ok {
			if err = d.Set("tc_server_port", vv); err != nil {
				return fmt.Errorf("Error reading tc_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tc_server_port: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerLwProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerLwProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileCupsApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerLwProfileCupsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileCupsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileLwProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileTcApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerLwProfileTcServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileTcServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerLwProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerLwProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cups_api_key"); ok || d.HasChange("cups_api_key") {
		t, err := expandWirelessControllerLwProfileCupsApiKey(d, v, "cups_api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-api-key"] = t
		}
	}

	if v, ok := d.GetOk("cups_server"); ok || d.HasChange("cups_server") {
		t, err := expandWirelessControllerLwProfileCupsServer(d, v, "cups_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-server"] = t
		}
	}

	if v, ok := d.GetOk("cups_server_port"); ok || d.HasChange("cups_server_port") {
		t, err := expandWirelessControllerLwProfileCupsServerPort(d, v, "cups_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-server-port"] = t
		}
	}

	if v, ok := d.GetOk("lw_protocol"); ok || d.HasChange("lw_protocol") {
		t, err := expandWirelessControllerLwProfileLwProtocol(d, v, "lw_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lw-protocol"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerLwProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tc_api_key"); ok || d.HasChange("tc_api_key") {
		t, err := expandWirelessControllerLwProfileTcApiKey(d, v, "tc_api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-api-key"] = t
		}
	}

	if v, ok := d.GetOk("tc_server"); ok || d.HasChange("tc_server") {
		t, err := expandWirelessControllerLwProfileTcServer(d, v, "tc_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-server"] = t
		}
	}

	if v, ok := d.GetOk("tc_server_port"); ok || d.HasChange("tc_server_port") {
		t, err := expandWirelessControllerLwProfileTcServerPort(d, v, "tc_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-server-port"] = t
		}
	}

	return &obj, nil
}
