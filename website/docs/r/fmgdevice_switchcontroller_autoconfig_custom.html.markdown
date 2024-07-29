---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_autoconfig_custom"
description: |-
  Policies which can override the 'default' for specific ISL/ICL/FortiLink interface.
---

# fmgdevice_switchcontroller_autoconfig_custom
Policies which can override the 'default' for specific ISL/ICL/FortiLink interface.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `switch_binding`: `fmgdevice_switchcontroller_autoconfig_custom_switchbinding`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_autoconfig_custom" "trname" {
  name = "your own value"
  switch_binding {
    policy    = ["your own value"]
    switch_id = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Auto-Config FortiLink or ISL/ICL interface name.
* `switch_binding` - Switch-Binding. The structure of `switch_binding` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `switch_binding` block supports:

* `policy` - Custom auto-config policy.
* `switch_id` - Switch name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController AutoConfigCustom can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_autoconfig_custom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

