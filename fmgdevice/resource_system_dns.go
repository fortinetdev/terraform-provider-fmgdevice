// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsUpdate,
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

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
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_over_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fqdn_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fqdn_max_refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fqdn_min_refresh": &schema.Schema{
				Type:     schema.TypeInt,
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
			"log": &schema.Schema{
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
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
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
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemDns")

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDns(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsAltPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsAltSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDnsOverTls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsFqdnCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsFqdnMaxRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsFqdnMinRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsServerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsServerSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("alt_primary", flattenSystemDnsAltPrimary(o["alt-primary"], d, "alt_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-primary"], "SystemDns-AltPrimary"); ok {
			if err = d.Set("alt_primary", vv); err != nil {
				return fmt.Errorf("Error reading alt_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_primary: %v", err)
		}
	}

	if err = d.Set("alt_secondary", flattenSystemDnsAltSecondary(o["alt-secondary"], d, "alt_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["alt-secondary"], "SystemDns-AltSecondary"); ok {
			if err = d.Set("alt_secondary", vv); err != nil {
				return fmt.Errorf("Error reading alt_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alt_secondary: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", flattenSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-notfound-responses"], "SystemDns-CacheNotfoundResponses"); ok {
			if err = d.Set("cache_notfound_responses", vv); err != nil {
				return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", flattenSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-cache-limit"], "SystemDns-DnsCacheLimit"); ok {
			if err = d.Set("dns_cache_limit", vv); err != nil {
				return fmt.Errorf("Error reading dns_cache_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", flattenSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-cache-ttl"], "SystemDns-DnsCacheTtl"); ok {
			if err = d.Set("dns_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("dns_over_tls", flattenSystemDnsDnsOverTls(o["dns-over-tls"], d, "dns_over_tls")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-over-tls"], "SystemDns-DnsOverTls"); ok {
			if err = d.Set("dns_over_tls", vv); err != nil {
				return fmt.Errorf("Error reading dns_over_tls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_over_tls: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemDns-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("fqdn_cache_ttl", flattenSystemDnsFqdnCacheTtl(o["fqdn-cache-ttl"], d, "fqdn_cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-cache-ttl"], "SystemDns-FqdnCacheTtl"); ok {
			if err = d.Set("fqdn_cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_cache_ttl: %v", err)
		}
	}

	if err = d.Set("fqdn_max_refresh", flattenSystemDnsFqdnMaxRefresh(o["fqdn-max-refresh"], d, "fqdn_max_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-max-refresh"], "SystemDns-FqdnMaxRefresh"); ok {
			if err = d.Set("fqdn_max_refresh", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_max_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_max_refresh: %v", err)
		}
	}

	if err = d.Set("fqdn_min_refresh", flattenSystemDnsFqdnMinRefresh(o["fqdn-min-refresh"], d, "fqdn_min_refresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn-min-refresh"], "SystemDns-FqdnMinRefresh"); ok {
			if err = d.Set("fqdn_min_refresh", vv); err != nil {
				return fmt.Errorf("Error reading fqdn_min_refresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn_min_refresh: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDnsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemDns-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemDnsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystemDns-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip6_primary", flattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-primary"], "SystemDns-Ip6Primary"); ok {
			if err = d.Set("ip6_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-secondary"], "SystemDns-Ip6Secondary"); ok {
			if err = d.Set("ip6_secondary", vv); err != nil {
				return fmt.Errorf("Error reading ip6_secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("log", flattenSystemDnsLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "SystemDns-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemDnsPrimary(o["primary"], d, "primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary"], "SystemDns-Primary"); ok {
			if err = d.Set("primary", vv); err != nil {
				return fmt.Errorf("Error reading primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemDnsProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemDns-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("retry", flattenSystemDnsRetry(o["retry"], d, "retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry"], "SystemDns-Retry"); ok {
			if err = d.Set("retry", vv); err != nil {
				return fmt.Errorf("Error reading retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary"], "SystemDns-Secondary"); ok {
			if err = d.Set("secondary", vv); err != nil {
				return fmt.Errorf("Error reading secondary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if err = d.Set("server_hostname", flattenSystemDnsServerHostname(o["server-hostname"], d, "server_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-hostname"], "SystemDns-ServerHostname"); ok {
			if err = d.Set("server_hostname", vv); err != nil {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_hostname: %v", err)
		}
	}

	if err = d.Set("server_select_method", flattenSystemDnsServerSelectMethod(o["server-select-method"], d, "server_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-select-method"], "SystemDns-ServerSelectMethod"); ok {
			if err = d.Set("server_select_method", vv); err != nil {
				return fmt.Errorf("Error reading server_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_select_method: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemDnsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemDns-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemDnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-certificate"], "SystemDns-SslCertificate"); ok {
			if err = d.Set("ssl_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemDnsTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "SystemDns-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDnsAltPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsAltSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsCacheNotfoundResponses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsFqdnCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsFqdnMaxRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsFqdnMinRefresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsServerSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alt_primary"); ok || d.HasChange("alt_primary") {
		t, err := expandSystemDnsAltPrimary(d, v, "alt_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-primary"] = t
		}
	}

	if v, ok := d.GetOk("alt_secondary"); ok || d.HasChange("alt_secondary") {
		t, err := expandSystemDnsAltSecondary(d, v, "alt_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-secondary"] = t
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok || d.HasChange("cache_notfound_responses") {
		t, err := expandSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-notfound-responses"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_limit"); ok || d.HasChange("dns_cache_limit") {
		t, err := expandSystemDnsDnsCacheLimit(d, v, "dns_cache_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-limit"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok || d.HasChange("dns_cache_ttl") {
		t, err := expandSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("dns_over_tls"); ok || d.HasChange("dns_over_tls") {
		t, err := expandSystemDnsDnsOverTls(d, v, "dns_over_tls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-over-tls"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemDnsDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_cache_ttl"); ok || d.HasChange("fqdn_cache_ttl") {
		t, err := expandSystemDnsFqdnCacheTtl(d, v, "fqdn_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_max_refresh"); ok || d.HasChange("fqdn_max_refresh") {
		t, err := expandSystemDnsFqdnMaxRefresh(d, v, "fqdn_max_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-max-refresh"] = t
		}
	}

	if v, ok := d.GetOk("fqdn_min_refresh"); ok || d.HasChange("fqdn_min_refresh") {
		t, err := expandSystemDnsFqdnMinRefresh(d, v, "fqdn_min_refresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn-min-refresh"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemDnsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystemDnsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok || d.HasChange("ip6_primary") {
		t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok || d.HasChange("ip6_secondary") {
		t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandSystemDnsLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("primary"); ok || d.HasChange("primary") {
		t, err := expandSystemDnsPrimary(d, v, "primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemDnsProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok || d.HasChange("retry") {
		t, err := expandSystemDnsRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok || d.HasChange("secondary") {
		t, err := expandSystemDnsSecondary(d, v, "secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok || d.HasChange("server_hostname") {
		t, err := expandSystemDnsServerHostname(d, v, "server_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-hostname"] = t
		}
	}

	if v, ok := d.GetOk("server_select_method"); ok || d.HasChange("server_select_method") {
		t, err := expandSystemDnsServerSelectMethod(d, v, "server_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-select-method"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemDnsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandSystemDnsSslCertificate(d, v, "ssl_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandSystemDnsTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
