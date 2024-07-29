---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_consoleserver_entries"
description: |-
  Entry used by console server.
---

# fmgdevice_system_consoleserver_entries
Entry used by console server.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_system_consoleserver`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_consoleserver_entries" "trname" {
  port        = 10
  slot_id     = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `port` - Port.
* `slot_id` - Slot ID of the FPC to connect.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{slot_id}}.

## Import

System ConsoleServerEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_consoleserver_entries.labelname {{slot_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

