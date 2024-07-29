// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CAPWAP timers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerTimers() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerTimersUpdate,
		Read:   resourceWirelessControllerTimersRead,
		Update: resourceWirelessControllerTimersUpdate,
		Delete: resourceWirelessControllerTimersDelete,

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
			"ap_reboot_wait_interval1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_reboot_wait_interval2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_reboot_wait_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ble_device_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ble_scan_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client_idle_rehome_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"discovery_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"drma_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fake_ap_log": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_intf_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nat_session_keep_alive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"radio_stats_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rogue_ap_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rogue_ap_log": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rogue_sta_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sta_cap_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sta_capability_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sta_locate_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sta_stats_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vap_stats_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerTimersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectWirelessControllerTimers(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerTimers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerTimers")

	return resourceWirelessControllerTimersRead(d, m)
}

func resourceWirelessControllerTimersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteWirelessControllerTimers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerTimers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerTimersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadWirelessControllerTimers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerTimers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerTimersApRebootWaitInterval1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersApRebootWaitInterval2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersApRebootWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersBleDeviceCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersBleScanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersClientIdleRehomeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersClientIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDiscoveryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDrmaInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersFakeApLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersIpsecIntfCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersNatSessionKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRadioStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueApCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueApLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueStaCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaCapCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaCapabilityInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaLocateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersVapStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerTimers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ap_reboot_wait_interval1", flattenWirelessControllerTimersApRebootWaitInterval1(o["ap-reboot-wait-interval1"], d, "ap_reboot_wait_interval1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-reboot-wait-interval1"], "WirelessControllerTimers-ApRebootWaitInterval1"); ok {
			if err = d.Set("ap_reboot_wait_interval1", vv); err != nil {
				return fmt.Errorf("Error reading ap_reboot_wait_interval1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_reboot_wait_interval1: %v", err)
		}
	}

	if err = d.Set("ap_reboot_wait_interval2", flattenWirelessControllerTimersApRebootWaitInterval2(o["ap-reboot-wait-interval2"], d, "ap_reboot_wait_interval2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-reboot-wait-interval2"], "WirelessControllerTimers-ApRebootWaitInterval2"); ok {
			if err = d.Set("ap_reboot_wait_interval2", vv); err != nil {
				return fmt.Errorf("Error reading ap_reboot_wait_interval2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_reboot_wait_interval2: %v", err)
		}
	}

	if err = d.Set("ap_reboot_wait_time", flattenWirelessControllerTimersApRebootWaitTime(o["ap-reboot-wait-time"], d, "ap_reboot_wait_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-reboot-wait-time"], "WirelessControllerTimers-ApRebootWaitTime"); ok {
			if err = d.Set("ap_reboot_wait_time", vv); err != nil {
				return fmt.Errorf("Error reading ap_reboot_wait_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_reboot_wait_time: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenWirelessControllerTimersAuthTimeout(o["auth-timeout"], d, "auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-timeout"], "WirelessControllerTimers-AuthTimeout"); ok {
			if err = d.Set("auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("ble_device_cleanup", flattenWirelessControllerTimersBleDeviceCleanup(o["ble-device-cleanup"], d, "ble_device_cleanup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-device-cleanup"], "WirelessControllerTimers-BleDeviceCleanup"); ok {
			if err = d.Set("ble_device_cleanup", vv); err != nil {
				return fmt.Errorf("Error reading ble_device_cleanup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_device_cleanup: %v", err)
		}
	}

	if err = d.Set("ble_scan_report_intv", flattenWirelessControllerTimersBleScanReportIntv(o["ble-scan-report-intv"], d, "ble_scan_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-scan-report-intv"], "WirelessControllerTimers-BleScanReportIntv"); ok {
			if err = d.Set("ble_scan_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading ble_scan_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_scan_report_intv: %v", err)
		}
	}

	if err = d.Set("client_idle_rehome_timeout", flattenWirelessControllerTimersClientIdleRehomeTimeout(o["client-idle-rehome-timeout"], d, "client_idle_rehome_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-idle-rehome-timeout"], "WirelessControllerTimers-ClientIdleRehomeTimeout"); ok {
			if err = d.Set("client_idle_rehome_timeout", vv); err != nil {
				return fmt.Errorf("Error reading client_idle_rehome_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_idle_rehome_timeout: %v", err)
		}
	}

	if err = d.Set("client_idle_timeout", flattenWirelessControllerTimersClientIdleTimeout(o["client-idle-timeout"], d, "client_idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-idle-timeout"], "WirelessControllerTimers-ClientIdleTimeout"); ok {
			if err = d.Set("client_idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading client_idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_idle_timeout: %v", err)
		}
	}

	if err = d.Set("discovery_interval", flattenWirelessControllerTimersDiscoveryInterval(o["discovery-interval"], d, "discovery_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["discovery-interval"], "WirelessControllerTimers-DiscoveryInterval"); ok {
			if err = d.Set("discovery_interval", vv); err != nil {
				return fmt.Errorf("Error reading discovery_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discovery_interval: %v", err)
		}
	}

	if err = d.Set("drma_interval", flattenWirelessControllerTimersDrmaInterval(o["drma-interval"], d, "drma_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["drma-interval"], "WirelessControllerTimers-DrmaInterval"); ok {
			if err = d.Set("drma_interval", vv); err != nil {
				return fmt.Errorf("Error reading drma_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drma_interval: %v", err)
		}
	}

	if err = d.Set("echo_interval", flattenWirelessControllerTimersEchoInterval(o["echo-interval"], d, "echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-interval"], "WirelessControllerTimers-EchoInterval"); ok {
			if err = d.Set("echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_interval: %v", err)
		}
	}

	if err = d.Set("fake_ap_log", flattenWirelessControllerTimersFakeApLog(o["fake-ap-log"], d, "fake_ap_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fake-ap-log"], "WirelessControllerTimers-FakeApLog"); ok {
			if err = d.Set("fake_ap_log", vv); err != nil {
				return fmt.Errorf("Error reading fake_ap_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fake_ap_log: %v", err)
		}
	}

	if err = d.Set("ipsec_intf_cleanup", flattenWirelessControllerTimersIpsecIntfCleanup(o["ipsec-intf-cleanup"], d, "ipsec_intf_cleanup")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-intf-cleanup"], "WirelessControllerTimers-IpsecIntfCleanup"); ok {
			if err = d.Set("ipsec_intf_cleanup", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_intf_cleanup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_intf_cleanup: %v", err)
		}
	}

	if err = d.Set("nat_session_keep_alive", flattenWirelessControllerTimersNatSessionKeepAlive(o["nat-session-keep-alive"], d, "nat_session_keep_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-session-keep-alive"], "WirelessControllerTimers-NatSessionKeepAlive"); ok {
			if err = d.Set("nat_session_keep_alive", vv); err != nil {
				return fmt.Errorf("Error reading nat_session_keep_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_session_keep_alive: %v", err)
		}
	}

	if err = d.Set("radio_stats_interval", flattenWirelessControllerTimersRadioStatsInterval(o["radio-stats-interval"], d, "radio_stats_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-stats-interval"], "WirelessControllerTimers-RadioStatsInterval"); ok {
			if err = d.Set("radio_stats_interval", vv); err != nil {
				return fmt.Errorf("Error reading radio_stats_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_stats_interval: %v", err)
		}
	}

	if err = d.Set("rogue_ap_cleanup", flattenWirelessControllerTimersRogueApCleanup(o["rogue-ap-cleanup"], d, "rogue_ap_cleanup")); err != nil {
		if vv, ok := fortiAPIPatch(o["rogue-ap-cleanup"], "WirelessControllerTimers-RogueApCleanup"); ok {
			if err = d.Set("rogue_ap_cleanup", vv); err != nil {
				return fmt.Errorf("Error reading rogue_ap_cleanup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rogue_ap_cleanup: %v", err)
		}
	}

	if err = d.Set("rogue_ap_log", flattenWirelessControllerTimersRogueApLog(o["rogue-ap-log"], d, "rogue_ap_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["rogue-ap-log"], "WirelessControllerTimers-RogueApLog"); ok {
			if err = d.Set("rogue_ap_log", vv); err != nil {
				return fmt.Errorf("Error reading rogue_ap_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rogue_ap_log: %v", err)
		}
	}

	if err = d.Set("rogue_sta_cleanup", flattenWirelessControllerTimersRogueStaCleanup(o["rogue-sta-cleanup"], d, "rogue_sta_cleanup")); err != nil {
		if vv, ok := fortiAPIPatch(o["rogue-sta-cleanup"], "WirelessControllerTimers-RogueStaCleanup"); ok {
			if err = d.Set("rogue_sta_cleanup", vv); err != nil {
				return fmt.Errorf("Error reading rogue_sta_cleanup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rogue_sta_cleanup: %v", err)
		}
	}

	if err = d.Set("sta_cap_cleanup", flattenWirelessControllerTimersStaCapCleanup(o["sta-cap-cleanup"], d, "sta_cap_cleanup")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-cap-cleanup"], "WirelessControllerTimers-StaCapCleanup"); ok {
			if err = d.Set("sta_cap_cleanup", vv); err != nil {
				return fmt.Errorf("Error reading sta_cap_cleanup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_cap_cleanup: %v", err)
		}
	}

	if err = d.Set("sta_capability_interval", flattenWirelessControllerTimersStaCapabilityInterval(o["sta-capability-interval"], d, "sta_capability_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-capability-interval"], "WirelessControllerTimers-StaCapabilityInterval"); ok {
			if err = d.Set("sta_capability_interval", vv); err != nil {
				return fmt.Errorf("Error reading sta_capability_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_capability_interval: %v", err)
		}
	}

	if err = d.Set("sta_locate_timer", flattenWirelessControllerTimersStaLocateTimer(o["sta-locate-timer"], d, "sta_locate_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-locate-timer"], "WirelessControllerTimers-StaLocateTimer"); ok {
			if err = d.Set("sta_locate_timer", vv); err != nil {
				return fmt.Errorf("Error reading sta_locate_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_locate_timer: %v", err)
		}
	}

	if err = d.Set("sta_stats_interval", flattenWirelessControllerTimersStaStatsInterval(o["sta-stats-interval"], d, "sta_stats_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sta-stats-interval"], "WirelessControllerTimers-StaStatsInterval"); ok {
			if err = d.Set("sta_stats_interval", vv); err != nil {
				return fmt.Errorf("Error reading sta_stats_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sta_stats_interval: %v", err)
		}
	}

	if err = d.Set("vap_stats_interval", flattenWirelessControllerTimersVapStatsInterval(o["vap-stats-interval"], d, "vap_stats_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap-stats-interval"], "WirelessControllerTimers-VapStatsInterval"); ok {
			if err = d.Set("vap_stats_interval", vv); err != nil {
				return fmt.Errorf("Error reading vap_stats_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap_stats_interval: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerTimersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerTimersApRebootWaitInterval1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersApRebootWaitInterval2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersApRebootWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersBleDeviceCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersBleScanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersClientIdleRehomeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersClientIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDiscoveryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDrmaInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersFakeApLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersIpsecIntfCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersNatSessionKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRadioStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueApCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueApLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueStaCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaCapCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaCapabilityInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaLocateTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersVapStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerTimers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ap_reboot_wait_interval1"); ok || d.HasChange("ap_reboot_wait_interval1") {
		t, err := expandWirelessControllerTimersApRebootWaitInterval1(d, v, "ap_reboot_wait_interval1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-reboot-wait-interval1"] = t
		}
	}

	if v, ok := d.GetOk("ap_reboot_wait_interval2"); ok || d.HasChange("ap_reboot_wait_interval2") {
		t, err := expandWirelessControllerTimersApRebootWaitInterval2(d, v, "ap_reboot_wait_interval2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-reboot-wait-interval2"] = t
		}
	}

	if v, ok := d.GetOk("ap_reboot_wait_time"); ok || d.HasChange("ap_reboot_wait_time") {
		t, err := expandWirelessControllerTimersApRebootWaitTime(d, v, "ap_reboot_wait_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-reboot-wait-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok || d.HasChange("auth_timeout") {
		t, err := expandWirelessControllerTimersAuthTimeout(d, v, "auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ble_device_cleanup"); ok || d.HasChange("ble_device_cleanup") {
		t, err := expandWirelessControllerTimersBleDeviceCleanup(d, v, "ble_device_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-device-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("ble_scan_report_intv"); ok || d.HasChange("ble_scan_report_intv") {
		t, err := expandWirelessControllerTimersBleScanReportIntv(d, v, "ble_scan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("client_idle_rehome_timeout"); ok || d.HasChange("client_idle_rehome_timeout") {
		t, err := expandWirelessControllerTimersClientIdleRehomeTimeout(d, v, "client_idle_rehome_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-idle-rehome-timeout"] = t
		}
	}

	if v, ok := d.GetOk("client_idle_timeout"); ok || d.HasChange("client_idle_timeout") {
		t, err := expandWirelessControllerTimersClientIdleTimeout(d, v, "client_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("discovery_interval"); ok || d.HasChange("discovery_interval") {
		t, err := expandWirelessControllerTimersDiscoveryInterval(d, v, "discovery_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-interval"] = t
		}
	}

	if v, ok := d.GetOk("drma_interval"); ok || d.HasChange("drma_interval") {
		t, err := expandWirelessControllerTimersDrmaInterval(d, v, "drma_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drma-interval"] = t
		}
	}

	if v, ok := d.GetOk("echo_interval"); ok || d.HasChange("echo_interval") {
		t, err := expandWirelessControllerTimersEchoInterval(d, v, "echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("fake_ap_log"); ok || d.HasChange("fake_ap_log") {
		t, err := expandWirelessControllerTimersFakeApLog(d, v, "fake_ap_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fake-ap-log"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_intf_cleanup"); ok || d.HasChange("ipsec_intf_cleanup") {
		t, err := expandWirelessControllerTimersIpsecIntfCleanup(d, v, "ipsec_intf_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-intf-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("nat_session_keep_alive"); ok || d.HasChange("nat_session_keep_alive") {
		t, err := expandWirelessControllerTimersNatSessionKeepAlive(d, v, "nat_session_keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-session-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("radio_stats_interval"); ok || d.HasChange("radio_stats_interval") {
		t, err := expandWirelessControllerTimersRadioStatsInterval(d, v, "radio_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-stats-interval"] = t
		}
	}

	if v, ok := d.GetOk("rogue_ap_cleanup"); ok || d.HasChange("rogue_ap_cleanup") {
		t, err := expandWirelessControllerTimersRogueApCleanup(d, v, "rogue_ap_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-ap-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("rogue_ap_log"); ok || d.HasChange("rogue_ap_log") {
		t, err := expandWirelessControllerTimersRogueApLog(d, v, "rogue_ap_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-ap-log"] = t
		}
	}

	if v, ok := d.GetOk("rogue_sta_cleanup"); ok || d.HasChange("rogue_sta_cleanup") {
		t, err := expandWirelessControllerTimersRogueStaCleanup(d, v, "rogue_sta_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-sta-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("sta_cap_cleanup"); ok || d.HasChange("sta_cap_cleanup") {
		t, err := expandWirelessControllerTimersStaCapCleanup(d, v, "sta_cap_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-cap-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("sta_capability_interval"); ok || d.HasChange("sta_capability_interval") {
		t, err := expandWirelessControllerTimersStaCapabilityInterval(d, v, "sta_capability_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-capability-interval"] = t
		}
	}

	if v, ok := d.GetOk("sta_locate_timer"); ok || d.HasChange("sta_locate_timer") {
		t, err := expandWirelessControllerTimersStaLocateTimer(d, v, "sta_locate_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-locate-timer"] = t
		}
	}

	if v, ok := d.GetOk("sta_stats_interval"); ok || d.HasChange("sta_stats_interval") {
		t, err := expandWirelessControllerTimersStaStatsInterval(d, v, "sta_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-stats-interval"] = t
		}
	}

	if v, ok := d.GetOk("vap_stats_interval"); ok || d.HasChange("vap_stats_interval") {
		t, err := expandWirelessControllerTimersVapStatsInterval(d, v, "vap_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap-stats-interval"] = t
		}
	}

	return &obj, nil
}
