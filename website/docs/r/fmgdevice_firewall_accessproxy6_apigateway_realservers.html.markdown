---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_accessproxy6_apigateway_realservers"
description: |-
  <i>This object will be purged after policy copy and install.</i> Select the real servers that this Access Proxy will distribute traffic to.
---

# fmgdevice_firewall_accessproxy6_apigateway_realservers
<i>This object will be purged after policy copy and install.</i> Select the real servers that this Access Proxy will distribute traffic to.

~> This resource is a sub resource for variable `realservers` of resource `fmgdevice_firewall_accessproxy6_apigateway`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `access_proxy6` - Access Proxy6.
* `api_gateway` - Api Gateway.

* `addr_type` - Type of address. Valid values: `fqdn`, `ip`.

* `address` - Address or address group of the real server.
* `domain` - Wildcard domain name of the real server.
* `external_auth` - Enable/disable use of external browser as user-agent for SAML user authentication. Valid values: `disable`, `enable`.

* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.

* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.

* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `disable`, `enable`.

* `http_host` - HTTP server domain name in HTTP header.
* `fosid` - Real server ID.
* `ip` - IP address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key` - One or more server host key.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.

* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `tunnel_encryption` - Tunnel encryption. Valid values: `disable`, `enable`.

* `type` - TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.

* `verify_cert` - Enable/disable certificate verification of the real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall AccessProxy6ApiGatewayRealservers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "access_proxy6=YOUR_VALUE", "api_gateway=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_accessproxy6_apigateway_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

