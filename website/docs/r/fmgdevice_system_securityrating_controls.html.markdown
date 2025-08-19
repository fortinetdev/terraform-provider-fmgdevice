---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_securityrating_controls"
description: |-
  Settings for individual Security Rating controls.
---

# fmgdevice_system_securityrating_controls
Settings for individual Security Rating controls.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `display_insight` - Enable/disable displaying the Security Rating control as an insight across the GUI. Valid values: `disable`, `enable`.

* `display_report` - Enable/disable displaying the Security Rating control in the default report. Valid values: `disable`, `enable`.

* `name` - Security Rating control name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SecurityRatingControls can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_securityrating_controls.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

