// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure firewall IP-translation.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIpTranslation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpTranslationCreate,
		Read:   resourceFirewallIpTranslationRead,
		Update: resourceFirewallIpTranslationUpdate,
		Delete: resourceFirewallIpTranslationDelete,

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
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"map_startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIpTranslationCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIpTranslation(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIpTranslation resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallIpTranslation(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating FirewallIpTranslation resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "transid")))

	return resourceFirewallIpTranslationRead(d, m)
}

func resourceFirewallIpTranslationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIpTranslation(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpTranslation resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallIpTranslation(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpTranslation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "transid")))

	return resourceFirewallIpTranslationRead(d, m)
}

func resourceFirewallIpTranslationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallIpTranslation(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpTranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpTranslationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIpTranslation(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpTranslation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpTranslation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpTranslation resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpTranslationEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpTranslationMapStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpTranslationStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpTranslationTransid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpTranslationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIpTranslation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("endip", flattenFirewallIpTranslationEndip(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "FirewallIpTranslation-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("map_startip", flattenFirewallIpTranslationMapStartip(o["map-startip"], d, "map_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["map-startip"], "FirewallIpTranslation-MapStartip"); ok {
			if err = d.Set("map_startip", vv); err != nil {
				return fmt.Errorf("Error reading map_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading map_startip: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIpTranslationStartip(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "FirewallIpTranslation-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("transid", flattenFirewallIpTranslationTransid(o["transid"], d, "transid")); err != nil {
		if vv, ok := fortiAPIPatch(o["transid"], "FirewallIpTranslation-Transid"); ok {
			if err = d.Set("transid", vv); err != nil {
				return fmt.Errorf("Error reading transid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transid: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallIpTranslationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallIpTranslation-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpTranslationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIpTranslationEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationMapStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationTransid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpTranslation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandFirewallIpTranslationEndip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("map_startip"); ok || d.HasChange("map_startip") {
		t, err := expandFirewallIpTranslationMapStartip(d, v, "map_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map-startip"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandFirewallIpTranslationStartip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("transid"); ok || d.HasChange("transid") {
		t, err := expandFirewallIpTranslationTransid(d, v, "transid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transid"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallIpTranslationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
