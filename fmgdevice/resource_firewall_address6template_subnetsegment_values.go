// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Subnet segment values.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6TemplateSubnetSegmentValues() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6TemplateSubnetSegmentValuesCreate,
		Read:   resourceFirewallAddress6TemplateSubnetSegmentValuesRead,
		Update: resourceFirewallAddress6TemplateSubnetSegmentValuesUpdate,
		Delete: resourceFirewallAddress6TemplateSubnetSegmentValuesDelete,

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
			"address6_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
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

func resourceFirewallAddress6TemplateSubnetSegmentValuesCreate(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	obj, err := getObjectFirewallAddress6TemplateSubnetSegmentValues(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6TemplateSubnetSegmentValues resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6TemplateSubnetSegmentValues(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6TemplateSubnetSegmentValues(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6TemplateSubnetSegmentValues(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6TemplateSubnetSegmentValuesRead(d, m)
}

func resourceFirewallAddress6TemplateSubnetSegmentValuesUpdate(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	obj, err := getObjectFirewallAddress6TemplateSubnetSegmentValues(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegmentValues resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6TemplateSubnetSegmentValues(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6TemplateSubnetSegmentValuesRead(d, m)
}

func resourceFirewallAddress6TemplateSubnetSegmentValuesDelete(d *schema.ResourceData, m interface{}) error {
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
	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAddress6TemplateSubnetSegmentValues(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6TemplateSubnetSegmentValuesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
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
	if address6_template == "" {
		address6_template = importOptionChecking(m.(*FortiClient).Cfg, "address6_template")
		if address6_template == "" {
			return fmt.Errorf("Parameter address6_template is missing")
		}
		if err = d.Set("address6_template", address6_template); err != nil {
			return fmt.Errorf("Error set params address6_template: %v", err)
		}
	}
	if subnet_segment == "" {
		subnet_segment = importOptionChecking(m.(*FortiClient).Cfg, "subnet_segment")
		if subnet_segment == "" {
			return fmt.Errorf("Parameter subnet_segment is missing")
		}
		if err = d.Set("subnet_segment", subnet_segment); err != nil {
			return fmt.Errorf("Error set params subnet_segment: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	o, err := c.ReadFirewallAddress6TemplateSubnetSegmentValues(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6TemplateSubnetSegmentValues(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6TemplateSubnetSegmentValues resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallAddress6TemplateSubnetSegmentValuesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress6TemplateSubnetSegmentValues-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "FirewallAddress6TemplateSubnetSegmentValues-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6TemplateSubnetSegmentValuesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentValuesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
