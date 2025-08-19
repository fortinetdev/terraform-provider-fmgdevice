// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure RIPng.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRipng() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipngUpdate,
		Read:   resourceRouterRipngRead,
		Update: resourceRouterRipngUpdate,
		Delete: resourceRouterRipngDelete,

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
			"aggregate_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"listname": &schema.Schema{
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
			"garbage_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"split_horizon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"split_horizon_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"max_out_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"offset_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"offset": &schema.Schema{
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
			"passive_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
			"timeout_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"update_timer": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceRouterRipngUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterRipng(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipng resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRipng(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipng resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterRipng")

	return resourceRouterRipngRead(d, m)
}

func resourceRouterRipngDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterRipng(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRipng resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipngRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRipng(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipng resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRipng(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipng resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipngAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipngAggregateAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-AggregateAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterRipngAggregateAddressPrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterRipng-AggregateAddress-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngAggregateAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {
			v := flattenRouterRipngDistanceAccessList6(i["access-list6"], d, pre_append)
			tmp["access_list6"] = fortiAPISubPartPatch(v, "RouterRipng-Distance-AccessList6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			v := flattenRouterRipngDistanceDistance(i["distance"], d, pre_append)
			tmp["distance"] = fortiAPISubPartPatch(v, "RouterRipng-Distance-Distance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipngDistanceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-Distance-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterRipngDistancePrefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterRipng-Distance-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngDistanceAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistancePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipngDistributeListDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterRipng-DistributeList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipngDistributeListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-DistributeList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterRipngDistributeListInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterRipng-DistributeList-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			v := flattenRouterRipngDistributeListListname(i["listname"], d, pre_append)
			tmp["listname"] = fortiAPISubPartPatch(v, "RouterRipng-DistributeList-Listname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterRipngDistributeListStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterRipng-DistributeList-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenRouterRipngInterfaceFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "RouterRipng-Interface-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterRipngInterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterRipng-Interface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			v := flattenRouterRipngInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
			tmp["split_horizon"] = fortiAPISubPartPatch(v, "RouterRipng-Interface-SplitHorizon")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			v := flattenRouterRipngInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
			tmp["split_horizon_status"] = fortiAPISubPartPatch(v, "RouterRipng-Interface-SplitHorizonStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipngNeighborId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-Neighbor-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterRipngNeighborInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterRipng-Neighbor-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenRouterRipngNeighborIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "RouterRipng-Neighbor-Ip6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipngNetworkId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-Network-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterRipngNetworkPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterRipng-Network-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {
			v := flattenRouterRipngOffsetListAccessList6(i["access-list6"], d, pre_append)
			tmp["access_list6"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-AccessList6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenRouterRipngOffsetListDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipngOffsetListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterRipngOffsetListInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			v := flattenRouterRipngOffsetListOffset(i["offset"], d, pre_append)
			tmp["offset"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-Offset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterRipngOffsetListStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterRipng-OffsetList-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngOffsetListAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngPassiveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			v := flattenRouterRipngRedistributeMetric(i["metric"], d, pre_append)
			tmp["metric"] = fortiAPISubPartPatch(v, "RouterRipng-Redistribute-Metric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterRipngRedistributeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterRipng-Redistribute-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			v := flattenRouterRipngRedistributeRoutemap(i["routemap"], d, pre_append)
			tmp["routemap"] = fortiAPISubPartPatch(v, "RouterRipng-Redistribute-Routemap")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterRipngRedistributeStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterRipng-Redistribute-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipngRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipngRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRipng(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("aggregate_address", flattenRouterRipngAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["aggregate-address"], "RouterRipng-AggregateAddress"); ok {
				if err = d.Set("aggregate_address", vv); err != nil {
					return fmt.Errorf("Error reading aggregate_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading aggregate_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address"); ok {
			if err = d.Set("aggregate_address", flattenRouterRipngAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["aggregate-address"], "RouterRipng-AggregateAddress"); ok {
					if err = d.Set("aggregate_address", vv); err != nil {
						return fmt.Errorf("Error reading aggregate_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading aggregate_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_information_originate", flattenRouterRipngDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-originate"], "RouterRipng-DefaultInformationOriginate"); ok {
			if err = d.Set("default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterRipngDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-metric"], "RouterRipng-DefaultMetric"); ok {
			if err = d.Set("default_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("distance", flattenRouterRipngDistance(o["distance"], d, "distance")); err != nil {
			if vv, ok := fortiAPIPatch(o["distance"], "RouterRipng-Distance"); ok {
				if err = d.Set("distance", vv); err != nil {
					return fmt.Errorf("Error reading distance: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distance"); ok {
			if err = d.Set("distance", flattenRouterRipngDistance(o["distance"], d, "distance")); err != nil {
				if vv, ok := fortiAPIPatch(o["distance"], "RouterRipng-Distance"); ok {
					if err = d.Set("distance", vv); err != nil {
						return fmt.Errorf("Error reading distance: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading distance: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("distribute_list", flattenRouterRipngDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterRipng-DistributeList"); ok {
				if err = d.Set("distribute_list", vv); err != nil {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterRipngDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterRipng-DistributeList"); ok {
					if err = d.Set("distribute_list", vv); err != nil {
						return fmt.Errorf("Error reading distribute_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("garbage_timer", flattenRouterRipngGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["garbage-timer"], "RouterRipng-GarbageTimer"); ok {
			if err = d.Set("garbage_timer", vv); err != nil {
				return fmt.Errorf("Error reading garbage_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterRipngInterface(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "RouterRipng-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterRipngInterface(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "RouterRipng-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_out_metric", flattenRouterRipngMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-out-metric"], "RouterRipng-MaxOutMetric"); ok {
			if err = d.Set("max_out_metric", vv); err != nil {
				return fmt.Errorf("Error reading max_out_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_out_metric: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterRipngNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterRipng-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterRipngNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterRipng-Neighbor"); ok {
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
		if err = d.Set("network", flattenRouterRipngNetwork(o["network"], d, "network")); err != nil {
			if vv, ok := fortiAPIPatch(o["network"], "RouterRipng-Network"); ok {
				if err = d.Set("network", vv); err != nil {
					return fmt.Errorf("Error reading network: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterRipngNetwork(o["network"], d, "network")); err != nil {
				if vv, ok := fortiAPIPatch(o["network"], "RouterRipng-Network"); ok {
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
		if err = d.Set("offset_list", flattenRouterRipngOffsetList(o["offset-list"], d, "offset_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["offset-list"], "RouterRipng-OffsetList"); ok {
				if err = d.Set("offset_list", vv); err != nil {
					return fmt.Errorf("Error reading offset_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading offset_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offset_list"); ok {
			if err = d.Set("offset_list", flattenRouterRipngOffsetList(o["offset-list"], d, "offset_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["offset-list"], "RouterRipng-OffsetList"); ok {
					if err = d.Set("offset_list", vv); err != nil {
						return fmt.Errorf("Error reading offset_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading offset_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("passive_interface", flattenRouterRipngPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-interface"], "RouterRipng-PassiveInterface"); ok {
			if err = d.Set("passive_interface", vv); err != nil {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterRipngRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterRipng-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterRipngRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterRipng-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("timeout_timer", flattenRouterRipngTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-timer"], "RouterRipng-TimeoutTimer"); ok {
			if err = d.Set("timeout_timer", vv); err != nil {
				return fmt.Errorf("Error reading timeout_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("update_timer", flattenRouterRipngUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-timer"], "RouterRipng-UpdateTimer"); ok {
			if err = d.Set("update_timer", vv); err != nil {
				return fmt.Errorf("Error reading update_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	return nil
}

func flattenRouterRipngFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRipngAggregateAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRipngAggregateAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterRipngAggregateAddressPrefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngAggregateAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngAggregateAddressPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-list6"], _ = expandRouterRipngDistanceAccessList6(d, i["access_list6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distance"], _ = expandRouterRipngDistanceDistance(d, i["distance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipngDistanceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterRipngDistancePrefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngDistanceAccessList6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngDistanceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistanceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistancePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["direction"], _ = expandRouterRipngDistributeListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipngDistributeListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterRipngDistributeListInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listname"], _ = expandRouterRipngDistributeListListname(d, i["listname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterRipngDistributeListStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngDistributeListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngDistributeListListname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngDistributeListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngGarbageTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandRouterRipngInterfaceFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterRipngInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-horizon"], _ = expandRouterRipngInterfaceSplitHorizon(d, i["split_horizon"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-horizon-status"], _ = expandRouterRipngInterfaceSplitHorizonStatus(d, i["split_horizon_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngInterfaceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngInterfaceSplitHorizon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceSplitHorizonStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngMaxOutMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRipngNeighborId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterRipngNeighborInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandRouterRipngNeighborIp6(d, i["ip6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngNeighborId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNeighborInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngNeighborIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRipngNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterRipngNetworkPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-list6"], _ = expandRouterRipngOffsetListAccessList6(d, i["access_list6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandRouterRipngOffsetListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipngOffsetListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterRipngOffsetListInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["offset"], _ = expandRouterRipngOffsetListOffset(d, i["offset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterRipngOffsetListStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngOffsetListAccessList6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngOffsetListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngOffsetListOffset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngPassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["metric"], _ = expandRouterRipngRedistributeMetric(d, i["metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterRipngRedistributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["routemap"], _ = expandRouterRipngRedistributeRoutemap(d, i["routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterRipngRedistributeStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipngRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipngRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngTimeoutTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngUpdateTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRipng(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggregate_address"); ok || d.HasChange("aggregate_address") {
		t, err := expandRouterRipngAggregateAddress(d, v, "aggregate_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address"] = t
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok || d.HasChange("default_information_originate") {
		t, err := expandRouterRipngDefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok || d.HasChange("default_metric") {
		t, err := expandRouterRipngDefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandRouterRipngDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok || d.HasChange("distribute_list") {
		t, err := expandRouterRipngDistributeList(d, v, "distribute_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list"] = t
		}
	}

	if v, ok := d.GetOk("garbage_timer"); ok || d.HasChange("garbage_timer") {
		t, err := expandRouterRipngGarbageTimer(d, v, "garbage_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["garbage-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterRipngInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("max_out_metric"); ok || d.HasChange("max_out_metric") {
		t, err := expandRouterRipngMaxOutMetric(d, v, "max_out_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-out-metric"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterRipngNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		t, err := expandRouterRipngNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("offset_list"); ok || d.HasChange("offset_list") {
		t, err := expandRouterRipngOffsetList(d, v, "offset_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offset-list"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		t, err := expandRouterRipngPassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterRipngRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("timeout_timer"); ok || d.HasChange("timeout_timer") {
		t, err := expandRouterRipngTimeoutTimer(d, v, "timeout_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-timer"] = t
		}
	}

	if v, ok := d.GetOk("update_timer"); ok || d.HasChange("update_timer") {
		t, err := expandRouterRipngUpdateTimer(d, v, "update_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-timer"] = t
		}
	}

	return &obj, nil
}
