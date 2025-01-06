---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_trafficforwardproxyreverseservice_remoteservers"
description: |-
  Connector Remote server
---

# fmgdevice_ztna_trafficforwardproxyreverseservice_remoteservers
Connector Remote server

~> This resource is a sub resource for variable `remote_servers` of resource `fmgdevice_ztna_trafficforwardproxyreverseservice`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

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
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna TrafficForwardProxyReverseServiceRemoteServers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_trafficforwardproxyreverseservice_remoteservers.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

