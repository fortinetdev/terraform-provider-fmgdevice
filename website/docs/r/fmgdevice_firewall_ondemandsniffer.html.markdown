---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ondemandsniffer"
description: |-
  Configure on-demand packet sniffer.
---

# fmgdevice_firewall_ondemandsniffer
Configure on-demand packet sniffer.

## Example Usage

```hcl
resource "fmgdevice_firewall_ondemandsniffer" "trname" {
  advanced_filter  = "your own value"
  hosts            = ["your own value"]
  interface        = ["port2"]
  max_packet_count = 10
  name             = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `advanced_filter` - Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
* `hosts` - IPv4 or IPv6 hosts to filter in this traffic sniffer.
* `interface` - Interface name that on-demand packet sniffer will take place.
* `max_packet_count` - Maximum number of packets to capture per on-demand packet sniffer.
* `name` - On-demand packet sniffer name.
* `non_ip_packet` - Include non-IP packets. Valid values: `disable`, `enable`.

* `ports` - Ports to filter for in this traffic sniffer.
* `protocols` - Protocols to filter in this traffic sniffer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall OnDemandSniffer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ondemandsniffer.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

