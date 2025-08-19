// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port policies with matching criteria and actions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerDynamicPortPolicyPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerDynamicPortPolicyPolicyCreate,
		Read:   resourceSwitchControllerDynamicPortPolicyPolicyRead,
		Update: resourceSwitchControllerDynamicPortPolicyPolicyUpdate,
		Delete: resourceSwitchControllerDynamicPortPolicyPolicyDelete,

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
			"dynamic_port_policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"n802_1x": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bounce_port_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bounce_port_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lldp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"poe_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_policy": &schema.Schema{
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerDynamicPortPolicyPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dynamic_port_policy"] = dynamic_port_policy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerDynamicPortPolicyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicyPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerDynamicPortPolicyPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerDynamicPortPolicyPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dynamic_port_policy"] = dynamic_port_policy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerDynamicPortPolicyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicyPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerDynamicPortPolicyPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerDynamicPortPolicyPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dynamic_port_policy"] = dynamic_port_policy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerDynamicPortPolicyPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerDynamicPortPolicyPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
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
	if dynamic_port_policy == "" {
		dynamic_port_policy = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_port_policy")
		if dynamic_port_policy == "" {
			return fmt.Errorf("Parameter dynamic_port_policy is missing")
		}
		if err = d.Set("dynamic_port_policy", dynamic_port_policy); err != nil {
			return fmt.Errorf("Error set params dynamic_port_policy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["dynamic_port_policy"] = dynamic_port_policy

	o, err := c.ReadSwitchControllerDynamicPortPolicyPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerDynamicPortPolicyPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicyPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerDynamicPortPolicyPolicy8021X2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyBouncePortDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyFamily2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyMatchPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyMatchType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyPoeReset2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerDynamicPortPolicyPolicyStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n802_1x", flattenSwitchControllerDynamicPortPolicyPolicy8021X2edl(o["802-1x"], d, "n802_1x")); err != nil {
		if vv, ok := fortiAPIPatch(o["802-1x"], "SwitchControllerDynamicPortPolicyPolicy-8021X"); ok {
			if err = d.Set("n802_1x", vv); err != nil {
				return fmt.Errorf("Error reading n802_1x: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n802_1x: %v", err)
		}
	}

	if err = d.Set("bounce_port_duration", flattenSwitchControllerDynamicPortPolicyPolicyBouncePortDuration2edl(o["bounce-port-duration"], d, "bounce_port_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-port-duration"], "SwitchControllerDynamicPortPolicyPolicy-BouncePortDuration"); ok {
			if err = d.Set("bounce_port_duration", vv); err != nil {
				return fmt.Errorf("Error reading bounce_port_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_port_duration: %v", err)
		}
	}

	if err = d.Set("bounce_port_link", flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(o["bounce-port-link"], d, "bounce_port_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-port-link"], "SwitchControllerDynamicPortPolicyPolicy-BouncePortLink"); ok {
			if err = d.Set("bounce_port_link", vv); err != nil {
				return fmt.Errorf("Error reading bounce_port_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_port_link: %v", err)
		}
	}

	if err = d.Set("category", flattenSwitchControllerDynamicPortPolicyPolicyCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "SwitchControllerDynamicPortPolicyPolicy-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerDynamicPortPolicyPolicyDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerDynamicPortPolicyPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("family", flattenSwitchControllerDynamicPortPolicyPolicyFamily2edl(o["family"], d, "family")); err != nil {
		if vv, ok := fortiAPIPatch(o["family"], "SwitchControllerDynamicPortPolicyPolicy-Family"); ok {
			if err = d.Set("family", vv); err != nil {
				return fmt.Errorf("Error reading family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("host", flattenSwitchControllerDynamicPortPolicyPolicyHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "SwitchControllerDynamicPortPolicyPolicy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "SwitchControllerDynamicPortPolicyPolicy-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("interface_tags", flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(o["interface-tags"], d, "interface_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-tags"], "SwitchControllerDynamicPortPolicyPolicy-InterfaceTags"); ok {
			if err = d.Set("interface_tags", vv); err != nil {
				return fmt.Errorf("Error reading interface_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_tags: %v", err)
		}
	}

	if err = d.Set("lldp_profile", flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(o["lldp-profile"], d, "lldp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-profile"], "SwitchControllerDynamicPortPolicyPolicy-LldpProfile"); ok {
			if err = d.Set("lldp_profile", vv); err != nil {
				return fmt.Errorf("Error reading lldp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_profile: %v", err)
		}
	}

	if err = d.Set("mac", flattenSwitchControllerDynamicPortPolicyPolicyMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "SwitchControllerDynamicPortPolicyPolicy-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("match_period", flattenSwitchControllerDynamicPortPolicyPolicyMatchPeriod2edl(o["match-period"], d, "match_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-period"], "SwitchControllerDynamicPortPolicyPolicy-MatchPeriod"); ok {
			if err = d.Set("match_period", vv); err != nil {
				return fmt.Errorf("Error reading match_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_period: %v", err)
		}
	}

	if err = d.Set("match_type", flattenSwitchControllerDynamicPortPolicyPolicyMatchType2edl(o["match-type"], d, "match_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-type"], "SwitchControllerDynamicPortPolicyPolicy-MatchType"); ok {
			if err = d.Set("match_type", vv); err != nil {
				return fmt.Errorf("Error reading match_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerDynamicPortPolicyPolicyName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerDynamicPortPolicyPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("poe_reset", flattenSwitchControllerDynamicPortPolicyPolicyPoeReset2edl(o["poe-reset"], d, "poe_reset")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-reset"], "SwitchControllerDynamicPortPolicyPolicy-PoeReset"); ok {
			if err = d.Set("poe_reset", vv); err != nil {
				return fmt.Errorf("Error reading poe_reset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_reset: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(o["qos-policy"], d, "qos_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-policy"], "SwitchControllerDynamicPortPolicyPolicy-QosPolicy"); ok {
			if err = d.Set("qos_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerDynamicPortPolicyPolicyStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SwitchControllerDynamicPortPolicyPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerDynamicPortPolicyPolicyType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SwitchControllerDynamicPortPolicyPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vlan_policy", flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(o["vlan-policy"], d, "vlan_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-policy"], "SwitchControllerDynamicPortPolicyPolicy-VlanPolicy"); ok {
			if err = d.Set("vlan_policy", vv); err != nil {
				return fmt.Errorf("Error reading vlan_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerDynamicPortPolicyPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerDynamicPortPolicyPolicy8021X2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyBouncePortDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyFamily2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMatchPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMatchType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyPoeReset2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerDynamicPortPolicyPolicyStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n802_1x"); ok || d.HasChange("n802_1x") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicy8021X2edl(d, v, "n802_1x")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802-1x"] = t
		}
	}

	if v, ok := d.GetOk("bounce_port_duration"); ok || d.HasChange("bounce_port_duration") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyBouncePortDuration2edl(d, v, "bounce_port_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-duration"] = t
		}
	}

	if v, ok := d.GetOk("bounce_port_link"); ok || d.HasChange("bounce_port_link") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(d, v, "bounce_port_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-link"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok || d.HasChange("family") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyFamily2edl(d, v, "family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("interface_tags"); ok || d.HasChange("interface_tags") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(d, v, "interface_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-tags"] = t
		}
	}

	if v, ok := d.GetOk("lldp_profile"); ok || d.HasChange("lldp_profile") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(d, v, "lldp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-profile"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("match_period"); ok || d.HasChange("match_period") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyMatchPeriod2edl(d, v, "match_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-period"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok || d.HasChange("match_type") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyMatchType2edl(d, v, "match_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("poe_reset"); ok || d.HasChange("poe_reset") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyPoeReset2edl(d, v, "poe_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-reset"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok || d.HasChange("qos_policy") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(d, v, "qos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vlan_policy"); ok || d.HasChange("vlan_policy") {
		t, err := expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(d, v, "vlan_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-policy"] = t
		}
	}

	return &obj, nil
}
