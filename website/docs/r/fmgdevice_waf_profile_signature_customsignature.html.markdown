---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_signature_customsignature"
description: |-
  <i>This object will be purged after policy copy and install.</i> Custom signature.
---

# fmgdevice_waf_profile_signature_customsignature
<i>This object will be purged after policy copy and install.</i> Custom signature.

~> This resource is a sub resource for variable `custom_signature` of resource `fmgdevice_waf_profile_signature`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `direction` - Traffic direction. Valid values: `request`, `response`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `target` - Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Waf ProfileSignatureCustomSignature can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_signature_customsignature.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

