---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_system"
description: |-
  Configure system-wide switch controller settings.
---

# fmgdevice_switchcontroller_system
Configure system-wide switch controller settings.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_system" "trname" {
  caputp_echo_interval      = 10
  caputp_max_retransmit     = 10
  data_sync_interval        = 10
  dynamic_periodic_interval = 10
  iot_holdoff               = 10
  device_name               = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `caputp_echo_interval` - Echo interval for the caputp echo requests from swtp.
* `caputp_max_retransmit` - Maximum retransmission count for the caputp tunnel packets.
* `data_sync_interval` - Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
* `dynamic_periodic_interval` - Periodic time interval to run Dynamic port policy engine (5 - 180 sec, default = 60).
* `iot_holdoff` - MAC entry's creation time. Time must be greater than this value for an entry to be created (0 - 10080 mins, default = 5 mins).
* `iot_mac_idle` - MAC entry's idle time. MAC entry is removed after this value (0 - 10080 mins, default = 1440 mins).
* `iot_scan_interval` - IoT scan interval (2 - 10080 mins, default = 60 mins, 0 = disable).
* `iot_weight_threshold` - MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
* `nac_periodic_interval` - Periodic time interval to run NAC engine (5 - 180 sec, default = 60).
* `parallel_process` - Maximum number of parallel processes.
* `parallel_process_override` - Enable/disable parallel process override. Valid values: `disable`, `enable`.

* `tunnel_mode` - Compatible/strict tunnel mode. Valid values: `compatible`, `strict`, `moderate`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController System can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_system.labelname SwitchControllerSystem
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

