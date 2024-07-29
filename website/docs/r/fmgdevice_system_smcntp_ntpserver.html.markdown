---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_smcntp_ntpserver"
description: |-
  Configure the FortiGate SMC to connect to an NTP server.
---

# fmgdevice_system_smcntp_ntpserver
Configure the FortiGate SMC to connect to an NTP server.

~> This resource is a sub resource for variable `ntpserver` of resource `fmgdevice_system_smcntp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_smcntp_ntpserver" "trname" {
  fosid       = 10
  server      = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - NTP server ID.
* `server` - IP address of the NTP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SmcNtpNtpserver can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_smcntp_ntpserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

