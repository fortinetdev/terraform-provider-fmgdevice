---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_speedtestserver_host"
description: |-
  Hosts of the server.
---

# fmgdevice_system_speedtestserver_host
Hosts of the server.

~> This resource is a sub resource for variable `host` of resource `fmgdevice_system_speedtestserver`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_speedyour valueserver_host" "trname" {
  distance    = 10
  fosid       = 10
  ip          = "your own value"
  latitude    = "your own value"
  longitude   = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `speed_test_server` - Speed Test Server.

* `distance` - Speed test host distance.
* `fosid` - Server host ID.
* `ip` - Server host IPv4 address.
* `latitude` - Speed test host latitude.
* `longitude` - Speed test host longitude.
* `password` - Speed test host password.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SpeedTestServerHost can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "speed_test_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_speedtestserver_host.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

