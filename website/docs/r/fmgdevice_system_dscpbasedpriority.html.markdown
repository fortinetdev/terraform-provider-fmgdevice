---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dscpbasedpriority"
description: |-
  Configure DSCP based priority table.
---

# fmgdevice_system_dscpbasedpriority
Configure DSCP based priority table.

## Example Usage

```hcl
resource "fmgdevice_system_dscpbasedpriority" "trname" {
  ds          = 10
  fosid       = 10
  priority    = "high"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ds` - DSCP(DiffServ) DS value (0 - 63).
* `fosid` - Item ID.
* `priority` - DSCP based priority level. Valid values: `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DscpBasedPriority can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dscpbasedpriority.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

