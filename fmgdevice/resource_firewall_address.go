// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IPv4 addresses.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddressCreate,
		Read:   resourceFirewallAddressRead,
		Update: resourceFirewallAddressUpdate,
		Delete: resourceFirewallAddressDelete,

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
			"agent_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"allow_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"clearpass_spt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dirty": &schema.Schema{
				Type:     schema.TypeString,
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
						"agent_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"allow_routing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"associated_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cache_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"clearpass_spt": &schema.Schema{
							Type:     schema.TypeString,
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
						"dirty": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"fsso_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"global_object": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"hw_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hw_vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"macaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"node_ip_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obj_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obj_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"organization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pattern_end": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pattern_start": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"policy_group": &schema.Schema{
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
						"sso_attribute_value": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sub_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"subnet_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sw_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_detection_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
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
						"url": &schema.Schema{
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
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"wildcard_fqdn": &schema.Schema{
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
			"fsso_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hw_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"node_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_group": &schema.Schema{
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
			"sso_attribute_value": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sub_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"subnet_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_detection_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"wildcard": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pattern_start": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceFirewallAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddressRead(d, m)
}

func resourceFirewallAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallAddressRead(d, m)
}

func resourceFirewallAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAddress(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddressImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressAgentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressClearpassSpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddressDynamicMappingImageBase64(i["_image-base64"], d, pre_append)
			tmp["_image_base64"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-ImageBase64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenFirewallAddressDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_id"
		if _, ok := i["agent-id"]; ok {
			v := flattenFirewallAddressDynamicMappingAgentId(i["agent-id"], d, pre_append)
			tmp["agent_id"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-AgentId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := i["allow-routing"]; ok {
			v := flattenFirewallAddressDynamicMappingAllowRouting(i["allow-routing"], d, pre_append)
			tmp["allow_routing"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-AllowRouting")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := i["associated-interface"]; ok {
			v := flattenFirewallAddressDynamicMappingAssociatedInterface(i["associated-interface"], d, pre_append)
			tmp["associated_interface"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-AssociatedInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := i["cache-ttl"]; ok {
			v := flattenFirewallAddressDynamicMappingCacheTtl(i["cache-ttl"], d, pre_append)
			tmp["cache_ttl"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-CacheTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clearpass_spt"
		if _, ok := i["clearpass-spt"]; ok {
			v := flattenFirewallAddressDynamicMappingClearpassSpt(i["clearpass-spt"], d, pre_append)
			tmp["clearpass_spt"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-ClearpassSpt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenFirewallAddressDynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallAddressDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {
			v := flattenFirewallAddressDynamicMappingCountry(i["country"], d, pre_append)
			tmp["country"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Country")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dirty"
		if _, ok := i["dirty"]; ok {
			v := flattenFirewallAddressDynamicMappingDirty(i["dirty"], d, pre_append)
			tmp["dirty"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Dirty")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenFirewallAddressDynamicMappingEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := i["end-mac"]; ok {
			v := flattenFirewallAddressDynamicMappingEndMac(i["end-mac"], d, pre_append)
			tmp["end_mac"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-EndMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := i["epg-name"]; ok {
			v := flattenFirewallAddressDynamicMappingEpgName(i["epg-name"], d, pre_append)
			tmp["epg_name"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-EpgName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenFirewallAddressDynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenFirewallAddressDynamicMappingFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenFirewallAddressDynamicMappingFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fsso_group"
		if _, ok := i["fsso-group"]; ok {
			v := flattenFirewallAddressDynamicMappingFssoGroup(i["fsso-group"], d, pre_append)
			tmp["fsso_group"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-FssoGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := i["global-object"]; ok {
			v := flattenFirewallAddressDynamicMappingGlobalObject(i["global-object"], d, pre_append)
			tmp["global_object"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-GlobalObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_model"
		if _, ok := i["hw-model"]; ok {
			v := flattenFirewallAddressDynamicMappingHwModel(i["hw-model"], d, pre_append)
			tmp["hw_model"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-HwModel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := i["hw-vendor"]; ok {
			v := flattenFirewallAddressDynamicMappingHwVendor(i["hw-vendor"], d, pre_append)
			tmp["hw_vendor"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-HwVendor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenFirewallAddressDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := i["macaddr"]; ok {
			v := flattenFirewallAddressDynamicMappingMacaddr(i["macaddr"], d, pre_append)
			tmp["macaddr"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Macaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "node_ip_only"
		if _, ok := i["node-ip-only"]; ok {
			v := flattenFirewallAddressDynamicMappingNodeIpOnly(i["node-ip-only"], d, pre_append)
			tmp["node_ip_only"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-NodeIpOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenFirewallAddressDynamicMappingObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-ObjId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_tag"
		if _, ok := i["obj-tag"]; ok {
			v := flattenFirewallAddressDynamicMappingObjTag(i["obj-tag"], d, pre_append)
			tmp["obj_tag"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-ObjTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_type"
		if _, ok := i["obj-type"]; ok {
			v := flattenFirewallAddressDynamicMappingObjType(i["obj-type"], d, pre_append)
			tmp["obj_type"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-ObjType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "organization"
		if _, ok := i["organization"]; ok {
			v := flattenFirewallAddressDynamicMappingOrganization(i["organization"], d, pre_append)
			tmp["organization"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Organization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {
			v := flattenFirewallAddressDynamicMappingOs(i["os"], d, pre_append)
			tmp["os"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Os")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_end"
		if _, ok := i["pattern-end"]; ok {
			v := flattenFirewallAddressDynamicMappingPatternEnd(i["pattern-end"], d, pre_append)
			tmp["pattern_end"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-PatternEnd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_start"
		if _, ok := i["pattern-start"]; ok {
			v := flattenFirewallAddressDynamicMappingPatternStart(i["pattern-start"], d, pre_append)
			tmp["pattern_start"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-PatternStart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_group"
		if _, ok := i["policy-group"]; ok {
			v := flattenFirewallAddressDynamicMappingPolicyGroup(i["policy-group"], d, pre_append)
			tmp["policy_group"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-PolicyGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenFirewallAddressDynamicMappingRouteTag(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := i["sdn"]; ok {
			v := flattenFirewallAddressDynamicMappingSdn(i["sdn"], d, pre_append)
			tmp["sdn"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Sdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := i["sdn-addr-type"]; ok {
			v := flattenFirewallAddressDynamicMappingSdnAddrType(i["sdn-addr-type"], d, pre_append)
			tmp["sdn_addr_type"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SdnAddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := i["sdn-tag"]; ok {
			v := flattenFirewallAddressDynamicMappingSdnTag(i["sdn-tag"], d, pre_append)
			tmp["sdn_tag"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SdnTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := i["sso-attribute-value"]; ok {
			v := flattenFirewallAddressDynamicMappingSsoAttributeValue(i["sso-attribute-value"], d, pre_append)
			tmp["sso_attribute_value"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SsoAttributeValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenFirewallAddressDynamicMappingStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := i["start-mac"]; ok {
			v := flattenFirewallAddressDynamicMappingStartMac(i["start-mac"], d, pre_append)
			tmp["start_mac"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-StartMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_type"
		if _, ok := i["sub-type"]; ok {
			v := flattenFirewallAddressDynamicMappingSubType(i["sub-type"], d, pre_append)
			tmp["sub_type"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SubType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenFirewallAddressDynamicMappingSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_name"
		if _, ok := i["subnet-name"]; ok {
			v := flattenFirewallAddressDynamicMappingSubnetName(i["subnet-name"], d, pre_append)
			tmp["subnet_name"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SubnetName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_version"
		if _, ok := i["sw-version"]; ok {
			v := flattenFirewallAddressDynamicMappingSwVersion(i["sw-version"], d, pre_append)
			tmp["sw_version"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-SwVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_detection_level"
		if _, ok := i["tag-detection-level"]; ok {
			v := flattenFirewallAddressDynamicMappingTagDetectionLevel(i["tag-detection-level"], d, pre_append)
			tmp["tag_detection_level"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-TagDetectionLevel")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_type"
		if _, ok := i["tag-type"]; ok {
			v := flattenFirewallAddressDynamicMappingTagType(i["tag-type"], d, pre_append)
			tmp["tag_type"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-TagType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_uuid"
		if _, ok := i["tag-uuid"]; ok {
			v := flattenFirewallAddressDynamicMappingTagUuid(i["tag-uuid"], d, pre_append)
			tmp["tag_uuid"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-TagUuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddressDynamicMappingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := i["tenant"]; ok {
			v := flattenFirewallAddressDynamicMappingTenant(i["tenant"], d, pre_append)
			tmp["tenant"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Tenant")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenFirewallAddressDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenFirewallAddressDynamicMappingUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenFirewallAddressDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := i["visibility"]; ok {
			v := flattenFirewallAddressDynamicMappingVisibility(i["visibility"], d, pre_append)
			tmp["visibility"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Visibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := i["wildcard"]; ok {
			v := flattenFirewallAddressDynamicMappingWildcard(i["wildcard"], d, pre_append)
			tmp["wildcard"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-Wildcard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := i["wildcard-fqdn"]; ok {
			v := flattenFirewallAddressDynamicMappingWildcardFqdn(i["wildcard-fqdn"], d, pre_append)
			tmp["wildcard_fqdn"] = fortiAPISubPartPatch(v, "FirewallAddress-DynamicMapping-WildcardFqdn")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddressDynamicMappingImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddressDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddressDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenFirewallAddressDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "FirewallAddressDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddressDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingAgentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingAllowRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingClearpassSpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingEpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingFssoGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingHwModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingNodeIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingObjTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingObjType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingPatternEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingPatternStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingPolicyGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingSdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingSubnetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingSwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingTagDetectionLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingTagUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingTenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressDynamicMappingWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressDynamicMappingWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressEpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressFssoGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressHwModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddressListIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FirewallAddress-List-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := i["net-id"]; ok {
			v := flattenFirewallAddressListNetId(i["net-id"], d, pre_append)
			tmp["net_id"] = fortiAPISubPartPatch(v, "FirewallAddress-List-NetId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenFirewallAddressListObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "FirewallAddress-List-ObjId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddressListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressListNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressListObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressNodeIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressObjTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressObjType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressPolicyGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressSdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressSubnetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressSwVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressTagDetectionLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressTagUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFirewallAddressTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "FirewallAddress-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallAddressTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallAddress-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenFirewallAddressTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "FirewallAddress-Tagging-Tags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressTenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAddressWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressPatternEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddressPatternStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenFirewallAddressImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "FirewallAddress-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("agent_id", flattenFirewallAddressAgentId(o["agent-id"], d, "agent_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["agent-id"], "FirewallAddress-AgentId"); ok {
			if err = d.Set("agent_id", vv); err != nil {
				return fmt.Errorf("Error reading agent_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agent_id: %v", err)
		}
	}

	if err = d.Set("allow_routing", flattenFirewallAddressAllowRouting(o["allow-routing"], d, "allow_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-routing"], "FirewallAddress-AllowRouting"); ok {
			if err = d.Set("allow_routing", vv); err != nil {
				return fmt.Errorf("Error reading allow_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_routing: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenFirewallAddressAssociatedInterface(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "FirewallAddress-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenFirewallAddressCacheTtl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "FirewallAddress-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("clearpass_spt", flattenFirewallAddressClearpassSpt(o["clearpass-spt"], d, "clearpass_spt")); err != nil {
		if vv, ok := fortiAPIPatch(o["clearpass-spt"], "FirewallAddress-ClearpassSpt"); ok {
			if err = d.Set("clearpass_spt", vv); err != nil {
				return fmt.Errorf("Error reading clearpass_spt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clearpass_spt: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallAddressColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallAddress-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallAddressComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallAddress-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenFirewallAddressCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "FirewallAddress-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("dirty", flattenFirewallAddressDirty(o["dirty"], d, "dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["dirty"], "FirewallAddress-Dirty"); ok {
			if err = d.Set("dirty", vv); err != nil {
				return fmt.Errorf("Error reading dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dirty: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenFirewallAddressDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddress-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenFirewallAddressDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "FirewallAddress-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_ip", flattenFirewallAddressEndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "FirewallAddress-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenFirewallAddressEndMac(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "FirewallAddress-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenFirewallAddressEpgName(o["epg-name"], d, "epg_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["epg-name"], "FirewallAddress-EpgName"); ok {
			if err = d.Set("epg_name", vv); err != nil {
				return fmt.Errorf("Error reading epg_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallAddressFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallAddress-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("filter", flattenFirewallAddressFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "FirewallAddress-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallAddressFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "FirewallAddress-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("fsso_group", flattenFirewallAddressFssoGroup(o["fsso-group"], d, "fsso_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-group"], "FirewallAddress-FssoGroup"); ok {
			if err = d.Set("fsso_group", vv); err != nil {
				return fmt.Errorf("Error reading fsso_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_group: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallAddressGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallAddress-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("hw_model", flattenFirewallAddressHwModel(o["hw-model"], d, "hw_model")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-model"], "FirewallAddress-HwModel"); ok {
			if err = d.Set("hw_model", vv); err != nil {
				return fmt.Errorf("Error reading hw_model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_model: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenFirewallAddressHwVendor(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "FirewallAddress-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallAddressInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "FirewallAddress-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenFirewallAddressList(o["list"], d, "list")); err != nil {
			if vv, ok := fortiAPIPatch(o["list"], "FirewallAddress-List"); ok {
				if err = d.Set("list", vv); err != nil {
					return fmt.Errorf("Error reading list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenFirewallAddressList(o["list"], d, "list")); err != nil {
				if vv, ok := fortiAPIPatch(o["list"], "FirewallAddress-List"); ok {
					if err = d.Set("list", vv); err != nil {
						return fmt.Errorf("Error reading list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if err = d.Set("macaddr", flattenFirewallAddressMacaddr(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "FirewallAddress-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("node_ip_only", flattenFirewallAddressNodeIpOnly(o["node-ip-only"], d, "node_ip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["node-ip-only"], "FirewallAddress-NodeIpOnly"); ok {
			if err = d.Set("node_ip_only", vv); err != nil {
				return fmt.Errorf("Error reading node_ip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading node_ip_only: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenFirewallAddressObjId(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "FirewallAddress-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("obj_tag", flattenFirewallAddressObjTag(o["obj-tag"], d, "obj_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-tag"], "FirewallAddress-ObjTag"); ok {
			if err = d.Set("obj_tag", vv); err != nil {
				return fmt.Errorf("Error reading obj_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_tag: %v", err)
		}
	}

	if err = d.Set("obj_type", flattenFirewallAddressObjType(o["obj-type"], d, "obj_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-type"], "FirewallAddress-ObjType"); ok {
			if err = d.Set("obj_type", vv); err != nil {
				return fmt.Errorf("Error reading obj_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_type: %v", err)
		}
	}

	if err = d.Set("organization", flattenFirewallAddressOrganization(o["organization"], d, "organization")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization"], "FirewallAddress-Organization"); ok {
			if err = d.Set("organization", vv); err != nil {
				return fmt.Errorf("Error reading organization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("os", flattenFirewallAddressOs(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "FirewallAddress-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("policy_group", flattenFirewallAddressPolicyGroup(o["policy-group"], d, "policy_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-group"], "FirewallAddress-PolicyGroup"); ok {
			if err = d.Set("policy_group", vv); err != nil {
				return fmt.Errorf("Error reading policy_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_group: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenFirewallAddressRouteTag(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "FirewallAddress-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("sdn", flattenFirewallAddressSdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "FirewallAddress-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenFirewallAddressSdnAddrType(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-addr-type"], "FirewallAddress-SdnAddrType"); ok {
			if err = d.Set("sdn_addr_type", vv); err != nil {
				return fmt.Errorf("Error reading sdn_addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenFirewallAddressSdnTag(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn-tag"], "FirewallAddress-SdnTag"); ok {
			if err = d.Set("sdn_tag", vv); err != nil {
				return fmt.Errorf("Error reading sdn_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenFirewallAddressSsoAttributeValue(o["sso-attribute-value"], d, "sso_attribute_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value"], "FirewallAddress-SsoAttributeValue"); ok {
			if err = d.Set("sso_attribute_value", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenFirewallAddressStartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "FirewallAddress-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenFirewallAddressStartMac(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "FirewallAddress-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if err = d.Set("sub_type", flattenFirewallAddressSubType(o["sub-type"], d, "sub_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-type"], "FirewallAddress-SubType"); ok {
			if err = d.Set("sub_type", vv); err != nil {
				return fmt.Errorf("Error reading sub_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_type: %v", err)
		}
	}

	if err = d.Set("subnet", flattenFirewallAddressSubnet(o["subnet"], d, "subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet"], "FirewallAddress-Subnet"); ok {
			if err = d.Set("subnet", vv); err != nil {
				return fmt.Errorf("Error reading subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("subnet_name", flattenFirewallAddressSubnetName(o["subnet-name"], d, "subnet_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-name"], "FirewallAddress-SubnetName"); ok {
			if err = d.Set("subnet_name", vv); err != nil {
				return fmt.Errorf("Error reading subnet_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_name: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenFirewallAddressSwVersion(o["sw-version"], d, "sw_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-version"], "FirewallAddress-SwVersion"); ok {
			if err = d.Set("sw_version", vv); err != nil {
				return fmt.Errorf("Error reading sw_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("tag_detection_level", flattenFirewallAddressTagDetectionLevel(o["tag-detection-level"], d, "tag_detection_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-detection-level"], "FirewallAddress-TagDetectionLevel"); ok {
			if err = d.Set("tag_detection_level", vv); err != nil {
				return fmt.Errorf("Error reading tag_detection_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_detection_level: %v", err)
		}
	}

	if err = d.Set("tag_type", flattenFirewallAddressTagType(o["tag-type"], d, "tag_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-type"], "FirewallAddress-TagType"); ok {
			if err = d.Set("tag_type", vv); err != nil {
				return fmt.Errorf("Error reading tag_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_type: %v", err)
		}
	}

	if err = d.Set("tag_uuid", flattenFirewallAddressTagUuid(o["tag-uuid"], d, "tag_uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag-uuid"], "FirewallAddress-TagUuid"); ok {
			if err = d.Set("tag_uuid", vv); err != nil {
				return fmt.Errorf("Error reading tag_uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag_uuid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallAddressTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddress-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallAddressTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "FirewallAddress-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("tenant", flattenFirewallAddressTenant(o["tenant"], d, "tenant")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant"], "FirewallAddress-Tenant"); ok {
			if err = d.Set("tenant", vv); err != nil {
				return fmt.Errorf("Error reading tenant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAddressType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddressUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallAddress-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenFirewallAddressWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "FirewallAddress-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenFirewallAddressWildcardFqdn(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "FirewallAddress-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	if err = d.Set("pattern_end", flattenFirewallAddressPatternEnd(o["pattern-end"], d, "pattern_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-end"], "FirewallAddress-PatternEnd"); ok {
			if err = d.Set("pattern_end", vv); err != nil {
				return fmt.Errorf("Error reading pattern_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_end: %v", err)
		}
	}

	if err = d.Set("pattern_start", flattenFirewallAddressPatternStart(o["pattern-start"], d, "pattern_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-start"], "FirewallAddress-PatternStart"); ok {
			if err = d.Set("pattern_start", vv); err != nil {
				return fmt.Errorf("Error reading pattern_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_start: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddressImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressAgentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressClearpassSpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_image-base64"], _ = expandFirewallAddressDynamicMappingImageBase64(d, i["_image_base64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandFirewallAddressDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["agent-id"], _ = expandFirewallAddressDynamicMappingAgentId(d, i["agent_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_routing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allow-routing"], _ = expandFirewallAddressDynamicMappingAllowRouting(d, i["allow_routing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["associated-interface"], _ = expandFirewallAddressDynamicMappingAssociatedInterface(d, i["associated_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cache-ttl"], _ = expandFirewallAddressDynamicMappingCacheTtl(d, i["cache_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clearpass_spt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["clearpass-spt"], _ = expandFirewallAddressDynamicMappingClearpassSpt(d, i["clearpass_spt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandFirewallAddressDynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallAddressDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["country"], _ = expandFirewallAddressDynamicMappingCountry(d, i["country"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dirty"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dirty"], _ = expandFirewallAddressDynamicMappingDirty(d, i["dirty"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandFirewallAddressDynamicMappingEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-mac"], _ = expandFirewallAddressDynamicMappingEndMac(d, i["end_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "epg_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["epg-name"], _ = expandFirewallAddressDynamicMappingEpgName(d, i["epg_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandFirewallAddressDynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandFirewallAddressDynamicMappingFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandFirewallAddressDynamicMappingFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fsso_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fsso-group"], _ = expandFirewallAddressDynamicMappingFssoGroup(d, i["fsso_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["global-object"], _ = expandFirewallAddressDynamicMappingGlobalObject(d, i["global_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_model"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hw-model"], _ = expandFirewallAddressDynamicMappingHwModel(d, i["hw_model"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hw-vendor"], _ = expandFirewallAddressDynamicMappingHwVendor(d, i["hw_vendor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandFirewallAddressDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["macaddr"], _ = expandFirewallAddressDynamicMappingMacaddr(d, i["macaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "node_ip_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["node-ip-only"], _ = expandFirewallAddressDynamicMappingNodeIpOnly(d, i["node_ip_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandFirewallAddressDynamicMappingObjId(d, i["obj_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-tag"], _ = expandFirewallAddressDynamicMappingObjTag(d, i["obj_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-type"], _ = expandFirewallAddressDynamicMappingObjType(d, i["obj_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "organization"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["organization"], _ = expandFirewallAddressDynamicMappingOrganization(d, i["organization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os"], _ = expandFirewallAddressDynamicMappingOs(d, i["os"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_end"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern-end"], _ = expandFirewallAddressDynamicMappingPatternEnd(d, i["pattern_end"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_start"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern-start"], _ = expandFirewallAddressDynamicMappingPatternStart(d, i["pattern_start"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policy-group"], _ = expandFirewallAddressDynamicMappingPolicyGroup(d, i["policy_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandFirewallAddressDynamicMappingRouteTag(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn"], _ = expandFirewallAddressDynamicMappingSdn(d, i["sdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn-addr-type"], _ = expandFirewallAddressDynamicMappingSdnAddrType(d, i["sdn_addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn-tag"], _ = expandFirewallAddressDynamicMappingSdnTag(d, i["sdn_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute-value"], _ = expandFirewallAddressDynamicMappingSsoAttributeValue(d, i["sso_attribute_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandFirewallAddressDynamicMappingStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-mac"], _ = expandFirewallAddressDynamicMappingStartMac(d, i["start_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sub-type"], _ = expandFirewallAddressDynamicMappingSubType(d, i["sub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandFirewallAddressDynamicMappingSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet-name"], _ = expandFirewallAddressDynamicMappingSubnetName(d, i["subnet_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sw-version"], _ = expandFirewallAddressDynamicMappingSwVersion(d, i["sw_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_detection_level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tag-detection-level"], _ = expandFirewallAddressDynamicMappingTagDetectionLevel(d, i["tag_detection_level"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tag-type"], _ = expandFirewallAddressDynamicMappingTagType(d, i["tag_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tag-uuid"], _ = expandFirewallAddressDynamicMappingTagUuid(d, i["tag_uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddressDynamicMappingTags(d, i["tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tenant"], _ = expandFirewallAddressDynamicMappingTenant(d, i["tenant"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandFirewallAddressDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandFirewallAddressDynamicMappingUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandFirewallAddressDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["visibility"], _ = expandFirewallAddressDynamicMappingVisibility(d, i["visibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wildcard"], _ = expandFirewallAddressDynamicMappingWildcard(d, i["wildcard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wildcard-fqdn"], _ = expandFirewallAddressDynamicMappingWildcardFqdn(d, i["wildcard_fqdn"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddressDynamicMappingImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddressDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandFirewallAddressDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddressDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingAgentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingAllowRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingClearpassSpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingEpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingFssoGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingHwModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingNodeIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingObjTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingObjType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingPatternEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingPatternStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingPolicyGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingSdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallAddressDynamicMappingSubnetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingSwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingTagDetectionLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingTagUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressDynamicMappingTenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressDynamicMappingWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallAddressDynamicMappingWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEpgName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFssoGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressHwModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["ip"], _ = expandFirewallAddressListIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["net-id"], _ = expandFirewallAddressListNetId(d, i["net_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandFirewallAddressListObjId(d, i["obj_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddressListIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressListNetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressListObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressNodeIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressPolicyGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressSdnAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSdnTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallAddressSubnetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSwVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagDetectionLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["category"], _ = expandFirewallAddressTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallAddressTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandFirewallAddressTaggingTags(d, i["tags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAddressTenant(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandFirewallAddressWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressPatternEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressPatternStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandFirewallAddressImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("agent_id"); ok || d.HasChange("agent_id") {
		t, err := expandFirewallAddressAgentId(d, v, "agent_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-id"] = t
		}
	}

	if v, ok := d.GetOk("allow_routing"); ok || d.HasChange("allow_routing") {
		t, err := expandFirewallAddressAllowRouting(d, v, "allow_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-routing"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandFirewallAddressAssociatedInterface(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok || d.HasChange("cache_ttl") {
		t, err := expandFirewallAddressCacheTtl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("clearpass_spt"); ok || d.HasChange("clearpass_spt") {
		t, err := expandFirewallAddressClearpassSpt(d, v, "clearpass_spt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clearpass-spt"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallAddressColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallAddressComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandFirewallAddressCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("dirty"); ok || d.HasChange("dirty") {
		t, err := expandFirewallAddressDirty(d, v, "dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dirty"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandFirewallAddressDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandFirewallAddressEndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok || d.HasChange("end_mac") {
		t, err := expandFirewallAddressEndMac(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok || d.HasChange("epg_name") {
		t, err := expandFirewallAddressEpgName(d, v, "epg_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallAddressFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandFirewallAddressFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandFirewallAddressFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fsso_group"); ok || d.HasChange("fsso_group") {
		t, err := expandFirewallAddressFssoGroup(d, v, "fsso_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-group"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallAddressGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("hw_model"); ok || d.HasChange("hw_model") {
		t, err := expandFirewallAddressHwModel(d, v, "hw_model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-model"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandFirewallAddressHwVendor(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallAddressInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok || d.HasChange("list") {
		t, err := expandFirewallAddressList(d, v, "list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandFirewallAddressMacaddr(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("node_ip_only"); ok || d.HasChange("node_ip_only") {
		t, err := expandFirewallAddressNodeIpOnly(d, v, "node_ip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandFirewallAddressObjId(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("obj_tag"); ok || d.HasChange("obj_tag") {
		t, err := expandFirewallAddressObjTag(d, v, "obj_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-tag"] = t
		}
	}

	if v, ok := d.GetOk("obj_type"); ok || d.HasChange("obj_type") {
		t, err := expandFirewallAddressObjType(d, v, "obj_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-type"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok || d.HasChange("organization") {
		t, err := expandFirewallAddressOrganization(d, v, "organization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandFirewallAddressOs(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("policy_group"); ok || d.HasChange("policy_group") {
		t, err := expandFirewallAddressPolicyGroup(d, v, "policy_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-group"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandFirewallAddressRouteTag(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandFirewallAddressSdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok || d.HasChange("sdn_addr_type") {
		t, err := expandFirewallAddressSdnAddrType(d, v, "sdn_addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok || d.HasChange("sdn_tag") {
		t, err := expandFirewallAddressSdnTag(d, v, "sdn_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok || d.HasChange("sso_attribute_value") {
		t, err := expandFirewallAddressSsoAttributeValue(d, v, "sso_attribute_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandFirewallAddressStartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok || d.HasChange("start_mac") {
		t, err := expandFirewallAddressStartMac(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("sub_type"); ok || d.HasChange("sub_type") {
		t, err := expandFirewallAddressSubType(d, v, "sub_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-type"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok || d.HasChange("subnet") {
		t, err := expandFirewallAddressSubnet(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("subnet_name"); ok || d.HasChange("subnet_name") {
		t, err := expandFirewallAddressSubnetName(d, v, "subnet_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-name"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok || d.HasChange("sw_version") {
		t, err := expandFirewallAddressSwVersion(d, v, "sw_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("tag_detection_level"); ok || d.HasChange("tag_detection_level") {
		t, err := expandFirewallAddressTagDetectionLevel(d, v, "tag_detection_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-detection-level"] = t
		}
	}

	if v, ok := d.GetOk("tag_type"); ok || d.HasChange("tag_type") {
		t, err := expandFirewallAddressTagType(d, v, "tag_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-type"] = t
		}
	}

	if v, ok := d.GetOk("tag_uuid"); ok || d.HasChange("tag_uuid") {
		t, err := expandFirewallAddressTagUuid(d, v, "tag_uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-uuid"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallAddressTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok || d.HasChange("tenant") {
		t, err := expandFirewallAddressTenant(d, v, "tenant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallAddressType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallAddressUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandFirewallAddressWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok || d.HasChange("wildcard_fqdn") {
		t, err := expandFirewallAddressWildcardFqdn(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("pattern_end"); ok || d.HasChange("pattern_end") {
		t, err := expandFirewallAddressPatternEnd(d, v, "pattern_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-end"] = t
		}
	}

	if v, ok := d.GetOk("pattern_start"); ok || d.HasChange("pattern_start") {
		t, err := expandFirewallAddressPatternStart(d, v, "pattern_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-start"] = t
		}
	}

	return &obj, nil
}
