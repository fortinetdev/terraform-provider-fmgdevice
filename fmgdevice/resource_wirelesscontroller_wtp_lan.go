// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WTP LAN port mapping.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpLan() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpLanUpdate,
		Read:   resourceWirelessControllerWtpLanRead,
		Update: resourceWirelessControllerWtpLanUpdate,
		Delete: resourceWirelessControllerWtpLanDelete,

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
			"wtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_esl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_esl_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port1_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port1_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port2_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port2_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port3_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port3_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port4_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port4_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port5_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port5_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port6_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port6_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port7_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port7_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port8_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port8_ssid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpLanUpdate(d *schema.ResourceData, m interface{}) error {
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
	wtp := d.Get("wtp").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	obj, err := getObjectWirelessControllerWtpLan(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpLan resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpLan(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpLan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpLan")

	return resourceWirelessControllerWtpLanRead(d, m)
}

func resourceWirelessControllerWtpLanDelete(d *schema.ResourceData, m interface{}) error {
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
	wtp := d.Get("wtp").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	err = c.DeleteWirelessControllerWtpLan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpLan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpLanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	wtp := d.Get("wtp").(string)
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
	if wtp == "" {
		wtp = importOptionChecking(m.(*FortiClient).Cfg, "wtp")
		if wtp == "" {
			return fmt.Errorf("Parameter wtp is missing")
		}
		if err = d.Set("wtp", wtp); err != nil {
			return fmt.Errorf("Error set params wtp: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp"] = wtp

	o, err := c.ReadWirelessControllerWtpLan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpLan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpLan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpLan resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpLanPortEslMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortEslSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPortMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort1Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort2Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort3Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort4Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort5Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort6Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort7Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort8Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Ssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWirelessControllerWtpLan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port_esl_mode", flattenWirelessControllerWtpLanPortEslMode2edl(o["port-esl-mode"], d, "port_esl_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-esl-mode"], "WirelessControllerWtpLan-PortEslMode"); ok {
			if err = d.Set("port_esl_mode", vv); err != nil {
				return fmt.Errorf("Error reading port_esl_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_esl_mode: %v", err)
		}
	}

	if err = d.Set("port_esl_ssid", flattenWirelessControllerWtpLanPortEslSsid2edl(o["port-esl-ssid"], d, "port_esl_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-esl-ssid"], "WirelessControllerWtpLan-PortEslSsid"); ok {
			if err = d.Set("port_esl_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port_esl_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_esl_ssid: %v", err)
		}
	}

	if err = d.Set("port_mode", flattenWirelessControllerWtpLanPortMode2edl(o["port-mode"], d, "port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-mode"], "WirelessControllerWtpLan-PortMode"); ok {
			if err = d.Set("port_mode", vv); err != nil {
				return fmt.Errorf("Error reading port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_mode: %v", err)
		}
	}

	if err = d.Set("port_ssid", flattenWirelessControllerWtpLanPortSsid2edl(o["port-ssid"], d, "port_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-ssid"], "WirelessControllerWtpLan-PortSsid"); ok {
			if err = d.Set("port_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_ssid: %v", err)
		}
	}

	if err = d.Set("port1_mode", flattenWirelessControllerWtpLanPort1Mode2edl(o["port1-mode"], d, "port1_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port1-mode"], "WirelessControllerWtpLan-Port1Mode"); ok {
			if err = d.Set("port1_mode", vv); err != nil {
				return fmt.Errorf("Error reading port1_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port1_mode: %v", err)
		}
	}

	if err = d.Set("port1_ssid", flattenWirelessControllerWtpLanPort1Ssid2edl(o["port1-ssid"], d, "port1_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port1-ssid"], "WirelessControllerWtpLan-Port1Ssid"); ok {
			if err = d.Set("port1_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port1_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port1_ssid: %v", err)
		}
	}

	if err = d.Set("port2_mode", flattenWirelessControllerWtpLanPort2Mode2edl(o["port2-mode"], d, "port2_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2-mode"], "WirelessControllerWtpLan-Port2Mode"); ok {
			if err = d.Set("port2_mode", vv); err != nil {
				return fmt.Errorf("Error reading port2_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2_mode: %v", err)
		}
	}

	if err = d.Set("port2_ssid", flattenWirelessControllerWtpLanPort2Ssid2edl(o["port2-ssid"], d, "port2_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2-ssid"], "WirelessControllerWtpLan-Port2Ssid"); ok {
			if err = d.Set("port2_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port2_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2_ssid: %v", err)
		}
	}

	if err = d.Set("port3_mode", flattenWirelessControllerWtpLanPort3Mode2edl(o["port3-mode"], d, "port3_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3-mode"], "WirelessControllerWtpLan-Port3Mode"); ok {
			if err = d.Set("port3_mode", vv); err != nil {
				return fmt.Errorf("Error reading port3_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3_mode: %v", err)
		}
	}

	if err = d.Set("port3_ssid", flattenWirelessControllerWtpLanPort3Ssid2edl(o["port3-ssid"], d, "port3_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3-ssid"], "WirelessControllerWtpLan-Port3Ssid"); ok {
			if err = d.Set("port3_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port3_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3_ssid: %v", err)
		}
	}

	if err = d.Set("port4_mode", flattenWirelessControllerWtpLanPort4Mode2edl(o["port4-mode"], d, "port4_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4-mode"], "WirelessControllerWtpLan-Port4Mode"); ok {
			if err = d.Set("port4_mode", vv); err != nil {
				return fmt.Errorf("Error reading port4_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4_mode: %v", err)
		}
	}

	if err = d.Set("port4_ssid", flattenWirelessControllerWtpLanPort4Ssid2edl(o["port4-ssid"], d, "port4_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port4-ssid"], "WirelessControllerWtpLan-Port4Ssid"); ok {
			if err = d.Set("port4_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port4_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port4_ssid: %v", err)
		}
	}

	if err = d.Set("port5_mode", flattenWirelessControllerWtpLanPort5Mode2edl(o["port5-mode"], d, "port5_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5-mode"], "WirelessControllerWtpLan-Port5Mode"); ok {
			if err = d.Set("port5_mode", vv); err != nil {
				return fmt.Errorf("Error reading port5_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5_mode: %v", err)
		}
	}

	if err = d.Set("port5_ssid", flattenWirelessControllerWtpLanPort5Ssid2edl(o["port5-ssid"], d, "port5_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port5-ssid"], "WirelessControllerWtpLan-Port5Ssid"); ok {
			if err = d.Set("port5_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port5_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port5_ssid: %v", err)
		}
	}

	if err = d.Set("port6_mode", flattenWirelessControllerWtpLanPort6Mode2edl(o["port6-mode"], d, "port6_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port6-mode"], "WirelessControllerWtpLan-Port6Mode"); ok {
			if err = d.Set("port6_mode", vv); err != nil {
				return fmt.Errorf("Error reading port6_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port6_mode: %v", err)
		}
	}

	if err = d.Set("port6_ssid", flattenWirelessControllerWtpLanPort6Ssid2edl(o["port6-ssid"], d, "port6_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port6-ssid"], "WirelessControllerWtpLan-Port6Ssid"); ok {
			if err = d.Set("port6_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port6_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port6_ssid: %v", err)
		}
	}

	if err = d.Set("port7_mode", flattenWirelessControllerWtpLanPort7Mode2edl(o["port7-mode"], d, "port7_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port7-mode"], "WirelessControllerWtpLan-Port7Mode"); ok {
			if err = d.Set("port7_mode", vv); err != nil {
				return fmt.Errorf("Error reading port7_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port7_mode: %v", err)
		}
	}

	if err = d.Set("port7_ssid", flattenWirelessControllerWtpLanPort7Ssid2edl(o["port7-ssid"], d, "port7_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port7-ssid"], "WirelessControllerWtpLan-Port7Ssid"); ok {
			if err = d.Set("port7_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port7_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port7_ssid: %v", err)
		}
	}

	if err = d.Set("port8_mode", flattenWirelessControllerWtpLanPort8Mode2edl(o["port8-mode"], d, "port8_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["port8-mode"], "WirelessControllerWtpLan-Port8Mode"); ok {
			if err = d.Set("port8_mode", vv); err != nil {
				return fmt.Errorf("Error reading port8_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port8_mode: %v", err)
		}
	}

	if err = d.Set("port8_ssid", flattenWirelessControllerWtpLanPort8Ssid2edl(o["port8-ssid"], d, "port8_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["port8-ssid"], "WirelessControllerWtpLan-Port8Ssid"); ok {
			if err = d.Set("port8_ssid", vv); err != nil {
				return fmt.Errorf("Error reading port8_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port8_ssid: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpLanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpLanPortEslMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortEslSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPortMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort1Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort2Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort3Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort4Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort5Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort6Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort7Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort8Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Ssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWirelessControllerWtpLan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port_esl_mode"); ok || d.HasChange("port_esl_mode") {
		t, err := expandWirelessControllerWtpLanPortEslMode2edl(d, v, "port_esl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-esl-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_esl_ssid"); ok || d.HasChange("port_esl_ssid") {
		t, err := expandWirelessControllerWtpLanPortEslSsid2edl(d, v, "port_esl_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-esl-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port_mode"); ok || d.HasChange("port_mode") {
		t, err := expandWirelessControllerWtpLanPortMode2edl(d, v, "port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_ssid"); ok || d.HasChange("port_ssid") {
		t, err := expandWirelessControllerWtpLanPortSsid2edl(d, v, "port_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port1_mode"); ok || d.HasChange("port1_mode") {
		t, err := expandWirelessControllerWtpLanPort1Mode2edl(d, v, "port1_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port1-mode"] = t
		}
	}

	if v, ok := d.GetOk("port1_ssid"); ok || d.HasChange("port1_ssid") {
		t, err := expandWirelessControllerWtpLanPort1Ssid2edl(d, v, "port1_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port1-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port2_mode"); ok || d.HasChange("port2_mode") {
		t, err := expandWirelessControllerWtpLanPort2Mode2edl(d, v, "port2_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2-mode"] = t
		}
	}

	if v, ok := d.GetOk("port2_ssid"); ok || d.HasChange("port2_ssid") {
		t, err := expandWirelessControllerWtpLanPort2Ssid2edl(d, v, "port2_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port3_mode"); ok || d.HasChange("port3_mode") {
		t, err := expandWirelessControllerWtpLanPort3Mode2edl(d, v, "port3_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3-mode"] = t
		}
	}

	if v, ok := d.GetOk("port3_ssid"); ok || d.HasChange("port3_ssid") {
		t, err := expandWirelessControllerWtpLanPort3Ssid2edl(d, v, "port3_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port4_mode"); ok || d.HasChange("port4_mode") {
		t, err := expandWirelessControllerWtpLanPort4Mode2edl(d, v, "port4_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4-mode"] = t
		}
	}

	if v, ok := d.GetOk("port4_ssid"); ok || d.HasChange("port4_ssid") {
		t, err := expandWirelessControllerWtpLanPort4Ssid2edl(d, v, "port4_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port5_mode"); ok || d.HasChange("port5_mode") {
		t, err := expandWirelessControllerWtpLanPort5Mode2edl(d, v, "port5_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5-mode"] = t
		}
	}

	if v, ok := d.GetOk("port5_ssid"); ok || d.HasChange("port5_ssid") {
		t, err := expandWirelessControllerWtpLanPort5Ssid2edl(d, v, "port5_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port6_mode"); ok || d.HasChange("port6_mode") {
		t, err := expandWirelessControllerWtpLanPort6Mode2edl(d, v, "port6_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port6-mode"] = t
		}
	}

	if v, ok := d.GetOk("port6_ssid"); ok || d.HasChange("port6_ssid") {
		t, err := expandWirelessControllerWtpLanPort6Ssid2edl(d, v, "port6_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port6-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port7_mode"); ok || d.HasChange("port7_mode") {
		t, err := expandWirelessControllerWtpLanPort7Mode2edl(d, v, "port7_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port7-mode"] = t
		}
	}

	if v, ok := d.GetOk("port7_ssid"); ok || d.HasChange("port7_ssid") {
		t, err := expandWirelessControllerWtpLanPort7Ssid2edl(d, v, "port7_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port7-ssid"] = t
		}
	}

	if v, ok := d.GetOk("port8_mode"); ok || d.HasChange("port8_mode") {
		t, err := expandWirelessControllerWtpLanPort8Mode2edl(d, v, "port8_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port8-mode"] = t
		}
	}

	if v, ok := d.GetOk("port8_ssid"); ok || d.HasChange("port8_ssid") {
		t, err := expandWirelessControllerWtpLanPort8Ssid2edl(d, v, "port8_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port8-ssid"] = t
		}
	}

	return &obj, nil
}
