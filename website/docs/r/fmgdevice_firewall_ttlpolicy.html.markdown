---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ttlpolicy"
description: |-
  Configure TTL policies.
---

# fmgdevice_firewall_ttlpolicy
Configure TTL policies.

## Example Usage

```hcl
resource "fmgdevice_firewall_ttlpolicy" "trname" {
  action      = "deny"
  fosid       = 10
  schedule    = ["your own value"]
  service     = ["your own value"]
  srcaddr     = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Action to be performed on traffic matching this policy (default = deny). Valid values: `deny`, `accept`.

* `fosid` - ID.
* `schedule` - Schedule object from available options.
* `service` - Service object(s) from available options. Separate multiple names with a space.
* `srcaddr` - Source address object(s) from available options. Separate multiple names with a space.
* `srcintf` - Source interface name from available interfaces.
* `status` - Enable/disable this TTL policy. Valid values: `disable`, `enable`.

* `ttl` - Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall TtlPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ttlpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

