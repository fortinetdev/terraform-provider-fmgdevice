---
subcategory: "LLM"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_llm_profile"
description: |-
  Device LlmProfile
---

# fmgdevice_llm_profile
Device LlmProfile

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `chat`: `fmgdevice_llm_profile_chat`
>- `image_gen`: `fmgdevice_llm_profile_imagegen`
>- `list_models`: `fmgdevice_llm_profile_listmodels`
>- `response`: `fmgdevice_llm_profile_response`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `chat` - Chat. The structure of `chat` block is documented below.
* `image_gen` - Image-Gen. The structure of `image_gen` block is documented below.
* `list_models` - List-Models. The structure of `list_models` block is documented below.
* `log` - Log. Valid values: `none`, `all`, `blocked`.

* `name` - Name.
* `response` - Response. The structure of `response` block is documented below.
* `unknown_api` - Unknown-Api. Valid values: `disable`, `enable`.


The `chat` block supports:

* `max_completion_tokens` - Max-Completion-Tokens.
* `max_req_len` - Max-Req-Len.
* `status` - Status. Valid values: `disable`, `enable`.

* `stream` - Stream. Valid values: `block`, `bypass`.

* `system_prompt` - System-Prompt.
* `system_prompt_mode` - System-Prompt-Mode. Valid values: `bypass`, `replace`, `prepend`, `append`.


The `image_gen` block supports:

* `status` - Status. Valid values: `disable`, `enable`.


The `list_models` block supports:

* `status` - Status. Valid values: `disable`, `enable`.


The `response` block supports:

* `instructions` - Instructions.
* `instructions_mode` - Instructions-Mode. Valid values: `bypass`, `replace`, `prepend`, `append`.

* `max_output_tokens` - Max-Output-Tokens.
* `max_req_len` - Max-Req-Len.
* `status` - Status. Valid values: `disable`, `enable`.

* `stream` - Stream. Valid values: `block`, `bypass`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Llm Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_llm_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

