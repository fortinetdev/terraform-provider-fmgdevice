// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Setup for self orchestrated fabric auto discovery VPN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFabricVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFabricVpnUpdate,
		Read:   resourceSystemFabricVpnRead,
		Update: resourceSystemFabricVpnUpdate,
		Delete: resourceSystemFabricVpnDelete,

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
			"advertised_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bgp_network": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"firewall_address": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"policies": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
			"bgp_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"branch_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_checks": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"loopback_address_block": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"loopback_advertised_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"loopback_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"overlays": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_neighbor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bgp_neighbor_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bgp_neighbor_range": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"bgp_network": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ipsec_phase1": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"overlay_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"overlay_tunnel_block": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"remote_gw": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sdwan_member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"policy_rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"populated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			"sync_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_role": &schema.Schema{
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

func resourceSystemFabricVpnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemFabricVpn(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpn resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFabricVpn(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFabricVpn")

	return resourceSystemFabricVpnRead(d, m)
}

func resourceSystemFabricVpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemFabricVpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFabricVpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFabricVpnRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemFabricVpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFabricVpn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemFabricVpnAdvertisedSubnets(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access"
		if _, ok := i["access"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsAccess(i["access"], d, pre_append)
			tmp["access"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-Access")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := i["bgp-network"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork(i["bgp-network"], d, pre_append)
			tmp["bgp_network"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-BgpNetwork")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "firewall_address"
		if _, ok := i["firewall-address"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress(i["firewall-address"], d, pre_append)
			tmp["firewall_address"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-FirewallAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policies"
		if _, ok := i["policies"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsPolicies(i["policies"], d, pre_append)
			tmp["policies"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-Policies")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemFabricVpnAdvertisedSubnetsPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemFabricVpn-AdvertisedSubnets-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemFabricVpnAdvertisedSubnetsAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsPolicies(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnBranchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnHealthChecks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnLoopbackAddressBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnLoopbackAdvertisedSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnLoopbackInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlays(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor"
		if _, ok := i["bgp-neighbor"]; ok {
			v := flattenSystemFabricVpnOverlaysBgpNeighbor(i["bgp-neighbor"], d, pre_append)
			tmp["bgp_neighbor"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-BgpNeighbor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_group"
		if _, ok := i["bgp-neighbor-group"]; ok {
			v := flattenSystemFabricVpnOverlaysBgpNeighborGroup(i["bgp-neighbor-group"], d, pre_append)
			tmp["bgp_neighbor_group"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-BgpNeighborGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_range"
		if _, ok := i["bgp-neighbor-range"]; ok {
			v := flattenSystemFabricVpnOverlaysBgpNeighborRange(i["bgp-neighbor-range"], d, pre_append)
			tmp["bgp_neighbor_range"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-BgpNeighborRange")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := i["bgp-network"]; ok {
			v := flattenSystemFabricVpnOverlaysBgpNetwork(i["bgp-network"], d, pre_append)
			tmp["bgp_network"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-BgpNetwork")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemFabricVpnOverlaysInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_phase1"
		if _, ok := i["ipsec-phase1"]; ok {
			v := flattenSystemFabricVpnOverlaysIpsecPhase1(i["ipsec-phase1"], d, pre_append)
			tmp["ipsec_phase1"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-IpsecPhase1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemFabricVpnOverlaysName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_policy"
		if _, ok := i["overlay-policy"]; ok {
			v := flattenSystemFabricVpnOverlaysOverlayPolicy(i["overlay-policy"], d, pre_append)
			tmp["overlay_policy"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-OverlayPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_tunnel_block"
		if _, ok := i["overlay-tunnel-block"]; ok {
			v := flattenSystemFabricVpnOverlaysOverlayTunnelBlock(i["overlay-tunnel-block"], d, pre_append)
			tmp["overlay_tunnel_block"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-OverlayTunnelBlock")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := i["remote-gw"]; ok {
			v := flattenSystemFabricVpnOverlaysRemoteGw(i["remote-gw"], d, pre_append)
			tmp["remote_gw"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-RemoteGw")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_policy"
		if _, ok := i["route-policy"]; ok {
			v := flattenSystemFabricVpnOverlaysRoutePolicy(i["route-policy"], d, pre_append)
			tmp["route_policy"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-RoutePolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdwan_member"
		if _, ok := i["sdwan-member"]; ok {
			v := flattenSystemFabricVpnOverlaysSdwanMember(i["sdwan-member"], d, pre_append)
			tmp["sdwan_member"] = fortiAPISubPartPatch(v, "SystemFabricVpn-Overlays-SdwanMember")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemFabricVpnOverlaysBgpNeighbor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysIpsecPhase1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysOverlayPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysOverlayTunnelBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysRoutePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysSdwanMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnPolicyRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnPopulated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnSdwanZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnSyncMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnVpnRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFabricVpn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("advertised_subnets", flattenSystemFabricVpnAdvertisedSubnets(o["advertised-subnets"], d, "advertised_subnets")); err != nil {
			if vv, ok := fortiAPIPatch(o["advertised-subnets"], "SystemFabricVpn-AdvertisedSubnets"); ok {
				if err = d.Set("advertised_subnets", vv); err != nil {
					return fmt.Errorf("Error reading advertised_subnets: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading advertised_subnets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("advertised_subnets"); ok {
			if err = d.Set("advertised_subnets", flattenSystemFabricVpnAdvertisedSubnets(o["advertised-subnets"], d, "advertised_subnets")); err != nil {
				if vv, ok := fortiAPIPatch(o["advertised-subnets"], "SystemFabricVpn-AdvertisedSubnets"); ok {
					if err = d.Set("advertised_subnets", vv); err != nil {
						return fmt.Errorf("Error reading advertised_subnets: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading advertised_subnets: %v", err)
				}
			}
		}
	}

	if err = d.Set("bgp_as", flattenSystemFabricVpnBgpAs(o["bgp-as"], d, "bgp_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-as"], "SystemFabricVpn-BgpAs"); ok {
			if err = d.Set("bgp_as", vv); err != nil {
				return fmt.Errorf("Error reading bgp_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_as: %v", err)
		}
	}

	if err = d.Set("branch_name", flattenSystemFabricVpnBranchName(o["branch-name"], d, "branch_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["branch-name"], "SystemFabricVpn-BranchName"); ok {
			if err = d.Set("branch_name", vv); err != nil {
				return fmt.Errorf("Error reading branch_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading branch_name: %v", err)
		}
	}

	if err = d.Set("health_checks", flattenSystemFabricVpnHealthChecks(o["health-checks"], d, "health_checks")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-checks"], "SystemFabricVpn-HealthChecks"); ok {
			if err = d.Set("health_checks", vv); err != nil {
				return fmt.Errorf("Error reading health_checks: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_checks: %v", err)
		}
	}

	if err = d.Set("loopback_address_block", flattenSystemFabricVpnLoopbackAddressBlock(o["loopback-address-block"], d, "loopback_address_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["loopback-address-block"], "SystemFabricVpn-LoopbackAddressBlock"); ok {
			if err = d.Set("loopback_address_block", vv); err != nil {
				return fmt.Errorf("Error reading loopback_address_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loopback_address_block: %v", err)
		}
	}

	if err = d.Set("loopback_advertised_subnet", flattenSystemFabricVpnLoopbackAdvertisedSubnet(o["loopback-advertised-subnet"], d, "loopback_advertised_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["loopback-advertised-subnet"], "SystemFabricVpn-LoopbackAdvertisedSubnet"); ok {
			if err = d.Set("loopback_advertised_subnet", vv); err != nil {
				return fmt.Errorf("Error reading loopback_advertised_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loopback_advertised_subnet: %v", err)
		}
	}

	if err = d.Set("loopback_interface", flattenSystemFabricVpnLoopbackInterface(o["loopback-interface"], d, "loopback_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["loopback-interface"], "SystemFabricVpn-LoopbackInterface"); ok {
			if err = d.Set("loopback_interface", vv); err != nil {
				return fmt.Errorf("Error reading loopback_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loopback_interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("overlays", flattenSystemFabricVpnOverlays(o["overlays"], d, "overlays")); err != nil {
			if vv, ok := fortiAPIPatch(o["overlays"], "SystemFabricVpn-Overlays"); ok {
				if err = d.Set("overlays", vv); err != nil {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading overlays: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("overlays"); ok {
			if err = d.Set("overlays", flattenSystemFabricVpnOverlays(o["overlays"], d, "overlays")); err != nil {
				if vv, ok := fortiAPIPatch(o["overlays"], "SystemFabricVpn-Overlays"); ok {
					if err = d.Set("overlays", vv); err != nil {
						return fmt.Errorf("Error reading overlays: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			}
		}
	}

	if err = d.Set("policy_rule", flattenSystemFabricVpnPolicyRule(o["policy-rule"], d, "policy_rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-rule"], "SystemFabricVpn-PolicyRule"); ok {
			if err = d.Set("policy_rule", vv); err != nil {
				return fmt.Errorf("Error reading policy_rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_rule: %v", err)
		}
	}

	if err = d.Set("populated", flattenSystemFabricVpnPopulated(o["populated"], d, "populated")); err != nil {
		if vv, ok := fortiAPIPatch(o["populated"], "SystemFabricVpn-Populated"); ok {
			if err = d.Set("populated", vv); err != nil {
				return fmt.Errorf("Error reading populated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading populated: %v", err)
		}
	}

	if err = d.Set("sdwan_zone", flattenSystemFabricVpnSdwanZone(o["sdwan-zone"], d, "sdwan_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan-zone"], "SystemFabricVpn-SdwanZone"); ok {
			if err = d.Set("sdwan_zone", vv); err != nil {
				return fmt.Errorf("Error reading sdwan_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan_zone: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFabricVpnStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFabricVpn-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sync_mode", flattenSystemFabricVpnSyncMode(o["sync-mode"], d, "sync_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-mode"], "SystemFabricVpn-SyncMode"); ok {
			if err = d.Set("sync_mode", vv); err != nil {
				return fmt.Errorf("Error reading sync_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_mode: %v", err)
		}
	}

	if err = d.Set("vpn_role", flattenSystemFabricVpnVpnRole(o["vpn-role"], d, "vpn_role")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-role"], "SystemFabricVpn-VpnRole"); ok {
			if err = d.Set("vpn_role", vv); err != nil {
				return fmt.Errorf("Error reading vpn_role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_role: %v", err)
		}
	}

	return nil
}

func flattenSystemFabricVpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFabricVpnAdvertisedSubnets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access"], _ = expandSystemFabricVpnAdvertisedSubnetsAccess(d, i["access"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bgp-network"], _ = expandSystemFabricVpnAdvertisedSubnetsBgpNetwork(d, i["bgp_network"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "firewall_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["firewall-address"], _ = expandSystemFabricVpnAdvertisedSubnetsFirewallAddress(d, i["firewall_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemFabricVpnAdvertisedSubnetsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policies"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["policies"], _ = expandSystemFabricVpnAdvertisedSubnetsPolicies(d, i["policies"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemFabricVpnAdvertisedSubnetsPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemFabricVpnAdvertisedSubnetsAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsBgpNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsFirewallAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsPolicies(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemFabricVpnBgpAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnBranchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnHealthChecks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnLoopbackAddressBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemFabricVpnLoopbackAdvertisedSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnLoopbackInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bgp-neighbor"], _ = expandSystemFabricVpnOverlaysBgpNeighbor(d, i["bgp_neighbor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bgp-neighbor-group"], _ = expandSystemFabricVpnOverlaysBgpNeighborGroup(d, i["bgp_neighbor_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bgp-neighbor-range"], _ = expandSystemFabricVpnOverlaysBgpNeighborRange(d, i["bgp_neighbor_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bgp-network"], _ = expandSystemFabricVpnOverlaysBgpNetwork(d, i["bgp_network"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemFabricVpnOverlaysInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_phase1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipsec-phase1"], _ = expandSystemFabricVpnOverlaysIpsecPhase1(d, i["ipsec_phase1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemFabricVpnOverlaysName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["overlay-policy"], _ = expandSystemFabricVpnOverlaysOverlayPolicy(d, i["overlay_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_tunnel_block"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["overlay-tunnel-block"], _ = expandSystemFabricVpnOverlaysOverlayTunnelBlock(d, i["overlay_tunnel_block"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-gw"], _ = expandSystemFabricVpnOverlaysRemoteGw(d, i["remote_gw"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-policy"], _ = expandSystemFabricVpnOverlaysRoutePolicy(d, i["route_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdwan_member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdwan-member"], _ = expandSystemFabricVpnOverlaysSdwanMember(d, i["sdwan_member"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemFabricVpnOverlaysBgpNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysIpsecPhase1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysOverlayPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysOverlayTunnelBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemFabricVpnOverlaysRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysRoutePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysSdwanMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnPolicyRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnPopulated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnPsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnSdwanZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnSyncMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnVpnRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFabricVpn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertised_subnets"); ok || d.HasChange("advertised_subnets") {
		t, err := expandSystemFabricVpnAdvertisedSubnets(d, v, "advertised_subnets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertised-subnets"] = t
		}
	}

	if v, ok := d.GetOk("bgp_as"); ok || d.HasChange("bgp_as") {
		t, err := expandSystemFabricVpnBgpAs(d, v, "bgp_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-as"] = t
		}
	}

	if v, ok := d.GetOk("branch_name"); ok || d.HasChange("branch_name") {
		t, err := expandSystemFabricVpnBranchName(d, v, "branch_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["branch-name"] = t
		}
	}

	if v, ok := d.GetOk("health_checks"); ok || d.HasChange("health_checks") {
		t, err := expandSystemFabricVpnHealthChecks(d, v, "health_checks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-checks"] = t
		}
	}

	if v, ok := d.GetOk("loopback_address_block"); ok || d.HasChange("loopback_address_block") {
		t, err := expandSystemFabricVpnLoopbackAddressBlock(d, v, "loopback_address_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-address-block"] = t
		}
	}

	if v, ok := d.GetOk("loopback_advertised_subnet"); ok || d.HasChange("loopback_advertised_subnet") {
		t, err := expandSystemFabricVpnLoopbackAdvertisedSubnet(d, v, "loopback_advertised_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-advertised-subnet"] = t
		}
	}

	if v, ok := d.GetOk("loopback_interface"); ok || d.HasChange("loopback_interface") {
		t, err := expandSystemFabricVpnLoopbackInterface(d, v, "loopback_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-interface"] = t
		}
	}

	if v, ok := d.GetOk("overlays"); ok || d.HasChange("overlays") {
		t, err := expandSystemFabricVpnOverlays(d, v, "overlays")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlays"] = t
		}
	}

	if v, ok := d.GetOk("policy_rule"); ok || d.HasChange("policy_rule") {
		t, err := expandSystemFabricVpnPolicyRule(d, v, "policy_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-rule"] = t
		}
	}

	if v, ok := d.GetOk("populated"); ok || d.HasChange("populated") {
		t, err := expandSystemFabricVpnPopulated(d, v, "populated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["populated"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok || d.HasChange("psksecret") {
		t, err := expandSystemFabricVpnPsksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok || d.HasChange("sdwan_zone") {
		t, err := expandSystemFabricVpnSdwanZone(d, v, "sdwan_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-zone"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFabricVpnStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sync_mode"); ok || d.HasChange("sync_mode") {
		t, err := expandSystemFabricVpnSyncMode(d, v, "sync_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-mode"] = t
		}
	}

	if v, ok := d.GetOk("vpn_role"); ok || d.HasChange("vpn_role") {
		t, err := expandSystemFabricVpnVpnRole(d, v, "vpn_role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-role"] = t
		}
	}

	return &obj, nil
}
