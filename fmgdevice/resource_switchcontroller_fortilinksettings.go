// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure integrated FortiLink settings for FortiSwitch.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerFortilinkSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerFortilinkSettingsCreate,
		Read:   resourceSwitchControllerFortilinkSettingsRead,
		Update: resourceSwitchControllerFortilinkSettingsUpdate,
		Delete: resourceSwitchControllerFortilinkSettingsDelete,

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
			"access_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"inactive_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_down_flush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bounce_nac_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lan_segment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"member_change": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"nac_lan_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"nac_segment_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"onboarding_vlan": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"parent_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerFortilinkSettingsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerFortilinkSettings(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerFortilinkSettings(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceSwitchControllerFortilinkSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerFortilinkSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerFortilinkSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceSwitchControllerFortilinkSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerFortilinkSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFortilinkSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerFortilinkSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerFortilinkSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerFortilinkSettingsAccessVlanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsInactiveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsLinkDownFlush(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := i["bounce-nac-port"]; ok {
		result["bounce_nac_port"] = flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort(i["bounce-nac-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_segment"
	if _, ok := i["lan-segment"]; ok {
		result["lan_segment"] = flattenSwitchControllerFortilinkSettingsNacPortsLanSegment(i["lan-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "member_change"
	if _, ok := i["member-change"]; ok {
		result["member_change"] = flattenSwitchControllerFortilinkSettingsNacPortsMemberChange(i["member-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := i["nac-lan-interface"]; ok {
		result["nac_lan_interface"] = flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface(i["nac-lan-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := i["nac-segment-vlans"]; ok {
		result["nac_segment_vlans"] = flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(i["nac-segment-vlans"], d, pre_append)
	}

	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := i["onboarding-vlan"]; ok {
		result["onboarding_vlan"] = flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(i["onboarding-vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {
		result["parent_key"] = flattenSwitchControllerFortilinkSettingsNacPortsParentKey(i["parent-key"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsLanSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsMemberChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsParentKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_vlan_mode", flattenSwitchControllerFortilinkSettingsAccessVlanMode(o["access-vlan-mode"], d, "access_vlan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-vlan-mode"], "SwitchControllerFortilinkSettings-AccessVlanMode"); ok {
			if err = d.Set("access_vlan_mode", vv); err != nil {
				return fmt.Errorf("Error reading access_vlan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_vlan_mode: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerFortilinkSettingsFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "SwitchControllerFortilinkSettings-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("inactive_timer", flattenSwitchControllerFortilinkSettingsInactiveTimer(o["inactive-timer"], d, "inactive_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["inactive-timer"], "SwitchControllerFortilinkSettings-InactiveTimer"); ok {
			if err = d.Set("inactive_timer", vv); err != nil {
				return fmt.Errorf("Error reading inactive_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inactive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_flush", flattenSwitchControllerFortilinkSettingsLinkDownFlush(o["link-down-flush"], d, "link_down_flush")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-flush"], "SwitchControllerFortilinkSettings-LinkDownFlush"); ok {
			if err = d.Set("link_down_flush", vv); err != nil {
				return fmt.Errorf("Error reading link_down_flush: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_flush: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nac_ports", flattenSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-ports"], "SwitchControllerFortilinkSettings-NacPorts"); ok {
				if err = d.Set("nac_ports", vv); err != nil {
					return fmt.Errorf("Error reading nac_ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_ports"); ok {
			if err = d.Set("nac_ports", flattenSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-ports"], "SwitchControllerFortilinkSettings-NacPorts"); ok {
					if err = d.Set("nac_ports", vv); err != nil {
						return fmt.Errorf("Error reading nac_ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSwitchControllerFortilinkSettingsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerFortilinkSettings-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerFortilinkSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerFortilinkSettingsAccessVlanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsInactiveTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsLinkDownFlush(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bounce-nac-port"], _ = expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d, i["bounce_nac_port"], pre_append)
	}
	pre_append = pre + ".0." + "lan_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-segment"], _ = expandSwitchControllerFortilinkSettingsNacPortsLanSegment(d, i["lan_segment"], pre_append)
	}
	pre_append = pre + ".0." + "member_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["member-change"], _ = expandSwitchControllerFortilinkSettingsNacPortsMemberChange(d, i["member_change"], pre_append)
	}
	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nac-lan-interface"], _ = expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d, i["nac_lan_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nac-segment-vlans"], _ = expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d, i["nac_segment_vlans"], pre_append)
	}
	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["onboarding-vlan"], _ = expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d, i["onboarding_vlan"], pre_append)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["parent-key"], _ = expandSwitchControllerFortilinkSettingsNacPortsParentKey(d, i["parent_key"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsLanSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsMemberChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsParentKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerFortilinkSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_vlan_mode"); ok || d.HasChange("access_vlan_mode") {
		t, err := expandSwitchControllerFortilinkSettingsAccessVlanMode(d, v, "access_vlan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-vlan-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandSwitchControllerFortilinkSettingsFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("inactive_timer"); ok || d.HasChange("inactive_timer") {
		t, err := expandSwitchControllerFortilinkSettingsInactiveTimer(d, v, "inactive_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inactive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_flush"); ok || d.HasChange("link_down_flush") {
		t, err := expandSwitchControllerFortilinkSettingsLinkDownFlush(d, v, "link_down_flush")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-flush"] = t
		}
	}

	if v, ok := d.GetOk("nac_ports"); ok || d.HasChange("nac_ports") {
		t, err := expandSwitchControllerFortilinkSettingsNacPorts(d, v, "nac_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-ports"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerFortilinkSettingsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
