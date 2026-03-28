// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Set up DLP rules for this profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpProfileRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpProfileRuleCreate,
		Read:   resourceDlpProfileRuleRead,
		Update: resourceDlpProfileRuleUpdate,
		Delete: resourceDlpProfileRuleDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceDlpProfileRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDlpProfileRule(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpProfileRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpProfileRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpProfileRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpProfileRule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpProfileRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpProfileRule resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpProfileRuleRead(d, m)
}

func resourceDlpProfileRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectDlpProfileRule(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpProfileRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpProfileRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpProfileRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceDlpProfileRuleRead(d, m)
}

func resourceDlpProfileRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteDlpProfileRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpProfileRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpProfileRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadDlpProfileRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpProfileRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpProfileRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpProfileRule resource from API: %v", err)
	}
	return nil
}

func flattenDlpProfileRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleArchive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleFileSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleFilterBy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleLabel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleMatchPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSensor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpProfileRuleSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpProfileRuleType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpProfileRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenDlpProfileRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "DlpProfileRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("archive", flattenDlpProfileRuleArchive2edl(o["archive"], d, "archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive"], "DlpProfileRule-Archive"); ok {
			if err = d.Set("archive", vv); err != nil {
				return fmt.Errorf("Error reading archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive: %v", err)
		}
	}

	if err = d.Set("expiry", flattenDlpProfileRuleExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "DlpProfileRule-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("file_size", flattenDlpProfileRuleFileSize2edl(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "DlpProfileRule-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("file_type", flattenDlpProfileRuleFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "DlpProfileRule-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter_by", flattenDlpProfileRuleFilterBy2edl(o["filter-by"], d, "filter_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-by"], "DlpProfileRule-FilterBy"); ok {
			if err = d.Set("filter_by", vv); err != nil {
				return fmt.Errorf("Error reading filter_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_by: %v", err)
		}
	}

	if err = d.Set("fosid", flattenDlpProfileRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "DlpProfileRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("label", flattenDlpProfileRuleLabel2edl(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "DlpProfileRule-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("match_percentage", flattenDlpProfileRuleMatchPercentage2edl(o["match-percentage"], d, "match_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-percentage"], "DlpProfileRule-MatchPercentage"); ok {
			if err = d.Set("match_percentage", vv); err != nil {
				return fmt.Errorf("Error reading match_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_percentage: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpProfileRuleName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpProfileRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proto", flattenDlpProfileRuleProto2edl(o["proto"], d, "proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["proto"], "DlpProfileRule-Proto"); ok {
			if err = d.Set("proto", vv); err != nil {
				return fmt.Errorf("Error reading proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenDlpProfileRuleSensitivity2edl(o["sensitivity"], d, "sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensitivity"], "DlpProfileRule-Sensitivity"); ok {
			if err = d.Set("sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("sensor", flattenDlpProfileRuleSensor2edl(o["sensor"], d, "sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensor"], "DlpProfileRule-Sensor"); ok {
			if err = d.Set("sensor", vv); err != nil {
				return fmt.Errorf("Error reading sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensor: %v", err)
		}
	}

	if err = d.Set("severity", flattenDlpProfileRuleSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "DlpProfileRule-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("type", flattenDlpProfileRuleType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DlpProfileRule-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDlpProfileRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpProfileRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleArchive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleFileSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleFilterBy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleLabel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleMatchPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSensor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpProfileRuleSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpProfileRuleType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpProfileRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandDlpProfileRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("archive"); ok || d.HasChange("archive") {
		t, err := expandDlpProfileRuleArchive2edl(d, v, "archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive"] = t
		}
	}

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandDlpProfileRuleExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandDlpProfileRuleFileSize2edl(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandDlpProfileRuleFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter_by"); ok || d.HasChange("filter_by") {
		t, err := expandDlpProfileRuleFilterBy2edl(d, v, "filter_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-by"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandDlpProfileRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandDlpProfileRuleLabel2edl(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("match_percentage"); ok || d.HasChange("match_percentage") {
		t, err := expandDlpProfileRuleMatchPercentage2edl(d, v, "match_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-percentage"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpProfileRuleName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proto"); ok || d.HasChange("proto") {
		t, err := expandDlpProfileRuleProto2edl(d, v, "proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proto"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok || d.HasChange("sensitivity") {
		t, err := expandDlpProfileRuleSensitivity2edl(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("sensor"); ok || d.HasChange("sensor") {
		t, err := expandDlpProfileRuleSensor2edl(d, v, "sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensor"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandDlpProfileRuleSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandDlpProfileRuleType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
