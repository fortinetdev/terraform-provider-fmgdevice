// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Virtual cluster table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaVcluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaVclusterCreate,
		Read:   resourceSystemHaVclusterRead,
		Update: resourceSystemHaVclusterUpdate,
		Delete: resourceSystemHaVclusterDelete,

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
			"monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pingserver_failover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pingserver_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pingserver_monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"pingserver_secondary_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingserver_slave_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaVclusterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemHaVcluster(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaVcluster resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaVcluster(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaVcluster resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vcluster_id")))

	return resourceSystemHaVclusterRead(d, m)
}

func resourceSystemHaVclusterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemHaVcluster(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaVcluster resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaVcluster(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaVcluster resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vcluster_id")))

	return resourceSystemHaVclusterRead(d, m)
}

func resourceSystemHaVclusterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemHaVcluster(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaVcluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaVclusterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHaVcluster(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaVcluster resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaVcluster(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaVcluster resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaVclusterMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaVclusterOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterOverrideWaitTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFailoverThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFlipTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverMonitorInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemHaVclusterPingserverSecondaryForceReset2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverSlaveForceReset2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterVclusterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemHaVcluster(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("monitor", flattenSystemHaVclusterMonitor2edl(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "SystemHaVcluster-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemHaVclusterOverride2edl(o["override"], d, "override")); err != nil {
		if vv, ok := fortiAPIPatch(o["override"], "SystemHaVcluster-Override"); ok {
			if err = d.Set("override", vv); err != nil {
				return fmt.Errorf("Error reading override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("override_wait_time", flattenSystemHaVclusterOverrideWaitTime2edl(o["override-wait-time"], d, "override_wait_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-wait-time"], "SystemHaVcluster-OverrideWaitTime"); ok {
			if err = d.Set("override_wait_time", vv); err != nil {
				return fmt.Errorf("Error reading override_wait_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_wait_time: %v", err)
		}
	}

	if err = d.Set("pingserver_failover_threshold", flattenSystemHaVclusterPingserverFailoverThreshold2edl(o["pingserver-failover-threshold"], d, "pingserver_failover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-failover-threshold"], "SystemHaVcluster-PingserverFailoverThreshold"); ok {
			if err = d.Set("pingserver_failover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
		}
	}

	if err = d.Set("pingserver_flip_timeout", flattenSystemHaVclusterPingserverFlipTimeout2edl(o["pingserver-flip-timeout"], d, "pingserver_flip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-flip-timeout"], "SystemHaVcluster-PingserverFlipTimeout"); ok {
			if err = d.Set("pingserver_flip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
		}
	}

	if err = d.Set("pingserver_monitor_interface", flattenSystemHaVclusterPingserverMonitorInterface2edl(o["pingserver-monitor-interface"], d, "pingserver_monitor_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-monitor-interface"], "SystemHaVcluster-PingserverMonitorInterface"); ok {
			if err = d.Set("pingserver_monitor_interface", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
		}
	}

	if err = d.Set("pingserver_secondary_force_reset", flattenSystemHaVclusterPingserverSecondaryForceReset2edl(o["pingserver-secondary-force-reset"], d, "pingserver_secondary_force_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-secondary-force-reset"], "SystemHaVcluster-PingserverSecondaryForceReset"); ok {
			if err = d.Set("pingserver_secondary_force_reset", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_slave_force_reset", flattenSystemHaVclusterPingserverSlaveForceReset2edl(o["pingserver-slave-force-reset"], d, "pingserver_slave_force_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["pingserver-slave-force-reset"], "SystemHaVcluster-PingserverSlaveForceReset"); ok {
			if err = d.Set("pingserver_slave_force_reset", vv); err != nil {
				return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaVclusterPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemHaVcluster-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemHaVclusterVclusterId2edl(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcluster-id"], "SystemHaVcluster-VclusterId"); ok {
			if err = d.Set("vcluster_id", vv); err != nil {
				return fmt.Errorf("Error reading vcluster_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemHaVclusterVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemHaVcluster-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemHaVclusterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaVclusterMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaVclusterOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterOverrideWaitTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFailoverThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFlipTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverMonitorInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemHaVclusterPingserverSecondaryForceReset2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverSlaveForceReset2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterVclusterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemHaVcluster(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandSystemHaVclusterMonitor2edl(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandSystemHaVclusterOverride2edl(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("override_wait_time"); ok || d.HasChange("override_wait_time") {
		t, err := expandSystemHaVclusterOverrideWaitTime2edl(d, v, "override_wait_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wait-time"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_failover_threshold"); ok || d.HasChange("pingserver_failover_threshold") {
		t, err := expandSystemHaVclusterPingserverFailoverThreshold2edl(d, v, "pingserver_failover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-failover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_flip_timeout"); ok || d.HasChange("pingserver_flip_timeout") {
		t, err := expandSystemHaVclusterPingserverFlipTimeout2edl(d, v, "pingserver_flip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-flip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_monitor_interface"); ok || d.HasChange("pingserver_monitor_interface") {
		t, err := expandSystemHaVclusterPingserverMonitorInterface2edl(d, v, "pingserver_monitor_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-monitor-interface"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_secondary_force_reset"); ok || d.HasChange("pingserver_secondary_force_reset") {
		t, err := expandSystemHaVclusterPingserverSecondaryForceReset2edl(d, v, "pingserver_secondary_force_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-secondary-force-reset"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_slave_force_reset"); ok || d.HasChange("pingserver_slave_force_reset") {
		t, err := expandSystemHaVclusterPingserverSlaveForceReset2edl(d, v, "pingserver_slave_force_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-slave-force-reset"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemHaVclusterPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("vcluster_id"); ok || d.HasChange("vcluster_id") {
		t, err := expandSystemHaVclusterVclusterId2edl(d, v, "vcluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemHaVclusterVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
