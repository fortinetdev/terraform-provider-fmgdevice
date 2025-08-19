// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure 802.1x MAC Authentication Bypass (MAB) policies.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicy8021X() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicy8021XCreate,
		Read:   resourceSwitchControllerSecurityPolicy8021XRead,
		Update: resourceSwitchControllerSecurityPolicy8021XUpdate,
		Delete: resourceSwitchControllerSecurityPolicy8021XDelete,

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
			"auth_fail_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_fail_vlan_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_order": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authserver_timeout_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authserver_timeout_tagged": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authserver_timeout_tagged_vlanid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"authserver_timeout_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authserver_timeout_vlanid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_auto_untagged_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_passthru": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"framevid_apply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_auth_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"guest_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_vlan_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"open_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_timeout_overwrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicy8021XCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSecurityPolicy8021X(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceSwitchControllerSecurityPolicy8021XUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSecurityPolicy8021X(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceSwitchControllerSecurityPolicy8021XDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSecurityPolicy8021X(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicy8021XRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSecurityPolicy8021X(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicy8021X(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicy8021X resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicy8021XAuthFailVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthFailVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicy8021XAuthOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutTagged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutTaggedVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicy8021XDacl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XEapPassthru(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XFramevidApply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestAuthDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerSecurityPolicy8021XMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XOpenAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_fail_vlan", flattenSwitchControllerSecurityPolicy8021XAuthFailVlan(o["auth-fail-vlan"], d, "auth_fail_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-fail-vlan"], "SwitchControllerSecurityPolicy8021X-AuthFailVlan"); ok {
			if err = d.Set("auth_fail_vlan", vv); err != nil {
				return fmt.Errorf("Error reading auth_fail_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_fail_vlan: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlan_id", flattenSwitchControllerSecurityPolicy8021XAuthFailVlanId(o["auth-fail-vlan-id"], d, "auth_fail_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-fail-vlan-id"], "SwitchControllerSecurityPolicy8021X-AuthFailVlanId"); ok {
			if err = d.Set("auth_fail_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading auth_fail_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_fail_vlan_id: %v", err)
		}
	}

	if err = d.Set("auth_order", flattenSwitchControllerSecurityPolicy8021XAuthOrder(o["auth-order"], d, "auth_order")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-order"], "SwitchControllerSecurityPolicy8021X-AuthOrder"); ok {
			if err = d.Set("auth_order", vv); err != nil {
				return fmt.Errorf("Error reading auth_order: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_order: %v", err)
		}
	}

	if err = d.Set("auth_priority", flattenSwitchControllerSecurityPolicy8021XAuthPriority(o["auth-priority"], d, "auth_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-priority"], "SwitchControllerSecurityPolicy8021X-AuthPriority"); ok {
			if err = d.Set("auth_priority", vv); err != nil {
				return fmt.Errorf("Error reading auth_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_priority: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_period", flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(o["authserver-timeout-period"], d, "authserver_timeout_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-period"], "SwitchControllerSecurityPolicy8021X-AuthserverTimeoutPeriod"); ok {
			if err = d.Set("authserver_timeout_period", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_period: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_tagged", flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutTagged(o["authserver-timeout-tagged"], d, "authserver_timeout_tagged")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-tagged"], "SwitchControllerSecurityPolicy8021X-AuthserverTimeoutTagged"); ok {
			if err = d.Set("authserver_timeout_tagged", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_tagged: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_tagged: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_tagged_vlanid", flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutTaggedVlanid(o["authserver-timeout-tagged-vlanid"], d, "authserver_timeout_tagged_vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-tagged-vlanid"], "SwitchControllerSecurityPolicy8021X-AuthserverTimeoutTaggedVlanid"); ok {
			if err = d.Set("authserver_timeout_tagged_vlanid", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_tagged_vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_tagged_vlanid: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_vlan", flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(o["authserver-timeout-vlan"], d, "authserver_timeout_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-vlan"], "SwitchControllerSecurityPolicy8021X-AuthserverTimeoutVlan"); ok {
			if err = d.Set("authserver_timeout_vlan", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_vlan: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_vlanid", flattenSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(o["authserver-timeout-vlanid"], d, "authserver_timeout_vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-vlanid"], "SwitchControllerSecurityPolicy8021X-AuthserverTimeoutVlanid"); ok {
			if err = d.Set("authserver_timeout_vlanid", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_vlanid: %v", err)
		}
	}

	if err = d.Set("dacl", flattenSwitchControllerSecurityPolicy8021XDacl(o["dacl"], d, "dacl")); err != nil {
		if vv, ok := fortiAPIPatch(o["dacl"], "SwitchControllerSecurityPolicy8021X-Dacl"); ok {
			if err = d.Set("dacl", vv); err != nil {
				return fmt.Errorf("Error reading dacl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dacl: %v", err)
		}
	}

	if err = d.Set("eap_auto_untagged_vlans", flattenSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(o["eap-auto-untagged-vlans"], d, "eap_auto_untagged_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-auto-untagged-vlans"], "SwitchControllerSecurityPolicy8021X-EapAutoUntaggedVlans"); ok {
			if err = d.Set("eap_auto_untagged_vlans", vv); err != nil {
				return fmt.Errorf("Error reading eap_auto_untagged_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_auto_untagged_vlans: %v", err)
		}
	}

	if err = d.Set("eap_passthru", flattenSwitchControllerSecurityPolicy8021XEapPassthru(o["eap-passthru"], d, "eap_passthru")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-passthru"], "SwitchControllerSecurityPolicy8021X-EapPassthru"); ok {
			if err = d.Set("eap_passthru", vv); err != nil {
				return fmt.Errorf("Error reading eap_passthru: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_passthru: %v", err)
		}
	}

	if err = d.Set("framevid_apply", flattenSwitchControllerSecurityPolicy8021XFramevidApply(o["framevid-apply"], d, "framevid_apply")); err != nil {
		if vv, ok := fortiAPIPatch(o["framevid-apply"], "SwitchControllerSecurityPolicy8021X-FramevidApply"); ok {
			if err = d.Set("framevid_apply", vv); err != nil {
				return fmt.Errorf("Error reading framevid_apply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading framevid_apply: %v", err)
		}
	}

	if err = d.Set("guest_auth_delay", flattenSwitchControllerSecurityPolicy8021XGuestAuthDelay(o["guest-auth-delay"], d, "guest_auth_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-auth-delay"], "SwitchControllerSecurityPolicy8021X-GuestAuthDelay"); ok {
			if err = d.Set("guest_auth_delay", vv); err != nil {
				return fmt.Errorf("Error reading guest_auth_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_auth_delay: %v", err)
		}
	}

	if err = d.Set("guest_vlan", flattenSwitchControllerSecurityPolicy8021XGuestVlan(o["guest-vlan"], d, "guest_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-vlan"], "SwitchControllerSecurityPolicy8021X-GuestVlan"); ok {
			if err = d.Set("guest_vlan", vv); err != nil {
				return fmt.Errorf("Error reading guest_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_vlan: %v", err)
		}
	}

	if err = d.Set("guest_vlan_id", flattenSwitchControllerSecurityPolicy8021XGuestVlanId(o["guest-vlan-id"], d, "guest_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-vlan-id"], "SwitchControllerSecurityPolicy8021X-GuestVlanId"); ok {
			if err = d.Set("guest_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading guest_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_vlan_id: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenSwitchControllerSecurityPolicy8021XMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-auth-bypass"], "SwitchControllerSecurityPolicy8021X-MacAuthBypass"); ok {
			if err = d.Set("mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSecurityPolicy8021XName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSecurityPolicy8021X-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("open_auth", flattenSwitchControllerSecurityPolicy8021XOpenAuth(o["open-auth"], d, "open_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-auth"], "SwitchControllerSecurityPolicy8021X-OpenAuth"); ok {
			if err = d.Set("open_auth", vv); err != nil {
				return fmt.Errorf("Error reading open_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_auth: %v", err)
		}
	}

	if err = d.Set("policy_type", flattenSwitchControllerSecurityPolicy8021XPolicyType(o["policy-type"], d, "policy_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-type"], "SwitchControllerSecurityPolicy8021X-PolicyType"); ok {
			if err = d.Set("policy_type", vv); err != nil {
				return fmt.Errorf("Error reading policy_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_type: %v", err)
		}
	}

	if err = d.Set("radius_timeout_overwrite", flattenSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(o["radius-timeout-overwrite"], d, "radius_timeout_overwrite")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-timeout-overwrite"], "SwitchControllerSecurityPolicy8021X-RadiusTimeoutOverwrite"); ok {
			if err = d.Set("radius_timeout_overwrite", vv); err != nil {
				return fmt.Errorf("Error reading radius_timeout_overwrite: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_timeout_overwrite: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSwitchControllerSecurityPolicy8021XSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "SwitchControllerSecurityPolicy8021X-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("user_group", flattenSwitchControllerSecurityPolicy8021XUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "SwitchControllerSecurityPolicy8021X-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicy8021XFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSecurityPolicy8021XAuthFailVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthFailVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicy8021XAuthOrder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutTagged(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutTaggedVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicy8021XDacl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XEapPassthru(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XFramevidApply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestAuthDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSecurityPolicy8021XMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XOpenAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_fail_vlan"); ok || d.HasChange("auth_fail_vlan") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthFailVlan(d, v, "auth_fail_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan"] = t
		}
	}

	if v, ok := d.GetOk("auth_fail_vlan_id"); ok || d.HasChange("auth_fail_vlan_id") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthFailVlanId(d, v, "auth_fail_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("auth_order"); ok || d.HasChange("auth_order") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthOrder(d, v, "auth_order")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-order"] = t
		}
	}

	if v, ok := d.GetOk("auth_priority"); ok || d.HasChange("auth_priority") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthPriority(d, v, "auth_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-priority"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_period"); ok || d.HasChange("authserver_timeout_period") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(d, v, "authserver_timeout_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-period"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_tagged"); ok || d.HasChange("authserver_timeout_tagged") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutTagged(d, v, "authserver_timeout_tagged")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-tagged"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_tagged_vlanid"); ok || d.HasChange("authserver_timeout_tagged_vlanid") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutTaggedVlanid(d, v, "authserver_timeout_tagged_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-tagged-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_vlan"); ok || d.HasChange("authserver_timeout_vlan") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(d, v, "authserver_timeout_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-vlan"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_vlanid"); ok || d.HasChange("authserver_timeout_vlanid") {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(d, v, "authserver_timeout_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("dacl"); ok || d.HasChange("dacl") {
		t, err := expandSwitchControllerSecurityPolicy8021XDacl(d, v, "dacl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dacl"] = t
		}
	}

	if v, ok := d.GetOk("eap_auto_untagged_vlans"); ok || d.HasChange("eap_auto_untagged_vlans") {
		t, err := expandSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(d, v, "eap_auto_untagged_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-auto-untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("eap_passthru"); ok || d.HasChange("eap_passthru") {
		t, err := expandSwitchControllerSecurityPolicy8021XEapPassthru(d, v, "eap_passthru")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-passthru"] = t
		}
	}

	if v, ok := d.GetOk("framevid_apply"); ok || d.HasChange("framevid_apply") {
		t, err := expandSwitchControllerSecurityPolicy8021XFramevidApply(d, v, "framevid_apply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["framevid-apply"] = t
		}
	}

	if v, ok := d.GetOk("guest_auth_delay"); ok || d.HasChange("guest_auth_delay") {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestAuthDelay(d, v, "guest_auth_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-auth-delay"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan"); ok || d.HasChange("guest_vlan") {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestVlan(d, v, "guest_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan_id"); ok || d.HasChange("guest_vlan_id") {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestVlanId(d, v, "guest_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok || d.HasChange("mac_auth_bypass") {
		t, err := expandSwitchControllerSecurityPolicy8021XMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSecurityPolicy8021XName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("open_auth"); ok || d.HasChange("open_auth") {
		t, err := expandSwitchControllerSecurityPolicy8021XOpenAuth(d, v, "open_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-auth"] = t
		}
	}

	if v, ok := d.GetOk("policy_type"); ok || d.HasChange("policy_type") {
		t, err := expandSwitchControllerSecurityPolicy8021XPolicyType(d, v, "policy_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-type"] = t
		}
	}

	if v, ok := d.GetOk("radius_timeout_overwrite"); ok || d.HasChange("radius_timeout_overwrite") {
		t, err := expandSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d, v, "radius_timeout_overwrite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-timeout-overwrite"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandSwitchControllerSecurityPolicy8021XSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandSwitchControllerSecurityPolicy8021XUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
