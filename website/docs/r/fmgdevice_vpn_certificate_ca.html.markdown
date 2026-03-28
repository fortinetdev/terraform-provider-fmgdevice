---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_ca"
description: |-
  CA certificate.
---

# fmgdevice_vpn_certificate_ca
CA certificate.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_private_key` - _Private_Key.
* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `ca` - CA certificate as a PEM file.
* `ca_identifier` - CA identifier of the SCEP server.
* `est_url` - URL of the EST server.
* `fabric_ca` - Enable/disable synchronization of CA across Security Fabric. Valid values: `disable`, `enable`.

* `last_updated` - Time at which CA was last updated.
* `name` - Name.
* `non_fabric_name` - Name used prior to becoming a Security Fabric synchronized CA.
* `obsolete` - Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.

* `range` - Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.

* `scep_url` - URL of the SCEP server.
* `source` - CA certificate source type. Valid values: `factory`, `user`, `bundle`.

* `source_ip` - Source IP address for communications to the SCEP server.
* `ssl_inspection_trusted` - Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateCa can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_ca.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

