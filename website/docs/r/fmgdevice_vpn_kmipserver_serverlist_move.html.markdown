---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_kmipserver_serverlist_move"
description: |-
  KMIP server list.
---

# fmgdevice_vpn_kmipserver_serverlist_move
KMIP server list.

## Example Usage

```hcl
resource "fmgdevice_vpn_kmipserver_serverlist_move" "trname" {
  target      = 11
  option      = "after"
  server_list = 12
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  kmip_server = fmgdevice_vpn_kmipserver.trname1.name
  depends_on  = [fmgdevice_vpn_kmipserver_serverlist.trname2, fmgdevice_vpn_kmipserver_serverlist.trname1]
}

resource "fmgdevice_vpn_kmipserver_serverlist" "trname2" {
  fosid       = 12
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  kmip_server = fmgdevice_vpn_kmipserver.trname1.name
  depends_on  = [fmgdevice_vpn_kmipserver.trname1]
}

resource "fmgdevice_vpn_kmipserver_serverlist" "trname1" {
  fosid       = 11
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  kmip_server = fmgdevice_vpn_kmipserver.trname1.name
  depends_on  = [fmgdevice_vpn_kmipserver.trname1]
}

resource "fmgdevice_vpn_kmipserver" "trname1" {
  name        = "test"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `kmip_server` - Kmip Server.
* `server_list` - Server List.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the server list changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of server lists.
