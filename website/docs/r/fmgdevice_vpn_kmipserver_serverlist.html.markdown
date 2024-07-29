---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_kmipserver_serverlist"
description: |-
  KMIP server list.
---

# fmgdevice_vpn_kmipserver_serverlist
KMIP server list.

~> This resource is a sub resource for variable `server_list` of resource `fmgdevice_vpn_kmipserver`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_kmipserver_serverlist" "trname" {
  cert        = ["your own value"]
  fosid       = 10
  port        = 10
  server      = "your own value"
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `kmip_server` - Kmip Server.

* `cert` - Client certificate to use for connectivity to the KMIP server.
* `fosid` - ID
* `port` - KMIP server port.
* `server` - KMIP server FQDN or IP address.
* `status` - Enable/disable KMIP server. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Vpn KmipServerServerList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "kmip_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_kmipserver_serverlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

