---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_externalresource"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure external resource.
---

# fmgdevice_system_externalresource
<i>This object will be purged after policy copy and install.</i> Configure external resource.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_system_externalresource_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `address_comment_field` - JSON Path to address description in generic address entry.
* `address_data_field` - JSON Path to address data in generic address entry.
* `address_name_field` - JSON Path to address name in generic address entry.
* `category` - User resource category.
* `client_cert` - Client certificate name.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - External resource name.
* `namespace` - Generic external connector address namespace.
* `object_array_path` - JSON Path to array of generic addresses in resource.
* `password` - HTTP basic authentication password.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `resource` - URL of external resource.
* `server_identity_check` - Certificate verification option. Valid values: `none`, `basic`, `full`.

* `source_ip` - Source IPv4 address used to communicate with server.
* `source_ip_interface` - IPv4 Source interface for communication with the server.
* `status` - Enable/disable user resource. Valid values: `disable`, `enable`.

* `type` - User resource type. Valid values: `category`, `address`, `domain`, `malware`, `mac-address`, `data`, `generic-address`.

* `update_method` - External resource update method. Valid values: `feed`, `push`.

* `user_agent` - HTTP User-Agent header (default = 'curl/7.58.0').
* `username` - HTTP basic authentication user name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vrf_select` - VRF ID used for connection to server.
* `proxy` - Proxy.
* `proxy_password` - Proxy-Password.
* `proxy_port` - Proxy-Port.
* `proxy_username` - Proxy-Username.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `source_ip` - Source-Ip.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ExternalResource can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_externalresource.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

