---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_hsmlocal"
description: |-
  Local certificates whose keys are stored on HSM.
---

# fmgdevice_vpn_certificate_hsmlocal
Local certificates whose keys are stored on HSM.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `api_version` - API version for communicating with HSM. Valid values: `unknown`, `gch-default`.

* `certificate` - PEM format certificate.
* `comments` - Comment.
* `gch_cloud_service_name` - Cloud service config name to generate access token.
* `gch_cryptokey` - Google Cloud HSM cryptokey.
* `gch_cryptokey_algorithm` - Google Cloud HSM cryptokey algorithm. Valid values: `rsa-sign-pkcs1-2048-sha256`, `rsa-sign-pkcs1-3072-sha256`, `rsa-sign-pkcs1-4096-sha256`, `rsa-sign-pkcs1-4096-sha512`, `rsa-sign-pss-2048-sha256`, `rsa-sign-pss-3072-sha256`, `rsa-sign-pss-4096-sha256`, `rsa-sign-pss-4096-sha512`, `ec-sign-p256-sha256`, `ec-sign-p384-sha384`, `ec-sign-secp256k1-sha256`.

* `gch_cryptokey_version` - Google Cloud HSM cryptokey version.
* `gch_keyring` - Google Cloud HSM keyring.
* `gch_location` - Google Cloud HSM location.
* `gch_project` - Google Cloud HSM project ID.
* `gch_url` - Gch-Url.
* `name` - Name.
* `range` - Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.

* `source` - Certificate source type. Valid values: `factory`, `user`, `bundle`.

* `tmp_cert_file` - Temporary certificate file.
* `vendor` - HSM vendor. Valid values: `unknown`, `gch`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateHsmLocal can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_hsmlocal.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

