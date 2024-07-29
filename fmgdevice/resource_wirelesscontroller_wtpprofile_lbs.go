// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Set various location based service (LBS) options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpProfileLbs() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfileLbsUpdate,
		Read:   resourceWirelessControllerWtpProfileLbsRead,
		Update: resourceWirelessControllerWtpProfileLbsUpdate,
		Delete: resourceWirelessControllerWtpProfileLbsDelete,

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
			"aeroscout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aeroscout_ap_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aeroscout_mmu_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aeroscout_mu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aeroscout_mu_factor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"aeroscout_mu_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"aeroscout_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aeroscout_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ekahau_blink_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ekahau_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"erc_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"erc_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortipresence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortipresence_ble": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortipresence_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortipresence_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortipresence_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_rogue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"fortipresence_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_server_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortipresence_unassoc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_accumulation_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_asset_addrgrp_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"polestar_asset_uuid_list1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_asset_uuid_list4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"polestar_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_reporting_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_server_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polestar_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"polestar_server_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"station_locate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileLbsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWtpProfileLbs(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileLbs resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpProfileLbs(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileLbs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerWtpProfileLbs")

	return resourceWirelessControllerWtpProfileLbsRead(d, m)
}

func resourceWirelessControllerWtpProfileLbsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerWtpProfileLbs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfileLbs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileLbsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWtpProfileLbs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileLbs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfileLbs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileLbs resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileLbsAeroscout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutApMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsEkahauTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceBle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresencePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceProject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceRogue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestar2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarServerPath2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarServerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsPolestarServerToken2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsStationLocate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfileLbs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aeroscout", flattenWirelessControllerWtpProfileLbsAeroscout2edl(o["aeroscout"], d, "aeroscout")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout"], "WirelessControllerWtpProfileLbs-Aeroscout"); ok {
			if err = d.Set("aeroscout", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout: %v", err)
		}
	}

	if err = d.Set("aeroscout_ap_mac", flattenWirelessControllerWtpProfileLbsAeroscoutApMac2edl(o["aeroscout-ap-mac"], d, "aeroscout_ap_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-ap-mac"], "WirelessControllerWtpProfileLbs-AeroscoutApMac"); ok {
			if err = d.Set("aeroscout_ap_mac", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_ap_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_ap_mac: %v", err)
		}
	}

	if err = d.Set("aeroscout_mmu_report", flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(o["aeroscout-mmu-report"], d, "aeroscout_mmu_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mmu-report"], "WirelessControllerWtpProfileLbs-AeroscoutMmuReport"); ok {
			if err = d.Set("aeroscout_mmu_report", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mmu_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mmu_report: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu", flattenWirelessControllerWtpProfileLbsAeroscoutMu2edl(o["aeroscout-mu"], d, "aeroscout_mu")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu"], "WirelessControllerWtpProfileLbs-AeroscoutMu"); ok {
			if err = d.Set("aeroscout_mu", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu_factor", flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(o["aeroscout-mu-factor"], d, "aeroscout_mu_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu-factor"], "WirelessControllerWtpProfileLbs-AeroscoutMuFactor"); ok {
			if err = d.Set("aeroscout_mu_factor", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu_factor: %v", err)
		}
	}

	if err = d.Set("aeroscout_mu_timeout", flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(o["aeroscout-mu-timeout"], d, "aeroscout_mu_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-mu-timeout"], "WirelessControllerWtpProfileLbs-AeroscoutMuTimeout"); ok {
			if err = d.Set("aeroscout_mu_timeout", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_mu_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_mu_timeout: %v", err)
		}
	}

	if err = d.Set("aeroscout_server_ip", flattenWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(o["aeroscout-server-ip"], d, "aeroscout_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-server-ip"], "WirelessControllerWtpProfileLbs-AeroscoutServerIp"); ok {
			if err = d.Set("aeroscout_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_server_ip: %v", err)
		}
	}

	if err = d.Set("aeroscout_server_port", flattenWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(o["aeroscout-server-port"], d, "aeroscout_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["aeroscout-server-port"], "WirelessControllerWtpProfileLbs-AeroscoutServerPort"); ok {
			if err = d.Set("aeroscout_server_port", vv); err != nil {
				return fmt.Errorf("Error reading aeroscout_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aeroscout_server_port: %v", err)
		}
	}

	if err = d.Set("ekahau_blink_mode", flattenWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(o["ekahau-blink-mode"], d, "ekahau_blink_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ekahau-blink-mode"], "WirelessControllerWtpProfileLbs-EkahauBlinkMode"); ok {
			if err = d.Set("ekahau_blink_mode", vv); err != nil {
				return fmt.Errorf("Error reading ekahau_blink_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ekahau_blink_mode: %v", err)
		}
	}

	if err = d.Set("ekahau_tag", flattenWirelessControllerWtpProfileLbsEkahauTag2edl(o["ekahau-tag"], d, "ekahau_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ekahau-tag"], "WirelessControllerWtpProfileLbs-EkahauTag"); ok {
			if err = d.Set("ekahau_tag", vv); err != nil {
				return fmt.Errorf("Error reading ekahau_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ekahau_tag: %v", err)
		}
	}

	if err = d.Set("erc_server_ip", flattenWirelessControllerWtpProfileLbsErcServerIp2edl(o["erc-server-ip"], d, "erc_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["erc-server-ip"], "WirelessControllerWtpProfileLbs-ErcServerIp"); ok {
			if err = d.Set("erc_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading erc_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading erc_server_ip: %v", err)
		}
	}

	if err = d.Set("erc_server_port", flattenWirelessControllerWtpProfileLbsErcServerPort2edl(o["erc-server-port"], d, "erc_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["erc-server-port"], "WirelessControllerWtpProfileLbs-ErcServerPort"); ok {
			if err = d.Set("erc_server_port", vv); err != nil {
				return fmt.Errorf("Error reading erc_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading erc_server_port: %v", err)
		}
	}

	if err = d.Set("fortipresence", flattenWirelessControllerWtpProfileLbsFortipresence2edl(o["fortipresence"], d, "fortipresence")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence"], "WirelessControllerWtpProfileLbs-Fortipresence"); ok {
			if err = d.Set("fortipresence", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence: %v", err)
		}
	}

	if err = d.Set("fortipresence_ble", flattenWirelessControllerWtpProfileLbsFortipresenceBle2edl(o["fortipresence-ble"], d, "fortipresence_ble")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-ble"], "WirelessControllerWtpProfileLbs-FortipresenceBle"); ok {
			if err = d.Set("fortipresence_ble", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_ble: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_ble: %v", err)
		}
	}

	if err = d.Set("fortipresence_frequency", flattenWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(o["fortipresence-frequency"], d, "fortipresence_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-frequency"], "WirelessControllerWtpProfileLbs-FortipresenceFrequency"); ok {
			if err = d.Set("fortipresence_frequency", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_frequency: %v", err)
		}
	}

	if err = d.Set("fortipresence_port", flattenWirelessControllerWtpProfileLbsFortipresencePort2edl(o["fortipresence-port"], d, "fortipresence_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-port"], "WirelessControllerWtpProfileLbs-FortipresencePort"); ok {
			if err = d.Set("fortipresence_port", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_port: %v", err)
		}
	}

	if err = d.Set("fortipresence_project", flattenWirelessControllerWtpProfileLbsFortipresenceProject2edl(o["fortipresence-project"], d, "fortipresence_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-project"], "WirelessControllerWtpProfileLbs-FortipresenceProject"); ok {
			if err = d.Set("fortipresence_project", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_project: %v", err)
		}
	}

	if err = d.Set("fortipresence_rogue", flattenWirelessControllerWtpProfileLbsFortipresenceRogue2edl(o["fortipresence-rogue"], d, "fortipresence_rogue")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-rogue"], "WirelessControllerWtpProfileLbs-FortipresenceRogue"); ok {
			if err = d.Set("fortipresence_rogue", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_rogue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_rogue: %v", err)
		}
	}

	if err = d.Set("fortipresence_server", flattenWirelessControllerWtpProfileLbsFortipresenceServer2edl(o["fortipresence-server"], d, "fortipresence_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server"], "WirelessControllerWtpProfileLbs-FortipresenceServer"); ok {
			if err = d.Set("fortipresence_server", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server: %v", err)
		}
	}

	if err = d.Set("fortipresence_server_addr_type", flattenWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(o["fortipresence-server-addr-type"], d, "fortipresence_server_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server-addr-type"], "WirelessControllerWtpProfileLbs-FortipresenceServerAddrType"); ok {
			if err = d.Set("fortipresence_server_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server_addr_type: %v", err)
		}
	}

	if err = d.Set("fortipresence_server_fqdn", flattenWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(o["fortipresence-server-fqdn"], d, "fortipresence_server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-server-fqdn"], "WirelessControllerWtpProfileLbs-FortipresenceServerFqdn"); ok {
			if err = d.Set("fortipresence_server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_server_fqdn: %v", err)
		}
	}

	if err = d.Set("fortipresence_unassoc", flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(o["fortipresence-unassoc"], d, "fortipresence_unassoc")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortipresence-unassoc"], "WirelessControllerWtpProfileLbs-FortipresenceUnassoc"); ok {
			if err = d.Set("fortipresence_unassoc", vv); err != nil {
				return fmt.Errorf("Error reading fortipresence_unassoc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortipresence_unassoc: %v", err)
		}
	}

	if err = d.Set("polestar", flattenWirelessControllerWtpProfileLbsPolestar2edl(o["polestar"], d, "polestar")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar"], "WirelessControllerWtpProfileLbs-Polestar"); ok {
			if err = d.Set("polestar", vv); err != nil {
				return fmt.Errorf("Error reading polestar: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar: %v", err)
		}
	}

	if err = d.Set("polestar_accumulation_interval", flattenWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(o["polestar-accumulation-interval"], d, "polestar_accumulation_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-accumulation-interval"], "WirelessControllerWtpProfileLbs-PolestarAccumulationInterval"); ok {
			if err = d.Set("polestar_accumulation_interval", vv); err != nil {
				return fmt.Errorf("Error reading polestar_accumulation_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_accumulation_interval: %v", err)
		}
	}

	if err = d.Set("polestar_asset_addrgrp_list", flattenWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(o["polestar-asset-addrgrp-list"], d, "polestar_asset_addrgrp_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-addrgrp-list"], "WirelessControllerWtpProfileLbs-PolestarAssetAddrgrpList"); ok {
			if err = d.Set("polestar_asset_addrgrp_list", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_addrgrp_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_addrgrp_list: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list1", flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(o["polestar-asset-uuid-list1"], d, "polestar_asset_uuid_list1")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list1"], "WirelessControllerWtpProfileLbs-PolestarAssetUuidList1"); ok {
			if err = d.Set("polestar_asset_uuid_list1", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list1: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list2", flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(o["polestar-asset-uuid-list2"], d, "polestar_asset_uuid_list2")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list2"], "WirelessControllerWtpProfileLbs-PolestarAssetUuidList2"); ok {
			if err = d.Set("polestar_asset_uuid_list2", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list2: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list3", flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(o["polestar-asset-uuid-list3"], d, "polestar_asset_uuid_list3")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list3"], "WirelessControllerWtpProfileLbs-PolestarAssetUuidList3"); ok {
			if err = d.Set("polestar_asset_uuid_list3", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list3: %v", err)
		}
	}

	if err = d.Set("polestar_asset_uuid_list4", flattenWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(o["polestar-asset-uuid-list4"], d, "polestar_asset_uuid_list4")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-asset-uuid-list4"], "WirelessControllerWtpProfileLbs-PolestarAssetUuidList4"); ok {
			if err = d.Set("polestar_asset_uuid_list4", vv); err != nil {
				return fmt.Errorf("Error reading polestar_asset_uuid_list4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_asset_uuid_list4: %v", err)
		}
	}

	if err = d.Set("polestar_protocol", flattenWirelessControllerWtpProfileLbsPolestarProtocol2edl(o["polestar-protocol"], d, "polestar_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-protocol"], "WirelessControllerWtpProfileLbs-PolestarProtocol"); ok {
			if err = d.Set("polestar_protocol", vv); err != nil {
				return fmt.Errorf("Error reading polestar_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_protocol: %v", err)
		}
	}

	if err = d.Set("polestar_reporting_interval", flattenWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(o["polestar-reporting-interval"], d, "polestar_reporting_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-reporting-interval"], "WirelessControllerWtpProfileLbs-PolestarReportingInterval"); ok {
			if err = d.Set("polestar_reporting_interval", vv); err != nil {
				return fmt.Errorf("Error reading polestar_reporting_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_reporting_interval: %v", err)
		}
	}

	if err = d.Set("polestar_server_fqdn", flattenWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(o["polestar-server-fqdn"], d, "polestar_server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-fqdn"], "WirelessControllerWtpProfileLbs-PolestarServerFqdn"); ok {
			if err = d.Set("polestar_server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_fqdn: %v", err)
		}
	}

	if err = d.Set("polestar_server_path", flattenWirelessControllerWtpProfileLbsPolestarServerPath2edl(o["polestar-server-path"], d, "polestar_server_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-path"], "WirelessControllerWtpProfileLbs-PolestarServerPath"); ok {
			if err = d.Set("polestar_server_path", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_path: %v", err)
		}
	}

	if err = d.Set("polestar_server_port", flattenWirelessControllerWtpProfileLbsPolestarServerPort2edl(o["polestar-server-port"], d, "polestar_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-port"], "WirelessControllerWtpProfileLbs-PolestarServerPort"); ok {
			if err = d.Set("polestar_server_port", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_port: %v", err)
		}
	}

	if err = d.Set("polestar_server_token", flattenWirelessControllerWtpProfileLbsPolestarServerToken2edl(o["polestar-server-token"], d, "polestar_server_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["polestar-server-token"], "WirelessControllerWtpProfileLbs-PolestarServerToken"); ok {
			if err = d.Set("polestar_server_token", vv); err != nil {
				return fmt.Errorf("Error reading polestar_server_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polestar_server_token: %v", err)
		}
	}

	if err = d.Set("station_locate", flattenWirelessControllerWtpProfileLbsStationLocate2edl(o["station-locate"], d, "station_locate")); err != nil {
		if vv, ok := fortiAPIPatch(o["station-locate"], "WirelessControllerWtpProfileLbs-StationLocate"); ok {
			if err = d.Set("station_locate", vv); err != nil {
				return fmt.Errorf("Error reading station_locate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading station_locate: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileLbsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfileLbsAeroscout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutApMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsEkahauTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceBle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresencePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceProject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceRogue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestar2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarServerPath2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarServerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsPolestarServerToken2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsStationLocate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfileLbs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aeroscout"); ok || d.HasChange("aeroscout") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscout2edl(d, v, "aeroscout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_ap_mac"); ok || d.HasChange("aeroscout_ap_mac") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutApMac2edl(d, v, "aeroscout_ap_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-ap-mac"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mmu_report"); ok || d.HasChange("aeroscout_mmu_report") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutMmuReport2edl(d, v, "aeroscout_mmu_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mmu-report"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu"); ok || d.HasChange("aeroscout_mu") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutMu2edl(d, v, "aeroscout_mu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu_factor"); ok || d.HasChange("aeroscout_mu_factor") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutMuFactor2edl(d, v, "aeroscout_mu_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu-factor"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_mu_timeout"); ok || d.HasChange("aeroscout_mu_timeout") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout2edl(d, v, "aeroscout_mu_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-mu-timeout"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_server_ip"); ok || d.HasChange("aeroscout_server_ip") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutServerIp2edl(d, v, "aeroscout_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("aeroscout_server_port"); ok || d.HasChange("aeroscout_server_port") {
		t, err := expandWirelessControllerWtpProfileLbsAeroscoutServerPort2edl(d, v, "aeroscout_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aeroscout-server-port"] = t
		}
	}

	if v, ok := d.GetOk("ekahau_blink_mode"); ok || d.HasChange("ekahau_blink_mode") {
		t, err := expandWirelessControllerWtpProfileLbsEkahauBlinkMode2edl(d, v, "ekahau_blink_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ekahau-blink-mode"] = t
		}
	}

	if v, ok := d.GetOk("ekahau_tag"); ok || d.HasChange("ekahau_tag") {
		t, err := expandWirelessControllerWtpProfileLbsEkahauTag2edl(d, v, "ekahau_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ekahau-tag"] = t
		}
	}

	if v, ok := d.GetOk("erc_server_ip"); ok || d.HasChange("erc_server_ip") {
		t, err := expandWirelessControllerWtpProfileLbsErcServerIp2edl(d, v, "erc_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erc-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("erc_server_port"); ok || d.HasChange("erc_server_port") {
		t, err := expandWirelessControllerWtpProfileLbsErcServerPort2edl(d, v, "erc_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erc-server-port"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence"); ok || d.HasChange("fortipresence") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresence2edl(d, v, "fortipresence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_ble"); ok || d.HasChange("fortipresence_ble") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceBle2edl(d, v, "fortipresence_ble")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-ble"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_frequency"); ok || d.HasChange("fortipresence_frequency") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceFrequency2edl(d, v, "fortipresence_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-frequency"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_port"); ok || d.HasChange("fortipresence_port") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresencePort2edl(d, v, "fortipresence_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-port"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_project"); ok || d.HasChange("fortipresence_project") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceProject2edl(d, v, "fortipresence_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-project"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_rogue"); ok || d.HasChange("fortipresence_rogue") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceRogue2edl(d, v, "fortipresence_rogue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-rogue"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_secret"); ok || d.HasChange("fortipresence_secret") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceSecret2edl(d, v, "fortipresence_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-secret"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server"); ok || d.HasChange("fortipresence_server") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceServer2edl(d, v, "fortipresence_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server_addr_type"); ok || d.HasChange("fortipresence_server_addr_type") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceServerAddrType2edl(d, v, "fortipresence_server_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_server_fqdn"); ok || d.HasChange("fortipresence_server_fqdn") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceServerFqdn2edl(d, v, "fortipresence_server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fortipresence_unassoc"); ok || d.HasChange("fortipresence_unassoc") {
		t, err := expandWirelessControllerWtpProfileLbsFortipresenceUnassoc2edl(d, v, "fortipresence_unassoc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortipresence-unassoc"] = t
		}
	}

	if v, ok := d.GetOk("polestar"); ok || d.HasChange("polestar") {
		t, err := expandWirelessControllerWtpProfileLbsPolestar2edl(d, v, "polestar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar"] = t
		}
	}

	if v, ok := d.GetOk("polestar_accumulation_interval"); ok || d.HasChange("polestar_accumulation_interval") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAccumulationInterval2edl(d, v, "polestar_accumulation_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-accumulation-interval"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_addrgrp_list"); ok || d.HasChange("polestar_asset_addrgrp_list") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAssetAddrgrpList2edl(d, v, "polestar_asset_addrgrp_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-addrgrp-list"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list1"); ok || d.HasChange("polestar_asset_uuid_list1") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAssetUuidList12edl(d, v, "polestar_asset_uuid_list1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list1"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list2"); ok || d.HasChange("polestar_asset_uuid_list2") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAssetUuidList22edl(d, v, "polestar_asset_uuid_list2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list2"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list3"); ok || d.HasChange("polestar_asset_uuid_list3") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAssetUuidList32edl(d, v, "polestar_asset_uuid_list3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list3"] = t
		}
	}

	if v, ok := d.GetOk("polestar_asset_uuid_list4"); ok || d.HasChange("polestar_asset_uuid_list4") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarAssetUuidList42edl(d, v, "polestar_asset_uuid_list4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-asset-uuid-list4"] = t
		}
	}

	if v, ok := d.GetOk("polestar_protocol"); ok || d.HasChange("polestar_protocol") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarProtocol2edl(d, v, "polestar_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-protocol"] = t
		}
	}

	if v, ok := d.GetOk("polestar_reporting_interval"); ok || d.HasChange("polestar_reporting_interval") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarReportingInterval2edl(d, v, "polestar_reporting_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-reporting-interval"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_fqdn"); ok || d.HasChange("polestar_server_fqdn") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarServerFqdn2edl(d, v, "polestar_server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_path"); ok || d.HasChange("polestar_server_path") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarServerPath2edl(d, v, "polestar_server_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-path"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_port"); ok || d.HasChange("polestar_server_port") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarServerPort2edl(d, v, "polestar_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-port"] = t
		}
	}

	if v, ok := d.GetOk("polestar_server_token"); ok || d.HasChange("polestar_server_token") {
		t, err := expandWirelessControllerWtpProfileLbsPolestarServerToken2edl(d, v, "polestar_server_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polestar-server-token"] = t
		}
	}

	if v, ok := d.GetOk("station_locate"); ok || d.HasChange("station_locate") {
		t, err := expandWirelessControllerWtpProfileLbsStationLocate2edl(d, v, "station_locate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["station-locate"] = t
		}
	}

	return &obj, nil
}
