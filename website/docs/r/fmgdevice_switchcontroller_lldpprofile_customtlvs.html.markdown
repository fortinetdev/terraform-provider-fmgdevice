---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_lldpprofile_customtlvs"
description: |-
  Configuration method to edit custom TLV entries.
---

# fmgdevice_switchcontroller_lldpprofile_customtlvs
Configuration method to edit custom TLV entries.

~> This resource is a sub resource for variable `custom_tlvs` of resource `fmgdevice_switchcontroller_lldpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_lldpprofile_customtlvs" "trname" {
  information_string = "your own value"
  name               = "your own value"
  oui                = "your own value"
  subtype            = 10
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `lldp_profile` - Lldp Profile.

* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).
* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController LldpProfileCustomTlvs can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "lldp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_lldpprofile_customtlvs.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

