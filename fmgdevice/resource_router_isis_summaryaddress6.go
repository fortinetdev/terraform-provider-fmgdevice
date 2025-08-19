// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IS-IS IPv6 summary address.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsisSummaryAddress6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisSummaryAddress6Create,
		Read:   resourceRouterIsisSummaryAddress6Read,
		Update: resourceRouterIsisSummaryAddress6Update,
		Delete: resourceRouterIsisSummaryAddress6Delete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterIsisSummaryAddress6Create(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterIsisSummaryAddress6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterIsisSummaryAddress6 resource while getting object: %v", err)
	}

	_, err = c.CreateRouterIsisSummaryAddress6(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterIsisSummaryAddress6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterIsisSummaryAddress6Read(d, m)
}

func resourceRouterIsisSummaryAddress6Update(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterIsisSummaryAddress6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisSummaryAddress6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsisSummaryAddress6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisSummaryAddress6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterIsisSummaryAddress6Read(d, m)
}

func resourceRouterIsisSummaryAddress6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterIsisSummaryAddress6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterIsisSummaryAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisSummaryAddress6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterIsisSummaryAddress6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisSummaryAddress6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsisSummaryAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisSummaryAddress6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisSummaryAddress6Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Level2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Prefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterIsisSummaryAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenRouterIsisSummaryAddress6Id2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterIsisSummaryAddress6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("level", flattenRouterIsisSummaryAddress6Level2edl(o["level"], d, "level")); err != nil {
		if vv, ok := fortiAPIPatch(o["level"], "RouterIsisSummaryAddress6-Level"); ok {
			if err = d.Set("level", vv); err != nil {
				return fmt.Errorf("Error reading level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterIsisSummaryAddress6Prefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "RouterIsisSummaryAddress6-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenRouterIsisSummaryAddress6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterIsisSummaryAddress6Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Level2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Prefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsisSummaryAddress6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterIsisSummaryAddress6Id2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandRouterIsisSummaryAddress6Level2edl(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandRouterIsisSummaryAddress6Prefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
