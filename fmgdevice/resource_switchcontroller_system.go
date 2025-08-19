// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system-wide switch controller settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSystemUpdate,
		Read:   resourceSwitchControllerSystemRead,
		Update: resourceSwitchControllerSystemUpdate,
		Delete: resourceSwitchControllerSystemDelete,

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
			"caputp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"caputp_max_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"data_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamic_periodic_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_holdoff": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_mac_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_weight_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nac_periodic_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"parallel_process": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"parallel_process_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSystemUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSystem(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSystem(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerSystem")

	return resourceSwitchControllerSystemRead(d, m)
}

func resourceSwitchControllerSystemDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSystem(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSystem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSystemRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSystem(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSystem(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSystemCaputpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemCaputpMaxRetransmit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemDataSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemDynamicPeriodicInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotHoldoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotMacIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotScanInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotWeightThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemNacPeriodicInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemParallelProcess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemParallelProcessOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSystem(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("caputp_echo_interval", flattenSwitchControllerSystemCaputpEchoInterval(o["caputp-echo-interval"], d, "caputp_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["caputp-echo-interval"], "SwitchControllerSystem-CaputpEchoInterval"); ok {
			if err = d.Set("caputp_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading caputp_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caputp_echo_interval: %v", err)
		}
	}

	if err = d.Set("caputp_max_retransmit", flattenSwitchControllerSystemCaputpMaxRetransmit(o["caputp-max-retransmit"], d, "caputp_max_retransmit")); err != nil {
		if vv, ok := fortiAPIPatch(o["caputp-max-retransmit"], "SwitchControllerSystem-CaputpMaxRetransmit"); ok {
			if err = d.Set("caputp_max_retransmit", vv); err != nil {
				return fmt.Errorf("Error reading caputp_max_retransmit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caputp_max_retransmit: %v", err)
		}
	}

	if err = d.Set("data_sync_interval", flattenSwitchControllerSystemDataSyncInterval(o["data-sync-interval"], d, "data_sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-sync-interval"], "SwitchControllerSystem-DataSyncInterval"); ok {
			if err = d.Set("data_sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading data_sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_sync_interval: %v", err)
		}
	}

	if err = d.Set("dynamic_periodic_interval", flattenSwitchControllerSystemDynamicPeriodicInterval(o["dynamic-periodic-interval"], d, "dynamic_periodic_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-periodic-interval"], "SwitchControllerSystem-DynamicPeriodicInterval"); ok {
			if err = d.Set("dynamic_periodic_interval", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_periodic_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_periodic_interval: %v", err)
		}
	}

	if err = d.Set("iot_holdoff", flattenSwitchControllerSystemIotHoldoff(o["iot-holdoff"], d, "iot_holdoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-holdoff"], "SwitchControllerSystem-IotHoldoff"); ok {
			if err = d.Set("iot_holdoff", vv); err != nil {
				return fmt.Errorf("Error reading iot_holdoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_holdoff: %v", err)
		}
	}

	if err = d.Set("iot_mac_idle", flattenSwitchControllerSystemIotMacIdle(o["iot-mac-idle"], d, "iot_mac_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-mac-idle"], "SwitchControllerSystem-IotMacIdle"); ok {
			if err = d.Set("iot_mac_idle", vv); err != nil {
				return fmt.Errorf("Error reading iot_mac_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_mac_idle: %v", err)
		}
	}

	if err = d.Set("iot_scan_interval", flattenSwitchControllerSystemIotScanInterval(o["iot-scan-interval"], d, "iot_scan_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-scan-interval"], "SwitchControllerSystem-IotScanInterval"); ok {
			if err = d.Set("iot_scan_interval", vv); err != nil {
				return fmt.Errorf("Error reading iot_scan_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_scan_interval: %v", err)
		}
	}

	if err = d.Set("iot_weight_threshold", flattenSwitchControllerSystemIotWeightThreshold(o["iot-weight-threshold"], d, "iot_weight_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-weight-threshold"], "SwitchControllerSystem-IotWeightThreshold"); ok {
			if err = d.Set("iot_weight_threshold", vv); err != nil {
				return fmt.Errorf("Error reading iot_weight_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_weight_threshold: %v", err)
		}
	}

	if err = d.Set("nac_periodic_interval", flattenSwitchControllerSystemNacPeriodicInterval(o["nac-periodic-interval"], d, "nac_periodic_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-periodic-interval"], "SwitchControllerSystem-NacPeriodicInterval"); ok {
			if err = d.Set("nac_periodic_interval", vv); err != nil {
				return fmt.Errorf("Error reading nac_periodic_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_periodic_interval: %v", err)
		}
	}

	if err = d.Set("parallel_process", flattenSwitchControllerSystemParallelProcess(o["parallel-process"], d, "parallel_process")); err != nil {
		if vv, ok := fortiAPIPatch(o["parallel-process"], "SwitchControllerSystem-ParallelProcess"); ok {
			if err = d.Set("parallel_process", vv); err != nil {
				return fmt.Errorf("Error reading parallel_process: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parallel_process: %v", err)
		}
	}

	if err = d.Set("parallel_process_override", flattenSwitchControllerSystemParallelProcessOverride(o["parallel-process-override"], d, "parallel_process_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["parallel-process-override"], "SwitchControllerSystem-ParallelProcessOverride"); ok {
			if err = d.Set("parallel_process_override", vv); err != nil {
				return fmt.Errorf("Error reading parallel_process_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parallel_process_override: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenSwitchControllerSystemTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mode"], "SwitchControllerSystem-TunnelMode"); ok {
			if err = d.Set("tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSystemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSystemCaputpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemCaputpMaxRetransmit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemDataSyncInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemDynamicPeriodicInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotHoldoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotMacIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotScanInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotWeightThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemNacPeriodicInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemParallelProcess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemParallelProcessOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSystem(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("caputp_echo_interval"); ok || d.HasChange("caputp_echo_interval") {
		t, err := expandSwitchControllerSystemCaputpEchoInterval(d, v, "caputp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caputp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("caputp_max_retransmit"); ok || d.HasChange("caputp_max_retransmit") {
		t, err := expandSwitchControllerSystemCaputpMaxRetransmit(d, v, "caputp_max_retransmit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caputp-max-retransmit"] = t
		}
	}

	if v, ok := d.GetOk("data_sync_interval"); ok || d.HasChange("data_sync_interval") {
		t, err := expandSwitchControllerSystemDataSyncInterval(d, v, "data_sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-sync-interval"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_periodic_interval"); ok || d.HasChange("dynamic_periodic_interval") {
		t, err := expandSwitchControllerSystemDynamicPeriodicInterval(d, v, "dynamic_periodic_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-periodic-interval"] = t
		}
	}

	if v, ok := d.GetOk("iot_holdoff"); ok || d.HasChange("iot_holdoff") {
		t, err := expandSwitchControllerSystemIotHoldoff(d, v, "iot_holdoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-holdoff"] = t
		}
	}

	if v, ok := d.GetOk("iot_mac_idle"); ok || d.HasChange("iot_mac_idle") {
		t, err := expandSwitchControllerSystemIotMacIdle(d, v, "iot_mac_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-mac-idle"] = t
		}
	}

	if v, ok := d.GetOk("iot_scan_interval"); ok || d.HasChange("iot_scan_interval") {
		t, err := expandSwitchControllerSystemIotScanInterval(d, v, "iot_scan_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("iot_weight_threshold"); ok || d.HasChange("iot_weight_threshold") {
		t, err := expandSwitchControllerSystemIotWeightThreshold(d, v, "iot_weight_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-weight-threshold"] = t
		}
	}

	if v, ok := d.GetOk("nac_periodic_interval"); ok || d.HasChange("nac_periodic_interval") {
		t, err := expandSwitchControllerSystemNacPeriodicInterval(d, v, "nac_periodic_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-periodic-interval"] = t
		}
	}

	if v, ok := d.GetOk("parallel_process"); ok || d.HasChange("parallel_process") {
		t, err := expandSwitchControllerSystemParallelProcess(d, v, "parallel_process")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parallel-process"] = t
		}
	}

	if v, ok := d.GetOk("parallel_process_override"); ok || d.HasChange("parallel_process_override") {
		t, err := expandSwitchControllerSystemParallelProcessOverride(d, v, "parallel_process_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parallel-process-override"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok || d.HasChange("tunnel_mode") {
		t, err := expandSwitchControllerSystemTunnelMode(d, v, "tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	return &obj, nil
}
