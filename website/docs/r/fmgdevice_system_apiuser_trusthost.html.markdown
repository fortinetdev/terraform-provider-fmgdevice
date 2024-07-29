---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_apiuser_trusthost"
description: |-
  Trusthost.
---

# fmgdevice_system_apiuser_trusthost
Trusthost.

~> This resource is a sub resource for variable `trusthost` of resource `fmgdevice_system_apiuser`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_apiuser_trusthost" "trname" {
  fosid          = 10
  ipv4_trusthost = ["your own value"]
  ipv6_trusthost = "your own value"
  type           = "ipv6-trusthost"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `api_user` - Api User.

* `fosid` - ID.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.
* `type` - Trusthost type. Valid values: `ipv4-trusthost`, `ipv6-trusthost`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System ApiUserTrusthost can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "api_user=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_apiuser_trusthost.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

