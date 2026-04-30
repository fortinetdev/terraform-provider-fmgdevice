---
subcategory: "Wireless Controller"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_lwprofile"
description: |-
  Configure LoRaWAN profile.
---

# fmgdevice_wirelesscontroller_lwprofile
Configure LoRaWAN profile.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `cups_api_key` - CUPS API key of LoRaWAN device.
* `cups_server` - CUPS (Configuration and Update Server) domain name or IP address of LoRaWAN device.
* `cups_server_port` - CUPS Port value of LoRaWAN device.
* `lw_protocol` - Configure LoRaWAN protocol (default = basics-station) Valid values: `basics-station`, `packet-forwarder`.

* `name` - LoRaWAN profile name.
* `tc_api_key` - TC API key of LoRaWAN device.
* `tc_server` - TC (Traffic Controller) domain name or IP address of LoRaWAN device.
* `tc_server_port` - TC Port value of LoRaWAN device.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController LwProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_lwprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

