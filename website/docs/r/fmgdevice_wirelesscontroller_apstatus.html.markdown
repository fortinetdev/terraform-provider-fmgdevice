---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_apstatus"
description: |-
  Configure access point status (rogue | accepted | suppressed).
---

# fmgdevice_wirelesscontroller_apstatus
Configure access point status (rogue | accepted | suppressed).

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_apstatus" "trname" {
  bssid       = "your own value"
  fosid       = 10
  ssid        = "your own value"
  status      = "rogue"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bssid` - Access Point's (AP's) BSSID.
* `fosid` - AP ID.
* `ssid` - Access Point's (AP's) SSID.
* `status` - Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController ApStatus can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_apstatus.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

