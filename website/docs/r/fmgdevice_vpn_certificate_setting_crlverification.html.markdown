---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_setting_crlverification"
description: |-
  CRL verification options.
---

# fmgdevice_vpn_certificate_setting_crlverification
CRL verification options.

~> This resource is a sub resource for variable `crl_verification` of resource `fmgdevice_vpn_certificate_setting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_certificate_setting_crlverification" "trname" {
  chain_crl_absence = "revoke"
  expiry            = "revoke"
  leaf_crl_absence  = "revoke"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `chain_crl_absence` - CRL verification option when CRL of any certificate in chain is absent (default = ignore). Valid values: `ignore`, `revoke`.

* `expiry` - CRL verification option when CRL is expired (default = ignore). Valid values: `ignore`, `revoke`.

* `leaf_crl_absence` - CRL verification option when leaf CRL is absent (default = ignore). Valid values: `ignore`, `revoke`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn CertificateSettingCrlVerification can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_setting_crlverification.labelname VpnCertificateSettingCrlVerification
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

