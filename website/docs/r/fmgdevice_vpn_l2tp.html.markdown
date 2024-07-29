---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_l2tp"
description: |-
  Configure L2TP.
---

# fmgdevice_vpn_l2tp
Configure L2TP.

## Example Usage

```hcl
resource "fmgdevice_vpn_l2tp" "trname" {
  compress          = "enable"
  eip               = "your own value"
  enforce_ipsec     = "disable"
  hello_interval    = 10
  lcp_echo_interval = 10
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `compress` - Enable/disable data compression. Valid values: `disable`, `enable`.

* `eip` - End IP.
* `enforce_ipsec` - Enable/disable IPsec enforcement. Valid values: `disable`, `enable`.

* `hello_interval` - L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum number of missed LCP echo messages before disconnect.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a L2TP gateway. Valid values: `disable`, `enable`.

* `usrgrp` - User group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn L2Tp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_l2tp.labelname VpnL2Tp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

