// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 of interface.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Update,
		Read:   resourceSystemInterfaceIpv6Read,
		Update: resourceSystemInterfaceIpv6Update,
		Delete: resourceSystemInterfaceIpv6Delete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autoconf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_conn6_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_options": &schema.Schema{
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
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dhcp6_client_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp6_iapd_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix_hint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix_hint_plt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix_hint_vlt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dhcp6_information_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_delegation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_prefix_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_prefix_hint_plt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp6_prefix_hint_vlt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dhcp6_relay_interface_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp6_relay_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_source_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp6_send_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_adv_rio": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip6_default_life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_delegated_prefix_iaid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_delegated_prefix_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autonomous_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"delegated_prefix_iaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"onlink_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rdnss": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rdnss_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"upstream_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip6_dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_dnssl_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dnssl_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_extra_addr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_hop_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_link_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_manage_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_max_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_min_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_other_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_prefix_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autonomous_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dnssl": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"onlink_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"preferred_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rdnss": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"valid_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ip6_prefix_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_rdnss_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rdnss": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rdnss_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ip6_reachable_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_retrans_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_route_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_life_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"route_pref": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_route_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_send_adv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_upstream_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nd_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nd_cga_modifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nd_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nd_security_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nd_timestamp_delta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nd_timestamp_fuzz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ra_send_mtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unique_autoconf_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrip6_link_local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrrp_virtual_mac6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrrp6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"adv_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ignore_default_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrdst_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrdst6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vrgrp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vrip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemInterfaceIpv6Update(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemInterfaceIpv6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemInterfaceIpv6")

	return resourceSystemInterfaceIpv6Read(d, m)
}

