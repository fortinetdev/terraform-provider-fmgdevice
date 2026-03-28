// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure sensors used by DLP blocking.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSensorCreate,
		Read:   resourceDlpSensorRead,
		Update: resourceDlpSensorUpdate,
		Delete: resourceDlpSensorDelete,

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
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dictionary": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"eval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgd_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dlp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"archive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"company_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"regexp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sensitivity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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

func resourceDlpSensorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpSensor resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpSensor(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpSensor(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpSensor resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpSensor(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpSensor resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpSensorRead(d, m)
}

func resourceDlpSensorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensor resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSensor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpSensorRead(d, m)
}

func resourceDlpSensorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpSensor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSensorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpSensor(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensor resource from API: %v", err)
	}
	return nil
}

func flattenDlpSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if _, ok := i["count"]; ok {
			v := flattenDlpSensorEntriesCount(i["count"], d, pre_append)
			tmp["count"] = fortiAPISubPartPatch(v, "DlpSensor-Entries-Count")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dictionary"
		if _, ok := i["dictionary"]; ok {
			v := flattenDlpSensorEntriesDictionary(i["dictionary"], d, pre_append)
			tmp["dictionary"] = fortiAPISubPartPatch(v, "DlpSensor-Entries-Dictionary")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDlpSensorEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DlpSensor-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenDlpSensorEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "DlpSensor-Entries-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpSensorEntriesCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEntriesDictionary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorEval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFgdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorMatchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorDlpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDlpSensorFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := i["archive"]; ok {
			v := flattenDlpSensorFilterArchive(i["archive"], d, pre_append)
			tmp["archive"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Archive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := i["company-identifier"]; ok {
			v := flattenDlpSensorFilterCompanyIdentifier(i["company-identifier"], d, pre_append)
			tmp["company_identifier"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-CompanyIdentifier")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := i["expiry"]; ok {
			v := flattenDlpSensorFilterExpiry(i["expiry"], d, pre_append)
			tmp["expiry"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Expiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := i["file-size"]; ok {
			v := flattenDlpSensorFilterFileSize(i["file-size"], d, pre_append)
			tmp["file_size"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-FileSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenDlpSensorFilterFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := i["filter-by"]; ok {
			v := flattenDlpSensorFilterFilterBy(i["filter-by"], d, pre_append)
			tmp["filter_by"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-FilterBy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenDlpSensorFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := i["match-percentage"]; ok {
			v := flattenDlpSensorFilterMatchPercentage(i["match-percentage"], d, pre_append)
			tmp["match_percentage"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-MatchPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDlpSensorFilterName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := i["proto"]; ok {
			v := flattenDlpSensorFilterProto(i["proto"], d, pre_append)
			tmp["proto"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Proto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			v := flattenDlpSensorFilterRegexp(i["regexp"], d, pre_append)
			tmp["regexp"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Regexp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := i["sensitivity"]; ok {
			v := flattenDlpSensorFilterSensitivity(i["sensitivity"], d, pre_append)
			tmp["sensitivity"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Sensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenDlpSensorFilterSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenDlpSensorFilterType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "DlpSensor-Filter-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDlpSensorFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterCompanyIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterFilterBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterMatchPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFullArchiveProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorSummaryProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDlpSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenDlpSensorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpSensor-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenDlpSensorEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "DlpSensor-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenDlpSensorEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "DlpSensor-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("eval", flattenDlpSensorEval(o["eval"], d, "eval")); err != nil {
		if vv, ok := fortiAPIPatch(o["eval"], "DlpSensor-Eval"); ok {
			if err = d.Set("eval", vv); err != nil {
				return fmt.Errorf("Error reading eval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eval: %v", err)
		}
	}

	if err = d.Set("fgd_id", flattenDlpSensorFgdId(o["fgd-id"], d, "fgd_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-id"], "DlpSensor-FgdId"); ok {
			if err = d.Set("fgd_id", vv); err != nil {
				return fmt.Errorf("Error reading fgd_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_id: %v", err)
		}
	}

	if err = d.Set("match_type", flattenDlpSensorMatchType(o["match-type"], d, "match_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-type"], "DlpSensor-MatchType"); ok {
			if err = d.Set("match_type", vv); err != nil {
				return fmt.Errorf("Error reading match_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpSensorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpSensor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dlp_log", flattenDlpSensorDlpLog(o["dlp-log"], d, "dlp_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-log"], "DlpSensor-DlpLog"); ok {
			if err = d.Set("dlp_log", vv); err != nil {
				return fmt.Errorf("Error reading dlp_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenDlpSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "DlpSensor-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenDlpSensorFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "DlpSensor-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["filter"], "DlpSensor-Filter"); ok {
				if err = d.Set("filter", vv); err != nil {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["filter"], "DlpSensor-Filter"); ok {
					if err = d.Set("filter", vv); err != nil {
						return fmt.Errorf("Error reading filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("full_archive_proto", flattenDlpSensorFullArchiveProto(o["full-archive-proto"], d, "full_archive_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-archive-proto"], "DlpSensor-FullArchiveProto"); ok {
			if err = d.Set("full_archive_proto", vv); err != nil {
				return fmt.Errorf("Error reading full_archive_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_archive_proto: %v", err)
		}
	}

	if err = d.Set("nac_quar_log", flattenDlpSensorNacQuarLog(o["nac-quar-log"], d, "nac_quar_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-quar-log"], "DlpSensor-NacQuarLog"); ok {
			if err = d.Set("nac_quar_log", vv); err != nil {
				return fmt.Errorf("Error reading nac_quar_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_quar_log: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenDlpSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "DlpSensor-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("summary_proto", flattenDlpSensorSummaryProto(o["summary-proto"], d, "summary_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["summary-proto"], "DlpSensor-SummaryProto"); ok {
			if err = d.Set("summary_proto", vv); err != nil {
				return fmt.Errorf("Error reading summary_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading summary_proto: %v", err)
		}
	}

	return nil
}

func flattenDlpSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["count"], _ = expandDlpSensorEntriesCount(d, i["count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dictionary"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dictionary"], _ = expandDlpSensorEntriesDictionary(d, i["dictionary"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDlpSensorEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandDlpSensorEntriesStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpSensorEntriesCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesDictionary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFgdId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorMatchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorDlpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandDlpSensorFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["archive"], _ = expandDlpSensorFilterArchive(d, i["archive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company-identifier"], _ = expandDlpSensorFilterCompanyIdentifier(d, i["company_identifier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiry"], _ = expandDlpSensorFilterExpiry(d, i["expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-size"], _ = expandDlpSensorFilterFileSize(d, i["file_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandDlpSensorFilterFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-by"], _ = expandDlpSensorFilterFilterBy(d, i["filter_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandDlpSensorFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-percentage"], _ = expandDlpSensorFilterMatchPercentage(d, i["match_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandDlpSensorFilterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proto"], _ = expandDlpSensorFilterProto(d, i["proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regexp"], _ = expandDlpSensorFilterRegexp(d, i["regexp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sensitivity"], _ = expandDlpSensorFilterSensitivity(d, i["sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandDlpSensorFilterSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandDlpSensorFilterType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterCompanyIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterFilterBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterMatchPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterRegexp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFullArchiveProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorSummaryProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDlpSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandDlpSensorEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("eval"); ok || d.HasChange("eval") {
		t, err := expandDlpSensorEval(d, v, "eval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eval"] = t
		}
	}

	if v, ok := d.GetOk("fgd_id"); ok || d.HasChange("fgd_id") {
		t, err := expandDlpSensorFgdId(d, v, "fgd_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-id"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok || d.HasChange("match_type") {
		t, err := expandDlpSensorMatchType(d, v, "match_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("dlp_log"); ok || d.HasChange("dlp_log") {
		t, err := expandDlpSensorDlpLog(d, v, "dlp_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandDlpSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandDlpSensorFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandDlpSensorFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("full_archive_proto"); ok || d.HasChange("full_archive_proto") {
		t, err := expandDlpSensorFullArchiveProto(d, v, "full_archive_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-archive-proto"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar_log"); ok || d.HasChange("nac_quar_log") {
		t, err := expandDlpSensorNacQuarLog(d, v, "nac_quar_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar-log"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandDlpSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("summary_proto"); ok || d.HasChange("summary_proto") {
		t, err := expandDlpSensorSummaryProto(d, v, "summary_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-proto"] = t
		}
	}

	return &obj, nil
}
