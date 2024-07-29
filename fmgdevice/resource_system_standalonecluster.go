// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterUpdate,
		Read:   resourceSystemStandaloneClusterRead,
		Update: resourceSystemStandaloneClusterUpdate,
		Delete: resourceSystemStandaloneClusterDelete,

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
			"asymmetric_traffic_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"down_intfs_before_sess_sync": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"hb_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hb_lost_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ike_heartbeat_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ike_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ike_monitor_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ike_use_rfc6311": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_tunnel_sync": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"peerip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"peervd": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_add_ipsec_routes": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"session_sync_filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_service": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_port_range": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"src_port_range": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"dstaddr": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"dstaddr6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstintf": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"srcaddr": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"srcaddr6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcintf": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"sync_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"syncvd": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"data_intf_session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_member_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"layer2_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"standalone_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemStandaloneClusterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemStandaloneCluster(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneCluster(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemStandaloneCluster")

	return resourceSystemStandaloneClusterRead(d, m)
}

func resourceSystemStandaloneClusterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemStandaloneCluster(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemStandaloneCluster(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneCluster(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterAsymmetricTrafficControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "down_intfs_before_sess_sync"
		if _, ok := i["down-intfs-before-sess-sync"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(i["down-intfs-before-sess-sync"], d, pre_append)
			tmp["down_intfs_before_sess_sync"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-DownIntfsBeforeSessSync")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_interval"
		if _, ok := i["hb-interval"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerHbInterval(i["hb-interval"], d, pre_append)
			tmp["hb_interval"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-HbInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_lost_threshold"
		if _, ok := i["hb-lost-threshold"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerHbLostThreshold(i["hb-lost-threshold"], d, pre_append)
			tmp["hb_lost_threshold"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-HbLostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_heartbeat_interval"
		if _, ok := i["ike-heartbeat-interval"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerIkeHeartbeatInterval(i["ike-heartbeat-interval"], d, pre_append)
			tmp["ike_heartbeat_interval"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-IkeHeartbeatInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_monitor"
		if _, ok := i["ike-monitor"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerIkeMonitor(i["ike-monitor"], d, pre_append)
			tmp["ike_monitor"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-IkeMonitor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_monitor_interval"
		if _, ok := i["ike-monitor-interval"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerIkeMonitorInterval(i["ike-monitor-interval"], d, pre_append)
			tmp["ike_monitor_interval"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-IkeMonitorInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_use_rfc6311"
		if _, ok := i["ike-use-rfc6311"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerIkeUseRfc6311(i["ike-use-rfc6311"], d, pre_append)
			tmp["ike_use_rfc6311"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-IkeUseRfc6311")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_tunnel_sync"
		if _, ok := i["ipsec-tunnel-sync"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync(i["ipsec-tunnel-sync"], d, pre_append)
			tmp["ipsec_tunnel_sync"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-IpsecTunnelSync")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peerip"
		if _, ok := i["peerip"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerPeerip(i["peerip"], d, pre_append)
			tmp["peerip"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-Peerip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peervd"
		if _, ok := i["peervd"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerPeervd(i["peervd"], d, pre_append)
			tmp["peervd"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-Peervd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_add_ipsec_routes"
		if _, ok := i["secondary-add-ipsec-routes"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(i["secondary-add-ipsec-routes"], d, pre_append)
			tmp["secondary_add_ipsec_routes"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-SecondaryAddIpsecRoutes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_sync_filter"
		if _, ok := i["session-sync-filter"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(i["session-sync-filter"], d, pre_append)
			tmp["session_sync_filter"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-SessionSyncFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sync_id"
		if _, ok := i["sync-id"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSyncId(i["sync-id"], d, pre_append)
			tmp["sync_id"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-SyncId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syncvd"
		if _, ok := i["syncvd"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSyncvd(i["syncvd"], d, pre_append)
			tmp["syncvd"] = fortiAPISubPartPatch(v, "SystemStandaloneCluster-ClusterPeer-Syncvd")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeUseRfc6311(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerPeervd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(i["custom-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(i["dstintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(i["srcintf"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := i["dst-port-range"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(i["dst-port-range"], d, pre_append)
			tmp["dst_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-DstPortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(i["src-port-range"], d, pre_append)
			tmp["src_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-SrcPortRange")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSyncId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSyncvd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterDataIntfSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterGroupMemberId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterLayer2Connection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterStandaloneGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemStandaloneCluster(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("asymmetric_traffic_control", flattenSystemStandaloneClusterAsymmetricTrafficControl(o["asymmetric-traffic-control"], d, "asymmetric_traffic_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["asymmetric-traffic-control"], "SystemStandaloneCluster-AsymmetricTrafficControl"); ok {
			if err = d.Set("asymmetric_traffic_control", vv); err != nil {
				return fmt.Errorf("Error reading asymmetric_traffic_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asymmetric_traffic_control: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cluster_peer", flattenSystemStandaloneClusterClusterPeer(o["cluster-peer"], d, "cluster_peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["cluster-peer"], "SystemStandaloneCluster-ClusterPeer"); ok {
				if err = d.Set("cluster_peer", vv); err != nil {
					return fmt.Errorf("Error reading cluster_peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cluster_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cluster_peer"); ok {
			if err = d.Set("cluster_peer", flattenSystemStandaloneClusterClusterPeer(o["cluster-peer"], d, "cluster_peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["cluster-peer"], "SystemStandaloneCluster-ClusterPeer"); ok {
					if err = d.Set("cluster_peer", vv); err != nil {
						return fmt.Errorf("Error reading cluster_peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cluster_peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("data_intf_session_sync_dev", flattenSystemStandaloneClusterDataIntfSessionSyncDev(o["data-intf-session-sync-dev"], d, "data_intf_session_sync_dev")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-intf-session-sync-dev"], "SystemStandaloneCluster-DataIntfSessionSyncDev"); ok {
			if err = d.Set("data_intf_session_sync_dev", vv); err != nil {
				return fmt.Errorf("Error reading data_intf_session_sync_dev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_intf_session_sync_dev: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemStandaloneClusterEncryption(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "SystemStandaloneCluster-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("group_member_id", flattenSystemStandaloneClusterGroupMemberId(o["group-member-id"], d, "group_member_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-member-id"], "SystemStandaloneCluster-GroupMemberId"); ok {
			if err = d.Set("group_member_id", vv); err != nil {
				return fmt.Errorf("Error reading group_member_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_member_id: %v", err)
		}
	}

	if err = d.Set("layer2_connection", flattenSystemStandaloneClusterLayer2Connection(o["layer2-connection"], d, "layer2_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["layer2-connection"], "SystemStandaloneCluster-Layer2Connection"); ok {
			if err = d.Set("layer2_connection", vv); err != nil {
				return fmt.Errorf("Error reading layer2_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading layer2_connection: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemStandaloneClusterSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-sync-dev"], "SystemStandaloneCluster-SessionSyncDev"); ok {
			if err = d.Set("session_sync_dev", vv); err != nil {
				return fmt.Errorf("Error reading session_sync_dev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("standalone_group_id", flattenSystemStandaloneClusterStandaloneGroupId(o["standalone-group-id"], d, "standalone_group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-group-id"], "SystemStandaloneCluster-StandaloneGroupId"); ok {
			if err = d.Set("standalone_group_id", vv); err != nil {
				return fmt.Errorf("Error reading standalone_group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_group_id: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStandaloneClusterAsymmetricTrafficControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "down_intfs_before_sess_sync"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["down-intfs-before-sess-sync"], _ = expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d, i["down_intfs_before_sess_sync"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hb-interval"], _ = expandSystemStandaloneClusterClusterPeerHbInterval(d, i["hb_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_lost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hb-lost-threshold"], _ = expandSystemStandaloneClusterClusterPeerHbLostThreshold(d, i["hb_lost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_heartbeat_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ike-heartbeat-interval"], _ = expandSystemStandaloneClusterClusterPeerIkeHeartbeatInterval(d, i["ike_heartbeat_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_monitor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ike-monitor"], _ = expandSystemStandaloneClusterClusterPeerIkeMonitor(d, i["ike_monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_monitor_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ike-monitor-interval"], _ = expandSystemStandaloneClusterClusterPeerIkeMonitorInterval(d, i["ike_monitor_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ike_use_rfc6311"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ike-use-rfc6311"], _ = expandSystemStandaloneClusterClusterPeerIkeUseRfc6311(d, i["ike_use_rfc6311"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_tunnel_sync"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-tunnel-sync"], _ = expandSystemStandaloneClusterClusterPeerIpsecTunnelSync(d, i["ipsec_tunnel_sync"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peerip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peerip"], _ = expandSystemStandaloneClusterClusterPeerPeerip(d, i["peerip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peervd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peervd"], _ = expandSystemStandaloneClusterClusterPeerPeervd(d, i["peervd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_add_ipsec_routes"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-add-ipsec-routes"], _ = expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(d, i["secondary_add_ipsec_routes"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_sync_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d, i["session_sync_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["session-sync-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sync_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sync-id"], _ = expandSystemStandaloneClusterClusterPeerSyncId(d, i["sync_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syncvd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["syncvd"], _ = expandSystemStandaloneClusterClusterPeerSyncvd(d, i["syncvd"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeUseRfc6311(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIpsecTunnelSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeerip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeervd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d, i["custom_service"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-service"] = t
		}
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(d, i["dstaddr"], pre_append)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(d, i["dstaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(d, i["dstintf"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(d, i["srcaddr"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(d, i["srcaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(d, i["srcintf"], pre_append)
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(d, i["dst_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(d, i["src_port_range"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSyncId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSyncvd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterDataIntfSessionSyncDev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterGroupMemberId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterLayer2Connection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterPsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterSessionSyncDev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterStandaloneGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStandaloneCluster(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("asymmetric_traffic_control"); ok || d.HasChange("asymmetric_traffic_control") {
		t, err := expandSystemStandaloneClusterAsymmetricTrafficControl(d, v, "asymmetric_traffic_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymmetric-traffic-control"] = t
		}
	}

	if v, ok := d.GetOk("cluster_peer"); ok || d.HasChange("cluster_peer") {
		t, err := expandSystemStandaloneClusterClusterPeer(d, v, "cluster_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cluster-peer"] = t
		}
	}

	if v, ok := d.GetOk("data_intf_session_sync_dev"); ok || d.HasChange("data_intf_session_sync_dev") {
		t, err := expandSystemStandaloneClusterDataIntfSessionSyncDev(d, v, "data_intf_session_sync_dev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-intf-session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandSystemStandaloneClusterEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("group_member_id"); ok || d.HasChange("group_member_id") {
		t, err := expandSystemStandaloneClusterGroupMemberId(d, v, "group_member_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-id"] = t
		}
	}

	if v, ok := d.GetOk("layer2_connection"); ok || d.HasChange("layer2_connection") {
		t, err := expandSystemStandaloneClusterLayer2Connection(d, v, "layer2_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer2-connection"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok || d.HasChange("psksecret") {
		t, err := expandSystemStandaloneClusterPsksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok || d.HasChange("session_sync_dev") {
		t, err := expandSystemStandaloneClusterSessionSyncDev(d, v, "session_sync_dev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("standalone_group_id"); ok || d.HasChange("standalone_group_id") {
		t, err := expandSystemStandaloneClusterStandaloneGroupId(d, v, "standalone_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-group-id"] = t
		}
	}

	return &obj, nil
}
