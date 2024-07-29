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
			"asleap_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"assoc_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"assoc_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"null_ssid_probe_resp": &schema.Schema{
				Type:     schema.TypeString,
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
			"weak_wep_iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_bridge": &schema.Schema{
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

	obj, err := getObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWidsProfile(obj, paradict)

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

	obj, err := getObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWidsProfile(obj, mkey, paradict)
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

	err = c.DeleteWirelessControllerWidsProfile(mkey, paradict)
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

func flattenWirelessControllerWidsProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenWirelessControllerWidsProfileInvalidMacOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNullSsidProbeResp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileSensorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileSpoofedDeauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWeakWepIv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWirelessBridge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWidsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

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

	if err = d.Set("comment", flattenWirelessControllerWidsProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerWidsProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
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

	if err = d.Set("name", flattenWirelessControllerWidsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerWidsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
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

	if err = d.Set("weak_wep_iv", flattenWirelessControllerWidsProfileWeakWepIv(o["weak-wep-iv"], d, "weak_wep_iv")); err != nil {
		if vv, ok := fortiAPIPatch(o["weak-wep-iv"], "WirelessControllerWidsProfile-WeakWepIv"); ok {
			if err = d.Set("weak_wep_iv", vv); err != nil {
				return fmt.Errorf("Error reading weak_wep_iv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weak_wep_iv: %v", err)
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

	return nil
}

func flattenWirelessControllerWidsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
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

func expandWirelessControllerWidsProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthUnknownSrcThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

func expandWirelessControllerWidsProfileInvalidMacOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNullSsidProbeResp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSensorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSpoofedDeauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWeakWepIv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWirelessBridge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWidsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

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

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerWidsProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
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

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerWidsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
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

	if v, ok := d.GetOk("weak_wep_iv"); ok || d.HasChange("weak_wep_iv") {
		t, err := expandWirelessControllerWidsProfileWeakWepIv(d, v, "weak_wep_iv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weak-wep-iv"] = t
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

	return &obj, nil
}
