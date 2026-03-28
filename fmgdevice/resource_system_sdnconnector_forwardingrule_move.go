// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure GCP forwarding rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnectorForwardingRuleMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorForwardingRuleMoveUpdate,
		Read:   resourceSystemSdnConnectorForwardingRuleMoveRead,
		Update: resourceSystemSdnConnectorForwardingRuleMoveUpdate,
		Delete: resourceSystemSdnConnectorForwardingRuleMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"forwarding_rule": &schema.Schema{
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

func resourceSystemSdnConnectorForwardingRuleMoveUpdate(d *schema.ResourceData, m interface{}) error {
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
	sdn_connector := d.Get("sdn_connector").(string)
	forwarding_rule := d.Get("forwarding_rule").(string)
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector
	paradict["forwarding_rule"] = forwarding_rule

	target := d.Get("target").(string)
	obj, err := getObjectSystemSdnConnectorForwardingRuleMove(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorForwardingRuleMove resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnectorForwardingRuleMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorForwardingRuleMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSdnConnectorForwardingRuleMove" + "_" + device_name + "_" + sdn_connector + "_" + forwarding_rule + "_" + target)

	return resourceSystemSdnConnectorForwardingRuleMoveRead(d, m)
}

func resourceSystemSdnConnectorForwardingRuleMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorForwardingRuleMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	sid := d.Get("forwarding_rule").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	device_name, err := getVariable(cfg, d, "device_name")
	sdn_connector := d.Get("sdn_connector").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	o, err := c.ReadSystemSdnConnectorForwardingRuleMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnectorForwardingRuleMove resource: %v", err)
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
				if _, ok := v["rule-name"]; !ok {
					return fmt.Errorf("Error reading SystemSdnConnectorForwardingRuleMove resource: rule_name doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["rule-name"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading SystemSdnConnectorForwardingRuleMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "rule_name(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "rule_name(" + sid + ") was deleted"
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
					state_pos = "rule_name(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "rule_name(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenSystemSdnConnectorForwardingRuleMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorForwardingRuleMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRuleMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnectorForwardingRuleMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandSystemSdnConnectorForwardingRuleMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandSystemSdnConnectorForwardingRuleMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
