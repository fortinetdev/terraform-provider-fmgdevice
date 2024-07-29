---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_saml"
description: |-
  Global settings for SAML authentication.
---

# fmgdevice_system_saml
Global settings for SAML authentication.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `service_providers`: `fmgdevice_system_saml_serviceproviders`



## Example Usage

```hcl
resource "fmgdevice_system_saml" "trname" {
  artifact_resolution_url = "your own value"
  binding_protocol        = "redirect"
  cert                    = ["your own value"]
  default_login_page      = "sso"
  default_profile         = ["your own value"]
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `artifact_resolution_url` - SP artifact resolution URL.
* `binding_protocol` - IdP Binding protocol. Valid values: `post`, `redirect`.

* `cert` - Certificate to sign SAML messages.
* `default_login_page` - Choose default login page. Valid values: `normal`, `sso`.

* `default_profile` - Default profile for new SSO admin.
* `entity_id` - SP entity ID.
* `idp_artifact_resolution_url` - IDP artifact resolution URL.
* `idp_cert` - IDP certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `life` - Length of the range of time when the assertion is valid (in minutes).
* `portal_url` - Portal-Url.
* `role` - SAML role. Valid values: `IDP`, `SP`, `identity-provider`, `service-provider`.

* `server_address` - Server address.
* `service_providers` - Service-Providers. The structure of `service_providers` block is documented below.
* `single_logout_url` - Single-Logout-Url.
* `single_sign_on_url` - Single-Sign-On-Url.
* `status` - Enable/disable SAML authentication (default = disable). Valid values: `disable`, `enable`.

* `tolerance` - Tolerance to the range of time when the assertion is valid (in minutes).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `service_providers` block supports:

* `assertion_attributes` - Assertion-Attributes. The structure of `assertion_attributes` block is documented below.
* `idp_artifact_resolution_url` - IDP artifact resolution URL.
* `idp_entity_id` - Idp-Entity-Id.
* `idp_single_logout_url` - Idp-Single-Logout-Url.
* `idp_single_sign_on_url` - Idp-Single-Sign-On-Url.
* `name` - Name.
* `prefix` - Prefix.
* `sp_artifact_resolution_url` - SP artifact resolution URL.
* `sp_binding_protocol` - SP binding protocol. Valid values: `post`, `redirect`.

* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP entity ID.
* `sp_portal_url` - SP portal URL.
* `sp_single_logout_url` - SP single logout URL.
* `sp_single_sign_on_url` - SP single sign-on URL.

The `assertion_attributes` block supports:

* `name` - Name.
* `type` - Type. Valid values: `username`, `email`, `privilege`, `profile-name`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Saml can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_saml.labelname SystemSaml
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

