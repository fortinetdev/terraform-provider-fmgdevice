---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdwan_service_move"
description: |-
  Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.
---

# fmgdevice_system_sdwan_service_move
Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

## Example Usage

```hcl
resource "fmgdevice_system_sdwan_service_move" "trname" {
  target      = 11
  option      = "after"
  service     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
  depends_on  = [fmgdevice_system_sdwan_service.trname1, fmgdevice_system_sdwan_service.trname2]
}

resource "fmgdevice_system_sdwan_service" "trname2" {
  fosid         = 11
  priority_zone = ["virtual-wan-link"]
  dst           = ["none"]
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}

resource "fmgdevice_system_sdwan_service" "trname1" {
  fosid         = 10
  priority_zone = ["virtual-wan-link"]
  dst           = ["none"]
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `service` - Service.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the service changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of services.
