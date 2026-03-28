---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_exchange"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure MS Exchange server entries.
---

# fmgdevice_user_exchange
<i>This object will be purged after policy copy and install.</i> Configure MS Exchange server entries.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_level` - Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.

* `auth_type` - Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.

* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses. Valid values: `disable`, `enable`.

* `connect_protocol` - Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.

* `domain_name` - MS Exchange server fully qualified domain name.
* `http_auth_type` - Authentication security type used for the HTTP transport. Valid values: `ntlm`, `basic`.

* `ip` - Server IPv4 address.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication.
* `name` - MS Exchange server entry name.
* `password` - Password for the specified username.
* `server_name` - MS Exchange server hostname.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `username` - User name used to sign in to the server. Must have proper permissions for service.
* `validate_server_certificate` - Enable/disable exchange server certificate validation. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Exchange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_exchange.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

