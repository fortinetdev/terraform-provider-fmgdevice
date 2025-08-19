// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure quarantine options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusQuarantineUpdate,
		Read:   resourceAntivirusQuarantineRead,
		Update: resourceAntivirusQuarantineUpdate,
		Delete: resourceAntivirusQuarantineDelete,

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
			"agelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_blocked": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drop_heuristic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drop_infected": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drop_intercepted": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"drop_machine_learning": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lowspace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"maxfilesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"quarantine_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"store_blocked": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"store_heuristic": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"store_infected": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"store_intercepted": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"store_machine_learning": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectAntivirusQuarantine(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource while getting object: %v", err)
	}

	_, err = c.UpdateAntivirusQuarantine(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusQuarantine")

	return resourceAntivirusQuarantineRead(d, m)
}

func resourceAntivirusQuarantineDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAntivirusQuarantine(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusQuarantineRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusQuarantine(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusQuarantine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusQuarantineAgelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineDropHeuristic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineDropInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineDropIntercepted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineDropMachineLearning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineLowspace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineMaxfilesize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineQuarantineQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineStoreHeuristic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineStoreInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineStoreIntercepted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusQuarantineStoreMachineLearning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectAntivirusQuarantine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("agelimit", flattenAntivirusQuarantineAgelimit(o["agelimit"], d, "agelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["agelimit"], "AntivirusQuarantine-Agelimit"); ok {
			if err = d.Set("agelimit", vv); err != nil {
				return fmt.Errorf("Error reading agelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agelimit: %v", err)
		}
	}

	if err = d.Set("destination", flattenAntivirusQuarantineDestination(o["destination"], d, "destination")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination"], "AntivirusQuarantine-Destination"); ok {
			if err = d.Set("destination", vv); err != nil {
				return fmt.Errorf("Error reading destination: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("drop_blocked", flattenAntivirusQuarantineDropBlocked(o["drop-blocked"], d, "drop_blocked")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-blocked"], "AntivirusQuarantine-DropBlocked"); ok {
			if err = d.Set("drop_blocked", vv); err != nil {
				return fmt.Errorf("Error reading drop_blocked: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_blocked: %v", err)
		}
	}

	if err = d.Set("drop_heuristic", flattenAntivirusQuarantineDropHeuristic(o["drop-heuristic"], d, "drop_heuristic")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-heuristic"], "AntivirusQuarantine-DropHeuristic"); ok {
			if err = d.Set("drop_heuristic", vv); err != nil {
				return fmt.Errorf("Error reading drop_heuristic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_heuristic: %v", err)
		}
	}

	if err = d.Set("drop_infected", flattenAntivirusQuarantineDropInfected(o["drop-infected"], d, "drop_infected")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-infected"], "AntivirusQuarantine-DropInfected"); ok {
			if err = d.Set("drop_infected", vv); err != nil {
				return fmt.Errorf("Error reading drop_infected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_infected: %v", err)
		}
	}

	if err = d.Set("drop_intercepted", flattenAntivirusQuarantineDropIntercepted(o["drop-intercepted"], d, "drop_intercepted")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-intercepted"], "AntivirusQuarantine-DropIntercepted"); ok {
			if err = d.Set("drop_intercepted", vv); err != nil {
				return fmt.Errorf("Error reading drop_intercepted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_intercepted: %v", err)
		}
	}

	if err = d.Set("drop_machine_learning", flattenAntivirusQuarantineDropMachineLearning(o["drop-machine-learning"], d, "drop_machine_learning")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-machine-learning"], "AntivirusQuarantine-DropMachineLearning"); ok {
			if err = d.Set("drop_machine_learning", vv); err != nil {
				return fmt.Errorf("Error reading drop_machine_learning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_machine_learning: %v", err)
		}
	}

	if err = d.Set("lowspace", flattenAntivirusQuarantineLowspace(o["lowspace"], d, "lowspace")); err != nil {
		if vv, ok := fortiAPIPatch(o["lowspace"], "AntivirusQuarantine-Lowspace"); ok {
			if err = d.Set("lowspace", vv); err != nil {
				return fmt.Errorf("Error reading lowspace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lowspace: %v", err)
		}
	}

	if err = d.Set("maxfilesize", flattenAntivirusQuarantineMaxfilesize(o["maxfilesize"], d, "maxfilesize")); err != nil {
		if vv, ok := fortiAPIPatch(o["maxfilesize"], "AntivirusQuarantine-Maxfilesize"); ok {
			if err = d.Set("maxfilesize", vv); err != nil {
				return fmt.Errorf("Error reading maxfilesize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maxfilesize: %v", err)
		}
	}

	if err = d.Set("quarantine_quota", flattenAntivirusQuarantineQuarantineQuota(o["quarantine-quota"], d, "quarantine_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-quota"], "AntivirusQuarantine-QuarantineQuota"); ok {
			if err = d.Set("quarantine_quota", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_quota: %v", err)
		}
	}

	if err = d.Set("store_blocked", flattenAntivirusQuarantineStoreBlocked(o["store-blocked"], d, "store_blocked")); err != nil {
		if vv, ok := fortiAPIPatch(o["store-blocked"], "AntivirusQuarantine-StoreBlocked"); ok {
			if err = d.Set("store_blocked", vv); err != nil {
				return fmt.Errorf("Error reading store_blocked: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading store_blocked: %v", err)
		}
	}

	if err = d.Set("store_heuristic", flattenAntivirusQuarantineStoreHeuristic(o["store-heuristic"], d, "store_heuristic")); err != nil {
		if vv, ok := fortiAPIPatch(o["store-heuristic"], "AntivirusQuarantine-StoreHeuristic"); ok {
			if err = d.Set("store_heuristic", vv); err != nil {
				return fmt.Errorf("Error reading store_heuristic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading store_heuristic: %v", err)
		}
	}

	if err = d.Set("store_infected", flattenAntivirusQuarantineStoreInfected(o["store-infected"], d, "store_infected")); err != nil {
		if vv, ok := fortiAPIPatch(o["store-infected"], "AntivirusQuarantine-StoreInfected"); ok {
			if err = d.Set("store_infected", vv); err != nil {
				return fmt.Errorf("Error reading store_infected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading store_infected: %v", err)
		}
	}

	if err = d.Set("store_intercepted", flattenAntivirusQuarantineStoreIntercepted(o["store-intercepted"], d, "store_intercepted")); err != nil {
		if vv, ok := fortiAPIPatch(o["store-intercepted"], "AntivirusQuarantine-StoreIntercepted"); ok {
			if err = d.Set("store_intercepted", vv); err != nil {
				return fmt.Errorf("Error reading store_intercepted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading store_intercepted: %v", err)
		}
	}

	if err = d.Set("store_machine_learning", flattenAntivirusQuarantineStoreMachineLearning(o["store-machine-learning"], d, "store_machine_learning")); err != nil {
		if vv, ok := fortiAPIPatch(o["store-machine-learning"], "AntivirusQuarantine-StoreMachineLearning"); ok {
			if err = d.Set("store_machine_learning", vv); err != nil {
				return fmt.Errorf("Error reading store_machine_learning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading store_machine_learning: %v", err)
		}
	}

	return nil
}

func flattenAntivirusQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusQuarantineAgelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineDropHeuristic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineDropInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineDropIntercepted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineDropMachineLearning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineLowspace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineMaxfilesize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineQuarantineQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineStoreHeuristic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineStoreInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineStoreIntercepted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusQuarantineStoreMachineLearning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectAntivirusQuarantine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("agelimit"); ok || d.HasChange("agelimit") {
		t, err := expandAntivirusQuarantineAgelimit(d, v, "agelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agelimit"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok || d.HasChange("destination") {
		t, err := expandAntivirusQuarantineDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("drop_blocked"); ok || d.HasChange("drop_blocked") {
		t, err := expandAntivirusQuarantineDropBlocked(d, v, "drop_blocked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-blocked"] = t
		}
	}

	if v, ok := d.GetOk("drop_heuristic"); ok || d.HasChange("drop_heuristic") {
		t, err := expandAntivirusQuarantineDropHeuristic(d, v, "drop_heuristic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-heuristic"] = t
		}
	}

	if v, ok := d.GetOk("drop_infected"); ok || d.HasChange("drop_infected") {
		t, err := expandAntivirusQuarantineDropInfected(d, v, "drop_infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-infected"] = t
		}
	}

	if v, ok := d.GetOk("drop_intercepted"); ok || d.HasChange("drop_intercepted") {
		t, err := expandAntivirusQuarantineDropIntercepted(d, v, "drop_intercepted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-intercepted"] = t
		}
	}

	if v, ok := d.GetOk("drop_machine_learning"); ok || d.HasChange("drop_machine_learning") {
		t, err := expandAntivirusQuarantineDropMachineLearning(d, v, "drop_machine_learning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-machine-learning"] = t
		}
	}

	if v, ok := d.GetOk("lowspace"); ok || d.HasChange("lowspace") {
		t, err := expandAntivirusQuarantineLowspace(d, v, "lowspace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lowspace"] = t
		}
	}

	if v, ok := d.GetOk("maxfilesize"); ok || d.HasChange("maxfilesize") {
		t, err := expandAntivirusQuarantineMaxfilesize(d, v, "maxfilesize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maxfilesize"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_quota"); ok || d.HasChange("quarantine_quota") {
		t, err := expandAntivirusQuarantineQuarantineQuota(d, v, "quarantine_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-quota"] = t
		}
	}

	if v, ok := d.GetOk("store_blocked"); ok || d.HasChange("store_blocked") {
		t, err := expandAntivirusQuarantineStoreBlocked(d, v, "store_blocked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-blocked"] = t
		}
	}

	if v, ok := d.GetOk("store_heuristic"); ok || d.HasChange("store_heuristic") {
		t, err := expandAntivirusQuarantineStoreHeuristic(d, v, "store_heuristic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-heuristic"] = t
		}
	}

	if v, ok := d.GetOk("store_infected"); ok || d.HasChange("store_infected") {
		t, err := expandAntivirusQuarantineStoreInfected(d, v, "store_infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-infected"] = t
		}
	}

	if v, ok := d.GetOk("store_intercepted"); ok || d.HasChange("store_intercepted") {
		t, err := expandAntivirusQuarantineStoreIntercepted(d, v, "store_intercepted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-intercepted"] = t
		}
	}

	if v, ok := d.GetOk("store_machine_learning"); ok || d.HasChange("store_machine_learning") {
		t, err := expandAntivirusQuarantineStoreMachineLearning(d, v, "store_machine_learning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-machine-learning"] = t
		}
	}

	return &obj, nil
}
