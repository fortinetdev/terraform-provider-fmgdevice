// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure global Web cache settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptWebcache() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptWebcacheUpdate,
		Read:   resourceWanoptWebcacheRead,
		Update: resourceWanoptWebcacheUpdate,
		Delete: resourceWanoptWebcacheDelete,

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
			"always_revalidate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourceWanoptWebcacheUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWanoptWebcache(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptWebcache resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptWebcache(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptWebcache resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptWebcache")

	return resourceWanoptWebcacheRead(d, m)
}

func resourceWanoptWebcacheDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptWebcache(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptWebcache resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptWebcacheRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptWebcache(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptWebcache resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptWebcache(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptWebcache resource from API: %v", err)
	}
	return nil
}

func flattenWanoptWebcacheAlwaysRevalidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheCacheByDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheCacheCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheCacheExpired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheDefaultTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheFreshFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheHostValidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheIgnoreConditional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheIgnoreIeReload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheIgnoreIms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheIgnorePnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheMaxObjectSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheMaxTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheMinTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheNegRespTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptWebcacheRevalPnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptWebcache(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("always_revalidate", flattenWanoptWebcacheAlwaysRevalidate(o["always-revalidate"], d, "always_revalidate")); err != nil {
		if vv, ok := fortiAPIPatch(o["always-revalidate"], "WanoptWebcache-AlwaysRevalidate"); ok {
			if err = d.Set("always_revalidate", vv); err != nil {
				return fmt.Errorf("Error reading always_revalidate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading always_revalidate: %v", err)
		}
	}

	if err = d.Set("cache_by_default", flattenWanoptWebcacheCacheByDefault(o["cache-by-default"], d, "cache_by_default")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-by-default"], "WanoptWebcache-CacheByDefault"); ok {
			if err = d.Set("cache_by_default", vv); err != nil {
				return fmt.Errorf("Error reading cache_by_default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_by_default: %v", err)
		}
	}

	if err = d.Set("cache_cookie", flattenWanoptWebcacheCacheCookie(o["cache-cookie"], d, "cache_cookie")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-cookie"], "WanoptWebcache-CacheCookie"); ok {
			if err = d.Set("cache_cookie", vv); err != nil {
				return fmt.Errorf("Error reading cache_cookie: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_cookie: %v", err)
		}
	}

	if err = d.Set("cache_expired", flattenWanoptWebcacheCacheExpired(o["cache-expired"], d, "cache_expired")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-expired"], "WanoptWebcache-CacheExpired"); ok {
			if err = d.Set("cache_expired", vv); err != nil {
				return fmt.Errorf("Error reading cache_expired: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_expired: %v", err)
		}
	}

	if err = d.Set("default_ttl", flattenWanoptWebcacheDefaultTtl(o["default-ttl"], d, "default_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-ttl"], "WanoptWebcache-DefaultTtl"); ok {
			if err = d.Set("default_ttl", vv); err != nil {
				return fmt.Errorf("Error reading default_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_ttl: %v", err)
		}
	}

	if err = d.Set("external", flattenWanoptWebcacheExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "WanoptWebcache-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("fresh_factor", flattenWanoptWebcacheFreshFactor(o["fresh-factor"], d, "fresh_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["fresh-factor"], "WanoptWebcache-FreshFactor"); ok {
			if err = d.Set("fresh_factor", vv); err != nil {
				return fmt.Errorf("Error reading fresh_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fresh_factor: %v", err)
		}
	}

	if err = d.Set("host_validate", flattenWanoptWebcacheHostValidate(o["host-validate"], d, "host_validate")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-validate"], "WanoptWebcache-HostValidate"); ok {
			if err = d.Set("host_validate", vv); err != nil {
				return fmt.Errorf("Error reading host_validate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_validate: %v", err)
		}
	}

	if err = d.Set("ignore_conditional", flattenWanoptWebcacheIgnoreConditional(o["ignore-conditional"], d, "ignore_conditional")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-conditional"], "WanoptWebcache-IgnoreConditional"); ok {
			if err = d.Set("ignore_conditional", vv); err != nil {
				return fmt.Errorf("Error reading ignore_conditional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_conditional: %v", err)
		}
	}

	if err = d.Set("ignore_ie_reload", flattenWanoptWebcacheIgnoreIeReload(o["ignore-ie-reload"], d, "ignore_ie_reload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-ie-reload"], "WanoptWebcache-IgnoreIeReload"); ok {
			if err = d.Set("ignore_ie_reload", vv); err != nil {
				return fmt.Errorf("Error reading ignore_ie_reload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_ie_reload: %v", err)
		}
	}

	if err = d.Set("ignore_ims", flattenWanoptWebcacheIgnoreIms(o["ignore-ims"], d, "ignore_ims")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-ims"], "WanoptWebcache-IgnoreIms"); ok {
			if err = d.Set("ignore_ims", vv); err != nil {
				return fmt.Errorf("Error reading ignore_ims: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_ims: %v", err)
		}
	}

	if err = d.Set("ignore_pnc", flattenWanoptWebcacheIgnorePnc(o["ignore-pnc"], d, "ignore_pnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-pnc"], "WanoptWebcache-IgnorePnc"); ok {
			if err = d.Set("ignore_pnc", vv); err != nil {
				return fmt.Errorf("Error reading ignore_pnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_pnc: %v", err)
		}
	}

	if err = d.Set("max_object_size", flattenWanoptWebcacheMaxObjectSize(o["max-object-size"], d, "max_object_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-object-size"], "WanoptWebcache-MaxObjectSize"); ok {
			if err = d.Set("max_object_size", vv); err != nil {
				return fmt.Errorf("Error reading max_object_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_object_size: %v", err)
		}
	}

	if err = d.Set("max_ttl", flattenWanoptWebcacheMaxTtl(o["max-ttl"], d, "max_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-ttl"], "WanoptWebcache-MaxTtl"); ok {
			if err = d.Set("max_ttl", vv); err != nil {
				return fmt.Errorf("Error reading max_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_ttl: %v", err)
		}
	}

	if err = d.Set("min_ttl", flattenWanoptWebcacheMinTtl(o["min-ttl"], d, "min_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-ttl"], "WanoptWebcache-MinTtl"); ok {
			if err = d.Set("min_ttl", vv); err != nil {
				return fmt.Errorf("Error reading min_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_ttl: %v", err)
		}
	}

	if err = d.Set("neg_resp_time", flattenWanoptWebcacheNegRespTime(o["neg-resp-time"], d, "neg_resp_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neg-resp-time"], "WanoptWebcache-NegRespTime"); ok {
			if err = d.Set("neg_resp_time", vv); err != nil {
				return fmt.Errorf("Error reading neg_resp_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neg_resp_time: %v", err)
		}
	}

	if err = d.Set("reval_pnc", flattenWanoptWebcacheRevalPnc(o["reval-pnc"], d, "reval_pnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["reval-pnc"], "WanoptWebcache-RevalPnc"); ok {
			if err = d.Set("reval_pnc", vv); err != nil {
				return fmt.Errorf("Error reading reval_pnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reval_pnc: %v", err)
		}
	}

	return nil
}

func flattenWanoptWebcacheFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptWebcacheAlwaysRevalidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheCacheByDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheCacheCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheCacheExpired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheDefaultTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheFreshFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheHostValidate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheIgnoreConditional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheIgnoreIeReload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheIgnoreIms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheIgnorePnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheMaxObjectSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheMaxTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheMinTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheNegRespTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptWebcacheRevalPnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptWebcache(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("always_revalidate"); ok || d.HasChange("always_revalidate") {
		t, err := expandWanoptWebcacheAlwaysRevalidate(d, v, "always_revalidate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-revalidate"] = t
		}
	}

	if v, ok := d.GetOk("cache_by_default"); ok || d.HasChange("cache_by_default") {
		t, err := expandWanoptWebcacheCacheByDefault(d, v, "cache_by_default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-by-default"] = t
		}
	}

	if v, ok := d.GetOk("cache_cookie"); ok || d.HasChange("cache_cookie") {
		t, err := expandWanoptWebcacheCacheCookie(d, v, "cache_cookie")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-cookie"] = t
		}
	}

	if v, ok := d.GetOk("cache_expired"); ok || d.HasChange("cache_expired") {
		t, err := expandWanoptWebcacheCacheExpired(d, v, "cache_expired")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-expired"] = t
		}
	}

	if v, ok := d.GetOk("default_ttl"); ok || d.HasChange("default_ttl") {
		t, err := expandWanoptWebcacheDefaultTtl(d, v, "default_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-ttl"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandWanoptWebcacheExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("fresh_factor"); ok || d.HasChange("fresh_factor") {
		t, err := expandWanoptWebcacheFreshFactor(d, v, "fresh_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fresh-factor"] = t
		}
	}

	if v, ok := d.GetOk("host_validate"); ok || d.HasChange("host_validate") {
		t, err := expandWanoptWebcacheHostValidate(d, v, "host_validate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-validate"] = t
		}
	}

	if v, ok := d.GetOk("ignore_conditional"); ok || d.HasChange("ignore_conditional") {
		t, err := expandWanoptWebcacheIgnoreConditional(d, v, "ignore_conditional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-conditional"] = t
		}
	}

	if v, ok := d.GetOk("ignore_ie_reload"); ok || d.HasChange("ignore_ie_reload") {
		t, err := expandWanoptWebcacheIgnoreIeReload(d, v, "ignore_ie_reload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-ie-reload"] = t
		}
	}

	if v, ok := d.GetOk("ignore_ims"); ok || d.HasChange("ignore_ims") {
		t, err := expandWanoptWebcacheIgnoreIms(d, v, "ignore_ims")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-ims"] = t
		}
	}

	if v, ok := d.GetOk("ignore_pnc"); ok || d.HasChange("ignore_pnc") {
		t, err := expandWanoptWebcacheIgnorePnc(d, v, "ignore_pnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-pnc"] = t
		}
	}

	if v, ok := d.GetOk("max_object_size"); ok || d.HasChange("max_object_size") {
		t, err := expandWanoptWebcacheMaxObjectSize(d, v, "max_object_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-object-size"] = t
		}
	}

	if v, ok := d.GetOk("max_ttl"); ok || d.HasChange("max_ttl") {
		t, err := expandWanoptWebcacheMaxTtl(d, v, "max_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-ttl"] = t
		}
	}

	if v, ok := d.GetOk("min_ttl"); ok || d.HasChange("min_ttl") {
		t, err := expandWanoptWebcacheMinTtl(d, v, "min_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-ttl"] = t
		}
	}

	if v, ok := d.GetOk("neg_resp_time"); ok || d.HasChange("neg_resp_time") {
		t, err := expandWanoptWebcacheNegRespTime(d, v, "neg_resp_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neg-resp-time"] = t
		}
	}

	if v, ok := d.GetOk("reval_pnc"); ok || d.HasChange("reval_pnc") {
		t, err := expandWanoptWebcacheRevalPnc(d, v, "reval_pnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reval-pnc"] = t
		}
	}

	return &obj, nil
}
