---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_apiuser"
description: |-
  Configure API users.
---

# fmgdevice_system_apiuser
Configure API users.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `trusthost`: `fmgdevice_system_apiuser_trusthost`



## Example Usage

```hcl
resource "fmgdevice_system_apiuser" "trname" {
  accprofile        = ["your own value"]
  api_key           = ["your own value"]
  comments          = "your own value"
  cors_allow_origin = "your own value"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - Admin user access profile.
* `api_key` - Admin user password.
* `comments` - Comment.
* `cors_allow_origin` - Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
* `name` - User name.
* `peer_auth` - Enable/disable peer authentication. Valid values: `disable`, `enable`.

* `peer_group` - Peer group name.
* `schedule` - Schedule name.
* `trusthost` - Trusthost. The structure of `trusthost` block is documented below.
* `vdom` - Virtual domains.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `trusthost` block supports:

* `id` - ID.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.
* `type` - Trusthost type. Valid values: `ipv4-trusthost`, `ipv6-trusthost`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ApiUser can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_apiuser.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

