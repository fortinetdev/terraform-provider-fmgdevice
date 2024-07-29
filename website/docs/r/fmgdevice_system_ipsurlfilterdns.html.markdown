---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipsurlfilterdns"
description: |-
  Configure IPS URL filter DNS servers.
---

# fmgdevice_system_ipsurlfilterdns
Configure IPS URL filter DNS servers.

## Example Usage

```hcl
resource "fmgdevice_system_ipsurlfilterdns" "trname" {
  address         = "your own value"
  ipv6_capability = "enable"
  status          = "disable"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `address` - DNS server IP address.
* `ipv6_capability` - Enable/disable this server for IPv6 queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{address}}.

## Import

System IpsUrlfilterDns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipsurlfilterdns.labelname {{address}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

