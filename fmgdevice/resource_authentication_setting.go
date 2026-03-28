// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure authentication setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationSettingUpdate,
		Read:   resourceAuthenticationSettingRead,
		Update: resourceAuthenticationSettingUpdate,
		Delete: resourceAuthenticationSettingDelete,

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
			"active_auth_scheme": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captive_portal_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ssl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"captive_portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_captive_portal": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cert_captive_portal_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_captive_portal_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookie_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookie_refresh_div": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dev_range": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistent_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_auth_scheme": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"update_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_cert_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_auth_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewrite_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAuthenticationSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAuthenticationSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AuthenticationSetting")

	return resourceAuthenticationSettingRead(d, m)
}

func resourceAuthenticationSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAuthenticationSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationSetting(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AuthenticationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationSetting resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationSettingActiveAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingAuthHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCertCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingCertCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCertCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCookieMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCookieRefreshDiv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingDevRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingIpAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingPersistentCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingSsoAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingUpdateTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingUserCertCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationSettingLogAuthRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingRewriteHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAuthenticationSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("active_auth_scheme", flattenAuthenticationSettingActiveAuthScheme(o["active-auth-scheme"], d, "active_auth_scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-auth-scheme"], "AuthenticationSetting-ActiveAuthScheme"); ok {
			if err = d.Set("active_auth_scheme", vv); err != nil {
				return fmt.Errorf("Error reading active_auth_scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_auth_scheme: %v", err)
		}
	}

	if err = d.Set("auth_https", flattenAuthenticationSettingAuthHttps(o["auth-https"], d, "auth_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-https"], "AuthenticationSetting-AuthHttps"); ok {
			if err = d.Set("auth_https", vv); err != nil {
				return fmt.Errorf("Error reading auth_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_https: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenAuthenticationSettingCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal"], "AuthenticationSetting-CaptivePortal"); ok {
			if err = d.Set("captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip", flattenAuthenticationSettingCaptivePortalIp(o["captive-portal-ip"], d, "captive_portal_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ip"], "AuthenticationSetting-CaptivePortalIp"); ok {
			if err = d.Set("captive_portal_ip", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip6", flattenAuthenticationSettingCaptivePortalIp6(o["captive-portal-ip6"], d, "captive_portal_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ip6"], "AuthenticationSetting-CaptivePortalIp6"); ok {
			if err = d.Set("captive_portal_ip6", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
		}
	}

	if err = d.Set("captive_portal_port", flattenAuthenticationSettingCaptivePortalPort(o["captive-portal-port"], d, "captive_portal_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-port"], "AuthenticationSetting-CaptivePortalPort"); ok {
			if err = d.Set("captive_portal_port", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_port: %v", err)
		}
	}

	if err = d.Set("captive_portal_ssl_port", flattenAuthenticationSettingCaptivePortalSslPort(o["captive-portal-ssl-port"], d, "captive_portal_ssl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ssl-port"], "AuthenticationSetting-CaptivePortalSslPort"); ok {
			if err = d.Set("captive_portal_ssl_port", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
		}
	}

	if err = d.Set("captive_portal_type", flattenAuthenticationSettingCaptivePortalType(o["captive-portal-type"], d, "captive_portal_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-type"], "AuthenticationSetting-CaptivePortalType"); ok {
			if err = d.Set("captive_portal_type", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_type: %v", err)
		}
	}

	if err = d.Set("captive_portal6", flattenAuthenticationSettingCaptivePortal6(o["captive-portal6"], d, "captive_portal6")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal6"], "AuthenticationSetting-CaptivePortal6"); ok {
			if err = d.Set("captive_portal6", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal6: %v", err)
		}
	}

	if err = d.Set("cert_auth", flattenAuthenticationSettingCertAuth(o["cert-auth"], d, "cert_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-auth"], "AuthenticationSetting-CertAuth"); ok {
			if err = d.Set("cert_auth", vv); err != nil {
				return fmt.Errorf("Error reading cert_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_auth: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal", flattenAuthenticationSettingCertCaptivePortal(o["cert-captive-portal"], d, "cert_captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-captive-portal"], "AuthenticationSetting-CertCaptivePortal"); ok {
			if err = d.Set("cert_captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading cert_captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_captive_portal: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal_ip", flattenAuthenticationSettingCertCaptivePortalIp(o["cert-captive-portal-ip"], d, "cert_captive_portal_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-captive-portal-ip"], "AuthenticationSetting-CertCaptivePortalIp"); ok {
			if err = d.Set("cert_captive_portal_ip", vv); err != nil {
				return fmt.Errorf("Error reading cert_captive_portal_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal_port", flattenAuthenticationSettingCertCaptivePortalPort(o["cert-captive-portal-port"], d, "cert_captive_portal_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-captive-portal-port"], "AuthenticationSetting-CertCaptivePortalPort"); ok {
			if err = d.Set("cert_captive_portal_port", vv); err != nil {
				return fmt.Errorf("Error reading cert_captive_portal_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_captive_portal_port: %v", err)
		}
	}

	if err = d.Set("cookie_max_age", flattenAuthenticationSettingCookieMaxAge(o["cookie-max-age"], d, "cookie_max_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["cookie-max-age"], "AuthenticationSetting-CookieMaxAge"); ok {
			if err = d.Set("cookie_max_age", vv); err != nil {
				return fmt.Errorf("Error reading cookie_max_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cookie_max_age: %v", err)
		}
	}

	if err = d.Set("cookie_refresh_div", flattenAuthenticationSettingCookieRefreshDiv(o["cookie-refresh-div"], d, "cookie_refresh_div")); err != nil {
		if vv, ok := fortiAPIPatch(o["cookie-refresh-div"], "AuthenticationSetting-CookieRefreshDiv"); ok {
			if err = d.Set("cookie_refresh_div", vv); err != nil {
				return fmt.Errorf("Error reading cookie_refresh_div: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cookie_refresh_div: %v", err)
		}
	}

	if err = d.Set("dev_range", flattenAuthenticationSettingDevRange(o["dev-range"], d, "dev_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-range"], "AuthenticationSetting-DevRange"); ok {
			if err = d.Set("dev_range", vv); err != nil {
				return fmt.Errorf("Error reading dev_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_range: %v", err)
		}
	}

	if err = d.Set("ip_auth_cookie", flattenAuthenticationSettingIpAuthCookie(o["ip-auth-cookie"], d, "ip_auth_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-auth-cookie"], "AuthenticationSetting-IpAuthCookie"); ok {
			if err = d.Set("ip_auth_cookie", vv); err != nil {
				return fmt.Errorf("Error reading ip_auth_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_auth_cookie: %v", err)
		}
	}

	if err = d.Set("persistent_cookie", flattenAuthenticationSettingPersistentCookie(o["persistent-cookie"], d, "persistent_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["persistent-cookie"], "AuthenticationSetting-PersistentCookie"); ok {
			if err = d.Set("persistent_cookie", vv); err != nil {
				return fmt.Errorf("Error reading persistent_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading persistent_cookie: %v", err)
		}
	}

	if err = d.Set("sso_auth_scheme", flattenAuthenticationSettingSsoAuthScheme(o["sso-auth-scheme"], d, "sso_auth_scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-auth-scheme"], "AuthenticationSetting-SsoAuthScheme"); ok {
			if err = d.Set("sso_auth_scheme", vv); err != nil {
				return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
		}
	}

	if err = d.Set("update_time", flattenAuthenticationSettingUpdateTime(o["update-time"], d, "update_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-time"], "AuthenticationSetting-UpdateTime"); ok {
			if err = d.Set("update_time", vv); err != nil {
				return fmt.Errorf("Error reading update_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_time: %v", err)
		}
	}

	if err = d.Set("user_cert_ca", flattenAuthenticationSettingUserCertCa(o["user-cert-ca"], d, "user_cert_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-cert-ca"], "AuthenticationSetting-UserCertCa"); ok {
			if err = d.Set("user_cert_ca", vv); err != nil {
				return fmt.Errorf("Error reading user_cert_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_cert_ca: %v", err)
		}
	}

	if err = d.Set("log_auth_request", flattenAuthenticationSettingLogAuthRequest(o["log-auth-request"], d, "log_auth_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-auth-request"], "AuthenticationSetting-LogAuthRequest"); ok {
			if err = d.Set("log_auth_request", vv); err != nil {
				return fmt.Errorf("Error reading log_auth_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_auth_request: %v", err)
		}
	}

	if err = d.Set("rewrite_https_port", flattenAuthenticationSettingRewriteHttpsPort(o["rewrite-https-port"], d, "rewrite_https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rewrite-https-port"], "AuthenticationSetting-RewriteHttpsPort"); ok {
			if err = d.Set("rewrite_https_port", vv); err != nil {
				return fmt.Errorf("Error reading rewrite_https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rewrite_https_port: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAuthenticationSettingActiveAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingAuthHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingCertAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingCertCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCookieMaxAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCookieRefreshDiv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingDevRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingIpAuthCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingPersistentCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingSsoAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingUpdateTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingUserCertCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationSettingLogAuthRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingRewriteHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("active_auth_scheme"); ok || d.HasChange("active_auth_scheme") {
		t, err := expandAuthenticationSettingActiveAuthScheme(d, v, "active_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("auth_https"); ok || d.HasChange("auth_https") {
		t, err := expandAuthenticationSettingAuthHttps(d, v, "auth_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-https"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok || d.HasChange("captive_portal") {
		t, err := expandAuthenticationSettingCaptivePortal(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip"); ok || d.HasChange("captive_portal_ip") {
		t, err := expandAuthenticationSettingCaptivePortalIp(d, v, "captive_portal_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip6"); ok || d.HasChange("captive_portal_ip6") {
		t, err := expandAuthenticationSettingCaptivePortalIp6(d, v, "captive_portal_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip6"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_port"); ok || d.HasChange("captive_portal_port") {
		t, err := expandAuthenticationSettingCaptivePortalPort(d, v, "captive_portal_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-port"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ssl_port"); ok || d.HasChange("captive_portal_ssl_port") {
		t, err := expandAuthenticationSettingCaptivePortalSslPort(d, v, "captive_portal_ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ssl-port"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_type"); ok || d.HasChange("captive_portal_type") {
		t, err := expandAuthenticationSettingCaptivePortalType(d, v, "captive_portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-type"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal6"); ok || d.HasChange("captive_portal6") {
		t, err := expandAuthenticationSettingCaptivePortal6(d, v, "captive_portal6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal6"] = t
		}
	}

	if v, ok := d.GetOk("cert_auth"); ok || d.HasChange("cert_auth") {
		t, err := expandAuthenticationSettingCertAuth(d, v, "cert_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("cert_captive_portal"); ok || d.HasChange("cert_captive_portal") {
		t, err := expandAuthenticationSettingCertCaptivePortal(d, v, "cert_captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("cert_captive_portal_ip"); ok || d.HasChange("cert_captive_portal_ip") {
		t, err := expandAuthenticationSettingCertCaptivePortalIp(d, v, "cert_captive_portal_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-captive-portal-ip"] = t
		}
	}

	if v, ok := d.GetOk("cert_captive_portal_port"); ok || d.HasChange("cert_captive_portal_port") {
		t, err := expandAuthenticationSettingCertCaptivePortalPort(d, v, "cert_captive_portal_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-captive-portal-port"] = t
		}
	}

	if v, ok := d.GetOk("cookie_max_age"); ok || d.HasChange("cookie_max_age") {
		t, err := expandAuthenticationSettingCookieMaxAge(d, v, "cookie_max_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-max-age"] = t
		}
	}

	if v, ok := d.GetOk("cookie_refresh_div"); ok || d.HasChange("cookie_refresh_div") {
		t, err := expandAuthenticationSettingCookieRefreshDiv(d, v, "cookie_refresh_div")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-refresh-div"] = t
		}
	}

	if v, ok := d.GetOk("dev_range"); ok || d.HasChange("dev_range") {
		t, err := expandAuthenticationSettingDevRange(d, v, "dev_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-range"] = t
		}
	}

	if v, ok := d.GetOk("ip_auth_cookie"); ok || d.HasChange("ip_auth_cookie") {
		t, err := expandAuthenticationSettingIpAuthCookie(d, v, "ip_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("persistent_cookie"); ok || d.HasChange("persistent_cookie") {
		t, err := expandAuthenticationSettingPersistentCookie(d, v, "persistent_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistent-cookie"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_scheme"); ok || d.HasChange("sso_auth_scheme") {
		t, err := expandAuthenticationSettingSsoAuthScheme(d, v, "sso_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("update_time"); ok || d.HasChange("update_time") {
		t, err := expandAuthenticationSettingUpdateTime(d, v, "update_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-time"] = t
		}
	}

	if v, ok := d.GetOk("user_cert_ca"); ok || d.HasChange("user_cert_ca") {
		t, err := expandAuthenticationSettingUserCertCa(d, v, "user_cert_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-cert-ca"] = t
		}
	}

	if v, ok := d.GetOk("log_auth_request"); ok || d.HasChange("log_auth_request") {
		t, err := expandAuthenticationSettingLogAuthRequest(d, v, "log_auth_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-auth-request"] = t
		}
	}

	if v, ok := d.GetOk("rewrite_https_port"); ok || d.HasChange("rewrite_https_port") {
		t, err := expandAuthenticationSettingRewriteHttpsPort(d, v, "rewrite_https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite-https-port"] = t
		}
	}

	return &obj, nil
}
