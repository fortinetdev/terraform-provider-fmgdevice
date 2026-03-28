---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ztna_trafficforwardproxy_urlroute"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_ztna_trafficforwardproxy_urlroute
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `url_route` of resource `fmgdevice_ztna_trafficforwardproxy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `traffic_forward_proxy` - Traffic Forward Proxy.

* `name` - Name.
* `service_connector` - Service-Connector.
* `url_pattern` - Url-Pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna TrafficForwardProxyUrlRoute can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "traffic_forward_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ztna_trafficforwardproxy_urlroute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

