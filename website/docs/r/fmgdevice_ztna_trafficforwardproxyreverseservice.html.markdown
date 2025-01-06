---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_trafficforwardproxyreverseservice"
description: |-
  Configure ZTNA traffic forward proxy reverse service.
---

# fmgdevice_ztna_trafficforwardproxyreverseservice
Configure ZTNA traffic forward proxy reverse service.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `remote_servers`: `fmgdevice_ztna_trafficforwardproxyreverseservice_remoteservers`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `remote_servers` - Remote-Servers. The structure of `remote_servers` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `remote_servers` block supports:

* `address` - Server adress(IP or FQDN).
* `certificate` - The name of the certificate to use for SSL handshake.
* `health_check_interval` - Health check interval in seconds (0 - 600, default = 60, 0 = disable).
* `name` - Remote server name
* `port` - Port number that traffic uses to connect to remote server (0 - 65535;).
* `ssl_max_version` - Highest TLS version acceptable from a server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `status` - Remote server status. Valid values: `disable`, `enable`.

* `trusted_server_ca` - Trusted Server CA certificate used by SSL connection.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna TrafficForwardProxyReverseService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_trafficforwardproxyreverseservice.labelname ZtnaTrafficForwardProxyReverseService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

