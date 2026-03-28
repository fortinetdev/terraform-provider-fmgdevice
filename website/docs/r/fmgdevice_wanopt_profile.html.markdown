---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure WAN optimization profiles.
---

# fmgdevice_wanopt_profile
<i>This object will be purged after policy copy and install.</i> Configure WAN optimization profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cifs`: `fmgdevice_wanopt_profile_cifs`
>- `ftp`: `fmgdevice_wanopt_profile_ftp`
>- `http`: `fmgdevice_wanopt_profile_http`
>- `mapi`: `fmgdevice_wanopt_profile_mapi`
>- `tcp`: `fmgdevice_wanopt_profile_tcp`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_group` - Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
* `cifs` - Cifs. The structure of `cifs` block is documented below.
* `comments` - Comment.
* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `http` - Http. The structure of `http` block is documented below.
* `mapi` - Mapi. The structure of `mapi` block is documented below.
* `name` - Profile name.
* `tcp` - Tcp. The structure of `tcp` block is documented below.
* `transparent` - Enable/disable transparent mode. Valid values: `disable`, `enable`.


The `cifs` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `ftp` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `http` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic`, `fix`.

* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol`, `tcp`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `mapi` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.


The `tcp` block supports:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `disable`, `enable`.

* `byte_caching_opt` - Select whether TCP byte-caching uses system memory only or both memory and disk space. Valid values: `mem-only`, `mem-disk`.

* `log_traffic` - Enable/disable logging. Valid values: `disable`, `enable`.

* `port` - Port numbers or port number ranges for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `disable`, `enable`.

* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `disable`, `enable`.

* `ssl_port` - Port numbers or port number ranges on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable WAN Optimization. Valid values: `disable`, `enable`.

* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wanopt Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

