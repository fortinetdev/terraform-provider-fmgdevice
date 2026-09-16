// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Fabric policy related attributes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallPolicyFabricPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallPolicyFabricPolicyUpdate,
		Read:   resourceFirewallPolicyFabricPolicyRead,
		Update: resourceFirewallPolicyFabricPolicyUpdate,
		Delete: resourceFirewallPolicyFabricPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"to": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallPolicyFabricPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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
	policy := d.Get("policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["policy"] = policy

	obj, err := getObjectFirewallPolicyFabricPolicy(d, false)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicyFabricPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallPolicyFabricPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicyFabricPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallPolicyFabricPolicy")

	return resourceFirewallPolicyFabricPolicyRead(d, m)
}

func resourceFirewallPolicyFabricPolicyDelete(d *schema.ResourceData, m interface{}) error {
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
	policy := d.Get("policy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["policy"] = policy

	obj, err := getObjectFirewallPolicyFabricPolicy(d, true)

	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicyFabricPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallPolicyFabricPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallPolicyFabricPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicyFabricPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	policy := d.Get("policy").(string)
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
	if policy == "" {
		policy = importOptionChecking(m.(*FortiClient).Cfg, "policy")
		if policy == "" {
			return fmt.Errorf("Parameter policy is missing")
		}
		if err = d.Set("policy", policy); err != nil {
			return fmt.Errorf("Error set params policy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["policy"] = policy

	o, err := c.ReadFirewallPolicyFabricPolicy(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallPolicyFabricPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicyFabricPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicyFabricPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicyFabricPolicyFrom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallPolicyFabricPolicyTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallPolicyFabricPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("from", flattenFirewallPolicyFabricPolicyFrom2edl(o["from"], d, "from")); err != nil {
		if vv, ok := fortiAPIPatch(o["from"], "FirewallPolicyFabricPolicy-From"); ok {
			if err = d.Set("from", vv); err != nil {
				return fmt.Errorf("Error reading from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if err = d.Set("to", flattenFirewallPolicyFabricPolicyTo2edl(o["to"], d, "to")); err != nil {
		if vv, ok := fortiAPIPatch(o["to"], "FirewallPolicyFabricPolicy-To"); ok {
			if err = d.Set("to", vv); err != nil {
				return fmt.Errorf("Error reading to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading to: %v", err)
		}
	}

	return nil
}

func flattenFirewallPolicyFabricPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallPolicyFabricPolicyFrom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallPolicyFabricPolicyTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallPolicyFabricPolicy(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("from"); ok || d.HasChange("from") {
		t, err := expandFirewallPolicyFabricPolicyFrom2edl(d, v, "from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from"] = t
		}
	}

	if v, ok := d.GetOk("to"); ok || d.HasChange("to") {
		t, err := expandFirewallPolicyFabricPolicyTo2edl(d, v, "to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["to"] = t
		}
	}

	return &obj, nil
}
