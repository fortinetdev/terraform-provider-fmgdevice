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

func resourceUserOidc() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserOidcCreate,
		Read:   resourceUserOidcRead,
		Update: resourceUserOidcUpdate,
		Delete: resourceUserOidcDelete,

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
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"discovery_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icon_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwks_uri": &schema.Schema{
				Type:     schema.TypeString,
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
			"private_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"token_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserOidcCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserOidc(d)
	if err != nil {
		return fmt.Errorf("Error creating UserOidc resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserOidc(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserOidc(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserOidc resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserOidc(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserOidc resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserOidcRead(d, m)
}

func resourceUserOidcUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserOidc(d)
	if err != nil {
		return fmt.Errorf("Error updating UserOidc resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserOidc(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserOidc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserOidcRead(d, m)
}

func resourceUserOidcDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserOidc(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserOidc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserOidcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserOidc(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserOidc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserOidc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserOidc resource from API: %v", err)
	}
	return nil
}

func flattenUserOidcAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcAuthorizationUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcDiscoveryUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcDisplayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcDomainHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcGroupAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcIconUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcIssuer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcJwksUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserOidcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserOidcTokenUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcUserAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcUserRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserOidcVerifyIssuer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserOidc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_method", flattenUserOidcAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-method"], "UserOidc-AuthMethod"); ok {
			if err = d.Set("auth_method", vv); err != nil {
				return fmt.Errorf("Error reading auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserOidcAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "UserOidc-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("authorization_url", flattenUserOidcAuthorizationUrl(o["authorization-url"], d, "authorization_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization-url"], "UserOidc-AuthorizationUrl"); ok {
			if err = d.Set("authorization_url", vv); err != nil {
				return fmt.Errorf("Error reading authorization_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization_url: %v", err)
		}
	}

	if err = d.Set("client_id", flattenUserOidcClientId(o["client-id"], d, "client_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-id"], "UserOidc-ClientId"); ok {
			if err = d.Set("client_id", vv); err != nil {
				return fmt.Errorf("Error reading client_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("client_secret", flattenUserOidcClientSecret(o["client-secret"], d, "client_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret"], "UserOidc-ClientSecret"); ok {
			if err = d.Set("client_secret", vv); err != nil {
				return fmt.Errorf("Error reading client_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenUserOidcClockTolerance(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["clock-tolerance"], "UserOidc-ClockTolerance"); ok {
			if err = d.Set("clock_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading clock_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("discovery_url", flattenUserOidcDiscoveryUrl(o["discovery-url"], d, "discovery_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["discovery-url"], "UserOidc-DiscoveryUrl"); ok {
			if err = d.Set("discovery_url", vv); err != nil {
				return fmt.Errorf("Error reading discovery_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discovery_url: %v", err)
		}
	}

	if err = d.Set("display_name", flattenUserOidcDisplayName(o["display-name"], d, "display_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-name"], "UserOidc-DisplayName"); ok {
			if err = d.Set("display_name", vv); err != nil {
				return fmt.Errorf("Error reading display_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_name: %v", err)
		}
	}

	if err = d.Set("domain_hint", flattenUserOidcDomainHint(o["domain-hint"], d, "domain_hint")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-hint"], "UserOidc-DomainHint"); ok {
			if err = d.Set("domain_hint", vv); err != nil {
				return fmt.Errorf("Error reading domain_hint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_hint: %v", err)
		}
	}

	if err = d.Set("group_attr_name", flattenUserOidcGroupAttrName(o["group-attr-name"], d, "group_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-attr-name"], "UserOidc-GroupAttrName"); ok {
			if err = d.Set("group_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading group_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_attr_name: %v", err)
		}
	}

	if err = d.Set("icon_url", flattenUserOidcIconUrl(o["icon-url"], d, "icon_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-url"], "UserOidc-IconUrl"); ok {
			if err = d.Set("icon_url", vv); err != nil {
				return fmt.Errorf("Error reading icon_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_url: %v", err)
		}
	}

	if err = d.Set("issuer", flattenUserOidcIssuer(o["issuer"], d, "issuer")); err != nil {
		if vv, ok := fortiAPIPatch(o["issuer"], "UserOidc-Issuer"); ok {
			if err = d.Set("issuer", vv); err != nil {
				return fmt.Errorf("Error reading issuer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("jwks_uri", flattenUserOidcJwksUri(o["jwks-uri"], d, "jwks_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["jwks-uri"], "UserOidc-JwksUri"); ok {
			if err = d.Set("jwks_uri", vv); err != nil {
				return fmt.Errorf("Error reading jwks_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jwks_uri: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserOidcLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "UserOidc-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenUserOidcName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserOidc-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("private_key", flattenUserOidcPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "UserOidc-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("token_url", flattenUserOidcTokenUrl(o["token-url"], d, "token_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["token-url"], "UserOidc-TokenUrl"); ok {
			if err = d.Set("token_url", vv); err != nil {
				return fmt.Errorf("Error reading token_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading token_url: %v", err)
		}
	}

	if err = d.Set("type", flattenUserOidcType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "UserOidc-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_attr_name", flattenUserOidcUserAttrName(o["user-attr-name"], d, "user_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-attr-name"], "UserOidc-UserAttrName"); ok {
			if err = d.Set("user_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading user_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_attr_name: %v", err)
		}
	}

	if err = d.Set("user_regex", flattenUserOidcUserRegex(o["user-regex"], d, "user_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-regex"], "UserOidc-UserRegex"); ok {
			if err = d.Set("user_regex", vv); err != nil {
				return fmt.Errorf("Error reading user_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_regex: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenUserOidcVerifyCert(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "UserOidc-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("verify_issuer", flattenUserOidcVerifyIssuer(o["verify-issuer"], d, "verify_issuer")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-issuer"], "UserOidc-VerifyIssuer"); ok {
			if err = d.Set("verify_issuer", vv); err != nil {
				return fmt.Errorf("Error reading verify_issuer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_issuer: %v", err)
		}
	}

	return nil
}

func flattenUserOidcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserOidcAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcAuthorizationUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClientSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcDiscoveryUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcDisplayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcDomainHint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcGroupAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcIconUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcIssuer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcJwksUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserOidcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserOidcTokenUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcUserAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcUserRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserOidcVerifyIssuer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserOidc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_method"); ok || d.HasChange("auth_method") {
		t, err := expandUserOidcAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandUserOidcAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization_url"); ok || d.HasChange("authorization_url") {
		t, err := expandUserOidcAuthorizationUrl(d, v, "authorization_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-url"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok || d.HasChange("client_id") {
		t, err := expandUserOidcClientId(d, v, "client_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok || d.HasChange("client_secret") {
		t, err := expandUserOidcClientSecret(d, v, "client_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("clock_tolerance"); ok || d.HasChange("clock_tolerance") {
		t, err := expandUserOidcClockTolerance(d, v, "clock_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("discovery_url"); ok || d.HasChange("discovery_url") {
		t, err := expandUserOidcDiscoveryUrl(d, v, "discovery_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-url"] = t
		}
	}

	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		t, err := expandUserOidcDisplayName(d, v, "display_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-name"] = t
		}
	}

	if v, ok := d.GetOk("domain_hint"); ok || d.HasChange("domain_hint") {
		t, err := expandUserOidcDomainHint(d, v, "domain_hint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-hint"] = t
		}
	}

	if v, ok := d.GetOk("group_attr_name"); ok || d.HasChange("group_attr_name") {
		t, err := expandUserOidcGroupAttrName(d, v, "group_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("icon_url"); ok || d.HasChange("icon_url") {
		t, err := expandUserOidcIconUrl(d, v, "icon_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-url"] = t
		}
	}

	if v, ok := d.GetOk("issuer"); ok || d.HasChange("issuer") {
		t, err := expandUserOidcIssuer(d, v, "issuer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["issuer"] = t
		}
	}

	if v, ok := d.GetOk("jwks_uri"); ok || d.HasChange("jwks_uri") {
		t, err := expandUserOidcJwksUri(d, v, "jwks_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jwks-uri"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserOidcLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserOidcName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok || d.HasChange("private_key") {
		t, err := expandUserOidcPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("token_url"); ok || d.HasChange("token_url") {
		t, err := expandUserOidcTokenUrl(d, v, "token_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-url"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandUserOidcType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_attr_name"); ok || d.HasChange("user_attr_name") {
		t, err := expandUserOidcUserAttrName(d, v, "user_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("user_regex"); ok || d.HasChange("user_regex") {
		t, err := expandUserOidcUserRegex(d, v, "user_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-regex"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandUserOidcVerifyCert(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("verify_issuer"); ok || d.HasChange("verify_issuer") {
		t, err := expandUserOidcVerifyIssuer(d, v, "verify_issuer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-issuer"] = t
		}
	}

	return &obj, nil
}
