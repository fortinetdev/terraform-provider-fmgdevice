// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device LlmServer

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmServerCreate,
		Read:   resourceLlmServerRead,
		Update: resourceLlmServerUpdate,
		Delete: resourceLlmServerDelete,

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
			"accept_custom_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anthropic_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"azure_api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"built_in_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"chat_completions_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_model_allow_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_model_block_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_point": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_gen_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"models": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLlmServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmServer(d)
	if err != nil {
		return fmt.Errorf("Error creating LlmServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLlmServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLlmServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LlmServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateLlmServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LlmServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceLlmServerRead(d, m)
}

func resourceLlmServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmServer(d)
	if err != nil {
		return fmt.Errorf("Error updating LlmServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLlmServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LlmServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceLlmServerRead(d, m)
}

func resourceLlmServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLlmServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LlmServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LlmServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LlmServer resource from API: %v", err)
	}
	return nil
}

func flattenLlmServerAcceptCustomModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerAnthropicVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmServerAzureApiVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerAzureResourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerBuiltInServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerChatCompletionsApi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerCustomModelAllowRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerCustomModelBlockRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerDisplayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerEndPoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerImageGenApi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerModels(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenLlmServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLlmServerVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLlmServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accept_custom_model", flattenLlmServerAcceptCustomModel(o["accept-custom-model"], d, "accept_custom_model")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-custom-model"], "LlmServer-AcceptCustomModel"); ok {
			if err = d.Set("accept_custom_model", vv); err != nil {
				return fmt.Errorf("Error reading accept_custom_model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_custom_model: %v", err)
		}
	}

	if err = d.Set("anthropic_version", flattenLlmServerAnthropicVersion(o["anthropic-version"], d, "anthropic_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["anthropic-version"], "LlmServer-AnthropicVersion"); ok {
			if err = d.Set("anthropic_version", vv); err != nil {
				return fmt.Errorf("Error reading anthropic_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading anthropic_version: %v", err)
		}
	}

	if err = d.Set("api_key", flattenLlmServerApiKey(o["api-key"], d, "api_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-key"], "LlmServer-ApiKey"); ok {
			if err = d.Set("api_key", vv); err != nil {
				return fmt.Errorf("Error reading api_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_key: %v", err)
		}
	}

	if err = d.Set("azure_api_version", flattenLlmServerAzureApiVersion(o["azure-api-version"], d, "azure_api_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-api-version"], "LlmServer-AzureApiVersion"); ok {
			if err = d.Set("azure_api_version", vv); err != nil {
				return fmt.Errorf("Error reading azure_api_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_api_version: %v", err)
		}
	}

	if err = d.Set("azure_resource_name", flattenLlmServerAzureResourceName(o["azure-resource-name"], d, "azure_resource_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-resource-name"], "LlmServer-AzureResourceName"); ok {
			if err = d.Set("azure_resource_name", vv); err != nil {
				return fmt.Errorf("Error reading azure_resource_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_resource_name: %v", err)
		}
	}

	if err = d.Set("built_in_server", flattenLlmServerBuiltInServer(o["built-in-server"], d, "built_in_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["built-in-server"], "LlmServer-BuiltInServer"); ok {
			if err = d.Set("built_in_server", vv); err != nil {
				return fmt.Errorf("Error reading built_in_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading built_in_server: %v", err)
		}
	}

	if err = d.Set("chat_completions_api", flattenLlmServerChatCompletionsApi(o["chat-completions-api"], d, "chat_completions_api")); err != nil {
		if vv, ok := fortiAPIPatch(o["chat-completions-api"], "LlmServer-ChatCompletionsApi"); ok {
			if err = d.Set("chat_completions_api", vv); err != nil {
				return fmt.Errorf("Error reading chat_completions_api: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading chat_completions_api: %v", err)
		}
	}

	if err = d.Set("custom_model_allow_regex", flattenLlmServerCustomModelAllowRegex(o["custom-model-allow-regex"], d, "custom_model_allow_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-model-allow-regex"], "LlmServer-CustomModelAllowRegex"); ok {
			if err = d.Set("custom_model_allow_regex", vv); err != nil {
				return fmt.Errorf("Error reading custom_model_allow_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_model_allow_regex: %v", err)
		}
	}

	if err = d.Set("custom_model_block_regex", flattenLlmServerCustomModelBlockRegex(o["custom-model-block-regex"], d, "custom_model_block_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-model-block-regex"], "LlmServer-CustomModelBlockRegex"); ok {
			if err = d.Set("custom_model_block_regex", vv); err != nil {
				return fmt.Errorf("Error reading custom_model_block_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_model_block_regex: %v", err)
		}
	}

	if err = d.Set("display_name", flattenLlmServerDisplayName(o["display-name"], d, "display_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-name"], "LlmServer-DisplayName"); ok {
			if err = d.Set("display_name", vv); err != nil {
				return fmt.Errorf("Error reading display_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_name: %v", err)
		}
	}

	if err = d.Set("end_point", flattenLlmServerEndPoint(o["end-point"], d, "end_point")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-point"], "LlmServer-EndPoint"); ok {
			if err = d.Set("end_point", vv); err != nil {
				return fmt.Errorf("Error reading end_point: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_point: %v", err)
		}
	}

	if err = d.Set("image_gen_api", flattenLlmServerImageGenApi(o["image-gen-api"], d, "image_gen_api")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-gen-api"], "LlmServer-ImageGenApi"); ok {
			if err = d.Set("image_gen_api", vv); err != nil {
				return fmt.Errorf("Error reading image_gen_api: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_gen_api: %v", err)
		}
	}

	if err = d.Set("models", flattenLlmServerModels(o["models"], d, "models")); err != nil {
		if vv, ok := fortiAPIPatch(o["models"], "LlmServer-Models"); ok {
			if err = d.Set("models", vv); err != nil {
				return fmt.Errorf("Error reading models: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading models: %v", err)
		}
	}

	if err = d.Set("name", flattenLlmServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "LlmServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenLlmServerType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "LlmServer-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenLlmServerVerifyCert(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "LlmServer-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	return nil
}

func flattenLlmServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLlmServerAcceptCustomModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAnthropicVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmServerAzureApiVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAzureResourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerBuiltInServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerChatCompletionsApi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerCustomModelAllowRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerCustomModelBlockRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerDisplayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerEndPoint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerImageGenApi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerModels(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandLlmServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLlmServerVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLlmServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_custom_model"); ok || d.HasChange("accept_custom_model") {
		t, err := expandLlmServerAcceptCustomModel(d, v, "accept_custom_model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-custom-model"] = t
		}
	}

	if v, ok := d.GetOk("anthropic_version"); ok || d.HasChange("anthropic_version") {
		t, err := expandLlmServerAnthropicVersion(d, v, "anthropic_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anthropic-version"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok || d.HasChange("api_key") {
		t, err := expandLlmServerApiKey(d, v, "api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("azure_api_version"); ok || d.HasChange("azure_api_version") {
		t, err := expandLlmServerAzureApiVersion(d, v, "azure_api_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-api-version"] = t
		}
	}

	if v, ok := d.GetOk("azure_resource_name"); ok || d.HasChange("azure_resource_name") {
		t, err := expandLlmServerAzureResourceName(d, v, "azure_resource_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-resource-name"] = t
		}
	}

	if v, ok := d.GetOk("built_in_server"); ok || d.HasChange("built_in_server") {
		t, err := expandLlmServerBuiltInServer(d, v, "built_in_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["built-in-server"] = t
		}
	}

	if v, ok := d.GetOk("chat_completions_api"); ok || d.HasChange("chat_completions_api") {
		t, err := expandLlmServerChatCompletionsApi(d, v, "chat_completions_api")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chat-completions-api"] = t
		}
	}

	if v, ok := d.GetOk("custom_model_allow_regex"); ok || d.HasChange("custom_model_allow_regex") {
		t, err := expandLlmServerCustomModelAllowRegex(d, v, "custom_model_allow_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-model-allow-regex"] = t
		}
	}

	if v, ok := d.GetOk("custom_model_block_regex"); ok || d.HasChange("custom_model_block_regex") {
		t, err := expandLlmServerCustomModelBlockRegex(d, v, "custom_model_block_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-model-block-regex"] = t
		}
	}

	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		t, err := expandLlmServerDisplayName(d, v, "display_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-name"] = t
		}
	}

	if v, ok := d.GetOk("end_point"); ok || d.HasChange("end_point") {
		t, err := expandLlmServerEndPoint(d, v, "end_point")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-point"] = t
		}
	}

	if v, ok := d.GetOk("image_gen_api"); ok || d.HasChange("image_gen_api") {
		t, err := expandLlmServerImageGenApi(d, v, "image_gen_api")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-gen-api"] = t
		}
	}

	if v, ok := d.GetOk("models"); ok || d.HasChange("models") {
		t, err := expandLlmServerModels(d, v, "models")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["models"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandLlmServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandLlmServerType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandLlmServerVerifyCert(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	return &obj, nil
}
