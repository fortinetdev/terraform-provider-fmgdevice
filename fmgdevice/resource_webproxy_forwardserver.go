// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure forward-server addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyForwardServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyForwardServerCreate,
		Read:   resourceWebProxyForwardServerRead,
		Update: resourceWebProxyForwardServerUpdate,
		Delete: resourceWebProxyForwardServerDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"masquerade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_down_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyForwardServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyForwardServer(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebProxyForwardServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebProxyForwardServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebProxyForwardServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebProxyForwardServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebProxyForwardServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyForwardServerRead(d, m)
}

func resourceWebProxyForwardServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyForwardServer(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebProxyForwardServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebProxyForwardServerRead(d, m)
}

func resourceWebProxyForwardServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebProxyForwardServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyForwardServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyForwardServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebProxyForwardServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyForwardServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServer resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyForwardServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyForwardServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerMasquerade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerServerDownOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebProxyForwardServerProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebProxyForwardServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenWebProxyForwardServerAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "WebProxyForwardServer-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyForwardServerComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebProxyForwardServer-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenWebProxyForwardServerFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "WebProxyForwardServer-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenWebProxyForwardServerHealthcheck(o["healthcheck"], d, "healthcheck")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck"], "WebProxyForwardServer-Healthcheck"); ok {
			if err = d.Set("healthcheck", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("interface", flattenWebProxyForwardServerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "WebProxyForwardServer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenWebProxyForwardServerInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "WebProxyForwardServer-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebProxyForwardServerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WebProxyForwardServer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenWebProxyForwardServerIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "WebProxyForwardServer-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("masquerade", flattenWebProxyForwardServerMasquerade(o["masquerade"], d, "masquerade")); err != nil {
		if vv, ok := fortiAPIPatch(o["masquerade"], "WebProxyForwardServer-Masquerade"); ok {
			if err = d.Set("masquerade", vv); err != nil {
				return fmt.Errorf("Error reading masquerade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading masquerade: %v", err)
		}
	}

	if err = d.Set("monitor", flattenWebProxyForwardServerMonitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "WebProxyForwardServer-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("name", flattenWebProxyForwardServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebProxyForwardServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenWebProxyForwardServerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WebProxyForwardServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server_down_option", flattenWebProxyForwardServerServerDownOption(o["server-down-option"], d, "server_down_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-down-option"], "WebProxyForwardServer-ServerDownOption"); ok {
			if err = d.Set("server_down_option", vv); err != nil {
				return fmt.Errorf("Error reading server_down_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_down_option: %v", err)
		}
	}

	if err = d.Set("username", flattenWebProxyForwardServerUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "WebProxyForwardServer-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenWebProxyForwardServerVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "WebProxyForwardServer-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("ippool", flattenWebProxyForwardServerIppool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "WebProxyForwardServer-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWebProxyForwardServerProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WebProxyForwardServer-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenWebProxyForwardServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyForwardServerAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyForwardServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerMasquerade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyForwardServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerServerDownOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebProxyForwardServerProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebProxyForwardServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandWebProxyForwardServerAddrType(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebProxyForwardServerComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandWebProxyForwardServerFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok || d.HasChange("healthcheck") {
		t, err := expandWebProxyForwardServerHealthcheck(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandWebProxyForwardServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandWebProxyForwardServerInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWebProxyForwardServerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandWebProxyForwardServerIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("masquerade"); ok || d.HasChange("masquerade") {
		t, err := expandWebProxyForwardServerMasquerade(d, v, "masquerade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["masquerade"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandWebProxyForwardServerMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebProxyForwardServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandWebProxyForwardServerPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWebProxyForwardServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server_down_option"); ok || d.HasChange("server_down_option") {
		t, err := expandWebProxyForwardServerServerDownOption(d, v, "server_down_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-down-option"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandWebProxyForwardServerUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandWebProxyForwardServerVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok || d.HasChange("ippool") {
		t, err := expandWebProxyForwardServerIppool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWebProxyForwardServerProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
