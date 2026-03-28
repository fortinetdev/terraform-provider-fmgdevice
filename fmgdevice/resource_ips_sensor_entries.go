// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> IPS sensor filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsSensorEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSensorEntriesCreate,
		Read:   resourceIpsSensorEntriesRead,
		Update: resourceIpsSensorEntriesUpdate,
		Delete: resourceIpsSensorEntriesDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceIpsSensorEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	obj, err := getObjectIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsSensorEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIpsSensorEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIpsSensorEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IpsSensorEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIpsSensorEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IpsSensorEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIpsSensorEntriesRead(d, m)
}

func resourceIpsSensorEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	obj, err := getObjectIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensorEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIpsSensorEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensorEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceIpsSensorEntriesRead(d, m)
}

func resourceIpsSensorEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	sensor := d.Get("sensor").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	wsParams["adom"] = adomv

	err = c.DeleteIpsSensorEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSensorEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSensorEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	sensor := d.Get("sensor").(string)
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
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["sensor"] = sensor

	o, err := c.ReadIpsSensorEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IpsSensorEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSensorEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsSensorEntries resource from API: %v", err)
	}
	return nil
}

func flattenIpsSensorEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesCve2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesDefaultStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenIpsSensorEntriesExemptIpDstIp2edl(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenIpsSensorEntriesExemptIpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenIpsSensorEntriesExemptIpSrcIp2edl(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "IpsSensorEntries-ExemptIp-SrcIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesExemptIpDstIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesExemptIpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIpSrcIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLastModified2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesLocation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogAttackContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesOs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesRule2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIpsSensorEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSensorEntriesVulnType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectIpsSensorEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenIpsSensorEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "IpsSensorEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenIpsSensorEntriesApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "IpsSensorEntries-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("cve", flattenIpsSensorEntriesCve2edl(o["cve"], d, "cve")); err != nil {
		if vv, ok := fortiAPIPatch(o["cve"], "IpsSensorEntries-Cve"); ok {
			if err = d.Set("cve", vv); err != nil {
				return fmt.Errorf("Error reading cve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cve: %v", err)
		}
	}

	if err = d.Set("default_action", flattenIpsSensorEntriesDefaultAction2edl(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "IpsSensorEntries-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if err = d.Set("default_status", flattenIpsSensorEntriesDefaultStatus2edl(o["default-status"], d, "default_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-status"], "IpsSensorEntries-DefaultStatus"); ok {
			if err = d.Set("default_status", vv); err != nil {
				return fmt.Errorf("Error reading default_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exempt_ip", flattenIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["exempt-ip"], "IpsSensorEntries-ExemptIp"); ok {
				if err = d.Set("exempt_ip", vv); err != nil {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exempt_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exempt_ip"); ok {
			if err = d.Set("exempt_ip", flattenIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["exempt-ip"], "IpsSensorEntries-ExemptIp"); ok {
					if err = d.Set("exempt_ip", vv); err != nil {
						return fmt.Errorf("Error reading exempt_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenIpsSensorEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "IpsSensorEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("last_modified", flattenIpsSensorEntriesLastModified2edl(o["last-modified"], d, "last_modified")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-modified"], "IpsSensorEntries-LastModified"); ok {
			if err = d.Set("last_modified", vv); err != nil {
				return fmt.Errorf("Error reading last_modified: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_modified: %v", err)
		}
	}

	if err = d.Set("location", flattenIpsSensorEntriesLocation2edl(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "IpsSensorEntries-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("log", flattenIpsSensorEntriesLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "IpsSensorEntries-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_attack_context", flattenIpsSensorEntriesLogAttackContext2edl(o["log-attack-context"], d, "log_attack_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-attack-context"], "IpsSensorEntries-LogAttackContext"); ok {
			if err = d.Set("log_attack_context", vv); err != nil {
				return fmt.Errorf("Error reading log_attack_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_attack_context: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenIpsSensorEntriesLogPacket2edl(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "IpsSensorEntries-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("os", flattenIpsSensorEntriesOs2edl(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "IpsSensorEntries-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("protocol", flattenIpsSensorEntriesProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "IpsSensorEntries-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenIpsSensorEntriesQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "IpsSensorEntries-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenIpsSensorEntriesQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "IpsSensorEntries-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenIpsSensorEntriesQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "IpsSensorEntries-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("rate_count", flattenIpsSensorEntriesRateCount2edl(o["rate-count"], d, "rate_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-count"], "IpsSensorEntries-RateCount"); ok {
			if err = d.Set("rate_count", vv); err != nil {
				return fmt.Errorf("Error reading rate_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_count: %v", err)
		}
	}

	if err = d.Set("rate_duration", flattenIpsSensorEntriesRateDuration2edl(o["rate-duration"], d, "rate_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-duration"], "IpsSensorEntries-RateDuration"); ok {
			if err = d.Set("rate_duration", vv); err != nil {
				return fmt.Errorf("Error reading rate_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_duration: %v", err)
		}
	}

	if err = d.Set("rate_mode", flattenIpsSensorEntriesRateMode2edl(o["rate-mode"], d, "rate_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-mode"], "IpsSensorEntries-RateMode"); ok {
			if err = d.Set("rate_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_mode: %v", err)
		}
	}

	if err = d.Set("rate_track", flattenIpsSensorEntriesRateTrack2edl(o["rate-track"], d, "rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-track"], "IpsSensorEntries-RateTrack"); ok {
			if err = d.Set("rate_track", vv); err != nil {
				return fmt.Errorf("Error reading rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_track: %v", err)
		}
	}

	if err = d.Set("rule", flattenIpsSensorEntriesRule2edl(o["rule"], d, "rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule"], "IpsSensorEntries-Rule"); ok {
			if err = d.Set("rule", vv); err != nil {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	if err = d.Set("severity", flattenIpsSensorEntriesSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "IpsSensorEntries-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenIpsSensorEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "IpsSensorEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vuln_type", flattenIpsSensorEntriesVulnType2edl(o["vuln-type"], d, "vuln_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vuln-type"], "IpsSensorEntries-VulnType"); ok {
			if err = d.Set("vuln_type", vv); err != nil {
				return fmt.Errorf("Error reading vuln_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vuln_type: %v", err)
		}
	}

	return nil
}

func flattenIpsSensorEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsSensorEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesCve2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesDefaultStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dst-ip"], _ = expandIpsSensorEntriesExemptIpDstIp2edl(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandIpsSensorEntriesExemptIpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandIpsSensorEntriesExemptIpSrcIp2edl(d, i["src_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesExemptIpDstIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandIpsSensorEntriesExemptIpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIpSrcIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandIpsSensorEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLastModified2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesLocation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogAttackContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesOs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIpsSensorEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesVulnType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectIpsSensorEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandIpsSensorEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandIpsSensorEntriesApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("cve"); ok || d.HasChange("cve") {
		t, err := expandIpsSensorEntriesCve2edl(d, v, "cve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cve"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandIpsSensorEntriesDefaultAction2edl(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("default_status"); ok || d.HasChange("default_status") {
		t, err := expandIpsSensorEntriesDefaultStatus2edl(d, v, "default_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-status"] = t
		}
	}

	if v, ok := d.GetOk("exempt_ip"); ok || d.HasChange("exempt_ip") {
		t, err := expandIpsSensorEntriesExemptIp2edl(d, v, "exempt_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandIpsSensorEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("last_modified"); ok || d.HasChange("last_modified") {
		t, err := expandIpsSensorEntriesLastModified2edl(d, v, "last_modified")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-modified"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandIpsSensorEntriesLocation2edl(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandIpsSensorEntriesLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_attack_context"); ok || d.HasChange("log_attack_context") {
		t, err := expandIpsSensorEntriesLogAttackContext2edl(d, v, "log_attack_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-attack-context"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandIpsSensorEntriesLogPacket2edl(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandIpsSensorEntriesOs2edl(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandIpsSensorEntriesProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandIpsSensorEntriesQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandIpsSensorEntriesQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandIpsSensorEntriesQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_count"); ok || d.HasChange("rate_count") {
		t, err := expandIpsSensorEntriesRateCount2edl(d, v, "rate_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-count"] = t
		}
	}

	if v, ok := d.GetOk("rate_duration"); ok || d.HasChange("rate_duration") {
		t, err := expandIpsSensorEntriesRateDuration2edl(d, v, "rate_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-duration"] = t
		}
	}

	if v, ok := d.GetOk("rate_mode"); ok || d.HasChange("rate_mode") {
		t, err := expandIpsSensorEntriesRateMode2edl(d, v, "rate_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_track"); ok || d.HasChange("rate_track") {
		t, err := expandIpsSensorEntriesRateTrack2edl(d, v, "rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-track"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandIpsSensorEntriesRule2edl(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandIpsSensorEntriesSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandIpsSensorEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vuln_type"); ok || d.HasChange("vuln_type") {
		t, err := expandIpsSensorEntriesVulnType2edl(d, v, "vuln_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vuln-type"] = t
		}
	}

	return &obj, nil
}
