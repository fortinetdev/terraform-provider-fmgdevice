---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fortindr"
description: |-
  Configure FortiNDR.
---

# fmgdevice_system_fortindr
Configure FortiNDR.

## Example Usage

```hcl
resource "fmgdevice_system_fortindr" "trname" {
  interface               = ["port2"]
  interface_select_method = "auto"
  source_ip               = "your own value"
  status                  = "enable"
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for communications to FortiNDR.
* `status` - Enable/disable FortiNDR. Valid values: `disable`, `enable`.

* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortindr can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fortindr.labelname SystemFortindr
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

