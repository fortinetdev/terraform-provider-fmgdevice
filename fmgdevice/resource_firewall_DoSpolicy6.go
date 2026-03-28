// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 DoS policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDosPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDosPolicy6Create,
		Read:   resourceFirewallDosPolicy6Read,
		Update: resourceFirewallDosPolicy6Update,
		Delete: resourceFirewallDosPolicy6Delete,

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

func resourceFirewallDosPolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallDosPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDosPolicy6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("policyid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallDosPolicy6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallDosPolicy6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallDosPolicy6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallDosPolicy6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallDosPolicy6 resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallDosPolicy6Read(d, m)
}

func resourceFirewallDosPolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallDosPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallDosPolicy6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourceFirewallDosPolicy6Read(d, m)
}

func resourceFirewallDosPolicy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallDosPolicy6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDosPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDosPolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallDosPolicy6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallDosPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDosPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDosPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDosPolicy6Anomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallDosPolicy6AnomalyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenFirewallDosPolicy6AnomalyLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallDosPolicy6AnomalyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenFirewallDosPolicy6AnomalyQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenFirewallDosPolicy6AnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenFirewallDosPolicy6AnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenFirewallDosPolicy6AnomalyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := i["synproxy-tcp-mss"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTcpMss(i["synproxy-tcp-mss"], d, pre_append)
			tmp["synproxy_tcp_mss"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTcpMss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := i["synproxy-tcp-sack"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTcpSack(i["synproxy-tcp-sack"], d, pre_append)
			tmp["synproxy_tcp_sack"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTcpSack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := i["synproxy-tcp-timestamp"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTcpTimestamp(i["synproxy-tcp-timestamp"], d, pre_append)
			tmp["synproxy_tcp_timestamp"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTcpTimestamp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := i["synproxy-tcp-window"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTcpWindow(i["synproxy-tcp-window"], d, pre_append)
			tmp["synproxy_tcp_window"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTcpWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := i["synproxy-tcp-windowscale"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTcpWindowscale(i["synproxy-tcp-windowscale"], d, pre_append)
			tmp["synproxy_tcp_windowscale"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTcpWindowscale")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := i["synproxy-tos"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTos(i["synproxy-tos"], d, pre_append)
			tmp["synproxy_tos"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := i["synproxy-ttl"]; ok {
			v := flattenFirewallDosPolicy6AnomalySynproxyTtl(i["synproxy-ttl"], d, pre_append)
			tmp["synproxy_ttl"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-SynproxyTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenFirewallDosPolicy6AnomalyThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-Threshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			v := flattenFirewallDosPolicy6AnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
			tmp["thresholddefault"] = fortiAPISubPartPatch(v, "FirewallDosPolicy6-Anomaly-ThresholdDefault")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallDosPolicy6AnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpSack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTcpWindowscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalySynproxyTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6AnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicy6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDosPolicy6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallDosPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallDosPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("anomaly", flattenFirewallDosPolicy6Anomaly(o["anomaly"], d, "anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallDosPolicy6-Anomaly"); ok {
				if err = d.Set("anomaly", vv); err != nil {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenFirewallDosPolicy6Anomaly(o["anomaly"], d, "anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["anomaly"], "FirewallDosPolicy6-Anomaly"); ok {
					if err = d.Set("anomaly", vv); err != nil {
						return fmt.Errorf("Error reading anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenFirewallDosPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallDosPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallDosPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallDosPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallDosPolicy6Interface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallDosPolicy6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallDosPolicy6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallDosPolicy6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallDosPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "FirewallDosPolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallDosPolicy6Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallDosPolicy6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallDosPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallDosPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallDosPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallDosPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFirewallDosPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDosPolicy6Anomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandFirewallDosPolicy6AnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandFirewallDosPolicy6AnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallDosPolicy6AnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandFirewallDosPolicy6AnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandFirewallDosPolicy6AnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandFirewallDosPolicy6AnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandFirewallDosPolicy6AnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-mss"], _ = expandFirewallDosPolicy6AnomalySynproxyTcpMss(d, i["synproxy_tcp_mss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-sack"], _ = expandFirewallDosPolicy6AnomalySynproxyTcpSack(d, i["synproxy_tcp_sack"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-timestamp"], _ = expandFirewallDosPolicy6AnomalySynproxyTcpTimestamp(d, i["synproxy_tcp_timestamp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-window"], _ = expandFirewallDosPolicy6AnomalySynproxyTcpWindow(d, i["synproxy_tcp_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-windowscale"], _ = expandFirewallDosPolicy6AnomalySynproxyTcpWindowscale(d, i["synproxy_tcp_windowscale"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tos"], _ = expandFirewallDosPolicy6AnomalySynproxyTos(d, i["synproxy_tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-ttl"], _ = expandFirewallDosPolicy6AnomalySynproxyTtl(d, i["synproxy_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold"], _ = expandFirewallDosPolicy6AnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold(default)"], _ = expandFirewallDosPolicy6AnomalyThresholdDefault(d, i["thresholddefault"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicy6AnomalyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpSack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTcpWindowscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalySynproxyTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6AnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicy6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicy6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicy6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallDosPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDosPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandFirewallDosPolicy6Anomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallDosPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallDosPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallDosPolicy6Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallDosPolicy6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandFirewallDosPolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallDosPolicy6Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallDosPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallDosPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
