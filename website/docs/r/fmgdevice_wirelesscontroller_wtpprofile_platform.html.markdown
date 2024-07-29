---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpprofile_platform"
description: |-
  WTP, FortiAP, or AP platform.
---

# fmgdevice_wirelesscontroller_wtpprofile_platform
WTP, FortiAP, or AP platform.

~> This resource is a sub resource for variable `platform` of resource `fmgdevice_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtpprofile_platform" "trname" {
  _local_platform_str = "your own value"
  ddscan              = "enable"
  mode                = "dual-5G"
  type                = "234G"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp_profile` - Wtp Profile.

* `_local_platform_str` - _Local_Platform_Str.
* `ddscan` - Enable/disable use of one radio for dedicated full-band scanning to detect RF characterization and wireless threat management. Valid values: `disable`, `enable`.

* `mode` - Configure operation mode of 5G radios (default = single-5G). Valid values: `dual-5G`, `single-5G`.

* `type` - WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile. Valid values: `30B-50B`, `60B`, `80CM-81CM`, `220A`, `220B`, `210B`, `60C`, `222B`, `112B`, `320B`, `11C`, `14C`, `223B`, `28C`, `320C`, `221C`, `25D`, `222C`, `224D`, `214B`, `21D`, `24D`, `112D`, `223C`, `321C`, `C220C`, `C225C`, `S321C`, `S323C`, `FWF`, `S311C`, `S313C`, `AP-11N`, `S322C`, `S321CR`, `S322CR`, `S323CR`, `S421E`, `S422E`, `S423E`, `421E`, `423E`, `C221E`, `C226E`, `C23JD`, `C24JE`, `C21D`, `U421E`, `U423E`, `221E`, `222E`, `223E`, `S221E`, `S223E`, `U221EV`, `U223EV`, `U321EV`, `U323EV`, `224E`, `U422EV`, `U24JEV`, `321E`, `U431F`, `U433F`, `231E`, `431F`, `433F`, `231F`, `432F`, `234F`, `23JF`, `U231F`, `831F`, `U234F`, `U432F`, `431FL`, `432FR`, `433FL`, `231FL`, `231G`, `233G`, `431G`, `433G`, `U231G`, `U441G`, `234G`, `432G`, `441K`, `443K`, `241K`, `243K`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController WtpProfilePlatform can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpprofile_platform.labelname WirelessControllerWtpProfilePlatform
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

