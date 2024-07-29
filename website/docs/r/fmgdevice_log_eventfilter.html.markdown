---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_eventfilter"
description: |-
  Configure log event filters.
---

# fmgdevice_log_eventfilter
Configure log event filters.

## Example Usage

```hcl
resource "fmgdevice_log_eventfilter" "trname" {
  cifs          = "enable"
  connector     = "disable"
  endpoint      = "disable"
  event         = "enable"
  fortiextender = "enable"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cifs` - Enable/disable CIFS logging. Valid values: `disable`, `enable`.

* `connector` - Enable/disable SDN connector logging. Valid values: `disable`, `enable`.

* `endpoint` - Enable/disable endpoint event logging. Valid values: `disable`, `enable`.

* `event` - Enable/disable event logging. Valid values: `disable`, `enable`.

* `fortiextender` - Enable/disable FortiExtender logging. Valid values: `disable`, `enable`.

* `ha` - Enable/disable ha event logging. Valid values: `disable`, `enable`.

* `rest_api` - Enable/disable REST API logging. Valid values: `disable`, `enable`.

* `router` - Enable/disable router event logging. Valid values: `disable`, `enable`.

* `sdwan` - Enable/disable SD-WAN logging. Valid values: `disable`, `enable`.

* `security_rating` - Enable/disable Security Rating result logging. Valid values: `disable`, `enable`.

* `switch_controller` - Enable/disable Switch-Controller logging. Valid values: `disable`, `enable`.

* `system` - Enable/disable system event logging. Valid values: `disable`, `enable`.

* `user` - Enable/disable user authentication event logging. Valid values: `disable`, `enable`.

* `vpn` - Enable/disable VPN event logging. Valid values: `disable`, `enable`.

* `wan_opt` - Enable/disable WAN optimization event logging. Valid values: `disable`, `enable`.

* `webproxy` - Enable/disable web proxy event logging. Valid values: `disable`, `enable`.

* `wireless_activity` - Enable/disable wireless event logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Eventfilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_eventfilter.labelname LogEventfilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

