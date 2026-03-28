---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_group"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure firewall application groups.
---

# fmgdevice_application_group
<i>This object will be purged after policy copy and install.</i> Configure firewall application groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `application` - Application ID list.
* `behavior` - Application behavior filter.
* `category` - Application category ID list.
* `comment` - Comments.
* `name` - Application group name.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.

* `protocols` - Application protocol filter.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
* `technology` - Application technology filter.
* `type` - Application group type. Valid values: `application`, `filter`.

* `vendor` - Application vendor filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Group can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

