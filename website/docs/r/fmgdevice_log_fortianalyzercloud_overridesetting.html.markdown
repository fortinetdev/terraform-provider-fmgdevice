---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortianalyzercloud_overridesetting"
description: |-
  Override FortiAnalyzer Cloud settings.
---

# fmgdevice_log_fortianalyzercloud_overridesetting
Override FortiAnalyzer Cloud settings.

## Example Usage

```hcl
resource "fmgdevice_log_fortianalyzercloud_overridesetting" "trname" {
  override    = "disable"
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `override` - Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `disable`, `enable`.

* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log FortianalyzerCloudOverrideSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortianalyzercloud_overridesetting.labelname LogFortianalyzerCloudOverrideSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

