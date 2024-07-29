---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_proxyarp"
description: |-
  Configure proxy-ARP.
---

# fmgdevice_system_proxyarp
Configure proxy-ARP.

## Example Usage

```hcl
resource "fmgdevice_system_proxyarp" "trname" {
  end_ip      = "your own value"
  fosid       = 10
  interface   = ["port2"]
  ip          = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `end_ip` - End IP of IP range to be proxied.
* `fosid` - Unique integer ID of the entry.
* `interface` - Interface acting proxy-ARP.
* `ip` - IP address or start IP to be proxied.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System ProxyArp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_proxyarp.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

