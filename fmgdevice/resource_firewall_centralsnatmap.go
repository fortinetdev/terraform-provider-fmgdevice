// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 and IPv6 central SNAT policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallCentralSnatMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallCentralSnatMapCreate,
		Read:   resourceFirewallCentralSnatMapRead,
		Update: resourceFirewallCentralSnatMapUpdate,
		Delete: resourceFirewallCentralSnatMapDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_addr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_ippool": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat_ippool6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nat_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"orig_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"orig_addr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"orig_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"port_preserve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_random": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_addr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallCentralSnatMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallCentralSnatMap(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallCentralSnatMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallCentralSnatMap(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallCentralSnatMap(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallCentralSnatMap resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallCentralSnatMap(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallCentralSnatMap resource: %v", err)
		}

		if v != nil && v["policyid"] != nil {
			if vidn, ok := v["policyid"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallCentralSnatMapRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallCentralSnatMap resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallCentralSnatMapRead(d, m)
}

func resourceFirewallCentralSnatMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallCentralSnatMap(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallCentralSnatMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallCentralSnatMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallCentralSnatMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallCentralSnatMapRead(d, m)
}

func resourceFirewallCentralSnatMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallCentralSnatMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallCentralSnatMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallCentralSnatMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallCentralSnatMap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallCentralSnatMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallCentralSnatMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallCentralSnatMap resource from API: %v", err)
	}
	return nil
}

func flattenFirewallCentralSnatMapComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapDstAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapDstAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapDstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapNatIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapNatIppool6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapNatPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapOrigAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapOrigAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapOrigPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapPortPreserve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapPortRandom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallCentralSnatMapSrcAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallCentralSnatMapSrcAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallCentralSnatMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comments", flattenFirewallCentralSnatMapComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallCentralSnatMap-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dst_addr", flattenFirewallCentralSnatMapDstAddr(o["dst-addr"], d, "dst_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr"], "FirewallCentralSnatMap-DstAddr"); ok {
			if err = d.Set("dst_addr", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr: %v", err)
		}
	}

	if err = d.Set("dst_addr6", flattenFirewallCentralSnatMapDstAddr6(o["dst-addr6"], d, "dst_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-addr6"], "FirewallCentralSnatMap-DstAddr6"); ok {
			if err = d.Set("dst_addr6", vv); err != nil {
				return fmt.Errorf("Error reading dst_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_addr6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenFirewallCentralSnatMapDstPort(o["dst-port"], d, "dst_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port"], "FirewallCentralSnatMap-DstPort"); ok {
			if err = d.Set("dst_port", vv); err != nil {
				return fmt.Errorf("Error reading dst_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallCentralSnatMapDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallCentralSnatMap-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("nat", flattenFirewallCentralSnatMapNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "FirewallCentralSnatMap-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("nat_ippool", flattenFirewallCentralSnatMapNatIppool(o["nat-ippool"], d, "nat_ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-ippool"], "FirewallCentralSnatMap-NatIppool"); ok {
			if err = d.Set("nat_ippool", vv); err != nil {
				return fmt.Errorf("Error reading nat_ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_ippool: %v", err)
		}
	}

	if err = d.Set("nat_ippool6", flattenFirewallCentralSnatMapNatIppool6(o["nat-ippool6"], d, "nat_ippool6")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-ippool6"], "FirewallCentralSnatMap-NatIppool6"); ok {
			if err = d.Set("nat_ippool6", vv); err != nil {
				return fmt.Errorf("Error reading nat_ippool6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_ippool6: %v", err)
		}
	}

	if err = d.Set("nat_port", flattenFirewallCentralSnatMapNatPort(o["nat-port"], d, "nat_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-port"], "FirewallCentralSnatMap-NatPort"); ok {
			if err = d.Set("nat_port", vv); err != nil {
				return fmt.Errorf("Error reading nat_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_port: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallCentralSnatMapNat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "FirewallCentralSnatMap-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallCentralSnatMapNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "FirewallCentralSnatMap-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("orig_addr", flattenFirewallCentralSnatMapOrigAddr(o["orig-addr"], d, "orig_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-addr"], "FirewallCentralSnatMap-OrigAddr"); ok {
			if err = d.Set("orig_addr", vv); err != nil {
				return fmt.Errorf("Error reading orig_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_addr: %v", err)
		}
	}

	if err = d.Set("orig_addr6", flattenFirewallCentralSnatMapOrigAddr6(o["orig-addr6"], d, "orig_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-addr6"], "FirewallCentralSnatMap-OrigAddr6"); ok {
			if err = d.Set("orig_addr6", vv); err != nil {
				return fmt.Errorf("Error reading orig_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_addr6: %v", err)
		}
	}

	if err = d.Set("orig_port", flattenFirewallCentralSnatMapOrigPort(o["orig-port"], d, "orig_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["orig-port"], "FirewallCentralSnatMap-OrigPort"); ok {
			if err = d.Set("orig_port", vv); err != nil {
				return fmt.Errorf("Error reading orig_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading orig_port: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallCentralSnatMapPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallCentralSnatMap-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("port_preserve", flattenFirewallCentralSnatMapPortPreserve(o["port-preserve"], d, "port_preserve")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-preserve"], "FirewallCentralSnatMap-PortPreserve"); ok {
			if err = d.Set("port_preserve", vv); err != nil {
				return fmt.Errorf("Error reading port_preserve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_preserve: %v", err)
		}
	}

	if err = d.Set("port_random", flattenFirewallCentralSnatMapPortRandom(o["port-random"], d, "port_random")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-random"], "FirewallCentralSnatMap-PortRandom"); ok {
			if err = d.Set("port_random", vv); err != nil {
				return fmt.Errorf("Error reading port_random: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_random: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallCentralSnatMapProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallCentralSnatMap-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallCentralSnatMapSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallCentralSnatMap-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallCentralSnatMapStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallCentralSnatMap-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallCentralSnatMapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallCentralSnatMap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallCentralSnatMapUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallCentralSnatMap-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallCentralSnatMapAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallCentralSnatMap-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenFirewallCentralSnatMapIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "FirewallCentralSnatMap-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("src_addr", flattenFirewallCentralSnatMapSrcAddr(o["src-addr"], d, "src_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-addr"], "FirewallCentralSnatMap-SrcAddr"); ok {
			if err = d.Set("src_addr", vv); err != nil {
				return fmt.Errorf("Error reading src_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_addr: %v", err)
		}
	}

	if err = d.Set("src_addr6", flattenFirewallCentralSnatMapSrcAddr6(o["src-addr6"], d, "src_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-addr6"], "FirewallCentralSnatMap-SrcAddr6"); ok {
			if err = d.Set("src_addr6", vv); err != nil {
				return fmt.Errorf("Error reading src_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_addr6: %v", err)
		}
	}

	return nil
}

func flattenFirewallCentralSnatMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallCentralSnatMapComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapDstAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapDstAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapDstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapNatIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapNatIppool6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapNatPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapOrigAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapOrigAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapOrigPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapPortPreserve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapPortRandom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallCentralSnatMapSrcAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallCentralSnatMapSrcAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallCentralSnatMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallCentralSnatMapComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr"); ok || d.HasChange("dst_addr") {
		t, err := expandFirewallCentralSnatMapDstAddr(d, v, "dst_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr6"); ok || d.HasChange("dst_addr6") {
		t, err := expandFirewallCentralSnatMapDstAddr6(d, v, "dst_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr6"] = t
		}
	}

	if v, ok := d.GetOk("dst_port"); ok || d.HasChange("dst_port") {
		t, err := expandFirewallCentralSnatMapDstPort(d, v, "dst_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallCentralSnatMapDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandFirewallCentralSnatMapNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("nat_ippool"); ok || d.HasChange("nat_ippool") {
		t, err := expandFirewallCentralSnatMapNatIppool(d, v, "nat_ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-ippool"] = t
		}
	}

	if v, ok := d.GetOk("nat_ippool6"); ok || d.HasChange("nat_ippool6") {
		t, err := expandFirewallCentralSnatMapNatIppool6(d, v, "nat_ippool6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-ippool6"] = t
		}
	}

	if v, ok := d.GetOk("nat_port"); ok || d.HasChange("nat_port") {
		t, err := expandFirewallCentralSnatMapNatPort(d, v, "nat_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-port"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandFirewallCentralSnatMapNat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandFirewallCentralSnatMapNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("orig_addr"); ok || d.HasChange("orig_addr") {
		t, err := expandFirewallCentralSnatMapOrigAddr(d, v, "orig_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-addr"] = t
		}
	}

	if v, ok := d.GetOk("orig_addr6"); ok || d.HasChange("orig_addr6") {
		t, err := expandFirewallCentralSnatMapOrigAddr6(d, v, "orig_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-addr6"] = t
		}
	}

	if v, ok := d.GetOk("orig_port"); ok || d.HasChange("orig_port") {
		t, err := expandFirewallCentralSnatMapOrigPort(d, v, "orig_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["orig-port"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallCentralSnatMapPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("port_preserve"); ok || d.HasChange("port_preserve") {
		t, err := expandFirewallCentralSnatMapPortPreserve(d, v, "port_preserve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-preserve"] = t
		}
	}

	if v, ok := d.GetOk("port_random"); ok || d.HasChange("port_random") {
		t, err := expandFirewallCentralSnatMapPortRandom(d, v, "port_random")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-random"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallCentralSnatMapProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallCentralSnatMapSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallCentralSnatMapStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallCentralSnatMapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallCentralSnatMapUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallCentralSnatMapAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandFirewallCentralSnatMapIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("src_addr"); ok || d.HasChange("src_addr") {
		t, err := expandFirewallCentralSnatMapSrcAddr(d, v, "src_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr"] = t
		}
	}

	if v, ok := d.GetOk("src_addr6"); ok || d.HasChange("src_addr6") {
		t, err := expandFirewallCentralSnatMapSrcAddr6(d, v, "src_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr6"] = t
		}
	}

	return &obj, nil
}
