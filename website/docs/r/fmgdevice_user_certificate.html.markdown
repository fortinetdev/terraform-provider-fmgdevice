---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_certificate"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure certificate users.
---

# fmgdevice_user_certificate
<i>This object will be purged after policy copy and install.</i> Configure certificate users.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `common_name` - Certificate common name.
* `fosid` - Id.
* `issuer` - CA certificate used for client certificate verification.
* `name` - User name.
* `status` - Enable/disable allowing the certificate user to authenticate with the FortiGate unit. Valid values: `disable`, `enable`.

* `type` - Type of certificate authentication method. Valid values: `single-certificate`, `trusted-issuer`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Certificate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_certificate.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

