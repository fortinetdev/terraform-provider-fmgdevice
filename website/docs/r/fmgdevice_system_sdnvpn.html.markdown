---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdnvpn"
description: |-
  Configure public cloud VPN service.
---

# fmgdevice_system_sdnvpn
Configure public cloud VPN service.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `bgp_as` - BGP Router AS number.
* `bgp_from` - BGP source. Valid values: `unset`, `config`, `cli`.

* `bgp_seq` - BGP sequence number.
* `cgw_gateway` - Public IP address of the customer gateway.
* `cgw_name` - AWS customer gateway name to be created.
* `cgw_id` - Customer gateway id.
* `code` - Code.
* `internal_interface` - Internal interface with local subnet.
* `local_cidr` - Local subnet address and subnet mask.
* `name` - Public cloud VPN name.
* `nat_traversal` - Enable/disable use for NAT traversal. Please enable if your FortiGate device is behind a NAT/PAT device. Valid values: `disable`, `enable`.

* `psksecret` - Pre-shared secret for PSK authentication. Auto-generated if not specified
* `remote_cidr` - Remote subnet address and subnet mask.
* `remote_type` - Type of remote device. Valid values: `vgw`, `tgw`.

* `routing_type` - Type of routing. Valid values: `static`, `dynamic`.

* `sdn` - SDN connector name.
* `status` - Status.
* `subnet_id` - AWS subnet id for TGW route propagation.
* `tgw_id` - Transit gateway id.
* `tgw_vpn_rtbl_id` - Transit gateway route table id.
* `trtbl_attachment` - Transit gateway route table attachment id.
* `tunnel_interface` - Tunnel interface with public IP.
* `type` - Type.
* `vgw_id` - Virtual private gateway id.
* `vpn_id` - VPN connection id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnVpn can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdnvpn.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

