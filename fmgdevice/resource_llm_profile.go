// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LlmProfile

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProfileCreate,
		Read:   resourceLlmProfileRead,
		Update: resourceLlmProfileUpdate,
		Delete: resourceLlmProfileDelete,

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
			"chat": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_completion_tokens": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_req_len": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_prompt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_prompt_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"image_gen": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"list_models": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instructions": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instructions_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_output_tokens": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_req_len": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"unknown_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLlmProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating LlmProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLlmProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLlmProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LlmProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateLlmProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LlmProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceLlmProfileRead(d, m)
}

func resourceLlmProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceLlmProfileRead(d, m)
}

func resourceLlmProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLlmProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LlmProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LlmProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LlmProfile resource from API: %v", err)
	}
	return nil
}

func flattenLlmProfileChat(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_completion_tokens"
	if _, ok := i["max-completion-tokens"]; ok {
		result["max_completion_tokens"] = flattenLlmProfileChatMaxCompletionTokens(i["max-completion-tokens"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_req_len"
	if _, ok := i["max-req-len"]; ok {
		result["max_req_len"] = flattenLlmProfileChatMaxReqLen(i["max-req-len"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileChatStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream"
	if _, ok := i["stream"]; ok {
		result["stream"] = flattenLlmProfileChatStream(i["stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_prompt"
	if _, ok := i["system-prompt"]; ok {
		result["system_prompt"] = flattenLlmProfileChatSystemPrompt(i["system-prompt"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_prompt_mode"
	if _, ok := i["system-prompt-mode"]; ok {
		result["system_prompt_mode"] = flattenLlmProfileChatSystemPromptMode(i["system-prompt-mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileChatMaxCompletionTokens(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatMaxReqLen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatSystemPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatSystemPromptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileImageGen(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileImageGenStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileImageGenStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileListModels(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileListModelsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileListModelsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponse(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "instructions"
	if _, ok := i["instructions"]; ok {
		result["instructions"] = flattenLlmProfileResponseInstructions(i["instructions"], d, pre_append)
	}

	pre_append = pre + ".0." + "instructions_mode"
	if _, ok := i["instructions-mode"]; ok {
		result["instructions_mode"] = flattenLlmProfileResponseInstructionsMode(i["instructions-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_output_tokens"
	if _, ok := i["max-output-tokens"]; ok {
		result["max_output_tokens"] = flattenLlmProfileResponseMaxOutputTokens(i["max-output-tokens"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_req_len"
	if _, ok := i["max-req-len"]; ok {
		result["max_req_len"] = flattenLlmProfileResponseMaxReqLen(i["max-req-len"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileResponseStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream"
	if _, ok := i["stream"]; ok {
		result["stream"] = flattenLlmProfileResponseStream(i["stream"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileResponseInstructions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseInstructionsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseMaxOutputTokens(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseMaxReqLen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileUnknownApi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLlmProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("chat", flattenLlmProfileChat(o["chat"], d, "chat")); err != nil {
			if vv, ok := fortiAPIPatch(o["chat"], "LlmProfile-Chat"); ok {
				if err = d.Set("chat", vv); err != nil {
					return fmt.Errorf("Error reading chat: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading chat: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("chat"); ok {
			if err = d.Set("chat", flattenLlmProfileChat(o["chat"], d, "chat")); err != nil {
				if vv, ok := fortiAPIPatch(o["chat"], "LlmProfile-Chat"); ok {
					if err = d.Set("chat", vv); err != nil {
						return fmt.Errorf("Error reading chat: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading chat: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("image_gen", flattenLlmProfileImageGen(o["image-gen"], d, "image_gen")); err != nil {
			if vv, ok := fortiAPIPatch(o["image-gen"], "LlmProfile-ImageGen"); ok {
				if err = d.Set("image_gen", vv); err != nil {
					return fmt.Errorf("Error reading image_gen: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading image_gen: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("image_gen"); ok {
			if err = d.Set("image_gen", flattenLlmProfileImageGen(o["image-gen"], d, "image_gen")); err != nil {
				if vv, ok := fortiAPIPatch(o["image-gen"], "LlmProfile-ImageGen"); ok {
					if err = d.Set("image_gen", vv); err != nil {
						return fmt.Errorf("Error reading image_gen: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading image_gen: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("list_models", flattenLlmProfileListModels(o["list-models"], d, "list_models")); err != nil {
			if vv, ok := fortiAPIPatch(o["list-models"], "LlmProfile-ListModels"); ok {
				if err = d.Set("list_models", vv); err != nil {
					return fmt.Errorf("Error reading list_models: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list_models: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list_models"); ok {
			if err = d.Set("list_models", flattenLlmProfileListModels(o["list-models"], d, "list_models")); err != nil {
				if vv, ok := fortiAPIPatch(o["list-models"], "LlmProfile-ListModels"); ok {
					if err = d.Set("list_models", vv); err != nil {
						return fmt.Errorf("Error reading list_models: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list_models: %v", err)
				}
			}
		}
	}

	if err = d.Set("log", flattenLlmProfileLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "LlmProfile-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenLlmProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LlmProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("response", flattenLlmProfileResponse(o["response"], d, "response")); err != nil {
			if vv, ok := fortiAPIPatch(o["response"], "LlmProfile-Response"); ok {
				if err = d.Set("response", vv); err != nil {
					return fmt.Errorf("Error reading response: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading response: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("response"); ok {
			if err = d.Set("response", flattenLlmProfileResponse(o["response"], d, "response")); err != nil {
				if vv, ok := fortiAPIPatch(o["response"], "LlmProfile-Response"); ok {
					if err = d.Set("response", vv); err != nil {
						return fmt.Errorf("Error reading response: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading response: %v", err)
				}
			}
		}
	}

	if err = d.Set("unknown_api", flattenLlmProfileUnknownApi(o["unknown-api"], d, "unknown_api")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-api"], "LlmProfile-UnknownApi"); ok {
			if err = d.Set("unknown_api", vv); err != nil {
				return fmt.Errorf("Error reading unknown_api: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_api: %v", err)
		}
	}

	return nil
}

func flattenLlmProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLlmProfileChat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_completion_tokens"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-completion-tokens"], _ = expandLlmProfileChatMaxCompletionTokens(d, i["max_completion_tokens"], pre_append)
	}
	pre_append = pre + ".0." + "max_req_len"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-req-len"], _ = expandLlmProfileChatMaxReqLen(d, i["max_req_len"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandLlmProfileChatStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stream"], _ = expandLlmProfileChatStream(d, i["stream"], pre_append)
	}
	pre_append = pre + ".0." + "system_prompt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-prompt"], _ = expandLlmProfileChatSystemPrompt(d, i["system_prompt"], pre_append)
	}
	pre_append = pre + ".0." + "system_prompt_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-prompt-mode"], _ = expandLlmProfileChatSystemPromptMode(d, i["system_prompt_mode"], pre_append)
	}

	return result, nil
}

func expandLlmProfileChatMaxCompletionTokens(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatMaxReqLen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPrompt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPromptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileImageGen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandLlmProfileImageGenStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandLlmProfileImageGenStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileListModels(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandLlmProfileListModelsStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandLlmProfileListModelsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "instructions"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["instructions"], _ = expandLlmProfileResponseInstructions(d, i["instructions"], pre_append)
	}
	pre_append = pre + ".0." + "instructions_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["instructions-mode"], _ = expandLlmProfileResponseInstructionsMode(d, i["instructions_mode"], pre_append)
	}
	pre_append = pre + ".0." + "max_output_tokens"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-output-tokens"], _ = expandLlmProfileResponseMaxOutputTokens(d, i["max_output_tokens"], pre_append)
	}
	pre_append = pre + ".0." + "max_req_len"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-req-len"], _ = expandLlmProfileResponseMaxReqLen(d, i["max_req_len"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandLlmProfileResponseStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stream"], _ = expandLlmProfileResponseStream(d, i["stream"], pre_append)
	}

	return result, nil
}

func expandLlmProfileResponseInstructions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseInstructionsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseMaxOutputTokens(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseMaxReqLen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileUnknownApi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("chat"); ok || d.HasChange("chat") {
		t, err := expandLlmProfileChat(d, v, "chat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chat"] = t
		}
	}

	if v, ok := d.GetOk("image_gen"); ok || d.HasChange("image_gen") {
		t, err := expandLlmProfileImageGen(d, v, "image_gen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-gen"] = t
		}
	}

	if v, ok := d.GetOk("list_models"); ok || d.HasChange("list_models") {
		t, err := expandLlmProfileListModels(d, v, "list_models")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list-models"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandLlmProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLlmProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("response"); ok || d.HasChange("response") {
		t, err := expandLlmProfileResponse(d, v, "response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response"] = t
		}
	}

	if v, ok := d.GetOk("unknown_api"); ok || d.HasChange("unknown_api") {
		t, err := expandLlmProfileUnknownApi(d, v, "unknown_api")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-api"] = t
		}
	}

	return &obj, nil
}
