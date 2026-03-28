---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_implicitproxyrule"
description: |-
  Device WebProxyImplicitProxyRule
---

# fmgdevice_webproxy_implicitproxyrule
Device WebProxyImplicitProxyRule

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `proxy_servers`: `fmgdevice_webproxy_implicitproxyrule_proxyservers`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bypass_list` - Bypass-List.
* `load_balance_method` - Load-Balance-Method. Valid values: `static`, `round-robin`, `weighted`, `first-alive`.

* `name` - Name.
* `proxy_servers` - Proxy-Servers. The structure of `proxy_servers` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `proxy_servers` block supports:

* `id` - Id.
* `ip` - Ip.
* `port` - Port.
* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ImplicitProxyRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_implicitproxyrule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

