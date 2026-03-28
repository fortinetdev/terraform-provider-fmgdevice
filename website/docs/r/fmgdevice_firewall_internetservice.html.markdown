---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservice"
description: |-
  <i>This object will be purged after policy copy and install.</i> Show Internet Service application.
---

# fmgdevice_firewall_internetservice
<i>This object will be purged after policy copy and install.</i> Show Internet Service application.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `database` - Database. Valid values: `isdb`, `irdb`.

* `direction` - Direction. Valid values: `src`, `dst`, `both`.

* `extra_ip_range_number` - Extra-Ip-Range-Number.
* `extra_ip6_range_number` - Extra-Ip6-Range-Number.
* `icon_id` - Icon-Id.
* `fosid` - Id.
* `ip_number` - Ip-Number.
* `ip_range_number` - Ip-Range-Number.
* `ip6_range_number` - Ip6-Range-Number.
* `name` - Name.
* `obsolete` - Obsolete.
* `singularity` - Singularity.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall InternetService can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservice.labelname FirewallInternetService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

