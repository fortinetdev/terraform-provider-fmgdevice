// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure application control lists.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationList() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationListCreate,
		Read:   resourceApplicationListRead,
		Update: resourceApplicationListUpdate,
		Delete: resourceApplicationListDelete,

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
			"app_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"control_default_network_services": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deep_app_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_network_services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"services": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"violation_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"behavior": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"exclusion": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"members": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"per_ip_shaper": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"popularity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protocols": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rate_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"risk": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"session_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"shaper": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"shaper_reverse": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sub_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"technology": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_inclusion_ssl_di_sigs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"other_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_application_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"p2p_black_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"p2p_block_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"unknown_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_application_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceApplicationListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationList(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationList(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationList resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateApplicationList(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ApplicationList resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceApplicationListRead(d, m)
}

func resourceApplicationListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateApplicationList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceApplicationListRead(d, m)
}

func resourceApplicationListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteApplicationList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationList(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ApplicationList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationList resource from API: %v", err)
	}
	return nil
}

func flattenApplicationListAppReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListControlDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDeepAppInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenApplicationListDefaultNetworkServicesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationList-DefaultNetworkServices-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenApplicationListDefaultNetworkServicesPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ApplicationList-DefaultNetworkServices-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := i["services"]; ok {
			v := flattenApplicationListDefaultNetworkServicesServices(i["services"], d, pre_append)
			tmp["services"] = fortiAPISubPartPatch(v, "ApplicationList-DefaultNetworkServices-Services")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if _, ok := i["violation-action"]; ok {
			v := flattenApplicationListDefaultNetworkServicesViolationAction(i["violation-action"], d, pre_append)
			tmp["violation_action"] = fortiAPISubPartPatch(v, "ApplicationList-DefaultNetworkServices-ViolationAction")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListDefaultNetworkServicesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServicesPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServicesServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListDefaultNetworkServicesViolationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenApplicationListEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenApplicationListEntriesApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := i["behavior"]; ok {
			v := flattenApplicationListEntriesBehavior(i["behavior"], d, pre_append)
			tmp["behavior"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Behavior")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenApplicationListEntriesCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if _, ok := i["exclusion"]; ok {
			v := flattenApplicationListEntriesExclusion(i["exclusion"], d, pre_append)
			tmp["exclusion"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Exclusion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenApplicationListEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenApplicationListEntriesLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenApplicationListEntriesLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := i["parameters"]; ok {
			v := flattenApplicationListEntriesParameters(i["parameters"], d, pre_append)
			tmp["parameters"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Parameters")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := i["per-ip-shaper"]; ok {
			v := flattenApplicationListEntriesPerIpShaper(i["per-ip-shaper"], d, pre_append)
			tmp["per_ip_shaper"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-PerIpShaper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := i["popularity"]; ok {
			v := flattenApplicationListEntriesPopularity(i["popularity"], d, pre_append)
			tmp["popularity"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Popularity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := i["protocols"]; ok {
			v := flattenApplicationListEntriesProtocols(i["protocols"], d, pre_append)
			tmp["protocols"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Protocols")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenApplicationListEntriesQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenApplicationListEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenApplicationListEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			v := flattenApplicationListEntriesRateCount(i["rate-count"], d, pre_append)
			tmp["rate_count"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-RateCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			v := flattenApplicationListEntriesRateDuration(i["rate-duration"], d, pre_append)
			tmp["rate_duration"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-RateDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			v := flattenApplicationListEntriesRateMode(i["rate-mode"], d, pre_append)
			tmp["rate_mode"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-RateMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			v := flattenApplicationListEntriesRateTrack(i["rate-track"], d, pre_append)
			tmp["rate_track"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-RateTrack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := i["risk"]; ok {
			v := flattenApplicationListEntriesRisk(i["risk"], d, pre_append)
			tmp["risk"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Risk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := i["session-ttl"]; ok {
			v := flattenApplicationListEntriesSessionTtl(i["session-ttl"], d, pre_append)
			tmp["session_ttl"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-SessionTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := i["shaper"]; ok {
			v := flattenApplicationListEntriesShaper(i["shaper"], d, pre_append)
			tmp["shaper"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Shaper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := i["shaper-reverse"]; ok {
			v := flattenApplicationListEntriesShaperReverse(i["shaper-reverse"], d, pre_append)
			tmp["shaper_reverse"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-ShaperReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := i["sub-category"]; ok {
			v := flattenApplicationListEntriesSubCategory(i["sub-category"], d, pre_append)
			tmp["sub_category"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-SubCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := i["technology"]; ok {
			v := flattenApplicationListEntriesTechnology(i["technology"], d, pre_append)
			tmp["technology"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Technology")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := i["vendor"]; ok {
			v := flattenApplicationListEntriesVendor(i["vendor"], d, pre_append)
			tmp["vendor"] = fortiAPISubPartPatch(v, "ApplicationList-Entries-Vendor")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesExclusion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParameters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenApplicationListEntriesParametersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationListEntries-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenApplicationListEntriesParametersMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "ApplicationListEntries-Parameters-Members")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListEntriesParametersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenApplicationListEntriesParametersMembersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenApplicationListEntriesParametersMembersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenApplicationListEntriesParametersMembersValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListEntriesParametersMembersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesPopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesSubCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListForceInclusionSslDiSigs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListOtherApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListOtherApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListP2PBlackList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListP2PBlockList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListUnknownApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListUnknownApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("app_replacemsg", flattenApplicationListAppReplacemsg(o["app-replacemsg"], d, "app_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-replacemsg"], "ApplicationList-AppReplacemsg"); ok {
			if err = d.Set("app_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading app_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_replacemsg: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ApplicationList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("control_default_network_services", flattenApplicationListControlDefaultNetworkServices(o["control-default-network-services"], d, "control_default_network_services")); err != nil {
		if vv, ok := fortiAPIPatch(o["control-default-network-services"], "ApplicationList-ControlDefaultNetworkServices"); ok {
			if err = d.Set("control_default_network_services", vv); err != nil {
				return fmt.Errorf("Error reading control_default_network_services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading control_default_network_services: %v", err)
		}
	}

	if err = d.Set("deep_app_inspection", flattenApplicationListDeepAppInspection(o["deep-app-inspection"], d, "deep_app_inspection")); err != nil {
		if vv, ok := fortiAPIPatch(o["deep-app-inspection"], "ApplicationList-DeepAppInspection"); ok {
			if err = d.Set("deep_app_inspection", vv); err != nil {
				return fmt.Errorf("Error reading deep_app_inspection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deep_app_inspection: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("default_network_services", flattenApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services")); err != nil {
			if vv, ok := fortiAPIPatch(o["default-network-services"], "ApplicationList-DefaultNetworkServices"); ok {
				if err = d.Set("default_network_services", vv); err != nil {
					return fmt.Errorf("Error reading default_network_services: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading default_network_services: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("default_network_services"); ok {
			if err = d.Set("default_network_services", flattenApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services")); err != nil {
				if vv, ok := fortiAPIPatch(o["default-network-services"], "ApplicationList-DefaultNetworkServices"); ok {
					if err = d.Set("default_network_services", vv); err != nil {
						return fmt.Errorf("Error reading default_network_services: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading default_network_services: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_default_app_port", flattenApplicationListEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-default-app-port"], "ApplicationList-EnforceDefaultAppPort"); ok {
			if err = d.Set("enforce_default_app_port", vv); err != nil {
				return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ApplicationList-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ApplicationList-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenApplicationListExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ApplicationList-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("force_inclusion_ssl_di_sigs", flattenApplicationListForceInclusionSslDiSigs(o["force-inclusion-ssl-di-sigs"], d, "force_inclusion_ssl_di_sigs")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-inclusion-ssl-di-sigs"], "ApplicationList-ForceInclusionSslDiSigs"); ok {
			if err = d.Set("force_inclusion_ssl_di_sigs", vv); err != nil {
				return fmt.Errorf("Error reading force_inclusion_ssl_di_sigs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_inclusion_ssl_di_sigs: %v", err)
		}
	}

	if err = d.Set("name", flattenApplicationListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ApplicationList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenApplicationListOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ApplicationList-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("other_application_action", flattenApplicationListOtherApplicationAction(o["other-application-action"], d, "other_application_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["other-application-action"], "ApplicationList-OtherApplicationAction"); ok {
			if err = d.Set("other_application_action", vv); err != nil {
				return fmt.Errorf("Error reading other_application_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading other_application_action: %v", err)
		}
	}

	if err = d.Set("other_application_log", flattenApplicationListOtherApplicationLog(o["other-application-log"], d, "other_application_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["other-application-log"], "ApplicationList-OtherApplicationLog"); ok {
			if err = d.Set("other_application_log", vv); err != nil {
				return fmt.Errorf("Error reading other_application_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading other_application_log: %v", err)
		}
	}

	if err = d.Set("p2p_black_list", flattenApplicationListP2PBlackList(o["p2p-black-list"], d, "p2p_black_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["p2p-black-list"], "ApplicationList-P2PBlackList"); ok {
			if err = d.Set("p2p_black_list", vv); err != nil {
				return fmt.Errorf("Error reading p2p_black_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading p2p_black_list: %v", err)
		}
	}

	if err = d.Set("p2p_block_list", flattenApplicationListP2PBlockList(o["p2p-block-list"], d, "p2p_block_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["p2p-block-list"], "ApplicationList-P2PBlockList"); ok {
			if err = d.Set("p2p_block_list", vv); err != nil {
				return fmt.Errorf("Error reading p2p_block_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading p2p_block_list: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenApplicationListReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ApplicationList-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("unknown_application_action", flattenApplicationListUnknownApplicationAction(o["unknown-application-action"], d, "unknown_application_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-application-action"], "ApplicationList-UnknownApplicationAction"); ok {
			if err = d.Set("unknown_application_action", vv); err != nil {
				return fmt.Errorf("Error reading unknown_application_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_application_action: %v", err)
		}
	}

	if err = d.Set("unknown_application_log", flattenApplicationListUnknownApplicationLog(o["unknown-application-log"], d, "unknown_application_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-application-log"], "ApplicationList-UnknownApplicationLog"); ok {
			if err = d.Set("unknown_application_log", vv); err != nil {
				return fmt.Errorf("Error reading unknown_application_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_application_log: %v", err)
		}
	}

	return nil
}

func flattenApplicationListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationListAppReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListControlDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDeepAppInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandApplicationListDefaultNetworkServicesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandApplicationListDefaultNetworkServicesPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["services"], _ = expandApplicationListDefaultNetworkServicesServices(d, i["services"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["violation-action"], _ = expandApplicationListDefaultNetworkServicesViolationAction(d, i["violation_action"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListDefaultNetworkServicesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListDefaultNetworkServicesViolationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandApplicationListEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandApplicationListEntriesApplication(d, i["application"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["behavior"], _ = expandApplicationListEntriesBehavior(d, i["behavior"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandApplicationListEntriesCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclusion"], _ = expandApplicationListEntriesExclusion(d, i["exclusion"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandApplicationListEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandApplicationListEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandApplicationListEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandApplicationListEntriesParameters(d, i["parameters"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["parameters"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["per-ip-shaper"], _ = expandApplicationListEntriesPerIpShaper(d, i["per_ip_shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["popularity"], _ = expandApplicationListEntriesPopularity(d, i["popularity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocols"], _ = expandApplicationListEntriesProtocols(d, i["protocols"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandApplicationListEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandApplicationListEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandApplicationListEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-count"], _ = expandApplicationListEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-duration"], _ = expandApplicationListEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-mode"], _ = expandApplicationListEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-track"], _ = expandApplicationListEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["risk"], _ = expandApplicationListEntriesRisk(d, i["risk"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["session-ttl"], _ = expandApplicationListEntriesSessionTtl(d, i["session_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shaper"], _ = expandApplicationListEntriesShaper(d, i["shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shaper-reverse"], _ = expandApplicationListEntriesShaperReverse(d, i["shaper_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sub-category"], _ = expandApplicationListEntriesSubCategory(d, i["sub_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["technology"], _ = expandApplicationListEntriesTechnology(d, i["technology"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vendor"], _ = expandApplicationListEntriesVendor(d, i["vendor"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesExclusion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandApplicationListEntriesParametersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandApplicationListEntriesParametersMembers(d, i["members"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["members"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandApplicationListEntriesParametersMembersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandApplicationListEntriesParametersMembersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandApplicationListEntriesParametersMembersValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersMembersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesPopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListForceInclusionSslDiSigs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListOtherApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOtherApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListP2PBlackList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListP2PBlockList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListUnknownApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListUnknownApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_replacemsg"); ok || d.HasChange("app_replacemsg") {
		t, err := expandApplicationListAppReplacemsg(d, v, "app_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandApplicationListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("control_default_network_services"); ok || d.HasChange("control_default_network_services") {
		t, err := expandApplicationListControlDefaultNetworkServices(d, v, "control_default_network_services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-default-network-services"] = t
		}
	}

	if v, ok := d.GetOk("deep_app_inspection"); ok || d.HasChange("deep_app_inspection") {
		t, err := expandApplicationListDeepAppInspection(d, v, "deep_app_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-inspection"] = t
		}
	}

	if v, ok := d.GetOk("default_network_services"); ok || d.HasChange("default_network_services") {
		t, err := expandApplicationListDefaultNetworkServices(d, v, "default_network_services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-network-services"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok || d.HasChange("enforce_default_app_port") {
		t, err := expandApplicationListEnforceDefaultAppPort(d, v, "enforce_default_app_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandApplicationListEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandApplicationListExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("force_inclusion_ssl_di_sigs"); ok || d.HasChange("force_inclusion_ssl_di_sigs") {
		t, err := expandApplicationListForceInclusionSslDiSigs(d, v, "force_inclusion_ssl_di_sigs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-inclusion-ssl-di-sigs"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandApplicationListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandApplicationListOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("other_application_action"); ok || d.HasChange("other_application_action") {
		t, err := expandApplicationListOtherApplicationAction(d, v, "other_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-action"] = t
		}
	}

	if v, ok := d.GetOk("other_application_log"); ok || d.HasChange("other_application_log") {
		t, err := expandApplicationListOtherApplicationLog(d, v, "other_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-log"] = t
		}
	}

	if v, ok := d.GetOk("p2p_black_list"); ok || d.HasChange("p2p_black_list") {
		t, err := expandApplicationListP2PBlackList(d, v, "p2p_black_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-black-list"] = t
		}
	}

	if v, ok := d.GetOk("p2p_block_list"); ok || d.HasChange("p2p_block_list") {
		t, err := expandApplicationListP2PBlockList(d, v, "p2p_block_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-block-list"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandApplicationListReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_action"); ok || d.HasChange("unknown_application_action") {
		t, err := expandApplicationListUnknownApplicationAction(d, v, "unknown_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-action"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_log"); ok || d.HasChange("unknown_application_log") {
		t, err := expandApplicationListUnknownApplicationLog(d, v, "unknown_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-log"] = t
		}
	}

	return &obj, nil
}
