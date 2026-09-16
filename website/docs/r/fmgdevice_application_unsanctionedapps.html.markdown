---
subcategory: "Application"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_application_unsanctionedapps"
description: |-
  Configure unsanctioned applications.
---

# fmgdevice_application_unsanctionedapps
Configure unsanctioned applications.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `app` - Application ID.
* `category` - Application category.
* `fosid` - Entry ID.
* `status` - Status. Valid values: `unsanctioned`, `sanctioned`.

* `type` - Type. Valid values: `app`, `category`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Application UnsanctionedApps can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_application_unsanctionedapps.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

