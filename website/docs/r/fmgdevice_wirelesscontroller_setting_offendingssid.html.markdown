---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_setting_offendingssid"
description: |-
  Configure offending SSID.
---

# fmgdevice_wirelesscontroller_setting_offendingssid
Configure offending SSID.

~> This resource is a sub resource for variable `offending_ssid` of resource `fmgdevice_wirelesscontroller_setting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_setting_offendingssid" "trname" {
  action       = ["log"]
  fosid        = 10
  ssid_pattern = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Actions taken for detected offending SSID. Valid values: `log`, `suppress`.

* `fosid` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive). For example, word, word*, *word, wo*rd.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController SettingOffendingSsid can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_setting_offendingssid.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

