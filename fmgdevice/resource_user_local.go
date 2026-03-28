// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure local users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLocalCreate,
		Read:   resourceUserLocalRead,
		Update: resourceUserLocalUpdate,
		Delete: resourceUserLocalDelete,

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
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortitoken": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"passwd_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ppk_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"qkd_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sms_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tacacs_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_insensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username_case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"workstation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"history0": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history1": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history10": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history11": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history12": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history13": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history14": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history15": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history16": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history17": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history18": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history19": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history3": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history4": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history5": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history6": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history7": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history8": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history9": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceUserLocalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating UserLocal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserLocal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserLocal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserLocal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserLocal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserLocal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserLocalRead(d, m)
}

func resourceUserLocalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating UserLocal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserLocal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserLocalRead(d, m)
}

func resourceUserLocalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserLocal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLocalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserLocal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserLocal resource from API: %v", err)
	}
	return nil
}

func flattenUserLocalAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalPasswdPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalPpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalQkdProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserLocalTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalUsernameCaseInsensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalUsernameCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalUsernameSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLocalWorkstation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_concurrent_override", flattenUserLocalAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "UserLocal-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenUserLocalAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "UserLocal-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenUserLocalAuthtimeout(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "UserLocal-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("email_to", flattenUserLocalEmailTo(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "UserLocal-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenUserLocalFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken"], "UserLocal-Fortitoken"); ok {
			if err = d.Set("fortitoken", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserLocalId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "UserLocal-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserLocalLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserLocal-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenUserLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("passwd_policy", flattenUserLocalPasswdPolicy(o["passwd-policy"], d, "passwd_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-policy"], "UserLocal-PasswdPolicy"); ok {
			if err = d.Set("passwd_policy", vv); err != nil {
				return fmt.Errorf("Error reading passwd_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_policy: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenUserLocalPpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "UserLocal-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenUserLocalQkdProfile(o["qkd-profile"], d, "qkd_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd-profile"], "UserLocal-QkdProfile"); ok {
			if err = d.Set("qkd_profile", vv); err != nil {
				return fmt.Errorf("Error reading qkd_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenUserLocalRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "UserLocal-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenUserLocalSamlServer(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "UserLocal-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenUserLocalSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "UserLocal-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenUserLocalSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-phone"], "UserLocal-SmsPhone"); ok {
			if err = d.Set("sms_phone", vv); err != nil {
				return fmt.Errorf("Error reading sms_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenUserLocalSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "UserLocal-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("status", flattenUserLocalStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "UserLocal-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tacacs_server", flattenUserLocalTacacsServer(o["tacacs+-server"], d, "tacacs_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs+-server"], "UserLocal-TacacsServer"); ok {
			if err = d.Set("tacacs_server", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserLocalTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "UserLocal-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenUserLocalTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "UserLocal-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenUserLocalTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "UserLocal-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLocalType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserLocal-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username_case_insensitivity", flattenUserLocalUsernameCaseInsensitivity(o["username-case-insensitivity"], d, "username_case_insensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-insensitivity"], "UserLocal-UsernameCaseInsensitivity"); ok {
			if err = d.Set("username_case_insensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_sensitivity", flattenUserLocalUsernameCaseSensitivity(o["username-case-sensitivity"], d, "username_case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitivity"], "UserLocal-UsernameCaseSensitivity"); ok {
			if err = d.Set("username_case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
		}
	}

	if err = d.Set("username_sensitivity", flattenUserLocalUsernameSensitivity(o["username-sensitivity"], d, "username_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-sensitivity"], "UserLocal-UsernameSensitivity"); ok {
			if err = d.Set("username_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_sensitivity: %v", err)
		}
	}

	if err = d.Set("workstation", flattenUserLocalWorkstation(o["workstation"], d, "workstation")); err != nil {
		if vv, ok := fortiAPIPatch(o["workstation"], "UserLocal-Workstation"); ok {
			if err = d.Set("workstation", vv); err != nil {
				return fmt.Errorf("Error reading workstation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workstation: %v", err)
		}
	}

	return nil
}

func flattenUserLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserLocalAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalPasswdPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalPpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalQkdProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameCaseInsensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalWorkstation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLocalHistory0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory13(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserLocalHistory9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectUserLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandUserLocalAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandUserLocalAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandUserLocalAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandUserLocalEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok || d.HasChange("fortitoken") {
		t, err := expandUserLocalFortitoken(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandUserLocalId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserLocalLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandUserLocalPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("passwd_policy"); ok || d.HasChange("passwd_policy") {
		t, err := expandUserLocalPasswdPolicy(d, v, "passwd_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-policy"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandUserLocalPpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandUserLocalPpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok || d.HasChange("qkd_profile") {
		t, err := expandUserLocalQkdProfile(d, v, "qkd_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandUserLocalRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandUserLocalSamlServer(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandUserLocalSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok || d.HasChange("sms_phone") {
		t, err := expandUserLocalSmsPhone(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandUserLocalSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandUserLocalStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_server"); ok || d.HasChange("tacacs_server") {
		t, err := expandUserLocalTacacsServer(d, v, "tacacs_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandUserLocalTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandUserLocalTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandUserLocalTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserLocalType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username_case_insensitivity"); ok || d.HasChange("username_case_insensitivity") {
		t, err := expandUserLocalUsernameCaseInsensitivity(d, v, "username_case_insensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-insensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitivity"); ok || d.HasChange("username_case_sensitivity") {
		t, err := expandUserLocalUsernameCaseSensitivity(d, v, "username_case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_sensitivity"); ok || d.HasChange("username_sensitivity") {
		t, err := expandUserLocalUsernameSensitivity(d, v, "username_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("workstation"); ok || d.HasChange("workstation") {
		t, err := expandUserLocalWorkstation(d, v, "workstation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workstation"] = t
		}
	}

	if v, ok := d.GetOk("history0"); ok || d.HasChange("history0") {
		t, err := expandUserLocalHistory0(d, v, "history0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history0"] = t
		}
	}

	if v, ok := d.GetOk("history1"); ok || d.HasChange("history1") {
		t, err := expandUserLocalHistory1(d, v, "history1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history1"] = t
		}
	}

	if v, ok := d.GetOk("history10"); ok || d.HasChange("history10") {
		t, err := expandUserLocalHistory10(d, v, "history10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history10"] = t
		}
	}

	if v, ok := d.GetOk("history11"); ok || d.HasChange("history11") {
		t, err := expandUserLocalHistory11(d, v, "history11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history11"] = t
		}
	}

	if v, ok := d.GetOk("history12"); ok || d.HasChange("history12") {
		t, err := expandUserLocalHistory12(d, v, "history12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history12"] = t
		}
	}

	if v, ok := d.GetOk("history13"); ok || d.HasChange("history13") {
		t, err := expandUserLocalHistory13(d, v, "history13")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history13"] = t
		}
	}

	if v, ok := d.GetOk("history14"); ok || d.HasChange("history14") {
		t, err := expandUserLocalHistory14(d, v, "history14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history14"] = t
		}
	}

	if v, ok := d.GetOk("history15"); ok || d.HasChange("history15") {
		t, err := expandUserLocalHistory15(d, v, "history15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history15"] = t
		}
	}

	if v, ok := d.GetOk("history16"); ok || d.HasChange("history16") {
		t, err := expandUserLocalHistory16(d, v, "history16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history16"] = t
		}
	}

	if v, ok := d.GetOk("history17"); ok || d.HasChange("history17") {
		t, err := expandUserLocalHistory17(d, v, "history17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history17"] = t
		}
	}

	if v, ok := d.GetOk("history18"); ok || d.HasChange("history18") {
		t, err := expandUserLocalHistory18(d, v, "history18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history18"] = t
		}
	}

	if v, ok := d.GetOk("history19"); ok || d.HasChange("history19") {
		t, err := expandUserLocalHistory19(d, v, "history19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history19"] = t
		}
	}

	if v, ok := d.GetOk("history2"); ok || d.HasChange("history2") {
		t, err := expandUserLocalHistory2(d, v, "history2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history2"] = t
		}
	}

	if v, ok := d.GetOk("history3"); ok || d.HasChange("history3") {
		t, err := expandUserLocalHistory3(d, v, "history3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history3"] = t
		}
	}

	if v, ok := d.GetOk("history4"); ok || d.HasChange("history4") {
		t, err := expandUserLocalHistory4(d, v, "history4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history4"] = t
		}
	}

	if v, ok := d.GetOk("history5"); ok || d.HasChange("history5") {
		t, err := expandUserLocalHistory5(d, v, "history5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history5"] = t
		}
	}

	if v, ok := d.GetOk("history6"); ok || d.HasChange("history6") {
		t, err := expandUserLocalHistory6(d, v, "history6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history6"] = t
		}
	}

	if v, ok := d.GetOk("history7"); ok || d.HasChange("history7") {
		t, err := expandUserLocalHistory7(d, v, "history7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history7"] = t
		}
	}

	if v, ok := d.GetOk("history8"); ok || d.HasChange("history8") {
		t, err := expandUserLocalHistory8(d, v, "history8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history8"] = t
		}
	}

	if v, ok := d.GetOk("history9"); ok || d.HasChange("history9") {
		t, err := expandUserLocalHistory9(d, v, "history9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history9"] = t
		}
	}

	return &obj, nil
}
