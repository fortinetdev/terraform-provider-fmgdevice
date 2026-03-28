// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure peer users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPeerCreate,
		Read:   resourceUserPeerRead,
		Update: resourceUserPeerUpdate,
		Delete: resourceUserPeerDelete,

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
			"ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mandatory_ca_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"mfa_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mfa_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ocsp_override_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ldap_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserPeerCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating UserPeer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserPeer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserPeer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserPeer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserPeer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserPeer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserPeer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteUserPeer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserPeer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserPeer resource from API: %v", err)
	}
	return nil
}

func flattenUserPeerCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserPeerCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerCnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerMandatoryCaVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerMfaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerMfaServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserPeerMfaUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerOcspOverrideServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserPeerSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserPeerLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ca", flattenUserPeerCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "UserPeer-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("cn", flattenUserPeerCn(o["cn"], d, "cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn"], "UserPeer-Cn"); ok {
			if err = d.Set("cn", vv); err != nil {
				return fmt.Errorf("Error reading cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn: %v", err)
		}
	}

	if err = d.Set("cn_type", flattenUserPeerCnType(o["cn-type"], d, "cn_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn-type"], "UserPeer-CnType"); ok {
			if err = d.Set("cn_type", vv); err != nil {
				return fmt.Errorf("Error reading cn_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn_type: %v", err)
		}
	}

	if err = d.Set("mandatory_ca_verify", flattenUserPeerMandatoryCaVerify(o["mandatory-ca-verify"], d, "mandatory_ca_verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["mandatory-ca-verify"], "UserPeer-MandatoryCaVerify"); ok {
			if err = d.Set("mandatory_ca_verify", vv); err != nil {
				return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
		}
	}

	if err = d.Set("mfa_mode", flattenUserPeerMfaMode(o["mfa-mode"], d, "mfa_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-mode"], "UserPeer-MfaMode"); ok {
			if err = d.Set("mfa_mode", vv); err != nil {
				return fmt.Errorf("Error reading mfa_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_mode: %v", err)
		}
	}

	if err = d.Set("mfa_server", flattenUserPeerMfaServer(o["mfa-server"], d, "mfa_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-server"], "UserPeer-MfaServer"); ok {
			if err = d.Set("mfa_server", vv); err != nil {
				return fmt.Errorf("Error reading mfa_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_server: %v", err)
		}
	}

	if err = d.Set("mfa_username", flattenUserPeerMfaUsername(o["mfa-username"], d, "mfa_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-username"], "UserPeer-MfaUsername"); ok {
			if err = d.Set("mfa_username", vv); err != nil {
				return fmt.Errorf("Error reading mfa_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_username: %v", err)
		}
	}

	if err = d.Set("name", flattenUserPeerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserPeer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocsp_override_server", flattenUserPeerOcspOverrideServer(o["ocsp-override-server"], d, "ocsp_override_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-override-server"], "UserPeer-OcspOverrideServer"); ok {
			if err = d.Set("ocsp_override_server", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_override_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_override_server: %v", err)
		}
	}

	if err = d.Set("subject", flattenUserPeerSubject(o["subject"], d, "subject")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject"], "UserPeer-Subject"); ok {
			if err = d.Set("subject", vv); err != nil {
				return fmt.Errorf("Error reading subject: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserPeerTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "UserPeer-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("ldap_mode", flattenUserPeerLdapMode(o["ldap-mode"], d, "ldap_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-mode"], "UserPeer-LdapMode"); ok {
			if err = d.Set("ldap_mode", vv); err != nil {
				return fmt.Errorf("Error reading ldap_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_mode: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserPeerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserPeer-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenUserPeerLdapUsername(o["ldap-username"], d, "ldap_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-username"], "UserPeer-LdapUsername"); ok {
			if err = d.Set("ldap_username", vv); err != nil {
				return fmt.Errorf("Error reading ldap_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	return nil
}

func flattenUserPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserPeerCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMandatoryCaVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerMfaServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerMfaUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerOcspOverrideServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserPeerLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok || d.HasChange("ca") {
		t, err := expandUserPeerCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("cn"); ok || d.HasChange("cn") {
		t, err := expandUserPeerCn(d, v, "cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn"] = t
		}
	}

	if v, ok := d.GetOk("cn_type"); ok || d.HasChange("cn_type") {
		t, err := expandUserPeerCnType(d, v, "cn_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-type"] = t
		}
	}

	if v, ok := d.GetOk("mandatory_ca_verify"); ok || d.HasChange("mandatory_ca_verify") {
		t, err := expandUserPeerMandatoryCaVerify(d, v, "mandatory_ca_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mandatory-ca-verify"] = t
		}
	}

	if v, ok := d.GetOk("mfa_mode"); ok || d.HasChange("mfa_mode") {
		t, err := expandUserPeerMfaMode(d, v, "mfa_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-mode"] = t
		}
	}

	if v, ok := d.GetOk("mfa_password"); ok || d.HasChange("mfa_password") {
		t, err := expandUserPeerMfaPassword(d, v, "mfa_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-password"] = t
		}
	}

	if v, ok := d.GetOk("mfa_server"); ok || d.HasChange("mfa_server") {
		t, err := expandUserPeerMfaServer(d, v, "mfa_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-server"] = t
		}
	}

	if v, ok := d.GetOk("mfa_username"); ok || d.HasChange("mfa_username") {
		t, err := expandUserPeerMfaUsername(d, v, "mfa_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-username"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserPeerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_override_server"); ok || d.HasChange("ocsp_override_server") {
		t, err := expandUserPeerOcspOverrideServer(d, v, "ocsp_override_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-override-server"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandUserPeerPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("subject"); ok || d.HasChange("subject") {
		t, err := expandUserPeerSubject(d, v, "subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandUserPeerTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("ldap_mode"); ok || d.HasChange("ldap_mode") {
		t, err := expandUserPeerLdapMode(d, v, "ldap_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-mode"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok || d.HasChange("ldap_password") {
		t, err := expandUserPeerLdapPassword(d, v, "ldap_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserPeerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok || d.HasChange("ldap_username") {
		t, err := expandUserPeerLdapUsername(d, v, "ldap_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	return &obj, nil
}
