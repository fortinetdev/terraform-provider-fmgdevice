// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate connector LAN extension configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerFortigateProfileLanExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerFortigateProfileLanExtensionUpdate,
		Read:   resourceExtensionControllerFortigateProfileLanExtensionRead,
		Update: resourceExtensionControllerFortigateProfileLanExtensionUpdate,
		Delete: resourceExtensionControllerFortigateProfileLanExtensionDelete,

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
			"fortigate_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backhaul_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"backhaul_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceExtensionControllerFortigateProfileLanExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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
	fortigate_profile := d.Get("fortigate_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortigate_profile"] = fortigate_profile

	obj, err := getObjectExtensionControllerFortigateProfileLanExtension(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfileLanExtension resource while getting object: %v", err)
	}

	_, err = c.UpdateExtensionControllerFortigateProfileLanExtension(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfileLanExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExtensionControllerFortigateProfileLanExtension")

	return resourceExtensionControllerFortigateProfileLanExtensionRead(d, m)
}

func resourceExtensionControllerFortigateProfileLanExtensionDelete(d *schema.ResourceData, m interface{}) error {
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
	fortigate_profile := d.Get("fortigate_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortigate_profile"] = fortigate_profile

	err = c.DeleteExtensionControllerFortigateProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerFortigateProfileLanExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerFortigateProfileLanExtensionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	fortigate_profile := d.Get("fortigate_profile").(string)
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
	if fortigate_profile == "" {
		fortigate_profile = importOptionChecking(m.(*FortiClient).Cfg, "fortigate_profile")
		if fortigate_profile == "" {
			return fmt.Errorf("Parameter fortigate_profile is missing")
		}
		if err = d.Set("fortigate_profile", fortigate_profile); err != nil {
			return fmt.Errorf("Error set params fortigate_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortigate_profile"] = fortigate_profile

	o, err := c.ReadExtensionControllerFortigateProfileLanExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfileLanExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerFortigateProfileLanExtension(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfileLanExtension resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerFortigateProfileLanExtension(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("backhaul_interface", flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface2edl(o["backhaul-interface"], d, "backhaul_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-interface"], "ExtensionControllerFortigateProfileLanExtension-BackhaulInterface"); ok {
			if err = d.Set("backhaul_interface", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_interface: %v", err)
		}
	}

	if err = d.Set("backhaul_ip", flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp2edl(o["backhaul-ip"], d, "backhaul_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["backhaul-ip"], "ExtensionControllerFortigateProfileLanExtension-BackhaulIp"); ok {
			if err = d.Set("backhaul_ip", vv); err != nil {
				return fmt.Errorf("Error reading backhaul_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backhaul_ip: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel", flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel2edl(o["ipsec-tunnel"], d, "ipsec_tunnel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-tunnel"], "ExtensionControllerFortigateProfileLanExtension-IpsecTunnel"); ok {
			if err = d.Set("ipsec_tunnel", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_tunnel: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerFortigateProfileLanExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerFortigateProfileLanExtension(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backhaul_interface"); ok || d.HasChange("backhaul_interface") {
		t, err := expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface2edl(d, v, "backhaul_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-interface"] = t
		}
	}

	if v, ok := d.GetOk("backhaul_ip"); ok || d.HasChange("backhaul_ip") {
		t, err := expandExtensionControllerFortigateProfileLanExtensionBackhaulIp2edl(d, v, "backhaul_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backhaul-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel"); ok || d.HasChange("ipsec_tunnel") {
		t, err := expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel2edl(d, v, "ipsec_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel"] = t
		}
	}

	return &obj, nil
}
