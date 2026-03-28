// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSystemDhcpServerOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsCreate,
		Read:   resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsRead,
		Update: resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsUpdate,
		Delete: resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"system_dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	system_dhcp_server := d.Get("system_dhcp_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["system_dhcp_server"] = system_dhcp_server

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServerOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServerOptions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerManagedSwitchSystemDhcpServerOptions(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServerOptions(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerOptions resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerManagedSwitchSystemDhcpServerOptions(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServerOptions resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	system_dhcp_server := d.Get("system_dhcp_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["system_dhcp_server"] = system_dhcp_server

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServerOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerOptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServerOptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	system_dhcp_server := d.Get("system_dhcp_server").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["system_dhcp_server"] = system_dhcp_server

	wsParams["adom"] = adomv

	err = c.DeleteSwitchControllerManagedSwitchSystemDhcpServerOptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSystemDhcpServerOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
	system_dhcp_server := d.Get("system_dhcp_server").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	if system_dhcp_server == "" {
		system_dhcp_server = importOptionChecking(m.(*FortiClient).Cfg, "system_dhcp_server")
		if system_dhcp_server == "" {
			return fmt.Errorf("Parameter system_dhcp_server is missing")
		}
		if err = d.Set("system_dhcp_server", system_dhcp_server); err != nil {
			return fmt.Errorf("Error set params system_dhcp_server: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch
	paradict["system_dhcp_server"] = system_dhcp_server

	o, err := c.ReadSwitchControllerManagedSwitchSystemDhcpServerOptions(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServerOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSystemDhcpServerOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServerOptions resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsCode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchSystemDhcpServerOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("code", flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsCode3rdl(o["code"], d, "code")); err != nil {
		if vv, ok := fortiAPIPatch(o["code"], "SwitchControllerManagedSwitchSystemDhcpServerOptions-Code"); ok {
			if err = d.Set("code", vv); err != nil {
				return fmt.Errorf("Error reading code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading code: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerManagedSwitchSystemDhcpServerOptions-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerManagedSwitchSystemDhcpServerOptions-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SwitchControllerManagedSwitchSystemDhcpServerOptions-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("value", flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SwitchControllerManagedSwitchSystemDhcpServerOptions-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerOptionsValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchSystemDhcpServerOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("code"); ok || d.HasChange("code") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptionsCode3rdl(d, v, "code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["code"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptionsId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptionsIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptionsType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerOptionsValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
