// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Extra servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserDomainControllerExtraServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDomainControllerExtraServerCreate,
		Read:   resourceUserDomainControllerExtraServerRead,
		Update: resourceUserDomainControllerExtraServerUpdate,
		Delete: resourceUserDomainControllerExtraServerDelete,

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
			"domain_controller": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceUserDomainControllerExtraServerCreate(d *schema.ResourceData, m interface{}) error {
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
	domain_controller := d.Get("domain_controller").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_controller"] = domain_controller

	obj, err := getObjectUserDomainControllerExtraServer(d)
	if err != nil {
		return fmt.Errorf("Error creating UserDomainControllerExtraServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserDomainControllerExtraServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserDomainControllerExtraServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserDomainControllerExtraServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserDomainControllerExtraServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserDomainControllerExtraServer resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserDomainControllerExtraServerRead(d, m)
}

func resourceUserDomainControllerExtraServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	domain_controller := d.Get("domain_controller").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_controller"] = domain_controller

	obj, err := getObjectUserDomainControllerExtraServer(d)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainControllerExtraServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserDomainControllerExtraServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainControllerExtraServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceUserDomainControllerExtraServerRead(d, m)
}

func resourceUserDomainControllerExtraServerDelete(d *schema.ResourceData, m interface{}) error {
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
	domain_controller := d.Get("domain_controller").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_controller"] = domain_controller

	wsParams["adom"] = adomv

	err = c.DeleteUserDomainControllerExtraServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserDomainControllerExtraServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDomainControllerExtraServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	domain_controller := d.Get("domain_controller").(string)
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
	if domain_controller == "" {
		domain_controller = importOptionChecking(m.(*FortiClient).Cfg, "domain_controller")
		if domain_controller == "" {
			return fmt.Errorf("Parameter domain_controller is missing")
		}
		if err = d.Set("domain_controller", domain_controller); err != nil {
			return fmt.Errorf("Error set params domain_controller: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_controller"] = domain_controller

	o, err := c.ReadUserDomainControllerExtraServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserDomainControllerExtraServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDomainControllerExtraServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainControllerExtraServer resource from API: %v", err)
	}
	return nil
}

func flattenUserDomainControllerExtraServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerSourceIpAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerSourcePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserDomainControllerExtraServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenUserDomainControllerExtraServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserDomainControllerExtraServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenUserDomainControllerExtraServerIpAddress2edl(o["ip-address"], d, "ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-address"], "UserDomainControllerExtraServer-IpAddress"); ok {
			if err = d.Set("ip_address", vv); err != nil {
				return fmt.Errorf("Error reading ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("port", flattenUserDomainControllerExtraServerPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserDomainControllerExtraServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("source_ip_address", flattenUserDomainControllerExtraServerSourceIpAddress2edl(o["source-ip-address"], d, "source_ip_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-address"], "UserDomainControllerExtraServer-SourceIpAddress"); ok {
			if err = d.Set("source_ip_address", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_address: %v", err)
		}
	}

	if err = d.Set("source_port", flattenUserDomainControllerExtraServerSourcePort2edl(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "UserDomainControllerExtraServer-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	return nil
}

func flattenUserDomainControllerExtraServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserDomainControllerExtraServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourceIpAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourcePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserDomainControllerExtraServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserDomainControllerExtraServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok || d.HasChange("ip_address") {
		t, err := expandUserDomainControllerExtraServerIpAddress2edl(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserDomainControllerExtraServerPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_address"); ok || d.HasChange("source_ip_address") {
		t, err := expandUserDomainControllerExtraServerSourceIpAddress2edl(d, v, "source_ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandUserDomainControllerExtraServerSourcePort2edl(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	return &obj, nil
}
