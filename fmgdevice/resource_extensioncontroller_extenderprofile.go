// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender extender profile configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCreate,
		Read:   resourceExtensionControllerExtenderProfileRead,
		Update: resourceExtensionControllerExtenderProfileUpdate,
		Delete: resourceExtensionControllerExtenderProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"_is_factory_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cellular": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"backhaul_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"backhaul_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"downlinks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pvid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vids": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_loadbalance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"traffic_split_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"service": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"vsdb": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"login_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"country": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_1": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"n80211d": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"band": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bandwidth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"beacon_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"bss_color": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"bss_color_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"channel": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"extension_channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"guard_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"lan_ext_vap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"local_vaps": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"max_clients": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"operating_standard": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_level": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"radio_id": &schema.Schema{
										Type:     schema.TypeInt,
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
						"radio_2": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"n80211d": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"band": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bandwidth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"beacon_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"bss_color": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"bss_color_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"channel": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"extension_channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"guard_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"lan_ext_vap": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"local_vaps": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"max_clients": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"operating_standard": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_level": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"radio_id": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectExtensionControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadExtensionControllerExtenderProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateExtensionControllerExtenderProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ExtensionControllerExtenderProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateExtensionControllerExtenderProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ExtensionControllerExtenderProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileRead(d, m)
}

func resourceExtensionControllerExtenderProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectExtensionControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateExtensionControllerExtenderProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderProfileRead(d, m)
}

func resourceExtensionControllerExtenderProfileDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteExtensionControllerExtenderProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfile resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellular(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := i["controller-report"]; ok {
		result["controller_report"] = flattenExtensionControllerExtenderProfileCellularControllerReport(i["controller-report"], d, pre_append)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1"
	if _, ok := i["modem1"]; ok {
		result["modem1"] = flattenExtensionControllerExtenderProfileCellularModem1(i["modem1"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2"
	if _, ok := i["modem2"]; ok {
		result["modem2"] = flattenExtensionControllerExtenderProfileCellularModem2(i["modem2"], d, pre_append)
	}

	pre_append = pre + ".0." + "sms_notification"
	if _, ok := i["sms-notification"]; ok {
		result["sms_notification"] = flattenExtensionControllerExtenderProfileCellularSmsNotification(i["sms-notification"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularControllerReport(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtensionControllerExtenderProfileCellularControllerReportInterval(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularControllerReportStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularControllerReportInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem1ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem1DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem1Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtensionControllerExtenderProfileCellularModem1ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn(i["multiple-PDN"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(i["pdn1-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(i["pdn2-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(i["pdn3-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(i["pdn4-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem2ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem2DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem2Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenExtensionControllerExtenderProfileCellularModem2ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn(i["multiple-PDN"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(i["pdn1-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(i["pdn2-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(i["pdn3-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(i["pdn4-dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotification(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert(i["alert"], d, pre_append)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver(i["receiver"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileExtension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := i["backhaul"]; ok {
		result["backhaul"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaul(i["backhaul"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "downlinks"
	if _, ok := i["downlinks"]; ok {
		result["downlinks"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinks(i["downlinks"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenExtensionControllerExtenderProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := i["link-loadbalance"]; ok {
		result["link_loadbalance"] = flattenExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(i["link-loadbalance"], d, pre_append)
	}

	pre_append = pre + ".0." + "traffic_split_services"
	if _, ok := i["traffic-split-services"]; ok {
		result["traffic_split_services"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(i["traffic-split-services"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaul(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionBackhaulName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Backhaul-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Backhaul-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Backhaul-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Backhaul-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if _, ok := i["pvid"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid(i["pvid"], d, pre_append)
			tmp["pvid"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Pvid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if _, ok := i["vap"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap(i["vap"], d, pre_append)
			tmp["vap"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Vap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vids"
		if _, ok := i["vids"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionDownlinksVids(i["vids"], d, pre_append)
			tmp["vids"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-Downlinks-Vids")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksVids(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(i["address"], d, pre_append)
			tmp["address"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-TrafficSplitServices-Address")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-TrafficSplitServices-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-TrafficSplitServices-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vsdb"
		if _, ok := i["vsdb"]; ok {
			v := flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(i["vsdb"], d, pre_append)
			tmp["vsdb"] = fortiAPISubPartPatch(v, "ExtensionControllerExtenderProfileLanExtension-TrafficSplitServices-Vsdb")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dfs"
	if _, ok := i["DFS"]; ok {
		result["dfs"] = flattenExtensionControllerExtenderProfileWifiDfs(i["DFS"], d, pre_append)
	}

	pre_append = pre + ".0." + "country"
	if _, ok := i["country"]; ok {
		result["country"] = flattenExtensionControllerExtenderProfileWifiCountry(i["country"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_1"
	if _, ok := i["radio-1"]; ok {
		result["radio_1"] = flattenExtensionControllerExtenderProfileWifiRadio1(i["radio-1"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_2"
	if _, ok := i["radio-2"]; ok {
		result["radio_2"] = flattenExtensionControllerExtenderProfileWifiRadio2(i["radio-2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiDfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenExtensionControllerExtenderProfileWifiRadio180211D(i["80211d"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenExtensionControllerExtenderProfileWifiRadio1Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth(i["bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenExtensionControllerExtenderProfileWifiRadio1BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode(i["bss-color-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenExtensionControllerExtenderProfileWifiRadio1Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(i["extension-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval(i["guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap(i["lan-ext-vap"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps(i["local-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenExtensionControllerExtenderProfileWifiRadio1MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenExtensionControllerExtenderProfileWifiRadio1Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard(i["operating-standard"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenExtensionControllerExtenderProfileWifiRadio1RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileWifiRadio1Status(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiRadio180211D(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenExtensionControllerExtenderProfileWifiRadio280211D(i["80211d"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenExtensionControllerExtenderProfileWifiRadio2Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenExtensionControllerExtenderProfileWifiRadio2Bandwidth(i["bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenExtensionControllerExtenderProfileWifiRadio2BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenExtensionControllerExtenderProfileWifiRadio2BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenExtensionControllerExtenderProfileWifiRadio2BssColorMode(i["bss-color-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenExtensionControllerExtenderProfileWifiRadio2Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(i["extension-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenExtensionControllerExtenderProfileWifiRadio2GuardInterval(i["guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenExtensionControllerExtenderProfileWifiRadio2LanExtVap(i["lan-ext-vap"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenExtensionControllerExtenderProfileWifiRadio2LocalVaps(i["local-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenExtensionControllerExtenderProfileWifiRadio2MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenExtensionControllerExtenderProfileWifiRadio2Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenExtensionControllerExtenderProfileWifiRadio2OperatingStandard(i["operating-standard"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenExtensionControllerExtenderProfileWifiRadio2PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenExtensionControllerExtenderProfileWifiRadio2RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileWifiRadio2Status(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiRadio280211D(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Bandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2BssColorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2GuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2LanExtVap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2LocalVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2OperatingStandard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_is_factory_setting", flattenExtensionControllerExtenderProfileIsFactorySetting(o["_is_factory_setting"], d, "_is_factory_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["_is_factory_setting"], "ExtensionControllerExtenderProfile-IsFactorySetting"); ok {
			if err = d.Set("_is_factory_setting", vv); err != nil {
				return fmt.Errorf("Error reading _is_factory_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _is_factory_setting: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtensionControllerExtenderProfileAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ExtensionControllerExtenderProfile-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtensionControllerExtenderProfileBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-limit"], "ExtensionControllerExtenderProfile-BandwidthLimit"); ok {
			if err = d.Set("bandwidth_limit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cellular", flattenExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
			if vv, ok := fortiAPIPatch(o["cellular"], "ExtensionControllerExtenderProfile-Cellular"); ok {
				if err = d.Set("cellular", vv); err != nil {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cellular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cellular"); ok {
			if err = d.Set("cellular", flattenExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
				if vv, ok := fortiAPIPatch(o["cellular"], "ExtensionControllerExtenderProfile-Cellular"); ok {
					if err = d.Set("cellular", vv); err != nil {
						return fmt.Errorf("Error reading cellular: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtensionControllerExtenderProfileEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-bandwidth"], "ExtensionControllerExtenderProfile-EnforceBandwidth"); ok {
			if err = d.Set("enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("extension", flattenExtensionControllerExtenderProfileExtension(o["extension"], d, "extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension"], "ExtensionControllerExtenderProfile-Extension"); ok {
			if err = d.Set("extension", vv); err != nil {
				return fmt.Errorf("Error reading extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerExtenderProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ExtensionControllerExtenderProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan_extension", flattenExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan-extension"], "ExtensionControllerExtenderProfile-LanExtension"); ok {
				if err = d.Set("lan_extension", vv); err != nil {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan-extension"], "ExtensionControllerExtenderProfile-LanExtension"); ok {
					if err = d.Set("lan_extension", vv); err != nil {
						return fmt.Errorf("Error reading lan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("login_password_change", flattenExtensionControllerExtenderProfileLoginPasswordChange(o["login-password-change"], d, "login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password-change"], "ExtensionControllerExtenderProfile-LoginPasswordChange"); ok {
			if err = d.Set("login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("model", flattenExtensionControllerExtenderProfileModel(o["model"], d, "model")); err != nil {
		if vv, ok := fortiAPIPatch(o["model"], "ExtensionControllerExtenderProfile-Model"); ok {
			if err = d.Set("model", vv); err != nil {
				return fmt.Errorf("Error reading model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("name", flattenExtensionControllerExtenderProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtenderProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wifi", flattenExtensionControllerExtenderProfileWifi(o["wifi"], d, "wifi")); err != nil {
			if vv, ok := fortiAPIPatch(o["wifi"], "ExtensionControllerExtenderProfile-Wifi"); ok {
				if err = d.Set("wifi", vv); err != nil {
					return fmt.Errorf("Error reading wifi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading wifi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wifi"); ok {
			if err = d.Set("wifi", flattenExtensionControllerExtenderProfileWifi(o["wifi"], d, "wifi")); err != nil {
				if vv, ok := fortiAPIPatch(o["wifi"], "ExtensionControllerExtenderProfile-Wifi"); ok {
					if err = d.Set("wifi", vv); err != nil {
						return fmt.Errorf("Error reading wifi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading wifi: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderProfileIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileBandwidthLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularControllerReport(d, i["controller_report"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["controller-report"] = t
		}
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "modem1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem1(d, i["modem1"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem1"] = t
		}
	}
	pre_append = pre + ".0." + "modem2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem2(d, i["modem2"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem2"] = t
		}
	}
	pre_append = pre + ".0." + "sms_notification"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotification(d, i["sms_notification"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["sms-notification"] = t
		}
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandExtensionControllerExtenderProfileCellularControllerReportInterval(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularControllerReportStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem1AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem1ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem1DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem1Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtensionControllerExtenderProfileCellularModem1ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem1MultiplePdn(d, i["multiple_pdn"], pre_append)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(d, i["pdn1_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(d, i["pdn2_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(d, i["pdn3_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(d, i["pdn4_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1MultiplePdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularModem2AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem2ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem2DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem2Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandExtensionControllerExtenderProfileCellularModem2ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem2MultiplePdn(d, i["multiple_pdn"], pre_append)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(d, i["pdn1_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(d, i["pdn2_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(d, i["pdn3_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(d, i["pdn4_dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2MultiplePdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationAlert(d, i["alert"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["alert"] = t
		}
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d, i["receiver"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["receiver"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileLanExtensionBackhaul(d, i["backhaul"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["backhaul"] = t
		}
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-interface"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-ip"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append)
	}
	pre_append = pre + ".0." + "downlinks"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileLanExtensionDownlinks(d, i["downlinks"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["downlinks"] = t
		}
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-tunnel"], _ = expandExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-loadbalance"], _ = expandExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d, i["link_loadbalance"], pre_append)
	}
	pre_append = pre + ".0." + "traffic_split_services"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d, i["traffic_split_services"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["traffic-split-services"] = t
		}
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pvid"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid(d, i["pvid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vap"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksVap(d, i["vap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vids"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vids"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksVids(d, i["vids"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksVap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksVids(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["address"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vsdb"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vsdb"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(d, i["vsdb"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dfs"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["DFS"], _ = expandExtensionControllerExtenderProfileWifiDfs(d, i["dfs"], pre_append)
	}
	pre_append = pre + ".0." + "country"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["country"], _ = expandExtensionControllerExtenderProfileWifiCountry(d, i["country"], pre_append)
	}
	pre_append = pre + ".0." + "radio_1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileWifiRadio1(d, i["radio_1"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["radio-1"] = t
		}
	}
	pre_append = pre + ".0." + "radio_2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandExtensionControllerExtenderProfileWifiRadio2(d, i["radio_2"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["radio-2"] = t
		}
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiDfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["80211d"], _ = expandExtensionControllerExtenderProfileWifiRadio180211D(d, i["n80211d"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandExtensionControllerExtenderProfileWifiRadio1Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bandwidth"], _ = expandExtensionControllerExtenderProfileWifiRadio1Bandwidth(d, i["bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["beacon-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color"], _ = expandExtensionControllerExtenderProfileWifiRadio1BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color-mode"], _ = expandExtensionControllerExtenderProfileWifiRadio1BssColorMode(d, i["bss_color_mode"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandExtensionControllerExtenderProfileWifiRadio1Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["extension-channel"], _ = expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(d, i["extension_channel"], pre_append)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["guard-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio1GuardInterval(d, i["guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-ext-vap"], _ = expandExtensionControllerExtenderProfileWifiRadio1LanExtVap(d, i["lan_ext_vap"], pre_append)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-vaps"], _ = expandExtensionControllerExtenderProfileWifiRadio1LocalVaps(d, i["local_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-clients"], _ = expandExtensionControllerExtenderProfileWifiRadio1MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandExtensionControllerExtenderProfileWifiRadio1Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["operating-standard"], _ = expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard(d, i["operating_standard"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandExtensionControllerExtenderProfileWifiRadio1PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandExtensionControllerExtenderProfileWifiRadio1RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileWifiRadio1Status(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio180211D(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Bandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1GuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LanExtVap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LocalVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio1MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["80211d"], _ = expandExtensionControllerExtenderProfileWifiRadio280211D(d, i["n80211d"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandExtensionControllerExtenderProfileWifiRadio2Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bandwidth"], _ = expandExtensionControllerExtenderProfileWifiRadio2Bandwidth(d, i["bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["beacon-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio2BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color"], _ = expandExtensionControllerExtenderProfileWifiRadio2BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color-mode"], _ = expandExtensionControllerExtenderProfileWifiRadio2BssColorMode(d, i["bss_color_mode"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandExtensionControllerExtenderProfileWifiRadio2Channel(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["extension-channel"], _ = expandExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(d, i["extension_channel"], pre_append)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["guard-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio2GuardInterval(d, i["guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-ext-vap"], _ = expandExtensionControllerExtenderProfileWifiRadio2LanExtVap(d, i["lan_ext_vap"], pre_append)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-vaps"], _ = expandExtensionControllerExtenderProfileWifiRadio2LocalVaps(d, i["local_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-clients"], _ = expandExtensionControllerExtenderProfileWifiRadio2MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandExtensionControllerExtenderProfileWifiRadio2Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["operating-standard"], _ = expandExtensionControllerExtenderProfileWifiRadio2OperatingStandard(d, i["operating_standard"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandExtensionControllerExtenderProfileWifiRadio2PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandExtensionControllerExtenderProfileWifiRadio2RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandExtensionControllerExtenderProfileWifiRadio2Status(d, i["status"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio280211D(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Bandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BssColorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2GuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2LanExtVap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio2LocalVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderProfileWifiRadio2MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2OperatingStandard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_is_factory_setting"); ok || d.HasChange("_is_factory_setting") {
		t, err := expandExtensionControllerExtenderProfileIsFactorySetting(d, v, "_is_factory_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_is_factory_setting"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandExtensionControllerExtenderProfileAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok || d.HasChange("bandwidth_limit") {
		t, err := expandExtensionControllerExtenderProfileBandwidthLimit(d, v, "bandwidth_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("cellular"); ok || d.HasChange("cellular") {
		t, err := expandExtensionControllerExtenderProfileCellular(d, v, "cellular")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cellular"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok || d.HasChange("enforce_bandwidth") {
		t, err := expandExtensionControllerExtenderProfileEnforceBandwidth(d, v, "enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok || d.HasChange("extension") {
		t, err := expandExtensionControllerExtenderProfileExtension(d, v, "extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandExtensionControllerExtenderProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok || d.HasChange("lan_extension") {
		t, err := expandExtensionControllerExtenderProfileLanExtension(d, v, "lan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok || d.HasChange("login_password") {
		t, err := expandExtensionControllerExtenderProfileLoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok || d.HasChange("login_password_change") {
		t, err := expandExtensionControllerExtenderProfileLoginPasswordChange(d, v, "login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok || d.HasChange("model") {
		t, err := expandExtensionControllerExtenderProfileModel(d, v, "model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("wifi"); ok || d.HasChange("wifi") {
		t, err := expandExtensionControllerExtenderProfileWifi(d, v, "wifi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi"] = t
		}
	}

	return &obj, nil
}
