// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IP address summary configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfSummaryAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfSummaryAddressCreate,
		Read:   resourceRouterOspfSummaryAddressRead,
		Update: resourceRouterOspfSummaryAddressUpdate,
		Delete: resourceRouterOspfSummaryAddressDelete,

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
			"advertise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceRouterOspfSummaryAddressCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfSummaryAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfSummaryAddress resource while getting object: %v", err)
	}

	v, err := c.CreateRouterOspfSummaryAddress(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfSummaryAddress resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceRouterOspfSummaryAddressRead(d, m)
		} else {
			return fmt.Errorf("Error creating RouterOspfSummaryAddress resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfSummaryAddressRead(d, m)
}

func resourceRouterOspfSummaryAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfSummaryAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfSummaryAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfSummaryAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfSummaryAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfSummaryAddressRead(d, m)
}

func resourceRouterOspfSummaryAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfSummaryAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfSummaryAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfSummaryAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfSummaryAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfSummaryAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfSummaryAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfSummaryAddress resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfSummaryAddressAdvertise2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfSummaryAddressTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfSummaryAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advertise", flattenRouterOspfSummaryAddressAdvertise2edl(o["advertise"], d, "advertise")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertise"], "RouterOspfSummaryAddress-Advertise"); ok {
			if err = d.Set("advertise", vv); err != nil {
				return fmt.Errorf("Error reading advertise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertise: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterOspfSummaryAddressId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspfSummaryAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterOspfSummaryAddressPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterOspfSummaryAddress-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("tag", flattenRouterOspfSummaryAddressTag2edl(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "RouterOspfSummaryAddress-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfSummaryAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfSummaryAddressAdvertise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterOspfSummaryAddressTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfSummaryAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertise"); ok || d.HasChange("advertise") {
		t, err := expandRouterOspfSummaryAddressAdvertise2edl(d, v, "advertise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertise"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspfSummaryAddressId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterOspfSummaryAddressPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandRouterOspfSummaryAddressTag2edl(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	return &obj, nil
}
