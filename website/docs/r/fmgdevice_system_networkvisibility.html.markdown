---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_networkvisibility"
description: |-
  Configure network visibility settings.
---

# fmgdevice_system_networkvisibility
Configure network visibility settings.

## Example Usage

```hcl
resource "fmgdevice_system_networkvisibility" "trname" {
  destination_hostname_visibility = "enable"
  destination_location            = "disable"
  destination_visibility          = "disable"
  hostname_limit                  = 10
  hostname_ttl                    = 10
  device_name                     = var.device_name # not required if setting is at provider
  device_vdom                     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `destination_hostname_visibility` - Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.

* `destination_location` - Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.

* `destination_visibility` - Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.

* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `source_location` - Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NetworkVisibility can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_networkvisibility.labelname SystemNetworkVisibility
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

