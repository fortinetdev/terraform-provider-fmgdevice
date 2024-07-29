---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ipmacbinding_setting"
description: |-
  Configure IP to MAC binding settings.
---

# fmgdevice_firewall_ipmacbinding_setting
Configure IP to MAC binding settings.

## Example Usage

```hcl
resource "fmgdevice_firewall_ipmacbinding_setting" "trname" {
  bindthroughfw = "enable"
  bindtofw      = "disable"
  undefinedhost = "block"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bindthroughfw` - Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `disable`, `enable`.

* `bindtofw` - Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `disable`, `enable`.

* `undefinedhost` - Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `block`, `allow`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall IpmacbindingSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ipmacbinding_setting.labelname FirewallIpmacbindingSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

