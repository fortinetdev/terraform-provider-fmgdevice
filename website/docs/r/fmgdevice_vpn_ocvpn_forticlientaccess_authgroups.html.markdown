---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ocvpn_forticlientaccess_authgroups"
description: |-
  FortiClient user authentication groups.
---

# fmgdevice_vpn_ocvpn_forticlientaccess_authgroups
FortiClient user authentication groups.

~> This resource is a sub resource for variable `auth_groups` of resource `fmgdevice_vpn_ocvpn_forticlientaccess`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_ocvpn_forticlientaccess_authgroups" "trname" {
  auth_group  = ["your own value"]
  name        = "your own value"
  overlays    = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_group` - Authentication user group for FortiClient access.
* `name` - Group name.
* `overlays` - OCVPN overlays to allow access to.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn OcvpnForticlientAccessAuthGroups can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ocvpn_forticlientaccess_authgroups.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

