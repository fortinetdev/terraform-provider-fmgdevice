// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpCreate,
		Read:   resourceWirelessControllerWtpRead,
		Update: resourceWirelessControllerWtpUpdate,
		Delete: resourceWirelessControllerWtpDelete,

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
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"apcfg_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ble_major_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ble_minor_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bonjour_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"coordinate_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"coordinate_longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_provision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_fragment_preventing": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_passwd_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesh_bridge_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_ip_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_login_passwd_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_split_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_wan_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
			"radio_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
			"radio_3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
			"radio_4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
			"region": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"region_x": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region_y": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ip": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"split_tunneling_acl_local_ap_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tun_mtu_downlink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tun_mtu_uplink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wtp_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"wtp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerWtpCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerWtp(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtp resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWtp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtp resource: %v", err)
	}

	d.SetId(getStringKey(d, "wtp_id"))

	return resourceWirelessControllerWtpRead(d, m)
}

func resourceWirelessControllerWtpUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerWtp(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtp resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "wtp_id"))

	return resourceWirelessControllerWtpRead(d, m)
}

func resourceWirelessControllerWtpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerWtp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpApcfgProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpBleMajorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpBleMinorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpBonjourProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpFirmwareProvision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpImageDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := i["port-esl-mode"]; ok {
		result["port_esl_mode"] = flattenWirelessControllerWtpLanPortEslMode(i["port-esl-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := i["port-esl-ssid"]; ok {
		result["port_esl_ssid"] = flattenWirelessControllerWtpLanPortEslSsid(i["port-esl-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {
		result["port_mode"] = flattenWirelessControllerWtpLanPortMode(i["port-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {
		result["port_ssid"] = flattenWirelessControllerWtpLanPortSsid(i["port-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {
		result["port1_mode"] = flattenWirelessControllerWtpLanPort1Mode(i["port1-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {
		result["port1_ssid"] = flattenWirelessControllerWtpLanPort1Ssid(i["port1-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {
		result["port2_mode"] = flattenWirelessControllerWtpLanPort2Mode(i["port2-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {
		result["port2_ssid"] = flattenWirelessControllerWtpLanPort2Ssid(i["port2-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {
		result["port3_mode"] = flattenWirelessControllerWtpLanPort3Mode(i["port3-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {
		result["port3_ssid"] = flattenWirelessControllerWtpLanPort3Ssid(i["port3-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {
		result["port4_mode"] = flattenWirelessControllerWtpLanPort4Mode(i["port4-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {
		result["port4_ssid"] = flattenWirelessControllerWtpLanPort4Ssid(i["port4-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {
		result["port5_mode"] = flattenWirelessControllerWtpLanPort5Mode(i["port5-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {
		result["port5_ssid"] = flattenWirelessControllerWtpLanPort5Ssid(i["port5-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {
		result["port6_mode"] = flattenWirelessControllerWtpLanPort6Mode(i["port6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {
		result["port6_ssid"] = flattenWirelessControllerWtpLanPort6Ssid(i["port6-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {
		result["port7_mode"] = flattenWirelessControllerWtpLanPort7Mode(i["port7-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {
		result["port7_ssid"] = flattenWirelessControllerWtpLanPort7Ssid(i["port7-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {
		result["port8_mode"] = flattenWirelessControllerWtpLanPort8Mode(i["port8-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {
		result["port8_ssid"] = flattenWirelessControllerWtpLanPort8Ssid(i["port8-ssid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpLanPortEslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortEslSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort3Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort4Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort5Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort7Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLanPort8Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpMeshBridgeEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideIpFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideSplitTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpPurdueLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio1AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpRadio1AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio1Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio1Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {
		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio1DrmaManualMode(i["drma-manual-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio1OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio1OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio1OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio1OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio1OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio1PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpRadio1PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpRadio1PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio1RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio1VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenWirelessControllerWtpRadio1Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenWirelessControllerWtpRadio1Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenWirelessControllerWtpRadio1Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenWirelessControllerWtpRadio1Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenWirelessControllerWtpRadio1Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenWirelessControllerWtpRadio1Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenWirelessControllerWtpRadio1Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenWirelessControllerWtpRadio1Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio1Vaps(i["vaps"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpRadio1Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio1DrmaManualMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio2AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpRadio2AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio2Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio2Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {
		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio2DrmaManualMode(i["drma-manual-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio2OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio2OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio2OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio2OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio2OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio2PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpRadio2PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpRadio2PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio2RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio2VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenWirelessControllerWtpRadio2Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenWirelessControllerWtpRadio2Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenWirelessControllerWtpRadio2Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenWirelessControllerWtpRadio2Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenWirelessControllerWtpRadio2Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenWirelessControllerWtpRadio2Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenWirelessControllerWtpRadio2Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenWirelessControllerWtpRadio2Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio2Vaps(i["vaps"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpRadio2Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio2DrmaManualMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio3AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio3AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio3AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpRadio3AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio3Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio3Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {
		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio3DrmaManualMode(i["drma-manual-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio3OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio3OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio3OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio3OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio3OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio3PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpRadio3PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpRadio3PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio3RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio3SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio3VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenWirelessControllerWtpRadio3Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenWirelessControllerWtpRadio3Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenWirelessControllerWtpRadio3Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenWirelessControllerWtpRadio3Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenWirelessControllerWtpRadio3Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenWirelessControllerWtpRadio3Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenWirelessControllerWtpRadio3Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenWirelessControllerWtpRadio3Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio3Vaps(i["vaps"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio3AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpRadio3Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio3DrmaManualMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio4(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio4AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio4AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio4AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpRadio4AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio4Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio4Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {
		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio4DrmaManualMode(i["drma-manual-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio4OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio4OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio4OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio4OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio4OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio4PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpRadio4PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpRadio4PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio4RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio4SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio4VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenWirelessControllerWtpRadio4Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenWirelessControllerWtpRadio4Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenWirelessControllerWtpRadio4Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenWirelessControllerWtpRadio4Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenWirelessControllerWtpRadio4Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenWirelessControllerWtpRadio4Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenWirelessControllerWtpRadio4Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenWirelessControllerWtpRadio4Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio4Vaps(i["vaps"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio4AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenWirelessControllerWtpRadio4Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRadio4DrmaManualMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpRegionX(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegionY(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {
			v := flattenWirelessControllerWtpSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append)
			tmp["dest_ip"] = fortiAPISubPartPatch(v, "WirelessControllerWtp-SplitTunnelingAcl-DestIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWirelessControllerWtpSplitTunnelingAclId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WirelessControllerWtp-SplitTunnelingAcl-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWirelessControllerWtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("admin", flattenWirelessControllerWtpAdmin(o["admin"], d, "admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin"], "WirelessControllerWtp-Admin"); ok {
			if err = d.Set("admin", vv); err != nil {
				return fmt.Errorf("Error reading admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenWirelessControllerWtpAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "WirelessControllerWtp-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("apcfg_profile", flattenWirelessControllerWtpApcfgProfile(o["apcfg-profile"], d, "apcfg_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["apcfg-profile"], "WirelessControllerWtp-ApcfgProfile"); ok {
			if err = d.Set("apcfg_profile", vv); err != nil {
				return fmt.Errorf("Error reading apcfg_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apcfg_profile: %v", err)
		}
	}

	if err = d.Set("ble_major_id", flattenWirelessControllerWtpBleMajorId(o["ble-major-id"], d, "ble_major_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-major-id"], "WirelessControllerWtp-BleMajorId"); ok {
			if err = d.Set("ble_major_id", vv); err != nil {
				return fmt.Errorf("Error reading ble_major_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_major_id: %v", err)
		}
	}

	if err = d.Set("ble_minor_id", flattenWirelessControllerWtpBleMinorId(o["ble-minor-id"], d, "ble_minor_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-minor-id"], "WirelessControllerWtp-BleMinorId"); ok {
			if err = d.Set("ble_minor_id", vv); err != nil {
				return fmt.Errorf("Error reading ble_minor_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_minor_id: %v", err)
		}
	}

	if err = d.Set("bonjour_profile", flattenWirelessControllerWtpBonjourProfile(o["bonjour-profile"], d, "bonjour_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["bonjour-profile"], "WirelessControllerWtp-BonjourProfile"); ok {
			if err = d.Set("bonjour_profile", vv); err != nil {
				return fmt.Errorf("Error reading bonjour_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bonjour_profile: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWtpComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerWtp-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("coordinate_latitude", flattenWirelessControllerWtpCoordinateLatitude(o["coordinate-latitude"], d, "coordinate_latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["coordinate-latitude"], "WirelessControllerWtp-CoordinateLatitude"); ok {
			if err = d.Set("coordinate_latitude", vv); err != nil {
				return fmt.Errorf("Error reading coordinate_latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coordinate_latitude: %v", err)
		}
	}

	if err = d.Set("coordinate_longitude", flattenWirelessControllerWtpCoordinateLongitude(o["coordinate-longitude"], d, "coordinate_longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["coordinate-longitude"], "WirelessControllerWtp-CoordinateLongitude"); ok {
			if err = d.Set("coordinate_longitude", vv); err != nil {
				return fmt.Errorf("Error reading coordinate_longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading coordinate_longitude: %v", err)
		}
	}

	if err = d.Set("firmware_provision", flattenWirelessControllerWtpFirmwareProvision(o["firmware-provision"], d, "firmware_provision")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision"], "WirelessControllerWtp-FirmwareProvision"); ok {
			if err = d.Set("firmware_provision", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenWirelessControllerWtpFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-latest"], "WirelessControllerWtp-FirmwareProvisionLatest"); ok {
			if err = d.Set("firmware_provision_latest", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if err = d.Set("image_download", flattenWirelessControllerWtpImageDownload(o["image-download"], d, "image_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-download"], "WirelessControllerWtp-ImageDownload"); ok {
			if err = d.Set("image_download", vv); err != nil {
				return fmt.Errorf("Error reading image_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_download: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerWtpIndex(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerWtp-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenWirelessControllerWtpIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-fragment-preventing"], "WirelessControllerWtp-IpFragmentPreventing"); ok {
			if err = d.Set("ip_fragment_preventing", vv); err != nil {
				return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan"], "WirelessControllerWtp-Lan"); ok {
				if err = d.Set("lan", vv); err != nil {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan"], "WirelessControllerWtp-Lan"); ok {
					if err = d.Set("lan", vv); err != nil {
						return fmt.Errorf("Error reading lan: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if err = d.Set("led_state", flattenWirelessControllerWtpLedState(o["led-state"], d, "led_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["led-state"], "WirelessControllerWtp-LedState"); ok {
			if err = d.Set("led_state", vv); err != nil {
				return fmt.Errorf("Error reading led_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if err = d.Set("location", flattenWirelessControllerWtpLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "WirelessControllerWtp-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenWirelessControllerWtpLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-passwd-change"], "WirelessControllerWtp-LoginPasswdChange"); ok {
			if err = d.Set("login_passwd_change", vv); err != nil {
				return fmt.Errorf("Error reading login_passwd_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if err = d.Set("mesh_bridge_enable", flattenWirelessControllerWtpMeshBridgeEnable(o["mesh-bridge-enable"], d, "mesh_bridge_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-bridge-enable"], "WirelessControllerWtp-MeshBridgeEnable"); ok {
			if err = d.Set("mesh_bridge_enable", vv); err != nil {
				return fmt.Errorf("Error reading mesh_bridge_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_bridge_enable: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWtpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerWtp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_allowaccess", flattenWirelessControllerWtpOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-allowaccess"], "WirelessControllerWtp-OverrideAllowaccess"); ok {
			if err = d.Set("override_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading override_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("override_ip_fragment", flattenWirelessControllerWtpOverrideIpFragment(o["override-ip-fragment"], d, "override_ip_fragment")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-ip-fragment"], "WirelessControllerWtp-OverrideIpFragment"); ok {
			if err = d.Set("override_ip_fragment", vv); err != nil {
				return fmt.Errorf("Error reading override_ip_fragment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_ip_fragment: %v", err)
		}
	}

	if err = d.Set("override_lan", flattenWirelessControllerWtpOverrideLan(o["override-lan"], d, "override_lan")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-lan"], "WirelessControllerWtp-OverrideLan"); ok {
			if err = d.Set("override_lan", vv); err != nil {
				return fmt.Errorf("Error reading override_lan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_lan: %v", err)
		}
	}

	if err = d.Set("override_led_state", flattenWirelessControllerWtpOverrideLedState(o["override-led-state"], d, "override_led_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-led-state"], "WirelessControllerWtp-OverrideLedState"); ok {
			if err = d.Set("override_led_state", vv); err != nil {
				return fmt.Errorf("Error reading override_led_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_led_state: %v", err)
		}
	}

	if err = d.Set("override_login_passwd_change", flattenWirelessControllerWtpOverrideLoginPasswdChange(o["override-login-passwd-change"], d, "override_login_passwd_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-login-passwd-change"], "WirelessControllerWtp-OverrideLoginPasswdChange"); ok {
			if err = d.Set("override_login_passwd_change", vv); err != nil {
				return fmt.Errorf("Error reading override_login_passwd_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_login_passwd_change: %v", err)
		}
	}

	if err = d.Set("override_split_tunnel", flattenWirelessControllerWtpOverrideSplitTunnel(o["override-split-tunnel"], d, "override_split_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-split-tunnel"], "WirelessControllerWtp-OverrideSplitTunnel"); ok {
			if err = d.Set("override_split_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading override_split_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_split_tunnel: %v", err)
		}
	}

	if err = d.Set("override_wan_port_mode", flattenWirelessControllerWtpOverrideWanPortMode(o["override-wan-port-mode"], d, "override_wan_port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-wan-port-mode"], "WirelessControllerWtp-OverrideWanPortMode"); ok {
			if err = d.Set("override_wan_port_mode", vv); err != nil {
				return fmt.Errorf("Error reading override_wan_port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_wan_port_mode: %v", err)
		}
	}

	if err = d.Set("purdue_level", flattenWirelessControllerWtpPurdueLevel(o["purdue-level"], d, "purdue_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["purdue-level"], "WirelessControllerWtp-PurdueLevel"); ok {
			if err = d.Set("purdue_level", vv); err != nil {
				return fmt.Errorf("Error reading purdue_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading purdue_level: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-1"], "WirelessControllerWtp-Radio1"); ok {
				if err = d.Set("radio_1", vv); err != nil {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-1"], "WirelessControllerWtp-Radio1"); ok {
					if err = d.Set("radio_1", vv); err != nil {
						return fmt.Errorf("Error reading radio_1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-2"], "WirelessControllerWtp-Radio2"); ok {
				if err = d.Set("radio_2", vv); err != nil {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-2"], "WirelessControllerWtp-Radio2"); ok {
					if err = d.Set("radio_2", vv); err != nil {
						return fmt.Errorf("Error reading radio_2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_3", flattenWirelessControllerWtpRadio3(o["radio-3"], d, "radio_3")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-3"], "WirelessControllerWtp-Radio3"); ok {
				if err = d.Set("radio_3", vv); err != nil {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_3"); ok {
			if err = d.Set("radio_3", flattenWirelessControllerWtpRadio3(o["radio-3"], d, "radio_3")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-3"], "WirelessControllerWtp-Radio3"); ok {
					if err = d.Set("radio_3", vv); err != nil {
						return fmt.Errorf("Error reading radio_3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_4", flattenWirelessControllerWtpRadio4(o["radio-4"], d, "radio_4")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-4"], "WirelessControllerWtp-Radio4"); ok {
				if err = d.Set("radio_4", vv); err != nil {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_4"); ok {
			if err = d.Set("radio_4", flattenWirelessControllerWtpRadio4(o["radio-4"], d, "radio_4")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-4"], "WirelessControllerWtp-Radio4"); ok {
					if err = d.Set("radio_4", vv); err != nil {
						return fmt.Errorf("Error reading radio_4: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			}
		}
	}

	if err = d.Set("region", flattenWirelessControllerWtpRegion(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "WirelessControllerWtp-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("region_x", flattenWirelessControllerWtpRegionX(o["region-x"], d, "region_x")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-x"], "WirelessControllerWtp-RegionX"); ok {
			if err = d.Set("region_x", vv); err != nil {
				return fmt.Errorf("Error reading region_x: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_x: %v", err)
		}
	}

	if err = d.Set("region_y", flattenWirelessControllerWtpRegionY(o["region-y"], d, "region_y")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-y"], "WirelessControllerWtp-RegionY"); ok {
			if err = d.Set("region_y", vv); err != nil {
				return fmt.Errorf("Error reading region_y: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_y: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
			if vv, ok := fortiAPIPatch(o["split-tunneling-acl"], "WirelessControllerWtp-SplitTunnelingAcl"); ok {
				if err = d.Set("split_tunneling_acl", vv); err != nil {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
				if vv, ok := fortiAPIPatch(o["split-tunneling-acl"], "WirelessControllerWtp-SplitTunnelingAcl"); ok {
					if err = d.Set("split_tunneling_acl", vv); err != nil {
						return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"], "WirelessControllerWtp-SplitTunnelingAclLocalApSubnet"); ok {
			if err = d.Set("split_tunneling_acl_local_ap_subnet", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenWirelessControllerWtpSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-acl-path"], "WirelessControllerWtp-SplitTunnelingAclPath"); ok {
			if err = d.Set("split_tunneling_acl_path", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenWirelessControllerWtpTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["tun-mtu-downlink"], "WirelessControllerWtp-TunMtuDownlink"); ok {
			if err = d.Set("tun_mtu_downlink", vv); err != nil {
				return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenWirelessControllerWtpTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink")); err != nil {
		if vv, ok := fortiAPIPatch(o["tun-mtu-uplink"], "WirelessControllerWtp-TunMtuUplink"); ok {
			if err = d.Set("tun_mtu_uplink", vv); err != nil {
				return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("uuid", flattenWirelessControllerWtpUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "WirelessControllerWtp-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenWirelessControllerWtpWanPortMode(o["wan-port-mode"], d, "wan_port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-port-mode"], "WirelessControllerWtp-WanPortMode"); ok {
			if err = d.Set("wan_port_mode", vv); err != nil {
				return fmt.Errorf("Error reading wan_port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	if err = d.Set("wtp_id", flattenWirelessControllerWtpWtpId(o["wtp-id"], d, "wtp_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-id"], "WirelessControllerWtp-WtpId"); ok {
			if err = d.Set("wtp_id", vv); err != nil {
				return fmt.Errorf("Error reading wtp_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_id: %v", err)
		}
	}

	if err = d.Set("wtp_mode", flattenWirelessControllerWtpWtpMode(o["wtp-mode"], d, "wtp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-mode"], "WirelessControllerWtp-WtpMode"); ok {
			if err = d.Set("wtp_mode", vv); err != nil {
				return fmt.Errorf("Error reading wtp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_mode: %v", err)
		}
	}

	if err = d.Set("wtp_profile", flattenWirelessControllerWtpWtpProfile(o["wtp-profile"], d, "wtp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-profile"], "WirelessControllerWtp-WtpProfile"); ok {
			if err = d.Set("wtp_profile", vv); err != nil {
				return fmt.Errorf("Error reading wtp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_profile: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpApcfgProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpBleMajorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpBleMinorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpBonjourProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpFirmwareProvision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpImageDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port-esl-mode"], _ = expandWirelessControllerWtpLanPortEslMode(d, i["port_esl_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port-esl-ssid"], _ = expandWirelessControllerWtpLanPortEslSsid(d, i["port_esl_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port-mode"], _ = expandWirelessControllerWtpLanPortMode(d, i["port_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port-ssid"], _ = expandWirelessControllerWtpLanPortSsid(d, i["port_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port1-mode"], _ = expandWirelessControllerWtpLanPort1Mode(d, i["port1_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port1-ssid"], _ = expandWirelessControllerWtpLanPort1Ssid(d, i["port1_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2-mode"], _ = expandWirelessControllerWtpLanPort2Mode(d, i["port2_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2-ssid"], _ = expandWirelessControllerWtpLanPort2Ssid(d, i["port2_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3-mode"], _ = expandWirelessControllerWtpLanPort3Mode(d, i["port3_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3-ssid"], _ = expandWirelessControllerWtpLanPort3Ssid(d, i["port3_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port4-mode"], _ = expandWirelessControllerWtpLanPort4Mode(d, i["port4_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port4-ssid"], _ = expandWirelessControllerWtpLanPort4Ssid(d, i["port4_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port5-mode"], _ = expandWirelessControllerWtpLanPort5Mode(d, i["port5_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port5-ssid"], _ = expandWirelessControllerWtpLanPort5Ssid(d, i["port5_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port6-mode"], _ = expandWirelessControllerWtpLanPort6Mode(d, i["port6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port6-ssid"], _ = expandWirelessControllerWtpLanPort6Ssid(d, i["port6_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port7-mode"], _ = expandWirelessControllerWtpLanPort7Mode(d, i["port7_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port7-ssid"], _ = expandWirelessControllerWtpLanPort7Ssid(d, i["port7_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port8-mode"], _ = expandWirelessControllerWtpLanPort8Mode(d, i["port8_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port8-ssid"], _ = expandWirelessControllerWtpLanPort8Ssid(d, i["port8_ssid"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpLanPortEslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortEslSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort3Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort4Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort5Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort7Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLanPort8Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLoginPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpMeshBridgeEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideIpFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideSplitTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpPurdueLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio1AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-target"], _ = expandWirelessControllerWtpRadio1AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandWirelessControllerWtpRadio1Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandWirelessControllerWtpRadio1Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio1DrmaManualMode(d, i["drma_manual_mode"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio1OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-band"], _ = expandWirelessControllerWtpRadio1OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-channel"], _ = expandWirelessControllerWtpRadio1OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio1OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio1OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandWirelessControllerWtpRadio1PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-mode"], _ = expandWirelessControllerWtpRadio1PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-value"], _ = expandWirelessControllerWtpRadio1PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandWirelessControllerWtpRadio1RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap-all"], _ = expandWirelessControllerWtpRadio1VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap1"], _ = expandWirelessControllerWtpRadio1Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap2"], _ = expandWirelessControllerWtpRadio1Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap3"], _ = expandWirelessControllerWtpRadio1Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap4"], _ = expandWirelessControllerWtpRadio1Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap5"], _ = expandWirelessControllerWtpRadio1Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap6"], _ = expandWirelessControllerWtpRadio1Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap7"], _ = expandWirelessControllerWtpRadio1Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap8"], _ = expandWirelessControllerWtpRadio1Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vaps"], _ = expandWirelessControllerWtpRadio1Vaps(d, i["vaps"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio1Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio1DrmaManualMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio2AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-target"], _ = expandWirelessControllerWtpRadio2AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandWirelessControllerWtpRadio2Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandWirelessControllerWtpRadio2Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio2DrmaManualMode(d, i["drma_manual_mode"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio2OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-band"], _ = expandWirelessControllerWtpRadio2OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-channel"], _ = expandWirelessControllerWtpRadio2OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio2OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio2OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandWirelessControllerWtpRadio2PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-mode"], _ = expandWirelessControllerWtpRadio2PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-value"], _ = expandWirelessControllerWtpRadio2PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandWirelessControllerWtpRadio2RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap-all"], _ = expandWirelessControllerWtpRadio2VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap1"], _ = expandWirelessControllerWtpRadio2Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap2"], _ = expandWirelessControllerWtpRadio2Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap3"], _ = expandWirelessControllerWtpRadio2Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap4"], _ = expandWirelessControllerWtpRadio2Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap5"], _ = expandWirelessControllerWtpRadio2Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap6"], _ = expandWirelessControllerWtpRadio2Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap7"], _ = expandWirelessControllerWtpRadio2Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap8"], _ = expandWirelessControllerWtpRadio2Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vaps"], _ = expandWirelessControllerWtpRadio2Vaps(d, i["vaps"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio2Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio2DrmaManualMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio3AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio3AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio3AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-target"], _ = expandWirelessControllerWtpRadio3AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandWirelessControllerWtpRadio3Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandWirelessControllerWtpRadio3Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio3DrmaManualMode(d, i["drma_manual_mode"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio3OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-band"], _ = expandWirelessControllerWtpRadio3OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-channel"], _ = expandWirelessControllerWtpRadio3OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio3OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio3OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandWirelessControllerWtpRadio3PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-mode"], _ = expandWirelessControllerWtpRadio3PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-value"], _ = expandWirelessControllerWtpRadio3PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandWirelessControllerWtpRadio3RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio3SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap-all"], _ = expandWirelessControllerWtpRadio3VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap1"], _ = expandWirelessControllerWtpRadio3Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap2"], _ = expandWirelessControllerWtpRadio3Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap3"], _ = expandWirelessControllerWtpRadio3Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap4"], _ = expandWirelessControllerWtpRadio3Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap5"], _ = expandWirelessControllerWtpRadio3Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap6"], _ = expandWirelessControllerWtpRadio3Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap7"], _ = expandWirelessControllerWtpRadio3Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap8"], _ = expandWirelessControllerWtpRadio3Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vaps"], _ = expandWirelessControllerWtpRadio3Vaps(d, i["vaps"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio3AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio3Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio3DrmaManualMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio4AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio4AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio4AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["auto-power-target"], _ = expandWirelessControllerWtpRadio4AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandWirelessControllerWtpRadio4Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandWirelessControllerWtpRadio4Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio4DrmaManualMode(d, i["drma_manual_mode"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio4OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-band"], _ = expandWirelessControllerWtpRadio4OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-channel"], _ = expandWirelessControllerWtpRadio4OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio4OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio4OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandWirelessControllerWtpRadio4PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-mode"], _ = expandWirelessControllerWtpRadio4PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-value"], _ = expandWirelessControllerWtpRadio4PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandWirelessControllerWtpRadio4RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio4SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap-all"], _ = expandWirelessControllerWtpRadio4VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap1"], _ = expandWirelessControllerWtpRadio4Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap2"], _ = expandWirelessControllerWtpRadio4Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap3"], _ = expandWirelessControllerWtpRadio4Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap4"], _ = expandWirelessControllerWtpRadio4Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap5"], _ = expandWirelessControllerWtpRadio4Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap6"], _ = expandWirelessControllerWtpRadio4Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap7"], _ = expandWirelessControllerWtpRadio4Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vap8"], _ = expandWirelessControllerWtpRadio4Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vaps"], _ = expandWirelessControllerWtpRadio4Vaps(d, i["vaps"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio4AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio4Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRadio4DrmaManualMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerWtpRegionX(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegionY(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dest-ip"], _ = expandWirelessControllerWtpSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWirelessControllerWtpSplitTunnelingAclId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerWtpSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWirelessControllerWtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin"); ok || d.HasChange("admin") {
		t, err := expandWirelessControllerWtpAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandWirelessControllerWtpAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("apcfg_profile"); ok || d.HasChange("apcfg_profile") {
		t, err := expandWirelessControllerWtpApcfgProfile(d, v, "apcfg_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apcfg-profile"] = t
		}
	}

	if v, ok := d.GetOk("ble_major_id"); ok || d.HasChange("ble_major_id") {
		t, err := expandWirelessControllerWtpBleMajorId(d, v, "ble_major_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-major-id"] = t
		}
	}

	if v, ok := d.GetOk("ble_minor_id"); ok || d.HasChange("ble_minor_id") {
		t, err := expandWirelessControllerWtpBleMinorId(d, v, "ble_minor_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-minor-id"] = t
		}
	}

	if v, ok := d.GetOk("bonjour_profile"); ok || d.HasChange("bonjour_profile") {
		t, err := expandWirelessControllerWtpBonjourProfile(d, v, "bonjour_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bonjour-profile"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerWtpComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_latitude"); ok || d.HasChange("coordinate_latitude") {
		t, err := expandWirelessControllerWtpCoordinateLatitude(d, v, "coordinate_latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-latitude"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_longitude"); ok || d.HasChange("coordinate_longitude") {
		t, err := expandWirelessControllerWtpCoordinateLongitude(d, v, "coordinate_longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-longitude"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision"); ok || d.HasChange("firmware_provision") {
		t, err := expandWirelessControllerWtpFirmwareProvision(d, v, "firmware_provision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok || d.HasChange("firmware_provision_latest") {
		t, err := expandWirelessControllerWtpFirmwareProvisionLatest(d, v, "firmware_provision_latest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("image_download"); ok || d.HasChange("image_download") {
		t, err := expandWirelessControllerWtpImageDownload(d, v, "image_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-download"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerWtpIndex(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok || d.HasChange("ip_fragment_preventing") {
		t, err := expandWirelessControllerWtpIpFragmentPreventing(d, v, "ip_fragment_preventing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok || d.HasChange("lan") {
		t, err := expandWirelessControllerWtpLan(d, v, "lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok || d.HasChange("led_state") {
		t, err := expandWirelessControllerWtpLedState(d, v, "led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandWirelessControllerWtpLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok || d.HasChange("login_passwd") {
		t, err := expandWirelessControllerWtpLoginPasswd(d, v, "login_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok || d.HasChange("login_passwd_change") {
		t, err := expandWirelessControllerWtpLoginPasswdChange(d, v, "login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("mesh_bridge_enable"); ok || d.HasChange("mesh_bridge_enable") {
		t, err := expandWirelessControllerWtpMeshBridgeEnable(d, v, "mesh_bridge_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-bridge-enable"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerWtpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok || d.HasChange("override_allowaccess") {
		t, err := expandWirelessControllerWtpOverrideAllowaccess(d, v, "override_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_ip_fragment"); ok || d.HasChange("override_ip_fragment") {
		t, err := expandWirelessControllerWtpOverrideIpFragment(d, v, "override_ip_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-ip-fragment"] = t
		}
	}

	if v, ok := d.GetOk("override_lan"); ok || d.HasChange("override_lan") {
		t, err := expandWirelessControllerWtpOverrideLan(d, v, "override_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-lan"] = t
		}
	}

	if v, ok := d.GetOk("override_led_state"); ok || d.HasChange("override_led_state") {
		t, err := expandWirelessControllerWtpOverrideLedState(d, v, "override_led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-led-state"] = t
		}
	}

	if v, ok := d.GetOk("override_login_passwd_change"); ok || d.HasChange("override_login_passwd_change") {
		t, err := expandWirelessControllerWtpOverrideLoginPasswdChange(d, v, "override_login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("override_split_tunnel"); ok || d.HasChange("override_split_tunnel") {
		t, err := expandWirelessControllerWtpOverrideSplitTunnel(d, v, "override_split_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-split-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("override_wan_port_mode"); ok || d.HasChange("override_wan_port_mode") {
		t, err := expandWirelessControllerWtpOverrideWanPortMode(d, v, "override_wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("purdue_level"); ok || d.HasChange("purdue_level") {
		t, err := expandWirelessControllerWtpPurdueLevel(d, v, "purdue_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["purdue-level"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok || d.HasChange("radio_1") {
		t, err := expandWirelessControllerWtpRadio1(d, v, "radio_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok || d.HasChange("radio_2") {
		t, err := expandWirelessControllerWtpRadio2(d, v, "radio_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("radio_3"); ok || d.HasChange("radio_3") {
		t, err := expandWirelessControllerWtpRadio3(d, v, "radio_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-3"] = t
		}
	}

	if v, ok := d.GetOk("radio_4"); ok || d.HasChange("radio_4") {
		t, err := expandWirelessControllerWtpRadio4(d, v, "radio_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-4"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok || d.HasChange("region") {
		t, err := expandWirelessControllerWtpRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("region_x"); ok || d.HasChange("region_x") {
		t, err := expandWirelessControllerWtpRegionX(d, v, "region_x")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-x"] = t
		}
	}

	if v, ok := d.GetOk("region_y"); ok || d.HasChange("region_y") {
		t, err := expandWirelessControllerWtpRegionY(d, v, "region_y")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-y"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok || d.HasChange("split_tunneling_acl") {
		t, err := expandWirelessControllerWtpSplitTunnelingAcl(d, v, "split_tunneling_acl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok || d.HasChange("split_tunneling_acl_local_ap_subnet") {
		t, err := expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok || d.HasChange("split_tunneling_acl_path") {
		t, err := expandWirelessControllerWtpSplitTunnelingAclPath(d, v, "split_tunneling_acl_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok || d.HasChange("tun_mtu_downlink") {
		t, err := expandWirelessControllerWtpTunMtuDownlink(d, v, "tun_mtu_downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok || d.HasChange("tun_mtu_uplink") {
		t, err := expandWirelessControllerWtpTunMtuUplink(d, v, "tun_mtu_uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandWirelessControllerWtpUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok || d.HasChange("wan_port_mode") {
		t, err := expandWirelessControllerWtpWanPortMode(d, v, "wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("wtp_id"); ok || d.HasChange("wtp_id") {
		t, err := expandWirelessControllerWtpWtpId(d, v, "wtp_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-id"] = t
		}
	}

	if v, ok := d.GetOk("wtp_mode"); ok || d.HasChange("wtp_mode") {
		t, err := expandWirelessControllerWtpWtpMode(d, v, "wtp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-mode"] = t
		}
	}

	if v, ok := d.GetOk("wtp_profile"); ok || d.HasChange("wtp_profile") {
		t, err := expandWirelessControllerWtpWtpProfile(d, v, "wtp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-profile"] = t
		}
	}

	return &obj, nil
}
