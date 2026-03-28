---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_peer"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure WAN optimization peers.
---

# fmgdevice_wanopt_peer
<i>This object will be purged after policy copy and install.</i> Configure WAN optimization peers.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ip` - Peer IP address.
* `peer_host_id` - Peer host ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{peer_host_id}}.

## Import

Wanopt Peer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_peer.labelname {{peer_host_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

