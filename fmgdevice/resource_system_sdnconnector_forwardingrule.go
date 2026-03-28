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

func resourceSystemSdnConnectorForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorForwardingRuleCreate,
		Read:   resourceSystemSdnConnectorForwardingRuleRead,
		Update: resourceSystemSdnConnectorForwardingRuleUpdate,
		Delete: resourceSystemSdnConnectorForwardingRuleDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSdnConnectorForwardingRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectSystemSdnConnectorForwardingRule(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnectorForwardingRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("rule_name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSdnConnectorForwardingRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSdnConnectorForwardingRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemSdnConnectorForwardingRule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemSdnConnectorForwardingRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemSdnConnectorForwardingRule resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "rule_name"))

	return resourceSystemSdnConnectorForwardingRuleRead(d, m)
}

func resourceSystemSdnConnectorForwardingRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	obj, err := getObjectSystemSdnConnectorForwardingRule(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorForwardingRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSdnConnectorForwardingRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnectorForwardingRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "rule_name"))

	return resourceSystemSdnConnectorForwardingRuleRead(d, m)
}

func resourceSystemSdnConnectorForwardingRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["sdn_connector"] = sdn_connector

	wsParams["adom"] = adomv

	err = c.DeleteSystemSdnConnectorForwardingRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnectorForwardingRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorForwardingRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

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

	o, err := c.ReadSystemSdnConnectorForwardingRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemSdnConnectorForwardingRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnectorForwardingRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnectorForwardingRule resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorForwardingRuleRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnConnectorForwardingRuleTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnConnectorForwardingRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("rule_name", flattenSystemSdnConnectorForwardingRuleRuleName2edl(o["rule-name"], d, "rule_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule-name"], "SystemSdnConnectorForwardingRule-RuleName"); ok {
			if err = d.Set("rule_name", vv); err != nil {
				return fmt.Errorf("Error reading rule_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule_name: %v", err)
		}
	}

	if err = d.Set("target", flattenSystemSdnConnectorForwardingRuleTarget2edl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "SystemSdnConnectorForwardingRule-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorForwardingRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnConnectorForwardingRuleRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRuleTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnectorForwardingRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rule_name"); ok || d.HasChange("rule_name") {
		t, err := expandSystemSdnConnectorForwardingRuleRuleName2edl(d, v, "rule_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-name"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandSystemSdnConnectorForwardingRuleTarget2edl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
