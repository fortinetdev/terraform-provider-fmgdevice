---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_scim"
description: |-
  Configure SCIM client entries.
---

# fmgdevice_user_scim
Configure SCIM client entries.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_method` - TLS client authentication methods (default = bearer token). Valid values: `token`, `base`.

* `base_url` - Server URL to receive SCIM create, read, update, delete (CRUD) requests.
* `certificate` - Certificate name.
* `client_authentication_method` - TLS client authentication methods (default = bearer token). Valid values: `token`, `base`.

* `client_identity_check` - Enable/disable client identity check. Valid values: `disable`, `enable`.

* `client_secret_token` - Client secret token for authentication.
* `fosid` - SCIM client ID.
* `name` - SCIM client name.
* `secret` - Secret for token verification or base authentication.
* `status` - Enable/disable System for Cross-domain Identity Management (SCIM). Valid values: `disable`, `enable`.

* `token_certificate` - Certificate for token verification.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Scim can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_scim.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

