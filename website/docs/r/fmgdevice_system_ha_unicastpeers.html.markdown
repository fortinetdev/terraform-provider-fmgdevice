---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ha_unicastpeers"
description: |-
  Number of unicast peers.
---

# fmgdevice_system_ha_unicastpeers
Number of unicast peers.

~> This resource is a sub resource for variable `unicast_peers` of resource `fmgdevice_system_ha`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ha_unicastpeers" "trname" {
  fosid       = 10
  peer_ip     = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Table ID.
* `peer_ip` - Unicast peer IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaUnicastPeers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ha_unicastpeers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

