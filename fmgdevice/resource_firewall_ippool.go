// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 IP pools.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIppool() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIppoolCreate,
		Read:   resourceFirewallIppoolRead,
		Update: resourceFirewallIppoolUpdate,
		Delete: resourceFirewallIppoolDelete,

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
			"add_nat64_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_intf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"block_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_block_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_client_endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_client_ipv6shift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_client_startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_fixedalloc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_overload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_port_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_port_start": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_spa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_prefix_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"add_nat64_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arp_intf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"associated_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"block_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cgn_block_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cgn_client_endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cgn_client_ipv6shift": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cgn_client_startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cgn_fixedalloc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cgn_overload": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cgn_port_end": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cgn_port_start": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cgn_spa": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_prefix_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"exclude_ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"icmp_session_quota": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"nat64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_blocks_per_user": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pba_interim_log": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pba_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"permit_any_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_per_user": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"privileged_port_use_pba": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"startport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"subnet_broadcast_in_ippool": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tcp_session_quota": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"udp_session_quota": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"utilization_alarm_clear": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"utilization_alarm_raise": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"endport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"exclude_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icmp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_blocks_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pba_interim_log": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pba_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"privileged_port_use_pba": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"startport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_broadcast_in_ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"utilization_alarm_clear": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"utilization_alarm_raise": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallIppoolCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIppool resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallIppool(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallIppool(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallIppool resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallIppool(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallIppool resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallIppoolRead(d, m)
}

func resourceFirewallIppoolUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallIppool(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallIppoolRead(d, m)
}

func resourceFirewallIppoolDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallIppool(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIppool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppoolRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIppool(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallIppool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIppoolAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolArpIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIppoolArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIppoolBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnClientEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnClientIpv6Shift(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnClientStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnFixedalloc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnOverload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnPortEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnPortStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolCgnSpa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolClientPrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallIppoolDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := i["add-nat64-route"]; ok {
			v := flattenFirewallIppoolDynamicMappingAddNat64Route(i["add-nat64-route"], d, pre_append)
			tmp["add_nat64_route"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-AddNat64Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_intf"
		if _, ok := i["arp-intf"]; ok {
			v := flattenFirewallIppoolDynamicMappingArpIntf(i["arp-intf"], d, pre_append)
			tmp["arp_intf"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-ArpIntf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenFirewallIppoolDynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := i["associated-interface"]; ok {
			v := flattenFirewallIppoolDynamicMappingAssociatedInterface(i["associated-interface"], d, pre_append)
			tmp["associated_interface"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-AssociatedInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "block_size"
		if _, ok := i["block-size"]; ok {
			v := flattenFirewallIppoolDynamicMappingBlockSize(i["block-size"], d, pre_append)
			tmp["block_size"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-BlockSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_block_size"
		if _, ok := i["cgn-block-size"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnBlockSize(i["cgn-block-size"], d, pre_append)
			tmp["cgn_block_size"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnBlockSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_endip"
		if _, ok := i["cgn-client-endip"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnClientEndip(i["cgn-client-endip"], d, pre_append)
			tmp["cgn_client_endip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnClientEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_ipv6shift"
		if _, ok := i["cgn-client-ipv6shift"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnClientIpv6Shift(i["cgn-client-ipv6shift"], d, pre_append)
			tmp["cgn_client_ipv6shift"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnClientIpv6Shift")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_startip"
		if _, ok := i["cgn-client-startip"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnClientStartip(i["cgn-client-startip"], d, pre_append)
			tmp["cgn_client_startip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnClientStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_fixedalloc"
		if _, ok := i["cgn-fixedalloc"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnFixedalloc(i["cgn-fixedalloc"], d, pre_append)
			tmp["cgn_fixedalloc"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnFixedalloc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_overload"
		if _, ok := i["cgn-overload"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnOverload(i["cgn-overload"], d, pre_append)
			tmp["cgn_overload"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnOverload")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_end"
		if _, ok := i["cgn-port-end"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnPortEnd(i["cgn-port-end"], d, pre_append)
			tmp["cgn_port_end"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnPortEnd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_start"
		if _, ok := i["cgn-port-start"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnPortStart(i["cgn-port-start"], d, pre_append)
			tmp["cgn_port_start"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnPortStart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_spa"
		if _, ok := i["cgn-spa"]; ok {
			v := flattenFirewallIppoolDynamicMappingCgnSpa(i["cgn-spa"], d, pre_append)
			tmp["cgn_spa"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-CgnSpa")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_prefix_length"
		if _, ok := i["client-prefix-length"]; ok {
			v := flattenFirewallIppoolDynamicMappingClientPrefixLength(i["client-prefix-length"], d, pre_append)
			tmp["client_prefix_length"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-ClientPrefixLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenFirewallIppoolDynamicMappingComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := i["endip"]; ok {
			v := flattenFirewallIppoolDynamicMappingEndip(i["endip"], d, pre_append)
			tmp["endip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Endip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endport"
		if _, ok := i["endport"]; ok {
			v := flattenFirewallIppoolDynamicMappingEndport(i["endport"], d, pre_append)
			tmp["endport"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Endport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_ip"
		if _, ok := i["exclude-ip"]; ok {
			v := flattenFirewallIppoolDynamicMappingExcludeIp(i["exclude-ip"], d, pre_append)
			tmp["exclude_ip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-ExcludeIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "icmp_session_quota"
		if _, ok := i["icmp-session-quota"]; ok {
			v := flattenFirewallIppoolDynamicMappingIcmpSessionQuota(i["icmp-session-quota"], d, pre_append)
			tmp["icmp_session_quota"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-IcmpSessionQuota")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := i["nat64"]; ok {
			v := flattenFirewallIppoolDynamicMappingNat64(i["nat64"], d, pre_append)
			tmp["nat64"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Nat64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_blocks_per_user"
		if _, ok := i["num-blocks-per-user"]; ok {
			v := flattenFirewallIppoolDynamicMappingNumBlocksPerUser(i["num-blocks-per-user"], d, pre_append)
			tmp["num_blocks_per_user"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-NumBlocksPerUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_interim_log"
		if _, ok := i["pba-interim-log"]; ok {
			v := flattenFirewallIppoolDynamicMappingPbaInterimLog(i["pba-interim-log"], d, pre_append)
			tmp["pba_interim_log"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-PbaInterimLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_timeout"
		if _, ok := i["pba-timeout"]; ok {
			v := flattenFirewallIppoolDynamicMappingPbaTimeout(i["pba-timeout"], d, pre_append)
			tmp["pba_timeout"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-PbaTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_any_host"
		if _, ok := i["permit-any-host"]; ok {
			v := flattenFirewallIppoolDynamicMappingPermitAnyHost(i["permit-any-host"], d, pre_append)
			tmp["permit_any_host"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-PermitAnyHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_per_user"
		if _, ok := i["port-per-user"]; ok {
			v := flattenFirewallIppoolDynamicMappingPortPerUser(i["port-per-user"], d, pre_append)
			tmp["port_per_user"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-PortPerUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "privileged_port_use_pba"
		if _, ok := i["privileged-port-use-pba"]; ok {
			v := flattenFirewallIppoolDynamicMappingPrivilegedPortUsePba(i["privileged-port-use-pba"], d, pre_append)
			tmp["privileged_port_use_pba"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-PrivilegedPortUsePba")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_endip"
		if _, ok := i["source-endip"]; ok {
			v := flattenFirewallIppoolDynamicMappingSourceEndip(i["source-endip"], d, pre_append)
			tmp["source_endip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-SourceEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_prefix6"
		if _, ok := i["source-prefix6"]; ok {
			v := flattenFirewallIppoolDynamicMappingSourcePrefix6(i["source-prefix6"], d, pre_append)
			tmp["source_prefix6"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-SourcePrefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_startip"
		if _, ok := i["source-startip"]; ok {
			v := flattenFirewallIppoolDynamicMappingSourceStartip(i["source-startip"], d, pre_append)
			tmp["source_startip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-SourceStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := i["startip"]; ok {
			v := flattenFirewallIppoolDynamicMappingStartip(i["startip"], d, pre_append)
			tmp["startip"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Startip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startport"
		if _, ok := i["startport"]; ok {
			v := flattenFirewallIppoolDynamicMappingStartport(i["startport"], d, pre_append)
			tmp["startport"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Startport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_broadcast_in_ippool"
		if _, ok := i["subnet-broadcast-in-ippool"]; ok {
			v := flattenFirewallIppoolDynamicMappingSubnetBroadcastInIppool(i["subnet-broadcast-in-ippool"], d, pre_append)
			tmp["subnet_broadcast_in_ippool"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-SubnetBroadcastInIppool")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_session_quota"
		if _, ok := i["tcp-session-quota"]; ok {
			v := flattenFirewallIppoolDynamicMappingTcpSessionQuota(i["tcp-session-quota"], d, pre_append)
			tmp["tcp_session_quota"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-TcpSessionQuota")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallIppoolDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_session_quota"
		if _, ok := i["udp-session-quota"]; ok {
			v := flattenFirewallIppoolDynamicMappingUdpSessionQuota(i["udp-session-quota"], d, pre_append)
			tmp["udp_session_quota"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-UdpSessionQuota")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_clear"
		if _, ok := i["utilization-alarm-clear"]; ok {
			v := flattenFirewallIppoolDynamicMappingUtilizationAlarmClear(i["utilization-alarm-clear"], d, pre_append)
			tmp["utilization_alarm_clear"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-UtilizationAlarmClear")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_raise"
		if _, ok := i["utilization-alarm-raise"]; ok {
			v := flattenFirewallIppoolDynamicMappingUtilizationAlarmRaise(i["utilization-alarm-raise"], d, pre_append)
			tmp["utilization_alarm_raise"] = fortiAPISubPartPatch(v, "FirewallIppool-DynamicMapping-UtilizationAlarmRaise")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallIppoolDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallIppoolDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallIppoolDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallIppoolDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallIppoolDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallIppoolDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingArpIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnClientEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnClientIpv6Shift(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnClientStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnFixedalloc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnOverload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnPortEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnPortStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingCgnSpa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingClientPrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingEndport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingExcludeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIppoolDynamicMappingIcmpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingNumBlocksPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingPbaInterimLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingPbaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingPortPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingPrivilegedPortUsePba(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingSourceEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingSourcePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingSourceStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingStartport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingSubnetBroadcastInIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingTcpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingUdpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingUtilizationAlarmClear(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolDynamicMappingUtilizationAlarmRaise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolEndport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolExcludeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIppoolIcmpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolNumBlocksPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolPbaInterimLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolPbaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolPortPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolPrivilegedPortUsePba(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolSourceEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolSourcePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolSourceStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolStartport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolSubnetBroadcastInIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolTcpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolUdpSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolUtilizationAlarmClear(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppoolUtilizationAlarmRaise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIppool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat64_route", flattenFirewallIppoolAddNat64Route(o["add-nat64-route"], d, "add_nat64_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat64-route"], "FirewallIppool-AddNat64Route"); ok {
			if err = d.Set("add_nat64_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat64_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if err = d.Set("arp_intf", flattenFirewallIppoolArpIntf(o["arp-intf"], d, "arp_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-intf"], "FirewallIppool-ArpIntf"); ok {
			if err = d.Set("arp_intf", vv); err != nil {
				return fmt.Errorf("Error reading arp_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_intf: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallIppoolArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "FirewallIppool-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenFirewallIppoolAssociatedInterface(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "FirewallIppool-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("block_size", flattenFirewallIppoolBlockSize(o["block-size"], d, "block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-size"], "FirewallIppool-BlockSize"); ok {
			if err = d.Set("block_size", vv); err != nil {
				return fmt.Errorf("Error reading block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_size: %v", err)
		}
	}

	if err = d.Set("cgn_block_size", flattenFirewallIppoolCgnBlockSize(o["cgn-block-size"], d, "cgn_block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-block-size"], "FirewallIppool-CgnBlockSize"); ok {
			if err = d.Set("cgn_block_size", vv); err != nil {
				return fmt.Errorf("Error reading cgn_block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_block_size: %v", err)
		}
	}

	if err = d.Set("cgn_client_endip", flattenFirewallIppoolCgnClientEndip(o["cgn-client-endip"], d, "cgn_client_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-endip"], "FirewallIppool-CgnClientEndip"); ok {
			if err = d.Set("cgn_client_endip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_endip: %v", err)
		}
	}

	if err = d.Set("cgn_client_ipv6shift", flattenFirewallIppoolCgnClientIpv6Shift(o["cgn-client-ipv6shift"], d, "cgn_client_ipv6shift")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-ipv6shift"], "FirewallIppool-CgnClientIpv6Shift"); ok {
			if err = d.Set("cgn_client_ipv6shift", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
		}
	}

	if err = d.Set("cgn_client_startip", flattenFirewallIppoolCgnClientStartip(o["cgn-client-startip"], d, "cgn_client_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-startip"], "FirewallIppool-CgnClientStartip"); ok {
			if err = d.Set("cgn_client_startip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_startip: %v", err)
		}
	}

	if err = d.Set("cgn_fixedalloc", flattenFirewallIppoolCgnFixedalloc(o["cgn-fixedalloc"], d, "cgn_fixedalloc")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-fixedalloc"], "FirewallIppool-CgnFixedalloc"); ok {
			if err = d.Set("cgn_fixedalloc", vv); err != nil {
				return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
		}
	}

	if err = d.Set("cgn_overload", flattenFirewallIppoolCgnOverload(o["cgn-overload"], d, "cgn_overload")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-overload"], "FirewallIppool-CgnOverload"); ok {
			if err = d.Set("cgn_overload", vv); err != nil {
				return fmt.Errorf("Error reading cgn_overload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_overload: %v", err)
		}
	}

	if err = d.Set("cgn_port_end", flattenFirewallIppoolCgnPortEnd(o["cgn-port-end"], d, "cgn_port_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-end"], "FirewallIppool-CgnPortEnd"); ok {
			if err = d.Set("cgn_port_end", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_end: %v", err)
		}
	}

	if err = d.Set("cgn_port_start", flattenFirewallIppoolCgnPortStart(o["cgn-port-start"], d, "cgn_port_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-start"], "FirewallIppool-CgnPortStart"); ok {
			if err = d.Set("cgn_port_start", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_start: %v", err)
		}
	}

	if err = d.Set("cgn_spa", flattenFirewallIppoolCgnSpa(o["cgn-spa"], d, "cgn_spa")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-spa"], "FirewallIppool-CgnSpa"); ok {
			if err = d.Set("cgn_spa", vv); err != nil {
				return fmt.Errorf("Error reading cgn_spa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_spa: %v", err)
		}
	}

	if err = d.Set("client_prefix_length", flattenFirewallIppoolClientPrefixLength(o["client-prefix-length"], d, "client_prefix_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-prefix-length"], "FirewallIppool-ClientPrefixLength"); ok {
			if err = d.Set("client_prefix_length", vv); err != nil {
				return fmt.Errorf("Error reading client_prefix_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_prefix_length: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppoolComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallIppool-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallIppoolDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallIppool-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallIppoolDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallIppool-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("endip", flattenFirewallIppoolEndip(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "FirewallIppool-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("endport", flattenFirewallIppoolEndport(o["endport"], d, "endport")); err != nil {
		if vv, ok := fortiAPIPatch(o["endport"], "FirewallIppool-Endport"); ok {
			if err = d.Set("endport", vv); err != nil {
				return fmt.Errorf("Error reading endport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endport: %v", err)
		}
	}

	if err = d.Set("exclude_ip", flattenFirewallIppoolExcludeIp(o["exclude-ip"], d, "exclude_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-ip"], "FirewallIppool-ExcludeIp"); ok {
			if err = d.Set("exclude_ip", vv); err != nil {
				return fmt.Errorf("Error reading exclude_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_ip: %v", err)
		}
	}

	if err = d.Set("icmp_session_quota", flattenFirewallIppoolIcmpSessionQuota(o["icmp-session-quota"], d, "icmp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-session-quota"], "FirewallIppool-IcmpSessionQuota"); ok {
			if err = d.Set("icmp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading icmp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_session_quota: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallIppoolName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallIppool-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallIppoolNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "FirewallIppool-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("num_blocks_per_user", flattenFirewallIppoolNumBlocksPerUser(o["num-blocks-per-user"], d, "num_blocks_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["num-blocks-per-user"], "FirewallIppool-NumBlocksPerUser"); ok {
			if err = d.Set("num_blocks_per_user", vv); err != nil {
				return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
		}
	}

	if err = d.Set("pba_interim_log", flattenFirewallIppoolPbaInterimLog(o["pba-interim-log"], d, "pba_interim_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-interim-log"], "FirewallIppool-PbaInterimLog"); ok {
			if err = d.Set("pba_interim_log", vv); err != nil {
				return fmt.Errorf("Error reading pba_interim_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_interim_log: %v", err)
		}
	}

	if err = d.Set("pba_timeout", flattenFirewallIppoolPbaTimeout(o["pba-timeout"], d, "pba_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-timeout"], "FirewallIppool-PbaTimeout"); ok {
			if err = d.Set("pba_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pba_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_timeout: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenFirewallIppoolPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "FirewallIppool-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("port_per_user", flattenFirewallIppoolPortPerUser(o["port-per-user"], d, "port_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-per-user"], "FirewallIppool-PortPerUser"); ok {
			if err = d.Set("port_per_user", vv); err != nil {
				return fmt.Errorf("Error reading port_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_per_user: %v", err)
		}
	}

	if err = d.Set("privileged_port_use_pba", flattenFirewallIppoolPrivilegedPortUsePba(o["privileged-port-use-pba"], d, "privileged_port_use_pba")); err != nil {
		if vv, ok := fortiAPIPatch(o["privileged-port-use-pba"], "FirewallIppool-PrivilegedPortUsePba"); ok {
			if err = d.Set("privileged_port_use_pba", vv); err != nil {
				return fmt.Errorf("Error reading privileged_port_use_pba: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading privileged_port_use_pba: %v", err)
		}
	}

	if err = d.Set("source_endip", flattenFirewallIppoolSourceEndip(o["source-endip"], d, "source_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-endip"], "FirewallIppool-SourceEndip"); ok {
			if err = d.Set("source_endip", vv); err != nil {
				return fmt.Errorf("Error reading source_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_endip: %v", err)
		}
	}

	if err = d.Set("source_prefix6", flattenFirewallIppoolSourcePrefix6(o["source-prefix6"], d, "source_prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-prefix6"], "FirewallIppool-SourcePrefix6"); ok {
			if err = d.Set("source_prefix6", vv); err != nil {
				return fmt.Errorf("Error reading source_prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_prefix6: %v", err)
		}
	}

	if err = d.Set("source_startip", flattenFirewallIppoolSourceStartip(o["source-startip"], d, "source_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-startip"], "FirewallIppool-SourceStartip"); ok {
			if err = d.Set("source_startip", vv); err != nil {
				return fmt.Errorf("Error reading source_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_startip: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppoolStartip(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "FirewallIppool-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("startport", flattenFirewallIppoolStartport(o["startport"], d, "startport")); err != nil {
		if vv, ok := fortiAPIPatch(o["startport"], "FirewallIppool-Startport"); ok {
			if err = d.Set("startport", vv); err != nil {
				return fmt.Errorf("Error reading startport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startport: %v", err)
		}
	}

	if err = d.Set("subnet_broadcast_in_ippool", flattenFirewallIppoolSubnetBroadcastInIppool(o["subnet-broadcast-in-ippool"], d, "subnet_broadcast_in_ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-broadcast-in-ippool"], "FirewallIppool-SubnetBroadcastInIppool"); ok {
			if err = d.Set("subnet_broadcast_in_ippool", vv); err != nil {
				return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
		}
	}

	if err = d.Set("tcp_session_quota", flattenFirewallIppoolTcpSessionQuota(o["tcp-session-quota"], d, "tcp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-quota"], "FirewallIppool-TcpSessionQuota"); ok {
			if err = d.Set("tcp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_quota: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallIppoolType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallIppool-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("udp_session_quota", flattenFirewallIppoolUdpSessionQuota(o["udp-session-quota"], d, "udp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-session-quota"], "FirewallIppool-UdpSessionQuota"); ok {
			if err = d.Set("udp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading udp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_session_quota: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_clear", flattenFirewallIppoolUtilizationAlarmClear(o["utilization-alarm-clear"], d, "utilization_alarm_clear")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-clear"], "FirewallIppool-UtilizationAlarmClear"); ok {
			if err = d.Set("utilization_alarm_clear", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_raise", flattenFirewallIppoolUtilizationAlarmRaise(o["utilization-alarm-raise"], d, "utilization_alarm_raise")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-raise"], "FirewallIppool-UtilizationAlarmRaise"); ok {
			if err = d.Set("utilization_alarm_raise", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
		}
	}

	return nil
}

func flattenFirewallIppoolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIppoolAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolArpIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIppoolArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIppoolBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnClientEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnClientIpv6Shift(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnClientStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnFixedalloc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnOverload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnPortEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnPortStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolCgnSpa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolClientPrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallIppoolDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat64-route"], _ = expandFirewallIppoolDynamicMappingAddNat64Route(d, i["add_nat64_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-intf"], _ = expandFirewallIppoolDynamicMappingArpIntf(d, i["arp_intf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandFirewallIppoolDynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["associated-interface"], _ = expandFirewallIppoolDynamicMappingAssociatedInterface(d, i["associated_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "block_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["block-size"], _ = expandFirewallIppoolDynamicMappingBlockSize(d, i["block_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_block_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-block-size"], _ = expandFirewallIppoolDynamicMappingCgnBlockSize(d, i["cgn_block_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-endip"], _ = expandFirewallIppoolDynamicMappingCgnClientEndip(d, i["cgn_client_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_ipv6shift"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-ipv6shift"], _ = expandFirewallIppoolDynamicMappingCgnClientIpv6Shift(d, i["cgn_client_ipv6shift"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-startip"], _ = expandFirewallIppoolDynamicMappingCgnClientStartip(d, i["cgn_client_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_fixedalloc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-fixedalloc"], _ = expandFirewallIppoolDynamicMappingCgnFixedalloc(d, i["cgn_fixedalloc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_overload"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-overload"], _ = expandFirewallIppoolDynamicMappingCgnOverload(d, i["cgn_overload"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_end"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-port-end"], _ = expandFirewallIppoolDynamicMappingCgnPortEnd(d, i["cgn_port_end"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_start"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-port-start"], _ = expandFirewallIppoolDynamicMappingCgnPortStart(d, i["cgn_port_start"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_spa"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-spa"], _ = expandFirewallIppoolDynamicMappingCgnSpa(d, i["cgn_spa"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_prefix_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-prefix-length"], _ = expandFirewallIppoolDynamicMappingClientPrefixLength(d, i["client_prefix_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandFirewallIppoolDynamicMappingComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endip"], _ = expandFirewallIppoolDynamicMappingEndip(d, i["endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endport"], _ = expandFirewallIppoolDynamicMappingEndport(d, i["endport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclude-ip"], _ = expandFirewallIppoolDynamicMappingExcludeIp(d, i["exclude_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "icmp_session_quota"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["icmp-session-quota"], _ = expandFirewallIppoolDynamicMappingIcmpSessionQuota(d, i["icmp_session_quota"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat64"], _ = expandFirewallIppoolDynamicMappingNat64(d, i["nat64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_blocks_per_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["num-blocks-per-user"], _ = expandFirewallIppoolDynamicMappingNumBlocksPerUser(d, i["num_blocks_per_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_interim_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pba-interim-log"], _ = expandFirewallIppoolDynamicMappingPbaInterimLog(d, i["pba_interim_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pba-timeout"], _ = expandFirewallIppoolDynamicMappingPbaTimeout(d, i["pba_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_any_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["permit-any-host"], _ = expandFirewallIppoolDynamicMappingPermitAnyHost(d, i["permit_any_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_per_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-per-user"], _ = expandFirewallIppoolDynamicMappingPortPerUser(d, i["port_per_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "privileged_port_use_pba"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["privileged-port-use-pba"], _ = expandFirewallIppoolDynamicMappingPrivilegedPortUsePba(d, i["privileged_port_use_pba"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-endip"], _ = expandFirewallIppoolDynamicMappingSourceEndip(d, i["source_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-prefix6"], _ = expandFirewallIppoolDynamicMappingSourcePrefix6(d, i["source_prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-startip"], _ = expandFirewallIppoolDynamicMappingSourceStartip(d, i["source_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startip"], _ = expandFirewallIppoolDynamicMappingStartip(d, i["startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startport"], _ = expandFirewallIppoolDynamicMappingStartport(d, i["startport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_broadcast_in_ippool"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet-broadcast-in-ippool"], _ = expandFirewallIppoolDynamicMappingSubnetBroadcastInIppool(d, i["subnet_broadcast_in_ippool"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_session_quota"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tcp-session-quota"], _ = expandFirewallIppoolDynamicMappingTcpSessionQuota(d, i["tcp_session_quota"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallIppoolDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_session_quota"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["udp-session-quota"], _ = expandFirewallIppoolDynamicMappingUdpSessionQuota(d, i["udp_session_quota"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_clear"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utilization-alarm-clear"], _ = expandFirewallIppoolDynamicMappingUtilizationAlarmClear(d, i["utilization_alarm_clear"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_raise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utilization-alarm-raise"], _ = expandFirewallIppoolDynamicMappingUtilizationAlarmRaise(d, i["utilization_alarm_raise"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallIppoolDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallIppoolDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallIppoolDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallIppoolDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingArpIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnClientEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnClientIpv6Shift(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnClientStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnFixedalloc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnOverload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnPortEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnPortStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingCgnSpa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingClientPrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingEndport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingExcludeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIppoolDynamicMappingIcmpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingNumBlocksPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingPbaInterimLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingPbaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingPortPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingPrivilegedPortUsePba(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingSourceEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingSourcePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingSourceStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingStartport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingSubnetBroadcastInIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingTcpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingUdpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingUtilizationAlarmClear(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolDynamicMappingUtilizationAlarmRaise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolEndport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolExcludeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIppoolIcmpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolNumBlocksPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPbaInterimLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPbaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPortPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolPrivilegedPortUsePba(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSourceEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSourcePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSourceStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolStartport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolSubnetBroadcastInIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolTcpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolUdpSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolUtilizationAlarmClear(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppoolUtilizationAlarmRaise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIppool(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat64_route"); ok || d.HasChange("add_nat64_route") {
		t, err := expandFirewallIppoolAddNat64Route(d, v, "add_nat64_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_intf"); ok || d.HasChange("arp_intf") {
		t, err := expandFirewallIppoolArpIntf(d, v, "arp_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-intf"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandFirewallIppoolArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandFirewallIppoolAssociatedInterface(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("block_size"); ok || d.HasChange("block_size") {
		t, err := expandFirewallIppoolBlockSize(d, v, "block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_block_size"); ok || d.HasChange("cgn_block_size") {
		t, err := expandFirewallIppoolCgnBlockSize(d, v, "cgn_block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_endip"); ok || d.HasChange("cgn_client_endip") {
		t, err := expandFirewallIppoolCgnClientEndip(d, v, "cgn_client_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-endip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_ipv6shift"); ok || d.HasChange("cgn_client_ipv6shift") {
		t, err := expandFirewallIppoolCgnClientIpv6Shift(d, v, "cgn_client_ipv6shift")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-ipv6shift"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_startip"); ok || d.HasChange("cgn_client_startip") {
		t, err := expandFirewallIppoolCgnClientStartip(d, v, "cgn_client_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-startip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_fixedalloc"); ok || d.HasChange("cgn_fixedalloc") {
		t, err := expandFirewallIppoolCgnFixedalloc(d, v, "cgn_fixedalloc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-fixedalloc"] = t
		}
	}

	if v, ok := d.GetOk("cgn_overload"); ok || d.HasChange("cgn_overload") {
		t, err := expandFirewallIppoolCgnOverload(d, v, "cgn_overload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-overload"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_end"); ok || d.HasChange("cgn_port_end") {
		t, err := expandFirewallIppoolCgnPortEnd(d, v, "cgn_port_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-end"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_start"); ok || d.HasChange("cgn_port_start") {
		t, err := expandFirewallIppoolCgnPortStart(d, v, "cgn_port_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-start"] = t
		}
	}

	if v, ok := d.GetOk("cgn_spa"); ok || d.HasChange("cgn_spa") {
		t, err := expandFirewallIppoolCgnSpa(d, v, "cgn_spa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-spa"] = t
		}
	}

	if v, ok := d.GetOk("client_prefix_length"); ok || d.HasChange("client_prefix_length") {
		t, err := expandFirewallIppoolClientPrefixLength(d, v, "client_prefix_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-prefix-length"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallIppoolComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallIppoolDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandFirewallIppoolEndip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("endport"); ok || d.HasChange("endport") {
		t, err := expandFirewallIppoolEndport(d, v, "endport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endport"] = t
		}
	}

	if v, ok := d.GetOk("exclude_ip"); ok || d.HasChange("exclude_ip") {
		t, err := expandFirewallIppoolExcludeIp(d, v, "exclude_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-ip"] = t
		}
	}

	if v, ok := d.GetOk("icmp_session_quota"); ok || d.HasChange("icmp_session_quota") {
		t, err := expandFirewallIppoolIcmpSessionQuota(d, v, "icmp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallIppoolName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandFirewallIppoolNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("num_blocks_per_user"); ok || d.HasChange("num_blocks_per_user") {
		t, err := expandFirewallIppoolNumBlocksPerUser(d, v, "num_blocks_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["num-blocks-per-user"] = t
		}
	}

	if v, ok := d.GetOk("pba_interim_log"); ok || d.HasChange("pba_interim_log") {
		t, err := expandFirewallIppoolPbaInterimLog(d, v, "pba_interim_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-interim-log"] = t
		}
	}

	if v, ok := d.GetOk("pba_timeout"); ok || d.HasChange("pba_timeout") {
		t, err := expandFirewallIppoolPbaTimeout(d, v, "pba_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-timeout"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok || d.HasChange("permit_any_host") {
		t, err := expandFirewallIppoolPermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("port_per_user"); ok || d.HasChange("port_per_user") {
		t, err := expandFirewallIppoolPortPerUser(d, v, "port_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-per-user"] = t
		}
	}

	if v, ok := d.GetOk("privileged_port_use_pba"); ok || d.HasChange("privileged_port_use_pba") {
		t, err := expandFirewallIppoolPrivilegedPortUsePba(d, v, "privileged_port_use_pba")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["privileged-port-use-pba"] = t
		}
	}

	if v, ok := d.GetOk("source_endip"); ok || d.HasChange("source_endip") {
		t, err := expandFirewallIppoolSourceEndip(d, v, "source_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-endip"] = t
		}
	}

	if v, ok := d.GetOk("source_prefix6"); ok || d.HasChange("source_prefix6") {
		t, err := expandFirewallIppoolSourcePrefix6(d, v, "source_prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-prefix6"] = t
		}
	}

	if v, ok := d.GetOk("source_startip"); ok || d.HasChange("source_startip") {
		t, err := expandFirewallIppoolSourceStartip(d, v, "source_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-startip"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandFirewallIppoolStartip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("startport"); ok || d.HasChange("startport") {
		t, err := expandFirewallIppoolStartport(d, v, "startport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startport"] = t
		}
	}

	if v, ok := d.GetOk("subnet_broadcast_in_ippool"); ok || d.HasChange("subnet_broadcast_in_ippool") {
		t, err := expandFirewallIppoolSubnetBroadcastInIppool(d, v, "subnet_broadcast_in_ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-broadcast-in-ippool"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_quota"); ok || d.HasChange("tcp_session_quota") {
		t, err := expandFirewallIppoolTcpSessionQuota(d, v, "tcp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallIppoolType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("udp_session_quota"); ok || d.HasChange("udp_session_quota") {
		t, err := expandFirewallIppoolUdpSessionQuota(d, v, "udp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_clear"); ok || d.HasChange("utilization_alarm_clear") {
		t, err := expandFirewallIppoolUtilizationAlarmClear(d, v, "utilization_alarm_clear")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-clear"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_raise"); ok || d.HasChange("utilization_alarm_raise") {
		t, err := expandFirewallIppoolUtilizationAlarmRaise(d, v, "utilization_alarm_raise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-raise"] = t
		}
	}

	return &obj, nil
}
