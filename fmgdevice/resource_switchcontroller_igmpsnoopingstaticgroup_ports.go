// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure static group in switches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIgmpSnoopingStaticGroupPorts() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIgmpSnoopingStaticGroupPortsCreate,
		Read:   resourceSwitchControllerIgmpSnoopingStaticGroupPortsRead,
		Update: resourceSwitchControllerIgmpSnoopingStaticGroupPortsUpdate,
		Delete: resourceSwitchControllerIgmpSnoopingStaticGroupPortsDelete,

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
			"igmp_snooping_static_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerIgmpSnoopingStaticGroupPortsCreate(d *schema.ResourceData, m interface{}) error {
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
	igmp_snooping_static_group := d.Get("igmp_snooping_static_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["igmp_snooping_static_group"] = igmp_snooping_static_group

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroupPorts(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroupPorts resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerIgmpSnoopingStaticGroupPorts(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerIgmpSnoopingStaticGroupPorts(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroupPorts resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerIgmpSnoopingStaticGroupPorts(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroupPorts resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerIgmpSnoopingStaticGroupPortsRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupPortsUpdate(d *schema.ResourceData, m interface{}) error {
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
	igmp_snooping_static_group := d.Get("igmp_snooping_static_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["igmp_snooping_static_group"] = igmp_snooping_static_group

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroupPorts(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroupPorts resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerIgmpSnoopingStaticGroupPorts(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroupPorts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerIgmpSnoopingStaticGroupPortsRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupPortsDelete(d *schema.ResourceData, m interface{}) error {
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
	igmp_snooping_static_group := d.Get("igmp_snooping_static_group").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["igmp_snooping_static_group"] = igmp_snooping_static_group

	wsParams["adom"] = adomv

	err = c.DeleteSwitchControllerIgmpSnoopingStaticGroupPorts(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIgmpSnoopingStaticGroupPorts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIgmpSnoopingStaticGroupPortsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	igmp_snooping_static_group := d.Get("igmp_snooping_static_group").(string)
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
	if igmp_snooping_static_group == "" {
		igmp_snooping_static_group = importOptionChecking(m.(*FortiClient).Cfg, "igmp_snooping_static_group")
		if igmp_snooping_static_group == "" {
			return fmt.Errorf("Parameter igmp_snooping_static_group is missing")
		}
		if err = d.Set("igmp_snooping_static_group", igmp_snooping_static_group); err != nil {
			return fmt.Errorf("Error set params igmp_snooping_static_group: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["igmp_snooping_static_group"] = igmp_snooping_static_group

	o, err := c.ReadSwitchControllerIgmpSnoopingStaticGroupPorts(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroupPorts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIgmpSnoopingStaticGroupPorts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroupPorts resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerIgmpSnoopingStaticGroupPorts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSwitchControllerIgmpSnoopingStaticGroupPortsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerIgmpSnoopingStaticGroupPorts-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ports", flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "SwitchControllerIgmpSnoopingStaticGroupPorts-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerIgmpSnoopingStaticGroupPorts-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerIgmpSnoopingStaticGroupPorts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupPortsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	return &obj, nil
}
