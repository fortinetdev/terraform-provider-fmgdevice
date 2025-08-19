// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Bonjour policy list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerBonjourProfilePolicyList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerBonjourProfilePolicyListCreate,
		Read:   resourceWirelessControllerBonjourProfilePolicyListRead,
		Update: resourceWirelessControllerBonjourProfilePolicyListUpdate,
		Delete: resourceWirelessControllerBonjourProfilePolicyListDelete,

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
			"bonjour_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"to_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerBonjourProfilePolicyListCreate(d *schema.ResourceData, m interface{}) error {
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
	bonjour_profile := d.Get("bonjour_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bonjour_profile"] = bonjour_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerBonjourProfilePolicyList(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfilePolicyList resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerBonjourProfilePolicyList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policy_id")))

	return resourceWirelessControllerBonjourProfilePolicyListRead(d, m)
}

func resourceWirelessControllerBonjourProfilePolicyListUpdate(d *schema.ResourceData, m interface{}) error {
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
	bonjour_profile := d.Get("bonjour_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bonjour_profile"] = bonjour_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectWirelessControllerBonjourProfilePolicyList(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfilePolicyList resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerBonjourProfilePolicyList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policy_id")))

	return resourceWirelessControllerBonjourProfilePolicyListRead(d, m)
}

func resourceWirelessControllerBonjourProfilePolicyListDelete(d *schema.ResourceData, m interface{}) error {
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
	bonjour_profile := d.Get("bonjour_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bonjour_profile"] = bonjour_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteWirelessControllerBonjourProfilePolicyList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBonjourProfilePolicyListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	bonjour_profile := d.Get("bonjour_profile").(string)
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
	if bonjour_profile == "" {
		bonjour_profile = importOptionChecking(m.(*FortiClient).Cfg, "bonjour_profile")
		if bonjour_profile == "" {
			return fmt.Errorf("Parameter bonjour_profile is missing")
		}
		if err = d.Set("bonjour_profile", bonjour_profile); err != nil {
			return fmt.Errorf("Error set params bonjour_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["bonjour_profile"] = bonjour_profile

	o, err := c.ReadWirelessControllerBonjourProfilePolicyList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBonjourProfilePolicyList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfilePolicyList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBonjourProfilePolicyListDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListFromVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListPolicyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListServices2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerBonjourProfilePolicyListToVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenWirelessControllerBonjourProfilePolicyListDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "WirelessControllerBonjourProfilePolicyList-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("from_vlan", flattenWirelessControllerBonjourProfilePolicyListFromVlan2edl(o["from-vlan"], d, "from_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["from-vlan"], "WirelessControllerBonjourProfilePolicyList-FromVlan"); ok {
			if err = d.Set("from_vlan", vv); err != nil {
				return fmt.Errorf("Error reading from_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from_vlan: %v", err)
		}
	}

	if err = d.Set("policy_id", flattenWirelessControllerBonjourProfilePolicyListPolicyId2edl(o["policy-id"], d, "policy_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-id"], "WirelessControllerBonjourProfilePolicyList-PolicyId"); ok {
			if err = d.Set("policy_id", vv); err != nil {
				return fmt.Errorf("Error reading policy_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_id: %v", err)
		}
	}

	if err = d.Set("services", flattenWirelessControllerBonjourProfilePolicyListServices2edl(o["services"], d, "services")); err != nil {
		if vv, ok := fortiAPIPatch(o["services"], "WirelessControllerBonjourProfilePolicyList-Services"); ok {
			if err = d.Set("services", vv); err != nil {
				return fmt.Errorf("Error reading services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading services: %v", err)
		}
	}

	if err = d.Set("to_vlan", flattenWirelessControllerBonjourProfilePolicyListToVlan2edl(o["to-vlan"], d, "to_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["to-vlan"], "WirelessControllerBonjourProfilePolicyList-ToVlan"); ok {
			if err = d.Set("to_vlan", vv); err != nil {
				return fmt.Errorf("Error reading to_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading to_vlan: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerBonjourProfilePolicyListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerBonjourProfilePolicyListDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListFromVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListPolicyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListServices2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWirelessControllerBonjourProfilePolicyListToVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandWirelessControllerBonjourProfilePolicyListDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("from_vlan"); ok || d.HasChange("from_vlan") {
		t, err := expandWirelessControllerBonjourProfilePolicyListFromVlan2edl(d, v, "from_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from-vlan"] = t
		}
	}

	if v, ok := d.GetOk("policy_id"); ok || d.HasChange("policy_id") {
		t, err := expandWirelessControllerBonjourProfilePolicyListPolicyId2edl(d, v, "policy_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-id"] = t
		}
	}

	if v, ok := d.GetOk("services"); ok || d.HasChange("services") {
		t, err := expandWirelessControllerBonjourProfilePolicyListServices2edl(d, v, "services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["services"] = t
		}
	}

	if v, ok := d.GetOk("to_vlan"); ok || d.HasChange("to_vlan") {
		t, err := expandWirelessControllerBonjourProfilePolicyListToVlan2edl(d, v, "to_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["to-vlan"] = t
		}
	}

	return &obj, nil
}
