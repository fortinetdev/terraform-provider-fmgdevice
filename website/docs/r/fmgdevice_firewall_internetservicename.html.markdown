---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservicename"
description: |-
  <i>This object will be purged after policy copy and install.</i> Define internet service names.
---

# fmgdevice_firewall_internetservicename
<i>This object will be purged after policy copy and install.</i> Define internet service names.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `city_id` - City ID.
* `country_id` - Country or Area ID.
* `internet_service_id` - Internet Service ID.
* `name` - Internet Service name.
* `region_id` - Region ID.
* `type` - Internet Service name type. Valid values: `default`, `location`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceName can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservicename.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

