// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration options for radio 2.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpProfileRadio2() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfileRadio2Update,
		Read:   resourceWirelessControllerWtpProfileRadio2Read,
		Update: resourceWirelessControllerWtpProfileRadio2Update,
		Delete: resourceWirelessControllerWtpProfileRadio2Delete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"n80211d": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n80211mc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"airtime_fairness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"amsdu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_handoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_sniffer_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_bufsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_sniffer_chan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_sniffer_chan_width": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_ctl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_mgmt_beacon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_mgmt_other": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_sniffer_mgmt_probe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arrp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_power_high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_power_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_power_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_power_target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"band": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"band_5g_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bss_color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bss_color_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"channel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"channel_bonding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"channel_bonding_ext": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"channel_utilization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"coexistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"darrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drma": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drma_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtim": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"frag_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"frequency_handoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iperf_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iperf_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mimo_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"optional_antenna": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"optional_antenna_gain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"power_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"powersave_optimize": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rts_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sam_bssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_ca_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sam_captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_client_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sam_cwp_failure_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_cwp_match_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_cwp_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sam_cwp_success_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_cwp_test_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_cwp_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_eap_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sam_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sam_private_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sam_private_key_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sam_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sam_security_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_test": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sam_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"short_guard_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spectrum_analysis": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transmit_optimize": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vap_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vap1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vap8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vaps": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wids_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"zero_wait_dfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileRadio2Update(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectWirelessControllerWtpProfileRadio2(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileRadio2 resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpProfileRadio2(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileRadio2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpProfileRadio2")

	return resourceWirelessControllerWtpProfileRadio2Read(d, m)
}

func resourceWirelessControllerWtpProfileRadio2Delete(d *schema.ResourceData, m interface{}) error {
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
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	err = c.DeleteWirelessControllerWtpProfileRadio2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfileRadio2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileRadio2Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	wtp_profile := d.Get("wtp_profile").(string)
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
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadWirelessControllerWtpProfileRadio2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileRadio2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfileRadio2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileRadio2 resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileRadio280211D2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio280211Mc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AirtimeFairness2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Amsdu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApHandoff2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferChan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferChanWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferCtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferData2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ArrpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerHigh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Band2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpProfileRadio2Band5GType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthCapacity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BeaconInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BssColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BssColorMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallAdmissionControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallCapacity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Channel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2ChannelBonding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelBondingExt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelUtilization2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Coexistence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Darrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Drma2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2DrmaSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Dtim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FragThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FrequencyHandoff2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2IperfProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2IperfServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxDistance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MimoMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2OptionalAntenna2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2OptionalAntennaGain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowersaveOptimize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2ProtectionMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2RadioId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2RtsThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamBssid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCaCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2SamCaptivePortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamClientCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2SamCwpFailureString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpMatchString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpSuccessString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpTestUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamEapMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamPrivateKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2SamReportIntv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamSecurityType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamSsid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamTest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ShortGuardInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2TransmitOptimize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2VapAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vap82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vaps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2WidsProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileRadio2ZeroWaitDfs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfileRadio2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n80211d", flattenWirelessControllerWtpProfileRadio280211D2edl(o["80211d"], d, "n80211d")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211d"], "WirelessControllerWtpProfileRadio2-80211D"); ok {
			if err = d.Set("n80211d", vv); err != nil {
				return fmt.Errorf("Error reading n80211d: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211d: %v", err)
		}
	}

	if err = d.Set("n80211mc", flattenWirelessControllerWtpProfileRadio280211Mc2edl(o["80211mc"], d, "n80211mc")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211mc"], "WirelessControllerWtpProfileRadio2-80211Mc"); ok {
			if err = d.Set("n80211mc", vv); err != nil {
				return fmt.Errorf("Error reading n80211mc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211mc: %v", err)
		}
	}

	if err = d.Set("airtime_fairness", flattenWirelessControllerWtpProfileRadio2AirtimeFairness2edl(o["airtime-fairness"], d, "airtime_fairness")); err != nil {
		if vv, ok := fortiAPIPatch(o["airtime-fairness"], "WirelessControllerWtpProfileRadio2-AirtimeFairness"); ok {
			if err = d.Set("airtime_fairness", vv); err != nil {
				return fmt.Errorf("Error reading airtime_fairness: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading airtime_fairness: %v", err)
		}
	}

	if err = d.Set("amsdu", flattenWirelessControllerWtpProfileRadio2Amsdu2edl(o["amsdu"], d, "amsdu")); err != nil {
		if vv, ok := fortiAPIPatch(o["amsdu"], "WirelessControllerWtpProfileRadio2-Amsdu"); ok {
			if err = d.Set("amsdu", vv); err != nil {
				return fmt.Errorf("Error reading amsdu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading amsdu: %v", err)
		}
	}

	if err = d.Set("ap_handoff", flattenWirelessControllerWtpProfileRadio2ApHandoff2edl(o["ap-handoff"], d, "ap_handoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-handoff"], "WirelessControllerWtpProfileRadio2-ApHandoff"); ok {
			if err = d.Set("ap_handoff", vv); err != nil {
				return fmt.Errorf("Error reading ap_handoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_handoff: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_addr", flattenWirelessControllerWtpProfileRadio2ApSnifferAddr2edl(o["ap-sniffer-addr"], d, "ap_sniffer_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-addr"], "WirelessControllerWtpProfileRadio2-ApSnifferAddr"); ok {
			if err = d.Set("ap_sniffer_addr", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_addr: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_bufsize", flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize2edl(o["ap-sniffer-bufsize"], d, "ap_sniffer_bufsize")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-bufsize"], "WirelessControllerWtpProfileRadio2-ApSnifferBufsize"); ok {
			if err = d.Set("ap_sniffer_bufsize", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_bufsize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_bufsize: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_chan", flattenWirelessControllerWtpProfileRadio2ApSnifferChan2edl(o["ap-sniffer-chan"], d, "ap_sniffer_chan")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-chan"], "WirelessControllerWtpProfileRadio2-ApSnifferChan"); ok {
			if err = d.Set("ap_sniffer_chan", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_chan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_chan: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_chan_width", flattenWirelessControllerWtpProfileRadio2ApSnifferChanWidth2edl(o["ap-sniffer-chan-width"], d, "ap_sniffer_chan_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-chan-width"], "WirelessControllerWtpProfileRadio2-ApSnifferChanWidth"); ok {
			if err = d.Set("ap_sniffer_chan_width", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_chan_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_chan_width: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_ctl", flattenWirelessControllerWtpProfileRadio2ApSnifferCtl2edl(o["ap-sniffer-ctl"], d, "ap_sniffer_ctl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-ctl"], "WirelessControllerWtpProfileRadio2-ApSnifferCtl"); ok {
			if err = d.Set("ap_sniffer_ctl", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_ctl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_ctl: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_data", flattenWirelessControllerWtpProfileRadio2ApSnifferData2edl(o["ap-sniffer-data"], d, "ap_sniffer_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-data"], "WirelessControllerWtpProfileRadio2-ApSnifferData"); ok {
			if err = d.Set("ap_sniffer_data", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_data: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_mgmt_beacon", flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon2edl(o["ap-sniffer-mgmt-beacon"], d, "ap_sniffer_mgmt_beacon")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-mgmt-beacon"], "WirelessControllerWtpProfileRadio2-ApSnifferMgmtBeacon"); ok {
			if err = d.Set("ap_sniffer_mgmt_beacon", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_mgmt_beacon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_mgmt_beacon: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_mgmt_other", flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther2edl(o["ap-sniffer-mgmt-other"], d, "ap_sniffer_mgmt_other")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-mgmt-other"], "WirelessControllerWtpProfileRadio2-ApSnifferMgmtOther"); ok {
			if err = d.Set("ap_sniffer_mgmt_other", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_mgmt_other: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_mgmt_other: %v", err)
		}
	}

	if err = d.Set("ap_sniffer_mgmt_probe", flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe2edl(o["ap-sniffer-mgmt-probe"], d, "ap_sniffer_mgmt_probe")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-sniffer-mgmt-probe"], "WirelessControllerWtpProfileRadio2-ApSnifferMgmtProbe"); ok {
			if err = d.Set("ap_sniffer_mgmt_probe", vv); err != nil {
				return fmt.Errorf("Error reading ap_sniffer_mgmt_probe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_sniffer_mgmt_probe: %v", err)
		}
	}

	if err = d.Set("arrp_profile", flattenWirelessControllerWtpProfileRadio2ArrpProfile2edl(o["arrp-profile"], d, "arrp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["arrp-profile"], "WirelessControllerWtpProfileRadio2-ArrpProfile"); ok {
			if err = d.Set("arrp_profile", vv); err != nil {
				return fmt.Errorf("Error reading arrp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arrp_profile: %v", err)
		}
	}

	if err = d.Set("auto_power_high", flattenWirelessControllerWtpProfileRadio2AutoPowerHigh2edl(o["auto-power-high"], d, "auto_power_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-high"], "WirelessControllerWtpProfileRadio2-AutoPowerHigh"); ok {
			if err = d.Set("auto_power_high", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_high: %v", err)
		}
	}

	if err = d.Set("auto_power_level", flattenWirelessControllerWtpProfileRadio2AutoPowerLevel2edl(o["auto-power-level"], d, "auto_power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-level"], "WirelessControllerWtpProfileRadio2-AutoPowerLevel"); ok {
			if err = d.Set("auto_power_level", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_level: %v", err)
		}
	}

	if err = d.Set("auto_power_low", flattenWirelessControllerWtpProfileRadio2AutoPowerLow2edl(o["auto-power-low"], d, "auto_power_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-low"], "WirelessControllerWtpProfileRadio2-AutoPowerLow"); ok {
			if err = d.Set("auto_power_low", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_low: %v", err)
		}
	}

	if err = d.Set("auto_power_target", flattenWirelessControllerWtpProfileRadio2AutoPowerTarget2edl(o["auto-power-target"], d, "auto_power_target")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-power-target"], "WirelessControllerWtpProfileRadio2-AutoPowerTarget"); ok {
			if err = d.Set("auto_power_target", vv); err != nil {
				return fmt.Errorf("Error reading auto_power_target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_power_target: %v", err)
		}
	}

	if err = d.Set("band", flattenWirelessControllerWtpProfileRadio2Band2edl(o["band"], d, "band")); err != nil {
		if vv, ok := fortiAPIPatch(o["band"], "WirelessControllerWtpProfileRadio2-Band"); ok {
			if err = d.Set("band", vv); err != nil {
				return fmt.Errorf("Error reading band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading band: %v", err)
		}
	}

	if err = d.Set("band_5g_type", flattenWirelessControllerWtpProfileRadio2Band5GType2edl(o["band-5g-type"], d, "band_5g_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["band-5g-type"], "WirelessControllerWtpProfileRadio2-Band5GType"); ok {
			if err = d.Set("band_5g_type", vv); err != nil {
				return fmt.Errorf("Error reading band_5g_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading band_5g_type: %v", err)
		}
	}

	if err = d.Set("bandwidth_admission_control", flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl2edl(o["bandwidth-admission-control"], d, "bandwidth_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-admission-control"], "WirelessControllerWtpProfileRadio2-BandwidthAdmissionControl"); ok {
			if err = d.Set("bandwidth_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
		}
	}

	if err = d.Set("bandwidth_capacity", flattenWirelessControllerWtpProfileRadio2BandwidthCapacity2edl(o["bandwidth-capacity"], d, "bandwidth_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-capacity"], "WirelessControllerWtpProfileRadio2-BandwidthCapacity"); ok {
			if err = d.Set("bandwidth_capacity", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenWirelessControllerWtpProfileRadio2BeaconInterval2edl(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-interval"], "WirelessControllerWtpProfileRadio2-BeaconInterval"); ok {
			if err = d.Set("beacon_interval", vv); err != nil {
				return fmt.Errorf("Error reading beacon_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("bss_color", flattenWirelessControllerWtpProfileRadio2BssColor2edl(o["bss-color"], d, "bss_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color"], "WirelessControllerWtpProfileRadio2-BssColor"); ok {
			if err = d.Set("bss_color", vv); err != nil {
				return fmt.Errorf("Error reading bss_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color: %v", err)
		}
	}

	if err = d.Set("bss_color_mode", flattenWirelessControllerWtpProfileRadio2BssColorMode2edl(o["bss-color-mode"], d, "bss_color_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-mode"], "WirelessControllerWtpProfileRadio2-BssColorMode"); ok {
			if err = d.Set("bss_color_mode", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_mode: %v", err)
		}
	}

	if err = d.Set("call_admission_control", flattenWirelessControllerWtpProfileRadio2CallAdmissionControl2edl(o["call-admission-control"], d, "call_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-admission-control"], "WirelessControllerWtpProfileRadio2-CallAdmissionControl"); ok {
			if err = d.Set("call_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading call_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_admission_control: %v", err)
		}
	}

	if err = d.Set("call_capacity", flattenWirelessControllerWtpProfileRadio2CallCapacity2edl(o["call-capacity"], d, "call_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-capacity"], "WirelessControllerWtpProfileRadio2-CallCapacity"); ok {
			if err = d.Set("call_capacity", vv); err != nil {
				return fmt.Errorf("Error reading call_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_capacity: %v", err)
		}
	}

	if err = d.Set("channel", flattenWirelessControllerWtpProfileRadio2Channel2edl(o["channel"], d, "channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel"], "WirelessControllerWtpProfileRadio2-Channel"); ok {
			if err = d.Set("channel", vv); err != nil {
				return fmt.Errorf("Error reading channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel: %v", err)
		}
	}

	if err = d.Set("channel_bonding", flattenWirelessControllerWtpProfileRadio2ChannelBonding2edl(o["channel-bonding"], d, "channel_bonding")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel-bonding"], "WirelessControllerWtpProfileRadio2-ChannelBonding"); ok {
			if err = d.Set("channel_bonding", vv); err != nil {
				return fmt.Errorf("Error reading channel_bonding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel_bonding: %v", err)
		}
	}

	if err = d.Set("channel_bonding_ext", flattenWirelessControllerWtpProfileRadio2ChannelBondingExt2edl(o["channel-bonding-ext"], d, "channel_bonding_ext")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel-bonding-ext"], "WirelessControllerWtpProfileRadio2-ChannelBondingExt"); ok {
			if err = d.Set("channel_bonding_ext", vv); err != nil {
				return fmt.Errorf("Error reading channel_bonding_ext: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel_bonding_ext: %v", err)
		}
	}

	if err = d.Set("channel_utilization", flattenWirelessControllerWtpProfileRadio2ChannelUtilization2edl(o["channel-utilization"], d, "channel_utilization")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel-utilization"], "WirelessControllerWtpProfileRadio2-ChannelUtilization"); ok {
			if err = d.Set("channel_utilization", vv); err != nil {
				return fmt.Errorf("Error reading channel_utilization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel_utilization: %v", err)
		}
	}

	if err = d.Set("coexistence", flattenWirelessControllerWtpProfileRadio2Coexistence2edl(o["coexistence"], d, "coexistence")); err != nil {
		if vv, ok := fortiAPIPatch(o["coexistence"], "WirelessControllerWtpProfileRadio2-Coexistence"); ok {
			if err = d.Set("coexistence", vv); err != nil {
				return fmt.Errorf("Error reading coexistence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coexistence: %v", err)
		}
	}

	if err = d.Set("darrp", flattenWirelessControllerWtpProfileRadio2Darrp2edl(o["darrp"], d, "darrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp"], "WirelessControllerWtpProfileRadio2-Darrp"); ok {
			if err = d.Set("darrp", vv); err != nil {
				return fmt.Errorf("Error reading darrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp: %v", err)
		}
	}

	if err = d.Set("drma", flattenWirelessControllerWtpProfileRadio2Drma2edl(o["drma"], d, "drma")); err != nil {
		if vv, ok := fortiAPIPatch(o["drma"], "WirelessControllerWtpProfileRadio2-Drma"); ok {
			if err = d.Set("drma", vv); err != nil {
				return fmt.Errorf("Error reading drma: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drma: %v", err)
		}
	}

	if err = d.Set("drma_sensitivity", flattenWirelessControllerWtpProfileRadio2DrmaSensitivity2edl(o["drma-sensitivity"], d, "drma_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["drma-sensitivity"], "WirelessControllerWtpProfileRadio2-DrmaSensitivity"); ok {
			if err = d.Set("drma_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading drma_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drma_sensitivity: %v", err)
		}
	}

	if err = d.Set("dtim", flattenWirelessControllerWtpProfileRadio2Dtim2edl(o["dtim"], d, "dtim")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtim"], "WirelessControllerWtpProfileRadio2-Dtim"); ok {
			if err = d.Set("dtim", vv); err != nil {
				return fmt.Errorf("Error reading dtim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtim: %v", err)
		}
	}

	if err = d.Set("frag_threshold", flattenWirelessControllerWtpProfileRadio2FragThreshold2edl(o["frag-threshold"], d, "frag_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["frag-threshold"], "WirelessControllerWtpProfileRadio2-FragThreshold"); ok {
			if err = d.Set("frag_threshold", vv); err != nil {
				return fmt.Errorf("Error reading frag_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frag_threshold: %v", err)
		}
	}

	if err = d.Set("frequency_handoff", flattenWirelessControllerWtpProfileRadio2FrequencyHandoff2edl(o["frequency-handoff"], d, "frequency_handoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["frequency-handoff"], "WirelessControllerWtpProfileRadio2-FrequencyHandoff"); ok {
			if err = d.Set("frequency_handoff", vv); err != nil {
				return fmt.Errorf("Error reading frequency_handoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frequency_handoff: %v", err)
		}
	}

	if err = d.Set("iperf_protocol", flattenWirelessControllerWtpProfileRadio2IperfProtocol2edl(o["iperf-protocol"], d, "iperf_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["iperf-protocol"], "WirelessControllerWtpProfileRadio2-IperfProtocol"); ok {
			if err = d.Set("iperf_protocol", vv); err != nil {
				return fmt.Errorf("Error reading iperf_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iperf_protocol: %v", err)
		}
	}

	if err = d.Set("iperf_server_port", flattenWirelessControllerWtpProfileRadio2IperfServerPort2edl(o["iperf-server-port"], d, "iperf_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["iperf-server-port"], "WirelessControllerWtpProfileRadio2-IperfServerPort"); ok {
			if err = d.Set("iperf_server_port", vv); err != nil {
				return fmt.Errorf("Error reading iperf_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iperf_server_port: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerWtpProfileRadio2MaxClients2edl(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "WirelessControllerWtpProfileRadio2-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_distance", flattenWirelessControllerWtpProfileRadio2MaxDistance2edl(o["max-distance"], d, "max_distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-distance"], "WirelessControllerWtpProfileRadio2-MaxDistance"); ok {
			if err = d.Set("max_distance", vv); err != nil {
				return fmt.Errorf("Error reading max_distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_distance: %v", err)
		}
	}

	if err = d.Set("mimo_mode", flattenWirelessControllerWtpProfileRadio2MimoMode2edl(o["mimo-mode"], d, "mimo_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mimo-mode"], "WirelessControllerWtpProfileRadio2-MimoMode"); ok {
			if err = d.Set("mimo_mode", vv); err != nil {
				return fmt.Errorf("Error reading mimo_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mimo_mode: %v", err)
		}
	}

	if err = d.Set("mode", flattenWirelessControllerWtpProfileRadio2Mode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "WirelessControllerWtpProfileRadio2-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("optional_antenna", flattenWirelessControllerWtpProfileRadio2OptionalAntenna2edl(o["optional-antenna"], d, "optional_antenna")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional-antenna"], "WirelessControllerWtpProfileRadio2-OptionalAntenna"); ok {
			if err = d.Set("optional_antenna", vv); err != nil {
				return fmt.Errorf("Error reading optional_antenna: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional_antenna: %v", err)
		}
	}

	if err = d.Set("optional_antenna_gain", flattenWirelessControllerWtpProfileRadio2OptionalAntennaGain2edl(o["optional-antenna-gain"], d, "optional_antenna_gain")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional-antenna-gain"], "WirelessControllerWtpProfileRadio2-OptionalAntennaGain"); ok {
			if err = d.Set("optional_antenna_gain", vv); err != nil {
				return fmt.Errorf("Error reading optional_antenna_gain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional_antenna_gain: %v", err)
		}
	}

	if err = d.Set("power_level", flattenWirelessControllerWtpProfileRadio2PowerLevel2edl(o["power-level"], d, "power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-level"], "WirelessControllerWtpProfileRadio2-PowerLevel"); ok {
			if err = d.Set("power_level", vv); err != nil {
				return fmt.Errorf("Error reading power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_level: %v", err)
		}
	}

	if err = d.Set("power_mode", flattenWirelessControllerWtpProfileRadio2PowerMode2edl(o["power-mode"], d, "power_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-mode"], "WirelessControllerWtpProfileRadio2-PowerMode"); ok {
			if err = d.Set("power_mode", vv); err != nil {
				return fmt.Errorf("Error reading power_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_mode: %v", err)
		}
	}

	if err = d.Set("power_value", flattenWirelessControllerWtpProfileRadio2PowerValue2edl(o["power-value"], d, "power_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-value"], "WirelessControllerWtpProfileRadio2-PowerValue"); ok {
			if err = d.Set("power_value", vv); err != nil {
				return fmt.Errorf("Error reading power_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_value: %v", err)
		}
	}

	if err = d.Set("powersave_optimize", flattenWirelessControllerWtpProfileRadio2PowersaveOptimize2edl(o["powersave-optimize"], d, "powersave_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["powersave-optimize"], "WirelessControllerWtpProfileRadio2-PowersaveOptimize"); ok {
			if err = d.Set("powersave_optimize", vv); err != nil {
				return fmt.Errorf("Error reading powersave_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading powersave_optimize: %v", err)
		}
	}

	if err = d.Set("protection_mode", flattenWirelessControllerWtpProfileRadio2ProtectionMode2edl(o["protection-mode"], d, "protection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["protection-mode"], "WirelessControllerWtpProfileRadio2-ProtectionMode"); ok {
			if err = d.Set("protection_mode", vv); err != nil {
				return fmt.Errorf("Error reading protection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protection_mode: %v", err)
		}
	}

	if err = d.Set("radio_id", flattenWirelessControllerWtpProfileRadio2RadioId2edl(o["radio-id"], d, "radio_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-id"], "WirelessControllerWtpProfileRadio2-RadioId"); ok {
			if err = d.Set("radio_id", vv); err != nil {
				return fmt.Errorf("Error reading radio_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_id: %v", err)
		}
	}

	if err = d.Set("rts_threshold", flattenWirelessControllerWtpProfileRadio2RtsThreshold2edl(o["rts-threshold"], d, "rts_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["rts-threshold"], "WirelessControllerWtpProfileRadio2-RtsThreshold"); ok {
			if err = d.Set("rts_threshold", vv); err != nil {
				return fmt.Errorf("Error reading rts_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rts_threshold: %v", err)
		}
	}

	if err = d.Set("sam_bssid", flattenWirelessControllerWtpProfileRadio2SamBssid2edl(o["sam-bssid"], d, "sam_bssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-bssid"], "WirelessControllerWtpProfileRadio2-SamBssid"); ok {
			if err = d.Set("sam_bssid", vv); err != nil {
				return fmt.Errorf("Error reading sam_bssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_bssid: %v", err)
		}
	}

	if err = d.Set("sam_ca_certificate", flattenWirelessControllerWtpProfileRadio2SamCaCertificate2edl(o["sam-ca-certificate"], d, "sam_ca_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-ca-certificate"], "WirelessControllerWtpProfileRadio2-SamCaCertificate"); ok {
			if err = d.Set("sam_ca_certificate", vv); err != nil {
				return fmt.Errorf("Error reading sam_ca_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_ca_certificate: %v", err)
		}
	}

	if err = d.Set("sam_captive_portal", flattenWirelessControllerWtpProfileRadio2SamCaptivePortal2edl(o["sam-captive-portal"], d, "sam_captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-captive-portal"], "WirelessControllerWtpProfileRadio2-SamCaptivePortal"); ok {
			if err = d.Set("sam_captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading sam_captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_captive_portal: %v", err)
		}
	}

	if err = d.Set("sam_client_certificate", flattenWirelessControllerWtpProfileRadio2SamClientCertificate2edl(o["sam-client-certificate"], d, "sam_client_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-client-certificate"], "WirelessControllerWtpProfileRadio2-SamClientCertificate"); ok {
			if err = d.Set("sam_client_certificate", vv); err != nil {
				return fmt.Errorf("Error reading sam_client_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_client_certificate: %v", err)
		}
	}

	if err = d.Set("sam_cwp_failure_string", flattenWirelessControllerWtpProfileRadio2SamCwpFailureString2edl(o["sam-cwp-failure-string"], d, "sam_cwp_failure_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-cwp-failure-string"], "WirelessControllerWtpProfileRadio2-SamCwpFailureString"); ok {
			if err = d.Set("sam_cwp_failure_string", vv); err != nil {
				return fmt.Errorf("Error reading sam_cwp_failure_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_cwp_failure_string: %v", err)
		}
	}

	if err = d.Set("sam_cwp_match_string", flattenWirelessControllerWtpProfileRadio2SamCwpMatchString2edl(o["sam-cwp-match-string"], d, "sam_cwp_match_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-cwp-match-string"], "WirelessControllerWtpProfileRadio2-SamCwpMatchString"); ok {
			if err = d.Set("sam_cwp_match_string", vv); err != nil {
				return fmt.Errorf("Error reading sam_cwp_match_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_cwp_match_string: %v", err)
		}
	}

	if err = d.Set("sam_cwp_success_string", flattenWirelessControllerWtpProfileRadio2SamCwpSuccessString2edl(o["sam-cwp-success-string"], d, "sam_cwp_success_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-cwp-success-string"], "WirelessControllerWtpProfileRadio2-SamCwpSuccessString"); ok {
			if err = d.Set("sam_cwp_success_string", vv); err != nil {
				return fmt.Errorf("Error reading sam_cwp_success_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_cwp_success_string: %v", err)
		}
	}

	if err = d.Set("sam_cwp_test_url", flattenWirelessControllerWtpProfileRadio2SamCwpTestUrl2edl(o["sam-cwp-test-url"], d, "sam_cwp_test_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-cwp-test-url"], "WirelessControllerWtpProfileRadio2-SamCwpTestUrl"); ok {
			if err = d.Set("sam_cwp_test_url", vv); err != nil {
				return fmt.Errorf("Error reading sam_cwp_test_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_cwp_test_url: %v", err)
		}
	}

	if err = d.Set("sam_cwp_username", flattenWirelessControllerWtpProfileRadio2SamCwpUsername2edl(o["sam-cwp-username"], d, "sam_cwp_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-cwp-username"], "WirelessControllerWtpProfileRadio2-SamCwpUsername"); ok {
			if err = d.Set("sam_cwp_username", vv); err != nil {
				return fmt.Errorf("Error reading sam_cwp_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_cwp_username: %v", err)
		}
	}

	if err = d.Set("sam_eap_method", flattenWirelessControllerWtpProfileRadio2SamEapMethod2edl(o["sam-eap-method"], d, "sam_eap_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-eap-method"], "WirelessControllerWtpProfileRadio2-SamEapMethod"); ok {
			if err = d.Set("sam_eap_method", vv); err != nil {
				return fmt.Errorf("Error reading sam_eap_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_eap_method: %v", err)
		}
	}

	if err = d.Set("sam_private_key", flattenWirelessControllerWtpProfileRadio2SamPrivateKey2edl(o["sam-private-key"], d, "sam_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-private-key"], "WirelessControllerWtpProfileRadio2-SamPrivateKey"); ok {
			if err = d.Set("sam_private_key", vv); err != nil {
				return fmt.Errorf("Error reading sam_private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_private_key: %v", err)
		}
	}

	if err = d.Set("sam_report_intv", flattenWirelessControllerWtpProfileRadio2SamReportIntv2edl(o["sam-report-intv"], d, "sam_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-report-intv"], "WirelessControllerWtpProfileRadio2-SamReportIntv"); ok {
			if err = d.Set("sam_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading sam_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_report_intv: %v", err)
		}
	}

	if err = d.Set("sam_security_type", flattenWirelessControllerWtpProfileRadio2SamSecurityType2edl(o["sam-security-type"], d, "sam_security_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-security-type"], "WirelessControllerWtpProfileRadio2-SamSecurityType"); ok {
			if err = d.Set("sam_security_type", vv); err != nil {
				return fmt.Errorf("Error reading sam_security_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_security_type: %v", err)
		}
	}

	if err = d.Set("sam_server", flattenWirelessControllerWtpProfileRadio2SamServer2edl(o["sam-server"], d, "sam_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-server"], "WirelessControllerWtpProfileRadio2-SamServer"); ok {
			if err = d.Set("sam_server", vv); err != nil {
				return fmt.Errorf("Error reading sam_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_server: %v", err)
		}
	}

	if err = d.Set("sam_server_fqdn", flattenWirelessControllerWtpProfileRadio2SamServerFqdn2edl(o["sam-server-fqdn"], d, "sam_server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-server-fqdn"], "WirelessControllerWtpProfileRadio2-SamServerFqdn"); ok {
			if err = d.Set("sam_server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading sam_server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_server_fqdn: %v", err)
		}
	}

	if err = d.Set("sam_server_ip", flattenWirelessControllerWtpProfileRadio2SamServerIp2edl(o["sam-server-ip"], d, "sam_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-server-ip"], "WirelessControllerWtpProfileRadio2-SamServerIp"); ok {
			if err = d.Set("sam_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading sam_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_server_ip: %v", err)
		}
	}

	if err = d.Set("sam_server_type", flattenWirelessControllerWtpProfileRadio2SamServerType2edl(o["sam-server-type"], d, "sam_server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-server-type"], "WirelessControllerWtpProfileRadio2-SamServerType"); ok {
			if err = d.Set("sam_server_type", vv); err != nil {
				return fmt.Errorf("Error reading sam_server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_server_type: %v", err)
		}
	}

	if err = d.Set("sam_ssid", flattenWirelessControllerWtpProfileRadio2SamSsid2edl(o["sam-ssid"], d, "sam_ssid")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-ssid"], "WirelessControllerWtpProfileRadio2-SamSsid"); ok {
			if err = d.Set("sam_ssid", vv); err != nil {
				return fmt.Errorf("Error reading sam_ssid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_ssid: %v", err)
		}
	}

	if err = d.Set("sam_test", flattenWirelessControllerWtpProfileRadio2SamTest2edl(o["sam-test"], d, "sam_test")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-test"], "WirelessControllerWtpProfileRadio2-SamTest"); ok {
			if err = d.Set("sam_test", vv); err != nil {
				return fmt.Errorf("Error reading sam_test: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_test: %v", err)
		}
	}

	if err = d.Set("sam_username", flattenWirelessControllerWtpProfileRadio2SamUsername2edl(o["sam-username"], d, "sam_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sam-username"], "WirelessControllerWtpProfileRadio2-SamUsername"); ok {
			if err = d.Set("sam_username", vv); err != nil {
				return fmt.Errorf("Error reading sam_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sam_username: %v", err)
		}
	}

	if err = d.Set("short_guard_interval", flattenWirelessControllerWtpProfileRadio2ShortGuardInterval2edl(o["short-guard-interval"], d, "short_guard_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["short-guard-interval"], "WirelessControllerWtpProfileRadio2-ShortGuardInterval"); ok {
			if err = d.Set("short_guard_interval", vv); err != nil {
				return fmt.Errorf("Error reading short_guard_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading short_guard_interval: %v", err)
		}
	}

	if err = d.Set("spectrum_analysis", flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis2edl(o["spectrum-analysis"], d, "spectrum_analysis")); err != nil {
		if vv, ok := fortiAPIPatch(o["spectrum-analysis"], "WirelessControllerWtpProfileRadio2-SpectrumAnalysis"); ok {
			if err = d.Set("spectrum_analysis", vv); err != nil {
				return fmt.Errorf("Error reading spectrum_analysis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spectrum_analysis: %v", err)
		}
	}

	if err = d.Set("transmit_optimize", flattenWirelessControllerWtpProfileRadio2TransmitOptimize2edl(o["transmit-optimize"], d, "transmit_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["transmit-optimize"], "WirelessControllerWtpProfileRadio2-TransmitOptimize"); ok {
			if err = d.Set("transmit_optimize", vv); err != nil {
				return fmt.Errorf("Error reading transmit_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transmit_optimize: %v", err)
		}
	}

	if err = d.Set("vap_all", flattenWirelessControllerWtpProfileRadio2VapAll2edl(o["vap-all"], d, "vap_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap-all"], "WirelessControllerWtpProfileRadio2-VapAll"); ok {
			if err = d.Set("vap_all", vv); err != nil {
				return fmt.Errorf("Error reading vap_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap_all: %v", err)
		}
	}

	if err = d.Set("vap1", flattenWirelessControllerWtpProfileRadio2Vap12edl(o["vap1"], d, "vap1")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap1"], "WirelessControllerWtpProfileRadio2-Vap1"); ok {
			if err = d.Set("vap1", vv); err != nil {
				return fmt.Errorf("Error reading vap1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap1: %v", err)
		}
	}

	if err = d.Set("vap2", flattenWirelessControllerWtpProfileRadio2Vap22edl(o["vap2"], d, "vap2")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap2"], "WirelessControllerWtpProfileRadio2-Vap2"); ok {
			if err = d.Set("vap2", vv); err != nil {
				return fmt.Errorf("Error reading vap2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap2: %v", err)
		}
	}

	if err = d.Set("vap3", flattenWirelessControllerWtpProfileRadio2Vap32edl(o["vap3"], d, "vap3")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap3"], "WirelessControllerWtpProfileRadio2-Vap3"); ok {
			if err = d.Set("vap3", vv); err != nil {
				return fmt.Errorf("Error reading vap3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap3: %v", err)
		}
	}

	if err = d.Set("vap4", flattenWirelessControllerWtpProfileRadio2Vap42edl(o["vap4"], d, "vap4")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap4"], "WirelessControllerWtpProfileRadio2-Vap4"); ok {
			if err = d.Set("vap4", vv); err != nil {
				return fmt.Errorf("Error reading vap4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap4: %v", err)
		}
	}

	if err = d.Set("vap5", flattenWirelessControllerWtpProfileRadio2Vap52edl(o["vap5"], d, "vap5")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap5"], "WirelessControllerWtpProfileRadio2-Vap5"); ok {
			if err = d.Set("vap5", vv); err != nil {
				return fmt.Errorf("Error reading vap5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap5: %v", err)
		}
	}

	if err = d.Set("vap6", flattenWirelessControllerWtpProfileRadio2Vap62edl(o["vap6"], d, "vap6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap6"], "WirelessControllerWtpProfileRadio2-Vap6"); ok {
			if err = d.Set("vap6", vv); err != nil {
				return fmt.Errorf("Error reading vap6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap6: %v", err)
		}
	}

	if err = d.Set("vap7", flattenWirelessControllerWtpProfileRadio2Vap72edl(o["vap7"], d, "vap7")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap7"], "WirelessControllerWtpProfileRadio2-Vap7"); ok {
			if err = d.Set("vap7", vv); err != nil {
				return fmt.Errorf("Error reading vap7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap7: %v", err)
		}
	}

	if err = d.Set("vap8", flattenWirelessControllerWtpProfileRadio2Vap82edl(o["vap8"], d, "vap8")); err != nil {
		if vv, ok := fortiAPIPatch(o["vap8"], "WirelessControllerWtpProfileRadio2-Vap8"); ok {
			if err = d.Set("vap8", vv); err != nil {
				return fmt.Errorf("Error reading vap8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vap8: %v", err)
		}
	}

	if err = d.Set("vaps", flattenWirelessControllerWtpProfileRadio2Vaps2edl(o["vaps"], d, "vaps")); err != nil {
		if vv, ok := fortiAPIPatch(o["vaps"], "WirelessControllerWtpProfileRadio2-Vaps"); ok {
			if err = d.Set("vaps", vv); err != nil {
				return fmt.Errorf("Error reading vaps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vaps: %v", err)
		}
	}

	if err = d.Set("wids_profile", flattenWirelessControllerWtpProfileRadio2WidsProfile2edl(o["wids-profile"], d, "wids_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["wids-profile"], "WirelessControllerWtpProfileRadio2-WidsProfile"); ok {
			if err = d.Set("wids_profile", vv); err != nil {
				return fmt.Errorf("Error reading wids_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wids_profile: %v", err)
		}
	}

	if err = d.Set("zero_wait_dfs", flattenWirelessControllerWtpProfileRadio2ZeroWaitDfs2edl(o["zero-wait-dfs"], d, "zero_wait_dfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["zero-wait-dfs"], "WirelessControllerWtpProfileRadio2-ZeroWaitDfs"); ok {
			if err = d.Set("zero_wait_dfs", vv); err != nil {
				return fmt.Errorf("Error reading zero_wait_dfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zero_wait_dfs: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileRadio2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfileRadio280211D2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio280211Mc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AirtimeFairness2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Amsdu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApHandoff2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferBufsize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferChan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferChanWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferCtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ArrpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerHigh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Band2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2Band5GType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthCapacity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BeaconInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BssColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BssColorMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallAdmissionControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallCapacity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Channel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2ChannelBonding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelBondingExt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelUtilization2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Coexistence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Darrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Drma2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2DrmaSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Dtim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FragThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FrequencyHandoff2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2IperfProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2IperfServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxDistance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MimoMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2OptionalAntenna2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2OptionalAntennaGain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowersaveOptimize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2ProtectionMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2RadioId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2RtsThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamBssid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCaCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamCaptivePortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamClientCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpFailureString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpMatchString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpSuccessString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpTestUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamEapMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamPrivateKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamPrivateKeyPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2SamReportIntv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamSecurityType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamSsid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamTest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ShortGuardInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SpectrumAnalysis2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2TransmitOptimize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2VapAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vap82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vaps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2WidsProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileRadio2ZeroWaitDfs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfileRadio2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n80211d"); ok || d.HasChange("n80211d") {
		t, err := expandWirelessControllerWtpProfileRadio280211D2edl(d, v, "n80211d")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211d"] = t
		}
	}

	if v, ok := d.GetOk("n80211mc"); ok || d.HasChange("n80211mc") {
		t, err := expandWirelessControllerWtpProfileRadio280211Mc2edl(d, v, "n80211mc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211mc"] = t
		}
	}

	if v, ok := d.GetOk("airtime_fairness"); ok || d.HasChange("airtime_fairness") {
		t, err := expandWirelessControllerWtpProfileRadio2AirtimeFairness2edl(d, v, "airtime_fairness")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["airtime-fairness"] = t
		}
	}

	if v, ok := d.GetOk("amsdu"); ok || d.HasChange("amsdu") {
		t, err := expandWirelessControllerWtpProfileRadio2Amsdu2edl(d, v, "amsdu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["amsdu"] = t
		}
	}

	if v, ok := d.GetOk("ap_handoff"); ok || d.HasChange("ap_handoff") {
		t, err := expandWirelessControllerWtpProfileRadio2ApHandoff2edl(d, v, "ap_handoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-handoff"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_addr"); ok || d.HasChange("ap_sniffer_addr") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferAddr2edl(d, v, "ap_sniffer_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-addr"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_bufsize"); ok || d.HasChange("ap_sniffer_bufsize") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferBufsize2edl(d, v, "ap_sniffer_bufsize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-bufsize"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_chan"); ok || d.HasChange("ap_sniffer_chan") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferChan2edl(d, v, "ap_sniffer_chan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-chan"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_chan_width"); ok || d.HasChange("ap_sniffer_chan_width") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferChanWidth2edl(d, v, "ap_sniffer_chan_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-chan-width"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_ctl"); ok || d.HasChange("ap_sniffer_ctl") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferCtl2edl(d, v, "ap_sniffer_ctl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-ctl"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_data"); ok || d.HasChange("ap_sniffer_data") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferData2edl(d, v, "ap_sniffer_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-data"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_mgmt_beacon"); ok || d.HasChange("ap_sniffer_mgmt_beacon") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon2edl(d, v, "ap_sniffer_mgmt_beacon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-mgmt-beacon"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_mgmt_other"); ok || d.HasChange("ap_sniffer_mgmt_other") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther2edl(d, v, "ap_sniffer_mgmt_other")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-mgmt-other"] = t
		}
	}

	if v, ok := d.GetOk("ap_sniffer_mgmt_probe"); ok || d.HasChange("ap_sniffer_mgmt_probe") {
		t, err := expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe2edl(d, v, "ap_sniffer_mgmt_probe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-sniffer-mgmt-probe"] = t
		}
	}

	if v, ok := d.GetOk("arrp_profile"); ok || d.HasChange("arrp_profile") {
		t, err := expandWirelessControllerWtpProfileRadio2ArrpProfile2edl(d, v, "arrp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arrp-profile"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_high"); ok || d.HasChange("auto_power_high") {
		t, err := expandWirelessControllerWtpProfileRadio2AutoPowerHigh2edl(d, v, "auto_power_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-high"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_level"); ok || d.HasChange("auto_power_level") {
		t, err := expandWirelessControllerWtpProfileRadio2AutoPowerLevel2edl(d, v, "auto_power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-level"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_low"); ok || d.HasChange("auto_power_low") {
		t, err := expandWirelessControllerWtpProfileRadio2AutoPowerLow2edl(d, v, "auto_power_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-low"] = t
		}
	}

	if v, ok := d.GetOk("auto_power_target"); ok || d.HasChange("auto_power_target") {
		t, err := expandWirelessControllerWtpProfileRadio2AutoPowerTarget2edl(d, v, "auto_power_target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-power-target"] = t
		}
	}

	if v, ok := d.GetOk("band"); ok || d.HasChange("band") {
		t, err := expandWirelessControllerWtpProfileRadio2Band2edl(d, v, "band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["band"] = t
		}
	}

	if v, ok := d.GetOk("band_5g_type"); ok || d.HasChange("band_5g_type") {
		t, err := expandWirelessControllerWtpProfileRadio2Band5GType2edl(d, v, "band_5g_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["band-5g-type"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_admission_control"); ok || d.HasChange("bandwidth_admission_control") {
		t, err := expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl2edl(d, v, "bandwidth_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_capacity"); ok || d.HasChange("bandwidth_capacity") {
		t, err := expandWirelessControllerWtpProfileRadio2BandwidthCapacity2edl(d, v, "bandwidth_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-capacity"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok || d.HasChange("beacon_interval") {
		t, err := expandWirelessControllerWtpProfileRadio2BeaconInterval2edl(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("bss_color"); ok || d.HasChange("bss_color") {
		t, err := expandWirelessControllerWtpProfileRadio2BssColor2edl(d, v, "bss_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_mode"); ok || d.HasChange("bss_color_mode") {
		t, err := expandWirelessControllerWtpProfileRadio2BssColorMode2edl(d, v, "bss_color_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-mode"] = t
		}
	}

	if v, ok := d.GetOk("call_admission_control"); ok || d.HasChange("call_admission_control") {
		t, err := expandWirelessControllerWtpProfileRadio2CallAdmissionControl2edl(d, v, "call_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("call_capacity"); ok || d.HasChange("call_capacity") {
		t, err := expandWirelessControllerWtpProfileRadio2CallCapacity2edl(d, v, "call_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-capacity"] = t
		}
	}

	if v, ok := d.GetOk("channel"); ok || d.HasChange("channel") {
		t, err := expandWirelessControllerWtpProfileRadio2Channel2edl(d, v, "channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel"] = t
		}
	}

	if v, ok := d.GetOk("channel_bonding"); ok || d.HasChange("channel_bonding") {
		t, err := expandWirelessControllerWtpProfileRadio2ChannelBonding2edl(d, v, "channel_bonding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel-bonding"] = t
		}
	}

	if v, ok := d.GetOk("channel_bonding_ext"); ok || d.HasChange("channel_bonding_ext") {
		t, err := expandWirelessControllerWtpProfileRadio2ChannelBondingExt2edl(d, v, "channel_bonding_ext")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel-bonding-ext"] = t
		}
	}

	if v, ok := d.GetOk("channel_utilization"); ok || d.HasChange("channel_utilization") {
		t, err := expandWirelessControllerWtpProfileRadio2ChannelUtilization2edl(d, v, "channel_utilization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel-utilization"] = t
		}
	}

	if v, ok := d.GetOk("coexistence"); ok || d.HasChange("coexistence") {
		t, err := expandWirelessControllerWtpProfileRadio2Coexistence2edl(d, v, "coexistence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coexistence"] = t
		}
	}

	if v, ok := d.GetOk("darrp"); ok || d.HasChange("darrp") {
		t, err := expandWirelessControllerWtpProfileRadio2Darrp2edl(d, v, "darrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp"] = t
		}
	}

	if v, ok := d.GetOk("drma"); ok || d.HasChange("drma") {
		t, err := expandWirelessControllerWtpProfileRadio2Drma2edl(d, v, "drma")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drma"] = t
		}
	}

	if v, ok := d.GetOk("drma_sensitivity"); ok || d.HasChange("drma_sensitivity") {
		t, err := expandWirelessControllerWtpProfileRadio2DrmaSensitivity2edl(d, v, "drma_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drma-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("dtim"); ok || d.HasChange("dtim") {
		t, err := expandWirelessControllerWtpProfileRadio2Dtim2edl(d, v, "dtim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtim"] = t
		}
	}

	if v, ok := d.GetOk("frag_threshold"); ok || d.HasChange("frag_threshold") {
		t, err := expandWirelessControllerWtpProfileRadio2FragThreshold2edl(d, v, "frag_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frag-threshold"] = t
		}
	}

	if v, ok := d.GetOk("frequency_handoff"); ok || d.HasChange("frequency_handoff") {
		t, err := expandWirelessControllerWtpProfileRadio2FrequencyHandoff2edl(d, v, "frequency_handoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frequency-handoff"] = t
		}
	}

	if v, ok := d.GetOk("iperf_protocol"); ok || d.HasChange("iperf_protocol") {
		t, err := expandWirelessControllerWtpProfileRadio2IperfProtocol2edl(d, v, "iperf_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iperf-protocol"] = t
		}
	}

	if v, ok := d.GetOk("iperf_server_port"); ok || d.HasChange("iperf_server_port") {
		t, err := expandWirelessControllerWtpProfileRadio2IperfServerPort2edl(d, v, "iperf_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iperf-server-port"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandWirelessControllerWtpProfileRadio2MaxClients2edl(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("max_distance"); ok || d.HasChange("max_distance") {
		t, err := expandWirelessControllerWtpProfileRadio2MaxDistance2edl(d, v, "max_distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-distance"] = t
		}
	}

	if v, ok := d.GetOk("mimo_mode"); ok || d.HasChange("mimo_mode") {
		t, err := expandWirelessControllerWtpProfileRadio2MimoMode2edl(d, v, "mimo_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mimo-mode"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandWirelessControllerWtpProfileRadio2Mode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("optional_antenna"); ok || d.HasChange("optional_antenna") {
		t, err := expandWirelessControllerWtpProfileRadio2OptionalAntenna2edl(d, v, "optional_antenna")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional-antenna"] = t
		}
	}

	if v, ok := d.GetOk("optional_antenna_gain"); ok || d.HasChange("optional_antenna_gain") {
		t, err := expandWirelessControllerWtpProfileRadio2OptionalAntennaGain2edl(d, v, "optional_antenna_gain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional-antenna-gain"] = t
		}
	}

	if v, ok := d.GetOk("power_level"); ok || d.HasChange("power_level") {
		t, err := expandWirelessControllerWtpProfileRadio2PowerLevel2edl(d, v, "power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-level"] = t
		}
	}

	if v, ok := d.GetOk("power_mode"); ok || d.HasChange("power_mode") {
		t, err := expandWirelessControllerWtpProfileRadio2PowerMode2edl(d, v, "power_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-mode"] = t
		}
	}

	if v, ok := d.GetOk("power_value"); ok || d.HasChange("power_value") {
		t, err := expandWirelessControllerWtpProfileRadio2PowerValue2edl(d, v, "power_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-value"] = t
		}
	}

	if v, ok := d.GetOk("powersave_optimize"); ok || d.HasChange("powersave_optimize") {
		t, err := expandWirelessControllerWtpProfileRadio2PowersaveOptimize2edl(d, v, "powersave_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["powersave-optimize"] = t
		}
	}

	if v, ok := d.GetOk("protection_mode"); ok || d.HasChange("protection_mode") {
		t, err := expandWirelessControllerWtpProfileRadio2ProtectionMode2edl(d, v, "protection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protection-mode"] = t
		}
	}

	if v, ok := d.GetOk("radio_id"); ok || d.HasChange("radio_id") {
		t, err := expandWirelessControllerWtpProfileRadio2RadioId2edl(d, v, "radio_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-id"] = t
		}
	}

	if v, ok := d.GetOk("rts_threshold"); ok || d.HasChange("rts_threshold") {
		t, err := expandWirelessControllerWtpProfileRadio2RtsThreshold2edl(d, v, "rts_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-threshold"] = t
		}
	}

	if v, ok := d.GetOk("sam_bssid"); ok || d.HasChange("sam_bssid") {
		t, err := expandWirelessControllerWtpProfileRadio2SamBssid2edl(d, v, "sam_bssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-bssid"] = t
		}
	}

	if v, ok := d.GetOk("sam_ca_certificate"); ok || d.HasChange("sam_ca_certificate") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCaCertificate2edl(d, v, "sam_ca_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-ca-certificate"] = t
		}
	}

	if v, ok := d.GetOk("sam_captive_portal"); ok || d.HasChange("sam_captive_portal") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCaptivePortal2edl(d, v, "sam_captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("sam_client_certificate"); ok || d.HasChange("sam_client_certificate") {
		t, err := expandWirelessControllerWtpProfileRadio2SamClientCertificate2edl(d, v, "sam_client_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-client-certificate"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_failure_string"); ok || d.HasChange("sam_cwp_failure_string") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpFailureString2edl(d, v, "sam_cwp_failure_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-failure-string"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_match_string"); ok || d.HasChange("sam_cwp_match_string") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpMatchString2edl(d, v, "sam_cwp_match_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-match-string"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_password"); ok || d.HasChange("sam_cwp_password") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpPassword2edl(d, v, "sam_cwp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-password"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_success_string"); ok || d.HasChange("sam_cwp_success_string") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpSuccessString2edl(d, v, "sam_cwp_success_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-success-string"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_test_url"); ok || d.HasChange("sam_cwp_test_url") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpTestUrl2edl(d, v, "sam_cwp_test_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-test-url"] = t
		}
	}

	if v, ok := d.GetOk("sam_cwp_username"); ok || d.HasChange("sam_cwp_username") {
		t, err := expandWirelessControllerWtpProfileRadio2SamCwpUsername2edl(d, v, "sam_cwp_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-cwp-username"] = t
		}
	}

	if v, ok := d.GetOk("sam_eap_method"); ok || d.HasChange("sam_eap_method") {
		t, err := expandWirelessControllerWtpProfileRadio2SamEapMethod2edl(d, v, "sam_eap_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-eap-method"] = t
		}
	}

	if v, ok := d.GetOk("sam_password"); ok || d.HasChange("sam_password") {
		t, err := expandWirelessControllerWtpProfileRadio2SamPassword2edl(d, v, "sam_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-password"] = t
		}
	}

	if v, ok := d.GetOk("sam_private_key"); ok || d.HasChange("sam_private_key") {
		t, err := expandWirelessControllerWtpProfileRadio2SamPrivateKey2edl(d, v, "sam_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-private-key"] = t
		}
	}

	if v, ok := d.GetOk("sam_private_key_password"); ok || d.HasChange("sam_private_key_password") {
		t, err := expandWirelessControllerWtpProfileRadio2SamPrivateKeyPassword2edl(d, v, "sam_private_key_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-private-key-password"] = t
		}
	}

	if v, ok := d.GetOk("sam_report_intv"); ok || d.HasChange("sam_report_intv") {
		t, err := expandWirelessControllerWtpProfileRadio2SamReportIntv2edl(d, v, "sam_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("sam_security_type"); ok || d.HasChange("sam_security_type") {
		t, err := expandWirelessControllerWtpProfileRadio2SamSecurityType2edl(d, v, "sam_security_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-security-type"] = t
		}
	}

	if v, ok := d.GetOk("sam_server"); ok || d.HasChange("sam_server") {
		t, err := expandWirelessControllerWtpProfileRadio2SamServer2edl(d, v, "sam_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-server"] = t
		}
	}

	if v, ok := d.GetOk("sam_server_fqdn"); ok || d.HasChange("sam_server_fqdn") {
		t, err := expandWirelessControllerWtpProfileRadio2SamServerFqdn2edl(d, v, "sam_server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("sam_server_ip"); ok || d.HasChange("sam_server_ip") {
		t, err := expandWirelessControllerWtpProfileRadio2SamServerIp2edl(d, v, "sam_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("sam_server_type"); ok || d.HasChange("sam_server_type") {
		t, err := expandWirelessControllerWtpProfileRadio2SamServerType2edl(d, v, "sam_server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-server-type"] = t
		}
	}

	if v, ok := d.GetOk("sam_ssid"); ok || d.HasChange("sam_ssid") {
		t, err := expandWirelessControllerWtpProfileRadio2SamSsid2edl(d, v, "sam_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-ssid"] = t
		}
	}

	if v, ok := d.GetOk("sam_test"); ok || d.HasChange("sam_test") {
		t, err := expandWirelessControllerWtpProfileRadio2SamTest2edl(d, v, "sam_test")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-test"] = t
		}
	}

	if v, ok := d.GetOk("sam_username"); ok || d.HasChange("sam_username") {
		t, err := expandWirelessControllerWtpProfileRadio2SamUsername2edl(d, v, "sam_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sam-username"] = t
		}
	}

	if v, ok := d.GetOk("short_guard_interval"); ok || d.HasChange("short_guard_interval") {
		t, err := expandWirelessControllerWtpProfileRadio2ShortGuardInterval2edl(d, v, "short_guard_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["short-guard-interval"] = t
		}
	}

	if v, ok := d.GetOk("spectrum_analysis"); ok || d.HasChange("spectrum_analysis") {
		t, err := expandWirelessControllerWtpProfileRadio2SpectrumAnalysis2edl(d, v, "spectrum_analysis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spectrum-analysis"] = t
		}
	}

	if v, ok := d.GetOk("transmit_optimize"); ok || d.HasChange("transmit_optimize") {
		t, err := expandWirelessControllerWtpProfileRadio2TransmitOptimize2edl(d, v, "transmit_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-optimize"] = t
		}
	}

	if v, ok := d.GetOk("vap_all"); ok || d.HasChange("vap_all") {
		t, err := expandWirelessControllerWtpProfileRadio2VapAll2edl(d, v, "vap_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap-all"] = t
		}
	}

	if v, ok := d.GetOk("vap1"); ok || d.HasChange("vap1") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap12edl(d, v, "vap1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap1"] = t
		}
	}

	if v, ok := d.GetOk("vap2"); ok || d.HasChange("vap2") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap22edl(d, v, "vap2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap2"] = t
		}
	}

	if v, ok := d.GetOk("vap3"); ok || d.HasChange("vap3") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap32edl(d, v, "vap3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap3"] = t
		}
	}

	if v, ok := d.GetOk("vap4"); ok || d.HasChange("vap4") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap42edl(d, v, "vap4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap4"] = t
		}
	}

	if v, ok := d.GetOk("vap5"); ok || d.HasChange("vap5") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap52edl(d, v, "vap5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap5"] = t
		}
	}

	if v, ok := d.GetOk("vap6"); ok || d.HasChange("vap6") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap62edl(d, v, "vap6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap6"] = t
		}
	}

	if v, ok := d.GetOk("vap7"); ok || d.HasChange("vap7") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap72edl(d, v, "vap7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap7"] = t
		}
	}

	if v, ok := d.GetOk("vap8"); ok || d.HasChange("vap8") {
		t, err := expandWirelessControllerWtpProfileRadio2Vap82edl(d, v, "vap8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap8"] = t
		}
	}

	if v, ok := d.GetOk("vaps"); ok || d.HasChange("vaps") {
		t, err := expandWirelessControllerWtpProfileRadio2Vaps2edl(d, v, "vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vaps"] = t
		}
	}

	if v, ok := d.GetOk("wids_profile"); ok || d.HasChange("wids_profile") {
		t, err := expandWirelessControllerWtpProfileRadio2WidsProfile2edl(d, v, "wids_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wids-profile"] = t
		}
	}

	if v, ok := d.GetOk("zero_wait_dfs"); ok || d.HasChange("zero_wait_dfs") {
		t, err := expandWirelessControllerWtpProfileRadio2ZeroWaitDfs2edl(d, v, "zero_wait_dfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zero-wait-dfs"] = t
		}
	}

	return &obj, nil
}
