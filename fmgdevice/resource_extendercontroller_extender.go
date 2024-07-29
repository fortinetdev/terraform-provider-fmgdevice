// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device ExtenderControllerExtender

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtenderControllerExtender() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtenderCreate,
		Read:   resourceExtenderControllerExtenderRead,
		Update: resourceExtenderControllerExtenderUpdate,
		Delete: resourceExtenderControllerExtenderDelete,

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
			"_dataplan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"_template": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"aaa_shared_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"access_point_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"at_dial_script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"billing_start_day": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cdma_aaa_spi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdma_ha_spi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdma_nai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
						},
						"signal_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dial_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dial_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_shared_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"extension_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ifname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"initiated_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modem_passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"modem_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
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
						"_sim_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
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
									},
									"disconnect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"disconnect_period": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"disconnect_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"signal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"switch_back_timer": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
						"ifname": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"_sim_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
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
									},
									"disconnect": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"disconnect_period": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"disconnect_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"signal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"switch_back_timer": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
						"ifname": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"multi_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ppp_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppp_echo_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppp_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ppp_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quota_limit_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redundant_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sim_pin": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"override_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wimax_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wimax_carrier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wimax_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_extension": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceExtenderControllerExtenderCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource while getting object: %v", err)
	}

	_, err = c.CreateExtenderControllerExtender(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource while getting object: %v", err)
	}

	_, err = c.UpdateExtenderControllerExtender(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtenderControllerExtender(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtenderControllerExtender(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtender(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtenderDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderAccessPointName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAtDialScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderAuthorized(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderBillingStartDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaAaaSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaHaSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaNai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderBandwidthLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReport(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtenderControllerExtenderControllerReportInterval(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtenderControllerExtenderControllerReportSignalThreshold(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderControllerReportStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderControllerReportInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderExtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderIfname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderInitiatedUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModemType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_sim_profile"
	if _, ok := i["_sim_profile"]; ok {
		result["_sim_profile"] = flattenExtenderControllerExtenderModem1SimProfile(i["_sim_profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtenderModem1AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtenderModem1ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtenderModem1DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtenderModem1Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {
		result["ifname"] = flattenExtenderControllerExtenderModem1Ifname(i["ifname"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtenderControllerExtenderModem1ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtenderModem1PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtenderModem1RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtenderModem1RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtenderModem1Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtenderModem1Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderModem1Status(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem1SimProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtenderModem1AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtenderModem1AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderModem1AutoSwitchStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem1ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_sim_profile"
	if _, ok := i["_sim_profile"]; ok {
		result["_sim_profile"] = flattenExtenderControllerExtenderModem2SimProfile(i["_sim_profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtenderModem2AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtenderModem2ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtenderModem2DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtenderModem2Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {
		result["ifname"] = flattenExtenderControllerExtenderModem2Ifname(i["ifname"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtenderControllerExtenderModem2ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtenderModem2PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtenderModem2RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtenderModem2RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtenderModem2Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtenderModem2Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderModem2Status(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem2SimProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtenderModem2AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtenderModem2AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderModem2AutoSwitchStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderModem2ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMultiMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppAuthProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppEchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPrimaryHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderQuotaLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderSecondaryHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxAuthProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := i["modem1-extension"]; ok {
		result["modem1_extension"] = flattenExtenderControllerExtenderWanExtensionModem1Extension(i["modem1-extension"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := i["modem2-extension"]; ok {
		result["modem2_extension"] = flattenExtenderControllerExtenderWanExtensionModem2Extension(i["modem2-extension"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderWanExtensionModem1Extension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtenderControllerExtenderWanExtensionModem2Extension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectExtenderControllerExtender(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_dataplan", flattenExtenderControllerExtenderDataplan(o["_dataplan"], d, "_dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dataplan"], "ExtenderControllerExtender-Dataplan"); ok {
			if err = d.Set("_dataplan", vv); err != nil {
				return fmt.Errorf("Error reading _dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dataplan: %v", err)
		}
	}

	if err = d.Set("_template", flattenExtenderControllerExtenderTemplate(o["_template"], d, "_template")); err != nil {
		if vv, ok := fortiAPIPatch(o["_template"], "ExtenderControllerExtender-Template"); ok {
			if err = d.Set("_template", vv); err != nil {
				return fmt.Errorf("Error reading _template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _template: %v", err)
		}
	}

	if err = d.Set("access_point_name", flattenExtenderControllerExtenderAccessPointName(o["access-point-name"], d, "access_point_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-point-name"], "ExtenderControllerExtender-AccessPointName"); ok {
			if err = d.Set("access_point_name", vv); err != nil {
				return fmt.Errorf("Error reading access_point_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_point_name: %v", err)
		}
	}

	if err = d.Set("admin", flattenExtenderControllerExtenderAdmin(o["admin"], d, "admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin"], "ExtenderControllerExtender-Admin"); ok {
			if err = d.Set("admin", vv); err != nil {
				return fmt.Errorf("Error reading admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("at_dial_script", flattenExtenderControllerExtenderAtDialScript(o["at-dial-script"], d, "at_dial_script")); err != nil {
		if vv, ok := fortiAPIPatch(o["at-dial-script"], "ExtenderControllerExtender-AtDialScript"); ok {
			if err = d.Set("at_dial_script", vv); err != nil {
				return fmt.Errorf("Error reading at_dial_script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading at_dial_script: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtenderControllerExtenderAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ExtenderControllerExtender-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtenderControllerExtenderAuthorized(o["authorized"], d, "authorized")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized"], "ExtenderControllerExtender-Authorized"); ok {
			if err = d.Set("authorized", vv); err != nil {
				return fmt.Errorf("Error reading authorized: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("billing_start_day", flattenExtenderControllerExtenderBillingStartDay(o["billing-start-day"], d, "billing_start_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-start-day"], "ExtenderControllerExtender-BillingStartDay"); ok {
			if err = d.Set("billing_start_day", vv); err != nil {
				return fmt.Errorf("Error reading billing_start_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_start_day: %v", err)
		}
	}

	if err = d.Set("cdma_aaa_spi", flattenExtenderControllerExtenderCdmaAaaSpi(o["cdma-aaa-spi"], d, "cdma_aaa_spi")); err != nil {
		if vv, ok := fortiAPIPatch(o["cdma-aaa-spi"], "ExtenderControllerExtender-CdmaAaaSpi"); ok {
			if err = d.Set("cdma_aaa_spi", vv); err != nil {
				return fmt.Errorf("Error reading cdma_aaa_spi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cdma_aaa_spi: %v", err)
		}
	}

	if err = d.Set("cdma_ha_spi", flattenExtenderControllerExtenderCdmaHaSpi(o["cdma-ha-spi"], d, "cdma_ha_spi")); err != nil {
		if vv, ok := fortiAPIPatch(o["cdma-ha-spi"], "ExtenderControllerExtender-CdmaHaSpi"); ok {
			if err = d.Set("cdma_ha_spi", vv); err != nil {
				return fmt.Errorf("Error reading cdma_ha_spi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cdma_ha_spi: %v", err)
		}
	}

	if err = d.Set("cdma_nai", flattenExtenderControllerExtenderCdmaNai(o["cdma-nai"], d, "cdma_nai")); err != nil {
		if vv, ok := fortiAPIPatch(o["cdma-nai"], "ExtenderControllerExtender-CdmaNai"); ok {
			if err = d.Set("cdma_nai", vv); err != nil {
				return fmt.Errorf("Error reading cdma_nai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cdma_nai: %v", err)
		}
	}

	if err = d.Set("conn_status", flattenExtenderControllerExtenderConnStatus(o["conn-status"], d, "conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-status"], "ExtenderControllerExtender-ConnStatus"); ok {
			if err = d.Set("conn_status", vv); err != nil {
				return fmt.Errorf("Error reading conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtenderControllerExtenderBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-limit"], "ExtenderControllerExtender-BandwidthLimit"); ok {
			if err = d.Set("bandwidth_limit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("controller_report", flattenExtenderControllerExtenderControllerReport(o["controller-report"], d, "controller_report")); err != nil {
			if vv, ok := fortiAPIPatch(o["controller-report"], "ExtenderControllerExtender-ControllerReport"); ok {
				if err = d.Set("controller_report", vv); err != nil {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading controller_report: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("controller_report"); ok {
			if err = d.Set("controller_report", flattenExtenderControllerExtenderControllerReport(o["controller-report"], d, "controller_report")); err != nil {
				if vv, ok := fortiAPIPatch(o["controller-report"], "ExtenderControllerExtender-ControllerReport"); ok {
					if err = d.Set("controller_report", vv); err != nil {
						return fmt.Errorf("Error reading controller_report: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenExtenderControllerExtenderDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ExtenderControllerExtender-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dial_mode", flattenExtenderControllerExtenderDialMode(o["dial-mode"], d, "dial_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dial-mode"], "ExtenderControllerExtender-DialMode"); ok {
			if err = d.Set("dial_mode", vv); err != nil {
				return fmt.Errorf("Error reading dial_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dial_mode: %v", err)
		}
	}

	if err = d.Set("dial_status", flattenExtenderControllerExtenderDialStatus(o["dial-status"], d, "dial_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dial-status"], "ExtenderControllerExtender-DialStatus"); ok {
			if err = d.Set("dial_status", vv); err != nil {
				return fmt.Errorf("Error reading dial_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dial_status: %v", err)
		}
	}

	if err = d.Set("device_id", flattenExtenderControllerExtenderDeviceId(o["device-id"], d, "device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-id"], "ExtenderControllerExtender-DeviceId"); ok {
			if err = d.Set("device_id", vv); err != nil {
				return fmt.Errorf("Error reading device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtenderControllerExtenderEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-bandwidth"], "ExtenderControllerExtender-EnforceBandwidth"); ok {
			if err = d.Set("enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtenderControllerExtenderExtName(o["ext-name"], d, "ext_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-name"], "ExtenderControllerExtender-ExtName"); ok {
			if err = d.Set("ext_name", vv); err != nil {
				return fmt.Errorf("Error reading ext_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("extension_type", flattenExtenderControllerExtenderExtensionType(o["extension-type"], d, "extension_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-type"], "ExtenderControllerExtender-ExtensionType"); ok {
			if err = d.Set("extension_type", vv); err != nil {
				return fmt.Errorf("Error reading extension_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtenderControllerExtenderId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ExtenderControllerExtender-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ifname", flattenExtenderControllerExtenderIfname(o["ifname"], d, "ifname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ifname"], "ExtenderControllerExtender-Ifname"); ok {
			if err = d.Set("ifname", vv); err != nil {
				return fmt.Errorf("Error reading ifname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ifname: %v", err)
		}
	}

	if err = d.Set("initiated_update", flattenExtenderControllerExtenderInitiatedUpdate(o["initiated-update"], d, "initiated_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["initiated-update"], "ExtenderControllerExtender-InitiatedUpdate"); ok {
			if err = d.Set("initiated_update", vv); err != nil {
				return fmt.Errorf("Error reading initiated_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading initiated_update: %v", err)
		}
	}

	if err = d.Set("mode", flattenExtenderControllerExtenderMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ExtenderControllerExtender-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("modem_type", flattenExtenderControllerExtenderModemType(o["modem-type"], d, "modem_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-type"], "ExtenderControllerExtender-ModemType"); ok {
			if err = d.Set("modem_type", vv); err != nil {
				return fmt.Errorf("Error reading modem_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_type: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtenderControllerExtenderLoginPasswordChange(o["login-password-change"], d, "login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password-change"], "ExtenderControllerExtender-LoginPasswordChange"); ok {
			if err = d.Set("login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("modem1", flattenExtenderControllerExtenderModem1(o["modem1"], d, "modem1")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem1"], "ExtenderControllerExtender-Modem1"); ok {
				if err = d.Set("modem1", vv); err != nil {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem1"); ok {
			if err = d.Set("modem1", flattenExtenderControllerExtenderModem1(o["modem1"], d, "modem1")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem1"], "ExtenderControllerExtender-Modem1"); ok {
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
		if err = d.Set("modem2", flattenExtenderControllerExtenderModem2(o["modem2"], d, "modem2")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem2"], "ExtenderControllerExtender-Modem2"); ok {
				if err = d.Set("modem2", vv); err != nil {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem2"); ok {
			if err = d.Set("modem2", flattenExtenderControllerExtenderModem2(o["modem2"], d, "modem2")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem2"], "ExtenderControllerExtender-Modem2"); ok {
					if err = d.Set("modem2", vv); err != nil {
						return fmt.Errorf("Error reading modem2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			}
		}
	}

	if err = d.Set("multi_mode", flattenExtenderControllerExtenderMultiMode(o["multi-mode"], d, "multi_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["multi-mode"], "ExtenderControllerExtender-MultiMode"); ok {
			if err = d.Set("multi_mode", vv); err != nil {
				return fmt.Errorf("Error reading multi_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multi_mode: %v", err)
		}
	}

	if err = d.Set("name", flattenExtenderControllerExtenderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtenderControllerExtender-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ppp_auth_protocol", flattenExtenderControllerExtenderPppAuthProtocol(o["ppp-auth-protocol"], d, "ppp_auth_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppp-auth-protocol"], "ExtenderControllerExtender-PppAuthProtocol"); ok {
			if err = d.Set("ppp_auth_protocol", vv); err != nil {
				return fmt.Errorf("Error reading ppp_auth_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppp_auth_protocol: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request", flattenExtenderControllerExtenderPppEchoRequest(o["ppp-echo-request"], d, "ppp_echo_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppp-echo-request"], "ExtenderControllerExtender-PppEchoRequest"); ok {
			if err = d.Set("ppp_echo_request", vv); err != nil {
				return fmt.Errorf("Error reading ppp_echo_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppp_echo_request: %v", err)
		}
	}

	if err = d.Set("ppp_username", flattenExtenderControllerExtenderPppUsername(o["ppp-username"], d, "ppp_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppp-username"], "ExtenderControllerExtender-PppUsername"); ok {
			if err = d.Set("ppp_username", vv); err != nil {
				return fmt.Errorf("Error reading ppp_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppp_username: %v", err)
		}
	}

	if err = d.Set("primary_ha", flattenExtenderControllerExtenderPrimaryHa(o["primary-ha"], d, "primary_ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-ha"], "ExtenderControllerExtender-PrimaryHa"); ok {
			if err = d.Set("primary_ha", vv); err != nil {
				return fmt.Errorf("Error reading primary_ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_ha: %v", err)
		}
	}

	if err = d.Set("quota_limit_mb", flattenExtenderControllerExtenderQuotaLimitMb(o["quota-limit-mb"], d, "quota_limit_mb")); err != nil {
		if vv, ok := fortiAPIPatch(o["quota-limit-mb"], "ExtenderControllerExtender-QuotaLimitMb"); ok {
			if err = d.Set("quota_limit_mb", vv); err != nil {
				return fmt.Errorf("Error reading quota_limit_mb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quota_limit_mb: %v", err)
		}
	}

	if err = d.Set("redial", flattenExtenderControllerExtenderRedial(o["redial"], d, "redial")); err != nil {
		if vv, ok := fortiAPIPatch(o["redial"], "ExtenderControllerExtender-Redial"); ok {
			if err = d.Set("redial", vv); err != nil {
				return fmt.Errorf("Error reading redial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redial: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenExtenderControllerExtenderRedundantIntf(o["redundant-intf"], d, "redundant_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-intf"], "ExtenderControllerExtender-RedundantIntf"); ok {
			if err = d.Set("redundant_intf", vv); err != nil {
				return fmt.Errorf("Error reading redundant_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("roaming", flattenExtenderControllerExtenderRoaming(o["roaming"], d, "roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["roaming"], "ExtenderControllerExtender-Roaming"); ok {
			if err = d.Set("roaming", vv); err != nil {
				return fmt.Errorf("Error reading roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roaming: %v", err)
		}
	}

	if err = d.Set("role", flattenExtenderControllerExtenderRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ExtenderControllerExtender-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("secondary_ha", flattenExtenderControllerExtenderSecondaryHa(o["secondary-ha"], d, "secondary_ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-ha"], "ExtenderControllerExtender-SecondaryHa"); ok {
			if err = d.Set("secondary_ha", vv); err != nil {
				return fmt.Errorf("Error reading secondary_ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_ha: %v", err)
		}
	}

	if err = d.Set("override_allowaccess", flattenExtenderControllerExtenderOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-allowaccess"], "ExtenderControllerExtender-OverrideAllowaccess"); ok {
			if err = d.Set("override_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading override_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("override_enforce_bandwidth", flattenExtenderControllerExtenderOverrideEnforceBandwidth(o["override-enforce-bandwidth"], d, "override_enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-enforce-bandwidth"], "ExtenderControllerExtender-OverrideEnforceBandwidth"); ok {
			if err = d.Set("override_enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("override_login_password_change", flattenExtenderControllerExtenderOverrideLoginPasswordChange(o["override-login-password-change"], d, "override_login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-login-password-change"], "ExtenderControllerExtender-OverrideLoginPasswordChange"); ok {
			if err = d.Set("override_login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading override_login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_login_password_change: %v", err)
		}
	}

	if err = d.Set("profile", flattenExtenderControllerExtenderProfile(o["profile"], d, "profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile"], "ExtenderControllerExtender-Profile"); ok {
			if err = d.Set("profile", vv); err != nil {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtenderControllerExtenderVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ExtenderControllerExtender-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("wimax_auth_protocol", flattenExtenderControllerExtenderWimaxAuthProtocol(o["wimax-auth-protocol"], d, "wimax_auth_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["wimax-auth-protocol"], "ExtenderControllerExtender-WimaxAuthProtocol"); ok {
			if err = d.Set("wimax_auth_protocol", vv); err != nil {
				return fmt.Errorf("Error reading wimax_auth_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wimax_auth_protocol: %v", err)
		}
	}

	if err = d.Set("wimax_carrier", flattenExtenderControllerExtenderWimaxCarrier(o["wimax-carrier"], d, "wimax_carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["wimax-carrier"], "ExtenderControllerExtender-WimaxCarrier"); ok {
			if err = d.Set("wimax_carrier", vv); err != nil {
				return fmt.Errorf("Error reading wimax_carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wimax_carrier: %v", err)
		}
	}

	if err = d.Set("wimax_realm", flattenExtenderControllerExtenderWimaxRealm(o["wimax-realm"], d, "wimax_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["wimax-realm"], "ExtenderControllerExtender-WimaxRealm"); ok {
			if err = d.Set("wimax_realm", vv); err != nil {
				return fmt.Errorf("Error reading wimax_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wimax_realm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wan_extension", flattenExtenderControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["wan-extension"], "ExtenderControllerExtender-WanExtension"); ok {
				if err = d.Set("wan_extension", vv); err != nil {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading wan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wan_extension"); ok {
			if err = d.Set("wan_extension", flattenExtenderControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["wan-extension"], "ExtenderControllerExtender-WanExtension"); ok {
					if err = d.Set("wan_extension", vv); err != nil {
						return fmt.Errorf("Error reading wan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtenderControllerExtenderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtenderControllerExtenderDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderAaaSharedSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderAccessPointName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAtDialScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderAuthorized(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderBillingStartDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaAaaSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaHaSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaNai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderBandwidthLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandExtenderControllerExtenderControllerReportInterval(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandExtenderControllerExtenderControllerReportSignalThreshold(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtenderControllerExtenderControllerReportStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderControllerReportInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderExtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderHaSharedSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderIfname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderInitiatedUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderLoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModemPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModemType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_sim_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["_sim_profile"], _ = expandExtenderControllerExtenderModem1SimProfile(d, i["_sim_profile"], pre_append)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtenderControllerExtenderModem1AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtenderControllerExtenderModem1ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtenderControllerExtenderModem1DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtenderControllerExtenderModem1Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ifname"], _ = expandExtenderControllerExtenderModem1Ifname(d, i["ifname"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtenderControllerExtenderModem1ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtenderControllerExtenderModem1PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtenderControllerExtenderModem1RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtenderControllerExtenderModem1RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtenderControllerExtenderModem1Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtenderControllerExtenderModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtenderControllerExtenderModem1Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtenderControllerExtenderModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtenderControllerExtenderModem1Status(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem1SimProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtenderControllerExtenderModem1AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtenderControllerExtenderModem1AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtenderControllerExtenderModem1AutoSwitchStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem1ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_sim_profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["_sim_profile"], _ = expandExtenderControllerExtenderModem2SimProfile(d, i["_sim_profile"], pre_append)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtenderControllerExtenderModem2AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtenderControllerExtenderModem2ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtenderControllerExtenderModem2DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtenderControllerExtenderModem2Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ifname"], _ = expandExtenderControllerExtenderModem2Ifname(d, i["ifname"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtenderControllerExtenderModem2ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtenderControllerExtenderModem2PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtenderControllerExtenderModem2RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtenderControllerExtenderModem2RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtenderControllerExtenderModem2Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtenderControllerExtenderModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtenderControllerExtenderModem2Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtenderControllerExtenderModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtenderControllerExtenderModem2Status(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem2SimProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtenderControllerExtenderModem2AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtenderControllerExtenderModem2AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtenderControllerExtenderModem2AutoSwitchStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem2ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderModem2Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderMultiMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppAuthProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppEchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderPppUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPrimaryHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderQuotaLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSecondaryHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSimPin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderOverrideEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderOverrideLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxAuthProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-extension"], _ = expandExtenderControllerExtenderWanExtensionModem1Extension(d, i["modem1_extension"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-extension"], _ = expandExtenderControllerExtenderWanExtensionModem2Extension(d, i["modem2_extension"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtenderWanExtensionModem1Extension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtenderControllerExtenderWanExtensionModem2Extension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectExtenderControllerExtender(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dataplan"); ok || d.HasChange("_dataplan") {
		t, err := expandExtenderControllerExtenderDataplan(d, v, "_dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dataplan"] = t
		}
	}

	if v, ok := d.GetOk("_template"); ok || d.HasChange("_template") {
		t, err := expandExtenderControllerExtenderTemplate(d, v, "_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_template"] = t
		}
	}

	if v, ok := d.GetOk("aaa_shared_secret"); ok || d.HasChange("aaa_shared_secret") {
		t, err := expandExtenderControllerExtenderAaaSharedSecret(d, v, "aaa_shared_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aaa-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("access_point_name"); ok || d.HasChange("access_point_name") {
		t, err := expandExtenderControllerExtenderAccessPointName(d, v, "access_point_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-point-name"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok || d.HasChange("admin") {
		t, err := expandExtenderControllerExtenderAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("at_dial_script"); ok || d.HasChange("at_dial_script") {
		t, err := expandExtenderControllerExtenderAtDialScript(d, v, "at_dial_script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["at-dial-script"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandExtenderControllerExtenderAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok || d.HasChange("authorized") {
		t, err := expandExtenderControllerExtenderAuthorized(d, v, "authorized")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("billing_start_day"); ok || d.HasChange("billing_start_day") {
		t, err := expandExtenderControllerExtenderBillingStartDay(d, v, "billing_start_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-start-day"] = t
		}
	}

	if v, ok := d.GetOk("cdma_aaa_spi"); ok || d.HasChange("cdma_aaa_spi") {
		t, err := expandExtenderControllerExtenderCdmaAaaSpi(d, v, "cdma_aaa_spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-aaa-spi"] = t
		}
	}

	if v, ok := d.GetOk("cdma_ha_spi"); ok || d.HasChange("cdma_ha_spi") {
		t, err := expandExtenderControllerExtenderCdmaHaSpi(d, v, "cdma_ha_spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-ha-spi"] = t
		}
	}

	if v, ok := d.GetOk("cdma_nai"); ok || d.HasChange("cdma_nai") {
		t, err := expandExtenderControllerExtenderCdmaNai(d, v, "cdma_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-nai"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok || d.HasChange("conn_status") {
		t, err := expandExtenderControllerExtenderConnStatus(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok || d.HasChange("bandwidth_limit") {
		t, err := expandExtenderControllerExtenderBandwidthLimit(d, v, "bandwidth_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("controller_report"); ok || d.HasChange("controller_report") {
		t, err := expandExtenderControllerExtenderControllerReport(d, v, "controller_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["controller-report"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandExtenderControllerExtenderDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dial_mode"); ok || d.HasChange("dial_mode") {
		t, err := expandExtenderControllerExtenderDialMode(d, v, "dial_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-mode"] = t
		}
	}

	if v, ok := d.GetOk("dial_status"); ok || d.HasChange("dial_status") {
		t, err := expandExtenderControllerExtenderDialStatus(d, v, "dial_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-status"] = t
		}
	}

	if v, ok := d.GetOk("device_id"); ok || d.HasChange("device_id") {
		t, err := expandExtenderControllerExtenderDeviceId(d, v, "device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok || d.HasChange("enforce_bandwidth") {
		t, err := expandExtenderControllerExtenderEnforceBandwidth(d, v, "enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok || d.HasChange("ext_name") {
		t, err := expandExtenderControllerExtenderExtName(d, v, "ext_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("ha_shared_secret"); ok || d.HasChange("ha_shared_secret") {
		t, err := expandExtenderControllerExtenderHaSharedSecret(d, v, "ha_shared_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("extension_type"); ok || d.HasChange("extension_type") {
		t, err := expandExtenderControllerExtenderExtensionType(d, v, "extension_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandExtenderControllerExtenderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ifname"); ok || d.HasChange("ifname") {
		t, err := expandExtenderControllerExtenderIfname(d, v, "ifname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ifname"] = t
		}
	}

	if v, ok := d.GetOk("initiated_update"); ok || d.HasChange("initiated_update") {
		t, err := expandExtenderControllerExtenderInitiatedUpdate(d, v, "initiated_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiated-update"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok || d.HasChange("login_password") {
		t, err := expandExtenderControllerExtenderLoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandExtenderControllerExtenderMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("modem_passwd"); ok || d.HasChange("modem_passwd") {
		t, err := expandExtenderControllerExtenderModemPasswd(d, v, "modem_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-passwd"] = t
		}
	}

	if v, ok := d.GetOk("modem_type"); ok || d.HasChange("modem_type") {
		t, err := expandExtenderControllerExtenderModemType(d, v, "modem_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-type"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok || d.HasChange("login_password_change") {
		t, err := expandExtenderControllerExtenderLoginPasswordChange(d, v, "login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("modem1"); ok || d.HasChange("modem1") {
		t, err := expandExtenderControllerExtenderModem1(d, v, "modem1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1"] = t
		}
	}

	if v, ok := d.GetOk("modem2"); ok || d.HasChange("modem2") {
		t, err := expandExtenderControllerExtenderModem2(d, v, "modem2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2"] = t
		}
	}

	if v, ok := d.GetOk("multi_mode"); ok || d.HasChange("multi_mode") {
		t, err := expandExtenderControllerExtenderMultiMode(d, v, "multi_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtenderControllerExtenderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ppp_auth_protocol"); ok || d.HasChange("ppp_auth_protocol") {
		t, err := expandExtenderControllerExtenderPppAuthProtocol(d, v, "ppp_auth_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ppp_echo_request"); ok || d.HasChange("ppp_echo_request") {
		t, err := expandExtenderControllerExtenderPppEchoRequest(d, v, "ppp_echo_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-echo-request"] = t
		}
	}

	if v, ok := d.GetOk("ppp_password"); ok || d.HasChange("ppp_password") {
		t, err := expandExtenderControllerExtenderPppPassword(d, v, "ppp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-password"] = t
		}
	}

	if v, ok := d.GetOk("ppp_username"); ok || d.HasChange("ppp_username") {
		t, err := expandExtenderControllerExtenderPppUsername(d, v, "ppp_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-username"] = t
		}
	}

	if v, ok := d.GetOk("primary_ha"); ok || d.HasChange("primary_ha") {
		t, err := expandExtenderControllerExtenderPrimaryHa(d, v, "primary_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-ha"] = t
		}
	}

	if v, ok := d.GetOk("quota_limit_mb"); ok || d.HasChange("quota_limit_mb") {
		t, err := expandExtenderControllerExtenderQuotaLimitMb(d, v, "quota_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("redial"); ok || d.HasChange("redial") {
		t, err := expandExtenderControllerExtenderRedial(d, v, "redial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redial"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok || d.HasChange("redundant_intf") {
		t, err := expandExtenderControllerExtenderRedundantIntf(d, v, "redundant_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOk("roaming"); ok || d.HasChange("roaming") {
		t, err := expandExtenderControllerExtenderRoaming(d, v, "roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandExtenderControllerExtenderRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ha"); ok || d.HasChange("secondary_ha") {
		t, err := expandExtenderControllerExtenderSecondaryHa(d, v, "secondary_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-ha"] = t
		}
	}

	if v, ok := d.GetOk("sim_pin"); ok || d.HasChange("sim_pin") {
		t, err := expandExtenderControllerExtenderSimPin(d, v, "sim_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-pin"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok || d.HasChange("override_allowaccess") {
		t, err := expandExtenderControllerExtenderOverrideAllowaccess(d, v, "override_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_enforce_bandwidth"); ok || d.HasChange("override_enforce_bandwidth") {
		t, err := expandExtenderControllerExtenderOverrideEnforceBandwidth(d, v, "override_enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("override_login_password_change"); ok || d.HasChange("override_login_password_change") {
		t, err := expandExtenderControllerExtenderOverrideLoginPasswordChange(d, v, "override_login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandExtenderControllerExtenderProfile(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandExtenderControllerExtenderVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("wimax_auth_protocol"); ok || d.HasChange("wimax_auth_protocol") {
		t, err := expandExtenderControllerExtenderWimaxAuthProtocol(d, v, "wimax_auth_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("wimax_carrier"); ok || d.HasChange("wimax_carrier") {
		t, err := expandExtenderControllerExtenderWimaxCarrier(d, v, "wimax_carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-carrier"] = t
		}
	}

	if v, ok := d.GetOk("wimax_realm"); ok || d.HasChange("wimax_realm") {
		t, err := expandExtenderControllerExtenderWimaxRealm(d, v, "wimax_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-realm"] = t
		}
	}

	if v, ok := d.GetOk("wan_extension"); ok || d.HasChange("wan_extension") {
		t, err := expandExtenderControllerExtenderWanExtension(d, v, "wan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-extension"] = t
		}
	}

	return &obj, nil
}
