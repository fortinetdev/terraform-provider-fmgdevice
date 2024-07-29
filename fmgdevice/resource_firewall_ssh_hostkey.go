// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSH proxy host public keys.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSshHostKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSshHostKeyCreate,
		Read:   resourceFirewallSshHostKeyRead,
		Update: resourceFirewallSshHostKeyUpdate,
		Delete: resourceFirewallSshHostKeyDelete,

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
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
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
			"usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSshHostKeyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectFirewallSshHostKey(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSshHostKey resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallSshHostKey(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating FirewallSshHostKey resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSshHostKeyRead(d, m)
}

func resourceFirewallSshHostKeyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectFirewallSshHostKey(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshHostKey resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallSshHostKey(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshHostKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallSshHostKeyRead(d, m)
}

func resourceFirewallSshHostKeyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteFirewallSshHostKey(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSshHostKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshHostKeyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSshHostKey(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshHostKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSshHostKey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshHostKey resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSshHostKeyHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyNid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyPublicKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSshHostKey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hostname", flattenFirewallSshHostKeyHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "FirewallSshHostKey-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallSshHostKeyIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallSshHostKey-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallSshHostKeyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallSshHostKey-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nid", flattenFirewallSshHostKeyNid(o["nid"], d, "nid")); err != nil {
		if vv, ok := fortiAPIPatch(o["nid"], "FirewallSshHostKey-Nid"); ok {
			if err = d.Set("nid", vv); err != nil {
				return fmt.Errorf("Error reading nid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nid: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallSshHostKeyPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallSshHostKey-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("public_key", flattenFirewallSshHostKeyPublicKey(o["public-key"], d, "public_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-key"], "FirewallSshHostKey-PublicKey"); ok {
			if err = d.Set("public_key", vv); err != nil {
				return fmt.Errorf("Error reading public_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_key: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSshHostKeyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSshHostKey-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallSshHostKeyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallSshHostKey-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("usage", flattenFirewallSshHostKeyUsage(o["usage"], d, "usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["usage"], "FirewallSshHostKey-Usage"); ok {
			if err = d.Set("usage", vv); err != nil {
				return fmt.Errorf("Error reading usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usage: %v", err)
		}
	}

	return nil
}

func flattenFirewallSshHostKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSshHostKeyHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyNid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyPublicKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSshHostKey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandFirewallSshHostKeyHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallSshHostKeyIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallSshHostKeyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nid"); ok || d.HasChange("nid") {
		t, err := expandFirewallSshHostKeyNid(d, v, "nid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nid"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallSshHostKeyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("public_key"); ok || d.HasChange("public_key") {
		t, err := expandFirewallSshHostKeyPublicKey(d, v, "public_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-key"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSshHostKeyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallSshHostKeyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("usage"); ok || d.HasChange("usage") {
		t, err := expandFirewallSshHostKeyUsage(d, v, "usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usage"] = t
		}
	}

	return &obj, nil
}
