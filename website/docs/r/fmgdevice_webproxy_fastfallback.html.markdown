---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_fastfallback"
description: |-
  Proxy destination connection fast-fallback.
---

# fmgdevice_webproxy_fastfallback
Proxy destination connection fast-fallback.

## Example Usage

```hcl
resource "fmgdevice_webproxy_fastfallback" "trname" {
  connection_mode    = "sequentially"
  connection_timeout = 10
  name               = "your own value"
  protocol           = "IPv4-only"
  status             = "enable"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `connection_mode` - Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.

* `connection_timeout` - Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
* `name` - Configure a name for the fast-fallback entry.
* `protocol` - Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.

* `status` - Enable/disable the fast-fallback entry. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy FastFallback can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_fastfallback.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

