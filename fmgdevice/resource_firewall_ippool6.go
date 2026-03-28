// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 IP pools.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIppool6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIppool6Create,
		Read:   resourceFirewallIppool6Read,
		Update: resourceFirewallIppool6Update,
		Delete: resourceFirewallIppool6Delete,

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
			"add_nat46_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"add_nat46_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"internal_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nat46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"startip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
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

func resourceFirewallIppool6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIppool6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallIppool6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallIppool6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallIppool6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallIppool6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallIppool6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallIppool6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallIppool6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIppool6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppool6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIppool6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallIppool6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIppool6AddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallIppool6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := i["add-nat46-route"]; ok {
			v := flattenFirewallIppool6DynamicMappingAddNat46Route(i["add-nat46-route"], d, pre_append)
			tmp["add_nat46_route"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-AddNat46Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenFirewallIppool6DynamicMappingComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := i["endip"]; ok {
			v := flattenFirewallIppool6DynamicMappingEndip(i["endip"], d, pre_append)
			tmp["endip"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Endip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_prefix"
		if _, ok := i["external-prefix"]; ok {
			v := flattenFirewallIppool6DynamicMappingExternalPrefix(i["external-prefix"], d, pre_append)
			tmp["external_prefix"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-ExternalPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internal_prefix"
		if _, ok := i["internal-prefix"]; ok {
			v := flattenFirewallIppool6DynamicMappingInternalPrefix(i["internal-prefix"], d, pre_append)
			tmp["internal_prefix"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-InternalPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := i["nat46"]; ok {
			v := flattenFirewallIppool6DynamicMappingNat46(i["nat46"], d, pre_append)
			tmp["nat46"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Nat46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := i["startip"]; ok {
			v := flattenFirewallIppool6DynamicMappingStartip(i["startip"], d, pre_append)
			tmp["startip"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Startip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallIppool6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallIppool6-DynamicMapping-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallIppool6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallIppool6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallIppool6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallIppool6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallIppool6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallIppool6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingAddNat46Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingExternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingInternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Endip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6ExternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6InternalPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Nat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Startip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIppool6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat46_route", flattenFirewallIppool6AddNat46Route(o["add-nat46-route"], d, "add_nat46_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat46-route"], "FirewallIppool6-AddNat46Route"); ok {
			if err = d.Set("add_nat46_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat46_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat46_route: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppool6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallIppool6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallIppool6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallIppool6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallIppool6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallIppool6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("endip", flattenFirewallIppool6Endip(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "FirewallIppool6-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("external_prefix", flattenFirewallIppool6ExternalPrefix(o["external-prefix"], d, "external_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-prefix"], "FirewallIppool6-ExternalPrefix"); ok {
			if err = d.Set("external_prefix", vv); err != nil {
				return fmt.Errorf("Error reading external_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_prefix: %v", err)
		}
	}

	if err = d.Set("internal_prefix", flattenFirewallIppool6InternalPrefix(o["internal-prefix"], d, "internal_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-prefix"], "FirewallIppool6-InternalPrefix"); ok {
			if err = d.Set("internal_prefix", vv); err != nil {
				return fmt.Errorf("Error reading internal_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_prefix: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallIppool6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallIppool6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallIppool6Nat46(o["nat46"], d, "nat46")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46"], "FirewallIppool6-Nat46"); ok {
			if err = d.Set("nat46", vv); err != nil {
				return fmt.Errorf("Error reading nat46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppool6Startip(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "FirewallIppool6-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallIppool6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallIppool6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenFirewallIppool6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIppool6AddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallIppool6DynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat46_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat46-route"], _ = expandFirewallIppool6DynamicMappingAddNat46Route(d, i["add_nat46_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandFirewallIppool6DynamicMappingComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endip"], _ = expandFirewallIppool6DynamicMappingEndip(d, i["endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["external-prefix"], _ = expandFirewallIppool6DynamicMappingExternalPrefix(d, i["external_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internal_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internal-prefix"], _ = expandFirewallIppool6DynamicMappingInternalPrefix(d, i["internal_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat46"], _ = expandFirewallIppool6DynamicMappingNat46(d, i["nat46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startip"], _ = expandFirewallIppool6DynamicMappingStartip(d, i["startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallIppool6DynamicMappingType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallIppool6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallIppool6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallIppool6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallIppool6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingAddNat46Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingExternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingInternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingNat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Endip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6ExternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6InternalPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Nat46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Startip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIppool6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat46_route"); ok || d.HasChange("add_nat46_route") {
		t, err := expandFirewallIppool6AddNat46Route(d, v, "add_nat46_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat46-route"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallIppool6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallIppool6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandFirewallIppool6Endip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("external_prefix"); ok || d.HasChange("external_prefix") {
		t, err := expandFirewallIppool6ExternalPrefix(d, v, "external_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-prefix"] = t
		}
	}

	if v, ok := d.GetOk("internal_prefix"); ok || d.HasChange("internal_prefix") {
		t, err := expandFirewallIppool6InternalPrefix(d, v, "internal_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-prefix"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallIppool6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok || d.HasChange("nat46") {
		t, err := expandFirewallIppool6Nat46(d, v, "nat46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandFirewallIppool6Startip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallIppool6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
