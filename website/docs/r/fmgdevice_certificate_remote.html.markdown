---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_certificate_remote"
description: |-
  Remote certificate as a PEM file.
---

# fmgdevice_certificate_remote
Remote certificate as a PEM file.

## Example Usage

```hcl
resource "fmgdevice_certificate_remote" "trname" {
  name        = "your own value"
  range       = "vdom"
  remote      = "your own value"
  source      = "bundle"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Name.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `vdom`, `global`.

* `remote` - Remote certificate.
* `source` - Remote certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Certificate Remote can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_certificate_remote.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

