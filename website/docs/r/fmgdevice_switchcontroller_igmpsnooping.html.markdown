---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_igmpsnooping"
description: |-
  Configure FortiSwitch IGMP snooping global settings.
---

# fmgdevice_switchcontroller_igmpsnooping
Configure FortiSwitch IGMP snooping global settings.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_igmpsnooping" "trname" {
  aging_time              = 10
  flood_unknown_multicast = "disable"
  query_interval          = 10
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `aging_time` - Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding. Valid values: `disable`, `enable`.

* `query_interval` - Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController IgmpSnooping can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_igmpsnooping.labelname SwitchControllerIgmpSnooping
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

