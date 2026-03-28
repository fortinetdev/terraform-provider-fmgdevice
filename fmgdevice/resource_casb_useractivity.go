// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure CASB user activity.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityCreate,
		Read:   resourceCasbUserActivityRead,
		Update: resourceCasbUserActivityUpdate,
		Delete: resourceCasbUserActivityDelete,

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
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"casb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"control_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"operations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"case_sensitive": &schema.Schema{
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"search_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"search_pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_from_input": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_name_from_input": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"values": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"match_strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceCasbUserActivityCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbUserActivity(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivity resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbUserActivity(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbUserActivity(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbUserActivity resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbUserActivity(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbUserActivity resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbUserActivityRead(d, m)
}

func resourceCasbUserActivityUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbUserActivity(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivity resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbUserActivity(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbUserActivityRead(d, m)
}

func resourceCasbUserActivityDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbUserActivity(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbUserActivity(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbUserActivity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivity resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityCasbName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbUserActivityControlOptionsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbUserActivity-ControlOptions-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "operations"
		if _, ok := i["operations"]; ok {
			v := flattenCasbUserActivityControlOptionsOperations(i["operations"], d, pre_append)
			tmp["operations"] = fortiAPISubPartPatch(v, "CasbUserActivity-ControlOptions-Operations")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenCasbUserActivityControlOptionsStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "CasbUserActivity-ControlOptions-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityControlOptionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperations(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbUserActivityControlOptionsOperationsAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsCaseSensitive(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
		if _, ok := i["search-key"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsSearchKey(i["search-key"], d, pre_append)
			tmp["search_key"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-SearchKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
		if _, ok := i["search-pattern"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsSearchPattern(i["search-pattern"], d, pre_append)
			tmp["search_pattern"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-SearchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsTarget(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-Target")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
		if _, ok := i["value-from-input"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsValueFromInput(i["value-from-input"], d, pre_append)
			tmp["value_from_input"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-ValueFromInput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_name_from_input"
		if _, ok := i["value-name-from-input"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsValueNameFromInput(i["value-name-from-input"], d, pre_append)
			tmp["value_name_from_input"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-ValueNameFromInput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			v := flattenCasbUserActivityControlOptionsOperationsValues(i["values"], d, pre_append)
			tmp["values"] = fortiAPISubPartPatch(v, "CasbUserActivityControlOptions-Operations-Values")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityControlOptionsOperationsAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsSearchKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsSearchPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValueFromInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValueNameFromInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValues(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityControlOptionsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivity-Match-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rules"
		if _, ok := i["rules"]; ok {
			v := flattenCasbUserActivityMatchRules(i["rules"], d, pre_append)
			tmp["rules"] = fortiAPISubPartPatch(v, "CasbUserActivity-Match-Rules")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strategy"
		if _, ok := i["strategy"]; ok {
			v := flattenCasbUserActivityMatchStrategy(i["strategy"], d, pre_append)
			tmp["strategy"] = fortiAPISubPartPatch(v, "CasbUserActivity-Match-Strategy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_extraction"
		if _, ok := i["tenant-extraction"]; ok {
			v := flattenCasbUserActivityMatchTenantExtraction(i["tenant-extraction"], d, pre_append)
			tmp["tenant_extraction"] = fortiAPISubPartPatch(v, "CasbUserActivity-Match-TenantExtraction")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbUserActivityMatchRulesBodyType(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenCasbUserActivityMatchRulesCaseSensitive(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := i["domains"]; ok {
			v := flattenCasbUserActivityMatchRulesDomains(i["domains"], d, pre_append)
			tmp["domains"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Domains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchRulesHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchRulesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := i["jq"]; ok {
			v := flattenCasbUserActivityMatchRulesJq(i["jq"], d, pre_append)
			tmp["jq"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Jq")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := i["match-pattern"]; ok {
			v := flattenCasbUserActivityMatchRulesMatchPattern(i["match-pattern"], d, pre_append)
			tmp["match_pattern"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-MatchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := i["match-value"]; ok {
			v := flattenCasbUserActivityMatchRulesMatchValue(i["match-value"], d, pre_append)
			tmp["match_value"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-MatchValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := i["methods"]; ok {
			v := flattenCasbUserActivityMatchRulesMethods(i["methods"], d, pre_append)
			tmp["methods"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Methods")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			v := flattenCasbUserActivityMatchRulesNegate(i["negate"], d, pre_append)
			tmp["negate"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Negate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenCasbUserActivityMatchRulesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatch-Rules-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchRulesBodyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesJq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbUserActivityMatchRulesNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchStrategy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtraction(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenCasbUserActivityMatchTenantExtractionFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "jq"
	if _, ok := i["jq"]; ok {
		result["jq"] = flattenCasbUserActivityMatchTenantExtractionJq(i["jq"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenCasbUserActivityMatchTenantExtractionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenCasbUserActivityMatchTenantExtractionType(i["type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenCasbUserActivityMatchTenantExtractionFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbUserActivityMatchTenantExtractionFiltersBodyType(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenCasbUserActivityMatchTenantExtractionFiltersPlace(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "CasbUserActivityMatchTenantExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbUserActivityMatchTenantExtractionFiltersBodyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersPlace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionJq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityMatchStrategyU(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbUserActivityUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCasbUserActivity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("application", flattenCasbUserActivityApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "CasbUserActivity-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("casb_name", flattenCasbUserActivityCasbName(o["casb-name"], d, "casb_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-name"], "CasbUserActivity-CasbName"); ok {
			if err = d.Set("casb_name", vv); err != nil {
				return fmt.Errorf("Error reading casb_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_name: %v", err)
		}
	}

	if err = d.Set("category", flattenCasbUserActivityCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "CasbUserActivity-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("control_options", flattenCasbUserActivityControlOptions(o["control-options"], d, "control_options")); err != nil {
			if vv, ok := fortiAPIPatch(o["control-options"], "CasbUserActivity-ControlOptions"); ok {
				if err = d.Set("control_options", vv); err != nil {
					return fmt.Errorf("Error reading control_options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading control_options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("control_options"); ok {
			if err = d.Set("control_options", flattenCasbUserActivityControlOptions(o["control-options"], d, "control_options")); err != nil {
				if vv, ok := fortiAPIPatch(o["control-options"], "CasbUserActivity-ControlOptions"); ok {
					if err = d.Set("control_options", vv); err != nil {
						return fmt.Errorf("Error reading control_options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading control_options: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenCasbUserActivityDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "CasbUserActivity-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("match", flattenCasbUserActivityMatch(o["match"], d, "match")); err != nil {
			if vv, ok := fortiAPIPatch(o["match"], "CasbUserActivity-Match"); ok {
				if err = d.Set("match", vv); err != nil {
					return fmt.Errorf("Error reading match: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenCasbUserActivityMatch(o["match"], d, "match")); err != nil {
				if vv, ok := fortiAPIPatch(o["match"], "CasbUserActivity-Match"); ok {
					if err = d.Set("match", vv); err != nil {
						return fmt.Errorf("Error reading match: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if err = d.Set("match_strategy", flattenCasbUserActivityMatchStrategyU(o["match-strategy"], d, "match_strategy")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-strategy"], "CasbUserActivity-MatchStrategy"); ok {
			if err = d.Set("match_strategy", vv); err != nil {
				return fmt.Errorf("Error reading match_strategy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_strategy: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbUserActivityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbUserActivity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbUserActivityStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "CasbUserActivity-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbUserActivityType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "CasbUserActivity-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenCasbUserActivityUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "CasbUserActivity-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenCasbUserActivityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbUserActivityApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityCasbName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbUserActivityControlOptionsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "operations"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbUserActivityControlOptionsOperations(d, i["operations"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["operations"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandCasbUserActivityControlOptionsStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityControlOptionsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperations(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbUserActivityControlOptionsOperationsAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandCasbUserActivityControlOptionsOperationsCaseSensitive(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandCasbUserActivityControlOptionsOperationsDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityControlOptionsOperationsHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbUserActivityControlOptionsOperationsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-key"], _ = expandCasbUserActivityControlOptionsOperationsSearchKey(d, i["search_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-pattern"], _ = expandCasbUserActivityControlOptionsOperationsSearchPattern(d, i["search_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandCasbUserActivityControlOptionsOperationsTarget(d, i["target"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value-from-input"], _ = expandCasbUserActivityControlOptionsOperationsValueFromInput(d, i["value_from_input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_name_from_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value-name-from-input"], _ = expandCasbUserActivityControlOptionsOperationsValueNameFromInput(d, i["value_name_from_input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["values"], _ = expandCasbUserActivityControlOptionsOperationsValues(d, i["values"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityControlOptionsOperationsAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsSearchKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsSearchPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValueFromInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValueNameFromInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityControlOptionsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rules"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbUserActivityMatchRules(d, i["rules"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["rules"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strategy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["strategy"], _ = expandCasbUserActivityMatchStrategy(d, i["strategy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_extraction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbUserActivityMatchTenantExtraction(d, i["tenant_extraction"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["tenant-extraction"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["body-type"], _ = expandCasbUserActivityMatchRulesBodyType(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandCasbUserActivityMatchRulesCaseSensitive(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domains"], _ = expandCasbUserActivityMatchRulesDomains(d, i["domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchRulesHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchRulesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jq"], _ = expandCasbUserActivityMatchRulesJq(d, i["jq"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-pattern"], _ = expandCasbUserActivityMatchRulesMatchPattern(d, i["match_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-value"], _ = expandCasbUserActivityMatchRulesMatchValue(d, i["match_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["methods"], _ = expandCasbUserActivityMatchRulesMethods(d, i["methods"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["negate"], _ = expandCasbUserActivityMatchRulesNegate(d, i["negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandCasbUserActivityMatchRulesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchRulesBodyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesJq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbUserActivityMatchRulesNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchStrategy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtraction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandCasbUserActivityMatchTenantExtractionFilters(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}
	pre_append = pre + ".0." + "jq"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["jq"], _ = expandCasbUserActivityMatchTenantExtractionJq(d, i["jq"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandCasbUserActivityMatchTenantExtractionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["type"], _ = expandCasbUserActivityMatchTenantExtractionType(d, i["type"], pre_append)
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["body-type"], _ = expandCasbUserActivityMatchTenantExtractionFiltersBodyType(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandCasbUserActivityMatchTenantExtractionFiltersDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandCasbUserActivityMatchTenantExtractionFiltersHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbUserActivityMatchTenantExtractionFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandCasbUserActivityMatchTenantExtractionFiltersPlace(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersBodyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersPlace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionJq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchStrategyU(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandCasbUserActivityApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("casb_name"); ok || d.HasChange("casb_name") {
		t, err := expandCasbUserActivityCasbName(d, v, "casb_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-name"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandCasbUserActivityCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("control_options"); ok || d.HasChange("control_options") {
		t, err := expandCasbUserActivityControlOptions(d, v, "control_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-options"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandCasbUserActivityDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandCasbUserActivityMatch(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("match_strategy"); ok || d.HasChange("match_strategy") {
		t, err := expandCasbUserActivityMatchStrategyU(d, v, "match_strategy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-strategy"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbUserActivityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandCasbUserActivityStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandCasbUserActivityType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandCasbUserActivityUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
