// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB profile custom control.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfileSaasApplicationCustomControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileSaasApplicationCustomControlCreate,
		Read:   resourceCasbProfileSaasApplicationCustomControlRead,
		Update: resourceCasbProfileSaasApplicationCustomControlUpdate,
		Delete: resourceCasbProfileSaasApplicationCustomControlDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attribute_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_match": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_input": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceCasbProfileSaasApplicationCustomControlCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectCasbProfileSaasApplicationCustomControl(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfileSaasApplicationCustomControl resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfileSaasApplicationCustomControl(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfileSaasApplicationCustomControl(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControl resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfileSaasApplicationCustomControl(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfileSaasApplicationCustomControl resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationCustomControlRead(d, m)
}

func resourceCasbProfileSaasApplicationCustomControlUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectCasbProfileSaasApplicationCustomControl(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControl resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfileSaasApplicationCustomControl(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationCustomControlRead(d, m)
}

func resourceCasbProfileSaasApplicationCustomControlDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	wsParams["adom"] = adomv

	err = c.DeleteCasbProfileSaasApplicationCustomControl(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileSaasApplicationCustomControlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
	saas_application := d.Get("saas_application").(string)
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
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	o, err := c.ReadCasbProfileSaasApplicationCustomControl(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfileSaasApplicationCustomControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfileSaasApplicationCustomControl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfileSaasApplicationCustomControl resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilter3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction3rdl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch3rdl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationCustomControlOption3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOptionName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlOptionName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("attribute_filter", flattenCasbProfileSaasApplicationCustomControlAttributeFilter3rdl(o["attribute-filter"], d, "attribute_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["attribute-filter"], "CasbProfileSaasApplicationCustomControl-AttributeFilter"); ok {
				if err = d.Set("attribute_filter", vv); err != nil {
					return fmt.Errorf("Error reading attribute_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading attribute_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("attribute_filter"); ok {
			if err = d.Set("attribute_filter", flattenCasbProfileSaasApplicationCustomControlAttributeFilter3rdl(o["attribute-filter"], d, "attribute_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["attribute-filter"], "CasbProfileSaasApplicationCustomControl-AttributeFilter"); ok {
					if err = d.Set("attribute_filter", vv); err != nil {
						return fmt.Errorf("Error reading attribute_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading attribute_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenCasbProfileSaasApplicationCustomControlName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbProfileSaasApplicationCustomControl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("option", flattenCasbProfileSaasApplicationCustomControlOption3rdl(o["option"], d, "option")); err != nil {
			if vv, ok := fortiAPIPatch(o["option"], "CasbProfileSaasApplicationCustomControl-Option"); ok {
				if err = d.Set("option", vv); err != nil {
					return fmt.Errorf("Error reading option: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading option: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("option"); ok {
			if err = d.Set("option", flattenCasbProfileSaasApplicationCustomControlOption3rdl(o["option"], d, "option")); err != nil {
				if vv, ok := fortiAPIPatch(o["option"], "CasbProfileSaasApplicationCustomControl-Option"); ok {
					if err = d.Set("option", vv); err != nil {
						return fmt.Errorf("Error reading option: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading option: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbProfileSaasApplicationCustomControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilter3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAction3rdl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch3rdl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterId3rdl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlOptionName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInput3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectCasbProfileSaasApplicationCustomControl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attribute_filter"); ok || d.HasChange("attribute_filter") {
		t, err := expandCasbProfileSaasApplicationCustomControlAttributeFilter3rdl(d, v, "attribute_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-filter"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbProfileSaasApplicationCustomControlName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandCasbProfileSaasApplicationCustomControlOption3rdl(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
