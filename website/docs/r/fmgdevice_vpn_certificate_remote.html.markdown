---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_certificate_remote"
description: |-
  <i>This object will be purged after policy copy and install.</i> Remote certificate as a PEM file.
---

# fmgdevice_vpn_certificate_remote
<i>This object will be purged after policy copy and install.</i> Remote certificate as a PEM file.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Name.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.

* `remote` - Remote certificate.
* `source` - Remote certificate source type. Valid values: `factory`, `user`, `bundle`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn CertificateRemote can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_certificate_remote.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

