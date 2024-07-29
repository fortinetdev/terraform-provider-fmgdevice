---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_settings"
description: |-
  Configure AntiVirus settings.
---

# fmgdevice_antivirus_settings
Configure AntiVirus settings.

## Example Usage

```hcl
resource "fmgdevice_antivirus_settings" "trname" {
  default_db                 = "flow-based"
  cache_clean_result         = "disable"
  cache_infected_result      = "enable"
  grayware                   = "disable"
  machine_learning_detection = "enable"
  device_name                = var.device_name # not required if setting is at provider
  device_vdom                = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_db` - Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `flow-based`, `extreme`.

* `cache_clean_result` - Cache-Clean-Result. Valid values: `disable`, `enable`.

* `cache_infected_result` - Enable/disable cache of infected scan results (default = enable). Valid values: `disable`, `enable`.

* `grayware` - Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `disable`, `enable`.

* `machine_learning_detection` - Use machine learning based malware detection. Valid values: `disable`, `enable`, `monitor`.

* `override_timeout` - Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
* `use_extreme_db` - Enable/disable the use of Extreme AVDB. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_settings.labelname AntivirusSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

