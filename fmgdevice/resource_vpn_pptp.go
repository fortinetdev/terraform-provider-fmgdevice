// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure PPTP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnPptp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnPptpUpdate,
		Read:   resourceVpnPptpRead,
		Update: resourceVpnPptpUpdate,
		Delete: resourceVpnPptpDelete,

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
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceVpnPptpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnPptp(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnPptp resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnPptp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnPptp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnPptp")

	return resourceVpnPptpRead(d, m)
}

func resourceVpnPptpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnPptp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnPptp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnPptpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnPptp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnPptp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnPptp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnPptp resource from API: %v", err)
	}
	return nil
}

func flattenVpnPptpEip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnPptpIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnPptpLocalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnPptpSip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnPptpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnPptpUsrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnPptp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("eip", flattenVpnPptpEip(o["eip"], d, "eip")); err != nil {
		if vv, ok := fortiAPIPatch(o["eip"], "VpnPptp-Eip"); ok {
			if err = d.Set("eip", vv); err != nil {
				return fmt.Errorf("Error reading eip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenVpnPptpIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "VpnPptp-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("local_ip", flattenVpnPptpLocalIp(o["local-ip"], d, "local_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-ip"], "VpnPptp-LocalIp"); ok {
			if err = d.Set("local_ip", vv); err != nil {
				return fmt.Errorf("Error reading local_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_ip: %v", err)
		}
	}

	if err = d.Set("sip", flattenVpnPptpSip(o["sip"], d, "sip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sip"], "VpnPptp-Sip"); ok {
			if err = d.Set("sip", vv); err != nil {
				return fmt.Errorf("Error reading sip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sip: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnPptpStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VpnPptp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnPptpUsrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["usrgrp"], "VpnPptp-Usrgrp"); ok {
			if err = d.Set("usrgrp", vv); err != nil {
				return fmt.Errorf("Error reading usrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	return nil
}

func flattenVpnPptpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnPptpEip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpLocalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpSip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpUsrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnPptp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eip"); ok || d.HasChange("eip") {
		t, err := expandVpnPptpEip(d, v, "eip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandVpnPptpIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("local_ip"); ok || d.HasChange("local_ip") {
		t, err := expandVpnPptpLocalIp(d, v, "local_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-ip"] = t
		}
	}

	if v, ok := d.GetOk("sip"); ok || d.HasChange("sip") {
		t, err := expandVpnPptpSip(d, v, "sip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVpnPptpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok || d.HasChange("usrgrp") {
		t, err := expandVpnPptpUsrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	return &obj, nil
}
