---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_virtualportpool"
description: |-
  Configure virtual pool.
---

# fmgdevice_switchcontroller_virtualportpool
Configure virtual pool.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_virtualportpool" "trname" {
  description = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Virtual switch pool description.
* `name` - Virtual switch pool name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController VirtualPortPool can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_virtualportpool.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

