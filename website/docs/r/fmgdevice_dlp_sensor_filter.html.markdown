---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_sensor_filter"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_dlp_sensor_filter
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `filter` of resource `fmgdevice_dlp_sensor`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `sensor` - Sensor.

* `action` - Action. Valid values: `log-only`, `block`, `quarantine-ip`, `allow`.

* `archive` - Archive. Valid values: `disable`, `enable`.

* `company_identifier` - Company-Identifier.
* `expiry` - Expiry.
* `file_size` - File-Size.
* `file_type` - File-Type.
* `filter_by` - Filter-By. Valid values: `credit-card`, `ssn`, `regexp`, `file-type`, `file-size`, `fingerprint`, `watermark`, `encrypted`, `file-type-and-size`.

* `fosid` - Id.
* `match_percentage` - Match-Percentage.
* `name` - Name.
* `proto` - Proto. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mapi`, `http-get`, `http-post`, `cifs`, `ssh`.

* `regexp` - Regexp.
* `sensitivity` - Sensitivity.
* `severity` - Severity. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Type. Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp SensorFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "sensor=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_sensor_filter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

