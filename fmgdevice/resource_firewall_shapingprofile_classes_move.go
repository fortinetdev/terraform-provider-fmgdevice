// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShapingProfileClassesMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingProfileClassesMoveUpdate,
		Read:   resourceFirewallShapingProfileClassesMoveRead,
		Update: resourceFirewallShapingProfileClassesMoveUpdate,
		Delete: resourceFirewallShapingProfileClassesMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
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
			"shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"classes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallShapingProfileClassesMoveUpdate(d *schema.ResourceData, m interface{}) error {
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
	shaping_profile := d.Get("shaping_profile").(string)
	classes := d.Get("classes").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile
	paradict["classes"] = classes

	target := d.Get("target").(string)
	obj, err := getObjectFirewallShapingProfileClassesMove(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfileClassesMove resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallShapingProfileClassesMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfileClassesMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallShapingProfileClassesMove" + "_" + device_name + "_" + device_vdom + "_" + shaping_profile + "_" + classes + "_" + target)

	return resourceFirewallShapingProfileClassesMoveRead(d, m)
}

func resourceFirewallShapingProfileClassesMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceFirewallShapingProfileClassesMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	sid := d.Get("classes").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	shaping_profile := d.Get("shaping_profile").(string)
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
	if shaping_profile == "" {
		shaping_profile = importOptionChecking(m.(*FortiClient).Cfg, "shaping_profile")
		if shaping_profile == "" {
			return fmt.Errorf("Parameter shaping_profile is missing")
		}
		if err = d.Set("shaping_profile", shaping_profile); err != nil {
			return fmt.Errorf("Error set params shaping_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["shaping_profile"] = shaping_profile

	o, err := c.ReadFirewallShapingProfileClassesMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallShapingProfileClassesMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["name"]; !ok {
					return fmt.Errorf("Error reading FirewallShapingProfileClassesMove resource: name doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["name"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading FirewallShapingProfileClassesMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "name(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "name(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenFirewallShapingProfileClassesMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShapingProfileClassesMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileClassesMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingProfileClassesMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandFirewallShapingProfileClassesMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandFirewallShapingProfileClassesMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
