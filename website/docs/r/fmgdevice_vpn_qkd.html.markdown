---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_qkd"
description: |-
  Configure Quantum Key Distribution servers
---

# fmgdevice_vpn_qkd
Configure Quantum Key Distribution servers

## Example Usage

```hcl
resource "fmgdevice_vpn_qkd" "trname" {
  certificate = ["your own value"]
  comment     = "your own value"
  fosid       = "your own value"
  name        = "your own value"
  peer        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `certificate` - Names of up to 4 certificates to offer to the KME.
* `comment` - Comment.
* `fosid` - Quantum Key Distribution ID assigned by the KME.
* `name` - Quantum Key Distribution configuration name.
* `peer` - Authenticate Quantum Key Device's certificate with the peer/peergrp.
* `port` - Port to connect to on the KME.
* `server` - IPv4, IPv6 or DNS address of the KME.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn Qkd can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_qkd.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

