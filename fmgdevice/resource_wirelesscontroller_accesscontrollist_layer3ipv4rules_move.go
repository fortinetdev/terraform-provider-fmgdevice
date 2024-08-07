// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AP ACL layer3 ipv4 rule list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerAccessControlListLayer3Ipv4RulesMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveUpdate,
		Read:   resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveRead,
		Update: resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveUpdate,
		Delete: resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"access_control_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"layer3_ipv4_rules": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_control_list := d.Get("access_control_list").(string)
	layer3_ipv4_rules := d.Get("layer3_ipv4_rules").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_control_list"] = access_control_list
	paradict["layer3_ipv4_rules"] = layer3_ipv4_rules

	target := d.Get("target").(string)
	obj, err := getObjectWirelessControllerAccessControlListLayer3Ipv4RulesMove(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlListLayer3Ipv4RulesMove resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerAccessControlListLayer3Ipv4RulesMove(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlListLayer3Ipv4RulesMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WirelessControllerAccessControlListLayer3Ipv4RulesMove" + "_" + device_name + "_" + device_vdom + "_" + access_control_list + "_" + layer3_ipv4_rules + "_" + target)

	return resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveRead(d, m)
}

func resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceWirelessControllerAccessControlListLayer3Ipv4RulesMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	sid := d.Get("layer3_ipv4_rules").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_control_list := d.Get("access_control_list").(string)
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
	if access_control_list == "" {
		access_control_list = importOptionChecking(m.(*FortiClient).Cfg, "access_control_list")
		if access_control_list == "" {
			return fmt.Errorf("Parameter access_control_list is missing")
		}
		if err = d.Set("access_control_list", access_control_list); err != nil {
			return fmt.Errorf("Error set params access_control_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_control_list"] = access_control_list

	o, err := c.ReadWirelessControllerAccessControlListLayer3Ipv4RulesMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WirelessControllerAccessControlListLayer3Ipv4RulesMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["rule-id"]; !ok {
					return fmt.Errorf("Error reading WirelessControllerAccessControlListLayer3Ipv4RulesMove resource: rule_id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["rule-id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading WirelessControllerAccessControlListLayer3Ipv4RulesMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "rule_id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "rule_id(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "rule_id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "rule_id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerAccessControlListLayer3Ipv4RulesMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv4RulesMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv4RulesMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
