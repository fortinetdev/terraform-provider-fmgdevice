// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure Email Filter profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterProfileCreate,
		Read:   resourceEmailfilterProfileRead,
		Update: resourceEmailfilterProfileUpdate,
		Delete: resourceEmailfilterProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"msn_hotmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"other_webmails": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hdrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spam_bwl_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spam_bal_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spam_bword_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spam_bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_filtering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_iptrust_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spam_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_log_fortiguard_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spam_mheader_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spam_rbl_table": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEmailfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEmailfilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEmailfilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EmailfilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEmailfilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EmailfilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceEmailfilterProfileRead(d, m)
}

func resourceEmailfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEmailfilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceEmailfilterProfileRead(d, m)
}

func resourceEmailfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteEmailfilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EmailfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileGmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileGmailLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileGmailLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileGmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileGmailLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileImapAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileImapLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileImapLogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfileImapTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfileImapTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileImapAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileImapLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileImapLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileImapTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileImapTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileMapiAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileMapiLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileMapiLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileMapiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileMapiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileMapiLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileMsnHotmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileMsnHotmailLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileMsnHotmailLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileMsnHotmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileMsnHotmailLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileOtherWebmails(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileOtherWebmailsLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileOtherWebmailsLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfilePop3Action(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfilePop3Log(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfilePop3LogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfilePop3TagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfilePop3TagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfilePop3Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3LogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileSmtpAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "hdrip"
	if _, ok := i["hdrip"]; ok {
		result["hdrip"] = flattenEmailfilterProfileSmtpHdrip(i["hdrip"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenEmailfilterProfileSmtpLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileSmtpLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileSmtpLogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfileSmtpTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfileSmtpTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileSmtpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpHdrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamBwlTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamBalTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamFiltering(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamIptrustTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamLogFortiguardResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamMheaderTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEmailfilterProfileSpamRblTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectEmailfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenEmailfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "EmailfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("external", flattenEmailfilterProfileExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "EmailfilterProfile-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenEmailfilterProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "EmailfilterProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gmail", flattenEmailfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["gmail"], "EmailfilterProfile-Gmail"); ok {
				if err = d.Set("gmail", vv); err != nil {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading gmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gmail"); ok {
			if err = d.Set("gmail", flattenEmailfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["gmail"], "EmailfilterProfile-Gmail"); ok {
					if err = d.Set("gmail", vv); err != nil {
						return fmt.Errorf("Error reading gmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenEmailfilterProfileImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "EmailfilterProfile-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenEmailfilterProfileImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "EmailfilterProfile-Imap"); ok {
					if err = d.Set("imap", vv); err != nil {
						return fmt.Errorf("Error reading imap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenEmailfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "EmailfilterProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenEmailfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "EmailfilterProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("msn_hotmail", flattenEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["msn-hotmail"], "EmailfilterProfile-MsnHotmail"); ok {
				if err = d.Set("msn_hotmail", vv); err != nil {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading msn_hotmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msn_hotmail"); ok {
			if err = d.Set("msn_hotmail", flattenEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["msn-hotmail"], "EmailfilterProfile-MsnHotmail"); ok {
					if err = d.Set("msn_hotmail", vv); err != nil {
						return fmt.Errorf("Error reading msn_hotmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenEmailfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "EmailfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenEmailfilterProfileOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "EmailfilterProfile-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("other_webmails", flattenEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails")); err != nil {
			if vv, ok := fortiAPIPatch(o["other-webmails"], "EmailfilterProfile-OtherWebmails"); ok {
				if err = d.Set("other_webmails", vv); err != nil {
					return fmt.Errorf("Error reading other_webmails: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading other_webmails: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("other_webmails"); ok {
			if err = d.Set("other_webmails", flattenEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails")); err != nil {
				if vv, ok := fortiAPIPatch(o["other-webmails"], "EmailfilterProfile-OtherWebmails"); ok {
					if err = d.Set("other_webmails", vv); err != nil {
						return fmt.Errorf("Error reading other_webmails: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading other_webmails: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenEmailfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "EmailfilterProfile-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenEmailfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "EmailfilterProfile-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenEmailfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "EmailfilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenEmailfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "EmailfilterProfile-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenEmailfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "EmailfilterProfile-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_bwl_table", flattenEmailfilterProfileSpamBwlTable(o["spam-bwl-table"], d, "spam_bwl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bwl-table"], "EmailfilterProfile-SpamBwlTable"); ok {
			if err = d.Set("spam_bwl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bwl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bwl_table: %v", err)
		}
	}

	if err = d.Set("spam_bal_table", flattenEmailfilterProfileSpamBalTable(o["spam-bal-table"], d, "spam_bal_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bal-table"], "EmailfilterProfile-SpamBalTable"); ok {
			if err = d.Set("spam_bal_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bal_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bal_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_table", flattenEmailfilterProfileSpamBwordTable(o["spam-bword-table"], d, "spam_bword_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-table"], "EmailfilterProfile-SpamBwordTable"); ok {
			if err = d.Set("spam_bword_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_threshold", flattenEmailfilterProfileSpamBwordThreshold(o["spam-bword-threshold"], d, "spam_bword_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-threshold"], "EmailfilterProfile-SpamBwordThreshold"); ok {
			if err = d.Set("spam_bword_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
		}
	}

	if err = d.Set("spam_filtering", flattenEmailfilterProfileSpamFiltering(o["spam-filtering"], d, "spam_filtering")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-filtering"], "EmailfilterProfile-SpamFiltering"); ok {
			if err = d.Set("spam_filtering", vv); err != nil {
				return fmt.Errorf("Error reading spam_filtering: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_filtering: %v", err)
		}
	}

	if err = d.Set("spam_iptrust_table", flattenEmailfilterProfileSpamIptrustTable(o["spam-iptrust-table"], d, "spam_iptrust_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-iptrust-table"], "EmailfilterProfile-SpamIptrustTable"); ok {
			if err = d.Set("spam_iptrust_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
		}
	}

	if err = d.Set("spam_log", flattenEmailfilterProfileSpamLog(o["spam-log"], d, "spam_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log"], "EmailfilterProfile-SpamLog"); ok {
			if err = d.Set("spam_log", vv); err != nil {
				return fmt.Errorf("Error reading spam_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log: %v", err)
		}
	}

	if err = d.Set("spam_log_fortiguard_response", flattenEmailfilterProfileSpamLogFortiguardResponse(o["spam-log-fortiguard-response"], d, "spam_log_fortiguard_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log-fortiguard-response"], "EmailfilterProfile-SpamLogFortiguardResponse"); ok {
			if err = d.Set("spam_log_fortiguard_response", vv); err != nil {
				return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if err = d.Set("spam_mheader_table", flattenEmailfilterProfileSpamMheaderTable(o["spam-mheader-table"], d, "spam_mheader_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-mheader-table"], "EmailfilterProfile-SpamMheaderTable"); ok {
			if err = d.Set("spam_mheader_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_mheader_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_mheader_table: %v", err)
		}
	}

	if err = d.Set("spam_rbl_table", flattenEmailfilterProfileSpamRblTable(o["spam-rbl-table"], d, "spam_rbl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-rbl-table"], "EmailfilterProfile-SpamRblTable"); ok {
			if err = d.Set("spam_rbl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_rbl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_rbl_table: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEmailfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfileGmailLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileGmailLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileGmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileGmailLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandEmailfilterProfileImapAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfileImapLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileImapLogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandEmailfilterProfileImapTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandEmailfilterProfileImapTagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileImapAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandEmailfilterProfileMapiAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfileMapiLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileMapiLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileMapiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMapiLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMapiLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfileMsnHotmailLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileMsnHotmailLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileMsnHotmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMsnHotmailLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileOtherWebmails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileOtherWebmailsLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileOtherWebmailsLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandEmailfilterProfilePop3Action(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfilePop3Log(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfilePop3LogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandEmailfilterProfilePop3TagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandEmailfilterProfilePop3TagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfilePop3Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3LogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandEmailfilterProfileSmtpAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "hdrip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hdrip"], _ = expandEmailfilterProfileSmtpHdrip(d, i["hdrip"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-override"], _ = expandEmailfilterProfileSmtpLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandEmailfilterProfileSmtpLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-all"], _ = expandEmailfilterProfileSmtpLogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-msg"], _ = expandEmailfilterProfileSmtpTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag-type"], _ = expandEmailfilterProfileSmtpTagType(d, i["tag_type"], pre_append)
	}

	return result, nil
}

func expandEmailfilterProfileSmtpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpHdrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamBwlTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamBalTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamFiltering(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamIptrustTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamLogFortiguardResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamMheaderTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEmailfilterProfileSpamRblTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectEmailfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandEmailfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok || d.HasChange("external") {
		t, err := expandEmailfilterProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandEmailfilterProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("gmail"); ok || d.HasChange("gmail") {
		t, err := expandEmailfilterProfileGmail(d, v, "gmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gmail"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok || d.HasChange("imap") {
		t, err := expandEmailfilterProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok || d.HasChange("mapi") {
		t, err := expandEmailfilterProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("msn_hotmail"); ok || d.HasChange("msn_hotmail") {
		t, err := expandEmailfilterProfileMsnHotmail(d, v, "msn_hotmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msn-hotmail"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandEmailfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandEmailfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("other_webmails"); ok || d.HasChange("other_webmails") {
		t, err := expandEmailfilterProfileOtherWebmails(d, v, "other_webmails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-webmails"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok || d.HasChange("pop3") {
		t, err := expandEmailfilterProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandEmailfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok || d.HasChange("smtp") {
		t, err := expandEmailfilterProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("spam_bwl_table"); ok || d.HasChange("spam_bwl_table") {
		t, err := expandEmailfilterProfileSpamBwlTable(d, v, "spam_bwl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bwl-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bal_table"); ok || d.HasChange("spam_bal_table") {
		t, err := expandEmailfilterProfileSpamBalTable(d, v, "spam_bal_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bal-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_table"); ok || d.HasChange("spam_bword_table") {
		t, err := expandEmailfilterProfileSpamBwordTable(d, v, "spam_bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_threshold"); ok || d.HasChange("spam_bword_threshold") {
		t, err := expandEmailfilterProfileSpamBwordThreshold(d, v, "spam_bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spam_filtering"); ok || d.HasChange("spam_filtering") {
		t, err := expandEmailfilterProfileSpamFiltering(d, v, "spam_filtering")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-filtering"] = t
		}
	}

	if v, ok := d.GetOk("spam_iptrust_table"); ok || d.HasChange("spam_iptrust_table") {
		t, err := expandEmailfilterProfileSpamIptrustTable(d, v, "spam_iptrust_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-iptrust-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_log"); ok || d.HasChange("spam_log") {
		t, err := expandEmailfilterProfileSpamLog(d, v, "spam_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log"] = t
		}
	}

	if v, ok := d.GetOk("spam_log_fortiguard_response"); ok || d.HasChange("spam_log_fortiguard_response") {
		t, err := expandEmailfilterProfileSpamLogFortiguardResponse(d, v, "spam_log_fortiguard_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log-fortiguard-response"] = t
		}
	}

	if v, ok := d.GetOk("spam_mheader_table"); ok || d.HasChange("spam_mheader_table") {
		t, err := expandEmailfilterProfileSpamMheaderTable(d, v, "spam_mheader_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-mheader-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_rbl_table"); ok || d.HasChange("spam_rbl_table") {
		t, err := expandEmailfilterProfileSpamRblTable(d, v, "spam_rbl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-rbl-table"] = t
		}
	}

	return &obj, nil
}
