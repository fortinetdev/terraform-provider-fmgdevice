---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_vip_gslbpublicips"
description: |-
  <i>This object will be purged after policy copy and install.</i> Publicly accessible IP addresses for the FortiGSLB service.
---

# fmgdevice_firewall_vip_gslbpublicips
<i>This object will be purged after policy copy and install.</i> Publicly accessible IP addresses for the FortiGSLB service.

~> This resource is a sub resource for variable `gslb_public_ips` of resource `fmgdevice_firewall_vip`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `vip` - Vip.

* `index` - Index of this public IP setting.
* `ip` - The publicly accessible IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

Firewall VipGslbPublicIps can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "vip=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_vip_gslbpublicips.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

