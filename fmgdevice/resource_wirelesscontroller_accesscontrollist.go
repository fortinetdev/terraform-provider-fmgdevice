// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi bridge access control list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerAccessControlList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerAccessControlListCreate,
		Read:   resourceWirelessControllerAccessControlListRead,
		Update: resourceWirelessControllerAccessControlListUpdate,
		Delete: resourceWirelessControllerAccessControlListDelete,

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
			"layer3_ipv4_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"layer3_ipv6_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerAccessControlListCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerAccessControlList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAccessControlList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerAccessControlList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerAccessControlListRead(d, m)
}

func resourceWirelessControllerAccessControlListUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerAccessControlList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerAccessControlList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerAccessControlListRead(d, m)
}

func resourceWirelessControllerAccessControlListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerAccessControlList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerAccessControlListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerAccessControlList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAccessControlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerAccessControlList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAccessControlList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerAccessControlListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4Rules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := i["dstport"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstport(i["dstport"], d, pre_append)
			tmp["dstport"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Dstport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(i["rule-id"], d, pre_append)
			tmp["rule_id"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-RuleId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := i["srcport"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(i["srcport"], d, pre_append)
			tmp["srcport"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv4Rules-Srcport")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6Rules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := i["dstport"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstport(i["dstport"], d, pre_append)
			tmp["dstport"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Dstport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(i["rule-id"], d, pre_append)
			tmp["rule_id"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-RuleId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := i["srcport"]; ok {
			v := flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(i["srcport"], d, pre_append)
			tmp["srcport"] = fortiAPISubPartPatch(v, "WirelessControllerAccessControlList-Layer3Ipv6Rules-Srcport")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerAccessControlList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenWirelessControllerAccessControlListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerAccessControlList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("layer3_ipv4_rules", flattenWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["layer3-ipv4-rules"], "WirelessControllerAccessControlList-Layer3Ipv4Rules"); ok {
				if err = d.Set("layer3_ipv4_rules", vv); err != nil {
					return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv4_rules"); ok {
			if err = d.Set("layer3_ipv4_rules", flattenWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["layer3-ipv4-rules"], "WirelessControllerAccessControlList-Layer3Ipv4Rules"); ok {
					if err = d.Set("layer3_ipv4_rules", vv); err != nil {
						return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("layer3_ipv6_rules", flattenWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules")); err != nil {
			if vv, ok := fortiAPIPatch(o["layer3-ipv6-rules"], "WirelessControllerAccessControlList-Layer3Ipv6Rules"); ok {
				if err = d.Set("layer3_ipv6_rules", vv); err != nil {
					return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv6_rules"); ok {
			if err = d.Set("layer3_ipv6_rules", flattenWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules")); err != nil {
				if vv, ok := fortiAPIPatch(o["layer3-ipv6-rules"], "WirelessControllerAccessControlList-Layer3Ipv6Rules"); ok {
					if err = d.Set("layer3_ipv6_rules", vv); err != nil {
						return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenWirelessControllerAccessControlListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerAccessControlList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerAccessControlListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerAccessControlListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstport"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d, i["dstport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-id"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d, i["rule_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcport"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d, i["srcport"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6Rules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstport"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d, i["dstport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-id"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d, i["rule_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcport"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d, i["srcport"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerAccessControlList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerAccessControlListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("layer3_ipv4_rules"); ok || d.HasChange("layer3_ipv4_rules") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv4Rules(d, v, "layer3_ipv4_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv4-rules"] = t
		}
	}

	if v, ok := d.GetOk("layer3_ipv6_rules"); ok || d.HasChange("layer3_ipv6_rules") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv6Rules(d, v, "layer3_ipv6_rules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerAccessControlListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
