// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DHCP IP range configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchSystemDhcpServerIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeCreate,
		Read:   resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeRead,
		Update: resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeUpdate,
		Delete: resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeDelete,

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
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServerIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServerIpRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerManagedSwitchSystemDhcpServerIpRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServerIpRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerIpRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerManagedSwitchSystemDhcpServerIpRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerManagedSwitchSystemDhcpServerIpRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerManagedSwitchSystemDhcpServerIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerIpRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerManagedSwitchSystemDhcpServerIpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchSystemDhcpServerIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeRead(d, m)
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerManagedSwitchSystemDhcpServerIpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchSystemDhcpServerIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchSystemDhcpServerIpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerManagedSwitchSystemDhcpServerIpRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServerIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchSystemDhcpServerIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchSystemDhcpServerIpRange resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitchSystemDhcpServerIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_ip", flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp3rdl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "SwitchControllerManagedSwitchSystemDhcpServerIpRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerManagedSwitchSystemDhcpServerIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp3rdl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "SwitchControllerManagedSwitchSystemDhcpServerIpRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchSystemDhcpServerIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitchSystemDhcpServerIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeEndIp3rdl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandSwitchControllerManagedSwitchSystemDhcpServerIpRangeStartIp3rdl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
