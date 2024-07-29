---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_concentrator"
description: |-
  Concentrator configuration.
---

# fmgdevice_vpn_ipsec_concentrator
Concentrator configuration.

## Example Usage

```hcl
resource "fmgdevice_vpn_ipsec_concentrator" "trname" {
  fosid       = 10
  member      = ["your own value"]
  name        = "your own value"
  src_check   = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Concentrator ID (1 - 65535).
* `member` - Names of up to 3 VPN tunnels to add to the concentrator.
* `name` - Concentrator name.
* `src_check` - Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Vpn IpsecConcentrator can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_concentrator.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

