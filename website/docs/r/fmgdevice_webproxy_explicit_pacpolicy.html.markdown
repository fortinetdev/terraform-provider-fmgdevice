---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_explicit_pacpolicy"
description: |-
  PAC policies.
---

# fmgdevice_webproxy_explicit_pacpolicy
PAC policies.

~> This resource is a sub resource for variable `pac_policy` of resource `fmgdevice_webproxy_explicit`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_webproxy_explicit_pacpolicy" "trname" {
  comments      = "your own value"
  dstaddr       = ["your own value"]
  pac_file_data = "your own value"
  pac_file_name = "your own value"
  policyid      = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Optional comments.
* `dstaddr` - Destination address objects.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_file_name` - Pac file name.
* `policyid` - Policy ID.
* `srcaddr` - Source address objects.
* `srcaddr6` - Source address6 objects.
* `status` - Enable/disable policy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

WebProxy ExplicitPacPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_explicit_pacpolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

