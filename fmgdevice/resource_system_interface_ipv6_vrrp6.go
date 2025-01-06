// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 VRRP configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6Vrrp6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Vrrp6Create,
		Read:   resourceSystemInterfaceIpv6Vrrp6Read,
		Update: resourceSystemInterfaceIpv6Vrrp6Update,
		Delete: resourceSystemInterfaceIpv6Vrrp6Delete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ignore_default_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrdst_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrdst6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vrip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Vrrp6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceIpv6Vrrp6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Vrrp6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemInterfaceIpv6Vrrp6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemInterfaceIpv6Vrrp6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceSystemInterfaceIpv6Vrrp6Read(d, m)
}

func resourceSystemInterfaceIpv6Vrrp6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceIpv6Vrrp6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Vrrp6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6Vrrp6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6Vrrp6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceSystemInterfaceIpv6Vrrp6Read(d, m)
}

func resourceSystemInterfaceIpv6Vrrp6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	err = c.DeleteSystemInterfaceIpv6Vrrp6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6Vrrp6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Vrrp6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	var_interface := d.Get("interface").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6Vrrp6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Vrrp6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6Vrrp6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6Vrrp6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Vrrp6AcceptMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6AdvInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Preempt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Priority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6StartTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Status3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6VrdstPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrdst63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Vrrp6Vrgrp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrip63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6Vrrp6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accept_mode", flattenSystemInterfaceIpv6Vrrp6AcceptMode3rdl(o["accept-mode"], d, "accept_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-mode"], "SystemInterfaceIpv6Vrrp6-AcceptMode"); ok {
			if err = d.Set("accept_mode", vv); err != nil {
				return fmt.Errorf("Error reading accept_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_mode: %v", err)
		}
	}

	if err = d.Set("adv_interval", flattenSystemInterfaceIpv6Vrrp6AdvInterval3rdl(o["adv-interval"], d, "adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-interval"], "SystemInterfaceIpv6Vrrp6-AdvInterval"); ok {
			if err = d.Set("adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_interval: %v", err)
		}
	}

	if err = d.Set("ignore_default_route", flattenSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute3rdl(o["ignore-default-route"], d, "ignore_default_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-default-route"], "SystemInterfaceIpv6Vrrp6-IgnoreDefaultRoute"); ok {
			if err = d.Set("ignore_default_route", vv); err != nil {
				return fmt.Errorf("Error reading ignore_default_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_default_route: %v", err)
		}
	}

	if err = d.Set("preempt", flattenSystemInterfaceIpv6Vrrp6Preempt3rdl(o["preempt"], d, "preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["preempt"], "SystemInterfaceIpv6Vrrp6-Preempt"); ok {
			if err = d.Set("preempt", vv); err != nil {
				return fmt.Errorf("Error reading preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemInterfaceIpv6Vrrp6Priority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemInterfaceIpv6Vrrp6-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("start_time", flattenSystemInterfaceIpv6Vrrp6StartTime3rdl(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "SystemInterfaceIpv6Vrrp6-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemInterfaceIpv6Vrrp6Status3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemInterfaceIpv6Vrrp6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vrdst_priority", flattenSystemInterfaceIpv6Vrrp6VrdstPriority3rdl(o["vrdst-priority"], d, "vrdst_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst-priority"], "SystemInterfaceIpv6Vrrp6-VrdstPriority"); ok {
			if err = d.Set("vrdst_priority", vv); err != nil {
				return fmt.Errorf("Error reading vrdst_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst_priority: %v", err)
		}
	}

	if err = d.Set("vrdst6", flattenSystemInterfaceIpv6Vrrp6Vrdst63rdl(o["vrdst6"], d, "vrdst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst6"], "SystemInterfaceIpv6Vrrp6-Vrdst6"); ok {
			if err = d.Set("vrdst6", vv); err != nil {
				return fmt.Errorf("Error reading vrdst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst6: %v", err)
		}
	}

	if err = d.Set("vrgrp", flattenSystemInterfaceIpv6Vrrp6Vrgrp3rdl(o["vrgrp"], d, "vrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrgrp"], "SystemInterfaceIpv6Vrrp6-Vrgrp"); ok {
			if err = d.Set("vrgrp", vv); err != nil {
				return fmt.Errorf("Error reading vrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrgrp: %v", err)
		}
	}

	if err = d.Set("vrid", flattenSystemInterfaceIpv6Vrrp6Vrid3rdl(o["vrid"], d, "vrid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrid"], "SystemInterfaceIpv6Vrrp6-Vrid"); ok {
			if err = d.Set("vrid", vv); err != nil {
				return fmt.Errorf("Error reading vrid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrid: %v", err)
		}
	}

	if err = d.Set("vrip6", flattenSystemInterfaceIpv6Vrrp6Vrip63rdl(o["vrip6"], d, "vrip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip6"], "SystemInterfaceIpv6Vrrp6-Vrip6"); ok {
			if err = d.Set("vrip6", vv); err != nil {
				return fmt.Errorf("Error reading vrip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip6: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6Vrrp6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Vrrp6AcceptMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6AdvInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Preempt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Priority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6StartTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Status3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6VrdstPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrdst63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Vrrp6Vrgrp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrip63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6Vrrp6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_mode"); ok || d.HasChange("accept_mode") {
		t, err := expandSystemInterfaceIpv6Vrrp6AcceptMode3rdl(d, v, "accept_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-mode"] = t
		}
	}

	if v, ok := d.GetOk("adv_interval"); ok || d.HasChange("adv_interval") {
		t, err := expandSystemInterfaceIpv6Vrrp6AdvInterval3rdl(d, v, "adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("ignore_default_route"); ok || d.HasChange("ignore_default_route") {
		t, err := expandSystemInterfaceIpv6Vrrp6IgnoreDefaultRoute3rdl(d, v, "ignore_default_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-default-route"] = t
		}
	}

	if v, ok := d.GetOk("preempt"); ok || d.HasChange("preempt") {
		t, err := expandSystemInterfaceIpv6Vrrp6Preempt3rdl(d, v, "preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemInterfaceIpv6Vrrp6Priority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandSystemInterfaceIpv6Vrrp6StartTime3rdl(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemInterfaceIpv6Vrrp6Status3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vrdst_priority"); ok || d.HasChange("vrdst_priority") {
		t, err := expandSystemInterfaceIpv6Vrrp6VrdstPriority3rdl(d, v, "vrdst_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst-priority"] = t
		}
	}

	if v, ok := d.GetOk("vrdst6"); ok || d.HasChange("vrdst6") {
		t, err := expandSystemInterfaceIpv6Vrrp6Vrdst63rdl(d, v, "vrdst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst6"] = t
		}
	}

	if v, ok := d.GetOk("vrgrp"); ok || d.HasChange("vrgrp") {
		t, err := expandSystemInterfaceIpv6Vrrp6Vrgrp3rdl(d, v, "vrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vrid"); ok || d.HasChange("vrid") {
		t, err := expandSystemInterfaceIpv6Vrrp6Vrid3rdl(d, v, "vrid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrid"] = t
		}
	}

	if v, ok := d.GetOk("vrip6"); ok || d.HasChange("vrip6") {
		t, err := expandSystemInterfaceIpv6Vrrp6Vrip63rdl(d, v, "vrip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip6"] = t
		}
	}

	return &obj, nil
}
