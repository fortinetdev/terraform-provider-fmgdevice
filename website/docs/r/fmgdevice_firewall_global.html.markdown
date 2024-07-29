---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_global"
description: |-
  Global firewall settings.
---

# fmgdevice_firewall_global
Global firewall settings.

## Example Usage

```hcl
resource "fmgdevice_firewall_global" "trname" {
  banned_ip_persistency = "permanent-only"
  device_name           = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `banned_ip_persistency` - Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_global.labelname FirewallGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

