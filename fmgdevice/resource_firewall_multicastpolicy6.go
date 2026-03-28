// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 multicast NAT policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallMulticastPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallMulticastPolicy6Create,
		Read:   resourceFirewallMulticastPolicy6Read,
		Update: resourceFirewallMulticastPolicy6Update,
		Delete: resourceFirewallMulticastPolicy6Delete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallMulticastPolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallMulticastPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallMulticastPolicy6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallMulticastPolicy6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallMulticastPolicy6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallMulticastPolicy6 resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallMulticastPolicy6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallMulticastPolicy6 resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallMulticastPolicy6Read(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallMulticastPolicy6 resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallMulticastPolicy6Read(d, m)
}

func resourceFirewallMulticastPolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallMulticastPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastPolicy6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallMulticastPolicy6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallMulticastPolicy6Read(d, m)
}

func resourceFirewallMulticastPolicy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallMulticastPolicy6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallMulticastPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastPolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallMulticastPolicy6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallMulticastPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallMulticastPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallMulticastPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallMulticastPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6AutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastPolicy6Dstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastPolicy6EndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastPolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastPolicy6Srcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastPolicy6StartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6UtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastPolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallMulticastPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenFirewallMulticastPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "FirewallMulticastPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenFirewallMulticastPolicy6AutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "FirewallMulticastPolicy6-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallMulticastPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallMulticastPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallMulticastPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallMulticastPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallMulticastPolicy6Dstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallMulticastPolicy6-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("end_port", flattenFirewallMulticastPolicy6EndPort(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "FirewallMulticastPolicy6-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallMulticastPolicy6Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallMulticastPolicy6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallMulticastPolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallMulticastPolicy6-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallMulticastPolicy6Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallMulticastPolicy6-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallMulticastPolicy6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallMulticastPolicy6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallMulticastPolicy6Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallMulticastPolicy6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallMulticastPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallMulticastPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallMulticastPolicy6Srcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallMulticastPolicy6-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("start_port", flattenFirewallMulticastPolicy6StartPort(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "FirewallMulticastPolicy6-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallMulticastPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallMulticastPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallMulticastPolicy6UtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "FirewallMulticastPolicy6-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallMulticastPolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallMulticastPolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallMulticastPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallMulticastPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6AutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastPolicy6Dstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastPolicy6EndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastPolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastPolicy6Srcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastPolicy6StartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6UtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastPolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallMulticastPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandFirewallMulticastPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandFirewallMulticastPolicy6AutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallMulticastPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallMulticastPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallMulticastPolicy6Dstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandFirewallMulticastPolicy6EndPort(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallMulticastPolicy6Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallMulticastPolicy6IpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallMulticastPolicy6Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallMulticastPolicy6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallMulticastPolicy6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallMulticastPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallMulticastPolicy6Srcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandFirewallMulticastPolicy6StartPort(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallMulticastPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok || d.HasChange("utm_status") {
		t, err := expandFirewallMulticastPolicy6UtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallMulticastPolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
