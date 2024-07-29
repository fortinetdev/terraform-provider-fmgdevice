---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_iptranslation"
description: |-
  Configure firewall IP-translation.
---

# fmgdevice_firewall_iptranslation
Configure firewall IP-translation.

## Example Usage

```hcl
resource "fmgdevice_firewall_iptranslation" "trname" {
  endip       = "your own value"
  map_startip = "your own value"
  startip     = "your own value"
  transid     = 10
  type        = "SCTP"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `endip` - Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `map_startip` - Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `transid` - IP translation ID.
* `type` - IP translation type (option: SCTP). Valid values: `SCTP`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{transid}}.

## Import

Firewall IpTranslation can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_iptranslation.labelname {{transid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

