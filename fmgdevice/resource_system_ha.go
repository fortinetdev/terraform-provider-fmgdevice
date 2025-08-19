// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure HA.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaUpdate,
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

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
			"arps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"arps_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_virtual_mac_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"backup_hbdev": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"board_failover_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"chassis_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"check_secondary_dev_health": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frup_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"active_switch_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"backup_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"evpn_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failover_hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ftp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gratuitous_arps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_mgmt_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ha_mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_port_dtag_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_port_outer_tpid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_uptime_diff_margin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_interval_in_milliseconds": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hbdev_second_vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hbdev_vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hc_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hello_holddown": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hw_session_hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hw_session_sync_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hw_session_sync_dev": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"imap_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_cluster_session_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_phase2_proposal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"l2ep_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_failed_signal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balance_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logical_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_based_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_compatible_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_failover_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_failover_monitor_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_failover_sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory_failover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minimum_worker_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"multicast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nntp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pingserver_failover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pingserver_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pingserver_monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pingserver_secondary_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingserver_slave_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pop3_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_switch_standby": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_wait_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pingserver_secondary_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vcluster_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"session_pickup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_connectionless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_expectation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"slave_switch_standby": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smtp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssd_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"standalone_config_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"standalone_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_packet_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb_peerip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_peers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"peer_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"unicast_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uninterruptible_primary_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uninterruptible_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcluster2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"upgrade_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_wait_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pingserver_flip_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pingserver_secondary_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vcluster_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"vcluster_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
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

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemHa")

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemHa(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHa(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaArpsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaAutoVirtualMacInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaBackupHbdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaBoardFailoverTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaChassisId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaCheckSecondaryDevHealth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFrup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFrupSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "active_interface"
	if _, ok := i["active-interface"]; ok {
		result["active_interface"] = flattenSystemHaFrupSettingsActiveInterface(i["active-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "active_switch_port"
	if _, ok := i["active-switch-port"]; ok {
		result["active_switch_port"] = flattenSystemHaFrupSettingsActiveSwitchPort(i["active-switch-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "backup_interface"
	if _, ok := i["backup-interface"]; ok {
		result["backup_interface"] = flattenSystemHaFrupSettingsBackupInterface(i["backup-interface"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemHaFrupSettingsActiveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaFrupSettingsActiveSwitchPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFrupSettingsBackupInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaEvpnTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFailoverHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGratuitousArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenSystemHaHaMgmtInterfacesDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "SystemHa-HaMgmtInterfaces-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenSystemHaHaMgmtInterfacesGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "SystemHa-HaMgmtInterfaces-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenSystemHaHaMgmtInterfacesGateway6(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "SystemHa-HaMgmtInterfaces-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemHaHaMgmtInterfacesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-HaMgmtInterfaces-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemHaHaMgmtInterfacesInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemHa-HaMgmtInterfaces-Interface")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaHaMgmtInterfacesDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaHaMgmtInterfacesGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaPortDtagMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaPortOuterTpid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaUptimeDiffMargin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbIntervalInMilliseconds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbdevSecondVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbdevVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHcEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHelloHolddown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHttpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHwSessionHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHwSessionSyncDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHwSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaImapProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaInterClusterSessionSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaIpsecPhase2Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaL2EpEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLinkFailedSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLoadBalanceAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLogicalSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryBasedFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryCompatibleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryFailoverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryFailoverMonitorPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryFailoverSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMinimumWorkerThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaNntpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPop3ProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondarySwitchStandby(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVcluster(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "monitor"
	if _, ok := i["monitor"]; ok {
		result["monitor"] = flattenSystemHaSecondaryVclusterMonitor(i["monitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "override"
	if _, ok := i["override"]; ok {
		result["override"] = flattenSystemHaSecondaryVclusterOverride(i["override"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := i["override-wait-time"]; ok {
		result["override_wait_time"] = flattenSystemHaSecondaryVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := i["pingserver-failover-threshold"]; ok {
		result["pingserver_failover_threshold"] = flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := i["pingserver-monitor-interface"]; ok {
		result["pingserver_monitor_interface"] = flattenSystemHaSecondaryVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_secondary_force_reset"
	if _, ok := i["pingserver-secondary-force-reset"]; ok {
		result["pingserver_secondary_force_reset"] = flattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(i["pingserver-secondary-force-reset"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := i["pingserver-slave-force-reset"]; ok {
		result["pingserver_slave_force_reset"] = flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemHaSecondaryVclusterPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := i["vcluster-id"]; ok {
		result["vcluster_id"] = flattenSystemHaSecondaryVclusterVclusterId(i["vcluster-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {
		result["vdom"] = flattenSystemHaSecondaryVclusterVdom(i["vdom"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemHaSecondaryVclusterMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaSecondaryVclusterOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaSessionPickup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupConnectionless(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupExpectation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaSlaveSwitchStandby(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSmtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSsdFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaStandaloneConfigSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaStandaloneMgmtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSyncConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSyncPacketBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHbNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHbPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastPeers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemHaUnicastPeersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemHa-UnicastPeers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := i["peer-ip"]; ok {
			v := flattenSystemHaUnicastPeersPeerIp(i["peer-ip"], d, pre_append)
			tmp["peer_ip"] = fortiAPISubPartPatch(v, "SystemHa-UnicastPeers-PeerIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaUnicastPeersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastPeersPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUninterruptiblePrimaryWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUninterruptibleUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVcluster2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaUpgradeMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVcluster(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			v := flattenSystemHaVclusterMonitor(i["monitor"], d, pre_append)
			tmp["monitor"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-Monitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override"
		if _, ok := i["override"]; ok {
			v := flattenSystemHaVclusterOverride(i["override"], d, pre_append)
			tmp["override"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-Override")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_wait_time"
		if _, ok := i["override-wait-time"]; ok {
			v := flattenSystemHaVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append)
			tmp["override_wait_time"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-OverrideWaitTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_failover_threshold"
		if _, ok := i["pingserver-failover-threshold"]; ok {
			v := flattenSystemHaVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append)
			tmp["pingserver_failover_threshold"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-PingserverFailoverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_flip_timeout"
		if _, ok := i["pingserver-flip-timeout"]; ok {
			v := flattenSystemHaVclusterPingserverFlipTimeout(i["pingserver-flip-timeout"], d, pre_append)
			tmp["pingserver_flip_timeout"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-PingserverFlipTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_monitor_interface"
		if _, ok := i["pingserver-monitor-interface"]; ok {
			v := flattenSystemHaVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append)
			tmp["pingserver_monitor_interface"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-PingserverMonitorInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_secondary_force_reset"
		if _, ok := i["pingserver-secondary-force-reset"]; ok {
			v := flattenSystemHaVclusterPingserverSecondaryForceReset(i["pingserver-secondary-force-reset"], d, pre_append)
			tmp["pingserver_secondary_force_reset"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-PingserverSecondaryForceReset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_slave_force_reset"
		if _, ok := i["pingserver-slave-force-reset"]; ok {
			v := flattenSystemHaVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append)
			tmp["pingserver_slave_force_reset"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-PingserverSlaveForceReset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSystemHaVclusterPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vcluster_id"
		if _, ok := i["vcluster-id"]; ok {
			v := flattenSystemHaVclusterVclusterId(i["vcluster-id"], d, pre_append)
			tmp["vcluster_id"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-VclusterId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSystemHaVclusterVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SystemHa-Vcluster-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemHaVclusterMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaVclusterOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaVclusterPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaVclusterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("arps", flattenSystemHaArps(o["arps"], d, "arps")); err != nil {
		if vv, ok := fortiAPIPatch(o["arps"], "SystemHa-Arps"); ok {
			if err = d.Set("arps", vv); err != nil {
				return fmt.Errorf("Error reading arps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arps: %v", err)
		}
	}

	if err = d.Set("arps_interval", flattenSystemHaArpsInterval(o["arps-interval"], d, "arps_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["arps-interval"], "SystemHa-ArpsInterval"); ok {
			if err = d.Set("arps_interval", vv); err != nil {
				return fmt.Errorf("Error reading arps_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arps_interval: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemHaAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystemHa-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("auto_virtual_mac_interface", flattenSystemHaAutoVirtualMacInterface(o["auto-virtual-mac-interface"], d, "auto_virtual_mac_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-virtual-mac-interface"], "SystemHa-AutoVirtualMacInterface"); ok {
			if err = d.Set("auto_virtual_mac_interface", vv); err != nil {
				return fmt.Errorf("Error reading auto_virtual_mac_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_virtual_mac_interface: %v", err)
		}
	}

	if err = d.Set("backup_hbdev", flattenSystemHaBackupHbdev(o["backup-hbdev"], d, "backup_hbdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["backup-hbdev"], "SystemHa-BackupHbdev"); ok {
			if err = d.Set("backup_hbdev", vv); err != nil {
				return fmt.Errorf("Error reading backup_hbdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backup_hbdev: %v", err)
		}
	}

	if err = d.Set("board_failover_tolerance", flattenSystemHaBoardFailoverTolerance(o["board-failover-tolerance"], d, "board_failover_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["board-failover-tolerance"], "SystemHa-BoardFailoverTolerance"); ok {
			if err = d.Set("board_failover_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading board_failover_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading board_failover_tolerance: %v", err)
		}
	}

	if err = d.Set("chassis_id", flattenSystemHaChassisId(o["chassis-id"], d, "chassis_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["chassis-id"], "SystemHa-ChassisId"); ok {
			if err = d.Set("chassis_id", vv); err != nil {
				return fmt.Errorf("Error reading chassis_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chassis_id: %v", err)
		}
	}

	if err = d.Set("check_secondary_dev_health", flattenSystemHaCheckSecondaryDevHealth(o["check-secondary-dev-health"], d, "check_secondary_dev_health")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-secondary-dev-health"], "SystemHa-CheckSecondaryDevHealth"); ok {
			if err = d.Set("check_secondary_dev_health", vv); err != nil {
				return fmt.Errorf("Error reading check_secondary_dev_health: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_secondary_dev_health: %v", err)
		}
	}

	if err = d.Set("cpu_threshold", flattenSystemHaCpuThreshold(o["cpu-threshold"], d, "cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu-threshold"], "SystemHa-CpuThreshold"); ok {
			if err = d.Set("cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu_threshold: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemHaEncryption(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "SystemHa-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("frup", flattenSystemHaFrup(o["frup"], d, "frup")); err != nil {
		if vv, ok := fortiAPIPatch(o["frup"], "SystemHa-Frup"); ok {
			if err = d.Set("frup", vv); err != nil {
				return fmt.Errorf("Error reading frup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frup: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("frup_settings", flattenSystemHaFrupSettings(o["frup-settings"], d, "frup_settings")); err != nil {
			if vv, ok := fortiAPIPatch(o["frup-settings"], "SystemHa-FrupSettings"); ok {
				if err = d.Set("frup_settings", vv); err != nil {
					return fmt.Errorf("Error reading frup_settings: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading frup_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("frup_settings"); ok {
			if err = d.Set("frup_settings", flattenSystemHaFrupSettings(o["frup-settings"], d, "frup_settings")); err != nil {
				if vv, ok := fortiAPIPatch(o["frup-settings"], "SystemHa-FrupSettings"); ok {
					if err = d.Set("frup_settings", vv); err != nil {
						return fmt.Errorf("Error reading frup_settings: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading frup_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("evpn_ttl", flattenSystemHaEvpnTtl(o["evpn-ttl"], d, "evpn_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["evpn-ttl"], "SystemHa-EvpnTtl"); ok {
			if err = d.Set("evpn_ttl", vv); err != nil {
				return fmt.Errorf("Error reading evpn_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading evpn_ttl: %v", err)
		}
	}

	if err = d.Set("failover_hold_time", flattenSystemHaFailoverHoldTime(o["failover-hold-time"], d, "failover_hold_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["failover-hold-time"], "SystemHa-FailoverHoldTime"); ok {
			if err = d.Set("failover_hold_time", vv); err != nil {
				return fmt.Errorf("Error reading failover_hold_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failover_hold_time: %v", err)
		}
	}

	if err = d.Set("ftp_proxy_threshold", flattenSystemHaFtpProxyThreshold(o["ftp-proxy-threshold"], d, "ftp_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-proxy-threshold"], "SystemHa-FtpProxyThreshold"); ok {
			if err = d.Set("ftp_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ftp_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("gratuitous_arps", flattenSystemHaGratuitousArps(o["gratuitous-arps"], d, "gratuitous_arps")); err != nil {
		if vv, ok := fortiAPIPatch(o["gratuitous-arps"], "SystemHa-GratuitousArps"); ok {
			if err = d.Set("gratuitous_arps", vv); err != nil {
				return fmt.Errorf("Error reading gratuitous_arps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gratuitous_arps: %v", err)
		}
	}

	if err = d.Set("group_id", flattenSystemHaGroupId(o["group-id"], d, "group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-id"], "SystemHa-GroupId"); ok {
			if err = d.Set("group_id", vv); err != nil {
				return fmt.Errorf("Error reading group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemHaGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "SystemHa-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("ha_direct", flattenSystemHaHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-direct"], "SystemHa-HaDirect"); ok {
			if err = d.Set("ha_direct", vv); err != nil {
				return fmt.Errorf("Error reading ha_direct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("ha_eth_type", flattenSystemHaHaEthType(o["ha-eth-type"], d, "ha_eth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-eth-type"], "SystemHa-HaEthType"); ok {
			if err = d.Set("ha_eth_type", vv); err != nil {
				return fmt.Errorf("Error reading ha_eth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_eth_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces")); err != nil {
			if vv, ok := fortiAPIPatch(o["ha-mgmt-interfaces"], "SystemHa-HaMgmtInterfaces"); ok {
				if err = d.Set("ha_mgmt_interfaces", vv); err != nil {
					return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ha_mgmt_interfaces"); ok {
			if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces")); err != nil {
				if vv, ok := fortiAPIPatch(o["ha-mgmt-interfaces"], "SystemHa-HaMgmtInterfaces"); ok {
					if err = d.Set("ha_mgmt_interfaces", vv); err != nil {
						return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_mgmt_status", flattenSystemHaHaMgmtStatus(o["ha-mgmt-status"], d, "ha_mgmt_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-mgmt-status"], "SystemHa-HaMgmtStatus"); ok {
			if err = d.Set("ha_mgmt_status", vv); err != nil {
				return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
		}
	}

	if err = d.Set("ha_port_dtag_mode", flattenSystemHaHaPortDtagMode(o["ha-port-dtag-mode"], d, "ha_port_dtag_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-port-dtag-mode"], "SystemHa-HaPortDtagMode"); ok {
			if err = d.Set("ha_port_dtag_mode", vv); err != nil {
				return fmt.Errorf("Error reading ha_port_dtag_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_port_dtag_mode: %v", err)
		}
	}

	if err = d.Set("ha_port_outer_tpid", flattenSystemHaHaPortOuterTpid(o["ha-port-outer-tpid"], d, "ha_port_outer_tpid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-port-outer-tpid"], "SystemHa-HaPortOuterTpid"); ok {
			if err = d.Set("ha_port_outer_tpid", vv); err != nil {
				return fmt.Errorf("Error reading ha_port_outer_tpid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_port_outer_tpid: %v", err)
		}
	}

	if err = d.Set("ha_uptime_diff_margin", flattenSystemHaHaUptimeDiffMargin(o["ha-uptime-diff-margin"], d, "ha_uptime_diff_margin")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-uptime-diff-margin"], "SystemHa-HaUptimeDiffMargin"); ok {
			if err = d.Set("ha_uptime_diff_margin", vv); err != nil {
				return fmt.Errorf("Error reading ha_uptime_diff_margin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_uptime_diff_margin: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemHa-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_interval_in_milliseconds", flattenSystemHaHbIntervalInMilliseconds(o["hb-interval-in-milliseconds"], d, "hb_interval_in_milliseconds")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval-in-milliseconds"], "SystemHa-HbIntervalInMilliseconds"); ok {
			if err = d.Set("hb_interval_in_milliseconds", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval_in_milliseconds: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval_in_milliseconds: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-lost-threshold"], "SystemHa-HbLostThreshold"); ok {
			if err = d.Set("hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("hbdev", flattenSystemHaHbdev(o["hbdev"], d, "hbdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["hbdev"], "SystemHa-Hbdev"); ok {
			if err = d.Set("hbdev", vv); err != nil {
				return fmt.Errorf("Error reading hbdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if err = d.Set("hbdev_second_vlan_id", flattenSystemHaHbdevSecondVlanId(o["hbdev-second-vlan-id"], d, "hbdev_second_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["hbdev-second-vlan-id"], "SystemHa-HbdevSecondVlanId"); ok {
			if err = d.Set("hbdev_second_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading hbdev_second_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hbdev_second_vlan_id: %v", err)
		}
	}

	if err = d.Set("hbdev_vlan_id", flattenSystemHaHbdevVlanId(o["hbdev-vlan-id"], d, "hbdev_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["hbdev-vlan-id"], "SystemHa-HbdevVlanId"); ok {
			if err = d.Set("hbdev_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading hbdev_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hbdev_vlan_id: %v", err)
		}
	}

	if err = d.Set("hc_eth_type", flattenSystemHaHcEthType(o["hc-eth-type"], d, "hc_eth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["hc-eth-type"], "SystemHa-HcEthType"); ok {
			if err = d.Set("hc_eth_type", vv); err != nil {
				return fmt.Errorf("Error reading hc_eth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hc_eth_type: %v", err)
		}
	}

	if err = d.Set("hello_holddown", flattenSystemHaHelloHolddown(o["hello-holddown"], d, "hello_holddown")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-holddown"], "SystemHa-HelloHolddown"); ok {
			if err = d.Set("hello_holddown", vv); err != nil {
				return fmt.Errorf("Error reading hello_holddown: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_holddown: %v", err)
		}
	}

	if err = d.Set("http_proxy_threshold", flattenSystemHaHttpProxyThreshold(o["http-proxy-threshold"], d, "http_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-proxy-threshold"], "SystemHa-HttpProxyThreshold"); ok {
			if err = d.Set("http_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading http_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("hw_session_hold_time", flattenSystemHaHwSessionHoldTime(o["hw-session-hold-time"], d, "hw_session_hold_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-session-hold-time"], "SystemHa-HwSessionHoldTime"); ok {
			if err = d.Set("hw_session_hold_time", vv); err != nil {
				return fmt.Errorf("Error reading hw_session_hold_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_session_hold_time: %v", err)
		}
	}

	if err = d.Set("hw_session_sync_delay", flattenSystemHaHwSessionSyncDelay(o["hw-session-sync-delay"], d, "hw_session_sync_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-session-sync-delay"], "SystemHa-HwSessionSyncDelay"); ok {
			if err = d.Set("hw_session_sync_delay", vv); err != nil {
				return fmt.Errorf("Error reading hw_session_sync_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_session_sync_delay: %v", err)
		}
	}

	if err = d.Set("hw_session_sync_dev", flattenSystemHaHwSessionSyncDev(o["hw-session-sync-dev"], d, "hw_session_sync_dev")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-session-sync-dev"], "SystemHa-HwSessionSyncDev"); ok {
			if err = d.Set("hw_session_sync_dev", vv); err != nil {
				return fmt.Errorf("Error reading hw_session_sync_dev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_session_sync_dev: %v", err)
		}
	}

	if err = d.Set("imap_proxy_threshold", flattenSystemHaImapProxyThreshold(o["imap-proxy-threshold"], d, "imap_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["imap-proxy-threshold"], "SystemHa-ImapProxyThreshold"); ok {
			if err = d.Set("imap_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading imap_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imap_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("inter_cluster_session_sync", flattenSystemHaInterClusterSessionSync(o["inter-cluster-session-sync"], d, "inter_cluster_session_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["inter-cluster-session-sync"], "SystemHa-InterClusterSessionSync"); ok {
			if err = d.Set("inter_cluster_session_sync", vv); err != nil {
				return fmt.Errorf("Error reading inter_cluster_session_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inter_cluster_session_sync: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_proposal", flattenSystemHaIpsecPhase2Proposal(o["ipsec-phase2-proposal"], d, "ipsec_phase2_proposal")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase2-proposal"], "SystemHa-IpsecPhase2Proposal"); ok {
			if err = d.Set("ipsec_phase2_proposal", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase2_proposal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase2_proposal: %v", err)
		}
	}

	if err = d.Set("l2ep_eth_type", flattenSystemHaL2EpEthType(o["l2ep-eth-type"], d, "l2ep_eth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2ep-eth-type"], "SystemHa-L2EpEthType"); ok {
			if err = d.Set("l2ep_eth_type", vv); err != nil {
				return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
		}
	}

	if err = d.Set("link_failed_signal", flattenSystemHaLinkFailedSignal(o["link-failed-signal"], d, "link_failed_signal")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-failed-signal"], "SystemHa-LinkFailedSignal"); ok {
			if err = d.Set("link_failed_signal", vv); err != nil {
				return fmt.Errorf("Error reading link_failed_signal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_failed_signal: %v", err)
		}
	}

	if err = d.Set("load_balance_all", flattenSystemHaLoadBalanceAll(o["load-balance-all"], d, "load_balance_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-all"], "SystemHa-LoadBalanceAll"); ok {
			if err = d.Set("load_balance_all", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_all: %v", err)
		}
	}

	if err = d.Set("logical_sn", flattenSystemHaLogicalSn(o["logical-sn"], d, "logical_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["logical-sn"], "SystemHa-LogicalSn"); ok {
			if err = d.Set("logical_sn", vv); err != nil {
				return fmt.Errorf("Error reading logical_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logical_sn: %v", err)
		}
	}

	if err = d.Set("memory_based_failover", flattenSystemHaMemoryBasedFailover(o["memory-based-failover"], d, "memory_based_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-based-failover"], "SystemHa-MemoryBasedFailover"); ok {
			if err = d.Set("memory_based_failover", vv); err != nil {
				return fmt.Errorf("Error reading memory_based_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_based_failover: %v", err)
		}
	}

	if err = d.Set("memory_compatible_mode", flattenSystemHaMemoryCompatibleMode(o["memory-compatible-mode"], d, "memory_compatible_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-compatible-mode"], "SystemHa-MemoryCompatibleMode"); ok {
			if err = d.Set("memory_compatible_mode", vv); err != nil {
				return fmt.Errorf("Error reading memory_compatible_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_compatible_mode: %v", err)
		}
	}

	if err = d.Set("memory_failover_flip_timeout", flattenSystemHaMemoryFailoverFlipTimeout(o["memory-failover-flip-timeout"], d, "memory_failover_flip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-failover-flip-timeout"], "SystemHa-MemoryFailoverFlipTimeout"); ok {
			if err = d.Set("memory_failover_flip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading memory_failover_flip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_failover_flip_timeout: %v", err)
		}
	}

	if err = d.Set("memory_failover_monitor_period", flattenSystemHaMemoryFailoverMonitorPeriod(o["memory-failover-monitor-period"], d, "memory_failover_monitor_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-failover-monitor-period"], "SystemHa-MemoryFailoverMonitorPeriod"); ok {
			if err = d.Set("memory_failover_monitor_period", vv); err != nil {
				return fmt.Errorf("Error reading memory_failover_monitor_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_failover_monitor_period: %v", err)
		}
	}

	if err = d.Set("memory_failover_sample_rate", flattenSystemHaMemoryFailoverSampleRate(o["memory-failover-sample-rate"], d, "memory_failover_sample_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-failover-sample-rate"], "SystemHa-MemoryFailoverSampleRate"); ok {
			if err = d.Set("memory_failover_sample_rate", vv); err != nil {
				return fmt.Errorf("Error reading memory_failover_sample_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_failover_sample_rate: %v", err)
		}
	}

	if err = d.Set("memory_failover_threshold", flattenSystemHaMemoryFailoverThreshold(o["memory-failover-threshold"], d, "memory_failover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-failover-threshold"], "SystemHa-MemoryFailoverThreshold"); ok {
			if err = d.Set("memory_failover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading memory_failover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_failover_threshold: %v", err)
		}
	}

	if err = d.Set("memory_threshold", flattenSystemHaMemoryThreshold(o["memory-threshold"], d, "memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["memory-threshold"], "SystemHa-MemoryThreshold"); ok {
			if err = d.Set("memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memory_threshold: %v", err)
		}
	}

	if err = d.Set("minimum_worker_threshold", flattenSystemHaMinimumWorkerThreshold(o["minimum-worker-threshold"], d, "minimum_worker_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-worker-threshold"], "SystemHa-MinimumWorkerThreshold"); ok {
			if err = d.Set("minimum_worker_threshold", vv); err != nil {
				return fmt.Errorf("Error reading minimum_worker_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_worker_threshold: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemHa-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("monitor", flattenSystemHaMonitor(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "SystemHa-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", flattenSystemHaMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-ttl"], "SystemHa-MulticastTtl"); ok {
			if err = d.Set("multicast_ttl", vv); err != nil {
				return fmt.Errorf("Error reading multicast_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("nntp_proxy_threshold", flattenSystemHaNntpProxyThreshold(o["nntp-proxy-threshold"], d, "nntp_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["nntp-proxy-threshold"], "SystemHa-NntpProxyThreshold"); ok {
			if err = d.Set("nntp_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading nntp_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nntp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemHaOverride(o["override"], d, "override")); err != nil {
		if vv, ok := fortiAPIPatch(o["override"], "SystemHa-Override"); ok {
			if err = d.Set("override", vv); err != nil {
				return fmt.Errorf("Error reading override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("override_wait_time", flattenSystemHaOverrideWaitTime(o["override-wait-time"], d, "override_wait_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-wait-time"], "SystemHa-OverrideWaitTime"); ok {
			if err = d.Set("override_wait_time", vv); err != nil {
				return fmt.Errorf("Error reading override_wait_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_wait_time: %v", err)
		}
	}

	if err = d.Set("pingserver_failover_threshold", flattenSystemHaPingserverFailoverThreshold(o["pingserver-failover-threshold"], d, "pingserver_failover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-failover-threshold"], "SystemHa-PingserverFailoverThreshold"); ok {
			if err = d.Set("pingserver_failover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
		}
	}

	if err = d.Set("pingserver_flip_timeout", flattenSystemHaPingserverFlipTimeout(o["pingserver-flip-timeout"], d, "pingserver_flip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-flip-timeout"], "SystemHa-PingserverFlipTimeout"); ok {
			if err = d.Set("pingserver_flip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
		}
	}

	if err = d.Set("pingserver_monitor_interface", flattenSystemHaPingserverMonitorInterface(o["pingserver-monitor-interface"], d, "pingserver_monitor_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-monitor-interface"], "SystemHa-PingserverMonitorInterface"); ok {
			if err = d.Set("pingserver_monitor_interface", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
		}
	}

	if err = d.Set("pingserver_secondary_force_reset", flattenSystemHaPingserverSecondaryForceReset(o["pingserver-secondary-force-reset"], d, "pingserver_secondary_force_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-secondary-force-reset"], "SystemHa-PingserverSecondaryForceReset"); ok {
			if err = d.Set("pingserver_secondary_force_reset", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_slave_force_reset", flattenSystemHaPingserverSlaveForceReset(o["pingserver-slave-force-reset"], d, "pingserver_slave_force_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-slave-force-reset"], "SystemHa-PingserverSlaveForceReset"); ok {
			if err = d.Set("pingserver_slave_force_reset", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if err = d.Set("pop3_proxy_threshold", flattenSystemHaPop3ProxyThreshold(o["pop3-proxy-threshold"], d, "pop3_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["pop3-proxy-threshold"], "SystemHa-Pop3ProxyThreshold"); ok {
			if err = d.Set("pop3_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading pop3_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pop3_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemHa-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("route_hold", flattenSystemHaRouteHold(o["route-hold"], d, "route_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-hold"], "SystemHa-RouteHold"); ok {
			if err = d.Set("route_hold", vv); err != nil {
				return fmt.Errorf("Error reading route_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_hold: %v", err)
		}
	}

	if err = d.Set("route_ttl", flattenSystemHaRouteTtl(o["route-ttl"], d, "route_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-ttl"], "SystemHa-RouteTtl"); ok {
			if err = d.Set("route_ttl", vv); err != nil {
				return fmt.Errorf("Error reading route_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_ttl: %v", err)
		}
	}

	if err = d.Set("route_wait", flattenSystemHaRouteWait(o["route-wait"], d, "route_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-wait"], "SystemHa-RouteWait"); ok {
			if err = d.Set("route_wait", vv); err != nil {
				return fmt.Errorf("Error reading route_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_wait: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemHaSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "SystemHa-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("secondary_switch_standby", flattenSystemHaSecondarySwitchStandby(o["secondary-switch-standby"], d, "secondary_switch_standby")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-switch-standby"], "SystemHa-SecondarySwitchStandby"); ok {
			if err = d.Set("secondary_switch_standby", vv); err != nil {
				return fmt.Errorf("Error reading secondary_switch_standby: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_switch_standby: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster")); err != nil {
			if vv, ok := fortiAPIPatch(o["secondary-vcluster"], "SystemHa-SecondaryVcluster"); ok {
				if err = d.Set("secondary_vcluster", vv); err != nil {
					return fmt.Errorf("Error reading secondary_vcluster: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading secondary_vcluster: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondary_vcluster"); ok {
			if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster")); err != nil {
				if vv, ok := fortiAPIPatch(o["secondary-vcluster"], "SystemHa-SecondaryVcluster"); ok {
					if err = d.Set("secondary_vcluster", vv); err != nil {
						return fmt.Errorf("Error reading secondary_vcluster: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading secondary_vcluster: %v", err)
				}
			}
		}
	}

	if err = d.Set("session_pickup", flattenSystemHaSessionPickup(o["session-pickup"], d, "session_pickup")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-pickup"], "SystemHa-SessionPickup"); ok {
			if err = d.Set("session_pickup", vv); err != nil {
				return fmt.Errorf("Error reading session_pickup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_pickup: %v", err)
		}
	}

	if err = d.Set("session_pickup_connectionless", flattenSystemHaSessionPickupConnectionless(o["session-pickup-connectionless"], d, "session_pickup_connectionless")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-pickup-connectionless"], "SystemHa-SessionPickupConnectionless"); ok {
			if err = d.Set("session_pickup_connectionless", vv); err != nil {
				return fmt.Errorf("Error reading session_pickup_connectionless: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_pickup_connectionless: %v", err)
		}
	}

	if err = d.Set("session_pickup_delay", flattenSystemHaSessionPickupDelay(o["session-pickup-delay"], d, "session_pickup_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-pickup-delay"], "SystemHa-SessionPickupDelay"); ok {
			if err = d.Set("session_pickup_delay", vv); err != nil {
				return fmt.Errorf("Error reading session_pickup_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_pickup_delay: %v", err)
		}
	}

	if err = d.Set("session_pickup_expectation", flattenSystemHaSessionPickupExpectation(o["session-pickup-expectation"], d, "session_pickup_expectation")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-pickup-expectation"], "SystemHa-SessionPickupExpectation"); ok {
			if err = d.Set("session_pickup_expectation", vv); err != nil {
				return fmt.Errorf("Error reading session_pickup_expectation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_pickup_expectation: %v", err)
		}
	}

	if err = d.Set("session_pickup_nat", flattenSystemHaSessionPickupNat(o["session-pickup-nat"], d, "session_pickup_nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-pickup-nat"], "SystemHa-SessionPickupNat"); ok {
			if err = d.Set("session_pickup_nat", vv); err != nil {
				return fmt.Errorf("Error reading session_pickup_nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_pickup_nat: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemHaSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-sync-dev"], "SystemHa-SessionSyncDev"); ok {
			if err = d.Set("session_sync_dev", vv); err != nil {
				return fmt.Errorf("Error reading session_sync_dev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("slave_switch_standby", flattenSystemHaSlaveSwitchStandby(o["slave-switch-standby"], d, "slave_switch_standby")); err != nil {
		if vv, ok := fortiAPIPatch(o["slave-switch-standby"], "SystemHa-SlaveSwitchStandby"); ok {
			if err = d.Set("slave_switch_standby", vv); err != nil {
				return fmt.Errorf("Error reading slave_switch_standby: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slave_switch_standby: %v", err)
		}
	}

	if err = d.Set("smtp_proxy_threshold", flattenSystemHaSmtpProxyThreshold(o["smtp-proxy-threshold"], d, "smtp_proxy_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["smtp-proxy-threshold"], "SystemHa-SmtpProxyThreshold"); ok {
			if err = d.Set("smtp_proxy_threshold", vv); err != nil {
				return fmt.Errorf("Error reading smtp_proxy_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smtp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("ssd_failover", flattenSystemHaSsdFailover(o["ssd-failover"], d, "ssd_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssd-failover"], "SystemHa-SsdFailover"); ok {
			if err = d.Set("ssd_failover", vv); err != nil {
				return fmt.Errorf("Error reading ssd_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssd_failover: %v", err)
		}
	}

	if err = d.Set("standalone_config_sync", flattenSystemHaStandaloneConfigSync(o["standalone-config-sync"], d, "standalone_config_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-config-sync"], "SystemHa-StandaloneConfigSync"); ok {
			if err = d.Set("standalone_config_sync", vv); err != nil {
				return fmt.Errorf("Error reading standalone_config_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_config_sync: %v", err)
		}
	}

	if err = d.Set("standalone_mgmt_vdom", flattenSystemHaStandaloneMgmtVdom(o["standalone-mgmt-vdom"], d, "standalone_mgmt_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-mgmt-vdom"], "SystemHa-StandaloneMgmtVdom"); ok {
			if err = d.Set("standalone_mgmt_vdom", vv); err != nil {
				return fmt.Errorf("Error reading standalone_mgmt_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("sync_config", flattenSystemHaSyncConfig(o["sync-config"], d, "sync_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-config"], "SystemHa-SyncConfig"); ok {
			if err = d.Set("sync_config", vv); err != nil {
				return fmt.Errorf("Error reading sync_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_config: %v", err)
		}
	}

	if err = d.Set("sync_packet_balance", flattenSystemHaSyncPacketBalance(o["sync-packet-balance"], d, "sync_packet_balance")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-packet-balance"], "SystemHa-SyncPacketBalance"); ok {
			if err = d.Set("sync_packet_balance", vv); err != nil {
				return fmt.Errorf("Error reading sync_packet_balance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_packet_balance: %v", err)
		}
	}

	if err = d.Set("unicast_gateway", flattenSystemHaUnicastGateway(o["unicast-gateway"], d, "unicast_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast-gateway"], "SystemHa-UnicastGateway"); ok {
			if err = d.Set("unicast_gateway", vv); err != nil {
				return fmt.Errorf("Error reading unicast_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast_gateway: %v", err)
		}
	}

	if err = d.Set("unicast_hb", flattenSystemHaUnicastHb(o["unicast-hb"], d, "unicast_hb")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast-hb"], "SystemHa-UnicastHb"); ok {
			if err = d.Set("unicast_hb", vv); err != nil {
				return fmt.Errorf("Error reading unicast_hb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast_hb: %v", err)
		}
	}

	if err = d.Set("unicast_hb_netmask", flattenSystemHaUnicastHbNetmask(o["unicast-hb-netmask"], d, "unicast_hb_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast-hb-netmask"], "SystemHa-UnicastHbNetmask"); ok {
			if err = d.Set("unicast_hb_netmask", vv); err != nil {
				return fmt.Errorf("Error reading unicast_hb_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast_hb_netmask: %v", err)
		}
	}

	if err = d.Set("unicast_hb_peerip", flattenSystemHaUnicastHbPeerip(o["unicast-hb-peerip"], d, "unicast_hb_peerip")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast-hb-peerip"], "SystemHa-UnicastHbPeerip"); ok {
			if err = d.Set("unicast_hb_peerip", vv); err != nil {
				return fmt.Errorf("Error reading unicast_hb_peerip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast_hb_peerip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("unicast_peers", flattenSystemHaUnicastPeers(o["unicast-peers"], d, "unicast_peers")); err != nil {
			if vv, ok := fortiAPIPatch(o["unicast-peers"], "SystemHa-UnicastPeers"); ok {
				if err = d.Set("unicast_peers", vv); err != nil {
					return fmt.Errorf("Error reading unicast_peers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading unicast_peers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("unicast_peers"); ok {
			if err = d.Set("unicast_peers", flattenSystemHaUnicastPeers(o["unicast-peers"], d, "unicast_peers")); err != nil {
				if vv, ok := fortiAPIPatch(o["unicast-peers"], "SystemHa-UnicastPeers"); ok {
					if err = d.Set("unicast_peers", vv); err != nil {
						return fmt.Errorf("Error reading unicast_peers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading unicast_peers: %v", err)
				}
			}
		}
	}

	if err = d.Set("unicast_status", flattenSystemHaUnicastStatus(o["unicast-status"], d, "unicast_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["unicast-status"], "SystemHa-UnicastStatus"); ok {
			if err = d.Set("unicast_status", vv); err != nil {
				return fmt.Errorf("Error reading unicast_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unicast_status: %v", err)
		}
	}

	if err = d.Set("uninterruptible_primary_wait", flattenSystemHaUninterruptiblePrimaryWait(o["uninterruptible-primary-wait"], d, "uninterruptible_primary_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["uninterruptible-primary-wait"], "SystemHa-UninterruptiblePrimaryWait"); ok {
			if err = d.Set("uninterruptible_primary_wait", vv); err != nil {
				return fmt.Errorf("Error reading uninterruptible_primary_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uninterruptible_primary_wait: %v", err)
		}
	}

	if err = d.Set("uninterruptible_upgrade", flattenSystemHaUninterruptibleUpgrade(o["uninterruptible-upgrade"], d, "uninterruptible_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["uninterruptible-upgrade"], "SystemHa-UninterruptibleUpgrade"); ok {
			if err = d.Set("uninterruptible_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading uninterruptible_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uninterruptible_upgrade: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemHaVclusterId(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcluster-id"], "SystemHa-VclusterId"); ok {
			if err = d.Set("vcluster_id", vv); err != nil {
				return fmt.Errorf("Error reading vcluster_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("vcluster2", flattenSystemHaVcluster2(o["vcluster2"], d, "vcluster2")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcluster2"], "SystemHa-Vcluster2"); ok {
			if err = d.Set("vcluster2", vv); err != nil {
				return fmt.Errorf("Error reading vcluster2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcluster2: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemHaVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemHa-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("upgrade_mode", flattenSystemHaUpgradeMode(o["upgrade-mode"], d, "upgrade_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["upgrade-mode"], "SystemHa-UpgradeMode"); ok {
			if err = d.Set("upgrade_mode", vv); err != nil {
				return fmt.Errorf("Error reading upgrade_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upgrade_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vcluster", flattenSystemHaVcluster(o["vcluster"], d, "vcluster")); err != nil {
			if vv, ok := fortiAPIPatch(o["vcluster"], "SystemHa-Vcluster"); ok {
				if err = d.Set("vcluster", vv); err != nil {
					return fmt.Errorf("Error reading vcluster: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vcluster: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vcluster"); ok {
			if err = d.Set("vcluster", flattenSystemHaVcluster(o["vcluster"], d, "vcluster")); err != nil {
				if vv, ok := fortiAPIPatch(o["vcluster"], "SystemHa-Vcluster"); ok {
					if err = d.Set("vcluster", vv); err != nil {
						return fmt.Errorf("Error reading vcluster: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vcluster: %v", err)
				}
			}
		}
	}

	if err = d.Set("vcluster_status", flattenSystemHaVclusterStatus(o["vcluster-status"], d, "vcluster_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcluster-status"], "SystemHa-VclusterStatus"); ok {
			if err = d.Set("vcluster_status", vv); err != nil {
				return fmt.Errorf("Error reading vcluster_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcluster_status: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemHaWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "SystemHa-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaArps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArpsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaAutoVirtualMacInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaBackupHbdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaBoardFailoverTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaChassisId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaCheckSecondaryDevHealth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFrup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFrupSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "active_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-interface"], _ = expandSystemHaFrupSettingsActiveInterface(d, i["active_interface"], pre_append)
	}
	pre_append = pre + ".0." + "active_switch_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-switch-port"], _ = expandSystemHaFrupSettingsActiveSwitchPort(d, i["active_switch_port"], pre_append)
	}
	pre_append = pre + ".0." + "backup_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backup-interface"], _ = expandSystemHaFrupSettingsBackupInterface(d, i["backup_interface"], pre_append)
	}

	return result, nil
}

func expandSystemHaFrupSettingsActiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaFrupSettingsActiveSwitchPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFrupSettingsBackupInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaEvpnTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverHoldTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGratuitousArps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandSystemHaHaMgmtInterfacesDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandSystemHaHaMgmtInterfacesGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandSystemHaHaMgmtInterfacesGateway6(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemHaHaMgmtInterfacesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemHaHaMgmtInterfacesInterface(d, i["interface"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaHaMgmtInterfacesDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemHaHaMgmtInterfacesGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaHaMgmtStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaPortDtagMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaPortOuterTpid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaUptimeDiffMargin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbIntervalInMilliseconds(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdevSecondVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdevVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHcEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHelloHolddown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHttpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHwSessionHoldTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHwSessionSyncDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHwSessionSyncDev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaImapProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInterClusterSessionSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaIpsecPhase2Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaL2EpEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLinkFailedSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLoadBalanceAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLogicalSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryBasedFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryCompatibleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverFlipTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverMonitorPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMinimumWorkerThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaMulticastTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaNntpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverFlipTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPop3ProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondarySwitchStandby(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVcluster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "monitor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["monitor"], _ = expandSystemHaSecondaryVclusterMonitor(d, i["monitor"], pre_append)
	}
	pre_append = pre + ".0." + "override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override"], _ = expandSystemHaSecondaryVclusterOverride(d, i["override"], pre_append)
	}
	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["override-wait-time"], _ = expandSystemHaSecondaryVclusterOverrideWaitTime(d, i["override_wait_time"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pingserver-failover-threshold"], _ = expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d, i["pingserver_failover_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pingserver-monitor-interface"], _ = expandSystemHaSecondaryVclusterPingserverMonitorInterface(d, i["pingserver_monitor_interface"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_secondary_force_reset"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pingserver-secondary-force-reset"], _ = expandSystemHaSecondaryVclusterPingserverSecondaryForceReset(d, i["pingserver_secondary_force_reset"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pingserver-slave-force-reset"], _ = expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d, i["pingserver_slave_force_reset"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemHaSecondaryVclusterPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vcluster-id"], _ = expandSystemHaSecondaryVclusterVclusterId(d, i["vcluster_id"], pre_append)
	}
	pre_append = pre + ".0." + "vdom"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom"], _ = expandSystemHaSecondaryVclusterVdom(d, i["vdom"], pre_append)
	}

	return result, nil
}

func expandSystemHaSecondaryVclusterMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaSecondaryVclusterOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaSecondaryVclusterPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaSessionPickup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupConnectionless(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupExpectation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionSyncDev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaSlaveSwitchStandby(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSmtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSsdFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneConfigSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneMgmtVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncPacketBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbPeerip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastPeers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemHaUnicastPeersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer-ip"], _ = expandSystemHaUnicastPeersPeerIp(d, i["peer_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaUnicastPeersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastPeersPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUninterruptiblePrimaryWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUninterruptibleUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVcluster2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaUpgradeMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVcluster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["monitor"], _ = expandSystemHaVclusterMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override"], _ = expandSystemHaVclusterOverride(d, i["override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_wait_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-wait-time"], _ = expandSystemHaVclusterOverrideWaitTime(d, i["override_wait_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_failover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pingserver-failover-threshold"], _ = expandSystemHaVclusterPingserverFailoverThreshold(d, i["pingserver_failover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_flip_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pingserver-flip-timeout"], _ = expandSystemHaVclusterPingserverFlipTimeout(d, i["pingserver_flip_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_monitor_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pingserver-monitor-interface"], _ = expandSystemHaVclusterPingserverMonitorInterface(d, i["pingserver_monitor_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_secondary_force_reset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pingserver-secondary-force-reset"], _ = expandSystemHaVclusterPingserverSecondaryForceReset(d, i["pingserver_secondary_force_reset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_slave_force_reset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pingserver-slave-force-reset"], _ = expandSystemHaVclusterPingserverSlaveForceReset(d, i["pingserver_slave_force_reset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSystemHaVclusterPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vcluster_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vcluster-id"], _ = expandSystemHaVclusterVclusterId(d, i["vcluster_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSystemHaVclusterVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemHaVclusterMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaVclusterOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFlipTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaVclusterPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaVclusterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("arps"); ok || d.HasChange("arps") {
		t, err := expandSystemHaArps(d, v, "arps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arps"] = t
		}
	}

	if v, ok := d.GetOk("arps_interval"); ok || d.HasChange("arps_interval") {
		t, err := expandSystemHaArpsInterval(d, v, "arps_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arps-interval"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystemHaAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("auto_virtual_mac_interface"); ok || d.HasChange("auto_virtual_mac_interface") {
		t, err := expandSystemHaAutoVirtualMacInterface(d, v, "auto_virtual_mac_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-virtual-mac-interface"] = t
		}
	}

	if v, ok := d.GetOk("backup_hbdev"); ok || d.HasChange("backup_hbdev") {
		t, err := expandSystemHaBackupHbdev(d, v, "backup_hbdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-hbdev"] = t
		}
	}

	if v, ok := d.GetOk("board_failover_tolerance"); ok || d.HasChange("board_failover_tolerance") {
		t, err := expandSystemHaBoardFailoverTolerance(d, v, "board_failover_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["board-failover-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("chassis_id"); ok || d.HasChange("chassis_id") {
		t, err := expandSystemHaChassisId(d, v, "chassis_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chassis-id"] = t
		}
	}

	if v, ok := d.GetOk("check_secondary_dev_health"); ok || d.HasChange("check_secondary_dev_health") {
		t, err := expandSystemHaCheckSecondaryDevHealth(d, v, "check_secondary_dev_health")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-secondary-dev-health"] = t
		}
	}

	if v, ok := d.GetOk("cpu_threshold"); ok || d.HasChange("cpu_threshold") {
		t, err := expandSystemHaCpuThreshold(d, v, "cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandSystemHaEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("frup"); ok || d.HasChange("frup") {
		t, err := expandSystemHaFrup(d, v, "frup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frup"] = t
		}
	}

	if v, ok := d.GetOk("frup_settings"); ok || d.HasChange("frup_settings") {
		t, err := expandSystemHaFrupSettings(d, v, "frup_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frup-settings"] = t
		}
	}

	if v, ok := d.GetOk("evpn_ttl"); ok || d.HasChange("evpn_ttl") {
		t, err := expandSystemHaEvpnTtl(d, v, "evpn_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["evpn-ttl"] = t
		}
	}

	if v, ok := d.GetOk("failover_hold_time"); ok || d.HasChange("failover_hold_time") {
		t, err := expandSystemHaFailoverHoldTime(d, v, "failover_hold_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failover-hold-time"] = t
		}
	}

	if v, ok := d.GetOk("ftp_proxy_threshold"); ok || d.HasChange("ftp_proxy_threshold") {
		t, err := expandSystemHaFtpProxyThreshold(d, v, "ftp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("gratuitous_arps"); ok || d.HasChange("gratuitous_arps") {
		t, err := expandSystemHaGratuitousArps(d, v, "gratuitous_arps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gratuitous-arps"] = t
		}
	}

	if v, ok := d.GetOk("group_id"); ok || d.HasChange("group_id") {
		t, err := expandSystemHaGroupId(d, v, "group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-id"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandSystemHaGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("ha_direct"); ok || d.HasChange("ha_direct") {
		t, err := expandSystemHaHaDirect(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("ha_eth_type"); ok || d.HasChange("ha_eth_type") {
		t, err := expandSystemHaHaEthType(d, v, "ha_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("ha_mgmt_interfaces"); ok || d.HasChange("ha_mgmt_interfaces") {
		t, err := expandSystemHaHaMgmtInterfaces(d, v, "ha_mgmt_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("ha_mgmt_status"); ok || d.HasChange("ha_mgmt_status") {
		t, err := expandSystemHaHaMgmtStatus(d, v, "ha_mgmt_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-status"] = t
		}
	}

	if v, ok := d.GetOk("ha_port_dtag_mode"); ok || d.HasChange("ha_port_dtag_mode") {
		t, err := expandSystemHaHaPortDtagMode(d, v, "ha_port_dtag_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-port-dtag-mode"] = t
		}
	}

	if v, ok := d.GetOk("ha_port_outer_tpid"); ok || d.HasChange("ha_port_outer_tpid") {
		t, err := expandSystemHaHaPortOuterTpid(d, v, "ha_port_outer_tpid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-port-outer-tpid"] = t
		}
	}

	if v, ok := d.GetOk("ha_uptime_diff_margin"); ok || d.HasChange("ha_uptime_diff_margin") {
		t, err := expandSystemHaHaUptimeDiffMargin(d, v, "ha_uptime_diff_margin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-uptime-diff-margin"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemHaHbInterval(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval_in_milliseconds"); ok || d.HasChange("hb_interval_in_milliseconds") {
		t, err := expandSystemHaHbIntervalInMilliseconds(d, v, "hb_interval_in_milliseconds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval-in-milliseconds"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok || d.HasChange("hb_lost_threshold") {
		t, err := expandSystemHaHbLostThreshold(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("hbdev"); ok || d.HasChange("hbdev") {
		t, err := expandSystemHaHbdev(d, v, "hbdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev"] = t
		}
	}

	if v, ok := d.GetOk("hbdev_second_vlan_id"); ok || d.HasChange("hbdev_second_vlan_id") {
		t, err := expandSystemHaHbdevSecondVlanId(d, v, "hbdev_second_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev-second-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("hbdev_vlan_id"); ok || d.HasChange("hbdev_vlan_id") {
		t, err := expandSystemHaHbdevVlanId(d, v, "hbdev_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("hc_eth_type"); ok || d.HasChange("hc_eth_type") {
		t, err := expandSystemHaHcEthType(d, v, "hc_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hc-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("hello_holddown"); ok || d.HasChange("hello_holddown") {
		t, err := expandSystemHaHelloHolddown(d, v, "hello_holddown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-holddown"] = t
		}
	}

	if v, ok := d.GetOk("http_proxy_threshold"); ok || d.HasChange("http_proxy_threshold") {
		t, err := expandSystemHaHttpProxyThreshold(d, v, "http_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("hw_session_hold_time"); ok || d.HasChange("hw_session_hold_time") {
		t, err := expandSystemHaHwSessionHoldTime(d, v, "hw_session_hold_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-session-hold-time"] = t
		}
	}

	if v, ok := d.GetOk("hw_session_sync_delay"); ok || d.HasChange("hw_session_sync_delay") {
		t, err := expandSystemHaHwSessionSyncDelay(d, v, "hw_session_sync_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-session-sync-delay"] = t
		}
	}

	if v, ok := d.GetOk("hw_session_sync_dev"); ok || d.HasChange("hw_session_sync_dev") {
		t, err := expandSystemHaHwSessionSyncDev(d, v, "hw_session_sync_dev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("imap_proxy_threshold"); ok || d.HasChange("imap_proxy_threshold") {
		t, err := expandSystemHaImapProxyThreshold(d, v, "imap_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("inter_cluster_session_sync"); ok || d.HasChange("inter_cluster_session_sync") {
		t, err := expandSystemHaInterClusterSessionSync(d, v, "inter_cluster_session_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-cluster-session-sync"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2_proposal"); ok || d.HasChange("ipsec_phase2_proposal") {
		t, err := expandSystemHaIpsecPhase2Proposal(d, v, "ipsec_phase2_proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2-proposal"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystemHaKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("l2ep_eth_type"); ok || d.HasChange("l2ep_eth_type") {
		t, err := expandSystemHaL2EpEthType(d, v, "l2ep_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2ep-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("link_failed_signal"); ok || d.HasChange("link_failed_signal") {
		t, err := expandSystemHaLinkFailedSignal(d, v, "link_failed_signal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-failed-signal"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_all"); ok || d.HasChange("load_balance_all") {
		t, err := expandSystemHaLoadBalanceAll(d, v, "load_balance_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-all"] = t
		}
	}

	if v, ok := d.GetOk("logical_sn"); ok || d.HasChange("logical_sn") {
		t, err := expandSystemHaLogicalSn(d, v, "logical_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logical-sn"] = t
		}
	}

	if v, ok := d.GetOk("memory_based_failover"); ok || d.HasChange("memory_based_failover") {
		t, err := expandSystemHaMemoryBasedFailover(d, v, "memory_based_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-based-failover"] = t
		}
	}

	if v, ok := d.GetOk("memory_compatible_mode"); ok || d.HasChange("memory_compatible_mode") {
		t, err := expandSystemHaMemoryCompatibleMode(d, v, "memory_compatible_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-compatible-mode"] = t
		}
	}

	if v, ok := d.GetOk("memory_failover_flip_timeout"); ok || d.HasChange("memory_failover_flip_timeout") {
		t, err := expandSystemHaMemoryFailoverFlipTimeout(d, v, "memory_failover_flip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-failover-flip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("memory_failover_monitor_period"); ok || d.HasChange("memory_failover_monitor_period") {
		t, err := expandSystemHaMemoryFailoverMonitorPeriod(d, v, "memory_failover_monitor_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-failover-monitor-period"] = t
		}
	}

	if v, ok := d.GetOk("memory_failover_sample_rate"); ok || d.HasChange("memory_failover_sample_rate") {
		t, err := expandSystemHaMemoryFailoverSampleRate(d, v, "memory_failover_sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-failover-sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("memory_failover_threshold"); ok || d.HasChange("memory_failover_threshold") {
		t, err := expandSystemHaMemoryFailoverThreshold(d, v, "memory_failover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-failover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("memory_threshold"); ok || d.HasChange("memory_threshold") {
		t, err := expandSystemHaMemoryThreshold(d, v, "memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-threshold"] = t
		}
	}

	if v, ok := d.GetOk("minimum_worker_threshold"); ok || d.HasChange("minimum_worker_threshold") {
		t, err := expandSystemHaMinimumWorkerThreshold(d, v, "minimum_worker_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-worker-threshold"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemHaMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandSystemHaMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl"); ok || d.HasChange("multicast_ttl") {
		t, err := expandSystemHaMulticastTtl(d, v, "multicast_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl"] = t
		}
	}

	if v, ok := d.GetOk("nntp_proxy_threshold"); ok || d.HasChange("nntp_proxy_threshold") {
		t, err := expandSystemHaNntpProxyThreshold(d, v, "nntp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandSystemHaOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("override_wait_time"); ok || d.HasChange("override_wait_time") {
		t, err := expandSystemHaOverrideWaitTime(d, v, "override_wait_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wait-time"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemHaPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_failover_threshold"); ok || d.HasChange("pingserver_failover_threshold") {
		t, err := expandSystemHaPingserverFailoverThreshold(d, v, "pingserver_failover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-failover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_flip_timeout"); ok || d.HasChange("pingserver_flip_timeout") {
		t, err := expandSystemHaPingserverFlipTimeout(d, v, "pingserver_flip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-flip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_monitor_interface"); ok || d.HasChange("pingserver_monitor_interface") {
		t, err := expandSystemHaPingserverMonitorInterface(d, v, "pingserver_monitor_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-monitor-interface"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_secondary_force_reset"); ok || d.HasChange("pingserver_secondary_force_reset") {
		t, err := expandSystemHaPingserverSecondaryForceReset(d, v, "pingserver_secondary_force_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-secondary-force-reset"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_slave_force_reset"); ok || d.HasChange("pingserver_slave_force_reset") {
		t, err := expandSystemHaPingserverSlaveForceReset(d, v, "pingserver_slave_force_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-slave-force-reset"] = t
		}
	}

	if v, ok := d.GetOk("pop3_proxy_threshold"); ok || d.HasChange("pop3_proxy_threshold") {
		t, err := expandSystemHaPop3ProxyThreshold(d, v, "pop3_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemHaPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("route_hold"); ok || d.HasChange("route_hold") {
		t, err := expandSystemHaRouteHold(d, v, "route_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-hold"] = t
		}
	}

	if v, ok := d.GetOk("route_ttl"); ok || d.HasChange("route_ttl") {
		t, err := expandSystemHaRouteTtl(d, v, "route_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-ttl"] = t
		}
	}

	if v, ok := d.GetOk("route_wait"); ok || d.HasChange("route_wait") {
		t, err := expandSystemHaRouteWait(d, v, "route_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-wait"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandSystemHaSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("secondary_switch_standby"); ok || d.HasChange("secondary_switch_standby") {
		t, err := expandSystemHaSecondarySwitchStandby(d, v, "secondary_switch_standby")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-switch-standby"] = t
		}
	}

	if v, ok := d.GetOk("secondary_vcluster"); ok || d.HasChange("secondary_vcluster") {
		t, err := expandSystemHaSecondaryVcluster(d, v, "secondary_vcluster")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-vcluster"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup"); ok || d.HasChange("session_pickup") {
		t, err := expandSystemHaSessionPickup(d, v, "session_pickup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_connectionless"); ok || d.HasChange("session_pickup_connectionless") {
		t, err := expandSystemHaSessionPickupConnectionless(d, v, "session_pickup_connectionless")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-connectionless"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_delay"); ok || d.HasChange("session_pickup_delay") {
		t, err := expandSystemHaSessionPickupDelay(d, v, "session_pickup_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-delay"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_expectation"); ok || d.HasChange("session_pickup_expectation") {
		t, err := expandSystemHaSessionPickupExpectation(d, v, "session_pickup_expectation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-expectation"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_nat"); ok || d.HasChange("session_pickup_nat") {
		t, err := expandSystemHaSessionPickupNat(d, v, "session_pickup_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-nat"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok || d.HasChange("session_sync_dev") {
		t, err := expandSystemHaSessionSyncDev(d, v, "session_sync_dev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("slave_switch_standby"); ok || d.HasChange("slave_switch_standby") {
		t, err := expandSystemHaSlaveSwitchStandby(d, v, "slave_switch_standby")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slave-switch-standby"] = t
		}
	}

	if v, ok := d.GetOk("smtp_proxy_threshold"); ok || d.HasChange("smtp_proxy_threshold") {
		t, err := expandSystemHaSmtpProxyThreshold(d, v, "smtp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ssd_failover"); ok || d.HasChange("ssd_failover") {
		t, err := expandSystemHaSsdFailover(d, v, "ssd_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-failover"] = t
		}
	}

	if v, ok := d.GetOk("standalone_config_sync"); ok || d.HasChange("standalone_config_sync") {
		t, err := expandSystemHaStandaloneConfigSync(d, v, "standalone_config_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-config-sync"] = t
		}
	}

	if v, ok := d.GetOk("standalone_mgmt_vdom"); ok || d.HasChange("standalone_mgmt_vdom") {
		t, err := expandSystemHaStandaloneMgmtVdom(d, v, "standalone_mgmt_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-mgmt-vdom"] = t
		}
	}

	if v, ok := d.GetOk("sync_config"); ok || d.HasChange("sync_config") {
		t, err := expandSystemHaSyncConfig(d, v, "sync_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-config"] = t
		}
	}

	if v, ok := d.GetOk("sync_packet_balance"); ok || d.HasChange("sync_packet_balance") {
		t, err := expandSystemHaSyncPacketBalance(d, v, "sync_packet_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-packet-balance"] = t
		}
	}

	if v, ok := d.GetOk("unicast_gateway"); ok || d.HasChange("unicast_gateway") {
		t, err := expandSystemHaUnicastGateway(d, v, "unicast_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-gateway"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb"); ok || d.HasChange("unicast_hb") {
		t, err := expandSystemHaUnicastHb(d, v, "unicast_hb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb_netmask"); ok || d.HasChange("unicast_hb_netmask") {
		t, err := expandSystemHaUnicastHbNetmask(d, v, "unicast_hb_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb-netmask"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb_peerip"); ok || d.HasChange("unicast_hb_peerip") {
		t, err := expandSystemHaUnicastHbPeerip(d, v, "unicast_hb_peerip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb-peerip"] = t
		}
	}

	if v, ok := d.GetOk("unicast_peers"); ok || d.HasChange("unicast_peers") {
		t, err := expandSystemHaUnicastPeers(d, v, "unicast_peers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-peers"] = t
		}
	}

	if v, ok := d.GetOk("unicast_status"); ok || d.HasChange("unicast_status") {
		t, err := expandSystemHaUnicastStatus(d, v, "unicast_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-status"] = t
		}
	}

	if v, ok := d.GetOk("uninterruptible_primary_wait"); ok || d.HasChange("uninterruptible_primary_wait") {
		t, err := expandSystemHaUninterruptiblePrimaryWait(d, v, "uninterruptible_primary_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uninterruptible-primary-wait"] = t
		}
	}

	if v, ok := d.GetOk("uninterruptible_upgrade"); ok || d.HasChange("uninterruptible_upgrade") {
		t, err := expandSystemHaUninterruptibleUpgrade(d, v, "uninterruptible_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uninterruptible-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("vcluster_id"); ok || d.HasChange("vcluster_id") {
		t, err := expandSystemHaVclusterId(d, v, "vcluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	if v, ok := d.GetOk("vcluster2"); ok || d.HasChange("vcluster2") {
		t, err := expandSystemHaVcluster2(d, v, "vcluster2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster2"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemHaVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_mode"); ok || d.HasChange("upgrade_mode") {
		t, err := expandSystemHaUpgradeMode(d, v, "upgrade_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-mode"] = t
		}
	}

	if v, ok := d.GetOk("vcluster"); ok || d.HasChange("vcluster") {
		t, err := expandSystemHaVcluster(d, v, "vcluster")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster"] = t
		}
	}

	if v, ok := d.GetOk("vcluster_status"); ok || d.HasChange("vcluster_status") {
		t, err := expandSystemHaVclusterStatus(d, v, "vcluster_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-status"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandSystemHaWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
