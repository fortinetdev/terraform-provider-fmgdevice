// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure per-IP traffic shaper.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShaperPerIpShaperCreate,
		Read:   resourceFirewallShaperPerIpShaperRead,
		Update: resourceFirewallShaperPerIpShaperUpdate,
		Delete: resourceFirewallShaperPerIpShaperDelete,

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
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_tcp_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_udp_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallShaperPerIpShaperCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShaperPerIpShaper(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperPerIpShaper resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallShaperPerIpShaper(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallShaperPerIpShaper(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallShaperPerIpShaper resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallShaperPerIpShaper(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallShaperPerIpShaper resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallShaperPerIpShaperRead(d, m)
}

func resourceFirewallShaperPerIpShaperUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShaperPerIpShaper(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperPerIpShaper resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallShaperPerIpShaper(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperPerIpShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallShaperPerIpShaperRead(d, m)
}

func resourceFirewallShaperPerIpShaperDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallShaperPerIpShaper(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShaperPerIpShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperPerIpShaperRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallShaperPerIpShaper(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallShaperPerIpShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShaperPerIpShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperPerIpShaper resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShaperPerIpShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShaperPerIpShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bandwidth_unit", flattenFirewallShaperPerIpShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-unit"], "FirewallShaperPerIpShaper-BandwidthUnit"); ok {
			if err = d.Set("bandwidth_unit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallShaperPerIpShaperDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "FirewallShaperPerIpShaper-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallShaperPerIpShaperDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "FirewallShaperPerIpShaper-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallShaperPerIpShaperDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "FirewallShaperPerIpShaper-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallShaperPerIpShaperDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "FirewallShaperPerIpShaper-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("max_bandwidth", flattenFirewallShaperPerIpShaperMaxBandwidth(o["max-bandwidth"], d, "max_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-bandwidth"], "FirewallShaperPerIpShaper-MaxBandwidth"); ok {
			if err = d.Set("max_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading max_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_bandwidth: %v", err)
		}
	}

	if err = d.Set("max_concurrent_session", flattenFirewallShaperPerIpShaperMaxConcurrentSession(o["max-concurrent-session"], d, "max_concurrent_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-session"], "FirewallShaperPerIpShaper-MaxConcurrentSession"); ok {
			if err = d.Set("max_concurrent_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_tcp_session", flattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(o["max-concurrent-tcp-session"], d, "max_concurrent_tcp_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-tcp-session"], "FirewallShaperPerIpShaper-MaxConcurrentTcpSession"); ok {
			if err = d.Set("max_concurrent_tcp_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_tcp_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_tcp_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_udp_session", flattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(o["max-concurrent-udp-session"], d, "max_concurrent_udp_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-udp-session"], "FirewallShaperPerIpShaper-MaxConcurrentUdpSession"); ok {
			if err = d.Set("max_concurrent_udp_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_udp_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_udp_session: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallShaperPerIpShaperName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallShaperPerIpShaper-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenFirewallShaperPerIpShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShaperPerIpShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentTcpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentUdpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShaperPerIpShaper(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_unit"); ok || d.HasChange("bandwidth_unit") {
		t, err := expandFirewallShaperPerIpShaperBandwidthUnit(d, v, "bandwidth_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandFirewallShaperPerIpShaperDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandFirewallShaperPerIpShaperDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandFirewallShaperPerIpShaperDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandFirewallShaperPerIpShaperDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("max_bandwidth"); ok || d.HasChange("max_bandwidth") {
		t, err := expandFirewallShaperPerIpShaperMaxBandwidth(d, v, "max_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_session"); ok || d.HasChange("max_concurrent_session") {
		t, err := expandFirewallShaperPerIpShaperMaxConcurrentSession(d, v, "max_concurrent_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-session"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_tcp_session"); ok || d.HasChange("max_concurrent_tcp_session") {
		t, err := expandFirewallShaperPerIpShaperMaxConcurrentTcpSession(d, v, "max_concurrent_tcp_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-tcp-session"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_udp_session"); ok || d.HasChange("max_concurrent_udp_session") {
		t, err := expandFirewallShaperPerIpShaperMaxConcurrentUdpSession(d, v, "max_concurrent_udp_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-udp-session"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallShaperPerIpShaperName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
