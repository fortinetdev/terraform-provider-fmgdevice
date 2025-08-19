// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure sniffer.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSnifferCreate,
		Read:   resourceFirewallSnifferRead,
		Update: resourceFirewallSnifferUpdate,
		Delete: resourceFirewallSnifferDelete,

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
			"anomaly": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tcp_mss": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tcp_sack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tcp_timestamp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tcp_window": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tcp_windowscale": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"synproxy_ttl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"thresholddefault": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application_list_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dlp_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_filter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_threatfeed": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_threatfeed_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_dos_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_packet_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"non_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceFirewallSnifferCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallSniffer(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSniffer resource while getting object: %v", err)
	}

	v, err := c.CreateFirewallSniffer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSniffer resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceFirewallSnifferRead(d, m)
		} else {
			return fmt.Errorf("Error creating FirewallSniffer resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSnifferRead(d, m)
}

func resourceFirewallSnifferUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallSniffer(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSniffer resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallSniffer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallSnifferRead(d, m)
}

func resourceFirewallSnifferDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSniffer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSnifferRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSniffer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSniffer resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSnifferAnomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenFirewallSnifferAnomalyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenFirewallSnifferAnomalyLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallSnifferAnomalyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenFirewallSnifferAnomalyQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenFirewallSnifferAnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenFirewallSnifferAnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallSnifferAnomalyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := i["synproxy-tcp-mss"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTcpMss(i["synproxy-tcp-mss"], d, pre_append)
			tmp["synproxy_tcp_mss"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTcpMss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := i["synproxy-tcp-sack"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTcpSack(i["synproxy-tcp-sack"], d, pre_append)
			tmp["synproxy_tcp_sack"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTcpSack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := i["synproxy-tcp-timestamp"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTcpTimestamp(i["synproxy-tcp-timestamp"], d, pre_append)
			tmp["synproxy_tcp_timestamp"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTcpTimestamp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := i["synproxy-tcp-window"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTcpWindow(i["synproxy-tcp-window"], d, pre_append)
			tmp["synproxy_tcp_window"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTcpWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := i["synproxy-tcp-windowscale"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTcpWindowscale(i["synproxy-tcp-windowscale"], d, pre_append)
			tmp["synproxy_tcp_windowscale"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTcpWindowscale")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := i["synproxy-tos"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTos(i["synproxy-tos"], d, pre_append)
			tmp["synproxy_tos"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := i["synproxy-ttl"]; ok {
			v := flattenFirewallSnifferAnomalySynproxyTtl(i["synproxy-ttl"], d, pre_append)
			tmp["synproxy_ttl"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-SynproxyTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenFirewallSnifferAnomalyThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-Threshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			v := flattenFirewallSnifferAnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
			tmp["thresholddefault"] = fortiAPISubPartPatch(v, "FirewallSniffer-Anomaly-ThresholdDefault")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallSnifferAnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTcpSack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTcpTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTcpWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTcpWindowscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalySynproxyTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferApplicationListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferAvProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferDlpSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferDlpProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferEmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferFileFilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferIpThreatfeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferIpThreatfeedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferIpsDosStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferIpsSensorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferMaxPacketCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferNonIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSnifferWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSnifferWebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSniffer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("anomaly", flattenFirewallSnifferAnomaly(o["anomaly"], d, "anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallSniffer-Anomaly"); ok {
				if err = d.Set("anomaly", vv); err != nil {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenFirewallSnifferAnomaly(o["anomaly"], d, "anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallSniffer-Anomaly"); ok {
					if err = d.Set("anomaly", vv); err != nil {
						return fmt.Errorf("Error reading anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("application_list", flattenFirewallSnifferApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallSniffer-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("application_list_status", flattenFirewallSnifferApplicationListStatus(o["application-list-status"], d, "application_list_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list-status"], "FirewallSniffer-ApplicationListStatus"); ok {
			if err = d.Set("application_list_status", vv); err != nil {
				return fmt.Errorf("Error reading application_list_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallSnifferAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallSniffer-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenFirewallSnifferAvProfileStatus(o["av-profile-status"], d, "av_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile-status"], "FirewallSniffer-AvProfileStatus"); ok {
			if err = d.Set("av_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading av_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallSnifferDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallSniffer-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenFirewallSnifferDlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor-status"], "FirewallSniffer-DlpSensorStatus"); ok {
			if err = d.Set("dlp_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallSnifferDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallSniffer-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile_status", flattenFirewallSnifferDlpProfileStatus(o["dlp-profile-status"], d, "dlp_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile-status"], "FirewallSniffer-DlpProfileStatus"); ok {
			if err = d.Set("dlp_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile_status: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallSnifferDsri(o["dsri"], d, "dsri")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsri"], "FirewallSniffer-Dsri"); ok {
			if err = d.Set("dsri", vv); err != nil {
				return fmt.Errorf("Error reading dsri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallSnifferEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallSniffer-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenFirewallSnifferEmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile-status"], "FirewallSniffer-EmailfilterProfileStatus"); ok {
			if err = d.Set("emailfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallSnifferFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "FirewallSniffer-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("file_filter_profile_status", flattenFirewallSnifferFileFilterProfileStatus(o["file-filter-profile-status"], d, "file_filter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile-status"], "FirewallSniffer-FileFilterProfileStatus"); ok {
			if err = d.Set("file_filter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile_status: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallSnifferHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallSniffer-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallSnifferId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallSniffer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallSnifferInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallSniffer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_threatfeed", flattenFirewallSnifferIpThreatfeed(o["ip-threatfeed"], d, "ip_threatfeed")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-threatfeed"], "FirewallSniffer-IpThreatfeed"); ok {
			if err = d.Set("ip_threatfeed", vv); err != nil {
				return fmt.Errorf("Error reading ip_threatfeed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_threatfeed: %v", err)
		}
	}

	if err = d.Set("ip_threatfeed_status", flattenFirewallSnifferIpThreatfeedStatus(o["ip-threatfeed-status"], d, "ip_threatfeed_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-threatfeed-status"], "FirewallSniffer-IpThreatfeedStatus"); ok {
			if err = d.Set("ip_threatfeed_status", vv); err != nil {
				return fmt.Errorf("Error reading ip_threatfeed_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_threatfeed_status: %v", err)
		}
	}

	if err = d.Set("ips_dos_status", flattenFirewallSnifferIpsDosStatus(o["ips-dos-status"], d, "ips_dos_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-dos-status"], "FirewallSniffer-IpsDosStatus"); ok {
			if err = d.Set("ips_dos_status", vv); err != nil {
				return fmt.Errorf("Error reading ips_dos_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_dos_status: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallSnifferIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallSniffer-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenFirewallSnifferIpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor-status"], "FirewallSniffer-IpsSensorStatus"); ok {
			if err = d.Set("ips_sensor_status", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenFirewallSnifferIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "FirewallSniffer-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallSnifferLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "FirewallSniffer-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("max_packet_count", flattenFirewallSnifferMaxPacketCount(o["max-packet-count"], d, "max_packet_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-packet-count"], "FirewallSniffer-MaxPacketCount"); ok {
			if err = d.Set("max_packet_count", vv); err != nil {
				return fmt.Errorf("Error reading max_packet_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_packet_count: %v", err)
		}
	}

	if err = d.Set("non_ip", flattenFirewallSnifferNonIp(o["non-ip"], d, "non_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["non-ip"], "FirewallSniffer-NonIp"); ok {
			if err = d.Set("non_ip", vv); err != nil {
				return fmt.Errorf("Error reading non_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading non_ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallSnifferPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallSniffer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallSnifferProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "FirewallSniffer-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSnifferStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallSniffer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallSnifferUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallSniffer-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("vlan", flattenFirewallSnifferVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "FirewallSniffer-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallSnifferWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallSniffer-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenFirewallSnifferWebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile-status"], "FirewallSniffer-WebfilterProfileStatus"); ok {
			if err = d.Set("webfilter_profile_status", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	return nil
}

func flattenFirewallSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSnifferAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandFirewallSnifferAnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandFirewallSnifferAnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallSnifferAnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandFirewallSnifferAnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandFirewallSnifferAnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandFirewallSnifferAnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallSnifferAnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-mss"], _ = expandFirewallSnifferAnomalySynproxyTcpMss(d, i["synproxy_tcp_mss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-sack"], _ = expandFirewallSnifferAnomalySynproxyTcpSack(d, i["synproxy_tcp_sack"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-timestamp"], _ = expandFirewallSnifferAnomalySynproxyTcpTimestamp(d, i["synproxy_tcp_timestamp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-window"], _ = expandFirewallSnifferAnomalySynproxyTcpWindow(d, i["synproxy_tcp_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-windowscale"], _ = expandFirewallSnifferAnomalySynproxyTcpWindowscale(d, i["synproxy_tcp_windowscale"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tos"], _ = expandFirewallSnifferAnomalySynproxyTos(d, i["synproxy_tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-ttl"], _ = expandFirewallSnifferAnomalySynproxyTtl(d, i["synproxy_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold"], _ = expandFirewallSnifferAnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold(default)"], _ = expandFirewallSnifferAnomalyThresholdDefault(d, i["thresholddefault"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallSnifferAnomalyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTcpSack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTcpTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTcpWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTcpWindowscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalySynproxyTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferApplicationListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferAvProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferDlpSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferDlpProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferDsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferEmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferFileFilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferIpThreatfeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferIpThreatfeedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferIpsDosStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferIpsSensorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferMaxPacketCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferNonIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSnifferWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSnifferWebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSniffer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandFirewallSnifferAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallSnifferApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok || d.HasChange("application_list_status") {
		t, err := expandFirewallSnifferApplicationListStatus(d, v, "application_list_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallSnifferAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok || d.HasChange("av_profile_status") {
		t, err := expandFirewallSnifferAvProfileStatus(d, v, "av_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallSnifferDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok || d.HasChange("dlp_sensor_status") {
		t, err := expandFirewallSnifferDlpSensorStatus(d, v, "dlp_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallSnifferDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile_status"); ok || d.HasChange("dlp_profile_status") {
		t, err := expandFirewallSnifferDlpProfileStatus(d, v, "dlp_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok || d.HasChange("dsri") {
		t, err := expandFirewallSnifferDsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallSnifferEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok || d.HasChange("emailfilter_profile_status") {
		t, err := expandFirewallSnifferEmailfilterProfileStatus(d, v, "emailfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandFirewallSnifferFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile_status"); ok || d.HasChange("file_filter_profile_status") {
		t, err := expandFirewallSnifferFileFilterProfileStatus(d, v, "file_filter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallSnifferHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallSnifferId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallSnifferInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_threatfeed"); ok || d.HasChange("ip_threatfeed") {
		t, err := expandFirewallSnifferIpThreatfeed(d, v, "ip_threatfeed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-threatfeed"] = t
		}
	}

	if v, ok := d.GetOk("ip_threatfeed_status"); ok || d.HasChange("ip_threatfeed_status") {
		t, err := expandFirewallSnifferIpThreatfeedStatus(d, v, "ip_threatfeed_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-threatfeed-status"] = t
		}
	}

	if v, ok := d.GetOk("ips_dos_status"); ok || d.HasChange("ips_dos_status") {
		t, err := expandFirewallSnifferIpsDosStatus(d, v, "ips_dos_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-dos-status"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallSnifferIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok || d.HasChange("ips_sensor_status") {
		t, err := expandFirewallSnifferIpsSensorStatus(d, v, "ips_sensor_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandFirewallSnifferIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok || d.HasChange("logtraffic") {
		t, err := expandFirewallSnifferLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("max_packet_count"); ok || d.HasChange("max_packet_count") {
		t, err := expandFirewallSnifferMaxPacketCount(d, v, "max_packet_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-packet-count"] = t
		}
	}

	if v, ok := d.GetOk("non_ip"); ok || d.HasChange("non_ip") {
		t, err := expandFirewallSnifferNonIp(d, v, "non_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["non-ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallSnifferPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandFirewallSnifferProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallSnifferStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallSnifferUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandFirewallSnifferVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallSnifferWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok || d.HasChange("webfilter_profile_status") {
		t, err := expandFirewallSnifferWebfilterProfileStatus(d, v, "webfilter_profile_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	return &obj, nil
}
