// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NAC specific configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerFortilinkSettingsNacPorts() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerFortilinkSettingsNacPortsUpdate,
		Read:   resourceSwitchControllerFortilinkSettingsNacPortsRead,
		Update: resourceSwitchControllerFortilinkSettingsNacPortsUpdate,
		Delete: resourceSwitchControllerFortilinkSettingsNacPortsDelete,

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
			"fortilink_settings": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
	}
}

func resourceSwitchControllerFortilinkSettingsNacPortsUpdate(d *schema.ResourceData, m interface{}) error {
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
	fortilink_settings := d.Get("fortilink_settings").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortilink_settings"] = fortilink_settings

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerFortilinkSettingsNacPorts(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettingsNacPorts resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerFortilinkSettingsNacPorts(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerFortilinkSettingsNacPorts")

	return resourceSwitchControllerFortilinkSettingsNacPortsRead(d, m)
}

func resourceSwitchControllerFortilinkSettingsNacPortsDelete(d *schema.ResourceData, m interface{}) error {
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
	fortilink_settings := d.Get("fortilink_settings").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortilink_settings"] = fortilink_settings

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerFortilinkSettingsNacPorts(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFortilinkSettingsNacPortsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	fortilink_settings := d.Get("fortilink_settings").(string)
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
	if fortilink_settings == "" {
		fortilink_settings = importOptionChecking(m.(*FortiClient).Cfg, "fortilink_settings")
		if fortilink_settings == "" {
			return fmt.Errorf("Parameter fortilink_settings is missing")
		}
		if err = d.Set("fortilink_settings", fortilink_settings); err != nil {
			return fmt.Errorf("Error set params fortilink_settings: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["fortilink_settings"] = fortilink_settings

	o, err := c.ReadSwitchControllerFortilinkSettingsNacPorts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerFortilinkSettingsNacPorts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettingsNacPorts resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFortilinkSettingsNacPortsParentKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bounce_nac_port", flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(o["bounce-nac-port"], d, "bounce_nac_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-nac-port"], "SwitchControllerFortilinkSettingsNacPorts-BounceNacPort"); ok {
			if err = d.Set("bounce_nac_port", vv); err != nil {
				return fmt.Errorf("Error reading bounce_nac_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_nac_port: %v", err)
		}
	}

	if err = d.Set("lan_segment", flattenSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(o["lan-segment"], d, "lan_segment")); err != nil {
		if vv, ok := fortiAPIPatch(o["lan-segment"], "SwitchControllerFortilinkSettingsNacPorts-LanSegment"); ok {
			if err = d.Set("lan_segment", vv); err != nil {
				return fmt.Errorf("Error reading lan_segment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lan_segment: %v", err)
		}
	}

	if err = d.Set("member_change", flattenSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(o["member-change"], d, "member_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-change"], "SwitchControllerFortilinkSettingsNacPorts-MemberChange"); ok {
			if err = d.Set("member_change", vv); err != nil {
				return fmt.Errorf("Error reading member_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_change: %v", err)
		}
	}

	if err = d.Set("nac_lan_interface", flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(o["nac-lan-interface"], d, "nac_lan_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-lan-interface"], "SwitchControllerFortilinkSettingsNacPorts-NacLanInterface"); ok {
			if err = d.Set("nac_lan_interface", vv); err != nil {
				return fmt.Errorf("Error reading nac_lan_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_lan_interface: %v", err)
		}
	}

	if err = d.Set("nac_segment_vlans", flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(o["nac-segment-vlans"], d, "nac_segment_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-segment-vlans"], "SwitchControllerFortilinkSettingsNacPorts-NacSegmentVlans"); ok {
			if err = d.Set("nac_segment_vlans", vv); err != nil {
				return fmt.Errorf("Error reading nac_segment_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_segment_vlans: %v", err)
		}
	}

	if err = d.Set("onboarding_vlan", flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(o["onboarding-vlan"], d, "onboarding_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["onboarding-vlan"], "SwitchControllerFortilinkSettingsNacPorts-OnboardingVlan"); ok {
			if err = d.Set("onboarding_vlan", vv); err != nil {
				return fmt.Errorf("Error reading onboarding_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onboarding_vlan: %v", err)
		}
	}

	if err = d.Set("parent_key", flattenSwitchControllerFortilinkSettingsNacPortsParentKey2edl(o["parent-key"], d, "parent_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent-key"], "SwitchControllerFortilinkSettingsNacPorts-ParentKey"); ok {
			if err = d.Set("parent_key", vv); err != nil {
				return fmt.Errorf("Error reading parent_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent_key: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerFortilinkSettingsNacPortsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerFortilinkSettingsNacPortsParentKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bounce_nac_port"); ok || d.HasChange("bounce_nac_port") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(d, v, "bounce_nac_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-nac-port"] = t
		}
	}

	if v, ok := d.GetOk("lan_segment"); ok || d.HasChange("lan_segment") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(d, v, "lan_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-segment"] = t
		}
	}

	if v, ok := d.GetOk("member_change"); ok || d.HasChange("member_change") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(d, v, "member_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-change"] = t
		}
	}

	if v, ok := d.GetOk("nac_lan_interface"); ok || d.HasChange("nac_lan_interface") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(d, v, "nac_lan_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-lan-interface"] = t
		}
	}

	if v, ok := d.GetOk("nac_segment_vlans"); ok || d.HasChange("nac_segment_vlans") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(d, v, "nac_segment_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-segment-vlans"] = t
		}
	}

	if v, ok := d.GetOk("onboarding_vlan"); ok || d.HasChange("onboarding_vlan") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(d, v, "onboarding_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onboarding-vlan"] = t
		}
	}

	if v, ok := d.GetOk("parent_key"); ok || d.HasChange("parent_key") {
		t, err := expandSwitchControllerFortilinkSettingsNacPortsParentKey2edl(d, v, "parent_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent-key"] = t
		}
	}

	return &obj, nil
}
