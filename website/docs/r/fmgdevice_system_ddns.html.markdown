---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ddns"
description: |-
  Configure DDNS.
---

# fmgdevice_system_ddns
Configure DDNS.

## Example Usage

```hcl
resource "fmgdevice_system_ddns" "trname" {
  addr_type   = "ipv6"
  bound_ip    = "your own value"
  clear_text  = "disable"
  ddns_auth   = "disable"
  ddns_domain = "your own value"
  ddnsid      = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `addr_type` - Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.

* `bound_ip` - Bound-Ip.
* `clear_text` - Enable/disable use of clear text connections. Valid values: `disable`, `enable`.

* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.

* `ddns_domain` - Your fully qualified domain name. For example, yourname.ddns.com.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_password` - DDNS password.
* `ddns_server` - Select a DDNS service provider. Valid values: `dhs.org`, `dyndns.org`, `dyns.net`, `tzo.com`, `ods.org`, `vavic.com`, `now.net.cn`, `dipdns.net`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.

* `ddns_server_addr` - Generic DDNS server IP/FQDN list.
* `ddns_server_ip` - Ddns-Server-Ip.
* `ddns_sn` - DDNS Serial Number.
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_username` - DDNS user name.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddnsid` - DDNS ID.
* `monitor_interface` - Monitored interface.
* `server_type` - Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.

* `ssl_certificate` - Name of local certificate for SSL connections.
* `update_interval` - DDNS update interval (60 - 2592000 sec, 0 means default).
* `use_public_ip` - Enable/disable use of public IP address. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ddnsid}}.

## Import

System Ddns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ddns.labelname {{ddnsid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

