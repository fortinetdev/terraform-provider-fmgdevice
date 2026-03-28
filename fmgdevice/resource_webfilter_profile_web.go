// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Web content filtering settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfileWeb() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileWebUpdate,
		Read:   resourceWebfilterProfileWebRead,
		Update: resourceWebfilterProfileWebUpdate,
		Delete: resourceWebfilterProfileWebDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"blacklist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowlist": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"blocklist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bword_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"content_header_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"keyword_match": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"log_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"safe_search": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"urlfilter_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vimeo_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"youtube_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"qwant_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterProfileWebUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectWebfilterProfileWeb(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileWeb resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfileWeb(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfileWeb resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WebfilterProfileWeb")

	return resourceWebfilterProfileWebRead(d, m)
}

func resourceWebfilterProfileWebDelete(d *schema.ResourceData, m interface{}) error {
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
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterProfileWeb(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfileWeb resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileWebRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
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
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadWebfilterProfileWeb(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfileWeb resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfileWeb(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfileWeb resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileWebBlacklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebAllowlist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebBwordTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebBwordThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentHeaderList2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebKeywordMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebLogSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebSafeSearch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebUrlfilterTable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebWhitelist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebVimeoRestrict2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebYoutubeRestrict2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebQwantRestrict2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfileWeb(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("blacklist", flattenWebfilterProfileWebBlacklist2edl(o["blacklist"], d, "blacklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["blacklist"], "WebfilterProfileWeb-Blacklist"); ok {
			if err = d.Set("blacklist", vv); err != nil {
				return fmt.Errorf("Error reading blacklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blacklist: %v", err)
		}
	}

	if err = d.Set("allowlist", flattenWebfilterProfileWebAllowlist2edl(o["allowlist"], d, "allowlist")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowlist"], "WebfilterProfileWeb-Allowlist"); ok {
			if err = d.Set("allowlist", vv); err != nil {
				return fmt.Errorf("Error reading allowlist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowlist: %v", err)
		}
	}

	if err = d.Set("blocklist", flattenWebfilterProfileWebBlocklist2edl(o["blocklist"], d, "blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocklist"], "WebfilterProfileWeb-Blocklist"); ok {
			if err = d.Set("blocklist", vv); err != nil {
				return fmt.Errorf("Error reading blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocklist: %v", err)
		}
	}

	if err = d.Set("bword_table", flattenWebfilterProfileWebBwordTable2edl(o["bword-table"], d, "bword_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["bword-table"], "WebfilterProfileWeb-BwordTable"); ok {
			if err = d.Set("bword_table", vv); err != nil {
				return fmt.Errorf("Error reading bword_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bword_table: %v", err)
		}
	}

	if err = d.Set("bword_threshold", flattenWebfilterProfileWebBwordThreshold2edl(o["bword-threshold"], d, "bword_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["bword-threshold"], "WebfilterProfileWeb-BwordThreshold"); ok {
			if err = d.Set("bword_threshold", vv); err != nil {
				return fmt.Errorf("Error reading bword_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bword_threshold: %v", err)
		}
	}

	if err = d.Set("content_header_list", flattenWebfilterProfileWebContentHeaderList2edl(o["content-header-list"], d, "content_header_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-header-list"], "WebfilterProfileWeb-ContentHeaderList"); ok {
			if err = d.Set("content_header_list", vv); err != nil {
				return fmt.Errorf("Error reading content_header_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_header_list: %v", err)
		}
	}

	if err = d.Set("keyword_match", flattenWebfilterProfileWebKeywordMatch2edl(o["keyword-match"], d, "keyword_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyword-match"], "WebfilterProfileWeb-KeywordMatch"); ok {
			if err = d.Set("keyword_match", vv); err != nil {
				return fmt.Errorf("Error reading keyword_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyword_match: %v", err)
		}
	}

	if err = d.Set("log_search", flattenWebfilterProfileWebLogSearch2edl(o["log-search"], d, "log_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-search"], "WebfilterProfileWeb-LogSearch"); ok {
			if err = d.Set("log_search", vv); err != nil {
				return fmt.Errorf("Error reading log_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_search: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenWebfilterProfileWebSafeSearch2edl(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "WebfilterProfileWeb-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("urlfilter_table", flattenWebfilterProfileWebUrlfilterTable2edl(o["urlfilter-table"], d, "urlfilter_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["urlfilter-table"], "WebfilterProfileWeb-UrlfilterTable"); ok {
			if err = d.Set("urlfilter_table", vv); err != nil {
				return fmt.Errorf("Error reading urlfilter_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading urlfilter_table: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenWebfilterProfileWebWhitelist2edl(o["whitelist"], d, "whitelist")); err != nil {
		if vv, ok := fortiAPIPatch(o["whitelist"], "WebfilterProfileWeb-Whitelist"); ok {
			if err = d.Set("whitelist", vv); err != nil {
				return fmt.Errorf("Error reading whitelist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if err = d.Set("vimeo_restrict", flattenWebfilterProfileWebVimeoRestrict2edl(o["vimeo-restrict"], d, "vimeo_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo-restrict"], "WebfilterProfileWeb-VimeoRestrict"); ok {
			if err = d.Set("vimeo_restrict", vv); err != nil {
				return fmt.Errorf("Error reading vimeo_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo_restrict: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenWebfilterProfileWebYoutubeRestrict2edl(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "WebfilterProfileWeb-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	if err = d.Set("qwant_restrict", flattenWebfilterProfileWebQwantRestrict2edl(o["qwant-restrict"], d, "qwant_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["qwant-restrict"], "WebfilterProfileWeb-QwantRestrict"); ok {
			if err = d.Set("qwant_restrict", vv); err != nil {
				return fmt.Errorf("Error reading qwant_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qwant_restrict: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileWebFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileWebBlacklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebAllowlist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBwordTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebBwordThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentHeaderList2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebKeywordMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebLogSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebSafeSearch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebUrlfilterTable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebWhitelist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebVimeoRestrict2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebYoutubeRestrict2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebQwantRestrict2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfileWeb(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("blacklist"); ok || d.HasChange("blacklist") {
		t, err := expandWebfilterProfileWebBlacklist2edl(d, v, "blacklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blacklist"] = t
		}
	}

	if v, ok := d.GetOk("allowlist"); ok || d.HasChange("allowlist") {
		t, err := expandWebfilterProfileWebAllowlist2edl(d, v, "allowlist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowlist"] = t
		}
	}

	if v, ok := d.GetOk("blocklist"); ok || d.HasChange("blocklist") {
		t, err := expandWebfilterProfileWebBlocklist2edl(d, v, "blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocklist"] = t
		}
	}

	if v, ok := d.GetOk("bword_table"); ok || d.HasChange("bword_table") {
		t, err := expandWebfilterProfileWebBwordTable2edl(d, v, "bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bword-table"] = t
		}
	}

	if v, ok := d.GetOk("bword_threshold"); ok || d.HasChange("bword_threshold") {
		t, err := expandWebfilterProfileWebBwordThreshold2edl(d, v, "bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bword-threshold"] = t
		}
	}

	if v, ok := d.GetOk("content_header_list"); ok || d.HasChange("content_header_list") {
		t, err := expandWebfilterProfileWebContentHeaderList2edl(d, v, "content_header_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-header-list"] = t
		}
	}

	if v, ok := d.GetOk("keyword_match"); ok || d.HasChange("keyword_match") {
		t, err := expandWebfilterProfileWebKeywordMatch2edl(d, v, "keyword_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyword-match"] = t
		}
	}

	if v, ok := d.GetOk("log_search"); ok || d.HasChange("log_search") {
		t, err := expandWebfilterProfileWebLogSearch2edl(d, v, "log_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-search"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok || d.HasChange("safe_search") {
		t, err := expandWebfilterProfileWebSafeSearch2edl(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("urlfilter_table"); ok || d.HasChange("urlfilter_table") {
		t, err := expandWebfilterProfileWebUrlfilterTable2edl(d, v, "urlfilter_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["urlfilter-table"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok || d.HasChange("whitelist") {
		t, err := expandWebfilterProfileWebWhitelist2edl(d, v, "whitelist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("vimeo_restrict"); ok || d.HasChange("vimeo_restrict") {
		t, err := expandWebfilterProfileWebVimeoRestrict2edl(d, v, "vimeo_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo-restrict"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok || d.HasChange("youtube_restrict") {
		t, err := expandWebfilterProfileWebYoutubeRestrict2edl(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	if v, ok := d.GetOk("qwant_restrict"); ok || d.HasChange("qwant_restrict") {
		t, err := expandWebfilterProfileWebQwantRestrict2edl(d, v, "qwant_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qwant-restrict"] = t
		}
	}

	return &obj, nil
}
