// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Sniffer ports to filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerTrafficSnifferTargetPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficSnifferTargetPortCreate,
		Read:   resourceSwitchControllerTrafficSnifferTargetPortRead,
		Update: resourceSwitchControllerTrafficSnifferTargetPortUpdate,
		Delete: resourceSwitchControllerTrafficSnifferTargetPortDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"out_ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerTrafficSnifferTargetPortCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerTrafficSnifferTargetPort(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetPort resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerTrafficSnifferTargetPort(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficSnifferTargetPort resource: %v", err)
	}

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerTrafficSnifferTargetPortRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetPortUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerTrafficSnifferTargetPort(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetPort resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerTrafficSnifferTargetPort(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSnifferTargetPort resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerTrafficSnifferTargetPortRead(d, m)
}

func resourceSwitchControllerTrafficSnifferTargetPortDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerTrafficSnifferTargetPort(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficSnifferTargetPort resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficSnifferTargetPortRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerTrafficSnifferTargetPort(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetPort resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficSnifferTargetPort(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSnifferTargetPort resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficSnifferTargetPortDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPortInPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerTrafficSnifferTargetPortOutPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerTrafficSnifferTargetPortSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficSnifferTargetPort(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenSwitchControllerTrafficSnifferTargetPortDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerTrafficSnifferTargetPort-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("in_ports", flattenSwitchControllerTrafficSnifferTargetPortInPorts2edl(o["in-ports"], d, "in_ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["in-ports"], "SwitchControllerTrafficSnifferTargetPort-InPorts"); ok {
			if err = d.Set("in_ports", vv); err != nil {
				return fmt.Errorf("Error reading in_ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading in_ports: %v", err)
		}
	}

	if err = d.Set("out_ports", flattenSwitchControllerTrafficSnifferTargetPortOutPorts2edl(o["out-ports"], d, "out_ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-ports"], "SwitchControllerTrafficSnifferTargetPort-OutPorts"); ok {
			if err = d.Set("out_ports", vv); err != nil {
				return fmt.Errorf("Error reading out_ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_ports: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerTrafficSnifferTargetPortSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerTrafficSnifferTargetPort-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerTrafficSnifferTargetPortFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficSnifferTargetPortDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPortInPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerTrafficSnifferTargetPortOutPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerTrafficSnifferTargetPortSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficSnifferTargetPort(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerTrafficSnifferTargetPortDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("in_ports"); ok || d.HasChange("in_ports") {
		t, err := expandSwitchControllerTrafficSnifferTargetPortInPorts2edl(d, v, "in_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["in-ports"] = t
		}
	}

	if v, ok := d.GetOk("out_ports"); ok || d.HasChange("out_ports") {
		t, err := expandSwitchControllerTrafficSnifferTargetPortOutPorts2edl(d, v, "out_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-ports"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerTrafficSnifferTargetPortSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	return &obj, nil
}
