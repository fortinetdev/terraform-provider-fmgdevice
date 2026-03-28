// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6DynamicMappingSubnetSegment() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6DynamicMappingSubnetSegmentCreate,
		Read:   resourceFirewallAddress6DynamicMappingSubnetSegmentRead,
		Update: resourceFirewallAddress6DynamicMappingSubnetSegmentUpdate,
		Delete: resourceFirewallAddress6DynamicMappingSubnetSegmentDelete,

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
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallAddress6DynamicMappingSubnetSegmentCreate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectFirewallAddress6DynamicMappingSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6DynamicMappingSubnetSegment resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6DynamicMappingSubnetSegment(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6DynamicMappingSubnetSegment(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6DynamicMappingSubnetSegment resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6DynamicMappingSubnetSegment(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6DynamicMappingSubnetSegment resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6DynamicMappingSubnetSegmentRead(d, m)
}

func resourceFirewallAddress6DynamicMappingSubnetSegmentUpdate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectFirewallAddress6DynamicMappingSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6DynamicMappingSubnetSegment resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6DynamicMappingSubnetSegment(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6DynamicMappingSubnetSegment resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6DynamicMappingSubnetSegmentRead(d, m)
}

func resourceFirewallAddress6DynamicMappingSubnetSegmentDelete(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAddress6DynamicMappingSubnetSegment(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6DynamicMappingSubnetSegment resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6DynamicMappingSubnetSegmentRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	address6 := d.Get("address6").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
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
	if address6 == "" {
		address6 = importOptionChecking(m.(*FortiClient).Cfg, "address6")
		if address6 == "" {
			return fmt.Errorf("Parameter address6 is missing")
		}
		if err = d.Set("address6", address6); err != nil {
			return fmt.Errorf("Error set params address6: %v", err)
		}
	}
	if dynamic_mapping_name == "" {
		dynamic_mapping_name = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_name")
		if dynamic_mapping_name == "" {
			return fmt.Errorf("Parameter dynamic_mapping_name is missing")
		}
		if err = d.Set("dynamic_mapping_name", dynamic_mapping_name); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_name: %v", err)
		}
	}
	if dynamic_mapping_vdom == "" {
		dynamic_mapping_vdom = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_vdom")
		if dynamic_mapping_vdom == "" {
			return fmt.Errorf("Parameter dynamic_mapping_vdom is missing")
		}
		if err = d.Set("dynamic_mapping_vdom", dynamic_mapping_vdom); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_vdom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	o, err := c.ReadFirewallAddress6DynamicMappingSubnetSegment(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6DynamicMappingSubnetSegment resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6DynamicMappingSubnetSegment(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6DynamicMappingSubnetSegment resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress6DynamicMappingSubnetSegment(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallAddress6DynamicMappingSubnetSegmentName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress6DynamicMappingSubnetSegment-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAddress6DynamicMappingSubnetSegmentType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAddress6DynamicMappingSubnetSegment-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("value", flattenFirewallAddress6DynamicMappingSubnetSegmentValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "FirewallAddress6DynamicMappingSubnetSegment-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6DynamicMappingSubnetSegmentName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6DynamicMappingSubnetSegment(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddress6DynamicMappingSubnetSegmentName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAddress6DynamicMappingSubnetSegmentType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandFirewallAddress6DynamicMappingSubnetSegmentValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
