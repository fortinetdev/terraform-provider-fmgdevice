// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 OSPF.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6Update,
		Read:   resourceRouterOspf6Read,
		Update: resourceRouterOspf6Update,
		Delete: resourceRouterOspf6Delete,

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
						"default_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"enc_key": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"key_rollover_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
									"prefix6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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
									"ipsec_auth_alg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipsec_enc_alg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipsec_keys": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_key": &schema.Schema{
													Type:      schema.TypeSet,
													Elem:      &schema.Schema{Type: schema.TypeString},
													Optional:  true,
													Sensitive: true,
													Computed:  true,
												},
												"enc_key": &schema.Schema{
													Type:      schema.TypeSet,
													Elem:      &schema.Schema{Type: schema.TypeString},
													Optional:  true,
													Sensitive: true,
													Computed:  true,
												},
												"spi": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"key_rollover_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dead_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"enc_key": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"key_rollover_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
						"neighbor": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
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
						},
						"prefix6": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterOspf6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterOspf6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterOspf6")

	return resourceRouterOspf6Read(d, m)
}

func resourceRouterOspf6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspf6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Area(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			v := flattenRouterOspf6AreaDefaultCost(i["default-cost"], d, pre_append)
			tmp["default_cost"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-DefaultCost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspf6AreaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			v := flattenRouterOspf6AreaIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
			tmp["ipsec_auth_alg"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-IpsecAuthAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			v := flattenRouterOspf6AreaIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
			tmp["ipsec_enc_alg"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-IpsecEncAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			v := flattenRouterOspf6AreaIpsecKeys(i["ipsec-keys"], d, pre_append)
			tmp["ipsec_keys"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-IpsecKeys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			v := flattenRouterOspf6AreaKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
			tmp["key_rollover_interval"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-KeyRolloverInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			v := flattenRouterOspf6AreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
			tmp["nssa_default_information_originate"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-NssaDefaultInformationOriginate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			v := flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
			tmp["nssa_default_information_originate_metric"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-NssaDefaultInformationOriginateMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			v := flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
			tmp["nssa_default_information_originate_metric_type"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-NssaDefaultInformationOriginateMetricType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			v := flattenRouterOspf6AreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
			tmp["nssa_redistribution"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-NssaRedistribution")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			v := flattenRouterOspf6AreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
			tmp["nssa_translator_role"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-NssaTranslatorRole")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			v := flattenRouterOspf6AreaRange(i["range"], d, pre_append)
			tmp["range"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-Range")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			v := flattenRouterOspf6AreaStubType(i["stub-type"], d, pre_append)
			tmp["stub_type"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-StubType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenRouterOspf6AreaType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			v := flattenRouterOspf6AreaVirtualLink(i["virtual-link"], d, pre_append)
			tmp["virtual_link"] = fortiAPISubPartPatch(v, "RouterOspf6-Area-VirtualLink")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			v := flattenRouterOspf6AreaIpsecKeysSpi(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6Area-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaRangeAdvertise(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspf6AreaRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterOspf6AreaRangePrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
			tmp["ipsec_auth_alg"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecAuthAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
			tmp["ipsec_enc_alg"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecEncAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecKeys(i["ipsec-keys"], d, pre_append)
			tmp["ipsec_keys"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecKeys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
			tmp["key_rollover_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-KeyRolloverInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkPeer(i["peer"], d, pre_append)
			tmp["peer"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Peer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6AreaVirtualLink-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspf6DefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6LogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := i["area-id"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceAreaId(i["area-id"], d, pre_append)
			tmp["area_id"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-AreaId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceBfd(i["bfd"], d, pre_append)
			tmp["bfd"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Bfd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceDeadInterval(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
			tmp["ipsec_auth_alg"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-IpsecAuthAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
			tmp["ipsec_enc_alg"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-IpsecEncAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceIpsecKeys(i["ipsec-keys"], d, pre_append)
			tmp["ipsec_keys"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-IpsecKeys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
			tmp["key_rollover_interval"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-KeyRolloverInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceMtu(i["mtu"], d, pre_append)
			tmp["mtu"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Mtu")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
			tmp["mtu_ignore"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-MtuIgnore")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := i["neighbor"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighbor(i["neighbor"], d, pre_append)
			tmp["neighbor"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Neighbor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNetworkType(i["network-type"], d, pre_append)
			tmp["network_type"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-NetworkType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenRouterOspf6Ospf6InterfacePriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspf6-Ospf6Interface-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6Ospf6InterfaceNeighborCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(i["poll-interval"], d, pre_append)
			tmp["poll_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-PollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6PassiveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspf6Redistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := i["metric"]; ok {
		result["metric"] = flattenRouterOspf6RedistributeMetric(i["metric"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric_type"
	if _, ok := i["metric-type"]; ok {
		result["metric_type"] = flattenRouterOspf6RedistributeMetricType(i["metric-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenRouterOspf6RedistributeName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "routemap"
	if _, ok := i["routemap"]; ok {
		result["routemap"] = flattenRouterOspf6RedistributeRoutemap(i["routemap"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenRouterOspf6RedistributeStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterOspf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RestartOnTopologyChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RestartPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenRouterOspf6SummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6SummaryAddressAdvertise(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspf6-SummaryAddress-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspf6SummaryAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf6-SummaryAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterOspf6SummaryAddressPrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterOspf6-SummaryAddress-Prefix6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			v := flattenRouterOspf6SummaryAddressTag(i["tag"], d, pre_append)
			tmp["tag"] = fortiAPISubPartPatch(v, "RouterOspf6-SummaryAddress-Tag")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6SummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("abr_type", flattenRouterOspf6AbrType(o["abr-type"], d, "abr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["abr-type"], "RouterOspf6-AbrType"); ok {
			if err = d.Set("abr_type", vv); err != nil {
				return fmt.Errorf("Error reading abr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area")); err != nil {
			if vv, ok := fortiAPIPatch(o["area"], "RouterOspf6-Area"); ok {
				if err = d.Set("area", vv); err != nil {
					return fmt.Errorf("Error reading area: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading area: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("area"); ok {
			if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area")); err != nil {
				if vv, ok := fortiAPIPatch(o["area"], "RouterOspf6-Area"); ok {
					if err = d.Set("area", vv); err != nil {
						return fmt.Errorf("Error reading area: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading area: %v", err)
				}
			}
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspf6AutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-cost-ref-bandwidth"], "RouterOspf6-AutoCostRefBandwidth"); ok {
			if err = d.Set("auto_cost_ref_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspf6Bfd(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterOspf6-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspf6DefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-metric"], "RouterOspf6-DefaultInformationMetric"); ok {
			if err = d.Set("default_information_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_information_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspf6DefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-metric-type"], "RouterOspf6-DefaultInformationMetricType"); ok {
			if err = d.Set("default_information_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading default_information_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspf6DefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-originate"], "RouterOspf6-DefaultInformationOriginate"); ok {
			if err = d.Set("default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspf6DefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-route-map"], "RouterOspf6-DefaultInformationRouteMap"); ok {
			if err = d.Set("default_information_route_map", vv); err != nil {
				return fmt.Errorf("Error reading default_information_route_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspf6DefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-metric"], "RouterOspf6-DefaultMetric"); ok {
			if err = d.Set("default_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspf6LogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-neighbour-changes"], "RouterOspf6-LogNeighbourChanges"); ok {
			if err = d.Set("log_neighbour_changes", vv); err != nil {
				return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["ospf6-interface"], "RouterOspf6-Ospf6Interface"); ok {
				if err = d.Set("ospf6_interface", vv); err != nil {
					return fmt.Errorf("Error reading ospf6_interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ospf6_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ospf6_interface"); ok {
			if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["ospf6-interface"], "RouterOspf6-Ospf6Interface"); ok {
					if err = d.Set("ospf6_interface", vv); err != nil {
						return fmt.Errorf("Error reading ospf6_interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ospf6_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-interface"], "RouterOspf6-PassiveInterface"); ok {
			if err = d.Set("passive_interface", vv); err != nil {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterOspf6-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterOspf6-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("restart_mode", flattenRouterOspf6RestartMode(o["restart-mode"], d, "restart_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-mode"], "RouterOspf6-RestartMode"); ok {
			if err = d.Set("restart_mode", vv); err != nil {
				return fmt.Errorf("Error reading restart_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_on_topology_change", flattenRouterOspf6RestartOnTopologyChange(o["restart-on-topology-change"], d, "restart_on_topology_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-on-topology-change"], "RouterOspf6-RestartOnTopologyChange"); ok {
			if err = d.Set("restart_on_topology_change", vv); err != nil {
				return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
		}
	}

	if err = d.Set("restart_period", flattenRouterOspf6RestartPeriod(o["restart-period"], d, "restart_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["restart-period"], "RouterOspf6-RestartPeriod"); ok {
			if err = d.Set("restart_period", vv); err != nil {
				return fmt.Errorf("Error reading restart_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspf6RouterId(o["router-id"], d, "router_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-id"], "RouterOspf6-RouterId"); ok {
			if err = d.Set("router_id", vv); err != nil {
				return fmt.Errorf("Error reading router_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspf6SpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if vv, ok := fortiAPIPatch(o["spf-timers"], "RouterOspf6-SpfTimers"); ok {
			if err = d.Set("spf_timers", vv); err != nil {
				return fmt.Errorf("Error reading spf_timers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["summary-address"], "RouterOspf6-SummaryAddress"); ok {
				if err = d.Set("summary_address", vv); err != nil {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["summary-address"], "RouterOspf6-SummaryAddress"); ok {
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

func flattenRouterOspf6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6AbrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Area(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["authentication"], _ = expandRouterOspf6AreaAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-cost"], _ = expandRouterOspf6AreaDefaultCost(d, i["default_cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspf6AreaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6AreaIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6AreaIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6AreaIpsecKeys(d, i["ipsec_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ipsec-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-rollover-interval"], _ = expandRouterOspf6AreaKeyRolloverInterval(d, i["key_rollover_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-redistribution"], _ = expandRouterOspf6AreaNssaRedistribution(d, i["nssa_redistribution"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nssa-translator-role"], _ = expandRouterOspf6AreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6AreaRange(d, i["range"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["range"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stub-type"], _ = expandRouterOspf6AreaStubType(d, i["stub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandRouterOspf6AreaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6AreaVirtualLink(d, i["virtual_link"], pre_append)
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

func expandRouterOspf6AreaAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaDefaultCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecKeys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-key"], _ = expandRouterOspf6AreaIpsecKeysAuthKey(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6AreaIpsecKeysEncKey(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6AreaIpsecKeysSpi(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise"], _ = expandRouterOspf6AreaRangeAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspf6AreaRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterOspf6AreaRangePrefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaStubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["authentication"], _ = expandRouterOspf6AreaVirtualLinkAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspf6AreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspf6AreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6AreaVirtualLinkIpsecKeys(d, i["ipsec_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ipsec-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-rollover-interval"], _ = expandRouterOspf6AreaVirtualLinkKeyRolloverInterval(d, i["key_rollover_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspf6AreaVirtualLinkName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer"], _ = expandRouterOspf6AreaVirtualLinkPeer(d, i["peer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspf6AreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspf6AreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysSpi(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Bfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6DefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6LogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["area-id"], _ = expandRouterOspf6Ospf6InterfaceAreaId(d, i["area_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandRouterOspf6Ospf6InterfaceAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bfd"], _ = expandRouterOspf6Ospf6InterfaceBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspf6Ospf6InterfaceDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspf6Ospf6InterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterOspf6Ospf6InterfaceInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6Ospf6InterfaceIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6Ospf6InterfaceIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6Ospf6InterfaceIpsecKeys(d, i["ipsec_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ipsec-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-rollover-interval"], _ = expandRouterOspf6Ospf6InterfaceKeyRolloverInterval(d, i["key_rollover_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mtu"], _ = expandRouterOspf6Ospf6InterfaceMtu(d, i["mtu"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mtu-ignore"], _ = expandRouterOspf6Ospf6InterfaceMtuIgnore(d, i["mtu_ignore"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspf6Ospf6InterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6Ospf6InterfaceNeighbor(d, i["neighbor"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["neighbor"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["network-type"], _ = expandRouterOspf6Ospf6InterfaceNetworkType(d, i["network_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfacePriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspf6Ospf6InterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterOspf6Ospf6InterfaceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspf6Ospf6InterfaceTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceAreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysSpi(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceNeighborCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandRouterOspf6Ospf6InterfaceNeighborIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poll-interval"], _ = expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d, i["poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfaceNeighborPriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6PassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Redistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric"], _ = expandRouterOspf6RedistributeMetric(d, i["metric"], pre_append)
	}
	pre_append = pre + ".0." + "metric_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric-type"], _ = expandRouterOspf6RedistributeMetricType(d, i["metric_type"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandRouterOspf6RedistributeName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "routemap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["routemap"], _ = expandRouterOspf6RedistributeRoutemap(d, i["routemap"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandRouterOspf6RedistributeStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandRouterOspf6RedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6RedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartOnTopologyChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SpfTimers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6SummaryAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise"], _ = expandRouterOspf6SummaryAddressAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspf6SummaryAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterOspf6SummaryAddressPrefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tag"], _ = expandRouterOspf6SummaryAddressTag(d, i["tag"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6SummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("abr_type"); ok || d.HasChange("abr_type") {
		t, err := expandRouterOspf6AbrType(d, v, "abr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["abr-type"] = t
		}
	}

	if v, ok := d.GetOk("area"); ok || d.HasChange("area") {
		t, err := expandRouterOspf6Area(d, v, "area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area"] = t
		}
	}

	if v, ok := d.GetOk("auto_cost_ref_bandwidth"); ok || d.HasChange("auto_cost_ref_bandwidth") {
		t, err := expandRouterOspf6AutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-cost-ref-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterOspf6Bfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok || d.HasChange("default_information_metric") {
		t, err := expandRouterOspf6DefaultInformationMetric(d, v, "default_information_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok || d.HasChange("default_information_metric_type") {
		t, err := expandRouterOspf6DefaultInformationMetricType(d, v, "default_information_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok || d.HasChange("default_information_originate") {
		t, err := expandRouterOspf6DefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_information_route_map"); ok || d.HasChange("default_information_route_map") {
		t, err := expandRouterOspf6DefaultInformationRouteMap(d, v, "default_information_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-route-map"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok || d.HasChange("default_metric") {
		t, err := expandRouterOspf6DefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok || d.HasChange("log_neighbour_changes") {
		t, err := expandRouterOspf6LogNeighbourChanges(d, v, "log_neighbour_changes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_interface"); ok || d.HasChange("ospf6_interface") {
		t, err := expandRouterOspf6Ospf6Interface(d, v, "ospf6_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6-interface"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		t, err := expandRouterOspf6PassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterOspf6Redistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("restart_mode"); ok || d.HasChange("restart_mode") {
		t, err := expandRouterOspf6RestartMode(d, v, "restart_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-mode"] = t
		}
	}

	if v, ok := d.GetOk("restart_on_topology_change"); ok || d.HasChange("restart_on_topology_change") {
		t, err := expandRouterOspf6RestartOnTopologyChange(d, v, "restart_on_topology_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-on-topology-change"] = t
		}
	}

	if v, ok := d.GetOk("restart_period"); ok || d.HasChange("restart_period") {
		t, err := expandRouterOspf6RestartPeriod(d, v, "restart_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-period"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok || d.HasChange("router_id") {
		t, err := expandRouterOspf6RouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("spf_timers"); ok || d.HasChange("spf_timers") {
		t, err := expandRouterOspf6SpfTimers(d, v, "spf_timers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-timers"] = t
		}
	}

	if v, ok := d.GetOk("summary_address"); ok || d.HasChange("summary_address") {
		t, err := expandRouterOspf6SummaryAddress(d, v, "summary_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address"] = t
		}
	}

	return &obj, nil
}
