---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vip_realservers"
description: |-
  <i>This object will be purged after policy copy and install.</i> Select the real servers that this server load balancing VIP will distribute traffic to.
---

# fmgdevice_firewall_vip_realservers
<i>This object will be purged after policy copy and install.</i> Select the real servers that this server load balancing VIP will distribute traffic to.

~> This resource is a sub resource for variable `realservers` of resource `fmgdevice_firewall_vip`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `vip` - Vip.

* `address` - Dynamic address of the real server.
* `client_ip` - Only clients in this IP range can connect to this real server.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Time in seconds that the system waits before re-activating a previously down active server in the active-standby mode. This is to prevent any flapping issues.
* `http_host` - HTTP server domain name in HTTP header.
* `fosid` - Real server ID.
* `ip` - IP address of the real server.
* `max_connections` - Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `type` - Type of address. Valid values: `ip`, `address`.

* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `health_check_proto` - Health-Check-Proto. Valid values: `ping`, `http`.

* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall VipRealservers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "vip=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vip_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

