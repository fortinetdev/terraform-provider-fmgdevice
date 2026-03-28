// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallResponseShapingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallResponseShapingPolicyCreate,
		Read:   resourceFirewallResponseShapingPolicyRead,
		Update: resourceFirewallResponseShapingPolicyUpdate,
		Delete: resourceFirewallResponseShapingPolicyDelete,

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
			"class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"class_id_reverse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
			"srcaddr": &schema.Schema{
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallResponseShapingPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallResponseShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallResponseShapingPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallResponseShapingPolicy(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallResponseShapingPolicy(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallResponseShapingPolicy resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallResponseShapingPolicy(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallResponseShapingPolicy resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallResponseShapingPolicyRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallResponseShapingPolicy resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallResponseShapingPolicyRead(d, m)
}

func resourceFirewallResponseShapingPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallResponseShapingPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallResponseShapingPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallResponseShapingPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallResponseShapingPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallResponseShapingPolicyRead(d, m)
}

func resourceFirewallResponseShapingPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallResponseShapingPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallResponseShapingPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallResponseShapingPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallResponseShapingPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallResponseShapingPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallResponseShapingPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallResponseShapingPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallResponseShapingPolicyClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyClassIdReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallResponseShapingPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallResponseShapingPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallResponseShapingPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("class_id", flattenFirewallResponseShapingPolicyClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "FirewallResponseShapingPolicy-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("class_id_reverse", flattenFirewallResponseShapingPolicyClassIdReverse(o["class-id-reverse"], d, "class_id_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id-reverse"], "FirewallResponseShapingPolicy-ClassIdReverse"); ok {
			if err = d.Set("class_id_reverse", vv); err != nil {
				return fmt.Errorf("Error reading class_id_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id_reverse: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallResponseShapingPolicyComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "FirewallResponseShapingPolicy-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenFirewallResponseShapingPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "FirewallResponseShapingPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenFirewallResponseShapingPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "FirewallResponseShapingPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallResponseShapingPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallResponseShapingPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenFirewallResponseShapingPolicyIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "FirewallResponseShapingPolicy-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallResponseShapingPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallResponseShapingPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallResponseShapingPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "FirewallResponseShapingPolicy-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallResponseShapingPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "FirewallResponseShapingPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenFirewallResponseShapingPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "FirewallResponseShapingPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallResponseShapingPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallResponseShapingPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallResponseShapingPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "FirewallResponseShapingPolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallResponseShapingPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "FirewallResponseShapingPolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallResponseShapingPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallResponseShapingPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallResponseShapingPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallResponseShapingPolicyClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyClassIdReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallResponseShapingPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallResponseShapingPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallResponseShapingPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandFirewallResponseShapingPolicyClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("class_id_reverse"); ok || d.HasChange("class_id_reverse") {
		t, err := expandFirewallResponseShapingPolicyClassIdReverse(d, v, "class_id_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id-reverse"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandFirewallResponseShapingPolicyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallResponseShapingPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallResponseShapingPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallResponseShapingPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandFirewallResponseShapingPolicyIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallResponseShapingPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandFirewallResponseShapingPolicyPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandFirewallResponseShapingPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallResponseShapingPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallResponseShapingPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandFirewallResponseShapingPolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandFirewallResponseShapingPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallResponseShapingPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
