// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch IGMP snooping global settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIgmpSnooping() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIgmpSnoopingUpdate,
		Read:   resourceSwitchControllerIgmpSnoopingRead,
		Update: resourceSwitchControllerIgmpSnoopingUpdate,
		Delete: resourceSwitchControllerIgmpSnoopingDelete,

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
			"aging_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flood_unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerIgmpSnoopingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerIgmpSnooping(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnooping resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerIgmpSnooping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnooping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerIgmpSnooping")

	return resourceSwitchControllerIgmpSnoopingRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

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

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerIgmpSnooping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIgmpSnooping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIgmpSnoopingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerIgmpSnooping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnooping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIgmpSnooping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnooping resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIgmpSnoopingAgingTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingFloodUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerIgmpSnooping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aging_time", flattenSwitchControllerIgmpSnoopingAgingTime(o["aging-time"], d, "aging_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["aging-time"], "SwitchControllerIgmpSnooping-AgingTime"); ok {
			if err = d.Set("aging_time", vv); err != nil {
				return fmt.Errorf("Error reading aging_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aging_time: %v", err)
		}
	}

	if err = d.Set("flood_unknown_multicast", flattenSwitchControllerIgmpSnoopingFloodUnknownMulticast(o["flood-unknown-multicast"], d, "flood_unknown_multicast")); err != nil {
		if vv, ok := fortiAPIPatch(o["flood-unknown-multicast"], "SwitchControllerIgmpSnooping-FloodUnknownMulticast"); ok {
			if err = d.Set("flood_unknown_multicast", vv); err != nil {
				return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
		}
	}

	if err = d.Set("query_interval", flattenSwitchControllerIgmpSnoopingQueryInterval(o["query-interval"], d, "query_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-interval"], "SwitchControllerIgmpSnooping-QueryInterval"); ok {
			if err = d.Set("query_interval", vv); err != nil {
				return fmt.Errorf("Error reading query_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_interval: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIgmpSnoopingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerIgmpSnoopingAgingTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingFloodUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingQueryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerIgmpSnooping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aging_time"); ok || d.HasChange("aging_time") {
		t, err := expandSwitchControllerIgmpSnoopingAgingTime(d, v, "aging_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aging-time"] = t
		}
	}

	if v, ok := d.GetOk("flood_unknown_multicast"); ok || d.HasChange("flood_unknown_multicast") {
		t, err := expandSwitchControllerIgmpSnoopingFloodUnknownMulticast(d, v, "flood_unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flood-unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("query_interval"); ok || d.HasChange("query_interval") {
		t, err := expandSwitchControllerIgmpSnoopingQueryInterval(d, v, "query_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-interval"] = t
		}
	}

	return &obj, nil
}
