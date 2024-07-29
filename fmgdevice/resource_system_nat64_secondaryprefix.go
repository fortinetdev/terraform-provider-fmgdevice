// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemNat64SecondaryPrefix

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNat64SecondaryPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNat64SecondaryPrefixCreate,
		Read:   resourceSystemNat64SecondaryPrefixRead,
		Update: resourceSystemNat64SecondaryPrefixUpdate,
		Delete: resourceSystemNat64SecondaryPrefixDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nat64_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemNat64SecondaryPrefixCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemNat64SecondaryPrefix(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNat64SecondaryPrefix resource while getting object: %v", err)
	}

	_, err = c.CreateSystemNat64SecondaryPrefix(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemNat64SecondaryPrefix resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNat64SecondaryPrefixRead(d, m)
}

func resourceSystemNat64SecondaryPrefixUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemNat64SecondaryPrefix(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNat64SecondaryPrefix resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNat64SecondaryPrefix(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemNat64SecondaryPrefix resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemNat64SecondaryPrefixRead(d, m)
}

func resourceSystemNat64SecondaryPrefixDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemNat64SecondaryPrefix(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNat64SecondaryPrefix resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNat64SecondaryPrefixRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNat64SecondaryPrefix(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNat64SecondaryPrefix resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNat64SecondaryPrefix(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNat64SecondaryPrefix resource from API: %v", err)
	}
	return nil
}

func flattenSystemNat64SecondaryPrefixName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNat64SecondaryPrefixNat64Prefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNat64SecondaryPrefix(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemNat64SecondaryPrefixName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemNat64SecondaryPrefix-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat64_prefix", flattenSystemNat64SecondaryPrefixNat64Prefix2edl(o["nat64-prefix"], d, "nat64_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64-prefix"], "SystemNat64SecondaryPrefix-Nat64Prefix"); ok {
			if err = d.Set("nat64_prefix", vv); err != nil {
				return fmt.Errorf("Error reading nat64_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64_prefix: %v", err)
		}
	}

	return nil
}

func flattenSystemNat64SecondaryPrefixFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNat64SecondaryPrefixName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64SecondaryPrefixNat64Prefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNat64SecondaryPrefix(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemNat64SecondaryPrefixName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat64_prefix"); ok || d.HasChange("nat64_prefix") {
		t, err := expandSystemNat64SecondaryPrefixNat64Prefix2edl(d, v, "nat64_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64-prefix"] = t
		}
	}

	return &obj, nil
}
