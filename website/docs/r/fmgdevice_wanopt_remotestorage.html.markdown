---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_remotestorage"
description: |-
  Configure a remote cache device as Web cache storage.
---

# fmgdevice_wanopt_remotestorage
Configure a remote cache device as Web cache storage.

## Example Usage

```hcl
resource "fmgdevice_wanopt_remoyour valueorage" "trname" {
  local_cache_id  = "your own value"
  remote_cache_id = "your own value"
  remote_cache_ip = "your own value"
  status          = "enable"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `local_cache_id` - ID that this device uses to connect to the remote device.
* `remote_cache_id` - ID of the remote device to which the device connects.
* `remote_cache_ip` - IP address of the remote device to which the device connects.
* `status` - Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt RemoteStorage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_remotestorage.labelname WanoptRemoteStorage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

