---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fortisandbox"
description: |-
  Configure FortiSandbox.
---

# fmgdevice_system_fortisandbox
Configure FortiSandbox.

## Example Usage

```hcl
resource "fmgdevice_system_fortisandbox" "trname" {
  email         = "your own value"
  enc_algorithm = "disable"
  forticloud    = "disable"
  inline_scan   = "disable"
  interface     = ["port2"]
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `email` - Notifier email address.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`, `disable`.

* `forticloud` - Enable/disable FortiSandbox Cloud. Valid values: `disable`, `enable`.

* `inline_scan` - Enable/disable FortiSandbox inline scan. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `server` - Server IP address or FQDN of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `status` - Enable/disable FortiSandbox. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortisandbox can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fortisandbox.labelname SystemFortisandbox
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

