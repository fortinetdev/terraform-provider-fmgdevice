---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dpdk_global"
description: |-
  Configure global DPDK options.
---

# fmgdevice_dpdk_global
Configure global DPDK options.

## Example Usage

```hcl
resource "fmgdevice_dpdk_global" "trname" {
  elasticbuffer       = "disable"
  hugepage_percentage = 10
  interface           = ["port2"]
  ipsec_offload       = "disable"
  mbufpool_percentage = 10
  device_name         = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `elasticbuffer` - Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.

* `hugepage_percentage` - Percentage of main memory allocated to hugepages, which are available for DPDK operation.
* `interface` - Physical interfaces that enable DPDK.
* `ipsec_offload` - Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.

* `mbufpool_percentage` - Percentage of main memory allocated to DPDK packet buffer.
* `multiqueue` - Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.

* `per_session_accounting` - Enable/disable per-session accounting. Valid values: `enable`, `disable`, `traffic-log-only`.

* `protects` - Special arguments for device
* `session_table_percentage` - Percentage of main memory allocated to DPDK session table.
* `sleep_on_idle` - Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.

* `status` - Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dpdk Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dpdk_global.labelname DpdkGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

