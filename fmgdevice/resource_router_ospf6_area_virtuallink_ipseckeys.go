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

func resourceRouterOspf6AreaVirtualLinkIpsecKeys() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6AreaVirtualLinkIpsecKeysCreate,
		Read:   resourceRouterOspf6AreaVirtualLinkIpsecKeysRead,
		Update: resourceRouterOspf6AreaVirtualLinkIpsecKeysUpdate,
		Delete: resourceRouterOspf6AreaVirtualLinkIpsecKeysDelete,

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
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_link": &schema.Schema{
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

func resourceRouterOspf6AreaVirtualLinkIpsecKeysCreate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaVirtualLinkIpsecKeys(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaVirtualLinkIpsecKeys resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6AreaVirtualLinkIpsecKeys(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6AreaVirtualLinkIpsecKeys resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "spi")))

	return resourceRouterOspf6AreaVirtualLinkIpsecKeysRead(d, m)
}

func resourceRouterOspf6AreaVirtualLinkIpsecKeysUpdate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterOspf6AreaVirtualLinkIpsecKeys(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaVirtualLinkIpsecKeys resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6AreaVirtualLinkIpsecKeys(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6AreaVirtualLinkIpsecKeys resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "spi")))

	return resourceRouterOspf6AreaVirtualLinkIpsecKeysRead(d, m)
}

func resourceRouterOspf6AreaVirtualLinkIpsecKeysDelete(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterOspf6AreaVirtualLinkIpsecKeys(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6AreaVirtualLinkIpsecKeys resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6AreaVirtualLinkIpsecKeysRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
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
	if area == "" {
		area = importOptionChecking(m.(*FortiClient).Cfg, "area")
		if area == "" {
			return fmt.Errorf("Parameter area is missing")
		}
		if err = d.Set("area", area); err != nil {
			return fmt.Errorf("Error set params area: %v", err)
		}
	}
	if virtual_link == "" {
		virtual_link = importOptionChecking(m.(*FortiClient).Cfg, "virtual_link")
		if virtual_link == "" {
			return fmt.Errorf("Parameter virtual_link is missing")
		}
		if err = d.Set("virtual_link", virtual_link); err != nil {
			return fmt.Errorf("Error set params virtual_link: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	o, err := c.ReadRouterOspf6AreaVirtualLinkIpsecKeys(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaVirtualLinkIpsecKeys resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6AreaVirtualLinkIpsecKeys(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6AreaVirtualLinkIpsecKeys resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6AreaVirtualLinkIpsecKeys(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("spi", flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi4thl(o["spi"], d, "spi")); err != nil {
		if vv, ok := fortiAPIPatch(o["spi"], "RouterOspf6AreaVirtualLinkIpsecKeys-Spi"); ok {
			if err = d.Set("spi", vv); err != nil {
				return fmt.Errorf("Error reading spi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spi: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysSpi4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6AreaVirtualLinkIpsecKeys(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_key"); ok || d.HasChange("auth_key") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey4thl(d, v, "auth_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-key"] = t
		}
	}

	if v, ok := d.GetOk("enc_key"); ok || d.HasChange("enc_key") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey4thl(d, v, "enc_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-key"] = t
		}
	}

	if v, ok := d.GetOk("spi"); ok || d.HasChange("spi") {
		t, err := expandRouterOspf6AreaVirtualLinkIpsecKeysSpi4thl(d, v, "spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spi"] = t
		}
	}

	return &obj, nil
}
