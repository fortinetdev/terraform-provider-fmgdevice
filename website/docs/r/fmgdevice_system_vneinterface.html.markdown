---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vneinterface"
description: |-
  Configure virtual network enabler tunnels.
---

# fmgdevice_system_vneinterface
Configure virtual network enabler tunnels.

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

* `name` - VNE tunnel name.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `update_url` - URL of provisioning server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VneInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vneinterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

