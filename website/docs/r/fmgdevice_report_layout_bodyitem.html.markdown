---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout_bodyitem"
description: |-
  Configure report body item.
---

# fmgdevice_report_layout_bodyitem
Configure report body item.

~> This resource is a sub resource for variable `body_item` of resource `fmgdevice_report_layout`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `list`: `fmgdevice_report_layout_bodyitem_list`
>- `parameters`: `fmgdevice_report_layout_bodyitem_parameters`



## Example Usage

```hcl
resource "fmgdevice_report_layout_bodyitem" "trname" {
  chart         = ["your own value"]
  chart_options = ["include-no-data"]
  column        = 10
  content       = "your own value"
  description   = "your own value"
  fosid         = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `layout` - Layout.

* `chart` - Report item chart name.
* `chart_options` - Report chart options. Valid values: `include-no-data`, `hide-title`, `show-caption`.

* `column` - Report section column number.
* `content` - Report item text content.
* `description` - Description.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `hide` - Enable/disable hide item in report. Valid values: `disable`, `enable`.

* `fosid` - Report item ID.
* `img_src` - Report item image file name.
* `list` - List. The structure of `list` block is documented below.
* `list_component` - Report item list component. Valid values: `bullet`, `numbered`.

* `misc_component` - Report item miscellaneous component. Valid values: `hline`, `page-break`, `column-break`, `section-start`.

* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `style` - Report item style.
* `table_caption_style` - Table chart caption style.
* `table_column_widths` - Report item table column widths.
* `table_even_row_style` - Table chart even row style.
* `table_head_style` - Table chart head style.
* `table_odd_row_style` - Table chart odd row style.
* `text_component` - Report item text component. Valid values: `text`, `heading1`, `heading2`, `heading3`.

* `title` - Report section title.
* `top_n` - Value of top.
* `type` - Report item type. Valid values: `text`, `image`, `chart`, `misc`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `list` block supports:

* `content` - List entry content.
* `id` - List entry ID.

The `parameters` block supports:

* `id` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Report LayoutBodyItem can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "layout=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_layout_bodyitem.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

