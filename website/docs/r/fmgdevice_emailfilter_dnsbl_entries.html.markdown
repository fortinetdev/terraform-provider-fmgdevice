---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_dnsbl_entries"
description: |-
  Spam filter DNSBL and ORBL server.
---

# fmgdevice_emailfilter_dnsbl_entries
Spam filter DNSBL and ORBL server.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_emailfilter_dnsbl`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `dnsbl` - Dnsbl.

* `action` - Reject connection or mark as spam email. Valid values: `spam`, `reject`.

* `fosid` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter DnsblEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "dnsbl=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_dnsbl_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

