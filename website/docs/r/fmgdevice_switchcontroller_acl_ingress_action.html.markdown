---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_acl_ingress_action"
description: |-
  ACL actions.
---

# fmgdevice_switchcontroller_acl_ingress_action
ACL actions.

~> This resource is a sub resource for variable `action` of resource `fmgdevice_switchcontroller_acl_ingress`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_acl_ingress_action" "trname" {
  count       = "disable"
  drop        = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ingress` - Ingress.

* `fmgcount` - Enable/disable count. Valid values: `disable`, `enable`.

* `drop` - Enable/disable drop. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController AclIngressAction can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ingress=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_acl_ingress_action.labelname SwitchControllerAclIngressAction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

