// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Switch binding list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigCustomSwitchBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigCustomSwitchBindingCreate,
		Read:   resourceSwitchControllerAutoConfigCustomSwitchBindingRead,
		Update: resourceSwitchControllerAutoConfigCustomSwitchBindingUpdate,
		Delete: resourceSwitchControllerAutoConfigCustomSwitchBindingDelete,

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
			"custom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigCustomSwitchBindingCreate(d *schema.ResourceData, m interface{}) error {
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
	custom := d.Get("custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["custom"] = custom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerAutoConfigCustomSwitchBinding(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustomSwitchBinding resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerAutoConfigCustomSwitchBinding(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustomSwitchBinding resource: %v", err)
	}

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerAutoConfigCustomSwitchBindingRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomSwitchBindingUpdate(d *schema.ResourceData, m interface{}) error {
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
	custom := d.Get("custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["custom"] = custom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerAutoConfigCustomSwitchBinding(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustomSwitchBinding resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAutoConfigCustomSwitchBinding(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustomSwitchBinding resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "switch_id"))

	return resourceSwitchControllerAutoConfigCustomSwitchBindingRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomSwitchBindingDelete(d *schema.ResourceData, m interface{}) error {
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
	custom := d.Get("custom").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["custom"] = custom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerAutoConfigCustomSwitchBinding(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigCustomSwitchBinding resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigCustomSwitchBindingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	custom := d.Get("custom").(string)
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
	if custom == "" {
		custom = importOptionChecking(m.(*FortiClient).Cfg, "custom")
		if custom == "" {
			return fmt.Errorf("Parameter custom is missing")
		}
		if err = d.Set("custom", custom); err != nil {
			return fmt.Errorf("Error set params custom: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["custom"] = custom

	o, err := c.ReadSwitchControllerAutoConfigCustomSwitchBinding(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustomSwitchBinding resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigCustomSwitchBinding(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustomSwitchBinding resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigCustomSwitchBinding(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("policy", flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy2edl(o["policy"], d, "policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy"], "SwitchControllerAutoConfigCustomSwitchBinding-Policy"); ok {
			if err = d.Set("policy", vv); err != nil {
				return fmt.Errorf("Error reading policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerAutoConfigCustomSwitchBinding-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigCustomSwitchBindingPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigCustomSwitchBinding(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("policy"); ok || d.HasChange("policy") {
		t, err := expandSwitchControllerAutoConfigCustomSwitchBindingPolicy2edl(d, v, "policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	return &obj, nil
}
