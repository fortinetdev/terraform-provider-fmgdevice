---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_accesscontrollist"
description: |-
  Configure WiFi bridge access control list.
---

# fmgdevice_wirelesscontroller_accesscontrollist
Configure WiFi bridge access control list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `layer3_ipv4_rules`: `fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules`
>- `layer3_ipv6_rules`: `fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv6rules`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_accesscontrollist" "trname" {
  comment = "your own value"
  layer3_ipv4_rules {
    action   = "deny"
    comment  = "your own value"
    dstaddr  = "your own value"
    dstport  = 10
    protocol = 10
    rule_id  = 10
    srcaddr  = "your own value"
    srcport  = 10
  }

  layer3_ipv6_rules {
    action   = "deny"
    comment  = "your own value"
    dstaddr  = "your own value"
    dstport  = 10
    protocol = 10
    rule_id  = 10
    srcaddr  = "your own value"
    srcport  = 10
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Description.
* `layer3_ipv4_rules` - Layer3-Ipv4-Rules. The structure of `layer3_ipv4_rules` block is documented below.
* `layer3_ipv6_rules` - Layer3-Ipv6-Rules. The structure of `layer3_ipv6_rules` block is documented below.
* `name` - AP access control list name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `layer3_ipv4_rules` block supports:

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).

The `layer3_ipv6_rules` block supports:

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `srcport` - Source port (0 - 65535, default = 0, meaning any).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController AccessControlList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_accesscontrollist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

