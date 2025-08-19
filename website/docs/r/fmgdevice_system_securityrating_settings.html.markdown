---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_securityrating_settings"
description: |-
  Settings for Security Rating.
---

# fmgdevice_system_securityrating_settings
Settings for Security Rating.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `override_sync` - Enable/disable overriding Security Rating control settings synced from the Security Fabric's root FortiGate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SecurityRatingSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_securityrating_settings.labelname SystemSecurityRatingSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

