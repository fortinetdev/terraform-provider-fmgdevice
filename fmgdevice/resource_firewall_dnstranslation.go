// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS translation.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDnstranslation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDnstranslationCreate,
		Read:   resourceFirewallDnstranslationRead,
		Update: resourceFirewallDnstranslationUpdate,
		Delete: resourceFirewallDnstranslationDelete,

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
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallDnstranslationCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallDnstranslation(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDnstranslation resource while getting object: %v", err)
	}

	_, err = c.CreateFirewallDnstranslation(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDnstranslation resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallDnstranslationRead(d, m)
}

func resourceFirewallDnstranslationUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallDnstranslation(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDnstranslation resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallDnstranslation(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDnstranslation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallDnstranslationRead(d, m)
}

func resourceFirewallDnstranslationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallDnstranslation(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDnstranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDnstranslationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallDnstranslation(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDnstranslation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDnstranslation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDnstranslation resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDnstranslationDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallDnstranslation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst", flattenFirewallDnstranslationDst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "FirewallDnstranslation-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallDnstranslationId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallDnstranslation-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("netmask", flattenFirewallDnstranslationNetmask(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "FirewallDnstranslation-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("src", flattenFirewallDnstranslationSrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "FirewallDnstranslation-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	return nil
}

func flattenFirewallDnstranslationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDnstranslationDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDnstranslation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandFirewallDnstranslationDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallDnstranslationId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandFirewallDnstranslationNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandFirewallDnstranslationSrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
