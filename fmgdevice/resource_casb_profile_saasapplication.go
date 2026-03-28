// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> CASB profile SaaS application.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfileSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileSaasApplicationCreate,
		Read:   resourceCasbProfileSaasApplicationRead,
		Update: resourceCasbProfileSaasApplicationUpdate,
		Delete: resourceCasbProfileSaasApplicationDelete,

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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceCasbProfileSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfileSaasApplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadCasbProfileSaasApplication(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateCasbProfileSaasApplication(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating CasbProfileSaasApplication resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateCasbProfileSaasApplication(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating CasbProfileSaasApplication resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationRead(d, m)
}

func resourceCasbProfileSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfileSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateCasbProfileSaasApplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfileSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceCasbProfileSaasApplicationRead(d, m)
}

func resourceCasbProfileSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteCasbProfileSaasApplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfileSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbProfileSaasApplication(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading CasbProfileSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfileSaasApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfileSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileSaasApplicationAccessRule2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAccessRuleAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := i["attribute-filter"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := i["bypass"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleBypass2edl(i["bypass"], d, pre_append)
			tmp["bypass"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Bypass")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AccessRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAccessRule-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleBypass2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAccessRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationAdvancedTenantControl2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(i["attribute"], d, pre_append)
			tmp["attribute"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AdvancedTenantControl-Attribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-AdvancedTenantControl-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(i["input"], d, pre_append)
			tmp["input"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAdvancedTenantControl-Attribute-Input")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationAdvancedTenantControl-Attribute-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationCustomControl2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilter2edl(i["attribute-filter"], d, pre_append)
			tmp["attribute_filter"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-AttributeFilter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := i["option"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOption2edl(i["option"], d, pre_append)
			tmp["option"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplication-CustomControl-Option")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := i["attribute-match"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(i["attribute-match"], d, pre_append)
			tmp["attribute_match"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-AttributeMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-AttributeFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationCustomControlOption2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenCasbProfileSaasApplicationCustomControlOptionName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := i["user-input"]; ok {
			v := flattenCasbProfileSaasApplicationCustomControlOptionUserInput2edl(i["user-input"], d, pre_append)
			tmp["user_input"] = fortiAPISubPartPatch(v, "CasbProfileSaasApplicationCustomControl-Option-UserInput")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenCasbProfileSaasApplicationCustomControlOptionName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationDomainControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationDomainControlDomains2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenCasbProfileSaasApplicationSafeSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationSafeSearchControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenCasbProfileSaasApplicationStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControlTenants2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectCasbProfileSaasApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("access_rule", flattenCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["access-rule"], "CasbProfileSaasApplication-AccessRule"); ok {
				if err = d.Set("access_rule", vv); err != nil {
					return fmt.Errorf("Error reading access_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading access_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_rule"); ok {
			if err = d.Set("access_rule", flattenCasbProfileSaasApplicationAccessRule2edl(o["access-rule"], d, "access_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["access-rule"], "CasbProfileSaasApplication-AccessRule"); ok {
					if err = d.Set("access_rule", vv); err != nil {
						return fmt.Errorf("Error reading access_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading access_rule: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("advanced_tenant_control", flattenCasbProfileSaasApplicationAdvancedTenantControl2edl(o["advanced-tenant-control"], d, "advanced_tenant_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["advanced-tenant-control"], "CasbProfileSaasApplication-AdvancedTenantControl"); ok {
				if err = d.Set("advanced_tenant_control", vv); err != nil {
					return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("advanced_tenant_control"); ok {
			if err = d.Set("advanced_tenant_control", flattenCasbProfileSaasApplicationAdvancedTenantControl2edl(o["advanced-tenant-control"], d, "advanced_tenant_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["advanced-tenant-control"], "CasbProfileSaasApplication-AdvancedTenantControl"); ok {
					if err = d.Set("advanced_tenant_control", vv); err != nil {
						return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading advanced_tenant_control: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("custom_control", flattenCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-control"], "CasbProfileSaasApplication-CustomControl"); ok {
				if err = d.Set("custom_control", vv); err != nil {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_control"); ok {
			if err = d.Set("custom_control", flattenCasbProfileSaasApplicationCustomControl2edl(o["custom-control"], d, "custom_control")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-control"], "CasbProfileSaasApplication-CustomControl"); ok {
					if err = d.Set("custom_control", vv); err != nil {
						return fmt.Errorf("Error reading custom_control: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_control: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain_control", flattenCasbProfileSaasApplicationDomainControl2edl(o["domain-control"], d, "domain_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control"], "CasbProfileSaasApplication-DomainControl"); ok {
			if err = d.Set("domain_control", vv); err != nil {
				return fmt.Errorf("Error reading domain_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control: %v", err)
		}
	}

	if err = d.Set("domain_control_domains", flattenCasbProfileSaasApplicationDomainControlDomains2edl(o["domain-control-domains"], d, "domain_control_domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-control-domains"], "CasbProfileSaasApplication-DomainControlDomains"); ok {
			if err = d.Set("domain_control_domains", vv); err != nil {
				return fmt.Errorf("Error reading domain_control_domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_control_domains: %v", err)
		}
	}

	if err = d.Set("log", flattenCasbProfileSaasApplicationLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "CasbProfileSaasApplication-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenCasbProfileSaasApplicationName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "CasbProfileSaasApplication-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenCasbProfileSaasApplicationSafeSearch2edl(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "CasbProfileSaasApplication-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("safe_search_control", flattenCasbProfileSaasApplicationSafeSearchControl2edl(o["safe-search-control"], d, "safe_search_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search-control"], "CasbProfileSaasApplication-SafeSearchControl"); ok {
			if err = d.Set("safe_search_control", vv); err != nil {
				return fmt.Errorf("Error reading safe_search_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search_control: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbProfileSaasApplicationStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "CasbProfileSaasApplication-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tenant_control", flattenCasbProfileSaasApplicationTenantControl2edl(o["tenant-control"], d, "tenant_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control"], "CasbProfileSaasApplication-TenantControl"); ok {
			if err = d.Set("tenant_control", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control: %v", err)
		}
	}

	if err = d.Set("tenant_control_tenants", flattenCasbProfileSaasApplicationTenantControlTenants2edl(o["tenant-control-tenants"], d, "tenant_control_tenants")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-control-tenants"], "CasbProfileSaasApplication-TenantControlTenants"); ok {
			if err = d.Set("tenant_control_tenants", vv); err != nil {
				return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_control_tenants: %v", err)
		}
	}

	return nil
}

func flattenCasbProfileSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCasbProfileSaasApplicationAccessRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bypass"], _ = expandCasbProfileSaasApplicationAccessRuleBypass2edl(d, i["bypass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAccessRuleName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleBypass2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAccessRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(d, i["attribute"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttribute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["input"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(d, i["input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandCasbProfileSaasApplicationCustomControlAttributeFilter2edl(d, i["attribute_filter"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["attribute-filter"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandCasbProfileSaasApplicationCustomControlOption2edl(d, i["option"], pre_append)
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

func expandCasbProfileSaasApplicationCustomControlAttributeFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(d, i["attribute_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlOptionName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-input"], _ = expandCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d, i["user_input"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationDomainControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationDomainControlDomains2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearchControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandCasbProfileSaasApplicationStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControlTenants2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectCasbProfileSaasApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_rule"); ok || d.HasChange("access_rule") {
		t, err := expandCasbProfileSaasApplicationAccessRule2edl(d, v, "access_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-rule"] = t
		}
	}

	if v, ok := d.GetOk("advanced_tenant_control"); ok || d.HasChange("advanced_tenant_control") {
		t, err := expandCasbProfileSaasApplicationAdvancedTenantControl2edl(d, v, "advanced_tenant_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advanced-tenant-control"] = t
		}
	}

	if v, ok := d.GetOk("custom_control"); ok || d.HasChange("custom_control") {
		t, err := expandCasbProfileSaasApplicationCustomControl2edl(d, v, "custom_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control"); ok || d.HasChange("domain_control") {
		t, err := expandCasbProfileSaasApplicationDomainControl2edl(d, v, "domain_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control"] = t
		}
	}

	if v, ok := d.GetOk("domain_control_domains"); ok || d.HasChange("domain_control_domains") {
		t, err := expandCasbProfileSaasApplicationDomainControlDomains2edl(d, v, "domain_control_domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-control-domains"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandCasbProfileSaasApplicationLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandCasbProfileSaasApplicationName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandCasbProfileSaasApplicationSafeSearch2edl(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("safe_search_control"); ok || d.HasChange("safe_search_control") {
		t, err := expandCasbProfileSaasApplicationSafeSearchControl2edl(d, v, "safe_search_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search-control"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandCasbProfileSaasApplicationStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control"); ok || d.HasChange("tenant_control") {
		t, err := expandCasbProfileSaasApplicationTenantControl2edl(d, v, "tenant_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control"] = t
		}
	}

	if v, ok := d.GetOk("tenant_control_tenants"); ok || d.HasChange("tenant_control_tenants") {
		t, err := expandCasbProfileSaasApplicationTenantControlTenants2edl(d, v, "tenant_control_tenants")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-control-tenants"] = t
		}
	}

	return &obj, nil
}
