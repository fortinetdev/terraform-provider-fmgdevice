// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemClusterSync

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemClusterSyncCreate,
		Read:   resourceSystemClusterSyncRead,
		Update: resourceSystemClusterSyncUpdate,
		Delete: resourceSystemClusterSyncDelete,

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
			},
			"ike_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ike_seqjump_speed": &schema.Schema{
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
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"slave_add_ike_routes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
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

func resourceSystemClusterSyncCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemClusterSync(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource while getting object: %v", err)
	}

	_, err = c.CreateSystemClusterSync(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "sync_id")))

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemClusterSync(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemClusterSync(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "sync_id")))

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemClusterSync(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemClusterSync resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemClusterSync(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemClusterSync(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource from API: %v", err)
	}
	return nil
}

func flattenSystemClusterSyncDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeSeqjumpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeUseRfc6311(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncIpsecTunnelSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncPeervd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSecondaryAddIpsecRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemClusterSyncSessionSyncFilterCustomService(i["custom-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemClusterSyncSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemClusterSyncSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemClusterSyncSessionSyncFilterDstintf(i["dstintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemClusterSyncSessionSyncFilterSrcintf(i["srcintf"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemClusterSyncSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(i["dst-port-range"], d, pre_append)
			tmp["dst_port_range"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-DstPortRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			v := flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(i["src-port-range"], d, pre_append)
			tmp["src_port_range"] = fortiAPISubPartPatch(v, "SystemClusterSyncSessionSyncFilter-CustomService-SrcPortRange")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSystemClusterSyncSlaveAddIkeRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSyncId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSyncvd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemClusterSync(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["down-intfs-before-sess-sync"], "SystemClusterSync-DownIntfsBeforeSessSync"); ok {
			if err = d.Set("down_intfs_before_sess_sync", vv); err != nil {
				return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemClusterSyncHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-interval"], "SystemClusterSync-HbInterval"); ok {
			if err = d.Set("hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemClusterSyncHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["hb-lost-threshold"], "SystemClusterSync-HbLostThreshold"); ok {
			if err = d.Set("hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("ike_heartbeat_interval", flattenSystemClusterSyncIkeHeartbeatInterval(o["ike-heartbeat-interval"], d, "ike_heartbeat_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-heartbeat-interval"], "SystemClusterSync-IkeHeartbeatInterval"); ok {
			if err = d.Set("ike_heartbeat_interval", vv); err != nil {
				return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("ike_monitor", flattenSystemClusterSyncIkeMonitor(o["ike-monitor"], d, "ike_monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-monitor"], "SystemClusterSync-IkeMonitor"); ok {
			if err = d.Set("ike_monitor", vv); err != nil {
				return fmt.Errorf("Error reading ike_monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_monitor: %v", err)
		}
	}

	if err = d.Set("ike_monitor_interval", flattenSystemClusterSyncIkeMonitorInterval(o["ike-monitor-interval"], d, "ike_monitor_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-monitor-interval"], "SystemClusterSync-IkeMonitorInterval"); ok {
			if err = d.Set("ike_monitor_interval", vv); err != nil {
				return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
		}
	}

	if err = d.Set("ike_seqjump_speed", flattenSystemClusterSyncIkeSeqjumpSpeed(o["ike-seqjump-speed"], d, "ike_seqjump_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-seqjump-speed"], "SystemClusterSync-IkeSeqjumpSpeed"); ok {
			if err = d.Set("ike_seqjump_speed", vv); err != nil {
				return fmt.Errorf("Error reading ike_seqjump_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_seqjump_speed: %v", err)
		}
	}

	if err = d.Set("ike_use_rfc6311", flattenSystemClusterSyncIkeUseRfc6311(o["ike-use-rfc6311"], d, "ike_use_rfc6311")); err != nil {
		if vv, ok := fortiAPIPatch(o["ike-use-rfc6311"], "SystemClusterSync-IkeUseRfc6311"); ok {
			if err = d.Set("ike_use_rfc6311", vv); err != nil {
				return fmt.Errorf("Error reading ike_use_rfc6311: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ike_use_rfc6311: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel_sync", flattenSystemClusterSyncIpsecTunnelSync(o["ipsec-tunnel-sync"], d, "ipsec_tunnel_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel-sync"], "SystemClusterSync-IpsecTunnelSync"); ok {
			if err = d.Set("ipsec_tunnel_sync", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
		}
	}

	if err = d.Set("peerip", flattenSystemClusterSyncPeerip(o["peerip"], d, "peerip")); err != nil {
		if vv, ok := fortiAPIPatch(o["peerip"], "SystemClusterSync-Peerip"); ok {
			if err = d.Set("peerip", vv); err != nil {
				return fmt.Errorf("Error reading peerip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peerip: %v", err)
		}
	}

	if err = d.Set("peervd", flattenSystemClusterSyncPeervd(o["peervd"], d, "peervd")); err != nil {
		if vv, ok := fortiAPIPatch(o["peervd"], "SystemClusterSync-Peervd"); ok {
			if err = d.Set("peervd", vv); err != nil {
				return fmt.Errorf("Error reading peervd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peervd: %v", err)
		}
	}

	if err = d.Set("secondary_add_ipsec_routes", flattenSystemClusterSyncSecondaryAddIpsecRoutes(o["secondary-add-ipsec-routes"], d, "secondary_add_ipsec_routes")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-add-ipsec-routes"], "SystemClusterSync-SecondaryAddIpsecRoutes"); ok {
			if err = d.Set("secondary_add_ipsec_routes", vv); err != nil {
				return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["session-sync-filter"], "SystemClusterSync-SessionSyncFilter"); ok {
				if err = d.Set("session_sync_filter", vv); err != nil {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading session_sync_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("session_sync_filter"); ok {
			if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["session-sync-filter"], "SystemClusterSync-SessionSyncFilter"); ok {
					if err = d.Set("session_sync_filter", vv); err != nil {
						return fmt.Errorf("Error reading session_sync_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("slave_add_ike_routes", flattenSystemClusterSyncSlaveAddIkeRoutes(o["slave-add-ike-routes"], d, "slave_add_ike_routes")); err != nil {
		if vv, ok := fortiAPIPatch(o["slave-add-ike-routes"], "SystemClusterSync-SlaveAddIkeRoutes"); ok {
			if err = d.Set("slave_add_ike_routes", vv); err != nil {
				return fmt.Errorf("Error reading slave_add_ike_routes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slave_add_ike_routes: %v", err)
		}
	}

	if err = d.Set("sync_id", flattenSystemClusterSyncSyncId(o["sync-id"], d, "sync_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-id"], "SystemClusterSync-SyncId"); ok {
			if err = d.Set("sync_id", vv); err != nil {
				return fmt.Errorf("Error reading sync_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_id: %v", err)
		}
	}

	if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd")); err != nil {
		if vv, ok := fortiAPIPatch(o["syncvd"], "SystemClusterSync-Syncvd"); ok {
			if err = d.Set("syncvd", vv); err != nil {
				return fmt.Errorf("Error reading syncvd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syncvd: %v", err)
		}
	}

	return nil
}

func flattenSystemClusterSyncFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemClusterSyncDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemClusterSyncHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeSeqjumpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeUseRfc6311(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIpsecTunnelSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeerip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeervd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemClusterSyncSecondaryAddIpsecRoutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandSystemClusterSyncSessionSyncFilterCustomService(d, i["custom_service"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-service"] = t
		}
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr(d, i["dstaddr"], pre_append)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstaddr6"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr6(d, i["dstaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstintf"], _ = expandSystemClusterSyncSessionSyncFilterDstintf(d, i["dstintf"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr(d, i["srcaddr"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcaddr6"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr6(d, i["srcaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcintf"], _ = expandSystemClusterSyncSessionSyncFilterSrcintf(d, i["srcintf"], pre_append)
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dst-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d, i["dst_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d, i["src_port_range"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSystemClusterSyncSlaveAddIkeRoutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSyncId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSyncvd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemClusterSync(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("down_intfs_before_sess_sync"); ok || d.HasChange("down_intfs_before_sess_sync") {
		t, err := expandSystemClusterSyncDownIntfsBeforeSessSync(d, v, "down_intfs_before_sess_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["down-intfs-before-sess-sync"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok || d.HasChange("hb_interval") {
		t, err := expandSystemClusterSyncHbInterval(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok || d.HasChange("hb_lost_threshold") {
		t, err := expandSystemClusterSyncHbLostThreshold(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ike_heartbeat_interval"); ok || d.HasChange("ike_heartbeat_interval") {
		t, err := expandSystemClusterSyncIkeHeartbeatInterval(d, v, "ike_heartbeat_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor"); ok || d.HasChange("ike_monitor") {
		t, err := expandSystemClusterSyncIkeMonitor(d, v, "ike_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor_interval"); ok || d.HasChange("ike_monitor_interval") {
		t, err := expandSystemClusterSyncIkeMonitorInterval(d, v, "ike_monitor_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor-interval"] = t
		}
	}

	if v, ok := d.GetOk("ike_seqjump_speed"); ok || d.HasChange("ike_seqjump_speed") {
		t, err := expandSystemClusterSyncIkeSeqjumpSpeed(d, v, "ike_seqjump_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-seqjump-speed"] = t
		}
	}

	if v, ok := d.GetOk("ike_use_rfc6311"); ok || d.HasChange("ike_use_rfc6311") {
		t, err := expandSystemClusterSyncIkeUseRfc6311(d, v, "ike_use_rfc6311")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-use-rfc6311"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel_sync"); ok || d.HasChange("ipsec_tunnel_sync") {
		t, err := expandSystemClusterSyncIpsecTunnelSync(d, v, "ipsec_tunnel_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel-sync"] = t
		}
	}

	if v, ok := d.GetOk("peerip"); ok || d.HasChange("peerip") {
		t, err := expandSystemClusterSyncPeerip(d, v, "peerip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerip"] = t
		}
	}

	if v, ok := d.GetOk("peervd"); ok || d.HasChange("peervd") {
		t, err := expandSystemClusterSyncPeervd(d, v, "peervd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peervd"] = t
		}
	}

	if v, ok := d.GetOk("secondary_add_ipsec_routes"); ok || d.HasChange("secondary_add_ipsec_routes") {
		t, err := expandSystemClusterSyncSecondaryAddIpsecRoutes(d, v, "secondary_add_ipsec_routes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-add-ipsec-routes"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_filter"); ok || d.HasChange("session_sync_filter") {
		t, err := expandSystemClusterSyncSessionSyncFilter(d, v, "session_sync_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-filter"] = t
		}
	}

	if v, ok := d.GetOk("slave_add_ike_routes"); ok || d.HasChange("slave_add_ike_routes") {
		t, err := expandSystemClusterSyncSlaveAddIkeRoutes(d, v, "slave_add_ike_routes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slave-add-ike-routes"] = t
		}
	}

	if v, ok := d.GetOk("sync_id"); ok || d.HasChange("sync_id") {
		t, err := expandSystemClusterSyncSyncId(d, v, "sync_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-id"] = t
		}
	}

	if v, ok := d.GetOk("syncvd"); ok || d.HasChange("syncvd") {
		t, err := expandSystemClusterSyncSyncvd(d, v, "syncvd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncvd"] = t
		}
	}

	return &obj, nil
}
