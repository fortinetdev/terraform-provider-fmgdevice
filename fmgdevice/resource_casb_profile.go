// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure CASB profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileCreate,
		Read:   resourceCasbProfileRead,
		Update: resourceCasbProfileUpdate,
		Delete: resourceCasbProfileDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"saas_application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_rule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
										Optional: true,
									},
								},
							},
						},
						"advanced_tenant_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"input": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"custom_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
								},
							},
						},
						"domain_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_control_domains": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"safe_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"safe_search_control": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant_control_tenants": &schema.Schema{
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

func resourceCasbProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileRead(d, m)
}

func resourceCasbProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileRead(d, m)
}

func resourceCasbProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfile resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if _, ok := i["access-rule"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRule(i["access-rule"], d, pre_append)
			tmp["access_rule"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-AccessRule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if _, ok := i["advanced-tenant-control"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControl(i["advanced-tenant-control"], d, pre_append)
			tmp["advanced_tenant_control"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-AdvancedTenantControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := i["custom-control"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControl(i["custom-control"], d, pre_append)
			tmp["custom_control"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-CustomControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := i["domain-control"]; ok {
			v := flattenCasbProfileSaasApplicationDomainControl(i["domain-control"], d, pre_append)
			tmp["domain_control"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-DomainControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := i["domain-control-domains"]; ok {
			v := flattenCasbProfileSaasApplicationDomainControlDomains(i["domain-control-domains"], d, pre_append)
			tmp["domain_control_domains"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-DomainControlDomains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenCasbProfileSaasApplicationLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := i["safe-search"]; ok {
			v := flattenCasbProfileSaasApplicationSafeSearch(i["safe-search"], d, pre_append)
			tmp["safe_search"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-SafeSearch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := i["safe-search-control"]; ok {
			v := flattenCasbProfileSaasApplicationSafeSearchControl(i["safe-search-control"], d, pre_append)
			tmp["safe_search_control"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-SafeSearchControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenCasbProfileSaasApplicationStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := i["tenant-control"]; ok {
			v := flattenCasbProfileSaasApplicationTenantControl(i["tenant-control"], d, pre_append)
			tmp["tenant_control"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-TenantControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := i["tenant-control-tenants"]; ok {
			v := flattenCasbProfileSaasApplicationTenantControlTenants(i["tenant-control-tenants"], d, pre_append)
			tmp["tenant_control_tenants"] = fortiAPISubPartPatch(v, "CasbProfile-SaasApplication-TenantControlTenants")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAccessRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilter(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleBypass(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationAdvancedTenantControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := i["attribute"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute(i["attribute"], d, pre_append)
			tmp["attribute"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AdvancedTenantControl-Attribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AdvancedTenantControl-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := i["input"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(i["input"], d, pre_append)
			tmp["input"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAdvancedTenantControl-Attribute-Input")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAdvancedTenantControl-Attribute-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationCustomControl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilter(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOption(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-Option")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationCustomControlOption(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlOptionName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOptionUserInput(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlOptionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationDomainControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationDomainControlDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationSafeSearchControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControlTenants(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectCasbProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenCasbProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "CasbProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("saas_application", flattenCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
			if vv, ok := fortiAPIPatch(o["saas-application"], "CasbProfile-SaasApplication"); ok {
				if err = d.Set("saas_application", vv); err != nil {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading saas_application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("saas_application"); ok {
			if err = d.Set("saas_application", flattenCasbProfileSaasApplication(o["saas-application"], d, "saas_application")); err != nil {
				if vv, ok := fortiAPIPatch(o["saas-application"], "CasbProfile-SaasApplication"); ok {
					if err = d.Set("saas_application", vv); err != nil {
						return fmt.Errorf("Error reading saas_application: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationAccessRule(d, i["access_rule"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["access-rule"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationAdvancedTenantControl(d, i["advanced_tenant_control"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["advanced-tenant-control"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationCustomControl(d, i["custom_control"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["custom-control"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control"], _ = expandCasbProfileSaasApplicationDomainControl(d, i["domain_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain-control-domains"], _ = expandCasbProfileSaasApplicationDomainControlDomains(d, i["domain_control_domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandCasbProfileSaasApplicationLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search"], _ = expandCasbProfileSaasApplicationSafeSearch(d, i["safe_search"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["safe-search-control"], _ = expandCasbProfileSaasApplicationSafeSearchControl(d, i["safe_search_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandCasbProfileSaasApplicationStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control"], _ = expandCasbProfileSaasApplicationTenantControl(d, i["tenant_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant-control-tenants"], _ = expandCasbProfileSaasApplicationTenantControlTenants(d, i["tenant_control_tenants"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilter(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandCasbProfileSaasApplicationAccessRuleBypass(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAccessRuleName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationAdvancedTenantControlAttribute(d, i["attribute"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d, i["input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationCustomControlAttributeFilter(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationCustomControlOption(d, i["option"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["option"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlOptionName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandCasbProfileSaasApplicationCustomControlOptionUserInput(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationDomainControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationDomainControlDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearchControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControlTenants(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectCasbProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandCasbProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("saas_application"); ok || d.HasChange("saas_application") {
		t, err := expandCasbProfileSaasApplication(d, v, "saas_application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saas-application"] = t
		}
	}

	return &obj, nil
}
