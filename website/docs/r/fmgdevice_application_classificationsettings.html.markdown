---
subcategory: "Application"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_classificationsettings"
description: |-
  Configure app classification setting.
---

# fmgdevice_application_classificationsettings
Configure app classification setting.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_app_classification` - Default application classification. Valid values: `unsanctioned`, `sanctioned`, `unclassified`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Application ClassificationSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_classificationsettings.labelname ApplicationClassificationSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

