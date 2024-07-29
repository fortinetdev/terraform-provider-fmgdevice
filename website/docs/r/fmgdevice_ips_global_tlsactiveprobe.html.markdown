---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_global_tlsactiveprobe"
description: |-
  TLS active probe configuration.
---

# fmgdevice_ips_global_tlsactiveprobe
TLS active probe configuration.

~> This resource is a sub resource for variable `tls_active_probe` of resource `fmgdevice_ips_global`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_ips_global_tlsactiveprobe" "trname" {
  interface               = ["port2"]
  interface_select_method = "auto"
  source_ip               = "your own value"
  source_ip6              = "your own value"
  vdom                    = ["your own value"]
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.
* `vdom` - Virtual domain name for TLS active probe.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips GlobalTlsActiveProbe can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_global_tlsactiveprobe.labelname IpsGlobalTlsActiveProbe
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

