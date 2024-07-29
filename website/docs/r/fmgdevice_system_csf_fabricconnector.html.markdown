---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf_fabricconnector"
description: |-
  Fabric connector configuration.
---

# fmgdevice_system_csf_fabricconnector
Fabric connector configuration.

~> This resource is a sub resource for variable `fabric_connector` of resource `fmgdevice_system_csf`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_csf_fabricconnector" "trname" {
  accprofile                 = ["your own value"]
  configuration_write_access = "disable"
  serial                     = "your own value"
  vdom                       = ["your own value"]
  device_name                = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration. Valid values: `disable`, `enable`.

* `serial` - Serial.
* `vdom` - Virtual domains that the connector has access to. If none are set, the connector will only have access to the VDOM that it joins the Security Fabric through.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial}}.

## Import

System CsfFabricConnector can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf_fabricconnector.labelname {{serial}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

