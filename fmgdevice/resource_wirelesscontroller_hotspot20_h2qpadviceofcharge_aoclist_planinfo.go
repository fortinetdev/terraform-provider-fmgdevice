// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Plan info.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCreate,
		Read:   resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoRead,
		Update: resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoDelete,

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
			"h2qp_advice_of_charge": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"aoc_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"currency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"info_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCreate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	aoc_list := d.Get("aoc_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge
	paradict["aoc_list"] = aoc_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoUpdate(d *schema.ResourceData, m interface{}) error {
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
	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	aoc_list := d.Get("aoc_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge
	paradict["aoc_list"] = aoc_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoDelete(d *schema.ResourceData, m interface{}) error {
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
	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	aoc_list := d.Get("aoc_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge
	paradict["aoc_list"] = aoc_list

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	aoc_list := d.Get("aoc_list").(string)
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
	if h2qp_advice_of_charge == "" {
		h2qp_advice_of_charge = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_advice_of_charge")
		if h2qp_advice_of_charge == "" {
			return fmt.Errorf("Parameter h2qp_advice_of_charge is missing")
		}
		if err = d.Set("h2qp_advice_of_charge", h2qp_advice_of_charge); err != nil {
			return fmt.Errorf("Error set params h2qp_advice_of_charge: %v", err)
		}
	}
	if aoc_list == "" {
		aoc_list = importOptionChecking(m.(*FortiClient).Cfg, "aoc_list")
		if aoc_list == "" {
			return fmt.Errorf("Parameter aoc_list is missing")
		}
		if err = d.Set("aoc_list", aoc_list); err != nil {
			return fmt.Errorf("Error set params aoc_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge
	paradict["aoc_list"] = aoc_list

	o, err := c.ReadWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("currency", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency3rdl(o["currency"], d, "currency")); err != nil {
		if vv, ok := fortiAPIPatch(o["currency"], "WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo-Currency"); ok {
			if err = d.Set("currency", vv); err != nil {
				return fmt.Errorf("Error reading currency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading currency: %v", err)
		}
	}

	if err = d.Set("info_file", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile3rdl(o["info-file"], d, "info_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["info-file"], "WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo-InfoFile"); ok {
			if err = d.Set("info_file", vv); err != nil {
				return fmt.Errorf("Error reading info_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading info_file: %v", err)
		}
	}

	if err = d.Set("lang", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang3rdl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("currency"); ok || d.HasChange("currency") {
		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency3rdl(d, v, "currency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["currency"] = t
		}
	}

	if v, ok := d.GetOk("info_file"); ok || d.HasChange("info_file") {
		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile3rdl(d, v, "info_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info-file"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang3rdl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
