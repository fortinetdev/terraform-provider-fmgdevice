// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VPN autokey tunnel.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecPhase2Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase2InterfaceCreate,
		Read:   resourceVpnIpsecPhase2InterfaceRead,
		Update: resourceVpnIpsecPhase2InterfaceUpdate,
		Delete: resourceVpnIpsecPhase2InterfaceDelete,

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
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_end_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_name6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dst_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_start_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_subnet6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encapsulation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbound_dscp_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initiator_ts_narrow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_df": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keylife_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keylifekbs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"keylifeseconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l2tp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phase1name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_end_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_name6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_start_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_subnet6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpnIpsecPhase2InterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase2Interface(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2Interface resource while getting object: %v", err)
	}

	_, err = c.CreateVpnIpsecPhase2Interface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2Interface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase2InterfaceRead(d, m)
}

func resourceVpnIpsecPhase2InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase2Interface(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2Interface resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnIpsecPhase2Interface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase2InterfaceRead(d, m)
}

func resourceVpnIpsecPhase2InterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnIpsecPhase2Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase2Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase2InterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecPhase2Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase2Interface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2Interface resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase2InterfaceAddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoDiscoverySender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDhcpIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDhgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstEndIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceDstName6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceDstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstStartIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceDstSubnet6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceEncapsulation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceInboundDscpCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceInitiatorTsNarrow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceIpv4Df(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifekbs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifeseconds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceL2Tp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfacePfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfacePhase1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceProposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceRouteOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSingleSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcEndIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceSrcName6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcStartIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2InterfaceSrcSubnet6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase2Interface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("add_route", flattenVpnIpsecPhase2InterfaceAddRoute(o["add-route"], d, "add_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-route"], "VpnIpsecPhase2Interface-AddRoute"); ok {
			if err = d.Set("add_route", vv); err != nil {
				return fmt.Errorf("Error reading add_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("auto_discovery_forwarder", flattenVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(o["auto-discovery-forwarder"], d, "auto_discovery_forwarder")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-forwarder"], "VpnIpsecPhase2Interface-AutoDiscoveryForwarder"); ok {
			if err = d.Set("auto_discovery_forwarder", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
		}
	}

	if err = d.Set("auto_discovery_sender", flattenVpnIpsecPhase2InterfaceAutoDiscoverySender(o["auto-discovery-sender"], d, "auto_discovery_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-sender"], "VpnIpsecPhase2Interface-AutoDiscoverySender"); ok {
			if err = d.Set("auto_discovery_sender", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase2InterfaceAutoNegotiate(o["auto-negotiate"], d, "auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-negotiate"], "VpnIpsecPhase2Interface-AutoNegotiate"); ok {
			if err = d.Set("auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase2InterfaceComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnIpsecPhase2Interface-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dhcp_ipsec", flattenVpnIpsecPhase2InterfaceDhcpIpsec(o["dhcp-ipsec"], d, "dhcp_ipsec")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ipsec"], "VpnIpsecPhase2Interface-DhcpIpsec"); ok {
			if err = d.Set("dhcp_ipsec", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase2InterfaceDhgrp(o["dhgrp"], d, "dhgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhgrp"], "VpnIpsecPhase2Interface-Dhgrp"); ok {
			if err = d.Set("dhgrp", vv); err != nil {
				return fmt.Errorf("Error reading dhgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenVpnIpsecPhase2InterfaceDiffserv(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "VpnIpsecPhase2Interface-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenVpnIpsecPhase2InterfaceDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "VpnIpsecPhase2Interface-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dst_addr_type", flattenVpnIpsecPhase2InterfaceDstAddrType(o["dst-addr-type"], d, "dst_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr-type"], "VpnIpsecPhase2Interface-DstAddrType"); ok {
			if err = d.Set("dst_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr_type: %v", err)
		}
	}

	if err = d.Set("dst_end_ip", flattenVpnIpsecPhase2InterfaceDstEndIp(o["dst-end-ip"], d, "dst_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-end-ip"], "VpnIpsecPhase2Interface-DstEndIp"); ok {
			if err = d.Set("dst_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_end_ip: %v", err)
		}
	}

	if err = d.Set("dst_end_ip6", flattenVpnIpsecPhase2InterfaceDstEndIp6(o["dst-end-ip6"], d, "dst_end_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-end-ip6"], "VpnIpsecPhase2Interface-DstEndIp6"); ok {
			if err = d.Set("dst_end_ip6", vv); err != nil {
				return fmt.Errorf("Error reading dst_end_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_end_ip6: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenVpnIpsecPhase2InterfaceDstName(o["dst-name"], d, "dst_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-name"], "VpnIpsecPhase2Interface-DstName"); ok {
			if err = d.Set("dst_name", vv); err != nil {
				return fmt.Errorf("Error reading dst_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_name6", flattenVpnIpsecPhase2InterfaceDstName6(o["dst-name6"], d, "dst_name6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-name6"], "VpnIpsecPhase2Interface-DstName6"); ok {
			if err = d.Set("dst_name6", vv); err != nil {
				return fmt.Errorf("Error reading dst_name6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenVpnIpsecPhase2InterfaceDstPort(o["dst-port"], d, "dst_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port"], "VpnIpsecPhase2Interface-DstPort"); ok {
			if err = d.Set("dst_port", vv); err != nil {
				return fmt.Errorf("Error reading dst_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	if err = d.Set("dst_start_ip", flattenVpnIpsecPhase2InterfaceDstStartIp(o["dst-start-ip"], d, "dst_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-start-ip"], "VpnIpsecPhase2Interface-DstStartIp"); ok {
			if err = d.Set("dst_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_start_ip: %v", err)
		}
	}

	if err = d.Set("dst_start_ip6", flattenVpnIpsecPhase2InterfaceDstStartIp6(o["dst-start-ip6"], d, "dst_start_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-start-ip6"], "VpnIpsecPhase2Interface-DstStartIp6"); ok {
			if err = d.Set("dst_start_ip6", vv); err != nil {
				return fmt.Errorf("Error reading dst_start_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_start_ip6: %v", err)
		}
	}

	if err = d.Set("dst_subnet", flattenVpnIpsecPhase2InterfaceDstSubnet(o["dst-subnet"], d, "dst_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-subnet"], "VpnIpsecPhase2Interface-DstSubnet"); ok {
			if err = d.Set("dst_subnet", vv); err != nil {
				return fmt.Errorf("Error reading dst_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_subnet: %v", err)
		}
	}

	if err = d.Set("dst_subnet6", flattenVpnIpsecPhase2InterfaceDstSubnet6(o["dst-subnet6"], d, "dst_subnet6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-subnet6"], "VpnIpsecPhase2Interface-DstSubnet6"); ok {
			if err = d.Set("dst_subnet6", vv); err != nil {
				return fmt.Errorf("Error reading dst_subnet6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_subnet6: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase2InterfaceEncapsulation(o["encapsulation"], d, "encapsulation")); err != nil {
		if vv, ok := fortiAPIPatch(o["encapsulation"], "VpnIpsecPhase2Interface-Encapsulation"); ok {
			if err = d.Set("encapsulation", vv); err != nil {
				return fmt.Errorf("Error reading encapsulation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase2InterfaceInboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "VpnIpsecPhase2Interface-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("initiator_ts_narrow", flattenVpnIpsecPhase2InterfaceInitiatorTsNarrow(o["initiator-ts-narrow"], d, "initiator_ts_narrow")); err != nil {
		if vv, ok := fortiAPIPatch(o["initiator-ts-narrow"], "VpnIpsecPhase2Interface-InitiatorTsNarrow"); ok {
			if err = d.Set("initiator_ts_narrow", vv); err != nil {
				return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
		}
	}

	if err = d.Set("ipv4_df", flattenVpnIpsecPhase2InterfaceIpv4Df(o["ipv4-df"], d, "ipv4_df")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-df"], "VpnIpsecPhase2Interface-Ipv4Df"); ok {
			if err = d.Set("ipv4_df", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_df: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_df: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase2InterfaceKeepalive(o["keepalive"], d, "keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive"], "VpnIpsecPhase2Interface-Keepalive"); ok {
			if err = d.Set("keepalive", vv); err != nil {
				return fmt.Errorf("Error reading keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("keylife_type", flattenVpnIpsecPhase2InterfaceKeylifeType(o["keylife-type"], d, "keylife_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylife-type"], "VpnIpsecPhase2Interface-KeylifeType"); ok {
			if err = d.Set("keylife_type", vv); err != nil {
				return fmt.Errorf("Error reading keylife_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylife_type: %v", err)
		}
	}

	if err = d.Set("keylifekbs", flattenVpnIpsecPhase2InterfaceKeylifekbs(o["keylifekbs"], d, "keylifekbs")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylifekbs"], "VpnIpsecPhase2Interface-Keylifekbs"); ok {
			if err = d.Set("keylifekbs", vv); err != nil {
				return fmt.Errorf("Error reading keylifekbs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylifekbs: %v", err)
		}
	}

	if err = d.Set("keylifeseconds", flattenVpnIpsecPhase2InterfaceKeylifeseconds(o["keylifeseconds"], d, "keylifeseconds")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylifeseconds"], "VpnIpsecPhase2Interface-Keylifeseconds"); ok {
			if err = d.Set("keylifeseconds", vv); err != nil {
				return fmt.Errorf("Error reading keylifeseconds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylifeseconds: %v", err)
		}
	}

	if err = d.Set("l2tp", flattenVpnIpsecPhase2InterfaceL2Tp(o["l2tp"], d, "l2tp")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tp"], "VpnIpsecPhase2Interface-L2Tp"); ok {
			if err = d.Set("l2tp", vv); err != nil {
				return fmt.Errorf("Error reading l2tp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tp: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnIpsecPhase2InterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnIpsecPhase2Interface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pfs", flattenVpnIpsecPhase2InterfacePfs(o["pfs"], d, "pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfs"], "VpnIpsecPhase2Interface-Pfs"); ok {
			if err = d.Set("pfs", vv); err != nil {
				return fmt.Errorf("Error reading pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfs: %v", err)
		}
	}

	if err = d.Set("phase1name", flattenVpnIpsecPhase2InterfacePhase1Name(o["phase1name"], d, "phase1name")); err != nil {
		if vv, ok := fortiAPIPatch(o["phase1name"], "VpnIpsecPhase2Interface-Phase1Name"); ok {
			if err = d.Set("phase1name", vv); err != nil {
				return fmt.Errorf("Error reading phase1name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phase1name: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase2InterfaceProposal(o["proposal"], d, "proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["proposal"], "VpnIpsecPhase2Interface-Proposal"); ok {
			if err = d.Set("proposal", vv); err != nil {
				return fmt.Errorf("Error reading proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("protocol", flattenVpnIpsecPhase2InterfaceProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "VpnIpsecPhase2Interface-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("replay", flattenVpnIpsecPhase2InterfaceReplay(o["replay"], d, "replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["replay"], "VpnIpsecPhase2Interface-Replay"); ok {
			if err = d.Set("replay", vv); err != nil {
				return fmt.Errorf("Error reading replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replay: %v", err)
		}
	}

	if err = d.Set("route_overlap", flattenVpnIpsecPhase2InterfaceRouteOverlap(o["route-overlap"], d, "route_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-overlap"], "VpnIpsecPhase2Interface-RouteOverlap"); ok {
			if err = d.Set("route_overlap", vv); err != nil {
				return fmt.Errorf("Error reading route_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_overlap: %v", err)
		}
	}

	if err = d.Set("single_source", flattenVpnIpsecPhase2InterfaceSingleSource(o["single-source"], d, "single_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-source"], "VpnIpsecPhase2Interface-SingleSource"); ok {
			if err = d.Set("single_source", vv); err != nil {
				return fmt.Errorf("Error reading single_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_source: %v", err)
		}
	}

	if err = d.Set("src_addr_type", flattenVpnIpsecPhase2InterfaceSrcAddrType(o["src-addr-type"], d, "src_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-addr-type"], "VpnIpsecPhase2Interface-SrcAddrType"); ok {
			if err = d.Set("src_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading src_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_addr_type: %v", err)
		}
	}

	if err = d.Set("src_end_ip", flattenVpnIpsecPhase2InterfaceSrcEndIp(o["src-end-ip"], d, "src_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-end-ip"], "VpnIpsecPhase2Interface-SrcEndIp"); ok {
			if err = d.Set("src_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_end_ip: %v", err)
		}
	}

	if err = d.Set("src_end_ip6", flattenVpnIpsecPhase2InterfaceSrcEndIp6(o["src-end-ip6"], d, "src_end_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-end-ip6"], "VpnIpsecPhase2Interface-SrcEndIp6"); ok {
			if err = d.Set("src_end_ip6", vv); err != nil {
				return fmt.Errorf("Error reading src_end_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_end_ip6: %v", err)
		}
	}

	if err = d.Set("src_name", flattenVpnIpsecPhase2InterfaceSrcName(o["src-name"], d, "src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-name"], "VpnIpsecPhase2Interface-SrcName"); ok {
			if err = d.Set("src_name", vv); err != nil {
				return fmt.Errorf("Error reading src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_name: %v", err)
		}
	}

	if err = d.Set("src_name6", flattenVpnIpsecPhase2InterfaceSrcName6(o["src-name6"], d, "src_name6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-name6"], "VpnIpsecPhase2Interface-SrcName6"); ok {
			if err = d.Set("src_name6", vv); err != nil {
				return fmt.Errorf("Error reading src_name6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_name6: %v", err)
		}
	}

	if err = d.Set("src_port", flattenVpnIpsecPhase2InterfaceSrcPort(o["src-port"], d, "src_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-port"], "VpnIpsecPhase2Interface-SrcPort"); ok {
			if err = d.Set("src_port", vv); err != nil {
				return fmt.Errorf("Error reading src_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_port: %v", err)
		}
	}

	if err = d.Set("src_start_ip", flattenVpnIpsecPhase2InterfaceSrcStartIp(o["src-start-ip"], d, "src_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-start-ip"], "VpnIpsecPhase2Interface-SrcStartIp"); ok {
			if err = d.Set("src_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_start_ip: %v", err)
		}
	}

	if err = d.Set("src_start_ip6", flattenVpnIpsecPhase2InterfaceSrcStartIp6(o["src-start-ip6"], d, "src_start_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-start-ip6"], "VpnIpsecPhase2Interface-SrcStartIp6"); ok {
			if err = d.Set("src_start_ip6", vv); err != nil {
				return fmt.Errorf("Error reading src_start_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_start_ip6: %v", err)
		}
	}

	if err = d.Set("src_subnet", flattenVpnIpsecPhase2InterfaceSrcSubnet(o["src-subnet"], d, "src_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-subnet"], "VpnIpsecPhase2Interface-SrcSubnet"); ok {
			if err = d.Set("src_subnet", vv); err != nil {
				return fmt.Errorf("Error reading src_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_subnet: %v", err)
		}
	}

	if err = d.Set("src_subnet6", flattenVpnIpsecPhase2InterfaceSrcSubnet6(o["src-subnet6"], d, "src_subnet6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-subnet6"], "VpnIpsecPhase2Interface-SrcSubnet6"); ok {
			if err = d.Set("src_subnet6", vv); err != nil {
				return fmt.Errorf("Error reading src_subnet6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_subnet6: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase2InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecPhase2InterfaceAddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoDiscoverySender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDhcpIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDhgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstEndIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceDstName6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceDstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstStartIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase2InterfaceDstSubnet6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceEncapsulation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceInboundDscpCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceInitiatorTsNarrow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceIpv4Df(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifeType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifekbs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifeseconds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceL2Tp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfacePfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfacePhase1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceProposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceReplay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceRouteOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSingleSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcEndIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceSrcName6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2InterfaceSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcStartIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase2InterfaceSrcSubnet6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase2Interface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_route"); ok || d.HasChange("add_route") {
		t, err := expandVpnIpsecPhase2InterfaceAddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_forwarder"); ok || d.HasChange("auto_discovery_forwarder") {
		t, err := expandVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(d, v, "auto_discovery_forwarder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-forwarder"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_sender"); ok || d.HasChange("auto_discovery_sender") {
		t, err := expandVpnIpsecPhase2InterfaceAutoDiscoverySender(d, v, "auto_discovery_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-sender"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok || d.HasChange("auto_negotiate") {
		t, err := expandVpnIpsecPhase2InterfaceAutoNegotiate(d, v, "auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnIpsecPhase2InterfaceComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ipsec"); ok || d.HasChange("dhcp_ipsec") {
		t, err := expandVpnIpsecPhase2InterfaceDhcpIpsec(d, v, "dhcp_ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok || d.HasChange("dhgrp") {
		t, err := expandVpnIpsecPhase2InterfaceDhgrp(d, v, "dhgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok || d.HasChange("diffserv") {
		t, err := expandVpnIpsecPhase2InterfaceDiffserv(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandVpnIpsecPhase2InterfaceDiffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr_type"); ok || d.HasChange("dst_addr_type") {
		t, err := expandVpnIpsecPhase2InterfaceDstAddrType(d, v, "dst_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip"); ok || d.HasChange("dst_end_ip") {
		t, err := expandVpnIpsecPhase2InterfaceDstEndIp(d, v, "dst_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip6"); ok || d.HasChange("dst_end_ip6") {
		t, err := expandVpnIpsecPhase2InterfaceDstEndIp6(d, v, "dst_end_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok || d.HasChange("dst_name") {
		t, err := expandVpnIpsecPhase2InterfaceDstName(d, v, "dst_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name"] = t
		}
	}

	if v, ok := d.GetOk("dst_name6"); ok || d.HasChange("dst_name6") {
		t, err := expandVpnIpsecPhase2InterfaceDstName6(d, v, "dst_name6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name6"] = t
		}
	}

	if v, ok := d.GetOk("dst_port"); ok || d.HasChange("dst_port") {
		t, err := expandVpnIpsecPhase2InterfaceDstPort(d, v, "dst_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip"); ok || d.HasChange("dst_start_ip") {
		t, err := expandVpnIpsecPhase2InterfaceDstStartIp(d, v, "dst_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip6"); ok || d.HasChange("dst_start_ip6") {
		t, err := expandVpnIpsecPhase2InterfaceDstStartIp6(d, v, "dst_start_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet"); ok || d.HasChange("dst_subnet") {
		t, err := expandVpnIpsecPhase2InterfaceDstSubnet(d, v, "dst_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet6"); ok || d.HasChange("dst_subnet6") {
		t, err := expandVpnIpsecPhase2InterfaceDstSubnet6(d, v, "dst_subnet6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet6"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok || d.HasChange("encapsulation") {
		t, err := expandVpnIpsecPhase2InterfaceEncapsulation(d, v, "encapsulation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok || d.HasChange("inbound_dscp_copy") {
		t, err := expandVpnIpsecPhase2InterfaceInboundDscpCopy(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("initiator_ts_narrow"); ok || d.HasChange("initiator_ts_narrow") {
		t, err := expandVpnIpsecPhase2InterfaceInitiatorTsNarrow(d, v, "initiator_ts_narrow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator-ts-narrow"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_df"); ok || d.HasChange("ipv4_df") {
		t, err := expandVpnIpsecPhase2InterfaceIpv4Df(d, v, "ipv4_df")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-df"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok || d.HasChange("keepalive") {
		t, err := expandVpnIpsecPhase2InterfaceKeepalive(d, v, "keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("keylife_type"); ok || d.HasChange("keylife_type") {
		t, err := expandVpnIpsecPhase2InterfaceKeylifeType(d, v, "keylife_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife-type"] = t
		}
	}

	if v, ok := d.GetOk("keylifekbs"); ok || d.HasChange("keylifekbs") {
		t, err := expandVpnIpsecPhase2InterfaceKeylifekbs(d, v, "keylifekbs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifekbs"] = t
		}
	}

	if v, ok := d.GetOk("keylifeseconds"); ok || d.HasChange("keylifeseconds") {
		t, err := expandVpnIpsecPhase2InterfaceKeylifeseconds(d, v, "keylifeseconds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifeseconds"] = t
		}
	}

	if v, ok := d.GetOk("l2tp"); ok || d.HasChange("l2tp") {
		t, err := expandVpnIpsecPhase2InterfaceL2Tp(d, v, "l2tp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnIpsecPhase2InterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pfs"); ok || d.HasChange("pfs") {
		t, err := expandVpnIpsecPhase2InterfacePfs(d, v, "pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfs"] = t
		}
	}

	if v, ok := d.GetOk("phase1name"); ok || d.HasChange("phase1name") {
		t, err := expandVpnIpsecPhase2InterfacePhase1Name(d, v, "phase1name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase1name"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok || d.HasChange("proposal") {
		t, err := expandVpnIpsecPhase2InterfaceProposal(d, v, "proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandVpnIpsecPhase2InterfaceProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("replay"); ok || d.HasChange("replay") {
		t, err := expandVpnIpsecPhase2InterfaceReplay(d, v, "replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay"] = t
		}
	}

	if v, ok := d.GetOk("route_overlap"); ok || d.HasChange("route_overlap") {
		t, err := expandVpnIpsecPhase2InterfaceRouteOverlap(d, v, "route_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-overlap"] = t
		}
	}

	if v, ok := d.GetOk("single_source"); ok || d.HasChange("single_source") {
		t, err := expandVpnIpsecPhase2InterfaceSingleSource(d, v, "single_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-source"] = t
		}
	}

	if v, ok := d.GetOk("src_addr_type"); ok || d.HasChange("src_addr_type") {
		t, err := expandVpnIpsecPhase2InterfaceSrcAddrType(d, v, "src_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip"); ok || d.HasChange("src_end_ip") {
		t, err := expandVpnIpsecPhase2InterfaceSrcEndIp(d, v, "src_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip6"); ok || d.HasChange("src_end_ip6") {
		t, err := expandVpnIpsecPhase2InterfaceSrcEndIp6(d, v, "src_end_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_name"); ok || d.HasChange("src_name") {
		t, err := expandVpnIpsecPhase2InterfaceSrcName(d, v, "src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name"] = t
		}
	}

	if v, ok := d.GetOk("src_name6"); ok || d.HasChange("src_name6") {
		t, err := expandVpnIpsecPhase2InterfaceSrcName6(d, v, "src_name6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name6"] = t
		}
	}

	if v, ok := d.GetOk("src_port"); ok || d.HasChange("src_port") {
		t, err := expandVpnIpsecPhase2InterfaceSrcPort(d, v, "src_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip"); ok || d.HasChange("src_start_ip") {
		t, err := expandVpnIpsecPhase2InterfaceSrcStartIp(d, v, "src_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip6"); ok || d.HasChange("src_start_ip6") {
		t, err := expandVpnIpsecPhase2InterfaceSrcStartIp6(d, v, "src_start_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet"); ok || d.HasChange("src_subnet") {
		t, err := expandVpnIpsecPhase2InterfaceSrcSubnet(d, v, "src_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet6"); ok || d.HasChange("src_subnet6") {
		t, err := expandVpnIpsecPhase2InterfaceSrcSubnet6(d, v, "src_subnet6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet6"] = t
		}
	}

	return &obj, nil
}
