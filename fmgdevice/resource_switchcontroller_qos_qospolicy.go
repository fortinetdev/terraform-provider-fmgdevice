// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerQosQosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosQosPolicyCreate,
		Read:   resourceSwitchControllerQosQosPolicyRead,
		Update: resourceSwitchControllerQosQosPolicyUpdate,
		Delete: resourceSwitchControllerQosQosPolicyDelete,

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
			"default_cos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"queue_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trust_dot1p_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trust_ip_dscp_map": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerQosQosPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerQosQosPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQosPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosQosPolicyRead(d, m)
}

func resourceSwitchControllerQosQosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQosQosPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerQosQosPolicyRead(d, m)
}

func resourceSwitchControllerQosQosPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerQosQosPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosQosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosQosPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerQosQosPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQosPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosQosPolicyDefaultCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyQueuePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosQosPolicyTrustDot1PMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerQosQosPolicyTrustIpDscpMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerQosQosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_cos", flattenSwitchControllerQosQosPolicyDefaultCos(o["default-cos"], d, "default_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-cos"], "SwitchControllerQosQosPolicy-DefaultCos"); ok {
			if err = d.Set("default_cos", vv); err != nil {
				return fmt.Errorf("Error reading default_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_cos: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerQosQosPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerQosQosPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("queue_policy", flattenSwitchControllerQosQosPolicyQueuePolicy(o["queue-policy"], d, "queue_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue-policy"], "SwitchControllerQosQosPolicy-QueuePolicy"); ok {
			if err = d.Set("queue_policy", vv); err != nil {
				return fmt.Errorf("Error reading queue_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue_policy: %v", err)
		}
	}

	if err = d.Set("trust_dot1p_map", flattenSwitchControllerQosQosPolicyTrustDot1PMap(o["trust-dot1p-map"], d, "trust_dot1p_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-dot1p-map"], "SwitchControllerQosQosPolicy-TrustDot1PMap"); ok {
			if err = d.Set("trust_dot1p_map", vv); err != nil {
				return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
		}
	}

	if err = d.Set("trust_ip_dscp_map", flattenSwitchControllerQosQosPolicyTrustIpDscpMap(o["trust-ip-dscp-map"], d, "trust_ip_dscp_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip-dscp-map"], "SwitchControllerQosQosPolicy-TrustIpDscpMap"); ok {
			if err = d.Set("trust_ip_dscp_map", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosQosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosQosPolicyDefaultCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyQueuePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosQosPolicyTrustDot1PMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerQosQosPolicyTrustIpDscpMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerQosQosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_cos"); ok || d.HasChange("default_cos") {
		t, err := expandSwitchControllerQosQosPolicyDefaultCos(d, v, "default_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cos"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerQosQosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("queue_policy"); ok || d.HasChange("queue_policy") {
		t, err := expandSwitchControllerQosQosPolicyQueuePolicy(d, v, "queue_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-policy"] = t
		}
	}

	if v, ok := d.GetOk("trust_dot1p_map"); ok || d.HasChange("trust_dot1p_map") {
		t, err := expandSwitchControllerQosQosPolicyTrustDot1PMap(d, v, "trust_dot1p_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-dot1p-map"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_dscp_map"); ok || d.HasChange("trust_ip_dscp_map") {
		t, err := expandSwitchControllerQosQosPolicyTrustIpDscpMap(d, v, "trust_ip_dscp_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-dscp-map"] = t
		}
	}

	return &obj, nil
}
