---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_service_custom"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure custom services.
---

# fmgdevice_firewall_service_custom
<i>This object will be purged after policy copy and install.</i> Configure custom services.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `app_category` - Application category ID.
* `app_service_type` - Application service type. Valid values: `disable`, `app-id`, `app-category`.

* `application` - Application ID.
* `category` - Service category.
* `check_reset_range` - Configure the type of ICMP error message verification. Valid values: `disable`, `default`, `strict`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fqdn` - Fully qualified domain name.
* `global_object` - Global-Object.
* `helper` - Helper name. Valid values: `disable`, `auto`, `ftp`, `tftp`, `ras`, `h323`, `tns`, `mms`, `sip`, `pptp`, `rtsp`, `dns-udp`, `dns-tcp`, `pmap`, `rsh`, `dcerpc`, `mgcp`, `gtp-c`, `gtp-u`, `gtp-b`, `pfcp`.

* `icmpcode` - ICMP code.
* `icmptype` - ICMP type.
* `iprange` - Start and end of the IP range associated with service.
* `name` - Custom service name.
* `protocol` - Protocol type based on IANA numbers. Valid values: `ICMP`, `IP`, `ICMP6`, `HTTP`, `FTP`, `CONNECT`, `ALL`, `SOCKS-TCP`, `SOCKS-UDP`, `TCP/UDP/UDP-Lite/SCTP`.

* `protocol_number` - IP protocol number.
* `proxy` - Enable/disable web proxy service. Valid values: `disable`, `enable`.

* `sctp_portrange` - Multiple SCTP port ranges.
* `session_ttl` - Session TTL (300 - 2764800, 0 = default).
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_portrange` - Multiple TCP port ranges.
* `tcp_rst_timer` - Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `udp_idle_timer` - Number of seconds before an idle UDP/UDP-Lite connection times out (0 - 86400 sec, 0 = default).
* `udp_portrange` - Multiple UDP port ranges.
* `visibility` - Enable/disable the visibility of the service on the GUI. Valid values: `disable`, `enable`.

* `udplite_portrange` - Multiple UDP-Lite port ranges.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ServiceCustom can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_service_custom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

