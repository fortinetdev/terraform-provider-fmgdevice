---
subcategory: "Firewall"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_policy_fabricpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i> Fabric policy related attributes.
---

# fmgdevice_firewall_policy_fabricpolicy
<i>This object will be purged after policy copy and install.</i> Fabric policy related attributes.

~> This resource is a sub resource for variable `fabric_policy` of resource `fmgdevice_firewall_policy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `policy` - Policy.

* `from` - From device.
* `to` - To device.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall PolicyFabricPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "policy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_policy_fabricpolicy.labelname FirewallPolicyFabricPolicy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

