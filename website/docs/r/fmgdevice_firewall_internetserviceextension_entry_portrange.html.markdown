---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetserviceextension_entry_portrange"
description: |-
  Port ranges in the custom entry.
---

# fmgdevice_firewall_internetserviceextension_entry_portrange
Port ranges in the custom entry.

~> This resource is a sub resource for variable `port_range` of resource `fmgdevice_firewall_internetserviceextension_entry`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_firewall_internetserviceextension_entry_portrange" "trname" {
  end_port    = 10
  fosid       = 10
  start_port  = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `internet_service_extension` - Internet Service Extension.
* `entry` - Entry.

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `fosid` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceExtensionEntryPortRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "internet_service_extension=YOUR_VALUE", "entry=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetserviceextension_entry_portrange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

