---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_pppoeinterface"
description: |-
  Configure the PPPoE interfaces.
---

# fmgdevice_system_pppoeinterface
Configure the PPPoE interfaces.

## Example Usage

```hcl
resource "fmgdevice_system_pppoeinterface" "trname" {
  ac_name            = "your own value"
  auth_type          = "mschapv1"
  device             = ["your own value"]
  dial_on_demand     = "disable"
  disc_retry_timeout = 10
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ac_name` - PPPoE AC name.
* `auth_type` - PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.

* `device` - Name for the physical interface.
* `dial_on_demand` - Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `disable`, `enable`.

* `disc_retry_timeout` - PPPoE discovery init timeout value in (0-4294967295 sec).
* `idle_timeout` - PPPoE auto disconnect after idle timeout (0-4294967295 sec).
* `ipunnumbered` - PPPoE unnumbered IP.
* `ipv6` - Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `disable`, `enable`.

* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `name` - Name of the PPPoE interface.
* `padt_retry_timeout` - PPPoE terminate timeout value in (0-4294967295 sec).
* `password` - Enter the password.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `disable`, `enable`.

* `service_name` - PPPoE service name.
* `username` - User name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System PppoeInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_pppoeinterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

