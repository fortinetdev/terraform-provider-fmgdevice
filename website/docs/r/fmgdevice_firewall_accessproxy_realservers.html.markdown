---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_accessproxy_realservers"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_firewall_accessproxy_realservers
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `realservers` of resource `fmgdevice_firewall_accessproxy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `access_proxy` - Access Proxy.

* `fosid` - Id.
* `ip` - Ip.
* `port` - Port.
* `status` - Status. Valid values: `active`, `standby`, `disable`.

* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall AccessProxyRealservers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "access_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_accessproxy_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

