// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSensorFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSensorFilterCreate,
		Read:   resourceDlpSensorFilterRead,
		Update: resourceDlpSensorFilterUpdate,
		Delete: resourceDlpSensorFilterDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceDlpSensorFilterCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSensorFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpSensorFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpSensorFilter(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpSensorFilter(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpSensorFilter resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpSensorFilter(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpSensorFilter resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpSensorFilterRead(d, m)
}

func resourceDlpSensorFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSensorFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensorFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpSensorFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensorFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpSensorFilterRead(d, m)
}

func resourceDlpSensorFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpSensorFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSensorFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSensorFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpSensorFilter(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpSensorFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSensorFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensorFilter resource from API: %v", err)
	}
	return nil
}

func flattenDlpSensorFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterArchive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterCompanyIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterFilterBy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterMatchPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterRegexp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpSensorFilterSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpSensorFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenDlpSensorFilterAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "DlpSensorFilter-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("archive", flattenDlpSensorFilterArchive2edl(o["archive"], d, "archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive"], "DlpSensorFilter-Archive"); ok {
			if err = d.Set("archive", vv); err != nil {
				return fmt.Errorf("Error reading archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive: %v", err)
		}
	}

	if err = d.Set("company_identifier", flattenDlpSensorFilterCompanyIdentifier2edl(o["company-identifier"], d, "company_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["company-identifier"], "DlpSensorFilter-CompanyIdentifier"); ok {
			if err = d.Set("company_identifier", vv); err != nil {
				return fmt.Errorf("Error reading company_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading company_identifier: %v", err)
		}
	}

	if err = d.Set("expiry", flattenDlpSensorFilterExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "DlpSensorFilter-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("file_size", flattenDlpSensorFilterFileSize2edl(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "DlpSensorFilter-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("file_type", flattenDlpSensorFilterFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "DlpSensorFilter-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter_by", flattenDlpSensorFilterFilterBy2edl(o["filter-by"], d, "filter_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-by"], "DlpSensorFilter-FilterBy"); ok {
			if err = d.Set("filter_by", vv); err != nil {
				return fmt.Errorf("Error reading filter_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_by: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDlpSensorFilterId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DlpSensorFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_percentage", flattenDlpSensorFilterMatchPercentage2edl(o["match-percentage"], d, "match_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-percentage"], "DlpSensorFilter-MatchPercentage"); ok {
			if err = d.Set("match_percentage", vv); err != nil {
				return fmt.Errorf("Error reading match_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_percentage: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpSensorFilterName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpSensorFilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proto", flattenDlpSensorFilterProto2edl(o["proto"], d, "proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["proto"], "DlpSensorFilter-Proto"); ok {
			if err = d.Set("proto", vv); err != nil {
				return fmt.Errorf("Error reading proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("regexp", flattenDlpSensorFilterRegexp2edl(o["regexp"], d, "regexp")); err != nil {
		if vv, ok := fortiAPIPatch(o["regexp"], "DlpSensorFilter-Regexp"); ok {
			if err = d.Set("regexp", vv); err != nil {
				return fmt.Errorf("Error reading regexp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regexp: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenDlpSensorFilterSensitivity2edl(o["sensitivity"], d, "sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensitivity"], "DlpSensorFilter-Sensitivity"); ok {
			if err = d.Set("sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("severity", flattenDlpSensorFilterSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "DlpSensorFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("type", flattenDlpSensorFilterType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DlpSensorFilter-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDlpSensorFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSensorFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterArchive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterCompanyIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterFilterBy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterMatchPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterRegexp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpSensorFilterSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSensorFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandDlpSensorFilterAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("archive"); ok || d.HasChange("archive") {
		t, err := expandDlpSensorFilterArchive2edl(d, v, "archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive"] = t
		}
	}

	if v, ok := d.GetOk("company_identifier"); ok || d.HasChange("company_identifier") {
		t, err := expandDlpSensorFilterCompanyIdentifier2edl(d, v, "company_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company-identifier"] = t
		}
	}

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandDlpSensorFilterExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandDlpSensorFilterFileSize2edl(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandDlpSensorFilterFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter_by"); ok || d.HasChange("filter_by") {
		t, err := expandDlpSensorFilterFilterBy2edl(d, v, "filter_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-by"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDlpSensorFilterId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_percentage"); ok || d.HasChange("match_percentage") {
		t, err := expandDlpSensorFilterMatchPercentage2edl(d, v, "match_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-percentage"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpSensorFilterName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proto"); ok || d.HasChange("proto") {
		t, err := expandDlpSensorFilterProto2edl(d, v, "proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proto"] = t
		}
	}

	if v, ok := d.GetOk("regexp"); ok || d.HasChange("regexp") {
		t, err := expandDlpSensorFilterRegexp2edl(d, v, "regexp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regexp"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok || d.HasChange("sensitivity") {
		t, err := expandDlpSensorFilterSensitivity2edl(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandDlpSensorFilterSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandDlpSensorFilterType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
