---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_speedtestsetting"
description: |-
  Configure speed test setting.
---

# fmgdevice_system_speedtestsetting
Configure speed test setting.

## Example Usage

```hcl
resource "fmgdevice_system_speedyour valuesetting" "trname" {
  latency_threshold   = 10
  multiple_tcp_stream = 10
  device_name         = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `latency_threshold` - Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
* `multiple_tcp_stream` - Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SpeedTestSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_speedtestsetting.labelname SystemSpeedTestSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

