---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_guidisplay"
description: |-
  Configure how log messages are displayed on the GUI.
---

# fmgdevice_log_guidisplay
Configure how log messages are displayed on the GUI.

## Example Usage

```hcl
resource "fmgdevice_log_guidisplay" "trname" {
  fortiview_unscanned_apps = "enable"
  resolve_apps             = "enable"
  resolve_hosts            = "enable"
  device_name              = var.device_name # not required if setting is at provider
  device_vdom              = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fortiview_unscanned_apps` - Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `disable`, `enable`.

* `resolve_apps` - Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `disable`, `enable`.

* `resolve_hosts` - Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log GuiDisplay can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_guidisplay.labelname LogGuiDisplay
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

