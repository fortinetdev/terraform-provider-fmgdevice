// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure on-demand packet sniffer.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallOnDemandSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallOnDemandSnifferCreate,
		Read:   resourceFirewallOnDemandSnifferRead,
		Update: resourceFirewallOnDemandSnifferUpdate,
		Delete: resourceFirewallOnDemandSnifferDelete,

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
			"advanced_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_packet_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_ip_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallOnDemandSnifferCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallOnDemandSniffer(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallOnDemandSniffer resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallOnDemandSniffer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating FirewallOnDemandSniffer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallOnDemandSnifferRead(d, m)
}

func resourceFirewallOnDemandSnifferUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallOnDemandSniffer(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallOnDemandSniffer resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallOnDemandSniffer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallOnDemandSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallOnDemandSnifferRead(d, m)
}

func resourceFirewallOnDemandSnifferDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallOnDemandSniffer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallOnDemandSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallOnDemandSnifferRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallOnDemandSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallOnDemandSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallOnDemandSniffer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallOnDemandSniffer resource from API: %v", err)
	}
	return nil
}

func flattenFirewallOnDemandSnifferAdvancedFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallOnDemandSnifferInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallOnDemandSnifferMaxPacketCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferNonIpPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallOnDemandSnifferProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectFirewallOnDemandSniffer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advanced_filter", flattenFirewallOnDemandSnifferAdvancedFilter(o["advanced-filter"], d, "advanced_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["advanced-filter"], "FirewallOnDemandSniffer-AdvancedFilter"); ok {
			if err = d.Set("advanced_filter", vv); err != nil {
				return fmt.Errorf("Error reading advanced_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advanced_filter: %v", err)
		}
	}

	if err = d.Set("hosts", flattenFirewallOnDemandSnifferHosts(o["hosts"], d, "hosts")); err != nil {
		if vv, ok := fortiAPIPatch(o["hosts"], "FirewallOnDemandSniffer-Hosts"); ok {
			if err = d.Set("hosts", vv); err != nil {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hosts: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallOnDemandSnifferInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallOnDemandSniffer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("max_packet_count", flattenFirewallOnDemandSnifferMaxPacketCount(o["max-packet-count"], d, "max_packet_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-packet-count"], "FirewallOnDemandSniffer-MaxPacketCount"); ok {
			if err = d.Set("max_packet_count", vv); err != nil {
				return fmt.Errorf("Error reading max_packet_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_packet_count: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallOnDemandSnifferName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallOnDemandSniffer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("non_ip_packet", flattenFirewallOnDemandSnifferNonIpPacket(o["non-ip-packet"], d, "non_ip_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["non-ip-packet"], "FirewallOnDemandSniffer-NonIpPacket"); ok {
			if err = d.Set("non_ip_packet", vv); err != nil {
				return fmt.Errorf("Error reading non_ip_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading non_ip_packet: %v", err)
		}
	}

	if err = d.Set("ports", flattenFirewallOnDemandSnifferPorts(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "FirewallOnDemandSniffer-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("protocols", flattenFirewallOnDemandSnifferProtocols(o["protocols"], d, "protocols")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocols"], "FirewallOnDemandSniffer-Protocols"); ok {
			if err = d.Set("protocols", vv); err != nil {
				return fmt.Errorf("Error reading protocols: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocols: %v", err)
		}
	}

	return nil
}

func flattenFirewallOnDemandSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallOnDemandSnifferAdvancedFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallOnDemandSnifferInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallOnDemandSnifferMaxPacketCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferNonIpPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallOnDemandSnifferProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectFirewallOnDemandSniffer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advanced_filter"); ok || d.HasChange("advanced_filter") {
		t, err := expandFirewallOnDemandSnifferAdvancedFilter(d, v, "advanced_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advanced-filter"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandFirewallOnDemandSnifferHosts(d, v, "hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallOnDemandSnifferInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("max_packet_count"); ok || d.HasChange("max_packet_count") {
		t, err := expandFirewallOnDemandSnifferMaxPacketCount(d, v, "max_packet_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-packet-count"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallOnDemandSnifferName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("non_ip_packet"); ok || d.HasChange("non_ip_packet") {
		t, err := expandFirewallOnDemandSnifferNonIpPacket(d, v, "non_ip_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["non-ip-packet"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallOnDemandSnifferPorts(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		t, err := expandFirewallOnDemandSnifferProtocols(d, v, "protocols")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	return &obj, nil
}
