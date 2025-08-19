---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationcondition"
description: |-
  Condition for automation stitches.
---

# fmgdevice_system_automationcondition
Condition for automation stitches.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `condition_type` - Condition type. Valid values: `cpu`, `memory`, `vpn`, `input`.

* `cpu_usage_percent` - CPU usage reaches specified percentage.
* `description` - Description.
* `input_id` - Input ID.
* `input_state` - Input state. Valid values: `close`, `open`.

* `mem_usage_percent` - Memory usage reaches specified percentage.
* `name` - Name.
* `vdom` - Virtual domain which the tunnel belongs to.
* `vpn_tunnel_name` - VPN tunnel name.
* `vpn_tunnel_state` - VPN tunnel state. Valid values: `tunnel-up`, `tunnel-down`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationCondition can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationcondition.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

