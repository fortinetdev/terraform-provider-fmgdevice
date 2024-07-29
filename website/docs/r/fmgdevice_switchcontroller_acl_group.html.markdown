---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_acl_group"
description: |-
  Configure ACL groups to be applied on managed FortiSwitch ports.
---

# fmgdevice_switchcontroller_acl_group
Configure ACL groups to be applied on managed FortiSwitch ports.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_acl_group" "trname" {
  ingress     = ["your own value"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ingress` - Configure ingress ACL policies in group.
* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController AclGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_acl_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

