---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_tacacsaccounting3_setting"
description: |-
  Settings for TACACS+ accounting.
---

# fmgdevice_log_tacacsaccounting3_setting
Settings for TACACS+ accounting.

## Example Usage

```hcl
resource "fmgdevice_log_tacacsaccounting3_setting" "trname" {
  interface               = ["port2"]
  interface_select_method = "specify"
  server                  = "your own value"
  server_key              = ["your own value"]
  source_ip               = "your own value"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `server` - Address of TACACS+ server.
* `server_key` - Key to access the TACACS+ server.
* `source_ip` - Source IP address for communication to TACACS+ server.
* `status` - Enable/disable TACACS+ accounting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log TacacsAccounting3Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_tacacsaccounting3_setting.labelname LogTacacsAccounting3Setting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

