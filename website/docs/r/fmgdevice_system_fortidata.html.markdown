---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fortidata"
description: |-
  Configure FortiData.
---

# fmgdevice_system_fortidata
Configure FortiData.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IPv4 address used to communicate with FortiData.
* `status` - Enable/disable FortiData. Valid values: `disable`, `enable`.

* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortidata can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fortidata.labelname SystemFortidata
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

