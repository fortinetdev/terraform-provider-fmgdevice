---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_fortigateprofile"
description: |-
  FortiGate connector profile configuration.
---

# fmgdevice_extensioncontroller_fortigateprofile
FortiGate connector profile configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `lan_extension`: `fmgdevice_extensioncontroller_fortigateprofile_lanextension`



## Example Usage

```hcl
resource "fmgdevice_extensioncontroller_fortigateprofile" "trname" {
  extension = "lan-extension"
  fosid     = 10
  lan_extension {
    backhaul_interface = ["port2"]
    backhaul_ip        = "your own value"
    ipsec_tunnel       = "your own value"
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `extension` - Extension. Valid values: `lan-extension`.

* `fosid` - ID.
* `lan_extension` - Lan-Extension. The structure of `lan_extension` block is documented below.
* `name` - FortiGate connector profile name.

The `lan_extension` block supports:

* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `ipsec_tunnel` - IPsec tunnel name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController FortigateProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_fortigateprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

