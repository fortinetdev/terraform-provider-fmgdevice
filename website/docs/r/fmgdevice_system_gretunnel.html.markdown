---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_gretunnel"
description: |-
  Configure GRE tunnel.
---

# fmgdevice_system_gretunnel
Configure GRE tunnel.

## Example Usage

```hcl
resource "fmgdevice_system_gretunnel" "trname" {
  auto_asic_offload     = "enable"
  checksum_reception    = "enable"
  checksum_transmission = "disable"
  diffservcode          = "your own value"
  dscp_copying          = "disable"
  name                  = "your own value"
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_asic_offload` - Enable/disable automatic ASIC offloading. Valid values: `disable`, `enable`.

* `checksum_reception` - Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.

* `checksum_transmission` - Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.

* `diffservcode` - DiffServ setting to be applied to GRE tunnel outer IP header.
* `dscp_copying` - Enable/disable DSCP copying. Valid values: `disable`, `enable`.

* `interface` - Interface name.
* `ip_version` - IP version to use for VPN interface. Valid values: `4`, `6`.

* `keepalive_failtimes` - Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
* `keepalive_interval` - Keepalive message interval (0 - 32767, 0 = disabled).
* `key_inbound` - Require received GRE packets contain this key (0 - 4294967295).
* `key_outbound` - Include this key in transmitted GRE packets (0 - 4294967295).
* `local_gw` - IP address of the local gateway.
* `local_gw6` - IPv6 address of the local gateway.
* `name` - Tunnel name.
* `remote_gw` - IP address of the remote gateway.
* `remote_gw6` - IPv6 address of the remote gateway.
* `sequence_number_reception` - Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.

* `sequence_number_transmission` - Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.

* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System GreTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_gretunnel.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

