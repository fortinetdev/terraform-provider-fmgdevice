// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WTP groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpGroupCreate,
		Read:   resourceWirelessControllerWtpGroupRead,
		Update: resourceWirelessControllerWtpGroupUpdate,
		Delete: resourceWirelessControllerWtpGroupDelete,

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
			"ble_major_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"platform_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wtps": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWtpGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWtpGroup(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWtpGroupRead(d, m)
}

func resourceWirelessControllerWtpGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWtpGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpGroup(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerWtpGroupRead(d, m)
}

func resourceWirelessControllerWtpGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerWtpGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWtpGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpGroupBleMajorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupPlatformType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupWtps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectWirelessControllerWtpGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ble_major_id", flattenWirelessControllerWtpGroupBleMajorId(o["ble-major-id"], d, "ble_major_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-major-id"], "WirelessControllerWtpGroup-BleMajorId"); ok {
			if err = d.Set("ble_major_id", vv); err != nil {
				return fmt.Errorf("Error reading ble_major_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_major_id: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWtpGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerWtpGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("platform_type", flattenWirelessControllerWtpGroupPlatformType(o["platform-type"], d, "platform_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["platform-type"], "WirelessControllerWtpGroup-PlatformType"); ok {
			if err = d.Set("platform_type", vv); err != nil {
				return fmt.Errorf("Error reading platform_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading platform_type: %v", err)
		}
	}

	if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o["wtps"], d, "wtps")); err != nil {
		if vv, ok := fortiAPIPatch(o["wtps"], "WirelessControllerWtpGroup-Wtps"); ok {
			if err = d.Set("wtps", vv); err != nil {
				return fmt.Errorf("Error reading wtps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wtps: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpGroupBleMajorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupPlatformType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupWtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWirelessControllerWtpGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ble_major_id"); ok || d.HasChange("ble_major_id") {
		t, err := expandWirelessControllerWtpGroupBleMajorId(d, v, "ble_major_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-major-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerWtpGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platform_type"); ok || d.HasChange("platform_type") {
		t, err := expandWirelessControllerWtpGroupPlatformType(d, v, "platform_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform-type"] = t
		}
	}

	if v, ok := d.GetOk("wtps"); ok || d.HasChange("wtps") {
		t, err := expandWirelessControllerWtpGroupWtps(d, v, "wtps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtps"] = t
		}
	}

	return &obj, nil
}
