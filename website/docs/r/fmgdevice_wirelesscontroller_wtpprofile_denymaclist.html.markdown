---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpprofile_denymaclist"
description: |-
  List of MAC addresses that are denied access to this WTP, FortiAP, or AP.
---

# fmgdevice_wirelesscontroller_wtpprofile_denymaclist
List of MAC addresses that are denied access to this WTP, FortiAP, or AP.

~> This resource is a sub resource for variable `deny_mac_list` of resource `fmgdevice_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtpprofile_denymaclist" "trname" {
  fosid       = 10
  mac         = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp_profile` - Wtp Profile.

* `fosid` - ID.
* `mac` - A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController WtpProfileDenyMacList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpprofile_denymaclist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

