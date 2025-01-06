// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Protocol Independent Multicast (PIM) interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast6Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6InterfaceCreate,
		Read:   resourceRouterMulticast6InterfaceRead,
		Update: resourceRouterMulticast6InterfaceUpdate,
		Delete: resourceRouterMulticast6InterfaceDelete,

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
			"hello_holdtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceRouterMulticast6InterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6Interface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticast6Interface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterMulticast6Interface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterMulticast6Interface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterMulticast6InterfaceRead(d, m)
}

func resourceRouterMulticast6InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6Interface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6Interface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast6Interface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterMulticast6InterfaceRead(d, m)
}

func resourceRouterMulticast6InterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterMulticast6Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticast6Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6InterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6Interface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6Interface resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6InterfaceHelloHoldtime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectRouterMulticast6Interface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hello_holdtime", flattenRouterMulticast6InterfaceHelloHoldtime2edl(o["hello-holdtime"], d, "hello_holdtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-holdtime"], "RouterMulticast6Interface-HelloHoldtime"); ok {
			if err = d.Set("hello_holdtime", vv); err != nil {
				return fmt.Errorf("Error reading hello_holdtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_holdtime: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterMulticast6InterfaceHelloInterval2edl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterMulticast6Interface-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterMulticast6InterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterMulticast6Interface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticast6InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6InterfaceHelloHoldtime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectRouterMulticast6Interface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hello_holdtime"); ok || d.HasChange("hello_holdtime") {
		t, err := expandRouterMulticast6InterfaceHelloHoldtime2edl(d, v, "hello_holdtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-holdtime"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterMulticast6InterfaceHelloInterval2edl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterMulticast6InterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
