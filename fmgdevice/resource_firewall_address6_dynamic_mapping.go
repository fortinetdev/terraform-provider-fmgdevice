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

func resourceFirewallAddress6DynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6DynamicMappingCreate,
		Read:   resourceFirewallAddress6DynamicMappingRead,
		Update: resourceFirewallAddress6DynamicMappingUpdate,
		Delete: resourceFirewallAddress6DynamicMappingDelete,

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
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAddress6DynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	obj, err := getObjectFirewallAddress6DynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6DynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("_scope")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6DynamicMapping(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6DynamicMapping(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6DynamicMapping resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6DynamicMapping(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6DynamicMapping resource: %v", err)
		}

	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceFirewallAddress6DynamicMappingRead(d, m)
}

func resourceFirewallAddress6DynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	obj, err := getObjectFirewallAddress6DynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6DynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6DynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6DynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceFirewallAddress6DynamicMappingRead(d, m)
}

func resourceFirewallAddress6DynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAddress6DynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6DynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6DynamicMappingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	address6 := d.Get("address6").(string)
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
	if address6 == "" {
		address6 = importOptionChecking(m.(*FortiClient).Cfg, "address6")
		if address6 == "" {
			return fmt.Errorf("Parameter address6 is missing")
		}
		if err = d.Set("address6", address6); err != nil {
			return fmt.Errorf("Error set params address6: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	o, err := c.ReadFirewallAddress6DynamicMapping(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6DynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6DynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6DynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6DynamicMappingImageBase642edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6DynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallAddress6DynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6DynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingCacheTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingCountry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingEndMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingEpgName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFabricObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingGlobalObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingHostType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingIp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingMacaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingObjId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingSdnAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSdnTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingStartMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegment2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenFirewallAddress6DynamicMappingSubnetSegmentValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "FirewallAddress6DynamicMapping-SubnetSegment-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingSubnetSegmentValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingTemplate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddress6DynamicMappingTenant2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingUuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingVisibility2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6DynamicMappingWildcard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectFirewallAddress6DynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallAddress6DynamicMappingImageBase642edl(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallAddress6DynamicMapping-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenFirewallAddress6DynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "FirewallAddress6DynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenFirewallAddress6DynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "FirewallAddress6DynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("cache_ttl", flattenFirewallAddress6DynamicMappingCacheTtl2edl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "FirewallAddress6DynamicMapping-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallAddress6DynamicMappingColor2edl(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallAddress6DynamicMapping-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallAddress6DynamicMappingComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallAddress6DynamicMapping-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenFirewallAddress6DynamicMappingCountry2edl(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "FirewallAddress6DynamicMapping-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenFirewallAddress6DynamicMappingEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "FirewallAddress6DynamicMapping-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenFirewallAddress6DynamicMappingEndMac2edl(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "FirewallAddress6DynamicMapping-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenFirewallAddress6DynamicMappingEpgName2edl(o["epg-name"], d, "epg_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["epg-name"], "FirewallAddress6DynamicMapping-EpgName"); ok {
			if err = d.Set("epg_name", vv); err != nil {
				return fmt.Errorf("Error reading epg_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallAddress6DynamicMappingFabricObject2edl(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallAddress6DynamicMapping-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("filter", flattenFirewallAddress6DynamicMappingFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "FirewallAddress6DynamicMapping-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallAddress6DynamicMappingFqdn2edl(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "FirewallAddress6DynamicMapping-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallAddress6DynamicMappingGlobalObject2edl(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallAddress6DynamicMapping-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallAddress6DynamicMappingHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "FirewallAddress6DynamicMapping-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_type", flattenFirewallAddress6DynamicMappingHostType2edl(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "FirewallAddress6DynamicMapping-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFirewallAddress6DynamicMappingIp62edl(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FirewallAddress6DynamicMapping-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("macaddr", flattenFirewallAddress6DynamicMappingMacaddr2edl(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "FirewallAddress6DynamicMapping-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenFirewallAddress6DynamicMappingObjId2edl(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "FirewallAddress6DynamicMapping-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenFirewallAddress6DynamicMappingRouteTag2edl(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "FirewallAddress6DynamicMapping-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("sdn", flattenFirewallAddress6DynamicMappingSdn2edl(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "FirewallAddress6DynamicMapping-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenFirewallAddress6DynamicMappingSdnAddrType2edl(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-addr-type"], "FirewallAddress6DynamicMapping-SdnAddrType"); ok {
			if err = d.Set("sdn_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading sdn_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenFirewallAddress6DynamicMappingSdnTag2edl(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-tag"], "FirewallAddress6DynamicMapping-SdnTag"); ok {
			if err = d.Set("sdn_tag", vv); err != nil {
				return fmt.Errorf("Error reading sdn_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenFirewallAddress6DynamicMappingStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "FirewallAddress6DynamicMapping-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenFirewallAddress6DynamicMappingStartMac2edl(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "FirewallAddress6DynamicMapping-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnet_segment", flattenFirewallAddress6DynamicMappingSubnetSegment2edl(o["subnet-segment"], d, "subnet_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6DynamicMapping-SubnetSegment"); ok {
				if err = d.Set("subnet_segment", vv); err != nil {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnet_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnet_segment"); ok {
			if err = d.Set("subnet_segment", flattenFirewallAddress6DynamicMappingSubnetSegment2edl(o["subnet-segment"], d, "subnet_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnet-segment"], "FirewallAddress6DynamicMapping-SubnetSegment"); ok {
					if err = d.Set("subnet_segment", vv); err != nil {
						return fmt.Errorf("Error reading subnet_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			}
		}
	}

	if err = d.Set("tags", flattenFirewallAddress6DynamicMappingTags2edl(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "FirewallAddress6DynamicMapping-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	if err = d.Set("template", flattenFirewallAddress6DynamicMappingTemplate2edl(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "FirewallAddress6DynamicMapping-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	if err = d.Set("tenant", flattenFirewallAddress6DynamicMappingTenant2edl(o["tenant"], d, "tenant")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant"], "FirewallAddress6DynamicMapping-Tenant"); ok {
			if err = d.Set("tenant", vv); err != nil {
				return fmt.Errorf("Error reading tenant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAddress6DynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAddress6DynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddress6DynamicMappingUuid2edl(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallAddress6DynamicMapping-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("visibility", flattenFirewallAddress6DynamicMappingVisibility2edl(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "FirewallAddress6DynamicMapping-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenFirewallAddress6DynamicMappingWildcard2edl(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "FirewallAddress6DynamicMapping-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6DynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6DynamicMappingImageBase642edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6DynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallAddress6DynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6DynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingCacheTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingCountry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingEndMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingEpgName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFabricObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingGlobalObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingHostType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingIp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingMacaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingObjId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingSdnAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSdnTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingStartMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandFirewallAddress6DynamicMappingSubnetSegmentValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingSubnetSegmentValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingTemplate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddress6DynamicMappingTenant2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingUuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingVisibility2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6DynamicMappingWildcard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6DynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallAddress6DynamicMappingImageBase642edl(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandFirewallAddress6DynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok || d.HasChange("cache_ttl") {
		t, err := expandFirewallAddress6DynamicMappingCacheTtl2edl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallAddress6DynamicMappingColor2edl(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallAddress6DynamicMappingComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandFirewallAddress6DynamicMappingCountry2edl(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandFirewallAddress6DynamicMappingEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok || d.HasChange("end_mac") {
		t, err := expandFirewallAddress6DynamicMappingEndMac2edl(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok || d.HasChange("epg_name") {
		t, err := expandFirewallAddress6DynamicMappingEpgName2edl(d, v, "epg_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallAddress6DynamicMappingFabricObject2edl(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandFirewallAddress6DynamicMappingFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandFirewallAddress6DynamicMappingFqdn2edl(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallAddress6DynamicMappingGlobalObject2edl(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandFirewallAddress6DynamicMappingHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandFirewallAddress6DynamicMappingHostType2edl(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandFirewallAddress6DynamicMappingIp62edl(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandFirewallAddress6DynamicMappingMacaddr2edl(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandFirewallAddress6DynamicMappingObjId2edl(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandFirewallAddress6DynamicMappingRouteTag2edl(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandFirewallAddress6DynamicMappingSdn2edl(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok || d.HasChange("sdn_addr_type") {
		t, err := expandFirewallAddress6DynamicMappingSdnAddrType2edl(d, v, "sdn_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok || d.HasChange("sdn_tag") {
		t, err := expandFirewallAddress6DynamicMappingSdnTag2edl(d, v, "sdn_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandFirewallAddress6DynamicMappingStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok || d.HasChange("start_mac") {
		t, err := expandFirewallAddress6DynamicMappingStartMac2edl(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok || d.HasChange("subnet_segment") {
		t, err := expandFirewallAddress6DynamicMappingSubnetSegment2edl(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandFirewallAddress6DynamicMappingTags2edl(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandFirewallAddress6DynamicMappingTemplate2edl(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok || d.HasChange("tenant") {
		t, err := expandFirewallAddress6DynamicMappingTenant2edl(d, v, "tenant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAddress6DynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallAddress6DynamicMappingUuid2edl(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandFirewallAddress6DynamicMappingVisibility2edl(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandFirewallAddress6DynamicMappingWildcard2edl(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	return &obj, nil
}
