---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_cacheservice"
description: |-
  Designate cache-service for wan-optimization and webcache.
---

# fmgdevice_wanopt_cacheservice
Designate cache-service for wan-optimization and webcache.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dst_peer`: `fmgdevice_wanopt_cacheservice_dstpeer`
>- `src_peer`: `fmgdevice_wanopt_cacheservice_srcpeer`



## Example Usage

```hcl
resource "fmgdevice_wanopt_cacheservice" "trname" {
  acceptable_connections = "peers"
  collaboration          = "enable"
  device_id              = "your own value"
  dst_peer {
    auth_type   = 10
    device_id   = "your own value"
    encode_type = 10
    ip          = "your own value"
    priority    = 10
  }

  prefer_scenario = "prefer-speed"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `acceptable_connections` - Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.

* `collaboration` - Enable/disable cache-collaboration between cache-service clusters. Valid values: `disable`, `enable`.

* `device_id` - Set identifier for this cache device.
* `dst_peer` - Dst-Peer. The structure of `dst_peer` block is documented below.
* `prefer_scenario` - Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.

* `src_peer` - Src-Peer. The structure of `src_peer` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dst_peer` block supports:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.

The `src_peer` block supports:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt CacheService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_cacheservice.labelname WanoptCacheService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

