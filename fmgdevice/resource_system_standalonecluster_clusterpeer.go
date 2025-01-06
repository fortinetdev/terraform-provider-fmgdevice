// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneClusterClusterPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterClusterPeerCreate,
		Read:   resourceSystemStandaloneClusterClusterPeerRead,
		Update: resourceSystemStandaloneClusterClusterPeerUpdate,
		Delete: resourceSystemStandaloneClusterClusterPeerDelete,

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
				Computed: true,
			},
			"ike_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ike_use_rfc6311": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				ForceNew: true,
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
	}
}

func resourceSystemStandaloneClusterClusterPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemStandaloneClusterClusterPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterClusterPeer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemStandaloneClusterClusterPeer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterClusterPeer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "sync_id")))

	return resourceSystemStandaloneClusterClusterPeerRead(d, m)
}

func resourceSystemStandaloneClusterClusterPeerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemStandaloneClusterClusterPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneClusterClusterPeer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "sync_id")))

	return resourceSystemStandaloneClusterClusterPeerRead(d, m)
}

func resourceSystemStandaloneClusterClusterPeerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemStandaloneClusterClusterPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneClusterClusterPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterClusterPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemStandaloneClusterClusterPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneClusterClusterPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeer resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerHbInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerHbLostThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeHeartbeatInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeMonitorInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIkeUseRfc63112edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerPeerip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerPeervd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService2edl(i["custom-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr2edl(i["dstaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr62edl(i["dstaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf2edl(i["dstintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr2edl(i["srcaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr62edl(i["srcaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf2edl(i["srcintf"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange2edl(i["dst-port-range"], d, pre_append)
			tmp["dst_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-DstPortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			v := flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange2edl(i["src-port-range"], d, pre_append)
			tmp["src_port_range"] = fortiAPISubPartPatch(v, "SystemStandaloneClusterClusterPeerSessionSyncFilter-CustomService-SrcPortRange")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemStandaloneClusterClusterPeerSyncId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSyncvd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemStandaloneClusterClusterPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("down_intfs_before_sess_sync", flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync2edl(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["down-intfs-before-sess-sync"], "SystemStandaloneClusterClusterPeer-DownIntfsBeforeSessSync"); ok {
			if err = d.Set("down_intfs_before_sess_sync", vv); err != nil {
				return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemStandaloneClusterClusterPeerHbInterval2edl(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemStandaloneClusterClusterPeer-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemStandaloneClusterClusterPeerHbLostThreshold2edl(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-lost-threshold"], "SystemStandaloneClusterClusterPeer-HbLostThreshold"); ok {
			if err = d.Set("hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("ike_heartbeat_interval", flattenSystemStandaloneClusterClusterPeerIkeHeartbeatInterval2edl(o["ike-heartbeat-interval"], d, "ike_heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-heartbeat-interval"], "SystemStandaloneClusterClusterPeer-IkeHeartbeatInterval"); ok {
			if err = d.Set("ike_heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("ike_monitor", flattenSystemStandaloneClusterClusterPeerIkeMonitor2edl(o["ike-monitor"], d, "ike_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-monitor"], "SystemStandaloneClusterClusterPeer-IkeMonitor"); ok {
			if err = d.Set("ike_monitor", vv); err != nil {
				return fmt.Errorf("Error reading ike_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_monitor: %v", err)
		}
	}

	if err = d.Set("ike_monitor_interval", flattenSystemStandaloneClusterClusterPeerIkeMonitorInterval2edl(o["ike-monitor-interval"], d, "ike_monitor_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-monitor-interval"], "SystemStandaloneClusterClusterPeer-IkeMonitorInterval"); ok {
			if err = d.Set("ike_monitor_interval", vv); err != nil {
				return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
		}
	}

	if err = d.Set("ike_use_rfc6311", flattenSystemStandaloneClusterClusterPeerIkeUseRfc63112edl(o["ike-use-rfc6311"], d, "ike_use_rfc6311")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-use-rfc6311"], "SystemStandaloneClusterClusterPeer-IkeUseRfc6311"); ok {
			if err = d.Set("ike_use_rfc6311", vv); err != nil {
				return fmt.Errorf("Error reading ike_use_rfc6311: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_use_rfc6311: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel_sync", flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync2edl(o["ipsec-tunnel-sync"], d, "ipsec_tunnel_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel-sync"], "SystemStandaloneClusterClusterPeer-IpsecTunnelSync"); ok {
			if err = d.Set("ipsec_tunnel_sync", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
		}
	}

	if err = d.Set("peerip", flattenSystemStandaloneClusterClusterPeerPeerip2edl(o["peerip"], d, "peerip")); err != nil {
		if vv, ok := fortiAPIPatch(o["peerip"], "SystemStandaloneClusterClusterPeer-Peerip"); ok {
			if err = d.Set("peerip", vv); err != nil {
				return fmt.Errorf("Error reading peerip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peerip: %v", err)
		}
	}

	if err = d.Set("peervd", flattenSystemStandaloneClusterClusterPeerPeervd2edl(o["peervd"], d, "peervd")); err != nil {
		if vv, ok := fortiAPIPatch(o["peervd"], "SystemStandaloneClusterClusterPeer-Peervd"); ok {
			if err = d.Set("peervd", vv); err != nil {
				return fmt.Errorf("Error reading peervd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peervd: %v", err)
		}
	}

	if err = d.Set("secondary_add_ipsec_routes", flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes2edl(o["secondary-add-ipsec-routes"], d, "secondary_add_ipsec_routes")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-add-ipsec-routes"], "SystemStandaloneClusterClusterPeer-SecondaryAddIpsecRoutes"); ok {
			if err = d.Set("secondary_add_ipsec_routes", vv); err != nil {
				return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("session_sync_filter", flattenSystemStandaloneClusterClusterPeerSessionSyncFilter2edl(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["session-sync-filter"], "SystemStandaloneClusterClusterPeer-SessionSyncFilter"); ok {
				if err = d.Set("session_sync_filter", vv); err != nil {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading session_sync_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("session_sync_filter"); ok {
			if err = d.Set("session_sync_filter", flattenSystemStandaloneClusterClusterPeerSessionSyncFilter2edl(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["session-sync-filter"], "SystemStandaloneClusterClusterPeer-SessionSyncFilter"); ok {
					if err = d.Set("session_sync_filter", vv); err != nil {
						return fmt.Errorf("Error reading session_sync_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("sync_id", flattenSystemStandaloneClusterClusterPeerSyncId2edl(o["sync-id"], d, "sync_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-id"], "SystemStandaloneClusterClusterPeer-SyncId"); ok {
			if err = d.Set("sync_id", vv); err != nil {
				return fmt.Errorf("Error reading sync_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_id: %v", err)
		}
	}

	if err = d.Set("syncvd", flattenSystemStandaloneClusterClusterPeerSyncvd2edl(o["syncvd"], d, "syncvd")); err != nil {
		if vv, ok := fortiAPIPatch(o["syncvd"], "SystemStandaloneClusterClusterPeer-Syncvd"); ok {
			if err = d.Set("syncvd", vv); err != nil {
				return fmt.Errorf("Error reading syncvd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syncvd: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterClusterPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerHbInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerHbLostThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeHeartbeatInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeMonitorInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIkeUseRfc63112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIpsecTunnelSync2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeerip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeervd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService2edl(d, i["custom_service"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-service"] = t
		}
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr2edl(d, i["dstaddr"], pre_append)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr62edl(d, i["dstaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf2edl(d, i["dstintf"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr2edl(d, i["srcaddr"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr62edl(d, i["srcaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf2edl(d, i["srcintf"], pre_append)
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dst-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange2edl(d, i["dst_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange2edl(d, i["src_port_range"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemStandaloneClusterClusterPeerSyncId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSyncvd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemStandaloneClusterClusterPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("down_intfs_before_sess_sync"); ok || d.HasChange("down_intfs_before_sess_sync") {
		t, err := expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync2edl(d, v, "down_intfs_before_sess_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["down-intfs-before-sess-sync"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemStandaloneClusterClusterPeerHbInterval2edl(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok || d.HasChange("hb_lost_threshold") {
		t, err := expandSystemStandaloneClusterClusterPeerHbLostThreshold2edl(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ike_heartbeat_interval"); ok || d.HasChange("ike_heartbeat_interval") {
		t, err := expandSystemStandaloneClusterClusterPeerIkeHeartbeatInterval2edl(d, v, "ike_heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor"); ok || d.HasChange("ike_monitor") {
		t, err := expandSystemStandaloneClusterClusterPeerIkeMonitor2edl(d, v, "ike_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor_interval"); ok || d.HasChange("ike_monitor_interval") {
		t, err := expandSystemStandaloneClusterClusterPeerIkeMonitorInterval2edl(d, v, "ike_monitor_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor-interval"] = t
		}
	}

	if v, ok := d.GetOk("ike_use_rfc6311"); ok || d.HasChange("ike_use_rfc6311") {
		t, err := expandSystemStandaloneClusterClusterPeerIkeUseRfc63112edl(d, v, "ike_use_rfc6311")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-use-rfc6311"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel_sync"); ok || d.HasChange("ipsec_tunnel_sync") {
		t, err := expandSystemStandaloneClusterClusterPeerIpsecTunnelSync2edl(d, v, "ipsec_tunnel_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel-sync"] = t
		}
	}

	if v, ok := d.GetOk("peerip"); ok || d.HasChange("peerip") {
		t, err := expandSystemStandaloneClusterClusterPeerPeerip2edl(d, v, "peerip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerip"] = t
		}
	}

	if v, ok := d.GetOk("peervd"); ok || d.HasChange("peervd") {
		t, err := expandSystemStandaloneClusterClusterPeerPeervd2edl(d, v, "peervd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peervd"] = t
		}
	}

	if v, ok := d.GetOk("secondary_add_ipsec_routes"); ok || d.HasChange("secondary_add_ipsec_routes") {
		t, err := expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes2edl(d, v, "secondary_add_ipsec_routes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-add-ipsec-routes"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_filter"); ok || d.HasChange("session_sync_filter") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilter2edl(d, v, "session_sync_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-filter"] = t
		}
	}

	if v, ok := d.GetOk("sync_id"); ok || d.HasChange("sync_id") {
		t, err := expandSystemStandaloneClusterClusterPeerSyncId2edl(d, v, "sync_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-id"] = t
		}
	}

	if v, ok := d.GetOk("syncvd"); ok || d.HasChange("syncvd") {
		t, err := expandSystemStandaloneClusterClusterPeerSyncvd2edl(d, v, "syncvd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncvd"] = t
		}
	}

	return &obj, nil
}
