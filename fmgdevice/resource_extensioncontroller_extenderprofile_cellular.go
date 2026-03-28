// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender cellular configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfileCellular() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCellularUpdate,
		Read:   resourceExtensionControllerExtenderProfileCellularRead,
		Update: resourceExtensionControllerExtenderProfileCellularUpdate,
		Delete: resourceExtensionControllerExtenderProfileCellularDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"controller_report": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"signal_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dataplan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"modem1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dataplan": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"disconnect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"disconnect_period": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"disconnect_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"signal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_back": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"switch_back_time": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_back_timer": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"conn_status": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"default_sim": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"modem_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"multiple_pdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdn1_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn2_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn3_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn4_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"preferred_carrier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redundant_intf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redundant_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sim1_pin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sim1_pin_code": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"sim2_pin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sim2_pin_code": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
			"modem2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dataplan": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"disconnect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"disconnect_period": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"disconnect_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"signal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_back": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"switch_back_time": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_back_timer": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"conn_status": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"default_sim": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"modem_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"multiple_pdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdn1_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn2_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn3_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"pdn4_dataplan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"preferred_carrier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redundant_intf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redundant_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sim1_pin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sim1_pin_code": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"sim2_pin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sim2_pin_code": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
			"sms_notification": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_exhausted": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"fgt_backup_mode_switch": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"low_signal_strength": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mode_switch": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"os_image_fallback": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"session_disconnect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"system_reboot": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"receiver": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCellularUpdate(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectExtensionControllerExtenderProfileCellular(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellular resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfileCellular(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerExtenderProfileCellular")

	return resourceExtensionControllerExtenderProfileCellularRead(d, m)
}

