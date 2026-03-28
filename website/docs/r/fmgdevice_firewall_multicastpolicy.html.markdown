---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_multicastpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure multicast NAT policies.
---

# fmgdevice_firewall_multicastpolicy
<i>This object will be purged after policy copy and install.</i> Configure multicast NAT policies.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dnat` - IPv4 DNAT address used for multicast destination addresses.
* `dstaddr` - Destination address objects.
* `dstintf` - Destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `fosid` - Policy ID ((0 - 4294967294).
* `ips_sensor` - Name of an existing IPS sensor.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `all`, `utm`.

* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `snat` - Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT). Valid values: `disable`, `enable`.

* `snat_ip` - IPv4 address to be used as the source address for NATed traffic.
* `srcaddr` - Source address objects.
* `srcintf` - Source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the multicast policy.
* `utm_status` - Enable to add an IPS security profile to the policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall MulticastPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_multicastpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

