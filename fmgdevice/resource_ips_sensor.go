// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPS sensor.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSensorCreate,
		Read:   resourceIpsSensorRead,
		Update: resourceIpsSensorUpdate,
		Delete: resourceIpsSensorDelete,

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
			"block_malicious_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"cve": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"last_modified": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_attack_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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
						"rate_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rate_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
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
						"vuln_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
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

func resourceIpsSensorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsSensor resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIpsSensor(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIpsSensor(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IpsSensor resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIpsSensor(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IpsSensor resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceIpsSensorRead(d, m)
}

func resourceIpsSensorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensor resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIpsSensor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceIpsSensorRead(d, m)
}

func resourceIpsSensorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIpsSensor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSensorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIpsSensor(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IpsSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsSensor resource from API: %v", err)
	}
	return nil
}

func flattenIpsSensorBlockMaliciousUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenIpsSensorEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenIpsSensorEntriesApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := i["cve"]; ok {
			v := flattenIpsSensorEntriesCve(i["cve"], d, pre_append)
			tmp["cve"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Cve")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := i["default-action"]; ok {
			v := flattenIpsSensorEntriesDefaultAction(i["default-action"], d, pre_append)
			tmp["default_action"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-DefaultAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := i["default-status"]; ok {
			v := flattenIpsSensorEntriesDefaultStatus(i["default-status"], d, pre_append)
			tmp["default_status"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-DefaultStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {
			v := flattenIpsSensorEntriesExemptIp(i["exempt-ip"], d, pre_append)
			tmp["exempt_ip"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-ExemptIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIpsSensorEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := i["last-modified"]; ok {
			v := flattenIpsSensorEntriesLastModified(i["last-modified"], d, pre_append)
			tmp["last_modified"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-LastModified")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {
			v := flattenIpsSensorEntriesLocation(i["location"], d, pre_append)
			tmp["location"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Location")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenIpsSensorEntriesLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := i["log-attack-context"]; ok {
			v := flattenIpsSensorEntriesLogAttackContext(i["log-attack-context"], d, pre_append)
			tmp["log_attack_context"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-LogAttackContext")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenIpsSensorEntriesLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {
			v := flattenIpsSensorEntriesOs(i["os"], d, pre_append)
			tmp["os"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Os")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenIpsSensorEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenIpsSensorEntriesQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenIpsSensorEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenIpsSensorEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			v := flattenIpsSensorEntriesRateCount(i["rate-count"], d, pre_append)
			tmp["rate_count"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-RateCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			v := flattenIpsSensorEntriesRateDuration(i["rate-duration"], d, pre_append)
			tmp["rate_duration"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-RateDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			v := flattenIpsSensorEntriesRateMode(i["rate-mode"], d, pre_append)
			tmp["rate_mode"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-RateMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			v := flattenIpsSensorEntriesRateTrack(i["rate-track"], d, pre_append)
			tmp["rate_track"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-RateTrack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {
			v := flattenIpsSensorEntriesRule(i["rule"], d, pre_append)
			tmp["rule"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Rule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenIpsSensorEntriesSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenIpsSensorEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := i["vuln-type"]; ok {
			v := flattenIpsSensorEntriesVulnType(i["vuln-type"], d, pre_append)
			tmp["vuln_type"] = fortiAPISubPartPatch(v, "IpsSensor-Entries-VulnType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesCve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesDefaultStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := i["dst-ip"]; ok {
			v := flattenIpsSensorEntriesExemptIpDstIp(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIpsSensorEntriesExemptIpId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenIpsSensorEntriesExemptIpSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-SrcIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesExemptIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLastModified(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogAttackContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesVulnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenIpsSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("block_malicious_url", flattenIpsSensorBlockMaliciousUrl(o["block-malicious-url"], d, "block_malicious_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-malicious-url"], "IpsSensor-BlockMaliciousUrl"); ok {
			if err = d.Set("block_malicious_url", vv); err != nil {
				return fmt.Errorf("Error reading block_malicious_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_malicious_url: %v", err)
		}
	}

	if err = d.Set("comment", flattenIpsSensorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "IpsSensor-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenIpsSensorEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "IpsSensor-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenIpsSensorEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "IpsSensor-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenIpsSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "IpsSensor-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("name", flattenIpsSensorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "IpsSensor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenIpsSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "IpsSensor-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenIpsSensorScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "IpsSensor-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	return nil
}

func flattenIpsSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsSensorBlockMaliciousUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandIpsSensorEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandIpsSensorEntriesApplication(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cve"], _ = expandIpsSensorEntriesCve(d, i["cve"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-action"], _ = expandIpsSensorEntriesDefaultAction(d, i["default_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-status"], _ = expandIpsSensorEntriesDefaultStatus(d, i["default_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandIpsSensorEntriesExemptIp(d, i["exempt_ip"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["exempt-ip"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIpsSensorEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["last-modified"], _ = expandIpsSensorEntriesLastModified(d, i["last_modified"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["location"], _ = expandIpsSensorEntriesLocation(d, i["location"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandIpsSensorEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-attack-context"], _ = expandIpsSensorEntriesLogAttackContext(d, i["log_attack_context"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandIpsSensorEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os"], _ = expandIpsSensorEntriesOs(d, i["os"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandIpsSensorEntriesProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandIpsSensorEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandIpsSensorEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandIpsSensorEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-count"], _ = expandIpsSensorEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-duration"], _ = expandIpsSensorEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-mode"], _ = expandIpsSensorEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-track"], _ = expandIpsSensorEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule"], _ = expandIpsSensorEntriesRule(d, i["rule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandIpsSensorEntriesSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandIpsSensorEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vuln-type"], _ = expandIpsSensorEntriesVulnType(d, i["vuln_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesCve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesDefaultStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-ip"], _ = expandIpsSensorEntriesExemptIpDstIp(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIpsSensorEntriesExemptIpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandIpsSensorEntriesExemptIpSrcIp(d, i["src_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandIpsSensorEntriesExemptIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandIpsSensorEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLastModified(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogAttackContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesVulnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandIpsSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_malicious_url"); ok || d.HasChange("block_malicious_url") {
		t, err := expandIpsSensorBlockMaliciousUrl(d, v, "block_malicious_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malicious-url"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandIpsSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandIpsSensorEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandIpsSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandIpsSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandIpsSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandIpsSensorScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	return &obj, nil
}
