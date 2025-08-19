---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_reverseconnector"
description: |-
  Configure ZTNA Reverse-Connector.
---

# fmgdevice_ztna_reverseconnector
Configure ZTNA Reverse-Connector.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `address` - Connector service edge adress(IP or FQDN).
* `certificate` - The name of the certificate to use for SSL handshake.
* `health_check_interval` - Health check interval in seconds (0 - 600, default = 60, 0 = disable).
* `name` - Reverse-Connector name
* `port` - Port number that traffic uses to connect to connector service edge(0 - 65535;).
* `ssl_max_version` - Highest TLS version acceptable from a server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Reverse-Connector status. Valid values: `disable`, `enable`.

* `trusted_server_ca` - Trusted Server CA certificate used by SSL connection.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna ReverseConnector can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_reverseconnector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

