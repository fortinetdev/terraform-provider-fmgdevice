---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_report_setting"
description: |-
  Report setting configuration.
---

# fmgdevice_report_setting
Report setting configuration.

## Example Usage

```hcl
resource "fmgdevice_report_setting" "trname" {
  fortiview              = "disable"
  pdf_report             = "disable"
  report_source          = ["local-deny-traffic"]
  top_n                  = 10
  web_browsing_threshold = 10
  device_name            = var.device_name # not required if setting is at provider
  device_vdom            = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fortiview` - Enable/disable historical FortiView. Valid values: `disable`, `enable`.

* `pdf_report` - Enable/disable PDF report. Valid values: `disable`, `enable`.

* `report_source` - Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.

* `top_n` - Number of items to populate (1000 - 20000).
* `web_browsing_threshold` - Web browsing time calculation threshold (3 - 15 min).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Report Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_report_setting.labelname ReportSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

