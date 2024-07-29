---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomproperty"
description: |-
  Configure VDOM property.
---

# fmgdevice_system_vdomproperty
Configure VDOM property.

## Example Usage

```hcl
resource "fmgdevice_system_vdomproperty" "trname" {
  custom_service   = 10
  description      = "your own value"
  dialup_tunnel    = 10
  firewall_address = 10
  firewall_addrgrp = 10
  name             = "your own value"
  device_name      = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `custom_service` - Maximum guaranteed number of firewall custom services.
* `description` - Description.
* `dialup_tunnel` - Maximum guaranteed number of dial-up tunnels.
* `firewall_address` - Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
* `firewall_addrgrp` - Maximum guaranteed number of firewall address groups (IPv4, IPv6).
* `firewall_policy` - Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
* `ipsec_phase1` - Maximum guaranteed number of VPN IPsec phase 1 tunnels.
* `ipsec_phase1_interface` - Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
* `ipsec_phase2` - Maximum guaranteed number of VPN IPsec phase 2 tunnels.
* `ipsec_phase2_interface` - Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
* `log_disk_quota` - Log disk quota in megabytes (MB). Range depends on how much disk space is available.
* `name` - VDOM name.
* `onetime_schedule` - Maximum guaranteed number of firewall one-time schedules.
* `proxy` - Maximum guaranteed number of concurrent proxy users.
* `recurring_schedule` - Maximum guaranteed number of firewall recurring schedules.
* `service_group` - Maximum guaranteed number of firewall service groups.
* `session` - Maximum guaranteed number of sessions.
* `snmp_index` - Permanent SNMP Index of the virtual domain (1 - 2147483647).
* `sslvpn` - Maximum guaranteed number of SSL-VPNs.
* `user` - Maximum guaranteed number of local users.
* `user_group` - Maximum guaranteed number of user groups.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomProperty can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomproperty.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

