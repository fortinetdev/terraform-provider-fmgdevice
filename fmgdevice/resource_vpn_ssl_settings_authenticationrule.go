// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Authentication rule for SSL-VPN.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslSettingsAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslSettingsAuthenticationRuleCreate,
		Read:   resourceVpnSslSettingsAuthenticationRuleRead,
		Update: resourceVpnSslSettingsAuthenticationRuleUpdate,
		Delete: resourceVpnSslSettingsAuthenticationRuleDelete,

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
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
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
				ForceNew: true,
				Optional: true,
			},
			"portal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_address6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_address6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"user_peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnSslSettingsAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslSettingsAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslSettingsAuthenticationRule resource while getting object: %v", err)
	}

	_, err = c.CreateVpnSslSettingsAuthenticationRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslSettingsAuthenticationRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnSslSettingsAuthenticationRuleRead(d, m)
}

func resourceVpnSslSettingsAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslSettingsAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettingsAuthenticationRule resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslSettingsAuthenticationRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettingsAuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceVpnSslSettingsAuthenticationRuleRead(d, m)
}

func resourceVpnSslSettingsAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnSslSettingsAuthenticationRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslSettingsAuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslSettingsAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslSettingsAuthenticationRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettingsAuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslSettingsAuthenticationRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettingsAuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslSettingsAuthenticationRuleAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleCipher2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleClientCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRulePortal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleUserPeer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnSslSettingsAuthenticationRuleUsers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnSslSettingsAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth", flattenVpnSslSettingsAuthenticationRuleAuth2edl(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "VpnSslSettingsAuthenticationRule-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("cipher", flattenVpnSslSettingsAuthenticationRuleCipher2edl(o["cipher"], d, "cipher")); err != nil {
		if vv, ok := fortiAPIPatch(o["cipher"], "VpnSslSettingsAuthenticationRule-Cipher"); ok {
			if err = d.Set("cipher", vv); err != nil {
				return fmt.Errorf("Error reading cipher: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cipher: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenVpnSslSettingsAuthenticationRuleClientCert2edl(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "VpnSslSettingsAuthenticationRule-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("groups", flattenVpnSslSettingsAuthenticationRuleGroups2edl(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "VpnSslSettingsAuthenticationRule-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", flattenVpnSslSettingsAuthenticationRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "VpnSslSettingsAuthenticationRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("portal", flattenVpnSslSettingsAuthenticationRulePortal2edl(o["portal"], d, "portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal"], "VpnSslSettingsAuthenticationRule-Portal"); ok {
			if err = d.Set("portal", vv); err != nil {
				return fmt.Errorf("Error reading portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal: %v", err)
		}
	}

	if err = d.Set("realm", flattenVpnSslSettingsAuthenticationRuleRealm2edl(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "VpnSslSettingsAuthenticationRule-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("source_address", flattenVpnSslSettingsAuthenticationRuleSourceAddress2edl(o["source-address"], d, "source_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address"], "VpnSslSettingsAuthenticationRule-SourceAddress"); ok {
			if err = d.Set("source_address", vv); err != nil {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_negate", flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate2edl(o["source-address-negate"], d, "source_address_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address-negate"], "VpnSslSettingsAuthenticationRule-SourceAddressNegate"); ok {
			if err = d.Set("source_address_negate", vv); err != nil {
				return fmt.Errorf("Error reading source_address_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address_negate: %v", err)
		}
	}

	if err = d.Set("source_address6", flattenVpnSslSettingsAuthenticationRuleSourceAddress62edl(o["source-address6"], d, "source_address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address6"], "VpnSslSettingsAuthenticationRule-SourceAddress6"); ok {
			if err = d.Set("source_address6", vv); err != nil {
				return fmt.Errorf("Error reading source_address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address6: %v", err)
		}
	}

	if err = d.Set("source_address6_negate", flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate2edl(o["source-address6-negate"], d, "source_address6_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-address6-negate"], "VpnSslSettingsAuthenticationRule-SourceAddress6Negate"); ok {
			if err = d.Set("source_address6_negate", vv); err != nil {
				return fmt.Errorf("Error reading source_address6_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("source_interface", flattenVpnSslSettingsAuthenticationRuleSourceInterface2edl(o["source-interface"], d, "source_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-interface"], "VpnSslSettingsAuthenticationRule-SourceInterface"); ok {
			if err = d.Set("source_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_interface: %v", err)
		}
	}

	if err = d.Set("user_peer", flattenVpnSslSettingsAuthenticationRuleUserPeer2edl(o["user-peer"], d, "user_peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-peer"], "VpnSslSettingsAuthenticationRule-UserPeer"); ok {
			if err = d.Set("user_peer", vv); err != nil {
				return fmt.Errorf("Error reading user_peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_peer: %v", err)
		}
	}

	if err = d.Set("users", flattenVpnSslSettingsAuthenticationRuleUsers2edl(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "VpnSslSettingsAuthenticationRule-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	return nil
}

func flattenVpnSslSettingsAuthenticationRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslSettingsAuthenticationRuleAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleCipher2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleClientCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRulePortal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddressNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleUserPeer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnSslSettingsAuthenticationRuleUsers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnSslSettingsAuthenticationRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandVpnSslSettingsAuthenticationRuleAuth2edl(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("cipher"); ok || d.HasChange("cipher") {
		t, err := expandVpnSslSettingsAuthenticationRuleCipher2edl(d, v, "cipher")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandVpnSslSettingsAuthenticationRuleClientCert2edl(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandVpnSslSettingsAuthenticationRuleGroups2edl(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandVpnSslSettingsAuthenticationRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("portal"); ok || d.HasChange("portal") {
		t, err := expandVpnSslSettingsAuthenticationRulePortal2edl(d, v, "portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandVpnSslSettingsAuthenticationRuleRealm2edl(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		t, err := expandVpnSslSettingsAuthenticationRuleSourceAddress2edl(d, v, "source_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address"] = t
		}
	}

	if v, ok := d.GetOk("source_address_negate"); ok || d.HasChange("source_address_negate") {
		t, err := expandVpnSslSettingsAuthenticationRuleSourceAddressNegate2edl(d, v, "source_address_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_address6"); ok || d.HasChange("source_address6") {
		t, err := expandVpnSslSettingsAuthenticationRuleSourceAddress62edl(d, v, "source_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6"] = t
		}
	}

	if v, ok := d.GetOk("source_address6_negate"); ok || d.HasChange("source_address6_negate") {
		t, err := expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate2edl(d, v, "source_address6_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-address6-negate"] = t
		}
	}

	if v, ok := d.GetOk("source_interface"); ok || d.HasChange("source_interface") {
		t, err := expandVpnSslSettingsAuthenticationRuleSourceInterface2edl(d, v, "source_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-interface"] = t
		}
	}

	if v, ok := d.GetOk("user_peer"); ok || d.HasChange("user_peer") {
		t, err := expandVpnSslSettingsAuthenticationRuleUserPeer2edl(d, v, "user_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-peer"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandVpnSslSettingsAuthenticationRuleUsers2edl(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
