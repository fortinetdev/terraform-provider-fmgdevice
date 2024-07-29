// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS servers for a non-management VDOM.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomDnsUpdate,
		Read:   resourceSystemVdomDnsRead,
		Update: resourceSystemVdomDnsUpdate,
		Delete: resourceSystemVdomDnsDelete,

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
			"dns_over_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alt_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alt_secondary": &schema.Schema{
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
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
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
			"vdom_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomDnsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVdomDns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomDns(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemVdomDns")

	return resourceSystemVdomDnsRead(d, m)
}

func resourceSystemVdomDnsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVdomDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomDnsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomDnsDnsOverTls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsAltPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsAltSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsServerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomDnsServerSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomDnsSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVdomDnsVdomDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_over_tls", flattenSystemVdomDnsDnsOverTls(o["dns-over-tls"], d, "dns_over_tls")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-over-tls"], "SystemVdomDns-DnsOverTls"); ok {
			if err = d.Set("dns_over_tls", vv); err != nil {
				return fmt.Errorf("Error reading dns_over_tls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_over_tls: %v", err)
		}
	}

	if err = d.Set("alt_primary", flattenSystemVdomDnsAltPrimary(o["alt-primary"], d, "alt_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-primary"], "SystemVdomDns-AltPrimary"); ok {
			if err = d.Set("alt_primary", vv); err != nil {
				return fmt.Errorf("Error reading alt_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_primary: %v", err)
		}
	}

	if err = d.Set("alt_secondary", flattenSystemVdomDnsAltSecondary(o["alt-secondary"], d, "alt_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-secondary"], "SystemVdomDns-AltSecondary"); ok {
			if err = d.Set("alt_secondary", vv); err != nil {
				return fmt.Errorf("Error reading alt_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_secondary: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVdomDnsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemVdomDns-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemVdomDnsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemVdomDns-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip6_primary", flattenSystemVdomDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-primary"], "SystemVdomDns-Ip6Primary"); ok {
			if err = d.Set("ip6_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemVdomDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-secondary"], "SystemVdomDns-Ip6Secondary"); ok {
			if err = d.Set("ip6_secondary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemVdomDnsPrimary(o["primary"], d, "primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary"], "SystemVdomDns-Primary"); ok {
			if err = d.Set("primary", vv); err != nil {
				return fmt.Errorf("Error reading primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemVdomDnsProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemVdomDns-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemVdomDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary"], "SystemVdomDns-Secondary"); ok {
			if err = d.Set("secondary", vv); err != nil {
				return fmt.Errorf("Error reading secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if err = d.Set("server_hostname", flattenSystemVdomDnsServerHostname(o["server-hostname"], d, "server_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-hostname"], "SystemVdomDns-ServerHostname"); ok {
			if err = d.Set("server_hostname", vv); err != nil {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_hostname: %v", err)
		}
	}

	if err = d.Set("server_select_method", flattenSystemVdomDnsServerSelectMethod(o["server-select-method"], d, "server_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-select-method"], "SystemVdomDns-ServerSelectMethod"); ok {
			if err = d.Set("server_select_method", vv); err != nil {
				return fmt.Errorf("Error reading server_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_select_method: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomDnsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemVdomDns-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemVdomDnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "SystemVdomDns-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("vdom_dns", flattenSystemVdomDnsVdomDns(o["vdom-dns"], d, "vdom_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-dns"], "SystemVdomDns-VdomDns"); ok {
			if err = d.Set("vdom_dns", vv); err != nil {
				return fmt.Errorf("Error reading vdom_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_dns: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsAltPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsAltSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomDnsSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsServerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomDnsServerSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVdomDnsVdomDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_over_tls"); ok || d.HasChange("dns_over_tls") {
		t, err := expandSystemVdomDnsDnsOverTls(d, v, "dns_over_tls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-over-tls"] = t
		}
	}

	if v, ok := d.GetOk("alt_primary"); ok || d.HasChange("alt_primary") {
		t, err := expandSystemVdomDnsAltPrimary(d, v, "alt_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-primary"] = t
		}
	}

	if v, ok := d.GetOk("alt_secondary"); ok || d.HasChange("alt_secondary") {
		t, err := expandSystemVdomDnsAltSecondary(d, v, "alt_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-secondary"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemVdomDnsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemVdomDnsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok || d.HasChange("ip6_primary") {
		t, err := expandSystemVdomDnsIp6Primary(d, v, "ip6_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok || d.HasChange("ip6_secondary") {
		t, err := expandSystemVdomDnsIp6Secondary(d, v, "ip6_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("primary"); ok || d.HasChange("primary") {
		t, err := expandSystemVdomDnsPrimary(d, v, "primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemVdomDnsProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok || d.HasChange("secondary") {
		t, err := expandSystemVdomDnsSecondary(d, v, "secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok || d.HasChange("server_hostname") {
		t, err := expandSystemVdomDnsServerHostname(d, v, "server_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-hostname"] = t
		}
	}

	if v, ok := d.GetOk("server_select_method"); ok || d.HasChange("server_select_method") {
		t, err := expandSystemVdomDnsServerSelectMethod(d, v, "server_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-select-method"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemVdomDnsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandSystemVdomDnsSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("vdom_dns"); ok || d.HasChange("vdom_dns") {
		t, err := expandSystemVdomDnsVdomDns(d, v, "vdom_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-dns"] = t
		}
	}

	return &obj, nil
}
