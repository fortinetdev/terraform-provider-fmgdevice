---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtp_lan"
description: |-
  WTP LAN port mapping.
---

# fmgdevice_wirelesscontroller_wtp_lan
WTP LAN port mapping.

~> This resource is a sub resource for variable `lan` of resource `fmgdevice_wirelesscontroller_wtp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtp_lan" "trname" {
  port_esl_mode = "offline"
  port_esl_ssid = ["your own value"]
  port_mode     = "bridge-to-wan"
  port_ssid     = ["your own value"]
  port1_mode    = "bridge-to-ssid"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp` - Wtp.

* `port_esl_mode` - ESL port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_esl_ssid` - Bridge ESL port to SSID.
* `port_mode` - LAN port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port8_ssid` - Bridge LAN port 8 to SSID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController WtpLan can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtp_lan.labelname WirelessControllerWtpLan
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

