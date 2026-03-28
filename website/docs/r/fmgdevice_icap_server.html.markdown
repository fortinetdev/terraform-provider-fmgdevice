---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_server"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure ICAP servers.
---

# fmgdevice_icap_server
<i>This object will be purged after policy copy and install.</i> Configure ICAP servers.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `addr_type` - Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `fqdn`, `ip4`, `ip6`.

* `fqdn` - ICAP remote server Fully Qualified Domain Name (FQDN).
* `healthcheck` - Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.

* `healthcheck_service` - ICAP Service name to use for health checks.
* `ip_address` - IPv4 address of the ICAP server.
* `ip_version` - IP version. Valid values: `4`, `6`.

* `ip6_address` - IPv6 address of the ICAP server.
* `max_connections` - Maximum number of concurrent connections to ICAP server (unlimited = 0, default = 100). Must not be less than wad-worker-count.
* `name` - Server name.
* `port` - ICAP server port.
* `secure` - Enable/disable secure connection to ICAP server. Valid values: `disable`, `enable`.

* `ssl_cert` - CA certificate name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap Server can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_server.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

