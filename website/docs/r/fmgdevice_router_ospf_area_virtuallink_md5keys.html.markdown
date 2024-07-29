---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_area_virtuallink_md5keys"
description: |-
  MD5 key.
---

# fmgdevice_router_ospf_area_virtuallink_md5keys
MD5 key.

~> This resource is a sub resource for variable `md5_keys` of resource `fmgdevice_router_ospf_area_virtuallink`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf_area_virtuallink_md5keys" "trname" {
  fosid       = 10
  key_string  = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `area` - Area.
* `virtual_link` - Virtual Link.

* `fosid` - Key ID (1 - 255).
* `key_string` - Password for the key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router OspfAreaVirtualLinkMd5Keys can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "area=YOUR_VALUE", "virtual_link=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_area_virtuallink_md5keys.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

