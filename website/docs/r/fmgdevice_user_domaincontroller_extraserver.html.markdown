---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_domaincontroller_extraserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> Extra servers.
---

# fmgdevice_user_domaincontroller_extraserver
<i>This object will be purged after policy copy and install.</i> Extra servers.

~> This resource is a sub resource for variable `extra_server` of resource `fmgdevice_user_domaincontroller`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `domain_controller` - Domain Controller.

* `fosid` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User DomainControllerExtraServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "domain_controller=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_domaincontroller_extraserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

