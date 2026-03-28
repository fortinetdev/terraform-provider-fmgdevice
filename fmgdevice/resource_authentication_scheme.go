// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Authentication Schemes.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationScheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationSchemeCreate,
		Read:   resourceAuthenticationSchemeRead,
		Update: resourceAuthenticationSchemeUpdate,
		Delete: resourceAuthenticationSchemeDelete,

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
			"digest_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"digest_rfc2069": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ems_device_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_idp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsso_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kerberos_keytab": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
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
			"negotiate_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"require_tfa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"saml_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssh_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"user_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_database": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oidc_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oidc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"search_all_ldap_databases": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAuthenticationSchemeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error creating AuthenticationScheme resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadAuthenticationScheme(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateAuthenticationScheme(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating AuthenticationScheme resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateAuthenticationScheme(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating AuthenticationScheme resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceAuthenticationSchemeRead(d, m)
}

func resourceAuthenticationSchemeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationScheme resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAuthenticationScheme(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationScheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceAuthenticationSchemeRead(d, m)
}

func resourceAuthenticationSchemeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAuthenticationScheme(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationScheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSchemeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationScheme(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AuthenticationScheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationScheme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationScheme resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationSchemeDigestAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeDigestRfc2069(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeEmsDeviceOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeExternalIdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeFssoGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeKerberosKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeNegotiateNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeRequireTfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeSamlTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeSshCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeUserCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeUserDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeOidcServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSchemeOidcTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeSearchAllLdapDatabases(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAuthenticationScheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("digest_algo", flattenAuthenticationSchemeDigestAlgo(o["digest-algo"], d, "digest_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-algo"], "AuthenticationScheme-DigestAlgo"); ok {
			if err = d.Set("digest_algo", vv); err != nil {
				return fmt.Errorf("Error reading digest_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_algo: %v", err)
		}
	}

	if err = d.Set("digest_rfc2069", flattenAuthenticationSchemeDigestRfc2069(o["digest-rfc2069"], d, "digest_rfc2069")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-rfc2069"], "AuthenticationScheme-DigestRfc2069"); ok {
			if err = d.Set("digest_rfc2069", vv); err != nil {
				return fmt.Errorf("Error reading digest_rfc2069: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_rfc2069: %v", err)
		}
	}

	if err = d.Set("domain_controller", flattenAuthenticationSchemeDomainController(o["domain-controller"], d, "domain_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-controller"], "AuthenticationScheme-DomainController"); ok {
			if err = d.Set("domain_controller", vv); err != nil {
				return fmt.Errorf("Error reading domain_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if err = d.Set("ems_device_owner", flattenAuthenticationSchemeEmsDeviceOwner(o["ems-device-owner"], d, "ems_device_owner")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-device-owner"], "AuthenticationScheme-EmsDeviceOwner"); ok {
			if err = d.Set("ems_device_owner", vv); err != nil {
				return fmt.Errorf("Error reading ems_device_owner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_device_owner: %v", err)
		}
	}

	if err = d.Set("external_idp", flattenAuthenticationSchemeExternalIdp(o["external-idp"], d, "external_idp")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-idp"], "AuthenticationScheme-ExternalIdp"); ok {
			if err = d.Set("external_idp", vv); err != nil {
				return fmt.Errorf("Error reading external_idp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_idp: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenAuthenticationSchemeFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-agent-for-ntlm"], "AuthenticationScheme-FssoAgentForNtlm"); ok {
			if err = d.Set("fsso_agent_for_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("fsso_guest", flattenAuthenticationSchemeFssoGuest(o["fsso-guest"], d, "fsso_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-guest"], "AuthenticationScheme-FssoGuest"); ok {
			if err = d.Set("fsso_guest", vv); err != nil {
				return fmt.Errorf("Error reading fsso_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_guest: %v", err)
		}
	}

	if err = d.Set("group_attr_type", flattenAuthenticationSchemeGroupAttrType(o["group-attr-type"], d, "group_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-attr-type"], "AuthenticationScheme-GroupAttrType"); ok {
			if err = d.Set("group_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading group_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_attr_type: %v", err)
		}
	}

	if err = d.Set("kerberos_keytab", flattenAuthenticationSchemeKerberosKeytab(o["kerberos-keytab"], d, "kerberos_keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["kerberos-keytab"], "AuthenticationScheme-KerberosKeytab"); ok {
			if err = d.Set("kerberos_keytab", vv); err != nil {
				return fmt.Errorf("Error reading kerberos_keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kerberos_keytab: %v", err)
		}
	}

	if err = d.Set("method", flattenAuthenticationSchemeMethod(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "AuthenticationScheme-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("name", flattenAuthenticationSchemeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "AuthenticationScheme-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("negotiate_ntlm", flattenAuthenticationSchemeNegotiateNtlm(o["negotiate-ntlm"], d, "negotiate_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["negotiate-ntlm"], "AuthenticationScheme-NegotiateNtlm"); ok {
			if err = d.Set("negotiate_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
		}
	}

	if err = d.Set("require_tfa", flattenAuthenticationSchemeRequireTfa(o["require-tfa"], d, "require_tfa")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-tfa"], "AuthenticationScheme-RequireTfa"); ok {
			if err = d.Set("require_tfa", vv); err != nil {
				return fmt.Errorf("Error reading require_tfa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_tfa: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenAuthenticationSchemeSamlServer(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "AuthenticationScheme-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("saml_timeout", flattenAuthenticationSchemeSamlTimeout(o["saml-timeout"], d, "saml_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-timeout"], "AuthenticationScheme-SamlTimeout"); ok {
			if err = d.Set("saml_timeout", vv); err != nil {
				return fmt.Errorf("Error reading saml_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_timeout: %v", err)
		}
	}

	if err = d.Set("ssh_ca", flattenAuthenticationSchemeSshCa(o["ssh-ca"], d, "ssh_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-ca"], "AuthenticationScheme-SshCa"); ok {
			if err = d.Set("ssh_ca", vv); err != nil {
				return fmt.Errorf("Error reading ssh_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_ca: %v", err)
		}
	}

	if err = d.Set("user_cert", flattenAuthenticationSchemeUserCert(o["user-cert"], d, "user_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-cert"], "AuthenticationScheme-UserCert"); ok {
			if err = d.Set("user_cert", vv); err != nil {
				return fmt.Errorf("Error reading user_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_cert: %v", err)
		}
	}

	if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-database"], "AuthenticationScheme-UserDatabase"); ok {
			if err = d.Set("user_database", vv); err != nil {
				return fmt.Errorf("Error reading user_database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_database: %v", err)
		}
	}

	if err = d.Set("oidc_server", flattenAuthenticationSchemeOidcServer(o["oidc-server"], d, "oidc_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["oidc-server"], "AuthenticationScheme-OidcServer"); ok {
			if err = d.Set("oidc_server", vv); err != nil {
				return fmt.Errorf("Error reading oidc_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oidc_server: %v", err)
		}
	}

	if err = d.Set("oidc_timeout", flattenAuthenticationSchemeOidcTimeout(o["oidc-timeout"], d, "oidc_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["oidc-timeout"], "AuthenticationScheme-OidcTimeout"); ok {
			if err = d.Set("oidc_timeout", vv); err != nil {
				return fmt.Errorf("Error reading oidc_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oidc_timeout: %v", err)
		}
	}

	if err = d.Set("search_all_ldap_databases", flattenAuthenticationSchemeSearchAllLdapDatabases(o["search-all-ldap-databases"], d, "search_all_ldap_databases")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-all-ldap-databases"], "AuthenticationScheme-SearchAllLdapDatabases"); ok {
			if err = d.Set("search_all_ldap_databases", vv); err != nil {
				return fmt.Errorf("Error reading search_all_ldap_databases: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_all_ldap_databases: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationSchemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAuthenticationSchemeDigestAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeDigestRfc2069(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeEmsDeviceOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeExternalIdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeFssoGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeKerberosKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeNegotiateNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeRequireTfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeSamlTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSshCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeUserCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeOidcServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSchemeOidcTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSearchAllLdapDatabases(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationScheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("digest_algo"); ok || d.HasChange("digest_algo") {
		t, err := expandAuthenticationSchemeDigestAlgo(d, v, "digest_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-algo"] = t
		}
	}

	if v, ok := d.GetOk("digest_rfc2069"); ok || d.HasChange("digest_rfc2069") {
		t, err := expandAuthenticationSchemeDigestRfc2069(d, v, "digest_rfc2069")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-rfc2069"] = t
		}
	}

	if v, ok := d.GetOk("domain_controller"); ok || d.HasChange("domain_controller") {
		t, err := expandAuthenticationSchemeDomainController(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("ems_device_owner"); ok || d.HasChange("ems_device_owner") {
		t, err := expandAuthenticationSchemeEmsDeviceOwner(d, v, "ems_device_owner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-device-owner"] = t
		}
	}

	if v, ok := d.GetOk("external_idp"); ok || d.HasChange("external_idp") {
		t, err := expandAuthenticationSchemeExternalIdp(d, v, "external_idp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-idp"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok || d.HasChange("fsso_agent_for_ntlm") {
		t, err := expandAuthenticationSchemeFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("fsso_guest"); ok || d.HasChange("fsso_guest") {
		t, err := expandAuthenticationSchemeFssoGuest(d, v, "fsso_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-guest"] = t
		}
	}

	if v, ok := d.GetOk("group_attr_type"); ok || d.HasChange("group_attr_type") {
		t, err := expandAuthenticationSchemeGroupAttrType(d, v, "group_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("kerberos_keytab"); ok || d.HasChange("kerberos_keytab") {
		t, err := expandAuthenticationSchemeKerberosKeytab(d, v, "kerberos_keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kerberos-keytab"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandAuthenticationSchemeMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandAuthenticationSchemeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_ntlm"); ok || d.HasChange("negotiate_ntlm") {
		t, err := expandAuthenticationSchemeNegotiateNtlm(d, v, "negotiate_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("require_tfa"); ok || d.HasChange("require_tfa") {
		t, err := expandAuthenticationSchemeRequireTfa(d, v, "require_tfa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-tfa"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandAuthenticationSchemeSamlServer(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("saml_timeout"); ok || d.HasChange("saml_timeout") {
		t, err := expandAuthenticationSchemeSamlTimeout(d, v, "saml_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssh_ca"); ok || d.HasChange("ssh_ca") {
		t, err := expandAuthenticationSchemeSshCa(d, v, "ssh_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-ca"] = t
		}
	}

	if v, ok := d.GetOk("user_cert"); ok || d.HasChange("user_cert") {
		t, err := expandAuthenticationSchemeUserCert(d, v, "user_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_database"); ok || d.HasChange("user_database") {
		t, err := expandAuthenticationSchemeUserDatabase(d, v, "user_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-database"] = t
		}
	}

	if v, ok := d.GetOk("oidc_server"); ok || d.HasChange("oidc_server") {
		t, err := expandAuthenticationSchemeOidcServer(d, v, "oidc_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oidc-server"] = t
		}
	}

	if v, ok := d.GetOk("oidc_timeout"); ok || d.HasChange("oidc_timeout") {
		t, err := expandAuthenticationSchemeOidcTimeout(d, v, "oidc_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oidc-timeout"] = t
		}
	}

	if v, ok := d.GetOk("search_all_ldap_databases"); ok || d.HasChange("search_all_ldap_databases") {
		t, err := expandAuthenticationSchemeSearchAllLdapDatabases(d, v, "search_all_ldap_databases")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-all-ldap-databases"] = t
		}
	}

	return &obj, nil
}
