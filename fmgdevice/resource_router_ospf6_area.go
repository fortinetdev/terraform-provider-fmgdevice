// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF6 area configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6Area() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6AreaCreate,
		Read:   resourceRouterOspf6AreaRead,
		Update: resourceRouterOspf6AreaUpdate,
		Delete: resourceRouterOspf6AreaDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterOspf6AreaCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspf6Area(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Area resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6Area(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Area resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceRouterOspf6AreaRead(d, m)
}

func resourceRouterOspf6AreaUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspf6Area(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Area resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6Area(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Area resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceRouterOspf6AreaRead(d, m)
}

func resourceRouterOspf6AreaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspf6Area(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6Area resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6AreaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf6Area(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Area resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6Area(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Area resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AreaAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaDefaultCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecAuthAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecEncAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecKeys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaIpsecKeysSpi2edl(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6Area-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaIpsecKeysSpi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaKeyRolloverInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaRedistribution2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaTranslatorRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaRangeAdvertise2edl(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspf6AreaRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterOspf6AreaRangePrefix62edl(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterOspf6Area-Range-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaRangeAdvertise2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaStubType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLink2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaVirtualLinkAuthentication2edl(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkDeadInterval2edl(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkHelloInterval2edl(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg2edl(i["ipsec-auth-alg"], d, pre_append)
			tmp["ipsec_auth_alg"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecAuthAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecEncAlg2edl(i["ipsec-enc-alg"], d, pre_append)
			tmp["ipsec_enc_alg"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecEncAlg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkIpsecKeys2edl(i["ipsec-keys"], d, pre_append)
			tmp["ipsec_keys"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-IpsecKeys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval2edl(i["key-rollover-interval"], d, pre_append)
			tmp["key_rollover_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-KeyRolloverInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkPeer2edl(i["peer"], d, pre_append)
			tmp["peer"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-Peer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkRetransmitInterval2edl(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspf6AreaVirtualLinkTransmitDelay2edl(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspf6Area-VirtualLink-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkDeadInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecEncAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi2edl(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6AreaVirtualLink-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkPeer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkRetransmitInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkTransmitDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6Area(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenRouterOspf6AreaAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspf6Area-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("default_cost", flattenRouterOspf6AreaDefaultCost2edl(o["default-cost"], d, "default_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-cost"], "RouterOspf6Area-DefaultCost"); ok {
			if err = d.Set("default_cost", vv); err != nil {
				return fmt.Errorf("Error reading default_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_cost: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterOspf6AreaId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspf6Area-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipsec_auth_alg", flattenRouterOspf6AreaIpsecAuthAlg2edl(o["ipsec-auth-alg"], d, "ipsec_auth_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-auth-alg"], "RouterOspf6Area-IpsecAuthAlg"); ok {
			if err = d.Set("ipsec_auth_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_alg", flattenRouterOspf6AreaIpsecEncAlg2edl(o["ipsec-enc-alg"], d, "ipsec_enc_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-enc-alg"], "RouterOspf6Area-IpsecEncAlg"); ok {
			if err = d.Set("ipsec_enc_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipsec_keys", flattenRouterOspf6AreaIpsecKeys2edl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6Area-IpsecKeys"); ok {
				if err = d.Set("ipsec_keys", vv); err != nil {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipsec_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipsec_keys"); ok {
			if err = d.Set("ipsec_keys", flattenRouterOspf6AreaIpsecKeys2edl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6Area-IpsecKeys"); ok {
					if err = d.Set("ipsec_keys", vv); err != nil {
						return fmt.Errorf("Error reading ipsec_keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("key_rollover_interval", flattenRouterOspf6AreaKeyRolloverInterval2edl(o["key-rollover-interval"], d, "key_rollover_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-rollover-interval"], "RouterOspf6Area-KeyRolloverInterval"); ok {
			if err = d.Set("key_rollover_interval", vv); err != nil {
				return fmt.Errorf("Error reading key_rollover_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_rollover_interval: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate", flattenRouterOspf6AreaNssaDefaultInformationOriginate2edl(o["nssa-default-information-originate"], d, "nssa_default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate"], "RouterOspf6Area-NssaDefaultInformationOriginate"); ok {
			if err = d.Set("nssa_default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric", flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric2edl(o["nssa-default-information-originate-metric"], d, "nssa_default_information_originate_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate-metric"], "RouterOspf6Area-NssaDefaultInformationOriginateMetric"); ok {
			if err = d.Set("nssa_default_information_originate_metric", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate_metric: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric_type", flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType2edl(o["nssa-default-information-originate-metric-type"], d, "nssa_default_information_originate_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate-metric-type"], "RouterOspf6Area-NssaDefaultInformationOriginateMetricType"); ok {
			if err = d.Set("nssa_default_information_originate_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate_metric_type: %v", err)
		}
	}

	if err = d.Set("nssa_redistribution", flattenRouterOspf6AreaNssaRedistribution2edl(o["nssa-redistribution"], d, "nssa_redistribution")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-redistribution"], "RouterOspf6Area-NssaRedistribution"); ok {
			if err = d.Set("nssa_redistribution", vv); err != nil {
				return fmt.Errorf("Error reading nssa_redistribution: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_redistribution: %v", err)
		}
	}

	if err = d.Set("nssa_translator_role", flattenRouterOspf6AreaNssaTranslatorRole2edl(o["nssa-translator-role"], d, "nssa_translator_role")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-translator-role"], "RouterOspf6Area-NssaTranslatorRole"); ok {
			if err = d.Set("nssa_translator_role", vv); err != nil {
				return fmt.Errorf("Error reading nssa_translator_role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_translator_role: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("range", flattenRouterOspf6AreaRange2edl(o["range"], d, "range")); err != nil {
			if vv, ok := fortiAPIPatch(o["range"], "RouterOspf6Area-Range"); ok {
				if err = d.Set("range", vv); err != nil {
					return fmt.Errorf("Error reading range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("range"); ok {
			if err = d.Set("range", flattenRouterOspf6AreaRange2edl(o["range"], d, "range")); err != nil {
				if vv, ok := fortiAPIPatch(o["range"], "RouterOspf6Area-Range"); ok {
					if err = d.Set("range", vv); err != nil {
						return fmt.Errorf("Error reading range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading range: %v", err)
				}
			}
		}
	}

	if err = d.Set("stub_type", flattenRouterOspf6AreaStubType2edl(o["stub-type"], d, "stub_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["stub-type"], "RouterOspf6Area-StubType"); ok {
			if err = d.Set("stub_type", vv); err != nil {
				return fmt.Errorf("Error reading stub_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stub_type: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterOspf6AreaType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "RouterOspf6Area-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("virtual_link", flattenRouterOspf6AreaVirtualLink2edl(o["virtual-link"], d, "virtual_link")); err != nil {
			if vv, ok := fortiAPIPatch(o["virtual-link"], "RouterOspf6Area-VirtualLink"); ok {
				if err = d.Set("virtual_link", vv); err != nil {
					return fmt.Errorf("Error reading virtual_link: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading virtual_link: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("virtual_link"); ok {
			if err = d.Set("virtual_link", flattenRouterOspf6AreaVirtualLink2edl(o["virtual-link"], d, "virtual_link")); err != nil {
				if vv, ok := fortiAPIPatch(o["virtual-link"], "RouterOspf6Area-VirtualLink"); ok {
					if err = d.Set("virtual_link", vv); err != nil {
						return fmt.Errorf("Error reading virtual_link: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading virtual_link: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterOspf6AreaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6AreaAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaDefaultCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecAuthAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecEncAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecKeys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["auth-key"], _ = expandRouterOspf6AreaIpsecKeysAuthKey2edl(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6AreaIpsecKeysEncKey2edl(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6AreaIpsecKeysSpi2edl(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaIpsecKeysAuthKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaIpsecKeysEncKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaIpsecKeysSpi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaKeyRolloverInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaRedistribution2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaTranslatorRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise"], _ = expandRouterOspf6AreaRangeAdvertise2edl(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspf6AreaRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterOspf6AreaRangePrefix62edl(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaRangeAdvertise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaStubType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["authentication"], _ = expandRouterOspf6AreaVirtualLinkAuthentication2edl(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspf6AreaVirtualLinkDeadInterval2edl(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspf6AreaVirtualLinkHelloInterval2edl(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecAuthAlg2edl(d, i["ipsec_auth_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecEncAlg2edl(d, i["ipsec_enc_alg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspf6AreaVirtualLinkIpsecKeys2edl(d, i["ipsec_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["ipsec-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-rollover-interval"], _ = expandRouterOspf6AreaVirtualLinkKeyRolloverInterval2edl(d, i["key_rollover_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspf6AreaVirtualLinkName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer"], _ = expandRouterOspf6AreaVirtualLinkPeer2edl(d, i["peer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspf6AreaVirtualLinkRetransmitInterval2edl(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspf6AreaVirtualLinkTransmitDelay2edl(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkDeadInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecAuthAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecEncAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["auth-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey2edl(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey2edl(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysSpi2edl(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysSpi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkKeyRolloverInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkPeer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkRetransmitInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkTransmitDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6Area(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspf6AreaAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("default_cost"); ok || d.HasChange("default_cost") {
		t, err := expandRouterOspf6AreaDefaultCost2edl(d, v, "default_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cost"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspf6AreaId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_auth_alg"); ok || d.HasChange("ipsec_auth_alg") {
		t, err := expandRouterOspf6AreaIpsecAuthAlg2edl(d, v, "ipsec_auth_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-auth-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_enc_alg"); ok || d.HasChange("ipsec_enc_alg") {
		t, err := expandRouterOspf6AreaIpsecEncAlg2edl(d, v, "ipsec_enc_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-enc-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_keys"); ok || d.HasChange("ipsec_keys") {
		t, err := expandRouterOspf6AreaIpsecKeys2edl(d, v, "ipsec_keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-keys"] = t
		}
	}

	if v, ok := d.GetOk("key_rollover_interval"); ok || d.HasChange("key_rollover_interval") {
		t, err := expandRouterOspf6AreaKeyRolloverInterval2edl(d, v, "key_rollover_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-rollover-interval"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate"); ok || d.HasChange("nssa_default_information_originate") {
		t, err := expandRouterOspf6AreaNssaDefaultInformationOriginate2edl(d, v, "nssa_default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric"); ok || d.HasChange("nssa_default_information_originate_metric") {
		t, err := expandRouterOspf6AreaNssaDefaultInformationOriginateMetric2edl(d, v, "nssa_default_information_originate_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric_type"); ok || d.HasChange("nssa_default_information_originate_metric_type") {
		t, err := expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType2edl(d, v, "nssa_default_information_originate_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("nssa_redistribution"); ok || d.HasChange("nssa_redistribution") {
		t, err := expandRouterOspf6AreaNssaRedistribution2edl(d, v, "nssa_redistribution")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-redistribution"] = t
		}
	}

	if v, ok := d.GetOk("nssa_translator_role"); ok || d.HasChange("nssa_translator_role") {
		t, err := expandRouterOspf6AreaNssaTranslatorRole2edl(d, v, "nssa_translator_role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-translator-role"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandRouterOspf6AreaRange2edl(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("stub_type"); ok || d.HasChange("stub_type") {
		t, err := expandRouterOspf6AreaStubType2edl(d, v, "stub_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stub-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandRouterOspf6AreaType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("virtual_link"); ok || d.HasChange("virtual_link") {
		t, err := expandRouterOspf6AreaVirtualLink2edl(d, v, "virtual_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-link"] = t
		}
	}

	return &obj, nil
}
