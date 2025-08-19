---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_webproxy_apigateway6_realservers"
description: |-
  Select the real servers that this Access Proxy will distribute traffic to.
---

# fmgdevice_ztna_webproxy_apigateway6_realservers
Select the real servers that this Access Proxy will distribute traffic to.

~> This resource is a sub resource for variable `realservers` of resource `fmgdevice_ztna_webproxy_apigateway6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `web_proxy` - Web Proxy.
* `api_gateway6` - Api Gateway6.

* `addr_type` - Type of address. Valid values: `fqdn`, `ip`.

* `address` - Address or address group of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.

* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.

* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `disable`, `enable`.

* `http_host` - HTTP server domain name in HTTP header.
* `fosid` - Real server ID.
* `ip` - IPv6 address of the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ztna WebProxyApiGateway6Realservers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "web_proxy=YOUR_VALUE", "api_gateway6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_webproxy_apigateway6_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

