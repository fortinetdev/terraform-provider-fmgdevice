// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Video Signaling.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLldpNetworkPolicyVideoSignaling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLldpNetworkPolicyVideoSignalingUpdate,
		Read:   resourceSystemLldpNetworkPolicyVideoSignalingRead,
		Update: resourceSystemLldpNetworkPolicyVideoSignalingUpdate,
		Delete: resourceSystemLldpNetworkPolicyVideoSignalingDelete,

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

func resourceSystemLldpNetworkPolicyVideoSignalingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemLldpNetworkPolicyVideoSignaling(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyVideoSignaling resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLldpNetworkPolicyVideoSignaling(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicyVideoSignaling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLldpNetworkPolicyVideoSignaling")

	return resourceSystemLldpNetworkPolicyVideoSignalingRead(d, m)
}

func resourceSystemLldpNetworkPolicyVideoSignalingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemLldpNetworkPolicyVideoSignaling(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLldpNetworkPolicyVideoSignaling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyVideoSignalingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemLldpNetworkPolicyVideoSignaling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyVideoSignaling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLldpNetworkPolicyVideoSignaling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicyVideoSignaling resource from API: %v", err)
	}
	return nil
}

func flattenSystemLldpNetworkPolicyVideoSignalingDscp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dscp", flattenSystemLldpNetworkPolicyVideoSignalingDscp2edl(o["dscp"], d, "dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp"], "SystemLldpNetworkPolicyVideoSignaling-Dscp"); ok {
			if err = d.Set("dscp", vv); err != nil {
				return fmt.Errorf("Error reading dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemLldpNetworkPolicyVideoSignalingPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemLldpNetworkPolicyVideoSignaling-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLldpNetworkPolicyVideoSignalingStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLldpNetworkPolicyVideoSignaling-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tag", flattenSystemLldpNetworkPolicyVideoSignalingTag2edl(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "SystemLldpNetworkPolicyVideoSignaling-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSystemLldpNetworkPolicyVideoSignalingVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SystemLldpNetworkPolicyVideoSignaling-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemLldpNetworkPolicyVideoSignalingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLldpNetworkPolicyVideoSignalingDscp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp"); ok || d.HasChange("dscp") {
		t, err := expandSystemLldpNetworkPolicyVideoSignalingDscp2edl(d, v, "dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemLldpNetworkPolicyVideoSignalingPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLldpNetworkPolicyVideoSignalingStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
		t, err := expandSystemLldpNetworkPolicyVideoSignalingTag2edl(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSystemLldpNetworkPolicyVideoSignalingVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
