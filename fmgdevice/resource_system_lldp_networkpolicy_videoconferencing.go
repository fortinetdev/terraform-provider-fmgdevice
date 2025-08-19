// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Video Conferencing.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLldpNetworkPolicyVideoConferencing() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLldpNetworkPolicyVideoConferencingUpdate,
		Read:   resourceSystemLldpNetworkPolicyVideoConferencingRead,
		Update: resourceSystemLldpNetworkPolicyVideoConferencingUpdate,
		Delete: resourceSystemLldpNetworkPolicyVideoConferencingDelete,

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

func resourceSystemLldpNetworkPolicyVideoConferencingUpdate(d *schema.ResourceData, m interface{}) error {
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
	network_policy := d.Get("network_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["network_policy"] = network_policy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemLldpNetworkPolicyVideoConferencing(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyVideoConferencing resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLldpNetworkPolicyVideoConferencing(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyVideoConferencing resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLldpNetworkPolicyVideoConferencing")

	return resourceSystemLldpNetworkPolicyVideoConferencingRead(d, m)
}

func resourceSystemLldpNetworkPolicyVideoConferencingDelete(d *schema.ResourceData, m interface{}) error {
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
	network_policy := d.Get("network_policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["network_policy"] = network_policy

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemLldpNetworkPolicyVideoConferencing(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLldpNetworkPolicyVideoConferencing resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyVideoConferencingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemLldpNetworkPolicyVideoConferencing(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyVideoConferencing resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLldpNetworkPolicyVideoConferencing(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyVideoConferencing resource from API: %v", err)
	}
	return nil
}

func flattenSystemLldpNetworkPolicyVideoConferencingDscp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dscp", flattenSystemLldpNetworkPolicyVideoConferencingDscp2edl(o["dscp"], d, "dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp"], "SystemLldpNetworkPolicyVideoConferencing-Dscp"); ok {
			if err = d.Set("dscp", vv); err != nil {
				return fmt.Errorf("Error reading dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemLldpNetworkPolicyVideoConferencingPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemLldpNetworkPolicyVideoConferencing-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLldpNetworkPolicyVideoConferencingStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLldpNetworkPolicyVideoConferencing-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenSystemLldpNetworkPolicyVideoConferencingTag2edl(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "SystemLldpNetworkPolicyVideoConferencing-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSystemLldpNetworkPolicyVideoConferencingVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SystemLldpNetworkPolicyVideoConferencing-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemLldpNetworkPolicyVideoConferencingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLldpNetworkPolicyVideoConferencingDscp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp"); ok || d.HasChange("dscp") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencingDscp2edl(d, v, "dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencingPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencingStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencingTag2edl(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencingVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
