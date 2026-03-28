---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_customlanguage"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure custom languages.
---

# fmgdevice_system_customlanguage
<i>This object will be purged after policy copy and install.</i> Configure custom languages.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `comments` - Comment.
* `filename` - Custom language file path.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CustomLanguage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_customlanguage.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

