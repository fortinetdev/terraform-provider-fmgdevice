// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: BGP aggregate address table.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpAggregateAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAggregateAddressCreate,
		Read:   resourceRouterBgpAggregateAddressRead,
		Update: resourceRouterBgpAggregateAddressUpdate,
		Delete: resourceRouterBgpAggregateAddressDelete,

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
			"as_set": &schema.Schema{
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
			"summary_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterBgpAggregateAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpAggregateAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterBgpAggregateAddress resource while getting object: %v", err)
	}

	v, err := c.CreateRouterBgpAggregateAddress(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterBgpAggregateAddress resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceRouterBgpAggregateAddressRead(d, m)
		} else {
			return fmt.Errorf("Error creating RouterBgpAggregateAddress resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpAggregateAddressRead(d, m)
}

func resourceRouterBgpAggregateAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgpAggregateAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpAggregateAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgpAggregateAddress(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgpAggregateAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterBgpAggregateAddressRead(d, m)
}

func resourceRouterBgpAggregateAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterBgpAggregateAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterBgpAggregateAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpAggregateAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgpAggregateAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpAggregateAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgpAggregateAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgpAggregateAddress resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpAggregateAddressAsSet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterBgpAggregateAddressSummaryOnly2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBgpAggregateAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("as_set", flattenRouterBgpAggregateAddressAsSet2edl(o["as-set"], d, "as_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-set"], "RouterBgpAggregateAddress-AsSet"); ok {
			if err = d.Set("as_set", vv); err != nil {
				return fmt.Errorf("Error reading as_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_set: %v", err)
		}
	}

	if err = d.Set("fosid", flattenRouterBgpAggregateAddressId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterBgpAggregateAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterBgpAggregateAddressPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "RouterBgpAggregateAddress-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("summary_only", flattenRouterBgpAggregateAddressSummaryOnly2edl(o["summary-only"], d, "summary_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["summary-only"], "RouterBgpAggregateAddress-SummaryOnly"); ok {
			if err = d.Set("summary_only", vv); err != nil {
				return fmt.Errorf("Error reading summary_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading summary_only: %v", err)
		}
	}

	return nil
}

func flattenRouterBgpAggregateAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBgpAggregateAddressAsSet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterBgpAggregateAddressSummaryOnly2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgpAggregateAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("as_set"); ok || d.HasChange("as_set") {
		t, err := expandRouterBgpAggregateAddressAsSet2edl(d, v, "as_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-set"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterBgpAggregateAddressId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandRouterBgpAggregateAddressPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("summary_only"); ok || d.HasChange("summary_only") {
		t, err := expandRouterBgpAggregateAddressSummaryOnly2edl(d, v, "summary_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-only"] = t
		}
	}

	return &obj, nil
}
