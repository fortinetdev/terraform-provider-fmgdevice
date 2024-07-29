// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure L2TP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnL2Tp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnL2TpUpdate,
		Read:   resourceVpnL2TpRead,
		Update: resourceVpnL2TpUpdate,
		Delete: resourceVpnL2TpDelete,

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
			"compress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usrgrp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnL2TpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnL2Tp(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnL2Tp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnL2Tp")

	return resourceVpnL2TpRead(d, m)
}

func resourceVpnL2TpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnL2Tp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnL2Tp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnL2TpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnL2Tp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnL2Tp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource from API: %v", err)
	}
	return nil
}

func flattenVpnL2TpCompress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpEip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpEnforceIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpSip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpUsrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnL2Tp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("compress", flattenVpnL2TpCompress(o["compress"], d, "compress")); err != nil {
		if vv, ok := fortiAPIPatch(o["compress"], "VpnL2Tp-Compress"); ok {
			if err = d.Set("compress", vv); err != nil {
				return fmt.Errorf("Error reading compress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compress: %v", err)
		}
	}

	if err = d.Set("eip", flattenVpnL2TpEip(o["eip"], d, "eip")); err != nil {
		if vv, ok := fortiAPIPatch(o["eip"], "VpnL2Tp-Eip"); ok {
			if err = d.Set("eip", vv); err != nil {
				return fmt.Errorf("Error reading eip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("enforce_ipsec", flattenVpnL2TpEnforceIpsec(o["enforce-ipsec"], d, "enforce_ipsec")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-ipsec"], "VpnL2Tp-EnforceIpsec"); ok {
			if err = d.Set("enforce_ipsec", vv); err != nil {
				return fmt.Errorf("Error reading enforce_ipsec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_ipsec: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenVpnL2TpHelloInterval(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "VpnL2Tp-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenVpnL2TpLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-echo-interval"], "VpnL2Tp-LcpEchoInterval"); ok {
			if err = d.Set("lcp_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenVpnL2TpLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-max-echo-fails"], "VpnL2Tp-LcpMaxEchoFails"); ok {
			if err = d.Set("lcp_max_echo_fails", vv); err != nil {
				return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("sip", flattenVpnL2TpSip(o["sip"], d, "sip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip"], "VpnL2Tp-Sip"); ok {
			if err = d.Set("sip", vv); err != nil {
				return fmt.Errorf("Error reading sip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnL2TpStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VpnL2Tp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnL2TpUsrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["usrgrp"], "VpnL2Tp-Usrgrp"); ok {
			if err = d.Set("usrgrp", vv); err != nil {
				return fmt.Errorf("Error reading usrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	return nil
}

func flattenVpnL2TpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnL2TpCompress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpEip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpEnforceIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpSip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpUsrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnL2Tp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("compress"); ok || d.HasChange("compress") {
		t, err := expandVpnL2TpCompress(d, v, "compress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compress"] = t
		}
	}

	if v, ok := d.GetOk("eip"); ok || d.HasChange("eip") {
		t, err := expandVpnL2TpEip(d, v, "eip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip"] = t
		}
	}

	if v, ok := d.GetOk("enforce_ipsec"); ok || d.HasChange("enforce_ipsec") {
		t, err := expandVpnL2TpEnforceIpsec(d, v, "enforce_ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandVpnL2TpHelloInterval(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("lcp_echo_interval"); ok || d.HasChange("lcp_echo_interval") {
		t, err := expandVpnL2TpLcpEchoInterval(d, v, "lcp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("lcp_max_echo_fails"); ok || d.HasChange("lcp_max_echo_fails") {
		t, err := expandVpnL2TpLcpMaxEchoFails(d, v, "lcp_max_echo_fails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-max-echo-fails"] = t
		}
	}

	if v, ok := d.GetOk("sip"); ok || d.HasChange("sip") {
		t, err := expandVpnL2TpSip(d, v, "sip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVpnL2TpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok || d.HasChange("usrgrp") {
		t, err := expandVpnL2TpUsrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	return &obj, nil
}
