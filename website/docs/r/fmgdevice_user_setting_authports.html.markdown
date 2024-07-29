---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_setting_authports"
description: |-
  Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET.
---

# fmgdevice_user_setting_authports
Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET.

~> This resource is a sub resource for variable `auth_ports` of resource `fmgdevice_user_setting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_user_setting_authports" "trname" {
  fosid       = 10
  port        = 10
  type        = "telnet"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - ID.
* `port` - Non-standard port for firewall user authentication.
* `type` - Service type. Valid values: `http`, `https`, `ftp`, `telnet`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User SettingAuthPorts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_setting_authports.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

