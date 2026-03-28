---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_halic"
description: |-
  Device SystemHaLic
---

# fmgdevice_system_halic
Device SystemHaLic

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `last_sync_seat` - Last-Sync-Seat.
* `last_sync_seat_dlp` - Last-Sync-Seat-Dlp.
* `last_sync_seat_fcas` - Last-Sync-Seat-Fcas.
* `last_sync_seat_fnbi` - Last-Sync-Seat-Fnbi.
* `last_sync_time` - Last-Sync-Time.
* `serial_num` - Serial-Num.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial_num}}.

## Import

System HaLic can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_halic.labelname {{serial_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

