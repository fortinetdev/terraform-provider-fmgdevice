---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_portal_landingpage"
description: |-
  <i>This object will be purged after policy copy and install.</i> Landing page options.
---

# fmgdevice_vpn_ssl_web_portal_landingpage
<i>This object will be purged after policy copy and install.</i> Landing page options.

~> This resource is a sub resource for variable `landing_page` of resource `fmgdevice_vpn_ssl_web_portal`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `form_data`: `fmgdevice_vpn_ssl_web_portal_landingpage_formdata`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `portal` - Portal.

* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - Landing page URL.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn SslWebPortalLandingPage can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "portal=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_portal_landingpage.labelname VpnSslWebPortalLandingPage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

