// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless intrusion detection system (WIDS) profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWidsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWidsProfileCreate,
		Read:   resourceWirelessControllerWidsProfileRead,
		Update: resourceWirelessControllerWidsProfileUpdate,
		Delete: resourceWirelessControllerWidsProfileDelete,

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
			"adhoc_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adhoc_valid_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"air_jack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_auto_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_disable_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_bgscan_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_fgscan_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_impersonation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan_channel_list_2g_5g": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_scan_channel_list_6g": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_scan_passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_scan_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_spoofing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asleap_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"assoc_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bcn_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bcn_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bcn_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"beacon_wrong_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_ack_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_ack_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"block_ack_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"chan_based_mitm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cts_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cts_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cts_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deauth_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deauth_unknown_src_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disassoc_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect_station": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_key_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_start_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_start_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_start_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fata_jack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_beacon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_probe_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_probe_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hotspotter_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ht_40mhz_intolerance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ht_greenfield": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_addr_combination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_mac_oui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"malformed_association": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_ht_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"netstumbler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netstumbler_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netstumbler_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"null_ssid_probe_resp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"omerta_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overflow_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probe_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pspoll_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pspoll_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pspoll_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pwsave_dos_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reassoc_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reassoc_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reassoc_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"risky_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rts_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rts_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rts_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sensor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spoofed_deauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unencrypted_valid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_client_misassociation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_ssid_misuse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weak_wep_iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wellenreiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wellenreiter_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wellenreiter_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"windows_bridge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_bridge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wpa_ft_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWidsProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWidsProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWidsProfileRead(d, m)
}

func resourceWirelessControllerWidsProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWidsProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWidsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWidsProfileRead(d, m)
}

func resourceWirelessControllerWidsProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerWidsProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWidsProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWidsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWidsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWidsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWidsProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWidsProfileAdhocNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAdhocValidSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAirJack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApAutoSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDisableSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWidsProfileApBgscanDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApFgscanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApImpersonation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScanChannelList2G5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWidsProfileApScanChannelList6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWidsProfileApScanPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScanThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApSpoofing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAsleapAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFrameFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAuthFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAuthFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAuthFrameFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBcnFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBcnFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBcnFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBeaconWrongChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBlockAckFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBlockAckFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBlockAckFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileChanBasedMitm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileClientFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileClientFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileClientFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileCtsFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileCtsFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileCtsFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDisassocBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDisconnectStation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolKeyOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolStartFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolStartIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolStartThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFataJack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedProbeRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedProbeResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHotspotterAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHt40MhzIntolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHtGreenfield(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileInvalidAddrCombination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileInvalidMacOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedAssociation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedHtIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNetstumbler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNetstumblerThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNetstumblerTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNullSsidProbeResp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileOmertaAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileOverflowIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileProbeFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileProbeFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileProbeFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePspollFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePspollFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePspollFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePwsaveDosAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileReassocFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileReassocFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileReassocFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRiskyEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRtsFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRtsFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRtsFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileSensorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileSpoofedDeauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileUnencryptedValid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileValidClientMisassociation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileValidSsidMisuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWeakWepIv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWellenreiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWellenreiterThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWellenreiterTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWindowsBridge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWirelessBridge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWpaFtAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWidsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("adhoc_network", flattenWirelessControllerWidsProfileAdhocNetwork(o["adhoc-network"], d, "adhoc_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["adhoc-network"], "WirelessControllerWidsProfile-AdhocNetwork"); ok {
			if err = d.Set("adhoc_network", vv); err != nil {
				return fmt.Errorf("Error reading adhoc_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adhoc_network: %v", err)
		}
	}

	if err = d.Set("adhoc_valid_ssid", flattenWirelessControllerWidsProfileAdhocValidSsid(o["adhoc-valid-ssid"], d, "adhoc_valid_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["adhoc-valid-ssid"], "WirelessControllerWidsProfile-AdhocValidSsid"); ok {
			if err = d.Set("adhoc_valid_ssid", vv); err != nil {
				return fmt.Errorf("Error reading adhoc_valid_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adhoc_valid_ssid: %v", err)
		}
	}

	if err = d.Set("air_jack", flattenWirelessControllerWidsProfileAirJack(o["air-jack"], d, "air_jack")); err != nil {
		if vv, ok := fortiAPIPatch(o["air-jack"], "WirelessControllerWidsProfile-AirJack"); ok {
			if err = d.Set("air_jack", vv); err != nil {
				return fmt.Errorf("Error reading air_jack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading air_jack: %v", err)
		}
	}

	if err = d.Set("ap_auto_suppress", flattenWirelessControllerWidsProfileApAutoSuppress(o["ap-auto-suppress"], d, "ap_auto_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-auto-suppress"], "WirelessControllerWidsProfile-ApAutoSuppress"); ok {
			if err = d.Set("ap_auto_suppress", vv); err != nil {
				return fmt.Errorf("Error reading ap_auto_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_auto_suppress: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_schedules", flattenWirelessControllerWidsProfileApBgscanDisableSchedules(o["ap-bgscan-disable-schedules"], d, "ap_bgscan_disable_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-disable-schedules"], "WirelessControllerWidsProfile-ApBgscanDisableSchedules"); ok {
			if err = d.Set("ap_bgscan_disable_schedules", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_duration", flattenWirelessControllerWidsProfileApBgscanDuration(o["ap-bgscan-duration"], d, "ap_bgscan_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-duration"], "WirelessControllerWidsProfile-ApBgscanDuration"); ok {
			if err = d.Set("ap_bgscan_duration", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_duration: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_idle", flattenWirelessControllerWidsProfileApBgscanIdle(o["ap-bgscan-idle"], d, "ap_bgscan_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-idle"], "WirelessControllerWidsProfile-ApBgscanIdle"); ok {
			if err = d.Set("ap_bgscan_idle", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_idle: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_intv", flattenWirelessControllerWidsProfileApBgscanIntv(o["ap-bgscan-intv"], d, "ap_bgscan_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-intv"], "WirelessControllerWidsProfile-ApBgscanIntv"); ok {
			if err = d.Set("ap_bgscan_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_intv: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_period", flattenWirelessControllerWidsProfileApBgscanPeriod(o["ap-bgscan-period"], d, "ap_bgscan_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-period"], "WirelessControllerWidsProfile-ApBgscanPeriod"); ok {
			if err = d.Set("ap_bgscan_period", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_period: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_report_intv", flattenWirelessControllerWidsProfileApBgscanReportIntv(o["ap-bgscan-report-intv"], d, "ap_bgscan_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-report-intv"], "WirelessControllerWidsProfile-ApBgscanReportIntv"); ok {
			if err = d.Set("ap_bgscan_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_report_intv: %v", err)
		}
	}

	if err = d.Set("ap_fgscan_report_intv", flattenWirelessControllerWidsProfileApFgscanReportIntv(o["ap-fgscan-report-intv"], d, "ap_fgscan_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-fgscan-report-intv"], "WirelessControllerWidsProfile-ApFgscanReportIntv"); ok {
			if err = d.Set("ap_fgscan_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_fgscan_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_fgscan_report_intv: %v", err)
		}
	}

	if err = d.Set("ap_impersonation", flattenWirelessControllerWidsProfileApImpersonation(o["ap-impersonation"], d, "ap_impersonation")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-impersonation"], "WirelessControllerWidsProfile-ApImpersonation"); ok {
			if err = d.Set("ap_impersonation", vv); err != nil {
				return fmt.Errorf("Error reading ap_impersonation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_impersonation: %v", err)
		}
	}

	if err = d.Set("ap_scan", flattenWirelessControllerWidsProfileApScan(o["ap-scan"], d, "ap_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan"], "WirelessControllerWidsProfile-ApScan"); ok {
			if err = d.Set("ap_scan", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan: %v", err)
		}
	}

	if err = d.Set("ap_scan_channel_list_2g_5g", flattenWirelessControllerWidsProfileApScanChannelList2G5G(o["ap-scan-channel-list-2G-5G"], d, "ap_scan_channel_list_2g_5g")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-channel-list-2G-5G"], "WirelessControllerWidsProfile-ApScanChannelList2G5G"); ok {
			if err = d.Set("ap_scan_channel_list_2g_5g", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
		}
	}

	if err = d.Set("ap_scan_channel_list_6g", flattenWirelessControllerWidsProfileApScanChannelList6G(o["ap-scan-channel-list-6G"], d, "ap_scan_channel_list_6g")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-channel-list-6G"], "WirelessControllerWidsProfile-ApScanChannelList6G"); ok {
			if err = d.Set("ap_scan_channel_list_6g", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
		}
	}

	if err = d.Set("ap_scan_passive", flattenWirelessControllerWidsProfileApScanPassive(o["ap-scan-passive"], d, "ap_scan_passive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-passive"], "WirelessControllerWidsProfile-ApScanPassive"); ok {
			if err = d.Set("ap_scan_passive", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_passive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_passive: %v", err)
		}
	}

	if err = d.Set("ap_scan_threshold", flattenWirelessControllerWidsProfileApScanThreshold(o["ap-scan-threshold"], d, "ap_scan_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-threshold"], "WirelessControllerWidsProfile-ApScanThreshold"); ok {
			if err = d.Set("ap_scan_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_threshold: %v", err)
		}
	}

	if err = d.Set("ap_spoofing", flattenWirelessControllerWidsProfileApSpoofing(o["ap-spoofing"], d, "ap_spoofing")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-spoofing"], "WirelessControllerWidsProfile-ApSpoofing"); ok {
			if err = d.Set("ap_spoofing", vv); err != nil {
				return fmt.Errorf("Error reading ap_spoofing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_spoofing: %v", err)
		}
	}

	if err = d.Set("asleap_attack", flattenWirelessControllerWidsProfileAsleapAttack(o["asleap-attack"], d, "asleap_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["asleap-attack"], "WirelessControllerWidsProfile-AsleapAttack"); ok {
			if err = d.Set("asleap_attack", vv); err != nil {
				return fmt.Errorf("Error reading asleap_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asleap_attack: %v", err)
		}
	}

	if err = d.Set("assoc_flood_thresh", flattenWirelessControllerWidsProfileAssocFloodThresh(o["assoc-flood-thresh"], d, "assoc_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-flood-thresh"], "WirelessControllerWidsProfile-AssocFloodThresh"); ok {
			if err = d.Set("assoc_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading assoc_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_flood_thresh: %v", err)
		}
	}

	if err = d.Set("assoc_flood_time", flattenWirelessControllerWidsProfileAssocFloodTime(o["assoc-flood-time"], d, "assoc_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-flood-time"], "WirelessControllerWidsProfile-AssocFloodTime"); ok {
			if err = d.Set("assoc_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading assoc_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_flood_time: %v", err)
		}
	}

	if err = d.Set("assoc_frame_flood", flattenWirelessControllerWidsProfileAssocFrameFlood(o["assoc-frame-flood"], d, "assoc_frame_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-frame-flood"], "WirelessControllerWidsProfile-AssocFrameFlood"); ok {
			if err = d.Set("assoc_frame_flood", vv); err != nil {
				return fmt.Errorf("Error reading assoc_frame_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_frame_flood: %v", err)
		}
	}

	if err = d.Set("auth_flood_thresh", flattenWirelessControllerWidsProfileAuthFloodThresh(o["auth-flood-thresh"], d, "auth_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-flood-thresh"], "WirelessControllerWidsProfile-AuthFloodThresh"); ok {
			if err = d.Set("auth_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading auth_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_flood_thresh: %v", err)
		}
	}

	if err = d.Set("auth_flood_time", flattenWirelessControllerWidsProfileAuthFloodTime(o["auth-flood-time"], d, "auth_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-flood-time"], "WirelessControllerWidsProfile-AuthFloodTime"); ok {
			if err = d.Set("auth_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading auth_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_flood_time: %v", err)
		}
	}

	if err = d.Set("auth_frame_flood", flattenWirelessControllerWidsProfileAuthFrameFlood(o["auth-frame-flood"], d, "auth_frame_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-frame-flood"], "WirelessControllerWidsProfile-AuthFrameFlood"); ok {
			if err = d.Set("auth_frame_flood", vv); err != nil {
				return fmt.Errorf("Error reading auth_frame_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_frame_flood: %v", err)
		}
	}

	if err = d.Set("bcn_flood", flattenWirelessControllerWidsProfileBcnFlood(o["bcn-flood"], d, "bcn_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["bcn-flood"], "WirelessControllerWidsProfile-BcnFlood"); ok {
			if err = d.Set("bcn_flood", vv); err != nil {
				return fmt.Errorf("Error reading bcn_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bcn_flood: %v", err)
		}
	}

	if err = d.Set("bcn_flood_thresh", flattenWirelessControllerWidsProfileBcnFloodThresh(o["bcn-flood-thresh"], d, "bcn_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["bcn-flood-thresh"], "WirelessControllerWidsProfile-BcnFloodThresh"); ok {
			if err = d.Set("bcn_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading bcn_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bcn_flood_thresh: %v", err)
		}
	}

	if err = d.Set("bcn_flood_time", flattenWirelessControllerWidsProfileBcnFloodTime(o["bcn-flood-time"], d, "bcn_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["bcn-flood-time"], "WirelessControllerWidsProfile-BcnFloodTime"); ok {
			if err = d.Set("bcn_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading bcn_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bcn_flood_time: %v", err)
		}
	}

	if err = d.Set("beacon_wrong_channel", flattenWirelessControllerWidsProfileBeaconWrongChannel(o["beacon-wrong-channel"], d, "beacon_wrong_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-wrong-channel"], "WirelessControllerWidsProfile-BeaconWrongChannel"); ok {
			if err = d.Set("beacon_wrong_channel", vv); err != nil {
				return fmt.Errorf("Error reading beacon_wrong_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_wrong_channel: %v", err)
		}
	}

	if err = d.Set("block_ack_flood", flattenWirelessControllerWidsProfileBlockAckFlood(o["block_ack-flood"], d, "block_ack_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["block_ack-flood"], "WirelessControllerWidsProfile-BlockAckFlood"); ok {
			if err = d.Set("block_ack_flood", vv); err != nil {
				return fmt.Errorf("Error reading block_ack_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_ack_flood: %v", err)
		}
	}

	if err = d.Set("block_ack_flood_thresh", flattenWirelessControllerWidsProfileBlockAckFloodThresh(o["block_ack-flood-thresh"], d, "block_ack_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["block_ack-flood-thresh"], "WirelessControllerWidsProfile-BlockAckFloodThresh"); ok {
			if err = d.Set("block_ack_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading block_ack_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_ack_flood_thresh: %v", err)
		}
	}

	if err = d.Set("block_ack_flood_time", flattenWirelessControllerWidsProfileBlockAckFloodTime(o["block_ack-flood-time"], d, "block_ack_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["block_ack-flood-time"], "WirelessControllerWidsProfile-BlockAckFloodTime"); ok {
			if err = d.Set("block_ack_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading block_ack_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_ack_flood_time: %v", err)
		}
	}

	if err = d.Set("chan_based_mitm", flattenWirelessControllerWidsProfileChanBasedMitm(o["chan-based-mitm"], d, "chan_based_mitm")); err != nil {
		if vv, ok := fortiAPIPatch(o["chan-based-mitm"], "WirelessControllerWidsProfile-ChanBasedMitm"); ok {
			if err = d.Set("chan_based_mitm", vv); err != nil {
				return fmt.Errorf("Error reading chan_based_mitm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chan_based_mitm: %v", err)
		}
	}

	if err = d.Set("client_flood", flattenWirelessControllerWidsProfileClientFlood(o["client-flood"], d, "client_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-flood"], "WirelessControllerWidsProfile-ClientFlood"); ok {
			if err = d.Set("client_flood", vv); err != nil {
				return fmt.Errorf("Error reading client_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_flood: %v", err)
		}
	}

	if err = d.Set("client_flood_thresh", flattenWirelessControllerWidsProfileClientFloodThresh(o["client-flood-thresh"], d, "client_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-flood-thresh"], "WirelessControllerWidsProfile-ClientFloodThresh"); ok {
			if err = d.Set("client_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading client_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_flood_thresh: %v", err)
		}
	}

	if err = d.Set("client_flood_time", flattenWirelessControllerWidsProfileClientFloodTime(o["client-flood-time"], d, "client_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-flood-time"], "WirelessControllerWidsProfile-ClientFloodTime"); ok {
			if err = d.Set("client_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading client_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_flood_time: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWidsProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerWidsProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cts_flood", flattenWirelessControllerWidsProfileCtsFlood(o["cts-flood"], d, "cts_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["cts-flood"], "WirelessControllerWidsProfile-CtsFlood"); ok {
			if err = d.Set("cts_flood", vv); err != nil {
				return fmt.Errorf("Error reading cts_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cts_flood: %v", err)
		}
	}

	if err = d.Set("cts_flood_thresh", flattenWirelessControllerWidsProfileCtsFloodThresh(o["cts-flood-thresh"], d, "cts_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["cts-flood-thresh"], "WirelessControllerWidsProfile-CtsFloodThresh"); ok {
			if err = d.Set("cts_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading cts_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cts_flood_thresh: %v", err)
		}
	}

	if err = d.Set("cts_flood_time", flattenWirelessControllerWidsProfileCtsFloodTime(o["cts-flood-time"], d, "cts_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["cts-flood-time"], "WirelessControllerWidsProfile-CtsFloodTime"); ok {
			if err = d.Set("cts_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading cts_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cts_flood_time: %v", err)
		}
	}

	if err = d.Set("deauth_broadcast", flattenWirelessControllerWidsProfileDeauthBroadcast(o["deauth-broadcast"], d, "deauth_broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-broadcast"], "WirelessControllerWidsProfile-DeauthBroadcast"); ok {
			if err = d.Set("deauth_broadcast", vv); err != nil {
				return fmt.Errorf("Error reading deauth_broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_broadcast: %v", err)
		}
	}

	if err = d.Set("deauth_unknown_src_thresh", flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(o["deauth-unknown-src-thresh"], d, "deauth_unknown_src_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-unknown-src-thresh"], "WirelessControllerWidsProfile-DeauthUnknownSrcThresh"); ok {
			if err = d.Set("deauth_unknown_src_thresh", vv); err != nil {
				return fmt.Errorf("Error reading deauth_unknown_src_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_unknown_src_thresh: %v", err)
		}
	}

	if err = d.Set("disassoc_broadcast", flattenWirelessControllerWidsProfileDisassocBroadcast(o["disassoc-broadcast"], d, "disassoc_broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["disassoc-broadcast"], "WirelessControllerWidsProfile-DisassocBroadcast"); ok {
			if err = d.Set("disassoc_broadcast", vv); err != nil {
				return fmt.Errorf("Error reading disassoc_broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disassoc_broadcast: %v", err)
		}
	}

	if err = d.Set("disconnect_station", flattenWirelessControllerWidsProfileDisconnectStation(o["disconnect-station"], d, "disconnect_station")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-station"], "WirelessControllerWidsProfile-DisconnectStation"); ok {
			if err = d.Set("disconnect_station", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_station: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_station: %v", err)
		}
	}

	if err = d.Set("eapol_fail_flood", flattenWirelessControllerWidsProfileEapolFailFlood(o["eapol-fail-flood"], d, "eapol_fail_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-flood"], "WirelessControllerWidsProfile-EapolFailFlood"); ok {
			if err = d.Set("eapol_fail_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_fail_intv", flattenWirelessControllerWidsProfileEapolFailIntv(o["eapol-fail-intv"], d, "eapol_fail_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-intv"], "WirelessControllerWidsProfile-EapolFailIntv"); ok {
			if err = d.Set("eapol_fail_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_intv: %v", err)
		}
	}

	if err = d.Set("eapol_fail_thresh", flattenWirelessControllerWidsProfileEapolFailThresh(o["eapol-fail-thresh"], d, "eapol_fail_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-thresh"], "WirelessControllerWidsProfile-EapolFailThresh"); ok {
			if err = d.Set("eapol_fail_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_key_overflow", flattenWirelessControllerWidsProfileEapolKeyOverflow(o["eapol-key-overflow"], d, "eapol_key_overflow")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-key-overflow"], "WirelessControllerWidsProfile-EapolKeyOverflow"); ok {
			if err = d.Set("eapol_key_overflow", vv); err != nil {
				return fmt.Errorf("Error reading eapol_key_overflow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_key_overflow: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_flood", flattenWirelessControllerWidsProfileEapolLogoffFlood(o["eapol-logoff-flood"], d, "eapol_logoff_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-flood"], "WirelessControllerWidsProfile-EapolLogoffFlood"); ok {
			if err = d.Set("eapol_logoff_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_flood: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_intv", flattenWirelessControllerWidsProfileEapolLogoffIntv(o["eapol-logoff-intv"], d, "eapol_logoff_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-intv"], "WirelessControllerWidsProfile-EapolLogoffIntv"); ok {
			if err = d.Set("eapol_logoff_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_intv: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_thresh", flattenWirelessControllerWidsProfileEapolLogoffThresh(o["eapol-logoff-thresh"], d, "eapol_logoff_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-thresh"], "WirelessControllerWidsProfile-EapolLogoffThresh"); ok {
			if err = d.Set("eapol_logoff_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_flood", flattenWirelessControllerWidsProfileEapolPreFailFlood(o["eapol-pre-fail-flood"], d, "eapol_pre_fail_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-flood"], "WirelessControllerWidsProfile-EapolPreFailFlood"); ok {
			if err = d.Set("eapol_pre_fail_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_intv", flattenWirelessControllerWidsProfileEapolPreFailIntv(o["eapol-pre-fail-intv"], d, "eapol_pre_fail_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-intv"], "WirelessControllerWidsProfile-EapolPreFailIntv"); ok {
			if err = d.Set("eapol_pre_fail_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_thresh", flattenWirelessControllerWidsProfileEapolPreFailThresh(o["eapol-pre-fail-thresh"], d, "eapol_pre_fail_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-thresh"], "WirelessControllerWidsProfile-EapolPreFailThresh"); ok {
			if err = d.Set("eapol_pre_fail_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_flood", flattenWirelessControllerWidsProfileEapolPreSuccFlood(o["eapol-pre-succ-flood"], d, "eapol_pre_succ_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-flood"], "WirelessControllerWidsProfile-EapolPreSuccFlood"); ok {
			if err = d.Set("eapol_pre_succ_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_intv", flattenWirelessControllerWidsProfileEapolPreSuccIntv(o["eapol-pre-succ-intv"], d, "eapol_pre_succ_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-intv"], "WirelessControllerWidsProfile-EapolPreSuccIntv"); ok {
			if err = d.Set("eapol_pre_succ_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_thresh", flattenWirelessControllerWidsProfileEapolPreSuccThresh(o["eapol-pre-succ-thresh"], d, "eapol_pre_succ_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-thresh"], "WirelessControllerWidsProfile-EapolPreSuccThresh"); ok {
			if err = d.Set("eapol_pre_succ_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_start_flood", flattenWirelessControllerWidsProfileEapolStartFlood(o["eapol-start-flood"], d, "eapol_start_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-flood"], "WirelessControllerWidsProfile-EapolStartFlood"); ok {
			if err = d.Set("eapol_start_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_flood: %v", err)
		}
	}

	if err = d.Set("eapol_start_intv", flattenWirelessControllerWidsProfileEapolStartIntv(o["eapol-start-intv"], d, "eapol_start_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-intv"], "WirelessControllerWidsProfile-EapolStartIntv"); ok {
			if err = d.Set("eapol_start_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_intv: %v", err)
		}
	}

	if err = d.Set("eapol_start_thresh", flattenWirelessControllerWidsProfileEapolStartThresh(o["eapol-start-thresh"], d, "eapol_start_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-thresh"], "WirelessControllerWidsProfile-EapolStartThresh"); ok {
			if err = d.Set("eapol_start_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_succ_flood", flattenWirelessControllerWidsProfileEapolSuccFlood(o["eapol-succ-flood"], d, "eapol_succ_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-flood"], "WirelessControllerWidsProfile-EapolSuccFlood"); ok {
			if err = d.Set("eapol_succ_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_succ_intv", flattenWirelessControllerWidsProfileEapolSuccIntv(o["eapol-succ-intv"], d, "eapol_succ_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-intv"], "WirelessControllerWidsProfile-EapolSuccIntv"); ok {
			if err = d.Set("eapol_succ_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_succ_thresh", flattenWirelessControllerWidsProfileEapolSuccThresh(o["eapol-succ-thresh"], d, "eapol_succ_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-thresh"], "WirelessControllerWidsProfile-EapolSuccThresh"); ok {
			if err = d.Set("eapol_succ_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_thresh: %v", err)
		}
	}

	if err = d.Set("fata_jack", flattenWirelessControllerWidsProfileFataJack(o["fata-jack"], d, "fata_jack")); err != nil {
		if vv, ok := fortiAPIPatch(o["fata-jack"], "WirelessControllerWidsProfile-FataJack"); ok {
			if err = d.Set("fata_jack", vv); err != nil {
				return fmt.Errorf("Error reading fata_jack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fata_jack: %v", err)
		}
	}

	if err = d.Set("fuzzed_beacon", flattenWirelessControllerWidsProfileFuzzedBeacon(o["fuzzed-beacon"], d, "fuzzed_beacon")); err != nil {
		if vv, ok := fortiAPIPatch(o["fuzzed-beacon"], "WirelessControllerWidsProfile-FuzzedBeacon"); ok {
			if err = d.Set("fuzzed_beacon", vv); err != nil {
				return fmt.Errorf("Error reading fuzzed_beacon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fuzzed_beacon: %v", err)
		}
	}

	if err = d.Set("fuzzed_probe_request", flattenWirelessControllerWidsProfileFuzzedProbeRequest(o["fuzzed-probe-request"], d, "fuzzed_probe_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["fuzzed-probe-request"], "WirelessControllerWidsProfile-FuzzedProbeRequest"); ok {
			if err = d.Set("fuzzed_probe_request", vv); err != nil {
				return fmt.Errorf("Error reading fuzzed_probe_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fuzzed_probe_request: %v", err)
		}
	}

	if err = d.Set("fuzzed_probe_response", flattenWirelessControllerWidsProfileFuzzedProbeResponse(o["fuzzed-probe-response"], d, "fuzzed_probe_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["fuzzed-probe-response"], "WirelessControllerWidsProfile-FuzzedProbeResponse"); ok {
			if err = d.Set("fuzzed_probe_response", vv); err != nil {
				return fmt.Errorf("Error reading fuzzed_probe_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fuzzed_probe_response: %v", err)
		}
	}

	if err = d.Set("hotspotter_attack", flattenWirelessControllerWidsProfileHotspotterAttack(o["hotspotter-attack"], d, "hotspotter_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["hotspotter-attack"], "WirelessControllerWidsProfile-HotspotterAttack"); ok {
			if err = d.Set("hotspotter_attack", vv); err != nil {
				return fmt.Errorf("Error reading hotspotter_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hotspotter_attack: %v", err)
		}
	}

	if err = d.Set("ht_40mhz_intolerance", flattenWirelessControllerWidsProfileHt40MhzIntolerance(o["ht-40mhz-intolerance"], d, "ht_40mhz_intolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["ht-40mhz-intolerance"], "WirelessControllerWidsProfile-Ht40MhzIntolerance"); ok {
			if err = d.Set("ht_40mhz_intolerance", vv); err != nil {
				return fmt.Errorf("Error reading ht_40mhz_intolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ht_40mhz_intolerance: %v", err)
		}
	}

	if err = d.Set("ht_greenfield", flattenWirelessControllerWidsProfileHtGreenfield(o["ht-greenfield"], d, "ht_greenfield")); err != nil {
		if vv, ok := fortiAPIPatch(o["ht-greenfield"], "WirelessControllerWidsProfile-HtGreenfield"); ok {
			if err = d.Set("ht_greenfield", vv); err != nil {
				return fmt.Errorf("Error reading ht_greenfield: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ht_greenfield: %v", err)
		}
	}

	if err = d.Set("invalid_addr_combination", flattenWirelessControllerWidsProfileInvalidAddrCombination(o["invalid-addr-combination"], d, "invalid_addr_combination")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-addr-combination"], "WirelessControllerWidsProfile-InvalidAddrCombination"); ok {
			if err = d.Set("invalid_addr_combination", vv); err != nil {
				return fmt.Errorf("Error reading invalid_addr_combination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_addr_combination: %v", err)
		}
	}

	if err = d.Set("invalid_mac_oui", flattenWirelessControllerWidsProfileInvalidMacOui(o["invalid-mac-oui"], d, "invalid_mac_oui")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-mac-oui"], "WirelessControllerWidsProfile-InvalidMacOui"); ok {
			if err = d.Set("invalid_mac_oui", vv); err != nil {
				return fmt.Errorf("Error reading invalid_mac_oui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_mac_oui: %v", err)
		}
	}

	if err = d.Set("long_duration_attack", flattenWirelessControllerWidsProfileLongDurationAttack(o["long-duration-attack"], d, "long_duration_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-duration-attack"], "WirelessControllerWidsProfile-LongDurationAttack"); ok {
			if err = d.Set("long_duration_attack", vv); err != nil {
				return fmt.Errorf("Error reading long_duration_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_duration_attack: %v", err)
		}
	}

	if err = d.Set("long_duration_thresh", flattenWirelessControllerWidsProfileLongDurationThresh(o["long-duration-thresh"], d, "long_duration_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-duration-thresh"], "WirelessControllerWidsProfile-LongDurationThresh"); ok {
			if err = d.Set("long_duration_thresh", vv); err != nil {
				return fmt.Errorf("Error reading long_duration_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_duration_thresh: %v", err)
		}
	}

	if err = d.Set("malformed_association", flattenWirelessControllerWidsProfileMalformedAssociation(o["malformed-association"], d, "malformed_association")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-association"], "WirelessControllerWidsProfile-MalformedAssociation"); ok {
			if err = d.Set("malformed_association", vv); err != nil {
				return fmt.Errorf("Error reading malformed_association: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_association: %v", err)
		}
	}

	if err = d.Set("malformed_auth", flattenWirelessControllerWidsProfileMalformedAuth(o["malformed-auth"], d, "malformed_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-auth"], "WirelessControllerWidsProfile-MalformedAuth"); ok {
			if err = d.Set("malformed_auth", vv); err != nil {
				return fmt.Errorf("Error reading malformed_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_auth: %v", err)
		}
	}

	if err = d.Set("malformed_ht_ie", flattenWirelessControllerWidsProfileMalformedHtIe(o["malformed-ht-ie"], d, "malformed_ht_ie")); err != nil {
		if vv, ok := fortiAPIPatch(o["malformed-ht-ie"], "WirelessControllerWidsProfile-MalformedHtIe"); ok {
			if err = d.Set("malformed_ht_ie", vv); err != nil {
				return fmt.Errorf("Error reading malformed_ht_ie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malformed_ht_ie: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWidsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerWidsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("netstumbler", flattenWirelessControllerWidsProfileNetstumbler(o["netstumbler"], d, "netstumbler")); err != nil {
		if vv, ok := fortiAPIPatch(o["netstumbler"], "WirelessControllerWidsProfile-Netstumbler"); ok {
			if err = d.Set("netstumbler", vv); err != nil {
				return fmt.Errorf("Error reading netstumbler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netstumbler: %v", err)
		}
	}

	if err = d.Set("netstumbler_thresh", flattenWirelessControllerWidsProfileNetstumblerThresh(o["netstumbler-thresh"], d, "netstumbler_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["netstumbler-thresh"], "WirelessControllerWidsProfile-NetstumblerThresh"); ok {
			if err = d.Set("netstumbler_thresh", vv); err != nil {
				return fmt.Errorf("Error reading netstumbler_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netstumbler_thresh: %v", err)
		}
	}

	if err = d.Set("netstumbler_time", flattenWirelessControllerWidsProfileNetstumblerTime(o["netstumbler-time"], d, "netstumbler_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["netstumbler-time"], "WirelessControllerWidsProfile-NetstumblerTime"); ok {
			if err = d.Set("netstumbler_time", vv); err != nil {
				return fmt.Errorf("Error reading netstumbler_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netstumbler_time: %v", err)
		}
	}

	if err = d.Set("null_ssid_probe_resp", flattenWirelessControllerWidsProfileNullSsidProbeResp(o["null-ssid-probe-resp"], d, "null_ssid_probe_resp")); err != nil {
		if vv, ok := fortiAPIPatch(o["null-ssid-probe-resp"], "WirelessControllerWidsProfile-NullSsidProbeResp"); ok {
			if err = d.Set("null_ssid_probe_resp", vv); err != nil {
				return fmt.Errorf("Error reading null_ssid_probe_resp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading null_ssid_probe_resp: %v", err)
		}
	}

	if err = d.Set("omerta_attack", flattenWirelessControllerWidsProfileOmertaAttack(o["omerta-attack"], d, "omerta_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["omerta-attack"], "WirelessControllerWidsProfile-OmertaAttack"); ok {
			if err = d.Set("omerta_attack", vv); err != nil {
				return fmt.Errorf("Error reading omerta_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading omerta_attack: %v", err)
		}
	}

	if err = d.Set("overflow_ie", flattenWirelessControllerWidsProfileOverflowIe(o["overflow-ie"], d, "overflow_ie")); err != nil {
		if vv, ok := fortiAPIPatch(o["overflow-ie"], "WirelessControllerWidsProfile-OverflowIe"); ok {
			if err = d.Set("overflow_ie", vv); err != nil {
				return fmt.Errorf("Error reading overflow_ie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overflow_ie: %v", err)
		}
	}

	if err = d.Set("probe_flood", flattenWirelessControllerWidsProfileProbeFlood(o["probe-flood"], d, "probe_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-flood"], "WirelessControllerWidsProfile-ProbeFlood"); ok {
			if err = d.Set("probe_flood", vv); err != nil {
				return fmt.Errorf("Error reading probe_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_flood: %v", err)
		}
	}

	if err = d.Set("probe_flood_thresh", flattenWirelessControllerWidsProfileProbeFloodThresh(o["probe-flood-thresh"], d, "probe_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-flood-thresh"], "WirelessControllerWidsProfile-ProbeFloodThresh"); ok {
			if err = d.Set("probe_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading probe_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_flood_thresh: %v", err)
		}
	}

	if err = d.Set("probe_flood_time", flattenWirelessControllerWidsProfileProbeFloodTime(o["probe-flood-time"], d, "probe_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-flood-time"], "WirelessControllerWidsProfile-ProbeFloodTime"); ok {
			if err = d.Set("probe_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading probe_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_flood_time: %v", err)
		}
	}

	if err = d.Set("pspoll_flood", flattenWirelessControllerWidsProfilePspollFlood(o["pspoll-flood"], d, "pspoll_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["pspoll-flood"], "WirelessControllerWidsProfile-PspollFlood"); ok {
			if err = d.Set("pspoll_flood", vv); err != nil {
				return fmt.Errorf("Error reading pspoll_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pspoll_flood: %v", err)
		}
	}

	if err = d.Set("pspoll_flood_thresh", flattenWirelessControllerWidsProfilePspollFloodThresh(o["pspoll-flood-thresh"], d, "pspoll_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["pspoll-flood-thresh"], "WirelessControllerWidsProfile-PspollFloodThresh"); ok {
			if err = d.Set("pspoll_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading pspoll_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pspoll_flood_thresh: %v", err)
		}
	}

	if err = d.Set("pspoll_flood_time", flattenWirelessControllerWidsProfilePspollFloodTime(o["pspoll-flood-time"], d, "pspoll_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["pspoll-flood-time"], "WirelessControllerWidsProfile-PspollFloodTime"); ok {
			if err = d.Set("pspoll_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading pspoll_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pspoll_flood_time: %v", err)
		}
	}

	if err = d.Set("pwsave_dos_attack", flattenWirelessControllerWidsProfilePwsaveDosAttack(o["pwsave-dos-attack"], d, "pwsave_dos_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["pwsave-dos-attack"], "WirelessControllerWidsProfile-PwsaveDosAttack"); ok {
			if err = d.Set("pwsave_dos_attack", vv); err != nil {
				return fmt.Errorf("Error reading pwsave_dos_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pwsave_dos_attack: %v", err)
		}
	}

	if err = d.Set("reassoc_flood", flattenWirelessControllerWidsProfileReassocFlood(o["reassoc-flood"], d, "reassoc_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["reassoc-flood"], "WirelessControllerWidsProfile-ReassocFlood"); ok {
			if err = d.Set("reassoc_flood", vv); err != nil {
				return fmt.Errorf("Error reading reassoc_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reassoc_flood: %v", err)
		}
	}

	if err = d.Set("reassoc_flood_thresh", flattenWirelessControllerWidsProfileReassocFloodThresh(o["reassoc-flood-thresh"], d, "reassoc_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["reassoc-flood-thresh"], "WirelessControllerWidsProfile-ReassocFloodThresh"); ok {
			if err = d.Set("reassoc_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading reassoc_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reassoc_flood_thresh: %v", err)
		}
	}

	if err = d.Set("reassoc_flood_time", flattenWirelessControllerWidsProfileReassocFloodTime(o["reassoc-flood-time"], d, "reassoc_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["reassoc-flood-time"], "WirelessControllerWidsProfile-ReassocFloodTime"); ok {
			if err = d.Set("reassoc_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading reassoc_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reassoc_flood_time: %v", err)
		}
	}

	if err = d.Set("risky_encryption", flattenWirelessControllerWidsProfileRiskyEncryption(o["risky-encryption"], d, "risky_encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["risky-encryption"], "WirelessControllerWidsProfile-RiskyEncryption"); ok {
			if err = d.Set("risky_encryption", vv); err != nil {
				return fmt.Errorf("Error reading risky_encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risky_encryption: %v", err)
		}
	}

	if err = d.Set("rts_flood", flattenWirelessControllerWidsProfileRtsFlood(o["rts-flood"], d, "rts_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["rts-flood"], "WirelessControllerWidsProfile-RtsFlood"); ok {
			if err = d.Set("rts_flood", vv); err != nil {
				return fmt.Errorf("Error reading rts_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rts_flood: %v", err)
		}
	}

	if err = d.Set("rts_flood_thresh", flattenWirelessControllerWidsProfileRtsFloodThresh(o["rts-flood-thresh"], d, "rts_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["rts-flood-thresh"], "WirelessControllerWidsProfile-RtsFloodThresh"); ok {
			if err = d.Set("rts_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading rts_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rts_flood_thresh: %v", err)
		}
	}

	if err = d.Set("rts_flood_time", flattenWirelessControllerWidsProfileRtsFloodTime(o["rts-flood-time"], d, "rts_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["rts-flood-time"], "WirelessControllerWidsProfile-RtsFloodTime"); ok {
			if err = d.Set("rts_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading rts_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rts_flood_time: %v", err)
		}
	}

	if err = d.Set("sensor_mode", flattenWirelessControllerWidsProfileSensorMode(o["sensor-mode"], d, "sensor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensor-mode"], "WirelessControllerWidsProfile-SensorMode"); ok {
			if err = d.Set("sensor_mode", vv); err != nil {
				return fmt.Errorf("Error reading sensor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensor_mode: %v", err)
		}
	}

	if err = d.Set("spoofed_deauth", flattenWirelessControllerWidsProfileSpoofedDeauth(o["spoofed-deauth"], d, "spoofed_deauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["spoofed-deauth"], "WirelessControllerWidsProfile-SpoofedDeauth"); ok {
			if err = d.Set("spoofed_deauth", vv); err != nil {
				return fmt.Errorf("Error reading spoofed_deauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spoofed_deauth: %v", err)
		}
	}

	if err = d.Set("unencrypted_valid", flattenWirelessControllerWidsProfileUnencryptedValid(o["unencrypted-valid"], d, "unencrypted_valid")); err != nil {
		if vv, ok := fortiAPIPatch(o["unencrypted-valid"], "WirelessControllerWidsProfile-UnencryptedValid"); ok {
			if err = d.Set("unencrypted_valid", vv); err != nil {
				return fmt.Errorf("Error reading unencrypted_valid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unencrypted_valid: %v", err)
		}
	}

	if err = d.Set("valid_client_misassociation", flattenWirelessControllerWidsProfileValidClientMisassociation(o["valid-client-misassociation"], d, "valid_client_misassociation")); err != nil {
		if vv, ok := fortiAPIPatch(o["valid-client-misassociation"], "WirelessControllerWidsProfile-ValidClientMisassociation"); ok {
			if err = d.Set("valid_client_misassociation", vv); err != nil {
				return fmt.Errorf("Error reading valid_client_misassociation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading valid_client_misassociation: %v", err)
		}
	}

	if err = d.Set("valid_ssid_misuse", flattenWirelessControllerWidsProfileValidSsidMisuse(o["valid-ssid-misuse"], d, "valid_ssid_misuse")); err != nil {
		if vv, ok := fortiAPIPatch(o["valid-ssid-misuse"], "WirelessControllerWidsProfile-ValidSsidMisuse"); ok {
			if err = d.Set("valid_ssid_misuse", vv); err != nil {
				return fmt.Errorf("Error reading valid_ssid_misuse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading valid_ssid_misuse: %v", err)
		}
	}

	if err = d.Set("weak_wep_iv", flattenWirelessControllerWidsProfileWeakWepIv(o["weak-wep-iv"], d, "weak_wep_iv")); err != nil {
		if vv, ok := fortiAPIPatch(o["weak-wep-iv"], "WirelessControllerWidsProfile-WeakWepIv"); ok {
			if err = d.Set("weak_wep_iv", vv); err != nil {
				return fmt.Errorf("Error reading weak_wep_iv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weak_wep_iv: %v", err)
		}
	}

	if err = d.Set("wellenreiter", flattenWirelessControllerWidsProfileWellenreiter(o["wellenreiter"], d, "wellenreiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["wellenreiter"], "WirelessControllerWidsProfile-Wellenreiter"); ok {
			if err = d.Set("wellenreiter", vv); err != nil {
				return fmt.Errorf("Error reading wellenreiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wellenreiter: %v", err)
		}
	}

	if err = d.Set("wellenreiter_thresh", flattenWirelessControllerWidsProfileWellenreiterThresh(o["wellenreiter-thresh"], d, "wellenreiter_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["wellenreiter-thresh"], "WirelessControllerWidsProfile-WellenreiterThresh"); ok {
			if err = d.Set("wellenreiter_thresh", vv); err != nil {
				return fmt.Errorf("Error reading wellenreiter_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wellenreiter_thresh: %v", err)
		}
	}

	if err = d.Set("wellenreiter_time", flattenWirelessControllerWidsProfileWellenreiterTime(o["wellenreiter-time"], d, "wellenreiter_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wellenreiter-time"], "WirelessControllerWidsProfile-WellenreiterTime"); ok {
			if err = d.Set("wellenreiter_time", vv); err != nil {
				return fmt.Errorf("Error reading wellenreiter_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wellenreiter_time: %v", err)
		}
	}

	if err = d.Set("windows_bridge", flattenWirelessControllerWidsProfileWindowsBridge(o["windows-bridge"], d, "windows_bridge")); err != nil {
		if vv, ok := fortiAPIPatch(o["windows-bridge"], "WirelessControllerWidsProfile-WindowsBridge"); ok {
			if err = d.Set("windows_bridge", vv); err != nil {
				return fmt.Errorf("Error reading windows_bridge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading windows_bridge: %v", err)
		}
	}

	if err = d.Set("wireless_bridge", flattenWirelessControllerWidsProfileWirelessBridge(o["wireless-bridge"], d, "wireless_bridge")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-bridge"], "WirelessControllerWidsProfile-WirelessBridge"); ok {
			if err = d.Set("wireless_bridge", vv); err != nil {
				return fmt.Errorf("Error reading wireless_bridge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_bridge: %v", err)
		}
	}

	if err = d.Set("wpa_ft_attack", flattenWirelessControllerWidsProfileWpaFtAttack(o["wpa-ft-attack"], d, "wpa_ft_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["wpa-ft-attack"], "WirelessControllerWidsProfile-WpaFtAttack"); ok {
			if err = d.Set("wpa_ft_attack", vv); err != nil {
				return fmt.Errorf("Error reading wpa_ft_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wpa_ft_attack: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWidsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWidsProfileAdhocNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAdhocValidSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAirJack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApAutoSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWidsProfileApBgscanDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApFgscanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApImpersonation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanChannelList2G5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWidsProfileApScanChannelList6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWidsProfileApScanPassive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApSpoofing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAsleapAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFrameFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFrameFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBeaconWrongChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileChanBasedMitm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthUnknownSrcThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDisassocBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDisconnectStation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolKeyOverflow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFataJack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedProbeRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedProbeResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHotspotterAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHt40MhzIntolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHtGreenfield(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileInvalidAddrCombination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileInvalidMacOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedAssociation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedHtIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumbler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumblerThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumblerTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNullSsidProbeResp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileOmertaAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileOverflowIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePwsaveDosAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRiskyEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSensorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSpoofedDeauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileUnencryptedValid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileValidClientMisassociation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileValidSsidMisuse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWeakWepIv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiterThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiterTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWindowsBridge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWirelessBridge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWpaFtAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWidsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adhoc_network"); ok || d.HasChange("adhoc_network") {
		t, err := expandWirelessControllerWidsProfileAdhocNetwork(d, v, "adhoc_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adhoc-network"] = t
		}
	}

	if v, ok := d.GetOk("adhoc_valid_ssid"); ok || d.HasChange("adhoc_valid_ssid") {
		t, err := expandWirelessControllerWidsProfileAdhocValidSsid(d, v, "adhoc_valid_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adhoc-valid-ssid"] = t
		}
	}

	if v, ok := d.GetOk("air_jack"); ok || d.HasChange("air_jack") {
		t, err := expandWirelessControllerWidsProfileAirJack(d, v, "air_jack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["air-jack"] = t
		}
	}

	if v, ok := d.GetOk("ap_auto_suppress"); ok || d.HasChange("ap_auto_suppress") {
		t, err := expandWirelessControllerWidsProfileApAutoSuppress(d, v, "ap_auto_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-auto-suppress"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_schedules"); ok || d.HasChange("ap_bgscan_disable_schedules") {
		t, err := expandWirelessControllerWidsProfileApBgscanDisableSchedules(d, v, "ap_bgscan_disable_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-schedules"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_duration"); ok || d.HasChange("ap_bgscan_duration") {
		t, err := expandWirelessControllerWidsProfileApBgscanDuration(d, v, "ap_bgscan_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-duration"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_idle"); ok || d.HasChange("ap_bgscan_idle") {
		t, err := expandWirelessControllerWidsProfileApBgscanIdle(d, v, "ap_bgscan_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-idle"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_intv"); ok || d.HasChange("ap_bgscan_intv") {
		t, err := expandWirelessControllerWidsProfileApBgscanIntv(d, v, "ap_bgscan_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_period"); ok || d.HasChange("ap_bgscan_period") {
		t, err := expandWirelessControllerWidsProfileApBgscanPeriod(d, v, "ap_bgscan_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-period"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_report_intv"); ok || d.HasChange("ap_bgscan_report_intv") {
		t, err := expandWirelessControllerWidsProfileApBgscanReportIntv(d, v, "ap_bgscan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_fgscan_report_intv"); ok || d.HasChange("ap_fgscan_report_intv") {
		t, err := expandWirelessControllerWidsProfileApFgscanReportIntv(d, v, "ap_fgscan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-fgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_impersonation"); ok || d.HasChange("ap_impersonation") {
		t, err := expandWirelessControllerWidsProfileApImpersonation(d, v, "ap_impersonation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-impersonation"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan"); ok || d.HasChange("ap_scan") {
		t, err := expandWirelessControllerWidsProfileApScan(d, v, "ap_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_2g_5g"); ok || d.HasChange("ap_scan_channel_list_2g_5g") {
		t, err := expandWirelessControllerWidsProfileApScanChannelList2G5G(d, v, "ap_scan_channel_list_2g_5g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-2G-5G"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_6g"); ok || d.HasChange("ap_scan_channel_list_6g") {
		t, err := expandWirelessControllerWidsProfileApScanChannelList6G(d, v, "ap_scan_channel_list_6g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-6G"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_passive"); ok || d.HasChange("ap_scan_passive") {
		t, err := expandWirelessControllerWidsProfileApScanPassive(d, v, "ap_scan_passive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-passive"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_threshold"); ok || d.HasChange("ap_scan_threshold") {
		t, err := expandWirelessControllerWidsProfileApScanThreshold(d, v, "ap_scan_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ap_spoofing"); ok || d.HasChange("ap_spoofing") {
		t, err := expandWirelessControllerWidsProfileApSpoofing(d, v, "ap_spoofing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-spoofing"] = t
		}
	}

	if v, ok := d.GetOk("asleap_attack"); ok || d.HasChange("asleap_attack") {
		t, err := expandWirelessControllerWidsProfileAsleapAttack(d, v, "asleap_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asleap-attack"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_thresh"); ok || d.HasChange("assoc_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileAssocFloodThresh(d, v, "assoc_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_time"); ok || d.HasChange("assoc_flood_time") {
		t, err := expandWirelessControllerWidsProfileAssocFloodTime(d, v, "assoc_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("assoc_frame_flood"); ok || d.HasChange("assoc_frame_flood") {
		t, err := expandWirelessControllerWidsProfileAssocFrameFlood(d, v, "assoc_frame_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_thresh"); ok || d.HasChange("auth_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileAuthFloodThresh(d, v, "auth_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_time"); ok || d.HasChange("auth_flood_time") {
		t, err := expandWirelessControllerWidsProfileAuthFloodTime(d, v, "auth_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_frame_flood"); ok || d.HasChange("auth_frame_flood") {
		t, err := expandWirelessControllerWidsProfileAuthFrameFlood(d, v, "auth_frame_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood"); ok || d.HasChange("bcn_flood") {
		t, err := expandWirelessControllerWidsProfileBcnFlood(d, v, "bcn_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood_thresh"); ok || d.HasChange("bcn_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileBcnFloodThresh(d, v, "bcn_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood_time"); ok || d.HasChange("bcn_flood_time") {
		t, err := expandWirelessControllerWidsProfileBcnFloodTime(d, v, "bcn_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("beacon_wrong_channel"); ok || d.HasChange("beacon_wrong_channel") {
		t, err := expandWirelessControllerWidsProfileBeaconWrongChannel(d, v, "beacon_wrong_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-wrong-channel"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood"); ok || d.HasChange("block_ack_flood") {
		t, err := expandWirelessControllerWidsProfileBlockAckFlood(d, v, "block_ack_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood_thresh"); ok || d.HasChange("block_ack_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileBlockAckFloodThresh(d, v, "block_ack_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood_time"); ok || d.HasChange("block_ack_flood_time") {
		t, err := expandWirelessControllerWidsProfileBlockAckFloodTime(d, v, "block_ack_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("chan_based_mitm"); ok || d.HasChange("chan_based_mitm") {
		t, err := expandWirelessControllerWidsProfileChanBasedMitm(d, v, "chan_based_mitm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chan-based-mitm"] = t
		}
	}

	if v, ok := d.GetOk("client_flood"); ok || d.HasChange("client_flood") {
		t, err := expandWirelessControllerWidsProfileClientFlood(d, v, "client_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood"] = t
		}
	}

	if v, ok := d.GetOk("client_flood_thresh"); ok || d.HasChange("client_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileClientFloodThresh(d, v, "client_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("client_flood_time"); ok || d.HasChange("client_flood_time") {
		t, err := expandWirelessControllerWidsProfileClientFloodTime(d, v, "client_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerWidsProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood"); ok || d.HasChange("cts_flood") {
		t, err := expandWirelessControllerWidsProfileCtsFlood(d, v, "cts_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood_thresh"); ok || d.HasChange("cts_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileCtsFloodThresh(d, v, "cts_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood_time"); ok || d.HasChange("cts_flood_time") {
		t, err := expandWirelessControllerWidsProfileCtsFloodTime(d, v, "cts_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("deauth_broadcast"); ok || d.HasChange("deauth_broadcast") {
		t, err := expandWirelessControllerWidsProfileDeauthBroadcast(d, v, "deauth_broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("deauth_unknown_src_thresh"); ok || d.HasChange("deauth_unknown_src_thresh") {
		t, err := expandWirelessControllerWidsProfileDeauthUnknownSrcThresh(d, v, "deauth_unknown_src_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-unknown-src-thresh"] = t
		}
	}

	if v, ok := d.GetOk("disassoc_broadcast"); ok || d.HasChange("disassoc_broadcast") {
		t, err := expandWirelessControllerWidsProfileDisassocBroadcast(d, v, "disassoc_broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disassoc-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_station"); ok || d.HasChange("disconnect_station") {
		t, err := expandWirelessControllerWidsProfileDisconnectStation(d, v, "disconnect_station")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-station"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_flood"); ok || d.HasChange("eapol_fail_flood") {
		t, err := expandWirelessControllerWidsProfileEapolFailFlood(d, v, "eapol_fail_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_intv"); ok || d.HasChange("eapol_fail_intv") {
		t, err := expandWirelessControllerWidsProfileEapolFailIntv(d, v, "eapol_fail_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_thresh"); ok || d.HasChange("eapol_fail_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolFailThresh(d, v, "eapol_fail_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_overflow"); ok || d.HasChange("eapol_key_overflow") {
		t, err := expandWirelessControllerWidsProfileEapolKeyOverflow(d, v, "eapol_key_overflow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-overflow"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_flood"); ok || d.HasChange("eapol_logoff_flood") {
		t, err := expandWirelessControllerWidsProfileEapolLogoffFlood(d, v, "eapol_logoff_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_intv"); ok || d.HasChange("eapol_logoff_intv") {
		t, err := expandWirelessControllerWidsProfileEapolLogoffIntv(d, v, "eapol_logoff_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_thresh"); ok || d.HasChange("eapol_logoff_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolLogoffThresh(d, v, "eapol_logoff_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_flood"); ok || d.HasChange("eapol_pre_fail_flood") {
		t, err := expandWirelessControllerWidsProfileEapolPreFailFlood(d, v, "eapol_pre_fail_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_intv"); ok || d.HasChange("eapol_pre_fail_intv") {
		t, err := expandWirelessControllerWidsProfileEapolPreFailIntv(d, v, "eapol_pre_fail_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_thresh"); ok || d.HasChange("eapol_pre_fail_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolPreFailThresh(d, v, "eapol_pre_fail_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_flood"); ok || d.HasChange("eapol_pre_succ_flood") {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccFlood(d, v, "eapol_pre_succ_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_intv"); ok || d.HasChange("eapol_pre_succ_intv") {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccIntv(d, v, "eapol_pre_succ_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_thresh"); ok || d.HasChange("eapol_pre_succ_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccThresh(d, v, "eapol_pre_succ_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_flood"); ok || d.HasChange("eapol_start_flood") {
		t, err := expandWirelessControllerWidsProfileEapolStartFlood(d, v, "eapol_start_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_intv"); ok || d.HasChange("eapol_start_intv") {
		t, err := expandWirelessControllerWidsProfileEapolStartIntv(d, v, "eapol_start_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_thresh"); ok || d.HasChange("eapol_start_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolStartThresh(d, v, "eapol_start_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_flood"); ok || d.HasChange("eapol_succ_flood") {
		t, err := expandWirelessControllerWidsProfileEapolSuccFlood(d, v, "eapol_succ_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_intv"); ok || d.HasChange("eapol_succ_intv") {
		t, err := expandWirelessControllerWidsProfileEapolSuccIntv(d, v, "eapol_succ_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_thresh"); ok || d.HasChange("eapol_succ_thresh") {
		t, err := expandWirelessControllerWidsProfileEapolSuccThresh(d, v, "eapol_succ_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("fata_jack"); ok || d.HasChange("fata_jack") {
		t, err := expandWirelessControllerWidsProfileFataJack(d, v, "fata_jack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fata-jack"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_beacon"); ok || d.HasChange("fuzzed_beacon") {
		t, err := expandWirelessControllerWidsProfileFuzzedBeacon(d, v, "fuzzed_beacon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-beacon"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_probe_request"); ok || d.HasChange("fuzzed_probe_request") {
		t, err := expandWirelessControllerWidsProfileFuzzedProbeRequest(d, v, "fuzzed_probe_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-probe-request"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_probe_response"); ok || d.HasChange("fuzzed_probe_response") {
		t, err := expandWirelessControllerWidsProfileFuzzedProbeResponse(d, v, "fuzzed_probe_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-probe-response"] = t
		}
	}

	if v, ok := d.GetOk("hotspotter_attack"); ok || d.HasChange("hotspotter_attack") {
		t, err := expandWirelessControllerWidsProfileHotspotterAttack(d, v, "hotspotter_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspotter-attack"] = t
		}
	}

	if v, ok := d.GetOk("ht_40mhz_intolerance"); ok || d.HasChange("ht_40mhz_intolerance") {
		t, err := expandWirelessControllerWidsProfileHt40MhzIntolerance(d, v, "ht_40mhz_intolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ht-40mhz-intolerance"] = t
		}
	}

	if v, ok := d.GetOk("ht_greenfield"); ok || d.HasChange("ht_greenfield") {
		t, err := expandWirelessControllerWidsProfileHtGreenfield(d, v, "ht_greenfield")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ht-greenfield"] = t
		}
	}

	if v, ok := d.GetOk("invalid_addr_combination"); ok || d.HasChange("invalid_addr_combination") {
		t, err := expandWirelessControllerWidsProfileInvalidAddrCombination(d, v, "invalid_addr_combination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-addr-combination"] = t
		}
	}

	if v, ok := d.GetOk("invalid_mac_oui"); ok || d.HasChange("invalid_mac_oui") {
		t, err := expandWirelessControllerWidsProfileInvalidMacOui(d, v, "invalid_mac_oui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-mac-oui"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_attack"); ok || d.HasChange("long_duration_attack") {
		t, err := expandWirelessControllerWidsProfileLongDurationAttack(d, v, "long_duration_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-attack"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_thresh"); ok || d.HasChange("long_duration_thresh") {
		t, err := expandWirelessControllerWidsProfileLongDurationThresh(d, v, "long_duration_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-thresh"] = t
		}
	}

	if v, ok := d.GetOk("malformed_association"); ok || d.HasChange("malformed_association") {
		t, err := expandWirelessControllerWidsProfileMalformedAssociation(d, v, "malformed_association")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-association"] = t
		}
	}

	if v, ok := d.GetOk("malformed_auth"); ok || d.HasChange("malformed_auth") {
		t, err := expandWirelessControllerWidsProfileMalformedAuth(d, v, "malformed_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-auth"] = t
		}
	}

	if v, ok := d.GetOk("malformed_ht_ie"); ok || d.HasChange("malformed_ht_ie") {
		t, err := expandWirelessControllerWidsProfileMalformedHtIe(d, v, "malformed_ht_ie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-ht-ie"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerWidsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler"); ok || d.HasChange("netstumbler") {
		t, err := expandWirelessControllerWidsProfileNetstumbler(d, v, "netstumbler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler_thresh"); ok || d.HasChange("netstumbler_thresh") {
		t, err := expandWirelessControllerWidsProfileNetstumblerThresh(d, v, "netstumbler_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler-thresh"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler_time"); ok || d.HasChange("netstumbler_time") {
		t, err := expandWirelessControllerWidsProfileNetstumblerTime(d, v, "netstumbler_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler-time"] = t
		}
	}

	if v, ok := d.GetOk("null_ssid_probe_resp"); ok || d.HasChange("null_ssid_probe_resp") {
		t, err := expandWirelessControllerWidsProfileNullSsidProbeResp(d, v, "null_ssid_probe_resp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-ssid-probe-resp"] = t
		}
	}

	if v, ok := d.GetOk("omerta_attack"); ok || d.HasChange("omerta_attack") {
		t, err := expandWirelessControllerWidsProfileOmertaAttack(d, v, "omerta_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["omerta-attack"] = t
		}
	}

	if v, ok := d.GetOk("overflow_ie"); ok || d.HasChange("overflow_ie") {
		t, err := expandWirelessControllerWidsProfileOverflowIe(d, v, "overflow_ie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overflow-ie"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood"); ok || d.HasChange("probe_flood") {
		t, err := expandWirelessControllerWidsProfileProbeFlood(d, v, "probe_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood_thresh"); ok || d.HasChange("probe_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileProbeFloodThresh(d, v, "probe_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood_time"); ok || d.HasChange("probe_flood_time") {
		t, err := expandWirelessControllerWidsProfileProbeFloodTime(d, v, "probe_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood"); ok || d.HasChange("pspoll_flood") {
		t, err := expandWirelessControllerWidsProfilePspollFlood(d, v, "pspoll_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood_thresh"); ok || d.HasChange("pspoll_flood_thresh") {
		t, err := expandWirelessControllerWidsProfilePspollFloodThresh(d, v, "pspoll_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood_time"); ok || d.HasChange("pspoll_flood_time") {
		t, err := expandWirelessControllerWidsProfilePspollFloodTime(d, v, "pspoll_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("pwsave_dos_attack"); ok || d.HasChange("pwsave_dos_attack") {
		t, err := expandWirelessControllerWidsProfilePwsaveDosAttack(d, v, "pwsave_dos_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pwsave-dos-attack"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood"); ok || d.HasChange("reassoc_flood") {
		t, err := expandWirelessControllerWidsProfileReassocFlood(d, v, "reassoc_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood_thresh"); ok || d.HasChange("reassoc_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileReassocFloodThresh(d, v, "reassoc_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood_time"); ok || d.HasChange("reassoc_flood_time") {
		t, err := expandWirelessControllerWidsProfileReassocFloodTime(d, v, "reassoc_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("risky_encryption"); ok || d.HasChange("risky_encryption") {
		t, err := expandWirelessControllerWidsProfileRiskyEncryption(d, v, "risky_encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risky-encryption"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood"); ok || d.HasChange("rts_flood") {
		t, err := expandWirelessControllerWidsProfileRtsFlood(d, v, "rts_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood_thresh"); ok || d.HasChange("rts_flood_thresh") {
		t, err := expandWirelessControllerWidsProfileRtsFloodThresh(d, v, "rts_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood_time"); ok || d.HasChange("rts_flood_time") {
		t, err := expandWirelessControllerWidsProfileRtsFloodTime(d, v, "rts_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("sensor_mode"); ok || d.HasChange("sensor_mode") {
		t, err := expandWirelessControllerWidsProfileSensorMode(d, v, "sensor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensor-mode"] = t
		}
	}

	if v, ok := d.GetOk("spoofed_deauth"); ok || d.HasChange("spoofed_deauth") {
		t, err := expandWirelessControllerWidsProfileSpoofedDeauth(d, v, "spoofed_deauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoofed-deauth"] = t
		}
	}

	if v, ok := d.GetOk("unencrypted_valid"); ok || d.HasChange("unencrypted_valid") {
		t, err := expandWirelessControllerWidsProfileUnencryptedValid(d, v, "unencrypted_valid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unencrypted-valid"] = t
		}
	}

	if v, ok := d.GetOk("valid_client_misassociation"); ok || d.HasChange("valid_client_misassociation") {
		t, err := expandWirelessControllerWidsProfileValidClientMisassociation(d, v, "valid_client_misassociation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-client-misassociation"] = t
		}
	}

	if v, ok := d.GetOk("valid_ssid_misuse"); ok || d.HasChange("valid_ssid_misuse") {
		t, err := expandWirelessControllerWidsProfileValidSsidMisuse(d, v, "valid_ssid_misuse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-ssid-misuse"] = t
		}
	}

	if v, ok := d.GetOk("weak_wep_iv"); ok || d.HasChange("weak_wep_iv") {
		t, err := expandWirelessControllerWidsProfileWeakWepIv(d, v, "weak_wep_iv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weak-wep-iv"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter"); ok || d.HasChange("wellenreiter") {
		t, err := expandWirelessControllerWidsProfileWellenreiter(d, v, "wellenreiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter_thresh"); ok || d.HasChange("wellenreiter_thresh") {
		t, err := expandWirelessControllerWidsProfileWellenreiterThresh(d, v, "wellenreiter_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter-thresh"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter_time"); ok || d.HasChange("wellenreiter_time") {
		t, err := expandWirelessControllerWidsProfileWellenreiterTime(d, v, "wellenreiter_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter-time"] = t
		}
	}

	if v, ok := d.GetOk("windows_bridge"); ok || d.HasChange("windows_bridge") {
		t, err := expandWirelessControllerWidsProfileWindowsBridge(d, v, "windows_bridge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-bridge"] = t
		}
	}

	if v, ok := d.GetOk("wireless_bridge"); ok || d.HasChange("wireless_bridge") {
		t, err := expandWirelessControllerWidsProfileWirelessBridge(d, v, "wireless_bridge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-bridge"] = t
		}
	}

	if v, ok := d.GetOk("wpa_ft_attack"); ok || d.HasChange("wpa_ft_attack") {
		t, err := expandWirelessControllerWidsProfileWpaFtAttack(d, v, "wpa_ft_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wpa-ft-attack"] = t
		}
	}

	return &obj, nil
}
