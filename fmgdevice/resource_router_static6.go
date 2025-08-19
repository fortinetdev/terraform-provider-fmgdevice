// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 static routing tables.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterStatic6Create,
		Read:   resourceRouterStatic6Read,
		Update: resourceRouterStatic6Update,
		Delete: resourceRouterStatic6Delete,

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
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"devindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_monitor_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan_zone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceRouterStatic6Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterStatic6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic6 resource while getting object: %v", err)
	}

	v, err := c.CreateRouterStatic6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic6 resource: %v", err)
	}

	if v != nil && v["seq-num"] != nil {
		if vidn, ok := v["seq-num"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceRouterStatic6Read(d, m)
		} else {
			return fmt.Errorf("Error creating RouterStatic6 resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterStatic6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterStatic6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterStatic6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterStatic6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStatic6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterStatic6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterStatic6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Blackhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterStatic6Devindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterStatic6DynamicGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6LinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Sdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6SdwanZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterStatic6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Tag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6VirtualWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Vrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Weight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterStatic6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bfd", flattenRouterStatic6Bfd(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterStatic6-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("blackhole", flattenRouterStatic6Blackhole(o["blackhole"], d, "blackhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["blackhole"], "RouterStatic6-Blackhole"); ok {
			if err = d.Set("blackhole", vv); err != nil {
				return fmt.Errorf("Error reading blackhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("comment", flattenRouterStatic6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "RouterStatic6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStatic6Device(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "RouterStatic6-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("devindex", flattenRouterStatic6Devindex(o["devindex"], d, "devindex")); err != nil {
		if vv, ok := fortiAPIPatch(o["devindex"], "RouterStatic6-Devindex"); ok {
			if err = d.Set("devindex", vv); err != nil {
				return fmt.Errorf("Error reading devindex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStatic6Distance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "RouterStatic6-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterStatic6Dst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "RouterStatic6-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenRouterStatic6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "RouterStatic6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", flattenRouterStatic6DynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-gateway"], "RouterStatic6-DynamicGateway"); ok {
			if err = d.Set("dynamic_gateway", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStatic6Gateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "RouterStatic6-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", flattenRouterStatic6LinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-monitor-exempt"], "RouterStatic6-LinkMonitorExempt"); ok {
			if err = d.Set("link_monitor_exempt", vv); err != nil {
				return fmt.Errorf("Error reading link_monitor_exempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterStatic6Priority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "RouterStatic6-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenRouterStatic6Sdwan(o["sdwan"], d, "sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan"], "RouterStatic6-Sdwan"); ok {
			if err = d.Set("sdwan", vv); err != nil {
				return fmt.Errorf("Error reading sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("sdwan_zone", flattenRouterStatic6SdwanZone(o["sdwan-zone"], d, "sdwan_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan-zone"], "RouterStatic6-SdwanZone"); ok {
			if err = d.Set("sdwan_zone", vv); err != nil {
				return fmt.Errorf("Error reading sdwan_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan_zone: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenRouterStatic6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq-num"], "RouterStatic6-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterStatic6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterStatic6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenRouterStatic6Tag(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "RouterStatic6-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", flattenRouterStatic6VirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-wan-link"], "RouterStatic6-VirtualWanLink"); ok {
			if err = d.Set("virtual_wan_link", vv); err != nil {
				return fmt.Errorf("Error reading virtual_wan_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("vrf", flattenRouterStatic6Vrf(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "RouterStatic6-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterStatic6Weight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "RouterStatic6-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenRouterStatic6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterStatic6Bfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Blackhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Device(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterStatic6Devindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterStatic6DynamicGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Gateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6LinkMonitorExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Sdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6SdwanZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterStatic6SeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Tag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6VirtualWanLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Vrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Weight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterStatic6Bfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("blackhole"); ok || d.HasChange("blackhole") {
		t, err := expandRouterStatic6Blackhole(d, v, "blackhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandRouterStatic6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandRouterStatic6Device(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("devindex"); ok || d.HasChange("devindex") {
		t, err := expandRouterStatic6Devindex(d, v, "devindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devindex"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandRouterStatic6Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandRouterStatic6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandRouterStatic6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_gateway"); ok || d.HasChange("dynamic_gateway") {
		t, err := expandRouterStatic6DynamicGateway(d, v, "dynamic_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandRouterStatic6Gateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("link_monitor_exempt"); ok || d.HasChange("link_monitor_exempt") {
		t, err := expandRouterStatic6LinkMonitorExempt(d, v, "link_monitor_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor-exempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandRouterStatic6Priority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok || d.HasChange("sdwan") {
		t, err := expandRouterStatic6Sdwan(d, v, "sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok || d.HasChange("sdwan_zone") {
		t, err := expandRouterStatic6SdwanZone(d, v, "sdwan_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-zone"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandRouterStatic6SeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterStatic6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandRouterStatic6Tag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("virtual_wan_link"); ok || d.HasChange("virtual_wan_link") {
		t, err := expandRouterStatic6VirtualWanLink(d, v, "virtual_wan_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandRouterStatic6Vrf(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandRouterStatic6Weight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
