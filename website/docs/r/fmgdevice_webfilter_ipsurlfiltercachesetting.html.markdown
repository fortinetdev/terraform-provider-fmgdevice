---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_ipsurlfiltercachesetting"
description: |-
  Configure IPS URL filter cache settings.
---

# fmgdevice_webfilter_ipsurlfiltercachesetting
Configure IPS URL filter cache settings.

## Example Usage

```hcl
resource "fmgdevice_webfilter_ipsurlfiltercachesetting" "trname" {
  dns_retry_interval = 10
  extended_ttl       = 10
  device_name        = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `dns_retry_interval` - Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
* `extended_ttl` - Extend time to live beyond reported by DNS. Use of 0 means use DNS server's TTL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterCacheSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_ipsurlfiltercachesetting.labelname WebfilterIpsUrlfilterCacheSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

