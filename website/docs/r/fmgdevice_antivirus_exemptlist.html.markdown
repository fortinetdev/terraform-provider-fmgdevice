---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_exemptlist"
description: |-
  Configure a list of hashes to be exempt from AV scanning.
---

# fmgdevice_antivirus_exemptlist
Configure a list of hashes to be exempt from AV scanning.

## Example Usage

```hcl
resource "fmgdevice_antivirus_exemptlist" "trname" {
  comment     = "your own value"
  hash        = "your own value"
  hash_type   = "sha1"
  name        = "your own value"
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `hash` - Hash value to be matched.
* `hash_type` - Hash type. Valid values: `md5`, `sha1`, `sha256`.

* `name` - Table entry name.
* `status` - Enable/disable table entry. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Antivirus ExemptList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_exemptlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

