---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_rip_interface"
description: |-
  RIP interface configuration.
---

# fmgdevice_router_rip_interface
RIP interface configuration.

~> This resource is a sub resource for variable `interface` of resource `fmgdevice_router_rip`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_rip_interface" "trname" {
  auth_keychain = ["your own value"]
  auth_mode     = "md5"
  auth_string   = ["your own value"]
  flags         = 10
  name          = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_keychain` - Authentication key-chain name.
* `auth_mode` - Authentication mode. Valid values: `none`, `md5`, `text`.

* `auth_string` - Authentication string/password.
* `flags` - Flags.
* `name` - Interface name.
* `receive_version` - Receive version. Valid values: `1`, `2`.

* `send_version` - Send version. Valid values: `1`, `2`.

* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets. Valid values: `disable`, `enable`.

* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned`, `regular`.

* `split_horizon_status` - Enable/disable split horizon. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router RipInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_rip_interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

