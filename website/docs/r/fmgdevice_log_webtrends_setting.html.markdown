---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_webtrends_setting"
description: |-
  Settings for WebTrends.
---

# fmgdevice_log_webtrends_setting
Settings for WebTrends.

## Example Usage

```hcl
resource "fmgdevice_log_webtrends_setting" "trname" {
  server      = "your own value"
  status      = "enable"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `server` - Address of the remote WebTrends server.
* `status` - Enable/disable logging to WebTrends. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log WebtrendsSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_webtrends_setting.labelname LogWebtrendsSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

