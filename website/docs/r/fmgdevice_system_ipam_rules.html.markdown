---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipam_rules"
description: |-
  Configure IPAM allocation rules.
---

# fmgdevice_system_ipam_rules
Configure IPAM allocation rules.

~> This resource is a sub resource for variable `rules` of resource `fmgdevice_system_ipam`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ipam_rules" "trname" {
  description = "your own value"
  device      = ["your own value"]
  dhcp        = "enable"
  interface   = ["port2"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `description` - Description.
* `device` - Configure serial number or wildcard of FortiGate to match.
* `dhcp` - Enable/disable DHCP server for matching IPAM interfaces. Valid values: `disable`, `enable`.

* `interface` - Configure name or wildcard of interface to match.
* `name` - IPAM rule name.
* `pool` - Configure name of IPAM pool to use.
* `role` - Configure role of interface to match. Valid values: `any`, `lan`, `wan`, `dmz`, `undefined`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System IpamRules can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipam_rules.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

