---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ntp_ntpserver"
description: |-
  Configure the FortiGate to connect to any available third-party NTP server.
---

# fmgdevice_system_ntp_ntpserver
Configure the FortiGate to connect to any available third-party NTP server.

~> This resource is a sub resource for variable `ntpserver` of resource `fmgdevice_system_ntp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ntp_ntpserver" "trname" {
  authentication          = "disable"
  fosid                   = 10
  interface               = ["port2"]
  interface_select_method = "specify"
  ip_type                 = "IPv4"
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `authentication` - Enable/disable authentication. Valid values: `disable`, `enable`.

* `fosid` - NTP server ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip_type` - Choose to connect to IPv4 or/and IPv6 NTP server. Valid values: `IPv6`, `IPv4`, `Both`.

* `key` - Key for MD5(NTPv3)/SHA1(NTPv4)/SHA256(NTPv4) authentication.
* `key_id` - Key ID for authentication.
* `key_type` - Select NTP authentication type. Valid values: `SHA1`, `SHA256`, `MD5`.

* `ntpv3` - Enable to use NTPv3 instead of NTPv4. Valid values: `disable`, `enable`.

* `server` - IP address or hostname of the NTP Server.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System NtpNtpserver can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ntp_ntpserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

