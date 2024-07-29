---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_speedtestschedule"
description: |-
  Speed test schedule for each interface.
---

# fmgdevice_system_speedtestschedule
Speed test schedule for each interface.

## Example Usage

```hcl
resource "fmgdevice_system_speedyour valueschedule" "trname" {
  ctrl_port      = 10
  diffserv       = "your own value"
  dynamic_server = "disable"
  interface      = ["port2"]
  mode           = "Auto"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ctrl_port` - Port of the controller to get access token.
* `diffserv` - DSCP used for speed test.
* `dynamic_server` - Enable/disable dynamic server option. Valid values: `disable`, `enable`.

* `interface` - Interface name.
* `mode` - Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP`, `TCP`, `Auto`.

* `schedules` - Schedules for the interface.
* `server_name` - Speed test server name.
* `server_port` - Port of the server to run speed test.
* `status` - Enable/disable scheduled speed test. Valid values: `disable`, `enable`.

* `update_inbandwidth` - Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.

* `update_inbandwidth_maximum` - Maximum downloading bandwidth (kbps) to be used in a speed test.
* `update_inbandwidth_minimum` - Minimum downloading bandwidth (kbps) to be considered effective.
* `update_outbandwidth` - Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.

* `update_outbandwidth_maximum` - Maximum uploading bandwidth (kbps) to be used in a speed test.
* `update_outbandwidth_minimum` - Minimum uploading bandwidth (kbps) to be considered effective.
* `update_shaper` - Set egress shaper based on the test result. Valid values: `disable`, `local`, `remote`, `both`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

System SpeedTestSchedule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_speedtestschedule.labelname {{interface}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

