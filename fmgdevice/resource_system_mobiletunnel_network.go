// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NEMO network configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMobileTunnelNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMobileTunnelNetworkCreate,
		Read:   resourceSystemMobileTunnelNetworkRead,
		Update: resourceSystemMobileTunnelNetworkUpdate,
		Delete: resourceSystemMobileTunnelNetworkDelete,

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
			"mobile_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
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

func resourceSystemMobileTunnelNetworkCreate(d *schema.ResourceData, m interface{}) error {
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
	mobile_tunnel := d.Get("mobile_tunnel").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mobile_tunnel"] = mobile_tunnel

	obj, err := getObjectSystemMobileTunnelNetwork(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnelNetwork resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMobileTunnelNetwork(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnelNetwork resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemMobileTunnelNetworkRead(d, m)
}

func resourceSystemMobileTunnelNetworkUpdate(d *schema.ResourceData, m interface{}) error {
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
	mobile_tunnel := d.Get("mobile_tunnel").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mobile_tunnel"] = mobile_tunnel

	obj, err := getObjectSystemMobileTunnelNetwork(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnelNetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMobileTunnelNetwork(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnelNetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemMobileTunnelNetworkRead(d, m)
}

func resourceSystemMobileTunnelNetworkDelete(d *schema.ResourceData, m interface{}) error {
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
	mobile_tunnel := d.Get("mobile_tunnel").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mobile_tunnel"] = mobile_tunnel

	err = c.DeleteSystemMobileTunnelNetwork(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMobileTunnelNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMobileTunnelNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	mobile_tunnel := d.Get("mobile_tunnel").(string)
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
	if mobile_tunnel == "" {
		mobile_tunnel = importOptionChecking(m.(*FortiClient).Cfg, "mobile_tunnel")
		if mobile_tunnel == "" {
			return fmt.Errorf("Parameter mobile_tunnel is missing")
		}
		if err = d.Set("mobile_tunnel", mobile_tunnel); err != nil {
			return fmt.Errorf("Error set params mobile_tunnel: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["mobile_tunnel"] = mobile_tunnel

	o, err := c.ReadSystemMobileTunnelNetwork(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnelNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMobileTunnelNetwork(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnelNetwork resource from API: %v", err)
	}
	return nil
}

func flattenSystemMobileTunnelNetworkId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetworkInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMobileTunnelNetworkPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemMobileTunnelNetwork(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemMobileTunnelNetworkId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemMobileTunnelNetwork-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemMobileTunnelNetworkInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemMobileTunnelNetwork-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("prefix", flattenSystemMobileTunnelNetworkPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "SystemMobileTunnelNetwork-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenSystemMobileTunnelNetworkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMobileTunnelNetworkId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetworkInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMobileTunnelNetworkPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSystemMobileTunnelNetwork(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemMobileTunnelNetworkId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemMobileTunnelNetworkInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandSystemMobileTunnelNetworkPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
