---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dnsserver"
description: |-
  Configure DNS servers.
---

# fmgdevice_system_dnsserver
Configure DNS servers.

## Example Usage

```hcl
resource "fmgdevice_system_dnsserver" "trname" {
  dnsfilter_profile = ["your own value"]
  doh               = "disable"
  doh3              = "enable"
  doq               = "enable"
  mode              = "recursive"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dnsfilter_profile` - DNS filter profile.
* `doh` - Enable/disable DNS over HTTPS/443 (default = disable). Valid values: `disable`, `enable`.

* `doh3` - Enable/disable DNS over QUIC/HTTP3/443 (default = disable). Valid values: `disable`, `enable`.

* `doq` - Enable/disable DNS over QUIC/853 (default = disable). Valid values: `disable`, `enable`.

* `mode` - DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.

* `name` - DNS server name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dnsserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

