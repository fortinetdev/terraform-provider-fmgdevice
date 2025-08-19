---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_setting"
description: |-
  VPN certificate setting.
---

# fmgdevice_vpn_certificate_setting
VPN certificate setting.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `crl_verification`: `fmgdevice_vpn_certificate_setting_crlverification`



## Example Usage

```hcl
resource "fmgdevice_vpn_certificate_setting" "trname" {
  cert_expire_warning = 10
  certname_dsa1024    = ["your own value"]
  certname_dsa2048    = ["your own value"]
  certname_ecdsa256   = ["your own value"]
  certname_ecdsa384   = ["your own value"]
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cert_expire_warning` - Number of days before a certificate expires to send a warning. Set to 0 to disable sending of the warning (0 - 100, default = 14).
* `certname_dsa1024` - 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_dsa2048` - 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa256` - 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa384` - 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa521` - 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ed25519` - 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ed448` - 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa1024` - 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa2048` - 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa4096` - 4096 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `check_ca_cert` - Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable). Valid values: `disable`, `enable`.

* `check_ca_chain` - Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable). Valid values: `disable`, `enable`.

* `cmp_key_usage_checking` - Enable/disable server certificate key usage checking in CMP mode (default = enable). Valid values: `disable`, `enable`.

* `cmp_save_extra_certs` - Enable/disable saving extra certificates in CMP mode (default = disable). Valid values: `disable`, `enable`.

* `cn_allow_multi` - When searching for a matching certificate, allow multiple CN fields in certificate subject name (default = enable). Valid values: `disable`, `enable`.

* `cn_match` - When searching for a matching certificate, control how to do CN value matching with certificate subject name (default = substring). Valid values: `substring`, `value`.

* `crl_verification` - Crl-Verification. The structure of `crl_verification` block is documented below.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ocsp_default_server` - Default OCSP server.
* `ocsp_option` - Specify whether the OCSP URL is from certificate or configured OCSP server. Valid values: `certificate`, `server`.

* `ocsp_status` - Enable/disable receiving certificates using the OCSP. Valid values: `disable`, `enable`, `mandatory`.

* `proxy` - Proxy server FQDN or IP for OCSP/CA queries during certificate verification.
* `proxy_password` - Proxy server password.
* `proxy_port` - Proxy server port (1 - 65535, default = 8080).
* `proxy_username` - Proxy server user name.
* `source_ip` - Source IP address for dynamic AIA and OCSP queries.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `ssl_ocsp_source_ip` - Source IP address to use to communicate with the OCSP server.
* `strict_crl_check` - Strict-Crl-Check. Valid values: `disable`, `enable`.

* `strict_ocsp_check` - Enable/disable strict mode OCSP checking. Valid values: `disable`, `enable`.

* `subject_match` - When searching for a matching certificate, control how to do RDN value matching with certificate subject name (default = substring). Valid values: `substring`, `value`.

* `subject_set` - When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset). Valid values: `subset`, `superset`.

* `vrf_select` - VRF ID used for connection to server.

The `crl_verification` block supports:

* `chain_crl_absence` - CRL verification option when CRL of any certificate in chain is absent (default = ignore). Valid values: `ignore`, `revoke`.

* `expiry` - CRL verification option when CRL is expired (default = ignore). Valid values: `ignore`, `revoke`.

* `leaf_crl_absence` - CRL verification option when leaf CRL is absent (default = ignore). Valid values: `ignore`, `revoke`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn CertificateSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_setting.labelname VpnCertificateSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

