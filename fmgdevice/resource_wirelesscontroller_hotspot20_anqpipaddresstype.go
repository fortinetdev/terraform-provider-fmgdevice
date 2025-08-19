// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP address type availability.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpIpAddressType() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpIpAddressTypeCreate,
		Read:   resourceWirelessControllerHotspot20AnqpIpAddressTypeRead,
		Update: resourceWirelessControllerHotspot20AnqpIpAddressTypeUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpIpAddressTypeDelete,

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
			"ipv4_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20AnqpIpAddressType(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpIpAddressType(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWirelessControllerHotspot20AnqpIpAddressType(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpIpAddressType(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerHotspot20AnqpIpAddressType(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20AnqpIpAddressType(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpIpAddressType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpIpAddressType resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ipv4_address_type", flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(o["ipv4-address-type"], d, "ipv4_address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-address-type"], "WirelessControllerHotspot20AnqpIpAddressType-Ipv4AddressType"); ok {
			if err = d.Set("ipv4_address_type", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_address_type: %v", err)
		}
	}

	if err = d.Set("ipv6_address_type", flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(o["ipv6-address-type"], d, "ipv6_address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-address-type"], "WirelessControllerHotspot20AnqpIpAddressType-Ipv6AddressType"); ok {
			if err = d.Set("ipv6_address_type", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_address_type: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpIpAddressTypeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20AnqpIpAddressType-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ipv4_address_type"); ok || d.HasChange("ipv4_address_type") {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d, v, "ipv4_address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-address-type"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_address_type"); ok || d.HasChange("ipv6_address_type") {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d, v, "ipv6_address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-address-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
