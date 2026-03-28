---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_admin_trusthosts"
description: |-
  Device SystemAdminTrusthosts
---

# fmgdevice_system_admin_trusthosts
Device SystemAdminTrusthosts

~> This resource is a sub resource for variable `trusthosts` of resource `fmgdevice_system_admin`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `admin` - Admin.

* `fosid` - Id.
* `ipv4` - Ipv4.
* `ipv6` - Ipv6.
* `type` - Type. Valid values: `ipv4`, `ipv6`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AdminTrusthosts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "admin=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_admin_trusthosts.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

