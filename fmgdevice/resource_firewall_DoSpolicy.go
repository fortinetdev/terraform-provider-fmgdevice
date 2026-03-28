// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 DoS policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDosPolicyCreate,
		Read:   resourceFirewallDosPolicyRead,
		Update: resourceFirewallDosPolicyUpdate,
		Delete: resourceFirewallDosPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceFirewallDosPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectFirewallDosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDosPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallDosPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallDosPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallDosPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallDosPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallDosPolicy resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallDosPolicyRead(d, m)
}

func resourceFirewallDosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectFirewallDosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallDosPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallDosPolicyRead(d, m)
}

func resourceFirewallDosPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	wsParams["adom"] = adomv

	err = c.DeleteFirewallDosPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDosPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallDosPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallDosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDosPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDosPolicyAnomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallDosPolicyAnomalyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenFirewallDosPolicyAnomalyLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallDosPolicyAnomalyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenFirewallDosPolicyAnomalyQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenFirewallDosPolicyAnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenFirewallDosPolicyAnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallDosPolicyAnomalyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := i["synproxy-tcp-mss"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTcpMss(i["synproxy-tcp-mss"], d, pre_append)
			tmp["synproxy_tcp_mss"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTcpMss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := i["synproxy-tcp-sack"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTcpSack(i["synproxy-tcp-sack"], d, pre_append)
			tmp["synproxy_tcp_sack"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTcpSack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := i["synproxy-tcp-timestamp"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTcpTimestamp(i["synproxy-tcp-timestamp"], d, pre_append)
			tmp["synproxy_tcp_timestamp"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTcpTimestamp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := i["synproxy-tcp-window"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTcpWindow(i["synproxy-tcp-window"], d, pre_append)
			tmp["synproxy_tcp_window"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTcpWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := i["synproxy-tcp-windowscale"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTcpWindowscale(i["synproxy-tcp-windowscale"], d, pre_append)
			tmp["synproxy_tcp_windowscale"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTcpWindowscale")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := i["synproxy-tos"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTos(i["synproxy-tos"], d, pre_append)
			tmp["synproxy_tos"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := i["synproxy-ttl"]; ok {
			v := flattenFirewallDosPolicyAnomalySynproxyTtl(i["synproxy-ttl"], d, pre_append)
			tmp["synproxy_ttl"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-SynproxyTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenFirewallDosPolicyAnomalyThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-Threshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			v := flattenFirewallDosPolicyAnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
			tmp["thresholddefault"] = fortiAPISubPartPatch(v, "FirewallDosPolicy-Anomaly-ThresholdDefault")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallDosPolicyAnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTcpSack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTcpTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTcpWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTcpWindowscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalySynproxyTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallDosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("anomaly", flattenFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallDosPolicy-Anomaly"); ok {
				if err = d.Set("anomaly", vv); err != nil {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallDosPolicy-Anomaly"); ok {
					if err = d.Set("anomaly", vv); err != nil {
						return fmt.Errorf("Error reading anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenFirewallDosPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallDosPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallDosPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallDosPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallDosPolicyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallDosPolicy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallDosPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallDosPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallDosPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallDosPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallDosPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallDosPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallDosPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallDosPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallDosPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallDosPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFirewallDosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDosPolicyAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandFirewallDosPolicyAnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandFirewallDosPolicyAnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallDosPolicyAnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandFirewallDosPolicyAnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandFirewallDosPolicyAnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandFirewallDosPolicyAnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallDosPolicyAnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-mss"], _ = expandFirewallDosPolicyAnomalySynproxyTcpMss(d, i["synproxy_tcp_mss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-sack"], _ = expandFirewallDosPolicyAnomalySynproxyTcpSack(d, i["synproxy_tcp_sack"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-timestamp"], _ = expandFirewallDosPolicyAnomalySynproxyTcpTimestamp(d, i["synproxy_tcp_timestamp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-window"], _ = expandFirewallDosPolicyAnomalySynproxyTcpWindow(d, i["synproxy_tcp_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-windowscale"], _ = expandFirewallDosPolicyAnomalySynproxyTcpWindowscale(d, i["synproxy_tcp_windowscale"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tos"], _ = expandFirewallDosPolicyAnomalySynproxyTos(d, i["synproxy_tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-ttl"], _ = expandFirewallDosPolicyAnomalySynproxyTtl(d, i["synproxy_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold"], _ = expandFirewallDosPolicyAnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold(default)"], _ = expandFirewallDosPolicyAnomalyThresholdDefault(d, i["thresholddefault"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicyAnomalyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTcpSack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTcpTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTcpWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTcpWindowscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalySynproxyTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandFirewallDosPolicyAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallDosPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallDosPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallDosPolicyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallDosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallDosPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallDosPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallDosPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallDosPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
