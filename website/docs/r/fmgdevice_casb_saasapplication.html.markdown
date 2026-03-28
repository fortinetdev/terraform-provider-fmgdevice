---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_saasapplication"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure CASB SaaS application.
---

# fmgdevice_casb_saasapplication
<i>This object will be purged after policy copy and install.</i> Configure CASB SaaS application.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `input_attributes`: `fmgdevice_casb_saasapplication_inputattributes`
>- `output_attributes`: `fmgdevice_casb_saasapplication_outputattributes`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `casb_name` - SaaS application signature name.
* `description` - SaaS application description.
* `domains` - SaaS application domain list.
* `input_attributes` - Input-Attributes. The structure of `input_attributes` block is documented below.
* `name` - SaaS application name.
* `output_attributes` - Output-Attributes. The structure of `output_attributes` block is documented below.
* `status` - Enable/disable setting. Valid values: `disable`, `enable`.

* `type` - SaaS application type. Valid values: `built-in`, `customized`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `input_attributes` block supports:

* `attr_type` - Attr-Type. Valid values: `tenant`.

* `default` - CASB attribute default value. Valid values: `string`, `string-list`.

* `description` - CASB attribute description.
* `fallback_input` - CASB attribute legacy input. Valid values: `disable`, `enable`.

* `name` - CASB attribute name.
* `required` - CASB input attribute required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.


The `output_attributes` block supports:

* `attr_type` - Attr-Type. Valid values: `tenant`.

* `description` - CASB attribute description.
* `name` - CASB attribute name.
* `optional` - CASB output attribute optional. Valid values: `disable`, `enable`.

* `required` - Required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb SaasApplication can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_saasapplication.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

