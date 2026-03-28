---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_tacacs"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure TACACS+ server entries.
---

# fmgdevice_user_tacacs
<i>This object will be purged after policy copy and install.</i> Configure TACACS+ server entries.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_user_tacacs_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `authen_type` - Allowed authentication protocols/methods. Valid values: `auto`, `ascii`, `pap`, `chap`, `mschap`.

* `authorization` - Enable/disable TACACS+ authorization. Valid values: `disable`, `enable`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key to access the primary server.
* `name` - TACACS+ server entry name.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - Source IP address for communications to TACACS+ server.
* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `authen_type` - Allowed authentication protocols/methods. Valid values: `auto`, `ascii`, `pap`, `chap`, `mschap`.

* `authorization` - Enable/disable TACACS+ authorization. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `key` - Key to access the primary server.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - Source IP address for communications to TACACS+ server.
* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
* `vrf_select` - VRF ID used for connection to server.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Tacacs can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_tacacs.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

