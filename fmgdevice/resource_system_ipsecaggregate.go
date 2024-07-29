// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure an aggregate of IPsec tunnels.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpsecAggregate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsecAggregateCreate,
		Read:   resourceSystemIpsecAggregateRead,
		Update: resourceSystemIpsecAggregateUpdate,
		Delete: resourceSystemIpsecAggregateDelete,

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
			"add_slbc_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemIpsecAggregateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemIpsecAggregate(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsecAggregate resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpsecAggregate(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpsecAggregate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpsecAggregateRead(d, m)
}

func resourceSystemIpsecAggregateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemIpsecAggregate(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsecAggregate resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpsecAggregate(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsecAggregate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemIpsecAggregateRead(d, m)
}

func resourceSystemIpsecAggregateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemIpsecAggregate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsecAggregate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsecAggregateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemIpsecAggregate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsecAggregate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsecAggregate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsecAggregate resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsecAggregateAddSlbcRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsecAggregateAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsecAggregateMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemIpsecAggregateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpsecAggregate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("add_slbc_route", flattenSystemIpsecAggregateAddSlbcRoute(o["add-slbc-route"], d, "add_slbc_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-slbc-route"], "SystemIpsecAggregate-AddSlbcRoute"); ok {
			if err = d.Set("add_slbc_route", vv); err != nil {
				return fmt.Errorf("Error reading add_slbc_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_slbc_route: %v", err)
		}
	}

	if err = d.Set("algorithm", flattenSystemIpsecAggregateAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["algorithm"], "SystemIpsecAggregate-Algorithm"); ok {
			if err = d.Set("algorithm", vv); err != nil {
				return fmt.Errorf("Error reading algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("member", flattenSystemIpsecAggregateMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "SystemIpsecAggregate-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemIpsecAggregateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemIpsecAggregate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsecAggregateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpsecAggregateAddSlbcRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsecAggregateAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsecAggregateMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemIpsecAggregateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsecAggregate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_slbc_route"); ok || d.HasChange("add_slbc_route") {
		t, err := expandSystemIpsecAggregateAddSlbcRoute(d, v, "add_slbc_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-slbc-route"] = t
		}
	}

	if v, ok := d.GetOk("algorithm"); ok || d.HasChange("algorithm") {
		t, err := expandSystemIpsecAggregateAlgorithm(d, v, "algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["algorithm"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandSystemIpsecAggregateMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemIpsecAggregateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
