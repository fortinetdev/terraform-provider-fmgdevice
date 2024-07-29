---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_utmprofile"
description: |-
  Configure UTM (Unified Threat Management) profile.
---

# fmgdevice_wirelesscontroller_utmprofile
Configure UTM (Unified Threat Management) profile.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_utmprofile" "trname" {
  antivirus_profile = ["your own value"]
  application_list  = ["your own value"]
  comment           = "your own value"
  ips_sensor        = ["your own value"]
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `antivirus_profile` - AntiVirus profile name.
* `application_list` - Application control list name.
* `comment` - Comment.
* `ips_sensor` - IPS sensor name.
* `name` - UTM profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.

* `utm_log` - Enable/disable UTM logging. Valid values: `disable`, `enable`.

* `webfilter_profile` - WebFilter profile name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController UtmProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_utmprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

