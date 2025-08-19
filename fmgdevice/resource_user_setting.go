// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure user authentication setting.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSettingUpdate,
		Read:   resourceUserSettingRead,
		Update: resourceUserSettingUpdate,
		Delete: resourceUserSettingDelete,

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
			"auth_blackout_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_http_basic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_invalid_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_portal_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"auth_secure_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_ssl_allow_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_ssl_max_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_ssl_sigalgs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_timeout_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_user_password_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"per_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_ses_timeout_act": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectUserSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateUserSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("UserSetting")

	return resourceUserSettingRead(d, m)
}

func resourceUserSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource from API: %v", err)
	}
	return nil
}

func flattenUserSettingAuthBlackoutTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSettingAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSettingAuthHttpBasic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthInvalidMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthLockoutDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthLockoutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthOnDemand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthPortalTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenUserSettingAuthPortsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserSetting-AuthPorts-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserSettingAuthPortsPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserSetting-AuthPorts-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenUserSettingAuthPortsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "UserSetting-AuthPorts-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserSettingAuthPortsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthPortsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthPortsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSecureHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSrcMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSslAllowRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSslMaxProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthSslSigalgs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthTimeoutType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSettingDefaultUserPasswordPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSettingPerPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSettingRadiusSesTimeoutAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_blackout_time", flattenUserSettingAuthBlackoutTime(o["auth-blackout-time"], d, "auth_blackout_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-blackout-time"], "UserSetting-AuthBlackoutTime"); ok {
			if err = d.Set("auth_blackout_time", vv); err != nil {
				return fmt.Errorf("Error reading auth_blackout_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_blackout_time: %v", err)
		}
	}

	if err = d.Set("auth_ca_cert", flattenUserSettingAuthCaCert(o["auth-ca-cert"], d, "auth_ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ca-cert"], "UserSetting-AuthCaCert"); ok {
			if err = d.Set("auth_ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ca_cert: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenUserSettingAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-cert"], "UserSetting-AuthCert"); ok {
			if err = d.Set("auth_cert", vv); err != nil {
				return fmt.Errorf("Error reading auth_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_http_basic", flattenUserSettingAuthHttpBasic(o["auth-http-basic"], d, "auth_http_basic")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-http-basic"], "UserSetting-AuthHttpBasic"); ok {
			if err = d.Set("auth_http_basic", vv); err != nil {
				return fmt.Errorf("Error reading auth_http_basic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_http_basic: %v", err)
		}
	}

	if err = d.Set("auth_invalid_max", flattenUserSettingAuthInvalidMax(o["auth-invalid-max"], d, "auth_invalid_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-invalid-max"], "UserSetting-AuthInvalidMax"); ok {
			if err = d.Set("auth_invalid_max", vv); err != nil {
				return fmt.Errorf("Error reading auth_invalid_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_invalid_max: %v", err)
		}
	}

	if err = d.Set("auth_lockout_duration", flattenUserSettingAuthLockoutDuration(o["auth-lockout-duration"], d, "auth_lockout_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-lockout-duration"], "UserSetting-AuthLockoutDuration"); ok {
			if err = d.Set("auth_lockout_duration", vv); err != nil {
				return fmt.Errorf("Error reading auth_lockout_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_lockout_duration: %v", err)
		}
	}

	if err = d.Set("auth_lockout_threshold", flattenUserSettingAuthLockoutThreshold(o["auth-lockout-threshold"], d, "auth_lockout_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-lockout-threshold"], "UserSetting-AuthLockoutThreshold"); ok {
			if err = d.Set("auth_lockout_threshold", vv); err != nil {
				return fmt.Errorf("Error reading auth_lockout_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("auth_on_demand", flattenUserSettingAuthOnDemand(o["auth-on-demand"], d, "auth_on_demand")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-on-demand"], "UserSetting-AuthOnDemand"); ok {
			if err = d.Set("auth_on_demand", vv); err != nil {
				return fmt.Errorf("Error reading auth_on_demand: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_on_demand: %v", err)
		}
	}

	if err = d.Set("auth_portal_timeout", flattenUserSettingAuthPortalTimeout(o["auth-portal-timeout"], d, "auth_portal_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal-timeout"], "UserSetting-AuthPortalTimeout"); ok {
			if err = d.Set("auth_portal_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["auth-ports"], "UserSetting-AuthPorts"); ok {
				if err = d.Set("auth_ports", vv); err != nil {
					return fmt.Errorf("Error reading auth_ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auth_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth_ports"); ok {
			if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["auth-ports"], "UserSetting-AuthPorts"); ok {
					if err = d.Set("auth_ports", vv); err != nil {
						return fmt.Errorf("Error reading auth_ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auth_ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_secure_http", flattenUserSettingAuthSecureHttp(o["auth-secure-http"], d, "auth_secure_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-secure-http"], "UserSetting-AuthSecureHttp"); ok {
			if err = d.Set("auth_secure_http", vv); err != nil {
				return fmt.Errorf("Error reading auth_secure_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_secure_http: %v", err)
		}
	}

	if err = d.Set("auth_src_mac", flattenUserSettingAuthSrcMac(o["auth-src-mac"], d, "auth_src_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-src-mac"], "UserSetting-AuthSrcMac"); ok {
			if err = d.Set("auth_src_mac", vv); err != nil {
				return fmt.Errorf("Error reading auth_src_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_src_mac: %v", err)
		}
	}

	if err = d.Set("auth_ssl_allow_renegotiation", flattenUserSettingAuthSslAllowRenegotiation(o["auth-ssl-allow-renegotiation"], d, "auth_ssl_allow_renegotiation")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ssl-allow-renegotiation"], "UserSetting-AuthSslAllowRenegotiation"); ok {
			if err = d.Set("auth_ssl_allow_renegotiation", vv); err != nil {
				return fmt.Errorf("Error reading auth_ssl_allow_renegotiation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ssl_allow_renegotiation: %v", err)
		}
	}

	if err = d.Set("auth_ssl_max_proto_version", flattenUserSettingAuthSslMaxProtoVersion(o["auth-ssl-max-proto-version"], d, "auth_ssl_max_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ssl-max-proto-version"], "UserSetting-AuthSslMaxProtoVersion"); ok {
			if err = d.Set("auth_ssl_max_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading auth_ssl_max_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ssl_max_proto_version: %v", err)
		}
	}

	if err = d.Set("auth_ssl_min_proto_version", flattenUserSettingAuthSslMinProtoVersion(o["auth-ssl-min-proto-version"], d, "auth_ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ssl-min-proto-version"], "UserSetting-AuthSslMinProtoVersion"); ok {
			if err = d.Set("auth_ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading auth_ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("auth_ssl_sigalgs", flattenUserSettingAuthSslSigalgs(o["auth-ssl-sigalgs"], d, "auth_ssl_sigalgs")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-ssl-sigalgs"], "UserSetting-AuthSslSigalgs"); ok {
			if err = d.Set("auth_ssl_sigalgs", vv); err != nil {
				return fmt.Errorf("Error reading auth_ssl_sigalgs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_ssl_sigalgs: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenUserSettingAuthTimeout(o["auth-timeout"], d, "auth_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-timeout"], "UserSetting-AuthTimeout"); ok {
			if err = d.Set("auth_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auth_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("auth_timeout_type", flattenUserSettingAuthTimeoutType(o["auth-timeout-type"], d, "auth_timeout_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-timeout-type"], "UserSetting-AuthTimeoutType"); ok {
			if err = d.Set("auth_timeout_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_timeout_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_timeout_type: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserSettingAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "UserSetting-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("default_user_password_policy", flattenUserSettingDefaultUserPasswordPolicy(o["default-user-password-policy"], d, "default_user_password_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-user-password-policy"], "UserSetting-DefaultUserPasswordPolicy"); ok {
			if err = d.Set("default_user_password_policy", vv); err != nil {
				return fmt.Errorf("Error reading default_user_password_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_user_password_policy: %v", err)
		}
	}

	if err = d.Set("per_policy_disclaimer", flattenUserSettingPerPolicyDisclaimer(o["per-policy-disclaimer"], d, "per_policy_disclaimer")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy-disclaimer"], "UserSetting-PerPolicyDisclaimer"); ok {
			if err = d.Set("per_policy_disclaimer", vv); err != nil {
				return fmt.Errorf("Error reading per_policy_disclaimer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy_disclaimer: %v", err)
		}
	}

	if err = d.Set("radius_ses_timeout_act", flattenUserSettingRadiusSesTimeoutAct(o["radius-ses-timeout-act"], d, "radius_ses_timeout_act")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-ses-timeout-act"], "UserSetting-RadiusSesTimeoutAct"); ok {
			if err = d.Set("radius_ses_timeout_act", vv); err != nil {
				return fmt.Errorf("Error reading radius_ses_timeout_act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_ses_timeout_act: %v", err)
		}
	}

	return nil
}

func flattenUserSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserSettingAuthBlackoutTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSettingAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSettingAuthHttpBasic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthInvalidMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthLockoutDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthLockoutThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthOnDemand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortalTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandUserSettingAuthPortsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserSettingAuthPortsPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandUserSettingAuthPortsType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserSettingAuthPortsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSecureHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSrcMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslAllowRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslMaxProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslSigalgs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeoutType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSettingDefaultUserPasswordPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSettingPerPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSettingRadiusSesTimeoutAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_blackout_time"); ok || d.HasChange("auth_blackout_time") {
		t, err := expandUserSettingAuthBlackoutTime(d, v, "auth_blackout_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-blackout-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_ca_cert"); ok || d.HasChange("auth_ca_cert") {
		t, err := expandUserSettingAuthCaCert(d, v, "auth_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok || d.HasChange("auth_cert") {
		t, err := expandUserSettingAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_http_basic"); ok || d.HasChange("auth_http_basic") {
		t, err := expandUserSettingAuthHttpBasic(d, v, "auth_http_basic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-http-basic"] = t
		}
	}

	if v, ok := d.GetOk("auth_invalid_max"); ok || d.HasChange("auth_invalid_max") {
		t, err := expandUserSettingAuthInvalidMax(d, v, "auth_invalid_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-invalid-max"] = t
		}
	}

	if v, ok := d.GetOk("auth_lockout_duration"); ok || d.HasChange("auth_lockout_duration") {
		t, err := expandUserSettingAuthLockoutDuration(d, v, "auth_lockout_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOk("auth_lockout_threshold"); ok || d.HasChange("auth_lockout_threshold") {
		t, err := expandUserSettingAuthLockoutThreshold(d, v, "auth_lockout_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOk("auth_on_demand"); ok || d.HasChange("auth_on_demand") {
		t, err := expandUserSettingAuthOnDemand(d, v, "auth_on_demand")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-on-demand"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_timeout"); ok || d.HasChange("auth_portal_timeout") {
		t, err := expandUserSettingAuthPortalTimeout(d, v, "auth_portal_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auth_ports"); ok || d.HasChange("auth_ports") {
		t, err := expandUserSettingAuthPorts(d, v, "auth_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ports"] = t
		}
	}

	if v, ok := d.GetOk("auth_secure_http"); ok || d.HasChange("auth_secure_http") {
		t, err := expandUserSettingAuthSecureHttp(d, v, "auth_secure_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-secure-http"] = t
		}
	}

	if v, ok := d.GetOk("auth_src_mac"); ok || d.HasChange("auth_src_mac") {
		t, err := expandUserSettingAuthSrcMac(d, v, "auth_src_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-src-mac"] = t
		}
	}

	if v, ok := d.GetOk("auth_ssl_allow_renegotiation"); ok || d.HasChange("auth_ssl_allow_renegotiation") {
		t, err := expandUserSettingAuthSslAllowRenegotiation(d, v, "auth_ssl_allow_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-allow-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("auth_ssl_max_proto_version"); ok || d.HasChange("auth_ssl_max_proto_version") {
		t, err := expandUserSettingAuthSslMaxProtoVersion(d, v, "auth_ssl_max_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-max-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("auth_ssl_min_proto_version"); ok || d.HasChange("auth_ssl_min_proto_version") {
		t, err := expandUserSettingAuthSslMinProtoVersion(d, v, "auth_ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("auth_ssl_sigalgs"); ok || d.HasChange("auth_ssl_sigalgs") {
		t, err := expandUserSettingAuthSslSigalgs(d, v, "auth_ssl_sigalgs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-sigalgs"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok || d.HasChange("auth_timeout") {
		t, err := expandUserSettingAuthTimeout(d, v, "auth_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout_type"); ok || d.HasChange("auth_timeout_type") {
		t, err := expandUserSettingAuthTimeoutType(d, v, "auth_timeout_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandUserSettingAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("default_user_password_policy"); ok || d.HasChange("default_user_password_policy") {
		t, err := expandUserSettingDefaultUserPasswordPolicy(d, v, "default_user_password_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-user-password-policy"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_disclaimer"); ok || d.HasChange("per_policy_disclaimer") {
		t, err := expandUserSettingPerPolicyDisclaimer(d, v, "per_policy_disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("radius_ses_timeout_act"); ok || d.HasChange("radius_ses_timeout_act") {
		t, err := expandUserSettingRadiusSesTimeoutAct(d, v, "radius_ses_timeout_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-ses-timeout-act"] = t
		}
	}

	return &obj, nil
}
