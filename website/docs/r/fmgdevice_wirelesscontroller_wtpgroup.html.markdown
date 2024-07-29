---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpgroup"
description: |-
  Configure WTP groups.
---

# fmgdevice_wirelesscontroller_wtpgroup
Configure WTP groups.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtpgroup" "trname" {
  ble_major_id  = 10
  name          = "your own value"
  platform_type = "24D"
  wtps          = ["your own value"]
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ble_major_id` - Override BLE Major ID.
* `name` - WTP group name.
* `platform_type` - FortiAP models to define the WTP group platform type. Valid values: `220B`, `210B`, `222B`, `112B`, `320B`, `11C`, `14C`, `223B`, `28C`, `320C`, `221C`, `25D`, `222C`, `224D`, `214B`, `21D`, `24D`, `112D`, `223C`, `321C`, `C220C`, `C225C`, `S321C`, `S323C`, `FWF`, `S311C`, `S313C`, `AP-11N`, `S322C`, `S321CR`, `S322CR`, `S323CR`, `S421E`, `S422E`, `S423E`, `421E`, `423E`, `C221E`, `C226E`, `C23JD`, `C24JE`, `C21D`, `U421E`, `U423E`, `221E`, `222E`, `223E`, `S221E`, `S223E`, `U221EV`, `U223EV`, `U321EV`, `U323EV`, `224E`, `U422EV`, `U24JEV`, `321E`, `U431F`, `U433F`, `231E`, `431F`, `433F`, `231F`, `432F`, `234F`, `23JF`, `U231F`, `831F`, `U234F`, `U432F`, `431FL`, `432FR`, `433FL`, `231FL`, `231G`, `233G`, `431G`, `433G`, `U231G`, `U441G`, `234G`, `432G`, `441K`, `443K`, `241K`, `243K`.

* `wtps` - WTP list.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WtpGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpgroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

