// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Define shaping entries of this shaping profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShapingProfileShapingEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingProfileShapingEntriesCreate,
		Read:   resourceFirewallShapingProfileShapingEntriesRead,
		Update: resourceFirewallShapingProfileShapingEntriesUpdate,
		Delete: resourceFirewallShapingProfileShapingEntriesDelete,

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
			"shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"burst_in_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cburst_in_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"guaranteed_bandwidth_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_bandwidth_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"red_probability": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallShapingProfileShapingEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	shaping_profile := d.Get("shaping_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectFirewallShapingProfileShapingEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingProfileShapingEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallShapingProfileShapingEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallShapingProfileShapingEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallShapingProfileShapingEntries resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallShapingProfileShapingEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallShapingProfileShapingEntries resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallShapingProfileShapingEntriesRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallShapingProfileShapingEntries resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallShapingProfileShapingEntriesRead(d, m)
}

func resourceFirewallShapingProfileShapingEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	shaping_profile := d.Get("shaping_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectFirewallShapingProfileShapingEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfileShapingEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallShapingProfileShapingEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfileShapingEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallShapingProfileShapingEntriesRead(d, m)
}

func resourceFirewallShapingProfileShapingEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	shaping_profile := d.Get("shaping_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile

	wsParams["adom"] = adomv

	err = c.DeleteFirewallShapingProfileShapingEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShapingProfileShapingEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingProfileShapingEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	shaping_profile := d.Get("shaping_profile").(string)
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
	if shaping_profile == "" {
		shaping_profile = importOptionChecking(m.(*FortiClient).Cfg, "shaping_profile")
		if shaping_profile == "" {
			return fmt.Errorf("Parameter shaping_profile is missing")
		}
		if err = d.Set("shaping_profile", shaping_profile); err != nil {
			return fmt.Errorf("Error set params shaping_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile

	o, err := c.ReadFirewallShapingProfileShapingEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallShapingProfileShapingEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShapingProfileShapingEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfileShapingEntries resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShapingProfileShapingEntriesBurstInMsec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesCburstInMsec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMax2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesRedProbability2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShapingProfileShapingEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("burst_in_msec", flattenFirewallShapingProfileShapingEntriesBurstInMsec2edl(o["burst-in-msec"], d, "burst_in_msec")); err != nil {
		if vv, ok := fortiAPIPatch(o["burst-in-msec"], "FirewallShapingProfileShapingEntries-BurstInMsec"); ok {
			if err = d.Set("burst_in_msec", vv); err != nil {
				return fmt.Errorf("Error reading burst_in_msec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading burst_in_msec: %v", err)
		}
	}

	if err = d.Set("cburst_in_msec", flattenFirewallShapingProfileShapingEntriesCburstInMsec2edl(o["cburst-in-msec"], d, "cburst_in_msec")); err != nil {
		if vv, ok := fortiAPIPatch(o["cburst-in-msec"], "FirewallShapingProfileShapingEntries-CburstInMsec"); ok {
			if err = d.Set("cburst_in_msec", vv); err != nil {
				return fmt.Errorf("Error reading cburst_in_msec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cburst_in_msec: %v", err)
		}
	}

	if err = d.Set("class_id", flattenFirewallShapingProfileShapingEntriesClassId2edl(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "FirewallShapingProfileShapingEntries-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth_percentage", flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(o["guaranteed-bandwidth-percentage"], d, "guaranteed_bandwidth_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth-percentage"], "FirewallShapingProfileShapingEntries-GuaranteedBandwidthPercentage"); ok {
			if err = d.Set("guaranteed_bandwidth_percentage", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth_percentage: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallShapingProfileShapingEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallShapingProfileShapingEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("limit", flattenFirewallShapingProfileShapingEntriesLimit2edl(o["limit"], d, "limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit"], "FirewallShapingProfileShapingEntries-Limit"); ok {
			if err = d.Set("limit", vv); err != nil {
				return fmt.Errorf("Error reading limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit: %v", err)
		}
	}

	if err = d.Set("max", flattenFirewallShapingProfileShapingEntriesMax2edl(o["max"], d, "max")); err != nil {
		if vv, ok := fortiAPIPatch(o["max"], "FirewallShapingProfileShapingEntries-Max"); ok {
			if err = d.Set("max", vv); err != nil {
				return fmt.Errorf("Error reading max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth_percentage", flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(o["maximum-bandwidth-percentage"], d, "maximum_bandwidth_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-bandwidth-percentage"], "FirewallShapingProfileShapingEntries-MaximumBandwidthPercentage"); ok {
			if err = d.Set("maximum_bandwidth_percentage", vv); err != nil {
				return fmt.Errorf("Error reading maximum_bandwidth_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_bandwidth_percentage: %v", err)
		}
	}

	if err = d.Set("min", flattenFirewallShapingProfileShapingEntriesMin2edl(o["min"], d, "min")); err != nil {
		if vv, ok := fortiAPIPatch(o["min"], "FirewallShapingProfileShapingEntries-Min"); ok {
			if err = d.Set("min", vv); err != nil {
				return fmt.Errorf("Error reading min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallShapingProfileShapingEntriesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "FirewallShapingProfileShapingEntries-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("red_probability", flattenFirewallShapingProfileShapingEntriesRedProbability2edl(o["red-probability"], d, "red_probability")); err != nil {
		if vv, ok := fortiAPIPatch(o["red-probability"], "FirewallShapingProfileShapingEntries-RedProbability"); ok {
			if err = d.Set("red_probability", vv); err != nil {
				return fmt.Errorf("Error reading red_probability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading red_probability: %v", err)
		}
	}

	return nil
}

func flattenFirewallShapingProfileShapingEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShapingProfileShapingEntriesBurstInMsec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesCburstInMsec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMax2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesRedProbability2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingProfileShapingEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("burst_in_msec"); ok || d.HasChange("burst_in_msec") {
		t, err := expandFirewallShapingProfileShapingEntriesBurstInMsec2edl(d, v, "burst_in_msec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst-in-msec"] = t
		}
	}

	if v, ok := d.GetOk("cburst_in_msec"); ok || d.HasChange("cburst_in_msec") {
		t, err := expandFirewallShapingProfileShapingEntriesCburstInMsec2edl(d, v, "cburst_in_msec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cburst-in-msec"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandFirewallShapingProfileShapingEntriesClassId2edl(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth_percentage"); ok || d.HasChange("guaranteed_bandwidth_percentage") {
		t, err := expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(d, v, "guaranteed_bandwidth_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth-percentage"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallShapingProfileShapingEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("limit"); ok || d.HasChange("limit") {
		t, err := expandFirewallShapingProfileShapingEntriesLimit2edl(d, v, "limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit"] = t
		}
	}

	if v, ok := d.GetOk("max"); ok || d.HasChange("max") {
		t, err := expandFirewallShapingProfileShapingEntriesMax2edl(d, v, "max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max"] = t
		}
	}

	if v, ok := d.GetOk("maximum_bandwidth_percentage"); ok || d.HasChange("maximum_bandwidth_percentage") {
		t, err := expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(d, v, "maximum_bandwidth_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth-percentage"] = t
		}
	}

	if v, ok := d.GetOk("min"); ok || d.HasChange("min") {
		t, err := expandFirewallShapingProfileShapingEntriesMin2edl(d, v, "min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandFirewallShapingProfileShapingEntriesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("red_probability"); ok || d.HasChange("red_probability") {
		t, err := expandFirewallShapingProfileShapingEntriesRedProbability2edl(d, v, "red_probability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["red-probability"] = t
		}
	}

	return &obj, nil
}
