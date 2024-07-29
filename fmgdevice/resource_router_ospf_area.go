// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF area configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfAreaCreate,
		Read:   resourceRouterOspfAreaRead,
		Update: resourceRouterOspfAreaUpdate,
		Delete: resourceRouterOspfAreaDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterOspfAreaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterOspfArea(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfArea resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfArea(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterOspfArea resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceRouterOspfAreaRead(d, m)
}

func resourceRouterOspfAreaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterOspfArea(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfArea resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfArea(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfArea resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceRouterOspfAreaRead(d, m)
}

func resourceRouterOspfAreaDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfArea(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfArea resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfAreaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfArea(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfArea resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfArea(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfArea resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAreaAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaComments2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaDefaultCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterList2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaFilterListDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfAreaFilterListId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			v := flattenRouterOspfAreaFilterListList2edl(i["list"], d, pre_append)
			tmp["list"] = fortiAPISubPartPatch(v, "RouterOspfArea-FilterList-List")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaFilterListDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetric2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaRedistribution2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaTranslatorRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaRangeAdvertise2edl(i["advertise"], d, pre_append)
			tmp["advertise"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Advertise")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterOspfAreaRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterOspfAreaRangePrefix2edl(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			v := flattenRouterOspfAreaRangeSubstitute2edl(i["substitute"], d, pre_append)
			tmp["substitute"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-Substitute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			v := flattenRouterOspfAreaRangeSubstituteStatus2edl(i["substitute-status"], d, pre_append)
			tmp["substitute_status"] = fortiAPISubPartPatch(v, "RouterOspfArea-Range-SubstituteStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaRangeAdvertise2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangePrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstitute2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaRangeSubstituteStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaShortcut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaStubType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLink2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaVirtualLinkAuthentication2edl(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkDeadInterval2edl(i["dead-interval"], d, pre_append)
			tmp["dead_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-DeadInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkHelloInterval2edl(i["hello-interval"], d, pre_append)
			tmp["hello_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-HelloInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {
			v := flattenRouterOspfAreaVirtualLinkKeychain2edl(i["keychain"], d, pre_append)
			tmp["keychain"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			v := flattenRouterOspfAreaVirtualLinkMd5Keychain2edl(i["md5-keychain"], d, pre_append)
			tmp["md5_keychain"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Md5Keychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			v := flattenRouterOspfAreaVirtualLinkMd5Keys2edl(i["md5-keys"], d, pre_append)
			tmp["md5_keys"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Md5Keys")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterOspfAreaVirtualLinkName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			v := flattenRouterOspfAreaVirtualLinkPeer2edl(i["peer"], d, pre_append)
			tmp["peer"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-Peer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			v := flattenRouterOspfAreaVirtualLinkRetransmitInterval2edl(i["retransmit-interval"], d, pre_append)
			tmp["retransmit_interval"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-RetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			v := flattenRouterOspfAreaVirtualLinkTransmitDelay2edl(i["transmit-delay"], d, pre_append)
			tmp["transmit_delay"] = fortiAPISubPartPatch(v, "RouterOspfArea-VirtualLink-TransmitDelay")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkDeadInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkKeychain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keychain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfAreaVirtualLinkMd5Keys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfAreaVirtualLinkMd5KeysId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfAreaVirtualLink-Md5Keys-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkMd5KeysId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkPeer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkRetransmitInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkTransmitDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfArea(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenRouterOspfAreaAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspfArea-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterOspfAreaComments2edl(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterOspfArea-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("default_cost", flattenRouterOspfAreaDefaultCost2edl(o["default-cost"], d, "default_cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-cost"], "RouterOspfArea-DefaultCost"); ok {
			if err = d.Set("default_cost", vv); err != nil {
				return fmt.Errorf("Error reading default_cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_cost: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter_list", flattenRouterOspfAreaFilterList2edl(o["filter-list"], d, "filter_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["filter-list"], "RouterOspfArea-FilterList"); ok {
				if err = d.Set("filter_list", vv); err != nil {
					return fmt.Errorf("Error reading filter_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter_list"); ok {
			if err = d.Set("filter_list", flattenRouterOspfAreaFilterList2edl(o["filter-list"], d, "filter_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["filter-list"], "RouterOspfArea-FilterList"); ok {
					if err = d.Set("filter_list", vv); err != nil {
						return fmt.Errorf("Error reading filter_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenRouterOspfAreaId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspfArea-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate", flattenRouterOspfAreaNssaDefaultInformationOriginate2edl(o["nssa-default-information-originate"], d, "nssa_default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate"], "RouterOspfArea-NssaDefaultInformationOriginate"); ok {
			if err = d.Set("nssa_default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric", flattenRouterOspfAreaNssaDefaultInformationOriginateMetric2edl(o["nssa-default-information-originate-metric"], d, "nssa_default_information_originate_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate-metric"], "RouterOspfArea-NssaDefaultInformationOriginateMetric"); ok {
			if err = d.Set("nssa_default_information_originate_metric", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate_metric: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric_type", flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType2edl(o["nssa-default-information-originate-metric-type"], d, "nssa_default_information_originate_metric_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-default-information-originate-metric-type"], "RouterOspfArea-NssaDefaultInformationOriginateMetricType"); ok {
			if err = d.Set("nssa_default_information_originate_metric_type", vv); err != nil {
				return fmt.Errorf("Error reading nssa_default_information_originate_metric_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_default_information_originate_metric_type: %v", err)
		}
	}

	if err = d.Set("nssa_redistribution", flattenRouterOspfAreaNssaRedistribution2edl(o["nssa-redistribution"], d, "nssa_redistribution")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-redistribution"], "RouterOspfArea-NssaRedistribution"); ok {
			if err = d.Set("nssa_redistribution", vv); err != nil {
				return fmt.Errorf("Error reading nssa_redistribution: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_redistribution: %v", err)
		}
	}

	if err = d.Set("nssa_translator_role", flattenRouterOspfAreaNssaTranslatorRole2edl(o["nssa-translator-role"], d, "nssa_translator_role")); err != nil {
		if vv, ok := fortiAPIPatch(o["nssa-translator-role"], "RouterOspfArea-NssaTranslatorRole"); ok {
			if err = d.Set("nssa_translator_role", vv); err != nil {
				return fmt.Errorf("Error reading nssa_translator_role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nssa_translator_role: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("range", flattenRouterOspfAreaRange2edl(o["range"], d, "range")); err != nil {
			if vv, ok := fortiAPIPatch(o["range"], "RouterOspfArea-Range"); ok {
				if err = d.Set("range", vv); err != nil {
					return fmt.Errorf("Error reading range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("range"); ok {
			if err = d.Set("range", flattenRouterOspfAreaRange2edl(o["range"], d, "range")); err != nil {
				if vv, ok := fortiAPIPatch(o["range"], "RouterOspfArea-Range"); ok {
					if err = d.Set("range", vv); err != nil {
						return fmt.Errorf("Error reading range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading range: %v", err)
				}
			}
		}
	}

	if err = d.Set("shortcut", flattenRouterOspfAreaShortcut2edl(o["shortcut"], d, "shortcut")); err != nil {
		if vv, ok := fortiAPIPatch(o["shortcut"], "RouterOspfArea-Shortcut"); ok {
			if err = d.Set("shortcut", vv); err != nil {
				return fmt.Errorf("Error reading shortcut: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shortcut: %v", err)
		}
	}

	if err = d.Set("stub_type", flattenRouterOspfAreaStubType2edl(o["stub-type"], d, "stub_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["stub-type"], "RouterOspfArea-StubType"); ok {
			if err = d.Set("stub_type", vv); err != nil {
				return fmt.Errorf("Error reading stub_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stub_type: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterOspfAreaType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "RouterOspfArea-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("virtual_link", flattenRouterOspfAreaVirtualLink2edl(o["virtual-link"], d, "virtual_link")); err != nil {
			if vv, ok := fortiAPIPatch(o["virtual-link"], "RouterOspfArea-VirtualLink"); ok {
				if err = d.Set("virtual_link", vv); err != nil {
					return fmt.Errorf("Error reading virtual_link: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading virtual_link: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("virtual_link"); ok {
			if err = d.Set("virtual_link", flattenRouterOspfAreaVirtualLink2edl(o["virtual-link"], d, "virtual_link")); err != nil {
				if vv, ok := fortiAPIPatch(o["virtual-link"], "RouterOspfArea-VirtualLink"); ok {
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

func flattenRouterOspfAreaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAreaAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaComments2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaDefaultCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["direction"], _ = expandRouterOspfAreaFilterListDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfAreaFilterListId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["list"], _ = expandRouterOspfAreaFilterListList2edl(d, i["list"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaFilterListDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetric2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetricType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaRedistribution2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaTranslatorRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["advertise"], _ = expandRouterOspfAreaRangeAdvertise2edl(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterOspfAreaRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterOspfAreaRangePrefix2edl(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["substitute"], _ = expandRouterOspfAreaRangeSubstitute2edl(d, i["substitute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["substitute-status"], _ = expandRouterOspfAreaRangeSubstituteStatus2edl(d, i["substitute_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaRangeAdvertise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangePrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstitute2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfAreaRangeSubstituteStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaShortcut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaStubType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["authentication"], _ = expandRouterOspfAreaVirtualLinkAuthentication2edl(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication-key"], _ = expandRouterOspfAreaVirtualLinkAuthenticationKey2edl(d, i["authentication_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dead-interval"], _ = expandRouterOspfAreaVirtualLinkDeadInterval2edl(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval"], _ = expandRouterOspfAreaVirtualLinkHelloInterval2edl(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keychain"], _ = expandRouterOspfAreaVirtualLinkKeychain2edl(d, i["keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5-keychain"], _ = expandRouterOspfAreaVirtualLinkMd5Keychain2edl(d, i["md5_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandRouterOspfAreaVirtualLinkMd5Keys2edl(d, i["md5_keys"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["md5-keys"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterOspfAreaVirtualLinkName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["peer"], _ = expandRouterOspfAreaVirtualLinkPeer2edl(d, i["peer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["retransmit-interval"], _ = expandRouterOspfAreaVirtualLinkRetransmitInterval2edl(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transmit-delay"], _ = expandRouterOspfAreaVirtualLinkTransmitDelay2edl(d, i["transmit_delay"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthenticationKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkDeadInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkKeychain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keychain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkMd5Keys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaVirtualLinkMd5KeysId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterOspfAreaVirtualLinkMd5KeysKeyString2edl(d, i["key_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysKeyString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfAreaVirtualLinkName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkPeer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkRetransmitInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkTransmitDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfArea(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspfAreaAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterOspfAreaComments2edl(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("default_cost"); ok || d.HasChange("default_cost") {
		t, err := expandRouterOspfAreaDefaultCost2edl(d, v, "default_cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cost"] = t
		}
	}

	if v, ok := d.GetOk("filter_list"); ok || d.HasChange("filter_list") {
		t, err := expandRouterOspfAreaFilterList2edl(d, v, "filter_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspfAreaId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate"); ok || d.HasChange("nssa_default_information_originate") {
		t, err := expandRouterOspfAreaNssaDefaultInformationOriginate2edl(d, v, "nssa_default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric"); ok || d.HasChange("nssa_default_information_originate_metric") {
		t, err := expandRouterOspfAreaNssaDefaultInformationOriginateMetric2edl(d, v, "nssa_default_information_originate_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric_type"); ok || d.HasChange("nssa_default_information_originate_metric_type") {
		t, err := expandRouterOspfAreaNssaDefaultInformationOriginateMetricType2edl(d, v, "nssa_default_information_originate_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("nssa_redistribution"); ok || d.HasChange("nssa_redistribution") {
		t, err := expandRouterOspfAreaNssaRedistribution2edl(d, v, "nssa_redistribution")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-redistribution"] = t
		}
	}

	if v, ok := d.GetOk("nssa_translator_role"); ok || d.HasChange("nssa_translator_role") {
		t, err := expandRouterOspfAreaNssaTranslatorRole2edl(d, v, "nssa_translator_role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-translator-role"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandRouterOspfAreaRange2edl(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("shortcut"); ok || d.HasChange("shortcut") {
		t, err := expandRouterOspfAreaShortcut2edl(d, v, "shortcut")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut"] = t
		}
	}

	if v, ok := d.GetOk("stub_type"); ok || d.HasChange("stub_type") {
		t, err := expandRouterOspfAreaStubType2edl(d, v, "stub_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stub-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandRouterOspfAreaType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("virtual_link"); ok || d.HasChange("virtual_link") {
		t, err := expandRouterOspfAreaVirtualLink2edl(d, v, "virtual_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-link"] = t
		}
	}

	return &obj, nil
}
