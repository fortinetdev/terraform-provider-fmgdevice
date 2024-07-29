---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ftpproxy_explicit"
description: |-
  Configure explicit FTP proxy settings.
---

# fmgdevice_ftpproxy_explicit
Configure explicit FTP proxy settings.

## Example Usage

```hcl
resource "fmgdevice_ftpproxy_explicit" "trname" {
  incoming_ip        = "your own value"
  incoming_port      = ["your own value"]
  outgoing_ip        = ["your own value"]
  sec_default_action = "accept"
  server_data_mode   = "client"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `incoming_ip` - Accept incoming FTP requests from this IP address. An interface must have this IP address.
* `incoming_port` - Accept incoming FTP requests on one or more ports.
* `outgoing_ip` - Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
* `sec_default_action` - Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `deny`, `accept`.

* `server_data_mode` - Determine mode of data session on FTP server side. Valid values: `client`, `passive`.

* `ssl` - Enable/disable the explicit FTPS proxy. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `low`, `medium`.

* `ssl_cert` - List of certificate names to use for SSL connections to this server.
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.

* `status` - Enable/disable the explicit FTP proxy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

FtpProxy Explicit can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ftpproxy_explicit.labelname FtpProxyExplicit
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

