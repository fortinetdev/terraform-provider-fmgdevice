---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_name"
description: |-
  Configure application signatures.
---

# fmgdevice_application_name
Configure application signatures.

## Example Usage

```hcl
resource "fmgdevice_application_name" "trname" {
  behavior       = "your own value"
  category       = ["your own value"]
  fosid          = 10
  max_scan_range = 10
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `behavior` - Behavior.
* `category` - Category.
* `fosid` - Id.
* `max_scan_range` - Max-Scan-Range.
* `name` - Application name.
* `parameter` - Parameter.
* `popularity` - Popularity.
* `protocol` - Protocol.
* `risk` - Risk.
* `sub_category` - Sub-Category.
* `technology` - Technology.
* `vendor` - Vendor.
* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Name can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_name.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

