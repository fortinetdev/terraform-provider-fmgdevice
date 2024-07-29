---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ips"
description: |-
  Configure IPS system settings.
---

# fmgdevice_system_ips
Configure IPS system settings.

## Example Usage

```hcl
resource "fmgdevice_system_ips" "trname" {
  override_signature_hold_by_id = "disable"
  signature_hold_time           = "your own value"
  device_name                   = var.device_name # not required if setting is at provider
  device_vdom                   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `override_signature_hold_by_id` - Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `disable`, `enable`.

* `signature_hold_time` - Time to hold and monitor IPS signatures. Format &lt;#d##h&gt; (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ips can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ips.labelname SystemIps
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

