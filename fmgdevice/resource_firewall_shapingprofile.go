// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure shaping profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShapingProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingProfileCreate,
		Read:   resourceFirewallShapingProfileRead,
		Update: resourceFirewallShapingProfileUpdate,
		Delete: resourceFirewallShapingProfileDelete,

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
			"default_class_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"npu_offloading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shaping_entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
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
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"classes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"guaranteed_bandwidth": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"maximum_bandwidth": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"default_class": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallShapingProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("profile_name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallShapingProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallShapingProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallShapingProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallShapingProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallShapingProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "profile_name"))

	return resourceFirewallShapingProfileRead(d, m)
}

func resourceFirewallShapingProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallShapingProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "profile_name"))

	return resourceFirewallShapingProfileRead(d, m)
}

func resourceFirewallShapingProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallShapingProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShapingProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallShapingProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallShapingProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShapingProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShapingProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileDefaultClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingProfileNpuOffloading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := i["burst-in-msec"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesBurstInMsec(i["burst-in-msec"], d, pre_append)
			tmp["burst_in_msec"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-BurstInMsec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := i["cburst-in-msec"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesCburstInMsec(i["cburst-in-msec"], d, pre_append)
			tmp["cburst_in_msec"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-CburstInMsec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := i["guaranteed-bandwidth-percentage"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(i["guaranteed-bandwidth-percentage"], d, pre_append)
			tmp["guaranteed_bandwidth_percentage"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-GuaranteedBandwidthPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := i["limit"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesLimit(i["limit"], d, pre_append)
			tmp["limit"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-Limit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := i["max"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesMax(i["max"], d, pre_append)
			tmp["max"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-Max")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := i["maximum-bandwidth-percentage"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(i["maximum-bandwidth-percentage"], d, pre_append)
			tmp["maximum_bandwidth_percentage"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-MaximumBandwidthPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := i["min"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesMin(i["min"], d, pre_append)
			tmp["min"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-Min")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := i["red-probability"]; ok {
			v := flattenFirewallShapingProfileShapingEntriesRedProbability(i["red-probability"], d, pre_append)
			tmp["red_probability"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-ShapingEntries-RedProbability")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallShapingProfileShapingEntriesBurstInMsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesCburstInMsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesRedProbability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileClasses(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenFirewallShapingProfileClassesClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-Classes-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth"
		if _, ok := i["guaranteed-bandwidth"]; ok {
			v := flattenFirewallShapingProfileClassesGuaranteedBandwidth(i["guaranteed-bandwidth"], d, pre_append)
			tmp["guaranteed_bandwidth"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-Classes-GuaranteedBandwidth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth"
		if _, ok := i["maximum-bandwidth"]; ok {
			v := flattenFirewallShapingProfileClassesMaximumBandwidth(i["maximum-bandwidth"], d, pre_append)
			tmp["maximum_bandwidth"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-Classes-MaximumBandwidth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenFirewallShapingProfileClassesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-Classes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenFirewallShapingProfileClassesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "FirewallShapingProfile-Classes-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallShapingProfileClassesClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileClassesGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileClassesMaximumBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileClassesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileClassesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileDefaultClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShapingProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenFirewallShapingProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallShapingProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("default_class_id", flattenFirewallShapingProfileDefaultClassId(o["default-class-id"], d, "default_class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-class-id"], "FirewallShapingProfile-DefaultClassId"); ok {
			if err = d.Set("default_class_id", vv); err != nil {
				return fmt.Errorf("Error reading default_class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_class_id: %v", err)
		}
	}

	if err = d.Set("npu_offloading", flattenFirewallShapingProfileNpuOffloading(o["npu-offloading"], d, "npu_offloading")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offloading"], "FirewallShapingProfile-NpuOffloading"); ok {
			if err = d.Set("npu_offloading", vv); err != nil {
				return fmt.Errorf("Error reading npu_offloading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offloading: %v", err)
		}
	}

	if err = d.Set("profile_name", flattenFirewallShapingProfileProfileName(o["profile-name"], d, "profile_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-name"], "FirewallShapingProfile-ProfileName"); ok {
			if err = d.Set("profile_name", vv); err != nil {
				return fmt.Errorf("Error reading profile_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["shaping-entries"], "FirewallShapingProfile-ShapingEntries"); ok {
				if err = d.Set("shaping_entries", vv); err != nil {
					return fmt.Errorf("Error reading shaping_entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading shaping_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shaping_entries"); ok {
			if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["shaping-entries"], "FirewallShapingProfile-ShapingEntries"); ok {
					if err = d.Set("shaping_entries", vv); err != nil {
						return fmt.Errorf("Error reading shaping_entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading shaping_entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenFirewallShapingProfileType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallShapingProfile-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("classes", flattenFirewallShapingProfileClasses(o["classes"], d, "classes")); err != nil {
			if vv, ok := fortiAPIPatch(o["classes"], "FirewallShapingProfile-Classes"); ok {
				if err = d.Set("classes", vv); err != nil {
					return fmt.Errorf("Error reading classes: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading classes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classes"); ok {
			if err = d.Set("classes", flattenFirewallShapingProfileClasses(o["classes"], d, "classes")); err != nil {
				if vv, ok := fortiAPIPatch(o["classes"], "FirewallShapingProfile-Classes"); ok {
					if err = d.Set("classes", vv); err != nil {
						return fmt.Errorf("Error reading classes: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading classes: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_class", flattenFirewallShapingProfileDefaultClass(o["default-class"], d, "default_class")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-class"], "FirewallShapingProfile-DefaultClass"); ok {
			if err = d.Set("default_class", vv); err != nil {
				return fmt.Errorf("Error reading default_class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_class: %v", err)
		}
	}

	return nil
}

func flattenFirewallShapingProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShapingProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileDefaultClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingProfileNpuOffloading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["burst-in-msec"], _ = expandFirewallShapingProfileShapingEntriesBurstInMsec(d, i["burst_in_msec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cburst-in-msec"], _ = expandFirewallShapingProfileShapingEntriesCburstInMsec(d, i["cburst_in_msec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandFirewallShapingProfileShapingEntriesClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["guaranteed-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d, i["guaranteed_bandwidth_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFirewallShapingProfileShapingEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["limit"], _ = expandFirewallShapingProfileShapingEntriesLimit(d, i["limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max"], _ = expandFirewallShapingProfileShapingEntriesMax(d, i["max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d, i["maximum_bandwidth_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min"], _ = expandFirewallShapingProfileShapingEntriesMin(d, i["min"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallShapingProfileShapingEntriesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["red-probability"], _ = expandFirewallShapingProfileShapingEntriesRedProbability(d, i["red_probability"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallShapingProfileShapingEntriesBurstInMsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesCburstInMsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesRedProbability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClasses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandFirewallShapingProfileClassesClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["guaranteed-bandwidth"], _ = expandFirewallShapingProfileClassesGuaranteedBandwidth(d, i["guaranteed_bandwidth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-bandwidth"], _ = expandFirewallShapingProfileClassesMaximumBandwidth(d, i["maximum_bandwidth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandFirewallShapingProfileClassesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandFirewallShapingProfileClassesPriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallShapingProfileClassesClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClassesGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClassesMaximumBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClassesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClassesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileDefaultClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallShapingProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("default_class_id"); ok || d.HasChange("default_class_id") {
		t, err := expandFirewallShapingProfileDefaultClassId(d, v, "default_class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-class-id"] = t
		}
	}

	if v, ok := d.GetOk("npu_offloading"); ok || d.HasChange("npu_offloading") {
		t, err := expandFirewallShapingProfileNpuOffloading(d, v, "npu_offloading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offloading"] = t
		}
	}

	if v, ok := d.GetOk("profile_name"); ok || d.HasChange("profile_name") {
		t, err := expandFirewallShapingProfileProfileName(d, v, "profile_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("shaping_entries"); ok || d.HasChange("shaping_entries") {
		t, err := expandFirewallShapingProfileShapingEntries(d, v, "shaping_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaping-entries"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallShapingProfileType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("classes"); ok || d.HasChange("classes") {
		t, err := expandFirewallShapingProfileClasses(d, v, "classes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classes"] = t
		}
	}

	if v, ok := d.GetOk("default_class"); ok || d.HasChange("default_class") {
		t, err := expandFirewallShapingProfileDefaultClass(d, v, "default_class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-class"] = t
		}
	}

	return &obj, nil
}
