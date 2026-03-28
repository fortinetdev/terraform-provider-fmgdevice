---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_peer"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure peer users.
---

# fmgdevice_user_peer
<i>This object will be purged after policy copy and install.</i> Configure peer users.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ca` - Name of the CA certificate.
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.

* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `disable`, `enable`.

* `mfa_mode` - MFA mode for remote peer authentication/authorization. Valid values: `none`, `password`, `subject-identity`.

* `mfa_password` - Unified password for remote authentication. This field may be left empty when RADIUS authentication is used, in which case the FortiGate will use the RADIUS username as a password.
* `mfa_server` - Name of a remote authenticator. Performs client access right check.
* `mfa_username` - Unified username for remote authentication.
* `name` - Peer name.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval.
* `passwd` - Peer's password used for two-factor authentication.
* `subject` - Peer certificate name constraints.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `disable`, `enable`.

* `ldap_mode` - Ldap-Mode. Valid values: `password`, `principal-name`.

* `ldap_password` - Ldap-Password.
* `ldap_server` - Ldap-Server.
* `ldap_username` - Ldap-Username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_peer.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