func resourceExtensionControllerExtenderProfileCellularDelete(d *schema.ResourceData, m interface{}) error {
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
	extender_profile := d.Get("extender_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteExtensionControllerExtenderProfileCellular(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileCellularRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	extender_profile := d.Get("extender_profile").(string)
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
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadExtensionControllerExtenderProfileCellular(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfileCellular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfileCellular resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileCellularControllerReport2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtensionControllerExtenderProfileCellularControllerReportInterval2edl(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularControllerReportStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularControllerReportInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem12edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem1Gps2edl(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtensionControllerExtenderProfileCellularModem1ModemId2edl(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn2edl(i["multiple-PDN"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan2edl(i["pdn1-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan2edl(i["pdn2-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan2edl(i["pdn3-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan2edl(i["pdn4-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Gps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1ModemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem22edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem2Gps2edl(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtensionControllerExtenderProfileCellularModem2ModemId2edl(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn2edl(i["multiple-PDN"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan2edl(i["pdn1-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan2edl(i["pdn2-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan2edl(i["pdn3-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan2edl(i["pdn4-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Gps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2ModemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotification2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(i["alert"], d, pre_append)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(i["receiver"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfileCellular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("controller_report", flattenExtensionControllerExtenderProfileCellularControllerReport2edl(o["controller-report"], d, "controller_report")); err != nil {
			if vv, ok := fortiAPIPatch(o["controller-report"], "ExtensionControllerExtenderProfileCellular-ControllerReport"); ok {
				if err = d.Set("controller_report", vv); err != nil {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading controller_report: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("controller_report"); ok {
			if err = d.Set("controller_report", flattenExtensionControllerExtenderProfileCellularControllerReport2edl(o["controller-report"], d, "controller_report")); err != nil {
				if vv, ok := fortiAPIPatch(o["controller-report"], "ExtensionControllerExtenderProfileCellular-ControllerReport"); ok {
					if err = d.Set("controller_report", vv); err != nil {
						return fmt.Errorf("Error reading controller_report: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			}
		}
	}

	if err = d.Set("dataplan", flattenExtensionControllerExtenderProfileCellularDataplan2edl(o["dataplan"], d, "dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dataplan"], "ExtensionControllerExtenderProfileCellular-Dataplan"); ok {
			if err = d.Set("dataplan", vv); err != nil {
				return fmt.Errorf("Error reading dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dataplan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("modem1", flattenExtensionControllerExtenderProfileCellularModem12edl(o["modem1"], d, "modem1")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem1"], "ExtensionControllerExtenderProfileCellular-Modem1"); ok {
				if err = d.Set("modem1", vv); err != nil {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem1"); ok {
			if err = d.Set("modem1", flattenExtensionControllerExtenderProfileCellularModem12edl(o["modem1"], d, "modem1")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem1"], "ExtensionControllerExtenderProfileCellular-Modem1"); ok {
					if err = d.Set("modem1", vv); err != nil {
						return fmt.Errorf("Error reading modem1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem2", flattenExtensionControllerExtenderProfileCellularModem22edl(o["modem2"], d, "modem2")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem2"], "ExtensionControllerExtenderProfileCellular-Modem2"); ok {
				if err = d.Set("modem2", vv); err != nil {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem2"); ok {
			if err = d.Set("modem2", flattenExtensionControllerExtenderProfileCellularModem22edl(o["modem2"], d, "modem2")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem2"], "ExtensionControllerExtenderProfileCellular-Modem2"); ok {
					if err = d.Set("modem2", vv); err != nil {
						return fmt.Errorf("Error reading modem2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sms_notification", flattenExtensionControllerExtenderProfileCellularSmsNotification2edl(o["sms-notification"], d, "sms_notification")); err != nil {
			if vv, ok := fortiAPIPatch(o["sms-notification"], "ExtensionControllerExtenderProfileCellular-SmsNotification"); ok {
				if err = d.Set("sms_notification", vv); err != nil {
					return fmt.Errorf("Error reading sms_notification: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sms_notification: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sms_notification"); ok {
			if err = d.Set("sms_notification", flattenExtensionControllerExtenderProfileCellularSmsNotification2edl(o["sms-notification"], d, "sms_notification")); err != nil {
				if vv, ok := fortiAPIPatch(o["sms-notification"], "ExtensionControllerExtenderProfileCellular-SmsNotification"); ok {
					if err = d.Set("sms_notification", vv); err != nil {
						return fmt.Errorf("Error reading sms_notification: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sms_notification: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileCellularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileCellularControllerReport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandExtensionControllerExtenderProfileCellularControllerReportInterval2edl(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularControllerReportStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem1Gps2edl(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtensionControllerExtenderProfileCellularModem1ModemId2edl(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem1MultiplePdn2edl(d, i["multiple_pdn"], pre_append)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan2edl(d, i["pdn1_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan2edl(d, i["pdn2_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan2edl(d, i["pdn3_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan2edl(d, i["pdn4_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode2edl(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode2edl(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Gps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1ModemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1MultiplePdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem2Gps2edl(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtensionControllerExtenderProfileCellularModem2ModemId2edl(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem2MultiplePdn2edl(d, i["multiple_pdn"], pre_append)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan2edl(d, i["pdn1_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan2edl(d, i["pdn2_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan2edl(d, i["pdn3_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan2edl(d, i["pdn4_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode2edl(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode2edl(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Gps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2ModemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2MultiplePdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(d, i["alert"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["alert"] = t
		}
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(d, i["receiver"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["receiver"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfileCellular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("controller_report"); ok || d.HasChange("controller_report") {
		t, err := expandExtensionControllerExtenderProfileCellularControllerReport2edl(d, v, "controller_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["controller-report"] = t
		}
	}

	if v, ok := d.GetOk("dataplan"); ok || d.HasChange("dataplan") {
		t, err := expandExtensionControllerExtenderProfileCellularDataplan2edl(d, v, "dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataplan"] = t
		}
	}

	if v, ok := d.GetOk("modem1"); ok || d.HasChange("modem1") {
		t, err := expandExtensionControllerExtenderProfileCellularModem12edl(d, v, "modem1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1"] = t
		}
	}

	if v, ok := d.GetOk("modem2"); ok || d.HasChange("modem2") {
		t, err := expandExtensionControllerExtenderProfileCellularModem22edl(d, v, "modem2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2"] = t
		}
	}

	if v, ok := d.GetOk("sms_notification"); ok || d.HasChange("sms_notification") {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotification2edl(d, v, "sms_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-notification"] = t
		}
	}

	return &obj, nil
}
