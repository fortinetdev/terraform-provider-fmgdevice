---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_layout_bodyitem_move"
description: |-
  Move report body item.
---

# fmgdevice_report_layout_bodyitem_move
Move report body item.

## Example Usage

```hcl
resource "fmgdevice_report_layout_bodyitem_move" "trname" {
  target      = 13
  body_item   = 14
  option      = "before"
  layout      = fmgdevice_report_layout.trname1.name
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_report_layout_bodyitem.trname1, fmgdevice_report_layout_bodyitem.trname2]
}

resource "fmgdevice_report_layout_bodyitem" "trname2" {
  fosid       = 14
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_report_layout.trname1]
  layout      = fmgdevice_report_layout.trname1.name
}

resource "fmgdevice_report_layout_bodyitem" "trname1" {
  fosid       = 13
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_report_layout.trname1]
  layout      = fmgdevice_report_layout.trname1.name
}

resource "fmgdevice_report_layout" "trname1" {
  name        = "example"
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

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the body item changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of body items.
