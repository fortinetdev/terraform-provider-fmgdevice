// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch IGMP snooping static group settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIgmpSnoopingStaticGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIgmpSnoopingStaticGroupCreate,
		Read:   resourceSwitchControllerIgmpSnoopingStaticGroupRead,
		Update: resourceSwitchControllerIgmpSnoopingStaticGroupUpdate,
		Delete: resourceSwitchControllerIgmpSnoopingStaticGroupDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mcast_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerIgmpSnoopingStaticGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerIgmpSnoopingStaticGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerIgmpSnoopingStaticGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSwitchControllerIgmpSnoopingStaticGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerIgmpSnoopingStaticGroupRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSwitchControllerIgmpSnoopingStaticGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerIgmpSnoopingStaticGroupRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerIgmpSnoopingStaticGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIgmpSnoopingStaticGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerIgmpSnoopingStaticGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIgmpSnoopingStaticGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroup resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupMcastAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSwitchControllerIgmpSnoopingStaticGroupPortsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerIgmpSnoopingStaticGroup-Ports-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ports"
		if _, ok := i["ports"]; ok {
			v := flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts(i["ports"], d, pre_append)
			tmp["ports"] = fortiAPISubPartPatch(v, "SwitchControllerIgmpSnoopingStaticGroup-Ports-Ports")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {
			v := flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(i["switch-id"], d, pre_append)
			tmp["switch_id"] = fortiAPISubPartPatch(v, "SwitchControllerIgmpSnoopingStaticGroup-Ports-SwitchId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerIgmpSnoopingStaticGroupVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerIgmpSnoopingStaticGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenSwitchControllerIgmpSnoopingStaticGroupDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerIgmpSnoopingStaticGroup-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ignore_reports", flattenSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(o["ignore-reports"], d, "ignore_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-reports"], "SwitchControllerIgmpSnoopingStaticGroup-IgnoreReports"); ok {
			if err = d.Set("ignore_reports", vv); err != nil {
				return fmt.Errorf("Error reading ignore_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_reports: %v", err)
		}
	}

	if err = d.Set("mcast_addr", flattenSwitchControllerIgmpSnoopingStaticGroupMcastAddr(o["mcast-addr"], d, "mcast_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-addr"], "SwitchControllerIgmpSnoopingStaticGroup-McastAddr"); ok {
			if err = d.Set("mcast_addr", vv); err != nil {
				return fmt.Errorf("Error reading mcast_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_addr: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerIgmpSnoopingStaticGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerIgmpSnoopingStaticGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ports", flattenSwitchControllerIgmpSnoopingStaticGroupPorts(o["ports"], d, "ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["ports"], "SwitchControllerIgmpSnoopingStaticGroup-Ports"); ok {
				if err = d.Set("ports", vv); err != nil {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenSwitchControllerIgmpSnoopingStaticGroupPorts(o["ports"], d, "ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["ports"], "SwitchControllerIgmpSnoopingStaticGroup-Ports"); ok {
					if err = d.Set("ports", vv); err != nil {
						return fmt.Errorf("Error reading ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerIgmpSnoopingStaticGroupVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerIgmpSnoopingStaticGroup-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerIgmpSnoopingStaticGroupDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupMcastAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ports"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts(d, i["ports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-id"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(d, i["switch_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerIgmpSnoopingStaticGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ignore_reports"); ok || d.HasChange("ignore_reports") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(d, v, "ignore_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-reports"] = t
		}
	}

	if v, ok := d.GetOk("mcast_addr"); ok || d.HasChange("mcast_addr") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupMcastAddr(d, v, "mcast_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-addr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupPorts(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
