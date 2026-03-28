---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_saasapplication_inputattributes"
description: |-
  <i>This object will be purged after policy copy and install.</i> SaaS application input attributes.
---

# fmgdevice_casb_saasapplication_inputattributes
<i>This object will be purged after policy copy and install.</i> SaaS application input attributes.

~> This resource is a sub resource for variable `input_attributes` of resource `fmgdevice_casb_saasapplication`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `saas_application` - Saas Application.

* `attr_type` - Attr-Type. Valid values: `tenant`.

* `default` - CASB attribute default value. Valid values: `string`, `string-list`.

* `description` - CASB attribute description.
* `fallback_input` - CASB attribute legacy input. Valid values: `disable`, `enable`.

* `name` - CASB attribute name.
* `required` - CASB input attribute required. Valid values: `disable`, `enable`.

* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb SaasApplicationInputAttributes can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_saasapplication_inputattributes.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

