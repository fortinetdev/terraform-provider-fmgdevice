---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_ocspserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> OCSP server configuration.
---

# fmgdevice_vpn_certificate_ocspserver
<i>This object will be purged after policy copy and install.</i> OCSP server configuration.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cert` - OCSP server certificate.
* `name` - OCSP server entry name.
* `secondary_cert` - Secondary OCSP server certificate.
* `secondary_url` - Secondary OCSP server URL.
* `source_ip` - Source IP address for dynamic AIA and OCSP queries.
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.

* `url` - OCSP server URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateOcspServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_ocspserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

