---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_implicitproxyrule_proxyservers"
description: |-
  Device WebProxyImplicitProxyRuleProxyServers
---

# fmgdevice_webproxy_implicitproxyrule_proxyservers
Device WebProxyImplicitProxyRuleProxyServers

~> This resource is a sub resource for variable `proxy_servers` of resource `fmgdevice_webproxy_implicitproxyrule`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `implicit_proxy_rule` - Implicit Proxy Rule.

* `fosid` - Id.
* `ip` - Ip.
* `port` - Port.
* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WebProxy ImplicitProxyRuleProxyServers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "implicit_proxy_rule=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_implicitproxyrule_proxyservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

