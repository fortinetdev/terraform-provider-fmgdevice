---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fmgdevice_emailfilter_fortishield
Configure FortiGuard - AntiSpam.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text. Valid values: `disable`, `enable`.

* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter Fortishield can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_fortishield.labelname EmailfilterFortishield
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

