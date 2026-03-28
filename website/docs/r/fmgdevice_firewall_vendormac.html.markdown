---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vendormac"
description: |-
  <i>This object will be purged after policy copy and install.</i> Show vendor and the MAC address they have.
---

# fmgdevice_firewall_vendormac
<i>This object will be purged after policy copy and install.</i> Show vendor and the MAC address they have.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Id.
* `mac_number` - Mac-Number.
* `name` - Name.
* `obsolete` - Obsolete.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall VendorMac can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vendormac.labelname FirewallVendorMac
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

