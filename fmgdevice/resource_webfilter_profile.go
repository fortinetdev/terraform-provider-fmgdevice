// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Web filter profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileCreate,
		Read:   resourceWebfilterProfileRead,
		Update: resourceWebfilterProfileUpdate,
		Delete: resourceWebfilterProfileDelete,

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
			"antiphish": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"check_basic_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"check_uri": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"check_username_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain_controller": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"inspection_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fortiguard_category": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"ldap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"max_body_len": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftgd_wf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exempt_quota": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_usr_grp": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"category": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"warn_duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"warning_duration_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"warning_prompt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"max_quota_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ovrd": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"quota": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"unit": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"rate_crl_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_css_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_image_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate_javascript_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"risk": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"risk_level": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"https_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_all_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ovrd_cookie": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_dur": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_dur_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_scope": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_user_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"profile_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ovrd_perm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"post_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_extraction": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redirect_no_content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redirect_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"web": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"web_antiphishing_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_content_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_extended_all_action_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_activex_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_applet_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_command_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_removal_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_js_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_jscript_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_referer_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_unknown_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_filter_vbs_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_flow_log_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_quota_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_invalid_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_url_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"youtube_channel_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ia_categorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWebfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterProfileRead(d, m)
}

func resourceWebfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebfilterProfileRead(d, m)
}

func resourceWebfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWebfilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileAntiphish(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "authentication"
	if _, ok := i["authentication"]; ok {
		result["authentication"] = flattenWebfilterProfileAntiphishAuthentication(i["authentication"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := i["check-basic-auth"]; ok {
		result["check_basic_auth"] = flattenWebfilterProfileAntiphishCheckBasicAuth(i["check-basic-auth"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_uri"
	if _, ok := i["check-uri"]; ok {
		result["check_uri"] = flattenWebfilterProfileAntiphishCheckUri(i["check-uri"], d, pre_append)
	}

	pre_append = pre + ".0." + "check_username_only"
	if _, ok := i["check-username-only"]; ok {
		result["check_username_only"] = flattenWebfilterProfileAntiphishCheckUsernameOnly(i["check-username-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := i["custom-patterns"]; ok {
		result["custom_patterns"] = flattenWebfilterProfileAntiphishCustomPatterns(i["custom-patterns"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_action"
	if _, ok := i["default-action"]; ok {
		result["default_action"] = flattenWebfilterProfileAntiphishDefaultAction(i["default-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = flattenWebfilterProfileAntiphishDomainController(i["domain-controller"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := i["inspection-entries"]; ok {
		result["inspection_entries"] = flattenWebfilterProfileAntiphishInspectionEntries(i["inspection-entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "ldap"
	if _, ok := i["ldap"]; ok {
		result["ldap"] = flattenWebfilterProfileAntiphishLdap(i["ldap"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_body_len"
	if _, ok := i["max-body-len"]; ok {
		result["max_body_len"] = flattenWebfilterProfileAntiphishMaxBodyLen(i["max-body-len"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWebfilterProfileAntiphishStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileAntiphishAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckBasicAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckUsernameOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatterns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenWebfilterProfileAntiphishCustomPatternsCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-CustomPatterns-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenWebfilterProfileAntiphishCustomPatternsPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-CustomPatterns-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWebfilterProfileAntiphishCustomPatternsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-CustomPatterns-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileAntiphishCustomPatternsCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatternsPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatternsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileAntiphishInspectionEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWebfilterProfileAntiphishInspectionEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-InspectionEntries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			v := flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(i["fortiguard-category"], d, pre_append)
			tmp["fortiguard_category"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-InspectionEntries-FortiguardCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWebfilterProfileAntiphishInspectionEntriesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WebfilterProfileAntiphish-InspectionEntries-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileAntiphishInspectionEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileAntiphishInspectionEntriesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishLdap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileAntiphishMaxBodyLen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := i["exempt-quota"]; ok {
		result["exempt_quota"] = flattenWebfilterProfileFtgdWfExemptQuota(i["exempt-quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenWebfilterProfileFtgdWfFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := i["max-quota-timeout"]; ok {
		result["max_quota_timeout"] = flattenWebfilterProfileFtgdWfMaxQuotaTimeout(i["max-quota-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenWebfilterProfileFtgdWfOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd"
	if _, ok := i["ovrd"]; ok {
		result["ovrd"] = flattenWebfilterProfileFtgdWfOvrd(i["ovrd"], d, pre_append)
	}

	pre_append = pre + ".0." + "quota"
	if _, ok := i["quota"]; ok {
		result["quota"] = flattenWebfilterProfileFtgdWfQuota(i["quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := i["rate-crl-urls"]; ok {
		result["rate_crl_urls"] = flattenWebfilterProfileFtgdWfRateCrlUrls(i["rate-crl-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := i["rate-css-urls"]; ok {
		result["rate_css_urls"] = flattenWebfilterProfileFtgdWfRateCssUrls(i["rate-css-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := i["rate-image-urls"]; ok {
		result["rate_image_urls"] = flattenWebfilterProfileFtgdWfRateImageUrls(i["rate-image-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := i["rate-javascript-urls"]; ok {
		result["rate_javascript_urls"] = flattenWebfilterProfileFtgdWfRateJavascriptUrls(i["rate-javascript-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "risk"
	if _, ok := i["risk"]; ok {
		result["risk"] = flattenWebfilterProfileFtgdWfRisk(i["risk"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileFtgdWfExemptQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := i["auth-usr-grp"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(i["auth-usr-grp"], d, pre_append)
			tmp["auth_usr_grp"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-AuthUsrGrp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := i["warn-duration"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarnDuration(i["warn-duration"], d, pre_append)
			tmp["warn_duration"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarnDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := i["warning-duration-type"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarningDurationType(i["warning-duration-type"], d, pre_append)
			tmp["warning_duration_type"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarningDurationType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := i["warning-prompt"]; ok {
			v := flattenWebfilterProfileFtgdWfFiltersWarningPrompt(i["warning-prompt"], d, pre_append)
			tmp["warning_prompt"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Filters-WarningPrompt")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarnDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningDurationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfMaxQuotaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfOvrd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := i["duration"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaDuration(i["duration"], d, pre_append)
			tmp["duration"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Duration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
			tmp["override_replacemsg"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-OverrideReplacemsg")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := i["unit"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaUnit(i["unit"], d, pre_append)
			tmp["unit"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Unit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenWebfilterProfileFtgdWfQuotaValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Quota-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfQuotaCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileFtgdWfQuotaDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCrlUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCssUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateImageUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateJavascriptUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRisk(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk_level"
		if _, ok := i["risk-level"]; ok {
			v := flattenWebfilterProfileFtgdWfRiskRiskLevel(i["risk-level"], d, pre_append)
			tmp["risk_level"] = fortiAPISubPartPatch(v, "WebfilterProfileFtgdWf-Risk-RiskLevel")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfRiskAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRiskRiskLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileHttpsReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileLogAllUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := i["ovrd-cookie"]; ok {
		result["ovrd_cookie"] = flattenWebfilterProfileOverrideOvrdCookie(i["ovrd-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := i["ovrd-dur"]; ok {
		result["ovrd_dur"] = flattenWebfilterProfileOverrideOvrdDur(i["ovrd-dur"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := i["ovrd-dur-mode"]; ok {
		result["ovrd_dur_mode"] = flattenWebfilterProfileOverrideOvrdDurMode(i["ovrd-dur-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := i["ovrd-scope"]; ok {
		result["ovrd_scope"] = flattenWebfilterProfileOverrideOvrdScope(i["ovrd-scope"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := i["ovrd-user-group"]; ok {
		result["ovrd_user_group"] = flattenWebfilterProfileOverrideOvrdUserGroup(i["ovrd-user-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenWebfilterProfileOverrideProfile(i["profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := i["profile-attribute"]; ok {
		result["profile_attribute"] = flattenWebfilterProfileOverrideProfileAttribute(i["profile-attribute"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_type"
	if _, ok := i["profile-type"]; ok {
		result["profile_type"] = flattenWebfilterProfileOverrideProfileType(i["profile-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileOverrideOvrdCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDur(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDurMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileOverrideProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileOverrideProfileAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOvrdPerm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfilePostAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileUrlExtraction(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redirect_header"
	if _, ok := i["redirect-header"]; ok {
		result["redirect_header"] = flattenWebfilterProfileUrlExtractionRedirectHeader(i["redirect-header"], d, pre_append)
	}

	pre_append = pre + ".0." + "redirect_no_content"
	if _, ok := i["redirect-no-content"]; ok {
		result["redirect_no_content"] = flattenWebfilterProfileUrlExtractionRedirectNoContent(i["redirect-no-content"], d, pre_append)
	}

	pre_append = pre + ".0." + "redirect_url"
	if _, ok := i["redirect-url"]; ok {
		result["redirect_url"] = flattenWebfilterProfileUrlExtractionRedirectUrl(i["redirect-url"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_fqdn"
	if _, ok := i["server-fqdn"]; ok {
		result["server_fqdn"] = flattenWebfilterProfileUrlExtractionServerFqdn(i["server-fqdn"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWebfilterProfileUrlExtractionStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileUrlExtractionRedirectHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileUrlExtractionRedirectNoContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileUrlExtractionRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileUrlExtractionServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileUrlExtractionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWeb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blacklist"
	if _, ok := i["blacklist"]; ok {
		result["blacklist"] = flattenWebfilterProfileWebBlacklist(i["blacklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "allowlist"
	if _, ok := i["allowlist"]; ok {
		result["allowlist"] = flattenWebfilterProfileWebAllowlist(i["allowlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocklist"
	if _, ok := i["blocklist"]; ok {
		result["blocklist"] = flattenWebfilterProfileWebBlocklist(i["blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_table"
	if _, ok := i["bword-table"]; ok {
		result["bword_table"] = flattenWebfilterProfileWebBwordTable(i["bword-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := i["bword-threshold"]; ok {
		result["bword_threshold"] = flattenWebfilterProfileWebBwordThreshold(i["bword-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_header_list"
	if _, ok := i["content-header-list"]; ok {
		result["content_header_list"] = flattenWebfilterProfileWebContentHeaderList(i["content-header-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "keyword_match"
	if _, ok := i["keyword-match"]; ok {
		result["keyword_match"] = flattenWebfilterProfileWebKeywordMatch(i["keyword-match"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_search"
	if _, ok := i["log-search"]; ok {
		result["log_search"] = flattenWebfilterProfileWebLogSearch(i["log-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "safe_search"
	if _, ok := i["safe-search"]; ok {
		result["safe_search"] = flattenWebfilterProfileWebSafeSearch(i["safe-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := i["urlfilter-table"]; ok {
		result["urlfilter_table"] = flattenWebfilterProfileWebUrlfilterTable(i["urlfilter-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "whitelist"
	if _, ok := i["whitelist"]; ok {
		result["whitelist"] = flattenWebfilterProfileWebWhitelist(i["whitelist"], d, pre_append)
	}

	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := i["vimeo-restrict"]; ok {
		result["vimeo_restrict"] = flattenWebfilterProfileWebVimeoRestrict(i["vimeo-restrict"], d, pre_append)
	}

	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := i["youtube-restrict"]; ok {
		result["youtube_restrict"] = flattenWebfilterProfileWebYoutubeRestrict(i["youtube-restrict"], d, pre_append)
	}

	pre_append = pre + ".0." + "qwant_restrict"
	if _, ok := i["qwant-restrict"]; ok {
		result["qwant_restrict"] = flattenWebfilterProfileWebQwantRestrict(i["qwant-restrict"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileWebBlacklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebAllowlist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentHeaderList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebKeywordMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebLogSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebUrlfilterTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileWebVimeoRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebQwantRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebAntiphishingLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebExtendedAllActionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterActivexLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterAppletLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCommandBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieRemovalLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJscriptLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterRefererLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterUnknownLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterVbsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFlowLogEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdErrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdQuotaUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebInvalidDomainLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebUrlLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWisp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWispAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWispServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebfilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := i["channel-id"]; ok {
			v := flattenWebfilterProfileYoutubeChannelFilterChannelId(i["channel-id"], d, pre_append)
			tmp["channel_id"] = fortiAPISubPartPatch(v, "WebfilterProfile-YoutubeChannelFilter-ChannelId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWebfilterProfileYoutubeChannelFilterComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WebfilterProfile-YoutubeChannelFilter-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWebfilterProfileYoutubeChannelFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WebfilterProfile-YoutubeChannelFilter-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWebfilterProfileYoutubeChannelFilterChannelId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileIaCategorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("antiphish", flattenWebfilterProfileAntiphish(o["antiphish"], d, "antiphish")); err != nil {
			if vv, ok := fortiAPIPatch(o["antiphish"], "WebfilterProfile-Antiphish"); ok {
				if err = d.Set("antiphish", vv); err != nil {
					return fmt.Errorf("Error reading antiphish: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("antiphish"); ok {
			if err = d.Set("antiphish", flattenWebfilterProfileAntiphish(o["antiphish"], d, "antiphish")); err != nil {
				if vv, ok := fortiAPIPatch(o["antiphish"], "WebfilterProfile-Antiphish"); ok {
					if err = d.Set("antiphish", vv); err != nil {
						return fmt.Errorf("Error reading antiphish: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading antiphish: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenWebfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "WebfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenWebfilterProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "WebfilterProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenWebfilterProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "WebfilterProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftgd-wf"], "WebfilterProfile-FtgdWf"); ok {
				if err = d.Set("ftgd_wf", vv); err != nil {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftgd_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_wf"); ok {
			if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftgd-wf"], "WebfilterProfile-FtgdWf"); ok {
					if err = d.Set("ftgd_wf", vv); err != nil {
						return fmt.Errorf("Error reading ftgd_wf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			}
		}
	}

	if err = d.Set("https_replacemsg", flattenWebfilterProfileHttpsReplacemsg(o["https-replacemsg"], d, "https_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-replacemsg"], "WebfilterProfile-HttpsReplacemsg"); ok {
			if err = d.Set("https_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading https_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_replacemsg: %v", err)
		}
	}

	if err = d.Set("log_all_url", flattenWebfilterProfileLogAllUrl(o["log-all-url"], d, "log_all_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all-url"], "WebfilterProfile-LogAllUrl"); ok {
			if err = d.Set("log_all_url", vv); err != nil {
				return fmt.Errorf("Error reading log_all_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all_url: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenWebfilterProfileOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "WebfilterProfile-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override")); err != nil {
			if vv, ok := fortiAPIPatch(o["override"], "WebfilterProfile-Override"); ok {
				if err = d.Set("override", vv); err != nil {
					return fmt.Errorf("Error reading override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override")); err != nil {
				if vv, ok := fortiAPIPatch(o["override"], "WebfilterProfile-Override"); ok {
					if err = d.Set("override", vv); err != nil {
						return fmt.Errorf("Error reading override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	if err = d.Set("ovrd_perm", flattenWebfilterProfileOvrdPerm(o["ovrd-perm"], d, "ovrd_perm")); err != nil {
		if vv, ok := fortiAPIPatch(o["ovrd-perm"], "WebfilterProfile-OvrdPerm"); ok {
			if err = d.Set("ovrd_perm", vv); err != nil {
				return fmt.Errorf("Error reading ovrd_perm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ovrd_perm: %v", err)
		}
	}

	if err = d.Set("post_action", flattenWebfilterProfilePostAction(o["post-action"], d, "post_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["post-action"], "WebfilterProfile-PostAction"); ok {
			if err = d.Set("post_action", vv); err != nil {
				return fmt.Errorf("Error reading post_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading post_action: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenWebfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "WebfilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("url_extraction", flattenWebfilterProfileUrlExtraction(o["url-extraction"], d, "url_extraction")); err != nil {
			if vv, ok := fortiAPIPatch(o["url-extraction"], "WebfilterProfile-UrlExtraction"); ok {
				if err = d.Set("url_extraction", vv); err != nil {
					return fmt.Errorf("Error reading url_extraction: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading url_extraction: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_extraction"); ok {
			if err = d.Set("url_extraction", flattenWebfilterProfileUrlExtraction(o["url-extraction"], d, "url_extraction")); err != nil {
				if vv, ok := fortiAPIPatch(o["url-extraction"], "WebfilterProfile-UrlExtraction"); ok {
					if err = d.Set("url_extraction", vv); err != nil {
						return fmt.Errorf("Error reading url_extraction: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading url_extraction: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web")); err != nil {
			if vv, ok := fortiAPIPatch(o["web"], "WebfilterProfile-Web"); ok {
				if err = d.Set("web", vv); err != nil {
					return fmt.Errorf("Error reading web: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web")); err != nil {
				if vv, ok := fortiAPIPatch(o["web"], "WebfilterProfile-Web"); ok {
					if err = d.Set("web", vv); err != nil {
						return fmt.Errorf("Error reading web: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	if err = d.Set("web_antiphishing_log", flattenWebfilterProfileWebAntiphishingLog(o["web-antiphishing-log"], d, "web_antiphishing_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-antiphishing-log"], "WebfilterProfile-WebAntiphishingLog"); ok {
			if err = d.Set("web_antiphishing_log", vv); err != nil {
				return fmt.Errorf("Error reading web_antiphishing_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_antiphishing_log: %v", err)
		}
	}

	if err = d.Set("web_content_log", flattenWebfilterProfileWebContentLog(o["web-content-log"], d, "web_content_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-content-log"], "WebfilterProfile-WebContentLog"); ok {
			if err = d.Set("web_content_log", vv); err != nil {
				return fmt.Errorf("Error reading web_content_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_content_log: %v", err)
		}
	}

	if err = d.Set("web_extended_all_action_log", flattenWebfilterProfileWebExtendedAllActionLog(o["web-extended-all-action-log"], d, "web_extended_all_action_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-extended-all-action-log"], "WebfilterProfile-WebExtendedAllActionLog"); ok {
			if err = d.Set("web_extended_all_action_log", vv); err != nil {
				return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
		}
	}

	if err = d.Set("web_filter_activex_log", flattenWebfilterProfileWebFilterActivexLog(o["web-filter-activex-log"], d, "web_filter_activex_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-activex-log"], "WebfilterProfile-WebFilterActivexLog"); ok {
			if err = d.Set("web_filter_activex_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
		}
	}

	if err = d.Set("web_filter_applet_log", flattenWebfilterProfileWebFilterAppletLog(o["web-filter-applet-log"], d, "web_filter_applet_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-applet-log"], "WebfilterProfile-WebFilterAppletLog"); ok {
			if err = d.Set("web_filter_applet_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
		}
	}

	if err = d.Set("web_filter_command_block_log", flattenWebfilterProfileWebFilterCommandBlockLog(o["web-filter-command-block-log"], d, "web_filter_command_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-command-block-log"], "WebfilterProfile-WebFilterCommandBlockLog"); ok {
			if err = d.Set("web_filter_command_block_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_log", flattenWebfilterProfileWebFilterCookieLog(o["web-filter-cookie-log"], d, "web_filter_cookie_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-cookie-log"], "WebfilterProfile-WebFilterCookieLog"); ok {
			if err = d.Set("web_filter_cookie_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_removal_log", flattenWebfilterProfileWebFilterCookieRemovalLog(o["web-filter-cookie-removal-log"], d, "web_filter_cookie_removal_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-cookie-removal-log"], "WebfilterProfile-WebFilterCookieRemovalLog"); ok {
			if err = d.Set("web_filter_cookie_removal_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
		}
	}

	if err = d.Set("web_filter_js_log", flattenWebfilterProfileWebFilterJsLog(o["web-filter-js-log"], d, "web_filter_js_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-js-log"], "WebfilterProfile-WebFilterJsLog"); ok {
			if err = d.Set("web_filter_js_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_js_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_js_log: %v", err)
		}
	}

	if err = d.Set("web_filter_jscript_log", flattenWebfilterProfileWebFilterJscriptLog(o["web-filter-jscript-log"], d, "web_filter_jscript_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-jscript-log"], "WebfilterProfile-WebFilterJscriptLog"); ok {
			if err = d.Set("web_filter_jscript_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
		}
	}

	if err = d.Set("web_filter_referer_log", flattenWebfilterProfileWebFilterRefererLog(o["web-filter-referer-log"], d, "web_filter_referer_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-referer-log"], "WebfilterProfile-WebFilterRefererLog"); ok {
			if err = d.Set("web_filter_referer_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
		}
	}

	if err = d.Set("web_filter_unknown_log", flattenWebfilterProfileWebFilterUnknownLog(o["web-filter-unknown-log"], d, "web_filter_unknown_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-unknown-log"], "WebfilterProfile-WebFilterUnknownLog"); ok {
			if err = d.Set("web_filter_unknown_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
		}
	}

	if err = d.Set("web_filter_vbs_log", flattenWebfilterProfileWebFilterVbsLog(o["web-filter-vbs-log"], d, "web_filter_vbs_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-filter-vbs-log"], "WebfilterProfile-WebFilterVbsLog"); ok {
			if err = d.Set("web_filter_vbs_log", vv); err != nil {
				return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
		}
	}

	if err = d.Set("web_flow_log_encoding", flattenWebfilterProfileWebFlowLogEncoding(o["web-flow-log-encoding"], d, "web_flow_log_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-flow-log-encoding"], "WebfilterProfile-WebFlowLogEncoding"); ok {
			if err = d.Set("web_flow_log_encoding", vv); err != nil {
				return fmt.Errorf("Error reading web_flow_log_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_flow_log_encoding: %v", err)
		}
	}

	if err = d.Set("web_ftgd_err_log", flattenWebfilterProfileWebFtgdErrLog(o["web-ftgd-err-log"], d, "web_ftgd_err_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-ftgd-err-log"], "WebfilterProfile-WebFtgdErrLog"); ok {
			if err = d.Set("web_ftgd_err_log", vv); err != nil {
				return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_quota_usage", flattenWebfilterProfileWebFtgdQuotaUsage(o["web-ftgd-quota-usage"], d, "web_ftgd_quota_usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-ftgd-quota-usage"], "WebfilterProfile-WebFtgdQuotaUsage"); ok {
			if err = d.Set("web_ftgd_quota_usage", vv); err != nil {
				return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
		}
	}

	if err = d.Set("web_invalid_domain_log", flattenWebfilterProfileWebInvalidDomainLog(o["web-invalid-domain-log"], d, "web_invalid_domain_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-invalid-domain-log"], "WebfilterProfile-WebInvalidDomainLog"); ok {
			if err = d.Set("web_invalid_domain_log", vv); err != nil {
				return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
		}
	}

	if err = d.Set("web_url_log", flattenWebfilterProfileWebUrlLog(o["web-url-log"], d, "web_url_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-url-log"], "WebfilterProfile-WebUrlLog"); ok {
			if err = d.Set("web_url_log", vv); err != nil {
				return fmt.Errorf("Error reading web_url_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_url_log: %v", err)
		}
	}

	if err = d.Set("wisp", flattenWebfilterProfileWisp(o["wisp"], d, "wisp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp"], "WebfilterProfile-Wisp"); ok {
			if err = d.Set("wisp", vv); err != nil {
				return fmt.Errorf("Error reading wisp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp: %v", err)
		}
	}

	if err = d.Set("wisp_algorithm", flattenWebfilterProfileWispAlgorithm(o["wisp-algorithm"], d, "wisp_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp-algorithm"], "WebfilterProfile-WispAlgorithm"); ok {
			if err = d.Set("wisp_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading wisp_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp_algorithm: %v", err)
		}
	}

	if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["wisp-servers"], "WebfilterProfile-WispServers"); ok {
			if err = d.Set("wisp_servers", vv); err != nil {
				return fmt.Errorf("Error reading wisp_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wisp_servers: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "WebfilterProfile-YoutubeChannelFilter"); ok {
				if err = d.Set("youtube_channel_filter", vv); err != nil {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("youtube_channel_filter"); ok {
			if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "WebfilterProfile-YoutubeChannelFilter"); ok {
					if err = d.Set("youtube_channel_filter", vv); err != nil {
						return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("youtube_channel_status", flattenWebfilterProfileYoutubeChannelStatus(o["youtube-channel-status"], d, "youtube_channel_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-channel-status"], "WebfilterProfile-YoutubeChannelStatus"); ok {
			if err = d.Set("youtube_channel_status", vv); err != nil {
				return fmt.Errorf("Error reading youtube_channel_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_channel_status: %v", err)
		}
	}

	if err = d.Set("ia_categorization", flattenWebfilterProfileIaCategorization(o["ia-categorization"], d, "ia_categorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["ia-categorization"], "WebfilterProfile-IaCategorization"); ok {
			if err = d.Set("ia_categorization", vv); err != nil {
				return fmt.Errorf("Error reading ia_categorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ia_categorization: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileAntiphish(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "authentication"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["authentication"], _ = expandWebfilterProfileAntiphishAuthentication(d, i["authentication"], pre_append)
	}
	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-basic-auth"], _ = expandWebfilterProfileAntiphishCheckBasicAuth(d, i["check_basic_auth"], pre_append)
	}
	pre_append = pre + ".0." + "check_uri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-uri"], _ = expandWebfilterProfileAntiphishCheckUri(d, i["check_uri"], pre_append)
	}
	pre_append = pre + ".0." + "check_username_only"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-username-only"], _ = expandWebfilterProfileAntiphishCheckUsernameOnly(d, i["check_username_only"], pre_append)
	}
	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWebfilterProfileAntiphishCustomPatterns(d, i["custom_patterns"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["custom-patterns"] = t
		}
	}
	pre_append = pre + ".0." + "default_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-action"], _ = expandWebfilterProfileAntiphishDefaultAction(d, i["default_action"], pre_append)
	}
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["domain-controller"], _ = expandWebfilterProfileAntiphishDomainController(d, i["domain_controller"], pre_append)
	}
	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWebfilterProfileAntiphishInspectionEntries(d, i["inspection_entries"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["inspection-entries"] = t
		}
	}
	pre_append = pre + ".0." + "ldap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ldap"], _ = expandWebfilterProfileAntiphishLdap(d, i["ldap"], pre_append)
	}
	pre_append = pre + ".0." + "max_body_len"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-body-len"], _ = expandWebfilterProfileAntiphishMaxBodyLen(d, i["max_body_len"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWebfilterProfileAntiphishStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWebfilterProfileAntiphishAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckBasicAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckUsernameOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatterns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandWebfilterProfileAntiphishCustomPatternsCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWebfilterProfileAntiphishCustomPatternsPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWebfilterProfileAntiphishCustomPatternsType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileAntiphishCustomPatternsCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatternsPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatternsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWebfilterProfileAntiphishInspectionEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard-category"], _ = expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d, i["fortiguard_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWebfilterProfileAntiphishInspectionEntriesName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileAntiphishInspectionEntriesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishLdap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileAntiphishMaxBodyLen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["exempt-quota"], _ = expandWebfilterProfileFtgdWfExemptQuota(d, i["exempt_quota"], pre_append)
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWebfilterProfileFtgdWfFilters(d, i["filters"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["filters"] = t
		}
	}
	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-quota-timeout"], _ = expandWebfilterProfileFtgdWfMaxQuotaTimeout(d, i["max_quota_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandWebfilterProfileFtgdWfOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd"], _ = expandWebfilterProfileFtgdWfOvrd(d, i["ovrd"], pre_append)
	}
	pre_append = pre + ".0." + "quota"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWebfilterProfileFtgdWfQuota(d, i["quota"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["quota"] = t
		}
	}
	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-crl-urls"], _ = expandWebfilterProfileFtgdWfRateCrlUrls(d, i["rate_crl_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-css-urls"], _ = expandWebfilterProfileFtgdWfRateCssUrls(d, i["rate_css_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-image-urls"], _ = expandWebfilterProfileFtgdWfRateImageUrls(d, i["rate_image_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-javascript-urls"], _ = expandWebfilterProfileFtgdWfRateJavascriptUrls(d, i["rate_javascript_urls"], pre_append)
	}
	pre_append = pre + ".0." + "risk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandWebfilterProfileFtgdWfRisk(d, i["risk"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["risk"] = t
		}
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfExemptQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWebfilterProfileFtgdWfFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-usr-grp"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d, i["auth_usr_grp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandWebfilterProfileFtgdWfFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWebfilterProfileFtgdWfFiltersLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warn-duration"], _ = expandWebfilterProfileFtgdWfFiltersWarnDuration(d, i["warn_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-duration-type"], _ = expandWebfilterProfileFtgdWfFiltersWarningDurationType(d, i["warning_duration_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["warning-prompt"], _ = expandWebfilterProfileFtgdWfFiltersWarningPrompt(d, i["warning_prompt"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarnDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningDurationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningPrompt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfMaxQuotaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfOvrd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandWebfilterProfileFtgdWfQuotaCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["duration"], _ = expandWebfilterProfileFtgdWfQuotaDuration(d, i["duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfQuotaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandWebfilterProfileFtgdWfQuotaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["unit"], _ = expandWebfilterProfileFtgdWfQuotaUnit(d, i["unit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandWebfilterProfileFtgdWfQuotaValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfQuotaCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileFtgdWfQuotaDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCrlUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCssUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateImageUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateJavascriptUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandWebfilterProfileFtgdWfRiskAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileFtgdWfRiskId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandWebfilterProfileFtgdWfRiskLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk_level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["risk-level"], _ = expandWebfilterProfileFtgdWfRiskRiskLevel(d, i["risk_level"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfRiskAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRiskRiskLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileHttpsReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileLogAllUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-cookie"], _ = expandWebfilterProfileOverrideOvrdCookie(d, i["ovrd_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-dur"], _ = expandWebfilterProfileOverrideOvrdDur(d, i["ovrd_dur"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-dur-mode"], _ = expandWebfilterProfileOverrideOvrdDurMode(d, i["ovrd_dur_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-scope"], _ = expandWebfilterProfileOverrideOvrdScope(d, i["ovrd_scope"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ovrd-user-group"], _ = expandWebfilterProfileOverrideOvrdUserGroup(d, i["ovrd_user_group"], pre_append)
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile"], _ = expandWebfilterProfileOverrideProfile(d, i["profile"], pre_append)
	}
	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile-attribute"], _ = expandWebfilterProfileOverrideProfileAttribute(d, i["profile_attribute"], pre_append)
	}
	pre_append = pre + ".0." + "profile_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["profile-type"], _ = expandWebfilterProfileOverrideProfileType(d, i["profile_type"], pre_append)
	}

	return result, nil
}

func expandWebfilterProfileOverrideOvrdCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDur(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDurMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileOverrideProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileOverrideProfileAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOvrdPerm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfilePostAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileUrlExtraction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redirect_header"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-header"], _ = expandWebfilterProfileUrlExtractionRedirectHeader(d, i["redirect_header"], pre_append)
	}
	pre_append = pre + ".0." + "redirect_no_content"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-no-content"], _ = expandWebfilterProfileUrlExtractionRedirectNoContent(d, i["redirect_no_content"], pre_append)
	}
	pre_append = pre + ".0." + "redirect_url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redirect-url"], _ = expandWebfilterProfileUrlExtractionRedirectUrl(d, i["redirect_url"], pre_append)
	}
	pre_append = pre + ".0." + "server_fqdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-fqdn"], _ = expandWebfilterProfileUrlExtractionServerFqdn(d, i["server_fqdn"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandWebfilterProfileUrlExtractionStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandWebfilterProfileUrlExtractionRedirectHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileUrlExtractionRedirectNoContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileUrlExtractionRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileUrlExtractionServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileUrlExtractionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "blacklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blacklist"], _ = expandWebfilterProfileWebBlacklist(d, i["blacklist"], pre_append)
	}
	pre_append = pre + ".0." + "allowlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["allowlist"], _ = expandWebfilterProfileWebAllowlist(d, i["allowlist"], pre_append)
	}
	pre_append = pre + ".0." + "blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["blocklist"], _ = expandWebfilterProfileWebBlocklist(d, i["blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "bword_table"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-table"], _ = expandWebfilterProfileWebBwordTable(d, i["bword_table"], pre_append)
	}
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-threshold"], _ = expandWebfilterProfileWebBwordThreshold(d, i["bword_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "content_header_list"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-header-list"], _ = expandWebfilterProfileWebContentHeaderList(d, i["content_header_list"], pre_append)
	}
	pre_append = pre + ".0." + "keyword_match"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["keyword-match"], _ = expandWebfilterProfileWebKeywordMatch(d, i["keyword_match"], pre_append)
	}
	pre_append = pre + ".0." + "log_search"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-search"], _ = expandWebfilterProfileWebLogSearch(d, i["log_search"], pre_append)
	}
	pre_append = pre + ".0." + "safe_search"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["safe-search"], _ = expandWebfilterProfileWebSafeSearch(d, i["safe_search"], pre_append)
	}
	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["urlfilter-table"], _ = expandWebfilterProfileWebUrlfilterTable(d, i["urlfilter_table"], pre_append)
	}
	pre_append = pre + ".0." + "whitelist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["whitelist"], _ = expandWebfilterProfileWebWhitelist(d, i["whitelist"], pre_append)
	}
	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vimeo-restrict"], _ = expandWebfilterProfileWebVimeoRestrict(d, i["vimeo_restrict"], pre_append)
	}
	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["youtube-restrict"], _ = expandWebfilterProfileWebYoutubeRestrict(d, i["youtube_restrict"], pre_append)
	}
	pre_append = pre + ".0." + "qwant_restrict"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["qwant-restrict"], _ = expandWebfilterProfileWebQwantRestrict(d, i["qwant_restrict"], pre_append)
	}

	return result, nil
}

func expandWebfilterProfileWebBlacklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebAllowlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentHeaderList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebKeywordMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebLogSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebUrlfilterTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileWebVimeoRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebQwantRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebAntiphishingLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebExtendedAllActionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterActivexLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterAppletLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCommandBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieRemovalLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJscriptLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterRefererLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterUnknownLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterVbsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFlowLogEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdErrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdQuotaUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebInvalidDomainLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebUrlLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWisp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["channel-id"], _ = expandWebfilterProfileYoutubeChannelFilterChannelId(d, i["channel_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWebfilterProfileYoutubeChannelFilterComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWebfilterProfileYoutubeChannelFilterId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileYoutubeChannelFilterChannelId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileIaCategorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("antiphish"); ok || d.HasChange("antiphish") {
		t, err := expandWebfilterProfileAntiphish(d, v, "antiphish")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandWebfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandWebfilterProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandWebfilterProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_wf"); ok || d.HasChange("ftgd_wf") {
		t, err := expandWebfilterProfileFtgdWf(d, v, "ftgd_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-wf"] = t
		}
	}

	if v, ok := d.GetOk("https_replacemsg"); ok || d.HasChange("https_replacemsg") {
		t, err := expandWebfilterProfileHttpsReplacemsg(d, v, "https_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("log_all_url"); ok || d.HasChange("log_all_url") {
		t, err := expandWebfilterProfileLogAllUrl(d, v, "log_all_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandWebfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandWebfilterProfileOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_perm"); ok || d.HasChange("ovrd_perm") {
		t, err := expandWebfilterProfileOvrdPerm(d, v, "ovrd_perm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-perm"] = t
		}
	}

	if v, ok := d.GetOk("post_action"); ok || d.HasChange("post_action") {
		t, err := expandWebfilterProfilePostAction(d, v, "post_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-action"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandWebfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("url_extraction"); ok || d.HasChange("url_extraction") {
		t, err := expandWebfilterProfileUrlExtraction(d, v, "url_extraction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-extraction"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok || d.HasChange("web") {
		t, err := expandWebfilterProfileWeb(d, v, "web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web"] = t
		}
	}

	if v, ok := d.GetOk("web_antiphishing_log"); ok || d.HasChange("web_antiphishing_log") {
		t, err := expandWebfilterProfileWebAntiphishingLog(d, v, "web_antiphishing_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-antiphishing-log"] = t
		}
	}

	if v, ok := d.GetOk("web_content_log"); ok || d.HasChange("web_content_log") {
		t, err := expandWebfilterProfileWebContentLog(d, v, "web_content_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("web_extended_all_action_log"); ok || d.HasChange("web_extended_all_action_log") {
		t, err := expandWebfilterProfileWebExtendedAllActionLog(d, v, "web_extended_all_action_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-extended-all-action-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_activex_log"); ok || d.HasChange("web_filter_activex_log") {
		t, err := expandWebfilterProfileWebFilterActivexLog(d, v, "web_filter_activex_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-activex-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_applet_log"); ok || d.HasChange("web_filter_applet_log") {
		t, err := expandWebfilterProfileWebFilterAppletLog(d, v, "web_filter_applet_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-applet-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_command_block_log"); ok || d.HasChange("web_filter_command_block_log") {
		t, err := expandWebfilterProfileWebFilterCommandBlockLog(d, v, "web_filter_command_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-command-block-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_log"); ok || d.HasChange("web_filter_cookie_log") {
		t, err := expandWebfilterProfileWebFilterCookieLog(d, v, "web_filter_cookie_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_removal_log"); ok || d.HasChange("web_filter_cookie_removal_log") {
		t, err := expandWebfilterProfileWebFilterCookieRemovalLog(d, v, "web_filter_cookie_removal_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-removal-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_js_log"); ok || d.HasChange("web_filter_js_log") {
		t, err := expandWebfilterProfileWebFilterJsLog(d, v, "web_filter_js_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-js-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_jscript_log"); ok || d.HasChange("web_filter_jscript_log") {
		t, err := expandWebfilterProfileWebFilterJscriptLog(d, v, "web_filter_jscript_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-jscript-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_referer_log"); ok || d.HasChange("web_filter_referer_log") {
		t, err := expandWebfilterProfileWebFilterRefererLog(d, v, "web_filter_referer_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-referer-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_unknown_log"); ok || d.HasChange("web_filter_unknown_log") {
		t, err := expandWebfilterProfileWebFilterUnknownLog(d, v, "web_filter_unknown_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-unknown-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_vbs_log"); ok || d.HasChange("web_filter_vbs_log") {
		t, err := expandWebfilterProfileWebFilterVbsLog(d, v, "web_filter_vbs_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-vbs-log"] = t
		}
	}

	if v, ok := d.GetOk("web_flow_log_encoding"); ok || d.HasChange("web_flow_log_encoding") {
		t, err := expandWebfilterProfileWebFlowLogEncoding(d, v, "web_flow_log_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-flow-log-encoding"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_err_log"); ok || d.HasChange("web_ftgd_err_log") {
		t, err := expandWebfilterProfileWebFtgdErrLog(d, v, "web_ftgd_err_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_quota_usage"); ok || d.HasChange("web_ftgd_quota_usage") {
		t, err := expandWebfilterProfileWebFtgdQuotaUsage(d, v, "web_ftgd_quota_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-quota-usage"] = t
		}
	}

	if v, ok := d.GetOk("web_invalid_domain_log"); ok || d.HasChange("web_invalid_domain_log") {
		t, err := expandWebfilterProfileWebInvalidDomainLog(d, v, "web_invalid_domain_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-invalid-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("web_url_log"); ok || d.HasChange("web_url_log") {
		t, err := expandWebfilterProfileWebUrlLog(d, v, "web_url_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-url-log"] = t
		}
	}

	if v, ok := d.GetOk("wisp"); ok || d.HasChange("wisp") {
		t, err := expandWebfilterProfileWisp(d, v, "wisp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp"] = t
		}
	}

	if v, ok := d.GetOk("wisp_algorithm"); ok || d.HasChange("wisp_algorithm") {
		t, err := expandWebfilterProfileWispAlgorithm(d, v, "wisp_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("wisp_servers"); ok || d.HasChange("wisp_servers") {
		t, err := expandWebfilterProfileWispServers(d, v, "wisp_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-servers"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok || d.HasChange("youtube_channel_filter") {
		t, err := expandWebfilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_status"); ok || d.HasChange("youtube_channel_status") {
		t, err := expandWebfilterProfileYoutubeChannelStatus(d, v, "youtube_channel_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-status"] = t
		}
	}

	if v, ok := d.GetOk("ia_categorization"); ok || d.HasChange("ia_categorization") {
		t, err := expandWebfilterProfileIaCategorization(d, v, "ia_categorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ia-categorization"] = t
		}
	}

	return &obj, nil
}
