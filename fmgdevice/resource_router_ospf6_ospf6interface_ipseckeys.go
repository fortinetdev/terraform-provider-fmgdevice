// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPsec authentication and encryption keys.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6Ospf6InterfaceIpsecKeys() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6Ospf6InterfaceIpsecKeysCreate,
		Read:   resourceRouterOspf6Ospf6InterfaceIpsecKeysRead,
		Update: resourceRouterOspf6Ospf6InterfaceIpsecKeysUpdate,
		Delete: resourceRouterOspf6Ospf6InterfaceIpsecKeysDelete,

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
			"ospf6_interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auth_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"enc_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"spi": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceRouterOspf6Ospf6InterfaceIpsecKeysCreate(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6Ospf6InterfaceIpsecKeys(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6InterfaceIpsecKeys resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6Ospf6InterfaceIpsecKeys(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6InterfaceIpsecKeys resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "spi")))

	return resourceRouterOspf6Ospf6InterfaceIpsecKeysRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceIpsecKeysUpdate(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6Ospf6InterfaceIpsecKeys(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6InterfaceIpsecKeys resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6Ospf6InterfaceIpsecKeys(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6InterfaceIpsecKeys resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "spi")))

	return resourceRouterOspf6Ospf6InterfaceIpsecKeysRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceIpsecKeysDelete(d *schema.ResourceData, m interface{}) error {
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
	ospf6_interface := d.Get("ospf6_interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterOspf6Ospf6InterfaceIpsecKeys(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6Ospf6InterfaceIpsecKeys resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Ospf6InterfaceIpsecKeysRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ospf6_interface := d.Get("ospf6_interface").(string)
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
	if ospf6_interface == "" {
		ospf6_interface = importOptionChecking(m.(*FortiClient).Cfg, "ospf6_interface")
		if ospf6_interface == "" {
			return fmt.Errorf("Parameter ospf6_interface is missing")
		}
		if err = d.Set("ospf6_interface", ospf6_interface); err != nil {
			return fmt.Errorf("Error set params ospf6_interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ospf6_interface"] = ospf6_interface

	o, err := c.ReadRouterOspf6Ospf6InterfaceIpsecKeys(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6InterfaceIpsecKeys resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6Ospf6InterfaceIpsecKeys(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6InterfaceIpsecKeys resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6Ospf6InterfaceIpsecKeys(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("spi", flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi3rdl(o["spi"], d, "spi")); err != nil {
		if vv, ok := fortiAPIPatch(o["spi"], "RouterOspf6Ospf6InterfaceIpsecKeys-Spi"); ok {
			if err = d.Set("spi", vv); err != nil {
				return fmt.Errorf("Error reading spi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spi: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeysFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysSpi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6Ospf6InterfaceIpsecKeys(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_key"); ok || d.HasChange("auth_key") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey3rdl(d, v, "auth_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-key"] = t
		}
	}

	if v, ok := d.GetOk("enc_key"); ok || d.HasChange("enc_key") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey3rdl(d, v, "enc_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-key"] = t
		}
	}

	if v, ok := d.GetOk("spi"); ok || d.HasChange("spi") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecKeysSpi3rdl(d, v, "spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spi"] = t
		}
	}

	return &obj, nil
}
