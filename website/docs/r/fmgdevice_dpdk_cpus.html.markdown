---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dpdk_cpus"
description: |-
  Configure CPUs enabled to run engines in each DPDK stage.
---

# fmgdevice_dpdk_cpus
Configure CPUs enabled to run engines in each DPDK stage.

## Example Usage

```hcl
resource "fmgdevice_dpdk_cpus" "trname" {
  ips_cpus      = "your own value"
  isolated_cpus = "your own value"
  rx_cpus       = "your own value"
  tx_cpus       = "your own value"
  vnp_cpus      = "your own value"
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ips_cpus` - CPUs enabled to run DPDK IPS engines.
* `isolated_cpus` - CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
* `rx_cpus` - CPUs enabled to run DPDK RX engines.
* `tx_cpus` - CPUs enabled to run DPDK TX engines.
* `vnp_cpus` - CPUs enabled to run DPDK VNP engines.
* `vnpsp_cpus` - CPUs enabled to run DPDK VNP slow path.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dpdk Cpus can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dpdk_cpus.labelname DpdkCpus
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

