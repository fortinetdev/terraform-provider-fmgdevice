// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ztna web-portal.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalCreate,
		Read:   resourceZtnaWebPortalRead,
		Update: resourceZtnaWebPortalUpdate,
		Delete: resourceZtnaWebPortalDelete,

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
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_virtual_host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_window_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_window_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policy_auth_sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vip6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZtnaWebPortalCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortal resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebPortal(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortal resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalRead(d, m)
}

func resourceZtnaWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectZtnaWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortal resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebPortal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaWebPortalRead(d, m)
}

func resourceZtnaWebPortalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteZtnaWebPortal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebPortal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalAuthPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalAuthRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalCookieAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return fmt.Sprintf("%v", v)
}

func flattenZtnaWebPortalHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalPolicyAuthSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebPortalVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalVip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebPortal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_portal", flattenZtnaWebPortalAuthPortal(o["auth-portal"], d, "auth_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-portal"], "ZtnaWebPortal-AuthPortal"); ok {
			if err = d.Set("auth_portal", vv); err != nil {
				return fmt.Errorf("Error reading auth_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_rule", flattenZtnaWebPortalAuthRule(o["auth-rule"], d, "auth_rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-rule"], "ZtnaWebPortal-AuthRule"); ok {
			if err = d.Set("auth_rule", vv); err != nil {
				return fmt.Errorf("Error reading auth_rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_rule: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenZtnaWebPortalAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-virtual-host"], "ZtnaWebPortal-AuthVirtualHost"); ok {
			if err = d.Set("auth_virtual_host", vv); err != nil {
				return fmt.Errorf("Error reading auth_virtual_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenZtnaWebPortalClipboard(o["clipboard"], d, "clipboard")); err != nil {
		if vv, ok := fortiAPIPatch(o["clipboard"], "ZtnaWebPortal-Clipboard"); ok {
			if err = d.Set("clipboard", vv); err != nil {
				return fmt.Errorf("Error reading clipboard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("cookie_age", flattenZtnaWebPortalCookieAge(o["cookie-age"], d, "cookie_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["cookie-age"], "ZtnaWebPortal-CookieAge"); ok {
			if err = d.Set("cookie_age", vv); err != nil {
				return fmt.Errorf("Error reading cookie_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cookie_age: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenZtnaWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["customize-forticlient-download-url"], "ZtnaWebPortal-CustomizeForticlientDownloadUrl"); ok {
			if err = d.Set("customize_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenZtnaWebPortalDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if vv, ok := fortiAPIPatch(o["decrypted-traffic-mirror"], "ZtnaWebPortal-DecryptedTrafficMirror"); ok {
			if err = d.Set("decrypted_traffic_mirror", vv); err != nil {
				return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenZtnaWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-height"], "ZtnaWebPortal-DefaultWindowHeight"); ok {
			if err = d.Set("default_window_height", vv); err != nil {
				return fmt.Errorf("Error reading default_window_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenZtnaWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-width"], "ZtnaWebPortal-DefaultWindowWidth"); ok {
			if err = d.Set("default_window_width", vv); err != nil {
				return fmt.Errorf("Error reading default_window_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenZtnaWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-bookmark"], "ZtnaWebPortal-DisplayBookmark"); ok {
			if err = d.Set("display_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading display_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("display_history", flattenZtnaWebPortalDisplayHistory(o["display-history"], d, "display_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-history"], "ZtnaWebPortal-DisplayHistory"); ok {
			if err = d.Set("display_history", vv); err != nil {
				return fmt.Errorf("Error reading display_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("display_status", flattenZtnaWebPortalDisplayStatus(o["display-status"], d, "display_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-status"], "ZtnaWebPortal-DisplayStatus"); ok {
			if err = d.Set("display_status", vv); err != nil {
				return fmt.Errorf("Error reading display_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenZtnaWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["focus-bookmark"], "ZtnaWebPortal-FocusBookmark"); ok {
			if err = d.Set("focus_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading focus_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("forticlient_download", flattenZtnaWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download"], "ZtnaWebPortal-ForticlientDownload"); ok {
			if err = d.Set("forticlient_download", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenZtnaWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download-method"], "ZtnaWebPortal-ForticlientDownloadMethod"); ok {
			if err = d.Set("forticlient_download_method", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("heading", flattenZtnaWebPortalHeading(o["heading"], d, "heading")); err != nil {
		if vv, ok := fortiAPIPatch(o["heading"], "ZtnaWebPortal-Heading"); ok {
			if err = d.Set("heading", vv); err != nil {
				return fmt.Errorf("Error reading heading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaWebPortalHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ZtnaWebPortal-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaWebPortalLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-blocked-traffic"], "ZtnaWebPortal-LogBlockedTraffic"); ok {
			if err = d.Set("log_blocked_traffic", vv); err != nil {
				return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenZtnaWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["macos-forticlient-download-url"], "ZtnaWebPortal-MacosForticlientDownloadUrl"); ok {
			if err = d.Set("macos_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("name", flattenZtnaWebPortalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaWebPortal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policy_auth_sso", flattenZtnaWebPortalPolicyAuthSso(o["policy-auth-sso"], d, "policy_auth_sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-auth-sso"], "ZtnaWebPortal-PolicyAuthSso"); ok {
			if err = d.Set("policy_auth_sso", vv); err != nil {
				return fmt.Errorf("Error reading policy_auth_sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_auth_sso: %v", err)
		}
	}

	if err = d.Set("theme", flattenZtnaWebPortalTheme(o["theme"], d, "theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["theme"], "ZtnaWebPortal-Theme"); ok {
			if err = d.Set("theme", vv); err != nil {
				return fmt.Errorf("Error reading theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("vip", flattenZtnaWebPortalVip(o["vip"], d, "vip")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip"], "ZtnaWebPortal-Vip"); ok {
			if err = d.Set("vip", vv); err != nil {
				return fmt.Errorf("Error reading vip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("vip6", flattenZtnaWebPortalVip6(o["vip6"], d, "vip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vip6"], "ZtnaWebPortal-Vip6"); ok {
			if err = d.Set("vip6", vv); err != nil {
				return fmt.Errorf("Error reading vip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenZtnaWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["windows-forticlient-download-url"], "ZtnaWebPortal-WindowsForticlientDownloadUrl"); ok {
			if err = d.Set("windows_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebPortalAuthPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalAuthRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalCookieAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return fmt.Sprintf("%v", v), nil
}

func expandZtnaWebPortalHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalPolicyAuthSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalVip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebPortal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_portal"); ok || d.HasChange("auth_portal") {
		t, err := expandZtnaWebPortalAuthPortal(d, v, "auth_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_rule"); ok || d.HasChange("auth_rule") {
		t, err := expandZtnaWebPortalAuthRule(d, v, "auth_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-rule"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok || d.HasChange("auth_virtual_host") {
		t, err := expandZtnaWebPortalAuthVirtualHost(d, v, "auth_virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok || d.HasChange("clipboard") {
		t, err := expandZtnaWebPortalClipboard(d, v, "clipboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOk("cookie_age"); ok || d.HasChange("cookie_age") {
		t, err := expandZtnaWebPortalCookieAge(d, v, "cookie_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok || d.HasChange("customize_forticlient_download_url") {
		t, err := expandZtnaWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok || d.HasChange("decrypted_traffic_mirror") {
		t, err := expandZtnaWebPortalDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("default_window_height"); ok || d.HasChange("default_window_height") {
		t, err := expandZtnaWebPortalDefaultWindowHeight(d, v, "default_window_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOk("default_window_width"); ok || d.HasChange("default_window_width") {
		t, err := expandZtnaWebPortalDefaultWindowWidth(d, v, "default_window_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOk("display_bookmark"); ok || d.HasChange("display_bookmark") {
		t, err := expandZtnaWebPortalDisplayBookmark(d, v, "display_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok || d.HasChange("display_history") {
		t, err := expandZtnaWebPortalDisplayHistory(d, v, "display_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok || d.HasChange("display_status") {
		t, err := expandZtnaWebPortalDisplayStatus(d, v, "display_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok || d.HasChange("focus_bookmark") {
		t, err := expandZtnaWebPortalFocusBookmark(d, v, "focus_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok || d.HasChange("forticlient_download") {
		t, err := expandZtnaWebPortalForticlientDownload(d, v, "forticlient_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok || d.HasChange("forticlient_download_method") {
		t, err := expandZtnaWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok || d.HasChange("heading") {
		t, err := expandZtnaWebPortalHeading(d, v, "heading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandZtnaWebPortalHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok || d.HasChange("log_blocked_traffic") {
		t, err := expandZtnaWebPortalLogBlockedTraffic(d, v, "log_blocked_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok || d.HasChange("macos_forticlient_download_url") {
		t, err := expandZtnaWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaWebPortalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy_auth_sso"); ok || d.HasChange("policy_auth_sso") {
		t, err := expandZtnaWebPortalPolicyAuthSso(d, v, "policy_auth_sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-auth-sso"] = t
		}
	}

	if v, ok := d.GetOk("theme"); ok || d.HasChange("theme") {
		t, err := expandZtnaWebPortalTheme(d, v, "theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok || d.HasChange("vip") {
		t, err := expandZtnaWebPortalVip(d, v, "vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("vip6"); ok || d.HasChange("vip6") {
		t, err := expandZtnaWebPortalVip6(d, v, "vip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok || d.HasChange("windows_forticlient_download_url") {
		t, err := expandZtnaWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	}

	return &obj, nil
}
