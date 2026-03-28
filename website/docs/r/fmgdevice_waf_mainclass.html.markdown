---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_mainclass"
description: |-
  <i>This object will be purged after policy copy and install.</i> Hidden table for datasource.
---

# fmgdevice_waf_mainclass
<i>This object will be purged after policy copy and install.</i> Hidden table for datasource.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Main signature class ID.
* `name` - Main signature class name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf MainClass can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_mainclass.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

