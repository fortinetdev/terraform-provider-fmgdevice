---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_remoteserver"
description: |-
  Device IcapRemoteServer
---

# fmgdevice_icap_remoteserver
Device IcapRemoteServer

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `addr_type` - Addr-Type. Valid values: `fqdn`, `ip4`, `ip6`.

* `fqdn` - Fqdn.
* `healthcheck` - Healthcheck. Valid values: `disable`, `enable`.

* `healthcheck_service` - Healthcheck-Service.
* `ip_address` - Ip-Address.
* `ip6_address` - Ip6-Address.
* `max_connections` - Max-Connections.
* `name` - Name.
* `port` - Port.
* `secure` - Secure. Valid values: `disable`, `enable`.

* `ssl_cert` - Ssl-Cert.
* `validate_server_certificate` - Validate-Server-Certificate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap RemoteServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_remoteserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

