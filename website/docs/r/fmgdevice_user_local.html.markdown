---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_local"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure local users.
---

# fmgdevice_user_local
<i>This object will be purged after policy copy and install.</i> Configure local users.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `disable`, `enable`.

* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `email_to` - Two-factor recipient's email address.
* `fortitoken` - Two-factor recipient's FortiToken serial number.
* `fosid` - Id.
* `ldap_server` - Name of LDAP server with which the user must authenticate.
* `name` - Local user name.
* `passwd` - User's password.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy.
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `qkd_profile` - Quantum Key Distribution (QKD) profile.
* `radius_server` - Name of RADIUS server with which the user must authenticate.
* `saml_server` - Name of SAML server with which the user must authenticate.
* `sms_custom_server` - Two-factor recipient's SMS server.
* `sms_phone` - Two-factor recipient's mobile phone number.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.

* `status` - Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `disable`, `enable`.

* `tacacs_server` - Name of TACACS+ server with which the user must authenticate.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken`, `email`, `sms`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`, `saml`.

* `username_case_insensitivity` - Username-Case-Insensitivity. Valid values: `disable`, `enable`.

* `username_case_sensitivity` - Username-Case-Sensitivity. Valid values: `disable`, `enable`.

* `username_sensitivity` - Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.

* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
* `history0` - History0.
* `history1` - History1.
* `history10` - History10.
* `history11` - History11.
* `history12` - History12.
* `history13` - History13.
* `history14` - History14.
* `history15` - History15.
* `history16` - History16.
* `history17` - History17.
* `history18` - History18.
* `history19` - History19.
* `history2` - History2.
* `history3` - History3.
* `history4` - History4.
* `history5` - History5.
* `history6` - History6.
* `history7` - History7.
* `history8` - History8.
* `history9` - History9.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Local can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_local.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

