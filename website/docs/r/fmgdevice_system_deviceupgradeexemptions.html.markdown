---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_deviceupgradeexemptions"
description: |-
  Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI.
---

# fmgdevice_system_deviceupgradeexemptions
Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fortinet_device` - Fortinet extension device (FortiAP, FortiSwitch, FortiExtender).
* `fosid` - Device upgrade exemption ID (0 - 65535).
* `version` - Highest version of Fortinet firmware to exempt (format in Major.minor.patch, such as 7.6.4).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DeviceUpgradeExemptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_deviceupgradeexemptions.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

