// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Rule.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIdentityBasedRouteRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIdentityBasedRouteRuleCreate,
		Read:   resourceFirewallIdentityBasedRouteRuleRead,
		Update: resourceFirewallIdentityBasedRouteRuleUpdate,
		Delete: resourceFirewallIdentityBasedRouteRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
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
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIdentityBasedRouteRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	identity_based_route := d.Get("identity_based_route").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["identity_based_route"] = identity_based_route

	obj, err := getObjectFirewallIdentityBasedRouteRule(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIdentityBasedRouteRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallIdentityBasedRouteRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallIdentityBasedRouteRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallIdentityBasedRouteRule resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateFirewallIdentityBasedRouteRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallIdentityBasedRouteRule resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceFirewallIdentityBasedRouteRuleRead(d, m)
			} else {
				return fmt.Errorf("Error creating FirewallIdentityBasedRouteRule resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallIdentityBasedRouteRuleRead(d, m)
}

func resourceFirewallIdentityBasedRouteRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	identity_based_route := d.Get("identity_based_route").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["identity_based_route"] = identity_based_route

	obj, err := getObjectFirewallIdentityBasedRouteRule(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIdentityBasedRouteRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallIdentityBasedRouteRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIdentityBasedRouteRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallIdentityBasedRouteRuleRead(d, m)
}

func resourceFirewallIdentityBasedRouteRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	identity_based_route := d.Get("identity_based_route").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["identity_based_route"] = identity_based_route

	wsParams["adom"] = adomv

	err = c.DeleteFirewallIdentityBasedRouteRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIdentityBasedRouteRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIdentityBasedRouteRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	identity_based_route := d.Get("identity_based_route").(string)
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
	if identity_based_route == "" {
		identity_based_route = importOptionChecking(m.(*FortiClient).Cfg, "identity_based_route")
		if identity_based_route == "" {
			return fmt.Errorf("Parameter identity_based_route is missing")
		}
		if err = d.Set("identity_based_route", identity_based_route); err != nil {
			return fmt.Errorf("Error set params identity_based_route: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["identity_based_route"] = identity_based_route

	o, err := c.ReadFirewallIdentityBasedRouteRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallIdentityBasedRouteRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIdentityBasedRouteRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIdentityBasedRouteRule resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIdentityBasedRouteRuleDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIdentityBasedRouteRuleGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteRuleGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallIdentityBasedRouteRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIdentityBasedRouteRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenFirewallIdentityBasedRouteRuleDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "FirewallIdentityBasedRouteRule-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("gateway", flattenFirewallIdentityBasedRouteRuleGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "FirewallIdentityBasedRouteRule-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", flattenFirewallIdentityBasedRouteRuleGroups2edl(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallIdentityBasedRouteRule-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallIdentityBasedRouteRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallIdentityBasedRouteRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenFirewallIdentityBasedRouteRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIdentityBasedRouteRuleDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIdentityBasedRouteRuleGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteRuleGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallIdentityBasedRouteRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIdentityBasedRouteRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandFirewallIdentityBasedRouteRuleDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandFirewallIdentityBasedRouteRuleGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallIdentityBasedRouteRuleGroups2edl(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallIdentityBasedRouteRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
