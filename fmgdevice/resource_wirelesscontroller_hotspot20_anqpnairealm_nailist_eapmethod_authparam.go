// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: EAP auth param.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamCreate,
		Read:   resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead,
		Update: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamDelete,

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
			"anqp_nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nai_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"eap_method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamCreate(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamUpdate(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamDelete(d *schema.ResourceData, m interface{}) error {
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
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
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
	if anqp_nai_realm == "" {
		anqp_nai_realm = importOptionChecking(m.(*FortiClient).Cfg, "anqp_nai_realm")
		if anqp_nai_realm == "" {
			return fmt.Errorf("Parameter anqp_nai_realm is missing")
		}
		if err = d.Set("anqp_nai_realm", anqp_nai_realm); err != nil {
			return fmt.Errorf("Error set params anqp_nai_realm: %v", err)
		}
	}
	if nai_list == "" {
		nai_list = importOptionChecking(m.(*FortiClient).Cfg, "nai_list")
		if nai_list == "" {
			return fmt.Errorf("Parameter nai_list is missing")
		}
		if err = d.Set("nai_list", nai_list); err != nil {
			return fmt.Errorf("Error set params nai_list: %v", err)
		}
	}
	if eap_method == "" {
		eap_method = importOptionChecking(m.(*FortiClient).Cfg, "eap_method")
		if eap_method == "" {
			return fmt.Errorf("Parameter eap_method is missing")
		}
		if err = d.Set("eap_method", eap_method); err != nil {
			return fmt.Errorf("Error set params eap_method: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	o, err := c.ReadWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("val", flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(o["val"], d, "val")); err != nil {
		if vv, ok := fortiAPIPatch(o["val"], "WirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Val"); ok {
			if err = d.Set("val", vv); err != nil {
				return fmt.Errorf("Error reading val: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading val: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("val"); ok || d.HasChange("val") {
		t, err := expandWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(d, v, "val")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["val"] = t
		}
	}

	return &obj, nil
}
