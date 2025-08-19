// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure RIP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipUpdate,
		Read:   resourceRouterRipRead,
		Update: resourceRouterRipUpdate,
		Delete: resourceRouterRipDelete,

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
						"access_list": &schema.Schema{
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
							Computed: true,
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
			"distribute_list": &schema.Schema{
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
						"auth_keychain": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_string": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"receive_version": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"send_version": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"send_version2_broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"ip": &schema.Schema{
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
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
						"access_list": &schema.Schema{
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
			"recv_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"version": &schema.Schema{
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

func resourceRouterRipUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterRip(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRip resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRip(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterRip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterRip")

	return resourceRouterRipRead(d, m)
}

func resourceRouterRipDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterRip(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterRip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRip resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipDistanceAccessList(i["access-list"], d, pre_append)
			tmp["access_list"] = fortiAPISubPartPatch(v, "RouterRip-Distance-AccessList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			v := flattenRouterRipDistanceDistance(i["distance"], d, pre_append)
			tmp["distance"] = fortiAPISubPartPatch(v, "RouterRip-Distance-Distance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipDistanceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRip-Distance-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterRipDistancePrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterRip-Distance-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipDistanceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDistancePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipDistributeListDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterRip-DistributeList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipDistributeListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRip-DistributeList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterRipDistributeListInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterRip-DistributeList-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			v := flattenRouterRipDistributeListListname(i["listname"], d, pre_append)
			tmp["listname"] = fortiAPISubPartPatch(v, "RouterRip-DistributeList-Listname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterRipDistributeListStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterRip-DistributeList-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {
			v := flattenRouterRipInterfaceAuthKeychain(i["auth-keychain"], d, pre_append)
			tmp["auth_keychain"] = fortiAPISubPartPatch(v, "RouterRip-Interface-AuthKeychain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			v := flattenRouterRipInterfaceAuthMode(i["auth-mode"], d, pre_append)
			tmp["auth_mode"] = fortiAPISubPartPatch(v, "RouterRip-Interface-AuthMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenRouterRipInterfaceFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "RouterRip-Interface-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterRipInterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterRip-Interface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {
			v := flattenRouterRipInterfaceReceiveVersion(i["receive-version"], d, pre_append)
			tmp["receive_version"] = fortiAPISubPartPatch(v, "RouterRip-Interface-ReceiveVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {
			v := flattenRouterRipInterfaceSendVersion(i["send-version"], d, pre_append)
			tmp["send_version"] = fortiAPISubPartPatch(v, "RouterRip-Interface-SendVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {
			v := flattenRouterRipInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append)
			tmp["send_version2_broadcast"] = fortiAPISubPartPatch(v, "RouterRip-Interface-SendVersion2Broadcast")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			v := flattenRouterRipInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
			tmp["split_horizon"] = fortiAPISubPartPatch(v, "RouterRip-Interface-SplitHorizon")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			v := flattenRouterRipInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
			tmp["split_horizon_status"] = fortiAPISubPartPatch(v, "RouterRip-Interface-SplitHorizonStatus")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenRouterRipInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipNeighborId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRip-Neighbor-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenRouterRipNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "RouterRip-Neighbor-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipNetworkId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRip-Network-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterRipNetworkPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterRip-Network-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterRipOffsetListAccessList(i["access-list"], d, pre_append)
			tmp["access_list"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-AccessList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenRouterRipOffsetListDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterRipOffsetListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenRouterRipOffsetListInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			v := flattenRouterRipOffsetListOffset(i["offset"], d, pre_append)
			tmp["offset"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-Offset")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterRipOffsetListStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterRip-OffsetList-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterRipOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipPassiveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipRecvBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := i["metric"]; ok {
		result["metric"] = flattenRouterRipRedistributeMetric(i["metric"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenRouterRipRedistributeName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "routemap"
	if _, ok := i["routemap"]; ok {
		result["routemap"] = flattenRouterRipRedistributeRoutemap(i["routemap"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenRouterRipRedistributeStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterRipRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterRipRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterRip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default_information_originate", flattenRouterRipDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-information-originate"], "RouterRip-DefaultInformationOriginate"); ok {
			if err = d.Set("default_information_originate", vv); err != nil {
				return fmt.Errorf("Error reading default_information_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterRipDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-metric"], "RouterRip-DefaultMetric"); ok {
			if err = d.Set("default_metric", vv); err != nil {
				return fmt.Errorf("Error reading default_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("distance", flattenRouterRipDistance(o["distance"], d, "distance")); err != nil {
			if vv, ok := fortiAPIPatch(o["distance"], "RouterRip-Distance"); ok {
				if err = d.Set("distance", vv); err != nil {
					return fmt.Errorf("Error reading distance: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distance"); ok {
			if err = d.Set("distance", flattenRouterRipDistance(o["distance"], d, "distance")); err != nil {
				if vv, ok := fortiAPIPatch(o["distance"], "RouterRip-Distance"); ok {
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
		if err = d.Set("distribute_list", flattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterRip-DistributeList"); ok {
				if err = d.Set("distribute_list", vv); err != nil {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["distribute-list"], "RouterRip-DistributeList"); ok {
					if err = d.Set("distribute_list", vv); err != nil {
						return fmt.Errorf("Error reading distribute_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("garbage_timer", flattenRouterRipGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["garbage-timer"], "RouterRip-GarbageTimer"); ok {
			if err = d.Set("garbage_timer", vv); err != nil {
				return fmt.Errorf("Error reading garbage_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterRipInterface(o["interface"], d, "interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["interface"], "RouterRip-Interface"); ok {
				if err = d.Set("interface", vv); err != nil {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterRipInterface(o["interface"], d, "interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["interface"], "RouterRip-Interface"); ok {
					if err = d.Set("interface", vv); err != nil {
						return fmt.Errorf("Error reading interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_out_metric", flattenRouterRipMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-out-metric"], "RouterRip-MaxOutMetric"); ok {
			if err = d.Set("max_out_metric", vv); err != nil {
				return fmt.Errorf("Error reading max_out_metric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_out_metric: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterRipNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterRip-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterRipNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterRip-Neighbor"); ok {
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
		if err = d.Set("network", flattenRouterRipNetwork(o["network"], d, "network")); err != nil {
			if vv, ok := fortiAPIPatch(o["network"], "RouterRip-Network"); ok {
				if err = d.Set("network", vv); err != nil {
					return fmt.Errorf("Error reading network: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterRipNetwork(o["network"], d, "network")); err != nil {
				if vv, ok := fortiAPIPatch(o["network"], "RouterRip-Network"); ok {
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
		if err = d.Set("offset_list", flattenRouterRipOffsetList(o["offset-list"], d, "offset_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["offset-list"], "RouterRip-OffsetList"); ok {
				if err = d.Set("offset_list", vv); err != nil {
					return fmt.Errorf("Error reading offset_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading offset_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offset_list"); ok {
			if err = d.Set("offset_list", flattenRouterRipOffsetList(o["offset-list"], d, "offset_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["offset-list"], "RouterRip-OffsetList"); ok {
					if err = d.Set("offset_list", vv); err != nil {
						return fmt.Errorf("Error reading offset_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading offset_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("passive_interface", flattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-interface"], "RouterRip-PassiveInterface"); ok {
			if err = d.Set("passive_interface", vv); err != nil {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("recv_buffer_size", flattenRouterRipRecvBufferSize(o["recv-buffer-size"], d, "recv_buffer_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["recv-buffer-size"], "RouterRip-RecvBufferSize"); ok {
			if err = d.Set("recv_buffer_size", vv); err != nil {
				return fmt.Errorf("Error reading recv_buffer_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recv_buffer_size: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterRipRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterRip-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterRipRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterRip-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("timeout_timer", flattenRouterRipTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-timer"], "RouterRip-TimeoutTimer"); ok {
			if err = d.Set("timeout_timer", vv); err != nil {
				return fmt.Errorf("Error reading timeout_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("update_timer", flattenRouterRipUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-timer"], "RouterRip-UpdateTimer"); ok {
			if err = d.Set("update_timer", vv); err != nil {
				return fmt.Errorf("Error reading update_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	if err = d.Set("version", flattenRouterRipVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "RouterRip-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenRouterRipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterRipDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["access-list"], _ = expandRouterRipDistanceAccessList(d, i["access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["distance"], _ = expandRouterRipDistanceDistance(d, i["distance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipDistanceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterRipDistancePrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipDistanceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipDistanceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistanceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistancePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterRipDistributeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["direction"], _ = expandRouterRipDistributeListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipDistributeListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterRipDistributeListInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listname"], _ = expandRouterRipDistributeListListname(d, i["listname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterRipDistributeListStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipDistributeListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipDistributeListListname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipDistributeListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipGarbageTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-keychain"], _ = expandRouterRipInterfaceAuthKeychain(d, i["auth_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-mode"], _ = expandRouterRipInterfaceAuthMode(d, i["auth_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-string"], _ = expandRouterRipInterfaceAuthString(d, i["auth_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandRouterRipInterfaceFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterRipInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["receive-version"], _ = expandRouterRipInterfaceReceiveVersion(d, i["receive_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-version"], _ = expandRouterRipInterfaceSendVersion(d, i["send_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-version2-broadcast"], _ = expandRouterRipInterfaceSendVersion2Broadcast(d, i["send_version2_broadcast"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-horizon"], _ = expandRouterRipInterfaceSplitHorizon(d, i["split_horizon"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["split-horizon-status"], _ = expandRouterRipInterfaceSplitHorizonStatus(d, i["split_horizon_status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipInterfaceAuthKeychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceAuthMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceAuthString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandRouterRipInterfaceReceiveVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceSendVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipInterfaceSendVersion2Broadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizonStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipMaxOutMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRipNeighborId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandRouterRipNeighborIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipNeighborId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRipNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterRipNetworkPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterRipOffsetList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["access-list"], _ = expandRouterRipOffsetListAccessList(d, i["access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandRouterRipOffsetListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterRipOffsetListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandRouterRipOffsetListInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["offset"], _ = expandRouterRipOffsetListOffset(d, i["offset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterRipOffsetListStatus(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterRipOffsetListAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipOffsetListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipOffsetListOffset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipPassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipRecvBufferSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "metric"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric"], _ = expandRouterRipRedistributeMetric(d, i["metric"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandRouterRipRedistributeName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "routemap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["routemap"], _ = expandRouterRipRedistributeRoutemap(d, i["routemap"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandRouterRipRedistributeStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandRouterRipRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterRipRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipTimeoutTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipUpdateTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_information_originate"); ok || d.HasChange("default_information_originate") {
		t, err := expandRouterRipDefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok || d.HasChange("default_metric") {
		t, err := expandRouterRipDefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandRouterRipDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok || d.HasChange("distribute_list") {
		t, err := expandRouterRipDistributeList(d, v, "distribute_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list"] = t
		}
	}

	if v, ok := d.GetOk("garbage_timer"); ok || d.HasChange("garbage_timer") {
		t, err := expandRouterRipGarbageTimer(d, v, "garbage_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["garbage-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterRipInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("max_out_metric"); ok || d.HasChange("max_out_metric") {
		t, err := expandRouterRipMaxOutMetric(d, v, "max_out_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-out-metric"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterRipNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		t, err := expandRouterRipNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("offset_list"); ok || d.HasChange("offset_list") {
		t, err := expandRouterRipOffsetList(d, v, "offset_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offset-list"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		t, err := expandRouterRipPassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("recv_buffer_size"); ok || d.HasChange("recv_buffer_size") {
		t, err := expandRouterRipRecvBufferSize(d, v, "recv_buffer_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recv-buffer-size"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterRipRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("timeout_timer"); ok || d.HasChange("timeout_timer") {
		t, err := expandRouterRipTimeoutTimer(d, v, "timeout_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-timer"] = t
		}
	}

	if v, ok := d.GetOk("update_timer"); ok || d.HasChange("update_timer") {
		t, err := expandRouterRipUpdateTimer(d, v, "update_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-timer"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandRouterRipVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
