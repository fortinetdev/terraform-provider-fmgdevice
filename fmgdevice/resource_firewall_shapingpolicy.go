// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure shaping policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShapingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingPolicyCreate,
		Read:   resourceFirewallShapingPolicyRead,
		Update: resourceFirewallShapingPolicyUpdate,
		Delete: resourceFirewallShapingPolicyDelete,

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
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class_id_reverse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_response_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallShapingPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallShapingPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallShapingPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallShapingPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallShapingPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallShapingPolicy resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallShapingPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallShapingPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallShapingPolicyRead(d, m)
}

func resourceFirewallShapingPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallShapingPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallShapingPolicyRead(d, m)
}

func resourceFirewallShapingPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallShapingPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShapingPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallShapingPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallShapingPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShapingPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShapingPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallShapingPolicyClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyCosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceSrcFortiguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyTrafficType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallShapingPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyClassIdReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyHttpResponseMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingPolicyServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShapingPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("app_category", flattenFirewallShapingPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "FirewallShapingPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenFirewallShapingPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "FirewallShapingPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenFirewallShapingPolicyApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "FirewallShapingPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("class_id", flattenFirewallShapingPolicyClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "FirewallShapingPolicy-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallShapingPolicyComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallShapingPolicy-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cos", flattenFirewallShapingPolicyCos(o["cos"], d, "cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos"], "FirewallShapingPolicy-Cos"); ok {
			if err = d.Set("cos", vv); err != nil {
				return fmt.Errorf("Error reading cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos: %v", err)
		}
	}

	if err = d.Set("cos_mask", flattenFirewallShapingPolicyCosMask(o["cos-mask"], d, "cos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos-mask"], "FirewallShapingPolicy-CosMask"); ok {
			if err = d.Set("cos_mask", vv); err != nil {
				return fmt.Errorf("Error reading cos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos_mask: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallShapingPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "FirewallShapingPolicy-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallShapingPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "FirewallShapingPolicy-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallShapingPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "FirewallShapingPolicy-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallShapingPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "FirewallShapingPolicy-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallShapingPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallShapingPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallShapingPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallShapingPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenFirewallShapingPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "FirewallShapingPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("groups", flattenFirewallShapingPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallShapingPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallShapingPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallShapingPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallShapingPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "FirewallShapingPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenFirewallShapingPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "FirewallShapingPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenFirewallShapingPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "FirewallShapingPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_fortiguard", flattenFirewallShapingPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-fortiguard"], "FirewallShapingPolicy-InternetServiceFortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenFirewallShapingPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "FirewallShapingPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenFirewallShapingPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "FirewallShapingPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallShapingPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "FirewallShapingPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenFirewallShapingPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "FirewallShapingPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenFirewallShapingPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "FirewallShapingPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_fortiguard", flattenFirewallShapingPolicyInternetServiceSrcFortiguard(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-fortiguard"], "FirewallShapingPolicy-InternetServiceSrcFortiguard"); ok {
			if err = d.Set("internet_service_src_fortiguard", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenFirewallShapingPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "FirewallShapingPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenFirewallShapingPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "FirewallShapingPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenFirewallShapingPolicyIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "FirewallShapingPolicy-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallShapingPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallShapingPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallShapingPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "FirewallShapingPolicy-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallShapingPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallShapingPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenFirewallShapingPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "FirewallShapingPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallShapingPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallShapingPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenFirewallShapingPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "FirewallShapingPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallShapingPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "FirewallShapingPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallShapingPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallShapingPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tos", flattenFirewallShapingPolicyTos(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "FirewallShapingPolicy-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenFirewallShapingPolicyTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "FirewallShapingPolicy-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenFirewallShapingPolicyTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-negate"], "FirewallShapingPolicy-TosNegate"); ok {
			if err = d.Set("tos_negate", vv); err != nil {
				return fmt.Errorf("Error reading tos_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallShapingPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "FirewallShapingPolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallShapingPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "FirewallShapingPolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("traffic_type", flattenFirewallShapingPolicyTrafficType(o["traffic-type"], d, "traffic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-type"], "FirewallShapingPolicy-TrafficType"); ok {
			if err = d.Set("traffic_type", vv); err != nil {
				return fmt.Errorf("Error reading traffic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_type: %v", err)
		}
	}

	if err = d.Set("url_category", flattenFirewallShapingPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "FirewallShapingPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenFirewallShapingPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "FirewallShapingPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallShapingPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallShapingPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("class_id_reverse", flattenFirewallShapingPolicyClassIdReverse(o["class-id-reverse"], d, "class_id_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id-reverse"], "FirewallShapingPolicy-ClassIdReverse"); ok {
			if err = d.Set("class_id_reverse", vv); err != nil {
				return fmt.Errorf("Error reading class_id_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id_reverse: %v", err)
		}
	}

	if err = d.Set("http_response_match", flattenFirewallShapingPolicyHttpResponseMatch(o["http-response-match"], d, "http_response_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-response-match"], "FirewallShapingPolicy-HttpResponseMatch"); ok {
			if err = d.Set("http_response_match", vv); err != nil {
				return fmt.Errorf("Error reading http_response_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_response_match: %v", err)
		}
	}

	if err = d.Set("service_type", flattenFirewallShapingPolicyServiceType(o["service-type"], d, "service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-type"], "FirewallShapingPolicy-ServiceType"); ok {
			if err = d.Set("service_type", vv); err != nil {
				return fmt.Errorf("Error reading service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	return nil
}

func flattenFirewallShapingPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShapingPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyCosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceSrcFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyTosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyTrafficType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallShapingPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyClassIdReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyHttpResponseMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingPolicyServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallShapingPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandFirewallShapingPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallShapingPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandFirewallShapingPolicyClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallShapingPolicyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cos"); ok || d.HasChange("cos") {
		t, err := expandFirewallShapingPolicyCos(d, v, "cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos"] = t
		}
	}

	if v, ok := d.GetOk("cos_mask"); ok || d.HasChange("cos_mask") {
		t, err := expandFirewallShapingPolicyCosMask(d, v, "cos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-mask"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandFirewallShapingPolicyDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandFirewallShapingPolicyDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandFirewallShapingPolicyDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandFirewallShapingPolicyDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallShapingPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallShapingPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallShapingPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallShapingPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallShapingPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandFirewallShapingPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallShapingPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallShapingPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandFirewallShapingPolicyInternetServiceFortiguard(d, v, "internet_service_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallShapingPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallShapingPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok || d.HasChange("internet_service_src") {
		t, err := expandFirewallShapingPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallShapingPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallShapingPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_fortiguard"); ok || d.HasChange("internet_service_src_fortiguard") {
		t, err := expandFirewallShapingPolicyInternetServiceSrcFortiguard(d, v, "internet_service_src_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallShapingPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallShapingPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandFirewallShapingPolicyIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallShapingPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandFirewallShapingPolicyPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallShapingPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallShapingPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallShapingPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallShapingPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallShapingPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallShapingPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandFirewallShapingPolicyTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandFirewallShapingPolicyTosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok || d.HasChange("tos_negate") {
		t, err := expandFirewallShapingPolicyTosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandFirewallShapingPolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandFirewallShapingPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("traffic_type"); ok || d.HasChange("traffic_type") {
		t, err := expandFirewallShapingPolicyTrafficType(d, v, "traffic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-type"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandFirewallShapingPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallShapingPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallShapingPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("class_id_reverse"); ok || d.HasChange("class_id_reverse") {
		t, err := expandFirewallShapingPolicyClassIdReverse(d, v, "class_id_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id-reverse"] = t
		}
	}

	if v, ok := d.GetOk("http_response_match"); ok || d.HasChange("http_response_match") {
		t, err := expandFirewallShapingPolicyHttpResponseMatch(d, v, "http_response_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-response-match"] = t
		}
	}

	if v, ok := d.GetOk("service_type"); ok || d.HasChange("service_type") {
		t, err := expandFirewallShapingPolicyServiceType(d, v, "service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-type"] = t
		}
	}

	return &obj, nil
}
