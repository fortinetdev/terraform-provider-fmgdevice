// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure server load balancing health monitors.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallLdbMonitorCreate,
		Read:   resourceFirewallLdbMonitorRead,
		Update: resourceFirewallLdbMonitorUpdate,
		Delete: resourceFirewallLdbMonitorDelete,

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
			"dns_match_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_max_redirects": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallLdbMonitorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallLdbMonitor resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallLdbMonitor(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallLdbMonitor(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallLdbMonitor resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallLdbMonitor(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallLdbMonitor resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallLdbMonitor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallLdbMonitor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallLdbMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLdbMonitorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallLdbMonitor(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallLdbMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallLdbMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLdbMonitor resource from API: %v", err)
	}
	return nil
}

func flattenFirewallLdbMonitorDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorDnsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMaxRedirects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallLdbMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_match_ip", flattenFirewallLdbMonitorDnsMatchIp(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-match-ip"], "FirewallLdbMonitor-DnsMatchIp"); ok {
			if err = d.Set("dns_match_ip", vv); err != nil {
				return fmt.Errorf("Error reading dns_match_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_protocol", flattenFirewallLdbMonitorDnsProtocol(o["dns-protocol"], d, "dns_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-protocol"], "FirewallLdbMonitor-DnsProtocol"); ok {
			if err = d.Set("dns_protocol", vv); err != nil {
				return fmt.Errorf("Error reading dns_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_protocol: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenFirewallLdbMonitorDnsRequestDomain(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-request-domain"], "FirewallLdbMonitor-DnsRequestDomain"); ok {
			if err = d.Set("dns_request_domain", vv); err != nil {
				return fmt.Errorf("Error reading dns_request_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("http_get", flattenFirewallLdbMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "FirewallLdbMonitor-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenFirewallLdbMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "FirewallLdbMonitor-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("http_max_redirects", flattenFirewallLdbMonitorHttpMaxRedirects(o["http-max-redirects"], d, "http_max_redirects")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-max-redirects"], "FirewallLdbMonitor-HttpMaxRedirects"); ok {
			if err = d.Set("http_max_redirects", vv); err != nil {
				return fmt.Errorf("Error reading http_max_redirects: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_max_redirects: %v", err)
		}
	}

	if err = d.Set("interval", flattenFirewallLdbMonitorInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "FirewallLdbMonitor-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallLdbMonitorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallLdbMonitor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallLdbMonitorPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallLdbMonitor-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("retry", flattenFirewallLdbMonitorRetry(o["retry"], d, "retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry"], "FirewallLdbMonitor-Retry"); ok {
			if err = d.Set("retry", vv); err != nil {
				return fmt.Errorf("Error reading retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenFirewallLdbMonitorSrcIp(o["src-ip"], d, "src_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip"], "FirewallLdbMonitor-SrcIp"); ok {
			if err = d.Set("src_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	if err = d.Set("timeout", flattenFirewallLdbMonitorTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "FirewallLdbMonitor-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallLdbMonitorType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallLdbMonitor-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenFirewallLdbMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallLdbMonitorDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorDnsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMaxRedirects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallLdbMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_match_ip"); ok || d.HasChange("dns_match_ip") {
		t, err := expandFirewallLdbMonitorDnsMatchIp(d, v, "dns_match_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_protocol"); ok || d.HasChange("dns_protocol") {
		t, err := expandFirewallLdbMonitorDnsProtocol(d, v, "dns_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-protocol"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok || d.HasChange("dns_request_domain") {
		t, err := expandFirewallLdbMonitorDnsRequestDomain(d, v, "dns_request_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandFirewallLdbMonitorHttpGet(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandFirewallLdbMonitorHttpMatch(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("http_max_redirects"); ok || d.HasChange("http_max_redirects") {
		t, err := expandFirewallLdbMonitorHttpMaxRedirects(d, v, "http_max_redirects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-max-redirects"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandFirewallLdbMonitorInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallLdbMonitorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallLdbMonitorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok || d.HasChange("retry") {
		t, err := expandFirewallLdbMonitorRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok || d.HasChange("src_ip") {
		t, err := expandFirewallLdbMonitorSrcIp(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandFirewallLdbMonitorTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallLdbMonitorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
