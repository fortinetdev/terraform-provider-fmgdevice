---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_signature"
description: |-
  <i>This object will be purged after policy copy and install.</i> WAF signatures.
---

# fmgdevice_waf_profile_signature
<i>This object will be purged after policy copy and install.</i> WAF signatures.

~> This resource is a sub resource for variable `signature` of resource `fmgdevice_waf_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_signature`: `fmgdevice_waf_profile_signature_customsignature`
>- `main_class`: `fmgdevice_waf_profile_signature_mainclass`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom-Signature. The structure of `custom_signature` block is documented below.
* `disabled_signature` - Disabled signatures.
* `disabled_sub_class` - Disabled signature subclasses.
* `main_class` - Main-Class. The structure of `main_class` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_signature` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `direction` - Traffic direction. Valid values: `request`, `response`.

* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `target` - Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.


The `main_class` block supports:

* `action` - Action. Valid values: `allow`, `block`, `erase`.

* `id` - Main signature class ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Waf ProfileSignature can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_signature.labelname WafProfileSignature
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

