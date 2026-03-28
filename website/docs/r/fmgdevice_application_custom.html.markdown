---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_custom"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure custom application signatures.
---

# fmgdevice_application_custom
<i>This object will be purged after policy copy and install.</i> Configure custom application signatures.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `behavior` - Custom application signature behavior.
* `category` - Custom application category ID (use ? to view available options).
* `comment` - Comment.
* `fosid` - Id.
* `name` - Name.
* `protocol` - Custom application signature protocol.
* `signature` - The text that makes up the actual custom application signature.
* `tag` - Signature tag.
* `technology` - Custom application signature technology.
* `vendor` - Custom application signature vendor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{tag}}.

## Import

Application Custom can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_custom.labelname {{tag}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

