// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> FortiGuard Web Filter settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileFtgdWf() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileFtgdWfUpdate,
		Read:   resourceWebfilterProfileFtgdWfRead,
		Update: resourceWebfilterProfileFtgdWfUpdate,
		Delete: resourceWebfilterProfileFtgdWfDelete,

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
			"exempt_quota": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_usr_grp": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_replacemsg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warn_duration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warning_duration_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"warning_prompt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"max_quota_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ovrd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quota": &schema.Schema{
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
						"duration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"override_replacemsg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"rate_crl_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_css_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_image_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rate_javascript_urls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"risk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"risk_level": &schema.Schema{
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

func resourceWebfilterProfileFtgdWfUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfileFtgdWf(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWf resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileFtgdWf(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileFtgdWf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterProfileFtgdWf")

	return resourceWebfilterProfileFtgdWfRead(d, m)
}

func resourceWebfilterProfileFtgdWfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterProfileFtgdWf(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileFtgdWf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileFtgdWfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterProfileFtgdWf(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileFtgdWf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileFtgdWf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileFtgdWf resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileFtgdWfExemptQuota2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWebfilterProfileFtgdWfFiltersAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := i["auth-usr-grp"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(i["auth-usr-grp"], d, pre_append)
			tmp["auth_usr_grp"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-AuthUsrGrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := i["warn-duration"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarnDuration2edl(i["warn-duration"], d, pre_append)
			tmp["warn_duration"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarnDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := i["warning-duration-type"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarningDurationType2edl(i["warning-duration-type"], d, pre_append)
			tmp["warning_duration_type"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarningDurationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := i["warning-prompt"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarningPrompt2edl(i["warning-prompt"], d, pre_append)
			tmp["warning_prompt"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarningPrompt")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarnDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningDurationType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningPrompt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfMaxQuotaTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfOvrd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfQuota2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWebfilterProfileFtgdWfQuotaCategory2edl(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := i["duration"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaDuration2edl(i["duration"], d, pre_append)
			tmp["duration"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Duration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := i["unit"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaUnit2edl(i["unit"], d, pre_append)
			tmp["unit"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Unit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfQuotaCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfQuotaDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaUnit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCrlUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCssUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateImageUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateJavascriptUrls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRisk2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWebfilterProfileFtgdWfRiskAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk_level"
		if _, ok := i["risk-level"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskRiskLevel2edl(i["risk-level"], d, pre_append)
			tmp["risk_level"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-RiskLevel")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfRiskAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskRiskLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWebfilterProfileFtgdWf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("exempt_quota", flattenWebfilterProfileFtgdWfExemptQuota2edl(o["exempt-quota"], d, "exempt_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["exempt-quota"], "WebfilterProfileFtgdWf-ExemptQuota"); ok {
			if err = d.Set("exempt_quota", vv); err != nil {
				return fmt.Errorf("Error reading exempt_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exempt_quota: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filters", flattenWebfilterProfileFtgdWfFilters2edl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "WebfilterProfileFtgdWf-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenWebfilterProfileFtgdWfFilters2edl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "WebfilterProfileFtgdWf-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_quota_timeout", flattenWebfilterProfileFtgdWfMaxQuotaTimeout2edl(o["max-quota-timeout"], d, "max_quota_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-quota-timeout"], "WebfilterProfileFtgdWf-MaxQuotaTimeout"); ok {
			if err = d.Set("max_quota_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_quota_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_quota_timeout: %v", err)
		}
	}

	if err = d.Set("options", flattenWebfilterProfileFtgdWfOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "WebfilterProfileFtgdWf-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("ovrd", flattenWebfilterProfileFtgdWfOvrd2edl(o["ovrd"], d, "ovrd")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd"], "WebfilterProfileFtgdWf-Ovrd"); ok {
			if err = d.Set("ovrd", vv); err != nil {
				return fmt.Errorf("Error reading ovrd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("quota", flattenWebfilterProfileFtgdWfQuota2edl(o["quota"], d, "quota")); err != nil {
			if vv, ok := fortiAPIPatch(o["quota"], "WebfilterProfileFtgdWf-Quota"); ok {
				if err = d.Set("quota", vv); err != nil {
					return fmt.Errorf("Error reading quota: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quota: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quota"); ok {
			if err = d.Set("quota", flattenWebfilterProfileFtgdWfQuota2edl(o["quota"], d, "quota")); err != nil {
				if vv, ok := fortiAPIPatch(o["quota"], "WebfilterProfileFtgdWf-Quota"); ok {
					if err = d.Set("quota", vv); err != nil {
						return fmt.Errorf("Error reading quota: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading quota: %v", err)
				}
			}
		}
	}

	if err = d.Set("rate_crl_urls", flattenWebfilterProfileFtgdWfRateCrlUrls2edl(o["rate-crl-urls"], d, "rate_crl_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-crl-urls"], "WebfilterProfileFtgdWf-RateCrlUrls"); ok {
			if err = d.Set("rate_crl_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_crl_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_crl_urls: %v", err)
		}
	}

	if err = d.Set("rate_css_urls", flattenWebfilterProfileFtgdWfRateCssUrls2edl(o["rate-css-urls"], d, "rate_css_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-css-urls"], "WebfilterProfileFtgdWf-RateCssUrls"); ok {
			if err = d.Set("rate_css_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_css_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_css_urls: %v", err)
		}
	}

	if err = d.Set("rate_image_urls", flattenWebfilterProfileFtgdWfRateImageUrls2edl(o["rate-image-urls"], d, "rate_image_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-image-urls"], "WebfilterProfileFtgdWf-RateImageUrls"); ok {
			if err = d.Set("rate_image_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_image_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_image_urls: %v", err)
		}
	}

	if err = d.Set("rate_javascript_urls", flattenWebfilterProfileFtgdWfRateJavascriptUrls2edl(o["rate-javascript-urls"], d, "rate_javascript_urls")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-javascript-urls"], "WebfilterProfileFtgdWf-RateJavascriptUrls"); ok {
			if err = d.Set("rate_javascript_urls", vv); err != nil {
				return fmt.Errorf("Error reading rate_javascript_urls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_javascript_urls: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("risk", flattenWebfilterProfileFtgdWfRisk2edl(o["risk"], d, "risk")); err != nil {
			if vv, ok := fortiAPIPatch(o["risk"], "WebfilterProfileFtgdWf-Risk"); ok {
				if err = d.Set("risk", vv); err != nil {
					return fmt.Errorf("Error reading risk: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("risk"); ok {
			if err = d.Set("risk", flattenWebfilterProfileFtgdWfRisk2edl(o["risk"], d, "risk")); err != nil {
				if vv, ok := fortiAPIPatch(o["risk"], "WebfilterProfileFtgdWf-Risk"); ok {
					if err = d.Set("risk", vv); err != nil {
						return fmt.Errorf("Error reading risk: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading risk: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWebfilterProfileFtgdWfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileFtgdWfExemptQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWebfilterProfileFtgdWfFiltersAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-usr-grp"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(d, i["auth_usr_grp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandWebfilterProfileFtgdWfFiltersCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWebfilterProfileFtgdWfFiltersLog2edl(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warn-duration"], _ = expandWebfilterProfileFtgdWfFiltersWarnDuration2edl(d, i["warn_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-duration-type"], _ = expandWebfilterProfileFtgdWfFiltersWarningDurationType2edl(d, i["warning_duration_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-prompt"], _ = expandWebfilterProfileFtgdWfFiltersWarningPrompt2edl(d, i["warning_prompt"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarnDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningDurationType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningPrompt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfMaxQuotaTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfOvrd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandWebfilterProfileFtgdWfQuotaCategory2edl(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["duration"], _ = expandWebfilterProfileFtgdWfQuotaDuration2edl(d, i["duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfQuotaId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWebfilterProfileFtgdWfQuotaType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unit"], _ = expandWebfilterProfileFtgdWfQuotaUnit2edl(d, i["unit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandWebfilterProfileFtgdWfQuotaValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfQuotaCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfQuotaDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaUnit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCrlUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCssUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateImageUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateJavascriptUrls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRisk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWebfilterProfileFtgdWfRiskAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfRiskId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWebfilterProfileFtgdWfRiskLog2edl(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk_level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["risk-level"], _ = expandWebfilterProfileFtgdWfRiskRiskLevel2edl(d, i["risk_level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfRiskAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskRiskLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWebfilterProfileFtgdWf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("exempt_quota"); ok || d.HasChange("exempt_quota") {
		t, err := expandWebfilterProfileFtgdWfExemptQuota2edl(d, v, "exempt_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt-quota"] = t
		}
	}

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandWebfilterProfileFtgdWfFilters2edl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("max_quota_timeout"); ok || d.HasChange("max_quota_timeout") {
		t, err := expandWebfilterProfileFtgdWfMaxQuotaTimeout2edl(d, v, "max_quota_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-quota-timeout"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandWebfilterProfileFtgdWfOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("ovrd"); ok || d.HasChange("ovrd") {
		t, err := expandWebfilterProfileFtgdWfOvrd2edl(d, v, "ovrd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd"] = t
		}
	}

	if v, ok := d.GetOk("quota"); ok || d.HasChange("quota") {
		t, err := expandWebfilterProfileFtgdWfQuota2edl(d, v, "quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota"] = t
		}
	}

	if v, ok := d.GetOk("rate_crl_urls"); ok || d.HasChange("rate_crl_urls") {
		t, err := expandWebfilterProfileFtgdWfRateCrlUrls2edl(d, v, "rate_crl_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-crl-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_css_urls"); ok || d.HasChange("rate_css_urls") {
		t, err := expandWebfilterProfileFtgdWfRateCssUrls2edl(d, v, "rate_css_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-css-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_image_urls"); ok || d.HasChange("rate_image_urls") {
		t, err := expandWebfilterProfileFtgdWfRateImageUrls2edl(d, v, "rate_image_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-image-urls"] = t
		}
	}

	if v, ok := d.GetOk("rate_javascript_urls"); ok || d.HasChange("rate_javascript_urls") {
		t, err := expandWebfilterProfileFtgdWfRateJavascriptUrls2edl(d, v, "rate_javascript_urls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-javascript-urls"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {
		t, err := expandWebfilterProfileFtgdWfRisk2edl(d, v, "risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	return &obj, nil
}
