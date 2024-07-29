---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ndproxy"
description: |-
  Configure IPv6 neighbor discovery proxy (RFC4389).
---

# fmgdevice_system_ndproxy
Configure IPv6 neighbor discovery proxy (RFC4389).

## Example Usage

```hcl
resource "fmgdevice_system_ndproxy" "trname" {
  member      = ["your own value"]
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `member` - Interfaces using the neighbor discovery proxy.
* `status` - Enable/disable neighbor discovery proxy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NdProxy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ndproxy.labelname SystemNdProxy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

