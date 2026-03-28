---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnproxy"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SDN proxy.
---

# fmgdevice_system_sdnproxy
<i>This object will be purged after policy copy and install.</i> Configure SDN proxy.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - SDN proxy name.
* `password` - SDN proxy password.
* `server` - Server address of the SDN proxy.
* `server_port` - Port number of the SDN proxy.
* `type` - Type of SDN proxy. Valid values: `general`, `fortimanager`.

* `username` - SDN proxy username.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnProxy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnproxy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

