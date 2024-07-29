---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_debugurl"
description: |-
  Configure debug URL addresses.
---

# fmgdevice_webproxy_debugurl
Configure debug URL addresses.

## Example Usage

```hcl
resource "fmgdevice_webproxy_debugurl" "trname" {
  exact       = "enable"
  name        = "your own value"
  status      = "enable"
  url_pattern = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `exact` - Enable/disable matching the exact path. Valid values: `disable`, `enable`.

* `name` - Debug URL name.
* `status` - Enable/disable this URL exemption. Valid values: `disable`, `enable`.

* `url_pattern` - URL exemption pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy DebugUrl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_debugurl.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

