// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless controller global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerGlobalUpdate,
		Read:   resourceWirelessControllerGlobalRead,
		Update: resourceWirelessControllerGlobalUpdate,
		Delete: resourceWirelessControllerGlobalDelete,

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
			"acd_process_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_log_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_log_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_log_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"control_message_offload": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"data_ethernet_ii": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfs_lab_test": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovery_mc_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fiapp_eth_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"image_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_base_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_radio_vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_ble_device": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_rogue_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rogue_ap_wtp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_rogue_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sta_cap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sta_cap_wtp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_wids_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mesh_eth_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nac_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rogue_scan_mac_adjacency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rolling_wtp_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rolling_wtp_upgrade_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wpad_process_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wtp_share": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerGlobal")

	return resourceWirelessControllerGlobalRead(d, m)
}

func resourceWirelessControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerGlobalAcdProcessCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalControlMessageOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerGlobalDataEthernetIi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalDfsLabTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalDiscoveryMcAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalFiappEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalImageDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalIpsecBaseIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalLinkAggregation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalLocalRadioVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerGlobalLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxBleDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRetransmit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRogueAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRogueApWtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRogueSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxStaCap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxStaCapWtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxWidsEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMeshEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalNacInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRogueScanMacAdjacency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRollingWtpUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRollingWtpUpgradeThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalWpadProcessCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerGlobalWtpShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("acd_process_count", flattenWirelessControllerGlobalAcdProcessCount(o["acd-process-count"], d, "acd_process_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["acd-process-count"], "WirelessControllerGlobal-AcdProcessCount"); ok {
			if err = d.Set("acd_process_count", vv); err != nil {
				return fmt.Errorf("Error reading acd_process_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acd_process_count: %v", err)
		}
	}

	if err = d.Set("ap_log_server", flattenWirelessControllerGlobalApLogServer(o["ap-log-server"], d, "ap_log_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-log-server"], "WirelessControllerGlobal-ApLogServer"); ok {
			if err = d.Set("ap_log_server", vv); err != nil {
				return fmt.Errorf("Error reading ap_log_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_log_server: %v", err)
		}
	}

	if err = d.Set("ap_log_server_ip", flattenWirelessControllerGlobalApLogServerIp(o["ap-log-server-ip"], d, "ap_log_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-log-server-ip"], "WirelessControllerGlobal-ApLogServerIp"); ok {
			if err = d.Set("ap_log_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ap_log_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_log_server_ip: %v", err)
		}
	}

	if err = d.Set("ap_log_server_port", flattenWirelessControllerGlobalApLogServerPort(o["ap-log-server-port"], d, "ap_log_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-log-server-port"], "WirelessControllerGlobal-ApLogServerPort"); ok {
			if err = d.Set("ap_log_server_port", vv); err != nil {
				return fmt.Errorf("Error reading ap_log_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_log_server_port: %v", err)
		}
	}

	if err = d.Set("control_message_offload", flattenWirelessControllerGlobalControlMessageOffload(o["control-message-offload"], d, "control_message_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["control-message-offload"], "WirelessControllerGlobal-ControlMessageOffload"); ok {
			if err = d.Set("control_message_offload", vv); err != nil {
				return fmt.Errorf("Error reading control_message_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading control_message_offload: %v", err)
		}
	}

	if err = d.Set("data_ethernet_ii", flattenWirelessControllerGlobalDataEthernetIi(o["data-ethernet-II"], d, "data_ethernet_ii")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-ethernet-II"], "WirelessControllerGlobal-DataEthernetIi"); ok {
			if err = d.Set("data_ethernet_ii", vv); err != nil {
				return fmt.Errorf("Error reading data_ethernet_ii: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_ethernet_ii: %v", err)
		}
	}

	if err = d.Set("dfs_lab_test", flattenWirelessControllerGlobalDfsLabTest(o["dfs-lab-test"], d, "dfs_lab_test")); err != nil {
		if vv, ok := fortiAPIPatch(o["dfs-lab-test"], "WirelessControllerGlobal-DfsLabTest"); ok {
			if err = d.Set("dfs_lab_test", vv); err != nil {
				return fmt.Errorf("Error reading dfs_lab_test: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dfs_lab_test: %v", err)
		}
	}

	if err = d.Set("discovery_mc_addr", flattenWirelessControllerGlobalDiscoveryMcAddr(o["discovery-mc-addr"], d, "discovery_mc_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["discovery-mc-addr"], "WirelessControllerGlobal-DiscoveryMcAddr"); ok {
			if err = d.Set("discovery_mc_addr", vv); err != nil {
				return fmt.Errorf("Error reading discovery_mc_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discovery_mc_addr: %v", err)
		}
	}

	if err = d.Set("fiapp_eth_type", flattenWirelessControllerGlobalFiappEthType(o["fiapp-eth-type"], d, "fiapp_eth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["fiapp-eth-type"], "WirelessControllerGlobal-FiappEthType"); ok {
			if err = d.Set("fiapp_eth_type", vv); err != nil {
				return fmt.Errorf("Error reading fiapp_eth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fiapp_eth_type: %v", err)
		}
	}

	if err = d.Set("image_download", flattenWirelessControllerGlobalImageDownload(o["image-download"], d, "image_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-download"], "WirelessControllerGlobal-ImageDownload"); ok {
			if err = d.Set("image_download", vv); err != nil {
				return fmt.Errorf("Error reading image_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_download: %v", err)
		}
	}

	if err = d.Set("ipsec_base_ip", flattenWirelessControllerGlobalIpsecBaseIp(o["ipsec-base-ip"], d, "ipsec_base_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-base-ip"], "WirelessControllerGlobal-IpsecBaseIp"); ok {
			if err = d.Set("ipsec_base_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_base_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_base_ip: %v", err)
		}
	}

	if err = d.Set("link_aggregation", flattenWirelessControllerGlobalLinkAggregation(o["link-aggregation"], d, "link_aggregation")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-aggregation"], "WirelessControllerGlobal-LinkAggregation"); ok {
			if err = d.Set("link_aggregation", vv); err != nil {
				return fmt.Errorf("Error reading link_aggregation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_aggregation: %v", err)
		}
	}

	if err = d.Set("local_radio_vdom", flattenWirelessControllerGlobalLocalRadioVdom(o["local-radio-vdom"], d, "local_radio_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-radio-vdom"], "WirelessControllerGlobal-LocalRadioVdom"); ok {
			if err = d.Set("local_radio_vdom", vv); err != nil {
				return fmt.Errorf("Error reading local_radio_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_radio_vdom: %v", err)
		}
	}

	if err = d.Set("location", flattenWirelessControllerGlobalLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "WirelessControllerGlobal-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("max_ble_device", flattenWirelessControllerGlobalMaxBleDevice(o["max-ble-device"], d, "max_ble_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ble-device"], "WirelessControllerGlobal-MaxBleDevice"); ok {
			if err = d.Set("max_ble_device", vv); err != nil {
				return fmt.Errorf("Error reading max_ble_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ble_device: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerGlobalMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "WirelessControllerGlobal-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_retransmit", flattenWirelessControllerGlobalMaxRetransmit(o["max-retransmit"], d, "max_retransmit")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-retransmit"], "WirelessControllerGlobal-MaxRetransmit"); ok {
			if err = d.Set("max_retransmit", vv); err != nil {
				return fmt.Errorf("Error reading max_retransmit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_retransmit: %v", err)
		}
	}

	if err = d.Set("max_rogue_ap", flattenWirelessControllerGlobalMaxRogueAp(o["max-rogue-ap"], d, "max_rogue_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rogue-ap"], "WirelessControllerGlobal-MaxRogueAp"); ok {
			if err = d.Set("max_rogue_ap", vv); err != nil {
				return fmt.Errorf("Error reading max_rogue_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rogue_ap: %v", err)
		}
	}

	if err = d.Set("max_rogue_ap_wtp", flattenWirelessControllerGlobalMaxRogueApWtp(o["max-rogue-ap-wtp"], d, "max_rogue_ap_wtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rogue-ap-wtp"], "WirelessControllerGlobal-MaxRogueApWtp"); ok {
			if err = d.Set("max_rogue_ap_wtp", vv); err != nil {
				return fmt.Errorf("Error reading max_rogue_ap_wtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rogue_ap_wtp: %v", err)
		}
	}

	if err = d.Set("max_rogue_sta", flattenWirelessControllerGlobalMaxRogueSta(o["max-rogue-sta"], d, "max_rogue_sta")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rogue-sta"], "WirelessControllerGlobal-MaxRogueSta"); ok {
			if err = d.Set("max_rogue_sta", vv); err != nil {
				return fmt.Errorf("Error reading max_rogue_sta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rogue_sta: %v", err)
		}
	}

	if err = d.Set("max_sta_cap", flattenWirelessControllerGlobalMaxStaCap(o["max-sta-cap"], d, "max_sta_cap")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-sta-cap"], "WirelessControllerGlobal-MaxStaCap"); ok {
			if err = d.Set("max_sta_cap", vv); err != nil {
				return fmt.Errorf("Error reading max_sta_cap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_sta_cap: %v", err)
		}
	}

	if err = d.Set("max_sta_cap_wtp", flattenWirelessControllerGlobalMaxStaCapWtp(o["max-sta-cap-wtp"], d, "max_sta_cap_wtp")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-sta-cap-wtp"], "WirelessControllerGlobal-MaxStaCapWtp"); ok {
			if err = d.Set("max_sta_cap_wtp", vv); err != nil {
				return fmt.Errorf("Error reading max_sta_cap_wtp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_sta_cap_wtp: %v", err)
		}
	}

	if err = d.Set("max_wids_entry", flattenWirelessControllerGlobalMaxWidsEntry(o["max-wids-entry"], d, "max_wids_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-wids-entry"], "WirelessControllerGlobal-MaxWidsEntry"); ok {
			if err = d.Set("max_wids_entry", vv); err != nil {
				return fmt.Errorf("Error reading max_wids_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_wids_entry: %v", err)
		}
	}

	if err = d.Set("mesh_eth_type", flattenWirelessControllerGlobalMeshEthType(o["mesh-eth-type"], d, "mesh_eth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-eth-type"], "WirelessControllerGlobal-MeshEthType"); ok {
			if err = d.Set("mesh_eth_type", vv); err != nil {
				return fmt.Errorf("Error reading mesh_eth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_eth_type: %v", err)
		}
	}

	if err = d.Set("nac_interval", flattenWirelessControllerGlobalNacInterval(o["nac-interval"], d, "nac_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-interval"], "WirelessControllerGlobal-NacInterval"); ok {
			if err = d.Set("nac_interval", vv); err != nil {
				return fmt.Errorf("Error reading nac_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerGlobalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerGlobal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rogue_scan_mac_adjacency", flattenWirelessControllerGlobalRogueScanMacAdjacency(o["rogue-scan-mac-adjacency"], d, "rogue_scan_mac_adjacency")); err != nil {
		if vv, ok := fortiAPIPatch(o["rogue-scan-mac-adjacency"], "WirelessControllerGlobal-RogueScanMacAdjacency"); ok {
			if err = d.Set("rogue_scan_mac_adjacency", vv); err != nil {
				return fmt.Errorf("Error reading rogue_scan_mac_adjacency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rogue_scan_mac_adjacency: %v", err)
		}
	}

	if err = d.Set("rolling_wtp_upgrade", flattenWirelessControllerGlobalRollingWtpUpgrade(o["rolling-wtp-upgrade"], d, "rolling_wtp_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["rolling-wtp-upgrade"], "WirelessControllerGlobal-RollingWtpUpgrade"); ok {
			if err = d.Set("rolling_wtp_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading rolling_wtp_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rolling_wtp_upgrade: %v", err)
		}
	}

	if err = d.Set("rolling_wtp_upgrade_threshold", flattenWirelessControllerGlobalRollingWtpUpgradeThreshold(o["rolling-wtp-upgrade-threshold"], d, "rolling_wtp_upgrade_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["rolling-wtp-upgrade-threshold"], "WirelessControllerGlobal-RollingWtpUpgradeThreshold"); ok {
			if err = d.Set("rolling_wtp_upgrade_threshold", vv); err != nil {
				return fmt.Errorf("Error reading rolling_wtp_upgrade_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rolling_wtp_upgrade_threshold: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenWirelessControllerGlobalTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mode"], "WirelessControllerGlobal-TunnelMode"); ok {
			if err = d.Set("tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("wpad_process_count", flattenWirelessControllerGlobalWpadProcessCount(o["wpad-process-count"], d, "wpad_process_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["wpad-process-count"], "WirelessControllerGlobal-WpadProcessCount"); ok {
			if err = d.Set("wpad_process_count", vv); err != nil {
				return fmt.Errorf("Error reading wpad_process_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wpad_process_count: %v", err)
		}
	}

	if err = d.Set("wtp_share", flattenWirelessControllerGlobalWtpShare(o["wtp-share"], d, "wtp_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtp-share"], "WirelessControllerGlobal-WtpShare"); ok {
			if err = d.Set("wtp_share", vv); err != nil {
				return fmt.Errorf("Error reading wtp_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtp_share: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerGlobalAcdProcessCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalControlMessageOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerGlobalDataEthernetIi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalDfsLabTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalDiscoveryMcAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalFiappEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalImageDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalIpsecBaseIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalLinkAggregation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalLocalRadioVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerGlobalLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxBleDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRetransmit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueApWtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaCap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaCapWtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxWidsEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMeshEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalNacInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRogueScanMacAdjacency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRollingWtpUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRollingWtpUpgradeThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalWpadProcessCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalWtpShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acd_process_count"); ok || d.HasChange("acd_process_count") {
		t, err := expandWirelessControllerGlobalAcdProcessCount(d, v, "acd_process_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acd-process-count"] = t
		}
	}

	if v, ok := d.GetOk("ap_log_server"); ok || d.HasChange("ap_log_server") {
		t, err := expandWirelessControllerGlobalApLogServer(d, v, "ap_log_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-log-server"] = t
		}
	}

	if v, ok := d.GetOk("ap_log_server_ip"); ok || d.HasChange("ap_log_server_ip") {
		t, err := expandWirelessControllerGlobalApLogServerIp(d, v, "ap_log_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-log-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ap_log_server_port"); ok || d.HasChange("ap_log_server_port") {
		t, err := expandWirelessControllerGlobalApLogServerPort(d, v, "ap_log_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-log-server-port"] = t
		}
	}

	if v, ok := d.GetOk("control_message_offload"); ok || d.HasChange("control_message_offload") {
		t, err := expandWirelessControllerGlobalControlMessageOffload(d, v, "control_message_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-message-offload"] = t
		}
	}

	if v, ok := d.GetOk("data_ethernet_ii"); ok || d.HasChange("data_ethernet_ii") {
		t, err := expandWirelessControllerGlobalDataEthernetIi(d, v, "data_ethernet_ii")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-ethernet-II"] = t
		}
	}

	if v, ok := d.GetOk("dfs_lab_test"); ok || d.HasChange("dfs_lab_test") {
		t, err := expandWirelessControllerGlobalDfsLabTest(d, v, "dfs_lab_test")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dfs-lab-test"] = t
		}
	}

	if v, ok := d.GetOk("discovery_mc_addr"); ok || d.HasChange("discovery_mc_addr") {
		t, err := expandWirelessControllerGlobalDiscoveryMcAddr(d, v, "discovery_mc_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-mc-addr"] = t
		}
	}

	if v, ok := d.GetOk("fiapp_eth_type"); ok || d.HasChange("fiapp_eth_type") {
		t, err := expandWirelessControllerGlobalFiappEthType(d, v, "fiapp_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fiapp-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("image_download"); ok || d.HasChange("image_download") {
		t, err := expandWirelessControllerGlobalImageDownload(d, v, "image_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-download"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_base_ip"); ok || d.HasChange("ipsec_base_ip") {
		t, err := expandWirelessControllerGlobalIpsecBaseIp(d, v, "ipsec_base_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-base-ip"] = t
		}
	}

	if v, ok := d.GetOk("link_aggregation"); ok || d.HasChange("link_aggregation") {
		t, err := expandWirelessControllerGlobalLinkAggregation(d, v, "link_aggregation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-aggregation"] = t
		}
	}

	if v, ok := d.GetOk("local_radio_vdom"); ok || d.HasChange("local_radio_vdom") {
		t, err := expandWirelessControllerGlobalLocalRadioVdom(d, v, "local_radio_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-radio-vdom"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandWirelessControllerGlobalLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("max_ble_device"); ok || d.HasChange("max_ble_device") {
		t, err := expandWirelessControllerGlobalMaxBleDevice(d, v, "max_ble_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ble-device"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandWirelessControllerGlobalMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("max_retransmit"); ok || d.HasChange("max_retransmit") {
		t, err := expandWirelessControllerGlobalMaxRetransmit(d, v, "max_retransmit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-retransmit"] = t
		}
	}

	if v, ok := d.GetOk("max_rogue_ap"); ok || d.HasChange("max_rogue_ap") {
		t, err := expandWirelessControllerGlobalMaxRogueAp(d, v, "max_rogue_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rogue-ap"] = t
		}
	}

	if v, ok := d.GetOk("max_rogue_ap_wtp"); ok || d.HasChange("max_rogue_ap_wtp") {
		t, err := expandWirelessControllerGlobalMaxRogueApWtp(d, v, "max_rogue_ap_wtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rogue-ap-wtp"] = t
		}
	}

	if v, ok := d.GetOk("max_rogue_sta"); ok || d.HasChange("max_rogue_sta") {
		t, err := expandWirelessControllerGlobalMaxRogueSta(d, v, "max_rogue_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rogue-sta"] = t
		}
	}

	if v, ok := d.GetOk("max_sta_cap"); ok || d.HasChange("max_sta_cap") {
		t, err := expandWirelessControllerGlobalMaxStaCap(d, v, "max_sta_cap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-sta-cap"] = t
		}
	}

	if v, ok := d.GetOk("max_sta_cap_wtp"); ok || d.HasChange("max_sta_cap_wtp") {
		t, err := expandWirelessControllerGlobalMaxStaCapWtp(d, v, "max_sta_cap_wtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-sta-cap-wtp"] = t
		}
	}

	if v, ok := d.GetOk("max_wids_entry"); ok || d.HasChange("max_wids_entry") {
		t, err := expandWirelessControllerGlobalMaxWidsEntry(d, v, "max_wids_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-wids-entry"] = t
		}
	}

	if v, ok := d.GetOk("mesh_eth_type"); ok || d.HasChange("mesh_eth_type") {
		t, err := expandWirelessControllerGlobalMeshEthType(d, v, "mesh_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("nac_interval"); ok || d.HasChange("nac_interval") {
		t, err := expandWirelessControllerGlobalNacInterval(d, v, "nac_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerGlobalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rogue_scan_mac_adjacency"); ok || d.HasChange("rogue_scan_mac_adjacency") {
		t, err := expandWirelessControllerGlobalRogueScanMacAdjacency(d, v, "rogue_scan_mac_adjacency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-scan-mac-adjacency"] = t
		}
	}

	if v, ok := d.GetOk("rolling_wtp_upgrade"); ok || d.HasChange("rolling_wtp_upgrade") {
		t, err := expandWirelessControllerGlobalRollingWtpUpgrade(d, v, "rolling_wtp_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-wtp-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("rolling_wtp_upgrade_threshold"); ok || d.HasChange("rolling_wtp_upgrade_threshold") {
		t, err := expandWirelessControllerGlobalRollingWtpUpgradeThreshold(d, v, "rolling_wtp_upgrade_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-wtp-upgrade-threshold"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok || d.HasChange("tunnel_mode") {
		t, err := expandWirelessControllerGlobalTunnelMode(d, v, "tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("wpad_process_count"); ok || d.HasChange("wpad_process_count") {
		t, err := expandWirelessControllerGlobalWpadProcessCount(d, v, "wpad_process_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wpad-process-count"] = t
		}
	}

	if v, ok := d.GetOk("wtp_share"); ok || d.HasChange("wtp_share") {
		t, err := expandWirelessControllerGlobalWtpShare(d, v, "wtp_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-share"] = t
		}
	}

	return &obj, nil
}
