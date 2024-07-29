// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DDNS.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDdns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDdnsCreate,
		Read:   resourceSystemDdnsRead,
		Update: resourceSystemDdnsUpdate,
		Delete: resourceSystemDdnsDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bound_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clear_text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_keyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_server_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddnsid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDdnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemDdns(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDdns resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDdns(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemDdns resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "ddnsid")))

	return resourceSystemDdnsRead(d, m)
}

func resourceSystemDdnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemDdns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDdns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDdns(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDdns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "ddnsid")))

	return resourceSystemDdnsRead(d, m)
}

func resourceSystemDdnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemDdns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDdns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDdnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemDdns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDdns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDdns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDdns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDdnsAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsBoundIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsClearText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsServerAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDdnsDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsDdnsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDdnsServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDdnsUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDdnsUsePublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDdns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenSystemDdnsAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "SystemDdns-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("bound_ip", flattenSystemDdnsBoundIp(o["bound-ip"], d, "bound_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["bound-ip"], "SystemDdns-BoundIp"); ok {
			if err = d.Set("bound_ip", vv); err != nil {
				return fmt.Errorf("Error reading bound_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bound_ip: %v", err)
		}
	}

	if err = d.Set("clear_text", flattenSystemDdnsClearText(o["clear-text"], d, "clear_text")); err != nil {
		if vv, ok := fortiAPIPatch(o["clear-text"], "SystemDdns-ClearText"); ok {
			if err = d.Set("clear_text", vv); err != nil {
				return fmt.Errorf("Error reading clear_text: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clear_text: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDdnsDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "SystemDdns-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_domain", flattenSystemDdnsDdnsDomain(o["ddns-domain"], d, "ddns_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-domain"], "SystemDdns-DdnsDomain"); ok {
			if err = d.Set("ddns_domain", vv); err != nil {
				return fmt.Errorf("Error reading ddns_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_domain: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDdnsDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "SystemDdns-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server", flattenSystemDdnsDdnsServer(o["ddns-server"], d, "ddns_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server"], "SystemDdns-DdnsServer"); ok {
			if err = d.Set("ddns_server", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server: %v", err)
		}
	}

	if err = d.Set("ddns_server_addr", flattenSystemDdnsDdnsServerAddr(o["ddns-server-addr"], d, "ddns_server_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-addr"], "SystemDdns-DdnsServerAddr"); ok {
			if err = d.Set("ddns_server_addr", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_addr: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDdnsDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "SystemDdns-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_sn", flattenSystemDdnsDdnsSn(o["ddns-sn"], d, "ddns_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-sn"], "SystemDdns-DdnsSn"); ok {
			if err = d.Set("ddns_sn", vv); err != nil {
				return fmt.Errorf("Error reading ddns_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_sn: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDdnsDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "SystemDdns-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_username", flattenSystemDdnsDdnsUsername(o["ddns-username"], d, "ddns_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-username"], "SystemDdns-DdnsUsername"); ok {
			if err = d.Set("ddns_username", vv); err != nil {
				return fmt.Errorf("Error reading ddns_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_username: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDdnsDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "SystemDdns-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddnsid", flattenSystemDdnsDdnsid(o["ddnsid"], d, "ddnsid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddnsid"], "SystemDdns-Ddnsid"); ok {
			if err = d.Set("ddnsid", vv); err != nil {
				return fmt.Errorf("Error reading ddnsid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddnsid: %v", err)
		}
	}

	if err = d.Set("monitor_interface", flattenSystemDdnsMonitorInterface(o["monitor-interface"], d, "monitor_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-interface"], "SystemDdns-MonitorInterface"); ok {
			if err = d.Set("monitor_interface", vv); err != nil {
				return fmt.Errorf("Error reading monitor_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_interface: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemDdnsServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemDdns-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemDdnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "SystemDdns-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemDdnsUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "SystemDdns-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("use_public_ip", flattenSystemDdnsUsePublicIp(o["use-public-ip"], d, "use_public_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-public-ip"], "SystemDdns-UsePublicIp"); ok {
			if err = d.Set("use_public_ip", vv); err != nil {
				return fmt.Errorf("Error reading use_public_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_public_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemDdnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDdnsAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsBoundIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsClearText(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDdnsDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDdnsDdnsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsServerAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDdnsDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDdnsServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDdnsUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsUsePublicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDdns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandSystemDdnsAddrType(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("bound_ip"); ok || d.HasChange("bound_ip") {
		t, err := expandSystemDdnsBoundIp(d, v, "bound_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bound-ip"] = t
		}
	}

	if v, ok := d.GetOk("clear_text"); ok || d.HasChange("clear_text") {
		t, err := expandSystemDdnsClearText(d, v, "clear_text")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clear-text"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandSystemDdnsDdnsAuth(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_domain"); ok || d.HasChange("ddns_domain") {
		t, err := expandSystemDdnsDdnsDomain(d, v, "ddns_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-domain"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandSystemDdnsDdnsKey(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandSystemDdnsDdnsKeyname(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_password"); ok || d.HasChange("ddns_password") {
		t, err := expandSystemDdnsDdnsPassword(d, v, "ddns_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-password"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server"); ok || d.HasChange("ddns_server") {
		t, err := expandSystemDdnsDdnsServer(d, v, "ddns_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_addr"); ok || d.HasChange("ddns_server_addr") {
		t, err := expandSystemDdnsDdnsServerAddr(d, v, "ddns_server_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-addr"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandSystemDdnsDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_sn"); ok || d.HasChange("ddns_sn") {
		t, err := expandSystemDdnsDdnsSn(d, v, "ddns_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-sn"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandSystemDdnsDdnsTtl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_username"); ok || d.HasChange("ddns_username") {
		t, err := expandSystemDdnsDdnsUsername(d, v, "ddns_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-username"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandSystemDdnsDdnsZone(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("ddnsid"); ok || d.HasChange("ddnsid") {
		t, err := expandSystemDdnsDdnsid(d, v, "ddnsid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddnsid"] = t
		}
	}

	if v, ok := d.GetOk("monitor_interface"); ok || d.HasChange("monitor_interface") {
		t, err := expandSystemDdnsMonitorInterface(d, v, "monitor_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-interface"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemDdnsServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandSystemDdnsSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandSystemDdnsUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("use_public_ip"); ok || d.HasChange("use_public_ip") {
		t, err := expandSystemDdnsUsePublicIp(d, v, "use_public_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-public-ip"] = t
		}
	}

	return &obj, nil
}
