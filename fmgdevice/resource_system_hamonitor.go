// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure HA monitor.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaMonitorUpdate,
		Read:   resourceSystemHaMonitorRead,
		Update: resourceSystemHaMonitorUpdate,
		Delete: resourceSystemHaMonitorDelete,

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
			"monitor_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaMonitorUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemHaMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMonitor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemHaMonitor")

	return resourceSystemHaMonitorRead(d, m)
}

func resourceSystemHaMonitorDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemHaMonitor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemHaMonitor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaMonitorMonitorVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("monitor_vlan", flattenSystemHaMonitorMonitorVlan(o["monitor-vlan"], d, "monitor_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-vlan"], "SystemHaMonitor-MonitorVlan"); ok {
			if err = d.Set("monitor_vlan", vv); err != nil {
				return fmt.Errorf("Error reading monitor_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_hb_interval", flattenSystemHaMonitorVlanHbInterval(o["vlan-hb-interval"], d, "vlan_hb_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-hb-interval"], "SystemHaMonitor-VlanHbInterval"); ok {
			if err = d.Set("vlan_hb_interval", vv); err != nil {
				return fmt.Errorf("Error reading vlan_hb_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_hb_interval: %v", err)
		}
	}

	if err = d.Set("vlan_hb_lost_threshold", flattenSystemHaMonitorVlanHbLostThreshold(o["vlan-hb-lost-threshold"], d, "vlan_hb_lost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-hb-lost-threshold"], "SystemHaMonitor-VlanHbLostThreshold"); ok {
			if err = d.Set("vlan_hb_lost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading vlan_hb_lost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_hb_lost_threshold: %v", err)
		}
	}

	return nil
}

func flattenSystemHaMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaMonitorMonitorVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("monitor_vlan"); ok || d.HasChange("monitor_vlan") {
		t, err := expandSystemHaMonitorMonitorVlan(d, v, "monitor_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vlan_hb_interval"); ok || d.HasChange("vlan_hb_interval") {
		t, err := expandSystemHaMonitorVlanHbInterval(d, v, "vlan_hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("vlan_hb_lost_threshold"); ok || d.HasChange("vlan_hb_lost_threshold") {
		t, err := expandSystemHaMonitorVlanHbLostThreshold(d, v, "vlan_hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-hb-lost-threshold"] = t
		}
	}

	return &obj, nil
}
