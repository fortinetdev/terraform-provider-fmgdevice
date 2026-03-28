---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_profile_nacquar"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure AntiVirus quarantine settings.
---

# fmgdevice_antivirus_profile_nacquar
<i>This object will be purged after policy copy and install.</i> Configure AntiVirus quarantine settings.

~> This resource is a sub resource for variable `nac_quar` of resource `fmgdevice_antivirus_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none`, `quar-src-ip`.

* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus ProfileNacQuar can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_profile_nacquar.labelname AntivirusProfileNacQuar
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

