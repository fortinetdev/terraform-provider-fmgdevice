// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Authentication Rules.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationRuleCreate,
		Read:   resourceAuthenticationRuleRead,
		Update: resourceAuthenticationRuleUpdate,
		Delete: resourceAuthenticationRuleDelete,

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
			"active_auth_method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cert_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cors_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cors_stateful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sso_auth_method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transaction_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"form_auth_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error creating AuthenticationRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadAuthenticationRule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateAuthenticationRule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating AuthenticationRule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateAuthenticationRule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating AuthenticationRule resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceAuthenticationRuleRead(d, m)
}

func resourceAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationRule(d)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAuthenticationRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceAuthenticationRuleRead(d, m)
}

func resourceAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAuthenticationRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationRule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationRuleActiveAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleCertAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleCorsDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleCorsStateful(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleIpBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleSsoAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAuthenticationRuleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleTransactionBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleWebAuthCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleWebPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleFormAuthFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationRuleWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectAuthenticationRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("active_auth_method", flattenAuthenticationRuleActiveAuthMethod(o["active-auth-method"], d, "active_auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-auth-method"], "AuthenticationRule-ActiveAuthMethod"); ok {
			if err = d.Set("active_auth_method", vv); err != nil {
				return fmt.Errorf("Error reading active_auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_auth_method: %v", err)
		}
	}

	if err = d.Set("cert_auth_cookie", flattenAuthenticationRuleCertAuthCookie(o["cert-auth-cookie"], d, "cert_auth_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-auth-cookie"], "AuthenticationRule-CertAuthCookie"); ok {
			if err = d.Set("cert_auth_cookie", vv); err != nil {
				return fmt.Errorf("Error reading cert_auth_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_auth_cookie: %v", err)
		}
	}

	if err = d.Set("comments", flattenAuthenticationRuleComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "AuthenticationRule-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("cors_depth", flattenAuthenticationRuleCorsDepth(o["cors-depth"], d, "cors_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-depth"], "AuthenticationRule-CorsDepth"); ok {
			if err = d.Set("cors_depth", vv); err != nil {
				return fmt.Errorf("Error reading cors_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_depth: %v", err)
		}
	}

	if err = d.Set("cors_stateful", flattenAuthenticationRuleCorsStateful(o["cors-stateful"], d, "cors_stateful")); err != nil {
		if vv, ok := fortiAPIPatch(o["cors-stateful"], "AuthenticationRule-CorsStateful"); ok {
			if err = d.Set("cors_stateful", vv); err != nil {
				return fmt.Errorf("Error reading cors_stateful: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cors_stateful: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenAuthenticationRuleDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "AuthenticationRule-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenAuthenticationRuleDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "AuthenticationRule-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("ip_based", flattenAuthenticationRuleIpBased(o["ip-based"], d, "ip_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-based"], "AuthenticationRule-IpBased"); ok {
			if err = d.Set("ip_based", vv); err != nil {
				return fmt.Errorf("Error reading ip_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_based: %v", err)
		}
	}

	if err = d.Set("name", flattenAuthenticationRuleName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "AuthenticationRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenAuthenticationRuleProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "AuthenticationRule-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenAuthenticationRuleSrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "AuthenticationRule-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenAuthenticationRuleSrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "AuthenticationRule-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenAuthenticationRuleSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "AuthenticationRule-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("sso_auth_method", flattenAuthenticationRuleSsoAuthMethod(o["sso-auth-method"], d, "sso_auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-auth-method"], "AuthenticationRule-SsoAuthMethod"); ok {
			if err = d.Set("sso_auth_method", vv); err != nil {
				return fmt.Errorf("Error reading sso_auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_auth_method: %v", err)
		}
	}

	if err = d.Set("status", flattenAuthenticationRuleStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "AuthenticationRule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transaction_based", flattenAuthenticationRuleTransactionBased(o["transaction-based"], d, "transaction_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["transaction-based"], "AuthenticationRule-TransactionBased"); ok {
			if err = d.Set("transaction_based", vv); err != nil {
				return fmt.Errorf("Error reading transaction_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transaction_based: %v", err)
		}
	}

	if err = d.Set("web_auth_cookie", flattenAuthenticationRuleWebAuthCookie(o["web-auth-cookie"], d, "web_auth_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-auth-cookie"], "AuthenticationRule-WebAuthCookie"); ok {
			if err = d.Set("web_auth_cookie", vv); err != nil {
				return fmt.Errorf("Error reading web_auth_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_auth_cookie: %v", err)
		}
	}

	if err = d.Set("web_portal", flattenAuthenticationRuleWebPortal(o["web-portal"], d, "web_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-portal"], "AuthenticationRule-WebPortal"); ok {
			if err = d.Set("web_portal", vv); err != nil {
				return fmt.Errorf("Error reading web_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_portal: %v", err)
		}
	}

	if err = d.Set("form_auth_fallback", flattenAuthenticationRuleFormAuthFallback(o["form-auth-fallback"], d, "form_auth_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["form-auth-fallback"], "AuthenticationRule-FormAuthFallback"); ok {
			if err = d.Set("form_auth_fallback", vv); err != nil {
				return fmt.Errorf("Error reading form_auth_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading form_auth_fallback: %v", err)
		}
	}

	if err = d.Set("web_proxy", flattenAuthenticationRuleWebProxy(o["web-proxy"], d, "web_proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-proxy"], "AuthenticationRule-WebProxy"); ok {
			if err = d.Set("web_proxy", vv); err != nil {
				return fmt.Errorf("Error reading web_proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_proxy: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAuthenticationRuleActiveAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleCertAuthCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleCorsDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleCorsStateful(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleIpBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleSsoAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAuthenticationRuleStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleTransactionBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleWebAuthCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleWebPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleFormAuthFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectAuthenticationRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("active_auth_method"); ok || d.HasChange("active_auth_method") {
		t, err := expandAuthenticationRuleActiveAuthMethod(d, v, "active_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("cert_auth_cookie"); ok || d.HasChange("cert_auth_cookie") {
		t, err := expandAuthenticationRuleCertAuthCookie(d, v, "cert_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandAuthenticationRuleComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("cors_depth"); ok || d.HasChange("cors_depth") {
		t, err := expandAuthenticationRuleCorsDepth(d, v, "cors_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-depth"] = t
		}
	}

	if v, ok := d.GetOk("cors_stateful"); ok || d.HasChange("cors_stateful") {
		t, err := expandAuthenticationRuleCorsStateful(d, v, "cors_stateful")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-stateful"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandAuthenticationRuleDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandAuthenticationRuleDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("ip_based"); ok || d.HasChange("ip_based") {
		t, err := expandAuthenticationRuleIpBased(d, v, "ip_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-based"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandAuthenticationRuleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandAuthenticationRuleProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandAuthenticationRuleSrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandAuthenticationRuleSrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandAuthenticationRuleSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_method"); ok || d.HasChange("sso_auth_method") {
		t, err := expandAuthenticationRuleSsoAuthMethod(d, v, "sso_auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-method"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandAuthenticationRuleStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transaction_based"); ok || d.HasChange("transaction_based") {
		t, err := expandAuthenticationRuleTransactionBased(d, v, "transaction_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transaction-based"] = t
		}
	}

	if v, ok := d.GetOk("web_auth_cookie"); ok || d.HasChange("web_auth_cookie") {
		t, err := expandAuthenticationRuleWebAuthCookie(d, v, "web_auth_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("web_portal"); ok || d.HasChange("web_portal") {
		t, err := expandAuthenticationRuleWebPortal(d, v, "web_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-portal"] = t
		}
	}

	if v, ok := d.GetOk("form_auth_fallback"); ok || d.HasChange("form_auth_fallback") {
		t, err := expandAuthenticationRuleFormAuthFallback(d, v, "form_auth_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-auth-fallback"] = t
		}
	}

	if v, ok := d.GetOk("web_proxy"); ok || d.HasChange("web_proxy") {
		t, err := expandAuthenticationRuleWebProxy(d, v, "web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-proxy"] = t
		}
	}

	return &obj, nil
}
