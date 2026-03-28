// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 address groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddrgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddrgrpCreate,
		Read:   resourceFirewallAddrgrpRead,
		Update: resourceFirewallAddrgrpUpdate,
		Delete: resourceFirewallAddrgrpDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_image_base64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
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
						"allow_routing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude_member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"global_object": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"visibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"exclude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exclude_member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tagging": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
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

func resourceFirewallAddrgrpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddrgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddrgrp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddrgrp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddrgrp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddrgrp resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddrgrp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddrgrp resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddrgrpRead(d, m)
}

func resourceFirewallAddrgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddrgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddrgrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddrgrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddrgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddrgrpRead(d, m)
}

func resourceFirewallAddrgrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAddrgrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddrgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddrgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAddrgrp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddrgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddrgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddrgrp resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddrgrpImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := i["_image-base64"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingImageBase64(i["_image-base64"], d, pre_append)
			tmp["_image_base64"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-ImageBase64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := i["allow-routing"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingAllowRouting(i["allow-routing"], d, pre_append)
			tmp["allow_routing"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-AllowRouting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude"
		if _, ok := i["exclude"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingExclude(i["exclude"], d, pre_append)
			tmp["exclude"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Exclude")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_member"
		if _, ok := i["exclude-member"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingExcludeMember(i["exclude-member"], d, pre_append)
			tmp["exclude_member"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-ExcludeMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := i["global-object"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingGlobalObject(i["global-object"], d, pre_append)
			tmp["global_object"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-GlobalObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := i["visibility"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingVisibility(i["visibility"], d, pre_append)
			tmp["visibility"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-DynamicMapping-Visibility")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddrgrpDynamicMappingImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddrgrpDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddrgrpDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallAddrgrpDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallAddrgrpDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddrgrpDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingExcludeMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpDynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpDynamicMappingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpDynamicMappingVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpExcludeMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddrgrpTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAddrgrpTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddrgrpTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddrgrp-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddrgrpTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddrgrpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddrgrpUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddrgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallAddrgrpImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallAddrgrp-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("allow_routing", flattenFirewallAddrgrpAllowRouting(o["allow-routing"], d, "allow_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-routing"], "FirewallAddrgrp-AllowRouting"); ok {
			if err = d.Set("allow_routing", vv); err != nil {
				return fmt.Errorf("Error reading allow_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_routing: %v", err)
		}
	}

	if err = d.Set("category", flattenFirewallAddrgrpCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallAddrgrp-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallAddrgrpColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallAddrgrp-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallAddrgrpComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallAddrgrp-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallAddrgrpDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddrgrp-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallAddrgrpDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddrgrp-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("exclude", flattenFirewallAddrgrpExclude(o["exclude"], d, "exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude"], "FirewallAddrgrp-Exclude"); ok {
			if err = d.Set("exclude", vv); err != nil {
				return fmt.Errorf("Error reading exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude: %v", err)
		}
	}

	if err = d.Set("exclude_member", flattenFirewallAddrgrpExcludeMember(o["exclude-member"], d, "exclude_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-member"], "FirewallAddrgrp-ExcludeMember"); ok {
			if err = d.Set("exclude_member", vv); err != nil {
				return fmt.Errorf("Error reading exclude_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_member: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallAddrgrpFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallAddrgrp-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallAddrgrpGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallAddrgrp-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("member", flattenFirewallAddrgrpMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "FirewallAddrgrp-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAddrgrpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddrgrp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallAddrgrpTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddrgrp-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallAddrgrpTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddrgrp-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenFirewallAddrgrpType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAddrgrp-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddrgrpUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallAddrgrp-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddrgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddrgrpImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_image-base64"], _ = expandFirewallAddrgrpDynamicMappingImageBase64(d, i["_image_base64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAddrgrpDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allow-routing"], _ = expandFirewallAddrgrpDynamicMappingAllowRouting(d, i["allow_routing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandFirewallAddrgrpDynamicMappingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallAddrgrpDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallAddrgrpDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclude"], _ = expandFirewallAddrgrpDynamicMappingExclude(d, i["exclude"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclude-member"], _ = expandFirewallAddrgrpDynamicMappingExcludeMember(d, i["exclude_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandFirewallAddrgrpDynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["global-object"], _ = expandFirewallAddrgrpDynamicMappingGlobalObject(d, i["global_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandFirewallAddrgrpDynamicMappingMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddrgrpDynamicMappingTags(d, i["tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddrgrpDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallAddrgrpDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["visibility"], _ = expandFirewallAddrgrpDynamicMappingVisibility(d, i["visibility"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddrgrpDynamicMappingImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddrgrpDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallAddrgrpDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddrgrpDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingExcludeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpDynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpDynamicMappingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpDynamicMappingVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpExcludeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandFirewallAddrgrpTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAddrgrpTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddrgrpTaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddrgrpTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddrgrpType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddrgrpUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddrgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallAddrgrpImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("allow_routing"); ok || d.HasChange("allow_routing") {
		t, err := expandFirewallAddrgrpAllowRouting(d, v, "allow_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-routing"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallAddrgrpCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallAddrgrpColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallAddrgrpComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallAddrgrpDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("exclude"); ok || d.HasChange("exclude") {
		t, err := expandFirewallAddrgrpExclude(d, v, "exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude"] = t
		}
	}

	if v, ok := d.GetOk("exclude_member"); ok || d.HasChange("exclude_member") {
		t, err := expandFirewallAddrgrpExcludeMember(d, v, "exclude_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-member"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallAddrgrpFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallAddrgrpGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandFirewallAddrgrpMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddrgrpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallAddrgrpTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAddrgrpType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallAddrgrpUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
