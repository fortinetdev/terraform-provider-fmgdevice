---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_local"
description: |-
  Local keys and certificates.
---

# fmgdevice_vpn_certificate_local
Local keys and certificates.

## Example Usage

```hcl
resource "fmgdevice_vpn_certificate_local" "trname" {
  acme_ca_url       = "your own value"
  acme_domain       = "your own value"
  acme_email        = "your own value"
  acme_renew_window = 10
  acme_rsa_key_size = 10
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `acme_ca_url` - The URL for the ACME CA server (Let's Encrypt is the default provider).
* `acme_domain` - A valid domain that resolves to this FortiGate unit.
* `acme_email` - Contact email address that is required by some CAs like LetsEncrypt.
* `acme_renew_window` - Beginning of the renewal window (in days before certificate expiration, 30 by default).
* `acme_rsa_key_size` - Length of the RSA private key of the generated cert (Minimum 2048 bits).
* `auto_regenerate_days` - Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
* `auto_regenerate_days_warning` - Number of days to wait before an expiry warning message is generated (0 = disabled).
* `ca_identifier` - CA identifier of the CA server for signing via SCEP.
* `certificate` - PEM format certificate.
* `cmp_path` - Path location inside CMP server.
* `cmp_regeneration_method` - CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.

* `cmp_server` - Address and port for CMP server (format = address:port).
* `cmp_server_cert` - CMP server certificate.
* `comments` - Comment.
* `csr` - Certificate Signing Request.
* `enroll_protocol` - Certificate enrollment protocol. Valid values: `none`, `scep`, `cmpv2`, `acme2`, `est`.

* `est_ca_id` - CA identifier of the CA server for signing via EST.
* `est_client_cert` - Certificate used to authenticate this FortiGate to EST server.
* `est_http_password` - HTTP Authentication password for signing via EST.
* `est_http_username` - HTTP Authentication username for signing via EST.
* `est_regeneration_method` - EST behavioral options during re-enrollment. Valid values: `create-new-key`, `use-existing-key`.

* `est_server` - Address and port for EST server (e.g. https://example.com:1234).
* `est_server_cert` - EST server's certificate must be verifiable by this certificate to be authenticated.
* `est_srp_password` - EST SRP authentication password.
* `est_srp_username` - EST SRP authentication username.
* `ike_localid` - Local ID the FortiGate uses for authentication as a VPN client.
* `ike_localid_type` - IKE local ID type. Valid values: `fqdn`, `asn1dn`.

* `last_updated` - Time at which certificate was last updated.
* `name` - Name.
* `name_encoding` - Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.

* `password` - Password as a PEM file.
* `private_key` - PEM format key encrypted with a password.
* `private_key_retain` - Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `disable`, `enable`.

* `range` - Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.

* `scep_password` - SCEP server challenge password for auto-regeneration.
* `scep_url` - SCEP server URL.
* `source` - Certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.

* `source_ip` - Source IP address for communications to the SCEP server.
* `state` - State.
* `tmp_cert_file` - Temporary certificate file.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateLocal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_local.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

