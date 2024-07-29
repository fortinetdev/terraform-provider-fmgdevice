---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sessionttl"
description: |-
  Configure global session TTL timers for this FortiGate.
---

# fmgdevice_system_sessionttl
Configure global session TTL timers for this FortiGate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `port`: `fmgdevice_system_sessionttl_port`



## Example Usage

```hcl
resource "fmgdevice_system_sessionttl" "trname" {
  default = "your own value"
  port {
    end_port          = 10
    fosid             = 10
    protocol          = 10
    refresh_direction = "outgoing"
    start_port        = 10
    timeout           = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default` - Default timeout.
* `port` - Port. The structure of `port` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `port` block supports:

* `end_port` - End port number.
* `id` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `refresh_direction` - Configure refresh direction. Valid values: `both`, `outgoing`, `incoming`.

* `start_port` - Start port number.
* `timeout` - Session timeout (TTL).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SessionTtl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sessionttl.labelname SystemSessionTtl
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

