// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> WAF signatures.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWafProfileSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileSignatureUpdate,
		Read:   resourceWafProfileSignatureRead,
		Update: resourceWafProfileSignatureUpdate,
		Delete: resourceWafProfileSignatureDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"credit_card_detection_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"disabled_signature": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"disabled_sub_class": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"main_class": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
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
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
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

func resourceWafProfileSignatureUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWafProfileSignature(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignature resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWafProfileSignature(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WafProfileSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WafProfileSignature")

	return resourceWafProfileSignatureRead(d, m)
}

func resourceWafProfileSignatureDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWafProfileSignature(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfileSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileSignatureRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWafProfileSignature(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WafProfileSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfileSignature(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfileSignature resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileSignatureCreditCardDetectionThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignature2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWafProfileSignatureCustomSignatureAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenWafProfileSignatureCustomSignatureCaseSensitivity2edl(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenWafProfileSignatureCustomSignatureDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWafProfileSignatureCustomSignatureLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWafProfileSignatureCustomSignatureName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWafProfileSignatureCustomSignaturePattern2edl(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenWafProfileSignatureCustomSignatureSeverity2edl(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWafProfileSignatureCustomSignatureStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenWafProfileSignatureCustomSignatureTarget2edl(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "WafProfileSignature-CustomSignature-Target")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWafProfileSignatureCustomSignatureAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureCaseSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignaturePattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureDisabledSignature2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureDisabledSubClass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureMainClass2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileSignatureMainClassAction2edl(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "id"
	if _, ok := i["id"]; ok {
		result["id"] = flattenWafProfileSignatureMainClassId2edl(i["id"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileSignatureMainClassLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileSignatureMainClassSeverity2edl(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileSignatureMainClassStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileSignatureMainClassAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWafProfileSignatureMainClassLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfileSignature(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("credit_card_detection_threshold", flattenWafProfileSignatureCreditCardDetectionThreshold2edl(o["credit-card-detection-threshold"], d, "credit_card_detection_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["credit-card-detection-threshold"], "WafProfileSignature-CreditCardDetectionThreshold"); ok {
			if err = d.Set("credit_card_detection_threshold", vv); err != nil {
				return fmt.Errorf("Error reading credit_card_detection_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading credit_card_detection_threshold: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_signature", flattenWafProfileSignatureCustomSignature2edl(o["custom-signature"], d, "custom_signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-signature"], "WafProfileSignature-CustomSignature"); ok {
				if err = d.Set("custom_signature", vv); err != nil {
					return fmt.Errorf("Error reading custom_signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_signature"); ok {
			if err = d.Set("custom_signature", flattenWafProfileSignatureCustomSignature2edl(o["custom-signature"], d, "custom_signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-signature"], "WafProfileSignature-CustomSignature"); ok {
					if err = d.Set("custom_signature", vv); err != nil {
						return fmt.Errorf("Error reading custom_signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_signature: %v", err)
				}
			}
		}
	}

	if err = d.Set("disabled_signature", flattenWafProfileSignatureDisabledSignature2edl(o["disabled-signature"], d, "disabled_signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["disabled-signature"], "WafProfileSignature-DisabledSignature"); ok {
			if err = d.Set("disabled_signature", vv); err != nil {
				return fmt.Errorf("Error reading disabled_signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disabled_signature: %v", err)
		}
	}

	if err = d.Set("disabled_sub_class", flattenWafProfileSignatureDisabledSubClass2edl(o["disabled-sub-class"], d, "disabled_sub_class")); err != nil {
		if vv, ok := fortiAPIPatch(o["disabled-sub-class"], "WafProfileSignature-DisabledSubClass"); ok {
			if err = d.Set("disabled_sub_class", vv); err != nil {
				return fmt.Errorf("Error reading disabled_sub_class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disabled_sub_class: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("main_class", flattenWafProfileSignatureMainClass2edl(o["main-class"], d, "main_class")); err != nil {
			if vv, ok := fortiAPIPatch(o["main-class"], "WafProfileSignature-MainClass"); ok {
				if err = d.Set("main_class", vv); err != nil {
					return fmt.Errorf("Error reading main_class: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading main_class: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("main_class"); ok {
			if err = d.Set("main_class", flattenWafProfileSignatureMainClass2edl(o["main-class"], d, "main_class")); err != nil {
				if vv, ok := fortiAPIPatch(o["main-class"], "WafProfileSignature-MainClass"); ok {
					if err = d.Set("main_class", vv); err != nil {
						return fmt.Errorf("Error reading main_class: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading main_class: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWafProfileSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileSignatureCreditCardDetectionThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWafProfileSignatureCustomSignatureAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandWafProfileSignatureCustomSignatureCaseSensitivity2edl(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandWafProfileSignatureCustomSignatureDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWafProfileSignatureCustomSignatureLog2edl(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWafProfileSignatureCustomSignatureName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWafProfileSignatureCustomSignaturePattern2edl(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandWafProfileSignatureCustomSignatureSeverity2edl(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWafProfileSignatureCustomSignatureStatus2edl(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandWafProfileSignatureCustomSignatureTarget2edl(d, i["target"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureCustomSignatureAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureCaseSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignaturePattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureDisabledSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureDisabledSubClass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureMainClass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandWafProfileSignatureMainClassAction2edl(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["id"], _ = expandWafProfileSignatureMainClassId2edl(d, i["id"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandWafProfileSignatureMainClassLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["severity"], _ = expandWafProfileSignatureMainClassSeverity2edl(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWafProfileSignatureMainClassStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWafProfileSignatureMainClassAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWafProfileSignatureMainClassLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfileSignature(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("credit_card_detection_threshold"); ok || d.HasChange("credit_card_detection_threshold") {
		t, err := expandWafProfileSignatureCreditCardDetectionThreshold2edl(d, v, "credit_card_detection_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["credit-card-detection-threshold"] = t
		}
	}

	if v, ok := d.GetOk("custom_signature"); ok || d.HasChange("custom_signature") {
		t, err := expandWafProfileSignatureCustomSignature2edl(d, v, "custom_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-signature"] = t
		}
	}

	if v, ok := d.GetOk("disabled_signature"); ok || d.HasChange("disabled_signature") {
		t, err := expandWafProfileSignatureDisabledSignature2edl(d, v, "disabled_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disabled-signature"] = t
		}
	}

	if v, ok := d.GetOk("disabled_sub_class"); ok || d.HasChange("disabled_sub_class") {
		t, err := expandWafProfileSignatureDisabledSubClass2edl(d, v, "disabled_sub_class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disabled-sub-class"] = t
		}
	}

	if v, ok := d.GetOk("main_class"); ok || d.HasChange("main_class") {
		t, err := expandWafProfileSignatureMainClass2edl(d, v, "main_class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["main-class"] = t
		}
	}

	return &obj, nil
}
