// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigDefault() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigDefaultUpdate,
		Read:   resourceSwitchControllerAutoConfigDefaultRead,
		Update: resourceSwitchControllerAutoConfigDefaultUpdate,
		Delete: resourceSwitchControllerAutoConfigDefaultDelete,

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
			"fgt_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icl_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"isl_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigDefaultUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerAutoConfigDefault(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAutoConfigDefault(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerAutoConfigDefault")

	return resourceSwitchControllerAutoConfigDefaultRead(d, m)
}

func resourceSwitchControllerAutoConfigDefaultDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerAutoConfigDefault(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigDefault resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigDefaultRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerAutoConfigDefault(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigDefault(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigDefaultFgtPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenSwitchControllerAutoConfigDefaultIclPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAutoConfigDefaultIslPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fgt_policy", flattenSwitchControllerAutoConfigDefaultFgtPolicy(o["fgt-policy"], d, "fgt_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-policy"], "SwitchControllerAutoConfigDefault-FgtPolicy"); ok {
			if err = d.Set("fgt_policy", vv); err != nil {
				return fmt.Errorf("Error reading fgt_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_policy: %v", err)
		}
	}

	if err = d.Set("icl_policy", flattenSwitchControllerAutoConfigDefaultIclPolicy(o["icl-policy"], d, "icl_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["icl-policy"], "SwitchControllerAutoConfigDefault-IclPolicy"); ok {
			if err = d.Set("icl_policy", vv); err != nil {
				return fmt.Errorf("Error reading icl_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icl_policy: %v", err)
		}
	}

	if err = d.Set("isl_policy", flattenSwitchControllerAutoConfigDefaultIslPolicy(o["isl-policy"], d, "isl_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["isl-policy"], "SwitchControllerAutoConfigDefault-IslPolicy"); ok {
			if err = d.Set("isl_policy", vv); err != nil {
				return fmt.Errorf("Error reading isl_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isl_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigDefaultFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigDefaultFgtPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandSwitchControllerAutoConfigDefaultIclPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerAutoConfigDefaultIslPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fgt_policy"); ok || d.HasChange("fgt_policy") {
		t, err := expandSwitchControllerAutoConfigDefaultFgtPolicy(d, v, "fgt_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-policy"] = t
		}
	}

	if v, ok := d.GetOk("icl_policy"); ok || d.HasChange("icl_policy") {
		t, err := expandSwitchControllerAutoConfigDefaultIclPolicy(d, v, "icl_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icl-policy"] = t
		}
	}

	if v, ok := d.GetOk("isl_policy"); ok || d.HasChange("isl_policy") {
		t, err := expandSwitchControllerAutoConfigDefaultIslPolicy(d, v, "isl_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isl-policy"] = t
		}
	}

	return &obj, nil
}
