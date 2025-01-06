---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_settings"
description: |-
  Configure IPS VDOM parameter.
---

# fmgdevice_ips_settings
Configure IPS VDOM parameter.

## Example Usage

```hcl
resource "fmgdevice_ips_settings" "trname" {
  ips_packet_quota       = 10
  packet_log_history     = 10
  packet_log_memory      = 10
  packet_log_post_attack = 10
  proxy_inline_ips       = "enable"
  device_name            = var.device_name # not required if setting is at provider
  device_vdom            = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ha_session_pickup` - IPS HA failover session pickup preference. Valid values: `connectivity`, `security`.

* `ips_packet_quota` - Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
* `packet_log_history` - Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
* `packet_log_memory` - Maximum memory can be used by packet log (64 - 8192 kB).
* `packet_log_post_attack` - Number of packets to log after the IPS signature is detected (0 - 255).
* `proxy_inline_ips` - Enable/disable proxy-mode policy inline IPS support. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_settings.labelname IpsSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

