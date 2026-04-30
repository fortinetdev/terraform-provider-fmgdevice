// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LlmProfileResponse

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProfileResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProfileResponseUpdate,
		Read:   resourceLlmProfileResponseRead,
		Update: resourceLlmProfileResponseUpdate,
		Delete: resourceLlmProfileResponseDelete,

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
	}
}

func resourceLlmProfileResponseUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfileResponse(d, false)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfileResponse resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProfileResponse(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfileResponse resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LlmProfileResponse")

	return resourceLlmProfileResponseRead(d, m)
}

func resourceLlmProfileResponseDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfileResponse(d, true)

	if err != nil {
		return fmt.Errorf("Error updating LlmProfileResponse resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmProfileResponse(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing LlmProfileResponse resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProfileResponseRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProfileResponse(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LlmProfileResponse resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProfileResponse(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LlmProfileResponse resource from API: %v", err)
	}
	return nil
}

func flattenLlmProfileResponseInstructions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseInstructionsMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseMaxOutputTokens2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseMaxReqLen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmProfileResponseStream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLlmProfileResponse(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("instructions", flattenLlmProfileResponseInstructions2edl(o["instructions"], d, "instructions")); err != nil {
		if vv, ok := fortiAPIPatch(o["instructions"], "LlmProfileResponse-Instructions"); ok {
			if err = d.Set("instructions", vv); err != nil {
				return fmt.Errorf("Error reading instructions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading instructions: %v", err)
		}
	}

	if err = d.Set("instructions_mode", flattenLlmProfileResponseInstructionsMode2edl(o["instructions-mode"], d, "instructions_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["instructions-mode"], "LlmProfileResponse-InstructionsMode"); ok {
			if err = d.Set("instructions_mode", vv); err != nil {
				return fmt.Errorf("Error reading instructions_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading instructions_mode: %v", err)
		}
	}

	if err = d.Set("max_output_tokens", flattenLlmProfileResponseMaxOutputTokens2edl(o["max-output-tokens"], d, "max_output_tokens")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-output-tokens"], "LlmProfileResponse-MaxOutputTokens"); ok {
			if err = d.Set("max_output_tokens", vv); err != nil {
				return fmt.Errorf("Error reading max_output_tokens: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_output_tokens: %v", err)
		}
	}

	if err = d.Set("max_req_len", flattenLlmProfileResponseMaxReqLen2edl(o["max-req-len"], d, "max_req_len")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-req-len"], "LlmProfileResponse-MaxReqLen"); ok {
			if err = d.Set("max_req_len", vv); err != nil {
				return fmt.Errorf("Error reading max_req_len: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_req_len: %v", err)
		}
	}

	if err = d.Set("status", flattenLlmProfileResponseStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "LlmProfileResponse-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stream", flattenLlmProfileResponseStream2edl(o["stream"], d, "stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream"], "LlmProfileResponse-Stream"); ok {
			if err = d.Set("stream", vv); err != nil {
				return fmt.Errorf("Error reading stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream: %v", err)
		}
	}

	return nil
}

func flattenLlmProfileResponseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLlmProfileResponseInstructions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseInstructionsMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseMaxOutputTokens2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseMaxReqLen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileResponseStream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProfileResponse(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("instructions"); ok || d.HasChange("instructions") {
		t, err := expandLlmProfileResponseInstructions2edl(d, v, "instructions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["instructions"] = t
		}
	}

	if v, ok := d.GetOk("instructions_mode"); ok || d.HasChange("instructions_mode") {
		t, err := expandLlmProfileResponseInstructionsMode2edl(d, v, "instructions_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["instructions-mode"] = t
		}
	}

	if v, ok := d.GetOk("max_output_tokens"); ok || d.HasChange("max_output_tokens") {
		t, err := expandLlmProfileResponseMaxOutputTokens2edl(d, v, "max_output_tokens")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-output-tokens"] = t
		}
	}

	if v, ok := d.GetOk("max_req_len"); ok || d.HasChange("max_req_len") {
		t, err := expandLlmProfileResponseMaxReqLen2edl(d, v, "max_req_len")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-req-len"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandLlmProfileResponseStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stream"); ok || d.HasChange("stream") {
		t, err := expandLlmProfileResponseStream2edl(d, v, "stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream"] = t
		}
	}

	return &obj, nil
}
