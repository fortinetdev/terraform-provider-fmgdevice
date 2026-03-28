// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure DLP profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpProfileCreate,
		Read:   resourceDlpProfileRead,
		Update: resourceDlpProfileUpdate,
		Delete: resourceDlpProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortidata_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_archive_proto": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"nac_quar_log": &schema.Schema{
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
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"archive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"file_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filter_by": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"label": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"match_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"proto": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sensitivity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sensor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"summary_proto": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceDlpProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpProfileRead(d, m)
}

func resourceDlpProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpProfileRead(d, m)
}

func resourceDlpProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpProfile resource from API: %v", err)
	}
	return nil
}

func flattenDlpProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileDlpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileFortidataErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileFullArchiveProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDlpProfileRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := i["archive"]; ok {
			v := flattenDlpProfileRuleArchive(i["archive"], d, pre_append)
			tmp["archive"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Archive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := i["expiry"]; ok {
			v := flattenDlpProfileRuleExpiry(i["expiry"], d, pre_append)
			tmp["expiry"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Expiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := i["file-size"]; ok {
			v := flattenDlpProfileRuleFileSize(i["file-size"], d, pre_append)
			tmp["file_size"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-FileSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenDlpProfileRuleFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := i["filter-by"]; ok {
			v := flattenDlpProfileRuleFilterBy(i["filter-by"], d, pre_append)
			tmp["filter_by"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-FilterBy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDlpProfileRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "label"
		if _, ok := i["label"]; ok {
			v := flattenDlpProfileRuleLabel(i["label"], d, pre_append)
			tmp["label"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Label")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := i["match-percentage"]; ok {
			v := flattenDlpProfileRuleMatchPercentage(i["match-percentage"], d, pre_append)
			tmp["match_percentage"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-MatchPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDlpProfileRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := i["proto"]; ok {
			v := flattenDlpProfileRuleProto(i["proto"], d, pre_append)
			tmp["proto"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Proto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := i["sensitivity"]; ok {
			v := flattenDlpProfileRuleSensitivity(i["sensitivity"], d, pre_append)
			tmp["sensitivity"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Sensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensor"
		if _, ok := i["sensor"]; ok {
			v := flattenDlpProfileRuleSensor(i["sensor"], d, pre_append)
			tmp["sensor"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Sensor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenDlpProfileRuleSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenDlpProfileRuleType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "DlpProfile-Rule-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpProfileRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleFilterBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleMatchPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileSummaryProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDlpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenDlpProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dlp_log", flattenDlpProfileDlpLog(o["dlp-log"], d, "dlp_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-log"], "DlpProfile-DlpLog"); ok {
			if err = d.Set("dlp_log", vv); err != nil {
				return fmt.Errorf("Error reading dlp_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenDlpProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "DlpProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenDlpProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "DlpProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("fortidata_error_action", flattenDlpProfileFortidataErrorAction(o["fortidata-error-action"], d, "fortidata_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortidata-error-action"], "DlpProfile-FortidataErrorAction"); ok {
			if err = d.Set("fortidata_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortidata_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortidata_error_action: %v", err)
		}
	}

	if err = d.Set("full_archive_proto", flattenDlpProfileFullArchiveProto(o["full-archive-proto"], d, "full_archive_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-archive-proto"], "DlpProfile-FullArchiveProto"); ok {
			if err = d.Set("full_archive_proto", vv); err != nil {
				return fmt.Errorf("Error reading full_archive_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_archive_proto: %v", err)
		}
	}

	if err = d.Set("nac_quar_log", flattenDlpProfileNacQuarLog(o["nac-quar-log"], d, "nac_quar_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-quar-log"], "DlpProfile-NacQuarLog"); ok {
			if err = d.Set("nac_quar_log", vv); err != nil {
				return fmt.Errorf("Error reading nac_quar_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_quar_log: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenDlpProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "DlpProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenDlpProfileRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "DlpProfile-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenDlpProfileRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "DlpProfile-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("summary_proto", flattenDlpProfileSummaryProto(o["summary-proto"], d, "summary_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["summary-proto"], "DlpProfile-SummaryProto"); ok {
			if err = d.Set("summary_proto", vv); err != nil {
				return fmt.Errorf("Error reading summary_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading summary_proto: %v", err)
		}
	}

	return nil
}

func flattenDlpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileDlpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileFortidataErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileFullArchiveProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandDlpProfileRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["archive"], _ = expandDlpProfileRuleArchive(d, i["archive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiry"], _ = expandDlpProfileRuleExpiry(d, i["expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-size"], _ = expandDlpProfileRuleFileSize(d, i["file_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandDlpProfileRuleFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-by"], _ = expandDlpProfileRuleFilterBy(d, i["filter_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDlpProfileRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "label"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["label"], _ = expandDlpProfileRuleLabel(d, i["label"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-percentage"], _ = expandDlpProfileRuleMatchPercentage(d, i["match_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandDlpProfileRuleName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proto"], _ = expandDlpProfileRuleProto(d, i["proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sensitivity"], _ = expandDlpProfileRuleSensitivity(d, i["sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sensor"], _ = expandDlpProfileRuleSensor(d, i["sensor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandDlpProfileRuleSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandDlpProfileRuleType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpProfileRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleFilterBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleMatchPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileSummaryProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDlpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dlp_log"); ok || d.HasChange("dlp_log") {
		t, err := expandDlpProfileDlpLog(d, v, "dlp_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandDlpProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandDlpProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("fortidata_error_action"); ok || d.HasChange("fortidata_error_action") {
		t, err := expandDlpProfileFortidataErrorAction(d, v, "fortidata_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortidata-error-action"] = t
		}
	}

	if v, ok := d.GetOk("full_archive_proto"); ok || d.HasChange("full_archive_proto") {
		t, err := expandDlpProfileFullArchiveProto(d, v, "full_archive_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-archive-proto"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar_log"); ok || d.HasChange("nac_quar_log") {
		t, err := expandDlpProfileNacQuarLog(d, v, "nac_quar_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar-log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandDlpProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandDlpProfileRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("summary_proto"); ok || d.HasChange("summary_proto") {
		t, err := expandDlpProfileSummaryProto(d, v, "summary_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-proto"] = t
		}
	}

	return &obj, nil
}
