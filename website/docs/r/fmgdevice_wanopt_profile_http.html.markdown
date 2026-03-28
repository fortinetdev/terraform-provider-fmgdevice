---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_profile_http"
description: |-
  <i>This object will be purged after policy copy and install.</i> Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features.
---

# fmgdevice_wanopt_profile_http
<i>This object will be purged after policy copy and install.</i> Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features.

~> This resource is a sub resource for variable `http` of resource `fmgdevice_wanopt_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt ProfileHttp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_profile_http.labelname WanoptProfileHttp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

