---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_autoupdate_tunneling"
description: |-
  Configure web proxy tunneling for the FDN.
---

# fmgdevice_system_autoupdate_tunneling
Configure web proxy tunneling for the FDN.

## Example Usage

```hcl
resource "fmgdevice_system_autoupdate_tunneling" "trname" {
  address     = "your own value"
  password    = ["your own value"]
  port        = 10
  status      = "enable"
  username    = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `address` - Web proxy IP address or FQDN.
* `password` - Web proxy password.
* `port` - Web proxy port.
* `status` - Enable/disable web proxy tunneling. Valid values: `disable`, `enable`.

* `username` - Web proxy username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoupdateTunneling can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_autoupdate_tunneling.labelname SystemAutoupdateTunneling
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

