---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_clientoptions"
description: |-
  DHCP client options.
---

# fmgdevice_system_interface_clientoptions
DHCP client options.

~> This resource is a sub resource for variable `client_options` of resource `fmgdevice_system_interface`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_clientoptions" "trname" {
  code        = 10
  fosid       = 10
  ip          = ["your own value"]
  type        = "string"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `code` - DHCP client option code.
* `fosid` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP client option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP client option value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System InterfaceClientOptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_clientoptions.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

