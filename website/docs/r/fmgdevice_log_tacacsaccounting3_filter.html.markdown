---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_tacacsaccounting3_filter"
description: |-
  Settings for TACACS+ accounting events filter.
---

# fmgdevice_log_tacacsaccounting3_filter
Settings for TACACS+ accounting events filter.

## Example Usage

```hcl
resource "fmgdevice_log_tacacsaccounting3_filter" "trname" {
  cli_cmd_audit       = "disable"
  config_change_audit = "enable"
  login_audit         = "enable"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cli_cmd_audit` - Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `disable`, `enable`.

* `config_change_audit` - Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `disable`, `enable`.

* `login_audit` - Enable/disable TACACS+ accounting for login events audit. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log TacacsAccounting3Filter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_tacacsaccounting3_filter.labelname LogTacacsAccounting3Filter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

