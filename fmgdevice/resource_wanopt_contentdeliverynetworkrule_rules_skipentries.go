// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of entries to skip.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptContentDeliveryNetworkRuleRulesSkipEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesCreate,
		Read:   resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesRead,
		Update: resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesUpdate,
		Delete: resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesDelete,

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
			"content_delivery_network_rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	rules := d.Get("rules").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule
	paradict["rules"] = rules

	obj, err := getObjectWanoptContentDeliveryNetworkRuleRulesSkipEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRuleRulesSkipEntries resource while getting object: %v", err)
	}

	_, err = c.CreateWanoptContentDeliveryNetworkRuleRulesSkipEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRuleRulesSkipEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	rules := d.Get("rules").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule
	paradict["rules"] = rules

	obj, err := getObjectWanoptContentDeliveryNetworkRuleRulesSkipEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRulesSkipEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptContentDeliveryNetworkRuleRulesSkipEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRulesSkipEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	rules := d.Get("rules").(string)
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule
	paradict["rules"] = rules

	err = c.DeleteWanoptContentDeliveryNetworkRuleRulesSkipEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptContentDeliveryNetworkRuleRulesSkipEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRulesSkipEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	content_delivery_network_rule := d.Get("content_delivery_network_rule").(string)
	rules := d.Get("rules").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if content_delivery_network_rule == "" {
		content_delivery_network_rule = importOptionChecking(m.(*FortiClient).Cfg, "content_delivery_network_rule")
		if content_delivery_network_rule == "" {
			return fmt.Errorf("Parameter content_delivery_network_rule is missing")
		}
		if err = d.Set("content_delivery_network_rule", content_delivery_network_rule); err != nil {
			return fmt.Errorf("Error set params content_delivery_network_rule: %v", err)
		}
	}
	if rules == "" {
		rules = importOptionChecking(m.(*FortiClient).Cfg, "rules")
		if rules == "" {
			return fmt.Errorf("Parameter rules is missing")
		}
		if err = d.Set("rules", rules); err != nil {
			return fmt.Errorf("Error set params rules: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["content_delivery_network_rule"] = content_delivery_network_rule
	paradict["rules"] = rules

	o, err := c.ReadWanoptContentDeliveryNetworkRuleRulesSkipEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRulesSkipEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptContentDeliveryNetworkRuleRulesSkipEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRulesSkipEntries resource from API: %v", err)
	}
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WanoptContentDeliveryNetworkRuleRulesSkipEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("pattern", flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern3rdl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "WanoptContentDeliveryNetworkRuleRulesSkipEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("target", flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget3rdl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "WanoptContentDeliveryNetworkRuleRulesSkipEntries-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern3rdl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget3rdl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
