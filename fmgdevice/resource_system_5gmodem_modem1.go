// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure 5G Modem1.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystem5GModemModem1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystem5GModemModem1Update,
		Read:   resourceSystem5GModemModem1Read,
		Update: resourceSystem5GModemModem1Update,
		Delete: resourceSystem5GModemModem1Delete,

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
			"carrier_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gps_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intferface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim_data_plan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sim_pin": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sim_switch": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_sim_slot": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"by_connection_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"by_link_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"by_sim_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"modem_disconnection_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sim_switch_log_alert_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sim_switch_log_alert_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sim1_data_plan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sim1_pin": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sim2_data_plan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sim2_pin": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystem5GModemModem1Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystem5GModemModem1(d)
	if err != nil {
		return fmt.Errorf("Error updating System5GModemModem1 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystem5GModemModem1(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating System5GModemModem1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("System5GModemModem1")

	return resourceSystem5GModemModem1Read(d, m)
}

func resourceSystem5GModemModem1Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystem5GModemModem1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting System5GModemModem1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystem5GModemModem1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystem5GModemModem1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading System5GModemModem1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystem5GModemModem1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading System5GModemModem1 resource from API: %v", err)
	}
	return nil
}

func flattenSystem5GModemModem1CarrierConfig2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1CustomIpv4Netmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1DefaultGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1DefaultNetmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1GpsService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1Intferface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1ModemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1NetworkType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimDataPlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystem5GModemModem1SimSwitch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "active_sim_slot"
	if _, ok := i["active-sim-slot"]; ok {
		result["active_sim_slot"] = flattenSystem5GModemModem1SimSwitchActiveSimSlot2edl(i["active-sim-slot"], d, pre_append)
	}

	pre_append = pre + ".0." + "by_connection_state"
	if _, ok := i["by-connection-state"]; ok {
		result["by_connection_state"] = flattenSystem5GModemModem1SimSwitchByConnectionState2edl(i["by-connection-state"], d, pre_append)
	}

	pre_append = pre + ".0." + "by_link_monitor"
	if _, ok := i["by-link-monitor"]; ok {
		result["by_link_monitor"] = flattenSystem5GModemModem1SimSwitchByLinkMonitor2edl(i["by-link-monitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "by_sim_state"
	if _, ok := i["by-sim-state"]; ok {
		result["by_sim_state"] = flattenSystem5GModemModem1SimSwitchBySimState2edl(i["by-sim-state"], d, pre_append)
	}

	pre_append = pre + ".0." + "link_monitor"
	if _, ok := i["link-monitor"]; ok {
		result["link_monitor"] = flattenSystem5GModemModem1SimSwitchLinkMonitor2edl(i["link-monitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_disconnection_time"
	if _, ok := i["modem-disconnection-time"]; ok {
		result["modem_disconnection_time"] = flattenSystem5GModemModem1SimSwitchModemDisconnectionTime2edl(i["modem-disconnection-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim_switch_log_alert_interval"
	if _, ok := i["sim-switch-log-alert-interval"]; ok {
		result["sim_switch_log_alert_interval"] = flattenSystem5GModemModem1SimSwitchSimSwitchLogAlertInterval2edl(i["sim-switch-log-alert-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim_switch_log_alert_threshold"
	if _, ok := i["sim-switch-log-alert-threshold"]; ok {
		result["sim_switch_log_alert_threshold"] = flattenSystem5GModemModem1SimSwitchSimSwitchLogAlertThreshold2edl(i["sim-switch-log-alert-threshold"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystem5GModemModem1SimSwitchActiveSimSlot2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchByConnectionState2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchByLinkMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchBySimState2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchLinkMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchModemDisconnectionTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchSimSwitchLogAlertInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1SimSwitchSimSwitchLogAlertThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystem5GModemModem1Sim1DataPlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystem5GModemModem1Sim2DataPlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystem5GModemModem1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystem5GModemModem1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("carrier_config", flattenSystem5GModemModem1CarrierConfig2edl(o["carrier-config"], d, "carrier_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier-config"], "System5GModemModem1-CarrierConfig"); ok {
			if err = d.Set("carrier_config", vv); err != nil {
				return fmt.Errorf("Error reading carrier_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier_config: %v", err)
		}
	}

	if err = d.Set("custom_ipv4_netmask", flattenSystem5GModemModem1CustomIpv4Netmask2edl(o["custom-ipv4-netmask"], d, "custom_ipv4_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-ipv4-netmask"], "System5GModemModem1-CustomIpv4Netmask"); ok {
			if err = d.Set("custom_ipv4_netmask", vv); err != nil {
				return fmt.Errorf("Error reading custom_ipv4_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenSystem5GModemModem1DefaultGateway2edl(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "System5GModemModem1-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("default_netmask", flattenSystem5GModemModem1DefaultNetmask2edl(o["default-netmask"], d, "default_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-netmask"], "System5GModemModem1-DefaultNetmask"); ok {
			if err = d.Set("default_netmask", vv); err != nil {
				return fmt.Errorf("Error reading default_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_netmask: %v", err)
		}
	}

	if err = d.Set("gps_service", flattenSystem5GModemModem1GpsService2edl(o["gps-service"], d, "gps_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["gps-service"], "System5GModemModem1-GpsService"); ok {
			if err = d.Set("gps_service", vv); err != nil {
				return fmt.Errorf("Error reading gps_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gps_service: %v", err)
		}
	}

	if err = d.Set("intferface", flattenSystem5GModemModem1Intferface2edl(o["intferface"], d, "intferface")); err != nil {
		if vv, ok := fortiAPIPatch(o["intferface"], "System5GModemModem1-Intferface"); ok {
			if err = d.Set("intferface", vv); err != nil {
				return fmt.Errorf("Error reading intferface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intferface: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenSystem5GModemModem1ModemId2edl(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "System5GModemModem1-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("network_type", flattenSystem5GModemModem1NetworkType2edl(o["network-type"], d, "network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-type"], "System5GModemModem1-NetworkType"); ok {
			if err = d.Set("network_type", vv); err != nil {
				return fmt.Errorf("Error reading network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("sim_data_plan", flattenSystem5GModemModem1SimDataPlan2edl(o["sim-data-plan"], d, "sim_data_plan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim-data-plan"], "System5GModemModem1-SimDataPlan"); ok {
			if err = d.Set("sim_data_plan", vv); err != nil {
				return fmt.Errorf("Error reading sim_data_plan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim_data_plan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sim_switch", flattenSystem5GModemModem1SimSwitch2edl(o["sim-switch"], d, "sim_switch")); err != nil {
			if vv, ok := fortiAPIPatch(o["sim-switch"], "System5GModemModem1-SimSwitch"); ok {
				if err = d.Set("sim_switch", vv); err != nil {
					return fmt.Errorf("Error reading sim_switch: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sim_switch: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sim_switch"); ok {
			if err = d.Set("sim_switch", flattenSystem5GModemModem1SimSwitch2edl(o["sim-switch"], d, "sim_switch")); err != nil {
				if vv, ok := fortiAPIPatch(o["sim-switch"], "System5GModemModem1-SimSwitch"); ok {
					if err = d.Set("sim_switch", vv); err != nil {
						return fmt.Errorf("Error reading sim_switch: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sim_switch: %v", err)
				}
			}
		}
	}

	if err = d.Set("sim1_data_plan", flattenSystem5GModemModem1Sim1DataPlan2edl(o["sim1-data-plan"], d, "sim1_data_plan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim1-data-plan"], "System5GModemModem1-Sim1DataPlan"); ok {
			if err = d.Set("sim1_data_plan", vv); err != nil {
				return fmt.Errorf("Error reading sim1_data_plan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim1_data_plan: %v", err)
		}
	}

	if err = d.Set("sim2_data_plan", flattenSystem5GModemModem1Sim2DataPlan2edl(o["sim2-data-plan"], d, "sim2_data_plan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim2-data-plan"], "System5GModemModem1-Sim2DataPlan"); ok {
			if err = d.Set("sim2_data_plan", vv); err != nil {
				return fmt.Errorf("Error reading sim2_data_plan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim2_data_plan: %v", err)
		}
	}

	if err = d.Set("status", flattenSystem5GModemModem1Status2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "System5GModemModem1-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystem5GModemModem1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystem5GModemModem1CarrierConfig2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1CustomIpv4Netmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1DefaultGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1DefaultNetmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1GpsService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1Intferface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1ModemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1NetworkType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimDataPlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1SimPin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1SimSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "active_sim_slot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["active-sim-slot"], _ = expandSystem5GModemModem1SimSwitchActiveSimSlot2edl(d, i["active_sim_slot"], pre_append)
	}
	pre_append = pre + ".0." + "by_connection_state"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["by-connection-state"], _ = expandSystem5GModemModem1SimSwitchByConnectionState2edl(d, i["by_connection_state"], pre_append)
	}
	pre_append = pre + ".0." + "by_link_monitor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["by-link-monitor"], _ = expandSystem5GModemModem1SimSwitchByLinkMonitor2edl(d, i["by_link_monitor"], pre_append)
	}
	pre_append = pre + ".0." + "by_sim_state"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["by-sim-state"], _ = expandSystem5GModemModem1SimSwitchBySimState2edl(d, i["by_sim_state"], pre_append)
	}
	pre_append = pre + ".0." + "link_monitor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-monitor"], _ = expandSystem5GModemModem1SimSwitchLinkMonitor2edl(d, i["link_monitor"], pre_append)
	}
	pre_append = pre + ".0." + "modem_disconnection_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-disconnection-time"], _ = expandSystem5GModemModem1SimSwitchModemDisconnectionTime2edl(d, i["modem_disconnection_time"], pre_append)
	}
	pre_append = pre + ".0." + "sim_switch_log_alert_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim-switch-log-alert-interval"], _ = expandSystem5GModemModem1SimSwitchSimSwitchLogAlertInterval2edl(d, i["sim_switch_log_alert_interval"], pre_append)
	}
	pre_append = pre + ".0." + "sim_switch_log_alert_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim-switch-log-alert-threshold"], _ = expandSystem5GModemModem1SimSwitchSimSwitchLogAlertThreshold2edl(d, i["sim_switch_log_alert_threshold"], pre_append)
	}

	return result, nil
}

func expandSystem5GModemModem1SimSwitchActiveSimSlot2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchByConnectionState2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchByLinkMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchBySimState2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchLinkMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchModemDisconnectionTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchSimSwitchLogAlertInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1SimSwitchSimSwitchLogAlertThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystem5GModemModem1Sim1DataPlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1Sim1Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1Sim2DataPlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1Sim2Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystem5GModemModem1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystem5GModemModem1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("carrier_config"); ok || d.HasChange("carrier_config") {
		t, err := expandSystem5GModemModem1CarrierConfig2edl(d, v, "carrier_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-config"] = t
		}
	}

	if v, ok := d.GetOk("custom_ipv4_netmask"); ok || d.HasChange("custom_ipv4_netmask") {
		t, err := expandSystem5GModemModem1CustomIpv4Netmask2edl(d, v, "custom_ipv4_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandSystem5GModemModem1DefaultGateway2edl(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("default_netmask"); ok || d.HasChange("default_netmask") {
		t, err := expandSystem5GModemModem1DefaultNetmask2edl(d, v, "default_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-netmask"] = t
		}
	}

	if v, ok := d.GetOk("gps_service"); ok || d.HasChange("gps_service") {
		t, err := expandSystem5GModemModem1GpsService2edl(d, v, "gps_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gps-service"] = t
		}
	}

	if v, ok := d.GetOk("intferface"); ok || d.HasChange("intferface") {
		t, err := expandSystem5GModemModem1Intferface2edl(d, v, "intferface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intferface"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandSystem5GModemModem1ModemId2edl(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok || d.HasChange("network_type") {
		t, err := expandSystem5GModemModem1NetworkType2edl(d, v, "network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("sim_data_plan"); ok || d.HasChange("sim_data_plan") {
		t, err := expandSystem5GModemModem1SimDataPlan2edl(d, v, "sim_data_plan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-data-plan"] = t
		}
	}

	if v, ok := d.GetOk("sim_pin"); ok || d.HasChange("sim_pin") {
		t, err := expandSystem5GModemModem1SimPin2edl(d, v, "sim_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim_switch"); ok || d.HasChange("sim_switch") {
		t, err := expandSystem5GModemModem1SimSwitch2edl(d, v, "sim_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-switch"] = t
		}
	}

	if v, ok := d.GetOk("sim1_data_plan"); ok || d.HasChange("sim1_data_plan") {
		t, err := expandSystem5GModemModem1Sim1DataPlan2edl(d, v, "sim1_data_plan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-data-plan"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin"); ok || d.HasChange("sim1_pin") {
		t, err := expandSystem5GModemModem1Sim1Pin2edl(d, v, "sim1_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim2_data_plan"); ok || d.HasChange("sim2_data_plan") {
		t, err := expandSystem5GModemModem1Sim2DataPlan2edl(d, v, "sim2_data_plan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-data-plan"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin"); ok || d.HasChange("sim2_pin") {
		t, err := expandSystem5GModemModem1Sim2Pin2edl(d, v, "sim2_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystem5GModemModem1Status2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
