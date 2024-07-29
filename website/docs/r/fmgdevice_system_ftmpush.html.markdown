---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ftmpush"
description: |-
  Configure FortiToken Mobile push services.
---

# fmgdevice_system_ftmpush
Configure FortiToken Mobile push services.

## Example Usage

```hcl
resource "fmgdevice_system_ftmpush" "trname" {
  proxy       = "disable"
  server      = "your own value"
  server_cert = ["your own value"]
  server_ip   = "your own value"
  server_port = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `proxy` - Enable/disable communication to the proxy server in FortiGuard configuration. Valid values: `disable`, `enable`.

* `server` - IPv4 address or domain name of FortiToken Mobile push services server.
* `server_cert` - Name of the server certificate to be used for SSL.
* `server_ip` - IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
* `server_port` - Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
* `status` - Enable/disable the use of FortiToken Mobile push services. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FtmPush can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ftmpush.labelname SystemFtmPush
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

