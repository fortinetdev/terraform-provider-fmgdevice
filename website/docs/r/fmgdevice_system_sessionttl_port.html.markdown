---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sessionttl_port"
description: |-
  Session TTL port.
---

# fmgdevice_system_sessionttl_port
Session TTL port.

~> This resource is a sub resource for variable `port` of resource `fmgdevice_system_sessionttl`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_sessionttl_port" "trname" {
  end_port          = 10
  fosid             = 10
  protocol          = 10
  refresh_direction = "outgoing"
  start_port        = 10
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `end_port` - End port number.
* `fosid` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `refresh_direction` - Configure refresh direction. Valid values: `both`, `outgoing`, `incoming`.

* `start_port` - Start port number.
* `timeout` - Session timeout (TTL).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SessionTtlPort can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sessionttl_port.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

