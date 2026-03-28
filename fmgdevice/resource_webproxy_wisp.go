// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Websense Integrated Services Protocol (WISP) servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyWisp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyWispCreate,
		Read:   resourceWebProxyWispRead,
		Update: resourceWebProxyWispUpdate,
		Delete: resourceWebProxyWispDelete,

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
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"outgoing_ip": &schema.Schema{
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
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyWispCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyWisp(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyWisp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyWisp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyWisp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyWisp resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyWisp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyWisp resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyWispRead(d, m)
}

func resourceWebProxyWispUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyWisp(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyWisp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyWisp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyWisp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyWispRead(d, m)
}

func resourceWebProxyWispDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyWisp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyWisp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyWispRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyWisp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyWisp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyWisp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyWisp resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyWispComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyWispTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyWisp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWebProxyWispComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebProxyWisp-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenWebProxyWispMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "WebProxyWisp-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyWispName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyWisp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenWebProxyWispOutgoingIp(o["outgoing-ip"], d, "outgoing_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["outgoing-ip"], "WebProxyWisp-OutgoingIp"); ok {
			if err = d.Set("outgoing_ip", vv); err != nil {
				return fmt.Errorf("Error reading outgoing_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenWebProxyWispServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "WebProxyWisp-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server_port", flattenWebProxyWispServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "WebProxyWisp-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("timeout", flattenWebProxyWispTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "WebProxyWisp-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenWebProxyWispFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyWispComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyWispTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyWisp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebProxyWispComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandWebProxyWispMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyWispName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok || d.HasChange("outgoing_ip") {
		t, err := expandWebProxyWispOutgoingIp(d, v, "outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok || d.HasChange("server_ip") {
		t, err := expandWebProxyWispServerIp(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok || d.HasChange("server_port") {
		t, err := expandWebProxyWispServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandWebProxyWispTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
