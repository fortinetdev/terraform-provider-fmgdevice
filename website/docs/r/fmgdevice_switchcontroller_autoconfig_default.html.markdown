---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_autoconfig_default"
description: |-
  Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.
---

# fmgdevice_switchcontroller_autoconfig_default
Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_autoconfig_default" "trname" {
  fgt_policy  = "your own value"
  icl_policy  = ["your own value"]
  isl_policy  = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fgt_policy` - Default FortiLink auto-config policy.
* `icl_policy` - Default ICL auto-config policy.
* `isl_policy` - Default ISL auto-config policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController AutoConfigDefault can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_autoconfig_default.labelname SwitchControllerAutoConfigDefault
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

