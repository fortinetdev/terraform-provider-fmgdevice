---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_domaincontroller"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure domain controller entries.
---

# fmgdevice_user_domaincontroller
<i>This object will be purged after policy copy and install.</i> Configure domain controller entries.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `extra_server`: `fmgdevice_user_domaincontroller_extraserver`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ad_mode` - Set Active Directory mode. Valid values: `none`, `ds`, `lds`.

* `adlds_dn` - AD LDS distinguished name.
* `adlds_ip_address` - AD LDS IPv4 address.
* `adlds_ip6` - AD LDS IPv6 address.
* `adlds_port` - Port number of AD LDS service (default = 389).
* `change_detection` - Enable/disable detection of a configuration change in the Active Directory server. Valid values: `disable`, `enable`.

* `change_detection_period` - Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
* `dns_srv_lookup` - Enable/disable DNS service lookup. Valid values: `disable`, `enable`.

* `domain_name` - Domain DNS name.
* `extra_server` - Extra-Server. The structure of `extra_server` block is documented below.
* `hostname` - Hostname of the server to connect to.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip_address` - Domain controller IPv4 address.
* `ip6` - Domain controller IPv6 address.
* `ldap_server` - LDAP server name(s).
* `name` - Domain controller entry name.
* `password` - Password for specified username.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `replication_port` - Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_ip6` - FortiGate IPv6 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.
* `username` - User name to sign in with. Must have proper permissions for service.
* `domain_name_src` - Domain-Name-Src. Valid values: `server`, `client`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DomainController can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_domaincontroller.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

