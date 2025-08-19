// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerBonjourProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerBonjourProfileCreate,
		Read:   resourceWirelessControllerBonjourProfileRead,
		Update: resourceWirelessControllerBonjourProfileUpdate,
		Delete: resourceWirelessControllerBonjourProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"micro_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policy_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"services": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"to_vlan": &schema.Schema{
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

func resourceWirelessControllerBonjourProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerBonjourProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerBonjourProfileRead(d, m)
}

func resourceWirelessControllerBonjourProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerBonjourProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerBonjourProfileRead(d, m)
}

func resourceWirelessControllerBonjourProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerBonjourProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerBonjourProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBonjourProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerBonjourProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBonjourProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBonjourProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfileMicroLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenWirelessControllerBonjourProfilePolicyListDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "WirelessControllerBonjourProfile-PolicyList-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := i["from-vlan"]; ok {
			v := flattenWirelessControllerBonjourProfilePolicyListFromVlan(i["from-vlan"], d, pre_append)
			tmp["from_vlan"] = fortiAPISubPartPatch(v, "WirelessControllerBonjourProfile-PolicyList-FromVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := i["policy-id"]; ok {
			v := flattenWirelessControllerBonjourProfilePolicyListPolicyId(i["policy-id"], d, pre_append)
			tmp["policy_id"] = fortiAPISubPartPatch(v, "WirelessControllerBonjourProfile-PolicyList-PolicyId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := i["services"]; ok {
			v := flattenWirelessControllerBonjourProfilePolicyListServices(i["services"], d, pre_append)
			tmp["services"] = fortiAPISubPartPatch(v, "WirelessControllerBonjourProfile-PolicyList-Services")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := i["to-vlan"]; ok {
			v := flattenWirelessControllerBonjourProfilePolicyListToVlan(i["to-vlan"], d, pre_append)
			tmp["to_vlan"] = fortiAPISubPartPatch(v, "WirelessControllerBonjourProfile-PolicyList-ToVlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerBonjourProfilePolicyListDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListFromVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerBonjourProfilePolicyListToVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerBonjourProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenWirelessControllerBonjourProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerBonjourProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("micro_location", flattenWirelessControllerBonjourProfileMicroLocation(o["micro-location"], d, "micro_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["micro-location"], "WirelessControllerBonjourProfile-MicroLocation"); ok {
			if err = d.Set("micro_location", vv); err != nil {
				return fmt.Errorf("Error reading micro_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading micro_location: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerBonjourProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerBonjourProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy_list", flattenWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy-list"], "WirelessControllerBonjourProfile-PolicyList"); ok {
				if err = d.Set("policy_list", vv); err != nil {
					return fmt.Errorf("Error reading policy_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_list"); ok {
			if err = d.Set("policy_list", flattenWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy-list"], "WirelessControllerBonjourProfile-PolicyList"); ok {
					if err = d.Set("policy_list", vv); err != nil {
						return fmt.Errorf("Error reading policy_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerBonjourProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerBonjourProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfileMicroLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandWirelessControllerBonjourProfilePolicyListDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["from-vlan"], _ = expandWirelessControllerBonjourProfilePolicyListFromVlan(d, i["from_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy-id"], _ = expandWirelessControllerBonjourProfilePolicyListPolicyId(d, i["policy_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["services"], _ = expandWirelessControllerBonjourProfilePolicyListServices(d, i["services"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["to-vlan"], _ = expandWirelessControllerBonjourProfilePolicyListToVlan(d, i["to_vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerBonjourProfilePolicyListDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListFromVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerBonjourProfilePolicyListToVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBonjourProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerBonjourProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("micro_location"); ok || d.HasChange("micro_location") {
		t, err := expandWirelessControllerBonjourProfileMicroLocation(d, v, "micro_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["micro-location"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerBonjourProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy_list"); ok || d.HasChange("policy_list") {
		t, err := expandWirelessControllerBonjourProfilePolicyList(d, v, "policy_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-list"] = t
		}
	}

	return &obj, nil
}
