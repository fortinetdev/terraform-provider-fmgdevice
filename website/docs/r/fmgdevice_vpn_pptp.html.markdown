---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_pptp"
description: |-
  Configure PPTP.
---

# fmgdevice_vpn_pptp
Configure PPTP.

## Example Usage

```hcl
resource "fmgdevice_vpn_pptp" "trname" {
  eip         = "your own value"
  ip_mode     = "range"
  local_ip    = "your own value"
  sip         = "your own value"
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `eip` - End IP.
* `ip_mode` - IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.

* `local_ip` - Local IP to be used for peer's remote IP.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a PPTP gateway. Valid values: `disable`, `enable`.

* `usrgrp` - User group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn Pptp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_pptp.labelname VpnPptp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

