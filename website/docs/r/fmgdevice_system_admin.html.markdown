---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_admin"
description: |-
  Configure admin users.
---

# fmgdevice_system_admin
Configure admin users.

## Example Usage

```hcl
resource "fmgdevice_system_admin" "trname" {
  accprofile                 = ["your own value"]
  accprofile_override        = "disable"
  allow_remove_admin_session = "enable"
  comments                   = "your own value"
  email_to                   = "your own value"
  name                       = "your own value"
  device_name                = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `accprofile_override` - Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access. Valid values: `disable`, `enable`.

* `allow_remove_admin_session` - Enable/disable allow admin session to be removed by privileged admin users. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `email_to` - This administrator's email address.
* `force_password_change` - Enable/disable force password change on next login. Valid values: `disable`, `enable`.

* `fortitoken` - This administrator's FortiToken serial number.
* `guest_auth` - Enable/disable guest authentication. Valid values: `disable`, `enable`.

* `guest_lang` - Guest management portal language.
* `guest_usergroups` - Select guest user groups.
* `hidden` - Hidden.
* `gui_default_dashboard_template` - The default dashboard template.
* `gui_ignore_invalid_signature_version` - FortiOS image build version to ignore invalid signature warning for.
* `gui_ignore_release_overview_version` - FortiOS version to ignore release overview prompt for.
* `history0` - History0.
* `history1` - History1.
* `ip6_trusthost1` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost10` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost2` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost3` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost4` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost5` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost6` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost7` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost8` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost9` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `name` - User name.
* `old_password` - Admin user old password.
* `password` - Admin user password.
* `password_expire` - Password expire time.
* `peer_auth` - Set to enable peer certificate authentication (for HTTPS admin access). Valid values: `disable`, `enable`.

* `peer_group` - Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
* `radius_vdom_override` - Radius-Vdom-Override. Valid values: `disable`, `enable`.

* `remote_auth` - Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server. Valid values: `disable`, `enable`.

* `remote_group` - User group name used for remote auth.
* `schedule` - Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
* `sms_custom_server` - Custom SMS server to send SMS messages to.
* `sms_phone` - Phone number on which the administrator receives SMS messages.
* `sms_server` - Send SMS messages using the FortiGuard SMS server or a custom server. Valid values: `fortiguard`, `custom`.

* `ssh_certificate` - Select the certificate to be used by the FortiGate for authentication with an SSH client.
* `ssh_public_key1` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key2` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key3` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `trusthost1` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost10` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost2` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost3` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost4` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost5` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost6` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost7` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost8` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost9` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken`, `email`, `sms`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `vdom` - Virtual domain(s) that the administrator can access.
* `vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access. Valid values: `disable`, `enable`.

* `wildcard` - Enable/disable wildcard RADIUS authentication. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Admin can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_admin.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

