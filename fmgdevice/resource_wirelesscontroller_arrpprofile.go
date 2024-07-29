// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerArrpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerArrpProfileCreate,
		Read:   resourceWirelessControllerArrpProfileRead,
		Update: resourceWirelessControllerArrpProfileUpdate,
		Delete: resourceWirelessControllerArrpProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"darrp_optimize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"darrp_optimize_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"include_dfs_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_weather_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"override_darrp_optimize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selection_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_channel_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_noise_floor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_rx_errors": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_spectral_rssi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_tx_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_channel_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_dfs_channel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_managed_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_noise_floor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_rogue_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_spectral_rssi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_weather_channel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerArrpProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerArrpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerArrpProfile resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerArrpProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerArrpProfileRead(d, m)
}

func resourceWirelessControllerArrpProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerArrpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerArrpProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerArrpProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerArrpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerArrpProfileRead(d, m)
}

func resourceWirelessControllerArrpProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerArrpProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerArrpProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerArrpProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerArrpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerArrpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerArrpProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerArrpProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileDarrpOptimizeSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerArrpProfileIncludeDfsChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileIncludeWeatherChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileMonitorPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileOverrideDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileSelectionPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdChannelLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdNoiseFloor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdRxErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdSpectralRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdTxRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightChannelLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightDfsChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightManagedAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightNoiseFloor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightRogueAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightSpectralRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightWeatherChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerArrpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenWirelessControllerArrpProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WirelessControllerArrpProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenWirelessControllerArrpProfileDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize"], "WirelessControllerArrpProfile-DarrpOptimize"); ok {
			if err = d.Set("darrp_optimize", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if err = d.Set("darrp_optimize_schedules", flattenWirelessControllerArrpProfileDarrpOptimizeSchedules(o["darrp-optimize-schedules"], d, "darrp_optimize_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize-schedules"], "WirelessControllerArrpProfile-DarrpOptimizeSchedules"); ok {
			if err = d.Set("darrp_optimize_schedules", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
		}
	}

	if err = d.Set("include_dfs_channel", flattenWirelessControllerArrpProfileIncludeDfsChannel(o["include-dfs-channel"], d, "include_dfs_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-dfs-channel"], "WirelessControllerArrpProfile-IncludeDfsChannel"); ok {
			if err = d.Set("include_dfs_channel", vv); err != nil {
				return fmt.Errorf("Error reading include_dfs_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_dfs_channel: %v", err)
		}
	}

	if err = d.Set("include_weather_channel", flattenWirelessControllerArrpProfileIncludeWeatherChannel(o["include-weather-channel"], d, "include_weather_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-weather-channel"], "WirelessControllerArrpProfile-IncludeWeatherChannel"); ok {
			if err = d.Set("include_weather_channel", vv); err != nil {
				return fmt.Errorf("Error reading include_weather_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_weather_channel: %v", err)
		}
	}

	if err = d.Set("monitor_period", flattenWirelessControllerArrpProfileMonitorPeriod(o["monitor-period"], d, "monitor_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-period"], "WirelessControllerArrpProfile-MonitorPeriod"); ok {
			if err = d.Set("monitor_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_period: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerArrpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerArrpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_darrp_optimize", flattenWirelessControllerArrpProfileOverrideDarrpOptimize(o["override-darrp-optimize"], d, "override_darrp_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-darrp-optimize"], "WirelessControllerArrpProfile-OverrideDarrpOptimize"); ok {
			if err = d.Set("override_darrp_optimize", vv); err != nil {
				return fmt.Errorf("Error reading override_darrp_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_darrp_optimize: %v", err)
		}
	}

	if err = d.Set("selection_period", flattenWirelessControllerArrpProfileSelectionPeriod(o["selection-period"], d, "selection_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["selection-period"], "WirelessControllerArrpProfile-SelectionPeriod"); ok {
			if err = d.Set("selection_period", vv); err != nil {
				return fmt.Errorf("Error reading selection_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selection_period: %v", err)
		}
	}

	if err = d.Set("threshold_ap", flattenWirelessControllerArrpProfileThresholdAp(o["threshold-ap"], d, "threshold_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-ap"], "WirelessControllerArrpProfile-ThresholdAp"); ok {
			if err = d.Set("threshold_ap", vv); err != nil {
				return fmt.Errorf("Error reading threshold_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_ap: %v", err)
		}
	}

	if err = d.Set("threshold_channel_load", flattenWirelessControllerArrpProfileThresholdChannelLoad(o["threshold-channel-load"], d, "threshold_channel_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-channel-load"], "WirelessControllerArrpProfile-ThresholdChannelLoad"); ok {
			if err = d.Set("threshold_channel_load", vv); err != nil {
				return fmt.Errorf("Error reading threshold_channel_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_channel_load: %v", err)
		}
	}

	if err = d.Set("threshold_noise_floor", flattenWirelessControllerArrpProfileThresholdNoiseFloor(o["threshold-noise-floor"], d, "threshold_noise_floor")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-noise-floor"], "WirelessControllerArrpProfile-ThresholdNoiseFloor"); ok {
			if err = d.Set("threshold_noise_floor", vv); err != nil {
				return fmt.Errorf("Error reading threshold_noise_floor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_noise_floor: %v", err)
		}
	}

	if err = d.Set("threshold_rx_errors", flattenWirelessControllerArrpProfileThresholdRxErrors(o["threshold-rx-errors"], d, "threshold_rx_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-rx-errors"], "WirelessControllerArrpProfile-ThresholdRxErrors"); ok {
			if err = d.Set("threshold_rx_errors", vv); err != nil {
				return fmt.Errorf("Error reading threshold_rx_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_rx_errors: %v", err)
		}
	}

	if err = d.Set("threshold_spectral_rssi", flattenWirelessControllerArrpProfileThresholdSpectralRssi(o["threshold-spectral-rssi"], d, "threshold_spectral_rssi")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-spectral-rssi"], "WirelessControllerArrpProfile-ThresholdSpectralRssi"); ok {
			if err = d.Set("threshold_spectral_rssi", vv); err != nil {
				return fmt.Errorf("Error reading threshold_spectral_rssi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("threshold_tx_retries", flattenWirelessControllerArrpProfileThresholdTxRetries(o["threshold-tx-retries"], d, "threshold_tx_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-tx-retries"], "WirelessControllerArrpProfile-ThresholdTxRetries"); ok {
			if err = d.Set("threshold_tx_retries", vv); err != nil {
				return fmt.Errorf("Error reading threshold_tx_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_tx_retries: %v", err)
		}
	}

	if err = d.Set("weight_channel_load", flattenWirelessControllerArrpProfileWeightChannelLoad(o["weight-channel-load"], d, "weight_channel_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-channel-load"], "WirelessControllerArrpProfile-WeightChannelLoad"); ok {
			if err = d.Set("weight_channel_load", vv); err != nil {
				return fmt.Errorf("Error reading weight_channel_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_channel_load: %v", err)
		}
	}

	if err = d.Set("weight_dfs_channel", flattenWirelessControllerArrpProfileWeightDfsChannel(o["weight-dfs-channel"], d, "weight_dfs_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-dfs-channel"], "WirelessControllerArrpProfile-WeightDfsChannel"); ok {
			if err = d.Set("weight_dfs_channel", vv); err != nil {
				return fmt.Errorf("Error reading weight_dfs_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_dfs_channel: %v", err)
		}
	}

	if err = d.Set("weight_managed_ap", flattenWirelessControllerArrpProfileWeightManagedAp(o["weight-managed-ap"], d, "weight_managed_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-managed-ap"], "WirelessControllerArrpProfile-WeightManagedAp"); ok {
			if err = d.Set("weight_managed_ap", vv); err != nil {
				return fmt.Errorf("Error reading weight_managed_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_managed_ap: %v", err)
		}
	}

	if err = d.Set("weight_noise_floor", flattenWirelessControllerArrpProfileWeightNoiseFloor(o["weight-noise-floor"], d, "weight_noise_floor")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-noise-floor"], "WirelessControllerArrpProfile-WeightNoiseFloor"); ok {
			if err = d.Set("weight_noise_floor", vv); err != nil {
				return fmt.Errorf("Error reading weight_noise_floor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_noise_floor: %v", err)
		}
	}

	if err = d.Set("weight_rogue_ap", flattenWirelessControllerArrpProfileWeightRogueAp(o["weight-rogue-ap"], d, "weight_rogue_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-rogue-ap"], "WirelessControllerArrpProfile-WeightRogueAp"); ok {
			if err = d.Set("weight_rogue_ap", vv); err != nil {
				return fmt.Errorf("Error reading weight_rogue_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_rogue_ap: %v", err)
		}
	}

	if err = d.Set("weight_spectral_rssi", flattenWirelessControllerArrpProfileWeightSpectralRssi(o["weight-spectral-rssi"], d, "weight_spectral_rssi")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-spectral-rssi"], "WirelessControllerArrpProfile-WeightSpectralRssi"); ok {
			if err = d.Set("weight_spectral_rssi", vv); err != nil {
				return fmt.Errorf("Error reading weight_spectral_rssi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("weight_weather_channel", flattenWirelessControllerArrpProfileWeightWeatherChannel(o["weight-weather-channel"], d, "weight_weather_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-weather-channel"], "WirelessControllerArrpProfile-WeightWeatherChannel"); ok {
			if err = d.Set("weight_weather_channel", vv); err != nil {
				return fmt.Errorf("Error reading weight_weather_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_weather_channel: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerArrpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerArrpProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileDarrpOptimizeSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerArrpProfileIncludeDfsChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileIncludeWeatherChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileMonitorPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileOverrideDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileSelectionPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdChannelLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdNoiseFloor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdRxErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdSpectralRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdTxRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightChannelLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightDfsChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightManagedAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightNoiseFloor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightRogueAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightSpectralRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightWeatherChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerArrpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWirelessControllerArrpProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize"); ok || d.HasChange("darrp_optimize") {
		t, err := expandWirelessControllerArrpProfileDarrpOptimize(d, v, "darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize_schedules"); ok || d.HasChange("darrp_optimize_schedules") {
		t, err := expandWirelessControllerArrpProfileDarrpOptimizeSchedules(d, v, "darrp_optimize_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize-schedules"] = t
		}
	}

	if v, ok := d.GetOk("include_dfs_channel"); ok || d.HasChange("include_dfs_channel") {
		t, err := expandWirelessControllerArrpProfileIncludeDfsChannel(d, v, "include_dfs_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-dfs-channel"] = t
		}
	}

	if v, ok := d.GetOk("include_weather_channel"); ok || d.HasChange("include_weather_channel") {
		t, err := expandWirelessControllerArrpProfileIncludeWeatherChannel(d, v, "include_weather_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-weather-channel"] = t
		}
	}

	if v, ok := d.GetOk("monitor_period"); ok || d.HasChange("monitor_period") {
		t, err := expandWirelessControllerArrpProfileMonitorPeriod(d, v, "monitor_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-period"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerArrpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_darrp_optimize"); ok || d.HasChange("override_darrp_optimize") {
		t, err := expandWirelessControllerArrpProfileOverrideDarrpOptimize(d, v, "override_darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("selection_period"); ok || d.HasChange("selection_period") {
		t, err := expandWirelessControllerArrpProfileSelectionPeriod(d, v, "selection_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-period"] = t
		}
	}

	if v, ok := d.GetOk("threshold_ap"); ok || d.HasChange("threshold_ap") {
		t, err := expandWirelessControllerArrpProfileThresholdAp(d, v, "threshold_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-ap"] = t
		}
	}

	if v, ok := d.GetOk("threshold_channel_load"); ok || d.HasChange("threshold_channel_load") {
		t, err := expandWirelessControllerArrpProfileThresholdChannelLoad(d, v, "threshold_channel_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-channel-load"] = t
		}
	}

	if v, ok := d.GetOk("threshold_noise_floor"); ok || d.HasChange("threshold_noise_floor") {
		t, err := expandWirelessControllerArrpProfileThresholdNoiseFloor(d, v, "threshold_noise_floor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-noise-floor"] = t
		}
	}

	if v, ok := d.GetOk("threshold_rx_errors"); ok || d.HasChange("threshold_rx_errors") {
		t, err := expandWirelessControllerArrpProfileThresholdRxErrors(d, v, "threshold_rx_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-rx-errors"] = t
		}
	}

	if v, ok := d.GetOk("threshold_spectral_rssi"); ok || d.HasChange("threshold_spectral_rssi") {
		t, err := expandWirelessControllerArrpProfileThresholdSpectralRssi(d, v, "threshold_spectral_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOk("threshold_tx_retries"); ok || d.HasChange("threshold_tx_retries") {
		t, err := expandWirelessControllerArrpProfileThresholdTxRetries(d, v, "threshold_tx_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-tx-retries"] = t
		}
	}

	if v, ok := d.GetOk("weight_channel_load"); ok || d.HasChange("weight_channel_load") {
		t, err := expandWirelessControllerArrpProfileWeightChannelLoad(d, v, "weight_channel_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-channel-load"] = t
		}
	}

	if v, ok := d.GetOk("weight_dfs_channel"); ok || d.HasChange("weight_dfs_channel") {
		t, err := expandWirelessControllerArrpProfileWeightDfsChannel(d, v, "weight_dfs_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-dfs-channel"] = t
		}
	}

	if v, ok := d.GetOk("weight_managed_ap"); ok || d.HasChange("weight_managed_ap") {
		t, err := expandWirelessControllerArrpProfileWeightManagedAp(d, v, "weight_managed_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-managed-ap"] = t
		}
	}

	if v, ok := d.GetOk("weight_noise_floor"); ok || d.HasChange("weight_noise_floor") {
		t, err := expandWirelessControllerArrpProfileWeightNoiseFloor(d, v, "weight_noise_floor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-noise-floor"] = t
		}
	}

	if v, ok := d.GetOk("weight_rogue_ap"); ok || d.HasChange("weight_rogue_ap") {
		t, err := expandWirelessControllerArrpProfileWeightRogueAp(d, v, "weight_rogue_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-rogue-ap"] = t
		}
	}

	if v, ok := d.GetOk("weight_spectral_rssi"); ok || d.HasChange("weight_spectral_rssi") {
		t, err := expandWirelessControllerArrpProfileWeightSpectralRssi(d, v, "weight_spectral_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOk("weight_weather_channel"); ok || d.HasChange("weight_weather_channel") {
		t, err := expandWirelessControllerArrpProfileWeightWeatherChannel(d, v, "weight_weather_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-weather-channel"] = t
		}
	}

	return &obj, nil
}
