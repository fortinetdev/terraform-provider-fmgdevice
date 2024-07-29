---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_speedtestserver"
description: |-
  Configure speed test server list.
---

# fmgdevice_system_speedtestserver
Configure speed test server list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `host`: `fmgdevice_system_speedtestserver_host`



## Example Usage

```hcl
resource "fmgdevice_system_speedyour valueserver" "trname" {
  host {
    distance  = 10
    fosid     = 10
    ip        = "your own value"
    latitude  = "your own value"
    longitude = "your own value"
    password  = ["your own value"]
    port      = 10
    user      = "your own value"
  }

  name        = "your own value"
  timestamp   = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `host` - Host. The structure of `host` block is documented below.
* `name` - Speed test server name.
* `timestamp` - Speed test server timestamp.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `host` block supports:

* `distance` - Speed test host distance.
* `id` - Server host ID.
* `ip` - Server host IPv4 address.
* `latitude` - Speed test host latitude.
* `longitude` - Speed test host longitude.
* `password` - Speed test host password.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SpeedTestServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_speedtestserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

