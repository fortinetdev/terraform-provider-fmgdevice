---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_settings"
description: |-
  Configure WAN optimization settings.
---

# fmgdevice_wanopt_settings
Configure WAN optimization settings.

## Example Usage

```hcl
resource "fmgdevice_wanopt_settings" "trname" {
  auto_detect_algorithm = "diff-req-resp"
  host_id               = "your own value"
  tunnel_optimization   = "memory-usage"
  tunnel_ssl_algorithm  = "low"
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auto_detect_algorithm` - Auto detection algorithms used in tunnel negotiations. Valid values: `simple`, `diff-req-resp`.

* `host_id` - Local host ID (must also be entered in the remote FortiGate's peer list).
* `tunnel_optimization` - WANOpt tunnel optimization option. Valid values: `memory-usage`, `balanced`, `throughput`.

* `tunnel_ssl_algorithm` - Relative strength of encryption algorithms accepted during tunnel negotiation. Valid values: `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_settings.labelname WanoptSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

