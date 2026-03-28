---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_implicitproxysetting"
description: |-
  Device WebProxyImplicitProxySetting
---

# fmgdevice_webproxy_implicitproxysetting
Device WebProxyImplicitProxySetting

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `implicit_proxy_rule` - Implicit-Proxy-Rule.
* `interface` - Interface.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ImplicitProxySetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_implicitproxysetting.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

