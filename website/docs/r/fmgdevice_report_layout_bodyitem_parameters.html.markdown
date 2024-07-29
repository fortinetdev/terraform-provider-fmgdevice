---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout_bodyitem_parameters"
description: |-
  Parameters.
---

# fmgdevice_report_layout_bodyitem_parameters
Parameters.

~> This resource is a sub resource for variable `parameters` of resource `fmgdevice_report_layout_bodyitem`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_report_layout_bodyitem_parameters" "trname" {
  fosid       = 10
  name        = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `layout` - Layout.
* `body_item` - Body Item.

* `fosid` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Report LayoutBodyItemParameters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "layout=YOUR_VALUE", "body_item=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_layout_bodyitem_parameters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

