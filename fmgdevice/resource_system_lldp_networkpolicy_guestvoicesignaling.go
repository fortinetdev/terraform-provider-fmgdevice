// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Guest Voice Signaling.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLldpNetworkPolicyGuestVoiceSignaling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLldpNetworkPolicyGuestVoiceSignalingUpdate,
		Read:   resourceSystemLldpNetworkPolicyGuestVoiceSignalingRead,
		Update: resourceSystemLldpNetworkPolicyGuestVoiceSignalingUpdate,
		Delete: resourceSystemLldpNetworkPolicyGuestVoiceSignalingDelete,

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
			"network_policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemLldpNetworkPolicyGuestVoiceSignalingUpdate(d *schema.ResourceData, m interface{}) error {
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
	network_policy := d.Get("network_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["network_policy"] = network_policy

	obj, err := getObjectSystemLldpNetworkPolicyGuestVoiceSignaling(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyGuestVoiceSignaling resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLldpNetworkPolicyGuestVoiceSignaling(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyGuestVoiceSignaling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLldpNetworkPolicyGuestVoiceSignaling")

	return resourceSystemLldpNetworkPolicyGuestVoiceSignalingRead(d, m)
}

func resourceSystemLldpNetworkPolicyGuestVoiceSignalingDelete(d *schema.ResourceData, m interface{}) error {
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
	network_policy := d.Get("network_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["network_policy"] = network_policy

	err = c.DeleteSystemLldpNetworkPolicyGuestVoiceSignaling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLldpNetworkPolicyGuestVoiceSignaling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyGuestVoiceSignalingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	network_policy := d.Get("network_policy").(string)
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
	if network_policy == "" {
		network_policy = importOptionChecking(m.(*FortiClient).Cfg, "network_policy")
		if network_policy == "" {
			return fmt.Errorf("Parameter network_policy is missing")
		}
		if err = d.Set("network_policy", network_policy); err != nil {
			return fmt.Errorf("Error set params network_policy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["network_policy"] = network_policy

	o, err := c.ReadSystemLldpNetworkPolicyGuestVoiceSignaling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyGuestVoiceSignaling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLldpNetworkPolicyGuestVoiceSignaling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyGuestVoiceSignaling resource from API: %v", err)
	}
	return nil
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dscp", flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp2edl(o["dscp"], d, "dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp"], "SystemLldpNetworkPolicyGuestVoiceSignaling-Dscp"); ok {
			if err = d.Set("dscp", vv); err != nil {
				return fmt.Errorf("Error reading dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemLldpNetworkPolicyGuestVoiceSignaling-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLldpNetworkPolicyGuestVoiceSignaling-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag2edl(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "SystemLldpNetworkPolicyGuestVoiceSignaling-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SystemLldpNetworkPolicyGuestVoiceSignaling-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp"); ok || d.HasChange("dscp") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp2edl(d, v, "dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignalingTag2edl(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
