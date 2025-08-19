// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS global parameter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsGlobalUpdate,
		Read:   resourceIpsGlobalRead,
		Update: resourceIpsGlobalUpdate,
		Delete: resourceIpsGlobalDelete,

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
			"anomaly_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_mem_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cp_accel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deep_app_insp_db_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"deep_app_insp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"engine_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"exclude_signatures": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_open": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intelligent_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_reserve_cpu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ngfw_max_scan_range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"np_accel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"packet_log_queue_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"session_limit_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skype_client_public_ipaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"socket_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sync_session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_active_probe": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip6": &schema.Schema{
							Type:     schema.TypeString,
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
			"traffic_submit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectIpsGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateIpsGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("IpsGlobal")

	return resourceIpsGlobalRead(d, m)
}

func resourceIpsGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIpsGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIpsGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource from API: %v", err)
	}
	return nil
}

func flattenIpsGlobalAnomalyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalAvMemLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalCpAccelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspDbLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalEngineCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalExcludeSignatures(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalFailOpen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalIntelligentMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalIpsReserveCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalNgfwMaxScanRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalNpAccelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalPacketLogQueueDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSessionLimitMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSkypeClientPublicIpaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSocketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSyncSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbe(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interface"
	if _, ok := i["interface"]; ok {
		result["interface"] = flattenIpsGlobalTlsActiveProbeInterface(i["interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_select_method"
	if _, ok := i["interface-select-method"]; ok {
		result["interface_select_method"] = flattenIpsGlobalTlsActiveProbeInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip"
	if _, ok := i["source-ip"]; ok {
		result["source_ip"] = flattenIpsGlobalTlsActiveProbeSourceIp(i["source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip6"
	if _, ok := i["source-ip6"]; ok {
		result["source_ip6"] = flattenIpsGlobalTlsActiveProbeSourceIp6(i["source-ip6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {
		result["vdom"] = flattenIpsGlobalTlsActiveProbeVdom(i["vdom"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenIpsGlobalTlsActiveProbeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsGlobalTlsActiveProbeInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsGlobalTrafficSubmit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("anomaly_mode", flattenIpsGlobalAnomalyMode(o["anomaly-mode"], d, "anomaly_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["anomaly-mode"], "IpsGlobal-AnomalyMode"); ok {
			if err = d.Set("anomaly_mode", vv); err != nil {
				return fmt.Errorf("Error reading anomaly_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anomaly_mode: %v", err)
		}
	}

	if err = d.Set("av_mem_limit", flattenIpsGlobalAvMemLimit(o["av-mem-limit"], d, "av_mem_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-mem-limit"], "IpsGlobal-AvMemLimit"); ok {
			if err = d.Set("av_mem_limit", vv); err != nil {
				return fmt.Errorf("Error reading av_mem_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_mem_limit: %v", err)
		}
	}

	if err = d.Set("cp_accel_mode", flattenIpsGlobalCpAccelMode(o["cp-accel-mode"], d, "cp_accel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["cp-accel-mode"], "IpsGlobal-CpAccelMode"); ok {
			if err = d.Set("cp_accel_mode", vv); err != nil {
				return fmt.Errorf("Error reading cp_accel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cp_accel_mode: %v", err)
		}
	}

	if err = d.Set("database", flattenIpsGlobalDatabase(o["database"], d, "database")); err != nil {
		if vv, ok := fortiAPIPatch(o["database"], "IpsGlobal-Database"); ok {
			if err = d.Set("database", vv); err != nil {
				return fmt.Errorf("Error reading database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_db_limit", flattenIpsGlobalDeepAppInspDbLimit(o["deep-app-insp-db-limit"], d, "deep_app_insp_db_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["deep-app-insp-db-limit"], "IpsGlobal-DeepAppInspDbLimit"); ok {
			if err = d.Set("deep_app_insp_db_limit", vv); err != nil {
				return fmt.Errorf("Error reading deep_app_insp_db_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deep_app_insp_db_limit: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_timeout", flattenIpsGlobalDeepAppInspTimeout(o["deep-app-insp-timeout"], d, "deep_app_insp_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["deep-app-insp-timeout"], "IpsGlobal-DeepAppInspTimeout"); ok {
			if err = d.Set("deep_app_insp_timeout", vv); err != nil {
				return fmt.Errorf("Error reading deep_app_insp_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deep_app_insp_timeout: %v", err)
		}
	}

	if err = d.Set("engine_count", flattenIpsGlobalEngineCount(o["engine-count"], d, "engine_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-count"], "IpsGlobal-EngineCount"); ok {
			if err = d.Set("engine_count", vv); err != nil {
				return fmt.Errorf("Error reading engine_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_count: %v", err)
		}
	}

	if err = d.Set("exclude_signatures", flattenIpsGlobalExcludeSignatures(o["exclude-signatures"], d, "exclude_signatures")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-signatures"], "IpsGlobal-ExcludeSignatures"); ok {
			if err = d.Set("exclude_signatures", vv); err != nil {
				return fmt.Errorf("Error reading exclude_signatures: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_signatures: %v", err)
		}
	}

	if err = d.Set("fail_open", flattenIpsGlobalFailOpen(o["fail-open"], d, "fail_open")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-open"], "IpsGlobal-FailOpen"); ok {
			if err = d.Set("fail_open", vv); err != nil {
				return fmt.Errorf("Error reading fail_open: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_open: %v", err)
		}
	}

	if err = d.Set("intelligent_mode", flattenIpsGlobalIntelligentMode(o["intelligent-mode"], d, "intelligent_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["intelligent-mode"], "IpsGlobal-IntelligentMode"); ok {
			if err = d.Set("intelligent_mode", vv); err != nil {
				return fmt.Errorf("Error reading intelligent_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intelligent_mode: %v", err)
		}
	}

	if err = d.Set("ips_reserve_cpu", flattenIpsGlobalIpsReserveCpu(o["ips-reserve-cpu"], d, "ips_reserve_cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-reserve-cpu"], "IpsGlobal-IpsReserveCpu"); ok {
			if err = d.Set("ips_reserve_cpu", vv); err != nil {
				return fmt.Errorf("Error reading ips_reserve_cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_reserve_cpu: %v", err)
		}
	}

	if err = d.Set("ngfw_max_scan_range", flattenIpsGlobalNgfwMaxScanRange(o["ngfw-max-scan-range"], d, "ngfw_max_scan_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["ngfw-max-scan-range"], "IpsGlobal-NgfwMaxScanRange"); ok {
			if err = d.Set("ngfw_max_scan_range", vv); err != nil {
				return fmt.Errorf("Error reading ngfw_max_scan_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ngfw_max_scan_range: %v", err)
		}
	}

	if err = d.Set("np_accel_mode", flattenIpsGlobalNpAccelMode(o["np-accel-mode"], d, "np_accel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["np-accel-mode"], "IpsGlobal-NpAccelMode"); ok {
			if err = d.Set("np_accel_mode", vv); err != nil {
				return fmt.Errorf("Error reading np_accel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np_accel_mode: %v", err)
		}
	}

	if err = d.Set("packet_log_queue_depth", flattenIpsGlobalPacketLogQueueDepth(o["packet-log-queue-depth"], d, "packet_log_queue_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-log-queue-depth"], "IpsGlobal-PacketLogQueueDepth"); ok {
			if err = d.Set("packet_log_queue_depth", vv); err != nil {
				return fmt.Errorf("Error reading packet_log_queue_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_log_queue_depth: %v", err)
		}
	}

	if err = d.Set("session_limit_mode", flattenIpsGlobalSessionLimitMode(o["session-limit-mode"], d, "session_limit_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-limit-mode"], "IpsGlobal-SessionLimitMode"); ok {
			if err = d.Set("session_limit_mode", vv); err != nil {
				return fmt.Errorf("Error reading session_limit_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_limit_mode: %v", err)
		}
	}

	if err = d.Set("skype_client_public_ipaddr", flattenIpsGlobalSkypeClientPublicIpaddr(o["skype-client-public-ipaddr"], d, "skype_client_public_ipaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["skype-client-public-ipaddr"], "IpsGlobal-SkypeClientPublicIpaddr"); ok {
			if err = d.Set("skype_client_public_ipaddr", vv); err != nil {
				return fmt.Errorf("Error reading skype_client_public_ipaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skype_client_public_ipaddr: %v", err)
		}
	}

	if err = d.Set("socket_size", flattenIpsGlobalSocketSize(o["socket-size"], d, "socket_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["socket-size"], "IpsGlobal-SocketSize"); ok {
			if err = d.Set("socket_size", vv); err != nil {
				return fmt.Errorf("Error reading socket_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading socket_size: %v", err)
		}
	}

	if err = d.Set("sync_session_ttl", flattenIpsGlobalSyncSessionTtl(o["sync-session-ttl"], d, "sync_session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-session-ttl"], "IpsGlobal-SyncSessionTtl"); ok {
			if err = d.Set("sync_session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading sync_session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_session_ttl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tls_active_probe", flattenIpsGlobalTlsActiveProbe(o["tls-active-probe"], d, "tls_active_probe")); err != nil {
			if vv, ok := fortiAPIPatch(o["tls-active-probe"], "IpsGlobal-TlsActiveProbe"); ok {
				if err = d.Set("tls_active_probe", vv); err != nil {
					return fmt.Errorf("Error reading tls_active_probe: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tls_active_probe: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tls_active_probe"); ok {
			if err = d.Set("tls_active_probe", flattenIpsGlobalTlsActiveProbe(o["tls-active-probe"], d, "tls_active_probe")); err != nil {
				if vv, ok := fortiAPIPatch(o["tls-active-probe"], "IpsGlobal-TlsActiveProbe"); ok {
					if err = d.Set("tls_active_probe", vv); err != nil {
						return fmt.Errorf("Error reading tls_active_probe: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tls_active_probe: %v", err)
				}
			}
		}
	}

	if err = d.Set("traffic_submit", flattenIpsGlobalTrafficSubmit(o["traffic-submit"], d, "traffic_submit")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-submit"], "IpsGlobal-TrafficSubmit"); ok {
			if err = d.Set("traffic_submit", vv); err != nil {
				return fmt.Errorf("Error reading traffic_submit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_submit: %v", err)
		}
	}

	return nil
}

func flattenIpsGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsGlobalAnomalyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalAvMemLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalCpAccelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspDbLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalEngineCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalExcludeSignatures(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalFailOpen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalIntelligentMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalIpsReserveCpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalNgfwMaxScanRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalNpAccelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalPacketLogQueueDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSessionLimitMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSkypeClientPublicIpaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSocketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSyncSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface"], _ = expandIpsGlobalTlsActiveProbeInterface(d, i["interface"], pre_append)
	}
	pre_append = pre + ".0." + "interface_select_method"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interface-select-method"], _ = expandIpsGlobalTlsActiveProbeInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-ip"], _ = expandIpsGlobalTlsActiveProbeSourceIp(d, i["source_ip"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-ip6"], _ = expandIpsGlobalTlsActiveProbeSourceIp6(d, i["source_ip6"], pre_append)
	}
	pre_append = pre + ".0." + "vdom"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom"], _ = expandIpsGlobalTlsActiveProbeVdom(d, i["vdom"], pre_append)
	}

	return result, nil
}

func expandIpsGlobalTlsActiveProbeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsGlobalTlsActiveProbeInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsGlobalTrafficSubmit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly_mode"); ok || d.HasChange("anomaly_mode") {
		t, err := expandIpsGlobalAnomalyMode(d, v, "anomaly_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly-mode"] = t
		}
	}

	if v, ok := d.GetOk("av_mem_limit"); ok || d.HasChange("av_mem_limit") {
		t, err := expandIpsGlobalAvMemLimit(d, v, "av_mem_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-mem-limit"] = t
		}
	}

	if v, ok := d.GetOk("cp_accel_mode"); ok || d.HasChange("cp_accel_mode") {
		t, err := expandIpsGlobalCpAccelMode(d, v, "cp_accel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cp-accel-mode"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok || d.HasChange("database") {
		t, err := expandIpsGlobalDatabase(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("deep_app_insp_db_limit"); ok || d.HasChange("deep_app_insp_db_limit") {
		t, err := expandIpsGlobalDeepAppInspDbLimit(d, v, "deep_app_insp_db_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-db-limit"] = t
		}
	}

	if v, ok := d.GetOk("deep_app_insp_timeout"); ok || d.HasChange("deep_app_insp_timeout") {
		t, err := expandIpsGlobalDeepAppInspTimeout(d, v, "deep_app_insp_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-timeout"] = t
		}
	}

	if v, ok := d.GetOk("engine_count"); ok || d.HasChange("engine_count") {
		t, err := expandIpsGlobalEngineCount(d, v, "engine_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-count"] = t
		}
	}

	if v, ok := d.GetOk("exclude_signatures"); ok || d.HasChange("exclude_signatures") {
		t, err := expandIpsGlobalExcludeSignatures(d, v, "exclude_signatures")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-signatures"] = t
		}
	}

	if v, ok := d.GetOk("fail_open"); ok || d.HasChange("fail_open") {
		t, err := expandIpsGlobalFailOpen(d, v, "fail_open")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-open"] = t
		}
	}

	if v, ok := d.GetOk("intelligent_mode"); ok || d.HasChange("intelligent_mode") {
		t, err := expandIpsGlobalIntelligentMode(d, v, "intelligent_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intelligent-mode"] = t
		}
	}

	if v, ok := d.GetOk("ips_reserve_cpu"); ok || d.HasChange("ips_reserve_cpu") {
		t, err := expandIpsGlobalIpsReserveCpu(d, v, "ips_reserve_cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-reserve-cpu"] = t
		}
	}

	if v, ok := d.GetOk("ngfw_max_scan_range"); ok || d.HasChange("ngfw_max_scan_range") {
		t, err := expandIpsGlobalNgfwMaxScanRange(d, v, "ngfw_max_scan_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ngfw-max-scan-range"] = t
		}
	}

	if v, ok := d.GetOk("np_accel_mode"); ok || d.HasChange("np_accel_mode") {
		t, err := expandIpsGlobalNpAccelMode(d, v, "np_accel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-accel-mode"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_queue_depth"); ok || d.HasChange("packet_log_queue_depth") {
		t, err := expandIpsGlobalPacketLogQueueDepth(d, v, "packet_log_queue_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-queue-depth"] = t
		}
	}

	if v, ok := d.GetOk("session_limit_mode"); ok || d.HasChange("session_limit_mode") {
		t, err := expandIpsGlobalSessionLimitMode(d, v, "session_limit_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-limit-mode"] = t
		}
	}

	if v, ok := d.GetOk("skype_client_public_ipaddr"); ok || d.HasChange("skype_client_public_ipaddr") {
		t, err := expandIpsGlobalSkypeClientPublicIpaddr(d, v, "skype_client_public_ipaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skype-client-public-ipaddr"] = t
		}
	}

	if v, ok := d.GetOk("socket_size"); ok || d.HasChange("socket_size") {
		t, err := expandIpsGlobalSocketSize(d, v, "socket_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socket-size"] = t
		}
	}

	if v, ok := d.GetOk("sync_session_ttl"); ok || d.HasChange("sync_session_ttl") {
		t, err := expandIpsGlobalSyncSessionTtl(d, v, "sync_session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tls_active_probe"); ok || d.HasChange("tls_active_probe") {
		t, err := expandIpsGlobalTlsActiveProbe(d, v, "tls_active_probe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-active-probe"] = t
		}
	}

	if v, ok := d.GetOk("traffic_submit"); ok || d.HasChange("traffic_submit") {
		t, err := expandIpsGlobalTrafficSubmit(d, v, "traffic_submit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-submit"] = t
		}
	}

	return &obj, nil
}
