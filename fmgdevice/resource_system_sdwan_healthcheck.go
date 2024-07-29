// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanHealthCheckCreate,
		Read:   resourceSystemSdwanHealthCheckRead,
		Update: resourceSystemSdwanHealthCheckUpdate,
		Delete: resourceSystemSdwanHealthCheckDelete,

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
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"detect_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_match_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"embed_measured_health": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ftp_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ftp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mos_codec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"packet_size": &schema.Schema{
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
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"probe_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probe_packets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quality_measured_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sla": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"jitter_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_cost_factor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packetloss_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_in_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_out_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sla_fail_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla_id_redistribute": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla_pass_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_alert_jitter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_alert_latency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_alert_packetloss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_jitter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_latency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_packetloss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_cascade_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf": &schema.Schema{
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

func resourceSystemSdwanHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheck resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanHealthCheck(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheck resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdwanHealthCheckRead(d, m)
}

func resourceSystemSdwanHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheck resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanHealthCheck(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheck resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdwanHealthCheckRead(d, m)
}

func resourceSystemSdwanHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdwanHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanHealthCheckRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanHealthCheck(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheck resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanHealthCheckAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckDetectMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDiffservcode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsMatchIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsRequestDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckEmbedMeasuredHealth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFailtime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFtpFile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFtpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHaPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpAgent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemSdwanHealthCheckHttpGet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckMosCodec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPacketSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbePackets2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckQualityMeasuredMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckRecoverytime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSecurityMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckSla2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSdwanHealthCheckSlaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaJitterThreshold2edl(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaLatencyThreshold2edl(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenSystemSdwanHealthCheckSlaLinkCostFactor2edl(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaMosThreshold2edl(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPacketlossThreshold2edl(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPriorityInSla2edl(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPriorityOutSla2edl(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckSlaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaJitterThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLinkCostFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckSlaMosThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPacketlossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityInSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityOutSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaFailLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaIdRedistribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPassLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSource62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSystemDns2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertJitter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertLatency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertPacketloss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningJitter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningLatency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningPacketloss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateCascadeInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateStaticRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanHealthCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenSystemSdwanHealthCheckAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "SystemSdwanHealthCheck-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("class_id", flattenSystemSdwanHealthCheckClassId2edl(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "SystemSdwanHealthCheck-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("detect_mode", flattenSystemSdwanHealthCheckDetectMode2edl(o["detect-mode"], d, "detect_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-mode"], "SystemSdwanHealthCheck-DetectMode"); ok {
			if err = d.Set("detect_mode", vv); err != nil {
				return fmt.Errorf("Error reading detect_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemSdwanHealthCheckDiffservcode2edl(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "SystemSdwanHealthCheck-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", flattenSystemSdwanHealthCheckDnsMatchIp2edl(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-match-ip"], "SystemSdwanHealthCheck-DnsMatchIp"); ok {
			if err = d.Set("dns_match_ip", vv); err != nil {
				return fmt.Errorf("Error reading dns_match_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenSystemSdwanHealthCheckDnsRequestDomain2edl(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-request-domain"], "SystemSdwanHealthCheck-DnsRequestDomain"); ok {
			if err = d.Set("dns_request_domain", vv); err != nil {
				return fmt.Errorf("Error reading dns_request_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("embed_measured_health", flattenSystemSdwanHealthCheckEmbedMeasuredHealth2edl(o["embed-measured-health"], d, "embed_measured_health")); err != nil {
		if vv, ok := fortiAPIPatch(o["embed-measured-health"], "SystemSdwanHealthCheck-EmbedMeasuredHealth"); ok {
			if err = d.Set("embed_measured_health", vv); err != nil {
				return fmt.Errorf("Error reading embed_measured_health: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embed_measured_health: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemSdwanHealthCheckFailtime2edl(o["failtime"], d, "failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["failtime"], "SystemSdwanHealthCheck-Failtime"); ok {
			if err = d.Set("failtime", vv); err != nil {
				return fmt.Errorf("Error reading failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("ftp_file", flattenSystemSdwanHealthCheckFtpFile2edl(o["ftp-file"], d, "ftp_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-file"], "SystemSdwanHealthCheck-FtpFile"); ok {
			if err = d.Set("ftp_file", vv); err != nil {
				return fmt.Errorf("Error reading ftp_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_file: %v", err)
		}
	}

	if err = d.Set("ftp_mode", flattenSystemSdwanHealthCheckFtpMode2edl(o["ftp-mode"], d, "ftp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-mode"], "SystemSdwanHealthCheck-FtpMode"); ok {
			if err = d.Set("ftp_mode", vv); err != nil {
				return fmt.Errorf("Error reading ftp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenSystemSdwanHealthCheckHaPriority2edl(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "SystemSdwanHealthCheck-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenSystemSdwanHealthCheckHttpAgent2edl(o["http-agent"], d, "http_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-agent"], "SystemSdwanHealthCheck-HttpAgent"); ok {
			if err = d.Set("http_agent", vv); err != nil {
				return fmt.Errorf("Error reading http_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenSystemSdwanHealthCheckHttpGet2edl(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "SystemSdwanHealthCheck-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenSystemSdwanHealthCheckHttpMatch2edl(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "SystemSdwanHealthCheck-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemSdwanHealthCheckInterval2edl(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "SystemSdwanHealthCheck-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("members", flattenSystemSdwanHealthCheckMembers2edl(o["members"], d, "members")); err != nil {
		if vv, ok := fortiAPIPatch(o["members"], "SystemSdwanHealthCheck-Members"); ok {
			if err = d.Set("members", vv); err != nil {
				return fmt.Errorf("Error reading members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading members: %v", err)
		}
	}

	if err = d.Set("mos_codec", flattenSystemSdwanHealthCheckMosCodec2edl(o["mos-codec"], d, "mos_codec")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-codec"], "SystemSdwanHealthCheck-MosCodec"); ok {
			if err = d.Set("mos_codec", vv); err != nil {
				return fmt.Errorf("Error reading mos_codec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_codec: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdwanHealthCheckName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdwanHealthCheck-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenSystemSdwanHealthCheckPacketSize2edl(o["packet-size"], d, "packet_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-size"], "SystemSdwanHealthCheck-PacketSize"); ok {
			if err = d.Set("packet_size", vv); err != nil {
				return fmt.Errorf("Error reading packet_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSdwanHealthCheckPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSdwanHealthCheck-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", flattenSystemSdwanHealthCheckProbeCount2edl(o["probe-count"], d, "probe_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-count"], "SystemSdwanHealthCheck-ProbeCount"); ok {
			if err = d.Set("probe_count", vv); err != nil {
				return fmt.Errorf("Error reading probe_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_packets", flattenSystemSdwanHealthCheckProbePackets2edl(o["probe-packets"], d, "probe_packets")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-packets"], "SystemSdwanHealthCheck-ProbePackets"); ok {
			if err = d.Set("probe_packets", vv); err != nil {
				return fmt.Errorf("Error reading probe_packets: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenSystemSdwanHealthCheckProbeTimeout2edl(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-timeout"], "SystemSdwanHealthCheck-ProbeTimeout"); ok {
			if err = d.Set("probe_timeout", vv); err != nil {
				return fmt.Errorf("Error reading probe_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemSdwanHealthCheckProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemSdwanHealthCheck-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_measured_method", flattenSystemSdwanHealthCheckQualityMeasuredMethod2edl(o["quality-measured-method"], d, "quality_measured_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-measured-method"], "SystemSdwanHealthCheck-QualityMeasuredMethod"); ok {
			if err = d.Set("quality_measured_method", vv); err != nil {
				return fmt.Errorf("Error reading quality_measured_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_measured_method: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenSystemSdwanHealthCheckRecoverytime2edl(o["recoverytime"], d, "recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["recoverytime"], "SystemSdwanHealthCheck-Recoverytime"); ok {
			if err = d.Set("recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemSdwanHealthCheckSecurityMode2edl(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "SystemSdwanHealthCheck-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdwanHealthCheckServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemSdwanHealthCheck-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenSystemSdwanHealthCheckSla2edl(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanHealthCheck-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenSystemSdwanHealthCheckSla2edl(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanHealthCheck-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_fail_log_period", flattenSystemSdwanHealthCheckSlaFailLogPeriod2edl(o["sla-fail-log-period"], d, "sla_fail_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-fail-log-period"], "SystemSdwanHealthCheck-SlaFailLogPeriod"); ok {
			if err = d.Set("sla_fail_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_id_redistribute", flattenSystemSdwanHealthCheckSlaIdRedistribute2edl(o["sla-id-redistribute"], d, "sla_id_redistribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-id-redistribute"], "SystemSdwanHealthCheck-SlaIdRedistribute"); ok {
			if err = d.Set("sla_id_redistribute", vv); err != nil {
				return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", flattenSystemSdwanHealthCheckSlaPassLogPeriod2edl(o["sla-pass-log-period"], d, "sla_pass_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-pass-log-period"], "SystemSdwanHealthCheck-SlaPassLogPeriod"); ok {
			if err = d.Set("sla_pass_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSdwanHealthCheckSource2edl(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "SystemSdwanHealthCheck-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenSystemSdwanHealthCheckSource62edl(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "SystemSdwanHealthCheck-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("system_dns", flattenSystemSdwanHealthCheckSystemDns2edl(o["system-dns"], d, "system_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-dns"], "SystemSdwanHealthCheck-SystemDns"); ok {
			if err = d.Set("system_dns", vv); err != nil {
				return fmt.Errorf("Error reading system_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_dns: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", flattenSystemSdwanHealthCheckThresholdAlertJitter2edl(o["threshold-alert-jitter"], d, "threshold_alert_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-jitter"], "SystemSdwanHealthCheck-ThresholdAlertJitter"); ok {
			if err = d.Set("threshold_alert_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", flattenSystemSdwanHealthCheckThresholdAlertLatency2edl(o["threshold-alert-latency"], d, "threshold_alert_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-latency"], "SystemSdwanHealthCheck-ThresholdAlertLatency"); ok {
			if err = d.Set("threshold_alert_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", flattenSystemSdwanHealthCheckThresholdAlertPacketloss2edl(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-packetloss"], "SystemSdwanHealthCheck-ThresholdAlertPacketloss"); ok {
			if err = d.Set("threshold_alert_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", flattenSystemSdwanHealthCheckThresholdWarningJitter2edl(o["threshold-warning-jitter"], d, "threshold_warning_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-jitter"], "SystemSdwanHealthCheck-ThresholdWarningJitter"); ok {
			if err = d.Set("threshold_warning_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", flattenSystemSdwanHealthCheckThresholdWarningLatency2edl(o["threshold-warning-latency"], d, "threshold_warning_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-latency"], "SystemSdwanHealthCheck-ThresholdWarningLatency"); ok {
			if err = d.Set("threshold_warning_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", flattenSystemSdwanHealthCheckThresholdWarningPacketloss2edl(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-packetloss"], "SystemSdwanHealthCheck-ThresholdWarningPacketloss"); ok {
			if err = d.Set("threshold_warning_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenSystemSdwanHealthCheckUpdateCascadeInterface2edl(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-cascade-interface"], "SystemSdwanHealthCheck-UpdateCascadeInterface"); ok {
			if err = d.Set("update_cascade_interface", vv); err != nil {
				return fmt.Errorf("Error reading update_cascade_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenSystemSdwanHealthCheckUpdateStaticRoute2edl(o["update-static-route"], d, "update_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-static-route"], "SystemSdwanHealthCheck-UpdateStaticRoute"); ok {
			if err = d.Set("update_static_route", vv); err != nil {
				return fmt.Errorf("Error reading update_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemSdwanHealthCheckUser2edl(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemSdwanHealthCheck-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSystemSdwanHealthCheckVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "SystemSdwanHealthCheck-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanHealthCheckFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanHealthCheckAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckDetectMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDiffservcode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsMatchIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsRequestDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckEmbedMeasuredHealth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFailtime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFtpFile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFtpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHaPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpAgent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpGet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckMosCodec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPacketSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbePackets2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckQualityMeasuredMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckRecoverytime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSecurityMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSdwanHealthCheckSlaId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandSystemSdwanHealthCheckSlaJitterThreshold2edl(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandSystemSdwanHealthCheckSlaLatencyThreshold2edl(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandSystemSdwanHealthCheckSlaLinkCostFactor2edl(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandSystemSdwanHealthCheckSlaMosThreshold2edl(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandSystemSdwanHealthCheckSlaPacketlossThreshold2edl(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandSystemSdwanHealthCheckSlaPriorityInSla2edl(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandSystemSdwanHealthCheckSlaPriorityOutSla2edl(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckSlaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaJitterThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLinkCostFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckSlaMosThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPacketlossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityInSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityOutSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaFailLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaIdRedistribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPassLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSource62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSystemDns2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertJitter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertLatency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertPacketloss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningJitter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningLatency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningPacketloss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateCascadeInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateStaticRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanHealthCheck(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandSystemSdwanHealthCheckAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandSystemSdwanHealthCheckClassId2edl(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("detect_mode"); ok || d.HasChange("detect_mode") {
		t, err := expandSystemSdwanHealthCheckDetectMode2edl(d, v, "detect_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-mode"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandSystemSdwanHealthCheckDiffservcode2edl(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dns_match_ip"); ok || d.HasChange("dns_match_ip") {
		t, err := expandSystemSdwanHealthCheckDnsMatchIp2edl(d, v, "dns_match_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok || d.HasChange("dns_request_domain") {
		t, err := expandSystemSdwanHealthCheckDnsRequestDomain2edl(d, v, "dns_request_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("embed_measured_health"); ok || d.HasChange("embed_measured_health") {
		t, err := expandSystemSdwanHealthCheckEmbedMeasuredHealth2edl(d, v, "embed_measured_health")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embed-measured-health"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok || d.HasChange("failtime") {
		t, err := expandSystemSdwanHealthCheckFailtime2edl(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("ftp_file"); ok || d.HasChange("ftp_file") {
		t, err := expandSystemSdwanHealthCheckFtpFile2edl(d, v, "ftp_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-file"] = t
		}
	}

	if v, ok := d.GetOk("ftp_mode"); ok || d.HasChange("ftp_mode") {
		t, err := expandSystemSdwanHealthCheckFtpMode2edl(d, v, "ftp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-mode"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandSystemSdwanHealthCheckHaPriority2edl(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok || d.HasChange("http_agent") {
		t, err := expandSystemSdwanHealthCheckHttpAgent2edl(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandSystemSdwanHealthCheckHttpGet2edl(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandSystemSdwanHealthCheckHttpMatch2edl(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandSystemSdwanHealthCheckInterval2edl(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandSystemSdwanHealthCheckMembers2edl(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("mos_codec"); ok || d.HasChange("mos_codec") {
		t, err := expandSystemSdwanHealthCheckMosCodec2edl(d, v, "mos_codec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-codec"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdwanHealthCheckName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok || d.HasChange("packet_size") {
		t, err := expandSystemSdwanHealthCheckPacketSize2edl(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemSdwanHealthCheckPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSdwanHealthCheckPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_count"); ok || d.HasChange("probe_count") {
		t, err := expandSystemSdwanHealthCheckProbeCount2edl(d, v, "probe_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-count"] = t
		}
	}

	if v, ok := d.GetOk("probe_packets"); ok || d.HasChange("probe_packets") {
		t, err := expandSystemSdwanHealthCheckProbePackets2edl(d, v, "probe_packets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-packets"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok || d.HasChange("probe_timeout") {
		t, err := expandSystemSdwanHealthCheckProbeTimeout2edl(d, v, "probe_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemSdwanHealthCheckProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_measured_method"); ok || d.HasChange("quality_measured_method") {
		t, err := expandSystemSdwanHealthCheckQualityMeasuredMethod2edl(d, v, "quality_measured_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-measured-method"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok || d.HasChange("recoverytime") {
		t, err := expandSystemSdwanHealthCheckRecoverytime2edl(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandSystemSdwanHealthCheckSecurityMode2edl(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemSdwanHealthCheckServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandSystemSdwanHealthCheckSla2edl(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_fail_log_period"); ok || d.HasChange("sla_fail_log_period") {
		t, err := expandSystemSdwanHealthCheckSlaFailLogPeriod2edl(d, v, "sla_fail_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-fail-log-period"] = t
		}
	}

	if v, ok := d.GetOk("sla_id_redistribute"); ok || d.HasChange("sla_id_redistribute") {
		t, err := expandSystemSdwanHealthCheckSlaIdRedistribute2edl(d, v, "sla_id_redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id-redistribute"] = t
		}
	}

	if v, ok := d.GetOk("sla_pass_log_period"); ok || d.HasChange("sla_pass_log_period") {
		t, err := expandSystemSdwanHealthCheckSlaPassLogPeriod2edl(d, v, "sla_pass_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-pass-log-period"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandSystemSdwanHealthCheckSource2edl(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandSystemSdwanHealthCheckSource62edl(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("system_dns"); ok || d.HasChange("system_dns") {
		t, err := expandSystemSdwanHealthCheckSystemDns2edl(d, v, "system_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-dns"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_jitter"); ok || d.HasChange("threshold_alert_jitter") {
		t, err := expandSystemSdwanHealthCheckThresholdAlertJitter2edl(d, v, "threshold_alert_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_latency"); ok || d.HasChange("threshold_alert_latency") {
		t, err := expandSystemSdwanHealthCheckThresholdAlertLatency2edl(d, v, "threshold_alert_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_packetloss"); ok || d.HasChange("threshold_alert_packetloss") {
		t, err := expandSystemSdwanHealthCheckThresholdAlertPacketloss2edl(d, v, "threshold_alert_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_jitter"); ok || d.HasChange("threshold_warning_jitter") {
		t, err := expandSystemSdwanHealthCheckThresholdWarningJitter2edl(d, v, "threshold_warning_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_latency"); ok || d.HasChange("threshold_warning_latency") {
		t, err := expandSystemSdwanHealthCheckThresholdWarningLatency2edl(d, v, "threshold_warning_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_packetloss"); ok || d.HasChange("threshold_warning_packetloss") {
		t, err := expandSystemSdwanHealthCheckThresholdWarningPacketloss2edl(d, v, "threshold_warning_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok || d.HasChange("update_cascade_interface") {
		t, err := expandSystemSdwanHealthCheckUpdateCascadeInterface2edl(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok || d.HasChange("update_static_route") {
		t, err := expandSystemSdwanHealthCheckUpdateStaticRoute2edl(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemSdwanHealthCheckUser2edl(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandSystemSdwanHealthCheckVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
