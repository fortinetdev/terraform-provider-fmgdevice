---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationaction_formdata"
description: |-
  Form data parts for content type multipart/form-data.
---

# fmgdevice_system_automationaction_formdata
Form data parts for content type multipart/form-data.

~> This resource is a sub resource for variable `form_data` of resource `fmgdevice_system_automationaction`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `automation_action` - Automation Action.

* `fosid` - Entry ID.
* `key` - Key of the part of Multipart/form-data.
* `value` - Value of the part of Multipart/form-data.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AutomationActionFormData can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "automation_action=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationaction_formdata.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

