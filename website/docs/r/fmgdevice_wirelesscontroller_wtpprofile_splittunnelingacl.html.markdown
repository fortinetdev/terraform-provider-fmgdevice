---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpprofile_splittunnelingacl"
description: |-
  Split tunneling ACL filter list.
---

# fmgdevice_wirelesscontroller_wtpprofile_splittunnelingacl
Split tunneling ACL filter list.

~> This resource is a sub resource for variable `split_tunneling_acl` of resource `fmgdevice_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp_profile` - Wtp Profile.

* `dest_ip` - Destination IP and mask for the split-tunneling subnet.
* `fosid` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController WtpProfileSplitTunnelingAcl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpprofile_splittunnelingacl.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

