---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_fortiguard"
description: |-
  Device EmailfilterFortiguard
---

# fmgdevice_emailfilter_fortiguard
Device EmailfilterFortiguard

## Example Usage

```hcl
resource "fmgdevice_emailfilter_fortiguard" "trname" {
  spam_submit_force   = "enable"
  spam_submit_srv     = "your own value"
  spam_submit_txt2htm = "enable"
  device_name         = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `spam_submit_force` - Spam-Submit-Force. Valid values: `disable`, `enable`.

* `spam_submit_srv` - Spam-Submit-Srv.
* `spam_submit_txt2htm` - Spam-Submit-Txt2Htm. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter Fortiguard can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_fortiguard.labelname EmailfilterFortiguard
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

