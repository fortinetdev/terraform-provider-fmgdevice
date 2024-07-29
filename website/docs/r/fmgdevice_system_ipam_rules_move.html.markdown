---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipam_rules_move"
description: |-
  Move IPAM allocation rules.
---

# fmgdevice_system_ipam_rules_move
Move IPAM allocation rules.

## Example Usage

```hcl
resource "fmgdevice_system_ipam_rules_move" "trname" {
  target      = "test2"
  option      = "after"
  rules       = "test3"
  device_name = var.device_name # not required if setting is at provider
  depends_on  = [fmgdevice_system_ipam_rules.trname1, fmgdevice_system_ipam_rules.trname2]
}

resource "fmgdevice_system_ipam_rules" "trname1" {
  name        = "test2"
  device_name = var.device_name # not required if setting is at provider
}

resource "fmgdevice_system_ipam_rules" "trname2" {
  name        = "test3"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `rules` - Rules.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the rules changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of rule.
