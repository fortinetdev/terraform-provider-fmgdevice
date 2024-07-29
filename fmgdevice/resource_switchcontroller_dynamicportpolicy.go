// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerDynamicPortPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerDynamicPortPolicyCreate,
		Read:   resourceSwitchControllerDynamicPortPolicyRead,
		Update: resourceSwitchControllerDynamicPortPolicyUpdate,
		Delete: resourceSwitchControllerDynamicPortPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortilink": &schema.Schema{
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
			"policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"n802_1x": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bounce_port_link": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hw_vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"lldp_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"qos_policy": &schema.Schema{
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
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_policy": &schema.Schema{
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

func resourceSwitchControllerDynamicPortPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectSwitchControllerDynamicPortPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerDynamicPortPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	obj, err := getObjectSwitchControllerDynamicPortPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerDynamicPortPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
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

	err = c.DeleteSwitchControllerDynamicPortPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerDynamicPortPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerDynamicPortPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerDynamicPortPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerDynamicPortPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := i["802-1x"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicy8021X(i["802-1x"], d, pre_append)
			tmp["n802_1x"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-8021X")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := i["bounce-port-link"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink(i["bounce-port-link"], d, pre_append)
			tmp["bounce_port_link"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-BouncePortLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := i["family"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyFamily(i["family"], d, pre_append)
			tmp["family"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Family")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := i["hw-vendor"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyHwVendor(i["hw-vendor"], d, pre_append)
			tmp["hw_vendor"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-HwVendor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := i["interface-tags"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(i["interface-tags"], d, pre_append)
			tmp["interface_tags"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-InterfaceTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile(i["lldp-profile"], d, pre_append)
			tmp["lldp_profile"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-LldpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_period"
		if _, ok := i["match-period"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyMatchPeriod(i["match-period"], d, pre_append)
			tmp["match_period"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-MatchPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_type"
		if _, ok := i["match-type"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyMatchType(i["match-type"], d, pre_append)
			tmp["match_type"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-MatchType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy(i["qos-policy"], d, pre_append)
			tmp["qos_policy"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-QosPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := i["vlan-policy"]; ok {
			v := flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy(i["vlan-policy"], d, pre_append)
			tmp["vlan_policy"] = fortiAPISubPartPatch(v, "SwitchControllerDynamicPortPolicy-Policy-VlanPolicy")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerDynamicPortPolicyPolicy8021X(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyMatchPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyMatchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenSwitchControllerDynamicPortPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerDynamicPortPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerDynamicPortPolicyFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "SwitchControllerDynamicPortPolicy-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerDynamicPortPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerDynamicPortPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy", flattenSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy"], "SwitchControllerDynamicPortPolicy-Policy"); ok {
				if err = d.Set("policy", vv); err != nil {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy"); ok {
			if err = d.Set("policy", flattenSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy"], "SwitchControllerDynamicPortPolicy-Policy"); ok {
					if err = d.Set("policy", vv); err != nil {
						return fmt.Errorf("Error reading policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerDynamicPortPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerDynamicPortPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["802-1x"], _ = expandSwitchControllerDynamicPortPolicyPolicy8021X(d, i["n802_1x"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bounce-port-link"], _ = expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d, i["bounce_port_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandSwitchControllerDynamicPortPolicyPolicyCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandSwitchControllerDynamicPortPolicyPolicyDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["family"], _ = expandSwitchControllerDynamicPortPolicyPolicyFamily(d, i["family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandSwitchControllerDynamicPortPolicyPolicyHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hw-vendor"], _ = expandSwitchControllerDynamicPortPolicyPolicyHwVendor(d, i["hw_vendor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-tags"], _ = expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d, i["interface_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-profile"], _ = expandSwitchControllerDynamicPortPolicyPolicyLldpProfile(d, i["lldp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandSwitchControllerDynamicPortPolicyPolicyMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-period"], _ = expandSwitchControllerDynamicPortPolicyPolicyMatchPeriod(d, i["match_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-type"], _ = expandSwitchControllerDynamicPortPolicyPolicyMatchType(d, i["match_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerDynamicPortPolicyPolicyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-policy"], _ = expandSwitchControllerDynamicPortPolicyPolicyQosPolicy(d, i["qos_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSwitchControllerDynamicPortPolicyPolicyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSwitchControllerDynamicPortPolicyPolicyType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-policy"], _ = expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d, i["vlan_policy"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerDynamicPortPolicyPolicy8021X(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyLldpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMatchPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMatchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerDynamicPortPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandSwitchControllerDynamicPortPolicyFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerDynamicPortPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy"); ok || d.HasChange("policy") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicy(d, v, "policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	return &obj, nil
}
