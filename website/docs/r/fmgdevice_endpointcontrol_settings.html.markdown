---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_endpointcontrol_settings"
description: |-
  Configure endpoint control settings.
---

# fmgdevice_endpointcontrol_settings
Configure endpoint control settings.

## Example Usage

```hcl
resource "fmgdevice_endpointcontrol_settings" "trname" {
  forticlient_disconnect_unsupported_client = "disable"
  forticlient_keepalive_interval            = 10
  forticlient_sys_update_interval           = 10
  forticlient_user_avatar                   = "enable"
  override                                  = "enable"
  device_name                               = var.device_name # not required if setting is at provider
  device_vdom                               = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `forticlient_disconnect_unsupported_client` - Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `disable`, `enable`.

* `forticlient_keepalive_interval` - Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
* `forticlient_sys_update_interval` - Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
* `forticlient_user_avatar` - Enable/disable uploading FortiClient user avatars. Valid values: `disable`, `enable`.

* `override` - Override global EMS table for this VDOM. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

EndpointControl Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_endpointcontrol_settings.labelname EndpointControlSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

