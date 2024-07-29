---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_snmp_sysinfo"
description: |-
  SNMP system info configuration.
---

# fmgdevice_system_snmp_sysinfo
SNMP system info configuration.

## Example Usage

```hcl
resource "fmgdevice_system_snmp_sysinfo" "trname" {
  append_index   = "disable"
  contact_info   = "your own value"
  description    = "your own value"
  engine_id      = "your own value"
  engine_id_type = "hex"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `append_index` - Enable/disable allowance of appending vdom or interface index in some RFC tables. Valid values: `disable`, `enable`.

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engineID string (maximum 27 characters).
* `engine_id_type` - Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.

* `location` - System location.
* `status` - Enable/disable SNMP. Valid values: `disable`, `enable`.

* `trap_free_memory_threshold` - Free memory usage when trap is sent.
* `trap_freeable_memory_threshold` - Freeable memory usage when trap is sent.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SnmpSysinfo can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_snmp_sysinfo.labelname SystemSnmpSysinfo
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

