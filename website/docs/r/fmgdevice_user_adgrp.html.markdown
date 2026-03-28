---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_adgrp"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure FSSO groups.
---

# fmgdevice_user_adgrp
<i>This object will be purged after policy copy and install.</i> Configure FSSO groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `connector_source` - FSSO connector source.
* `fosid` - Id.
* `name` - Name.
* `server_name` - FSSO agent name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Adgrp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_adgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

