---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipsurlfilterdns6"
description: |-
  Configure IPS URL filter IPv6 DNS servers.
---

# fmgdevice_system_ipsurlfilterdns6
Configure IPS URL filter IPv6 DNS servers.

## Example Usage

```hcl
resource "fmgdevice_system_ipsurlfilterdns6" "trname" {
  address6    = "your own value"
  status      = "disable"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `address6` - IPv6 address of DNS server.
* `status` - Enable/disable this server for IPv6 DNS queries. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{address6}}.

## Import

System IpsUrlfilterDns6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipsurlfilterdns6.labelname {{address6}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

