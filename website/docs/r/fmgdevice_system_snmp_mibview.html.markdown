---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_snmp_mibview"
description: |-
  SNMP Access Control MIB View configuration.
---

# fmgdevice_system_snmp_mibview
SNMP Access Control MIB View configuration.

## Example Usage

```hcl
resource "fmgdevice_system_snmp_mibview" "trname" {
  exclude     = ["your own value"]
  include     = ["your own value"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `exclude` - OID subtrees to be excluded in the view. Maximum 64 allowed.
* `include` - OID subtrees to be included in the view. Maximum 16 allowed.
* `name` - MIB view name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SnmpMibView can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_snmp_mibview.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

