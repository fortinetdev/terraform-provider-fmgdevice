---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_policy_move"
description: |-
  Move IPv4 routing policies.
---

# fmgdevice_router_policy_move
Move IPv4 routing policies.

## Example Usage

```hcl
resource "fmgdevice_router_policy_move" "trname" {
  target      = 7
  policy      = 8
  option      = "after"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_router_policy.trname1, fmgdevice_router_policy.trname2]
}

resource "fmgdevice_router_policy" "trname1" {
  seq_num     = 7
  gateway     = "2.3.4.7"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}

resource "fmgdevice_router_policy" "trname2" {
  seq_num     = 8
  gateway     = "2.3.4.6"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `policy` - Policy.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the policy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of policies.
