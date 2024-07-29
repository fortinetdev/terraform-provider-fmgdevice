---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_userbookmark_bookmarks"
description: |-
  Bookmark table.
---

# fmgdevice_vpn_ssl_web_userbookmark_bookmarks
Bookmark table.

~> This resource is a sub resource for variable `bookmarks` of resource `fmgdevice_vpn_ssl_web_userbookmark`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `form_data`: `fmgdevice_vpn_ssl_web_userbookmark_bookmarks_formdata`



## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_web_userbookmark_bookmarks" "trname" {
  additional_params = "your own value"
  apptype           = "portforward"
  color_depth       = "32"
  description       = "your own value"
  domain            = "your own value"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_bookmark` - User Bookmark.

* `additional_params` - Additional parameters.
* `apptype` - Application type. Valid values: `web`, `ftp`, `telnet`, `smb`, `vnc`, `rdp`, `ssh`, `citrix`, `rdpnative`, `portforward`, `sftp`.

* `color_depth` - Color depth per pixel. Valid values: `8`, `16`, `32`.

* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `height` - Screen height (range from 0 - 65535, default = 0).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout. Valid values: `ar`, `da`, `de`, `de-ch`, `en-gb`, `en-uk`, `en-us`, `es`, `fi`, `fr`, `fr-be`, `fr-ca`, `fr-ch`, `hr`, `hu`, `it`, `ja`, `lt`, `lv`, `mk`, `no`, `pl`, `pt`, `pt-br`, `ru`, `sl`, `sv`, `tk`, `tr`, `fr-ca-m`, `wg`, `ar-101`, `ar-102`, `ar-102-azerty`, `can-mul`, `cz`, `cz-qwerty`, `cz-pr`, `nl`, `de-ibm`, `en-uk-ext`, `en-us-dvorak`, `es-var`, `fi-sami`, `hu-101`, `it-142`, `ko`, `lt-ibm`, `lt-std`, `lav-std`, `lav-leg`, `mk-std`, `no-sami`, `pol-214`, `pol-pr`, `pt-br-abnt2`, `ru-mne`, `ru-t`, `sv-sami`, `tuk`, `tur-f`, `tur-q`, `zh-sym-sg-us`, `zh-sym-us`, `zh-tr-hk`, `zh-tr-mo`, `zh-tr-us`, `fr-apple`, `la-am`, `ja-106`.

* `listening_port` - Listening-Port.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `remote_port` - Remote-Port.
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `disable`, `enable`.

* `security` - Security mode for RDP connection (default = any). Valid values: `rdp`, `nla`, `tls`, `any`.

* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `disable`, `enable`.

* `server_layout` - Server-Layout. Valid values: `en-us-qwerty`, `de-de-qwertz`, `fr-fr-azerty`, `it-it-qwerty`, `sv-se-qwerty`, `failsafe`, `en-gb-qwerty`, `es-es-qwerty`, `fr-ch-qwertz`, `ja-jp-qwerty`, `pt-br-qwerty`, `tr-tr-qwerty`, `fr-ca-qwerty`.

* `show_status_window` - Show-Status-Window. Valid values: `disable`, `enable`.

* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `disable`, `enable`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - URL parameter.
* `vnc_keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `es`, `fi`, `fr`, `fr-be`, `it`, `no`, `pt`, `sv`, `nl`, `en-uk-ext`, `it-142`, `pt-br-abnt2`, `default`, `fr-ca-mul`, `gd`, `us-intl`.

* `width` - Screen width (range from 0 - 65535, default = 0).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn SslWebUserBookmarkBookmarks can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_bookmark=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_userbookmark_bookmarks.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

