// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebcacheSettings

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebcacheSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebcacheSettingsUpdate,
		Read:   resourceWebcacheSettingsRead,
		Update: resourceWebcacheSettingsUpdate,
		Delete: resourceWebcacheSettingsDelete,

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
			"add_x_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"always_revalidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_by_default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_expired": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crawler_user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fresh_factor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"host_validate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_conditional": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_ie_reload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_ims": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_pnc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_object_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"neg_resp_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reval_pnc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x_cache_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebcacheSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebcacheSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating WebcacheSettings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebcacheSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebcacheSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebcacheSettings")

	return resourceWebcacheSettingsRead(d, m)
}

func resourceWebcacheSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebcacheSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebcacheSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebcacheSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebcacheSettings(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebcacheSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebcacheSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebcacheSettings resource from API: %v", err)
	}
	return nil
}

func flattenWebcacheSettingsAddXCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsAlwaysRevalidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsCacheByDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsCacheCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsCacheExpired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsCrawlerUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsDefaultTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsFreshFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsHostValidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsIgnoreConditional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsIgnoreIeReload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsIgnoreIms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsIgnorePnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsMaxObjectSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsMaxTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsMinTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsNegRespTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsRevalPnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheSettingsXCacheMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebcacheSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("add_x_cache", flattenWebcacheSettingsAddXCache(o["add-x-cache"], d, "add_x_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-x-cache"], "WebcacheSettings-AddXCache"); ok {
			if err = d.Set("add_x_cache", vv); err != nil {
				return fmt.Errorf("Error reading add_x_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_x_cache: %v", err)
		}
	}

	if err = d.Set("always_revalidate", flattenWebcacheSettingsAlwaysRevalidate(o["always-revalidate"], d, "always_revalidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["always-revalidate"], "WebcacheSettings-AlwaysRevalidate"); ok {
			if err = d.Set("always_revalidate", vv); err != nil {
				return fmt.Errorf("Error reading always_revalidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading always_revalidate: %v", err)
		}
	}

	if err = d.Set("cache_by_default", flattenWebcacheSettingsCacheByDefault(o["cache-by-default"], d, "cache_by_default")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-by-default"], "WebcacheSettings-CacheByDefault"); ok {
			if err = d.Set("cache_by_default", vv); err != nil {
				return fmt.Errorf("Error reading cache_by_default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_by_default: %v", err)
		}
	}

	if err = d.Set("cache_cookie", flattenWebcacheSettingsCacheCookie(o["cache-cookie"], d, "cache_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-cookie"], "WebcacheSettings-CacheCookie"); ok {
			if err = d.Set("cache_cookie", vv); err != nil {
				return fmt.Errorf("Error reading cache_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_cookie: %v", err)
		}
	}

	if err = d.Set("cache_expired", flattenWebcacheSettingsCacheExpired(o["cache-expired"], d, "cache_expired")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-expired"], "WebcacheSettings-CacheExpired"); ok {
			if err = d.Set("cache_expired", vv); err != nil {
				return fmt.Errorf("Error reading cache_expired: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_expired: %v", err)
		}
	}

	if err = d.Set("crawler_user_agent", flattenWebcacheSettingsCrawlerUserAgent(o["crawler-user-agent"], d, "crawler_user_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["crawler-user-agent"], "WebcacheSettings-CrawlerUserAgent"); ok {
			if err = d.Set("crawler_user_agent", vv); err != nil {
				return fmt.Errorf("Error reading crawler_user_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading crawler_user_agent: %v", err)
		}
	}

	if err = d.Set("default_ttl", flattenWebcacheSettingsDefaultTtl(o["default-ttl"], d, "default_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-ttl"], "WebcacheSettings-DefaultTtl"); ok {
			if err = d.Set("default_ttl", vv); err != nil {
				return fmt.Errorf("Error reading default_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_ttl: %v", err)
		}
	}

	if err = d.Set("external", flattenWebcacheSettingsExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "WebcacheSettings-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("fresh_factor", flattenWebcacheSettingsFreshFactor(o["fresh-factor"], d, "fresh_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["fresh-factor"], "WebcacheSettings-FreshFactor"); ok {
			if err = d.Set("fresh_factor", vv); err != nil {
				return fmt.Errorf("Error reading fresh_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fresh_factor: %v", err)
		}
	}

	if err = d.Set("host_validate", flattenWebcacheSettingsHostValidate(o["host-validate"], d, "host_validate")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-validate"], "WebcacheSettings-HostValidate"); ok {
			if err = d.Set("host_validate", vv); err != nil {
				return fmt.Errorf("Error reading host_validate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_validate: %v", err)
		}
	}

	if err = d.Set("ignore_conditional", flattenWebcacheSettingsIgnoreConditional(o["ignore-conditional"], d, "ignore_conditional")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-conditional"], "WebcacheSettings-IgnoreConditional"); ok {
			if err = d.Set("ignore_conditional", vv); err != nil {
				return fmt.Errorf("Error reading ignore_conditional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_conditional: %v", err)
		}
	}

	if err = d.Set("ignore_ie_reload", flattenWebcacheSettingsIgnoreIeReload(o["ignore-ie-reload"], d, "ignore_ie_reload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-ie-reload"], "WebcacheSettings-IgnoreIeReload"); ok {
			if err = d.Set("ignore_ie_reload", vv); err != nil {
				return fmt.Errorf("Error reading ignore_ie_reload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_ie_reload: %v", err)
		}
	}

	if err = d.Set("ignore_ims", flattenWebcacheSettingsIgnoreIms(o["ignore-ims"], d, "ignore_ims")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-ims"], "WebcacheSettings-IgnoreIms"); ok {
			if err = d.Set("ignore_ims", vv); err != nil {
				return fmt.Errorf("Error reading ignore_ims: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_ims: %v", err)
		}
	}

	if err = d.Set("ignore_pnc", flattenWebcacheSettingsIgnorePnc(o["ignore-pnc"], d, "ignore_pnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-pnc"], "WebcacheSettings-IgnorePnc"); ok {
			if err = d.Set("ignore_pnc", vv); err != nil {
				return fmt.Errorf("Error reading ignore_pnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_pnc: %v", err)
		}
	}

	if err = d.Set("max_object_size", flattenWebcacheSettingsMaxObjectSize(o["max-object-size"], d, "max_object_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-object-size"], "WebcacheSettings-MaxObjectSize"); ok {
			if err = d.Set("max_object_size", vv); err != nil {
				return fmt.Errorf("Error reading max_object_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_object_size: %v", err)
		}
	}

	if err = d.Set("max_ttl", flattenWebcacheSettingsMaxTtl(o["max-ttl"], d, "max_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ttl"], "WebcacheSettings-MaxTtl"); ok {
			if err = d.Set("max_ttl", vv); err != nil {
				return fmt.Errorf("Error reading max_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ttl: %v", err)
		}
	}

	if err = d.Set("min_ttl", flattenWebcacheSettingsMinTtl(o["min-ttl"], d, "min_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-ttl"], "WebcacheSettings-MinTtl"); ok {
			if err = d.Set("min_ttl", vv); err != nil {
				return fmt.Errorf("Error reading min_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_ttl: %v", err)
		}
	}

	if err = d.Set("neg_resp_time", flattenWebcacheSettingsNegRespTime(o["neg-resp-time"], d, "neg_resp_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neg-resp-time"], "WebcacheSettings-NegRespTime"); ok {
			if err = d.Set("neg_resp_time", vv); err != nil {
				return fmt.Errorf("Error reading neg_resp_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neg_resp_time: %v", err)
		}
	}

	if err = d.Set("reval_pnc", flattenWebcacheSettingsRevalPnc(o["reval-pnc"], d, "reval_pnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["reval-pnc"], "WebcacheSettings-RevalPnc"); ok {
			if err = d.Set("reval_pnc", vv); err != nil {
				return fmt.Errorf("Error reading reval_pnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reval_pnc: %v", err)
		}
	}

	if err = d.Set("x_cache_message", flattenWebcacheSettingsXCacheMessage(o["x-cache-message"], d, "x_cache_message")); err != nil {
		if vv, ok := fortiAPIPatch(o["x-cache-message"], "WebcacheSettings-XCacheMessage"); ok {
			if err = d.Set("x_cache_message", vv); err != nil {
				return fmt.Errorf("Error reading x_cache_message: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading x_cache_message: %v", err)
		}
	}

	return nil
}

func flattenWebcacheSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebcacheSettingsAddXCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsAlwaysRevalidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsCacheByDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsCacheCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsCacheExpired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsCrawlerUserAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsDefaultTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsFreshFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsHostValidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsIgnoreConditional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsIgnoreIeReload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsIgnoreIms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsIgnorePnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsMaxObjectSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsMaxTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsMinTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsNegRespTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsRevalPnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheSettingsXCacheMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebcacheSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_x_cache"); ok || d.HasChange("add_x_cache") {
		t, err := expandWebcacheSettingsAddXCache(d, v, "add_x_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-x-cache"] = t
		}
	}

	if v, ok := d.GetOk("always_revalidate"); ok || d.HasChange("always_revalidate") {
		t, err := expandWebcacheSettingsAlwaysRevalidate(d, v, "always_revalidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-revalidate"] = t
		}
	}

	if v, ok := d.GetOk("cache_by_default"); ok || d.HasChange("cache_by_default") {
		t, err := expandWebcacheSettingsCacheByDefault(d, v, "cache_by_default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-by-default"] = t
		}
	}

	if v, ok := d.GetOk("cache_cookie"); ok || d.HasChange("cache_cookie") {
		t, err := expandWebcacheSettingsCacheCookie(d, v, "cache_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-cookie"] = t
		}
	}

	if v, ok := d.GetOk("cache_expired"); ok || d.HasChange("cache_expired") {
		t, err := expandWebcacheSettingsCacheExpired(d, v, "cache_expired")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-expired"] = t
		}
	}

	if v, ok := d.GetOk("crawler_user_agent"); ok || d.HasChange("crawler_user_agent") {
		t, err := expandWebcacheSettingsCrawlerUserAgent(d, v, "crawler_user_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crawler-user-agent"] = t
		}
	}

	if v, ok := d.GetOk("default_ttl"); ok || d.HasChange("default_ttl") {
		t, err := expandWebcacheSettingsDefaultTtl(d, v, "default_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-ttl"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandWebcacheSettingsExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("fresh_factor"); ok || d.HasChange("fresh_factor") {
		t, err := expandWebcacheSettingsFreshFactor(d, v, "fresh_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fresh-factor"] = t
		}
	}

	if v, ok := d.GetOk("host_validate"); ok || d.HasChange("host_validate") {
		t, err := expandWebcacheSettingsHostValidate(d, v, "host_validate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-validate"] = t
		}
	}

	if v, ok := d.GetOk("ignore_conditional"); ok || d.HasChange("ignore_conditional") {
		t, err := expandWebcacheSettingsIgnoreConditional(d, v, "ignore_conditional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-conditional"] = t
		}
	}

	if v, ok := d.GetOk("ignore_ie_reload"); ok || d.HasChange("ignore_ie_reload") {
		t, err := expandWebcacheSettingsIgnoreIeReload(d, v, "ignore_ie_reload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-ie-reload"] = t
		}
	}

	if v, ok := d.GetOk("ignore_ims"); ok || d.HasChange("ignore_ims") {
		t, err := expandWebcacheSettingsIgnoreIms(d, v, "ignore_ims")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-ims"] = t
		}
	}

	if v, ok := d.GetOk("ignore_pnc"); ok || d.HasChange("ignore_pnc") {
		t, err := expandWebcacheSettingsIgnorePnc(d, v, "ignore_pnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-pnc"] = t
		}
	}

	if v, ok := d.GetOk("max_object_size"); ok || d.HasChange("max_object_size") {
		t, err := expandWebcacheSettingsMaxObjectSize(d, v, "max_object_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-object-size"] = t
		}
	}

	if v, ok := d.GetOk("max_ttl"); ok || d.HasChange("max_ttl") {
		t, err := expandWebcacheSettingsMaxTtl(d, v, "max_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ttl"] = t
		}
	}

	if v, ok := d.GetOk("min_ttl"); ok || d.HasChange("min_ttl") {
		t, err := expandWebcacheSettingsMinTtl(d, v, "min_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-ttl"] = t
		}
	}

	if v, ok := d.GetOk("neg_resp_time"); ok || d.HasChange("neg_resp_time") {
		t, err := expandWebcacheSettingsNegRespTime(d, v, "neg_resp_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neg-resp-time"] = t
		}
	}

	if v, ok := d.GetOk("reval_pnc"); ok || d.HasChange("reval_pnc") {
		t, err := expandWebcacheSettingsRevalPnc(d, v, "reval_pnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reval-pnc"] = t
		}
	}

	if v, ok := d.GetOk("x_cache_message"); ok || d.HasChange("x_cache_message") {
		t, err := expandWebcacheSettingsXCacheMessage(d, v, "x_cache_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["x-cache-message"] = t
		}
	}

	return &obj, nil
}
