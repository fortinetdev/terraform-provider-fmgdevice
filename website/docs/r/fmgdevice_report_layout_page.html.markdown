---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout_page"
description: |-
  Configure report page.
---

# fmgdevice_report_layout_page
Configure report page.

~> This resource is a sub resource for variable `page` of resource `fmgdevice_report_layout`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `footer`: `fmgdevice_report_layout_page_footer`
>- `header`: `fmgdevice_report_layout_page_header`



## Example Usage

```hcl
resource "fmgdevice_report_layout_page" "trname" {
  column_break_before = ["heading2"]
  footer {
    footer_item {
      content     = "your own value"
      description = "your own value"
      fosid       = 10
      img_src     = "your own value"
      style       = ["your own value"]
      type        = "text"
    }

    style = ["your own value"]
  }

  header {
    header_item {
      content     = "your own value"
      description = "your own value"
      fosid       = 10
      img_src     = "your own value"
      style       = ["your own value"]
      type        = "text"
    }

    style = ["your own value"]
  }

  options           = ["footer-on-first-page"]
  page_break_before = ["heading2"]
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `layout` - Layout.

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
* `id` - an identifier for the resource.

## Import

Report LayoutPage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "layout=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_layout_page.labelname ReportLayoutPage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

