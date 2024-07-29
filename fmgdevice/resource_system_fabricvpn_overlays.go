// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Local overlay interfaces table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFabricVpnOverlays() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFabricVpnOverlaysCreate,
		Read:   resourceSystemFabricVpnOverlaysRead,
		Update: resourceSystemFabricVpnOverlaysUpdate,
		Delete: resourceSystemFabricVpnOverlaysDelete,

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
				ForceNew: true,
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
	}
}

func resourceSystemFabricVpnOverlaysCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemFabricVpnOverlays(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemFabricVpnOverlays resource while getting object: %v", err)
	}

	_, err = c.CreateSystemFabricVpnOverlays(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemFabricVpnOverlays resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemFabricVpnOverlaysRead(d, m)
}

func resourceSystemFabricVpnOverlaysUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFabricVpnOverlays(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpnOverlays resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFabricVpnOverlays(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpnOverlays resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemFabricVpnOverlaysRead(d, m)
}

func resourceSystemFabricVpnOverlaysDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFabricVpnOverlays(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFabricVpnOverlays resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFabricVpnOverlaysRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFabricVpnOverlays(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpnOverlays resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFabricVpnOverlays(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpnOverlays resource from API: %v", err)
	}
	return nil
}

func flattenSystemFabricVpnOverlaysBgpNeighbor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNeighborGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNeighborRange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysBgpNetwork2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysIpsecPhase12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysOverlayPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysOverlayTunnelBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysRemoteGw2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysRoutePolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnOverlaysSdwanMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemFabricVpnOverlays(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bgp_neighbor", flattenSystemFabricVpnOverlaysBgpNeighbor2edl(o["bgp-neighbor"], d, "bgp_neighbor")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-neighbor"], "SystemFabricVpnOverlays-BgpNeighbor"); ok {
			if err = d.Set("bgp_neighbor", vv); err != nil {
				return fmt.Errorf("Error reading bgp_neighbor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_neighbor: %v", err)
		}
	}

	if err = d.Set("bgp_neighbor_group", flattenSystemFabricVpnOverlaysBgpNeighborGroup2edl(o["bgp-neighbor-group"], d, "bgp_neighbor_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-neighbor-group"], "SystemFabricVpnOverlays-BgpNeighborGroup"); ok {
			if err = d.Set("bgp_neighbor_group", vv); err != nil {
				return fmt.Errorf("Error reading bgp_neighbor_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_neighbor_group: %v", err)
		}
	}

	if err = d.Set("bgp_neighbor_range", flattenSystemFabricVpnOverlaysBgpNeighborRange2edl(o["bgp-neighbor-range"], d, "bgp_neighbor_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-neighbor-range"], "SystemFabricVpnOverlays-BgpNeighborRange"); ok {
			if err = d.Set("bgp_neighbor_range", vv); err != nil {
				return fmt.Errorf("Error reading bgp_neighbor_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_neighbor_range: %v", err)
		}
	}

	if err = d.Set("bgp_network", flattenSystemFabricVpnOverlaysBgpNetwork2edl(o["bgp-network"], d, "bgp_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-network"], "SystemFabricVpnOverlays-BgpNetwork"); ok {
			if err = d.Set("bgp_network", vv); err != nil {
				return fmt.Errorf("Error reading bgp_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_network: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemFabricVpnOverlaysInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemFabricVpnOverlays-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", flattenSystemFabricVpnOverlaysIpsecPhase12edl(o["ipsec-phase1"], d, "ipsec_phase1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-phase1"], "SystemFabricVpnOverlays-IpsecPhase1"); ok {
			if err = d.Set("ipsec_phase1", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_phase1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemFabricVpnOverlaysName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemFabricVpnOverlays-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overlay_policy", flattenSystemFabricVpnOverlaysOverlayPolicy2edl(o["overlay-policy"], d, "overlay_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["overlay-policy"], "SystemFabricVpnOverlays-OverlayPolicy"); ok {
			if err = d.Set("overlay_policy", vv); err != nil {
				return fmt.Errorf("Error reading overlay_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overlay_policy: %v", err)
		}
	}

	if err = d.Set("overlay_tunnel_block", flattenSystemFabricVpnOverlaysOverlayTunnelBlock2edl(o["overlay-tunnel-block"], d, "overlay_tunnel_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["overlay-tunnel-block"], "SystemFabricVpnOverlays-OverlayTunnelBlock"); ok {
			if err = d.Set("overlay_tunnel_block", vv); err != nil {
				return fmt.Errorf("Error reading overlay_tunnel_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overlay_tunnel_block: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenSystemFabricVpnOverlaysRemoteGw2edl(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "SystemFabricVpnOverlays-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("route_policy", flattenSystemFabricVpnOverlaysRoutePolicy2edl(o["route-policy"], d, "route_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-policy"], "SystemFabricVpnOverlays-RoutePolicy"); ok {
			if err = d.Set("route_policy", vv); err != nil {
				return fmt.Errorf("Error reading route_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_policy: %v", err)
		}
	}

	if err = d.Set("sdwan_member", flattenSystemFabricVpnOverlaysSdwanMember2edl(o["sdwan-member"], d, "sdwan_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan-member"], "SystemFabricVpnOverlays-SdwanMember"); ok {
			if err = d.Set("sdwan_member", vv); err != nil {
				return fmt.Errorf("Error reading sdwan_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan_member: %v", err)
		}
	}

	return nil
}

func flattenSystemFabricVpnOverlaysFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFabricVpnOverlaysBgpNeighbor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNeighborGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNeighborRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysBgpNetwork2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysIpsecPhase12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysOverlayPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysOverlayTunnelBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemFabricVpnOverlaysRemoteGw2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysRoutePolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnOverlaysSdwanMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemFabricVpnOverlays(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bgp_neighbor"); ok || d.HasChange("bgp_neighbor") {
		t, err := expandSystemFabricVpnOverlaysBgpNeighbor2edl(d, v, "bgp_neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-neighbor"] = t
		}
	}

	if v, ok := d.GetOk("bgp_neighbor_group"); ok || d.HasChange("bgp_neighbor_group") {
		t, err := expandSystemFabricVpnOverlaysBgpNeighborGroup2edl(d, v, "bgp_neighbor_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-neighbor-group"] = t
		}
	}

	if v, ok := d.GetOk("bgp_neighbor_range"); ok || d.HasChange("bgp_neighbor_range") {
		t, err := expandSystemFabricVpnOverlaysBgpNeighborRange2edl(d, v, "bgp_neighbor_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-neighbor-range"] = t
		}
	}

	if v, ok := d.GetOk("bgp_network"); ok || d.HasChange("bgp_network") {
		t, err := expandSystemFabricVpnOverlaysBgpNetwork2edl(d, v, "bgp_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-network"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemFabricVpnOverlaysInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1"); ok || d.HasChange("ipsec_phase1") {
		t, err := expandSystemFabricVpnOverlaysIpsecPhase12edl(d, v, "ipsec_phase1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemFabricVpnOverlaysName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overlay_policy"); ok || d.HasChange("overlay_policy") {
		t, err := expandSystemFabricVpnOverlaysOverlayPolicy2edl(d, v, "overlay_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlay-policy"] = t
		}
	}

	if v, ok := d.GetOk("overlay_tunnel_block"); ok || d.HasChange("overlay_tunnel_block") {
		t, err := expandSystemFabricVpnOverlaysOverlayTunnelBlock2edl(d, v, "overlay_tunnel_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overlay-tunnel-block"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandSystemFabricVpnOverlaysRemoteGw2edl(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("route_policy"); ok || d.HasChange("route_policy") {
		t, err := expandSystemFabricVpnOverlaysRoutePolicy2edl(d, v, "route_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-policy"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_member"); ok || d.HasChange("sdwan_member") {
		t, err := expandSystemFabricVpnOverlaysSdwanMember2edl(d, v, "sdwan_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-member"] = t
		}
	}

	return &obj, nil
}
