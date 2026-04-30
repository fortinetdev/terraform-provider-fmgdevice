// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LlmProfileChat

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProfileChat() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProfileChatUpdate,
		Read:   resourceLlmProfileChatRead,
		Update: resourceLlmProfileChatUpdate,
		Delete: resourceLlmProfileChatDelete,

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
	}
}

func resourceLlmProfileChatUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfileChat(d, false)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfileChat resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProfileChat(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfileChat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LlmProfileChat")

	return resourceLlmProfileChatRead(d, m)
}

func resourceLlmProfileChatDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfileChat(d, true)

	if err != nil {
		return fmt.Errorf("Error updating LlmProfileChat resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProfileChat(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing LlmProfileChat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProfileChatRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProfileChat(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LlmProfileChat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProfileChat(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LlmProfileChat resource from API: %v", err)
	}
	return nil
}

func flattenLlmProfileChatMaxCompletionTokens2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatMaxReqLen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatStream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatSystemPrompt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileChatSystemPromptMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLlmProfileChat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_completion_tokens", flattenLlmProfileChatMaxCompletionTokens2edl(o["max-completion-tokens"], d, "max_completion_tokens")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-completion-tokens"], "LlmProfileChat-MaxCompletionTokens"); ok {
			if err = d.Set("max_completion_tokens", vv); err != nil {
				return fmt.Errorf("Error reading max_completion_tokens: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_completion_tokens: %v", err)
		}
	}

	if err = d.Set("max_req_len", flattenLlmProfileChatMaxReqLen2edl(o["max-req-len"], d, "max_req_len")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-req-len"], "LlmProfileChat-MaxReqLen"); ok {
			if err = d.Set("max_req_len", vv); err != nil {
				return fmt.Errorf("Error reading max_req_len: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_req_len: %v", err)
		}
	}

	if err = d.Set("status", flattenLlmProfileChatStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LlmProfileChat-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stream", flattenLlmProfileChatStream2edl(o["stream"], d, "stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream"], "LlmProfileChat-Stream"); ok {
			if err = d.Set("stream", vv); err != nil {
				return fmt.Errorf("Error reading stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream: %v", err)
		}
	}

	if err = d.Set("system_prompt", flattenLlmProfileChatSystemPrompt2edl(o["system-prompt"], d, "system_prompt")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-prompt"], "LlmProfileChat-SystemPrompt"); ok {
			if err = d.Set("system_prompt", vv); err != nil {
				return fmt.Errorf("Error reading system_prompt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_prompt: %v", err)
		}
	}

	if err = d.Set("system_prompt_mode", flattenLlmProfileChatSystemPromptMode2edl(o["system-prompt-mode"], d, "system_prompt_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-prompt-mode"], "LlmProfileChat-SystemPromptMode"); ok {
			if err = d.Set("system_prompt_mode", vv); err != nil {
				return fmt.Errorf("Error reading system_prompt_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_prompt_mode: %v", err)
		}
	}

	return nil
}

func flattenLlmProfileChatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLlmProfileChatMaxCompletionTokens2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatMaxReqLen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatStream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPrompt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPromptMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProfileChat(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_completion_tokens"); ok || d.HasChange("max_completion_tokens") {
		t, err := expandLlmProfileChatMaxCompletionTokens2edl(d, v, "max_completion_tokens")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-completion-tokens"] = t
		}
	}

	if v, ok := d.GetOk("max_req_len"); ok || d.HasChange("max_req_len") {
		t, err := expandLlmProfileChatMaxReqLen2edl(d, v, "max_req_len")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-req-len"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLlmProfileChatStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stream"); ok || d.HasChange("stream") {
		t, err := expandLlmProfileChatStream2edl(d, v, "stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream"] = t
		}
	}

	if v, ok := d.GetOk("system_prompt"); ok || d.HasChange("system_prompt") {
		t, err := expandLlmProfileChatSystemPrompt2edl(d, v, "system_prompt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-prompt"] = t
		}
	}

	if v, ok := d.GetOk("system_prompt_mode"); ok || d.HasChange("system_prompt_mode") {
		t, err := expandLlmProfileChatSystemPromptMode2edl(d, v, "system_prompt_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-prompt-mode"] = t
		}
	}

	return &obj, nil
}
