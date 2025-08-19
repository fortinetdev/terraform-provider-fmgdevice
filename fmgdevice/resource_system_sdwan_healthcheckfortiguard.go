// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanHealthCheckFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanHealthCheckFortiguardCreate,
		Read:   resourceSystemSdwanHealthCheckFortiguardRead,
		Update: resourceSystemSdwanHealthCheckFortiguardUpdate,
		Delete: resourceSystemSdwanHealthCheckFortiguardDelete,

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
				Computed: true,
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
			"packet_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
							Computed: true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
				Computed: true,
			},
			"system_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
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

func resourceSystemSdwanHealthCheckFortiguardCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdwanHealthCheckFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckFortiguard resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanHealthCheckFortiguard(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanHealthCheckFortiguard resource: %v", err)
	}

	d.SetId(getStringKey(d, "target_name"))

	return resourceSystemSdwanHealthCheckFortiguardRead(d, m)
}

func resourceSystemSdwanHealthCheckFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdwanHealthCheckFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanHealthCheckFortiguard(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanHealthCheckFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "target_name"))

	return resourceSystemSdwanHealthCheckFortiguardRead(d, m)
}

func resourceSystemSdwanHealthCheckFortiguardDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdwanHealthCheckFortiguard(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanHealthCheckFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanHealthCheckFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanHealthCheckFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanHealthCheckFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanHealthCheckFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanHealthCheckFortiguardAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardDetectMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDiffservcode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDnsMatchIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDnsRequestDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFailtime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFtpFile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFtpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHaPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHttpAgent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemSdwanHealthCheckFortiguardHttpGet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHttpMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardMosCodec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardPacketSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbeCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbePackets2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbeTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardQualityMeasuredMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardRecoverytime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSecurityMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardSla2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSdwanHealthCheckFortiguardSlaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold2edl(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold2edl(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor2edl(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold2edl(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold2edl(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla2edl(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla2edl(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckFortiguardSlaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaFailLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaIdRedistribute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPassLogPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSource62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSystemDns2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardTargetName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertJitter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertLatency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningJitter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningLatency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUpdateCascadeInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUpdateStaticRoute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanHealthCheckFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenSystemSdwanHealthCheckFortiguardAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "SystemSdwanHealthCheckFortiguard-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("class_id", flattenSystemSdwanHealthCheckFortiguardClassId2edl(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "SystemSdwanHealthCheckFortiguard-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("detect_mode", flattenSystemSdwanHealthCheckFortiguardDetectMode2edl(o["detect-mode"], d, "detect_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-mode"], "SystemSdwanHealthCheckFortiguard-DetectMode"); ok {
			if err = d.Set("detect_mode", vv); err != nil {
				return fmt.Errorf("Error reading detect_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemSdwanHealthCheckFortiguardDiffservcode2edl(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "SystemSdwanHealthCheckFortiguard-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", flattenSystemSdwanHealthCheckFortiguardDnsMatchIp2edl(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-match-ip"], "SystemSdwanHealthCheckFortiguard-DnsMatchIp"); ok {
			if err = d.Set("dns_match_ip", vv); err != nil {
				return fmt.Errorf("Error reading dns_match_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenSystemSdwanHealthCheckFortiguardDnsRequestDomain2edl(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-request-domain"], "SystemSdwanHealthCheckFortiguard-DnsRequestDomain"); ok {
			if err = d.Set("dns_request_domain", vv); err != nil {
				return fmt.Errorf("Error reading dns_request_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("embed_measured_health", flattenSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth2edl(o["embed-measured-health"], d, "embed_measured_health")); err != nil {
		if vv, ok := fortiAPIPatch(o["embed-measured-health"], "SystemSdwanHealthCheckFortiguard-EmbedMeasuredHealth"); ok {
			if err = d.Set("embed_measured_health", vv); err != nil {
				return fmt.Errorf("Error reading embed_measured_health: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embed_measured_health: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemSdwanHealthCheckFortiguardFailtime2edl(o["failtime"], d, "failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["failtime"], "SystemSdwanHealthCheckFortiguard-Failtime"); ok {
			if err = d.Set("failtime", vv); err != nil {
				return fmt.Errorf("Error reading failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("ftp_file", flattenSystemSdwanHealthCheckFortiguardFtpFile2edl(o["ftp-file"], d, "ftp_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-file"], "SystemSdwanHealthCheckFortiguard-FtpFile"); ok {
			if err = d.Set("ftp_file", vv); err != nil {
				return fmt.Errorf("Error reading ftp_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_file: %v", err)
		}
	}

	if err = d.Set("ftp_mode", flattenSystemSdwanHealthCheckFortiguardFtpMode2edl(o["ftp-mode"], d, "ftp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-mode"], "SystemSdwanHealthCheckFortiguard-FtpMode"); ok {
			if err = d.Set("ftp_mode", vv); err != nil {
				return fmt.Errorf("Error reading ftp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenSystemSdwanHealthCheckFortiguardHaPriority2edl(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "SystemSdwanHealthCheckFortiguard-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenSystemSdwanHealthCheckFortiguardHttpAgent2edl(o["http-agent"], d, "http_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-agent"], "SystemSdwanHealthCheckFortiguard-HttpAgent"); ok {
			if err = d.Set("http_agent", vv); err != nil {
				return fmt.Errorf("Error reading http_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenSystemSdwanHealthCheckFortiguardHttpGet2edl(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "SystemSdwanHealthCheckFortiguard-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenSystemSdwanHealthCheckFortiguardHttpMatch2edl(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "SystemSdwanHealthCheckFortiguard-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemSdwanHealthCheckFortiguardInterval2edl(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "SystemSdwanHealthCheckFortiguard-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("members", flattenSystemSdwanHealthCheckFortiguardMembers2edl(o["members"], d, "members")); err != nil {
		if vv, ok := fortiAPIPatch(o["members"], "SystemSdwanHealthCheckFortiguard-Members"); ok {
			if err = d.Set("members", vv); err != nil {
				return fmt.Errorf("Error reading members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading members: %v", err)
		}
	}

	if err = d.Set("mos_codec", flattenSystemSdwanHealthCheckFortiguardMosCodec2edl(o["mos-codec"], d, "mos_codec")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-codec"], "SystemSdwanHealthCheckFortiguard-MosCodec"); ok {
			if err = d.Set("mos_codec", vv); err != nil {
				return fmt.Errorf("Error reading mos_codec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_codec: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenSystemSdwanHealthCheckFortiguardPacketSize2edl(o["packet-size"], d, "packet_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-size"], "SystemSdwanHealthCheckFortiguard-PacketSize"); ok {
			if err = d.Set("packet_size", vv); err != nil {
				return fmt.Errorf("Error reading packet_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSdwanHealthCheckFortiguardPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSdwanHealthCheckFortiguard-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", flattenSystemSdwanHealthCheckFortiguardProbeCount2edl(o["probe-count"], d, "probe_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-count"], "SystemSdwanHealthCheckFortiguard-ProbeCount"); ok {
			if err = d.Set("probe_count", vv); err != nil {
				return fmt.Errorf("Error reading probe_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_packets", flattenSystemSdwanHealthCheckFortiguardProbePackets2edl(o["probe-packets"], d, "probe_packets")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-packets"], "SystemSdwanHealthCheckFortiguard-ProbePackets"); ok {
			if err = d.Set("probe_packets", vv); err != nil {
				return fmt.Errorf("Error reading probe_packets: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenSystemSdwanHealthCheckFortiguardProbeTimeout2edl(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-timeout"], "SystemSdwanHealthCheckFortiguard-ProbeTimeout"); ok {
			if err = d.Set("probe_timeout", vv); err != nil {
				return fmt.Errorf("Error reading probe_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemSdwanHealthCheckFortiguardProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemSdwanHealthCheckFortiguard-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_measured_method", flattenSystemSdwanHealthCheckFortiguardQualityMeasuredMethod2edl(o["quality-measured-method"], d, "quality_measured_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-measured-method"], "SystemSdwanHealthCheckFortiguard-QualityMeasuredMethod"); ok {
			if err = d.Set("quality_measured_method", vv); err != nil {
				return fmt.Errorf("Error reading quality_measured_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_measured_method: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenSystemSdwanHealthCheckFortiguardRecoverytime2edl(o["recoverytime"], d, "recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["recoverytime"], "SystemSdwanHealthCheckFortiguard-Recoverytime"); ok {
			if err = d.Set("recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemSdwanHealthCheckFortiguardSecurityMode2edl(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "SystemSdwanHealthCheckFortiguard-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdwanHealthCheckFortiguardServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemSdwanHealthCheckFortiguard-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenSystemSdwanHealthCheckFortiguardSla2edl(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanHealthCheckFortiguard-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenSystemSdwanHealthCheckFortiguardSla2edl(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "SystemSdwanHealthCheckFortiguard-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_fail_log_period", flattenSystemSdwanHealthCheckFortiguardSlaFailLogPeriod2edl(o["sla-fail-log-period"], d, "sla_fail_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-fail-log-period"], "SystemSdwanHealthCheckFortiguard-SlaFailLogPeriod"); ok {
			if err = d.Set("sla_fail_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_id_redistribute", flattenSystemSdwanHealthCheckFortiguardSlaIdRedistribute2edl(o["sla-id-redistribute"], d, "sla_id_redistribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-id-redistribute"], "SystemSdwanHealthCheckFortiguard-SlaIdRedistribute"); ok {
			if err = d.Set("sla_id_redistribute", vv); err != nil {
				return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", flattenSystemSdwanHealthCheckFortiguardSlaPassLogPeriod2edl(o["sla-pass-log-period"], d, "sla_pass_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-pass-log-period"], "SystemSdwanHealthCheckFortiguard-SlaPassLogPeriod"); ok {
			if err = d.Set("sla_pass_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSdwanHealthCheckFortiguardSource2edl(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "SystemSdwanHealthCheckFortiguard-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenSystemSdwanHealthCheckFortiguardSource62edl(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "SystemSdwanHealthCheckFortiguard-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("system_dns", flattenSystemSdwanHealthCheckFortiguardSystemDns2edl(o["system-dns"], d, "system_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-dns"], "SystemSdwanHealthCheckFortiguard-SystemDns"); ok {
			if err = d.Set("system_dns", vv); err != nil {
				return fmt.Errorf("Error reading system_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_dns: %v", err)
		}
	}

	if err = d.Set("target_name", flattenSystemSdwanHealthCheckFortiguardTargetName2edl(o["target-name"], d, "target_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-name"], "SystemSdwanHealthCheckFortiguard-TargetName"); ok {
			if err = d.Set("target_name", vv); err != nil {
				return fmt.Errorf("Error reading target_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_name: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", flattenSystemSdwanHealthCheckFortiguardThresholdAlertJitter2edl(o["threshold-alert-jitter"], d, "threshold_alert_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-jitter"], "SystemSdwanHealthCheckFortiguard-ThresholdAlertJitter"); ok {
			if err = d.Set("threshold_alert_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", flattenSystemSdwanHealthCheckFortiguardThresholdAlertLatency2edl(o["threshold-alert-latency"], d, "threshold_alert_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-latency"], "SystemSdwanHealthCheckFortiguard-ThresholdAlertLatency"); ok {
			if err = d.Set("threshold_alert_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", flattenSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss2edl(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-packetloss"], "SystemSdwanHealthCheckFortiguard-ThresholdAlertPacketloss"); ok {
			if err = d.Set("threshold_alert_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", flattenSystemSdwanHealthCheckFortiguardThresholdWarningJitter2edl(o["threshold-warning-jitter"], d, "threshold_warning_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-jitter"], "SystemSdwanHealthCheckFortiguard-ThresholdWarningJitter"); ok {
			if err = d.Set("threshold_warning_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", flattenSystemSdwanHealthCheckFortiguardThresholdWarningLatency2edl(o["threshold-warning-latency"], d, "threshold_warning_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-latency"], "SystemSdwanHealthCheckFortiguard-ThresholdWarningLatency"); ok {
			if err = d.Set("threshold_warning_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", flattenSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss2edl(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-packetloss"], "SystemSdwanHealthCheckFortiguard-ThresholdWarningPacketloss"); ok {
			if err = d.Set("threshold_warning_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenSystemSdwanHealthCheckFortiguardUpdateCascadeInterface2edl(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-cascade-interface"], "SystemSdwanHealthCheckFortiguard-UpdateCascadeInterface"); ok {
			if err = d.Set("update_cascade_interface", vv); err != nil {
				return fmt.Errorf("Error reading update_cascade_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenSystemSdwanHealthCheckFortiguardUpdateStaticRoute2edl(o["update-static-route"], d, "update_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-static-route"], "SystemSdwanHealthCheckFortiguard-UpdateStaticRoute"); ok {
			if err = d.Set("update_static_route", vv); err != nil {
				return fmt.Errorf("Error reading update_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemSdwanHealthCheckFortiguardUser2edl(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemSdwanHealthCheckFortiguard-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSystemSdwanHealthCheckFortiguardVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "SystemSdwanHealthCheckFortiguard-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanHealthCheckFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanHealthCheckFortiguardAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardDetectMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDiffservcode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDnsMatchIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDnsRequestDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFailtime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFtpFile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFtpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHaPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpAgent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpGet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardMosCodec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardPacketSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbeCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbePackets2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbeTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardQualityMeasuredMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardRecoverytime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSecurityMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSdwanHealthCheckFortiguardSlaId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold2edl(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold2edl(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor2edl(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaMosThreshold2edl(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold2edl(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla2edl(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla2edl(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardSlaMosThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaFailLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaIdRedistribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPassLogPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSource62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSystemDns2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardTargetName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertJitter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertLatency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningJitter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningLatency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUpdateCascadeInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUpdateStaticRoute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanHealthCheckFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandSystemSdwanHealthCheckFortiguardAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandSystemSdwanHealthCheckFortiguardClassId2edl(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("detect_mode"); ok || d.HasChange("detect_mode") {
		t, err := expandSystemSdwanHealthCheckFortiguardDetectMode2edl(d, v, "detect_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-mode"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandSystemSdwanHealthCheckFortiguardDiffservcode2edl(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dns_match_ip"); ok || d.HasChange("dns_match_ip") {
		t, err := expandSystemSdwanHealthCheckFortiguardDnsMatchIp2edl(d, v, "dns_match_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok || d.HasChange("dns_request_domain") {
		t, err := expandSystemSdwanHealthCheckFortiguardDnsRequestDomain2edl(d, v, "dns_request_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("embed_measured_health"); ok || d.HasChange("embed_measured_health") {
		t, err := expandSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth2edl(d, v, "embed_measured_health")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embed-measured-health"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok || d.HasChange("failtime") {
		t, err := expandSystemSdwanHealthCheckFortiguardFailtime2edl(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("ftp_file"); ok || d.HasChange("ftp_file") {
		t, err := expandSystemSdwanHealthCheckFortiguardFtpFile2edl(d, v, "ftp_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-file"] = t
		}
	}

	if v, ok := d.GetOk("ftp_mode"); ok || d.HasChange("ftp_mode") {
		t, err := expandSystemSdwanHealthCheckFortiguardFtpMode2edl(d, v, "ftp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-mode"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandSystemSdwanHealthCheckFortiguardHaPriority2edl(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok || d.HasChange("http_agent") {
		t, err := expandSystemSdwanHealthCheckFortiguardHttpAgent2edl(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandSystemSdwanHealthCheckFortiguardHttpGet2edl(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandSystemSdwanHealthCheckFortiguardHttpMatch2edl(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandSystemSdwanHealthCheckFortiguardInterval2edl(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandSystemSdwanHealthCheckFortiguardMembers2edl(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("mos_codec"); ok || d.HasChange("mos_codec") {
		t, err := expandSystemSdwanHealthCheckFortiguardMosCodec2edl(d, v, "mos_codec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-codec"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok || d.HasChange("packet_size") {
		t, err := expandSystemSdwanHealthCheckFortiguardPacketSize2edl(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemSdwanHealthCheckFortiguardPassword2edl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSdwanHealthCheckFortiguardPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_count"); ok || d.HasChange("probe_count") {
		t, err := expandSystemSdwanHealthCheckFortiguardProbeCount2edl(d, v, "probe_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-count"] = t
		}
	}

	if v, ok := d.GetOk("probe_packets"); ok || d.HasChange("probe_packets") {
		t, err := expandSystemSdwanHealthCheckFortiguardProbePackets2edl(d, v, "probe_packets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-packets"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok || d.HasChange("probe_timeout") {
		t, err := expandSystemSdwanHealthCheckFortiguardProbeTimeout2edl(d, v, "probe_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemSdwanHealthCheckFortiguardProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_measured_method"); ok || d.HasChange("quality_measured_method") {
		t, err := expandSystemSdwanHealthCheckFortiguardQualityMeasuredMethod2edl(d, v, "quality_measured_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-measured-method"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok || d.HasChange("recoverytime") {
		t, err := expandSystemSdwanHealthCheckFortiguardRecoverytime2edl(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandSystemSdwanHealthCheckFortiguardSecurityMode2edl(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemSdwanHealthCheckFortiguardServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandSystemSdwanHealthCheckFortiguardSla2edl(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_fail_log_period"); ok || d.HasChange("sla_fail_log_period") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaFailLogPeriod2edl(d, v, "sla_fail_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-fail-log-period"] = t
		}
	}

	if v, ok := d.GetOk("sla_id_redistribute"); ok || d.HasChange("sla_id_redistribute") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaIdRedistribute2edl(d, v, "sla_id_redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id-redistribute"] = t
		}
	}

	if v, ok := d.GetOk("sla_pass_log_period"); ok || d.HasChange("sla_pass_log_period") {
		t, err := expandSystemSdwanHealthCheckFortiguardSlaPassLogPeriod2edl(d, v, "sla_pass_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-pass-log-period"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandSystemSdwanHealthCheckFortiguardSource2edl(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandSystemSdwanHealthCheckFortiguardSource62edl(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("system_dns"); ok || d.HasChange("system_dns") {
		t, err := expandSystemSdwanHealthCheckFortiguardSystemDns2edl(d, v, "system_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-dns"] = t
		}
	}

	if v, ok := d.GetOk("target_name"); ok || d.HasChange("target_name") {
		t, err := expandSystemSdwanHealthCheckFortiguardTargetName2edl(d, v, "target_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-name"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_jitter"); ok || d.HasChange("threshold_alert_jitter") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdAlertJitter2edl(d, v, "threshold_alert_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_latency"); ok || d.HasChange("threshold_alert_latency") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdAlertLatency2edl(d, v, "threshold_alert_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_packetloss"); ok || d.HasChange("threshold_alert_packetloss") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss2edl(d, v, "threshold_alert_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_jitter"); ok || d.HasChange("threshold_warning_jitter") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdWarningJitter2edl(d, v, "threshold_warning_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_latency"); ok || d.HasChange("threshold_warning_latency") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdWarningLatency2edl(d, v, "threshold_warning_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_packetloss"); ok || d.HasChange("threshold_warning_packetloss") {
		t, err := expandSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss2edl(d, v, "threshold_warning_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok || d.HasChange("update_cascade_interface") {
		t, err := expandSystemSdwanHealthCheckFortiguardUpdateCascadeInterface2edl(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok || d.HasChange("update_static_route") {
		t, err := expandSystemSdwanHealthCheckFortiguardUpdateStaticRoute2edl(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemSdwanHealthCheckFortiguardUser2edl(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandSystemSdwanHealthCheckFortiguardVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
