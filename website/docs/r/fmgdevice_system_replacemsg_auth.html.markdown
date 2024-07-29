---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_replacemsg_auth"
description: |-
  Replacement messages.
---

# fmgdevice_system_replacemsg_auth
Replacement messages.

## Example Usage

```hcl
resource "fmgdevice_system_replacemsg_auth" "trname" {
  buffer      = "your own value"
  format      = "html"
  header      = "none"
  msg_type    = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

System ReplacemsgAuth can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_replacemsg_auth.labelname {{msg_type}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

