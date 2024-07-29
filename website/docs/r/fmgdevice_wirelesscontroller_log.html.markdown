---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_log"
description: |-
  Configure wireless controller event log filters.
---

# fmgdevice_wirelesscontroller_log
Configure wireless controller event log filters.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_log" "trname" {
  addrgrp_log    = "notification"
  ble_log        = "critical"
  clb_log        = "information"
  dhcp_starv_log = "warning"
  led_sched_log  = "alert"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `addrgrp_log` - Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `ble_log` - Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `clb_log` - Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `dhcp_starv_log` - Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `led_sched_log` - Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `radio_event_log` - Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `rogue_event_log` - Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `sta_event_log` - Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `sta_locate_log` - Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable wireless event logging. Valid values: `disable`, `enable`.

* `wids_log` - Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `wtp_event_log` - Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `wtp_fips_event_log` - Lowest severity level to log FAP fips event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Log can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_log.labelname WirelessControllerLog
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

