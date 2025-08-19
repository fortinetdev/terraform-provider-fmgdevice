// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure OSPF.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfUpdate,
		Read:   resourceRouterOspfRead,
		Update: resourceRouterOspfUpdate,
		Delete: resourceRouterOspfDelete,

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
			"abr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"list": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nssa_default_information_originate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_redistribution": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_translator_role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"prefix": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"substitute": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"substitute_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stub_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_link": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"authentication_key": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"dead_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"hello_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"keychain": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"md5_keychain": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"md5_keys": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"key_string": &schema.Schema{
													Type:      schema.TypeSet,
													Elem:      &schema.Schema{Type: schema.TypeString},
													Optional:  true,
													Sensitive: true,
													Computed:  true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"peer": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"retransmit_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"transmit_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"auto_cost_ref_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_overflow_max_lsas": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"database_overflow_time_to_recover": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_route_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance_inter_area": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance_intra_area": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"distribute_list_in": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"distribute_route_map_in": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lsa_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poll_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ospf_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"database_filter_out": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dead_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"hello_multiplier": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keychain": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"linkdown_fast_failover": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"md5_keychain": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"md5_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"key_string": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"resync_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retransmit_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_on_topology_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rfc1583_compatible": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spf_timers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterOspfUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspf(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterOspf")

	return resourceRouterOspfRead(d, m)
}

func resourceRouterOspfDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspf(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfArea(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenRouterOspfAreaAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenRouterOspfAreaComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			v := flattenRouterOspfAreaDefaultCost(i["default-cost"], d, pre_append)
			tmp["default_cost"] = fortiAPISubPartPatch(v, "RouterOspf-Area-DefaultCost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := i["filter-list"]; ok {
			v := flattenRouterOspfAreaFilterList(i["filter-list"], d, pre_append)
			tmp["filter_list"] = fortiAPISubPartPatch(v, "RouterOspf-Area-FilterList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfAreaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			v := flattenRouterOspfAreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
			tmp["nssa_default_information_originate"] = fortiAPISubPartPatch(v, "RouterOspf-Area-NssaDefaultInformationOriginate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			v := flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
			tmp["nssa_default_information_originate_metric"] = fortiAPISubPartPatch(v, "RouterOspf-Area-NssaDefaultInformationOriginateMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			v := flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
			tmp["nssa_default_information_originate_metric_type"] = fortiAPISubPartPatch(v, "RouterOspf-Area-NssaDefaultInformationOriginateMetricType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			v := flattenRouterOspfAreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
			tmp["nssa_redistribution"] = fortiAPISubPartPatch(v, "RouterOspf-Area-NssaRedistribution")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			v := flattenRouterOspfAreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
			tmp["nssa_translator_role"] = fortiAPISubPartPatch(v, "RouterOspf-Area-NssaTranslatorRole")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			v := flattenRouterOspfAreaRange(i["range"], d, pre_append)
			tmp["range"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Range")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			v := flattenRouterOspfAreaShortcut(i["shortcut"], d, pre_append)
			tmp["shortcut"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Shortcut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			v := flattenRouterOspfAreaStubType(i["stub-type"], d, pre_append)
			tmp["stub_type"] = fortiAPISubPartPatch(v, "RouterOspf-Area-StubType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenRouterOspfAreaType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "RouterOspf-Area-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			v := flattenRouterOspfAreaVirtualLink(i["virtual-link"], d, pre_append)
			tmp["virtual_link"] = fortiAPISubPartPatch(v, "RouterOspf-Area-VirtualLink")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenRouterOspfAreaFilterListDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfAreaFilterListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			v := flattenRouterOspfAreaFilterListList(i["list"], d, pre_append)
			tmp["list"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-List")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			v := flattenRouterOspfAreaRangeAdvertise(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfAreaRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterOspfAreaRangePrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			v := flattenRouterOspfAreaRangeSubstitute(i["substitute"], d, pre_append)
			tmp["substitute"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Substitute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			v := flattenRouterOspfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
			tmp["substitute_status"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-SubstituteStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenRouterOspfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {
			v := flattenRouterOspfAreaVirtualLinkKeychain(i["keychain"], d, pre_append)
			tmp["keychain"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			v := flattenRouterOspfAreaVirtualLinkMd5Keychain(i["md5-keychain"], d, pre_append)
			tmp["md5_keychain"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Md5Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			v := flattenRouterOspfAreaVirtualLinkMd5Keys(i["md5-keys"], d, pre_append)
			tmp["md5_keys"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Md5Keys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspfAreaVirtualLinkName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			v := flattenRouterOspfAreaVirtualLinkPeer(i["peer"], d, pre_append)
			tmp["peer"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Peer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaVirtualLinkMd5KeysId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfAreaVirtualLink-Md5Keys-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflowMaxLsas(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			v := flattenRouterOspfDistributeListAccessList(i["access-list"], d, pre_append)
			tmp["access_list"] = fortiAPISubPartPatch(v, "RouterOspf-DistributeList-AccessList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfDistributeListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf-DistributeList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenRouterOspfDistributeListProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "RouterOspf-DistributeList-Protocol")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfDistributeRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfLsaRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenRouterOspfNeighborCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "RouterOspf-Neighbor-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfNeighborId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf-Neighbor-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenRouterOspfNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "RouterOspf-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			v := flattenRouterOspfNeighborPollInterval(i["poll-interval"], d, pre_append)
			tmp["poll_interval"] = fortiAPISubPartPatch(v, "RouterOspf-Neighbor-PollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenRouterOspfNeighborPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "RouterOspf-Neighbor-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			v := flattenRouterOspfNetworkArea(i["area"], d, pre_append)
			tmp["area"] = fortiAPISubPartPatch(v, "RouterOspf-Network-Area")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenRouterOspfNetworkComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "RouterOspf-Network-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfNetworkId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf-Network-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterOspfNetworkPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterOspf-Network-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfNetworkArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetworkComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenRouterOspfOspfInterfaceAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			v := flattenRouterOspfOspfInterfaceBfd(i["bfd"], d, pre_append)
			tmp["bfd"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Bfd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenRouterOspfOspfInterfaceComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenRouterOspfOspfInterfaceCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := i["database-filter-out"]; ok {
			v := flattenRouterOspfOspfInterfaceDatabaseFilterOut(i["database-filter-out"], d, pre_append)
			tmp["database_filter_out"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-DatabaseFilterOut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspfOspfInterfaceDeadInterval(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspfOspfInterfaceHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := i["hello-multiplier"]; ok {
			v := flattenRouterOspfOspfInterfaceHelloMultiplier(i["hello-multiplier"], d, pre_append)
			tmp["hello_multiplier"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-HelloMultiplier")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterOspfOspfInterfaceInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenRouterOspfOspfInterfaceIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {
			v := flattenRouterOspfOspfInterfaceKeychain(i["keychain"], d, pre_append)
			tmp["keychain"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "linkdown_fast_failover"
		if _, ok := i["linkdown-fast-failover"]; ok {
			v := flattenRouterOspfOspfInterfaceLinkdownFastFailover(i["linkdown-fast-failover"], d, pre_append)
			tmp["linkdown_fast_failover"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-LinkdownFastFailover")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			v := flattenRouterOspfOspfInterfaceMd5Keychain(i["md5-keychain"], d, pre_append)
			tmp["md5_keychain"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Md5Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			v := flattenRouterOspfOspfInterfaceMd5Keys(i["md5-keys"], d, pre_append)
			tmp["md5_keys"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Md5Keys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			v := flattenRouterOspfOspfInterfaceMtu(i["mtu"], d, pre_append)
			tmp["mtu"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Mtu")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			v := flattenRouterOspfOspfInterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
			tmp["mtu_ignore"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-MtuIgnore")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspfOspfInterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			v := flattenRouterOspfOspfInterfaceNetworkType(i["network-type"], d, pre_append)
			tmp["network_type"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-NetworkType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := i["prefix-length"]; ok {
			v := flattenRouterOspfOspfInterfacePrefixLength(i["prefix-length"], d, pre_append)
			tmp["prefix_length"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-PrefixLength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenRouterOspfOspfInterfacePriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := i["resync-timeout"]; ok {
			v := flattenRouterOspfOspfInterfaceResyncTimeout(i["resync-timeout"], d, pre_append)
			tmp["resync_timeout"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-ResyncTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspfOspfInterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterOspfOspfInterfaceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspfOspfInterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspf-OspfInterface-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceLinkdownFastFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfOspfInterfaceMd5KeysId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfOspfInterface-Md5Keys-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := i["metric"]; ok {
		result["metric"] = flattenRouterOspfRedistributeMetric(i["metric"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric_type"
	if _, ok := i["metric-type"]; ok {
		result["metric_type"] = flattenRouterOspfRedistributeMetricType(i["metric-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenRouterOspfRedistributeName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "routemap"
	if _, ok := i["routemap"]; ok {
		result["routemap"] = flattenRouterOspfRedistributeRoutemap(i["routemap"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenRouterOspfRedistributeStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenRouterOspfRedistributeTag(i["tag"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterOspfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRestartOnTopologyChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRestartPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenRouterOspfSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			v := flattenRouterOspfSummaryAddressAdvertise(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspf-SummaryAddress-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfSummaryAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf-SummaryAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterOspfSummaryAddressPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterOspf-SummaryAddress-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			v := flattenRouterOspfSummaryAddressTag(i["tag"], d, pre_append)
			tmp["tag"] = fortiAPISubPartPatch(v, "RouterOspf-SummaryAddress-Tag")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfSummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("abr_type", flattenRouterOspfAbrType(o["abr-type"], d, "abr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["abr-type"], "RouterOspf-AbrType"); ok {
			if err = d.Set("abr_type", vv); err != nil {
				return fmt.Errorf("Error reading abr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area")); err != nil {
			if vv, ok := fortiAPIPatch(o["area"], "RouterOspf-Area"); ok {
				if err = d.Set("area", vv); err != nil {
					return fmt.Errorf("Error reading area: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading area: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("area"); ok {
			if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area")); err != nil {
				if vv, ok := fortiAPIPatch(o["area"], "RouterOspf-Area"); ok {
					if err = d.Set("area", vv); err != nil {
						return fmt.Errorf("Error reading area: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading area: %v", err)
				}
			}
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspfAutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-cost-ref-bandwidth"], "RouterOspf-AutoCostRefBandwidth"); ok {
			if err = d.Set("auto_cost_ref_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspfBfd(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterOspf-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("database_overflow", flattenRouterOspfDatabaseOverflow(o["database-overflow"], d, "database_overflow")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-overflow"], "RouterOspf-DatabaseOverflow"); ok {
			if err = d.Set("database_overflow", vv); err != nil {
				return fmt.Errorf("Error reading database_overflow: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_overflow: %v", err)
		}
	}

	if err = d.Set("database_overflow_max_lsas", flattenRouterOspfDatabaseOverflowMaxLsas(o["database-overflow-max-lsas"], d, "database_overflow_max_lsas")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-overflow-max-lsas"], "RouterOspf-DatabaseOverflowMaxLsas"); ok {
			if err = d.Set("database_overflow_max_lsas", vv); err != nil {
				return fmt.Errorf("Error reading database_overflow_max_lsas: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_overflow_max_lsas: %v", err)
		}
	}

	if err = d.Set("database_overflow_time_to_recover", flattenRouterOspfDatabaseOverflowTimeToRecover(o["database-overflow-time-to-recover"], d, "database_overflow_time_to_recover")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-overflow-time-to-recover"], "RouterOspf-DatabaseOverflowTimeToRecover"); ok {
			if err = d.Set("database_overflow_time_to_recover", vv); err != nil {
				return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspfDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-metric"], "RouterOspf-DefaultInformationMetric"); ok {
			if err = d.Set("default_information_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_information_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspfDefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-metric-type"], "RouterOspf-DefaultInformationMetricType"); ok {
			if err = d.Set("default_information_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading default_information_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspfDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-originate"], "RouterOspf-DefaultInformationOriginate"); ok {
			if err = d.Set("default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspfDefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-route-map"], "RouterOspf-DefaultInformationRouteMap"); ok {
			if err = d.Set("default_information_route_map", vv); err != nil {
				return fmt.Errorf("Error reading default_information_route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspfDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-metric"], "RouterOspf-DefaultMetric"); ok {
			if err = d.Set("default_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterOspfDistance(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "RouterOspf-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterOspfDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-external"], "RouterOspf-DistanceExternal"); ok {
			if err = d.Set("distance_external", vv); err != nil {
				return fmt.Errorf("Error reading distance_external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_inter_area", flattenRouterOspfDistanceInterArea(o["distance-inter-area"], d, "distance_inter_area")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-inter-area"], "RouterOspf-DistanceInterArea"); ok {
			if err = d.Set("distance_inter_area", vv); err != nil {
				return fmt.Errorf("Error reading distance_inter_area: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_inter_area: %v", err)
		}
	}

	if err = d.Set("distance_intra_area", flattenRouterOspfDistanceIntraArea(o["distance-intra-area"], d, "distance_intra_area")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance-intra-area"], "RouterOspf-DistanceIntraArea"); ok {
			if err = d.Set("distance_intra_area", vv); err != nil {
				return fmt.Errorf("Error reading distance_intra_area: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance_intra_area: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterOspf-DistributeList"); ok {
				if err = d.Set("distribute_list", vv); err != nil {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterOspf-DistributeList"); ok {
					if err = d.Set("distribute_list", vv); err != nil {
						return fmt.Errorf("Error reading distribute_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterOspfDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-list-in"], "RouterOspf-DistributeListIn"); ok {
			if err = d.Set("distribute_list_in", vv); err != nil {
				return fmt.Errorf("Error reading distribute_list_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_route_map_in", flattenRouterOspfDistributeRouteMapIn(o["distribute-route-map-in"], d, "distribute_route_map_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["distribute-route-map-in"], "RouterOspf-DistributeRouteMapIn"); ok {
			if err = d.Set("distribute_route_map_in", vv); err != nil {
				return fmt.Errorf("Error reading distribute_route_map_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distribute_route_map_in: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspfLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-neighbour-changes"], "RouterOspf-LogNeighbourChanges"); ok {
			if err = d.Set("log_neighbour_changes", vv); err != nil {
				return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("lsa_refresh_interval", flattenRouterOspfLsaRefreshInterval(o["lsa-refresh-interval"], d, "lsa_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsa-refresh-interval"], "RouterOspf-LsaRefreshInterval"); ok {
			if err = d.Set("lsa_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading lsa_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsa_refresh_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterOspf-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterOspf-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
			if vv, ok := fortiAPIPatch(o["network"], "RouterOspf-Network"); ok {
				if err = d.Set("network", vv); err != nil {
					return fmt.Errorf("Error reading network: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
				if vv, ok := fortiAPIPatch(o["network"], "RouterOspf-Network"); ok {
					if err = d.Set("network", vv); err != nil {
						return fmt.Errorf("Error reading network: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["ospf-interface"], "RouterOspf-OspfInterface"); ok {
				if err = d.Set("ospf_interface", vv); err != nil {
					return fmt.Errorf("Error reading ospf_interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ospf_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ospf_interface"); ok {
			if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["ospf-interface"], "RouterOspf-OspfInterface"); ok {
					if err = d.Set("ospf_interface", vv); err != nil {
						return fmt.Errorf("Error reading ospf_interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ospf_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-interface"], "RouterOspf-PassiveInterface"); ok {
			if err = d.Set("passive_interface", vv); err != nil {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterOspf-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterOspf-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("restart_mode", flattenRouterOspfRestartMode(o["restart-mode"], d, "restart_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-mode"], "RouterOspf-RestartMode"); ok {
			if err = d.Set("restart_mode", vv); err != nil {
				return fmt.Errorf("Error reading restart_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_on_topology_change", flattenRouterOspfRestartOnTopologyChange(o["restart-on-topology-change"], d, "restart_on_topology_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-on-topology-change"], "RouterOspf-RestartOnTopologyChange"); ok {
			if err = d.Set("restart_on_topology_change", vv); err != nil {
				return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
		}
	}

	if err = d.Set("restart_period", flattenRouterOspfRestartPeriod(o["restart-period"], d, "restart_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-period"], "RouterOspf-RestartPeriod"); ok {
			if err = d.Set("restart_period", vv); err != nil {
				return fmt.Errorf("Error reading restart_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("rfc1583_compatible", flattenRouterOspfRfc1583Compatible(o["rfc1583-compatible"], d, "rfc1583_compatible")); err != nil {
		if vv, ok := fortiAPIPatch(o["rfc1583-compatible"], "RouterOspf-Rfc1583Compatible"); ok {
			if err = d.Set("rfc1583_compatible", vv); err != nil {
				return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspfRouterId(o["router-id"], d, "router_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-id"], "RouterOspf-RouterId"); ok {
			if err = d.Set("router_id", vv); err != nil {
				return fmt.Errorf("Error reading router_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspfSpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if vv, ok := fortiAPIPatch(o["spf-timers"], "RouterOspf-SpfTimers"); ok {
			if err = d.Set("spf_timers", vv); err != nil {
				return fmt.Errorf("Error reading spf_timers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["summary-address"], "RouterOspf-SummaryAddress"); ok {
				if err = d.Set("summary_address", vv); err != nil {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["summary-address"], "RouterOspf-SummaryAddress"); ok {
					if err = d.Set("summary_address", vv); err != nil {
						return fmt.Errorf("Error reading summary_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterOspfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAbrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandRouterOspfAreaAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandRouterOspfAreaComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-cost"], _ = expandRouterOspfAreaDefaultCost(d, i["default_cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfAreaFilterList(d, i["filter_list"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["filter-list"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfAreaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate"], _ = expandRouterOspfAreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-redistribution"], _ = expandRouterOspfAreaNssaRedistribution(d, i["nssa_redistribution"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-translator-role"], _ = expandRouterOspfAreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfAreaRange(d, i["range"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["range"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut"], _ = expandRouterOspfAreaShortcut(d, i["shortcut"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stub-type"], _ = expandRouterOspfAreaStubType(d, i["stub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandRouterOspfAreaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfAreaVirtualLink(d, i["virtual_link"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["virtual-link"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaDefaultCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandRouterOspfAreaFilterListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfAreaFilterListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["list"], _ = expandRouterOspfAreaFilterListList(d, i["list"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaFilterListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertise"], _ = expandRouterOspfAreaRangeAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfAreaRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterOspfAreaRangePrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["substitute"], _ = expandRouterOspfAreaRangeSubstitute(d, i["substitute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["substitute-status"], _ = expandRouterOspfAreaRangeSubstituteStatus(d, i["substitute_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstitute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstituteStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaShortcut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaStubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandRouterOspfAreaVirtualLinkAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication-key"], _ = expandRouterOspfAreaVirtualLinkAuthenticationKey(d, i["authentication_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspfAreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspfAreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keychain"], _ = expandRouterOspfAreaVirtualLinkKeychain(d, i["keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5-keychain"], _ = expandRouterOspfAreaVirtualLinkMd5Keychain(d, i["md5_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfAreaVirtualLinkMd5Keys(d, i["md5_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["md5-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspfAreaVirtualLinkName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer"], _ = expandRouterOspfAreaVirtualLinkPeer(d, i["peer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspfAreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspfAreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthenticationKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkKeychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaVirtualLinkMd5KeysId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterOspfAreaVirtualLinkMd5KeysKeyString(d, i["key_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowMaxLsas(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowTimeToRecover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfDefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceInterArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceIntraArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-list"], _ = expandRouterOspfDistributeListAccessList(d, i["access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfDistributeListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandRouterOspfDistributeListProtocol(d, i["protocol"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfDistributeListAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfDistributeListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfDistributeRouteMapIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfLsaRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandRouterOspfNeighborCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfNeighborId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandRouterOspfNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poll-interval"], _ = expandRouterOspfNeighborPollInterval(d, i["poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandRouterOspfNeighborPriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfNeighborCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["area"], _ = expandRouterOspfNetworkArea(d, i["area"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandRouterOspfNetworkComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterOspfNetworkPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfNetworkArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfOspfInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandRouterOspfOspfInterfaceAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication-key"], _ = expandRouterOspfOspfInterfaceAuthenticationKey(d, i["authentication_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd"], _ = expandRouterOspfOspfInterfaceBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandRouterOspfOspfInterfaceComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandRouterOspfOspfInterfaceCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["database-filter-out"], _ = expandRouterOspfOspfInterfaceDatabaseFilterOut(d, i["database_filter_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspfOspfInterfaceDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspfOspfInterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-multiplier"], _ = expandRouterOspfOspfInterfaceHelloMultiplier(d, i["hello_multiplier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterOspfOspfInterfaceInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandRouterOspfOspfInterfaceIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keychain"], _ = expandRouterOspfOspfInterfaceKeychain(d, i["keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "linkdown_fast_failover"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["linkdown-fast-failover"], _ = expandRouterOspfOspfInterfaceLinkdownFastFailover(d, i["linkdown_fast_failover"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5-keychain"], _ = expandRouterOspfOspfInterfaceMd5Keychain(d, i["md5_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfOspfInterfaceMd5Keys(d, i["md5_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["md5-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mtu"], _ = expandRouterOspfOspfInterfaceMtu(d, i["mtu"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mtu-ignore"], _ = expandRouterOspfOspfInterfaceMtuIgnore(d, i["mtu_ignore"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspfOspfInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["network-type"], _ = expandRouterOspfOspfInterfaceNetworkType(d, i["network_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix-length"], _ = expandRouterOspfOspfInterfacePrefixLength(d, i["prefix_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandRouterOspfOspfInterfacePriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["resync-timeout"], _ = expandRouterOspfOspfInterfaceResyncTimeout(d, i["resync_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspfOspfInterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterOspfOspfInterfaceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspfOspfInterfaceTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDatabaseFilterOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceKeychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceLinkdownFastFailover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Keychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfOspfInterfaceMd5KeysId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterOspfOspfInterfaceMd5KeysKeyString(d, i["key_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceMd5KeysId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceResyncTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfPassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric"], _ = expandRouterOspfRedistributeMetric(d, i["metric"], pre_append)
	}
	pre_append = pre + ".0." + "metric_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric-type"], _ = expandRouterOspfRedistributeMetricType(d, i["metric_type"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandRouterOspfRedistributeName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "routemap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["routemap"], _ = expandRouterOspfRedistributeRoutemap(d, i["routemap"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandRouterOspfRedistributeStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandRouterOspfRedistributeTag(d, i["tag"], pre_append)
	}

	return result, nil
}

func expandRouterOspfRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartOnTopologyChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRfc1583Compatible(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSpfTimers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandRouterOspfSummaryAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advertise"], _ = expandRouterOspfSummaryAddressAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfSummaryAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterOspfSummaryAddressPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tag"], _ = expandRouterOspfSummaryAddressTag(d, i["tag"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfSummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfSummaryAddressTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("abr_type"); ok || d.HasChange("abr_type") {
		t, err := expandRouterOspfAbrType(d, v, "abr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["abr-type"] = t
		}
	}

	if v, ok := d.GetOk("area"); ok || d.HasChange("area") {
		t, err := expandRouterOspfArea(d, v, "area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area"] = t
		}
	}

	if v, ok := d.GetOk("auto_cost_ref_bandwidth"); ok || d.HasChange("auto_cost_ref_bandwidth") {
		t, err := expandRouterOspfAutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-cost-ref-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterOspfBfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow"); ok || d.HasChange("database_overflow") {
		t, err := expandRouterOspfDatabaseOverflow(d, v, "database_overflow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow_max_lsas"); ok || d.HasChange("database_overflow_max_lsas") {
		t, err := expandRouterOspfDatabaseOverflowMaxLsas(d, v, "database_overflow_max_lsas")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow-max-lsas"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow_time_to_recover"); ok || d.HasChange("database_overflow_time_to_recover") {
		t, err := expandRouterOspfDatabaseOverflowTimeToRecover(d, v, "database_overflow_time_to_recover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow-time-to-recover"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok || d.HasChange("default_information_metric") {
		t, err := expandRouterOspfDefaultInformationMetric(d, v, "default_information_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok || d.HasChange("default_information_metric_type") {
		t, err := expandRouterOspfDefaultInformationMetricType(d, v, "default_information_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok || d.HasChange("default_information_originate") {
		t, err := expandRouterOspfDefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_information_route_map"); ok || d.HasChange("default_information_route_map") {
		t, err := expandRouterOspfDefaultInformationRouteMap(d, v, "default_information_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-route-map"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok || d.HasChange("default_metric") {
		t, err := expandRouterOspfDefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandRouterOspfDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("distance_external"); ok || d.HasChange("distance_external") {
		t, err := expandRouterOspfDistanceExternal(d, v, "distance_external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-external"] = t
		}
	}

	if v, ok := d.GetOk("distance_inter_area"); ok || d.HasChange("distance_inter_area") {
		t, err := expandRouterOspfDistanceInterArea(d, v, "distance_inter_area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-inter-area"] = t
		}
	}

	if v, ok := d.GetOk("distance_intra_area"); ok || d.HasChange("distance_intra_area") {
		t, err := expandRouterOspfDistanceIntraArea(d, v, "distance_intra_area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-intra-area"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok || d.HasChange("distribute_list") {
		t, err := expandRouterOspfDistributeList(d, v, "distribute_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok || d.HasChange("distribute_list_in") {
		t, err := expandRouterOspfDistributeListIn(d, v, "distribute_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_route_map_in"); ok || d.HasChange("distribute_route_map_in") {
		t, err := expandRouterOspfDistributeRouteMapIn(d, v, "distribute_route_map_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-route-map-in"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok || d.HasChange("log_neighbour_changes") {
		t, err := expandRouterOspfLogNeighbourChanges(d, v, "log_neighbour_changes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("lsa_refresh_interval"); ok || d.HasChange("lsa_refresh_interval") {
		t, err := expandRouterOspfLsaRefreshInterval(d, v, "lsa_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsa-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterOspfNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		t, err := expandRouterOspfNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("ospf_interface"); ok || d.HasChange("ospf_interface") {
		t, err := expandRouterOspfOspfInterface(d, v, "ospf_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf-interface"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		t, err := expandRouterOspfPassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterOspfRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("restart_mode"); ok || d.HasChange("restart_mode") {
		t, err := expandRouterOspfRestartMode(d, v, "restart_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-mode"] = t
		}
	}

	if v, ok := d.GetOk("restart_on_topology_change"); ok || d.HasChange("restart_on_topology_change") {
		t, err := expandRouterOspfRestartOnTopologyChange(d, v, "restart_on_topology_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-on-topology-change"] = t
		}
	}

	if v, ok := d.GetOk("restart_period"); ok || d.HasChange("restart_period") {
		t, err := expandRouterOspfRestartPeriod(d, v, "restart_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-period"] = t
		}
	}

	if v, ok := d.GetOk("rfc1583_compatible"); ok || d.HasChange("rfc1583_compatible") {
		t, err := expandRouterOspfRfc1583Compatible(d, v, "rfc1583_compatible")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc1583-compatible"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok || d.HasChange("router_id") {
		t, err := expandRouterOspfRouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("spf_timers"); ok || d.HasChange("spf_timers") {
		t, err := expandRouterOspfSpfTimers(d, v, "spf_timers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-timers"] = t
		}
	}

	if v, ok := d.GetOk("summary_address"); ok || d.HasChange("summary_address") {
		t, err := expandRouterOspfSummaryAddress(d, v, "summary_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address"] = t
		}
	}

	return &obj, nil
}
