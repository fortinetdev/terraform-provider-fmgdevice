// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv6 firewall addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6Create,
		Read:   resourceFirewallAddress6Read,
		Update: resourceFirewallAddress6Update,
		Delete: resourceFirewallAddress6Delete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_image_base64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"cache_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"country": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"epg_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"global_object": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"macaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sdn": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sdn_addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sdn_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_segment": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
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
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tenant": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"visibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wildcard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"epg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"net_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
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
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"template": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wildcard": &schema.Schema{
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

func resourceFirewallAddress6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6 resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6 resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6Read(d, m)
}

func resourceFirewallAddress6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddress6Read(d, m)
}

func resourceFirewallAddress6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAddress6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAddress6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6ImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6CacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Country(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := i["_image-base64"]; ok {
			v := flattenFirewallAddress6DynamicMappingImageBase64(i["_image-base64"], d, pre_append)
			tmp["_image_base64"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-ImageBase64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallAddress6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := i["cache-ttl"]; ok {
			v := flattenFirewallAddress6DynamicMappingCacheTtl(i["cache-ttl"], d, pre_append)
			tmp["cache_ttl"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-CacheTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallAddress6DynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallAddress6DynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {
			v := flattenFirewallAddress6DynamicMappingCountry(i["country"], d, pre_append)
			tmp["country"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Country")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenFirewallAddress6DynamicMappingEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := i["end-mac"]; ok {
			v := flattenFirewallAddress6DynamicMappingEndMac(i["end-mac"], d, pre_append)
			tmp["end_mac"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-EndMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := i["epg-name"]; ok {
			v := flattenFirewallAddress6DynamicMappingEpgName(i["epg-name"], d, pre_append)
			tmp["epg_name"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-EpgName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenFirewallAddress6DynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenFirewallAddress6DynamicMappingFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenFirewallAddress6DynamicMappingFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := i["global-object"]; ok {
			v := flattenFirewallAddress6DynamicMappingGlobalObject(i["global-object"], d, pre_append)
			tmp["global_object"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-GlobalObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenFirewallAddress6DynamicMappingHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			v := flattenFirewallAddress6DynamicMappingHostType(i["host-type"], d, pre_append)
			tmp["host_type"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-HostType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFirewallAddress6DynamicMappingIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := i["macaddr"]; ok {
			v := flattenFirewallAddress6DynamicMappingMacaddr(i["macaddr"], d, pre_append)
			tmp["macaddr"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Macaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenFirewallAddress6DynamicMappingObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-ObjId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenFirewallAddress6DynamicMappingRouteTag(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := i["sdn"]; ok {
			v := flattenFirewallAddress6DynamicMappingSdn(i["sdn"], d, pre_append)
			tmp["sdn"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Sdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := i["sdn-addr-type"]; ok {
			v := flattenFirewallAddress6DynamicMappingSdnAddrType(i["sdn-addr-type"], d, pre_append)
			tmp["sdn_addr_type"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-SdnAddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := i["sdn-tag"]; ok {
			v := flattenFirewallAddress6DynamicMappingSdnTag(i["sdn-tag"], d, pre_append)
			tmp["sdn_tag"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-SdnTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenFirewallAddress6DynamicMappingStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := i["start-mac"]; ok {
			v := flattenFirewallAddress6DynamicMappingStartMac(i["start-mac"], d, pre_append)
			tmp["start_mac"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-StartMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_segment"
		if _, ok := i["subnet-segment"]; ok {
			v := flattenFirewallAddress6DynamicMappingSubnetSegment(i["subnet-segment"], d, pre_append)
			tmp["subnet_segment"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-SubnetSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddress6DynamicMappingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := i["template"]; ok {
			v := flattenFirewallAddress6DynamicMappingTemplate(i["template"], d, pre_append)
			tmp["template"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Template")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := i["tenant"]; ok {
			v := flattenFirewallAddress6DynamicMappingTenant(i["tenant"], d, pre_append)
			tmp["tenant"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Tenant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddress6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallAddress6DynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := i["visibility"]; ok {
			v := flattenFirewallAddress6DynamicMappingVisibility(i["visibility"], d, pre_append)
			tmp["visibility"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Visibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := i["wildcard"]; ok {
			v := flattenFirewallAddress6DynamicMappingWildcard(i["wildcard"], d, pre_append)
			tmp["wildcard"] = fortiAPISubPartPatch(v, "FirewallAddress6-DynamicMapping-Wildcard")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6DynamicMappingImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallAddress6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingEpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingSdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingTenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenFirewallAddress6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6EndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6EpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6FabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Filter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Fqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6GlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Host(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6HostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Ip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6List(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFirewallAddress6ListIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAddress6-List-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := i["net-id"]; ok {
			v := flattenFirewallAddress6ListNetId(i["net-id"], d, pre_append)
			tmp["net_id"] = fortiAPISubPartPatch(v, "FirewallAddress6-List-NetId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenFirewallAddress6ListObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "FirewallAddress6-List-ObjId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6ListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6ListNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6ListObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Macaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6ObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6RouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Sdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6SdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6SdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6StartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6SubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6SubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddress6SubnetSegmentType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddress6-SubnetSegment-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenFirewallAddress6SubnetSegmentValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "FirewallAddress6-SubnetSegment-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6SubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6SubnetSegmentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6SubnetSegmentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Tagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenFirewallAddress6TaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallAddress6-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAddress6TaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddress6TaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddress6-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6TaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6TaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6Template(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6Tenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6Tags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6Wildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectFirewallAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallAddress6ImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallAddress6-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenFirewallAddress6CacheTtl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "FirewallAddress6-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallAddress6Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallAddress6-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallAddress6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallAddress6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenFirewallAddress6Country(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "FirewallAddress6-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallAddress6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddress6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallAddress6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddress6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_ip", flattenFirewallAddress6EndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "FirewallAddress6-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenFirewallAddress6EndMac(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "FirewallAddress6-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenFirewallAddress6EpgName(o["epg-name"], d, "epg_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["epg-name"], "FirewallAddress6-EpgName"); ok {
			if err = d.Set("epg_name", vv); err != nil {
				return fmt.Errorf("Error reading epg_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallAddress6FabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallAddress6-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("filter", flattenFirewallAddress6Filter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "FirewallAddress6-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallAddress6Fqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "FirewallAddress6-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallAddress6GlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallAddress6-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallAddress6Host(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallAddress6-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_type", flattenFirewallAddress6HostType(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "FirewallAddress6-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFirewallAddress6Ip6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FirewallAddress6-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenFirewallAddress6List(o["list"], d, "list")); err != nil {
			if vv, ok := fortiAPIPatch(o["list"], "FirewallAddress6-List"); ok {
				if err = d.Set("list", vv); err != nil {
					return fmt.Errorf("Error reading list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenFirewallAddress6List(o["list"], d, "list")); err != nil {
				if vv, ok := fortiAPIPatch(o["list"], "FirewallAddress6-List"); ok {
					if err = d.Set("list", vv); err != nil {
						return fmt.Errorf("Error reading list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if err = d.Set("macaddr", flattenFirewallAddress6Macaddr(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "FirewallAddress6-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAddress6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenFirewallAddress6ObjId(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "FirewallAddress6-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenFirewallAddress6RouteTag(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "FirewallAddress6-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("sdn", flattenFirewallAddress6Sdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "FirewallAddress6-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenFirewallAddress6SdnAddrType(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-addr-type"], "FirewallAddress6-SdnAddrType"); ok {
			if err = d.Set("sdn_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading sdn_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenFirewallAddress6SdnTag(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-tag"], "FirewallAddress6-SdnTag"); ok {
			if err = d.Set("sdn_tag", vv); err != nil {
				return fmt.Errorf("Error reading sdn_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenFirewallAddress6StartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "FirewallAddress6-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenFirewallAddress6StartMac(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "FirewallAddress6-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnet_segment", flattenFirewallAddress6SubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6-SubnetSegment"); ok {
				if err = d.Set("subnet_segment", vv); err != nil {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnet_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnet_segment"); ok {
			if err = d.Set("subnet_segment", flattenFirewallAddress6SubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6-SubnetSegment"); ok {
					if err = d.Set("subnet_segment", vv); err != nil {
						return fmt.Errorf("Error reading subnet_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddress6-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddress6-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("template", flattenFirewallAddress6Template(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "FirewallAddress6-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	if err = d.Set("tenant", flattenFirewallAddress6Tenant(o["tenant"], d, "tenant")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant"], "FirewallAddress6-Tenant"); ok {
			if err = d.Set("tenant", vv); err != nil {
				return fmt.Errorf("Error reading tenant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAddress6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAddress6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddress6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallAddress6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("tags", flattenFirewallAddress6Tags(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "FirewallAddress6-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenFirewallAddress6Wildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "FirewallAddress6-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6ImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6CacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Country(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_image-base64"], _ = expandFirewallAddress6DynamicMappingImageBase64(d, i["_image_base64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAddress6DynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cache-ttl"], _ = expandFirewallAddress6DynamicMappingCacheTtl(d, i["cache_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallAddress6DynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallAddress6DynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["country"], _ = expandFirewallAddress6DynamicMappingCountry(d, i["country"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandFirewallAddress6DynamicMappingEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-mac"], _ = expandFirewallAddress6DynamicMappingEndMac(d, i["end_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["epg-name"], _ = expandFirewallAddress6DynamicMappingEpgName(d, i["epg_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandFirewallAddress6DynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandFirewallAddress6DynamicMappingFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandFirewallAddress6DynamicMappingFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["global-object"], _ = expandFirewallAddress6DynamicMappingGlobalObject(d, i["global_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandFirewallAddress6DynamicMappingHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host-type"], _ = expandFirewallAddress6DynamicMappingHostType(d, i["host_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFirewallAddress6DynamicMappingIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["macaddr"], _ = expandFirewallAddress6DynamicMappingMacaddr(d, i["macaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandFirewallAddress6DynamicMappingObjId(d, i["obj_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandFirewallAddress6DynamicMappingRouteTag(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn"], _ = expandFirewallAddress6DynamicMappingSdn(d, i["sdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn-addr-type"], _ = expandFirewallAddress6DynamicMappingSdnAddrType(d, i["sdn_addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn-tag"], _ = expandFirewallAddress6DynamicMappingSdnTag(d, i["sdn_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandFirewallAddress6DynamicMappingStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-mac"], _ = expandFirewallAddress6DynamicMappingStartMac(d, i["start_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAddress6DynamicMappingSubnetSegment(d, i["subnet_segment"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["subnet-segment"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddress6DynamicMappingTags(d, i["tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["template"], _ = expandFirewallAddress6DynamicMappingTemplate(d, i["template"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant"], _ = expandFirewallAddress6DynamicMappingTenant(d, i["tenant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddress6DynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallAddress6DynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["visibility"], _ = expandFirewallAddress6DynamicMappingVisibility(d, i["visibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wildcard"], _ = expandFirewallAddress6DynamicMappingWildcard(d, i["wildcard"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6DynamicMappingImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallAddress6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingEpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingHostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingSdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingTenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6EndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6EpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6FabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Filter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Fqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6GlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Host(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6HostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Ip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFirewallAddress6ListIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["net-id"], _ = expandFirewallAddress6ListNetId(d, i["net_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandFirewallAddress6ListObjId(d, i["obj_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6ListIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6ListNetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6ListObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Macaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6ObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6RouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Sdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6SdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6SdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6StartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6SubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6SubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddress6SubnetSegmentType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandFirewallAddress6SubnetSegmentValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6SubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6SubnetSegmentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6SubnetSegmentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Tagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandFirewallAddress6TaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAddress6TaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddress6TaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6TaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6Template(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6Tenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6Tags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6Wildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallAddress6ImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok || d.HasChange("cache_ttl") {
		t, err := expandFirewallAddress6CacheTtl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallAddress6Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallAddress6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandFirewallAddress6Country(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallAddress6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandFirewallAddress6EndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok || d.HasChange("end_mac") {
		t, err := expandFirewallAddress6EndMac(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok || d.HasChange("epg_name") {
		t, err := expandFirewallAddress6EpgName(d, v, "epg_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallAddress6FabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandFirewallAddress6Filter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandFirewallAddress6Fqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallAddress6GlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallAddress6Host(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandFirewallAddress6HostType(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandFirewallAddress6Ip6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok || d.HasChange("list") {
		t, err := expandFirewallAddress6List(d, v, "list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandFirewallAddress6Macaddr(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddress6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandFirewallAddress6ObjId(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandFirewallAddress6RouteTag(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandFirewallAddress6Sdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok || d.HasChange("sdn_addr_type") {
		t, err := expandFirewallAddress6SdnAddrType(d, v, "sdn_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok || d.HasChange("sdn_tag") {
		t, err := expandFirewallAddress6SdnTag(d, v, "sdn_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandFirewallAddress6StartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok || d.HasChange("start_mac") {
		t, err := expandFirewallAddress6StartMac(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok || d.HasChange("subnet_segment") {
		t, err := expandFirewallAddress6SubnetSegment(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallAddress6Tagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandFirewallAddress6Template(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok || d.HasChange("tenant") {
		t, err := expandFirewallAddress6Tenant(d, v, "tenant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAddress6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallAddress6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandFirewallAddress6Tags(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandFirewallAddress6Wildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	return &obj, nil
}
