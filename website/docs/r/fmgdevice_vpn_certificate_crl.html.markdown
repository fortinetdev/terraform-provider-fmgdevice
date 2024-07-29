---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_crl"
description: |-
  Certificate Revocation List as a PEM file.
---

# fmgdevice_vpn_certificate_crl
Certificate Revocation List as a PEM file.

## Example Usage

```hcl
resource "fmgdevice_vpn_certificate_crl" "trname" {
  crl           = "your own value"
  http_url      = "your own value"
  last_updated  = 10
  ldap_password = ["your own value"]
  ldap_server   = ["your own value"]
  name          = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `crl` - Certificate Revocation List as a PEM file.
* `http_url` - HTTP server URL for CRL auto-update.
* `last_updated` - Time at which CRL was last updated.
* `ldap_password` - LDAP server user password.
* `ldap_server` - LDAP server name for CRL auto-update.
* `ldap_username` - LDAP server user name.
* `name` - Name.
* `range` - Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.

* `scep_cert` - Local certificate for SCEP communication for CRL auto-update.
* `scep_url` - SCEP server URL for CRL auto-update.
* `source` - Certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.

* `source_ip` - Source IP address for communications to a HTTP or SCEP CA server.
* `update_interval` - Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
* `update_vdom` - VDOM for CRL update.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateCrl can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_crl.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

