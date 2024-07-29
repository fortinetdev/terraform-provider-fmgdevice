---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout"
description: |-
  Report layout configuration.
---

# fmgdevice_report_layout
Report layout configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `body_item`: `fmgdevice_report_layout_bodyitem`
>- `page`: `fmgdevice_report_layout_page`



## Example Usage

```hcl
resource "fmgdevice_report_layout" "trname" {
  body_item {
    chart            = ["your own value"]
    chart_options    = ["hide-title"]
    column           = 10
    content          = "your own value"
    description      = "your own value"
    drill_down_items = "your own value"
    drill_down_types = "your own value"
    hide             = "enable"
    fosid            = 10
    img_src          = "your own value"
    list {
      content = "your own value"
      fosid   = 10
    }

    list_component = "numbered"
    misc_component = "column-break"
    parameters {
      fosid = 10
      name  = "your own value"
      value = "your own value"
    }

    style                = ["your own value"]
    table_caption_style  = ["your own value"]
    table_column_widths  = "your own value"
    table_even_row_style = ["your own value"]
    table_head_style     = ["your own value"]
    table_odd_row_style  = ["your own value"]
    text_component       = "heading2"
    title                = "your own value"
    top_n                = 10
    type                 = "text"
  }

  cutoff_option = "custom"
  cutoff_time   = "your own value"
  day           = "sunday"
  description   = "your own value"
  name          = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `body_item` - Body-Item. The structure of `body_item` block is documented below.
* `cutoff_option` - Cutoff-option is either run-time or custom. Valid values: `run-time`, `custom`.

* `cutoff_time` - Custom cutoff time to generate report (format = hh:mm).
* `day` - Schedule days of week to generate report. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `description` - Description.
* `email_recipients` - Email recipients for generated reports.
* `email_send` - Enable/disable sending emails after reports are generated. Valid values: `disable`, `enable`.

* `format` - Report format. Valid values: `html`, `pdf`.

* `max_pdf_report` - Maximum number of PDF reports to keep at one time (oldest report is overwritten).
* `name` - Report layout name.
* `options` - Report layout options. Valid values: `include-table-of-content`, `auto-numbering-heading`, `view-chart-as-heading`, `show-html-navbar-before-heading`, `dummy-option`.

* `page` - Page. The structure of `page` block is documented below.
* `schedule_type` - Report schedule type. Valid values: `daily`, `weekly`, `once`, `demand`.

* `style_theme` - Report style theme.
* `subtitle` - Report subtitle.
* `time` - Schedule time to generate report (format = hh:mm).
* `title` - Report title.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `body_item` block supports:

* `chart` - Report item chart name.
* `chart_options` - Report chart options. Valid values: `include-no-data`, `hide-title`, `show-caption`.

* `column` - Report section column number.
* `content` - Report item text content.
* `description` - Description.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `hide` - Enable/disable hide item in report. Valid values: `disable`, `enable`.

* `id` - Report item ID.
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


The `list` block supports:

* `content` - List entry content.
* `id` - List entry ID.

The `parameters` block supports:

* `id` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.

The `page` block supports:

* `column_break_before` - Report page auto column break before heading. Valid values: `heading1`, `heading2`, `heading3`.

* `footer` - Footer. The structure of `footer` block is documented below.
* `header` - Header. The structure of `header` block is documented below.
* `options` - Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.

* `page_break_before` - Report page auto page break before heading. Valid values: `heading1`, `heading2`, `heading3`.

* `paper` - Report page paper. Valid values: `a4`, `letter`.


The `footer` block supports:

* `footer_item` - Footer-Item. The structure of `footer_item` block is documented below.
* `style` - Report footer style.

The `footer_item` block supports:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type. Valid values: `text`, `image`.


The `header` block supports:

* `header_item` - Header-Item. The structure of `header_item` block is documented below.
* `style` - Report header style.

The `header_item` block supports:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type. Valid values: `text`, `image`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Layout can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_layout.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

