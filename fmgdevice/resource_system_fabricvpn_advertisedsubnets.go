// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Local advertised subnets.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFabricVpnAdvertisedSubnets() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFabricVpnAdvertisedSubnetsCreate,
		Read:   resourceSystemFabricVpnAdvertisedSubnetsRead,
		Update: resourceSystemFabricVpnAdvertisedSubnetsUpdate,
		Delete: resourceSystemFabricVpnAdvertisedSubnetsDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceSystemFabricVpnAdvertisedSubnetsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemFabricVpnAdvertisedSubnets(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemFabricVpnAdvertisedSubnets resource while getting object: %v", err)
	}

	_, err = c.CreateSystemFabricVpnAdvertisedSubnets(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemFabricVpnAdvertisedSubnets resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemFabricVpnAdvertisedSubnetsRead(d, m)
}

func resourceSystemFabricVpnAdvertisedSubnetsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFabricVpnAdvertisedSubnets(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpnAdvertisedSubnets resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFabricVpnAdvertisedSubnets(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpnAdvertisedSubnets resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemFabricVpnAdvertisedSubnetsRead(d, m)
}

func resourceSystemFabricVpnAdvertisedSubnetsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemFabricVpnAdvertisedSubnets(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFabricVpnAdvertisedSubnets resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFabricVpnAdvertisedSubnetsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFabricVpnAdvertisedSubnets(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpnAdvertisedSubnets resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFabricVpnAdvertisedSubnets(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpnAdvertisedSubnets resource from API: %v", err)
	}
	return nil
}

func flattenSystemFabricVpnAdvertisedSubnetsAccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsPolicies2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemFabricVpnAdvertisedSubnets(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access", flattenSystemFabricVpnAdvertisedSubnetsAccess2edl(o["access"], d, "access")); err != nil {
		if vv, ok := fortiAPIPatch(o["access"], "SystemFabricVpnAdvertisedSubnets-Access"); ok {
			if err = d.Set("access", vv); err != nil {
				return fmt.Errorf("Error reading access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access: %v", err)
		}
	}

	if err = d.Set("bgp_network", flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork2edl(o["bgp-network"], d, "bgp_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-network"], "SystemFabricVpnAdvertisedSubnets-BgpNetwork"); ok {
			if err = d.Set("bgp_network", vv); err != nil {
				return fmt.Errorf("Error reading bgp_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_network: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress2edl(o["firewall-address"], d, "firewall_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-address"], "SystemFabricVpnAdvertisedSubnets-FirewallAddress"); ok {
			if err = d.Set("firewall_address", vv); err != nil {
				return fmt.Errorf("Error reading firewall_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemFabricVpnAdvertisedSubnetsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemFabricVpnAdvertisedSubnets-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("policies", flattenSystemFabricVpnAdvertisedSubnetsPolicies2edl(o["policies"], d, "policies")); err != nil {
		if vv, ok := fortiAPIPatch(o["policies"], "SystemFabricVpnAdvertisedSubnets-Policies"); ok {
			if err = d.Set("policies", vv); err != nil {
				return fmt.Errorf("Error reading policies: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policies: %v", err)
		}
	}

	if err = d.Set("prefix", flattenSystemFabricVpnAdvertisedSubnetsPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "SystemFabricVpnAdvertisedSubnets-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenSystemFabricVpnAdvertisedSubnetsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFabricVpnAdvertisedSubnetsAccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsBgpNetwork2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsFirewallAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsPolicies2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemFabricVpnAdvertisedSubnetsPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSystemFabricVpnAdvertisedSubnets(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access"); ok || d.HasChange("access") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsAccess2edl(d, v, "access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access"] = t
		}
	}

	if v, ok := d.GetOk("bgp_network"); ok || d.HasChange("bgp_network") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsBgpNetwork2edl(d, v, "bgp_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-network"] = t
		}
	}

	if v, ok := d.GetOk("firewall_address"); ok || d.HasChange("firewall_address") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsFirewallAddress2edl(d, v, "firewall_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("policies"); ok || d.HasChange("policies") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsPolicies2edl(d, v, "policies")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policies"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandSystemFabricVpnAdvertisedSubnetsPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
