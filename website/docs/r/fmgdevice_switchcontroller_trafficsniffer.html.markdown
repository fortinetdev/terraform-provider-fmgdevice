---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_trafficsniffer"
description: |-
  Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.
---

# fmgdevice_switchcontroller_trafficsniffer
Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `target_ip`: `fmgdevice_switchcontroller_trafficsniffer_targetip`
>- `target_mac`: `fmgdevice_switchcontroller_trafficsniffer_targetmac`
>- `target_port`: `fmgdevice_switchcontroller_trafficsniffer_targetport`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_trafficsniffer" "trname" {
  erspan_ip = "your own value"
  mode      = "rspan"
  target_ip {
    description  = "your own value"
    dst_entry_id = 10
    ip           = "your own value"
    src_entry_id = 10
  }

  target_mac {
    description  = "your own value"
    dst_entry_id = 10
    mac          = "your own value"
    src_entry_id = 10
  }

  target_port {
    description = "your own value"
    in_ports    = ["your own value"]
    out_ports   = ["your own value"]
    switch_id   = ["your own value"]
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `erspan_ip` - Configure ERSPAN collector IP address.
* `mode` - Configure traffic sniffer mode. Valid values: `none`, `erspan-auto`, `rspan`.

* `target_ip` - Target-Ip. The structure of `target_ip` block is documented below.
* `target_mac` - Target-Mac. The structure of `target_mac` block is documented below.
* `target_port` - Target-Port. The structure of `target_port` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `target_ip` block supports:

* `description` - Description for the sniffer IP.
* `dst_entry_id` - FortiSwitch dest entry ID for the sniffer IP.
* `ip` - Sniffer IP.
* `src_entry_id` - FortiSwitch source entry ID for the sniffer IP.

The `target_mac` block supports:

* `description` - Description for the sniffer MAC.
* `dst_entry_id` - FortiSwitch dest entry ID for the sniffer MAC.
* `mac` - Sniffer MAC.
* `src_entry_id` - FortiSwitch source entry ID for the sniffer MAC.

The `target_port` block supports:

* `description` - Description for the sniffer port entry.
* `in_ports` - Configure source ingress port interfaces.
* `out_ports` - Configure source egress port interfaces.
* `switch_id` - Managed-switch ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController TrafficSniffer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_trafficsniffer.labelname SwitchControllerTrafficSniffer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

