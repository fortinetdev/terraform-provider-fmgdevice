---
subcategory: "LLM"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_llm_profile_response"
description: |-
  Device LlmProfileResponse
---

# fmgdevice_llm_profile_response
Device LlmProfileResponse

~> This resource is a sub resource for variable `response` of resource `fmgdevice_llm_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `instructions` - Instructions.
* `instructions_mode` - Instructions-Mode. Valid values: `bypass`, `replace`, `prepend`, `append`.

* `max_output_tokens` - Max-Output-Tokens.
* `max_req_len` - Max-Req-Len.
* `status` - Status. Valid values: `disable`, `enable`.

* `stream` - Stream. Valid values: `block`, `bypass`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Llm ProfileResponse can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_llm_profile_response.labelname LlmProfileResponse
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

