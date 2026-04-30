---
subcategory: "LLM"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_llm_server"
description: |-
  Device LlmServer
---

# fmgdevice_llm_server
Device LlmServer

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `accept_custom_model` - Accept-Custom-Model. Valid values: `disable`, `enable`.

* `anthropic_version` - Anthropic-Version.
* `api_key` - Api-Key.
* `azure_api_version` - Azure-Api-Version.
* `azure_resource_name` - Azure-Resource-Name.
* `built_in_server` - Built-In-Server. Valid values: `openai`, `azure`, `azure-openai`, `gemini`, `anthropic`, `grok`, `gemini-with-openai-api`, `anthropic-with-openai-api`.

* `chat_completions_api` - Chat-Completions-Api. Valid values: `none`, `openai`, `azure-openai`, `gemini`, `anthropic`.

* `custom_model_allow_regex` - Custom-Model-Allow-Regex.
* `custom_model_block_regex` - Custom-Model-Block-Regex.
* `display_name` - Display-Name.
* `end_point` - End-Point.
* `image_gen_api` - Image-Gen-Api. Valid values: `none`, `openai`, `azure-openai`.

* `models` - Models.
* `name` - Name.
* `type` - Type. Valid values: `built-in`, `customized`.

* `verify_cert` - Verify-Cert. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Llm Server can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_llm_server.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

