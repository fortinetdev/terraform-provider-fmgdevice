// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure firewall authentication portals.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAuthPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAuthPortalUpdate,
		Read:   resourceFirewallAuthPortalRead,
		Update: resourceFirewallAuthPortalUpdate,
		Delete: resourceFirewallAuthPortalDelete,

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
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"portal_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portal_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallAuthPortalUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectFirewallAuthPortal(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAuthPortal resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallAuthPortal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAuthPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallAuthPortal")

	return resourceFirewallAuthPortalRead(d, m)
}

func resourceFirewallAuthPortalDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteFirewallAuthPortal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAuthPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAuthPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAuthPortal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAuthPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAuthPortal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAuthPortal resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAuthPortalGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAuthPortalIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallAuthPortalPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAuthPortalPortalAddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAuthPortalProxyAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAuthPortal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("groups", flattenFirewallAuthPortalGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "FirewallAuthPortal-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenFirewallAuthPortalIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["identity-based-route"], "FirewallAuthPortal-IdentityBasedRoute"); ok {
			if err = d.Set("identity_based_route", vv); err != nil {
				return fmt.Errorf("Error reading identity_based_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("portal_addr", flattenFirewallAuthPortalPortalAddr(o["portal-addr"], d, "portal_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-addr"], "FirewallAuthPortal-PortalAddr"); ok {
			if err = d.Set("portal_addr", vv); err != nil {
				return fmt.Errorf("Error reading portal_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_addr: %v", err)
		}
	}

	if err = d.Set("portal_addr6", flattenFirewallAuthPortalPortalAddr6(o["portal-addr6"], d, "portal_addr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-addr6"], "FirewallAuthPortal-PortalAddr6"); ok {
			if err = d.Set("portal_addr6", vv); err != nil {
				return fmt.Errorf("Error reading portal_addr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_addr6: %v", err)
		}
	}

	if err = d.Set("proxy_auth", flattenFirewallAuthPortalProxyAuth(o["proxy-auth"], d, "proxy_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-auth"], "FirewallAuthPortal-ProxyAuth"); ok {
			if err = d.Set("proxy_auth", vv); err != nil {
				return fmt.Errorf("Error reading proxy_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_auth: %v", err)
		}
	}

	return nil
}

func flattenFirewallAuthPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAuthPortalGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAuthPortalIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallAuthPortalPortalAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAuthPortalPortalAddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAuthPortalProxyAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAuthPortal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallAuthPortalGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("identity_based_route"); ok || d.HasChange("identity_based_route") {
		t, err := expandFirewallAuthPortalIdentityBasedRoute(d, v, "identity_based_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	}

	if v, ok := d.GetOk("portal_addr"); ok || d.HasChange("portal_addr") {
		t, err := expandFirewallAuthPortalPortalAddr(d, v, "portal_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("portal_addr6"); ok || d.HasChange("portal_addr6") {
		t, err := expandFirewallAuthPortalPortalAddr6(d, v, "portal_addr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-addr6"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth"); ok || d.HasChange("proxy_auth") {
		t, err := expandFirewallAuthPortalProxyAuth(d, v, "proxy_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth"] = t
		}
	}

	return &obj, nil
}
