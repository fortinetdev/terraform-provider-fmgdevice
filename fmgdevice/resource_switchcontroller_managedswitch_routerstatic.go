// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure static routes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitchRouterStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchRouterStaticCreate,
		Read:   resourceSwitchControllerManagedSwitchRouterStaticRead,
		Update: resourceSwitchControllerManagedSwitchRouterStaticUpdate,
		Delete: resourceSwitchControllerManagedSwitchRouterStaticDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchRouterStaticCreate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchRouterStatic(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitchRouterStatic resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerManagedSwitchRouterStatic(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerManagedSwitchRouterStatic(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerManagedSwitchRouterStatic resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSwitchControllerManagedSwitchRouterStatic(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerManagedSwitchRouterStatic resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSwitchControllerManagedSwitchRouterStaticRead(d, m)
			} else {
				return fmt.Errorf("Error creating SwitchControllerManagedSwitchRouterStatic resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchRouterStaticRead(d, m)
}

func resourceSwitchControllerManagedSwitchRouterStaticUpdate(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectSwitchControllerManagedSwitchRouterStatic(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchRouterStatic resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerManagedSwitchRouterStatic(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitchRouterStatic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerManagedSwitchRouterStaticRead(d, m)
}

func resourceSwitchControllerManagedSwitchRouterStaticDelete(d *schema.ResourceData, m interface{}) error {
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
	managed_switch := d.Get("managed_switch").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteSwitchControllerManagedSwitchRouterStatic(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitchRouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchRouterStaticRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	managed_switch := d.Get("managed_switch").(string)
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
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadSwitchControllerManagedSwitchRouterStatic(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchRouterStatic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitchRouterStatic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitchRouterStatic resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchRouterStaticBlackhole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchRouterStaticDistance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchRouterStaticDynamicGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouterStaticSwitchId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerManagedSwitchRouterStaticVrf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerManagedSwitchRouterStatic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("blackhole", flattenSwitchControllerManagedSwitchRouterStaticBlackhole2edl(o["blackhole"], d, "blackhole")); err != nil {
		if vv, ok := fortiAPIPatch(o["blackhole"], "SwitchControllerManagedSwitchRouterStatic-Blackhole"); ok {
			if err = d.Set("blackhole", vv); err != nil {
				return fmt.Errorf("Error reading blackhole: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("comment", flattenSwitchControllerManagedSwitchRouterStaticComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SwitchControllerManagedSwitchRouterStatic-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("device", flattenSwitchControllerManagedSwitchRouterStaticDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SwitchControllerManagedSwitchRouterStatic-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("distance", flattenSwitchControllerManagedSwitchRouterStaticDistance2edl(o["distance"], d, "distance")); err != nil {
		if vv, ok := fortiAPIPatch(o["distance"], "SwitchControllerManagedSwitchRouterStatic-Distance"); ok {
			if err = d.Set("distance", vv); err != nil {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("dst", flattenSwitchControllerManagedSwitchRouterStaticDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SwitchControllerManagedSwitchRouterStatic-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", flattenSwitchControllerManagedSwitchRouterStaticDynamicGateway2edl(o["dynamic-gateway"], d, "dynamic_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-gateway"], "SwitchControllerManagedSwitchRouterStatic-DynamicGateway"); ok {
			if err = d.Set("dynamic_gateway", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSwitchControllerManagedSwitchRouterStaticGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SwitchControllerManagedSwitchRouterStatic-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerManagedSwitchRouterStaticId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerManagedSwitchRouterStatic-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerManagedSwitchRouterStaticStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerManagedSwitchRouterStatic-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchRouterStaticSwitchId2edl(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "SwitchControllerManagedSwitchRouterStatic-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSwitchControllerManagedSwitchRouterStaticVrf2edl(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "SwitchControllerManagedSwitchRouterStatic-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchRouterStaticFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerManagedSwitchRouterStaticBlackhole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchRouterStaticDistance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerManagedSwitchRouterStaticDynamicGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouterStaticSwitchId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerManagedSwitchRouterStaticVrf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerManagedSwitchRouterStatic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("blackhole"); ok || d.HasChange("blackhole") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticBlackhole2edl(d, v, "blackhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok || d.HasChange("distance") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticDistance2edl(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_gateway"); ok || d.HasChange("dynamic_gateway") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticDynamicGateway2edl(d, v, "dynamic_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticSwitchId2edl(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandSwitchControllerManagedSwitchRouterStaticVrf2edl(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
