---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_trafficsniffer_targetport"
description: |-
  Sniffer ports to filter.
---

# fmgdevice_switchcontroller_trafficsniffer_targetport
Sniffer ports to filter.

~> This resource is a sub resource for variable `target_port` of resource `fmgdevice_switchcontroller_trafficsniffer`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_trafficsniffer_targetport" "trname" {
  description = "your own value"
  in_ports    = ["your own value"]
  out_ports   = ["your own value"]
  switch_id   = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description for the sniffer port entry.
* `in_ports` - Configure source ingress port interfaces.
* `out_ports` - Configure source egress port interfaces.
* `switch_id` - Managed-switch ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

SwitchController TrafficSnifferTargetPort can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_trafficsniffer_targetport.labelname {{switch_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

