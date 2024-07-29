// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Content ID settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptContentDeliveryNetworkRuleRulesContentId() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptContentDeliveryNetworkRuleRulesContentIdUpdate,
		Read:   resourceWanoptContentDeliveryNetworkRuleRulesContentIdRead,
		Update: resourceWanoptContentDeliveryNetworkRuleRulesContentIdUpdate,
		Delete: resourceWanoptContentDeliveryNetworkRuleRulesContentIdDelete,

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
			"end_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_skip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"range_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_skip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptContentDeliveryNetworkRuleRulesContentIdUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptContentDeliveryNetworkRuleRulesContentId(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRulesContentId resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptContentDeliveryNetworkRuleRulesContentId(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRuleRulesContentId resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptContentDeliveryNetworkRuleRulesContentId")

	return resourceWanoptContentDeliveryNetworkRuleRulesContentIdRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleRulesContentIdDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptContentDeliveryNetworkRuleRulesContentId(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptContentDeliveryNetworkRuleRulesContentId resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRulesContentIdRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptContentDeliveryNetworkRuleRulesContentId(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRulesContentId resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptContentDeliveryNetworkRuleRulesContentId(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRuleRulesContentId resource from API: %v", err)
	}
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("end_direction", flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection3rdl(o["end-direction"], d, "end_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-direction"], "WanoptContentDeliveryNetworkRuleRulesContentId-EndDirection"); ok {
			if err = d.Set("end_direction", vv); err != nil {
				return fmt.Errorf("Error reading end_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_direction: %v", err)
		}
	}

	if err = d.Set("end_skip", flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip3rdl(o["end-skip"], d, "end_skip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-skip"], "WanoptContentDeliveryNetworkRuleRulesContentId-EndSkip"); ok {
			if err = d.Set("end_skip", vv); err != nil {
				return fmt.Errorf("Error reading end_skip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_skip: %v", err)
		}
	}

	if err = d.Set("end_str", flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr3rdl(o["end-str"], d, "end_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-str"], "WanoptContentDeliveryNetworkRuleRulesContentId-EndStr"); ok {
			if err = d.Set("end_str", vv); err != nil {
				return fmt.Errorf("Error reading end_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_str: %v", err)
		}
	}

	if err = d.Set("range_str", flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr3rdl(o["range-str"], d, "range_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["range-str"], "WanoptContentDeliveryNetworkRuleRulesContentId-RangeStr"); ok {
			if err = d.Set("range_str", vv); err != nil {
				return fmt.Errorf("Error reading range_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range_str: %v", err)
		}
	}

	if err = d.Set("start_direction", flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection3rdl(o["start-direction"], d, "start_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-direction"], "WanoptContentDeliveryNetworkRuleRulesContentId-StartDirection"); ok {
			if err = d.Set("start_direction", vv); err != nil {
				return fmt.Errorf("Error reading start_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_direction: %v", err)
		}
	}

	if err = d.Set("start_skip", flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip3rdl(o["start-skip"], d, "start_skip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-skip"], "WanoptContentDeliveryNetworkRuleRulesContentId-StartSkip"); ok {
			if err = d.Set("start_skip", vv); err != nil {
				return fmt.Errorf("Error reading start_skip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_skip: %v", err)
		}
	}

	if err = d.Set("start_str", flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr3rdl(o["start-str"], d, "start_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-str"], "WanoptContentDeliveryNetworkRuleRulesContentId-StartStr"); ok {
			if err = d.Set("start_str", vv); err != nil {
				return fmt.Errorf("Error reading start_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_str: %v", err)
		}
	}

	if err = d.Set("target", flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget3rdl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "WanoptContentDeliveryNetworkRuleRulesContentId-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	return nil
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_direction"); ok || d.HasChange("end_direction") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection3rdl(d, v, "end_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-direction"] = t
		}
	}

	if v, ok := d.GetOk("end_skip"); ok || d.HasChange("end_skip") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip3rdl(d, v, "end_skip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-skip"] = t
		}
	}

	if v, ok := d.GetOk("end_str"); ok || d.HasChange("end_str") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr3rdl(d, v, "end_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-str"] = t
		}
	}

	if v, ok := d.GetOk("range_str"); ok || d.HasChange("range_str") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr3rdl(d, v, "range_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-str"] = t
		}
	}

	if v, ok := d.GetOk("start_direction"); ok || d.HasChange("start_direction") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection3rdl(d, v, "start_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-direction"] = t
		}
	}

	if v, ok := d.GetOk("start_skip"); ok || d.HasChange("start_skip") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip3rdl(d, v, "start_skip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-skip"] = t
		}
	}

	if v, ok := d.GetOk("start_str"); ok || d.HasChange("start_str") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr3rdl(d, v, "start_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-str"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget3rdl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	return &obj, nil
}
