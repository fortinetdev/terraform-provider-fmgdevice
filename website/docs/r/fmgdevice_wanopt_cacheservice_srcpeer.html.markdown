---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_cacheservice_srcpeer"
description: |-
  Modify cache-service source peer list.
---

# fmgdevice_wanopt_cacheservice_srcpeer
Modify cache-service source peer list.

~> This resource is a sub resource for variable `src_peer` of resource `fmgdevice_wanopt_cacheservice`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wanopt_cacheservice_srcpeer" "trname" {
  auth_type   = 10
  device_id   = "your own value"
  encode_type = 10
  ip          = "your own value"
  priority    = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{device_id}}.

## Import

Wanopt CacheServiceSrcPeer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_cacheservice_srcpeer.labelname {{device_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

