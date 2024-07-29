---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqpipaddresstype"
description: |-
  Configure IP address type availability.
---

# fmgdevice_wirelesscontroller_hotspot20_anqpipaddresstype
Configure IP address type availability.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqpipaddresstype" "trname" {
  ipv4_address_type = "public"
  ipv6_address_type = "not-known"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ipv4_address_type` - IPv4 address type. Valid values: `not-available`, `not-known`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`.

* `ipv6_address_type` - IPv6 address type. Valid values: `not-available`, `available`, `not-known`.

* `name` - IP type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20AnqpIpAddressType can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqpipaddresstype.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

