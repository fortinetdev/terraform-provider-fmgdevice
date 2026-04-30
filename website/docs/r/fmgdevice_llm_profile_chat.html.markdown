---
subcategory: "LLM"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_llm_profile_chat"
description: |-
  Device LlmProfileChat
---

# fmgdevice_llm_profile_chat
Device LlmProfileChat

~> This resource is a sub resource for variable `chat` of resource `fmgdevice_llm_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `max_completion_tokens` - Max-Completion-Tokens.
* `max_req_len` - Max-Req-Len.
* `status` - Status. Valid values: `disable`, `enable`.

* `stream` - Stream. Valid values: `block`, `bypass`.

* `system_prompt` - System-Prompt.
* `system_prompt_mode` - System-Prompt-Mode. Valid values: `bypass`, `replace`, `prepend`, `append`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Llm ProfileChat can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_llm_profile_chat.labelname LlmProfileChat
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

