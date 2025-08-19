// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Wireless Termination Points (WTP) system log server profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSyslogProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSyslogProfileCreate,
		Read:   resourceWirelessControllerSyslogProfileRead,
		Update: resourceWirelessControllerSyslogProfileUpdate,
		Delete: resourceWirelessControllerSyslogProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"server_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerSyslogProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerSyslogProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerSyslogProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerSyslogProfileRead(d, m)
}

func resourceWirelessControllerSyslogProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerSyslogProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSyslogProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSyslogProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerSyslogProfileRead(d, m)
}

func resourceWirelessControllerSyslogProfileDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerSyslogProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSyslogProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSyslogProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSyslogProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSyslogProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSyslogProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSyslogProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileLogLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSyslogProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerSyslogProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerSyslogProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("log_level", flattenWirelessControllerSyslogProfileLogLevel(o["log-level"], d, "log_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-level"], "WirelessControllerSyslogProfile-LogLevel"); ok {
			if err = d.Set("log_level", vv); err != nil {
				return fmt.Errorf("Error reading log_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_level: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerSyslogProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerSyslogProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_addr_type", flattenWirelessControllerSyslogProfileServerAddrType(o["server-addr-type"], d, "server_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-addr-type"], "WirelessControllerSyslogProfile-ServerAddrType"); ok {
			if err = d.Set("server_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading server_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_addr_type: %v", err)
		}
	}

	if err = d.Set("server_fqdn", flattenWirelessControllerSyslogProfileServerFqdn(o["server-fqdn"], d, "server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-fqdn"], "WirelessControllerSyslogProfile-ServerFqdn"); ok {
			if err = d.Set("server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_fqdn: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenWirelessControllerSyslogProfileServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "WirelessControllerSyslogProfile-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server_port", flattenWirelessControllerSyslogProfileServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "WirelessControllerSyslogProfile-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("server_status", flattenWirelessControllerSyslogProfileServerStatus(o["server-status"], d, "server_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-status"], "WirelessControllerSyslogProfile-ServerStatus"); ok {
			if err = d.Set("server_status", vv); err != nil {
				return fmt.Errorf("Error reading server_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_status: %v", err)
		}
	}

	if err = d.Set("server_type", flattenWirelessControllerSyslogProfileServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "WirelessControllerSyslogProfile-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSyslogProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSyslogProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileLogLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSyslogProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerSyslogProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("log_level"); ok || d.HasChange("log_level") {
		t, err := expandWirelessControllerSyslogProfileLogLevel(d, v, "log_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-level"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerSyslogProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_addr_type"); ok || d.HasChange("server_addr_type") {
		t, err := expandWirelessControllerSyslogProfileServerAddrType(d, v, "server_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("server_fqdn"); ok || d.HasChange("server_fqdn") {
		t, err := expandWirelessControllerSyslogProfileServerFqdn(d, v, "server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok || d.HasChange("server_ip") {
		t, err := expandWirelessControllerSyslogProfileServerIp(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandWirelessControllerSyslogProfileServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("server_status"); ok || d.HasChange("server_status") {
		t, err := expandWirelessControllerSyslogProfileServerStatus(d, v, "server_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-status"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandWirelessControllerSyslogProfileServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	return &obj, nil
}
