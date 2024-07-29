// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SD-WAN zones.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwanZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanZoneCreate,
		Read:   resourceSystemSdwanZoneRead,
		Update: resourceSystemSdwanZoneUpdate,
		Delete: resourceSystemSdwanZoneDelete,

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
			"advpn_health_check": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"advpn_select": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minimum_sla_meet_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"service_sla_tie_break": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSdwanZoneCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanZone(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanZone resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdwanZone(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdwanZone resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdwanZoneRead(d, m)
}

func resourceSystemSdwanZoneUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdwanZone(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanZone resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwanZone(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwanZone resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdwanZoneRead(d, m)
}

func resourceSystemSdwanZoneDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdwanZone(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwanZone resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanZoneRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwanZone(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanZone resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwanZone(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwanZone resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanZoneAdvpnHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanZoneAdvpnSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneMinimumSlaMeetMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneServiceSlaTieBreak2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwanZone(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("advpn_health_check", flattenSystemSdwanZoneAdvpnHealthCheck2edl(o["advpn-health-check"], d, "advpn_health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["advpn-health-check"], "SystemSdwanZone-AdvpnHealthCheck"); ok {
			if err = d.Set("advpn_health_check", vv); err != nil {
				return fmt.Errorf("Error reading advpn_health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advpn_health_check: %v", err)
		}
	}

	if err = d.Set("advpn_select", flattenSystemSdwanZoneAdvpnSelect2edl(o["advpn-select"], d, "advpn_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["advpn-select"], "SystemSdwanZone-AdvpnSelect"); ok {
			if err = d.Set("advpn_select", vv); err != nil {
				return fmt.Errorf("Error reading advpn_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advpn_select: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", flattenSystemSdwanZoneMinimumSlaMeetMembers2edl(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-sla-meet-members"], "SystemSdwanZone-MinimumSlaMeetMembers"); ok {
			if err = d.Set("minimum_sla_meet_members", vv); err != nil {
				return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdwanZoneName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdwanZone-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service_sla_tie_break", flattenSystemSdwanZoneServiceSlaTieBreak2edl(o["service-sla-tie-break"], d, "service_sla_tie_break")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-sla-tie-break"], "SystemSdwanZone-ServiceSlaTieBreak"); ok {
			if err = d.Set("service_sla_tie_break", vv); err != nil {
				return fmt.Errorf("Error reading service_sla_tie_break: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_sla_tie_break: %v", err)
		}
	}

	return nil
}

func flattenSystemSdwanZoneFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanZoneAdvpnHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanZoneAdvpnSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneMinimumSlaMeetMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneServiceSlaTieBreak2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwanZone(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advpn_health_check"); ok || d.HasChange("advpn_health_check") {
		t, err := expandSystemSdwanZoneAdvpnHealthCheck2edl(d, v, "advpn_health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advpn-health-check"] = t
		}
	}

	if v, ok := d.GetOk("advpn_select"); ok || d.HasChange("advpn_select") {
		t, err := expandSystemSdwanZoneAdvpnSelect2edl(d, v, "advpn_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advpn-select"] = t
		}
	}

	if v, ok := d.GetOk("minimum_sla_meet_members"); ok || d.HasChange("minimum_sla_meet_members") {
		t, err := expandSystemSdwanZoneMinimumSlaMeetMembers2edl(d, v, "minimum_sla_meet_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-sla-meet-members"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdwanZoneName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_sla_tie_break"); ok || d.HasChange("service_sla_tie_break") {
		t, err := expandSystemSdwanZoneServiceSlaTieBreak2edl(d, v, "service_sla_tie_break")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-sla-tie-break"] = t
		}
	}

	return &obj, nil
}
