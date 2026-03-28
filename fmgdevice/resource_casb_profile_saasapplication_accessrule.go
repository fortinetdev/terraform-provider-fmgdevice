// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB profile access rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfileSaasApplicationAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileSaasApplicationAccessRuleCreate,
		Read:   resourceCasbProfileSaasApplicationAccessRuleRead,
		Update: resourceCasbProfileSaasApplicationAccessRuleUpdate,
		Delete: resourceCasbProfileSaasApplicationAccessRuleDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"bypass": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceCasbProfileSaasApplicationAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfileSaasApplicationAccessRule(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfileSaasApplicationAccessRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfileSaasApplicationAccessRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfileSaasApplicationAccessRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfileSaasApplicationAccessRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfileSaasApplicationAccessRule resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationAccessRuleRead(d, m)
}

func resourceCasbProfileSaasApplicationAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfileSaasApplicationAccessRule(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfileSaasApplicationAccessRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationAccessRuleRead(d, m)
}

func resourceCasbProfileSaasApplicationAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbProfileSaasApplicationAccessRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileSaasApplicationAccessRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbProfileSaasApplicationAccessRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfileSaasApplicationAccessRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfileSaasApplicationAccessRule resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileSaasApplicationAccessRuleAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilter3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction3rdl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch3rdl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleBypass3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenCasbProfileSaasApplicationAccessRuleAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "CasbProfileSaasApplicationAccessRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("attribute_filter", flattenCasbProfileSaasApplicationAccessRuleAttributeFilter3rdl(o["attribute-filter"], d, "attribute_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["attribute-filter"], "CasbProfileSaasApplicationAccessRule-AttributeFilter"); ok {
				if err = d.Set("attribute_filter", vv); err != nil {
					return fmt.Errorf("Error reading attribute_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading attribute_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("attribute_filter"); ok {
			if err = d.Set("attribute_filter", flattenCasbProfileSaasApplicationAccessRuleAttributeFilter3rdl(o["attribute-filter"], d, "attribute_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["attribute-filter"], "CasbProfileSaasApplicationAccessRule-AttributeFilter"); ok {
					if err = d.Set("attribute_filter", vv); err != nil {
						return fmt.Errorf("Error reading attribute_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading attribute_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("bypass", flattenCasbProfileSaasApplicationAccessRuleBypass3rdl(o["bypass"], d, "bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["bypass"], "CasbProfileSaasApplicationAccessRule-Bypass"); ok {
			if err = d.Set("bypass", vv); err != nil {
				return fmt.Errorf("Error reading bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bypass: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbProfileSaasApplicationAccessRuleName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbProfileSaasApplicationAccessRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenCasbProfileSaasApplicationAccessRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileSaasApplicationAccessRuleAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilter3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction3rdl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch3rdl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterId3rdl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleBypass3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbProfileSaasApplicationAccessRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandCasbProfileSaasApplicationAccessRuleAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("attribute_filter"); ok || d.HasChange("attribute_filter") {
		t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilter3rdl(d, v, "attribute_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-filter"] = t
		}
	}

	if v, ok := d.GetOk("bypass"); ok || d.HasChange("bypass") {
		t, err := expandCasbProfileSaasApplicationAccessRuleBypass3rdl(d, v, "bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bypass"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbProfileSaasApplicationAccessRuleName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
