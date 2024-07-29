---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_nacpolicy"
description: |-
  Configure NAC policy matching pattern to identify matching NAC devices.
---

# fmgdevice_user_nacpolicy
Configure NAC policy matching pattern to identify matching NAC devices.

## Example Usage

```hcl
resource "fmgdevice_user_nacpolicy" "trname" {
  category    = "vulnerability"
  description = "your own value"
  ems_tag     = ["your own value"]
  family      = "your own value"
  host        = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `category` - Category of NAC policy. Valid values: `device`, `firewall-user`, `ems-tag`, `vulnerability`, `fortivoice-tag`.

* `description` - Description for the NAC policy matching pattern.
* `ems_tag` - NAC policy matching EMS tag.
* `family` - NAC policy matching family.
* `host` - NAC policy matching host.
* `hw_vendor` - NAC policy matching hardware vendor.
* `hw_version` - NAC policy matching hardware version.
* `mac` - NAC policy matching MAC address.
* `name` - NAC policy name.
* `os` - NAC policy matching operating system.
* `src` - NAC policy matching source.
* `status` - Enable/disable NAC policy. Valid values: `disable`, `enable`.

* `sw_version` - NAC policy matching software version.
* `switch_auto_auth` - NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`, `global`.

* `switch_fortilink` - FortiLink interface for which this NAC policy belongs to.
* `switch_mac_policy` - switch-mac-policy to be applied on the matched NAC policy.
* `switch_port_policy` - switch-port-policy to be applied on the matched NAC policy.
* `switch_scope` - List of managed FortiSwitches on which NAC policy can be applied.
* `type` - NAC policy matching type.
* `user` - NAC policy matching user.
* `user_group` - NAC policy matching user group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User NacPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_nacpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

