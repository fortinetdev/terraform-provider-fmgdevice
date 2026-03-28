// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebPortalSplitDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalSplitDnsCreate,
		Read:   resourceVpnSslWebPortalSplitDnsRead,
		Update: resourceVpnSslWebPortalSplitDnsUpdate,
		Delete: resourceVpnSslWebPortalSplitDnsDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnSslWebPortalSplitDnsCreate(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	obj, err := getObjectVpnSslWebPortalSplitDns(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebPortalSplitDns resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadVpnSslWebPortalSplitDns(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateVpnSslWebPortalSplitDns(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating VpnSslWebPortalSplitDns resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateVpnSslWebPortalSplitDns(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating VpnSslWebPortalSplitDns resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceVpnSslWebPortalSplitDnsRead(d, m)
			} else {
				return fmt.Errorf("Error creating VpnSslWebPortalSplitDns resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnSslWebPortalSplitDnsRead(d, m)
}

func resourceVpnSslWebPortalSplitDnsUpdate(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	obj, err := getObjectVpnSslWebPortalSplitDns(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalSplitDns resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVpnSslWebPortalSplitDns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortalSplitDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnSslWebPortalSplitDnsRead(d, m)
}

func resourceVpnSslWebPortalSplitDnsDelete(d *schema.ResourceData, m interface{}) error {
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
	portal := d.Get("portal").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	wsParams["adom"] = adomv

	err = c.DeleteVpnSslWebPortalSplitDns(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortalSplitDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalSplitDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	portal := d.Get("portal").(string)
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
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["portal"] = portal

	o, err := c.ReadVpnSslWebPortalSplitDns(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VpnSslWebPortalSplitDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortalSplitDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortalSplitDns resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalSplitDnsDnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortalSplitDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dns_server1", flattenVpnSslWebPortalSplitDnsDnsServer12edl(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "VpnSslWebPortalSplitDns-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenVpnSslWebPortalSplitDnsDnsServer22edl(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "VpnSslWebPortalSplitDns-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("domains", flattenVpnSslWebPortalSplitDnsDomains2edl(o["domains"], d, "domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domains"], "VpnSslWebPortalSplitDns-Domains"); ok {
			if err = d.Set("domains", vv); err != nil {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}

	if err = d.Set("fosid", flattenVpnSslWebPortalSplitDnsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VpnSslWebPortalSplitDns-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnSslWebPortalSplitDnsIpv6DnsServer12edl(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "VpnSslWebPortalSplitDns-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnSslWebPortalSplitDnsIpv6DnsServer22edl(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "VpnSslWebPortalSplitDns-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebPortalSplitDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebPortalSplitDnsDnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortalSplitDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandVpnSslWebPortalSplitDnsDnsServer12edl(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandVpnSslWebPortalSplitDnsDnsServer22edl(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandVpnSslWebPortalSplitDnsDomains2edl(d, v, "domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVpnSslWebPortalSplitDnsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandVpnSslWebPortalSplitDnsIpv6DnsServer12edl(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandVpnSslWebPortalSplitDnsIpv6DnsServer22edl(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	return &obj, nil
}
