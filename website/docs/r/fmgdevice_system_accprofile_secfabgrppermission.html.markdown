---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_accprofile_secfabgrppermission"
description: |-
  Custom Security Fabric permissions.
---

# fmgdevice_system_accprofile_secfabgrppermission
Custom Security Fabric permissions.

~> This resource is a sub resource for variable `secfabgrp_permission` of resource `fmgdevice_system_accprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `accprofile` - Accprofile.

* `csffoo` - Fabric Overlay Orchestrator profiles and settings. Valid values: `none`, `read`, `read-write`.

* `csfsys` - Security Fabric system profiles and settings. Valid values: `none`, `read`, `read-write`.

* `mmsgtp` - UTM permission. Valid values: `none`, `read`, `read-write`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AccprofileSecfabgrpPermission can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "accprofile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_accprofile_secfabgrppermission.labelname SystemAccprofileSecfabgrpPermission
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

