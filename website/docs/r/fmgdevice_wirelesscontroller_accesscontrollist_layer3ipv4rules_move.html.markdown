---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules_move"
description: |-
  AP ACL layer3 ipv4 rule list.
---

# fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules_move
AP ACL layer3 ipv4 rule list.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules_move" "trname" {
  target              = 11
  layer3_ipv4_rules   = 12
  option              = "after"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
  access_control_list = fmgdevice_wirelesscontroller_accesscontrollist.trname1.name
  depends_on          = [fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules.trname1, fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules.trname2]
}

resource "fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules" "trname2" {
  rule_id             = 11
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
  access_control_list = fmgdevice_wirelesscontroller_accesscontrollist.trname1.name
  depends_on          = [fmgdevice_wirelesscontroller_accesscontrollist.trname1]
}

resource "fmgdevice_wirelesscontroller_accesscontrollist_layer3ipv4rules" "trname1" {
  rule_id             = 12
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
  access_control_list = fmgdevice_wirelesscontroller_accesscontrollist.trname1.name
  depends_on          = [fmgdevice_wirelesscontroller_accesscontrollist.trname1]
}

resource "fmgdevice_wirelesscontroller_accesscontrollist" "trname1" {
  name        = "test4"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `access_control_list` - Access Control List.
* `layer3_ipv4_rules` - Layer3 Ipv4 Rules.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_id}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the layer3 ipv4 rules changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of layer3 ipv4 ruless.
