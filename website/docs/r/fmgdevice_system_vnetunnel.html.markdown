---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vnetunnel"
description: |-
  Configure virtual network enabler tunnel.
---

# fmgdevice_system_vnetunnel
Configure virtual network enabler tunnel.

## Example Usage

```hcl
resource "fmgdevice_system_vnetunnel" "trname" {
  auto_asic_offload = "enable"
  bmr_hostname      = ["your own value"]
  br                = "your own value"
  http_password     = ["your own value"]
  http_username     = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `disable`, `enable`.

* `bmr_hostname` - BMR hostname.
* `br` - IPv6 address or FQDN of the border relay.
* `http_password` - HTTP authentication password.
* `http_username` - HTTP authentication user name.
* `interface` - Interface name.
* `ipv4_address` - Tunnel IPv4 address and netmask.
* `mode` - VNE tunnel mode. Valid values: `map-e`, `fixed-ip`, `ds-lite`.

* `ssl_certificate` - Name of local certificate for SSL connections.
* `status` - Enable/disable VNE tunnel. Valid values: `disable`, `enable`.

* `update_url` - URL of provisioning server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VneTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vnetunnel.labelname SystemVneTunnel
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

