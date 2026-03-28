// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure custom services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallServiceCustomCreate,
		Read:   resourceFirewallServiceCustomRead,
		Update: resourceFirewallServiceCustomUpdate,
		Delete: resourceFirewallServiceCustomDelete,

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
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"app_service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"helper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udplite_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallServiceCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallServiceCustom resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallServiceCustom(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallServiceCustom(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallServiceCustom resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallServiceCustom(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallServiceCustom resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallServiceCustomRead(d, m)
}

func resourceFirewallServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallServiceCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCustom resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallServiceCustom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallServiceCustomRead(d, m)
}

func resourceFirewallServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallServiceCustom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallServiceCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallServiceCustom(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallServiceCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallServiceCustom resource from API: %v", err)
	}
	return nil
}

func flattenFirewallServiceCustomAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallServiceCustomAppServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallServiceCustomCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallServiceCustomCheckResetRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomHelper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomIcmpcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomIcmptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomIprange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomProtocolNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomSctpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallServiceCustomSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallServiceCustomTcpRstTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomUdpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallServiceCustomVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomUdplitePortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallServiceCustomUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallServiceCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("app_category", flattenFirewallServiceCustomAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "FirewallServiceCustom-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_service_type", flattenFirewallServiceCustomAppServiceType(o["app-service-type"], d, "app_service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-service-type"], "FirewallServiceCustom-AppServiceType"); ok {
			if err = d.Set("app_service_type", vv); err != nil {
				return fmt.Errorf("Error reading app_service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_service_type: %v", err)
		}
	}

	if err = d.Set("application", flattenFirewallServiceCustomApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallServiceCustom-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("category", flattenFirewallServiceCustomCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallServiceCustom-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("check_reset_range", flattenFirewallServiceCustomCheckResetRange(o["check-reset-range"], d, "check_reset_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-reset-range"], "FirewallServiceCustom-CheckResetRange"); ok {
			if err = d.Set("check_reset_range", vv); err != nil {
				return fmt.Errorf("Error reading check_reset_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallServiceCustomColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallServiceCustom-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallServiceCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallServiceCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallServiceCustomFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallServiceCustom-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallServiceCustomFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "FirewallServiceCustom-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallServiceCustomGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallServiceCustom-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("helper", flattenFirewallServiceCustomHelper(o["helper"], d, "helper")); err != nil {
		if vv, ok := fortiAPIPatch(o["helper"], "FirewallServiceCustom-Helper"); ok {
			if err = d.Set("helper", vv); err != nil {
				return fmt.Errorf("Error reading helper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading helper: %v", err)
		}
	}

	if err = d.Set("icmpcode", flattenFirewallServiceCustomIcmpcode(o["icmpcode"], d, "icmpcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmpcode"], "FirewallServiceCustom-Icmpcode"); ok {
			if err = d.Set("icmpcode", vv); err != nil {
				return fmt.Errorf("Error reading icmpcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmpcode: %v", err)
		}
	}

	if err = d.Set("icmptype", flattenFirewallServiceCustomIcmptype(o["icmptype"], d, "icmptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmptype"], "FirewallServiceCustom-Icmptype"); ok {
			if err = d.Set("icmptype", vv); err != nil {
				return fmt.Errorf("Error reading icmptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmptype: %v", err)
		}
	}

	if err = d.Set("iprange", flattenFirewallServiceCustomIprange(o["iprange"], d, "iprange")); err != nil {
		if vv, ok := fortiAPIPatch(o["iprange"], "FirewallServiceCustom-Iprange"); ok {
			if err = d.Set("iprange", vv); err != nil {
				return fmt.Errorf("Error reading iprange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iprange: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallServiceCustomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallServiceCustom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallServiceCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallServiceCustom-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("protocol_number", flattenFirewallServiceCustomProtocolNumber(o["protocol-number"], d, "protocol_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol-number"], "FirewallServiceCustom-ProtocolNumber"); ok {
			if err = d.Set("protocol_number", vv); err != nil {
				return fmt.Errorf("Error reading protocol_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol_number: %v", err)
		}
	}

	if err = d.Set("proxy", flattenFirewallServiceCustomProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "FirewallServiceCustom-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("sctp_portrange", flattenFirewallServiceCustomSctpPortrange(o["sctp-portrange"], d, "sctp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-portrange"], "FirewallServiceCustom-SctpPortrange"); ok {
			if err = d.Set("sctp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading sctp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_portrange: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenFirewallServiceCustomSessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "FirewallServiceCustom-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", flattenFirewallServiceCustomTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfclose-timer"], "FirewallServiceCustom-TcpHalfcloseTimer"); ok {
			if err = d.Set("tcp_halfclose_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", flattenFirewallServiceCustomTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfopen-timer"], "FirewallServiceCustom-TcpHalfopenTimer"); ok {
			if err = d.Set("tcp_halfopen_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_portrange", flattenFirewallServiceCustomTcpPortrange(o["tcp-portrange"], d, "tcp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-portrange"], "FirewallServiceCustom-TcpPortrange"); ok {
			if err = d.Set("tcp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading tcp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_portrange: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", flattenFirewallServiceCustomTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst-timer"], "FirewallServiceCustom-TcpRstTimer"); ok {
			if err = d.Set("tcp_rst_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", flattenFirewallServiceCustomTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timewait-timer"], "FirewallServiceCustom-TcpTimewaitTimer"); ok {
			if err = d.Set("tcp_timewait_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", flattenFirewallServiceCustomUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-idle-timer"], "FirewallServiceCustom-UdpIdleTimer"); ok {
			if err = d.Set("udp_idle_timer", vv); err != nil {
				return fmt.Errorf("Error reading udp_idle_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("udp_portrange", flattenFirewallServiceCustomUdpPortrange(o["udp-portrange"], d, "udp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-portrange"], "FirewallServiceCustom-UdpPortrange"); ok {
			if err = d.Set("udp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading udp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_portrange: %v", err)
		}
	}

	if err = d.Set("visibility", flattenFirewallServiceCustomVisibility(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "FirewallServiceCustom-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("udplite_portrange", flattenFirewallServiceCustomUdplitePortrange(o["udplite-portrange"], d, "udplite_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["udplite-portrange"], "FirewallServiceCustom-UdplitePortrange"); ok {
			if err = d.Set("udplite_portrange", vv); err != nil {
				return fmt.Errorf("Error reading udplite_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udplite_portrange: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallServiceCustomUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallServiceCustom-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallServiceCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallServiceCustomAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomAppServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomCheckResetRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomHelper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIcmpcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIcmptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIprange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProtocolNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomSctpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpHalfcloseTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpHalfopenTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomTcpRstTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpTimewaitTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUdpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallServiceCustomVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUdplitePortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallServiceCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallServiceCustomAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_service_type"); ok || d.HasChange("app_service_type") {
		t, err := expandFirewallServiceCustomAppServiceType(d, v, "app_service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-service-type"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallServiceCustomApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallServiceCustomCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("check_reset_range"); ok || d.HasChange("check_reset_range") {
		t, err := expandFirewallServiceCustomCheckResetRange(d, v, "check_reset_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-reset-range"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallServiceCustomColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallServiceCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallServiceCustomFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandFirewallServiceCustomFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallServiceCustomGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("helper"); ok || d.HasChange("helper") {
		t, err := expandFirewallServiceCustomHelper(d, v, "helper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["helper"] = t
		}
	}

	if v, ok := d.GetOk("icmpcode"); ok || d.HasChange("icmpcode") {
		t, err := expandFirewallServiceCustomIcmpcode(d, v, "icmpcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmpcode"] = t
		}
	}

	if v, ok := d.GetOk("icmptype"); ok || d.HasChange("icmptype") {
		t, err := expandFirewallServiceCustomIcmptype(d, v, "icmptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmptype"] = t
		}
	}

	if v, ok := d.GetOk("iprange"); ok || d.HasChange("iprange") {
		t, err := expandFirewallServiceCustomIprange(d, v, "iprange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iprange"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallServiceCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallServiceCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("protocol_number"); ok || d.HasChange("protocol_number") {
		t, err := expandFirewallServiceCustomProtocolNumber(d, v, "protocol_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-number"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok || d.HasChange("proxy") {
		t, err := expandFirewallServiceCustomProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("sctp_portrange"); ok || d.HasChange("sctp_portrange") {
		t, err := expandFirewallServiceCustomSctpPortrange(d, v, "sctp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandFirewallServiceCustomSessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfclose_timer"); ok || d.HasChange("tcp_halfclose_timer") {
		t, err := expandFirewallServiceCustomTcpHalfcloseTimer(d, v, "tcp_halfclose_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfclose-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfopen_timer"); ok || d.HasChange("tcp_halfopen_timer") {
		t, err := expandFirewallServiceCustomTcpHalfopenTimer(d, v, "tcp_halfopen_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_portrange"); ok || d.HasChange("tcp_portrange") {
		t, err := expandFirewallServiceCustomTcpPortrange(d, v, "tcp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timer"); ok || d.HasChange("tcp_rst_timer") {
		t, err := expandFirewallServiceCustomTcpRstTimer(d, v, "tcp_rst_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timewait_timer"); ok || d.HasChange("tcp_timewait_timer") {
		t, err := expandFirewallServiceCustomTcpTimewaitTimer(d, v, "tcp_timewait_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timewait-timer"] = t
		}
	}

	if v, ok := d.GetOk("udp_idle_timer"); ok || d.HasChange("udp_idle_timer") {
		t, err := expandFirewallServiceCustomUdpIdleTimer(d, v, "udp_idle_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("udp_portrange"); ok || d.HasChange("udp_portrange") {
		t, err := expandFirewallServiceCustomUdpPortrange(d, v, "udp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandFirewallServiceCustomVisibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("udplite_portrange"); ok || d.HasChange("udplite_portrange") {
		t, err := expandFirewallServiceCustomUdplitePortrange(d, v, "udplite_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udplite-portrange"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallServiceCustomUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
