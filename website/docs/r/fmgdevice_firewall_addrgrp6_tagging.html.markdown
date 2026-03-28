---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_addrgrp6_tagging"
description: |-
  <i>This object will be purged after policy copy and install.</i> Config object tagging.
---

# fmgdevice_firewall_addrgrp6_tagging
<i>This object will be purged after policy copy and install.</i> Config object tagging.

~> This resource is a sub resource for variable `tagging` of resource `fmgdevice_firewall_addrgrp6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `addrgrp6` - Addrgrp6.

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Addrgrp6Tagging can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "addrgrp6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_addrgrp6_tagging.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

