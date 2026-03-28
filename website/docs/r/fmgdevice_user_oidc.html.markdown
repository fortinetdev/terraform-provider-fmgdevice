---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_oidc"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_user_oidc
<i>This object will be purged after policy copy and install.</i>

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_method` - Auth-Method. Valid values: `client_secret_basic`, `client_secret_post`, `private_key_jwt`.

* `auth_type` - Auth-Type. Valid values: `client-secret`, `private-key`.

* `authorization_url` - Authorization-Url.
* `client_id` - Client-Id.
* `client_secret` - Client-Secret.
* `clock_tolerance` - Clock-Tolerance.
* `discovery_url` - Discovery-Url.
* `display_name` - Display-Name.
* `domain_hint` - Domain-Hint.
* `group_attr_name` - Group-Attr-Name.
* `icon_url` - Icon-Url.
* `issuer` - Issuer.
* `jwks_uri` - Jwks-Uri.
* `ldap_server` - Ldap-Server.
* `name` - Name.
* `private_key` - Private-Key.
* `token_url` - Token-Url.
* `type` - Type. Valid values: `discovery`, `manual`.

* `user_attr_name` - User-Attr-Name. Valid values: `email`, `sub`, `preferred_username`.

* `user_regex` - User-Regex.
* `verify_cert` - Verify-Cert. Valid values: `disable`, `enable`.

* `verify_issuer` - Verify-Issuer. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Oidc can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_oidc.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

