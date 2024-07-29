---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ssoadmin"
description: |-
  Configure SSO admin users.
---

# fmgdevice_system_ssoadmin
Configure SSO admin users.

## Example Usage

```hcl
resource "fmgdevice_system_ssoadmin" "trname" {
  accprofile                           = ["your own value"]
  gui_default_dashboard_template       = "your own value"
  gui_ignore_invalid_signature_version = "your own value"
  gui_ignore_release_overview_version  = "your own value"
  name                                 = "your own value"
  device_name                          = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - SSO admin user access profile.
* `gui_default_dashboard_template` - The default dashboard template.
* `gui_ignore_invalid_signature_version` - FortiOS image build version to ignore invalid signature warning for.
* `gui_ignore_release_overview_version` - FortiOS version to ignore release overview prompt for.
* `name` - SSO admin name.
* `vdom` - Virtual domain(s) that the administrator can access.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SsoAdmin can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ssoadmin.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

