---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout_page_header"
description: |-
  Configure report page header.
---

# fmgdevice_report_layout_page_header
Configure report page header.

~> This resource is a sub resource for variable `header` of resource `fmgdevice_report_layout_page`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_item`: `fmgdevice_report_layout_page_header_headeritem`



## Example Usage

```hcl
resource "fmgdevice_report_layout_page_header" "trname" {
  header_item {
    content     = "your own value"
    description = "your own value"
    fosid       = 10
    img_src     = "your own value"
    style       = ["your own value"]
    type        = "text"
  }

  style       = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `layout` - Layout.

* `header_item` - Header-Item. The structure of `header_item` block is documented below.
* `style` - Report header style.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_item` block supports:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type. Valid values: `text`, `image`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Report LayoutPageHeader can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "layout=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_layout_page_header.labelname ReportLayoutPageHeader
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

