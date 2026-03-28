// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Realm.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebRealmCreate,
		Read:   resourceVpnSslWebRealmRead,
		Update: resourceVpnSslWebRealmUpdate,
		Delete: resourceVpnSslWebRealmDelete,

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
			"login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_concurrent_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_path": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"virtual_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_host_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_host_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnSslWebRealmCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebRealm resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("url_path")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnSslWebRealm(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnSslWebRealm(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnSslWebRealm resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateVpnSslWebRealm(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnSslWebRealm resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "url_path"))

	return resourceVpnSslWebRealmRead(d, m)
}

func resourceVpnSslWebRealmUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebRealm resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebRealm(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebRealm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "url_path"))

	return resourceVpnSslWebRealmRead(d, m)
}

func resourceVpnSslWebRealmDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslWebRealm(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebRealmRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebRealm(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebRealm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebRealm resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebRealmLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmMaxConcurrentUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslWebRealmUrlPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHostOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHostServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnSslWebRealm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("login_page", flattenVpnSslWebRealmLoginPage(o["login-page"], d, "login_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-page"], "VpnSslWebRealm-LoginPage"); ok {
			if err = d.Set("login_page", vv); err != nil {
				return fmt.Errorf("Error reading login_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_page: %v", err)
		}
	}

	if err = d.Set("max_concurrent_user", flattenVpnSslWebRealmMaxConcurrentUser(o["max-concurrent-user"], d, "max_concurrent_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-user"], "VpnSslWebRealm-MaxConcurrentUser"); ok {
			if err = d.Set("max_concurrent_user", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_user: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenVpnSslWebRealmNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "VpnSslWebRealm-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenVpnSslWebRealmRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "VpnSslWebRealm-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenVpnSslWebRealmRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "VpnSslWebRealm-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("url_path", flattenVpnSslWebRealmUrlPath(o["url-path"], d, "url_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-path"], "VpnSslWebRealm-UrlPath"); ok {
			if err = d.Set("url_path", vv); err != nil {
				return fmt.Errorf("Error reading url_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_path: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenVpnSslWebRealmVirtualHost(o["virtual-host"], d, "virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host"], "VpnSslWebRealm-VirtualHost"); ok {
			if err = d.Set("virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	if err = d.Set("virtual_host_only", flattenVpnSslWebRealmVirtualHostOnly(o["virtual-host-only"], d, "virtual_host_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host-only"], "VpnSslWebRealm-VirtualHostOnly"); ok {
			if err = d.Set("virtual_host_only", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host_only: %v", err)
		}
	}

	if err = d.Set("virtual_host_server_cert", flattenVpnSslWebRealmVirtualHostServerCert(o["virtual-host-server-cert"], d, "virtual_host_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-host-server-cert"], "VpnSslWebRealm-VirtualHostServerCert"); ok {
			if err = d.Set("virtual_host_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading virtual_host_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_host_server_cert: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebRealmLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmMaxConcurrentUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslWebRealmUrlPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHostOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHostServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnSslWebRealm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("login_page"); ok || d.HasChange("login_page") {
		t, err := expandVpnSslWebRealmLoginPage(d, v, "login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-page"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_user"); ok || d.HasChange("max_concurrent_user") {
		t, err := expandVpnSslWebRealmMaxConcurrentUser(d, v, "max_concurrent_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-user"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok || d.HasChange("nas_ip") {
		t, err := expandVpnSslWebRealmNasIp(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok || d.HasChange("radius_port") {
		t, err := expandVpnSslWebRealmRadiusPort(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandVpnSslWebRealmRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("url_path"); ok || d.HasChange("url_path") {
		t, err := expandVpnSslWebRealmUrlPath(d, v, "url_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-path"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok || d.HasChange("virtual_host") {
		t, err := expandVpnSslWebRealmVirtualHost(d, v, "virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host_only"); ok || d.HasChange("virtual_host_only") {
		t, err := expandVpnSslWebRealmVirtualHostOnly(d, v, "virtual_host_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host-only"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host_server_cert"); ok || d.HasChange("virtual_host_server_cert") {
		t, err := expandVpnSslWebRealmVirtualHostServerCert(d, v, "virtual_host_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host-server-cert"] = t
		}
	}

	return &obj, nil
}
