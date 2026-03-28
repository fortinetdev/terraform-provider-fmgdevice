// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> IP address list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAddress6List() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6ListCreate,
		Read:   resourceFirewallAddress6ListRead,
		Update: resourceFirewallAddress6ListUpdate,
		Delete: resourceFirewallAddress6ListDelete,

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
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"net_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallAddress6ListCreate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	obj, err := getObjectFirewallAddress6List(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6List resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("ip")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAddress6List(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAddress6List(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAddress6List resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAddress6List(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAddress6List resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "ip"))

	return resourceFirewallAddress6ListRead(d, m)
}

func resourceFirewallAddress6ListUpdate(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	obj, err := getObjectFirewallAddress6List(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6List resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAddress6List(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6List resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip"))

	return resourceFirewallAddress6ListRead(d, m)
}

func resourceFirewallAddress6ListDelete(d *schema.ResourceData, m interface{}) error {
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
	address6 := d.Get("address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAddress6List(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6List resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6ListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	address6 := d.Get("address6").(string)
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
	if address6 == "" {
		address6 = importOptionChecking(m.(*FortiClient).Cfg, "address6")
		if address6 == "" {
			return fmt.Errorf("Parameter address6 is missing")
		}
		if err = d.Set("address6", address6); err != nil {
			return fmt.Errorf("Error set params address6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["address6"] = address6

	o, err := c.ReadFirewallAddress6List(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAddress6List resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6List(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6List resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddress6ListIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6ListNetId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6ListObjId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAddress6List(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenFirewallAddress6ListIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallAddress6List-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("net_id", flattenFirewallAddress6ListNetId2edl(o["net-id"], d, "net_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["net-id"], "FirewallAddress6List-NetId"); ok {
			if err = d.Set("net_id", vv); err != nil {
				return fmt.Errorf("Error reading net_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading net_id: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenFirewallAddress6ListObjId2edl(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "FirewallAddress6List-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddress6ListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAddress6ListIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6ListNetId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6ListObjId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress6List(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallAddress6ListIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("net_id"); ok || d.HasChange("net_id") {
		t, err := expandFirewallAddress6ListNetId2edl(d, v, "net_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["net-id"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandFirewallAddress6ListObjId2edl(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	return &obj, nil
}
