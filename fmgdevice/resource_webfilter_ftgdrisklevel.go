// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter risk level.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFtgdRiskLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdRiskLevelCreate,
		Read:   resourceWebfilterFtgdRiskLevelRead,
		Update: resourceWebfilterFtgdRiskLevelUpdate,
		Delete: resourceWebfilterFtgdRiskLevelDelete,

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
			"high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceWebfilterFtgdRiskLevelCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebfilterFtgdRiskLevel(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdRiskLevel resource while getting object: %v", err)
	}

	_, err = c.CreateWebfilterFtgdRiskLevel(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdRiskLevel resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterFtgdRiskLevelRead(d, m)
}

func resourceWebfilterFtgdRiskLevelUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWebfilterFtgdRiskLevel(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdRiskLevel resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterFtgdRiskLevel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdRiskLevel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterFtgdRiskLevelRead(d, m)
}

func resourceWebfilterFtgdRiskLevelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterFtgdRiskLevel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdRiskLevel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdRiskLevelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterFtgdRiskLevel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdRiskLevel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdRiskLevel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdRiskLevel resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdRiskLevelHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdRiskLevelLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdRiskLevelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterFtgdRiskLevel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("high", flattenWebfilterFtgdRiskLevelHigh(o["high"], d, "high")); err != nil {
		if vv, ok := fortiAPIPatch(o["high"], "WebfilterFtgdRiskLevel-High"); ok {
			if err = d.Set("high", vv); err != nil {
				return fmt.Errorf("Error reading high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("low", flattenWebfilterFtgdRiskLevelLow(o["low"], d, "low")); err != nil {
		if vv, ok := fortiAPIPatch(o["low"], "WebfilterFtgdRiskLevel-Low"); ok {
			if err = d.Set("low", vv); err != nil {
				return fmt.Errorf("Error reading low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterFtgdRiskLevelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterFtgdRiskLevel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdRiskLevelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterFtgdRiskLevelHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdRiskLevelLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdRiskLevelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdRiskLevel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("high"); ok || d.HasChange("high") {
		t, err := expandWebfilterFtgdRiskLevelHigh(d, v, "high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	}

	if v, ok := d.GetOk("low"); ok || d.HasChange("low") {
		t, err := expandWebfilterFtgdRiskLevelLow(d, v, "low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterFtgdRiskLevelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
