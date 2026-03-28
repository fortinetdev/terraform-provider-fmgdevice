// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure threat weight settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogThreatWeight() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogThreatWeightUpdate,
		Read:   resourceLogThreatWeightRead,
		Update: resourceLogThreatWeightUpdate,
		Delete: resourceLogThreatWeightDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"blocked_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"botnet_connection_detected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failed_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geolocation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"high_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"info_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"low_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"medium_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"level": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"medium": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"malware": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_blocked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ems_threat_feed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"file_blocked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fsa_high_risk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsa_malicious": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsa_medium_risk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inline_block": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malware_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mimefragmented": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversized": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_file_type_executable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_infected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_scan_error": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_block_detected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLogThreatWeightUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogThreatWeight(d)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeight resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogThreatWeight(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeight resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogThreatWeight")

	return resourceLogThreatWeightRead(d, m)
}

func resourceLogThreatWeightDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogThreatWeight(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogThreatWeight resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogThreatWeight(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogThreatWeight resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogThreatWeight(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeight resource from API: %v", err)
	}
	return nil
}

func flattenLogThreatWeightApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenLogThreatWeightApplicationCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogThreatWeight-Application-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogThreatWeightApplicationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogThreatWeight-Application-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			v := flattenLogThreatWeightApplicationLevel(i["level"], d, pre_append)
			tmp["level"] = fortiAPISubPartPatch(v, "LogThreatWeight-Application-Level")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogThreatWeightApplicationCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogThreatWeightApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightApplicationLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightBlockedConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightBotnetConnectionDetected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightFailedConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {
			v := flattenLogThreatWeightGeolocationCountry(i["country"], d, pre_append)
			tmp["country"] = fortiAPISubPartPatch(v, "LogThreatWeight-Geolocation-Country")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogThreatWeightGeolocationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogThreatWeight-Geolocation-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			v := flattenLogThreatWeightGeolocationLevel(i["level"], d, pre_append)
			tmp["level"] = fortiAPISubPartPatch(v, "LogThreatWeight-Geolocation-Level")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogThreatWeightGeolocationCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogThreatWeightGeolocationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocationLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "critical_severity"
	if _, ok := i["critical-severity"]; ok {
		result["critical_severity"] = flattenLogThreatWeightIpsCriticalSeverity(i["critical-severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "high_severity"
	if _, ok := i["high-severity"]; ok {
		result["high_severity"] = flattenLogThreatWeightIpsHighSeverity(i["high-severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "info_severity"
	if _, ok := i["info-severity"]; ok {
		result["info_severity"] = flattenLogThreatWeightIpsInfoSeverity(i["info-severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_severity"
	if _, ok := i["low-severity"]; ok {
		result["low_severity"] = flattenLogThreatWeightIpsLowSeverity(i["low-severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "medium_severity"
	if _, ok := i["medium-severity"]; ok {
		result["medium_severity"] = flattenLogThreatWeightIpsMediumSeverity(i["medium-severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightIpsCriticalSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsHighSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsInfoSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsLowSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightIpsMediumSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "critical"
	if _, ok := i["critical"]; ok {
		result["critical"] = flattenLogThreatWeightLevelCritical(i["critical"], d, pre_append)
	}

	pre_append = pre + ".0." + "high"
	if _, ok := i["high"]; ok {
		result["high"] = flattenLogThreatWeightLevelHigh(i["high"], d, pre_append)
	}

	pre_append = pre + ".0." + "low"
	if _, ok := i["low"]; ok {
		result["low"] = flattenLogThreatWeightLevelLow(i["low"], d, pre_append)
	}

	pre_append = pre + ".0." + "medium"
	if _, ok := i["medium"]; ok {
		result["medium"] = flattenLogThreatWeightLevelMedium(i["medium"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightLevelCritical(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightLevelMedium(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalware(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "command_blocked"
	if _, ok := i["command-blocked"]; ok {
		result["command_blocked"] = flattenLogThreatWeightMalwareCommandBlocked(i["command-blocked"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenLogThreatWeightMalwareContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ems_threat_feed"
	if _, ok := i["ems-threat-feed"]; ok {
		result["ems_threat_feed"] = flattenLogThreatWeightMalwareEmsThreatFeed(i["ems-threat-feed"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_blocked"
	if _, ok := i["file-blocked"]; ok {
		result["file_blocked"] = flattenLogThreatWeightMalwareFileBlocked(i["file-blocked"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenLogThreatWeightMalwareFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fsa_high_risk"
	if _, ok := i["fsa-high-risk"]; ok {
		result["fsa_high_risk"] = flattenLogThreatWeightMalwareFsaHighRisk(i["fsa-high-risk"], d, pre_append)
	}

	pre_append = pre + ".0." + "fsa_malicious"
	if _, ok := i["fsa-malicious"]; ok {
		result["fsa_malicious"] = flattenLogThreatWeightMalwareFsaMalicious(i["fsa-malicious"], d, pre_append)
	}

	pre_append = pre + ".0." + "fsa_medium_risk"
	if _, ok := i["fsa-medium-risk"]; ok {
		result["fsa_medium_risk"] = flattenLogThreatWeightMalwareFsaMediumRisk(i["fsa-medium-risk"], d, pre_append)
	}

	pre_append = pre + ".0." + "inline_block"
	if _, ok := i["inline-block"]; ok {
		result["inline_block"] = flattenLogThreatWeightMalwareInlineBlock(i["inline-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_list"
	if _, ok := i["malware-list"]; ok {
		result["malware_list"] = flattenLogThreatWeightMalwareMalwareList(i["malware-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "mimefragmented"
	if _, ok := i["mimefragmented"]; ok {
		result["mimefragmented"] = flattenLogThreatWeightMalwareMimefragmented(i["mimefragmented"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversized"
	if _, ok := i["oversized"]; ok {
		result["oversized"] = flattenLogThreatWeightMalwareOversized(i["oversized"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_proto"
	if _, ok := i["switch-proto"]; ok {
		result["switch_proto"] = flattenLogThreatWeightMalwareSwitchProto(i["switch-proto"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_file_type_executable"
	if _, ok := i["virus-file-type-executable"]; ok {
		result["virus_file_type_executable"] = flattenLogThreatWeightMalwareVirusFileTypeExecutable(i["virus-file-type-executable"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_infected"
	if _, ok := i["virus-infected"]; ok {
		result["virus_infected"] = flattenLogThreatWeightMalwareVirusInfected(i["virus-infected"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_outbreak_prevention"
	if _, ok := i["virus-outbreak-prevention"]; ok {
		result["virus_outbreak_prevention"] = flattenLogThreatWeightMalwareVirusOutbreakPrevention(i["virus-outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_scan_error"
	if _, ok := i["virus-scan-error"]; ok {
		result["virus_scan_error"] = flattenLogThreatWeightMalwareVirusScanError(i["virus-scan-error"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenLogThreatWeightMalwareFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenLogThreatWeightMalwareFortisandbox(i["fortisandbox"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightMalwareCommandBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareEmsThreatFeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFileBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaHighRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaMalicious(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaMediumRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareInlineBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareMalwareList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareMimefragmented(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareOversized(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareSwitchProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusFileTypeExecutable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusScanError(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightUrlBlockDetected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightWeb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenLogThreatWeightWebCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "LogThreatWeight-Web-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenLogThreatWeightWebId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "LogThreatWeight-Web-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			v := flattenLogThreatWeightWebLevel(i["level"], d, pre_append)
			tmp["level"] = fortiAPISubPartPatch(v, "LogThreatWeight-Web-Level")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenLogThreatWeightWebCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLogThreatWeightWebId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogThreatWeightWebLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogThreatWeight(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("application", flattenLogThreatWeightApplication(o["application"], d, "application")); err != nil {
			if vv, ok := fortiAPIPatch(o["application"], "LogThreatWeight-Application"); ok {
				if err = d.Set("application", vv); err != nil {
					return fmt.Errorf("Error reading application: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenLogThreatWeightApplication(o["application"], d, "application")); err != nil {
				if vv, ok := fortiAPIPatch(o["application"], "LogThreatWeight-Application"); ok {
					if err = d.Set("application", vv); err != nil {
						return fmt.Errorf("Error reading application: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if err = d.Set("blocked_connection", flattenLogThreatWeightBlockedConnection(o["blocked-connection"], d, "blocked_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-connection"], "LogThreatWeight-BlockedConnection"); ok {
			if err = d.Set("blocked_connection", vv); err != nil {
				return fmt.Errorf("Error reading blocked_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_connection: %v", err)
		}
	}

	if err = d.Set("botnet_connection_detected", flattenLogThreatWeightBotnetConnectionDetected(o["botnet-connection-detected"], d, "botnet_connection_detected")); err != nil {
		if vv, ok := fortiAPIPatch(o["botnet-connection-detected"], "LogThreatWeight-BotnetConnectionDetected"); ok {
			if err = d.Set("botnet_connection_detected", vv); err != nil {
				return fmt.Errorf("Error reading botnet_connection_detected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading botnet_connection_detected: %v", err)
		}
	}

	if err = d.Set("failed_connection", flattenLogThreatWeightFailedConnection(o["failed-connection"], d, "failed_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["failed-connection"], "LogThreatWeight-FailedConnection"); ok {
			if err = d.Set("failed_connection", vv); err != nil {
				return fmt.Errorf("Error reading failed_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failed_connection: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("geolocation", flattenLogThreatWeightGeolocation(o["geolocation"], d, "geolocation")); err != nil {
			if vv, ok := fortiAPIPatch(o["geolocation"], "LogThreatWeight-Geolocation"); ok {
				if err = d.Set("geolocation", vv); err != nil {
					return fmt.Errorf("Error reading geolocation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading geolocation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("geolocation"); ok {
			if err = d.Set("geolocation", flattenLogThreatWeightGeolocation(o["geolocation"], d, "geolocation")); err != nil {
				if vv, ok := fortiAPIPatch(o["geolocation"], "LogThreatWeight-Geolocation"); ok {
					if err = d.Set("geolocation", vv); err != nil {
						return fmt.Errorf("Error reading geolocation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading geolocation: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ips", flattenLogThreatWeightIps(o["ips"], d, "ips")); err != nil {
			if vv, ok := fortiAPIPatch(o["ips"], "LogThreatWeight-Ips"); ok {
				if err = d.Set("ips", vv); err != nil {
					return fmt.Errorf("Error reading ips: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ips: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ips"); ok {
			if err = d.Set("ips", flattenLogThreatWeightIps(o["ips"], d, "ips")); err != nil {
				if vv, ok := fortiAPIPatch(o["ips"], "LogThreatWeight-Ips"); ok {
					if err = d.Set("ips", vv); err != nil {
						return fmt.Errorf("Error reading ips: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ips: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("level", flattenLogThreatWeightLevel(o["level"], d, "level")); err != nil {
			if vv, ok := fortiAPIPatch(o["level"], "LogThreatWeight-Level"); ok {
				if err = d.Set("level", vv); err != nil {
					return fmt.Errorf("Error reading level: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading level: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("level"); ok {
			if err = d.Set("level", flattenLogThreatWeightLevel(o["level"], d, "level")); err != nil {
				if vv, ok := fortiAPIPatch(o["level"], "LogThreatWeight-Level"); ok {
					if err = d.Set("level", vv); err != nil {
						return fmt.Errorf("Error reading level: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading level: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("malware", flattenLogThreatWeightMalware(o["malware"], d, "malware")); err != nil {
			if vv, ok := fortiAPIPatch(o["malware"], "LogThreatWeight-Malware"); ok {
				if err = d.Set("malware", vv); err != nil {
					return fmt.Errorf("Error reading malware: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading malware: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("malware"); ok {
			if err = d.Set("malware", flattenLogThreatWeightMalware(o["malware"], d, "malware")); err != nil {
				if vv, ok := fortiAPIPatch(o["malware"], "LogThreatWeight-Malware"); ok {
					if err = d.Set("malware", vv); err != nil {
						return fmt.Errorf("Error reading malware: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading malware: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenLogThreatWeightStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LogThreatWeight-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_block_detected", flattenLogThreatWeightUrlBlockDetected(o["url-block-detected"], d, "url_block_detected")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-block-detected"], "LogThreatWeight-UrlBlockDetected"); ok {
			if err = d.Set("url_block_detected", vv); err != nil {
				return fmt.Errorf("Error reading url_block_detected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_block_detected: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("web", flattenLogThreatWeightWeb(o["web"], d, "web")); err != nil {
			if vv, ok := fortiAPIPatch(o["web"], "LogThreatWeight-Web"); ok {
				if err = d.Set("web", vv); err != nil {
					return fmt.Errorf("Error reading web: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenLogThreatWeightWeb(o["web"], d, "web")); err != nil {
				if vv, ok := fortiAPIPatch(o["web"], "LogThreatWeight-Web"); ok {
					if err = d.Set("web", vv); err != nil {
						return fmt.Errorf("Error reading web: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenLogThreatWeightFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogThreatWeightApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandLogThreatWeightApplicationCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogThreatWeightApplicationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["level"], _ = expandLogThreatWeightApplicationLevel(d, i["level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightApplicationCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogThreatWeightApplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightApplicationLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightBlockedConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightBotnetConnectionDetected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightFailedConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["country"], _ = expandLogThreatWeightGeolocationCountry(d, i["country"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogThreatWeightGeolocationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["level"], _ = expandLogThreatWeightGeolocationLevel(d, i["level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightGeolocationCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogThreatWeightGeolocationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocationLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "critical_severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["critical-severity"], _ = expandLogThreatWeightIpsCriticalSeverity(d, i["critical_severity"], pre_append)
	}
	pre_append = pre + ".0." + "high_severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["high-severity"], _ = expandLogThreatWeightIpsHighSeverity(d, i["high_severity"], pre_append)
	}
	pre_append = pre + ".0." + "info_severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["info-severity"], _ = expandLogThreatWeightIpsInfoSeverity(d, i["info_severity"], pre_append)
	}
	pre_append = pre + ".0." + "low_severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-severity"], _ = expandLogThreatWeightIpsLowSeverity(d, i["low_severity"], pre_append)
	}
	pre_append = pre + ".0." + "medium_severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["medium-severity"], _ = expandLogThreatWeightIpsMediumSeverity(d, i["medium_severity"], pre_append)
	}

	return result, nil
}

func expandLogThreatWeightIpsCriticalSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsHighSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsInfoSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsLowSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsMediumSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "critical"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["critical"], _ = expandLogThreatWeightLevelCritical(d, i["critical"], pre_append)
	}
	pre_append = pre + ".0." + "high"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["high"], _ = expandLogThreatWeightLevelHigh(d, i["high"], pre_append)
	}
	pre_append = pre + ".0." + "low"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low"], _ = expandLogThreatWeightLevelLow(d, i["low"], pre_append)
	}
	pre_append = pre + ".0." + "medium"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["medium"], _ = expandLogThreatWeightLevelMedium(d, i["medium"], pre_append)
	}

	return result, nil
}

func expandLogThreatWeightLevelCritical(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelMedium(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "command_blocked"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["command-blocked"], _ = expandLogThreatWeightMalwareCommandBlocked(d, i["command_blocked"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandLogThreatWeightMalwareContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "ems_threat_feed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ems-threat-feed"], _ = expandLogThreatWeightMalwareEmsThreatFeed(d, i["ems_threat_feed"], pre_append)
	}
	pre_append = pre + ".0." + "file_blocked"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-blocked"], _ = expandLogThreatWeightMalwareFileBlocked(d, i["file_blocked"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandLogThreatWeightMalwareFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fsa_high_risk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fsa-high-risk"], _ = expandLogThreatWeightMalwareFsaHighRisk(d, i["fsa_high_risk"], pre_append)
	}
	pre_append = pre + ".0." + "fsa_malicious"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fsa-malicious"], _ = expandLogThreatWeightMalwareFsaMalicious(d, i["fsa_malicious"], pre_append)
	}
	pre_append = pre + ".0." + "fsa_medium_risk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fsa-medium-risk"], _ = expandLogThreatWeightMalwareFsaMediumRisk(d, i["fsa_medium_risk"], pre_append)
	}
	pre_append = pre + ".0." + "inline_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["inline-block"], _ = expandLogThreatWeightMalwareInlineBlock(d, i["inline_block"], pre_append)
	}
	pre_append = pre + ".0." + "malware_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-list"], _ = expandLogThreatWeightMalwareMalwareList(d, i["malware_list"], pre_append)
	}
	pre_append = pre + ".0." + "mimefragmented"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mimefragmented"], _ = expandLogThreatWeightMalwareMimefragmented(d, i["mimefragmented"], pre_append)
	}
	pre_append = pre + ".0." + "oversized"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["oversized"], _ = expandLogThreatWeightMalwareOversized(d, i["oversized"], pre_append)
	}
	pre_append = pre + ".0." + "switch_proto"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-proto"], _ = expandLogThreatWeightMalwareSwitchProto(d, i["switch_proto"], pre_append)
	}
	pre_append = pre + ".0." + "virus_file_type_executable"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-file-type-executable"], _ = expandLogThreatWeightMalwareVirusFileTypeExecutable(d, i["virus_file_type_executable"], pre_append)
	}
	pre_append = pre + ".0." + "virus_infected"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-infected"], _ = expandLogThreatWeightMalwareVirusInfected(d, i["virus_infected"], pre_append)
	}
	pre_append = pre + ".0." + "virus_outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-outbreak-prevention"], _ = expandLogThreatWeightMalwareVirusOutbreakPrevention(d, i["virus_outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "virus_scan_error"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-scan-error"], _ = expandLogThreatWeightMalwareVirusScanError(d, i["virus_scan_error"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandLogThreatWeightMalwareFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandLogThreatWeightMalwareFortisandbox(d, i["fortisandbox"], pre_append)
	}

	return result, nil
}

func expandLogThreatWeightMalwareCommandBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareEmsThreatFeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFileBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaHighRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaMalicious(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaMediumRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareInlineBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareMalwareList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareMimefragmented(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareOversized(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareSwitchProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusFileTypeExecutable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusScanError(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightUrlBlockDetected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandLogThreatWeightWebCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandLogThreatWeightWebId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["level"], _ = expandLogThreatWeightWebLevel(d, i["level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightWebCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLogThreatWeightWebId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightWebLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogThreatWeight(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandLogThreatWeightApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("blocked_connection"); ok || d.HasChange("blocked_connection") {
		t, err := expandLogThreatWeightBlockedConnection(d, v, "blocked_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-connection"] = t
		}
	}

	if v, ok := d.GetOk("botnet_connection_detected"); ok || d.HasChange("botnet_connection_detected") {
		t, err := expandLogThreatWeightBotnetConnectionDetected(d, v, "botnet_connection_detected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["botnet-connection-detected"] = t
		}
	}

	if v, ok := d.GetOk("failed_connection"); ok || d.HasChange("failed_connection") {
		t, err := expandLogThreatWeightFailedConnection(d, v, "failed_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failed-connection"] = t
		}
	}

	if v, ok := d.GetOk("geolocation"); ok || d.HasChange("geolocation") {
		t, err := expandLogThreatWeightGeolocation(d, v, "geolocation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geolocation"] = t
		}
	}

	if v, ok := d.GetOk("ips"); ok || d.HasChange("ips") {
		t, err := expandLogThreatWeightIps(d, v, "ips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips"] = t
		}
	}

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandLogThreatWeightLevel(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("malware"); ok || d.HasChange("malware") {
		t, err := expandLogThreatWeightMalware(d, v, "malware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malware"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLogThreatWeightStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_block_detected"); ok || d.HasChange("url_block_detected") {
		t, err := expandLogThreatWeightUrlBlockDetected(d, v, "url_block_detected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-block-detected"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok || d.HasChange("web") {
		t, err := expandLogThreatWeightWeb(d, v, "web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web"] = t
		}
	}

	return &obj, nil
}
