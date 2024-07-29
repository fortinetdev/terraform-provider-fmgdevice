---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ipmacbinding_table"
description: |-
  Configure IP to MAC address pairs in the IP/MAC binding table.
---

# fmgdevice_firewall_ipmacbinding_table
Configure IP to MAC address pairs in the IP/MAC binding table.

## Example Usage

```hcl
resource "fmgdevice_firewall_ipmacbinding_table" "trname" {
  ip          = "your own value"
  mac         = "your own value"
  name        = "your own value"
  seq_num     = 10
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ip` - IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
* `mac` - MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).
* `name` - Name of the pair (optional, default = no name).
* `seq_num` - Entry number.
* `status` - Enable/disable this IP-mac binding pair. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Firewall IpmacbindingTable can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ipmacbinding_table.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

