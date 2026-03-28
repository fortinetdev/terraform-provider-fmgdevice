// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB user activity match rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivityMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityMatchCreate,
		Read:   resourceCasbUserActivityMatchRead,
		Update: resourceCasbUserActivityMatchUpdate,
		Delete: resourceCasbUserActivityMatchDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domains": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"jq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"methods": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"negate": &schema.Schema{
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
			"strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_extraction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"body_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"header_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"place": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"jq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceCasbUserActivityMatchCreate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity

	obj, err := getObjectCasbUserActivityMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivityMatch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbUserActivityMatch(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbUserActivityMatch(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbUserActivityMatch resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateCasbUserActivityMatch(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbUserActivityMatch resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceCasbUserActivityMatchRead(d, m)
			} else {
				return fmt.Errorf("Error creating CasbUserActivityMatch resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchRead(d, m)
}

func resourceCasbUserActivityMatchUpdate(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity

	obj, err := getObjectCasbUserActivityMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivityMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivityMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceCasbUserActivityMatchRead(d, m)
}

func resourceCasbUserActivityMatchDelete(d *schema.ResourceData, m interface{}) error {
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
	user_activity := d.Get("user_activity").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity

	wsParams["adom"] = adomv

	err = c.DeleteCasbUserActivityMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivityMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityMatchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	user_activity := d.Get("user_activity").(string)
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
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["user_activity"] = user_activity

	o, err := c.ReadCasbUserActivityMatch(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivityMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivityMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivityMatch resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityMatchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRules2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenCasbUserActivityMatchRulesBodyType2edl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenCasbUserActivityMatchRulesCaseSensitive2edl(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := i["domains"]; ok {
			v := flattenCasbUserActivityMatchRulesDomains2edl(i["domains"], d, pre_append)
			tmp["domains"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Domains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchRulesHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchRulesId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := i["jq"]; ok {
			v := flattenCasbUserActivityMatchRulesJq2edl(i["jq"], d, pre_append)
			tmp["jq"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Jq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := i["match-pattern"]; ok {
			v := flattenCasbUserActivityMatchRulesMatchPattern2edl(i["match-pattern"], d, pre_append)
			tmp["match_pattern"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-MatchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := i["match-value"]; ok {
			v := flattenCasbUserActivityMatchRulesMatchValue2edl(i["match-value"], d, pre_append)
			tmp["match_value"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-MatchValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := i["methods"]; ok {
			v := flattenCasbUserActivityMatchRulesMethods2edl(i["methods"], d, pre_append)
			tmp["methods"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Methods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenCasbUserActivityMatchRulesNegate2edl(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenCasbUserActivityMatchRulesType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchRulesBodyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesJq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMethods2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchStrategy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtraction2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenCasbUserActivityMatchTenantExtractionFilters2edl(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "jq"
	if _, ok := i["jq"]; ok {
		result["jq"] = flattenCasbUserActivityMatchTenantExtractionJq2edl(i["jq"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenCasbUserActivityMatchTenantExtractionStatus2edl(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenCasbUserActivityMatchTenantExtractionType2edl(i["type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenCasbUserActivityMatchTenantExtractionFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersPlace2edl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersPlace2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionJq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivityMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fosid", flattenCasbUserActivityMatchId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "CasbUserActivityMatch-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rules", flattenCasbUserActivityMatchRules2edl(o["rules"], d, "rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["rules"], "CasbUserActivityMatch-Rules"); ok {
				if err = d.Set("rules", vv); err != nil {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rules"); ok {
			if err = d.Set("rules", flattenCasbUserActivityMatchRules2edl(o["rules"], d, "rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["rules"], "CasbUserActivityMatch-Rules"); ok {
					if err = d.Set("rules", vv); err != nil {
						return fmt.Errorf("Error reading rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("strategy", flattenCasbUserActivityMatchStrategy2edl(o["strategy"], d, "strategy")); err != nil {
		if vv, ok := fortiAPIPatch(o["strategy"], "CasbUserActivityMatch-Strategy"); ok {
			if err = d.Set("strategy", vv); err != nil {
				return fmt.Errorf("Error reading strategy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strategy: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tenant_extraction", flattenCasbUserActivityMatchTenantExtraction2edl(o["tenant-extraction"], d, "tenant_extraction")); err != nil {
			if vv, ok := fortiAPIPatch(o["tenant-extraction"], "CasbUserActivityMatch-TenantExtraction"); ok {
				if err = d.Set("tenant_extraction", vv); err != nil {
					return fmt.Errorf("Error reading tenant_extraction: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tenant_extraction: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tenant_extraction"); ok {
			if err = d.Set("tenant_extraction", flattenCasbUserActivityMatchTenantExtraction2edl(o["tenant-extraction"], d, "tenant_extraction")); err != nil {
				if vv, ok := fortiAPIPatch(o["tenant-extraction"], "CasbUserActivityMatch-TenantExtraction"); ok {
					if err = d.Set("tenant_extraction", vv); err != nil {
						return fmt.Errorf("Error reading tenant_extraction: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tenant_extraction: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbUserActivityMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityMatchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRules2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandCasbUserActivityMatchRulesBodyType2edl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandCasbUserActivityMatchRulesCaseSensitive2edl(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domains"], _ = expandCasbUserActivityMatchRulesDomains2edl(d, i["domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchRulesHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchRulesId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jq"], _ = expandCasbUserActivityMatchRulesJq2edl(d, i["jq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-pattern"], _ = expandCasbUserActivityMatchRulesMatchPattern2edl(d, i["match_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-value"], _ = expandCasbUserActivityMatchRulesMatchValue2edl(d, i["match_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["methods"], _ = expandCasbUserActivityMatchRulesMethods2edl(d, i["methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandCasbUserActivityMatchRulesNegate2edl(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandCasbUserActivityMatchRulesType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchRulesBodyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesJq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMethods2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchStrategy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtraction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandCasbUserActivityMatchTenantExtractionFilters2edl(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}
	pre_append = pre + ".0." + "jq"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["jq"], _ = expandCasbUserActivityMatchTenantExtractionJq2edl(d, i["jq"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandCasbUserActivityMatchTenantExtractionStatus2edl(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["type"], _ = expandCasbUserActivityMatchTenantExtractionType2edl(d, i["type"], pre_append)
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandCasbUserActivityMatchTenantExtractionFiltersDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchTenantExtractionFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandCasbUserActivityMatchTenantExtractionFiltersPlace2edl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersBodyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersPlace2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionJq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivityMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandCasbUserActivityMatchId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("rules"); ok || d.HasChange("rules") {
		t, err := expandCasbUserActivityMatchRules2edl(d, v, "rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rules"] = t
		}
	}

	if v, ok := d.GetOk("strategy"); ok || d.HasChange("strategy") {
		t, err := expandCasbUserActivityMatchStrategy2edl(d, v, "strategy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strategy"] = t
		}
	}

	if v, ok := d.GetOk("tenant_extraction"); ok || d.HasChange("tenant_extraction") {
		t, err := expandCasbUserActivityMatchTenantExtraction2edl(d, v, "tenant_extraction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-extraction"] = t
		}
	}

	return &obj, nil
}
