---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_isis_isisinterface"
description: |-
  IS-IS interface configuration.
---

# fmgdevice_router_isis_isisinterface
IS-IS interface configuration.

~> This resource is a sub resource for variable `isis_interface` of resource `fmgdevice_router_isis`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_isis_isisinterface" "trname" {
  auth_keychain_l1 = ["your own value"]
  auth_keychain_l2 = ["your own value"]
  auth_mode_l1     = "md5"
  auth_mode_l2     = "md5"
  auth_password_l1 = ["your own value"]
  name             = ["your own value"]
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5`, `password`.

* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5`, `password`.

* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs. Valid values: `disable`, `enable`.

* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs. Valid values: `disable`, `enable`.

* `circuit_type` - IS-IS interface's circuit type. Valid values: `level-1-2`, `level-1`, `level-2`.

* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets. Valid values: `disable`, `enable`.

* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `mesh_group` - Enable/disable IS-IS mesh group. Valid values: `disable`, `enable`.

* `mesh_group_id` - Mesh group ID &lt;0-4294967295&gt;, 0: mesh-group blocked.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `name` - IS-IS interface name.
* `network_type` - IS-IS interface's network type. Valid values: `broadcast`, `point-to-point`, `loopback`.

* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `status` - Enable/disable interface for IS-IS. Valid values: `disable`, `enable`.

* `status6` - Enable/disable IPv6 interface for IS-IS. Valid values: `disable`, `enable`.

* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router IsisIsisInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_isis_isisinterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

