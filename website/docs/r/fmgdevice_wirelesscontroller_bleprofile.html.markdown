---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_bleprofile"
description: |-
  Configure Bluetooth Low Energy profile.
---

# fmgdevice_wirelesscontroller_bleprofile
Configure Bluetooth Low Energy profile.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_bleprofile" "trname" {
  advertising        = ["eddystone-uid"]
  beacon_interval    = 10
  ble_scanning       = "disable"
  comment            = "your own value"
  eddystone_instance = "your own value"
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `advertising` - Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.

* `beacon_interval` - Beacon interval (default = 100 msec).
* `ble_scanning` - Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `eddystone_instance` - Eddystone instance ID.
* `eddystone_namespace` - Eddystone namespace ID.
* `eddystone_url` - Eddystone URL.
* `eddystone_url_encode_hex` - Eddystone encoded URL hexadecimal string
* `ibeacon_uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `major_id` - Major ID.
* `minor_id` - Minor ID.
* `name` - Bluetooth Low Energy profile name.
* `scan_interval` - Scan Interval (default = 50 msec).
* `scan_period` - Scan Period (default = 4000 msec).
* `scan_threshold` - Minimum signal level/threshold in dBm required for the AP to report detected BLE device (-95 to -20, default = -90).
* `scan_time` - Scan Time (default = 1000 msec).
* `scan_type` - Scan Type (default = active). Valid values: `active`, `passive`.

* `scan_window` - Scan Windows (default = 50 msec).
* `txpower` - Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController BleProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_bleprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

