---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_initialconfig_template"
description: |-
  Configure template for auto-generated VLANs.
---

# fmgdevice_switchcontroller_initialconfig_template
Configure template for auto-generated VLANs.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_initialconfig_template" "trname" {
  allowaccess = ["ftm"]
  auto_ip     = "enable"
  dhcp_server = "disable"
  ip          = ["your own value"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allowaccess` - Permitted types of management access to this interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `dnp`, `ftm`, `fabric`.

* `auto_ip` - Automatically allocate interface address and subnet block. Valid values: `disable`, `enable`.

* `dhcp_server` - Enable/disable a DHCP server on this interface. Valid values: `disable`, `enable`.

* `ip` - Interface IPv4 address and subnet mask.
* `name` - Initial config template name.
* `vlanid` - Unique VLAN ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController InitialConfigTemplate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_initialconfig_template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

