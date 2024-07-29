---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtp_move"
description: |-
  Move Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.
---

# fmgdevice_wirelesscontroller_wtp_move
Move Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtp_move" "trname" {
  target      = "test3"
  option      = "after"
  wtp         = "test4"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_wirelesscontroller_wtp.trname2, fmgdevice_wirelesscontroller_wtp.trname1]
}

resource "fmgdevice_wirelesscontroller_wtp" "trname2" {
  wtp_id      = "test3"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}

resource "fmgdevice_wirelesscontroller_wtp" "trname1" {
  wtp_id      = "test4"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp` - Wtp.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{wtp_id}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the wtp changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of wtps.
