// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCPv6 servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcp6Server() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcp6ServerCreate,
		Read:   resourceSystemDhcp6ServerRead,
		Update: resourceSystemDhcp6ServerUpdate,
		Delete: resourceSystemDhcp6ServerDelete,

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
			"delegated_prefix_iaid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dns_search_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"option1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"option2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"option3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"rapid_commit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemDhcp6ServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemDhcp6Server(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6Server resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDhcp6Server(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6Server resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerRead(d, m)
}

func resourceSystemDhcp6ServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemDhcp6Server(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6Server resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDhcp6Server(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6Server resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemDhcp6ServerRead(d, m)
}

func resourceSystemDhcp6ServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemDhcp6Server(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcp6Server resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcp6ServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcp6Server(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcp6Server(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcp6ServerDelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsSearchList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcp6ServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenSystemDhcp6ServerIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcp6ServerIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenSystemDhcp6ServerIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcp6ServerIpRangeVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcp6ServerIpRangeVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-IpRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcp6ServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcp6ServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			v := flattenSystemDhcp6ServerOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcp6ServerOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenSystemDhcp6ServerOptionsIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemDhcp6ServerOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemDhcp6ServerOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenSystemDhcp6ServerOptionsVciMatch(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenSystemDhcp6ServerOptionsVciString(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcp6ServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDhcp6ServerOption1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if _, ok := i["end-prefix"]; ok {
			v := flattenSystemDhcp6ServerPrefixRangeEndPrefix(i["end-prefix"], d, pre_append)
			tmp["end_prefix"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-PrefixRange-EndPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemDhcp6ServerPrefixRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-PrefixRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := i["prefix-length"]; ok {
			v := flattenSystemDhcp6ServerPrefixRangePrefixLength(i["prefix-length"], d, pre_append)
			tmp["prefix_length"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-PrefixRange-PrefixLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if _, ok := i["start-prefix"]; ok {
			v := flattenSystemDhcp6ServerPrefixRangeStartPrefix(i["start-prefix"], d, pre_append)
			tmp["start_prefix"] = fortiAPISubPartPatch(v, "SystemDhcp6Server-PrefixRange-StartPrefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemDhcp6ServerPrefixRangeEndPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangePrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeStartPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerRapidCommit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemDhcp6Server(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("delegated_prefix_iaid", flattenSystemDhcp6ServerDelegatedPrefixIaid(o["delegated-prefix-iaid"], d, "delegated_prefix_iaid")); err != nil {
		if vv, ok := fortiAPIPatch(o["delegated-prefix-iaid"], "SystemDhcp6Server-DelegatedPrefixIaid"); ok {
			if err = d.Set("delegated_prefix_iaid", vv); err != nil {
				return fmt.Errorf("Error reading delegated_prefix_iaid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delegated_prefix_iaid: %v", err)
		}
	}

	if err = d.Set("dns_search_list", flattenSystemDhcp6ServerDnsSearchList(o["dns-search-list"], d, "dns_search_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-search-list"], "SystemDhcp6Server-DnsSearchList"); ok {
			if err = d.Set("dns_search_list", vv); err != nil {
				return fmt.Errorf("Error reading dns_search_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_search_list: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcp6ServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "SystemDhcp6Server-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcp6ServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "SystemDhcp6Server-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcp6ServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "SystemDhcp6Server-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenSystemDhcp6ServerDnsServer4(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "SystemDhcp6Server-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcp6ServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "SystemDhcp6Server-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcp6ServerDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "SystemDhcp6Server-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemDhcp6ServerId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemDhcp6Server-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDhcp6ServerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemDhcp6Server-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenSystemDhcp6ServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "SystemDhcp6Server-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcp6Server-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "SystemDhcp6Server-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcp6ServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "SystemDhcp6Server-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenSystemDhcp6ServerOptions(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "SystemDhcp6Server-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSystemDhcp6ServerOptions(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "SystemDhcp6Server-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("option1", flattenSystemDhcp6ServerOption1(o["option1"], d, "option1")); err != nil {
		if vv, ok := fortiAPIPatch(o["option1"], "SystemDhcp6Server-Option1"); ok {
			if err = d.Set("option1", vv); err != nil {
				return fmt.Errorf("Error reading option1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option1: %v", err)
		}
	}

	if err = d.Set("option2", flattenSystemDhcp6ServerOption2(o["option2"], d, "option2")); err != nil {
		if vv, ok := fortiAPIPatch(o["option2"], "SystemDhcp6Server-Option2"); ok {
			if err = d.Set("option2", vv); err != nil {
				return fmt.Errorf("Error reading option2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option2: %v", err)
		}
	}

	if err = d.Set("option3", flattenSystemDhcp6ServerOption3(o["option3"], d, "option3")); err != nil {
		if vv, ok := fortiAPIPatch(o["option3"], "SystemDhcp6Server-Option3"); ok {
			if err = d.Set("option3", vv); err != nil {
				return fmt.Errorf("Error reading option3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option3: %v", err)
		}
	}

	if err = d.Set("prefix_mode", flattenSystemDhcp6ServerPrefixMode(o["prefix-mode"], d, "prefix_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-mode"], "SystemDhcp6Server-PrefixMode"); ok {
			if err = d.Set("prefix_mode", vv); err != nil {
				return fmt.Errorf("Error reading prefix_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["prefix-range"], "SystemDhcp6Server-PrefixRange"); ok {
				if err = d.Set("prefix_range", vv); err != nil {
					return fmt.Errorf("Error reading prefix_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading prefix_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("prefix_range"); ok {
			if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["prefix-range"], "SystemDhcp6Server-PrefixRange"); ok {
					if err = d.Set("prefix_range", vv); err != nil {
						return fmt.Errorf("Error reading prefix_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading prefix_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("rapid_commit", flattenSystemDhcp6ServerRapidCommit(o["rapid-commit"], d, "rapid_commit")); err != nil {
		if vv, ok := fortiAPIPatch(o["rapid-commit"], "SystemDhcp6Server-RapidCommit"); ok {
			if err = d.Set("rapid_commit", vv); err != nil {
				return fmt.Errorf("Error reading rapid_commit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rapid_commit: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDhcp6ServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemDhcp6Server-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("subnet", flattenSystemDhcp6ServerSubnet(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "SystemDhcp6Server-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenSystemDhcp6ServerUpstreamInterface(o["upstream-interface"], d, "upstream_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["upstream-interface"], "SystemDhcp6Server-UpstreamInterface"); ok {
			if err = d.Set("upstream_interface", vv); err != nil {
				return fmt.Errorf("Error reading upstream_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcp6ServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDhcp6ServerDelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsSearchList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcp6ServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandSystemDhcp6ServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcp6ServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandSystemDhcp6ServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcp6ServerIpRangeVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcp6ServerIpRangeVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcp6ServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandSystemDhcp6ServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcp6ServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandSystemDhcp6ServerOptionsIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemDhcp6ServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemDhcp6ServerOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandSystemDhcp6ServerOptionsVciMatch(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandSystemDhcp6ServerOptionsVciString(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDhcp6ServerOption1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-prefix"], _ = expandSystemDhcp6ServerPrefixRangeEndPrefix(d, i["end_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemDhcp6ServerPrefixRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-length"], _ = expandSystemDhcp6ServerPrefixRangePrefixLength(d, i["prefix_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-prefix"], _ = expandSystemDhcp6ServerPrefixRangeStartPrefix(d, i["start_prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerPrefixRangeEndPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangePrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeStartPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerRapidCommit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemDhcp6Server(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("delegated_prefix_iaid"); ok || d.HasChange("delegated_prefix_iaid") {
		t, err := expandSystemDhcp6ServerDelegatedPrefixIaid(d, v, "delegated_prefix_iaid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delegated-prefix-iaid"] = t
		}
	}

	if v, ok := d.GetOk("dns_search_list"); ok || d.HasChange("dns_search_list") {
		t, err := expandSystemDhcp6ServerDnsSearchList(d, v, "dns_search_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-search-list"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandSystemDhcp6ServerDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandSystemDhcp6ServerDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandSystemDhcp6ServerDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandSystemDhcp6ServerDnsServer4(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandSystemDhcp6ServerDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandSystemDhcp6ServerDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemDhcp6ServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemDhcp6ServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandSystemDhcp6ServerIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemDhcp6ServerIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandSystemDhcp6ServerLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSystemDhcp6ServerOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("option1"); ok || d.HasChange("option1") {
		t, err := expandSystemDhcp6ServerOption1(d, v, "option1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option1"] = t
		}
	}

	if v, ok := d.GetOk("option2"); ok || d.HasChange("option2") {
		t, err := expandSystemDhcp6ServerOption2(d, v, "option2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option2"] = t
		}
	}

	if v, ok := d.GetOk("option3"); ok || d.HasChange("option3") {
		t, err := expandSystemDhcp6ServerOption3(d, v, "option3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option3"] = t
		}
	}

	if v, ok := d.GetOk("prefix_mode"); ok || d.HasChange("prefix_mode") {
		t, err := expandSystemDhcp6ServerPrefixMode(d, v, "prefix_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-mode"] = t
		}
	}

	if v, ok := d.GetOk("prefix_range"); ok || d.HasChange("prefix_range") {
		t, err := expandSystemDhcp6ServerPrefixRange(d, v, "prefix_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-range"] = t
		}
	}

	if v, ok := d.GetOk("rapid_commit"); ok || d.HasChange("rapid_commit") {
		t, err := expandSystemDhcp6ServerRapidCommit(d, v, "rapid_commit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rapid-commit"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemDhcp6ServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandSystemDhcp6ServerSubnet(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok || d.HasChange("upstream_interface") {
		t, err := expandSystemDhcp6ServerUpstreamInterface(d, v, "upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface"] = t
		}
	}

	return &obj, nil
}