func resourceSystemInterfaceIpv6Delete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemInterfaceIpv6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Autoconf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6CliConn6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6ClientOptions2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemInterfaceIpv6ClientOptionsCode2edl(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-ClientOptions-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemInterfaceIpv6ClientOptionsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-ClientOptions-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenSystemInterfaceIpv6ClientOptionsIp62edl(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-ClientOptions-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemInterfaceIpv6ClientOptionsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-ClientOptions-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemInterfaceIpv6ClientOptionsValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-ClientOptions-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6ClientOptionsCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6ClientOptionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6ClientOptionsIp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6ClientOptionsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6ClientOptionsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6ClientOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Dhcp6IapdList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iaid"
		if _, ok := i["iaid"]; ok {
			v := flattenSystemInterfaceIpv6Dhcp6IapdListIaid2edl(i["iaid"], d, pre_append)
			tmp["iaid"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Dhcp6IapdList-Iaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint"
		if _, ok := i["prefix-hint"]; ok {
			v := flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint2edl(i["prefix-hint"], d, pre_append)
			tmp["prefix_hint"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Dhcp6IapdList-PrefixHint")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_plt"
		if _, ok := i["prefix-hint-plt"]; ok {
			v := flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt2edl(i["prefix-hint-plt"], d, pre_append)
			tmp["prefix_hint_plt"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Dhcp6IapdList-PrefixHintPlt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_vlt"
		if _, ok := i["prefix-hint-vlt"]; ok {
			v := flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt2edl(i["prefix-hint-vlt"], d, pre_append)
			tmp["prefix_hint_vlt"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Dhcp6IapdList-PrefixHintVlt")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Dhcp6IapdListIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6InformationRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixDelegation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHint2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHintPlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHintVlt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayInterfaceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Dhcp6RelayService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelaySourceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelaySourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Icmp6SendRedirect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6InterfaceIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6AdvRio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Allowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6DefaultLife2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(i["delegated-prefix-iaid"], d, pre_append)
			tmp["delegated_prefix_iaid"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-DelegatedPrefixIaid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(i["prefix-id"], d, pre_append)
			tmp["prefix_id"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-PrefixId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(i["rdnss-service"], d, pre_append)
			tmp["rdnss_service"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-RdnssService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(i["upstream-interface"], d, pre_append)
			tmp["upstream_interface"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DelegatedPrefixList-UpstreamInterface")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6DnsServerOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DnsslList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl_life_time"
		if _, ok := i["dnssl-life-time"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime2edl(i["dnssl-life-time"], d, pre_append)
			tmp["dnssl_life_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DnsslList-DnsslLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenSystemInterfaceIpv6Ip6DnsslListDomain2edl(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6DnsslList-Domain")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DnsslListDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ExtraAddr2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemInterfaceIpv6Ip6ExtraAddrPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6ExtraAddr-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6ExtraAddrPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6HopLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6LinkMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ManageFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6MaxInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6MinInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6OtherFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(i["autonomous-flag"], d, pre_append)
			tmp["autonomous_flag"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-AutonomousFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListDnssl2edl(i["dnssl"], d, pre_append)
			tmp["dnssl"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-Dnssl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(i["onlink-flag"], d, pre_append)
			tmp["onlink_flag"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-OnlinkFlag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(i["preferred-life-time"], d, pre_append)
			tmp["preferred_life_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-PreferredLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListPrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			v := flattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime2edl(i["valid-life-time"], d, pre_append)
			tmp["valid_life_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6PrefixList-ValidLifeTime")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListDnssl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RdnssList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			v := flattenSystemInterfaceIpv6Ip6RdnssListRdnss2edl(i["rdnss"], d, pre_append)
			tmp["rdnss"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6RdnssList-Rdnss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_life_time"
		if _, ok := i["rdnss-life-time"]; ok {
			v := flattenSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime2edl(i["rdnss-life-time"], d, pre_append)
			tmp["rdnss_life_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6RdnssList-RdnssLifeTime")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6RdnssListRdnss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ReachableTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RetransTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RouteList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {
			v := flattenSystemInterfaceIpv6Ip6RouteListRoute2edl(i["route"], d, pre_append)
			tmp["route"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6RouteList-Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_life_time"
		if _, ok := i["route-life-time"]; ok {
			v := flattenSystemInterfaceIpv6Ip6RouteListRouteLifeTime2edl(i["route-life-time"], d, pre_append)
			tmp["route_life_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6RouteList-RouteLifeTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_pref"
		if _, ok := i["route-pref"]; ok {
			v := flattenSystemInterfaceIpv6Ip6RouteListRoutePref2edl(i["route-pref"], d, pre_append)
			tmp["route_pref"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Ip6RouteList-RoutePref")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6RouteListRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RouteListRouteLifeTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RouteListRoutePref2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RoutePref2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6SendAdv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Subnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6UpstreamInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6NdCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6NdCgaModifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdSecurityLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdTimestampDelta2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdTimestampFuzz2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6RaSendMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6UniqueAutoconfAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrip6LinkLocal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6VrrpVirtualMac62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp62edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6AcceptMode2edl(i["accept-mode"], d, pre_append)
			tmp["accept_mode"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-AcceptMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6AdvInterval2edl(i["adv-interval"], d, pre_append)
			tmp["adv_interval"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-AdvInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(i["ignore-default-route"], d, pre_append)
			tmp["ignore_default_route"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-IgnoreDefaultRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Preempt2edl(i["preempt"], d, pre_append)
			tmp["preempt"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Preempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Priority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6StartTime2edl(i["start-time"], d, pre_append)
			tmp["start_time"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-StartTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Status2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6VrdstPriority2edl(i["vrdst-priority"], d, pre_append)
			tmp["vrdst_priority"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-VrdstPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Vrdst62edl(i["vrdst6"], d, pre_append)
			tmp["vrdst6"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Vrdst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Vrgrp2edl(i["vrgrp"], d, pre_append)
			tmp["vrgrp"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Vrgrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Vrid2edl(i["vrid"], d, pre_append)
			tmp["vrid"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Vrid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			v := flattenSystemInterfaceIpv6Vrrp6Vrip62edl(i["vrip6"], d, pre_append)
			tmp["vrip6"] = fortiAPISubPartPatch(v, "SystemInterfaceIpv6-Vrrp6-Vrip6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Vrrp6AcceptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6AdvInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Preempt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Priority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6StartTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6VrdstPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrdst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Vrrp6Vrgrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrip62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("autoconf", flattenSystemInterfaceIpv6Autoconf2edl(o["autoconf"], d, "autoconf")); err != nil {
		if vv, ok := fortiAPIPatch(o["autoconf"], "SystemInterfaceIpv6-Autoconf"); ok {
			if err = d.Set("autoconf", vv); err != nil {
				return fmt.Errorf("Error reading autoconf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading autoconf: %v", err)
		}
	}

	if err = d.Set("cli_conn6_status", flattenSystemInterfaceIpv6CliConn6Status2edl(o["cli-conn6-status"], d, "cli_conn6_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-conn6-status"], "SystemInterfaceIpv6-CliConn6Status"); ok {
			if err = d.Set("cli_conn6_status", vv); err != nil {
				return fmt.Errorf("Error reading cli_conn6_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_conn6_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("client_options", flattenSystemInterfaceIpv6ClientOptions2edl(o["client-options"], d, "client_options")); err != nil {
			if vv, ok := fortiAPIPatch(o["client-options"], "SystemInterfaceIpv6-ClientOptions"); ok {
				if err = d.Set("client_options", vv); err != nil {
					return fmt.Errorf("Error reading client_options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading client_options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("client_options"); ok {
			if err = d.Set("client_options", flattenSystemInterfaceIpv6ClientOptions2edl(o["client-options"], d, "client_options")); err != nil {
				if vv, ok := fortiAPIPatch(o["client-options"], "SystemInterfaceIpv6-ClientOptions"); ok {
					if err = d.Set("client_options", vv); err != nil {
						return fmt.Errorf("Error reading client_options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading client_options: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp6_client_options", flattenSystemInterfaceIpv6Dhcp6ClientOptions2edl(o["dhcp6-client-options"], d, "dhcp6_client_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-client-options"], "SystemInterfaceIpv6-Dhcp6ClientOptions"); ok {
			if err = d.Set("dhcp6_client_options", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_client_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_client_options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp6_iapd_list", flattenSystemInterfaceIpv6Dhcp6IapdList2edl(o["dhcp6-iapd-list"], d, "dhcp6_iapd_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp6-iapd-list"], "SystemInterfaceIpv6-Dhcp6IapdList"); ok {
				if err = d.Set("dhcp6_iapd_list", vv); err != nil {
					return fmt.Errorf("Error reading dhcp6_iapd_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp6_iapd_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp6_iapd_list"); ok {
			if err = d.Set("dhcp6_iapd_list", flattenSystemInterfaceIpv6Dhcp6IapdList2edl(o["dhcp6-iapd-list"], d, "dhcp6_iapd_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp6-iapd-list"], "SystemInterfaceIpv6-Dhcp6IapdList"); ok {
					if err = d.Set("dhcp6_iapd_list", vv); err != nil {
						return fmt.Errorf("Error reading dhcp6_iapd_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dhcp6_iapd_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp6_information_request", flattenSystemInterfaceIpv6Dhcp6InformationRequest2edl(o["dhcp6-information-request"], d, "dhcp6_information_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-information-request"], "SystemInterfaceIpv6-Dhcp6InformationRequest"); ok {
			if err = d.Set("dhcp6_information_request", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_information_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_information_request: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_delegation", flattenSystemInterfaceIpv6Dhcp6PrefixDelegation2edl(o["dhcp6-prefix-delegation"], d, "dhcp6_prefix_delegation")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-delegation"], "SystemInterfaceIpv6-Dhcp6PrefixDelegation"); ok {
			if err = d.Set("dhcp6_prefix_delegation", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_delegation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_delegation: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint", flattenSystemInterfaceIpv6Dhcp6PrefixHint2edl(o["dhcp6-prefix-hint"], d, "dhcp6_prefix_hint")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint"], "SystemInterfaceIpv6-Dhcp6PrefixHint"); ok {
			if err = d.Set("dhcp6_prefix_hint", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint_plt", flattenSystemInterfaceIpv6Dhcp6PrefixHintPlt2edl(o["dhcp6-prefix-hint-plt"], d, "dhcp6_prefix_hint_plt")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint-plt"], "SystemInterfaceIpv6-Dhcp6PrefixHintPlt"); ok {
			if err = d.Set("dhcp6_prefix_hint_plt", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint_plt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint_plt: %v", err)
		}
	}

	if err = d.Set("dhcp6_prefix_hint_vlt", flattenSystemInterfaceIpv6Dhcp6PrefixHintVlt2edl(o["dhcp6-prefix-hint-vlt"], d, "dhcp6_prefix_hint_vlt")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-prefix-hint-vlt"], "SystemInterfaceIpv6-Dhcp6PrefixHintVlt"); ok {
			if err = d.Set("dhcp6_prefix_hint_vlt", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_prefix_hint_vlt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_prefix_hint_vlt: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_interface_id", flattenSystemInterfaceIpv6Dhcp6RelayInterfaceId2edl(o["dhcp6-relay-interface-id"], d, "dhcp6_relay_interface_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-interface-id"], "SystemInterfaceIpv6-Dhcp6RelayInterfaceId"); ok {
			if err = d.Set("dhcp6_relay_interface_id", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_interface_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_interface_id: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_ip", flattenSystemInterfaceIpv6Dhcp6RelayIp2edl(o["dhcp6-relay-ip"], d, "dhcp6_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-ip"], "SystemInterfaceIpv6-Dhcp6RelayIp"); ok {
			if err = d.Set("dhcp6_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_service", flattenSystemInterfaceIpv6Dhcp6RelayService2edl(o["dhcp6-relay-service"], d, "dhcp6_relay_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-service"], "SystemInterfaceIpv6-Dhcp6RelayService"); ok {
			if err = d.Set("dhcp6_relay_service", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_service: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_source_interface", flattenSystemInterfaceIpv6Dhcp6RelaySourceInterface2edl(o["dhcp6-relay-source-interface"], d, "dhcp6_relay_source_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-source-interface"], "SystemInterfaceIpv6-Dhcp6RelaySourceInterface"); ok {
			if err = d.Set("dhcp6_relay_source_interface", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_source_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_source_interface: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_source_ip", flattenSystemInterfaceIpv6Dhcp6RelaySourceIp2edl(o["dhcp6-relay-source-ip"], d, "dhcp6_relay_source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-source-ip"], "SystemInterfaceIpv6-Dhcp6RelaySourceIp"); ok {
			if err = d.Set("dhcp6_relay_source_ip", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_source_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_relay_type", flattenSystemInterfaceIpv6Dhcp6RelayType2edl(o["dhcp6-relay-type"], d, "dhcp6_relay_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-relay-type"], "SystemInterfaceIpv6-Dhcp6RelayType"); ok {
			if err = d.Set("dhcp6_relay_type", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_relay_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_relay_type: %v", err)
		}
	}

	if err = d.Set("icmp6_send_redirect", flattenSystemInterfaceIpv6Icmp6SendRedirect2edl(o["icmp6-send-redirect"], d, "icmp6_send_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp6-send-redirect"], "SystemInterfaceIpv6-Icmp6SendRedirect"); ok {
			if err = d.Set("icmp6_send_redirect", vv); err != nil {
				return fmt.Errorf("Error reading icmp6_send_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp6_send_redirect: %v", err)
		}
	}

	if err = d.Set("interface_identifier", flattenSystemInterfaceIpv6InterfaceIdentifier2edl(o["interface-identifier"], d, "interface_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-identifier"], "SystemInterfaceIpv6-InterfaceIdentifier"); ok {
			if err = d.Set("interface_identifier", vv); err != nil {
				return fmt.Errorf("Error reading interface_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_identifier: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenSystemInterfaceIpv6Ip6Address2edl(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "SystemInterfaceIpv6-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("ip6_adv_rio", flattenSystemInterfaceIpv6Ip6AdvRio2edl(o["ip6-adv-rio"], d, "ip6_adv_rio")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-adv-rio"], "SystemInterfaceIpv6-Ip6AdvRio"); ok {
			if err = d.Set("ip6_adv_rio", vv); err != nil {
				return fmt.Errorf("Error reading ip6_adv_rio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_adv_rio: %v", err)
		}
	}

	if err = d.Set("ip6_allowaccess", flattenSystemInterfaceIpv6Ip6Allowaccess2edl(o["ip6-allowaccess"], d, "ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-allowaccess"], "SystemInterfaceIpv6-Ip6Allowaccess"); ok {
			if err = d.Set("ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("ip6_default_life", flattenSystemInterfaceIpv6Ip6DefaultLife2edl(o["ip6-default-life"], d, "ip6_default_life")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-default-life"], "SystemInterfaceIpv6-Ip6DefaultLife"); ok {
			if err = d.Set("ip6_default_life", vv); err != nil {
				return fmt.Errorf("Error reading ip6_default_life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_default_life: %v", err)
		}
	}

	if err = d.Set("ip6_delegated_prefix_iaid", flattenSystemInterfaceIpv6Ip6DelegatedPrefixIaid2edl(o["ip6-delegated-prefix-iaid"], d, "ip6_delegated_prefix_iaid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-iaid"], "SystemInterfaceIpv6-Ip6DelegatedPrefixIaid"); ok {
			if err = d.Set("ip6_delegated_prefix_iaid", vv); err != nil {
				return fmt.Errorf("Error reading ip6_delegated_prefix_iaid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_delegated_prefix_iaid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_delegated_prefix_list", flattenSystemInterfaceIpv6Ip6DelegatedPrefixList2edl(o["ip6-delegated-prefix-list"], d, "ip6_delegated_prefix_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-list"], "SystemInterfaceIpv6-Ip6DelegatedPrefixList"); ok {
				if err = d.Set("ip6_delegated_prefix_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_delegated_prefix_list"); ok {
			if err = d.Set("ip6_delegated_prefix_list", flattenSystemInterfaceIpv6Ip6DelegatedPrefixList2edl(o["ip6-delegated-prefix-list"], d, "ip6_delegated_prefix_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-delegated-prefix-list"], "SystemInterfaceIpv6-Ip6DelegatedPrefixList"); ok {
					if err = d.Set("ip6_delegated_prefix_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_delegated_prefix_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_dns_server_override", flattenSystemInterfaceIpv6Ip6DnsServerOverride2edl(o["ip6-dns-server-override"], d, "ip6_dns_server_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-dns-server-override"], "SystemInterfaceIpv6-Ip6DnsServerOverride"); ok {
			if err = d.Set("ip6_dns_server_override", vv); err != nil {
				return fmt.Errorf("Error reading ip6_dns_server_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_dns_server_override: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_dnssl_list", flattenSystemInterfaceIpv6Ip6DnsslList2edl(o["ip6-dnssl-list"], d, "ip6_dnssl_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-dnssl-list"], "SystemInterfaceIpv6-Ip6DnsslList"); ok {
				if err = d.Set("ip6_dnssl_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_dnssl_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_dnssl_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_dnssl_list"); ok {
			if err = d.Set("ip6_dnssl_list", flattenSystemInterfaceIpv6Ip6DnsslList2edl(o["ip6-dnssl-list"], d, "ip6_dnssl_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-dnssl-list"], "SystemInterfaceIpv6-Ip6DnsslList"); ok {
					if err = d.Set("ip6_dnssl_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_dnssl_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_dnssl_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_extra_addr", flattenSystemInterfaceIpv6Ip6ExtraAddr2edl(o["ip6-extra-addr"], d, "ip6_extra_addr")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-extra-addr"], "SystemInterfaceIpv6-Ip6ExtraAddr"); ok {
				if err = d.Set("ip6_extra_addr", vv); err != nil {
					return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_extra_addr"); ok {
			if err = d.Set("ip6_extra_addr", flattenSystemInterfaceIpv6Ip6ExtraAddr2edl(o["ip6-extra-addr"], d, "ip6_extra_addr")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-extra-addr"], "SystemInterfaceIpv6-Ip6ExtraAddr"); ok {
					if err = d.Set("ip6_extra_addr", vv); err != nil {
						return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_extra_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_hop_limit", flattenSystemInterfaceIpv6Ip6HopLimit2edl(o["ip6-hop-limit"], d, "ip6_hop_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-hop-limit"], "SystemInterfaceIpv6-Ip6HopLimit"); ok {
			if err = d.Set("ip6_hop_limit", vv); err != nil {
				return fmt.Errorf("Error reading ip6_hop_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_hop_limit: %v", err)
		}
	}

	if err = d.Set("ip6_link_mtu", flattenSystemInterfaceIpv6Ip6LinkMtu2edl(o["ip6-link-mtu"], d, "ip6_link_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-link-mtu"], "SystemInterfaceIpv6-Ip6LinkMtu"); ok {
			if err = d.Set("ip6_link_mtu", vv); err != nil {
				return fmt.Errorf("Error reading ip6_link_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_link_mtu: %v", err)
		}
	}

	if err = d.Set("ip6_manage_flag", flattenSystemInterfaceIpv6Ip6ManageFlag2edl(o["ip6-manage-flag"], d, "ip6_manage_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-manage-flag"], "SystemInterfaceIpv6-Ip6ManageFlag"); ok {
			if err = d.Set("ip6_manage_flag", vv); err != nil {
				return fmt.Errorf("Error reading ip6_manage_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_manage_flag: %v", err)
		}
	}

	if err = d.Set("ip6_max_interval", flattenSystemInterfaceIpv6Ip6MaxInterval2edl(o["ip6-max-interval"], d, "ip6_max_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-max-interval"], "SystemInterfaceIpv6-Ip6MaxInterval"); ok {
			if err = d.Set("ip6_max_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip6_max_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_max_interval: %v", err)
		}
	}

	if err = d.Set("ip6_min_interval", flattenSystemInterfaceIpv6Ip6MinInterval2edl(o["ip6-min-interval"], d, "ip6_min_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-min-interval"], "SystemInterfaceIpv6-Ip6MinInterval"); ok {
			if err = d.Set("ip6_min_interval", vv); err != nil {
				return fmt.Errorf("Error reading ip6_min_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_min_interval: %v", err)
		}
	}

	if err = d.Set("ip6_mode", flattenSystemInterfaceIpv6Ip6Mode2edl(o["ip6-mode"], d, "ip6_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-mode"], "SystemInterfaceIpv6-Ip6Mode"); ok {
			if err = d.Set("ip6_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip6_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_mode: %v", err)
		}
	}

	if err = d.Set("ip6_other_flag", flattenSystemInterfaceIpv6Ip6OtherFlag2edl(o["ip6-other-flag"], d, "ip6_other_flag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-other-flag"], "SystemInterfaceIpv6-Ip6OtherFlag"); ok {
			if err = d.Set("ip6_other_flag", vv); err != nil {
				return fmt.Errorf("Error reading ip6_other_flag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_other_flag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_prefix_list", flattenSystemInterfaceIpv6Ip6PrefixList2edl(o["ip6-prefix-list"], d, "ip6_prefix_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-prefix-list"], "SystemInterfaceIpv6-Ip6PrefixList"); ok {
				if err = d.Set("ip6_prefix_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_prefix_list"); ok {
			if err = d.Set("ip6_prefix_list", flattenSystemInterfaceIpv6Ip6PrefixList2edl(o["ip6-prefix-list"], d, "ip6_prefix_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-prefix-list"], "SystemInterfaceIpv6-Ip6PrefixList"); ok {
					if err = d.Set("ip6_prefix_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_prefix_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_prefix_mode", flattenSystemInterfaceIpv6Ip6PrefixMode2edl(o["ip6-prefix-mode"], d, "ip6_prefix_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-prefix-mode"], "SystemInterfaceIpv6-Ip6PrefixMode"); ok {
			if err = d.Set("ip6_prefix_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip6_prefix_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_prefix_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_rdnss_list", flattenSystemInterfaceIpv6Ip6RdnssList2edl(o["ip6-rdnss-list"], d, "ip6_rdnss_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-rdnss-list"], "SystemInterfaceIpv6-Ip6RdnssList"); ok {
				if err = d.Set("ip6_rdnss_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_rdnss_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_rdnss_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_rdnss_list"); ok {
			if err = d.Set("ip6_rdnss_list", flattenSystemInterfaceIpv6Ip6RdnssList2edl(o["ip6-rdnss-list"], d, "ip6_rdnss_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-rdnss-list"], "SystemInterfaceIpv6-Ip6RdnssList"); ok {
					if err = d.Set("ip6_rdnss_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_rdnss_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_rdnss_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_reachable_time", flattenSystemInterfaceIpv6Ip6ReachableTime2edl(o["ip6-reachable-time"], d, "ip6_reachable_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-reachable-time"], "SystemInterfaceIpv6-Ip6ReachableTime"); ok {
			if err = d.Set("ip6_reachable_time", vv); err != nil {
				return fmt.Errorf("Error reading ip6_reachable_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_reachable_time: %v", err)
		}
	}

	if err = d.Set("ip6_retrans_time", flattenSystemInterfaceIpv6Ip6RetransTime2edl(o["ip6-retrans-time"], d, "ip6_retrans_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-retrans-time"], "SystemInterfaceIpv6-Ip6RetransTime"); ok {
			if err = d.Set("ip6_retrans_time", vv); err != nil {
				return fmt.Errorf("Error reading ip6_retrans_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_retrans_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_route_list", flattenSystemInterfaceIpv6Ip6RouteList2edl(o["ip6-route-list"], d, "ip6_route_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-route-list"], "SystemInterfaceIpv6-Ip6RouteList"); ok {
				if err = d.Set("ip6_route_list", vv); err != nil {
					return fmt.Errorf("Error reading ip6_route_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_route_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_route_list"); ok {
			if err = d.Set("ip6_route_list", flattenSystemInterfaceIpv6Ip6RouteList2edl(o["ip6-route-list"], d, "ip6_route_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-route-list"], "SystemInterfaceIpv6-Ip6RouteList"); ok {
					if err = d.Set("ip6_route_list", vv); err != nil {
						return fmt.Errorf("Error reading ip6_route_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_route_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_route_pref", flattenSystemInterfaceIpv6Ip6RoutePref2edl(o["ip6-route-pref"], d, "ip6_route_pref")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-route-pref"], "SystemInterfaceIpv6-Ip6RoutePref"); ok {
			if err = d.Set("ip6_route_pref", vv); err != nil {
				return fmt.Errorf("Error reading ip6_route_pref: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_route_pref: %v", err)
		}
	}

	if err = d.Set("ip6_send_adv", flattenSystemInterfaceIpv6Ip6SendAdv2edl(o["ip6-send-adv"], d, "ip6_send_adv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-send-adv"], "SystemInterfaceIpv6-Ip6SendAdv"); ok {
			if err = d.Set("ip6_send_adv", vv); err != nil {
				return fmt.Errorf("Error reading ip6_send_adv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_send_adv: %v", err)
		}
	}

	if err = d.Set("ip6_subnet", flattenSystemInterfaceIpv6Ip6Subnet2edl(o["ip6-subnet"], d, "ip6_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-subnet"], "SystemInterfaceIpv6-Ip6Subnet"); ok {
			if err = d.Set("ip6_subnet", vv); err != nil {
				return fmt.Errorf("Error reading ip6_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_subnet: %v", err)
		}
	}

	if err = d.Set("ip6_upstream_interface", flattenSystemInterfaceIpv6Ip6UpstreamInterface2edl(o["ip6-upstream-interface"], d, "ip6_upstream_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-upstream-interface"], "SystemInterfaceIpv6-Ip6UpstreamInterface"); ok {
			if err = d.Set("ip6_upstream_interface", vv); err != nil {
				return fmt.Errorf("Error reading ip6_upstream_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_upstream_interface: %v", err)
		}
	}

	if err = d.Set("nd_cert", flattenSystemInterfaceIpv6NdCert2edl(o["nd-cert"], d, "nd_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-cert"], "SystemInterfaceIpv6-NdCert"); ok {
			if err = d.Set("nd_cert", vv); err != nil {
				return fmt.Errorf("Error reading nd_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_cert: %v", err)
		}
	}

	if err = d.Set("nd_cga_modifier", flattenSystemInterfaceIpv6NdCgaModifier2edl(o["nd-cga-modifier"], d, "nd_cga_modifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-cga-modifier"], "SystemInterfaceIpv6-NdCgaModifier"); ok {
			if err = d.Set("nd_cga_modifier", vv); err != nil {
				return fmt.Errorf("Error reading nd_cga_modifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_cga_modifier: %v", err)
		}
	}

	if err = d.Set("nd_mode", flattenSystemInterfaceIpv6NdMode2edl(o["nd-mode"], d, "nd_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-mode"], "SystemInterfaceIpv6-NdMode"); ok {
			if err = d.Set("nd_mode", vv); err != nil {
				return fmt.Errorf("Error reading nd_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_mode: %v", err)
		}
	}

	if err = d.Set("nd_security_level", flattenSystemInterfaceIpv6NdSecurityLevel2edl(o["nd-security-level"], d, "nd_security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-security-level"], "SystemInterfaceIpv6-NdSecurityLevel"); ok {
			if err = d.Set("nd_security_level", vv); err != nil {
				return fmt.Errorf("Error reading nd_security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_security_level: %v", err)
		}
	}

	if err = d.Set("nd_timestamp_delta", flattenSystemInterfaceIpv6NdTimestampDelta2edl(o["nd-timestamp-delta"], d, "nd_timestamp_delta")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-timestamp-delta"], "SystemInterfaceIpv6-NdTimestampDelta"); ok {
			if err = d.Set("nd_timestamp_delta", vv); err != nil {
				return fmt.Errorf("Error reading nd_timestamp_delta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_timestamp_delta: %v", err)
		}
	}

	if err = d.Set("nd_timestamp_fuzz", flattenSystemInterfaceIpv6NdTimestampFuzz2edl(o["nd-timestamp-fuzz"], d, "nd_timestamp_fuzz")); err != nil {
		if vv, ok := fortiAPIPatch(o["nd-timestamp-fuzz"], "SystemInterfaceIpv6-NdTimestampFuzz"); ok {
			if err = d.Set("nd_timestamp_fuzz", vv); err != nil {
				return fmt.Errorf("Error reading nd_timestamp_fuzz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nd_timestamp_fuzz: %v", err)
		}
	}

	if err = d.Set("ra_send_mtu", flattenSystemInterfaceIpv6RaSendMtu2edl(o["ra-send-mtu"], d, "ra_send_mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ra-send-mtu"], "SystemInterfaceIpv6-RaSendMtu"); ok {
			if err = d.Set("ra_send_mtu", vv); err != nil {
				return fmt.Errorf("Error reading ra_send_mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ra_send_mtu: %v", err)
		}
	}

	if err = d.Set("unique_autoconf_addr", flattenSystemInterfaceIpv6UniqueAutoconfAddr2edl(o["unique-autoconf-addr"], d, "unique_autoconf_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["unique-autoconf-addr"], "SystemInterfaceIpv6-UniqueAutoconfAddr"); ok {
			if err = d.Set("unique_autoconf_addr", vv); err != nil {
				return fmt.Errorf("Error reading unique_autoconf_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unique_autoconf_addr: %v", err)
		}
	}

	if err = d.Set("vrip6_link_local", flattenSystemInterfaceIpv6Vrip6LinkLocal2edl(o["vrip6_link_local"], d, "vrip6_link_local")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip6_link_local"], "SystemInterfaceIpv6-Vrip6LinkLocal"); ok {
			if err = d.Set("vrip6_link_local", vv); err != nil {
				return fmt.Errorf("Error reading vrip6_link_local: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip6_link_local: %v", err)
		}
	}

	if err = d.Set("vrrp_virtual_mac6", flattenSystemInterfaceIpv6VrrpVirtualMac62edl(o["vrrp-virtual-mac6"], d, "vrrp_virtual_mac6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrrp-virtual-mac6"], "SystemInterfaceIpv6-VrrpVirtualMac6"); ok {
			if err = d.Set("vrrp_virtual_mac6", vv); err != nil {
				return fmt.Errorf("Error reading vrrp_virtual_mac6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrrp_virtual_mac6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrrp6", flattenSystemInterfaceIpv6Vrrp62edl(o["vrrp6"], d, "vrrp6")); err != nil {
			if vv, ok := fortiAPIPatch(o["vrrp6"], "SystemInterfaceIpv6-Vrrp6"); ok {
				if err = d.Set("vrrp6", vv); err != nil {
					return fmt.Errorf("Error reading vrrp6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vrrp6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrrp6"); ok {
			if err = d.Set("vrrp6", flattenSystemInterfaceIpv6Vrrp62edl(o["vrrp6"], d, "vrrp6")); err != nil {
				if vv, ok := fortiAPIPatch(o["vrrp6"], "SystemInterfaceIpv6-Vrrp6"); ok {
					if err = d.Set("vrrp6", vv); err != nil {
						return fmt.Errorf("Error reading vrrp6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vrrp6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Autoconf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6CliConn6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6ClientOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["code"], _ = expandSystemInterfaceIpv6ClientOptionsCode2edl(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemInterfaceIpv6ClientOptionsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandSystemInterfaceIpv6ClientOptionsIp62edl(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemInterfaceIpv6ClientOptionsType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemInterfaceIpv6ClientOptionsValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6ClientOptionsCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6ClientOptionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6ClientOptionsIp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6ClientOptionsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6ClientOptionsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6ClientOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Dhcp6IapdList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["iaid"], _ = expandSystemInterfaceIpv6Dhcp6IapdListIaid2edl(d, i["iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-hint"], _ = expandSystemInterfaceIpv6Dhcp6IapdListPrefixHint2edl(d, i["prefix_hint"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_plt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-hint-plt"], _ = expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt2edl(d, i["prefix_hint_plt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_vlt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-hint-vlt"], _ = expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt2edl(d, i["prefix_hint_vlt"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHint2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6InformationRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixDelegation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHint2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHintPlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHintVlt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayInterfaceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Dhcp6RelayService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelaySourceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelaySourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Icmp6SendRedirect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6InterfaceIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6AdvRio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Allowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6DefaultLife2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delegated-prefix-iaid"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d, i["delegated_prefix_iaid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-id"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-service"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["upstream-interface"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d, i["upstream_interface"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6DnsServerOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DnsslList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl-life-time"], _ = expandSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime2edl(d, i["dnssl_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandSystemInterfaceIpv6Ip6DnsslListDomain2edl(d, i["domain"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6DnsslListDnsslLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DnsslListDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ExtraAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemInterfaceIpv6Ip6ExtraAddrPrefix2edl(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6ExtraAddrPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6HopLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6LinkMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ManageFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6MaxInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6MinInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6OtherFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["autonomous-flag"], _ = expandSystemInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dnssl"], _ = expandSystemInterfaceIpv6Ip6PrefixListDnssl2edl(d, i["dnssl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["onlink-flag"], _ = expandSystemInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-life-time"], _ = expandSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemInterfaceIpv6Ip6PrefixListPrefix2edl(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandSystemInterfaceIpv6Ip6PrefixListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["valid-life-time"], _ = expandSystemInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d, i["valid_life_time"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListAutonomousFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListDnssl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6PrefixListOnlinkFlag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6PrefixListValidLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RdnssList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss"], _ = expandSystemInterfaceIpv6Ip6RdnssListRdnss2edl(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rdnss-life-time"], _ = expandSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime2edl(d, i["rdnss_life_time"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6RdnssListRdnss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RdnssListRdnssLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ReachableTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RetransTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RouteList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route"], _ = expandSystemInterfaceIpv6Ip6RouteListRoute2edl(d, i["route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_life_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-life-time"], _ = expandSystemInterfaceIpv6Ip6RouteListRouteLifeTime2edl(d, i["route_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_pref"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-pref"], _ = expandSystemInterfaceIpv6Ip6RouteListRoutePref2edl(d, i["route_pref"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6RouteListRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RouteListRouteLifeTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RouteListRoutePref2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RoutePref2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6SendAdv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Subnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6UpstreamInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6NdCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6NdCgaModifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdSecurityLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdTimestampDelta2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdTimestampFuzz2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6RaSendMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6UniqueAutoconfAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrip6LinkLocal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6VrrpVirtualMac62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["accept-mode"], _ = expandSystemInterfaceIpv6Vrrp6AcceptMode2edl(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adv-interval"], _ = expandSystemInterfaceIpv6Vrrp6AdvInterval2edl(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ignore-default-route"], _ = expandSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(d, i["ignore_default_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preempt"], _ = expandSystemInterfaceIpv6Vrrp6Preempt2edl(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSystemInterfaceIpv6Vrrp6Priority2edl(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-time"], _ = expandSystemInterfaceIpv6Vrrp6StartTime2edl(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemInterfaceIpv6Vrrp6Status2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst-priority"], _ = expandSystemInterfaceIpv6Vrrp6VrdstPriority2edl(d, i["vrdst_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrdst6"], _ = expandSystemInterfaceIpv6Vrrp6Vrdst62edl(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrgrp"], _ = expandSystemInterfaceIpv6Vrrp6Vrgrp2edl(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrid"], _ = expandSystemInterfaceIpv6Vrrp6Vrid2edl(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrip6"], _ = expandSystemInterfaceIpv6Vrrp6Vrip62edl(d, i["vrip6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Vrrp6AcceptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6AdvInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Preempt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Priority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6StartTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6VrdstPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrdst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Vrrp6Vrgrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrip62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("autoconf"); ok || d.HasChange("autoconf") {
		t, err := expandSystemInterfaceIpv6Autoconf2edl(d, v, "autoconf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autoconf"] = t
		}
	}

	if v, ok := d.GetOk("cli_conn6_status"); ok || d.HasChange("cli_conn6_status") {
		t, err := expandSystemInterfaceIpv6CliConn6Status2edl(d, v, "cli_conn6_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-conn6-status"] = t
		}
	}

	if v, ok := d.GetOk("client_options"); ok || d.HasChange("client_options") {
		t, err := expandSystemInterfaceIpv6ClientOptions2edl(d, v, "client_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-options"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_client_options"); ok || d.HasChange("dhcp6_client_options") {
		t, err := expandSystemInterfaceIpv6Dhcp6ClientOptions2edl(d, v, "dhcp6_client_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-client-options"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_iapd_list"); ok || d.HasChange("dhcp6_iapd_list") {
		t, err := expandSystemInterfaceIpv6Dhcp6IapdList2edl(d, v, "dhcp6_iapd_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-iapd-list"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_information_request"); ok || d.HasChange("dhcp6_information_request") {
		t, err := expandSystemInterfaceIpv6Dhcp6InformationRequest2edl(d, v, "dhcp6_information_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-information-request"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_delegation"); ok || d.HasChange("dhcp6_prefix_delegation") {
		t, err := expandSystemInterfaceIpv6Dhcp6PrefixDelegation2edl(d, v, "dhcp6_prefix_delegation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-delegation"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint"); ok || d.HasChange("dhcp6_prefix_hint") {
		t, err := expandSystemInterfaceIpv6Dhcp6PrefixHint2edl(d, v, "dhcp6_prefix_hint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint_plt"); ok || d.HasChange("dhcp6_prefix_hint_plt") {
		t, err := expandSystemInterfaceIpv6Dhcp6PrefixHintPlt2edl(d, v, "dhcp6_prefix_hint_plt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint-plt"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_prefix_hint_vlt"); ok || d.HasChange("dhcp6_prefix_hint_vlt") {
		t, err := expandSystemInterfaceIpv6Dhcp6PrefixHintVlt2edl(d, v, "dhcp6_prefix_hint_vlt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-prefix-hint-vlt"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_interface_id"); ok || d.HasChange("dhcp6_relay_interface_id") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelayInterfaceId2edl(d, v, "dhcp6_relay_interface_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-interface-id"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_ip"); ok || d.HasChange("dhcp6_relay_ip") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelayIp2edl(d, v, "dhcp6_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_service"); ok || d.HasChange("dhcp6_relay_service") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelayService2edl(d, v, "dhcp6_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_source_interface"); ok || d.HasChange("dhcp6_relay_source_interface") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelaySourceInterface2edl(d, v, "dhcp6_relay_source_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-source-interface"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_source_ip"); ok || d.HasChange("dhcp6_relay_source_ip") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelaySourceIp2edl(d, v, "dhcp6_relay_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_relay_type"); ok || d.HasChange("dhcp6_relay_type") {
		t, err := expandSystemInterfaceIpv6Dhcp6RelayType2edl(d, v, "dhcp6_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("icmp6_send_redirect"); ok || d.HasChange("icmp6_send_redirect") {
		t, err := expandSystemInterfaceIpv6Icmp6SendRedirect2edl(d, v, "icmp6_send_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp6-send-redirect"] = t
		}
	}

	if v, ok := d.GetOk("interface_identifier"); ok || d.HasChange("interface_identifier") {
		t, err := expandSystemInterfaceIpv6InterfaceIdentifier2edl(d, v, "interface_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-identifier"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandSystemInterfaceIpv6Ip6Address2edl(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6_adv_rio"); ok || d.HasChange("ip6_adv_rio") {
		t, err := expandSystemInterfaceIpv6Ip6AdvRio2edl(d, v, "ip6_adv_rio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-adv-rio"] = t
		}
	}

	if v, ok := d.GetOk("ip6_allowaccess"); ok || d.HasChange("ip6_allowaccess") {
		t, err := expandSystemInterfaceIpv6Ip6Allowaccess2edl(d, v, "ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("ip6_default_life"); ok || d.HasChange("ip6_default_life") {
		t, err := expandSystemInterfaceIpv6Ip6DefaultLife2edl(d, v, "ip6_default_life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-default-life"] = t
		}
	}

	if v, ok := d.GetOk("ip6_delegated_prefix_iaid"); ok || d.HasChange("ip6_delegated_prefix_iaid") {
		t, err := expandSystemInterfaceIpv6Ip6DelegatedPrefixIaid2edl(d, v, "ip6_delegated_prefix_iaid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-delegated-prefix-iaid"] = t
		}
	}

	if v, ok := d.GetOk("ip6_delegated_prefix_list"); ok || d.HasChange("ip6_delegated_prefix_list") {
		t, err := expandSystemInterfaceIpv6Ip6DelegatedPrefixList2edl(d, v, "ip6_delegated_prefix_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-delegated-prefix-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_dns_server_override"); ok || d.HasChange("ip6_dns_server_override") {
		t, err := expandSystemInterfaceIpv6Ip6DnsServerOverride2edl(d, v, "ip6_dns_server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-dns-server-override"] = t
		}
	}

	if v, ok := d.GetOk("ip6_dnssl_list"); ok || d.HasChange("ip6_dnssl_list") {
		t, err := expandSystemInterfaceIpv6Ip6DnsslList2edl(d, v, "ip6_dnssl_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-dnssl-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_extra_addr"); ok || d.HasChange("ip6_extra_addr") {
		t, err := expandSystemInterfaceIpv6Ip6ExtraAddr2edl(d, v, "ip6_extra_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-extra-addr"] = t
		}
	}

	if v, ok := d.GetOk("ip6_hop_limit"); ok || d.HasChange("ip6_hop_limit") {
		t, err := expandSystemInterfaceIpv6Ip6HopLimit2edl(d, v, "ip6_hop_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-hop-limit"] = t
		}
	}

	if v, ok := d.GetOk("ip6_link_mtu"); ok || d.HasChange("ip6_link_mtu") {
		t, err := expandSystemInterfaceIpv6Ip6LinkMtu2edl(d, v, "ip6_link_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-link-mtu"] = t
		}
	}

	if v, ok := d.GetOk("ip6_manage_flag"); ok || d.HasChange("ip6_manage_flag") {
		t, err := expandSystemInterfaceIpv6Ip6ManageFlag2edl(d, v, "ip6_manage_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-manage-flag"] = t
		}
	}

	if v, ok := d.GetOk("ip6_max_interval"); ok || d.HasChange("ip6_max_interval") {
		t, err := expandSystemInterfaceIpv6Ip6MaxInterval2edl(d, v, "ip6_max_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-max-interval"] = t
		}
	}

	if v, ok := d.GetOk("ip6_min_interval"); ok || d.HasChange("ip6_min_interval") {
		t, err := expandSystemInterfaceIpv6Ip6MinInterval2edl(d, v, "ip6_min_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-min-interval"] = t
		}
	}

	if v, ok := d.GetOk("ip6_mode"); ok || d.HasChange("ip6_mode") {
		t, err := expandSystemInterfaceIpv6Ip6Mode2edl(d, v, "ip6_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip6_other_flag"); ok || d.HasChange("ip6_other_flag") {
		t, err := expandSystemInterfaceIpv6Ip6OtherFlag2edl(d, v, "ip6_other_flag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-other-flag"] = t
		}
	}

	if v, ok := d.GetOk("ip6_prefix_list"); ok || d.HasChange("ip6_prefix_list") {
		t, err := expandSystemInterfaceIpv6Ip6PrefixList2edl(d, v, "ip6_prefix_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-prefix-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_prefix_mode"); ok || d.HasChange("ip6_prefix_mode") {
		t, err := expandSystemInterfaceIpv6Ip6PrefixMode2edl(d, v, "ip6_prefix_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-prefix-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip6_rdnss_list"); ok || d.HasChange("ip6_rdnss_list") {
		t, err := expandSystemInterfaceIpv6Ip6RdnssList2edl(d, v, "ip6_rdnss_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-rdnss-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_reachable_time"); ok || d.HasChange("ip6_reachable_time") {
		t, err := expandSystemInterfaceIpv6Ip6ReachableTime2edl(d, v, "ip6_reachable_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-reachable-time"] = t
		}
	}

	if v, ok := d.GetOk("ip6_retrans_time"); ok || d.HasChange("ip6_retrans_time") {
		t, err := expandSystemInterfaceIpv6Ip6RetransTime2edl(d, v, "ip6_retrans_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-retrans-time"] = t
		}
	}

	if v, ok := d.GetOk("ip6_route_list"); ok || d.HasChange("ip6_route_list") {
		t, err := expandSystemInterfaceIpv6Ip6RouteList2edl(d, v, "ip6_route_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-route-list"] = t
		}
	}

	if v, ok := d.GetOk("ip6_route_pref"); ok || d.HasChange("ip6_route_pref") {
		t, err := expandSystemInterfaceIpv6Ip6RoutePref2edl(d, v, "ip6_route_pref")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-route-pref"] = t
		}
	}

	if v, ok := d.GetOk("ip6_send_adv"); ok || d.HasChange("ip6_send_adv") {
		t, err := expandSystemInterfaceIpv6Ip6SendAdv2edl(d, v, "ip6_send_adv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-send-adv"] = t
		}
	}

	if v, ok := d.GetOk("ip6_subnet"); ok || d.HasChange("ip6_subnet") {
		t, err := expandSystemInterfaceIpv6Ip6Subnet2edl(d, v, "ip6_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("ip6_upstream_interface"); ok || d.HasChange("ip6_upstream_interface") {
		t, err := expandSystemInterfaceIpv6Ip6UpstreamInterface2edl(d, v, "ip6_upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-upstream-interface"] = t
		}
	}

	if v, ok := d.GetOk("nd_cert"); ok || d.HasChange("nd_cert") {
		t, err := expandSystemInterfaceIpv6NdCert2edl(d, v, "nd_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-cert"] = t
		}
	}

	if v, ok := d.GetOk("nd_cga_modifier"); ok || d.HasChange("nd_cga_modifier") {
		t, err := expandSystemInterfaceIpv6NdCgaModifier2edl(d, v, "nd_cga_modifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-cga-modifier"] = t
		}
	}

	if v, ok := d.GetOk("nd_mode"); ok || d.HasChange("nd_mode") {
		t, err := expandSystemInterfaceIpv6NdMode2edl(d, v, "nd_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-mode"] = t
		}
	}

	if v, ok := d.GetOk("nd_security_level"); ok || d.HasChange("nd_security_level") {
		t, err := expandSystemInterfaceIpv6NdSecurityLevel2edl(d, v, "nd_security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-security-level"] = t
		}
	}

	if v, ok := d.GetOk("nd_timestamp_delta"); ok || d.HasChange("nd_timestamp_delta") {
		t, err := expandSystemInterfaceIpv6NdTimestampDelta2edl(d, v, "nd_timestamp_delta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-timestamp-delta"] = t
		}
	}

	if v, ok := d.GetOk("nd_timestamp_fuzz"); ok || d.HasChange("nd_timestamp_fuzz") {
		t, err := expandSystemInterfaceIpv6NdTimestampFuzz2edl(d, v, "nd_timestamp_fuzz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nd-timestamp-fuzz"] = t
		}
	}

	if v, ok := d.GetOk("ra_send_mtu"); ok || d.HasChange("ra_send_mtu") {
		t, err := expandSystemInterfaceIpv6RaSendMtu2edl(d, v, "ra_send_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ra-send-mtu"] = t
		}
	}

	if v, ok := d.GetOk("unique_autoconf_addr"); ok || d.HasChange("unique_autoconf_addr") {
		t, err := expandSystemInterfaceIpv6UniqueAutoconfAddr2edl(d, v, "unique_autoconf_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unique-autoconf-addr"] = t
		}
	}

	if v, ok := d.GetOk("vrip6_link_local"); ok || d.HasChange("vrip6_link_local") {
		t, err := expandSystemInterfaceIpv6Vrip6LinkLocal2edl(d, v, "vrip6_link_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip6_link_local"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_virtual_mac6"); ok || d.HasChange("vrrp_virtual_mac6") {
		t, err := expandSystemInterfaceIpv6VrrpVirtualMac62edl(d, v, "vrrp_virtual_mac6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-virtual-mac6"] = t
		}
	}

	if v, ok := d.GetOk("vrrp6"); ok || d.HasChange("vrrp6") {
		t, err := expandSystemInterfaceIpv6Vrrp62edl(d, v, "vrrp6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp6"] = t
		}
	}

	return &obj, nil
}
