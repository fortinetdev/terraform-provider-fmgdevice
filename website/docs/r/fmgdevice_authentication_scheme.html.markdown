---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_authentication_scheme"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Authentication Schemes.
---

# fmgdevice_authentication_scheme
<i>This object will be purged after policy copy and install.</i> Configure Authentication Schemes.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `digest_algo` - Digest Authentication Algorithms. Valid values: `md5`, `sha-256`.

* `digest_rfc2069` - Enable/disable support for the deprecated RFC2069 Digest Client (no cnonce field, default = disable). Valid values: `disable`, `enable`.

* `domain_controller` - Domain controller setting.
* `ems_device_owner` - Ems-Device-Owner. Valid values: `disable`, `enable`.

* `external_idp` - External identity provider configuration.
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable). Valid values: `disable`, `enable`.

* `group_attr_type` - Group attribute type used to match SCIM groups (default = display-name). Valid values: `display-name`, `external-id`.

* `kerberos_keytab` - Kerberos keytab setting.
* `method` - Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`, `saml`, `cert`, `entra-sso`.

* `name` - Authentication scheme name.
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `disable`, `enable`.

* `require_tfa` - Enable/disable two-factor authentication (default = disable). Valid values: `disable`, `enable`.

* `saml_server` - SAML configuration.
* `saml_timeout` - SAML authentication timeout in seconds.
* `ssh_ca` - SSH CA name.
* `user_cert` - Enable/disable authentication with user certificate (default = disable). Valid values: `disable`, `enable`.

* `user_database` - Authentication server to contain user information; "local-user-db" (default) or "123" (for LDAP).
* `oidc_server` - Oidc-Server.
* `oidc_timeout` - Oidc-Timeout.
* `search_all_ldap_databases` - Search-All-Ldap-Databases. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Authentication Scheme can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_authentication_scheme.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

