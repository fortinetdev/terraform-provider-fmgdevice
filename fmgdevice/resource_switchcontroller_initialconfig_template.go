// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure template for auto-generated VLANs.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerInitialConfigTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerInitialConfigTemplateCreate,
		Read:   resourceSwitchControllerInitialConfigTemplateRead,
		Update: resourceSwitchControllerInitialConfigTemplateUpdate,
		Delete: resourceSwitchControllerInitialConfigTemplateDelete,

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
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerInitialConfigTemplateCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerInitialConfigTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerInitialConfigTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerInitialConfigTemplate(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerInitialConfigTemplateRead(d, m)
}

func resourceSwitchControllerInitialConfigTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerInitialConfigTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerInitialConfigTemplate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerInitialConfigTemplateRead(d, m)
}

func resourceSwitchControllerInitialConfigTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerInitialConfigTemplate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerInitialConfigTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerInitialConfigTemplate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerInitialConfigTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigTemplate resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerInitialConfigTemplateAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigTemplateAutoIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateDhcpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerInitialConfigTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerInitialConfigTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowaccess", flattenSwitchControllerInitialConfigTemplateAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "SwitchControllerInitialConfigTemplate-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("auto_ip", flattenSwitchControllerInitialConfigTemplateAutoIp(o["auto-ip"], d, "auto_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-ip"], "SwitchControllerInitialConfigTemplate-AutoIp"); ok {
			if err = d.Set("auto_ip", vv); err != nil {
				return fmt.Errorf("Error reading auto_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenSwitchControllerInitialConfigTemplateDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server"], "SwitchControllerInitialConfigTemplate-DhcpServer"); ok {
			if err = d.Set("dhcp_server", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerInitialConfigTemplateIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerInitialConfigTemplate-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerInitialConfigTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerInitialConfigTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenSwitchControllerInitialConfigTemplateVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlanid"], "SwitchControllerInitialConfigTemplate-Vlanid"); ok {
			if err = d.Set("vlanid", vv); err != nil {
				return fmt.Errorf("Error reading vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerInitialConfigTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerInitialConfigTemplateAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerInitialConfigTemplateAutoIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerInitialConfigTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerInitialConfigTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandSwitchControllerInitialConfigTemplateAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("auto_ip"); ok || d.HasChange("auto_ip") {
		t, err := expandSwitchControllerInitialConfigTemplateAutoIp(d, v, "auto_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok || d.HasChange("dhcp_server") {
		t, err := expandSwitchControllerInitialConfigTemplateDhcpServer(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerInitialConfigTemplateIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerInitialConfigTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok || d.HasChange("vlanid") {
		t, err := expandSwitchControllerInitialConfigTemplateVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	return &obj, nil
}
