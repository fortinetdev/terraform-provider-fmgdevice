// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Overlay Controller VPN settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnOcvpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnOcvpnUpdate,
		Read:   resourceVpnOcvpnRead,
		Update: resourceVpnOcvpnUpdate,
		Delete: resourceVpnOcvpnDelete,

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
			"auto_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_shortcut_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"forticlient_access": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_group": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"overlays": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"psksecret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ha_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_allocation_block": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overlays": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"inter_overlay": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"overlay_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnets": &schema.Schema{
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
									"subnet": &schema.Schema{
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan_zone": &schema.Schema{
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
			"wan_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceVpnOcvpnUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnOcvpn(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpn resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnOcvpn(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VpnOcvpn")

	return resourceVpnOcvpnRead(d, m)
}

func resourceVpnOcvpnDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnOcvpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnOcvpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnOcvpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnOcvpn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpn resource from API: %v", err)
	}
	return nil
}

func flattenVpnOcvpnAutoDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnAutoDiscoveryShortcutMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnEap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnEapUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnForticlientAccess(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_groups"
	if _, ok := i["auth-groups"]; ok {
		result["auth_groups"] = flattenVpnOcvpnForticlientAccessAuthGroups(i["auth-groups"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenVpnOcvpnForticlientAccessStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnOcvpnForticlientAccessAuthGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_group"
		if _, ok := i["auth-group"]; ok {
			v := flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup(i["auth-group"], d, pre_append)
			tmp["auth_group"] = fortiAPISubPartPatch(v, "VpnOcvpnForticlientAccess-AuthGroups-AuthGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnOcvpnForticlientAccessAuthGroupsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnOcvpnForticlientAccess-AuthGroups-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlays"
		if _, ok := i["overlays"]; ok {
			v := flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(i["overlays"], d, pre_append)
			tmp["overlays"] = fortiAPISubPartPatch(v, "VpnOcvpnForticlientAccess-AuthGroups-Overlays")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnForticlientAccessAuthGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnForticlientAccessStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnHaAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnIpAllocationBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnMultipath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlays(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_ip"
		if _, ok := i["assign-ip"]; ok {
			v := flattenVpnOcvpnOverlaysAssignIp(i["assign-ip"], d, pre_append)
			tmp["assign_ip"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-AssignIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenVpnOcvpnOverlaysId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inter_overlay"
		if _, ok := i["inter-overlay"]; ok {
			v := flattenVpnOcvpnOverlaysInterOverlay(i["inter-overlay"], d, pre_append)
			tmp["inter_overlay"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-InterOverlay")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_end_ip"
		if _, ok := i["ipv4-end-ip"]; ok {
			v := flattenVpnOcvpnOverlaysIpv4EndIp(i["ipv4-end-ip"], d, pre_append)
			tmp["ipv4_end_ip"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-Ipv4EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_start_ip"
		if _, ok := i["ipv4-start-ip"]; ok {
			v := flattenVpnOcvpnOverlaysIpv4StartIp(i["ipv4-start-ip"], d, pre_append)
			tmp["ipv4_start_ip"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-Ipv4StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenVpnOcvpnOverlaysName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := i["overlay-name"]; ok {
			v := flattenVpnOcvpnOverlaysOverlayName(i["overlay-name"], d, pre_append)
			tmp["overlay_name"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-OverlayName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnets"
		if _, ok := i["subnets"]; ok {
			v := flattenVpnOcvpnOverlaysSubnets(i["subnets"], d, pre_append)
			tmp["subnets"] = fortiAPISubPartPatch(v, "VpnOcvpn-Overlays-Subnets")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnOcvpnOverlaysAssignIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysInterOverlay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysOverlayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnets(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenVpnOcvpnOverlaysSubnetsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsSubnet(i["subnet"], d, pre_append)
			tmp["subnet"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenVpnOcvpnOverlaysSubnetsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "VpnOcvpnOverlays-Subnets-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenVpnOcvpnOverlaysSubnetsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnetsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnOverlaysSubnetsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnSdwanZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnOcvpnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnOcvpnWanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnOcvpn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_discovery", flattenVpnOcvpnAutoDiscovery(o["auto-discovery"], d, "auto_discovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery"], "VpnOcvpn-AutoDiscovery"); ok {
			if err = d.Set("auto_discovery", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery: %v", err)
		}
	}

	if err = d.Set("auto_discovery_shortcut_mode", flattenVpnOcvpnAutoDiscoveryShortcutMode(o["auto-discovery-shortcut-mode"], d, "auto_discovery_shortcut_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-discovery-shortcut-mode"], "VpnOcvpn-AutoDiscoveryShortcutMode"); ok {
			if err = d.Set("auto_discovery_shortcut_mode", vv); err != nil {
				return fmt.Errorf("Error reading auto_discovery_shortcut_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_discovery_shortcut_mode: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnOcvpnEap(o["eap"], d, "eap")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap"], "VpnOcvpn-Eap"); ok {
			if err = d.Set("eap", vv); err != nil {
				return fmt.Errorf("Error reading eap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_users", flattenVpnOcvpnEapUsers(o["eap-users"], d, "eap_users")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-users"], "VpnOcvpn-EapUsers"); ok {
			if err = d.Set("eap_users", vv); err != nil {
				return fmt.Errorf("Error reading eap_users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_users: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("forticlient_access", flattenVpnOcvpnForticlientAccess(o["forticlient-access"], d, "forticlient_access")); err != nil {
			if vv, ok := fortiAPIPatch(o["forticlient-access"], "VpnOcvpn-ForticlientAccess"); ok {
				if err = d.Set("forticlient_access", vv); err != nil {
					return fmt.Errorf("Error reading forticlient_access: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading forticlient_access: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forticlient_access"); ok {
			if err = d.Set("forticlient_access", flattenVpnOcvpnForticlientAccess(o["forticlient-access"], d, "forticlient_access")); err != nil {
				if vv, ok := fortiAPIPatch(o["forticlient-access"], "VpnOcvpn-ForticlientAccess"); ok {
					if err = d.Set("forticlient_access", vv); err != nil {
						return fmt.Errorf("Error reading forticlient_access: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading forticlient_access: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_alias", flattenVpnOcvpnHaAlias(o["ha-alias"], d, "ha_alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-alias"], "VpnOcvpn-HaAlias"); ok {
			if err = d.Set("ha_alias", vv); err != nil {
				return fmt.Errorf("Error reading ha_alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_alias: %v", err)
		}
	}

	if err = d.Set("ip_allocation_block", flattenVpnOcvpnIpAllocationBlock(o["ip-allocation-block"], d, "ip_allocation_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-allocation-block"], "VpnOcvpn-IpAllocationBlock"); ok {
			if err = d.Set("ip_allocation_block", vv); err != nil {
				return fmt.Errorf("Error reading ip_allocation_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_allocation_block: %v", err)
		}
	}

	if err = d.Set("multipath", flattenVpnOcvpnMultipath(o["multipath"], d, "multipath")); err != nil {
		if vv, ok := fortiAPIPatch(o["multipath"], "VpnOcvpn-Multipath"); ok {
			if err = d.Set("multipath", vv); err != nil {
				return fmt.Errorf("Error reading multipath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multipath: %v", err)
		}
	}

	if err = d.Set("nat", flattenVpnOcvpnNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "VpnOcvpn-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("overlays", flattenVpnOcvpnOverlays(o["overlays"], d, "overlays")); err != nil {
			if vv, ok := fortiAPIPatch(o["overlays"], "VpnOcvpn-Overlays"); ok {
				if err = d.Set("overlays", vv); err != nil {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading overlays: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("overlays"); ok {
			if err = d.Set("overlays", flattenVpnOcvpnOverlays(o["overlays"], d, "overlays")); err != nil {
				if vv, ok := fortiAPIPatch(o["overlays"], "VpnOcvpn-Overlays"); ok {
					if err = d.Set("overlays", vv); err != nil {
						return fmt.Errorf("Error reading overlays: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			}
		}
	}

	if err = d.Set("poll_interval", flattenVpnOcvpnPollInterval(o["poll-interval"], d, "poll_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["poll-interval"], "VpnOcvpn-PollInterval"); ok {
			if err = d.Set("poll_interval", vv); err != nil {
				return fmt.Errorf("Error reading poll_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poll_interval: %v", err)
		}
	}

	if err = d.Set("role", flattenVpnOcvpnRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "VpnOcvpn-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenVpnOcvpnSdwan(o["sdwan"], d, "sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan"], "VpnOcvpn-Sdwan"); ok {
			if err = d.Set("sdwan", vv); err != nil {
				return fmt.Errorf("Error reading sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("sdwan_zone", flattenVpnOcvpnSdwanZone(o["sdwan-zone"], d, "sdwan_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan-zone"], "VpnOcvpn-SdwanZone"); ok {
			if err = d.Set("sdwan_zone", vv); err != nil {
				return fmt.Errorf("Error reading sdwan_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan_zone: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnOcvpnStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VpnOcvpn-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("wan_interface", flattenVpnOcvpnWanInterface(o["wan-interface"], d, "wan_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-interface"], "VpnOcvpn-WanInterface"); ok {
			if err = d.Set("wan_interface", vv); err != nil {
				return fmt.Errorf("Error reading wan_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_interface: %v", err)
		}
	}

	return nil
}

func flattenVpnOcvpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnOcvpnAutoDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnAutoDiscoveryShortcutMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnEap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnEapUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnForticlientAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_groups"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandVpnOcvpnForticlientAccessAuthGroups(d, i["auth_groups"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auth-groups"] = t
		}
	}
	pre_append = pre + ".0." + "psksecret"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["psksecret"], _ = expandVpnOcvpnForticlientAccessPsksecret(d, i["psksecret"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandVpnOcvpnForticlientAccessStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-group"], _ = expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup(d, i["auth_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnOcvpnForticlientAccessAuthGroupsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlays"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["overlays"], _ = expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d, i["overlays"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnForticlientAccessPsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnForticlientAccessStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnHaAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnIpAllocationBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnOcvpnMultipath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["assign-ip"], _ = expandVpnOcvpnOverlaysAssignIp(d, i["assign_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandVpnOcvpnOverlaysId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inter_overlay"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["inter-overlay"], _ = expandVpnOcvpnOverlaysInterOverlay(d, i["inter_overlay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-end-ip"], _ = expandVpnOcvpnOverlaysIpv4EndIp(d, i["ipv4_end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-start-ip"], _ = expandVpnOcvpnOverlaysIpv4StartIp(d, i["ipv4_start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandVpnOcvpnOverlaysName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["overlay-name"], _ = expandVpnOcvpnOverlaysOverlayName(d, i["overlay_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandVpnOcvpnOverlaysSubnets(d, i["subnets"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["subnets"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnOverlaysAssignIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysInterOverlay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysOverlayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnOcvpnOverlaysSubnetsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandVpnOcvpnOverlaysSubnetsInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet"], _ = expandVpnOcvpnOverlaysSubnetsSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandVpnOcvpnOverlaysSubnetsType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnOverlaysSubnetsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnOverlaysSubnetsSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandVpnOcvpnOverlaysSubnetsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnSdwanZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnOcvpnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnWanInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnOcvpn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_discovery"); ok || d.HasChange("auto_discovery") {
		t, err := expandVpnOcvpnAutoDiscovery(d, v, "auto_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_shortcut_mode"); ok || d.HasChange("auto_discovery_shortcut_mode") {
		t, err := expandVpnOcvpnAutoDiscoveryShortcutMode(d, v, "auto_discovery_shortcut_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-shortcut-mode"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok || d.HasChange("eap") {
		t, err := expandVpnOcvpnEap(d, v, "eap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_users"); ok || d.HasChange("eap_users") {
		t, err := expandVpnOcvpnEapUsers(d, v, "eap_users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-users"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_access"); ok || d.HasChange("forticlient_access") {
		t, err := expandVpnOcvpnForticlientAccess(d, v, "forticlient_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-access"] = t
		}
	}

	if v, ok := d.GetOk("ha_alias"); ok || d.HasChange("ha_alias") {
		t, err := expandVpnOcvpnHaAlias(d, v, "ha_alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-alias"] = t
		}
	}

	if v, ok := d.GetOk("ip_allocation_block"); ok || d.HasChange("ip_allocation_block") {
		t, err := expandVpnOcvpnIpAllocationBlock(d, v, "ip_allocation_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-allocation-block"] = t
		}
	}

	if v, ok := d.GetOk("multipath"); ok || d.HasChange("multipath") {
		t, err := expandVpnOcvpnMultipath(d, v, "multipath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipath"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandVpnOcvpnNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("overlays"); ok || d.HasChange("overlays") {
		t, err := expandVpnOcvpnOverlays(d, v, "overlays")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlays"] = t
		}
	}

	if v, ok := d.GetOk("poll_interval"); ok || d.HasChange("poll_interval") {
		t, err := expandVpnOcvpnPollInterval(d, v, "poll_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandVpnOcvpnRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok || d.HasChange("sdwan") {
		t, err := expandVpnOcvpnSdwan(d, v, "sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok || d.HasChange("sdwan_zone") {
		t, err := expandVpnOcvpnSdwanZone(d, v, "sdwan_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-zone"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVpnOcvpnStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("wan_interface"); ok || d.HasChange("wan_interface") {
		t, err := expandVpnOcvpnWanInterface(d, v, "wan_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-interface"] = t
		}
	}

	return &obj, nil
}
