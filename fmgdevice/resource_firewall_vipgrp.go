// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 virtual IP groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVipgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipgrpCreate,
		Read:   resourceFirewallVipgrpRead,
		Update: resourceFirewallVipgrpUpdate,
		Delete: resourceFirewallVipgrpDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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

func resourceFirewallVipgrpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVipgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVipgrp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVipgrp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVipgrp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVipgrp resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallVipgrp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVipgrp resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVipgrpRead(d, m)
}

func resourceFirewallVipgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVipgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipgrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVipgrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallVipgrpRead(d, m)
}

func resourceFirewallVipgrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallVipgrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVipgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallVipgrp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVipgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVipgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipgrp resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipgrpColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipgrpDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallVipgrpDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenFirewallVipgrpDynamicMappingComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenFirewallVipgrpDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenFirewallVipgrpDynamicMappingMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallVipgrpDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallVipgrp-DynamicMapping-Uuid")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipgrpDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallVipgrpDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallVipgrpDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallVipgrpDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallVipgrpDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallVipgrpDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMappingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpDynamicMappingMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipgrpDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipgrpMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrpUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVipgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("color", flattenFirewallVipgrpColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallVipgrp-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallVipgrpComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "FirewallVipgrp-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallVipgrpDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVipgrp-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallVipgrpDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallVipgrp-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenFirewallVipgrpInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallVipgrp-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("member", flattenFirewallVipgrpMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "FirewallVipgrp-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallVipgrpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallVipgrp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVipgrpUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallVipgrp-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVipgrpColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandFirewallVipgrpDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallVipgrpDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandFirewallVipgrpDynamicMappingComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandFirewallVipgrpDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandFirewallVipgrpDynamicMappingMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallVipgrpDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipgrpDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallVipgrpDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallVipgrpDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallVipgrpDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMappingComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpDynamicMappingMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipgrpDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipgrpMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrpUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVipgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallVipgrpColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandFirewallVipgrpComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallVipgrpDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallVipgrpInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandFirewallVipgrpMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallVipgrpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallVipgrpUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
