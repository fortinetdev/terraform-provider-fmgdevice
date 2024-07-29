---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_fortiguard"
description: |-
  Configure FortiGuard Web Filter service.
---

# fmgdevice_webfilter_fortiguard
Configure FortiGuard Web Filter service.

## Example Usage

```hcl
resource "fmgdevice_webfilter_fortiguard" "trname" {
  cache_mem_percent  = 10
  cache_mem_permille = 10
  cache_mode         = "ttl"
  cache_prefix_match = "enable"
  close_ports        = "enable"
  device_name        = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15).
* `cache_mem_permille` - Maximum permille of available memory allocated to caching (1 - 150).
* `cache_mode` - Cache entry expiration mode. Valid values: `ttl`, `db-ver`.

* `cache_prefix_match` - Enable/disable prefix matching in the cache. Valid values: `disable`, `enable`.

* `close_ports` - Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `disable`, `enable`.

* `embed_image` - Enable/disable embedding images into replacement messages (default = enable). Valid values: `disable`, `enable`.

* `ovrd_auth_https` - Enable/disable use of HTTPS for override authentication. Valid values: `disable`, `enable`.

* `ovrd_auth_port_http` - Port to use for FortiGuard Web Filter HTTP override authentication.
* `ovrd_auth_port_https` - Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
* `ovrd_auth_port_https_flow` - Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
* `ovrd_auth_port_warning` - Port to use for FortiGuard Web Filter Warning override authentication.
* `request_packet_size_limit` - Limit size of URL request packets sent to FortiGuard server (0 for default).
* `warn_auth_https` - Enable/disable use of HTTPS for warning and authentication. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter Fortiguard can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_fortiguard.labelname WebfilterFortiguard
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

