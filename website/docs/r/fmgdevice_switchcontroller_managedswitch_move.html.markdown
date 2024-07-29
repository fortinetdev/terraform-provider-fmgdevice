---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_move"
description: |-
  Move FortiSwitch devices that are managed by this FortiGate.
---

# fmgdevice_switchcontroller_managedswitch_move
Move FortiSwitch devices that are managed by this FortiGate.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_move" "trname" {
  target         = 54
  option         = "after"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
  managed_switch = 53
  depends_on     = [fmgdevice_switchcontroller_managedswitch.trname1, fmgdevice_switchcontroller_managedswitch.trname2]
}

resource "fmgdevice_switchcontroller_managedswitch" "trname1" {
  switch_id   = 53
  sn          = "S548DN4K16000230"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}

resource "fmgdevice_switchcontroller_managedswitch" "trname2" {
  switch_id   = 54
  sn          = "S548DN4K16000229"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the managed switch changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of managed switches.
