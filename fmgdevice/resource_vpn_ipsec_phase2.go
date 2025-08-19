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

func resourceVpnIpsecPhase2() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase2Create,
		Read:   resourceVpnIpsecPhase2Read,
		Update: resourceVpnIpsecPhase2Update,
		Delete: resourceVpnIpsecPhase2Delete,

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
			"addke1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke4": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke5": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"addke7": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Computed: true,
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
			"selector_match": &schema.Schema{
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
			"use_natip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecPhase2Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnIpsecPhase2(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2 resource while getting object: %v", err)
	}

	_, err = c.CreateVpnIpsecPhase2(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase2Read(d, m)
}

func resourceVpnIpsecPhase2Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnIpsecPhase2(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2 resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnIpsecPhase2(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceVpnIpsecPhase2Read(d, m)
}

func resourceVpnIpsecPhase2Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnIpsecPhase2(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase2Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecPhase2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2 resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase2AddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Addke1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Addke7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2AutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DhcpIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Dhgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Diffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Diffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstEndIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2DstName6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2DstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstStartIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2DstSubnet6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Encapsulation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InboundDscpCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InitiatorTsNarrow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Ipv4Df(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2KeylifeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keylifekbs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keylifeseconds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2L2Tp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Pfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Phase1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Replay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2RouteOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SelectorMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SingleSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcEndIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2SrcName6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2SrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcStartIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecPhase2SrcSubnet6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase2UseNatip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("add_route", flattenVpnIpsecPhase2AddRoute(o["add-route"], d, "add_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-route"], "VpnIpsecPhase2-AddRoute"); ok {
			if err = d.Set("add_route", vv); err != nil {
				return fmt.Errorf("Error reading add_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("addke1", flattenVpnIpsecPhase2Addke1(o["addke1"], d, "addke1")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke1"], "VpnIpsecPhase2-Addke1"); ok {
			if err = d.Set("addke1", vv); err != nil {
				return fmt.Errorf("Error reading addke1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke1: %v", err)
		}
	}

	if err = d.Set("addke2", flattenVpnIpsecPhase2Addke2(o["addke2"], d, "addke2")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke2"], "VpnIpsecPhase2-Addke2"); ok {
			if err = d.Set("addke2", vv); err != nil {
				return fmt.Errorf("Error reading addke2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke2: %v", err)
		}
	}

	if err = d.Set("addke3", flattenVpnIpsecPhase2Addke3(o["addke3"], d, "addke3")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke3"], "VpnIpsecPhase2-Addke3"); ok {
			if err = d.Set("addke3", vv); err != nil {
				return fmt.Errorf("Error reading addke3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke3: %v", err)
		}
	}

	if err = d.Set("addke4", flattenVpnIpsecPhase2Addke4(o["addke4"], d, "addke4")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke4"], "VpnIpsecPhase2-Addke4"); ok {
			if err = d.Set("addke4", vv); err != nil {
				return fmt.Errorf("Error reading addke4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke4: %v", err)
		}
	}

	if err = d.Set("addke5", flattenVpnIpsecPhase2Addke5(o["addke5"], d, "addke5")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke5"], "VpnIpsecPhase2-Addke5"); ok {
			if err = d.Set("addke5", vv); err != nil {
				return fmt.Errorf("Error reading addke5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke5: %v", err)
		}
	}

	if err = d.Set("addke6", flattenVpnIpsecPhase2Addke6(o["addke6"], d, "addke6")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke6"], "VpnIpsecPhase2-Addke6"); ok {
			if err = d.Set("addke6", vv); err != nil {
				return fmt.Errorf("Error reading addke6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke6: %v", err)
		}
	}

	if err = d.Set("addke7", flattenVpnIpsecPhase2Addke7(o["addke7"], d, "addke7")); err != nil {
		if vv, ok := fortiAPIPatch(o["addke7"], "VpnIpsecPhase2-Addke7"); ok {
			if err = d.Set("addke7", vv); err != nil {
				return fmt.Errorf("Error reading addke7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addke7: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase2AutoNegotiate(o["auto-negotiate"], d, "auto_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-negotiate"], "VpnIpsecPhase2-AutoNegotiate"); ok {
			if err = d.Set("auto_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading auto_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase2Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "VpnIpsecPhase2-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dhcp_ipsec", flattenVpnIpsecPhase2DhcpIpsec(o["dhcp-ipsec"], d, "dhcp_ipsec")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ipsec"], "VpnIpsecPhase2-DhcpIpsec"); ok {
			if err = d.Set("dhcp_ipsec", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase2Dhgrp(o["dhgrp"], d, "dhgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhgrp"], "VpnIpsecPhase2-Dhgrp"); ok {
			if err = d.Set("dhgrp", vv); err != nil {
				return fmt.Errorf("Error reading dhgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenVpnIpsecPhase2Diffserv(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "VpnIpsecPhase2-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenVpnIpsecPhase2Diffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "VpnIpsecPhase2-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dst_addr_type", flattenVpnIpsecPhase2DstAddrType(o["dst-addr-type"], d, "dst_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr-type"], "VpnIpsecPhase2-DstAddrType"); ok {
			if err = d.Set("dst_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr_type: %v", err)
		}
	}

	if err = d.Set("dst_end_ip", flattenVpnIpsecPhase2DstEndIp(o["dst-end-ip"], d, "dst_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-end-ip"], "VpnIpsecPhase2-DstEndIp"); ok {
			if err = d.Set("dst_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_end_ip: %v", err)
		}
	}

	if err = d.Set("dst_end_ip6", flattenVpnIpsecPhase2DstEndIp6(o["dst-end-ip6"], d, "dst_end_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-end-ip6"], "VpnIpsecPhase2-DstEndIp6"); ok {
			if err = d.Set("dst_end_ip6", vv); err != nil {
				return fmt.Errorf("Error reading dst_end_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_end_ip6: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenVpnIpsecPhase2DstName(o["dst-name"], d, "dst_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-name"], "VpnIpsecPhase2-DstName"); ok {
			if err = d.Set("dst_name", vv); err != nil {
				return fmt.Errorf("Error reading dst_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_name6", flattenVpnIpsecPhase2DstName6(o["dst-name6"], d, "dst_name6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-name6"], "VpnIpsecPhase2-DstName6"); ok {
			if err = d.Set("dst_name6", vv); err != nil {
				return fmt.Errorf("Error reading dst_name6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenVpnIpsecPhase2DstPort(o["dst-port"], d, "dst_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port"], "VpnIpsecPhase2-DstPort"); ok {
			if err = d.Set("dst_port", vv); err != nil {
				return fmt.Errorf("Error reading dst_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	if err = d.Set("dst_start_ip", flattenVpnIpsecPhase2DstStartIp(o["dst-start-ip"], d, "dst_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-start-ip"], "VpnIpsecPhase2-DstStartIp"); ok {
			if err = d.Set("dst_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_start_ip: %v", err)
		}
	}

	if err = d.Set("dst_start_ip6", flattenVpnIpsecPhase2DstStartIp6(o["dst-start-ip6"], d, "dst_start_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-start-ip6"], "VpnIpsecPhase2-DstStartIp6"); ok {
			if err = d.Set("dst_start_ip6", vv); err != nil {
				return fmt.Errorf("Error reading dst_start_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_start_ip6: %v", err)
		}
	}

	if err = d.Set("dst_subnet", flattenVpnIpsecPhase2DstSubnet(o["dst-subnet"], d, "dst_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-subnet"], "VpnIpsecPhase2-DstSubnet"); ok {
			if err = d.Set("dst_subnet", vv); err != nil {
				return fmt.Errorf("Error reading dst_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_subnet: %v", err)
		}
	}

	if err = d.Set("dst_subnet6", flattenVpnIpsecPhase2DstSubnet6(o["dst-subnet6"], d, "dst_subnet6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-subnet6"], "VpnIpsecPhase2-DstSubnet6"); ok {
			if err = d.Set("dst_subnet6", vv); err != nil {
				return fmt.Errorf("Error reading dst_subnet6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_subnet6: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase2Encapsulation(o["encapsulation"], d, "encapsulation")); err != nil {
		if vv, ok := fortiAPIPatch(o["encapsulation"], "VpnIpsecPhase2-Encapsulation"); ok {
			if err = d.Set("encapsulation", vv); err != nil {
				return fmt.Errorf("Error reading encapsulation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase2InboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "VpnIpsecPhase2-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("initiator_ts_narrow", flattenVpnIpsecPhase2InitiatorTsNarrow(o["initiator-ts-narrow"], d, "initiator_ts_narrow")); err != nil {
		if vv, ok := fortiAPIPatch(o["initiator-ts-narrow"], "VpnIpsecPhase2-InitiatorTsNarrow"); ok {
			if err = d.Set("initiator_ts_narrow", vv); err != nil {
				return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
		}
	}

	if err = d.Set("ipv4_df", flattenVpnIpsecPhase2Ipv4Df(o["ipv4-df"], d, "ipv4_df")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-df"], "VpnIpsecPhase2-Ipv4Df"); ok {
			if err = d.Set("ipv4_df", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_df: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_df: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase2Keepalive(o["keepalive"], d, "keepalive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive"], "VpnIpsecPhase2-Keepalive"); ok {
			if err = d.Set("keepalive", vv); err != nil {
				return fmt.Errorf("Error reading keepalive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("keylife_type", flattenVpnIpsecPhase2KeylifeType(o["keylife-type"], d, "keylife_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylife-type"], "VpnIpsecPhase2-KeylifeType"); ok {
			if err = d.Set("keylife_type", vv); err != nil {
				return fmt.Errorf("Error reading keylife_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylife_type: %v", err)
		}
	}

	if err = d.Set("keylifekbs", flattenVpnIpsecPhase2Keylifekbs(o["keylifekbs"], d, "keylifekbs")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylifekbs"], "VpnIpsecPhase2-Keylifekbs"); ok {
			if err = d.Set("keylifekbs", vv); err != nil {
				return fmt.Errorf("Error reading keylifekbs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylifekbs: %v", err)
		}
	}

	if err = d.Set("keylifeseconds", flattenVpnIpsecPhase2Keylifeseconds(o["keylifeseconds"], d, "keylifeseconds")); err != nil {
		if vv, ok := fortiAPIPatch(o["keylifeseconds"], "VpnIpsecPhase2-Keylifeseconds"); ok {
			if err = d.Set("keylifeseconds", vv); err != nil {
				return fmt.Errorf("Error reading keylifeseconds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keylifeseconds: %v", err)
		}
	}

	if err = d.Set("l2tp", flattenVpnIpsecPhase2L2Tp(o["l2tp"], d, "l2tp")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tp"], "VpnIpsecPhase2-L2Tp"); ok {
			if err = d.Set("l2tp", vv); err != nil {
				return fmt.Errorf("Error reading l2tp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tp: %v", err)
		}
	}

	if err = d.Set("name", flattenVpnIpsecPhase2Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "VpnIpsecPhase2-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pfs", flattenVpnIpsecPhase2Pfs(o["pfs"], d, "pfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfs"], "VpnIpsecPhase2-Pfs"); ok {
			if err = d.Set("pfs", vv); err != nil {
				return fmt.Errorf("Error reading pfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfs: %v", err)
		}
	}

	if err = d.Set("phase1name", flattenVpnIpsecPhase2Phase1Name(o["phase1name"], d, "phase1name")); err != nil {
		if vv, ok := fortiAPIPatch(o["phase1name"], "VpnIpsecPhase2-Phase1Name"); ok {
			if err = d.Set("phase1name", vv); err != nil {
				return fmt.Errorf("Error reading phase1name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phase1name: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase2Proposal(o["proposal"], d, "proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["proposal"], "VpnIpsecPhase2-Proposal"); ok {
			if err = d.Set("proposal", vv); err != nil {
				return fmt.Errorf("Error reading proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("protocol", flattenVpnIpsecPhase2Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "VpnIpsecPhase2-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("replay", flattenVpnIpsecPhase2Replay(o["replay"], d, "replay")); err != nil {
		if vv, ok := fortiAPIPatch(o["replay"], "VpnIpsecPhase2-Replay"); ok {
			if err = d.Set("replay", vv); err != nil {
				return fmt.Errorf("Error reading replay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replay: %v", err)
		}
	}

	if err = d.Set("route_overlap", flattenVpnIpsecPhase2RouteOverlap(o["route-overlap"], d, "route_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-overlap"], "VpnIpsecPhase2-RouteOverlap"); ok {
			if err = d.Set("route_overlap", vv); err != nil {
				return fmt.Errorf("Error reading route_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_overlap: %v", err)
		}
	}

	if err = d.Set("selector_match", flattenVpnIpsecPhase2SelectorMatch(o["selector-match"], d, "selector_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["selector-match"], "VpnIpsecPhase2-SelectorMatch"); ok {
			if err = d.Set("selector_match", vv); err != nil {
				return fmt.Errorf("Error reading selector_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selector_match: %v", err)
		}
	}

	if err = d.Set("single_source", flattenVpnIpsecPhase2SingleSource(o["single-source"], d, "single_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-source"], "VpnIpsecPhase2-SingleSource"); ok {
			if err = d.Set("single_source", vv); err != nil {
				return fmt.Errorf("Error reading single_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_source: %v", err)
		}
	}

	if err = d.Set("src_addr_type", flattenVpnIpsecPhase2SrcAddrType(o["src-addr-type"], d, "src_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-addr-type"], "VpnIpsecPhase2-SrcAddrType"); ok {
			if err = d.Set("src_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading src_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_addr_type: %v", err)
		}
	}

	if err = d.Set("src_end_ip", flattenVpnIpsecPhase2SrcEndIp(o["src-end-ip"], d, "src_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-end-ip"], "VpnIpsecPhase2-SrcEndIp"); ok {
			if err = d.Set("src_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_end_ip: %v", err)
		}
	}

	if err = d.Set("src_end_ip6", flattenVpnIpsecPhase2SrcEndIp6(o["src-end-ip6"], d, "src_end_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-end-ip6"], "VpnIpsecPhase2-SrcEndIp6"); ok {
			if err = d.Set("src_end_ip6", vv); err != nil {
				return fmt.Errorf("Error reading src_end_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_end_ip6: %v", err)
		}
	}

	if err = d.Set("src_name", flattenVpnIpsecPhase2SrcName(o["src-name"], d, "src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-name"], "VpnIpsecPhase2-SrcName"); ok {
			if err = d.Set("src_name", vv); err != nil {
				return fmt.Errorf("Error reading src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_name: %v", err)
		}
	}

	if err = d.Set("src_name6", flattenVpnIpsecPhase2SrcName6(o["src-name6"], d, "src_name6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-name6"], "VpnIpsecPhase2-SrcName6"); ok {
			if err = d.Set("src_name6", vv); err != nil {
				return fmt.Errorf("Error reading src_name6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_name6: %v", err)
		}
	}

	if err = d.Set("src_port", flattenVpnIpsecPhase2SrcPort(o["src-port"], d, "src_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-port"], "VpnIpsecPhase2-SrcPort"); ok {
			if err = d.Set("src_port", vv); err != nil {
				return fmt.Errorf("Error reading src_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_port: %v", err)
		}
	}

	if err = d.Set("src_start_ip", flattenVpnIpsecPhase2SrcStartIp(o["src-start-ip"], d, "src_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-start-ip"], "VpnIpsecPhase2-SrcStartIp"); ok {
			if err = d.Set("src_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_start_ip: %v", err)
		}
	}

	if err = d.Set("src_start_ip6", flattenVpnIpsecPhase2SrcStartIp6(o["src-start-ip6"], d, "src_start_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-start-ip6"], "VpnIpsecPhase2-SrcStartIp6"); ok {
			if err = d.Set("src_start_ip6", vv); err != nil {
				return fmt.Errorf("Error reading src_start_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_start_ip6: %v", err)
		}
	}

	if err = d.Set("src_subnet", flattenVpnIpsecPhase2SrcSubnet(o["src-subnet"], d, "src_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-subnet"], "VpnIpsecPhase2-SrcSubnet"); ok {
			if err = d.Set("src_subnet", vv); err != nil {
				return fmt.Errorf("Error reading src_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_subnet: %v", err)
		}
	}

	if err = d.Set("src_subnet6", flattenVpnIpsecPhase2SrcSubnet6(o["src-subnet6"], d, "src_subnet6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-subnet6"], "VpnIpsecPhase2-SrcSubnet6"); ok {
			if err = d.Set("src_subnet6", vv); err != nil {
				return fmt.Errorf("Error reading src_subnet6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_subnet6: %v", err)
		}
	}

	if err = d.Set("use_natip", flattenVpnIpsecPhase2UseNatip(o["use-natip"], d, "use_natip")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-natip"], "VpnIpsecPhase2-UseNatip"); ok {
			if err = d.Set("use_natip", vv); err != nil {
				return fmt.Errorf("Error reading use_natip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_natip: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecPhase2AddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Addke1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Addke7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2AutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DhcpIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Dhgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Diffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Diffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstEndIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2DstName6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2DstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstStartIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase2DstSubnet6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Encapsulation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InboundDscpCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InitiatorTsNarrow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Ipv4Df(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2KeylifeType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keylifekbs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keylifeseconds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2L2Tp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Pfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Phase1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Replay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2RouteOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SelectorMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SingleSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcEndIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2SrcName6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecPhase2SrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcStartIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnIpsecPhase2SrcSubnet6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2UseNatip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_route"); ok || d.HasChange("add_route") {
		t, err := expandVpnIpsecPhase2AddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("addke1"); ok || d.HasChange("addke1") {
		t, err := expandVpnIpsecPhase2Addke1(d, v, "addke1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke1"] = t
		}
	}

	if v, ok := d.GetOk("addke2"); ok || d.HasChange("addke2") {
		t, err := expandVpnIpsecPhase2Addke2(d, v, "addke2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke2"] = t
		}
	}

	if v, ok := d.GetOk("addke3"); ok || d.HasChange("addke3") {
		t, err := expandVpnIpsecPhase2Addke3(d, v, "addke3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke3"] = t
		}
	}

	if v, ok := d.GetOk("addke4"); ok || d.HasChange("addke4") {
		t, err := expandVpnIpsecPhase2Addke4(d, v, "addke4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke4"] = t
		}
	}

	if v, ok := d.GetOk("addke5"); ok || d.HasChange("addke5") {
		t, err := expandVpnIpsecPhase2Addke5(d, v, "addke5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke5"] = t
		}
	}

	if v, ok := d.GetOk("addke6"); ok || d.HasChange("addke6") {
		t, err := expandVpnIpsecPhase2Addke6(d, v, "addke6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke6"] = t
		}
	}

	if v, ok := d.GetOk("addke7"); ok || d.HasChange("addke7") {
		t, err := expandVpnIpsecPhase2Addke7(d, v, "addke7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke7"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok || d.HasChange("auto_negotiate") {
		t, err := expandVpnIpsecPhase2AutoNegotiate(d, v, "auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandVpnIpsecPhase2Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ipsec"); ok || d.HasChange("dhcp_ipsec") {
		t, err := expandVpnIpsecPhase2DhcpIpsec(d, v, "dhcp_ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok || d.HasChange("dhgrp") {
		t, err := expandVpnIpsecPhase2Dhgrp(d, v, "dhgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok || d.HasChange("diffserv") {
		t, err := expandVpnIpsecPhase2Diffserv(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandVpnIpsecPhase2Diffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr_type"); ok || d.HasChange("dst_addr_type") {
		t, err := expandVpnIpsecPhase2DstAddrType(d, v, "dst_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip"); ok || d.HasChange("dst_end_ip") {
		t, err := expandVpnIpsecPhase2DstEndIp(d, v, "dst_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip6"); ok || d.HasChange("dst_end_ip6") {
		t, err := expandVpnIpsecPhase2DstEndIp6(d, v, "dst_end_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok || d.HasChange("dst_name") {
		t, err := expandVpnIpsecPhase2DstName(d, v, "dst_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name"] = t
		}
	}

	if v, ok := d.GetOk("dst_name6"); ok || d.HasChange("dst_name6") {
		t, err := expandVpnIpsecPhase2DstName6(d, v, "dst_name6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name6"] = t
		}
	}

	if v, ok := d.GetOk("dst_port"); ok || d.HasChange("dst_port") {
		t, err := expandVpnIpsecPhase2DstPort(d, v, "dst_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip"); ok || d.HasChange("dst_start_ip") {
		t, err := expandVpnIpsecPhase2DstStartIp(d, v, "dst_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip6"); ok || d.HasChange("dst_start_ip6") {
		t, err := expandVpnIpsecPhase2DstStartIp6(d, v, "dst_start_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet"); ok || d.HasChange("dst_subnet") {
		t, err := expandVpnIpsecPhase2DstSubnet(d, v, "dst_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet6"); ok || d.HasChange("dst_subnet6") {
		t, err := expandVpnIpsecPhase2DstSubnet6(d, v, "dst_subnet6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet6"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok || d.HasChange("encapsulation") {
		t, err := expandVpnIpsecPhase2Encapsulation(d, v, "encapsulation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok || d.HasChange("inbound_dscp_copy") {
		t, err := expandVpnIpsecPhase2InboundDscpCopy(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("initiator_ts_narrow"); ok || d.HasChange("initiator_ts_narrow") {
		t, err := expandVpnIpsecPhase2InitiatorTsNarrow(d, v, "initiator_ts_narrow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator-ts-narrow"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_df"); ok || d.HasChange("ipv4_df") {
		t, err := expandVpnIpsecPhase2Ipv4Df(d, v, "ipv4_df")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-df"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok || d.HasChange("keepalive") {
		t, err := expandVpnIpsecPhase2Keepalive(d, v, "keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("keylife_type"); ok || d.HasChange("keylife_type") {
		t, err := expandVpnIpsecPhase2KeylifeType(d, v, "keylife_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife-type"] = t
		}
	}

	if v, ok := d.GetOk("keylifekbs"); ok || d.HasChange("keylifekbs") {
		t, err := expandVpnIpsecPhase2Keylifekbs(d, v, "keylifekbs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifekbs"] = t
		}
	}

	if v, ok := d.GetOk("keylifeseconds"); ok || d.HasChange("keylifeseconds") {
		t, err := expandVpnIpsecPhase2Keylifeseconds(d, v, "keylifeseconds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifeseconds"] = t
		}
	}

	if v, ok := d.GetOk("l2tp"); ok || d.HasChange("l2tp") {
		t, err := expandVpnIpsecPhase2L2Tp(d, v, "l2tp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandVpnIpsecPhase2Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pfs"); ok || d.HasChange("pfs") {
		t, err := expandVpnIpsecPhase2Pfs(d, v, "pfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfs"] = t
		}
	}

	if v, ok := d.GetOk("phase1name"); ok || d.HasChange("phase1name") {
		t, err := expandVpnIpsecPhase2Phase1Name(d, v, "phase1name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase1name"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok || d.HasChange("proposal") {
		t, err := expandVpnIpsecPhase2Proposal(d, v, "proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandVpnIpsecPhase2Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("replay"); ok || d.HasChange("replay") {
		t, err := expandVpnIpsecPhase2Replay(d, v, "replay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay"] = t
		}
	}

	if v, ok := d.GetOk("route_overlap"); ok || d.HasChange("route_overlap") {
		t, err := expandVpnIpsecPhase2RouteOverlap(d, v, "route_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-overlap"] = t
		}
	}

	if v, ok := d.GetOk("selector_match"); ok || d.HasChange("selector_match") {
		t, err := expandVpnIpsecPhase2SelectorMatch(d, v, "selector_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selector-match"] = t
		}
	}

	if v, ok := d.GetOk("single_source"); ok || d.HasChange("single_source") {
		t, err := expandVpnIpsecPhase2SingleSource(d, v, "single_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-source"] = t
		}
	}

	if v, ok := d.GetOk("src_addr_type"); ok || d.HasChange("src_addr_type") {
		t, err := expandVpnIpsecPhase2SrcAddrType(d, v, "src_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip"); ok || d.HasChange("src_end_ip") {
		t, err := expandVpnIpsecPhase2SrcEndIp(d, v, "src_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip6"); ok || d.HasChange("src_end_ip6") {
		t, err := expandVpnIpsecPhase2SrcEndIp6(d, v, "src_end_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_name"); ok || d.HasChange("src_name") {
		t, err := expandVpnIpsecPhase2SrcName(d, v, "src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name"] = t
		}
	}

	if v, ok := d.GetOk("src_name6"); ok || d.HasChange("src_name6") {
		t, err := expandVpnIpsecPhase2SrcName6(d, v, "src_name6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name6"] = t
		}
	}

	if v, ok := d.GetOk("src_port"); ok || d.HasChange("src_port") {
		t, err := expandVpnIpsecPhase2SrcPort(d, v, "src_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip"); ok || d.HasChange("src_start_ip") {
		t, err := expandVpnIpsecPhase2SrcStartIp(d, v, "src_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip6"); ok || d.HasChange("src_start_ip6") {
		t, err := expandVpnIpsecPhase2SrcStartIp6(d, v, "src_start_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet"); ok || d.HasChange("src_subnet") {
		t, err := expandVpnIpsecPhase2SrcSubnet(d, v, "src_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet6"); ok || d.HasChange("src_subnet6") {
		t, err := expandVpnIpsecPhase2SrcSubnet6(d, v, "src_subnet6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet6"] = t
		}
	}

	if v, ok := d.GetOk("use_natip"); ok || d.HasChange("use_natip") {
		t, err := expandVpnIpsecPhase2UseNatip(d, v, "use_natip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-natip"] = t
		}
	}

	return &obj, nil
}
