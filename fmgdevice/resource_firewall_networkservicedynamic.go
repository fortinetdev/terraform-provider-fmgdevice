// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Dynamic Network Services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallNetworkServiceDynamic() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallNetworkServiceDynamicCreate,
		Read:   resourceFirewallNetworkServiceDynamicRead,
		Update: resourceFirewallNetworkServiceDynamicUpdate,
		Delete: resourceFirewallNetworkServiceDynamicDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallNetworkServiceDynamicCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallNetworkServiceDynamic(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallNetworkServiceDynamic resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallNetworkServiceDynamic(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallNetworkServiceDynamic(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallNetworkServiceDynamic resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallNetworkServiceDynamic(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallNetworkServiceDynamic resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallNetworkServiceDynamicRead(d, m)
}

func resourceFirewallNetworkServiceDynamicUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallNetworkServiceDynamic(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallNetworkServiceDynamic resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallNetworkServiceDynamic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallNetworkServiceDynamic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallNetworkServiceDynamicRead(d, m)
}

func resourceFirewallNetworkServiceDynamicDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteFirewallNetworkServiceDynamic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallNetworkServiceDynamic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallNetworkServiceDynamicRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallNetworkServiceDynamic(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallNetworkServiceDynamic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallNetworkServiceDynamic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallNetworkServiceDynamic resource from API: %v", err)
	}
	return nil
}

func flattenFirewallNetworkServiceDynamicComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallNetworkServiceDynamicFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallNetworkServiceDynamicId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallNetworkServiceDynamicName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallNetworkServiceDynamicSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallNetworkServiceDynamic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenFirewallNetworkServiceDynamicComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallNetworkServiceDynamic-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("filter", flattenFirewallNetworkServiceDynamicFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "FirewallNetworkServiceDynamic-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallNetworkServiceDynamicId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallNetworkServiceDynamic-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallNetworkServiceDynamicName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallNetworkServiceDynamic-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sdn", flattenFirewallNetworkServiceDynamicSdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "FirewallNetworkServiceDynamic-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	return nil
}

func flattenFirewallNetworkServiceDynamicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallNetworkServiceDynamicComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallNetworkServiceDynamicFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallNetworkServiceDynamicId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallNetworkServiceDynamicName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallNetworkServiceDynamicSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallNetworkServiceDynamic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallNetworkServiceDynamicComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandFirewallNetworkServiceDynamicFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallNetworkServiceDynamicId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallNetworkServiceDynamicName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandFirewallNetworkServiceDynamicSdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	return &obj, nil
}
