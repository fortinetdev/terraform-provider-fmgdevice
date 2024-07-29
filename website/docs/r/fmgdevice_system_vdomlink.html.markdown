---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomlink"
description: |-
  Configure VDOM links.
---

# fmgdevice_system_vdomlink
Configure VDOM links.

## Example Usage

```hcl
resource "fmgdevice_system_vdomlink" "trname" {
  name        = "your own value"
  type        = "ppp"
  vcluster    = "vcluster2"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - VDOM link name (maximum = 11 characters).
* `type` - VDOM link type: PPP or Ethernet. Valid values: `ppp`, `ethernet`, `npupair`.

* `vcluster` - Virtual cluster. Valid values: `vcluster1`, `vcluster2`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomLink can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomlink.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

