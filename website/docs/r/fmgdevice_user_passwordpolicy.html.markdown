---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_passwordpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure user password policy.
---

# fmgdevice_user_passwordpolicy
<i>This object will be purged after policy copy and install.</i> Configure user password policy.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `expire_days` - Time in days before the user's password expires.
* `expire_status` - Enable/disable password expiration. Valid values: `disable`, `enable`.

* `expired_password_renewal` - Enable/disable renewal of a password that already is expired. Valid values: `disable`, `enable`.

* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `name` - Password policy name.
* `reuse_password` - Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `disable`, `enable`.

* `reuse_password_limit` - Number of times passwords can be reused (0 - 20, default = 0. If set to 0, can reuse password an unlimited number of times.).
* `warn_days` - Time in days before a password expiration warning message is displayed to the user upon login.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User PasswordPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_passwordpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

