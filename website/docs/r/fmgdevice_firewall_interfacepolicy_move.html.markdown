---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_interfacepolicy_move"
description: |-
  <i>This object will be purged after policy copy and install.</i> Move IPv4 interface policies.
---

# fmgdevice_firewall_interfacepolicy_move
<i>This object will be purged after policy copy and install.</i> Move IPv4 interface policies.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `interface_policy` - Interface Policy.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the interface policy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of interface policies.
