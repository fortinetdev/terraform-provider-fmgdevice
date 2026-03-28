---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_pacpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_webproxy_pacpolicy
<i>This object will be purged after policy copy and install.</i>

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comments.
* `dstaddr` - Dstaddr.
* `pac_file_data` - Pac-File-Data.
* `pac_file_name` - Pac-File-Name.
* `policyid` - Policyid.
* `srcaddr` - Srcaddr.
* `srcaddr6` - Srcaddr6.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

WebProxy PacPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_pacpolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

