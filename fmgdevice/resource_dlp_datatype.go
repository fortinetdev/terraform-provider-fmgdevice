// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure predefined data type used by DLP blocking.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpDataType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpDataTypeCreate,
		Read:   resourceDlpDataTypeRead,
		Update: resourceDlpDataTypeUpdate,
		Delete: resourceDlpDataTypeDelete,

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
			"fgd_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"look_ahead": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"look_back": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match_ahead": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match_around": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"match_back": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_transformed_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDlpDataTypeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDataType(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpDataType resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadDlpDataType(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateDlpDataType(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating DlpDataType resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateDlpDataType(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating DlpDataType resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceDlpDataTypeRead(d, m)
}

func resourceDlpDataTypeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDataType(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpDataType resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDlpDataType(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DlpDataType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDlpDataTypeRead(d, m)
}

func resourceDlpDataTypeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDlpDataType(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DlpDataType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpDataTypeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpDataType(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DlpDataType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpDataType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpDataType resource from API: %v", err)
	}
	return nil
}

func flattenDlpDataTypeComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeFgdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeLookAhead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeLookBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeMatchAhead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeMatchAround(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDlpDataTypeMatchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeTransform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeVerifyTransformedPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpDataTypeVerify2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpDataType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenDlpDataTypeComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "DlpDataType-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fgd_id", flattenDlpDataTypeFgdId(o["fgd-id"], d, "fgd_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-id"], "DlpDataType-FgdId"); ok {
			if err = d.Set("fgd_id", vv); err != nil {
				return fmt.Errorf("Error reading fgd_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_id: %v", err)
		}
	}

	if err = d.Set("look_ahead", flattenDlpDataTypeLookAhead(o["look-ahead"], d, "look_ahead")); err != nil {
		if vv, ok := fortiAPIPatch(o["look-ahead"], "DlpDataType-LookAhead"); ok {
			if err = d.Set("look_ahead", vv); err != nil {
				return fmt.Errorf("Error reading look_ahead: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading look_ahead: %v", err)
		}
	}

	if err = d.Set("look_back", flattenDlpDataTypeLookBack(o["look-back"], d, "look_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["look-back"], "DlpDataType-LookBack"); ok {
			if err = d.Set("look_back", vv); err != nil {
				return fmt.Errorf("Error reading look_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading look_back: %v", err)
		}
	}

	if err = d.Set("match_ahead", flattenDlpDataTypeMatchAhead(o["match-ahead"], d, "match_ahead")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-ahead"], "DlpDataType-MatchAhead"); ok {
			if err = d.Set("match_ahead", vv); err != nil {
				return fmt.Errorf("Error reading match_ahead: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_ahead: %v", err)
		}
	}

	if err = d.Set("match_around", flattenDlpDataTypeMatchAround(o["match-around"], d, "match_around")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-around"], "DlpDataType-MatchAround"); ok {
			if err = d.Set("match_around", vv); err != nil {
				return fmt.Errorf("Error reading match_around: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_around: %v", err)
		}
	}

	if err = d.Set("match_back", flattenDlpDataTypeMatchBack(o["match-back"], d, "match_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-back"], "DlpDataType-MatchBack"); ok {
			if err = d.Set("match_back", vv); err != nil {
				return fmt.Errorf("Error reading match_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_back: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpDataTypeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DlpDataType-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern", flattenDlpDataTypePattern(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "DlpDataType-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("transform", flattenDlpDataTypeTransform(o["transform"], d, "transform")); err != nil {
		if vv, ok := fortiAPIPatch(o["transform"], "DlpDataType-Transform"); ok {
			if err = d.Set("transform", vv); err != nil {
				return fmt.Errorf("Error reading transform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transform: %v", err)
		}
	}

	if err = d.Set("verify", flattenDlpDataTypeVerify(o["verify"], d, "verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify"], "DlpDataType-Verify"); ok {
			if err = d.Set("verify", vv); err != nil {
				return fmt.Errorf("Error reading verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify: %v", err)
		}
	}

	if err = d.Set("verify_transformed_pattern", flattenDlpDataTypeVerifyTransformedPattern(o["verify-transformed-pattern"], d, "verify_transformed_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-transformed-pattern"], "DlpDataType-VerifyTransformedPattern"); ok {
			if err = d.Set("verify_transformed_pattern", vv); err != nil {
				return fmt.Errorf("Error reading verify_transformed_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_transformed_pattern: %v", err)
		}
	}

	if err = d.Set("verify2", flattenDlpDataTypeVerify2(o["verify2"], d, "verify2")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify2"], "DlpDataType-Verify2"); ok {
			if err = d.Set("verify2", vv); err != nil {
				return fmt.Errorf("Error reading verify2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify2: %v", err)
		}
	}

	return nil
}

func flattenDlpDataTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpDataTypeComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeFgdId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeLookAhead(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeLookBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeMatchAhead(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeMatchAround(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDlpDataTypeMatchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeTransform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerifyTransformedPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerify2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpDataType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandDlpDataTypeComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fgd_id"); ok || d.HasChange("fgd_id") {
		t, err := expandDlpDataTypeFgdId(d, v, "fgd_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-id"] = t
		}
	}

	if v, ok := d.GetOk("look_ahead"); ok || d.HasChange("look_ahead") {
		t, err := expandDlpDataTypeLookAhead(d, v, "look_ahead")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-ahead"] = t
		}
	}

	if v, ok := d.GetOk("look_back"); ok || d.HasChange("look_back") {
		t, err := expandDlpDataTypeLookBack(d, v, "look_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-back"] = t
		}
	}

	if v, ok := d.GetOk("match_ahead"); ok || d.HasChange("match_ahead") {
		t, err := expandDlpDataTypeMatchAhead(d, v, "match_ahead")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ahead"] = t
		}
	}

	if v, ok := d.GetOk("match_around"); ok || d.HasChange("match_around") {
		t, err := expandDlpDataTypeMatchAround(d, v, "match_around")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-around"] = t
		}
	}

	if v, ok := d.GetOk("match_back"); ok || d.HasChange("match_back") {
		t, err := expandDlpDataTypeMatchBack(d, v, "match_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-back"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDlpDataTypeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandDlpDataTypePattern(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("transform"); ok || d.HasChange("transform") {
		t, err := expandDlpDataTypeTransform(d, v, "transform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform"] = t
		}
	}

	if v, ok := d.GetOk("verify"); ok || d.HasChange("verify") {
		t, err := expandDlpDataTypeVerify(d, v, "verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify"] = t
		}
	}

	if v, ok := d.GetOk("verify_transformed_pattern"); ok || d.HasChange("verify_transformed_pattern") {
		t, err := expandDlpDataTypeVerifyTransformedPattern(d, v, "verify_transformed_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-transformed-pattern"] = t
		}
	}

	if v, ok := d.GetOk("verify2"); ok || d.HasChange("verify2") {
		t, err := expandDlpDataTypeVerify2(d, v, "verify2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify2"] = t
		}
	}

	return &obj, nil
}
