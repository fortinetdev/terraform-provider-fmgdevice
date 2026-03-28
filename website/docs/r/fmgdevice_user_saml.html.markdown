---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_saml"
description: |-
  <i>This object will be purged after policy copy and install.</i> SAML server entry configuration.
---

# fmgdevice_user_saml
<i>This object will be purged after policy copy and install.</i> SAML server entry configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fmgdevice_user_saml_dynamic_mapping`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `disable`, `enable`.

* `auth_url` - Auth-Url.
* `cert` - Certificate to sign SAML messages.
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `digest_method` - Digest method algorithm. Valid values: `sha1`, `sha256`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `entity_id` - SP entity ID.
* `group_claim_type` - Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `disable`, `enable`.

* `name` - SAML server entry name.
* `reauth` - Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `disable`, `enable`.

* `scim_client` - SCIM client name.
* `scim_group_attr_type` - Group attribute type used to match SCIM groups (default = display-name). Valid values: `display-name`, `external-id`.

* `scim_user_attr_type` - User attribute type used to match SCIM users (default = user-name). Valid values: `display-name`, `external-id`, `user-name`.

* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_claim_type` - User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `user_name` - User name in assertion statement.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `disable`, `enable`.

* `auth_url` - Auth-Url.
* `cert` - Certificate to sign SAML messages.
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `digest_method` - Digest method algorithm. Valid values: `sha1`, `sha256`.

* `entity_id` - SP entity ID.
* `group_claim_type` - Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `disable`, `enable`.

* `reauth` - Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `disable`, `enable`.

* `scim_client` - SCIM client name.
* `scim_group_attr_type` - Group attribute type used to match SCIM groups (default = display-name). Valid values: `display-name`, `external-id`.

* `scim_user_attr_type` - User attribute type used to match SCIM users (default = user-name). Valid values: `display-name`, `external-id`, `user-name`.

* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_claim_type` - User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `user_name` - User name in assertion statement.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Saml can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_saml.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

