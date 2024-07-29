---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_fortigateprofile_lanextension"
description: |-
  FortiGate connector LAN extension configuration.
---

# fmgdevice_extensioncontroller_fortigateprofile_lanextension
FortiGate connector LAN extension configuration.

~> This resource is a sub resource for variable `lan_extension` of resource `fmgdevice_extensioncontroller_fortigateprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_extensioncontroller_fortigateprofile_lanextension" "trname" {
  backhaul_interface = ["port2"]
  backhaul_ip        = "your own value"
  ipsec_tunnel       = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `fortigate_profile` - Fortigate Profile.

* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `ipsec_tunnel` - IPsec tunnel name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ExtensionController FortigateProfileLanExtension can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "fortigate_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_fortigateprofile_lanextension.labelname ExtensionControllerFortigateProfileLanExtension
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

