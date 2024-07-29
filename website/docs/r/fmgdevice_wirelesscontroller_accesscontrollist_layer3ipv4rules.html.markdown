---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules"
description: |-
  AP ACL layer3 ipv4 rule list.
---

# fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules
AP ACL layer3 ipv4 rule list.

~> This resource is a sub resource for variable `layer3_ipv4_rules` of resource `fmgdevice_wirelesscontroller_accesscontrollist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules" "trname" {
  action      = "allow"
  comment     = "your own value"
  dstaddr     = "your own value"
  dstport     = 10
  protocol    = 10
  rule_id     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `access_control_list` - Access Control List.

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_id}}.

## Import

WirelessController AccessControlListLayer3Ipv4Rules can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "access_control_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules.labelname {{rule_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

