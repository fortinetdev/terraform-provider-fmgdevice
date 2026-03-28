// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Default network service entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationListDefaultNetworkServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationListDefaultNetworkServicesCreate,
		Read:   resourceApplicationListDefaultNetworkServicesRead,
		Update: resourceApplicationListDefaultNetworkServicesUpdate,
		Delete: resourceApplicationListDefaultNetworkServicesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"violation_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationListDefaultNetworkServicesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	obj, err := getObjectApplicationListDefaultNetworkServices(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationListDefaultNetworkServices resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationListDefaultNetworkServices(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationListDefaultNetworkServices(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationListDefaultNetworkServices resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateApplicationListDefaultNetworkServices(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ApplicationListDefaultNetworkServices resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceApplicationListDefaultNetworkServicesRead(d, m)
}

func resourceApplicationListDefaultNetworkServicesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	obj, err := getObjectApplicationListDefaultNetworkServices(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationListDefaultNetworkServices resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateApplicationListDefaultNetworkServices(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationListDefaultNetworkServices resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceApplicationListDefaultNetworkServicesRead(d, m)
}

func resourceApplicationListDefaultNetworkServicesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	wsParams["adom"] = adomv

	err = c.DeleteApplicationListDefaultNetworkServices(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationListDefaultNetworkServices resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationListDefaultNetworkServicesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	list := d.Get("list").(string)
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
	if list == "" {
		list = importOptionChecking(m.(*FortiClient).Cfg, "list")
		if list == "" {
			return fmt.Errorf("Parameter list is missing")
		}
		if err = d.Set("list", list); err != nil {
			return fmt.Errorf("Error set params list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	o, err := c.ReadApplicationListDefaultNetworkServices(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ApplicationListDefaultNetworkServices resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationListDefaultNetworkServices(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationListDefaultNetworkServices resource from API: %v", err)
	}
	return nil
}

func flattenApplicationListDefaultNetworkServicesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServicesPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServicesServices2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListDefaultNetworkServicesViolationAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationListDefaultNetworkServices(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenApplicationListDefaultNetworkServicesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ApplicationListDefaultNetworkServices-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("port", flattenApplicationListDefaultNetworkServicesPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ApplicationListDefaultNetworkServices-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("services", flattenApplicationListDefaultNetworkServicesServices2edl(o["services"], d, "services")); err != nil {
		if vv, ok := fortiAPIPatch(o["services"], "ApplicationListDefaultNetworkServices-Services"); ok {
			if err = d.Set("services", vv); err != nil {
				return fmt.Errorf("Error reading services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading services: %v", err)
		}
	}

	if err = d.Set("violation_action", flattenApplicationListDefaultNetworkServicesViolationAction2edl(o["violation-action"], d, "violation_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["violation-action"], "ApplicationListDefaultNetworkServices-ViolationAction"); ok {
			if err = d.Set("violation_action", vv); err != nil {
				return fmt.Errorf("Error reading violation_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading violation_action: %v", err)
		}
	}

	return nil
}

func flattenApplicationListDefaultNetworkServicesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationListDefaultNetworkServicesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesServices2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListDefaultNetworkServicesViolationAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationListDefaultNetworkServices(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandApplicationListDefaultNetworkServicesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandApplicationListDefaultNetworkServicesPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("services"); ok || d.HasChange("services") {
		t, err := expandApplicationListDefaultNetworkServicesServices2edl(d, v, "services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["services"] = t
		}
	}

	if v, ok := d.GetOk("violation_action"); ok || d.HasChange("violation_action") {
		t, err := expandApplicationListDefaultNetworkServicesViolationAction2edl(d, v, "violation_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["violation-action"] = t
		}
	}

	return &obj, nil
}
