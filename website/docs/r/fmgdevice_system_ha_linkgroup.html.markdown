---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ha_linkgroup"
description: |-
  Link group table.
---

# fmgdevice_system_ha_linkgroup
Link group table.

~> This resource is a sub resource for variable `link_group` of resource `fmgdevice_system_ha`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `member` - Member interface in this link group.
* `min_members` - Minimum number of members that must be up before this link group is considered up.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System HaLinkGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ha_linkgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

