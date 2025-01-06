---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_passwordpolicy"
description: |-
  Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
---

# fmgdevice_system_passwordpolicy
Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

## Example Usage

```hcl
resource "fmgdevice_system_passwordpolicy" "trname" {
  apply_to              = ["ipsec-preshared-key"]
  change_4_characters   = "enable"
  expire_day            = 10
  expire_status         = "disable"
  min_change_characters = 10
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `apply_to` - Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.

* `change_4_characters` - Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `disable`, `enable`.

* `expire_day` - Number of days after which passwords expire (1 - 999 days, default = 90).
* `expire_status` - Enable/disable password expiration. Valid values: `disable`, `enable`.

* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `reuse_password` - Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `disable`, `enable`.

* `reuse_password_limit` - Number of times passwords can be reused (0 - 20, default = 0. If set to 0, can reuse password an unlimited number of times.).
* `status` - Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_passwordpolicy.labelname SystemPasswordPolicy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

