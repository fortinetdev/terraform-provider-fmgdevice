// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless controller event log filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerLogUpdate,
		Read:   resourceWirelessControllerLogRead,
		Update: resourceWirelessControllerLogUpdate,
		Delete: resourceWirelessControllerLogDelete,

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
			"addrgrp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ble_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clb_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_starv_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_sched_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rogue_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sta_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sta_locate_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wids_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_fips_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerLogUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLog(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLog resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerLog(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerLog")

	return resourceWirelessControllerLogRead(d, m)
}

func resourceWirelessControllerLogDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerLog(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerLogRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerLog(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerLog(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLog resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerLogAddrgrpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogBleLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogClbLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogDhcpStarvLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogLedSchedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogRadioEventLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogRogueEventLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogStaEventLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogStaLocateLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogWidsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogWtpEventLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerLogWtpFipsEventLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerLog(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addrgrp_log", flattenWirelessControllerLogAddrgrpLog(o["addrgrp-log"], d, "addrgrp_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["addrgrp-log"], "WirelessControllerLog-AddrgrpLog"); ok {
			if err = d.Set("addrgrp_log", vv); err != nil {
				return fmt.Errorf("Error reading addrgrp_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addrgrp_log: %v", err)
		}
	}

	if err = d.Set("ble_log", flattenWirelessControllerLogBleLog(o["ble-log"], d, "ble_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-log"], "WirelessControllerLog-BleLog"); ok {
			if err = d.Set("ble_log", vv); err != nil {
				return fmt.Errorf("Error reading ble_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_log: %v", err)
		}
	}

	if err = d.Set("clb_log", flattenWirelessControllerLogClbLog(o["clb-log"], d, "clb_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["clb-log"], "WirelessControllerLog-ClbLog"); ok {
			if err = d.Set("clb_log", vv); err != nil {
				return fmt.Errorf("Error reading clb_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clb_log: %v", err)
		}
	}

	if err = d.Set("dhcp_starv_log", flattenWirelessControllerLogDhcpStarvLog(o["dhcp-starv-log"], d, "dhcp_starv_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-starv-log"], "WirelessControllerLog-DhcpStarvLog"); ok {
			if err = d.Set("dhcp_starv_log", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_starv_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_starv_log: %v", err)
		}
	}

	if err = d.Set("led_sched_log", flattenWirelessControllerLogLedSchedLog(o["led-sched-log"], d, "led_sched_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["led-sched-log"], "WirelessControllerLog-LedSchedLog"); ok {
			if err = d.Set("led_sched_log", vv); err != nil {
				return fmt.Errorf("Error reading led_sched_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading led_sched_log: %v", err)
		}
	}

	if err = d.Set("radio_event_log", flattenWirelessControllerLogRadioEventLog(o["radio-event-log"], d, "radio_event_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-event-log"], "WirelessControllerLog-RadioEventLog"); ok {
			if err = d.Set("radio_event_log", vv); err != nil {
				return fmt.Errorf("Error reading radio_event_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_event_log: %v", err)
		}
	}

	if err = d.Set("rogue_event_log", flattenWirelessControllerLogRogueEventLog(o["rogue-event-log"], d, "rogue_event_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["rogue-event-log"], "WirelessControllerLog-RogueEventLog"); ok {
			if err = d.Set("rogue_event_log", vv); err != nil {
				return fmt.Errorf("Error reading rogue_event_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rogue_event_log: %v", err)
		}
	}

	if err = d.Set("sta_event_log", flattenWirelessControllerLogStaEventLog(o["sta-event-log"], d, "sta_event_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-event-log"], "WirelessControllerLog-StaEventLog"); ok {
			if err = d.Set("sta_event_log", vv); err != nil {
				return fmt.Errorf("Error reading sta_event_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_event_log: %v", err)
		}
	}

	if err = d.Set("sta_locate_log", flattenWirelessControllerLogStaLocateLog(o["sta-locate-log"], d, "sta_locate_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-locate-log"], "WirelessControllerLog-StaLocateLog"); ok {
			if err = d.Set("sta_locate_log", vv); err != nil {
				return fmt.Errorf("Error reading sta_locate_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_locate_log: %v", err)
		}
	}

	if err = d.Set("status", flattenWirelessControllerLogStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WirelessControllerLog-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("wids_log", flattenWirelessControllerLogWidsLog(o["wids-log"], d, "wids_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wids-log"], "WirelessControllerLog-WidsLog"); ok {
			if err = d.Set("wids_log", vv); err != nil {
				return fmt.Errorf("Error reading wids_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wids_log: %v", err)
		}
	}

	if err = d.Set("wtp_event_log", flattenWirelessControllerLogWtpEventLog(o["wtp-event-log"], d, "wtp_event_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-event-log"], "WirelessControllerLog-WtpEventLog"); ok {
			if err = d.Set("wtp_event_log", vv); err != nil {
				return fmt.Errorf("Error reading wtp_event_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_event_log: %v", err)
		}
	}

	if err = d.Set("wtp_fips_event_log", flattenWirelessControllerLogWtpFipsEventLog(o["wtp-fips-event-log"], d, "wtp_fips_event_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-fips-event-log"], "WirelessControllerLog-WtpFipsEventLog"); ok {
			if err = d.Set("wtp_fips_event_log", vv); err != nil {
				return fmt.Errorf("Error reading wtp_fips_event_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_fips_event_log: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerLogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerLogAddrgrpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogBleLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogClbLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogDhcpStarvLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogLedSchedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogRadioEventLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogRogueEventLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogStaEventLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogStaLocateLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWidsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWtpEventLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWtpFipsEventLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerLog(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addrgrp_log"); ok || d.HasChange("addrgrp_log") {
		t, err := expandWirelessControllerLogAddrgrpLog(d, v, "addrgrp_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addrgrp-log"] = t
		}
	}

	if v, ok := d.GetOk("ble_log"); ok || d.HasChange("ble_log") {
		t, err := expandWirelessControllerLogBleLog(d, v, "ble_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-log"] = t
		}
	}

	if v, ok := d.GetOk("clb_log"); ok || d.HasChange("clb_log") {
		t, err := expandWirelessControllerLogClbLog(d, v, "clb_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clb-log"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_starv_log"); ok || d.HasChange("dhcp_starv_log") {
		t, err := expandWirelessControllerLogDhcpStarvLog(d, v, "dhcp_starv_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-starv-log"] = t
		}
	}

	if v, ok := d.GetOk("led_sched_log"); ok || d.HasChange("led_sched_log") {
		t, err := expandWirelessControllerLogLedSchedLog(d, v, "led_sched_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-sched-log"] = t
		}
	}

	if v, ok := d.GetOk("radio_event_log"); ok || d.HasChange("radio_event_log") {
		t, err := expandWirelessControllerLogRadioEventLog(d, v, "radio_event_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-event-log"] = t
		}
	}

	if v, ok := d.GetOk("rogue_event_log"); ok || d.HasChange("rogue_event_log") {
		t, err := expandWirelessControllerLogRogueEventLog(d, v, "rogue_event_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-event-log"] = t
		}
	}

	if v, ok := d.GetOk("sta_event_log"); ok || d.HasChange("sta_event_log") {
		t, err := expandWirelessControllerLogStaEventLog(d, v, "sta_event_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-event-log"] = t
		}
	}

	if v, ok := d.GetOk("sta_locate_log"); ok || d.HasChange("sta_locate_log") {
		t, err := expandWirelessControllerLogStaLocateLog(d, v, "sta_locate_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-locate-log"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWirelessControllerLogStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("wids_log"); ok || d.HasChange("wids_log") {
		t, err := expandWirelessControllerLogWidsLog(d, v, "wids_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wids-log"] = t
		}
	}

	if v, ok := d.GetOk("wtp_event_log"); ok || d.HasChange("wtp_event_log") {
		t, err := expandWirelessControllerLogWtpEventLog(d, v, "wtp_event_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-event-log"] = t
		}
	}

	if v, ok := d.GetOk("wtp_fips_event_log"); ok || d.HasChange("wtp_fips_event_log") {
		t, err := expandWirelessControllerLogWtpFipsEventLog(d, v, "wtp_fips_event_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-fips-event-log"] = t
		}
	}

	return &obj, nil
}
