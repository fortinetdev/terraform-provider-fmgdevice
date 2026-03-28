---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_spanport"
description: |-
  Device SystemSpanPort
---

# fmgdevice_system_spanport
Device SystemSpanPort

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Id.
* `span_dest_port` - Span-Dest-Port.
* `span_direction` - Span-Direction. Valid values: `rx`, `tx`, `both`.

* `span_source_port` - Span-Source-Port.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SpanPort can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_spanport.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

