---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_list_defaultnetworkservices"
description: |-
  <i>This object will be purged after policy copy and install.</i> Default network service entries.
---

# fmgdevice_application_list_defaultnetworkservices
<i>This object will be purged after policy copy and install.</i> Default network service entries.

~> This resource is a sub resource for variable `default_network_services` of resource `fmgdevice_application_list`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `list` - List.

* `fosid` - Entry ID.
* `port` - Port number.
* `services` - Network protocols. Valid values: `http`, `ssh`, `telnet`, `ftp`, `dns`, `smtp`, `pop3`, `imap`, `snmp`, `nntp`, `https`.

* `violation_action` - Action for protocols not in the allowlist for selected port. Valid values: `block`, `monitor`, `pass`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Application ListDefaultNetworkServices can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_list_defaultnetworkservices.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

